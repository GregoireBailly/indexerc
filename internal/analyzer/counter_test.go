package analyzer

import (
	"context"
	"errors"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProvider struct {
	mock.Mock
}

func (m *MockProvider) FetchLogs(ctx context.Context, lastNBlocks int, q ethereum.FilterQuery) ([]types.Log, error) {
	args := m.Called(ctx, lastNBlocks, q)
	return args.Get(0).([]types.Log), args.Error(1)
}

func TestCounter_Count_Success(t *testing.T) {
	mockProvider := new(MockProvider)
	counter := NewCounter(mockProvider)

	query := ethereum.FilterQuery{}
	expectedLogs := []types.Log{{}, {}, {}}

	mockProvider.On("FetchLogs", mock.Anything, 10, query).Return(expectedLogs, nil)

	count, err := counter.Count(context.Background(), 10, query)

	assert.NoError(t, err)
	assert.Equal(t, 3, count)
	mockProvider.AssertExpectations(t)
}

func TestCounter_Count_Error(t *testing.T) {
	mockProvider := new(MockProvider)
	counter := NewCounter(mockProvider)

	query := ethereum.FilterQuery{}
	mockProvider.On("FetchLogs", mock.Anything, 5, query).Return([]types.Log{}, errors.New("unavailable"))

	count, err := counter.Count(context.Background(), 5, query)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to fetch logs")
	assert.Equal(t, 0, count)
	mockProvider.AssertExpectations(t)
}
