// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cryptoPunksMarket

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CryptoPunksMarketABI is the input ABI used to generate the binding from.
const CryptoPunksMarketABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"punksOfferedForSale\",\"outputs\":[{\"name\":\"isForSale\",\"type\":\"bool\"},{\"name\":\"punkIndex\",\"type\":\"uint256\"},{\"name\":\"seller\",\"type\":\"address\"},{\"name\":\"minValue\",\"type\":\"uint256\"},{\"name\":\"onlySellTo\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"enterBidForPunk\",\"outputs\":[],\"payable\":true,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"punkIndex\",\"type\":\"uint256\"},{\"name\":\"minPrice\",\"type\":\"uint256\"}],\"name\":\"acceptBidForPunk\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addresses\",\"type\":\"address[]\"},{\"name\":\"indices\",\"type\":\"uint256[]\"}],\"name\":\"setInitialOwners\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"imageHash\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextPunkIndexToAssign\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"punkIndexToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"standard\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"punkBids\",\"outputs\":[{\"name\":\"hasBid\",\"type\":\"bool\"},{\"name\":\"punkIndex\",\"type\":\"uint256\"},{\"name\":\"bidder\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"allInitialOwnersAssigned\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allPunksAssigned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"buyPunk\",\"outputs\":[],\"payable\":true,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"transferPunk\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"withdrawBidForPunk\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"setInitialOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"punkIndex\",\"type\":\"uint256\"},{\"name\":\"minSalePriceInWei\",\"type\":\"uint256\"},{\"name\":\"toAddress\",\"type\":\"address\"}],\"name\":\"offerPunkForSaleToAddress\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"punksRemainingToAssign\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"punkIndex\",\"type\":\"uint256\"},{\"name\":\"minSalePriceInWei\",\"type\":\"uint256\"}],\"name\":\"offerPunkForSale\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"getPunk\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingWithdrawals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"punkNoLongerForSale\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"Assign\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"PunkTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"punkIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"minValue\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"toAddress\",\"type\":\"address\"}],\"name\":\"PunkOffered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"punkIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"fromAddress\",\"type\":\"address\"}],\"name\":\"PunkBidEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"punkIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"fromAddress\",\"type\":\"address\"}],\"name\":\"PunkBidWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"punkIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"fromAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"toAddress\",\"type\":\"address\"}],\"name\":\"PunkBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"PunkNoLongerForSale\",\"type\":\"event\"}]"

// CryptoPunksMarket is an auto generated Go binding around an Ethereum contract.
type CryptoPunksMarket struct {
	CryptoPunksMarketCaller     // Read-only binding to the contract
	CryptoPunksMarketTransactor // Write-only binding to the contract
	CryptoPunksMarketFilterer   // Log filterer for contract events
}

// CryptoPunksMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type CryptoPunksMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptoPunksMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CryptoPunksMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptoPunksMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CryptoPunksMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptoPunksMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CryptoPunksMarketSession struct {
	Contract     *CryptoPunksMarket // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CryptoPunksMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CryptoPunksMarketCallerSession struct {
	Contract *CryptoPunksMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// CryptoPunksMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CryptoPunksMarketTransactorSession struct {
	Contract     *CryptoPunksMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// CryptoPunksMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type CryptoPunksMarketRaw struct {
	Contract *CryptoPunksMarket // Generic contract binding to access the raw methods on
}

// CryptoPunksMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CryptoPunksMarketCallerRaw struct {
	Contract *CryptoPunksMarketCaller // Generic read-only contract binding to access the raw methods on
}

// CryptoPunksMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CryptoPunksMarketTransactorRaw struct {
	Contract *CryptoPunksMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCryptoPunksMarket creates a new instance of CryptoPunksMarket, bound to a specific deployed contract.
func NewCryptoPunksMarket(address common.Address, backend bind.ContractBackend) (*CryptoPunksMarket, error) {
	contract, err := bindCryptoPunksMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CryptoPunksMarket{CryptoPunksMarketCaller: CryptoPunksMarketCaller{contract: contract}, CryptoPunksMarketTransactor: CryptoPunksMarketTransactor{contract: contract}, CryptoPunksMarketFilterer: CryptoPunksMarketFilterer{contract: contract}}, nil
}

// NewCryptoPunksMarketCaller creates a new read-only instance of CryptoPunksMarket, bound to a specific deployed contract.
func NewCryptoPunksMarketCaller(address common.Address, caller bind.ContractCaller) (*CryptoPunksMarketCaller, error) {
	contract, err := bindCryptoPunksMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CryptoPunksMarketCaller{contract: contract}, nil
}

// NewCryptoPunksMarketTransactor creates a new write-only instance of CryptoPunksMarket, bound to a specific deployed contract.
func NewCryptoPunksMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*CryptoPunksMarketTransactor, error) {
	contract, err := bindCryptoPunksMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CryptoPunksMarketTransactor{contract: contract}, nil
}

// NewCryptoPunksMarketFilterer creates a new log filterer instance of CryptoPunksMarket, bound to a specific deployed contract.
func NewCryptoPunksMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*CryptoPunksMarketFilterer, error) {
	contract, err := bindCryptoPunksMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CryptoPunksMarketFilterer{contract: contract}, nil
}

// bindCryptoPunksMarket binds a generic wrapper to an already deployed contract.
func bindCryptoPunksMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CryptoPunksMarketABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CryptoPunksMarket *CryptoPunksMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CryptoPunksMarket.Contract.CryptoPunksMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CryptoPunksMarket *CryptoPunksMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.CryptoPunksMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CryptoPunksMarket *CryptoPunksMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.CryptoPunksMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CryptoPunksMarket *CryptoPunksMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CryptoPunksMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CryptoPunksMarket *CryptoPunksMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CryptoPunksMarket *CryptoPunksMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.contract.Transact(opts, method, params...)
}

