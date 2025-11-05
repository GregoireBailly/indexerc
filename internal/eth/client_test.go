package eth

import (
	"context"
	"errors"
	"testing"

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

func (m *MockRPC) Close() {}

// --- TESTS ---

func TestLatestBlock_Success(t *testing.T) {
	mockRPC := new(MockRPC)
	client := Client{rpc: mockRPC}

	mockRPC.On("BlockNumber", mock.Anything).Return(uint64(123456), nil)

	blockNum, err := client.LatestBlock(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, uint64(123456), blockNum)

	mockRPC.AssertExpectations(t)
}

func TestLatestBlock_Error(t *testing.T) {
	mockRPC := new(MockRPC)
	client := Client{rpc: mockRPC}

	mockErr := errors.New("Ethereum node unavailable")
	mockRPC.On("BlockNumber", mock.Anything).Return(uint64(0), mockErr)

	blockNum, err := client.LatestBlock(context.Background())
	assert.Error(t, err)
	assert.EqualError(t, err, "Ethereum node unavailable")
	assert.Zero(t, blockNum)

	mockRPC.AssertExpectations(t)
}
