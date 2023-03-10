// Parse event/log data based on the 'blur' ABI
//
// Code generated by "genabi"; DO NOT EDIT.
package blur

import (
	"github.com/indexsupply/x/abi"
	"github.com/indexsupply/x/abi/schema"
	"math/big"
)

type AdminChanged struct {
	item *abi.Item

	// Un-indexed:
	PreviousAdmin [20]byte
	NewAdmin      [20]byte
}

func (x AdminChanged) Done() {
	x.item.Done()
}

func decodeAdminChanged(item *abi.Item) AdminChanged {
	x := AdminChanged{}
	x.PreviousAdmin = item.At(0).Address()
	x.NewAdmin = item.At(1).Address()
	return x
}

var (
	adminChangedSignature  = [32]byte{0x7e, 0x64, 0x4d, 0x79, 0x42, 0x2f, 0x17, 0xc0, 0x1e, 0x48, 0x94, 0xb5, 0xf4, 0xf5, 0x88, 0xd3, 0x31, 0xeb, 0xfa, 0x28, 0x65, 0x3d, 0x42, 0xae, 0x83, 0x2d, 0xc5, 0x9e, 0x38, 0xc9, 0x79, 0x8f}
	adminChangedSchema     = schema.Parse("(address,address)")
	adminChangedNumIndexed = int(0)
)

// Event Signature:
//	AdminChanged(address,address)
// Checks the first log topic against the signature hash:
//	7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f
//
// Copies indexed event inputs from the remaining topics
// into [AdminChanged]
//
// Uses the the following abi schema to decode the un-indexed
// event inputs from the log's data field into [AdminChanged]:
//	(address,address)
func MatchAdminChanged(l abi.Log) (AdminChanged, error) {
	if len(l.Topics) == 0 {
		return AdminChanged{}, abi.NoTopics
	}
	if len(l.Topics) > 0 && adminChangedSignature != l.Topics[0] {
		return AdminChanged{}, abi.SigMismatch
	}
	if len(l.Topics[1:]) != adminChangedNumIndexed {
		return AdminChanged{}, abi.IndexMismatch
	}
	item, _, err := abi.Decode(l.Data, adminChangedSchema)
	if err != nil {
		return AdminChanged{}, err
	}
	res := decodeAdminChanged(item)
	res.item = item
	return res, nil
}

type BeaconUpgraded struct {
	item *abi.Item

	// Indexed:
	Beacon [20]byte
}

func (x BeaconUpgraded) Done() {
	x.item.Done()
}

func decodeBeaconUpgraded(item *abi.Item) BeaconUpgraded {
	x := BeaconUpgraded{}
	return x
}

var (
	beaconUpgradedSignature  = [32]byte{0x1c, 0xf3, 0xb0, 0x3a, 0x6c, 0xf1, 0x9f, 0xa2, 0xba, 0xba, 0x4d, 0xf1, 0x48, 0xe9, 0xdc, 0xab, 0xed, 0xea, 0x7f, 0x8a, 0x5c, 0x7, 0x84, 0xe, 0x20, 0x7e, 0x5c, 0x8, 0x9b, 0xe9, 0x5d, 0x3e}
	beaconUpgradedSchema     = schema.Parse("()")
	beaconUpgradedNumIndexed = int(1)
)

// Event Signature:
//	BeaconUpgraded(address)
// Checks the first log topic against the signature hash:
//	1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e
//
// Copies indexed event inputs from the remaining topics
// into [BeaconUpgraded]
//
// Uses the the following abi schema to decode the un-indexed
// event inputs from the log's data field into [BeaconUpgraded]:
//	()
func MatchBeaconUpgraded(l abi.Log) (BeaconUpgraded, error) {
	if len(l.Topics) == 0 {
		return BeaconUpgraded{}, abi.NoTopics
	}
	if len(l.Topics) > 0 && beaconUpgradedSignature != l.Topics[0] {
		return BeaconUpgraded{}, abi.SigMismatch
	}
	if len(l.Topics[1:]) != beaconUpgradedNumIndexed {
		return BeaconUpgraded{}, abi.IndexMismatch
	}
	res := BeaconUpgraded{}
	res.Beacon = abi.Bytes(l.Topics[1][:]).Address()
	return res, nil
}

type Closed struct {
	item *abi.Item
}

func (x Closed) Done() {
	x.item.Done()
}

func decodeClosed(item *abi.Item) Closed {
	x := Closed{}
	return x
}