// AllPunksAssigned is a free data retrieval call binding the contract method 0x8126c38a.
//
// Solidity: function allPunksAssigned() returns(bool)
func (_CryptoPunksMarket *CryptoPunksMarketCaller) AllPunksAssigned(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CryptoPunksMarket.contract.Call(opts, &out, "allPunksAssigned")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllPunksAssigned is a free data retrieval call binding the contract method 0x8126c38a.
//
// Solidity: function allPunksAssigned() returns(bool)
func (_CryptoPunksMarket *CryptoPunksMarketSession) AllPunksAssigned() (bool, error) {
	return _CryptoPunksMarket.Contract.AllPunksAssigned(&_CryptoPunksMarket.CallOpts)
}

// AllPunksAssigned is a free data retrieval call binding the contract method 0x8126c38a.
//
// Solidity: function allPunksAssigned() returns(bool)
func (_CryptoPunksMarket *CryptoPunksMarketCallerSession) AllPunksAssigned() (bool, error) {
	return _CryptoPunksMarket.Contract.AllPunksAssigned(&_CryptoPunksMarket.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CryptoPunksMarket.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CryptoPunksMarket.Contract.BalanceOf(&_CryptoPunksMarket.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CryptoPunksMarket.Contract.BalanceOf(&_CryptoPunksMarket.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint8)
func (_CryptoPunksMarket *CryptoPunksMarketCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CryptoPunksMarket.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint8)
func (_CryptoPunksMarket *CryptoPunksMarketSession) Decimals() (uint8, error) {
	return _CryptoPunksMarket.Contract.Decimals(&_CryptoPunksMarket.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint8)
func (_CryptoPunksMarket *CryptoPunksMarketCallerSession) Decimals() (uint8, error) {
	return _CryptoPunksMarket.Contract.Decimals(&_CryptoPunksMarket.CallOpts)
}

// ImageHash is a free data retrieval call binding the contract method 0x51605d80.
//
// Solidity: function imageHash() returns(string)
func (_CryptoPunksMarket *CryptoPunksMarketCaller) ImageHash(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CryptoPunksMarket.contract.Call(opts, &out, "imageHash")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ImageHash is a free data retrieval call binding the contract method 0x51605d80.
//
// Solidity: function imageHash() returns(string)
func (_CryptoPunksMarket *CryptoPunksMarketSession) ImageHash() (string, error) {
	return _CryptoPunksMarket.Contract.ImageHash(&_CryptoPunksMarket.CallOpts)
}

// ImageHash is a free data retrieval call binding the contract method 0x51605d80.
//
// Solidity: function imageHash() returns(string)
func (_CryptoPunksMarket *CryptoPunksMarketCallerSession) ImageHash() (string, error) {
	return _CryptoPunksMarket.Contract.ImageHash(&_CryptoPunksMarket.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_CryptoPunksMarket *CryptoPunksMarketCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CryptoPunksMarket.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_CryptoPunksMarket *CryptoPunksMarketSession) Name() (string, error) {
	return _CryptoPunksMarket.Contract.Name(&_CryptoPunksMarket.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_CryptoPunksMarket *CryptoPunksMarketCallerSession) Name() (string, error) {
	return _CryptoPunksMarket.Contract.Name(&_CryptoPunksMarket.CallOpts)
}

// NextPunkIndexToAssign is a free data retrieval call binding the contract method 0x52f29a25.
//
// Solidity: function nextPunkIndexToAssign() returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketCaller) NextPunkIndexToAssign(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CryptoPunksMarket.contract.Call(opts, &out, "nextPunkIndexToAssign")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextPunkIndexToAssign is a free data retrieval call binding the contract method 0x52f29a25.
//
// Solidity: function nextPunkIndexToAssign() returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketSession) NextPunkIndexToAssign() (*big.Int, error) {
	return _CryptoPunksMarket.Contract.NextPunkIndexToAssign(&_CryptoPunksMarket.CallOpts)
}

// NextPunkIndexToAssign is a free data retrieval call binding the contract method 0x52f29a25.
//
// Solidity: function nextPunkIndexToAssign() returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketCallerSession) NextPunkIndexToAssign() (*big.Int, error) {
	return _CryptoPunksMarket.Contract.NextPunkIndexToAssign(&_CryptoPunksMarket.CallOpts)
}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xf3f43703.
//
// Solidity: function pendingWithdrawals(address ) returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketCaller) PendingWithdrawals(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CryptoPunksMarket.contract.Call(opts, &out, "pendingWithdrawals", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xf3f43703.
//
// Solidity: function pendingWithdrawals(address ) returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketSession) PendingWithdrawals(arg0 common.Address) (*big.Int, error) {
	return _CryptoPunksMarket.Contract.PendingWithdrawals(&_CryptoPunksMarket.CallOpts, arg0)
}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xf3f43703.
//
// Solidity: function pendingWithdrawals(address ) returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketCallerSession) PendingWithdrawals(arg0 common.Address) (*big.Int, error) {
	return _CryptoPunksMarket.Contract.PendingWithdrawals(&_CryptoPunksMarket.CallOpts, arg0)
}

// PunkBids is a free data retrieval call binding the contract method 0x6e743fa9.
//
// Solidity: function punkBids(uint256 ) returns(bool hasBid, uint256 punkIndex, address bidder, uint256 value)
func (_CryptoPunksMarket *CryptoPunksMarketCaller) PunkBids(opts *bind.CallOpts, arg0 *big.Int) (struct {
	HasBid    bool
	PunkIndex *big.Int
	Bidder    common.Address
	Value     *big.Int
}, error) {
	var out []interface{}
	err := _CryptoPunksMarket.contract.Call(opts, &out, "punkBids", arg0)

	outstruct := new(struct {
		HasBid    bool
		PunkIndex *big.Int
		Bidder    common.Address
		Value     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.HasBid = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.PunkIndex = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Bidder = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Value = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PunkBids is a free data retrieval call binding the contract method 0x6e743fa9.
//
// Solidity: function punkBids(uint256 ) returns(bool hasBid, uint256 punkIndex, address bidder, uint256 value)
func (_CryptoPunksMarket *CryptoPunksMarketSession) PunkBids(arg0 *big.Int) (struct {
	HasBid    bool
	PunkIndex *big.Int
	Bidder    common.Address
	Value     *big.Int
}, error) {
	return _CryptoPunksMarket.Contract.PunkBids(&_CryptoPunksMarket.CallOpts, arg0)
}

// PunkBids is a free data retrieval call binding the contract method 0x6e743fa9.
//
// Solidity: function punkBids(uint256 ) returns(bool hasBid, uint256 punkIndex, address bidder, uint256 value)
func (_CryptoPunksMarket *CryptoPunksMarketCallerSession) PunkBids(arg0 *big.Int) (struct {
	HasBid    bool
	PunkIndex *big.Int
	Bidder    common.Address
	Value     *big.Int
}, error) {
	return _CryptoPunksMarket.Contract.PunkBids(&_CryptoPunksMarket.CallOpts, arg0)
}

// PunkIndexToAddress is a free data retrieval call binding the contract method 0x58178168.
//
// Solidity: function punkIndexToAddress(uint256 ) returns(address)
func (_CryptoPunksMarket *CryptoPunksMarketCaller) PunkIndexToAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CryptoPunksMarket.contract.Call(opts, &out, "punkIndexToAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PunkIndexToAddress is a free data retrieval call binding the contract method 0x58178168.
//
// Solidity: function punkIndexToAddress(uint256 ) returns(address)
func (_CryptoPunksMarket *CryptoPunksMarketSession) PunkIndexToAddress(arg0 *big.Int) (common.Address, error) {
	return _CryptoPunksMarket.Contract.PunkIndexToAddress(&_CryptoPunksMarket.CallOpts, arg0)
}

// PunkIndexToAddress is a free data retrieval call binding the contract method 0x58178168.
//
// Solidity: function punkIndexToAddress(uint256 ) returns(address)
func (_CryptoPunksMarket *CryptoPunksMarketCallerSession) PunkIndexToAddress(arg0 *big.Int) (common.Address, error) {
	return _CryptoPunksMarket.Contract.PunkIndexToAddress(&_CryptoPunksMarket.CallOpts, arg0)
}

// PunksOfferedForSale is a free data retrieval call binding the contract method 0x088f11f3.
//
// Solidity: function punksOfferedForSale(uint256 ) returns(bool isForSale, uint256 punkIndex, address seller, uint256 minValue, address onlySellTo)
func (_CryptoPunksMarket *CryptoPunksMarketCaller) PunksOfferedForSale(opts *bind.CallOpts, arg0 *big.Int) (struct {
	IsForSale  bool
	PunkIndex  *big.Int
	Seller     common.Address
	MinValue   *big.Int
	OnlySellTo common.Address
}, error) {
	var out []interface{}
	err := _CryptoPunksMarket.contract.Call(opts, &out, "punksOfferedForSale", arg0)

	outstruct := new(struct {
		IsForSale  bool
		PunkIndex  *big.Int
		Seller     common.Address
		MinValue   *big.Int
		OnlySellTo common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsForSale = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.PunkIndex = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Seller = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.MinValue = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.OnlySellTo = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// PunksOfferedForSale is a free data retrieval call binding the contract method 0x088f11f3.
//
// Solidity: function punksOfferedForSale(uint256 ) returns(bool isForSale, uint256 punkIndex, address seller, uint256 minValue, address onlySellTo)
func (_CryptoPunksMarket *CryptoPunksMarketSession) PunksOfferedForSale(arg0 *big.Int) (struct {
	IsForSale  bool
	PunkIndex  *big.Int
	Seller     common.Address
	MinValue   *big.Int
	OnlySellTo common.Address
}, error) {
	return _CryptoPunksMarket.Contract.PunksOfferedForSale(&_CryptoPunksMarket.CallOpts, arg0)
}

// PunksOfferedForSale is a free data retrieval call binding the contract method 0x088f11f3.
//
// Solidity: function punksOfferedForSale(uint256 ) returns(bool isForSale, uint256 punkIndex, address seller, uint256 minValue, address onlySellTo)
func (_CryptoPunksMarket *CryptoPunksMarketCallerSession) PunksOfferedForSale(arg0 *big.Int) (struct {
	IsForSale  bool
	PunkIndex  *big.Int
	Seller     common.Address
	MinValue   *big.Int
	OnlySellTo common.Address
}, error) {
	return _CryptoPunksMarket.Contract.PunksOfferedForSale(&_CryptoPunksMarket.CallOpts, arg0)
}

// PunksRemainingToAssign is a free data retrieval call binding the contract method 0xc0d6ce63.
//
// Solidity: function punksRemainingToAssign() returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketCaller) PunksRemainingToAssign(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CryptoPunksMarket.contract.Call(opts, &out, "punksRemainingToAssign")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PunksRemainingToAssign is a free data retrieval call binding the contract method 0xc0d6ce63.
//
// Solidity: function punksRemainingToAssign() returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketSession) PunksRemainingToAssign() (*big.Int, error) {
	return _CryptoPunksMarket.Contract.PunksRemainingToAssign(&_CryptoPunksMarket.CallOpts)
}

// PunksRemainingToAssign is a free data retrieval call binding the contract method 0xc0d6ce63.
//
// Solidity: function punksRemainingToAssign() returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketCallerSession) PunksRemainingToAssign() (*big.Int, error) {
	return _CryptoPunksMarket.Contract.PunksRemainingToAssign(&_CryptoPunksMarket.CallOpts)
}

// Standard is a free data retrieval call binding the contract method 0x5a3b7e42.
//
// Solidity: function standard() returns(string)
func (_CryptoPunksMarket *CryptoPunksMarketCaller) Standard(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CryptoPunksMarket.contract.Call(opts, &out, "standard")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Standard is a free data retrieval call binding the contract method 0x5a3b7e42.
//
// Solidity: function standard() returns(string)
func (_CryptoPunksMarket *CryptoPunksMarketSession) Standard() (string, error) {
	return _CryptoPunksMarket.Contract.Standard(&_CryptoPunksMarket.CallOpts)
}

// Standard is a free data retrieval call binding the contract method 0x5a3b7e42.
//
// Solidity: function standard() returns(string)
func (_CryptoPunksMarket *CryptoPunksMarketCallerSession) Standard() (string, error) {
	return _CryptoPunksMarket.Contract.Standard(&_CryptoPunksMarket.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_CryptoPunksMarket *CryptoPunksMarketCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CryptoPunksMarket.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_CryptoPunksMarket *CryptoPunksMarketSession) Symbol() (string, error) {
	return _CryptoPunksMarket.Contract.Symbol(&_CryptoPunksMarket.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_CryptoPunksMarket *CryptoPunksMarketCallerSession) Symbol() (string, error) {
	return _CryptoPunksMarket.Contract.Symbol(&_CryptoPunksMarket.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CryptoPunksMarket.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketSession) TotalSupply() (*big.Int, error) {
	return _CryptoPunksMarket.Contract.TotalSupply(&_CryptoPunksMarket.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketCallerSession) TotalSupply() (*big.Int, error) {
	return _CryptoPunksMarket.Contract.TotalSupply(&_CryptoPunksMarket.CallOpts)
}

// AcceptBidForPunk is a paid mutator transaction binding the contract method 0x23165b75.
//
// Solidity: function acceptBidForPunk(uint256 punkIndex, uint256 minPrice) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactor) AcceptBidForPunk(opts *bind.TransactOpts, punkIndex *big.Int, minPrice *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.contract.Transact(opts, "acceptBidForPunk", punkIndex, minPrice)
}

// AcceptBidForPunk is a paid mutator transaction binding the contract method 0x23165b75.
//
// Solidity: function acceptBidForPunk(uint256 punkIndex, uint256 minPrice) returns()
func (_CryptoPunksMarket *CryptoPunksMarketSession) AcceptBidForPunk(punkIndex *big.Int, minPrice *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.AcceptBidForPunk(&_CryptoPunksMarket.TransactOpts, punkIndex, minPrice)
}

// AcceptBidForPunk is a paid mutator transaction binding the contract method 0x23165b75.
//
// Solidity: function acceptBidForPunk(uint256 punkIndex, uint256 minPrice) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactorSession) AcceptBidForPunk(punkIndex *big.Int, minPrice *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.AcceptBidForPunk(&_CryptoPunksMarket.TransactOpts, punkIndex, minPrice)
}

// AllInitialOwnersAssigned is a paid mutator transaction binding the contract method 0x7ecedac9.
//
// Solidity: function allInitialOwnersAssigned() returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactor) AllInitialOwnersAssigned(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoPunksMarket.contract.Transact(opts, "allInitialOwnersAssigned")
}

// AllInitialOwnersAssigned is a paid mutator transaction binding the contract method 0x7ecedac9.
//
// Solidity: function allInitialOwnersAssigned() returns()
func (_CryptoPunksMarket *CryptoPunksMarketSession) AllInitialOwnersAssigned() (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.AllInitialOwnersAssigned(&_CryptoPunksMarket.TransactOpts)
}

// AllInitialOwnersAssigned is a paid mutator transaction binding the contract method 0x7ecedac9.
//
// Solidity: function allInitialOwnersAssigned() returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactorSession) AllInitialOwnersAssigned() (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.AllInitialOwnersAssigned(&_CryptoPunksMarket.TransactOpts)
}

// BuyPunk is a paid mutator transaction binding the contract method 0x8264fe98.
//
// Solidity: function buyPunk(uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactor) BuyPunk(opts *bind.TransactOpts, punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.contract.Transact(opts, "buyPunk", punkIndex)
}

// BuyPunk is a paid mutator transaction binding the contract method 0x8264fe98.
//
// Solidity: function buyPunk(uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketSession) BuyPunk(punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.BuyPunk(&_CryptoPunksMarket.TransactOpts, punkIndex)
}

// BuyPunk is a paid mutator transaction binding the contract method 0x8264fe98.
//
// Solidity: function buyPunk(uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactorSession) BuyPunk(punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.BuyPunk(&_CryptoPunksMarket.TransactOpts, punkIndex)
}

// EnterBidForPunk is a paid mutator transaction binding the contract method 0x091dbfd2.
//
// Solidity: function enterBidForPunk(uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactor) EnterBidForPunk(opts *bind.TransactOpts, punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.contract.Transact(opts, "enterBidForPunk", punkIndex)
}

// EnterBidForPunk is a paid mutator transaction binding the contract method 0x091dbfd2.
//
// Solidity: function enterBidForPunk(uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketSession) EnterBidForPunk(punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.EnterBidForPunk(&_CryptoPunksMarket.TransactOpts, punkIndex)
}

// EnterBidForPunk is a paid mutator transaction binding the contract method 0x091dbfd2.
//
// Solidity: function enterBidForPunk(uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactorSession) EnterBidForPunk(punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.EnterBidForPunk(&_CryptoPunksMarket.TransactOpts, punkIndex)
}

// GetPunk is a paid mutator transaction binding the contract method 0xc81d1d5b.
//
// Solidity: function getPunk(uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactor) GetPunk(opts *bind.TransactOpts, punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.contract.Transact(opts, "getPunk", punkIndex)
}

// GetPunk is a paid mutator transaction binding the contract method 0xc81d1d5b.
//
// Solidity: function getPunk(uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketSession) GetPunk(punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.GetPunk(&_CryptoPunksMarket.TransactOpts, punkIndex)
}

// GetPunk is a paid mutator transaction binding the contract method 0xc81d1d5b.
//
// Solidity: function getPunk(uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactorSession) GetPunk(punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.GetPunk(&_CryptoPunksMarket.TransactOpts, punkIndex)
}

// OfferPunkForSale is a paid mutator transaction binding the contract method 0xc44193c3.
//
// Solidity: function offerPunkForSale(uint256 punkIndex, uint256 minSalePriceInWei) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactor) OfferPunkForSale(opts *bind.TransactOpts, punkIndex *big.Int, minSalePriceInWei *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.contract.Transact(opts, "offerPunkForSale", punkIndex, minSalePriceInWei)
}

