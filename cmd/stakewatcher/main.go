package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/lib/pq"
	"github.com/markbates/pkger"
	"github.com/public-awesome/stakebird/app"
	"github.com/public-awesome/stakewatcher/client"
	"github.com/public-awesome/stakewatcher/models"
	"github.com/public-awesome/stakewatcher/workqueue"
	"github.com/rs/zerolog/log"
	migrate "github.com/rubenv/sql-migrate"
	tmtypes "github.com/tendermint/tendermint/types"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

const (
	envPrefix = "STAKEWATCHER"
)

func runMigrations(db *sql.DB) {
	migrate.SetTable("migrations")
	migrationSource := &migrate.HttpFileSystemMigrationSource{
		FileSystem: pkger.Dir("/db/migrations/"),
	}
	n, err := migrate.ExecMax(db, "postgres", migrationSource, migrate.Up, 0)

	if err != nil {
		panic(err)
	}

	log.Info().Msgf("migrations run %d", n)

}

func getEnv(key, defaultValue string) string {
	v := os.Getenv(fmt.Sprintf("%s_%s", envPrefix, key))
	if v == strings.TrimSpace("") {
		return defaultValue
	}
	return v
}

func main() {
	var (
		rpcEndpoint        string
		restServerEndpoint string
		autoMigrate        bool
	)
	fs := flag.NewFlagSet("stakewatcher", flag.ExitOnError)
	fs.StringVar(&rpcEndpoint, "rpc-endpoint", getEnv("RPC_ENDPOINT", "http://localhost:26657"), "--rpc-endpoint specify the rpc endpoint")
	fs.StringVar(&restServerEndpoint, "rest-server", getEnv("REST_SERVER", "http://localhost:1317"), "--rest-server specify the rest-server endpoint")
	fs.BoolVar(&autoMigrate, "auto-migrate", false, "--auto-migrate specificy if should perform database migration on start")
	err := fs.Parse(os.Args[1:])
	if err != nil {
		log.Fatal().Err(err).Msg("error parsing arguments")
	}

	dbSourceName := os.Getenv("DB_SOURCE")
	if dbSourceName == "" {
		log.Info().Msg("using default db settings")
		dbSourceName = "dbname=stakewatcher user=postgres password=postgres sslmode=disable"
	}
	// Open handle to database like normal
	db, err := sql.Open("postgres", dbSourceName)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to conect to db")
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal().Err(err).Msg("unable to conect to db")
	}

	if autoMigrate {
		runMigrations(db)
	}

	appCodec, cdc := app.MakeCodecs()

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(app.Bech32PrefixAccAddr, app.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(app.Bech32PrefixValAddr, app.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(app.Bech32PrefixConsAddr, app.Bech32PrefixConsPub)
	config.Seal()

	cp, err := client.NewProxy(rpcEndpoint, restServerEndpoint, cdc, appCodec)
	if err != nil {
		log.Fatal().Err(err).Msg("error initializing client")
	}
	log.Info().Str("rpc-endpoint", rpcEndpoint).Str("rest-server", restServerEndpoint).Msg("client settings")

	// context cancelation setup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()

	go func() {
		defer wg.Done()

		// Wait for signals (configurable per-platform) and then cancel the
		// context to indicate that the process should shut down.
		sigC := make(chan os.Signal, 1)
		signal.Notify(sigC, syscall.SIGINT, syscall.SIGTERM)

		s := <-sigC
		log.Info().Msgf("received %s, shutting down", s)
		cancel()

		// Stop handling signals at this point to allow the user to forcefully
		// terminate the binary.
		signal.Stop(sigC)
	}()

	exportQueue := make(chan int64, 100)
	go enqueueMissingBlocks(ctx, cp, db, exportQueue)
	wk := workqueue.NewWorker(cdc, appCodec, exportQueue, db, cp)
	go wk.Start(ctx)
	startNewBlockListener(ctx, cp, exportQueue, db)

}

// enqueueMissingBlocks enqueues jobs (block heights) for missed blocks starting
// at the startHeight up until the latest known height.
func enqueueMissingBlocks(ctx context.Context, cp *client.Proxy, db *sql.DB, exportQueue chan<- int64) {
	latestBlockHeight, err := cp.LatestHeight()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to fetch latest block from RPC endpoint")
	}
	var startHeight int64
	sl, err := models.SyncLogs(qm.OrderBy(fmt.Sprintf("%s DESC", models.SyncLogColumns.BlockHeight))).One(ctx, db)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	if err == sql.ErrNoRows {
		startHeight = 1
	}

	if err == nil {
		startHeight = sl.BlockHeight + 1
	}

	log.Info().Msg("syncing missing blocks")
	for i := startHeight; i <= latestBlockHeight; i++ {
		// TODO : check context cancelation status
		log.Info().Int64("height", i).Msg("enqueueing missing block")
		l := &models.SyncLog{
			BlockHeight: i,
		}
		err := l.Insert(ctx, db, boil.Infer())
		if err != nil {
			panic(err)
		}
		exportQueue <- i
	}
}

// startNewBlockListener subscribes to new block events via the Tendermint RPC
// and enqueues each new block height onto the provided queue. It blocks as new
// blocks are incoming.
func startNewBlockListener(ctx context.Context, cp *client.Proxy, exportQueue chan<- int64, db *sql.DB) {
	eventCh, cancel, err := cp.SubscribeNewBlocks("stakewatcher-client")
	defer cancel()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to subscribe to new blocks")
	}
	log.Info().Msg("listening for new block events")
	t := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-ctx.Done():
			log.Info().Msg("closing block listener")
			return
		case e := <-eventCh:
			newBlock := e.Data.(tmtypes.EventDataNewBlock).Block
			height := newBlock.Header.Height
			log.Info().Int64("height", height).Msg("enqueueing missing block")
			l := &models.SyncLog{
				BlockHeight: height,
			}
			err := l.Insert(ctx, db, boil.Infer())
			if err != nil {
				panic(err)
			}
			exportQueue <- height
		case <-t.C:
			log.Info().Msg("retry pending blocks")
			go retryBlocks(ctx, exportQueue, db)
		}
	}
}

func retryBlocks(ctx context.Context, exportQueue chan<- int64, db *sql.DB) {
	q := fmt.Sprintf("%s is null and %s < ?", models.SyncLogColumns.SyncedAt, models.SyncLogColumns.CreatedAt)
	blocks, err := models.SyncLogs(qm.Where(q, time.Now().Add(time.Second*-20))).All(ctx, db)
	if err != nil {
		log.Error().Err(err).Msg("error getting pending blocks")
	}
	for _, b := range blocks {
		exportQueue <- b.BlockHeight
	}
}