var (
	closedSignature  = [32]byte{0x1c, 0xdd, 0xe6, 0x7b, 0x72, 0xa9, 0xf, 0x19, 0x91, 0x9a, 0xc7, 0x32, 0xa4, 0x37, 0xac, 0x2f, 0x7a, 0x10, 0xfc, 0x12, 0x8d, 0x28, 0xc2, 0xa6, 0xe5, 0x25, 0xd8, 0x9c, 0xe5, 0xcd, 0x9d, 0x3a}
	closedSchema     = schema.Parse("()")
	closedNumIndexed = int(0)
)

// Event Signature:
//	Closed()
// Checks the first log topic against the signature hash:
//	1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a
//
// Copies indexed event inputs from the remaining topics
// into [Closed]
//
// Uses the the following abi schema to decode the un-indexed
// event inputs from the log's data field into [Closed]:
//	()
func MatchClosed(l abi.Log) (Closed, error) {
	if len(l.Topics) == 0 {
		return Closed{}, abi.NoTopics
	}
	if len(l.Topics) > 0 && closedSignature != l.Topics[0] {
		return Closed{}, abi.SigMismatch
	}
	if len(l.Topics[1:]) != closedNumIndexed {
		return Closed{}, abi.IndexMismatch
	}
	res := Closed{}
	return res, nil
}

type Initialized struct {
	item *abi.Item

	// Un-indexed:
	Version uint8
}

func (x Initialized) Done() {
	x.item.Done()
}

func decodeInitialized(item *abi.Item) Initialized {
	x := Initialized{}
	x.Version = item.At(0).Uint8()
	return x
}

var (
	initializedSignature  = [32]byte{0x7f, 0x26, 0xb8, 0x3f, 0xf9, 0x6e, 0x1f, 0x2b, 0x6a, 0x68, 0x2f, 0x13, 0x38, 0x52, 0xf6, 0x79, 0x8a, 0x9, 0xc4, 0x65, 0xda, 0x95, 0x92, 0x14, 0x60, 0xce, 0xfb, 0x38, 0x47, 0x40, 0x24, 0x98}
	initializedSchema     = schema.Parse("(uint8)")
	initializedNumIndexed = int(0)
)

// Event Signature:
//	Initialized(uint8)
// Checks the first log topic against the signature hash:
//	7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498
//
// Copies indexed event inputs from the remaining topics
// into [Initialized]
//
// Uses the the following abi schema to decode the un-indexed
// event inputs from the log's data field into [Initialized]:
//	(uint8)
func MatchInitialized(l abi.Log) (Initialized, error) {
	if len(l.Topics) == 0 {
		return Initialized{}, abi.NoTopics
	}
	if len(l.Topics) > 0 && initializedSignature != l.Topics[0] {
		return Initialized{}, abi.SigMismatch
	}
	if len(l.Topics[1:]) != initializedNumIndexed {
		return Initialized{}, abi.IndexMismatch
	}
	item, _, err := abi.Decode(l.Data, initializedSchema)
	if err != nil {
		return Initialized{}, err
	}
	res := decodeInitialized(item)
	res.item = item
	return res, nil
}

type NewBlockRange struct {
	item *abi.Item

	// Un-indexed:
	BlockRange *big.Int
}

func (x NewBlockRange) Done() {
	x.item.Done()
}

func decodeNewBlockRange(item *abi.Item) NewBlockRange {
	x := NewBlockRange{}
	x.BlockRange = item.At(0).BigInt()
	return x
}

var (
	newBlockRangeSignature  = [32]byte{0x77, 0x6, 0x17, 0x7c, 0x54, 0x1b, 0xa1, 0xb8, 0x58, 0x37, 0x1b, 0xfc, 0x56, 0x8a, 0xa7, 0x74, 0x50, 0xb4, 0x71, 0x3b, 0xbd, 0xbb, 0xa6, 0x2c, 0x73, 0xd, 0x44, 0x84, 0xab, 0x6c, 0x12, 0x51}
	newBlockRangeSchema     = schema.Parse("(uint256)")
	newBlockRangeNumIndexed = int(0)
)

// Event Signature:
//	NewBlockRange(uint256)
// Checks the first log topic against the signature hash:
//	7706177c541ba1b858371bfc568aa77450b4713bbdbba62c730d4484ab6c1251
//
// Copies indexed event inputs from the remaining topics
// into [NewBlockRange]
//
// Uses the the following abi schema to decode the un-indexed
// event inputs from the log's data field into [NewBlockRange]:
//	(uint256)
func MatchNewBlockRange(l abi.Log) (NewBlockRange, error) {
	if len(l.Topics) == 0 {
		return NewBlockRange{}, abi.NoTopics
	}
	if len(l.Topics) > 0 && newBlockRangeSignature != l.Topics[0] {
		return NewBlockRange{}, abi.SigMismatch
	}
	if len(l.Topics[1:]) != newBlockRangeNumIndexed {
		return NewBlockRange{}, abi.IndexMismatch
	}
	item, _, err := abi.Decode(l.Data, newBlockRangeSchema)
	if err != nil {
		return NewBlockRange{}, err
	}
	res := decodeNewBlockRange(item)
	res.item = item
	return res, nil
}

