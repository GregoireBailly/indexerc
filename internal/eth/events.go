package eth

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	// ERC20TransferTopic is the keccak256 hash of the ERC-20 Transfer event signature.
	ERC20TransferTopic = crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
)

func ERC20TransferQuery() ethereum.FilterQuery {
	return ethereum.FilterQuery{
		Topics: [][]common.Hash{{ERC20TransferTopic}},
	}
}
