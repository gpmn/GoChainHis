// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Escrow

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

// EscrowMetaData contains all meta data concerning the Escrow contract.
var EscrowMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depAmt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolAmt\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolAmt\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"OwnerAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TotalDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositTm\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositAmt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toTm\",\"type\":\"uint256\"}],\"name\":\"calcVotesTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"voteAmt\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toTm\",\"type\":\"uint256\"}],\"name\":\"getDepositInfoTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositTm\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteAmt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getDepositInfoTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositTm\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteAmt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getVoteAmt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"voteAmt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// EscrowABI is the input ABI used to generate the binding from.
// Deprecated: Use EscrowMetaData.ABI instead.
var EscrowABI = EscrowMetaData.ABI

// Escrow is an auto generated Go binding around an Ethereum contract.
type Escrow struct {
	EscrowCaller     // Read-only binding to the contract
	EscrowTransactor // Write-only binding to the contract
	EscrowFilterer   // Log filterer for contract events
}

// EscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type EscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EscrowSession struct {
	Contract     *Escrow           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EscrowCallerSession struct {
	Contract *EscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EscrowTransactorSession struct {
	Contract     *EscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type EscrowRaw struct {
	Contract *Escrow // Generic contract binding to access the raw methods on
}

// EscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EscrowCallerRaw struct {
	Contract *EscrowCaller // Generic read-only contract binding to access the raw methods on
}

// EscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EscrowTransactorRaw struct {
	Contract *EscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEscrow creates a new instance of Escrow, bound to a specific deployed contract.
func NewEscrow(address common.Address, backend bind.ContractBackend) (*Escrow, error) {
	contract, err := bindEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Escrow{EscrowCaller: EscrowCaller{contract: contract}, EscrowTransactor: EscrowTransactor{contract: contract}, EscrowFilterer: EscrowFilterer{contract: contract}}, nil
}

// NewEscrowCaller creates a new read-only instance of Escrow, bound to a specific deployed contract.
func NewEscrowCaller(address common.Address, caller bind.ContractCaller) (*EscrowCaller, error) {
	contract, err := bindEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EscrowCaller{contract: contract}, nil
}

// NewEscrowTransactor creates a new write-only instance of Escrow, bound to a specific deployed contract.
func NewEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*EscrowTransactor, error) {
	contract, err := bindEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EscrowTransactor{contract: contract}, nil
}

// NewEscrowFilterer creates a new log filterer instance of Escrow, bound to a specific deployed contract.
func NewEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*EscrowFilterer, error) {
	contract, err := bindEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EscrowFilterer{contract: contract}, nil
}

// bindEscrow binds a generic wrapper to an already deployed contract.
func bindEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EscrowABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Escrow *EscrowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Escrow.Contract.EscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Escrow *EscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrow.Contract.EscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Escrow *EscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Escrow.Contract.EscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Escrow *EscrowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Escrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Escrow *EscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Escrow *EscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Escrow.Contract.contract.Transact(opts, method, params...)
}

// OwnerAddr is a free data retrieval call binding the contract method 0xc1d39fc1.
//
// Solidity: function OwnerAddr() view returns(address)
func (_Escrow *EscrowCaller) OwnerAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "OwnerAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerAddr is a free data retrieval call binding the contract method 0xc1d39fc1.
//
// Solidity: function OwnerAddr() view returns(address)
func (_Escrow *EscrowSession) OwnerAddr() (common.Address, error) {
	return _Escrow.Contract.OwnerAddr(&_Escrow.CallOpts)
}

// OwnerAddr is a free data retrieval call binding the contract method 0xc1d39fc1.
//
// Solidity: function OwnerAddr() view returns(address)
func (_Escrow *EscrowCallerSession) OwnerAddr() (common.Address, error) {
	return _Escrow.Contract.OwnerAddr(&_Escrow.CallOpts)
}

// TotalDeposits is a free data retrieval call binding the contract method 0xef110f49.
//
// Solidity: function TotalDeposits() view returns(uint256)
func (_Escrow *EscrowCaller) TotalDeposits(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "TotalDeposits")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDeposits is a free data retrieval call binding the contract method 0xef110f49.
//
// Solidity: function TotalDeposits() view returns(uint256)
func (_Escrow *EscrowSession) TotalDeposits() (*big.Int, error) {
	return _Escrow.Contract.TotalDeposits(&_Escrow.CallOpts)
}

// TotalDeposits is a free data retrieval call binding the contract method 0xef110f49.
//
// Solidity: function TotalDeposits() view returns(uint256)
func (_Escrow *EscrowCallerSession) TotalDeposits() (*big.Int, error) {
	return _Escrow.Contract.TotalDeposits(&_Escrow.CallOpts)
}

