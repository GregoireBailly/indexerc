package eth

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// RPCAPI defines the subset of Ethereum client methods used by our app.
type RPCAPI interface {
	BlockNumber(ctx context.Context) (uint64, error)
	FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error)
	Close()
}

type Client struct {
	rpc RPCAPI
}

func New(ctx context.Context, rpcURL string) (*Client, error) {
	r, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum node: %w", err)
	}
	return &Client{rpc: r}, nil
}

func (c *Client) FetchLogs(ctx context.Context, lastNBlocks int, q ethereum.FilterQuery) ([]types.Log, error) {
	if lastNBlocks <= 0 {
		return nil, fmt.Errorf("invalid block range: %d", lastNBlocks)
	}
	
	latest, err := c.latestBlock(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest block: %w", err)
	}

	from := big.NewInt(0)
	if int64(latest) > int64(lastNBlocks) {
		from = big.NewInt(int64(latest) - int64(lastNBlocks))
	}

	q.FromBlock = from
	q.ToBlock = big.NewInt(int64(latest))

	return c.logs(ctx, q)
}

func (c *Client) Close() {
	c.rpc.Close()
}

// --- private helpers ---

func (c *Client) latestBlock(ctx context.Context) (uint64, error) {
	return c.rpc.BlockNumber(ctx)
}

func (c *Client) logs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	logs, err := c.rpc.FilterLogs(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch logs: %w", err)
	}
	return logs, nil
}
