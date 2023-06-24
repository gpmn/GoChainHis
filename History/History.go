// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package History

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// HistoryMetaData contains all meta data concerning the History contract.
var HistoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"daySlot\",\"type\":\"uint64\"}],\"name\":\"NewCard\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"daySlot\",\"type\":\"uint64\"},{\"internalType\":\"string[]\",\"name\":\"candidates\",\"type\":\"string[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"AirDrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"daySlot\",\"type\":\"uint64\"}],\"name\":\"Bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CardAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CardRewardsMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ClaimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ClaimedRewardsMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Echo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EscrowAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"daySlot\",\"type\":\"uint64\"}],\"name\":\"GetBigStories\",\"outputs\":[{\"internalType\":\"uint8[3]\",\"name\":\"bis\",\"type\":\"uint8[3]\"},{\"internalType\":\"string[3]\",\"name\":\"content\",\"type\":\"string[3]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"daySlot\",\"type\":\"uint64\"}],\"name\":\"GetHisRecBigIdxs\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"id0\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"id1\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"id2\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"daySlot\",\"type\":\"uint64\"}],\"name\":\"GetHisRecStoriesCnt\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"cnt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"daySlot\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"storyIdx\",\"type\":\"uint8\"}],\"name\":\"GetHisRecStoryAt\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"rsvWeight\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"voteSum\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"daySlot\",\"type\":\"uint64\"}],\"name\":\"GetStatusOf\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"daySlot\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"GetVoteInfo\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"prefer0\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"prefer1\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"prefer2\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"voteAmt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"daySlot\",\"type\":\"uint64\"}],\"name\":\"GetWinnerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"HistoryMap\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"Status\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"ShareOfDev\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"ShareOfSecretary\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"ShareOfCards\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"ShareOfVoters\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"VoterCnt\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"VoteEndTm\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"AucInitTm\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"AucInitPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"FinalPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"VoteSum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"WeightedVoteSum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"Winner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAAuctionPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"daySlot\",\"type\":\"uint64\"},{\"internalType\":\"string[]\",\"name\":\"festivals\",\"type\":\"string[]\"}],\"name\":\"MintAndAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"OpsRewardsMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"PendingRewardsMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"tokenId\",\"type\":\"uint64\"}],\"name\":\"QueryCurPrice\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"hisStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"curPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"daySlot\",\"type\":\"uint64\"}],\"name\":\"Resolve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"escAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cardAddr\",\"type\":\"address\"}],\"name\":\"SetAddrs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"SetSecretary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"mySlot\",\"type\":\"uint64\"},{\"internalType\":\"uint64[]\",\"name\":\"otherSlots\",\"type\":\"uint64[]\"}],\"name\":\"SettleCardReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"slots\",\"type\":\"uint64[]\"}],\"name\":\"SettleOpsReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"daySlots\",\"type\":\"uint64[]\"}],\"name\":\"SettleVoteReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"daySlot\",\"type\":\"uint64\"},{\"internalType\":\"string[]\",\"name\":\"candidates\",\"type\":\"string[]\"}],\"name\":\"SubmitCandidates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"daySlot\",\"type\":\"uint64\"},{\"internalType\":\"uint8[3]\",\"name\":\"prefers\",\"type\":\"uint8[3]\"}],\"name\":\"Vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"VoteInfoMap\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"Status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"VoteAmt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Reward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"secretary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// HistoryABI is the input ABI used to generate the binding from.
// Deprecated: Use HistoryMetaData.ABI instead.
var HistoryABI = HistoryMetaData.ABI

// History is an auto generated Go binding around an Ethereum contract.
type History struct {
	HistoryCaller     // Read-only binding to the contract
	HistoryTransactor // Write-only binding to the contract
	HistoryFilterer   // Log filterer for contract events
}

// HistoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type HistoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HistoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HistoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HistoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HistoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HistorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HistorySession struct {
	Contract     *History          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HistoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HistoryCallerSession struct {
	Contract *HistoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// HistoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HistoryTransactorSession struct {
	Contract     *HistoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// HistoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type HistoryRaw struct {
	Contract *History // Generic contract binding to access the raw methods on
}

// HistoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HistoryCallerRaw struct {
	Contract *HistoryCaller // Generic read-only contract binding to access the raw methods on
}

// HistoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HistoryTransactorRaw struct {
	Contract *HistoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHistory creates a new instance of History, bound to a specific deployed contract.
func NewHistory(address common.Address, backend bind.ContractBackend) (*History, error) {
	contract, err := bindHistory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &History{HistoryCaller: HistoryCaller{contract: contract}, HistoryTransactor: HistoryTransactor{contract: contract}, HistoryFilterer: HistoryFilterer{contract: contract}}, nil
}

// NewHistoryCaller creates a new read-only instance of History, bound to a specific deployed contract.
func NewHistoryCaller(address common.Address, caller bind.ContractCaller) (*HistoryCaller, error) {
	contract, err := bindHistory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HistoryCaller{contract: contract}, nil
}

// NewHistoryTransactor creates a new write-only instance of History, bound to a specific deployed contract.
func NewHistoryTransactor(address common.Address, transactor bind.ContractTransactor) (*HistoryTransactor, error) {
	contract, err := bindHistory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HistoryTransactor{contract: contract}, nil
}

// NewHistoryFilterer creates a new log filterer instance of History, bound to a specific deployed contract.
func NewHistoryFilterer(address common.Address, filterer bind.ContractFilterer) (*HistoryFilterer, error) {
	contract, err := bindHistory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HistoryFilterer{contract: contract}, nil
}

// bindHistory binds a generic wrapper to an already deployed contract.
func bindHistory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HistoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_History *HistoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _History.Contract.HistoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_History *HistoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _History.Contract.HistoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_History *HistoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _History.Contract.HistoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_History *HistoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _History.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_History *HistoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _History.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_History *HistoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _History.Contract.contract.Transact(opts, method, params...)
}