// CalcVotesTo is a free data retrieval call binding the contract method 0x2fe5d4a2.
//
// Solidity: function calcVotesTo(uint256 depositTm, uint256 depositAmt, uint256 toTm) pure returns(uint256 voteAmt)
func (_Escrow *EscrowCaller) CalcVotesTo(opts *bind.CallOpts, depositTm *big.Int, depositAmt *big.Int, toTm *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "calcVotesTo", depositTm, depositAmt, toTm)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcVotesTo is a free data retrieval call binding the contract method 0x2fe5d4a2.
//
// Solidity: function calcVotesTo(uint256 depositTm, uint256 depositAmt, uint256 toTm) pure returns(uint256 voteAmt)
func (_Escrow *EscrowSession) CalcVotesTo(depositTm *big.Int, depositAmt *big.Int, toTm *big.Int) (*big.Int, error) {
	return _Escrow.Contract.CalcVotesTo(&_Escrow.CallOpts, depositTm, depositAmt, toTm)
}

// CalcVotesTo is a free data retrieval call binding the contract method 0x2fe5d4a2.
//
// Solidity: function calcVotesTo(uint256 depositTm, uint256 depositAmt, uint256 toTm) pure returns(uint256 voteAmt)
func (_Escrow *EscrowCallerSession) CalcVotesTo(depositTm *big.Int, depositAmt *big.Int, toTm *big.Int) (*big.Int, error) {
	return _Escrow.Contract.CalcVotesTo(&_Escrow.CallOpts, depositTm, depositAmt, toTm)
}