type NewExecutionDelegate struct {
	item *abi.Item

	// Indexed:
	ExecutionDelegate [20]byte
}

func (x NewExecutionDelegate) Done() {
	x.item.Done()
}

func decodeNewExecutionDelegate(item *abi.Item) NewExecutionDelegate {
	x := NewExecutionDelegate{}
	return x
}

var (
	newExecutionDelegateSignature  = [32]byte{0xf9, 0xa0, 0xf3, 0x56, 0xa7, 0xef, 0x7, 0x93, 0x55, 0xde, 0x9, 0xd3, 0x2c, 0xe4, 0x5c, 0xc3, 0xcf, 0xab, 0xc8, 0xf1, 0x18, 0x68, 0x1c, 0x19, 0xa1, 0x75, 0x1, 0xf0, 0x5, 0xa3, 0x76, 0xac}
	newExecutionDelegateSchema     = schema.Parse("()")
	newExecutionDelegateNumIndexed = int(1)
)

// Event Signature:
//	NewExecutionDelegate(address)
// Checks the first log topic against the signature hash:
//	f9a0f356a7ef079355de09d32ce45cc3cfabc8f118681c19a17501f005a376ac
//
// Copies indexed event inputs from the remaining topics
// into [NewExecutionDelegate]
//
// Uses the the following abi schema to decode the un-indexed
// event inputs from the log's data field into [NewExecutionDelegate]:
//	()
func MatchNewExecutionDelegate(l abi.Log) (NewExecutionDelegate, error) {
	if len(l.Topics) == 0 {
		return NewExecutionDelegate{}, abi.NoTopics
	}
	if len(l.Topics) > 0 && newExecutionDelegateSignature != l.Topics[0] {
		return NewExecutionDelegate{}, abi.SigMismatch
	}
	if len(l.Topics[1:]) != newExecutionDelegateNumIndexed {
		return NewExecutionDelegate{}, abi.IndexMismatch
	}
	res := NewExecutionDelegate{}
	res.ExecutionDelegate = abi.Bytes(l.Topics[1][:]).Address()
	return res, nil
}

type NewFeeRate struct {
	item *abi.Item

	// Un-indexed:
	FeeRate *big.Int
}

func (x NewFeeRate) Done() {
	x.item.Done()
}

func decodeNewFeeRate(item *abi.Item) NewFeeRate {
	x := NewFeeRate{}
	x.FeeRate = item.At(0).BigInt()
	return x
}

var (
	newFeeRateSignature  = [32]byte{0x78, 0x89, 0x80, 0xe8, 0x2f, 0x46, 0x51, 0xcc, 0x86, 0xd1, 0xcc, 0x0, 0x91, 0x66, 0x85, 0x52, 0x8f, 0x16, 0xc9, 0xab, 0xb2, 0x1b, 0x2a, 0xfe, 0x72, 0x32, 0x54, 0x96, 0xc1, 0x8c, 0x94, 0xae}
	newFeeRateSchema     = schema.Parse("(uint256)")
	newFeeRateNumIndexed = int(0)
)

// Event Signature:
//	NewFeeRate(uint256)
// Checks the first log topic against the signature hash:
//	788980e82f4651cc86d1cc00916685528f16c9abb21b2afe72325496c18c94ae
//
// Copies indexed event inputs from the remaining topics
// into [NewFeeRate]
//
// Uses the the following abi schema to decode the un-indexed
// event inputs from the log's data field into [NewFeeRate]:
//	(uint256)
func MatchNewFeeRate(l abi.Log) (NewFeeRate, error) {
	if len(l.Topics) == 0 {
		return NewFeeRate{}, abi.NoTopics
	}
	if len(l.Topics) > 0 && newFeeRateSignature != l.Topics[0] {
		return NewFeeRate{}, abi.SigMismatch
	}
	if len(l.Topics[1:]) != newFeeRateNumIndexed {
		return NewFeeRate{}, abi.IndexMismatch
	}
	item, _, err := abi.Decode(l.Data, newFeeRateSchema)
	if err != nil {
		return NewFeeRate{}, err
	}
	res := decodeNewFeeRate(item)
	res.item = item
	return res, nil
}

type NewFeeRecipient struct {
	item *abi.Item

	// Un-indexed:
	FeeRecipient [20]byte
}

func (x NewFeeRecipient) Done() {
	x.item.Done()
}

