// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Card

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

// CardInfo is an auto generated low-level Go binding around an user-defined struct.
type CardInfo struct {
	BigStoriesIdx [3]uint8
	RenderOpt     uint8
	ClaimSec      uint64
	Greeting      string
	GreetingImg   string
	Festivals     []string
}

// CardMetaData contains all meta data concerning the Card contract.
var CardMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint8[3]\",\"name\":\"BigStoriesIdx\",\"type\":\"uint8[3]\"},{\"internalType\":\"uint8\",\"name\":\"RenderOpt\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"ClaimSec\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"Greeting\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"GreetingImg\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"Festivals\",\"type\":\"string[]\"}],\"internalType\":\"structCardInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"name\":\"AirDrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BaseImg\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"BaseImgMap\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"CardInfoMap\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"RenderOpt\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"ClaimSec\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"Greeting\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"GreetingImg\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"}],\"name\":\"Claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"renderOpt\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"greeting\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"greetingImg\",\"type\":\"string\"}],\"name\":\"Customize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Echo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"tokenId\",\"type\":\"uint64\"}],\"name\":\"Exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FirstAirDropFreeId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FirstTokenId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HistoryAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LastAirDropId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint8[3]\",\"name\":\"BigStoriesIdx\",\"type\":\"uint8[3]\"},{\"internalType\":\"uint8\",\"name\":\"RenderOpt\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"ClaimSec\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"Greeting\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"GreetingImg\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"Festivals\",\"type\":\"string[]\"}],\"internalType\":\"structCardInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"name\":\"Mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseImg\",\"type\":\"string\"}],\"name\":\"SetBaseImg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"baseImg\",\"type\":\"string\"}],\"name\":\"SetBaseImg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"historyAddr\",\"type\":\"address\"}],\"name\":\"SetHistoryAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// CardABI is the input ABI used to generate the binding from.
// Deprecated: Use CardMetaData.ABI instead.
var CardABI = CardMetaData.ABI

// Card is an auto generated Go binding around an Ethereum contract.
type Card struct {
	CardCaller     // Read-only binding to the contract
	CardTransactor // Write-only binding to the contract
	CardFilterer   // Log filterer for contract events
}