// GetDepositInfoTo is a free data retrieval call binding the contract method 0x58345aa7.
//
// Solidity: function getDepositInfoTo(address addr, uint256 toTm) view returns(uint256 amount, uint256 depositTm, uint256 voteAmt)
func (_Escrow *EscrowCaller) GetDepositInfoTo(opts *bind.CallOpts, addr common.Address, toTm *big.Int) (struct {
	Amount    *big.Int
	DepositTm *big.Int
	VoteAmt   *big.Int
}, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "getDepositInfoTo", addr, toTm)

	outstruct := new(struct {
		Amount    *big.Int
		DepositTm *big.Int
		VoteAmt   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DepositTm = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.VoteAmt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetDepositInfoTo is a free data retrieval call binding the contract method 0x58345aa7.
//
// Solidity: function getDepositInfoTo(address addr, uint256 toTm) view returns(uint256 amount, uint256 depositTm, uint256 voteAmt)
func (_Escrow *EscrowSession) GetDepositInfoTo(addr common.Address, toTm *big.Int) (struct {
	Amount    *big.Int
	DepositTm *big.Int
	VoteAmt   *big.Int
}, error) {
	return _Escrow.Contract.GetDepositInfoTo(&_Escrow.CallOpts, addr, toTm)
}

// GetDepositInfoTo is a free data retrieval call binding the contract method 0x58345aa7.
//
// Solidity: function getDepositInfoTo(address addr, uint256 toTm) view returns(uint256 amount, uint256 depositTm, uint256 voteAmt)
func (_Escrow *EscrowCallerSession) GetDepositInfoTo(addr common.Address, toTm *big.Int) (struct {
	Amount    *big.Int
	DepositTm *big.Int
	VoteAmt   *big.Int
}, error) {
	return _Escrow.Contract.GetDepositInfoTo(&_Escrow.CallOpts, addr, toTm)
}

// GetDepositInfoTo0 is a free data retrieval call binding the contract method 0xdf49663f.
//
// Solidity: function getDepositInfoTo(address addr) view returns(uint256 amount, uint256 depositTm, uint256 voteAmt)
func (_Escrow *EscrowCaller) GetDepositInfoTo0(opts *bind.CallOpts, addr common.Address) (struct {
	Amount    *big.Int
	DepositTm *big.Int
	VoteAmt   *big.Int
}, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "getDepositInfoTo0", addr)

	outstruct := new(struct {
		Amount    *big.Int
		DepositTm *big.Int
		VoteAmt   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DepositTm = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.VoteAmt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetDepositInfoTo0 is a free data retrieval call binding the contract method 0xdf49663f.
//
// Solidity: function getDepositInfoTo(address addr) view returns(uint256 amount, uint256 depositTm, uint256 voteAmt)
func (_Escrow *EscrowSession) GetDepositInfoTo0(addr common.Address) (struct {
	Amount    *big.Int
	DepositTm *big.Int
	VoteAmt   *big.Int
}, error) {
	return _Escrow.Contract.GetDepositInfoTo0(&_Escrow.CallOpts, addr)
}

// GetDepositInfoTo0 is a free data retrieval call binding the contract method 0xdf49663f.
//
// Solidity: function getDepositInfoTo(address addr) view returns(uint256 amount, uint256 depositTm, uint256 voteAmt)
func (_Escrow *EscrowCallerSession) GetDepositInfoTo0(addr common.Address) (struct {
	Amount    *big.Int
	DepositTm *big.Int
	VoteAmt   *big.Int
}, error) {
	return _Escrow.Contract.GetDepositInfoTo0(&_Escrow.CallOpts, addr)
}

// GetVoteAmt is a free data retrieval call binding the contract method 0x515ff6d2.
//
// Solidity: function getVoteAmt(address addr) view returns(uint256 voteAmt)
func (_Escrow *EscrowCaller) GetVoteAmt(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "getVoteAmt", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVoteAmt is a free data retrieval call binding the contract method 0x515ff6d2.
//
// Solidity: function getVoteAmt(address addr) view returns(uint256 voteAmt)
func (_Escrow *EscrowSession) GetVoteAmt(addr common.Address) (*big.Int, error) {
	return _Escrow.Contract.GetVoteAmt(&_Escrow.CallOpts, addr)
}

// GetVoteAmt is a free data retrieval call binding the contract method 0x515ff6d2.
//
// Solidity: function getVoteAmt(address addr) view returns(uint256 voteAmt)
func (_Escrow *EscrowCallerSession) GetVoteAmt(addr common.Address) (*big.Int, error) {
	return _Escrow.Contract.GetVoteAmt(&_Escrow.CallOpts, addr)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Escrow *EscrowTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Escrow *EscrowSession) Deposit() (*types.Transaction, error) {
	return _Escrow.Contract.Deposit(&_Escrow.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Escrow *EscrowTransactorSession) Deposit() (*types.Transaction, error) {
	return _Escrow.Contract.Deposit(&_Escrow.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amt) returns()
func (_Escrow *EscrowTransactor) Withdraw(opts *bind.TransactOpts, amt *big.Int) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "withdraw", amt)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amt) returns()
func (_Escrow *EscrowSession) Withdraw(amt *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.Withdraw(&_Escrow.TransactOpts, amt)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amt) returns()
func (_Escrow *EscrowTransactorSession) Withdraw(amt *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.Withdraw(&_Escrow.TransactOpts, amt)
}

// EscrowDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Escrow contract.
type EscrowDepositedIterator struct {
	Event *EscrowDeposited // Event containing the contract specifics and raw log

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
func (it *EscrowDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowDeposited)
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
		it.Event = new(EscrowDeposited)
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
func (it *EscrowDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowDeposited represents a Deposited event raised by the Escrow contract.
type EscrowDeposited struct {
	Addr    common.Address
	DepAmt  *big.Int
	PoolAmt *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed addr, uint256 depAmt, uint256 poolAmt)
func (_Escrow *EscrowFilterer) FilterDeposited(opts *bind.FilterOpts, addr []common.Address) (*EscrowDepositedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "Deposited", addrRule)
	if err != nil {
		return nil, err
	}
	return &EscrowDepositedIterator{contract: _Escrow.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed addr, uint256 depAmt, uint256 poolAmt)
func (_Escrow *EscrowFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *EscrowDeposited, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "Deposited", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowDeposited)
				if err := _Escrow.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed addr, uint256 depAmt, uint256 poolAmt)
func (_Escrow *EscrowFilterer) ParseDeposited(log types.Log) (*EscrowDeposited, error) {
	event := new(EscrowDeposited)
	if err := _Escrow.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EscrowWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Escrow contract.
type EscrowWithdrawnIterator struct {
	Event *EscrowWithdrawn // Event containing the contract specifics and raw log

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
func (it *EscrowWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowWithdrawn)
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
		it.Event = new(EscrowWithdrawn)
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
func (it *EscrowWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowWithdrawn represents a Withdrawn event raised by the Escrow contract.
type EscrowWithdrawn struct {
	Addr        common.Address
	WithdrawAmt *big.Int
	PoolAmt     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed addr, uint256 withdrawAmt, uint256 poolAmt)
func (_Escrow *EscrowFilterer) FilterWithdrawn(opts *bind.FilterOpts, addr []common.Address) (*EscrowWithdrawnIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "Withdrawn", addrRule)
	if err != nil {
		return nil, err
	}
	return &EscrowWithdrawnIterator{contract: _Escrow.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed addr, uint256 withdrawAmt, uint256 poolAmt)
func (_Escrow *EscrowFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *EscrowWithdrawn, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "Withdrawn", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowWithdrawn)
				if err := _Escrow.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed addr, uint256 withdrawAmt, uint256 poolAmt)
func (_Escrow *EscrowFilterer) ParseWithdrawn(log types.Log) (*EscrowWithdrawn, error) {
	event := new(EscrowWithdrawn)
	if err := _Escrow.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