func decodeNewFeeRecipient(item *abi.Item) NewFeeRecipient {
	x := NewFeeRecipient{}
	x.FeeRecipient = item.At(0).Address()
	return x
}

var (
	newFeeRecipientSignature  = [32]byte{0x41, 0x28, 0x71, 0x52, 0x9f, 0x3c, 0xed, 0xd6, 0xca, 0x6f, 0x10, 0x78, 0x42, 0x58, 0xf4, 0x96, 0x5a, 0x5d, 0x6e, 0x25, 0x41, 0x27, 0x59, 0x3f, 0xe3, 0x54, 0xe7, 0xa6, 0x2f, 0x6d, 0xa, 0x23}
	newFeeRecipientSchema     = schema.Parse("(address)")
	newFeeRecipientNumIndexed = int(0)
)

// Event Signature:
//	NewFeeRecipient(address)
// Checks the first log topic against the signature hash:
//	412871529f3cedd6ca6f10784258f4965a5d6e254127593fe354e7a62f6d0a23
//
// Copies indexed event inputs from the remaining topics
// into [NewFeeRecipient]
//
// Uses the the following abi schema to decode the un-indexed
// event inputs from the log's data field into [NewFeeRecipient]:
//	(address)
func MatchNewFeeRecipient(l abi.Log) (NewFeeRecipient, error) {
	if len(l.Topics) == 0 {
		return NewFeeRecipient{}, abi.NoTopics
	}
	if len(l.Topics) > 0 && newFeeRecipientSignature != l.Topics[0] {
		return NewFeeRecipient{}, abi.SigMismatch
	}
	if len(l.Topics[1:]) != newFeeRecipientNumIndexed {
		return NewFeeRecipient{}, abi.IndexMismatch
	}
	item, _, err := abi.Decode(l.Data, newFeeRecipientSchema)
	if err != nil {
		return NewFeeRecipient{}, err
	}
	res := decodeNewFeeRecipient(item)
	res.item = item
	return res, nil
}

type NewGovernor struct {
	item *abi.Item

	// Un-indexed:
	Governor [20]byte
}

func (x NewGovernor) Done() {
	x.item.Done()
}

func decodeNewGovernor(item *abi.Item) NewGovernor {
	x := NewGovernor{}
	x.Governor = item.At(0).Address()
	return x
}

var (
	newGovernorSignature  = [32]byte{0x54, 0x25, 0x36, 0x3a, 0x3, 0xf1, 0x82, 0x28, 0x11, 0x20, 0xf5, 0x91, 0x91, 0x7, 0xc4, 0x9c, 0x7a, 0x1a, 0x62, 0x3a, 0xcc, 0x1c, 0xbc, 0x6d, 0xf4, 0x68, 0xb6, 0xf0, 0xc1, 0x1f, 0xcf, 0x8c}
	newGovernorSchema     = schema.Parse("(address)")
	newGovernorNumIndexed = int(0)
)

// Event Signature:
//	NewGovernor(address)
// Checks the first log topic against the signature hash:
//	5425363a03f182281120f5919107c49c7a1a623acc1cbc6df468b6f0c11fcf8c
//
// Copies indexed event inputs from the remaining topics
// into [NewGovernor]
//
// Uses the the following abi schema to decode the un-indexed
// event inputs from the log's data field into [NewGovernor]:
//	(address)
func MatchNewGovernor(l abi.Log) (NewGovernor, error) {
	if len(l.Topics) == 0 {
		return NewGovernor{}, abi.NoTopics
	}
	if len(l.Topics) > 0 && newGovernorSignature != l.Topics[0] {
		return NewGovernor{}, abi.SigMismatch
	}
	if len(l.Topics[1:]) != newGovernorNumIndexed {
		return NewGovernor{}, abi.IndexMismatch
	}
	item, _, err := abi.Decode(l.Data, newGovernorSchema)
	if err != nil {
		return NewGovernor{}, err
	}
	res := decodeNewGovernor(item)
	res.item = item
	return res, nil
}

type NewOracle struct {
	item *abi.Item

	// Indexed:
	Oracle [20]byte
}

func (x NewOracle) Done() {
	x.item.Done()
}

func decodeNewOracle(item *abi.Item) NewOracle {
	x := NewOracle{}
	return x
}

var (
	newOracleSignature  = [32]byte{0xb3, 0xea, 0xcd, 0xe, 0x35, 0x1f, 0xaf, 0xdf, 0xef, 0xde, 0xc8, 0x4e, 0x1c, 0xd1, 0x96, 0x79, 0xb3, 0x8d, 0xbc, 0xd6, 0x3e, 0xa7, 0xc2, 0xc2, 0x4d, 0xa1, 0x7f, 0xd2, 0xbc, 0x3b, 0x3c, 0xe}
	newOracleSchema     = schema.Parse("()")
	newOracleNumIndexed = int(1)
)