// CardAddr is a free data retrieval call binding the contract method 0x4c704278.
//
// Solidity: function CardAddr() view returns(address)
func (_History *HistoryCaller) CardAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _History.contract.Call(opts, &out, "CardAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CardAddr is a free data retrieval call binding the contract method 0x4c704278.
//
// Solidity: function CardAddr() view returns(address)
func (_History *HistorySession) CardAddr() (common.Address, error) {
	return _History.Contract.CardAddr(&_History.CallOpts)
}

// CardAddr is a free data retrieval call binding the contract method 0x4c704278.
//
// Solidity: function CardAddr() view returns(address)
func (_History *HistoryCallerSession) CardAddr() (common.Address, error) {
	return _History.Contract.CardAddr(&_History.CallOpts)
}

// CardRewardsMap is a free data retrieval call binding the contract method 0x68e3c52d.
//
// Solidity: function CardRewardsMap(uint64 , uint256 ) view returns(uint256)
func (_History *HistoryCaller) CardRewardsMap(opts *bind.CallOpts, arg0 uint64, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _History.contract.Call(opts, &out, "CardRewardsMap", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CardRewardsMap is a free data retrieval call binding the contract method 0x68e3c52d.
//
// Solidity: function CardRewardsMap(uint64 , uint256 ) view returns(uint256)
func (_History *HistorySession) CardRewardsMap(arg0 uint64, arg1 *big.Int) (*big.Int, error) {
	return _History.Contract.CardRewardsMap(&_History.CallOpts, arg0, arg1)
}

// CardRewardsMap is a free data retrieval call binding the contract method 0x68e3c52d.
//
// Solidity: function CardRewardsMap(uint64 , uint256 ) view returns(uint256)
func (_History *HistoryCallerSession) CardRewardsMap(arg0 uint64, arg1 *big.Int) (*big.Int, error) {
	return _History.Contract.CardRewardsMap(&_History.CallOpts, arg0, arg1)
}

// ClaimedRewardsMap is a free data retrieval call binding the contract method 0x106330e6.
//
// Solidity: function ClaimedRewardsMap(address ) view returns(uint256)
func (_History *HistoryCaller) ClaimedRewardsMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _History.contract.Call(opts, &out, "ClaimedRewardsMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedRewardsMap is a free data retrieval call binding the contract method 0x106330e6.
//
// Solidity: function ClaimedRewardsMap(address ) view returns(uint256)
func (_History *HistorySession) ClaimedRewardsMap(arg0 common.Address) (*big.Int, error) {
	return _History.Contract.ClaimedRewardsMap(&_History.CallOpts, arg0)
}

// ClaimedRewardsMap is a free data retrieval call binding the contract method 0x106330e6.
//
// Solidity: function ClaimedRewardsMap(address ) view returns(uint256)
func (_History *HistoryCallerSession) ClaimedRewardsMap(arg0 common.Address) (*big.Int, error) {
	return _History.Contract.ClaimedRewardsMap(&_History.CallOpts, arg0)
}

// Echo is a free data retrieval call binding the contract method 0x24c4b0ef.
//
// Solidity: function Echo() pure returns(string)
func (_History *HistoryCaller) Echo(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _History.contract.Call(opts, &out, "Echo")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Echo is a free data retrieval call binding the contract method 0x24c4b0ef.
//
// Solidity: function Echo() pure returns(string)
func (_History *HistorySession) Echo() (string, error) {
	return _History.Contract.Echo(&_History.CallOpts)
}

// Echo is a free data retrieval call binding the contract method 0x24c4b0ef.
//
// Solidity: function Echo() pure returns(string)
func (_History *HistoryCallerSession) Echo() (string, error) {
	return _History.Contract.Echo(&_History.CallOpts)
}

// EscrowAddr is a free data retrieval call binding the contract method 0x72f5e8dc.
//
// Solidity: function EscrowAddr() view returns(address)
func (_History *HistoryCaller) EscrowAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _History.contract.Call(opts, &out, "EscrowAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EscrowAddr is a free data retrieval call binding the contract method 0x72f5e8dc.
//
// Solidity: function EscrowAddr() view returns(address)
func (_History *HistorySession) EscrowAddr() (common.Address, error) {
	return _History.Contract.EscrowAddr(&_History.CallOpts)
}

// EscrowAddr is a free data retrieval call binding the contract method 0x72f5e8dc.
//
// Solidity: function EscrowAddr() view returns(address)
func (_History *HistoryCallerSession) EscrowAddr() (common.Address, error) {
	return _History.Contract.EscrowAddr(&_History.CallOpts)
}

// GetBigStories is a free data retrieval call binding the contract method 0x797fca74.
//
// Solidity: function GetBigStories(uint64 daySlot) view returns(uint8[3] bis, string[3] content)
func (_History *HistoryCaller) GetBigStories(opts *bind.CallOpts, daySlot uint64) (struct {
	Bis     [3]uint8
	Content [3]string
}, error) {
	var out []interface{}
	err := _History.contract.Call(opts, &out, "GetBigStories", daySlot)

	outstruct := new(struct {
		Bis     [3]uint8
		Content [3]string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bis = *abi.ConvertType(out[0], new([3]uint8)).(*[3]uint8)
	outstruct.Content = *abi.ConvertType(out[1], new([3]string)).(*[3]string)

	return *outstruct, err

}

// GetBigStories is a free data retrieval call binding the contract method 0x797fca74.
//
// Solidity: function GetBigStories(uint64 daySlot) view returns(uint8[3] bis, string[3] content)
func (_History *HistorySession) GetBigStories(daySlot uint64) (struct {
	Bis     [3]uint8
	Content [3]string
}, error) {
	return _History.Contract.GetBigStories(&_History.CallOpts, daySlot)
}

// GetBigStories is a free data retrieval call binding the contract method 0x797fca74.
//
// Solidity: function GetBigStories(uint64 daySlot) view returns(uint8[3] bis, string[3] content)
func (_History *HistoryCallerSession) GetBigStories(daySlot uint64) (struct {
	Bis     [3]uint8
	Content [3]string
}, error) {
	return _History.Contract.GetBigStories(&_History.CallOpts, daySlot)
}

// GetHisRecBigIdxs is a free data retrieval call binding the contract method 0xe28555da.
//
// Solidity: function GetHisRecBigIdxs(uint64 daySlot) view returns(uint8 id0, uint8 id1, uint8 id2)
func (_History *HistoryCaller) GetHisRecBigIdxs(opts *bind.CallOpts, daySlot uint64) (struct {
	Id0 uint8
	Id1 uint8
	Id2 uint8
}, error) {
	var out []interface{}
	err := _History.contract.Call(opts, &out, "GetHisRecBigIdxs", daySlot)

	outstruct := new(struct {
		Id0 uint8
		Id1 uint8
		Id2 uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id0 = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Id1 = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Id2 = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// GetHisRecBigIdxs is a free data retrieval call binding the contract method 0xe28555da.
//
// Solidity: function GetHisRecBigIdxs(uint64 daySlot) view returns(uint8 id0, uint8 id1, uint8 id2)
func (_History *HistorySession) GetHisRecBigIdxs(daySlot uint64) (struct {
	Id0 uint8
	Id1 uint8
	Id2 uint8
}, error) {
	return _History.Contract.GetHisRecBigIdxs(&_History.CallOpts, daySlot)
}

// GetHisRecBigIdxs is a free data retrieval call binding the contract method 0xe28555da.
//
// Solidity: function GetHisRecBigIdxs(uint64 daySlot) view returns(uint8 id0, uint8 id1, uint8 id2)
func (_History *HistoryCallerSession) GetHisRecBigIdxs(daySlot uint64) (struct {
	Id0 uint8
	Id1 uint8
	Id2 uint8
}, error) {
	return _History.Contract.GetHisRecBigIdxs(&_History.CallOpts, daySlot)
}

// GetHisRecStoriesCnt is a free data retrieval call binding the contract method 0xf1bd8059.
//
// Solidity: function GetHisRecStoriesCnt(uint64 daySlot) view returns(uint64 cnt)
func (_History *HistoryCaller) GetHisRecStoriesCnt(opts *bind.CallOpts, daySlot uint64) (uint64, error) {
	var out []interface{}
	err := _History.contract.Call(opts, &out, "GetHisRecStoriesCnt", daySlot)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetHisRecStoriesCnt is a free data retrieval call binding the contract method 0xf1bd8059.
//
// Solidity: function GetHisRecStoriesCnt(uint64 daySlot) view returns(uint64 cnt)
func (_History *HistorySession) GetHisRecStoriesCnt(daySlot uint64) (uint64, error) {
	return _History.Contract.GetHisRecStoriesCnt(&_History.CallOpts, daySlot)
}

// GetHisRecStoriesCnt is a free data retrieval call binding the contract method 0xf1bd8059.
//
// Solidity: function GetHisRecStoriesCnt(uint64 daySlot) view returns(uint64 cnt)
func (_History *HistoryCallerSession) GetHisRecStoriesCnt(daySlot uint64) (uint64, error) {
	return _History.Contract.GetHisRecStoriesCnt(&_History.CallOpts, daySlot)
}

// GetHisRecStoryAt is a free data retrieval call binding the contract method 0xc9840f44.
//
// Solidity: function GetHisRecStoryAt(uint64 daySlot, uint8 storyIdx) view returns(uint8 rsvWeight, uint256 voteSum, string content)
func (_History *HistoryCaller) GetHisRecStoryAt(opts *bind.CallOpts, daySlot uint64, storyIdx uint8) (struct {
	RsvWeight uint8
	VoteSum   *big.Int
	Content   string
}, error) {
	var out []interface{}
	err := _History.contract.Call(opts, &out, "GetHisRecStoryAt", daySlot, storyIdx)

	outstruct := new(struct {
		RsvWeight uint8
		VoteSum   *big.Int
		Content   string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RsvWeight = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.VoteSum = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Content = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// GetHisRecStoryAt is a free data retrieval call binding the contract method 0xc9840f44.
//
// Solidity: function GetHisRecStoryAt(uint64 daySlot, uint8 storyIdx) view returns(uint8 rsvWeight, uint256 voteSum, string content)
func (_History *HistorySession) GetHisRecStoryAt(daySlot uint64, storyIdx uint8) (struct {
	RsvWeight uint8
	VoteSum   *big.Int
	Content   string
}, error) {
	return _History.Contract.GetHisRecStoryAt(&_History.CallOpts, daySlot, storyIdx)
}

// GetHisRecStoryAt is a free data retrieval call binding the contract method 0xc9840f44.
//
// Solidity: function GetHisRecStoryAt(uint64 daySlot, uint8 storyIdx) view returns(uint8 rsvWeight, uint256 voteSum, string content)
func (_History *HistoryCallerSession) GetHisRecStoryAt(daySlot uint64, storyIdx uint8) (struct {
	RsvWeight uint8
	VoteSum   *big.Int
	Content   string
}, error) {
	return _History.Contract.GetHisRecStoryAt(&_History.CallOpts, daySlot, storyIdx)
}

// GetStatusOf is a free data retrieval call binding the contract method 0x180d1d77.
//
// Solidity: function GetStatusOf(uint64 daySlot) view returns(uint8 status)
func (_History *HistoryCaller) GetStatusOf(opts *bind.CallOpts, daySlot uint64) (uint8, error) {
	var out []interface{}
	err := _History.contract.Call(opts, &out, "GetStatusOf", daySlot)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetStatusOf is a free data retrieval call binding the contract method 0x180d1d77.
//
// Solidity: function GetStatusOf(uint64 daySlot) view returns(uint8 status)
func (_History *HistorySession) GetStatusOf(daySlot uint64) (uint8, error) {
	return _History.Contract.GetStatusOf(&_History.CallOpts, daySlot)
}

// GetStatusOf is a free data retrieval call binding the contract method 0x180d1d77.
//
// Solidity: function GetStatusOf(uint64 daySlot) view returns(uint8 status)
func (_History *HistoryCallerSession) GetStatusOf(daySlot uint64) (uint8, error) {
	return _History.Contract.GetStatusOf(&_History.CallOpts, daySlot)
}

// GetVoteInfo is a free data retrieval call binding the contract method 0xe3cf9929.
//
// Solidity: function GetVoteInfo(uint64 daySlot, address voter) view returns(uint8 status, uint8 prefer0, uint8 prefer1, uint8 prefer2, uint256 voteAmt, uint256 reward)
func (_History *HistoryCaller) GetVoteInfo(opts *bind.CallOpts, daySlot uint64, voter common.Address) (struct {
	Status  uint8
	Prefer0 uint8
	Prefer1 uint8
	Prefer2 uint8
	VoteAmt *big.Int
	Reward  *big.Int
}, error) {
	var out []interface{}
	err := _History.contract.Call(opts, &out, "GetVoteInfo", daySlot, voter)

	outstruct := new(struct {
		Status  uint8
		Prefer0 uint8
		Prefer1 uint8
		Prefer2 uint8
		VoteAmt *big.Int
		Reward  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Prefer0 = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Prefer1 = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.Prefer2 = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.VoteAmt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Reward = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetVoteInfo is a free data retrieval call binding the contract method 0xe3cf9929.
//
// Solidity: function GetVoteInfo(uint64 daySlot, address voter) view returns(uint8 status, uint8 prefer0, uint8 prefer1, uint8 prefer2, uint256 voteAmt, uint256 reward)
func (_History *HistorySession) GetVoteInfo(daySlot uint64, voter common.Address) (struct {
	Status  uint8
	Prefer0 uint8
	Prefer1 uint8
	Prefer2 uint8
	VoteAmt *big.Int
	Reward  *big.Int
}, error) {
	return _History.Contract.GetVoteInfo(&_History.CallOpts, daySlot, voter)
}

// GetVoteInfo is a free data retrieval call binding the contract method 0xe3cf9929.
//
// Solidity: function GetVoteInfo(uint64 daySlot, address voter) view returns(uint8 status, uint8 prefer0, uint8 prefer1, uint8 prefer2, uint256 voteAmt, uint256 reward)
func (_History *HistoryCallerSession) GetVoteInfo(daySlot uint64, voter common.Address) (struct {
	Status  uint8
	Prefer0 uint8
	Prefer1 uint8
	Prefer2 uint8
	VoteAmt *big.Int
	Reward  *big.Int
}, error) {
	return _History.Contract.GetVoteInfo(&_History.CallOpts, daySlot, voter)
}

// GetWinnerOf is a free data retrieval call binding the contract method 0x245b3025.
//
// Solidity: function GetWinnerOf(uint64 daySlot) view returns(address winner)
func (_History *HistoryCaller) GetWinnerOf(opts *bind.CallOpts, daySlot uint64) (common.Address, error) {
	var out []interface{}
	err := _History.contract.Call(opts, &out, "GetWinnerOf", daySlot)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWinnerOf is a free data retrieval call binding the contract method 0x245b3025.
//
// Solidity: function GetWinnerOf(uint64 daySlot) view returns(address winner)
func (_History *HistorySession) GetWinnerOf(daySlot uint64) (common.Address, error) {
	return _History.Contract.GetWinnerOf(&_History.CallOpts, daySlot)
}

// GetWinnerOf is a free data retrieval call binding the contract method 0x245b3025.
//
// Solidity: function GetWinnerOf(uint64 daySlot) view returns(address winner)
func (_History *HistoryCallerSession) GetWinnerOf(daySlot uint64) (common.Address, error) {
	return _History.Contract.GetWinnerOf(&_History.CallOpts, daySlot)
}

// HistoryMap is a free data retrieval call binding the contract method 0x22cfc6ca.
//
// Solidity: function HistoryMap(uint64 ) view returns(uint8 Status, uint8 ShareOfDev, uint8 ShareOfSecretary, uint8 ShareOfCards, uint8 ShareOfVoters, uint32 VoterCnt, uint64 VoteEndTm, uint64 AucInitTm, uint256 AucInitPrice, uint256 FinalPrice, uint256 VoteSum, uint256 WeightedVoteSum, address Winner)
func (_History *HistoryCaller) HistoryMap(opts *bind.CallOpts, arg0 uint64) (struct {
	Status           uint8
	ShareOfDev       uint8
	ShareOfSecretary uint8
	ShareOfCards     uint8
	ShareOfVoters    uint8
	VoterCnt         uint32
	VoteEndTm        uint64
	AucInitTm        uint64
	AucInitPrice     *big.Int
	FinalPrice       *big.Int
	VoteSum          *big.Int
	WeightedVoteSum  *big.Int
	Winner           common.Address
}, error) {
	var out []interface{}
	err := _History.contract.Call(opts, &out, "HistoryMap", arg0)

	outstruct := new(struct {
		Status           uint8
		ShareOfDev       uint8
		ShareOfSecretary uint8
		ShareOfCards     uint8
		ShareOfVoters    uint8
		VoterCnt         uint32
		VoteEndTm        uint64
		AucInitTm        uint64
		AucInitPrice     *big.Int
		FinalPrice       *big.Int
		VoteSum          *big.Int
		WeightedVoteSum  *big.Int
		Winner           common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.ShareOfDev = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.ShareOfSecretary = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.ShareOfCards = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.ShareOfVoters = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.VoterCnt = *abi.ConvertType(out[5], new(uint32)).(*uint32)
	outstruct.VoteEndTm = *abi.ConvertType(out[6], new(uint64)).(*uint64)
	outstruct.AucInitTm = *abi.ConvertType(out[7], new(uint64)).(*uint64)
	outstruct.AucInitPrice = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.FinalPrice = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.VoteSum = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.WeightedVoteSum = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	outstruct.Winner = *abi.ConvertType(out[12], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// HistoryMap is a free data retrieval call binding the contract method 0x22cfc6ca.
//
// Solidity: function HistoryMap(uint64 ) view returns(uint8 Status, uint8 ShareOfDev, uint8 ShareOfSecretary, uint8 ShareOfCards, uint8 ShareOfVoters, uint32 VoterCnt, uint64 VoteEndTm, uint64 AucInitTm, uint256 AucInitPrice, uint256 FinalPrice, uint256 VoteSum, uint256 WeightedVoteSum, address Winner)
func (_History *HistorySession) HistoryMap(arg0 uint64) (struct {
	Status           uint8
	ShareOfDev       uint8
	ShareOfSecretary uint8
	ShareOfCards     uint8
	ShareOfVoters    uint8
	VoterCnt         uint32
	VoteEndTm        uint64
	AucInitTm        uint64
	AucInitPrice     *big.Int
	FinalPrice       *big.Int
	VoteSum          *big.Int
	WeightedVoteSum  *big.Int
	Winner           common.Address
}, error) {
	return _History.Contract.HistoryMap(&_History.CallOpts, arg0)
}

// HistoryMap is a free data retrieval call binding the contract method 0x22cfc6ca.
//
// Solidity: function HistoryMap(uint64 ) view returns(uint8 Status, uint8 ShareOfDev, uint8 ShareOfSecretary, uint8 ShareOfCards, uint8 ShareOfVoters, uint32 VoterCnt, uint64 VoteEndTm, uint64 AucInitTm, uint256 AucInitPrice, uint256 FinalPrice, uint256 VoteSum, uint256 WeightedVoteSum, address Winner)
func (_History *HistoryCallerSession) HistoryMap(arg0 uint64) (struct {
	Status           uint8
	ShareOfDev       uint8
	ShareOfSecretary uint8
	ShareOfCards     uint8
	ShareOfVoters    uint8
	VoterCnt         uint32
	VoteEndTm        uint64
	AucInitTm        uint64
	AucInitPrice     *big.Int
	FinalPrice       *big.Int
	VoteSum          *big.Int
	WeightedVoteSum  *big.Int
	Winner           common.Address
}, error) {
	return _History.Contract.HistoryMap(&_History.CallOpts, arg0)
}

// MAAuctionPrice is a free data retrieval call binding the contract method 0x2743c01e.
//
// Solidity: function MAAuctionPrice() view returns(uint256)
func (_History *HistoryCaller) MAAuctionPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _History.contract.Call(opts, &out, "MAAuctionPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAAuctionPrice is a free data retrieval call binding the contract method 0x2743c01e.
//
// Solidity: function MAAuctionPrice() view returns(uint256)
func (_History *HistorySession) MAAuctionPrice() (*big.Int, error) {
	return _History.Contract.MAAuctionPrice(&_History.CallOpts)
}

// MAAuctionPrice is a free data retrieval call binding the contract method 0x2743c01e.
//
// Solidity: function MAAuctionPrice() view returns(uint256)
func (_History *HistoryCallerSession) MAAuctionPrice() (*big.Int, error) {
	return _History.Contract.MAAuctionPrice(&_History.CallOpts)
}

// OpsRewardsMap is a free data retrieval call binding the contract method 0x9ee88f59.
//
// Solidity: function OpsRewardsMap(uint64 ) view returns(uint256)
func (_History *HistoryCaller) OpsRewardsMap(opts *bind.CallOpts, arg0 uint64) (*big.Int, error) {
	var out []interface{}
	err := _History.contract.Call(opts, &out, "OpsRewardsMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpsRewardsMap is a free data retrieval call binding the contract method 0x9ee88f59.
//
// Solidity: function OpsRewardsMap(uint64 ) view returns(uint256)
func (_History *HistorySession) OpsRewardsMap(arg0 uint64) (*big.Int, error) {
	return _History.Contract.OpsRewardsMap(&_History.CallOpts, arg0)
}

// OpsRewardsMap is a free data retrieval call binding the contract method 0x9ee88f59.
//
// Solidity: function OpsRewardsMap(uint64 ) view returns(uint256)
func (_History *HistoryCallerSession) OpsRewardsMap(arg0 uint64) (*big.Int, error) {
	return _History.Contract.OpsRewardsMap(&_History.CallOpts, arg0)
}

// PendingRewardsMap is a free data retrieval call binding the contract method 0xc48b317a.
//
// Solidity: function PendingRewardsMap(address ) view returns(uint256)
func (_History *HistoryCaller) PendingRewardsMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _History.contract.Call(opts, &out, "PendingRewardsMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingRewardsMap is a free data retrieval call binding the contract method 0xc48b317a.
//
// Solidity: function PendingRewardsMap(address ) view returns(uint256)
func (_History *HistorySession) PendingRewardsMap(arg0 common.Address) (*big.Int, error) {
	return _History.Contract.PendingRewardsMap(&_History.CallOpts, arg0)
}

// PendingRewardsMap is a free data retrieval call binding the contract method 0xc48b317a.
//
// Solidity: function PendingRewardsMap(address ) view returns(uint256)
func (_History *HistoryCallerSession) PendingRewardsMap(arg0 common.Address) (*big.Int, error) {
	return _History.Contract.PendingRewardsMap(&_History.CallOpts, arg0)
}

// QueryCurPrice is a free data retrieval call binding the contract method 0x0c0000bd.
//
// Solidity: function QueryCurPrice(uint64 tokenId) view returns(uint8 hisStatus, uint256 curPrice)
func (_History *HistoryCaller) QueryCurPrice(opts *bind.CallOpts, tokenId uint64) (struct {
	HisStatus uint8
	CurPrice  *big.Int
}, error) {
	var out []interface{}
	err := _History.contract.Call(opts, &out, "QueryCurPrice", tokenId)

	outstruct := new(struct {
		HisStatus uint8
		CurPrice  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.HisStatus = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.CurPrice = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QueryCurPrice is a free data retrieval call binding the contract method 0x0c0000bd.
//
// Solidity: function QueryCurPrice(uint64 tokenId) view returns(uint8 hisStatus, uint256 curPrice)
func (_History *HistorySession) QueryCurPrice(tokenId uint64) (struct {
	HisStatus uint8
	CurPrice  *big.Int
}, error) {
	return _History.Contract.QueryCurPrice(&_History.CallOpts, tokenId)
}

// QueryCurPrice is a free data retrieval call binding the contract method 0x0c0000bd.
//
// Solidity: function QueryCurPrice(uint64 tokenId) view returns(uint8 hisStatus, uint256 curPrice)
func (_History *HistoryCallerSession) QueryCurPrice(tokenId uint64) (struct {
	HisStatus uint8
	CurPrice  *big.Int
}, error) {
	return _History.Contract.QueryCurPrice(&_History.CallOpts, tokenId)
}

// VoteInfoMap is a free data retrieval call binding the contract method 0xdb1b7c54.
//
// Solidity: function VoteInfoMap(uint64 , address ) view returns(uint8 Status, uint256 VoteAmt, uint256 Reward)
func (_History *HistoryCaller) VoteInfoMap(opts *bind.CallOpts, arg0 uint64, arg1 common.Address) (struct {
	Status  uint8
	VoteAmt *big.Int
	Reward  *big.Int
}, error) {
	var out []interface{}
	err := _History.contract.Call(opts, &out, "VoteInfoMap", arg0, arg1)

	outstruct := new(struct {
		Status  uint8
		VoteAmt *big.Int
		Reward  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.VoteAmt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Reward = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// VoteInfoMap is a free data retrieval call binding the contract method 0xdb1b7c54.
//
// Solidity: function VoteInfoMap(uint64 , address ) view returns(uint8 Status, uint256 VoteAmt, uint256 Reward)
func (_History *HistorySession) VoteInfoMap(arg0 uint64, arg1 common.Address) (struct {
	Status  uint8
	VoteAmt *big.Int
	Reward  *big.Int
}, error) {
	return _History.Contract.VoteInfoMap(&_History.CallOpts, arg0, arg1)
}

// VoteInfoMap is a free data retrieval call binding the contract method 0xdb1b7c54.
//
// Solidity: function VoteInfoMap(uint64 , address ) view returns(uint8 Status, uint256 VoteAmt, uint256 Reward)
func (_History *HistoryCallerSession) VoteInfoMap(arg0 uint64, arg1 common.Address) (struct {
	Status  uint8
	VoteAmt *big.Int
	Reward  *big.Int
}, error) {
	return _History.Contract.VoteInfoMap(&_History.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_History *HistoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _History.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_History *HistorySession) Owner() (common.Address, error) {
	return _History.Contract.Owner(&_History.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_History *HistoryCallerSession) Owner() (common.Address, error) {
	return _History.Contract.Owner(&_History.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_History *HistoryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _History.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_History *HistorySession) Paused() (bool, error) {
	return _History.Contract.Paused(&_History.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_History *HistoryCallerSession) Paused() (bool, error) {
	return _History.Contract.Paused(&_History.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_History *HistoryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _History.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_History *HistorySession) ProxiableUUID() ([32]byte, error) {
	return _History.Contract.ProxiableUUID(&_History.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_History *HistoryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _History.Contract.ProxiableUUID(&_History.CallOpts)
}

// Secretary is a free data retrieval call binding the contract method 0x5495d2aa.
//
// Solidity: function secretary() view returns(address)
func (_History *HistoryCaller) Secretary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _History.contract.Call(opts, &out, "secretary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Secretary is a free data retrieval call binding the contract method 0x5495d2aa.
//
// Solidity: function secretary() view returns(address)
func (_History *HistorySession) Secretary() (common.Address, error) {
	return _History.Contract.Secretary(&_History.CallOpts)
}

// Secretary is a free data retrieval call binding the contract method 0x5495d2aa.
//
// Solidity: function secretary() view returns(address)
func (_History *HistoryCallerSession) Secretary() (common.Address, error) {
	return _History.Contract.Secretary(&_History.CallOpts)
}

// AirDrop is a paid mutator transaction binding the contract method 0xfbb21004.
//
// Solidity: function AirDrop(uint64 daySlot, string[] candidates, address to) returns()
func (_History *HistoryTransactor) AirDrop(opts *bind.TransactOpts, daySlot uint64, candidates []string, to common.Address) (*types.Transaction, error) {
	return _History.contract.Transact(opts, "AirDrop", daySlot, candidates, to)
}

// AirDrop is a paid mutator transaction binding the contract method 0xfbb21004.
//
// Solidity: function AirDrop(uint64 daySlot, string[] candidates, address to) returns()
func (_History *HistorySession) AirDrop(daySlot uint64, candidates []string, to common.Address) (*types.Transaction, error) {
	return _History.Contract.AirDrop(&_History.TransactOpts, daySlot, candidates, to)
}

// AirDrop is a paid mutator transaction binding the contract method 0xfbb21004.
//
// Solidity: function AirDrop(uint64 daySlot, string[] candidates, address to) returns()
func (_History *HistoryTransactorSession) AirDrop(daySlot uint64, candidates []string, to common.Address) (*types.Transaction, error) {
	return _History.Contract.AirDrop(&_History.TransactOpts, daySlot, candidates, to)
}

// Bid is a paid mutator transaction binding the contract method 0x82766e8c.
//
// Solidity: function Bid(uint64 daySlot) payable returns()
func (_History *HistoryTransactor) Bid(opts *bind.TransactOpts, daySlot uint64) (*types.Transaction, error) {
	return _History.contract.Transact(opts, "Bid", daySlot)
}

// Bid is a paid mutator transaction binding the contract method 0x82766e8c.
//
// Solidity: function Bid(uint64 daySlot) payable returns()
func (_History *HistorySession) Bid(daySlot uint64) (*types.Transaction, error) {
	return _History.Contract.Bid(&_History.TransactOpts, daySlot)
}

// Bid is a paid mutator transaction binding the contract method 0x82766e8c.
//
// Solidity: function Bid(uint64 daySlot) payable returns()
func (_History *HistoryTransactorSession) Bid(daySlot uint64) (*types.Transaction, error) {
	return _History.Contract.Bid(&_History.TransactOpts, daySlot)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x79372f9a.
//
// Solidity: function ClaimReward() returns()
func (_History *HistoryTransactor) ClaimReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _History.contract.Transact(opts, "ClaimReward")
}

// ClaimReward is a paid mutator transaction binding the contract method 0x79372f9a.
//
// Solidity: function ClaimReward() returns()
func (_History *HistorySession) ClaimReward() (*types.Transaction, error) {
	return _History.Contract.ClaimReward(&_History.TransactOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x79372f9a.
//
// Solidity: function ClaimReward() returns()
func (_History *HistoryTransactorSession) ClaimReward() (*types.Transaction, error) {
	return _History.Contract.ClaimReward(&_History.TransactOpts)
}

// MintAndAuction is a paid mutator transaction binding the contract method 0xc6ac1460.
//
// Solidity: function MintAndAuction(uint64 daySlot, string[] festivals) returns()
func (_History *HistoryTransactor) MintAndAuction(opts *bind.TransactOpts, daySlot uint64, festivals []string) (*types.Transaction, error) {
	return _History.contract.Transact(opts, "MintAndAuction", daySlot, festivals)
}

// MintAndAuction is a paid mutator transaction binding the contract method 0xc6ac1460.
//
// Solidity: function MintAndAuction(uint64 daySlot, string[] festivals) returns()
func (_History *HistorySession) MintAndAuction(daySlot uint64, festivals []string) (*types.Transaction, error) {
	return _History.Contract.MintAndAuction(&_History.TransactOpts, daySlot, festivals)
}

// MintAndAuction is a paid mutator transaction binding the contract method 0xc6ac1460.
//
// Solidity: function MintAndAuction(uint64 daySlot, string[] festivals) returns()
func (_History *HistoryTransactorSession) MintAndAuction(daySlot uint64, festivals []string) (*types.Transaction, error) {
	return _History.Contract.MintAndAuction(&_History.TransactOpts, daySlot, festivals)
}

// Resolve is a paid mutator transaction binding the contract method 0x333efd86.
//
// Solidity: function Resolve(uint64 daySlot) returns()
func (_History *HistoryTransactor) Resolve(opts *bind.TransactOpts, daySlot uint64) (*types.Transaction, error) {
	return _History.contract.Transact(opts, "Resolve", daySlot)
}

// Resolve is a paid mutator transaction binding the contract method 0x333efd86.
//
// Solidity: function Resolve(uint64 daySlot) returns()
func (_History *HistorySession) Resolve(daySlot uint64) (*types.Transaction, error) {
	return _History.Contract.Resolve(&_History.TransactOpts, daySlot)
}

// Resolve is a paid mutator transaction binding the contract method 0x333efd86.
//
// Solidity: function Resolve(uint64 daySlot) returns()
func (_History *HistoryTransactorSession) Resolve(daySlot uint64) (*types.Transaction, error) {
	return _History.Contract.Resolve(&_History.TransactOpts, daySlot)
}

// SetAddrs is a paid mutator transaction binding the contract method 0x17a210d0.
//
// Solidity: function SetAddrs(address escAddr, address cardAddr) returns()
func (_History *HistoryTransactor) SetAddrs(opts *bind.TransactOpts, escAddr common.Address, cardAddr common.Address) (*types.Transaction, error) {
	return _History.contract.Transact(opts, "SetAddrs", escAddr, cardAddr)
}

// SetAddrs is a paid mutator transaction binding the contract method 0x17a210d0.
//
// Solidity: function SetAddrs(address escAddr, address cardAddr) returns()
func (_History *HistorySession) SetAddrs(escAddr common.Address, cardAddr common.Address) (*types.Transaction, error) {
	return _History.Contract.SetAddrs(&_History.TransactOpts, escAddr, cardAddr)
}

// SetAddrs is a paid mutator transaction binding the contract method 0x17a210d0.
//
// Solidity: function SetAddrs(address escAddr, address cardAddr) returns()
func (_History *HistoryTransactorSession) SetAddrs(escAddr common.Address, cardAddr common.Address) (*types.Transaction, error) {
	return _History.Contract.SetAddrs(&_History.TransactOpts, escAddr, cardAddr)
}

// SetSecretary is a paid mutator transaction binding the contract method 0xd01ab2e2.
//
// Solidity: function SetSecretary(address addr) returns()
func (_History *HistoryTransactor) SetSecretary(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _History.contract.Transact(opts, "SetSecretary", addr)
}

// SetSecretary is a paid mutator transaction binding the contract method 0xd01ab2e2.
//
// Solidity: function SetSecretary(address addr) returns()
func (_History *HistorySession) SetSecretary(addr common.Address) (*types.Transaction, error) {
	return _History.Contract.SetSecretary(&_History.TransactOpts, addr)
}

// SetSecretary is a paid mutator transaction binding the contract method 0xd01ab2e2.
//
// Solidity: function SetSecretary(address addr) returns()
func (_History *HistoryTransactorSession) SetSecretary(addr common.Address) (*types.Transaction, error) {
	return _History.Contract.SetSecretary(&_History.TransactOpts, addr)
}

// SettleCardReward is a paid mutator transaction binding the contract method 0x63bcfa05.
//
// Solidity: function SettleCardReward(uint64 mySlot, uint64[] otherSlots) returns()
func (_History *HistoryTransactor) SettleCardReward(opts *bind.TransactOpts, mySlot uint64, otherSlots []uint64) (*types.Transaction, error) {
	return _History.contract.Transact(opts, "SettleCardReward", mySlot, otherSlots)
}

// SettleCardReward is a paid mutator transaction binding the contract method 0x63bcfa05.
//
// Solidity: function SettleCardReward(uint64 mySlot, uint64[] otherSlots) returns()
func (_History *HistorySession) SettleCardReward(mySlot uint64, otherSlots []uint64) (*types.Transaction, error) {
	return _History.Contract.SettleCardReward(&_History.TransactOpts, mySlot, otherSlots)
}

// SettleCardReward is a paid mutator transaction binding the contract method 0x63bcfa05.
//
// Solidity: function SettleCardReward(uint64 mySlot, uint64[] otherSlots) returns()
func (_History *HistoryTransactorSession) SettleCardReward(mySlot uint64, otherSlots []uint64) (*types.Transaction, error) {
	return _History.Contract.SettleCardReward(&_History.TransactOpts, mySlot, otherSlots)
}

// SettleOpsReward is a paid mutator transaction binding the contract method 0xffa66f28.
//
// Solidity: function SettleOpsReward(uint64[] slots) returns()
func (_History *HistoryTransactor) SettleOpsReward(opts *bind.TransactOpts, slots []uint64) (*types.Transaction, error) {
	return _History.contract.Transact(opts, "SettleOpsReward", slots)
}

// SettleOpsReward is a paid mutator transaction binding the contract method 0xffa66f28.
//
// Solidity: function SettleOpsReward(uint64[] slots) returns()
func (_History *HistorySession) SettleOpsReward(slots []uint64) (*types.Transaction, error) {
	return _History.Contract.SettleOpsReward(&_History.TransactOpts, slots)
}

// SettleOpsReward is a paid mutator transaction binding the contract method 0xffa66f28.
//
// Solidity: function SettleOpsReward(uint64[] slots) returns()
func (_History *HistoryTransactorSession) SettleOpsReward(slots []uint64) (*types.Transaction, error) {
	return _History.Contract.SettleOpsReward(&_History.TransactOpts, slots)
}

// SettleVoteReward is a paid mutator transaction binding the contract method 0x8ac43ff2.
//
// Solidity: function SettleVoteReward(uint64[] daySlots) returns()
func (_History *HistoryTransactor) SettleVoteReward(opts *bind.TransactOpts, daySlots []uint64) (*types.Transaction, error) {
	return _History.contract.Transact(opts, "SettleVoteReward", daySlots)
}

// SettleVoteReward is a paid mutator transaction binding the contract method 0x8ac43ff2.
//
// Solidity: function SettleVoteReward(uint64[] daySlots) returns()
func (_History *HistorySession) SettleVoteReward(daySlots []uint64) (*types.Transaction, error) {
	return _History.Contract.SettleVoteReward(&_History.TransactOpts, daySlots)
}

// SettleVoteReward is a paid mutator transaction binding the contract method 0x8ac43ff2.
//
// Solidity: function SettleVoteReward(uint64[] daySlots) returns()
func (_History *HistoryTransactorSession) SettleVoteReward(daySlots []uint64) (*types.Transaction, error) {
	return _History.Contract.SettleVoteReward(&_History.TransactOpts, daySlots)
}

// SubmitCandidates is a paid mutator transaction binding the contract method 0xc878d1c5.
//
// Solidity: function SubmitCandidates(uint64 daySlot, string[] candidates) returns()
func (_History *HistoryTransactor) SubmitCandidates(opts *bind.TransactOpts, daySlot uint64, candidates []string) (*types.Transaction, error) {
	return _History.contract.Transact(opts, "SubmitCandidates", daySlot, candidates)
}

// SubmitCandidates is a paid mutator transaction binding the contract method 0xc878d1c5.
//
// Solidity: function SubmitCandidates(uint64 daySlot, string[] candidates) returns()
func (_History *HistorySession) SubmitCandidates(daySlot uint64, candidates []string) (*types.Transaction, error) {
	return _History.Contract.SubmitCandidates(&_History.TransactOpts, daySlot, candidates)
}

// SubmitCandidates is a paid mutator transaction binding the contract method 0xc878d1c5.
//
// Solidity: function SubmitCandidates(uint64 daySlot, string[] candidates) returns()
func (_History *HistoryTransactorSession) SubmitCandidates(daySlot uint64, candidates []string) (*types.Transaction, error) {
	return _History.Contract.SubmitCandidates(&_History.TransactOpts, daySlot, candidates)
}

// Vote is a paid mutator transaction binding the contract method 0x1ea2d084.
//
// Solidity: function Vote(uint64 daySlot, uint8[3] prefers) returns()
func (_History *HistoryTransactor) Vote(opts *bind.TransactOpts, daySlot uint64, prefers [3]uint8) (*types.Transaction, error) {
	return _History.contract.Transact(opts, "Vote", daySlot, prefers)
}

// Vote is a paid mutator transaction binding the contract method 0x1ea2d084.
//
// Solidity: function Vote(uint64 daySlot, uint8[3] prefers) returns()
func (_History *HistorySession) Vote(daySlot uint64, prefers [3]uint8) (*types.Transaction, error) {
	return _History.Contract.Vote(&_History.TransactOpts, daySlot, prefers)
}

// Vote is a paid mutator transaction binding the contract method 0x1ea2d084.
//
// Solidity: function Vote(uint64 daySlot, uint8[3] prefers) returns()
func (_History *HistoryTransactorSession) Vote(daySlot uint64, prefers [3]uint8) (*types.Transaction, error) {
	return _History.Contract.Vote(&_History.TransactOpts, daySlot, prefers)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_History *HistoryTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _History.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_History *HistorySession) Initialize() (*types.Transaction, error) {
	return _History.Contract.Initialize(&_History.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_History *HistoryTransactorSession) Initialize() (*types.Transaction, error) {
	return _History.Contract.Initialize(&_History.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_History *HistoryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _History.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_History *HistorySession) Pause() (*types.Transaction, error) {
	return _History.Contract.Pause(&_History.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_History *HistoryTransactorSession) Pause() (*types.Transaction, error) {
	return _History.Contract.Pause(&_History.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_History *HistoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _History.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_History *HistorySession) RenounceOwnership() (*types.Transaction, error) {
	return _History.Contract.RenounceOwnership(&_History.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_History *HistoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _History.Contract.RenounceOwnership(&_History.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_History *HistoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _History.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_History *HistorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _History.Contract.TransferOwnership(&_History.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_History *HistoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _History.Contract.TransferOwnership(&_History.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_History *HistoryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _History.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_History *HistorySession) Unpause() (*types.Transaction, error) {
	return _History.Contract.Unpause(&_History.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_History *HistoryTransactorSession) Unpause() (*types.Transaction, error) {
	return _History.Contract.Unpause(&_History.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_History *HistoryTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _History.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_History *HistorySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _History.Contract.UpgradeTo(&_History.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_History *HistoryTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _History.Contract.UpgradeTo(&_History.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_History *HistoryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _History.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_History *HistorySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _History.Contract.UpgradeToAndCall(&_History.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_History *HistoryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _History.Contract.UpgradeToAndCall(&_History.TransactOpts, newImplementation, data)
}

// HistoryAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the History contract.
type HistoryAdminChangedIterator struct {
	Event *HistoryAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HistoryAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HistoryAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HistoryAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HistoryAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HistoryAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HistoryAdminChanged represents a AdminChanged event raised by the History contract.
type HistoryAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_History *HistoryFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*HistoryAdminChangedIterator, error) {

	logs, sub, err := _History.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &HistoryAdminChangedIterator{contract: _History.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_History *HistoryFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *HistoryAdminChanged) (event.Subscription, error) {

	logs, sub, err := _History.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HistoryAdminChanged)
				if err := _History.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_History *HistoryFilterer) ParseAdminChanged(log types.Log) (*HistoryAdminChanged, error) {
	event := new(HistoryAdminChanged)
	if err := _History.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HistoryBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the History contract.
type HistoryBeaconUpgradedIterator struct {
	Event *HistoryBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HistoryBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HistoryBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HistoryBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HistoryBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HistoryBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HistoryBeaconUpgraded represents a BeaconUpgraded event raised by the History contract.
type HistoryBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_History *HistoryFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*HistoryBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _History.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &HistoryBeaconUpgradedIterator{contract: _History.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_History *HistoryFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *HistoryBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _History.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HistoryBeaconUpgraded)
				if err := _History.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_History *HistoryFilterer) ParseBeaconUpgraded(log types.Log) (*HistoryBeaconUpgraded, error) {
	event := new(HistoryBeaconUpgraded)
	if err := _History.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HistoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the History contract.
type HistoryInitializedIterator struct {
	Event *HistoryInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HistoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HistoryInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HistoryInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HistoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HistoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HistoryInitialized represents a Initialized event raised by the History contract.
type HistoryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_History *HistoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*HistoryInitializedIterator, error) {

	logs, sub, err := _History.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &HistoryInitializedIterator{contract: _History.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_History *HistoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *HistoryInitialized) (event.Subscription, error) {

	logs, sub, err := _History.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HistoryInitialized)
				if err := _History.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_History *HistoryFilterer) ParseInitialized(log types.Log) (*HistoryInitialized, error) {
	event := new(HistoryInitialized)
	if err := _History.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HistoryNewCardIterator is returned from FilterNewCard and is used to iterate over the raw logs and unpacked data for NewCard events raised by the History contract.
type HistoryNewCardIterator struct {
	Event *HistoryNewCard // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HistoryNewCardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HistoryNewCard)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HistoryNewCard)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HistoryNewCardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HistoryNewCardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HistoryNewCard represents a NewCard event raised by the History contract.
type HistoryNewCard struct {
	DaySlot uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewCard is a free log retrieval operation binding the contract event 0xe1b8fb0326a90ea1dc6b90625c26b9e0e9f172bc0cd1c2d6d2900d4e5707979c.
//
// Solidity: event NewCard(uint64 daySlot)
func (_History *HistoryFilterer) FilterNewCard(opts *bind.FilterOpts) (*HistoryNewCardIterator, error) {

	logs, sub, err := _History.contract.FilterLogs(opts, "NewCard")
	if err != nil {
		return nil, err
	}
	return &HistoryNewCardIterator{contract: _History.contract, event: "NewCard", logs: logs, sub: sub}, nil
}

// WatchNewCard is a free log subscription operation binding the contract event 0xe1b8fb0326a90ea1dc6b90625c26b9e0e9f172bc0cd1c2d6d2900d4e5707979c.
//
// Solidity: event NewCard(uint64 daySlot)
func (_History *HistoryFilterer) WatchNewCard(opts *bind.WatchOpts, sink chan<- *HistoryNewCard) (event.Subscription, error) {

	logs, sub, err := _History.contract.WatchLogs(opts, "NewCard")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HistoryNewCard)
				if err := _History.contract.UnpackLog(event, "NewCard", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewCard is a log parse operation binding the contract event 0xe1b8fb0326a90ea1dc6b90625c26b9e0e9f172bc0cd1c2d6d2900d4e5707979c.
//
// Solidity: event NewCard(uint64 daySlot)
func (_History *HistoryFilterer) ParseNewCard(log types.Log) (*HistoryNewCard, error) {
	event := new(HistoryNewCard)
	if err := _History.contract.UnpackLog(event, "NewCard", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HistoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the History contract.
type HistoryOwnershipTransferredIterator struct {
	Event *HistoryOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HistoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HistoryOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HistoryOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HistoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HistoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HistoryOwnershipTransferred represents a OwnershipTransferred event raised by the History contract.
type HistoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_History *HistoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HistoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _History.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HistoryOwnershipTransferredIterator{contract: _History.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_History *HistoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HistoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _History.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HistoryOwnershipTransferred)
				if err := _History.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_History *HistoryFilterer) ParseOwnershipTransferred(log types.Log) (*HistoryOwnershipTransferred, error) {
	event := new(HistoryOwnershipTransferred)
	if err := _History.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HistoryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the History contract.
type HistoryPausedIterator struct {
	Event *HistoryPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HistoryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HistoryPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HistoryPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HistoryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HistoryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HistoryPaused represents a Paused event raised by the History contract.
type HistoryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_History *HistoryFilterer) FilterPaused(opts *bind.FilterOpts) (*HistoryPausedIterator, error) {

	logs, sub, err := _History.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &HistoryPausedIterator{contract: _History.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_History *HistoryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *HistoryPaused) (event.Subscription, error) {

	logs, sub, err := _History.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HistoryPaused)
				if err := _History.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_History *HistoryFilterer) ParsePaused(log types.Log) (*HistoryPaused, error) {
	event := new(HistoryPaused)
	if err := _History.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HistoryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the History contract.
type HistoryUnpausedIterator struct {
	Event *HistoryUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HistoryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HistoryUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HistoryUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HistoryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HistoryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HistoryUnpaused represents a Unpaused event raised by the History contract.
type HistoryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_History *HistoryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*HistoryUnpausedIterator, error) {

	logs, sub, err := _History.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &HistoryUnpausedIterator{contract: _History.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_History *HistoryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *HistoryUnpaused) (event.Subscription, error) {

	logs, sub, err := _History.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HistoryUnpaused)
				if err := _History.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_History *HistoryFilterer) ParseUnpaused(log types.Log) (*HistoryUnpaused, error) {
	event := new(HistoryUnpaused)
	if err := _History.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HistoryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the History contract.
type HistoryUpgradedIterator struct {
	Event *HistoryUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HistoryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HistoryUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HistoryUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HistoryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HistoryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HistoryUpgraded represents a Upgraded event raised by the History contract.
type HistoryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_History *HistoryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*HistoryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _History.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &HistoryUpgradedIterator{contract: _History.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_History *HistoryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *HistoryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _History.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HistoryUpgraded)
				if err := _History.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_History *HistoryFilterer) ParseUpgraded(log types.Log) (*HistoryUpgraded, error) {
	event := new(HistoryUpgraded)
	if err := _History.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
