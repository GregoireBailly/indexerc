package eth

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

func TestIsERC20Transfer(t *testing.T) {
	tests := []struct {
		name     string
		log      types.Log
		expected bool
	}{
		{
			name: "valid ERC20 Transfer event",
			log: types.Log{
				Address: common.HexToAddress("0x1234"),
				Topics:  []common.Hash{ERC20TransferTopic},
				Data:    big.NewInt(100).Bytes(),
			},
			expected: true,
		},
		{
			name:     "empty topics",
			log:      types.Log{},
			expected: false,
		},
		{
			name: "non-transfer topic",
			log: types.Log{
				Topics: []common.Hash{common.HexToHash("0xdeadbeef")},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, IsERC20Transfer(tt.log))
		})
	}
}
