package workqueue

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/public-awesome/stakewatcher/client"
	"github.com/public-awesome/stakewatcher/models"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/sqlboiler/v4/boil"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

type Worker struct {
	q        <-chan int64
	cdc      *codec.Codec
	appCodec *std.Codec
	db       *sql.DB
	cp       *client.Proxy
}

// NewWorker returns an intialized worker
func NewWorker(cdc *codec.Codec, appCodec *std.Codec, queue <-chan int64, db *sql.DB, cp *client.Proxy) *Worker {
	return &Worker{
		q:        queue,
		cdc:      cdc,
		appCodec: appCodec,
		db:       db,
		cp:       cp,
	}
}

func (w *Worker) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Info().Msg("closing worker listener")
			return
		case blockHeight := <-w.q:
			w.process(ctx, blockHeight)
		}
	}
}

func (w *Worker) process(ctx context.Context, height int64) error {
	// skip block 1
	if height == 1 {
		return nil
	}

	exists, err := models.BlockExists(ctx, w.db, height)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	block, err := w.cp.Block(height)
	if err != nil {
		log.Info().Err(err).Int64("height", height).Msg("failed to get block")
		return err
	}

	txs, err := w.cp.Txs(block)
	if err != nil {
		log.Info().Err(err).Int64("height", height).Msg("failed to get transactions for block")
		return err
	}

	vals, err := w.cp.Validators(block.Block.LastCommit.GetHeight())
	if err != nil {
		log.Info().Err(err).Int64("height", height).Msg("failed to get validators for block")
		return err
	}

	err = ExportBlockSignatures(ctx, block.Block.LastCommit, vals, w.db)
	if err != nil {
		log.Info().Err(err).Int64("height", height).Msg("failed to export precommits")
		return err
	}
	// if err := w.db.ExportBlockSignatures(block.Block.LastCommit, vals); err != nil {
	// 	return err
	// }
	err = ExportBlock(ctx, block, txs, vals, w.db)
	if err != nil {
		return err
	}
	// return w.db.ExportBlock(block, txs, vals)
	return nil
}

func ExportBlock(ctx context.Context, b *tmctypes.ResultBlock, txs []sdk.TxResponse, validators *tmctypes.ResultValidators, db *sql.DB) error {
	totalGas := sumGasTxs(txs)
	signatures := len(b.Block.LastCommit.Signatures)

	proposerAddress := sdk.ConsAddress(b.Block.ProposerAddress).String()
	val := findValidatorByAddr(proposerAddress, validators)
	if val == nil {
		err := fmt.Errorf("failed to find validator by address %s for block %d", proposerAddress, b.Block.Height)
		return err
	}

	block := &models.Block{
		Height:          b.Block.Height,
		NumTXS:          len(txs),
		TotalGas:        int64(totalGas),
		Signatures:      signatures,
		ProposerAddress: proposerAddress,
		Hash:            b.BlockID.Hash.String(),
		BlockTimestamp:  b.Block.Time,
	}
	err := block.Insert(ctx, db, boil.Infer())
	if err != nil {
		return err
	}

	for _, t := range txs {
		if err != nil {
			fmt.Println("error parsing time", err)
			continue
		}
		tx := &models.Transaction{
			Hash:      t.TxHash,
			GasWanted: int(t.GasWanted),
			GasUsed:   int(t.GasUsed),
		}
		err = tx.Insert(ctx, db, boil.Infer())
		if err != nil {
			fmt.Println("error inserting tx", err)
			continue
		}
	}
	return nil
}

// sumGasTxs returns the total gas consumed by a set of transactions.
func sumGasTxs(txs []sdk.TxResponse) uint64 {
	var totalGas uint64

	for _, tx := range txs {
		totalGas += uint64(tx.GasUsed)
	}

	return totalGas
}

// ExportBlockSignatures ...
func ExportBlockSignatures(ctx context.Context, commit *tmtypes.Commit, validators *tmctypes.ResultValidators, db *sql.DB) error {
	for _, sig := range commit.Signatures {
		if sig.Signature == nil {
			continue
		}
		addr := sdk.ConsAddress(sig.ValidatorAddress).String()
		val := findValidatorByAddr(addr, validators)
		if val == nil {
			err := fmt.Errorf("failed to find validator by address %s for block %d", addr, commit.GetHeight())
			return err
		}
		err := ExportValidator(ctx, val, db)
		if err != nil {
			return err
		}

		err = SetBlockSignature(ctx, commit, sig, val.VotingPower, val.ProposerPriority, db)
		if err != nil {
			return err
		}

		log.Info().Int64("voting-power", val.VotingPower).Msg("found validator")
	}
	return nil
}

// SetBlockSignature stores a block prevote
func SetBlockSignature(ctx context.Context, commit *tmtypes.Commit, sig tmtypes.CommitSig, vp, pp int64, db *sql.DB) error {
	s := &models.BlockSignature{
		Height:           commit.GetHeight(),
		Round:            commit.GetRound(),
		ValidatorAddress: sdk.ConsAddress(sig.ValidatorAddress).String(),
		VotingPower:      int(vp),
		ProposerPriority: int(pp),
		Flag:             int(sig.BlockIDFlag),
		Timestamp:        sig.Timestamp,
		Hash:             sig.BlockID(commit.BlockID).Hash.String(),
	}
	return s.Insert(ctx, db, boil.Infer())
}

// ExportValidator exports validator
func ExportValidator(ctx context.Context, val *tmtypes.Validator, db *sql.DB) error {
	address := sdk.ConsAddress(val.Address).String()
	consPubKey, err := sdk.Bech32ifyPubKey(sdk.Bech32PubKeyTypeConsPub, val.PubKey)

	if err != nil {
		log.Error().Err(err).Str("validator", address).Msg("failed to convert validator public key")
		return err
	}
	validator := &models.Validator{
		Address: address,
		PubKey:  consPubKey,
	}
	return validator.Upsert(ctx, db, false, []string{}, boil.Whitelist(), boil.Infer())
}

// findValidatorByAddr finds a validator by address given a set of
// Tendermint validators for a particular block. If no validator is found, nil
// is returned.
func findValidatorByAddr(address string, vals *tmctypes.ResultValidators) *tmtypes.Validator {
	for _, val := range vals.Validators {
		if strings.EqualFold(address, sdk.ConsAddress(val.Address).String()) {
			return val
		}
	}
	return nil
}