// Event Signature:
//	NewOracle(address)
// Checks the first log topic against the signature hash:
//	b3eacd0e351fafdfefdec84e1cd19679b38dbcd63ea7c2c24da17fd2bc3b3c0e
//
// Copies indexed event inputs from the remaining topics
// into [NewOracle]
//
// Uses the the following abi schema to decode the un-indexed
// event inputs from the log's data field into [NewOracle]:
//	()
func MatchNewOracle(l abi.Log) (NewOracle, error) {
	if len(l.Topics) == 0 {
		return NewOracle{}, abi.NoTopics
	}
	if len(l.Topics) > 0 && newOracleSignature != l.Topics[0] {
		return NewOracle{}, abi.SigMismatch
	}
	if len(l.Topics[1:]) != newOracleNumIndexed {
		return NewOracle{}, abi.IndexMismatch
	}
	res := NewOracle{}
	res.Oracle = abi.Bytes(l.Topics[1][:]).Address()
	return res, nil
}

type NewPolicyManager struct {
	item *abi.Item

	// Indexed:
	PolicyManager [20]byte
}

func (x NewPolicyManager) Done() {
	x.item.Done()
}

func decodeNewPolicyManager(item *abi.Item) NewPolicyManager {
	x := NewPolicyManager{}
	return x
}

var (
	newPolicyManagerSignature  = [32]byte{0xdb, 0xe1, 0x8f, 0x3f, 0xd9, 0x27, 0xcc, 0x2a, 0xef, 0xe3, 0x80, 0xff, 0xd2, 0xab, 0xfd, 0xb8, 0xe1, 0x3f, 0x2, 0x39, 0xc0, 0x25, 0x8c, 0xcf, 0xc8, 0x4c, 0x3d, 0x8f, 0xdd, 0x8c, 0x4, 0x18}
	newPolicyManagerSchema     = schema.Parse("()")
	newPolicyManagerNumIndexed = int(1)
)

// Event Signature:
//	NewPolicyManager(address)
// Checks the first log topic against the signature hash:
//	dbe18f3fd927cc2aefe380ffd2abfdb8e13f0239c0258ccfc84c3d8fdd8c0418
//
// Copies indexed event inputs from the remaining topics
// into [NewPolicyManager]
//
// Uses the the following abi schema to decode the un-indexed
// event inputs from the log's data field into [NewPolicyManager]:
//	()
func MatchNewPolicyManager(l abi.Log) (NewPolicyManager, error) {
	if len(l.Topics) == 0 {
		return NewPolicyManager{}, abi.NoTopics
	}
	if len(l.Topics) > 0 && newPolicyManagerSignature != l.Topics[0] {
		return NewPolicyManager{}, abi.SigMismatch
	}
	if len(l.Topics[1:]) != newPolicyManagerNumIndexed {
		return NewPolicyManager{}, abi.IndexMismatch
	}
	res := NewPolicyManager{}
	res.PolicyManager = abi.Bytes(l.Topics[1][:]).Address()
	return res, nil
}

type NonceIncremented struct {
	item *abi.Item

	// Indexed:
	Trader [20]byte
	// Un-indexed:
	NewNonce *big.Int
}

func (x NonceIncremented) Done() {
	x.item.Done()
}

func decodeNonceIncremented(item *abi.Item) NonceIncremented {
	x := NonceIncremented{}
	x.NewNonce = item.At(0).BigInt()
	return x
}

var (
	nonceIncrementedSignature  = [32]byte{0xa8, 0x2a, 0x64, 0x9b, 0xbd, 0x6, 0xc, 0x90, 0x99, 0xcd, 0x7b, 0x73, 0x26, 0xe2, 0xb0, 0xdc, 0x9e, 0x9a, 0xf0, 0x83, 0x64, 0x80, 0xe0, 0xf8, 0x49, 0xdc, 0x9e, 0xaa, 0x79, 0x71, 0xb, 0x3b}
	nonceIncrementedSchema     = schema.Parse("(uint256)")
	nonceIncrementedNumIndexed = int(1)
)