// CardCaller is an auto generated read-only Go binding around an Ethereum contract.
type CardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CardSession struct {
	Contract     *Card             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CardCallerSession struct {
	Contract *CardCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CardTransactorSession struct {
	Contract     *CardTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CardRaw is an auto generated low-level Go binding around an Ethereum contract.
type CardRaw struct {
	Contract *Card // Generic contract binding to access the raw methods on
}

// CardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CardCallerRaw struct {
	Contract *CardCaller // Generic read-only contract binding to access the raw methods on
}

// CardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CardTransactorRaw struct {
	Contract *CardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCard creates a new instance of Card, bound to a specific deployed contract.
func NewCard(address common.Address, backend bind.ContractBackend) (*Card, error) {
	contract, err := bindCard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Card{CardCaller: CardCaller{contract: contract}, CardTransactor: CardTransactor{contract: contract}, CardFilterer: CardFilterer{contract: contract}}, nil
}

// NewCardCaller creates a new read-only instance of Card, bound to a specific deployed contract.
func NewCardCaller(address common.Address, caller bind.ContractCaller) (*CardCaller, error) {
	contract, err := bindCard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CardCaller{contract: contract}, nil
}

// NewCardTransactor creates a new write-only instance of Card, bound to a specific deployed contract.
func NewCardTransactor(address common.Address, transactor bind.ContractTransactor) (*CardTransactor, error) {
	contract, err := bindCard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CardTransactor{contract: contract}, nil
}

// NewCardFilterer creates a new log filterer instance of Card, bound to a specific deployed contract.
func NewCardFilterer(address common.Address, filterer bind.ContractFilterer) (*CardFilterer, error) {
	contract, err := bindCard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CardFilterer{contract: contract}, nil
}

// bindCard binds a generic wrapper to an already deployed contract.
func bindCard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Card *CardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Card.Contract.CardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Card *CardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Card.Contract.CardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Card *CardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Card.Contract.CardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Card *CardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Card.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Card *CardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Card.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Card *CardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Card.Contract.contract.Transact(opts, method, params...)
}

// BaseImg is a free data retrieval call binding the contract method 0x73b3e22c.
//
// Solidity: function BaseImg() view returns(string)
func (_Card *CardCaller) BaseImg(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Card.contract.Call(opts, &out, "BaseImg")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseImg is a free data retrieval call binding the contract method 0x73b3e22c.
//
// Solidity: function BaseImg() view returns(string)
func (_Card *CardSession) BaseImg() (string, error) {
	return _Card.Contract.BaseImg(&_Card.CallOpts)
}

// BaseImg is a free data retrieval call binding the contract method 0x73b3e22c.
//
// Solidity: function BaseImg() view returns(string)
func (_Card *CardCallerSession) BaseImg() (string, error) {
	return _Card.Contract.BaseImg(&_Card.CallOpts)
}

// BaseImgMap is a free data retrieval call binding the contract method 0x4f3eb037.
//
// Solidity: function BaseImgMap(uint64 ) view returns(string)
func (_Card *CardCaller) BaseImgMap(opts *bind.CallOpts, arg0 uint64) (string, error) {
	var out []interface{}
	err := _Card.contract.Call(opts, &out, "BaseImgMap", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseImgMap is a free data retrieval call binding the contract method 0x4f3eb037.
//
// Solidity: function BaseImgMap(uint64 ) view returns(string)
func (_Card *CardSession) BaseImgMap(arg0 uint64) (string, error) {
	return _Card.Contract.BaseImgMap(&_Card.CallOpts, arg0)
}

// BaseImgMap is a free data retrieval call binding the contract method 0x4f3eb037.
//
// Solidity: function BaseImgMap(uint64 ) view returns(string)
func (_Card *CardCallerSession) BaseImgMap(arg0 uint64) (string, error) {
	return _Card.Contract.BaseImgMap(&_Card.CallOpts, arg0)
}

// CardInfoMap is a free data retrieval call binding the contract method 0x2d536432.
//
// Solidity: function CardInfoMap(uint64 ) view returns(uint8 RenderOpt, uint64 ClaimSec, string Greeting, string GreetingImg)
func (_Card *CardCaller) CardInfoMap(opts *bind.CallOpts, arg0 uint64) (struct {
	RenderOpt   uint8
	ClaimSec    uint64
	Greeting    string
	GreetingImg string
}, error) {
	var out []interface{}
	err := _Card.contract.Call(opts, &out, "CardInfoMap", arg0)

	outstruct := new(struct {
		RenderOpt   uint8
		ClaimSec    uint64
		Greeting    string
		GreetingImg string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RenderOpt = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.ClaimSec = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Greeting = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.GreetingImg = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// CardInfoMap is a free data retrieval call binding the contract method 0x2d536432.
//
// Solidity: function CardInfoMap(uint64 ) view returns(uint8 RenderOpt, uint64 ClaimSec, string Greeting, string GreetingImg)
func (_Card *CardSession) CardInfoMap(arg0 uint64) (struct {
	RenderOpt   uint8
	ClaimSec    uint64
	Greeting    string
	GreetingImg string
}, error) {
	return _Card.Contract.CardInfoMap(&_Card.CallOpts, arg0)
}

// CardInfoMap is a free data retrieval call binding the contract method 0x2d536432.
//
// Solidity: function CardInfoMap(uint64 ) view returns(uint8 RenderOpt, uint64 ClaimSec, string Greeting, string GreetingImg)
func (_Card *CardCallerSession) CardInfoMap(arg0 uint64) (struct {
	RenderOpt   uint8
	ClaimSec    uint64
	Greeting    string
	GreetingImg string
}, error) {
	return _Card.Contract.CardInfoMap(&_Card.CallOpts, arg0)
}

// Echo is a free data retrieval call binding the contract method 0x24c4b0ef.
//
// Solidity: function Echo() pure returns(string)
func (_Card *CardCaller) Echo(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Card.contract.Call(opts, &out, "Echo")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Echo is a free data retrieval call binding the contract method 0x24c4b0ef.
//
// Solidity: function Echo() pure returns(string)
func (_Card *CardSession) Echo() (string, error) {
	return _Card.Contract.Echo(&_Card.CallOpts)
}

// Echo is a free data retrieval call binding the contract method 0x24c4b0ef.
//
// Solidity: function Echo() pure returns(string)
func (_Card *CardCallerSession) Echo() (string, error) {
	return _Card.Contract.Echo(&_Card.CallOpts)
}

// Exists is a free data retrieval call binding the contract method 0x0876c881.
//
// Solidity: function Exists(uint64 tokenId) view returns(bool)
func (_Card *CardCaller) Exists(opts *bind.CallOpts, tokenId uint64) (bool, error) {
	var out []interface{}
	err := _Card.contract.Call(opts, &out, "Exists", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x0876c881.
//
// Solidity: function Exists(uint64 tokenId) view returns(bool)
func (_Card *CardSession) Exists(tokenId uint64) (bool, error) {
	return _Card.Contract.Exists(&_Card.CallOpts, tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x0876c881.
//
// Solidity: function Exists(uint64 tokenId) view returns(bool)
func (_Card *CardCallerSession) Exists(tokenId uint64) (bool, error) {
	return _Card.Contract.Exists(&_Card.CallOpts, tokenId)
}

// FirstAirDropFreeId is a free data retrieval call binding the contract method 0xd69b0a3a.
//
// Solidity: function FirstAirDropFreeId() view returns(uint64)
func (_Card *CardCaller) FirstAirDropFreeId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Card.contract.Call(opts, &out, "FirstAirDropFreeId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// FirstAirDropFreeId is a free data retrieval call binding the contract method 0xd69b0a3a.
//
// Solidity: function FirstAirDropFreeId() view returns(uint64)
func (_Card *CardSession) FirstAirDropFreeId() (uint64, error) {
	return _Card.Contract.FirstAirDropFreeId(&_Card.CallOpts)
}

// FirstAirDropFreeId is a free data retrieval call binding the contract method 0xd69b0a3a.
//
// Solidity: function FirstAirDropFreeId() view returns(uint64)
func (_Card *CardCallerSession) FirstAirDropFreeId() (uint64, error) {
	return _Card.Contract.FirstAirDropFreeId(&_Card.CallOpts)
}

// FirstTokenId is a free data retrieval call binding the contract method 0xcd087590.
//
// Solidity: function FirstTokenId() view returns(uint64)
func (_Card *CardCaller) FirstTokenId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Card.contract.Call(opts, &out, "FirstTokenId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// FirstTokenId is a free data retrieval call binding the contract method 0xcd087590.
//
// Solidity: function FirstTokenId() view returns(uint64)
func (_Card *CardSession) FirstTokenId() (uint64, error) {
	return _Card.Contract.FirstTokenId(&_Card.CallOpts)
}

// FirstTokenId is a free data retrieval call binding the contract method 0xcd087590.
//
// Solidity: function FirstTokenId() view returns(uint64)
func (_Card *CardCallerSession) FirstTokenId() (uint64, error) {
	return _Card.Contract.FirstTokenId(&_Card.CallOpts)
}

// HistoryAddr is a free data retrieval call binding the contract method 0xbe0cd318.
//
// Solidity: function HistoryAddr() view returns(address)
func (_Card *CardCaller) HistoryAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Card.contract.Call(opts, &out, "HistoryAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HistoryAddr is a free data retrieval call binding the contract method 0xbe0cd318.
//
// Solidity: function HistoryAddr() view returns(address)
func (_Card *CardSession) HistoryAddr() (common.Address, error) {
	return _Card.Contract.HistoryAddr(&_Card.CallOpts)
}

// HistoryAddr is a free data retrieval call binding the contract method 0xbe0cd318.
//
// Solidity: function HistoryAddr() view returns(address)
func (_Card *CardCallerSession) HistoryAddr() (common.Address, error) {
	return _Card.Contract.HistoryAddr(&_Card.CallOpts)
}

// LastAirDropId is a free data retrieval call binding the contract method 0xfb5927c6.
//
// Solidity: function LastAirDropId() view returns(uint64)
func (_Card *CardCaller) LastAirDropId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Card.contract.Call(opts, &out, "LastAirDropId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastAirDropId is a free data retrieval call binding the contract method 0xfb5927c6.
//
// Solidity: function LastAirDropId() view returns(uint64)
func (_Card *CardSession) LastAirDropId() (uint64, error) {
	return _Card.Contract.LastAirDropId(&_Card.CallOpts)
}

// LastAirDropId is a free data retrieval call binding the contract method 0xfb5927c6.
//
// Solidity: function LastAirDropId() view returns(uint64)
func (_Card *CardCallerSession) LastAirDropId() (uint64, error) {
	return _Card.Contract.LastAirDropId(&_Card.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Card *CardCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Card.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Card *CardSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Card.Contract.BalanceOf(&_Card.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Card *CardCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Card.Contract.BalanceOf(&_Card.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Card *CardCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Card.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Card *CardSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Card.Contract.GetApproved(&_Card.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Card *CardCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Card.Contract.GetApproved(&_Card.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Card *CardCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Card.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Card *CardSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Card.Contract.IsApprovedForAll(&_Card.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Card *CardCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Card.Contract.IsApprovedForAll(&_Card.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Card *CardCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Card.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Card *CardSession) Name() (string, error) {
	return _Card.Contract.Name(&_Card.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Card *CardCallerSession) Name() (string, error) {
	return _Card.Contract.Name(&_Card.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Card *CardCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Card.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Card *CardSession) Owner() (common.Address, error) {
	return _Card.Contract.Owner(&_Card.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Card *CardCallerSession) Owner() (common.Address, error) {
	return _Card.Contract.Owner(&_Card.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Card *CardCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Card.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Card *CardSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Card.Contract.OwnerOf(&_Card.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Card *CardCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Card.Contract.OwnerOf(&_Card.CallOpts, tokenId)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Card *CardCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Card.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Card *CardSession) ProxiableUUID() ([32]byte, error) {
	return _Card.Contract.ProxiableUUID(&_Card.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Card *CardCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Card.Contract.ProxiableUUID(&_Card.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Card *CardCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Card.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Card *CardSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Card.Contract.SupportsInterface(&_Card.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Card *CardCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Card.Contract.SupportsInterface(&_Card.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Card *CardCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Card.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Card *CardSession) Symbol() (string, error) {
	return _Card.Contract.Symbol(&_Card.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Card *CardCallerSession) Symbol() (string, error) {
	return _Card.Contract.Symbol(&_Card.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Card *CardCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Card.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Card *CardSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Card.Contract.TokenByIndex(&_Card.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Card *CardCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Card.Contract.TokenByIndex(&_Card.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Card *CardCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Card.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Card *CardSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Card.Contract.TokenOfOwnerByIndex(&_Card.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Card *CardCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Card.Contract.TokenOfOwnerByIndex(&_Card.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 slot) view returns(string)
func (_Card *CardCaller) TokenURI(opts *bind.CallOpts, slot *big.Int) (string, error) {
	var out []interface{}
	err := _Card.contract.Call(opts, &out, "tokenURI", slot)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 slot) view returns(string)
func (_Card *CardSession) TokenURI(slot *big.Int) (string, error) {
	return _Card.Contract.TokenURI(&_Card.CallOpts, slot)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 slot) view returns(string)
func (_Card *CardCallerSession) TokenURI(slot *big.Int) (string, error) {
	return _Card.Contract.TokenURI(&_Card.CallOpts, slot)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Card *CardCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Card.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Card *CardSession) TotalSupply() (*big.Int, error) {
	return _Card.Contract.TotalSupply(&_Card.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Card *CardCallerSession) TotalSupply() (*big.Int, error) {
	return _Card.Contract.TotalSupply(&_Card.CallOpts)
}

// AirDrop is a paid mutator transaction binding the contract method 0xa4f67d98.
//
// Solidity: function AirDrop(uint64 slot, (uint8[3],uint8,uint64,string,string,string[]) info) returns()
func (_Card *CardTransactor) AirDrop(opts *bind.TransactOpts, slot uint64, info CardInfo) (*types.Transaction, error) {
	return _Card.contract.Transact(opts, "AirDrop", slot, info)
}

// AirDrop is a paid mutator transaction binding the contract method 0xa4f67d98.
//
// Solidity: function AirDrop(uint64 slot, (uint8[3],uint8,uint64,string,string,string[]) info) returns()
func (_Card *CardSession) AirDrop(slot uint64, info CardInfo) (*types.Transaction, error) {
	return _Card.Contract.AirDrop(&_Card.TransactOpts, slot, info)
}

// AirDrop is a paid mutator transaction binding the contract method 0xa4f67d98.
//
// Solidity: function AirDrop(uint64 slot, (uint8[3],uint8,uint64,string,string,string[]) info) returns()
func (_Card *CardTransactorSession) AirDrop(slot uint64, info CardInfo) (*types.Transaction, error) {
	return _Card.Contract.AirDrop(&_Card.TransactOpts, slot, info)
}

// Claim is a paid mutator transaction binding the contract method 0xbaf0a95c.
//
// Solidity: function Claim(uint64 slot) returns()
func (_Card *CardTransactor) Claim(opts *bind.TransactOpts, slot uint64) (*types.Transaction, error) {
	return _Card.contract.Transact(opts, "Claim", slot)
}

// Claim is a paid mutator transaction binding the contract method 0xbaf0a95c.
//
// Solidity: function Claim(uint64 slot) returns()
func (_Card *CardSession) Claim(slot uint64) (*types.Transaction, error) {
	return _Card.Contract.Claim(&_Card.TransactOpts, slot)
}

// Claim is a paid mutator transaction binding the contract method 0xbaf0a95c.
//
// Solidity: function Claim(uint64 slot) returns()
func (_Card *CardTransactorSession) Claim(slot uint64) (*types.Transaction, error) {
	return _Card.Contract.Claim(&_Card.TransactOpts, slot)
}

// Customize is a paid mutator transaction binding the contract method 0x923cd5cb.
//
// Solidity: function Customize(uint64 slot, uint8 renderOpt, string greeting, string greetingImg) returns()
func (_Card *CardTransactor) Customize(opts *bind.TransactOpts, slot uint64, renderOpt uint8, greeting string, greetingImg string) (*types.Transaction, error) {
	return _Card.contract.Transact(opts, "Customize", slot, renderOpt, greeting, greetingImg)
}

// Customize is a paid mutator transaction binding the contract method 0x923cd5cb.
//
// Solidity: function Customize(uint64 slot, uint8 renderOpt, string greeting, string greetingImg) returns()
func (_Card *CardSession) Customize(slot uint64, renderOpt uint8, greeting string, greetingImg string) (*types.Transaction, error) {
	return _Card.Contract.Customize(&_Card.TransactOpts, slot, renderOpt, greeting, greetingImg)
}

// Customize is a paid mutator transaction binding the contract method 0x923cd5cb.
//
// Solidity: function Customize(uint64 slot, uint8 renderOpt, string greeting, string greetingImg) returns()
func (_Card *CardTransactorSession) Customize(slot uint64, renderOpt uint8, greeting string, greetingImg string) (*types.Transaction, error) {
	return _Card.Contract.Customize(&_Card.TransactOpts, slot, renderOpt, greeting, greetingImg)
}

// Mint is a paid mutator transaction binding the contract method 0x71703533.
//
// Solidity: function Mint(uint64 slot, (uint8[3],uint8,uint64,string,string,string[]) info) returns()
func (_Card *CardTransactor) Mint(opts *bind.TransactOpts, slot uint64, info CardInfo) (*types.Transaction, error) {
	return _Card.contract.Transact(opts, "Mint", slot, info)
}

// Mint is a paid mutator transaction binding the contract method 0x71703533.
//
// Solidity: function Mint(uint64 slot, (uint8[3],uint8,uint64,string,string,string[]) info) returns()
func (_Card *CardSession) Mint(slot uint64, info CardInfo) (*types.Transaction, error) {
	return _Card.Contract.Mint(&_Card.TransactOpts, slot, info)
}

// Mint is a paid mutator transaction binding the contract method 0x71703533.
//
// Solidity: function Mint(uint64 slot, (uint8[3],uint8,uint64,string,string,string[]) info) returns()
func (_Card *CardTransactorSession) Mint(slot uint64, info CardInfo) (*types.Transaction, error) {
	return _Card.Contract.Mint(&_Card.TransactOpts, slot, info)
}

// SetBaseImg is a paid mutator transaction binding the contract method 0x5dec4e20.
//
// Solidity: function SetBaseImg(string baseImg) returns()
func (_Card *CardTransactor) SetBaseImg(opts *bind.TransactOpts, baseImg string) (*types.Transaction, error) {
	return _Card.contract.Transact(opts, "SetBaseImg", baseImg)
}

// SetBaseImg is a paid mutator transaction binding the contract method 0x5dec4e20.
//
// Solidity: function SetBaseImg(string baseImg) returns()
func (_Card *CardSession) SetBaseImg(baseImg string) (*types.Transaction, error) {
	return _Card.Contract.SetBaseImg(&_Card.TransactOpts, baseImg)
}

// SetBaseImg is a paid mutator transaction binding the contract method 0x5dec4e20.
//
// Solidity: function SetBaseImg(string baseImg) returns()
func (_Card *CardTransactorSession) SetBaseImg(baseImg string) (*types.Transaction, error) {
	return _Card.Contract.SetBaseImg(&_Card.TransactOpts, baseImg)
}

// SetBaseImg0 is a paid mutator transaction binding the contract method 0xe0f18f3f.
//
// Solidity: function SetBaseImg(uint64 slot, string baseImg) returns()
func (_Card *CardTransactor) SetBaseImg0(opts *bind.TransactOpts, slot uint64, baseImg string) (*types.Transaction, error) {
	return _Card.contract.Transact(opts, "SetBaseImg0", slot, baseImg)
}

// SetBaseImg0 is a paid mutator transaction binding the contract method 0xe0f18f3f.
//
// Solidity: function SetBaseImg(uint64 slot, string baseImg) returns()
func (_Card *CardSession) SetBaseImg0(slot uint64, baseImg string) (*types.Transaction, error) {
	return _Card.Contract.SetBaseImg0(&_Card.TransactOpts, slot, baseImg)
}

// SetBaseImg0 is a paid mutator transaction binding the contract method 0xe0f18f3f.
//
// Solidity: function SetBaseImg(uint64 slot, string baseImg) returns()
func (_Card *CardTransactorSession) SetBaseImg0(slot uint64, baseImg string) (*types.Transaction, error) {
	return _Card.Contract.SetBaseImg0(&_Card.TransactOpts, slot, baseImg)
}

// SetHistoryAddr is a paid mutator transaction binding the contract method 0x84a3a035.
//
// Solidity: function SetHistoryAddr(address historyAddr) returns()
func (_Card *CardTransactor) SetHistoryAddr(opts *bind.TransactOpts, historyAddr common.Address) (*types.Transaction, error) {
	return _Card.contract.Transact(opts, "SetHistoryAddr", historyAddr)
}

// SetHistoryAddr is a paid mutator transaction binding the contract method 0x84a3a035.
//
// Solidity: function SetHistoryAddr(address historyAddr) returns()
func (_Card *CardSession) SetHistoryAddr(historyAddr common.Address) (*types.Transaction, error) {
	return _Card.Contract.SetHistoryAddr(&_Card.TransactOpts, historyAddr)
}

// SetHistoryAddr is a paid mutator transaction binding the contract method 0x84a3a035.
//
// Solidity: function SetHistoryAddr(address historyAddr) returns()
func (_Card *CardTransactorSession) SetHistoryAddr(historyAddr common.Address) (*types.Transaction, error) {
	return _Card.Contract.SetHistoryAddr(&_Card.TransactOpts, historyAddr)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Card *CardTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Card.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Card *CardSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Card.Contract.Approve(&_Card.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Card *CardTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Card.Contract.Approve(&_Card.TransactOpts, to, tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Card *CardTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Card.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Card *CardSession) Initialize() (*types.Transaction, error) {
	return _Card.Contract.Initialize(&_Card.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Card *CardTransactorSession) Initialize() (*types.Transaction, error) {
	return _Card.Contract.Initialize(&_Card.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Card *CardTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Card.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Card *CardSession) RenounceOwnership() (*types.Transaction, error) {
	return _Card.Contract.RenounceOwnership(&_Card.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Card *CardTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Card.Contract.RenounceOwnership(&_Card.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Card *CardTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Card.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Card *CardSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Card.Contract.SafeTransferFrom(&_Card.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Card *CardTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Card.Contract.SafeTransferFrom(&_Card.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Card *CardTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Card.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Card *CardSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Card.Contract.SafeTransferFrom0(&_Card.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Card *CardTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Card.Contract.SafeTransferFrom0(&_Card.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Card *CardTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Card.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Card *CardSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Card.Contract.SetApprovalForAll(&_Card.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Card *CardTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Card.Contract.SetApprovalForAll(&_Card.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Card *CardTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Card.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Card *CardSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Card.Contract.TransferFrom(&_Card.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Card *CardTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Card.Contract.TransferFrom(&_Card.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Card *CardTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Card.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Card *CardSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Card.Contract.TransferOwnership(&_Card.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Card *CardTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Card.Contract.TransferOwnership(&_Card.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Card *CardTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Card.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Card *CardSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Card.Contract.UpgradeTo(&_Card.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Card *CardTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Card.Contract.UpgradeTo(&_Card.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Card *CardTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Card.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Card *CardSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Card.Contract.UpgradeToAndCall(&_Card.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Card *CardTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Card.Contract.UpgradeToAndCall(&_Card.TransactOpts, newImplementation, data)
}

// CardAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Card contract.
type CardAdminChangedIterator struct {
	Event *CardAdminChanged // Event containing the contract specifics and raw log

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
func (it *CardAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardAdminChanged)
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
		it.Event = new(CardAdminChanged)
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
func (it *CardAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardAdminChanged represents a AdminChanged event raised by the Card contract.
type CardAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Card *CardFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*CardAdminChangedIterator, error) {

	logs, sub, err := _Card.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &CardAdminChangedIterator{contract: _Card.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Card *CardFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *CardAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Card.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardAdminChanged)
				if err := _Card.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Card *CardFilterer) ParseAdminChanged(log types.Log) (*CardAdminChanged, error) {
	event := new(CardAdminChanged)
	if err := _Card.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CardApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Card contract.
type CardApprovalIterator struct {
	Event *CardApproval // Event containing the contract specifics and raw log

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
func (it *CardApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardApproval)
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
		it.Event = new(CardApproval)
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
func (it *CardApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardApproval represents a Approval event raised by the Card contract.
type CardApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Card *CardFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*CardApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Card.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CardApprovalIterator{contract: _Card.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Card *CardFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CardApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Card.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardApproval)
				if err := _Card.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Card *CardFilterer) ParseApproval(log types.Log) (*CardApproval, error) {
	event := new(CardApproval)
	if err := _Card.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CardApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Card contract.
type CardApprovalForAllIterator struct {
	Event *CardApprovalForAll // Event containing the contract specifics and raw log

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
func (it *CardApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardApprovalForAll)
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
		it.Event = new(CardApprovalForAll)
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
func (it *CardApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardApprovalForAll represents a ApprovalForAll event raised by the Card contract.
type CardApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Card *CardFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*CardApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Card.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CardApprovalForAllIterator{contract: _Card.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Card *CardFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CardApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Card.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardApprovalForAll)
				if err := _Card.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Card *CardFilterer) ParseApprovalForAll(log types.Log) (*CardApprovalForAll, error) {
	event := new(CardApprovalForAll)
	if err := _Card.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CardBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Card contract.
type CardBeaconUpgradedIterator struct {
	Event *CardBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *CardBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardBeaconUpgraded)
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
		it.Event = new(CardBeaconUpgraded)
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
func (it *CardBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardBeaconUpgraded represents a BeaconUpgraded event raised by the Card contract.
type CardBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Card *CardFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*CardBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Card.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &CardBeaconUpgradedIterator{contract: _Card.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Card *CardFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *CardBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Card.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardBeaconUpgraded)
				if err := _Card.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Card *CardFilterer) ParseBeaconUpgraded(log types.Log) (*CardBeaconUpgraded, error) {
	event := new(CardBeaconUpgraded)
	if err := _Card.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CardInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Card contract.
type CardInitializedIterator struct {
	Event *CardInitialized // Event containing the contract specifics and raw log

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
func (it *CardInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardInitialized)
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
		it.Event = new(CardInitialized)
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
func (it *CardInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardInitialized represents a Initialized event raised by the Card contract.
type CardInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Card *CardFilterer) FilterInitialized(opts *bind.FilterOpts) (*CardInitializedIterator, error) {

	logs, sub, err := _Card.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CardInitializedIterator{contract: _Card.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Card *CardFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CardInitialized) (event.Subscription, error) {

	logs, sub, err := _Card.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardInitialized)
				if err := _Card.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Card *CardFilterer) ParseInitialized(log types.Log) (*CardInitialized, error) {
	event := new(CardInitialized)
	if err := _Card.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CardOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Card contract.
type CardOwnershipTransferredIterator struct {
	Event *CardOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CardOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardOwnershipTransferred)
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
		it.Event = new(CardOwnershipTransferred)
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
func (it *CardOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardOwnershipTransferred represents a OwnershipTransferred event raised by the Card contract.
type CardOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Card *CardFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CardOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Card.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CardOwnershipTransferredIterator{contract: _Card.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Card *CardFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CardOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Card.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardOwnershipTransferred)
				if err := _Card.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Card *CardFilterer) ParseOwnershipTransferred(log types.Log) (*CardOwnershipTransferred, error) {
	event := new(CardOwnershipTransferred)
	if err := _Card.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CardTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Card contract.
type CardTransferIterator struct {
	Event *CardTransfer // Event containing the contract specifics and raw log

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
func (it *CardTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardTransfer)
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
		it.Event = new(CardTransfer)
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
func (it *CardTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardTransfer represents a Transfer event raised by the Card contract.
type CardTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Card *CardFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*CardTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Card.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CardTransferIterator{contract: _Card.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Card *CardFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CardTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Card.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardTransfer)
				if err := _Card.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Card *CardFilterer) ParseTransfer(log types.Log) (*CardTransfer, error) {
	event := new(CardTransfer)
	if err := _Card.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CardUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Card contract.
type CardUpgradedIterator struct {
	Event *CardUpgraded // Event containing the contract specifics and raw log

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
func (it *CardUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardUpgraded)
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
		it.Event = new(CardUpgraded)
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
func (it *CardUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardUpgraded represents a Upgraded event raised by the Card contract.
type CardUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Card *CardFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*CardUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Card.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &CardUpgradedIterator{contract: _Card.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Card *CardFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CardUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Card.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardUpgraded)
				if err := _Card.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Card *CardFilterer) ParseUpgraded(log types.Log) (*CardUpgraded, error) {
	event := new(CardUpgraded)
	if err := _Card.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
