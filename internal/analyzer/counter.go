package analyzer

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
)

type LogProvider interface {
	FetchLogs(ctx context.Context, lastNBlocks int, query ethereum.FilterQuery) ([]types.Log, error)
}

type Counter struct {
	provider LogProvider
}

func NewCounter(provider LogProvider) *Counter {
	return &Counter{provider: provider}
}

func (c *Counter) Count(ctx context.Context, lastNBlocks int, query ethereum.FilterQuery) (int, error) {
	logs, err := c.provider.FetchLogs(ctx, lastNBlocks, query)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch logs: %w", err)
	}
	return len(logs), nil
}