// Event Signature:
//	NonceIncremented(address,uint256)
// Checks the first log topic against the signature hash:
//	a82a649bbd060c9099cd7b7326e2b0dc9e9af0836480e0f849dc9eaa79710b3b
//
// Copies indexed event inputs from the remaining topics
// into [NonceIncremented]
//
// Uses the the following abi schema to decode the un-indexed
// event inputs from the log's data field into [NonceIncremented]:
//	(uint256)
func MatchNonceIncremented(l abi.Log) (NonceIncremented, error) {
	if len(l.Topics) == 0 {
		return NonceIncremented{}, abi.NoTopics
	}
	if len(l.Topics) > 0 && nonceIncrementedSignature != l.Topics[0] {
		return NonceIncremented{}, abi.SigMismatch
	}
	if len(l.Topics[1:]) != nonceIncrementedNumIndexed {
		return NonceIncremented{}, abi.IndexMismatch
	}
	item, _, err := abi.Decode(l.Data, nonceIncrementedSchema)
	if err != nil {
		return NonceIncremented{}, err
	}
	res := decodeNonceIncremented(item)
	res.item = item
	res.Trader = abi.Bytes(l.Topics[1][:]).Address()
	return res, nil
}

type Opened struct {
	item *abi.Item
}

func (x Opened) Done() {
	x.item.Done()
}

func decodeOpened(item *abi.Item) Opened {
	x := Opened{}
	return x
}

var (
	openedSignature  = [32]byte{0xd1, 0xdc, 0xd0, 0x5, 0x34, 0x37, 0x3f, 0x20, 0x88, 0x2b, 0x79, 0xe6, 0xab, 0x68, 0x75, 0xa5, 0xc3, 0x58, 0xc5, 0xbd, 0x57, 0x64, 0x48, 0x75, 0x7e, 0xd5, 0xe, 0x63, 0x6, 0x9a, 0xb5, 0x18}
	openedSchema     = schema.Parse("()")
	openedNumIndexed = int(0)
)

// Event Signature:
//	Opened()
// Checks the first log topic against the signature hash:
//	d1dcd00534373f20882b79e6ab6875a5c358c5bd576448757ed50e63069ab518
//
// Copies indexed event inputs from the remaining topics
// into [Opened]
//
// Uses the the following abi schema to decode the un-indexed
// event inputs from the log's data field into [Opened]:
//	()
func MatchOpened(l abi.Log) (Opened, error) {
	if len(l.Topics) == 0 {
		return Opened{}, abi.NoTopics
	}
	if len(l.Topics) > 0 && openedSignature != l.Topics[0] {
		return Opened{}, abi.SigMismatch
	}
	if len(l.Topics[1:]) != openedNumIndexed {
		return Opened{}, abi.IndexMismatch
	}
	res := Opened{}
	return res, nil
}

type OrderCancelled struct {
	item *abi.Item

	// Un-indexed:
	Hash [32]byte
}

func (x OrderCancelled) Done() {
	x.item.Done()
}

func decodeOrderCancelled(item *abi.Item) OrderCancelled {
	x := OrderCancelled{}
	x.Hash = item.At(0).Bytes32()
	return x
}

var (
	orderCancelledSignature  = [32]byte{0x51, 0x52, 0xab, 0xf9, 0x59, 0xf6, 0x56, 0x46, 0x62, 0x35, 0x8c, 0x2e, 0x52, 0xb7, 0x2, 0x25, 0x9b, 0x78, 0xba, 0xc5, 0xee, 0x78, 0x42, 0xa0, 0xf0, 0x19, 0x37, 0xe6, 0x70, 0xef, 0xcc, 0x7d}
	orderCancelledSchema     = schema.Parse("(bytes32)")
	orderCancelledNumIndexed = int(0)
)

// Event Signature:
//	OrderCancelled(bytes32)
// Checks the first log topic against the signature hash:
//	5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d
//
// Copies indexed event inputs from the remaining topics
// into [OrderCancelled]
//
// Uses the the following abi schema to decode the un-indexed
// event inputs from the log's data field into [OrderCancelled]:
//	(bytes32)
func MatchOrderCancelled(l abi.Log) (OrderCancelled, error) {
	if len(l.Topics) == 0 {
		return OrderCancelled{}, abi.NoTopics
	}
	if len(l.Topics) > 0 && orderCancelledSignature != l.Topics[0] {
		return OrderCancelled{}, abi.SigMismatch
	}
	if len(l.Topics[1:]) != orderCancelledNumIndexed {
		return OrderCancelled{}, abi.IndexMismatch
	}
	item, _, err := abi.Decode(l.Data, orderCancelledSchema)
	if err != nil {
		return OrderCancelled{}, err
	}
	res := decodeOrderCancelled(item)
	res.item = item
	return res, nil
}

type OrdersMatched struct {
	item *abi.Item

	// Indexed:
	Maker [20]byte
	Taker [20]byte
	// Un-indexed:
	Sell     Sell
	SellHash [32]byte
	Buy      Buy
	BuyHash  [32]byte
}

func (x OrdersMatched) Done() {
	x.item.Done()
}

