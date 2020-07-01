package client

import (
	"context"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"

	rpcclient "github.com/tendermint/tendermint/rpc/client"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
	libclient "github.com/tendermint/tendermint/rpc/lib/client"
)

var clientTimeout = 5 * time.Second

// Proxy implements a wrapper around both a Tendermint RPC client and a
// cosmos SDK REST client that allows for essential data queries.
type Proxy struct {
	rpcClient  rpcclient.Client // Tendermint (RPC client) node
	httpClient *http.Client
	restNode   string // Full (REST client) node
	cdc        *codec.Codec
	appCodec   *std.Codec
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
func NewProxy(rpcNode, restNode string, cdc *codec.Codec, appCodec *std.Codec) (*Proxy, error) {
	rpcClient, err := newRPCClient(rpcNode)
	if err != nil {
		return nil, err
	}
	p := &Proxy{
		rpcClient: rpcClient,
		httpClient: &http.Client{
			Timeout: clientTimeout,
		},
		restNode: restNode,
		appCodec: appCodec,
		cdc:      cdc,
	}
	return p, nil
}

// LatestHeight returns the latest block height on the active chain. An error
// is returned if the query fails.
func (p *Proxy) LatestHeight() (int64, error) {
	status, err := p.rpcClient.Status()
	if err != nil {
		return -1, err
	}

	height := status.SyncInfo.LatestBlockHeight
	return height, nil
}

// Block queries for a block by height. An error is returned if the query fails.
func (p *Proxy) Block(height int64) (*tmctypes.ResultBlock, error) {
	return p.rpcClient.Block(&height)
}

// TendermintTx queries for a transaction by hash. An error is returned if the
// query fails.
func (p *Proxy) TendermintTx(hash string) (*tmctypes.ResultTx, error) {
	hashRaw, err := hex.DecodeString(hash)
	if err != nil {
		return nil, err
	}

	return p.rpcClient.Tx(hashRaw, false)
}

// Validators returns all the known Tendermint validators for a given block
// height. An error is returned if the query fails.
func (p *Proxy) Validators(height int64) (*tmctypes.ResultValidators, error) {
	return p.rpcClient.Validators(&height, 1, 10000)
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
func (p *Proxy) Tx(hash string) (sdk.TxResponse, error) {

	resp, err := p.httpClient.Get(fmt.Sprintf("%s/txs/%s", p.restNode, hash))
	if err != nil {
		return sdk.TxResponse{}, err
	}

	defer resp.Body.Close()

	bz, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return sdk.TxResponse{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return sdk.TxResponse{}, fmt.Errorf("response error fetching transaction with status [%d] -  %s ", resp.StatusCode, resp.Status)
	}
	var tx sdk.TxResponse
	if err := p.cdc.UnmarshalJSON(bz, &tx); err != nil {
		return sdk.TxResponse{}, err
	}
	return tx, nil
}

// Txs queries for all the transactions in a block. Transactions are returned
// in the sdk.TxResponse format which internally contains an sdk.Tx. An error is
// returned if any query fails.
func (p *Proxy) Txs(block *tmctypes.ResultBlock) ([]sdk.TxResponse, error) {
	txResponses := make([]sdk.TxResponse, len(block.Block.Txs), len(block.Block.Txs))

	for i, tmTx := range block.Block.Txs {
		txResponse, err := p.Tx(fmt.Sprintf("%X", tmTx.Hash()))
		if err != nil {
			return nil, err
		}
		txResponses[i] = txResponse
	}

	return txResponses, nil
}
