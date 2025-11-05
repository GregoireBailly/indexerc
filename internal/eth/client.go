package eth

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

type RPCAPI interface {
	BlockNumber(ctx context.Context) (uint64, error)
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

func (c *Client) LatestBlock(ctx context.Context) (uint64, error) {
	return c.rpc.BlockNumber(ctx)
}

func (c *Client) Close() {
	c.rpc.Close()
}