func decodeOrdersMatched(item *abi.Item) OrdersMatched {
	x := OrdersMatched{}
	x.Sell = decodeSell(item.At(0))
	x.SellHash = item.At(1).Bytes32()
	x.Buy = decodeBuy(item.At(2))
	x.BuyHash = item.At(3).Bytes32()
	return x
}

type Sell struct {
	item *abi.Item

	// Un-indexed:
	Trader         [20]byte
	Side           uint8
	MatchingPolicy [20]byte
	Collection     [20]byte
	TokenId        *big.Int
	Amount         *big.Int
	PaymentToken   [20]byte
	Price          *big.Int
	ListingTime    *big.Int
	ExpirationTime *big.Int
	Fees           []Fees
	Salt           *big.Int
	ExtraParams    []byte
}

func (x Sell) Done() {
	x.item.Done()
}

func decodeSell(item *abi.Item) Sell {
	x := Sell{}
	x.Trader = item.At(0).Address()
	x.Side = item.At(1).Uint8()
	x.MatchingPolicy = item.At(2).Address()
	x.Collection = item.At(3).Address()
	x.TokenId = item.At(4).BigInt()
	x.Amount = item.At(5).BigInt()
	x.PaymentToken = item.At(6).Address()
	x.Price = item.At(7).BigInt()
	x.ListingTime = item.At(8).BigInt()
	x.ExpirationTime = item.At(9).BigInt()
	feesItem0 := item.At(10)
	fees0 := make([]Fees, feesItem0.Len())
	for i0 := 0; i0 < feesItem0.Len(); i0++ {
		fees0[i0] = decodeFees(feesItem0.At(i0))
	}
	x.Fees = fees0
	x.Salt = item.At(11).BigInt()
	x.ExtraParams = item.At(12).Bytes()
	return x
}

type Fees struct {
	item *abi.Item

	// Un-indexed:
	Rate      uint16
	Recipient [20]byte
}

func (x Fees) Done() {
	x.item.Done()
}

func decodeFees(item *abi.Item) Fees {
	x := Fees{}
	x.Rate = item.At(0).Uint16()
	x.Recipient = item.At(1).Address()
	return x
}

type Buy struct {
	item *abi.Item

	// Un-indexed:
	Trader         [20]byte
	Side           uint8
	MatchingPolicy [20]byte
	Collection     [20]byte
	TokenId        *big.Int
	Amount         *big.Int
	PaymentToken   [20]byte
	Price          *big.Int
	ListingTime    *big.Int
	ExpirationTime *big.Int
	Fees           []Fees
	Salt           *big.Int
	ExtraParams    []byte
}

func (x Buy) Done() {
	x.item.Done()
}

func decodeBuy(item *abi.Item) Buy {
	x := Buy{}
	x.Trader = item.At(0).Address()
	x.Side = item.At(1).Uint8()
	x.MatchingPolicy = item.At(2).Address()
	x.Collection = item.At(3).Address()
	x.TokenId = item.At(4).BigInt()
	x.Amount = item.At(5).BigInt()
	x.PaymentToken = item.At(6).Address()
	x.Price = item.At(7).BigInt()
	x.ListingTime = item.At(8).BigInt()
	x.ExpirationTime = item.At(9).BigInt()
	feesItem0 := item.At(10)
	fees0 := make([]Fees, feesItem0.Len())
	for i0 := 0; i0 < feesItem0.Len(); i0++ {
		fees0[i0] = decodeFees(feesItem0.At(i0))
	}
	x.Fees = fees0
	x.Salt = item.At(11).BigInt()
	x.ExtraParams = item.At(12).Bytes()
	return x
}

var (
	ordersMatchedSignature  = [32]byte{0x61, 0xcb, 0xb2, 0xa3, 0xde, 0xe0, 0xb6, 0x6, 0x4c, 0x2e, 0x68, 0x1a, 0xad, 0xd6, 0x16, 0x77, 0xfb, 0x4e, 0xf3, 0x19, 0xf0, 0xb5, 0x47, 0x50, 0x8d, 0x49, 0x56, 0x26, 0xf5, 0xa6, 0x2f, 0x64}
	ordersMatchedSchema     = schema.Parse("((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),bytes32,(address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),bytes32)")
	ordersMatchedNumIndexed = int(2)
)

