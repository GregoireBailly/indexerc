package eth

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	// ERC20TransferTopic is the keccak256 hash of the ERC-20 Transfer event signature.
	ERC20TransferTopic = crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
)

func IsERC20Transfer(log types.Log) bool {
	return len(log.Topics) > 0 && log.Topics[0] == ERC20TransferTopic
}