// OfferPunkForSale is a paid mutator transaction binding the contract method 0xc44193c3.
//
// Solidity: function offerPunkForSale(uint256 punkIndex, uint256 minSalePriceInWei) returns()
func (_CryptoPunksMarket *CryptoPunksMarketSession) OfferPunkForSale(punkIndex *big.Int, minSalePriceInWei *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.OfferPunkForSale(&_CryptoPunksMarket.TransactOpts, punkIndex, minSalePriceInWei)
}

// OfferPunkForSale is a paid mutator transaction binding the contract method 0xc44193c3.
//
// Solidity: function offerPunkForSale(uint256 punkIndex, uint256 minSalePriceInWei) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactorSession) OfferPunkForSale(punkIndex *big.Int, minSalePriceInWei *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.OfferPunkForSale(&_CryptoPunksMarket.TransactOpts, punkIndex, minSalePriceInWei)
}

// OfferPunkForSaleToAddress is a paid mutator transaction binding the contract method 0xbf31196f.
//
// Solidity: function offerPunkForSaleToAddress(uint256 punkIndex, uint256 minSalePriceInWei, address toAddress) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactor) OfferPunkForSaleToAddress(opts *bind.TransactOpts, punkIndex *big.Int, minSalePriceInWei *big.Int, toAddress common.Address) (*types.Transaction, error) {
	return _CryptoPunksMarket.contract.Transact(opts, "offerPunkForSaleToAddress", punkIndex, minSalePriceInWei, toAddress)
}