// Event Signature:
//	OrdersMatched(address,address,(address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),bytes32,(address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),bytes32)
// Checks the first log topic against the signature hash:
//	61cbb2a3dee0b6064c2e681aadd61677fb4ef319f0b547508d495626f5a62f64
//
// Copies indexed event inputs from the remaining topics
// into [OrdersMatched]
//
// Uses the the following abi schema to decode the un-indexed
// event inputs from the log's data field into [OrdersMatched]:
//	((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),bytes32,(address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),bytes32)
func MatchOrdersMatched(l abi.Log) (OrdersMatched, error) {
	if len(l.Topics) == 0 {
		return OrdersMatched{}, abi.NoTopics
	}
	if len(l.Topics) > 0 && ordersMatchedSignature != l.Topics[0] {
		return OrdersMatched{}, abi.SigMismatch
	}
	if len(l.Topics[1:]) != ordersMatchedNumIndexed {
		return OrdersMatched{}, abi.IndexMismatch
	}
	item, _, err := abi.Decode(l.Data, ordersMatchedSchema)
	if err != nil {
		return OrdersMatched{}, err
	}
	res := decodeOrdersMatched(item)
	res.item = item
	res.Maker = abi.Bytes(l.Topics[1][:]).Address()
	res.Taker = abi.Bytes(l.Topics[2][:]).Address()
	return res, nil
}

type OwnershipTransferred struct {
	item *abi.Item

	// Indexed:
	PreviousOwner [20]byte
	NewOwner      [20]byte
}

func (x OwnershipTransferred) Done() {
	x.item.Done()
}

func decodeOwnershipTransferred(item *abi.Item) OwnershipTransferred {
	x := OwnershipTransferred{}
	return x
}

var (
	ownershipTransferredSignature  = [32]byte{0x8b, 0xe0, 0x7, 0x9c, 0x53, 0x16, 0x59, 0x14, 0x13, 0x44, 0xcd, 0x1f, 0xd0, 0xa4, 0xf2, 0x84, 0x19, 0x49, 0x7f, 0x97, 0x22, 0xa3, 0xda, 0xaf, 0xe3, 0xb4, 0x18, 0x6f, 0x6b, 0x64, 0x57, 0xe0}
	ownershipTransferredSchema     = schema.Parse("()")
	ownershipTransferredNumIndexed = int(2)
)

// Event Signature:
//	OwnershipTransferred(address,address)
// Checks the first log topic against the signature hash:
//	8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0
//
// Copies indexed event inputs from the remaining topics
// into [OwnershipTransferred]
//
// Uses the the following abi schema to decode the un-indexed
// event inputs from the log's data field into [OwnershipTransferred]:
//	()
func MatchOwnershipTransferred(l abi.Log) (OwnershipTransferred, error) {
	if len(l.Topics) == 0 {
		return OwnershipTransferred{}, abi.NoTopics
	}
	if len(l.Topics) > 0 && ownershipTransferredSignature != l.Topics[0] {
		return OwnershipTransferred{}, abi.SigMismatch
	}
	if len(l.Topics[1:]) != ownershipTransferredNumIndexed {
		return OwnershipTransferred{}, abi.IndexMismatch
	}
	res := OwnershipTransferred{}
	res.PreviousOwner = abi.Bytes(l.Topics[1][:]).Address()
	res.NewOwner = abi.Bytes(l.Topics[2][:]).Address()
	return res, nil
}

type Upgraded struct {
	item *abi.Item

	// Indexed:
	Implementation [20]byte
}

func (x Upgraded) Done() {
	x.item.Done()
}

func decodeUpgraded(item *abi.Item) Upgraded {
	x := Upgraded{}
	return x
}

var (
	upgradedSignature  = [32]byte{0xbc, 0x7c, 0xd7, 0x5a, 0x20, 0xee, 0x27, 0xfd, 0x9a, 0xde, 0xba, 0xb3, 0x20, 0x41, 0xf7, 0x55, 0x21, 0x4d, 0xbc, 0x6b, 0xff, 0xa9, 0xc, 0xc0, 0x22, 0x5b, 0x39, 0xda, 0x2e, 0x5c, 0x2d, 0x3b}
	upgradedSchema     = schema.Parse("()")
	upgradedNumIndexed = int(1)
)

// Event Signature:
//	Upgraded(address)
// Checks the first log topic against the signature hash:
//	bc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b
//
// Copies indexed event inputs from the remaining topics
// into [Upgraded]
//
// Uses the the following abi schema to decode the un-indexed
// event inputs from the log's data field into [Upgraded]:
//	()
func MatchUpgraded(l abi.Log) (Upgraded, error) {
	if len(l.Topics) == 0 {
		return Upgraded{}, abi.NoTopics
	}
	if len(l.Topics) > 0 && upgradedSignature != l.Topics[0] {
		return Upgraded{}, abi.SigMismatch
	}
	if len(l.Topics[1:]) != upgradedNumIndexed {
		return Upgraded{}, abi.IndexMismatch
	}
	res := Upgraded{}
	res.Implementation = abi.Bytes(l.Topics[1][:]).Address()
	return res, nil
}
