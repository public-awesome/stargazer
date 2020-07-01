package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	sdk "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/lib/pq"
	"github.com/public-awesome/stakebird/app"
	"github.com/public-awesome/stakewatcher/client"
	"github.com/public-awesome/stakewatcher/workqueue"
	"github.com/rs/zerolog/log"
	tmtypes "github.com/tendermint/tendermint/types"
)

func main() {
	flag.NewFlagSet("stakewatcher", flag.ExitOnError)
	flag.Parse()
	appCodec, cdc := app.MakeCodecs()

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(app.Bech32PrefixAccAddr, app.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(app.Bech32PrefixValAddr, app.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(app.Bech32PrefixConsAddr, app.Bech32PrefixConsPub)
	config.Seal()

	cp, err := client.NewProxy("http://localhost:26657", "http://localhost:1317", cdc, appCodec)

	if err != nil {
		log.Fatal().Err(fmt.Errorf("error init client %w", err))
	}

	dbSourceName := os.Getenv("DB_SOURCE")
	if dbSourceName == "" {
		fmt.Println("using default db")
		dbSourceName = "dbname=stakewatcher user=postgres password=postgres sslmode=disable"
	}

	// Open handle to database like normal
	db, err := sql.Open("postgres", dbSourceName)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
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
		log.Printf("received %s, shutting down", s)
		cancel()

		// Stop handling signals at this point to allow the user to forcefully
		// terminate the binary.
		signal.Stop(sigC)
	}()

	exportQueue := make(chan int64, 100)
	go enqueueMissingBlocks(ctx, cp, 1, exportQueue)
	wk := workqueue.NewWorker(cdc, appCodec, exportQueue, db, cp)
	go wk.Start(ctx)
	startNewBlockListener(ctx, cp, exportQueue)

}

// enqueueMissingBlocks enqueues jobs (block heights) for missed blocks starting
// at the startHeight up until the latest known height.
func enqueueMissingBlocks(ctx context.Context, cp *client.Proxy, startHeight int64, exportQueue chan<- int64) {
	latestBlockHeight, err := cp.LatestHeight()
	if err != nil {
		log.Fatal().Err(fmt.Errorf("failed to get lastest block from RPC client %w", err))
	}

	log.Info().Msg("syncing missing blocks")

	for i := startHeight; i <= latestBlockHeight; i++ {
		ctx.Err()
		log.Info().Int64("height", i).Msg("enqueueing missing block")
		exportQueue <- i
	}
}

// startNewBlockListener subscribes to new block events via the Tendermint RPC
// and enqueues each new block height onto the provided queue. It blocks as new
// blocks are incoming.
func startNewBlockListener(ctx context.Context, cp *client.Proxy, exportQueue chan<- int64) {
	eventCh, cancel, err := cp.SubscribeNewBlocks("juno-client")
	defer cancel()

	if err != nil {
		log.Fatal().Err(fmt.Errorf("failed to subscribe to new blocks %w", err))
	}

	log.Printf("listening for new block events...")

	for {
		select {
		case <-ctx.Done():
			log.Info().Msg("closing block listener")
			return
		case e := <-eventCh:
			newBlock := e.Data.(tmtypes.EventDataNewBlock).Block
			height := newBlock.Header.Height
			log.Info().Int64("height", height).Msg("enqueueing new block")
			exportQueue <- height
		}
	}
}