// OfferPunkForSaleToAddress is a paid mutator transaction binding the contract method 0xbf31196f.
//
// Solidity: function offerPunkForSaleToAddress(uint256 punkIndex, uint256 minSalePriceInWei, address toAddress) returns()
func (_CryptoPunksMarket *CryptoPunksMarketSession) OfferPunkForSaleToAddress(punkIndex *big.Int, minSalePriceInWei *big.Int, toAddress common.Address) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.OfferPunkForSaleToAddress(&_CryptoPunksMarket.TransactOpts, punkIndex, minSalePriceInWei, toAddress)
}

// OfferPunkForSaleToAddress is a paid mutator transaction binding the contract method 0xbf31196f.
//
// Solidity: function offerPunkForSaleToAddress(uint256 punkIndex, uint256 minSalePriceInWei, address toAddress) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactorSession) OfferPunkForSaleToAddress(punkIndex *big.Int, minSalePriceInWei *big.Int, toAddress common.Address) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.OfferPunkForSaleToAddress(&_CryptoPunksMarket.TransactOpts, punkIndex, minSalePriceInWei, toAddress)
}

// PunkNoLongerForSale is a paid mutator transaction binding the contract method 0xf6eeff1e.
//
// Solidity: function punkNoLongerForSale(uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactor) PunkNoLongerForSale(opts *bind.TransactOpts, punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.contract.Transact(opts, "punkNoLongerForSale", punkIndex)
}

