// Convert between Geth and Index Supply types
package gethconv

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/indexsupply/x/abi"
)

func Log(l *types.Log) abi.Log {
	var topics [4][32]byte
	for i := range l.Topics {
		topics[i] = [32]byte(l.Topics[i].Bytes())
	}
	return abi.Log{
		Address: l.Address,
		Topics:  topics,
		Data:    l.Data,
	}
}
