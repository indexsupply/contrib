package bench

import (
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/indexsupply/contrib/abis/erc721"
	"github.com/indexsupply/x/abi"
)

//go:generate abigen -abi ../abis/erc721/erc721.json --type erc721 --pkg bench --out erc721.go

func h2b(s string) []byte {
	b, _ := hex.DecodeString(s)
	return b
}

var log = types.Log{
	Address: [20]byte(h2b("0000000000a39bb272e79075ade125fd351887ac")),
	Topics: []common.Hash{
		[32]byte(h2b("0000000000000000000000002e0406b4fd09fa101800f43833039e50d5a0ad1e")),
		[32]byte(h2b("0000000000000000000000007b873602d8770303be7273d0023d89444d32f9ca")),
	},
	Data: h2b("00000000000000000000000000000000000000000000000000afdbfcdc61c000"),
}

var isnLog = abi.Log{
	Address: [20]byte(h2b("0000000000a39bb272e79075ade125fd351887ac")),
	Topics: [4][32]byte{
		[32]byte(h2b("0000000000000000000000002e0406b4fd09fa101800f43833039e50d5a0ad1e")),
		[32]byte(h2b("0000000000000000000000007b873602d8770303be7273d0023d89444d32f9ca")),
	},
	Data: h2b("00000000000000000000000000000000000000000000000000afdbfcdc61c000"),
}

func BenchmarkERC721(b *testing.B) {
	g721, err := NewErc721([20]byte{}, &ethclient.Client{})
	if err != nil {
		b.Fatal(err)
	}
	b.Run("geth", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			g721.ParseTransfer(log)
		}
	})
	b.Run("isn1", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			erc721.MatchTransfer(isnLog)
		}
	})
}
