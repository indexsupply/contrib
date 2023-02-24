// Parse event/log data based on the 'erc1155' ABI
//
// Code generated by "genabi"; DO NOT EDIT.
package erc1155

import (
	"github.com/indexsupply/x/abi"
	"github.com/indexsupply/x/abi/schema"
	"math/big"
)

type ApprovalForAll struct {
	item *abi.Item

	// Indexed:
	Account  [20]byte
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
func MatchApprovalForAll(l abi.Log) (ApprovalForAll, error) {
	if len(l.Topics) > 0 && approvalForAllSignature != l.Topics[0] {
		return ApprovalForAll{}, abi.SigMismatch
	}
	if len(l.Topics[1:]) != approvalForAllNumIndexed {
		return ApprovalForAll{}, abi.IndexMismatch
	}
	item, _, err := abi.Decode(l.Data, approvalForAllSchema)
	if err != nil {
		return ApprovalForAll{}, err
	}
	res := decodeApprovalForAll(item)
	res.item = item
	res.Account = abi.Bytes(l.Topics[1][:]).Address()
	res.Operator = abi.Bytes(l.Topics[2][:]).Address()
	return res, nil
}

type TransferBatch struct {
	item *abi.Item

	// Indexed:
	Operator [20]byte
	From     [20]byte
	To       [20]byte
	// Un-indexed:
	Ids    []*big.Int
	Values []*big.Int
}

func (x TransferBatch) Done() {
	x.item.Done()
}

func decodeTransferBatch(item *abi.Item) TransferBatch {
	x := TransferBatch{}
	idsItem0 := item.At(0)
	ids0 := make([]*big.Int, idsItem0.Len())
	for i0 := 0; i0 < idsItem0.Len(); i0++ {
		ids0[i0] = idsItem0.BigInt()
	}
	x.Ids = ids0
	valuesItem0 := item.At(1)
	values0 := make([]*big.Int, valuesItem0.Len())
	for i0 := 0; i0 < valuesItem0.Len(); i0++ {
		values0[i0] = valuesItem0.BigInt()
	}
	x.Values = values0
	return x
}

var (
	transferBatchSignature  = [32]byte{0x4a, 0x39, 0xdc, 0x6, 0xd4, 0xc0, 0xdb, 0xc6, 0x4b, 0x70, 0xaf, 0x90, 0xfd, 0x69, 0x8a, 0x23, 0x3a, 0x51, 0x8a, 0xa5, 0xd0, 0x7e, 0x59, 0x5d, 0x98, 0x3b, 0x8c, 0x5, 0x26, 0xc8, 0xf7, 0xfb}
	transferBatchSchema     = schema.Parse("(uint256[],uint256[])")
	transferBatchNumIndexed = int(3)
)

// Event Signature:
//	TransferBatch(address,address,address,uint256[],uint256[])
// Checks the first log topic against the signature hash:
//	4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb
//
// Copies indexed event inputs from the remaining topics
// into [TransferBatch]
//
// Uses the the following abi schema to decode the un-indexed
// event inputs from the log's data field into [TransferBatch]:
//	(uint256[],uint256[])
func MatchTransferBatch(l abi.Log) (TransferBatch, error) {
	if len(l.Topics) > 0 && transferBatchSignature != l.Topics[0] {
		return TransferBatch{}, abi.SigMismatch
	}
	if len(l.Topics[1:]) != transferBatchNumIndexed {
		return TransferBatch{}, abi.IndexMismatch
	}
	item, _, err := abi.Decode(l.Data, transferBatchSchema)
	if err != nil {
		return TransferBatch{}, err
	}
	res := decodeTransferBatch(item)
	res.item = item
	res.Operator = abi.Bytes(l.Topics[1][:]).Address()
	res.From = abi.Bytes(l.Topics[2][:]).Address()
	res.To = abi.Bytes(l.Topics[3][:]).Address()
	return res, nil
}

type TransferSingle struct {
	item *abi.Item

	// Indexed:
	Operator [20]byte
	From     [20]byte
	To       [20]byte
	// Un-indexed:
	Id    *big.Int
	Value *big.Int
}

func (x TransferSingle) Done() {
	x.item.Done()
}

func decodeTransferSingle(item *abi.Item) TransferSingle {
	x := TransferSingle{}
	x.Id = item.At(0).BigInt()
	x.Value = item.At(1).BigInt()
	return x
}

var (
	transferSingleSignature  = [32]byte{0xc3, 0xd5, 0x81, 0x68, 0xc5, 0xae, 0x73, 0x97, 0x73, 0x1d, 0x6, 0x3d, 0x5b, 0xbf, 0x3d, 0x65, 0x78, 0x54, 0x42, 0x73, 0x43, 0xf4, 0xc0, 0x83, 0x24, 0xf, 0x7a, 0xac, 0xaa, 0x2d, 0xf, 0x62}
	transferSingleSchema     = schema.Parse("(uint256,uint256)")
	transferSingleNumIndexed = int(3)
)

// Event Signature:
//	TransferSingle(address,address,address,uint256,uint256)
// Checks the first log topic against the signature hash:
//	c3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62
//
// Copies indexed event inputs from the remaining topics
// into [TransferSingle]
//
// Uses the the following abi schema to decode the un-indexed
// event inputs from the log's data field into [TransferSingle]:
//	(uint256,uint256)
func MatchTransferSingle(l abi.Log) (TransferSingle, error) {
	if len(l.Topics) > 0 && transferSingleSignature != l.Topics[0] {
		return TransferSingle{}, abi.SigMismatch
	}
	if len(l.Topics[1:]) != transferSingleNumIndexed {
		return TransferSingle{}, abi.IndexMismatch
	}
	item, _, err := abi.Decode(l.Data, transferSingleSchema)
	if err != nil {
		return TransferSingle{}, err
	}
	res := decodeTransferSingle(item)
	res.item = item
	res.Operator = abi.Bytes(l.Topics[1][:]).Address()
	res.From = abi.Bytes(l.Topics[2][:]).Address()
	res.To = abi.Bytes(l.Topics[3][:]).Address()
	return res, nil
}

type URI struct {
	item *abi.Item

	// Indexed:
	Id *big.Int
	// Un-indexed:
	Value string
}

func (x URI) Done() {
	x.item.Done()
}

func decodeURI(item *abi.Item) URI {
	x := URI{}
	x.Value = item.At(0).String()
	return x
}

var (
	uRISignature  = [32]byte{0x6b, 0xb7, 0xff, 0x70, 0x86, 0x19, 0xba, 0x6, 0x10, 0xcb, 0xa2, 0x95, 0xa5, 0x85, 0x92, 0xe0, 0x45, 0x1d, 0xee, 0x26, 0x22, 0x93, 0x8c, 0x87, 0x55, 0x66, 0x76, 0x88, 0xda, 0xf3, 0x52, 0x9b}
	uRISchema     = schema.Parse("(string)")
	uRINumIndexed = int(1)
)

// Event Signature:
//	URI(string,uint256)
// Checks the first log topic against the signature hash:
//	6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b
//
// Copies indexed event inputs from the remaining topics
// into [URI]
//
// Uses the the following abi schema to decode the un-indexed
// event inputs from the log's data field into [URI]:
//	(string)
func MatchURI(l abi.Log) (URI, error) {
	if len(l.Topics) > 0 && uRISignature != l.Topics[0] {
		return URI{}, abi.SigMismatch
	}
	if len(l.Topics[1:]) != uRINumIndexed {
		return URI{}, abi.IndexMismatch
	}
	item, _, err := abi.Decode(l.Data, uRISchema)
	if err != nil {
		return URI{}, err
	}
	res := decodeURI(item)
	res.item = item
	res.Id = abi.Bytes(l.Topics[1][:]).BigInt()
	return res, nil
}