// PunkNoLongerForSale is a paid mutator transaction binding the contract method 0xf6eeff1e.
//
// Solidity: function punkNoLongerForSale(uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketSession) PunkNoLongerForSale(punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.PunkNoLongerForSale(&_CryptoPunksMarket.TransactOpts, punkIndex)
}

// PunkNoLongerForSale is a paid mutator transaction binding the contract method 0xf6eeff1e.
//
// Solidity: function punkNoLongerForSale(uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactorSession) PunkNoLongerForSale(punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.PunkNoLongerForSale(&_CryptoPunksMarket.TransactOpts, punkIndex)
}

// SetInitialOwner is a paid mutator transaction binding the contract method 0xa75a9049.
//
// Solidity: function setInitialOwner(address to, uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactor) SetInitialOwner(opts *bind.TransactOpts, to common.Address, punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.contract.Transact(opts, "setInitialOwner", to, punkIndex)
}

// SetInitialOwner is a paid mutator transaction binding the contract method 0xa75a9049.
//
// Solidity: function setInitialOwner(address to, uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketSession) SetInitialOwner(to common.Address, punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.SetInitialOwner(&_CryptoPunksMarket.TransactOpts, to, punkIndex)
}

// SetInitialOwner is a paid mutator transaction binding the contract method 0xa75a9049.
//
// Solidity: function setInitialOwner(address to, uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactorSession) SetInitialOwner(to common.Address, punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.SetInitialOwner(&_CryptoPunksMarket.TransactOpts, to, punkIndex)
}

// SetInitialOwners is a paid mutator transaction binding the contract method 0x39c5dde6.
//
// Solidity: function setInitialOwners(address[] addresses, uint256[] indices) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactor) SetInitialOwners(opts *bind.TransactOpts, addresses []common.Address, indices []*big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.contract.Transact(opts, "setInitialOwners", addresses, indices)
}

// SetInitialOwners is a paid mutator transaction binding the contract method 0x39c5dde6.
//
// Solidity: function setInitialOwners(address[] addresses, uint256[] indices) returns()
func (_CryptoPunksMarket *CryptoPunksMarketSession) SetInitialOwners(addresses []common.Address, indices []*big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.SetInitialOwners(&_CryptoPunksMarket.TransactOpts, addresses, indices)
}

// SetInitialOwners is a paid mutator transaction binding the contract method 0x39c5dde6.
//
// Solidity: function setInitialOwners(address[] addresses, uint256[] indices) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactorSession) SetInitialOwners(addresses []common.Address, indices []*big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.SetInitialOwners(&_CryptoPunksMarket.TransactOpts, addresses, indices)
}

// TransferPunk is a paid mutator transaction binding the contract method 0x8b72a2ec.
//
// Solidity: function transferPunk(address to, uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactor) TransferPunk(opts *bind.TransactOpts, to common.Address, punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.contract.Transact(opts, "transferPunk", to, punkIndex)
}

// TransferPunk is a paid mutator transaction binding the contract method 0x8b72a2ec.
//
// Solidity: function transferPunk(address to, uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketSession) TransferPunk(to common.Address, punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.TransferPunk(&_CryptoPunksMarket.TransactOpts, to, punkIndex)
}

// TransferPunk is a paid mutator transaction binding the contract method 0x8b72a2ec.
//
// Solidity: function transferPunk(address to, uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactorSession) TransferPunk(to common.Address, punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.TransferPunk(&_CryptoPunksMarket.TransactOpts, to, punkIndex)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoPunksMarket.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_CryptoPunksMarket *CryptoPunksMarketSession) Withdraw() (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.Withdraw(&_CryptoPunksMarket.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactorSession) Withdraw() (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.Withdraw(&_CryptoPunksMarket.TransactOpts)
}

// WithdrawBidForPunk is a paid mutator transaction binding the contract method 0x979bc638.
//
// Solidity: function withdrawBidForPunk(uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactor) WithdrawBidForPunk(opts *bind.TransactOpts, punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.contract.Transact(opts, "withdrawBidForPunk", punkIndex)
}

// WithdrawBidForPunk is a paid mutator transaction binding the contract method 0x979bc638.
//
// Solidity: function withdrawBidForPunk(uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketSession) WithdrawBidForPunk(punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.WithdrawBidForPunk(&_CryptoPunksMarket.TransactOpts, punkIndex)
}

// WithdrawBidForPunk is a paid mutator transaction binding the contract method 0x979bc638.
//
// Solidity: function withdrawBidForPunk(uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactorSession) WithdrawBidForPunk(punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.WithdrawBidForPunk(&_CryptoPunksMarket.TransactOpts, punkIndex)
}

