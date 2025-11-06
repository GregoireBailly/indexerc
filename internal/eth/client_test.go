package eth

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRPC struct {
	mock.Mock
}

func (m *MockRPC) BlockNumber(ctx context.Context) (uint64, error) {
	args := m.Called(ctx)
	return args.Get(0).(uint64), args.Error(1)
}

func (m *MockRPC) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	args := m.Called(ctx, q)
	return args.Get(0).([]types.Log), args.Error(1)
}

func (m *MockRPC) Close() {}

// --- TESTS ---

func TestFetchLogs_Success(t *testing.T) {
	mockRPC := new(MockRPC)
	client := Client{rpc: mockRPC}

	latest := uint64(200)
	logs := []types.Log{{Index: 1}, {Index: 2}}

	mockRPC.On("BlockNumber", mock.Anything).Return(latest, nil)
	mockRPC.On("FilterLogs", mock.Anything, mock.MatchedBy(func(q ethereum.FilterQuery) bool {
		return q.FromBlock.Uint64() == 190 && q.ToBlock.Uint64() == 200
	})).Return(logs, nil)

	got, err := client.FetchLogs(context.Background(), 10, ethereum.FilterQuery{})
	assert.NoError(t, err)
	assert.Len(t, got, 2)
	mockRPC.AssertExpectations(t)
}

func TestFetchLogs_ClampFromBlock(t *testing.T) {
	mockRPC := new(MockRPC)
	client := Client{rpc: mockRPC}

	latest := uint64(5)
	mockRPC.On("BlockNumber", mock.Anything).Return(latest, nil)
	mockRPC.On("FilterLogs", mock.Anything, mock.MatchedBy(func(q ethereum.FilterQuery) bool {
		return q.FromBlock.Cmp(big.NewInt(0)) == 0 && q.ToBlock.Uint64() == 5
	})).Return([]types.Log{}, nil)

	_, err := client.FetchLogs(context.Background(), 10, ethereum.FilterQuery{})
	assert.NoError(t, err)
	mockRPC.AssertExpectations(t)
}

func TestFetchLogs_BlockNumberError(t *testing.T) {
	mockRPC := new(MockRPC)
	client := Client{rpc: mockRPC}

	mockRPC.On("BlockNumber", mock.Anything).Return(uint64(0), errors.New("unavailable"))

	_, err := client.FetchLogs(context.Background(), 10, ethereum.FilterQuery{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get latest block")
	mockRPC.AssertExpectations(t)
}

func TestFetchLogs_FilterLogsError(t *testing.T) {
	mockRPC := new(MockRPC)
	client := Client{rpc: mockRPC}

	mockRPC.On("BlockNumber", mock.Anything).Return(uint64(100), nil)
	mockRPC.On("FilterLogs", mock.Anything, mock.Anything).Return([]types.Log{}, errors.New("rpc timeout"))

	_, err := client.FetchLogs(context.Background(), 10, ethereum.FilterQuery{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to fetch logs")
	mockRPC.AssertExpectations(t)
}

func TestFetchLogs_ZeroRange(t *testing.T) {
	mockRPC := new(MockRPC)
    client := Client{rpc: mockRPC}

    _, err := client.FetchLogs(context.Background(), 0, ethereum.FilterQuery{})
    assert.Error(t, err)
    assert.Contains(t, err.Error(), "invalid block range")
    mockRPC.AssertExpectations(t)

}
