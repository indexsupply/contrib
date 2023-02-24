// Parse event/log data based on the 'erc721' ABI
//
// Code generated by "genabi"; DO NOT EDIT.
package erc721

import (
	"github.com/indexsupply/x/abi"
	"github.com/indexsupply/x/abi/schema"
	"math/big"
)

type Approval struct {
	item *abi.Item

	// Indexed:
	Owner    [20]byte
	Approved [20]byte
	TokenId  *big.Int
}

func (x Approval) Done() {
	x.item.Done()
}

func decodeApproval(item *abi.Item) Approval {
	x := Approval{}
	return x
}

var (
	approvalSignature  = [32]byte{0x8c, 0x5b, 0xe1, 0xe5, 0xeb, 0xec, 0x7d, 0x5b, 0xd1, 0x4f, 0x71, 0x42, 0x7d, 0x1e, 0x84, 0xf3, 0xdd, 0x3, 0x14, 0xc0, 0xf7, 0xb2, 0x29, 0x1e, 0x5b, 0x20, 0xa, 0xc8, 0xc7, 0xc3, 0xb9, 0x25}
	approvalSchema     = schema.Parse("()")
	approvalNumIndexed = int(3)
)

// Event Signature:
//	Approval(address,address,uint256)
// Checks the first log topic against the signature hash:
//	8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925
//
// Copies indexed event inputs from the remaining topics
// into [Approval]
//
// Uses the the following abi schema to decode the un-indexed
// event inputs from the log's data field into [Approval]:
//	()
func MatchApproval(l abi.Log) (Approval, bool) {
	if len(l.Topics) > 0 && approvalSignature != l.Topics[0] {
		return Approval{}, false
	}
	if len(l.Topics[1:]) != approvalNumIndexed {
		return Approval{}, false
	}
	res := Approval{}
	res.Owner = abi.Bytes(l.Topics[1][:]).Address()
	res.Approved = abi.Bytes(l.Topics[2][:]).Address()
	res.TokenId = abi.Bytes(l.Topics[3][:]).BigInt()
	return res, true
}

type ApprovalForAll struct {
	item *abi.Item

	// Indexed:
	Owner    [20]byte
	Operator [20]byte
	// Un-indexed:
	Approved bool
}

func (x ApprovalForAll) Done() {
	x.item.Done()
}

func decodeApprovalForAll(item *abi.Item) ApprovalForAll {
	x := ApprovalForAll{}
	x.Approved = item.At(0).Bool()
	return x
}

var (
	approvalForAllSignature  = [32]byte{0x17, 0x30, 0x7e, 0xab, 0x39, 0xab, 0x61, 0x7, 0xe8, 0x89, 0x98, 0x45, 0xad, 0x3d, 0x59, 0xbd, 0x96, 0x53, 0xf2, 0x0, 0xf2, 0x20, 0x92, 0x4, 0x89, 0xca, 0x2b, 0x59, 0x37, 0x69, 0x6c, 0x31}
	approvalForAllSchema     = schema.Parse("(bool)")
	approvalForAllNumIndexed = int(2)
)

// Event Signature:
//	ApprovalForAll(address,address,bool)
// Checks the first log topic against the signature hash:
//	17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31
//
// Copies indexed event inputs from the remaining topics
// into [ApprovalForAll]
//
// Uses the the following abi schema to decode the un-indexed
// event inputs from the log's data field into [ApprovalForAll]:
//	(bool)
func MatchApprovalForAll(l abi.Log) (ApprovalForAll, bool) {
	if len(l.Topics) > 0 && approvalForAllSignature != l.Topics[0] {
		return ApprovalForAll{}, false
	}
	if len(l.Topics[1:]) != approvalForAllNumIndexed {
		return ApprovalForAll{}, false
	}
	_, item := abi.Decode(l.Data, approvalForAllSchema)
	res := decodeApprovalForAll(item)
	res.item = item
	res.Owner = abi.Bytes(l.Topics[1][:]).Address()
	res.Operator = abi.Bytes(l.Topics[2][:]).Address()
	return res, true
}

type Transfer struct {
	item *abi.Item

	// Indexed:
	From    [20]byte
	To      [20]byte
	TokenId *big.Int
}

func (x Transfer) Done() {
	x.item.Done()
}

func decodeTransfer(item *abi.Item) Transfer {
	x := Transfer{}
	return x
}

var (
	transferSignature  = [32]byte{0xdd, 0xf2, 0x52, 0xad, 0x1b, 0xe2, 0xc8, 0x9b, 0x69, 0xc2, 0xb0, 0x68, 0xfc, 0x37, 0x8d, 0xaa, 0x95, 0x2b, 0xa7, 0xf1, 0x63, 0xc4, 0xa1, 0x16, 0x28, 0xf5, 0x5a, 0x4d, 0xf5, 0x23, 0xb3, 0xef}
	transferSchema     = schema.Parse("()")
	transferNumIndexed = int(3)
)

// Event Signature:
//	Transfer(address,address,uint256)
// Checks the first log topic against the signature hash:
//	ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef
//
// Copies indexed event inputs from the remaining topics
// into [Transfer]
//
// Uses the the following abi schema to decode the un-indexed
// event inputs from the log's data field into [Transfer]:
//	()
func MatchTransfer(l abi.Log) (Transfer, bool) {
	if len(l.Topics) > 0 && transferSignature != l.Topics[0] {
		return Transfer{}, false
	}
	if len(l.Topics[1:]) != transferNumIndexed {
		return Transfer{}, false
	}
	res := Transfer{}
	res.From = abi.Bytes(l.Topics[1][:]).Address()
	res.To = abi.Bytes(l.Topics[2][:]).Address()
	res.TokenId = abi.Bytes(l.Topics[3][:]).BigInt()
	return res, true
}