// CryptoPunksMarketAssignIterator is returned from FilterAssign and is used to iterate over the raw logs and unpacked data for Assign events raised by the CryptoPunksMarket contract.
type CryptoPunksMarketAssignIterator struct {
	Event *CryptoPunksMarketAssign // Event containing the contract specifics and raw log

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
func (it *CryptoPunksMarketAssignIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoPunksMarketAssign)
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
		it.Event = new(CryptoPunksMarketAssign)
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
func (it *CryptoPunksMarketAssignIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoPunksMarketAssignIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoPunksMarketAssign represents a Assign event raised by the CryptoPunksMarket contract.
type CryptoPunksMarketAssign struct {
	To        common.Address
	PunkIndex *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAssign is a free log retrieval operation binding the contract event 0x8a0e37b73a0d9c82e205d4d1a3ff3d0b57ce5f4d7bccf6bac03336dc101cb7ba.
//
// Solidity: event Assign(address indexed to, uint256 punkIndex)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) FilterAssign(opts *bind.FilterOpts, to []common.Address) (*CryptoPunksMarketAssignIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.FilterLogs(opts, "Assign", toRule)
	if err != nil {
		return nil, err
	}
	return &CryptoPunksMarketAssignIterator{contract: _CryptoPunksMarket.contract, event: "Assign", logs: logs, sub: sub}, nil
}

// WatchAssign is a free log subscription operation binding the contract event 0x8a0e37b73a0d9c82e205d4d1a3ff3d0b57ce5f4d7bccf6bac03336dc101cb7ba.
//
// Solidity: event Assign(address indexed to, uint256 punkIndex)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) WatchAssign(opts *bind.WatchOpts, sink chan<- *CryptoPunksMarketAssign, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.WatchLogs(opts, "Assign", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoPunksMarketAssign)
				if err := _CryptoPunksMarket.contract.UnpackLog(event, "Assign", log); err != nil {
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

// ParseAssign is a log parse operation binding the contract event 0x8a0e37b73a0d9c82e205d4d1a3ff3d0b57ce5f4d7bccf6bac03336dc101cb7ba.
//
// Solidity: event Assign(address indexed to, uint256 punkIndex)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) ParseAssign(log types.Log) (*CryptoPunksMarketAssign, error) {
	event := new(CryptoPunksMarketAssign)
	if err := _CryptoPunksMarket.contract.UnpackLog(event, "Assign", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoPunksMarketPunkBidEnteredIterator is returned from FilterPunkBidEntered and is used to iterate over the raw logs and unpacked data for PunkBidEntered events raised by the CryptoPunksMarket contract.
type CryptoPunksMarketPunkBidEnteredIterator struct {
	Event *CryptoPunksMarketPunkBidEntered // Event containing the contract specifics and raw log

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
func (it *CryptoPunksMarketPunkBidEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoPunksMarketPunkBidEntered)
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
		it.Event = new(CryptoPunksMarketPunkBidEntered)
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
func (it *CryptoPunksMarketPunkBidEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoPunksMarketPunkBidEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoPunksMarketPunkBidEntered represents a PunkBidEntered event raised by the CryptoPunksMarket contract.
type CryptoPunksMarketPunkBidEntered struct {
	PunkIndex   *big.Int
	Value       *big.Int
	FromAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPunkBidEntered is a free log retrieval operation binding the contract event 0x5b859394fabae0c1ba88baffe67e751ab5248d2e879028b8c8d6897b0519f56a.
//
// Solidity: event PunkBidEntered(uint256 indexed punkIndex, uint256 value, address indexed fromAddress)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) FilterPunkBidEntered(opts *bind.FilterOpts, punkIndex []*big.Int, fromAddress []common.Address) (*CryptoPunksMarketPunkBidEnteredIterator, error) {

	var punkIndexRule []interface{}
	for _, punkIndexItem := range punkIndex {
		punkIndexRule = append(punkIndexRule, punkIndexItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.FilterLogs(opts, "PunkBidEntered", punkIndexRule, fromAddressRule)
	if err != nil {
		return nil, err
	}
	return &CryptoPunksMarketPunkBidEnteredIterator{contract: _CryptoPunksMarket.contract, event: "PunkBidEntered", logs: logs, sub: sub}, nil
}

// WatchPunkBidEntered is a free log subscription operation binding the contract event 0x5b859394fabae0c1ba88baffe67e751ab5248d2e879028b8c8d6897b0519f56a.
//
// Solidity: event PunkBidEntered(uint256 indexed punkIndex, uint256 value, address indexed fromAddress)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) WatchPunkBidEntered(opts *bind.WatchOpts, sink chan<- *CryptoPunksMarketPunkBidEntered, punkIndex []*big.Int, fromAddress []common.Address) (event.Subscription, error) {

	var punkIndexRule []interface{}
	for _, punkIndexItem := range punkIndex {
		punkIndexRule = append(punkIndexRule, punkIndexItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.WatchLogs(opts, "PunkBidEntered", punkIndexRule, fromAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoPunksMarketPunkBidEntered)
				if err := _CryptoPunksMarket.contract.UnpackLog(event, "PunkBidEntered", log); err != nil {
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

// ParsePunkBidEntered is a log parse operation binding the contract event 0x5b859394fabae0c1ba88baffe67e751ab5248d2e879028b8c8d6897b0519f56a.
//
// Solidity: event PunkBidEntered(uint256 indexed punkIndex, uint256 value, address indexed fromAddress)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) ParsePunkBidEntered(log types.Log) (*CryptoPunksMarketPunkBidEntered, error) {
	event := new(CryptoPunksMarketPunkBidEntered)
	if err := _CryptoPunksMarket.contract.UnpackLog(event, "PunkBidEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoPunksMarketPunkBidWithdrawnIterator is returned from FilterPunkBidWithdrawn and is used to iterate over the raw logs and unpacked data for PunkBidWithdrawn events raised by the CryptoPunksMarket contract.
type CryptoPunksMarketPunkBidWithdrawnIterator struct {
	Event *CryptoPunksMarketPunkBidWithdrawn // Event containing the contract specifics and raw log

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
func (it *CryptoPunksMarketPunkBidWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoPunksMarketPunkBidWithdrawn)
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
		it.Event = new(CryptoPunksMarketPunkBidWithdrawn)
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
func (it *CryptoPunksMarketPunkBidWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoPunksMarketPunkBidWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoPunksMarketPunkBidWithdrawn represents a PunkBidWithdrawn event raised by the CryptoPunksMarket contract.
type CryptoPunksMarketPunkBidWithdrawn struct {
	PunkIndex   *big.Int
	Value       *big.Int
	FromAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPunkBidWithdrawn is a free log retrieval operation binding the contract event 0x6f30e1ee4d81dcc7a8a478577f65d2ed2edb120565960ac45fe7c50551c87932.
//
// Solidity: event PunkBidWithdrawn(uint256 indexed punkIndex, uint256 value, address indexed fromAddress)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) FilterPunkBidWithdrawn(opts *bind.FilterOpts, punkIndex []*big.Int, fromAddress []common.Address) (*CryptoPunksMarketPunkBidWithdrawnIterator, error) {

	var punkIndexRule []interface{}
	for _, punkIndexItem := range punkIndex {
		punkIndexRule = append(punkIndexRule, punkIndexItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.FilterLogs(opts, "PunkBidWithdrawn", punkIndexRule, fromAddressRule)
	if err != nil {
		return nil, err
	}
	return &CryptoPunksMarketPunkBidWithdrawnIterator{contract: _CryptoPunksMarket.contract, event: "PunkBidWithdrawn", logs: logs, sub: sub}, nil
}

// WatchPunkBidWithdrawn is a free log subscription operation binding the contract event 0x6f30e1ee4d81dcc7a8a478577f65d2ed2edb120565960ac45fe7c50551c87932.
//
// Solidity: event PunkBidWithdrawn(uint256 indexed punkIndex, uint256 value, address indexed fromAddress)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) WatchPunkBidWithdrawn(opts *bind.WatchOpts, sink chan<- *CryptoPunksMarketPunkBidWithdrawn, punkIndex []*big.Int, fromAddress []common.Address) (event.Subscription, error) {

	var punkIndexRule []interface{}
	for _, punkIndexItem := range punkIndex {
		punkIndexRule = append(punkIndexRule, punkIndexItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.WatchLogs(opts, "PunkBidWithdrawn", punkIndexRule, fromAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoPunksMarketPunkBidWithdrawn)
				if err := _CryptoPunksMarket.contract.UnpackLog(event, "PunkBidWithdrawn", log); err != nil {
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

// ParsePunkBidWithdrawn is a log parse operation binding the contract event 0x6f30e1ee4d81dcc7a8a478577f65d2ed2edb120565960ac45fe7c50551c87932.
//
// Solidity: event PunkBidWithdrawn(uint256 indexed punkIndex, uint256 value, address indexed fromAddress)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) ParsePunkBidWithdrawn(log types.Log) (*CryptoPunksMarketPunkBidWithdrawn, error) {
	event := new(CryptoPunksMarketPunkBidWithdrawn)
	if err := _CryptoPunksMarket.contract.UnpackLog(event, "PunkBidWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoPunksMarketPunkBoughtIterator is returned from FilterPunkBought and is used to iterate over the raw logs and unpacked data for PunkBought events raised by the CryptoPunksMarket contract.
type CryptoPunksMarketPunkBoughtIterator struct {
	Event *CryptoPunksMarketPunkBought // Event containing the contract specifics and raw log

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
func (it *CryptoPunksMarketPunkBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoPunksMarketPunkBought)
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
		it.Event = new(CryptoPunksMarketPunkBought)
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
func (it *CryptoPunksMarketPunkBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoPunksMarketPunkBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoPunksMarketPunkBought represents a PunkBought event raised by the CryptoPunksMarket contract.
type CryptoPunksMarketPunkBought struct {
	PunkIndex   *big.Int
	Value       *big.Int
	FromAddress common.Address
	ToAddress   common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPunkBought is a free log retrieval operation binding the contract event 0x58e5d5a525e3b40bc15abaa38b5882678db1ee68befd2f60bafe3a7fd06db9e3.
//
// Solidity: event PunkBought(uint256 indexed punkIndex, uint256 value, address indexed fromAddress, address indexed toAddress)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) FilterPunkBought(opts *bind.FilterOpts, punkIndex []*big.Int, fromAddress []common.Address, toAddress []common.Address) (*CryptoPunksMarketPunkBoughtIterator, error) {

	var punkIndexRule []interface{}
	for _, punkIndexItem := range punkIndex {
		punkIndexRule = append(punkIndexRule, punkIndexItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.FilterLogs(opts, "PunkBought", punkIndexRule, fromAddressRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return &CryptoPunksMarketPunkBoughtIterator{contract: _CryptoPunksMarket.contract, event: "PunkBought", logs: logs, sub: sub}, nil
}

// WatchPunkBought is a free log subscription operation binding the contract event 0x58e5d5a525e3b40bc15abaa38b5882678db1ee68befd2f60bafe3a7fd06db9e3.
//
// Solidity: event PunkBought(uint256 indexed punkIndex, uint256 value, address indexed fromAddress, address indexed toAddress)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) WatchPunkBought(opts *bind.WatchOpts, sink chan<- *CryptoPunksMarketPunkBought, punkIndex []*big.Int, fromAddress []common.Address, toAddress []common.Address) (event.Subscription, error) {

	var punkIndexRule []interface{}
	for _, punkIndexItem := range punkIndex {
		punkIndexRule = append(punkIndexRule, punkIndexItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.WatchLogs(opts, "PunkBought", punkIndexRule, fromAddressRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoPunksMarketPunkBought)
				if err := _CryptoPunksMarket.contract.UnpackLog(event, "PunkBought", log); err != nil {
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

// ParsePunkBought is a log parse operation binding the contract event 0x58e5d5a525e3b40bc15abaa38b5882678db1ee68befd2f60bafe3a7fd06db9e3.
//
// Solidity: event PunkBought(uint256 indexed punkIndex, uint256 value, address indexed fromAddress, address indexed toAddress)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) ParsePunkBought(log types.Log) (*CryptoPunksMarketPunkBought, error) {
	event := new(CryptoPunksMarketPunkBought)
	if err := _CryptoPunksMarket.contract.UnpackLog(event, "PunkBought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoPunksMarketPunkNoLongerForSaleIterator is returned from FilterPunkNoLongerForSale and is used to iterate over the raw logs and unpacked data for PunkNoLongerForSale events raised by the CryptoPunksMarket contract.
type CryptoPunksMarketPunkNoLongerForSaleIterator struct {
	Event *CryptoPunksMarketPunkNoLongerForSale // Event containing the contract specifics and raw log

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
func (it *CryptoPunksMarketPunkNoLongerForSaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoPunksMarketPunkNoLongerForSale)
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
		it.Event = new(CryptoPunksMarketPunkNoLongerForSale)
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
func (it *CryptoPunksMarketPunkNoLongerForSaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoPunksMarketPunkNoLongerForSaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoPunksMarketPunkNoLongerForSale represents a PunkNoLongerForSale event raised by the CryptoPunksMarket contract.
type CryptoPunksMarketPunkNoLongerForSale struct {
	PunkIndex *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPunkNoLongerForSale is a free log retrieval operation binding the contract event 0xb0e0a660b4e50f26f0b7ce75c24655fc76cc66e3334a54ff410277229fa10bd4.
//
// Solidity: event PunkNoLongerForSale(uint256 indexed punkIndex)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) FilterPunkNoLongerForSale(opts *bind.FilterOpts, punkIndex []*big.Int) (*CryptoPunksMarketPunkNoLongerForSaleIterator, error) {

	var punkIndexRule []interface{}
	for _, punkIndexItem := range punkIndex {
		punkIndexRule = append(punkIndexRule, punkIndexItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.FilterLogs(opts, "PunkNoLongerForSale", punkIndexRule)
	if err != nil {
		return nil, err
	}
	return &CryptoPunksMarketPunkNoLongerForSaleIterator{contract: _CryptoPunksMarket.contract, event: "PunkNoLongerForSale", logs: logs, sub: sub}, nil
}

// WatchPunkNoLongerForSale is a free log subscription operation binding the contract event 0xb0e0a660b4e50f26f0b7ce75c24655fc76cc66e3334a54ff410277229fa10bd4.
//
// Solidity: event PunkNoLongerForSale(uint256 indexed punkIndex)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) WatchPunkNoLongerForSale(opts *bind.WatchOpts, sink chan<- *CryptoPunksMarketPunkNoLongerForSale, punkIndex []*big.Int) (event.Subscription, error) {

	var punkIndexRule []interface{}
	for _, punkIndexItem := range punkIndex {
		punkIndexRule = append(punkIndexRule, punkIndexItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.WatchLogs(opts, "PunkNoLongerForSale", punkIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoPunksMarketPunkNoLongerForSale)
				if err := _CryptoPunksMarket.contract.UnpackLog(event, "PunkNoLongerForSale", log); err != nil {
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

// ParsePunkNoLongerForSale is a log parse operation binding the contract event 0xb0e0a660b4e50f26f0b7ce75c24655fc76cc66e3334a54ff410277229fa10bd4.
//
// Solidity: event PunkNoLongerForSale(uint256 indexed punkIndex)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) ParsePunkNoLongerForSale(log types.Log) (*CryptoPunksMarketPunkNoLongerForSale, error) {
	event := new(CryptoPunksMarketPunkNoLongerForSale)
	if err := _CryptoPunksMarket.contract.UnpackLog(event, "PunkNoLongerForSale", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoPunksMarketPunkOfferedIterator is returned from FilterPunkOffered and is used to iterate over the raw logs and unpacked data for PunkOffered events raised by the CryptoPunksMarket contract.
type CryptoPunksMarketPunkOfferedIterator struct {
	Event *CryptoPunksMarketPunkOffered // Event containing the contract specifics and raw log

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
func (it *CryptoPunksMarketPunkOfferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoPunksMarketPunkOffered)
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
		it.Event = new(CryptoPunksMarketPunkOffered)
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
func (it *CryptoPunksMarketPunkOfferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoPunksMarketPunkOfferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoPunksMarketPunkOffered represents a PunkOffered event raised by the CryptoPunksMarket contract.
type CryptoPunksMarketPunkOffered struct {
	PunkIndex *big.Int
	MinValue  *big.Int
	ToAddress common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPunkOffered is a free log retrieval operation binding the contract event 0x3c7b682d5da98001a9b8cbda6c647d2c63d698a4184fd1d55e2ce7b66f5d21eb.
//
// Solidity: event PunkOffered(uint256 indexed punkIndex, uint256 minValue, address indexed toAddress)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) FilterPunkOffered(opts *bind.FilterOpts, punkIndex []*big.Int, toAddress []common.Address) (*CryptoPunksMarketPunkOfferedIterator, error) {

	var punkIndexRule []interface{}
	for _, punkIndexItem := range punkIndex {
		punkIndexRule = append(punkIndexRule, punkIndexItem)
	}

	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.FilterLogs(opts, "PunkOffered", punkIndexRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return &CryptoPunksMarketPunkOfferedIterator{contract: _CryptoPunksMarket.contract, event: "PunkOffered", logs: logs, sub: sub}, nil
}

// WatchPunkOffered is a free log subscription operation binding the contract event 0x3c7b682d5da98001a9b8cbda6c647d2c63d698a4184fd1d55e2ce7b66f5d21eb.
//
// Solidity: event PunkOffered(uint256 indexed punkIndex, uint256 minValue, address indexed toAddress)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) WatchPunkOffered(opts *bind.WatchOpts, sink chan<- *CryptoPunksMarketPunkOffered, punkIndex []*big.Int, toAddress []common.Address) (event.Subscription, error) {

	var punkIndexRule []interface{}
	for _, punkIndexItem := range punkIndex {
		punkIndexRule = append(punkIndexRule, punkIndexItem)
	}

	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.WatchLogs(opts, "PunkOffered", punkIndexRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoPunksMarketPunkOffered)
				if err := _CryptoPunksMarket.contract.UnpackLog(event, "PunkOffered", log); err != nil {
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

// ParsePunkOffered is a log parse operation binding the contract event 0x3c7b682d5da98001a9b8cbda6c647d2c63d698a4184fd1d55e2ce7b66f5d21eb.
//
// Solidity: event PunkOffered(uint256 indexed punkIndex, uint256 minValue, address indexed toAddress)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) ParsePunkOffered(log types.Log) (*CryptoPunksMarketPunkOffered, error) {
	event := new(CryptoPunksMarketPunkOffered)
	if err := _CryptoPunksMarket.contract.UnpackLog(event, "PunkOffered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoPunksMarketPunkTransferIterator is returned from FilterPunkTransfer and is used to iterate over the raw logs and unpacked data for PunkTransfer events raised by the CryptoPunksMarket contract.
type CryptoPunksMarketPunkTransferIterator struct {
	Event *CryptoPunksMarketPunkTransfer // Event containing the contract specifics and raw log

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
func (it *CryptoPunksMarketPunkTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoPunksMarketPunkTransfer)
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
		it.Event = new(CryptoPunksMarketPunkTransfer)
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
func (it *CryptoPunksMarketPunkTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoPunksMarketPunkTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoPunksMarketPunkTransfer represents a PunkTransfer event raised by the CryptoPunksMarket contract.
type CryptoPunksMarketPunkTransfer struct {
	From      common.Address
	To        common.Address
	PunkIndex *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPunkTransfer is a free log retrieval operation binding the contract event 0x05af636b70da6819000c49f85b21fa82081c632069bb626f30932034099107d8.
//
// Solidity: event PunkTransfer(address indexed from, address indexed to, uint256 punkIndex)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) FilterPunkTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CryptoPunksMarketPunkTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.FilterLogs(opts, "PunkTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CryptoPunksMarketPunkTransferIterator{contract: _CryptoPunksMarket.contract, event: "PunkTransfer", logs: logs, sub: sub}, nil
}

// WatchPunkTransfer is a free log subscription operation binding the contract event 0x05af636b70da6819000c49f85b21fa82081c632069bb626f30932034099107d8.
//
// Solidity: event PunkTransfer(address indexed from, address indexed to, uint256 punkIndex)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) WatchPunkTransfer(opts *bind.WatchOpts, sink chan<- *CryptoPunksMarketPunkTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.WatchLogs(opts, "PunkTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoPunksMarketPunkTransfer)
				if err := _CryptoPunksMarket.contract.UnpackLog(event, "PunkTransfer", log); err != nil {
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

// ParsePunkTransfer is a log parse operation binding the contract event 0x05af636b70da6819000c49f85b21fa82081c632069bb626f30932034099107d8.
//
// Solidity: event PunkTransfer(address indexed from, address indexed to, uint256 punkIndex)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) ParsePunkTransfer(log types.Log) (*CryptoPunksMarketPunkTransfer, error) {
	event := new(CryptoPunksMarketPunkTransfer)
	if err := _CryptoPunksMarket.contract.UnpackLog(event, "PunkTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoPunksMarketTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CryptoPunksMarket contract.
type CryptoPunksMarketTransferIterator struct {
	Event *CryptoPunksMarketTransfer // Event containing the contract specifics and raw log

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
func (it *CryptoPunksMarketTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoPunksMarketTransfer)
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
		it.Event = new(CryptoPunksMarketTransfer)
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
func (it *CryptoPunksMarketTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoPunksMarketTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoPunksMarketTransfer represents a Transfer event raised by the CryptoPunksMarket contract.
type CryptoPunksMarketTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CryptoPunksMarketTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CryptoPunksMarketTransferIterator{contract: _CryptoPunksMarket.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CryptoPunksMarketTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoPunksMarketTransfer)
				if err := _CryptoPunksMarket.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) ParseTransfer(log types.Log) (*CryptoPunksMarketTransfer, error) {
	event := new(CryptoPunksMarketTransfer)
	if err := _CryptoPunksMarket.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
