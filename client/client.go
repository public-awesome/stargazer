package client

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/types/query"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	stargazeparams "github.com/public-awesome/stargaze/app/params"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
	libclient "github.com/tendermint/tendermint/rpc/jsonrpc/client"
)

var clientTimeout = 5 * time.Second

// Proxy implements a wrapper around both a Tendermint RPC client and a
// cosmos SDK REST client that allows for essential data queries.
type Proxy struct {
	rpcClient      rpcclient.Client // Tendermint (RPC client) node
	httpClient     *http.Client
	rpcNode        string
	encodingConfig stargazeparams.EncodingConfig
}

func newRPCClient(remote string) (*rpchttp.HTTP, error) {
	httpClient, err := libclient.DefaultHTTPClient(remote)
	if err != nil {
		return nil, err
	}
	httpClient.Timeout = clientTimeout
	rpcClient, err := rpchttp.NewWithClient(remote, "/websocket", httpClient)

	if err != nil {
		return nil, err
	}

	if err := rpcClient.Start(); err != nil {
		return nil, err
	}
	return rpcClient, nil
}

// NewProxy returns a new Proxy instance
func NewProxy(rpcNode string, encodingConfig stargazeparams.EncodingConfig) (*Proxy, error) {
	rpcClient, err := newRPCClient(rpcNode)
	if err != nil {
		return nil, err
	}
	p := &Proxy{
		rpcClient: rpcClient,
		httpClient: &http.Client{
			Timeout: clientTimeout,
		},
		rpcNode:        rpcNode,
		encodingConfig: encodingConfig,
	}
	return p, nil
}

// LatestHeight returns the latest block height on the active chain. An error
// is returned if the query fails.
func (p *Proxy) LatestHeight(ctx context.Context) (int64, error) {
	status, err := p.rpcClient.Status(ctx)
	if err != nil {
		return -1, err
	}

	height := status.SyncInfo.LatestBlockHeight
	return height, nil
}

// Block queries for a block by height. An error is returned if the query fails.
func (p *Proxy) Block(ctx context.Context, height int64) (*tmctypes.ResultBlock, error) {
	return p.rpcClient.Block(ctx, &height)
}

// BlockResults queries for block results by height. An error is returned if the query fails.
func (p *Proxy) BlockResults(ctx context.Context, height int64) (*tmctypes.ResultBlockResults, error) {
	return p.rpcClient.BlockResults(ctx, &height)
}

// Validators returns all the known Tendermint validators for a given block
// height. An error is returned if the query fails.
func (p *Proxy) Validators(ctx context.Context, height int64) (*tmctypes.ResultValidators, error) {
	page := 1
	perPage := 1000
	return p.rpcClient.Validators(ctx, &height, &page, &perPage)
}

// AppValidators returns the list of validators from the state machine
func (p *Proxy) AppValidators(ctx context.Context, height int64) ([]stakingtypes.Validator, error) {
	initClientCtx := client.Context{}.
		WithJSONMarshaler(p.encodingConfig.Marshaler).
		WithInterfaceRegistry(p.encodingConfig.InterfaceRegistry).
		WithTxConfig(p.encodingConfig.TxConfig).
		WithLegacyAmino(p.encodingConfig.Amino).
		WithInput(os.Stdin).
		WithAccountRetriever(types.AccountRetriever{}).
		WithBroadcastMode(flags.BroadcastBlock).
		WithNodeURI(p.rpcNode).
		WithClient(p.rpcClient).WithHeight(height)

	pageReq := &query.PageRequest{
		Limit: 200,
	}

	result, err := stakingtypes.NewQueryClient(initClientCtx).Validators(ctx, &stakingtypes.QueryValidatorsRequest{
		// Leaving status empty on purpose to query all validators.
		Pagination: pageReq,
	})
	if err != nil {
		return nil, err
	}

	return result.Validators, nil
}

// Genesis retrieves the genesis from tendermint
func (p *Proxy) Genesis(ctx context.Context) (*tmctypes.ResultGenesis, error) {
	return p.rpcClient.Genesis(ctx)
}

// Stop defers the node stop execution to the RPC client.
func (p *Proxy) Stop() error {
	return p.rpcClient.Stop()
}

// SubscribeNewBlocks subscribes to the new block event handler through the RPC
// client with the given subscriber name. An receiving only channel, context
// cancel function and an error is returned. It is up to the caller to cancel
// the context and handle any errors appropriately.
func (p *Proxy) SubscribeNewBlocks(subscriber string) (<-chan tmctypes.ResultEvent, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	eventCh, err := p.rpcClient.Subscribe(ctx, subscriber, "tm.event = 'NewBlock'")
	return eventCh, cancel, err
}

// Tx queries for a transaction from the REST client and decodes it into a sdk.Tx
// if the transaction exists. An error is returned if the tx doesn't exist or
// decoding fails.
func (p *Proxy) Tx(hash string) (*sdk.TxResponse, error) {
	initClientCtx := client.Context{}.
		WithJSONMarshaler(p.encodingConfig.Marshaler).
		WithInterfaceRegistry(p.encodingConfig.InterfaceRegistry).
		WithTxConfig(p.encodingConfig.TxConfig).
		WithLegacyAmino(p.encodingConfig.Amino).
		WithInput(os.Stdin).
		WithAccountRetriever(types.AccountRetriever{}).
		WithBroadcastMode(flags.BroadcastBlock).
		WithNodeURI(p.rpcNode).
		WithClient(p.rpcClient)
	return authclient.QueryTx(initClientCtx, hash)
}

// Txs queries for all the transactions in a block. Transactions are returned
// in the sdk.TxResponse format which internally contains an sdk.Tx. An error is
// returned if any query fails.
func (p *Proxy) Txs(block *tmctypes.ResultBlock) ([]*sdk.TxResponse, error) {
	txResponses := make([]*sdk.TxResponse, len(block.Block.Txs))
	for i, tmTx := range block.Block.Txs {
		txResponse, err := p.Tx(fmt.Sprintf("%X", tmTx.Hash()))
		if err != nil {
			return nil, err
		}
		txResponses[i] = txResponse
	}

	return txResponses, nil
}

// Status returns status of the rpc client
func (p *Proxy) Status(ctx context.Context) (*tmctypes.ResultStatus, error) {
	return p.rpcClient.Status(ctx)
}
