// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

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

// StoreABI is the input ABI used to generate the binding from.
const StoreABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"punksOfferedForSale\",\"outputs\":[{\"name\":\"isForSale\",\"type\":\"bool\"},{\"name\":\"punkIndex\",\"type\":\"uint256\"},{\"name\":\"seller\",\"type\":\"address\"},{\"name\":\"minValue\",\"type\":\"uint256\"},{\"name\":\"onlySellTo\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"enterBidForPunk\",\"outputs\":[],\"payable\":true,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"punkIndex\",\"type\":\"uint256\"},{\"name\":\"minPrice\",\"type\":\"uint256\"}],\"name\":\"acceptBidForPunk\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addresses\",\"type\":\"address[]\"},{\"name\":\"indices\",\"type\":\"uint256[]\"}],\"name\":\"setInitialOwners\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"imageHash\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextPunkIndexToAssign\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"punkIndexToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"standard\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"punkBids\",\"outputs\":[{\"name\":\"hasBid\",\"type\":\"bool\"},{\"name\":\"punkIndex\",\"type\":\"uint256\"},{\"name\":\"bidder\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"allInitialOwnersAssigned\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allPunksAssigned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"buyPunk\",\"outputs\":[],\"payable\":true,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"transferPunk\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"withdrawBidForPunk\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"setInitialOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"punkIndex\",\"type\":\"uint256\"},{\"name\":\"minSalePriceInWei\",\"type\":\"uint256\"},{\"name\":\"toAddress\",\"type\":\"address\"}],\"name\":\"offerPunkForSaleToAddress\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"punksRemainingToAssign\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"punkIndex\",\"type\":\"uint256\"},{\"name\":\"minSalePriceInWei\",\"type\":\"uint256\"}],\"name\":\"offerPunkForSale\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"getPunk\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingWithdrawals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"punkNoLongerForSale\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"Assign\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"PunkTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"punkIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"minValue\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"toAddress\",\"type\":\"address\"}],\"name\":\"PunkOffered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"punkIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"fromAddress\",\"type\":\"address\"}],\"name\":\"PunkBidEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"punkIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"fromAddress\",\"type\":\"address\"}],\"name\":\"PunkBidWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"punkIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"fromAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"toAddress\",\"type\":\"address\"}],\"name\":\"PunkBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"PunkNoLongerForSale\",\"type\":\"event\"}]"

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// AllPunksAssigned is a free data retrieval call binding the contract method 0x8126c38a.
//
// Solidity: function allPunksAssigned() returns(bool)
func (_Store *StoreCaller) AllPunksAssigned(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "allPunksAssigned")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllPunksAssigned is a free data retrieval call binding the contract method 0x8126c38a.
//
// Solidity: function allPunksAssigned() returns(bool)
func (_Store *StoreSession) AllPunksAssigned() (bool, error) {
	return _Store.Contract.AllPunksAssigned(&_Store.CallOpts)
}

// AllPunksAssigned is a free data retrieval call binding the contract method 0x8126c38a.
//
// Solidity: function allPunksAssigned() returns(bool)
func (_Store *StoreCallerSession) AllPunksAssigned() (bool, error) {
	return _Store.Contract.AllPunksAssigned(&_Store.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) returns(uint256)
func (_Store *StoreCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) returns(uint256)
func (_Store *StoreSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Store.Contract.BalanceOf(&_Store.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) returns(uint256)
func (_Store *StoreCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Store.Contract.BalanceOf(&_Store.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint8)
func (_Store *StoreCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint8)
func (_Store *StoreSession) Decimals() (uint8, error) {
	return _Store.Contract.Decimals(&_Store.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint8)
func (_Store *StoreCallerSession) Decimals() (uint8, error) {
	return _Store.Contract.Decimals(&_Store.CallOpts)
}

// ImageHash is a free data retrieval call binding the contract method 0x51605d80.
//
// Solidity: function imageHash() returns(string)
func (_Store *StoreCaller) ImageHash(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "imageHash")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ImageHash is a free data retrieval call binding the contract method 0x51605d80.
//
// Solidity: function imageHash() returns(string)
func (_Store *StoreSession) ImageHash() (string, error) {
	return _Store.Contract.ImageHash(&_Store.CallOpts)
}

// ImageHash is a free data retrieval call binding the contract method 0x51605d80.
//
// Solidity: function imageHash() returns(string)
func (_Store *StoreCallerSession) ImageHash() (string, error) {
	return _Store.Contract.ImageHash(&_Store.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_Store *StoreCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_Store *StoreSession) Name() (string, error) {
	return _Store.Contract.Name(&_Store.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_Store *StoreCallerSession) Name() (string, error) {
	return _Store.Contract.Name(&_Store.CallOpts)
}

// NextPunkIndexToAssign is a free data retrieval call binding the contract method 0x52f29a25.
//
// Solidity: function nextPunkIndexToAssign() returns(uint256)
func (_Store *StoreCaller) NextPunkIndexToAssign(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "nextPunkIndexToAssign")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextPunkIndexToAssign is a free data retrieval call binding the contract method 0x52f29a25.
//
// Solidity: function nextPunkIndexToAssign() returns(uint256)
func (_Store *StoreSession) NextPunkIndexToAssign() (*big.Int, error) {
	return _Store.Contract.NextPunkIndexToAssign(&_Store.CallOpts)
}

// NextPunkIndexToAssign is a free data retrieval call binding the contract method 0x52f29a25.
//
// Solidity: function nextPunkIndexToAssign() returns(uint256)
func (_Store *StoreCallerSession) NextPunkIndexToAssign() (*big.Int, error) {
	return _Store.Contract.NextPunkIndexToAssign(&_Store.CallOpts)
}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xf3f43703.
//
// Solidity: function pendingWithdrawals(address ) returns(uint256)
func (_Store *StoreCaller) PendingWithdrawals(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "pendingWithdrawals", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xf3f43703.
//
// Solidity: function pendingWithdrawals(address ) returns(uint256)
func (_Store *StoreSession) PendingWithdrawals(arg0 common.Address) (*big.Int, error) {
	return _Store.Contract.PendingWithdrawals(&_Store.CallOpts, arg0)
}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xf3f43703.
//
// Solidity: function pendingWithdrawals(address ) returns(uint256)
func (_Store *StoreCallerSession) PendingWithdrawals(arg0 common.Address) (*big.Int, error) {
	return _Store.Contract.PendingWithdrawals(&_Store.CallOpts, arg0)
}

// PunkBids is a free data retrieval call binding the contract method 0x6e743fa9.
//
// Solidity: function punkBids(uint256 ) returns(bool hasBid, uint256 punkIndex, address bidder, uint256 value)
func (_Store *StoreCaller) PunkBids(opts *bind.CallOpts, arg0 *big.Int) (struct {
	HasBid    bool
	PunkIndex *big.Int
	Bidder    common.Address
	Value     *big.Int
}, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "punkBids", arg0)

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
func (_Store *StoreSession) PunkBids(arg0 *big.Int) (struct {
	HasBid    bool
	PunkIndex *big.Int
	Bidder    common.Address
	Value     *big.Int
}, error) {
	return _Store.Contract.PunkBids(&_Store.CallOpts, arg0)
}

// PunkBids is a free data retrieval call binding the contract method 0x6e743fa9.
//
// Solidity: function punkBids(uint256 ) returns(bool hasBid, uint256 punkIndex, address bidder, uint256 value)
func (_Store *StoreCallerSession) PunkBids(arg0 *big.Int) (struct {
	HasBid    bool
	PunkIndex *big.Int
	Bidder    common.Address
	Value     *big.Int
}, error) {
	return _Store.Contract.PunkBids(&_Store.CallOpts, arg0)
}

// PunkIndexToAddress is a free data retrieval call binding the contract method 0x58178168.
//
// Solidity: function punkIndexToAddress(uint256 ) returns(address)
func (_Store *StoreCaller) PunkIndexToAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "punkIndexToAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PunkIndexToAddress is a free data retrieval call binding the contract method 0x58178168.
//
// Solidity: function punkIndexToAddress(uint256 ) returns(address)
func (_Store *StoreSession) PunkIndexToAddress(arg0 *big.Int) (common.Address, error) {
	return _Store.Contract.PunkIndexToAddress(&_Store.CallOpts, arg0)
}

// PunkIndexToAddress is a free data retrieval call binding the contract method 0x58178168.
//
// Solidity: function punkIndexToAddress(uint256 ) returns(address)
func (_Store *StoreCallerSession) PunkIndexToAddress(arg0 *big.Int) (common.Address, error) {
	return _Store.Contract.PunkIndexToAddress(&_Store.CallOpts, arg0)
}

// PunksOfferedForSale is a free data retrieval call binding the contract method 0x088f11f3.
//
// Solidity: function punksOfferedForSale(uint256 ) returns(bool isForSale, uint256 punkIndex, address seller, uint256 minValue, address onlySellTo)
func (_Store *StoreCaller) PunksOfferedForSale(opts *bind.CallOpts, arg0 *big.Int) (struct {
	IsForSale  bool
	PunkIndex  *big.Int
	Seller     common.Address
	MinValue   *big.Int
	OnlySellTo common.Address
}, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "punksOfferedForSale", arg0)

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
func (_Store *StoreSession) PunksOfferedForSale(arg0 *big.Int) (struct {
	IsForSale  bool
	PunkIndex  *big.Int
	Seller     common.Address
	MinValue   *big.Int
	OnlySellTo common.Address
}, error) {
	return _Store.Contract.PunksOfferedForSale(&_Store.CallOpts, arg0)
}

// PunksOfferedForSale is a free data retrieval call binding the contract method 0x088f11f3.
//
// Solidity: function punksOfferedForSale(uint256 ) returns(bool isForSale, uint256 punkIndex, address seller, uint256 minValue, address onlySellTo)
func (_Store *StoreCallerSession) PunksOfferedForSale(arg0 *big.Int) (struct {
	IsForSale  bool
	PunkIndex  *big.Int
	Seller     common.Address
	MinValue   *big.Int
	OnlySellTo common.Address
}, error) {
	return _Store.Contract.PunksOfferedForSale(&_Store.CallOpts, arg0)
}

// PunksRemainingToAssign is a free data retrieval call binding the contract method 0xc0d6ce63.
//
// Solidity: function punksRemainingToAssign() returns(uint256)
func (_Store *StoreCaller) PunksRemainingToAssign(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "punksRemainingToAssign")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PunksRemainingToAssign is a free data retrieval call binding the contract method 0xc0d6ce63.
//
// Solidity: function punksRemainingToAssign() returns(uint256)
func (_Store *StoreSession) PunksRemainingToAssign() (*big.Int, error) {
	return _Store.Contract.PunksRemainingToAssign(&_Store.CallOpts)
}

// PunksRemainingToAssign is a free data retrieval call binding the contract method 0xc0d6ce63.
//
// Solidity: function punksRemainingToAssign() returns(uint256)
func (_Store *StoreCallerSession) PunksRemainingToAssign() (*big.Int, error) {
	return _Store.Contract.PunksRemainingToAssign(&_Store.CallOpts)
}

// Standard is a free data retrieval call binding the contract method 0x5a3b7e42.
//
// Solidity: function standard() returns(string)
func (_Store *StoreCaller) Standard(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "standard")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Standard is a free data retrieval call binding the contract method 0x5a3b7e42.
//
// Solidity: function standard() returns(string)
func (_Store *StoreSession) Standard() (string, error) {
	return _Store.Contract.Standard(&_Store.CallOpts)
}

// Standard is a free data retrieval call binding the contract method 0x5a3b7e42.
//
// Solidity: function standard() returns(string)
func (_Store *StoreCallerSession) Standard() (string, error) {
	return _Store.Contract.Standard(&_Store.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_Store *StoreCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_Store *StoreSession) Symbol() (string, error) {
	return _Store.Contract.Symbol(&_Store.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_Store *StoreCallerSession) Symbol() (string, error) {
	return _Store.Contract.Symbol(&_Store.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_Store *StoreCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_Store *StoreSession) TotalSupply() (*big.Int, error) {
	return _Store.Contract.TotalSupply(&_Store.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_Store *StoreCallerSession) TotalSupply() (*big.Int, error) {
	return _Store.Contract.TotalSupply(&_Store.CallOpts)
}

// AcceptBidForPunk is a paid mutator transaction binding the contract method 0x23165b75.
//
// Solidity: function acceptBidForPunk(uint256 punkIndex, uint256 minPrice) returns()
func (_Store *StoreTransactor) AcceptBidForPunk(opts *bind.TransactOpts, punkIndex *big.Int, minPrice *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "acceptBidForPunk", punkIndex, minPrice)
}

// AcceptBidForPunk is a paid mutator transaction binding the contract method 0x23165b75.
//
// Solidity: function acceptBidForPunk(uint256 punkIndex, uint256 minPrice) returns()
func (_Store *StoreSession) AcceptBidForPunk(punkIndex *big.Int, minPrice *big.Int) (*types.Transaction, error) {
	return _Store.Contract.AcceptBidForPunk(&_Store.TransactOpts, punkIndex, minPrice)
}

// AcceptBidForPunk is a paid mutator transaction binding the contract method 0x23165b75.
//
// Solidity: function acceptBidForPunk(uint256 punkIndex, uint256 minPrice) returns()
func (_Store *StoreTransactorSession) AcceptBidForPunk(punkIndex *big.Int, minPrice *big.Int) (*types.Transaction, error) {
	return _Store.Contract.AcceptBidForPunk(&_Store.TransactOpts, punkIndex, minPrice)
}

// AllInitialOwnersAssigned is a paid mutator transaction binding the contract method 0x7ecedac9.
//
// Solidity: function allInitialOwnersAssigned() returns()
func (_Store *StoreTransactor) AllInitialOwnersAssigned(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "allInitialOwnersAssigned")
}

// AllInitialOwnersAssigned is a paid mutator transaction binding the contract method 0x7ecedac9.
//
// Solidity: function allInitialOwnersAssigned() returns()
func (_Store *StoreSession) AllInitialOwnersAssigned() (*types.Transaction, error) {
	return _Store.Contract.AllInitialOwnersAssigned(&_Store.TransactOpts)
}

// AllInitialOwnersAssigned is a paid mutator transaction binding the contract method 0x7ecedac9.
//
// Solidity: function allInitialOwnersAssigned() returns()
func (_Store *StoreTransactorSession) AllInitialOwnersAssigned() (*types.Transaction, error) {
	return _Store.Contract.AllInitialOwnersAssigned(&_Store.TransactOpts)
}

// BuyPunk is a paid mutator transaction binding the contract method 0x8264fe98.
//
// Solidity: function buyPunk(uint256 punkIndex) returns()
func (_Store *StoreTransactor) BuyPunk(opts *bind.TransactOpts, punkIndex *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "buyPunk", punkIndex)
}

// BuyPunk is a paid mutator transaction binding the contract method 0x8264fe98.
//
// Solidity: function buyPunk(uint256 punkIndex) returns()
func (_Store *StoreSession) BuyPunk(punkIndex *big.Int) (*types.Transaction, error) {
	return _Store.Contract.BuyPunk(&_Store.TransactOpts, punkIndex)
}

// BuyPunk is a paid mutator transaction binding the contract method 0x8264fe98.
//
// Solidity: function buyPunk(uint256 punkIndex) returns()
func (_Store *StoreTransactorSession) BuyPunk(punkIndex *big.Int) (*types.Transaction, error) {
	return _Store.Contract.BuyPunk(&_Store.TransactOpts, punkIndex)
}

// EnterBidForPunk is a paid mutator transaction binding the contract method 0x091dbfd2.
//
// Solidity: function enterBidForPunk(uint256 punkIndex) returns()
func (_Store *StoreTransactor) EnterBidForPunk(opts *bind.TransactOpts, punkIndex *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "enterBidForPunk", punkIndex)
}

// EnterBidForPunk is a paid mutator transaction binding the contract method 0x091dbfd2.
//
// Solidity: function enterBidForPunk(uint256 punkIndex) returns()
func (_Store *StoreSession) EnterBidForPunk(punkIndex *big.Int) (*types.Transaction, error) {
	return _Store.Contract.EnterBidForPunk(&_Store.TransactOpts, punkIndex)
}

// EnterBidForPunk is a paid mutator transaction binding the contract method 0x091dbfd2.
//
// Solidity: function enterBidForPunk(uint256 punkIndex) returns()
func (_Store *StoreTransactorSession) EnterBidForPunk(punkIndex *big.Int) (*types.Transaction, error) {
	return _Store.Contract.EnterBidForPunk(&_Store.TransactOpts, punkIndex)
}

// GetPunk is a paid mutator transaction binding the contract method 0xc81d1d5b.
//
// Solidity: function getPunk(uint256 punkIndex) returns()
func (_Store *StoreTransactor) GetPunk(opts *bind.TransactOpts, punkIndex *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "getPunk", punkIndex)
}

// GetPunk is a paid mutator transaction binding the contract method 0xc81d1d5b.
//
// Solidity: function getPunk(uint256 punkIndex) returns()
func (_Store *StoreSession) GetPunk(punkIndex *big.Int) (*types.Transaction, error) {
	return _Store.Contract.GetPunk(&_Store.TransactOpts, punkIndex)
}

// GetPunk is a paid mutator transaction binding the contract method 0xc81d1d5b.
//
// Solidity: function getPunk(uint256 punkIndex) returns()
func (_Store *StoreTransactorSession) GetPunk(punkIndex *big.Int) (*types.Transaction, error) {
	return _Store.Contract.GetPunk(&_Store.TransactOpts, punkIndex)
}

// OfferPunkForSale is a paid mutator transaction binding the contract method 0xc44193c3.
//
// Solidity: function offerPunkForSale(uint256 punkIndex, uint256 minSalePriceInWei) returns()
func (_Store *StoreTransactor) OfferPunkForSale(opts *bind.TransactOpts, punkIndex *big.Int, minSalePriceInWei *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "offerPunkForSale", punkIndex, minSalePriceInWei)
}

// OfferPunkForSale is a paid mutator transaction binding the contract method 0xc44193c3.
//
// Solidity: function offerPunkForSale(uint256 punkIndex, uint256 minSalePriceInWei) returns()
func (_Store *StoreSession) OfferPunkForSale(punkIndex *big.Int, minSalePriceInWei *big.Int) (*types.Transaction, error) {
	return _Store.Contract.OfferPunkForSale(&_Store.TransactOpts, punkIndex, minSalePriceInWei)
}

// OfferPunkForSale is a paid mutator transaction binding the contract method 0xc44193c3.
//
// Solidity: function offerPunkForSale(uint256 punkIndex, uint256 minSalePriceInWei) returns()
func (_Store *StoreTransactorSession) OfferPunkForSale(punkIndex *big.Int, minSalePriceInWei *big.Int) (*types.Transaction, error) {
	return _Store.Contract.OfferPunkForSale(&_Store.TransactOpts, punkIndex, minSalePriceInWei)
}

// OfferPunkForSaleToAddress is a paid mutator transaction binding the contract method 0xbf31196f.
//
// Solidity: function offerPunkForSaleToAddress(uint256 punkIndex, uint256 minSalePriceInWei, address toAddress) returns()
func (_Store *StoreTransactor) OfferPunkForSaleToAddress(opts *bind.TransactOpts, punkIndex *big.Int, minSalePriceInWei *big.Int, toAddress common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "offerPunkForSaleToAddress", punkIndex, minSalePriceInWei, toAddress)
}

// OfferPunkForSaleToAddress is a paid mutator transaction binding the contract method 0xbf31196f.
//
// Solidity: function offerPunkForSaleToAddress(uint256 punkIndex, uint256 minSalePriceInWei, address toAddress) returns()
func (_Store *StoreSession) OfferPunkForSaleToAddress(punkIndex *big.Int, minSalePriceInWei *big.Int, toAddress common.Address) (*types.Transaction, error) {
	return _Store.Contract.OfferPunkForSaleToAddress(&_Store.TransactOpts, punkIndex, minSalePriceInWei, toAddress)
}

// OfferPunkForSaleToAddress is a paid mutator transaction binding the contract method 0xbf31196f.
//
// Solidity: function offerPunkForSaleToAddress(uint256 punkIndex, uint256 minSalePriceInWei, address toAddress) returns()
func (_Store *StoreTransactorSession) OfferPunkForSaleToAddress(punkIndex *big.Int, minSalePriceInWei *big.Int, toAddress common.Address) (*types.Transaction, error) {
	return _Store.Contract.OfferPunkForSaleToAddress(&_Store.TransactOpts, punkIndex, minSalePriceInWei, toAddress)
}

// PunkNoLongerForSale is a paid mutator transaction binding the contract method 0xf6eeff1e.
//
// Solidity: function punkNoLongerForSale(uint256 punkIndex) returns()
func (_Store *StoreTransactor) PunkNoLongerForSale(opts *bind.TransactOpts, punkIndex *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "punkNoLongerForSale", punkIndex)
}

// PunkNoLongerForSale is a paid mutator transaction binding the contract method 0xf6eeff1e.
//
// Solidity: function punkNoLongerForSale(uint256 punkIndex) returns()
func (_Store *StoreSession) PunkNoLongerForSale(punkIndex *big.Int) (*types.Transaction, error) {
	return _Store.Contract.PunkNoLongerForSale(&_Store.TransactOpts, punkIndex)
}

// PunkNoLongerForSale is a paid mutator transaction binding the contract method 0xf6eeff1e.
//
// Solidity: function punkNoLongerForSale(uint256 punkIndex) returns()
func (_Store *StoreTransactorSession) PunkNoLongerForSale(punkIndex *big.Int) (*types.Transaction, error) {
	return _Store.Contract.PunkNoLongerForSale(&_Store.TransactOpts, punkIndex)
}

// SetInitialOwner is a paid mutator transaction binding the contract method 0xa75a9049.
//
// Solidity: function setInitialOwner(address to, uint256 punkIndex) returns()
func (_Store *StoreTransactor) SetInitialOwner(opts *bind.TransactOpts, to common.Address, punkIndex *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setInitialOwner", to, punkIndex)
}

// SetInitialOwner is a paid mutator transaction binding the contract method 0xa75a9049.
//
// Solidity: function setInitialOwner(address to, uint256 punkIndex) returns()
func (_Store *StoreSession) SetInitialOwner(to common.Address, punkIndex *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetInitialOwner(&_Store.TransactOpts, to, punkIndex)
}

// SetInitialOwner is a paid mutator transaction binding the contract method 0xa75a9049.
//
// Solidity: function setInitialOwner(address to, uint256 punkIndex) returns()
func (_Store *StoreTransactorSession) SetInitialOwner(to common.Address, punkIndex *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetInitialOwner(&_Store.TransactOpts, to, punkIndex)
}

// SetInitialOwners is a paid mutator transaction binding the contract method 0x39c5dde6.
//
// Solidity: function setInitialOwners(address[] addresses, uint256[] indices) returns()
func (_Store *StoreTransactor) SetInitialOwners(opts *bind.TransactOpts, addresses []common.Address, indices []*big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setInitialOwners", addresses, indices)
}

// SetInitialOwners is a paid mutator transaction binding the contract method 0x39c5dde6.
//
// Solidity: function setInitialOwners(address[] addresses, uint256[] indices) returns()
func (_Store *StoreSession) SetInitialOwners(addresses []common.Address, indices []*big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetInitialOwners(&_Store.TransactOpts, addresses, indices)
}

// SetInitialOwners is a paid mutator transaction binding the contract method 0x39c5dde6.
//
// Solidity: function setInitialOwners(address[] addresses, uint256[] indices) returns()
func (_Store *StoreTransactorSession) SetInitialOwners(addresses []common.Address, indices []*big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetInitialOwners(&_Store.TransactOpts, addresses, indices)
}

// TransferPunk is a paid mutator transaction binding the contract method 0x8b72a2ec.
//
// Solidity: function transferPunk(address to, uint256 punkIndex) returns()
func (_Store *StoreTransactor) TransferPunk(opts *bind.TransactOpts, to common.Address, punkIndex *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "transferPunk", to, punkIndex)
}

// TransferPunk is a paid mutator transaction binding the contract method 0x8b72a2ec.
//
// Solidity: function transferPunk(address to, uint256 punkIndex) returns()
func (_Store *StoreSession) TransferPunk(to common.Address, punkIndex *big.Int) (*types.Transaction, error) {
	return _Store.Contract.TransferPunk(&_Store.TransactOpts, to, punkIndex)
}

// TransferPunk is a paid mutator transaction binding the contract method 0x8b72a2ec.
//
// Solidity: function transferPunk(address to, uint256 punkIndex) returns()
func (_Store *StoreTransactorSession) TransferPunk(to common.Address, punkIndex *big.Int) (*types.Transaction, error) {
	return _Store.Contract.TransferPunk(&_Store.TransactOpts, to, punkIndex)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Store *StoreTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Store *StoreSession) Withdraw() (*types.Transaction, error) {
	return _Store.Contract.Withdraw(&_Store.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Store *StoreTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Store.Contract.Withdraw(&_Store.TransactOpts)
}

// WithdrawBidForPunk is a paid mutator transaction binding the contract method 0x979bc638.
//
// Solidity: function withdrawBidForPunk(uint256 punkIndex) returns()
func (_Store *StoreTransactor) WithdrawBidForPunk(opts *bind.TransactOpts, punkIndex *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "withdrawBidForPunk", punkIndex)
}

// WithdrawBidForPunk is a paid mutator transaction binding the contract method 0x979bc638.
//
// Solidity: function withdrawBidForPunk(uint256 punkIndex) returns()
func (_Store *StoreSession) WithdrawBidForPunk(punkIndex *big.Int) (*types.Transaction, error) {
	return _Store.Contract.WithdrawBidForPunk(&_Store.TransactOpts, punkIndex)
}

// WithdrawBidForPunk is a paid mutator transaction binding the contract method 0x979bc638.
//
// Solidity: function withdrawBidForPunk(uint256 punkIndex) returns()
func (_Store *StoreTransactorSession) WithdrawBidForPunk(punkIndex *big.Int) (*types.Transaction, error) {
	return _Store.Contract.WithdrawBidForPunk(&_Store.TransactOpts, punkIndex)
}

// StoreAssignIterator is returned from FilterAssign and is used to iterate over the raw logs and unpacked data for Assign events raised by the Store contract.
type StoreAssignIterator struct {
	Event *StoreAssign // Event containing the contract specifics and raw log

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
func (it *StoreAssignIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreAssign)
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
		it.Event = new(StoreAssign)
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
func (it *StoreAssignIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreAssignIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreAssign represents a Assign event raised by the Store contract.
type StoreAssign struct {
	To        common.Address
	PunkIndex *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAssign is a free log retrieval operation binding the contract event 0x8a0e37b73a0d9c82e205d4d1a3ff3d0b57ce5f4d7bccf6bac03336dc101cb7ba.
//
// Solidity: event Assign(address indexed to, uint256 punkIndex)
func (_Store *StoreFilterer) FilterAssign(opts *bind.FilterOpts, to []common.Address) (*StoreAssignIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "Assign", toRule)
	if err != nil {
		return nil, err
	}
	return &StoreAssignIterator{contract: _Store.contract, event: "Assign", logs: logs, sub: sub}, nil
}

// WatchAssign is a free log subscription operation binding the contract event 0x8a0e37b73a0d9c82e205d4d1a3ff3d0b57ce5f4d7bccf6bac03336dc101cb7ba.
//
// Solidity: event Assign(address indexed to, uint256 punkIndex)
func (_Store *StoreFilterer) WatchAssign(opts *bind.WatchOpts, sink chan<- *StoreAssign, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "Assign", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreAssign)
				if err := _Store.contract.UnpackLog(event, "Assign", log); err != nil {
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
func (_Store *StoreFilterer) ParseAssign(log types.Log) (*StoreAssign, error) {
	event := new(StoreAssign)
	if err := _Store.contract.UnpackLog(event, "Assign", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorePunkBidEnteredIterator is returned from FilterPunkBidEntered and is used to iterate over the raw logs and unpacked data for PunkBidEntered events raised by the Store contract.
type StorePunkBidEnteredIterator struct {
	Event *StorePunkBidEntered // Event containing the contract specifics and raw log

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
func (it *StorePunkBidEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorePunkBidEntered)
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
		it.Event = new(StorePunkBidEntered)
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
func (it *StorePunkBidEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorePunkBidEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorePunkBidEntered represents a PunkBidEntered event raised by the Store contract.
type StorePunkBidEntered struct {
	PunkIndex   *big.Int
	Value       *big.Int
	FromAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPunkBidEntered is a free log retrieval operation binding the contract event 0x5b859394fabae0c1ba88baffe67e751ab5248d2e879028b8c8d6897b0519f56a.
//
// Solidity: event PunkBidEntered(uint256 indexed punkIndex, uint256 value, address indexed fromAddress)
func (_Store *StoreFilterer) FilterPunkBidEntered(opts *bind.FilterOpts, punkIndex []*big.Int, fromAddress []common.Address) (*StorePunkBidEnteredIterator, error) {

	var punkIndexRule []interface{}
	for _, punkIndexItem := range punkIndex {
		punkIndexRule = append(punkIndexRule, punkIndexItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "PunkBidEntered", punkIndexRule, fromAddressRule)
	if err != nil {
		return nil, err
	}
	return &StorePunkBidEnteredIterator{contract: _Store.contract, event: "PunkBidEntered", logs: logs, sub: sub}, nil
}

// WatchPunkBidEntered is a free log subscription operation binding the contract event 0x5b859394fabae0c1ba88baffe67e751ab5248d2e879028b8c8d6897b0519f56a.
//
// Solidity: event PunkBidEntered(uint256 indexed punkIndex, uint256 value, address indexed fromAddress)
func (_Store *StoreFilterer) WatchPunkBidEntered(opts *bind.WatchOpts, sink chan<- *StorePunkBidEntered, punkIndex []*big.Int, fromAddress []common.Address) (event.Subscription, error) {

	var punkIndexRule []interface{}
	for _, punkIndexItem := range punkIndex {
		punkIndexRule = append(punkIndexRule, punkIndexItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "PunkBidEntered", punkIndexRule, fromAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorePunkBidEntered)
				if err := _Store.contract.UnpackLog(event, "PunkBidEntered", log); err != nil {
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
func (_Store *StoreFilterer) ParsePunkBidEntered(log types.Log) (*StorePunkBidEntered, error) {
	event := new(StorePunkBidEntered)
	if err := _Store.contract.UnpackLog(event, "PunkBidEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorePunkBidWithdrawnIterator is returned from FilterPunkBidWithdrawn and is used to iterate over the raw logs and unpacked data for PunkBidWithdrawn events raised by the Store contract.
type StorePunkBidWithdrawnIterator struct {
	Event *StorePunkBidWithdrawn // Event containing the contract specifics and raw log

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
func (it *StorePunkBidWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorePunkBidWithdrawn)
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
		it.Event = new(StorePunkBidWithdrawn)
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
func (it *StorePunkBidWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorePunkBidWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorePunkBidWithdrawn represents a PunkBidWithdrawn event raised by the Store contract.
type StorePunkBidWithdrawn struct {
	PunkIndex   *big.Int
	Value       *big.Int
	FromAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPunkBidWithdrawn is a free log retrieval operation binding the contract event 0x6f30e1ee4d81dcc7a8a478577f65d2ed2edb120565960ac45fe7c50551c87932.
//
// Solidity: event PunkBidWithdrawn(uint256 indexed punkIndex, uint256 value, address indexed fromAddress)
func (_Store *StoreFilterer) FilterPunkBidWithdrawn(opts *bind.FilterOpts, punkIndex []*big.Int, fromAddress []common.Address) (*StorePunkBidWithdrawnIterator, error) {

	var punkIndexRule []interface{}
	for _, punkIndexItem := range punkIndex {
		punkIndexRule = append(punkIndexRule, punkIndexItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "PunkBidWithdrawn", punkIndexRule, fromAddressRule)
	if err != nil {
		return nil, err
	}
	return &StorePunkBidWithdrawnIterator{contract: _Store.contract, event: "PunkBidWithdrawn", logs: logs, sub: sub}, nil
}

// WatchPunkBidWithdrawn is a free log subscription operation binding the contract event 0x6f30e1ee4d81dcc7a8a478577f65d2ed2edb120565960ac45fe7c50551c87932.
//
// Solidity: event PunkBidWithdrawn(uint256 indexed punkIndex, uint256 value, address indexed fromAddress)
func (_Store *StoreFilterer) WatchPunkBidWithdrawn(opts *bind.WatchOpts, sink chan<- *StorePunkBidWithdrawn, punkIndex []*big.Int, fromAddress []common.Address) (event.Subscription, error) {

	var punkIndexRule []interface{}
	for _, punkIndexItem := range punkIndex {
		punkIndexRule = append(punkIndexRule, punkIndexItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "PunkBidWithdrawn", punkIndexRule, fromAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorePunkBidWithdrawn)
				if err := _Store.contract.UnpackLog(event, "PunkBidWithdrawn", log); err != nil {
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
func (_Store *StoreFilterer) ParsePunkBidWithdrawn(log types.Log) (*StorePunkBidWithdrawn, error) {
	event := new(StorePunkBidWithdrawn)
	if err := _Store.contract.UnpackLog(event, "PunkBidWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorePunkBoughtIterator is returned from FilterPunkBought and is used to iterate over the raw logs and unpacked data for PunkBought events raised by the Store contract.
type StorePunkBoughtIterator struct {
	Event *StorePunkBought // Event containing the contract specifics and raw log

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
func (it *StorePunkBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorePunkBought)
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
		it.Event = new(StorePunkBought)
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
func (it *StorePunkBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorePunkBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorePunkBought represents a PunkBought event raised by the Store contract.
type StorePunkBought struct {
	PunkIndex   *big.Int
	Value       *big.Int
	FromAddress common.Address
	ToAddress   common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPunkBought is a free log retrieval operation binding the contract event 0x58e5d5a525e3b40bc15abaa38b5882678db1ee68befd2f60bafe3a7fd06db9e3.
//
// Solidity: event PunkBought(uint256 indexed punkIndex, uint256 value, address indexed fromAddress, address indexed toAddress)
func (_Store *StoreFilterer) FilterPunkBought(opts *bind.FilterOpts, punkIndex []*big.Int, fromAddress []common.Address, toAddress []common.Address) (*StorePunkBoughtIterator, error) {

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

	logs, sub, err := _Store.contract.FilterLogs(opts, "PunkBought", punkIndexRule, fromAddressRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return &StorePunkBoughtIterator{contract: _Store.contract, event: "PunkBought", logs: logs, sub: sub}, nil
}

// WatchPunkBought is a free log subscription operation binding the contract event 0x58e5d5a525e3b40bc15abaa38b5882678db1ee68befd2f60bafe3a7fd06db9e3.
//
// Solidity: event PunkBought(uint256 indexed punkIndex, uint256 value, address indexed fromAddress, address indexed toAddress)
func (_Store *StoreFilterer) WatchPunkBought(opts *bind.WatchOpts, sink chan<- *StorePunkBought, punkIndex []*big.Int, fromAddress []common.Address, toAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Store.contract.WatchLogs(opts, "PunkBought", punkIndexRule, fromAddressRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorePunkBought)
				if err := _Store.contract.UnpackLog(event, "PunkBought", log); err != nil {
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
func (_Store *StoreFilterer) ParsePunkBought(log types.Log) (*StorePunkBought, error) {
	event := new(StorePunkBought)
	if err := _Store.contract.UnpackLog(event, "PunkBought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorePunkNoLongerForSaleIterator is returned from FilterPunkNoLongerForSale and is used to iterate over the raw logs and unpacked data for PunkNoLongerForSale events raised by the Store contract.
type StorePunkNoLongerForSaleIterator struct {
	Event *StorePunkNoLongerForSale // Event containing the contract specifics and raw log

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
func (it *StorePunkNoLongerForSaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorePunkNoLongerForSale)
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
		it.Event = new(StorePunkNoLongerForSale)
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
func (it *StorePunkNoLongerForSaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorePunkNoLongerForSaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorePunkNoLongerForSale represents a PunkNoLongerForSale event raised by the Store contract.
type StorePunkNoLongerForSale struct {
	PunkIndex *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPunkNoLongerForSale is a free log retrieval operation binding the contract event 0xb0e0a660b4e50f26f0b7ce75c24655fc76cc66e3334a54ff410277229fa10bd4.
//
// Solidity: event PunkNoLongerForSale(uint256 indexed punkIndex)
func (_Store *StoreFilterer) FilterPunkNoLongerForSale(opts *bind.FilterOpts, punkIndex []*big.Int) (*StorePunkNoLongerForSaleIterator, error) {

	var punkIndexRule []interface{}
	for _, punkIndexItem := range punkIndex {
		punkIndexRule = append(punkIndexRule, punkIndexItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "PunkNoLongerForSale", punkIndexRule)
	if err != nil {
		return nil, err
	}
	return &StorePunkNoLongerForSaleIterator{contract: _Store.contract, event: "PunkNoLongerForSale", logs: logs, sub: sub}, nil
}

// WatchPunkNoLongerForSale is a free log subscription operation binding the contract event 0xb0e0a660b4e50f26f0b7ce75c24655fc76cc66e3334a54ff410277229fa10bd4.
//
// Solidity: event PunkNoLongerForSale(uint256 indexed punkIndex)
func (_Store *StoreFilterer) WatchPunkNoLongerForSale(opts *bind.WatchOpts, sink chan<- *StorePunkNoLongerForSale, punkIndex []*big.Int) (event.Subscription, error) {

	var punkIndexRule []interface{}
	for _, punkIndexItem := range punkIndex {
		punkIndexRule = append(punkIndexRule, punkIndexItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "PunkNoLongerForSale", punkIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorePunkNoLongerForSale)
				if err := _Store.contract.UnpackLog(event, "PunkNoLongerForSale", log); err != nil {
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
func (_Store *StoreFilterer) ParsePunkNoLongerForSale(log types.Log) (*StorePunkNoLongerForSale, error) {
	event := new(StorePunkNoLongerForSale)
	if err := _Store.contract.UnpackLog(event, "PunkNoLongerForSale", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorePunkOfferedIterator is returned from FilterPunkOffered and is used to iterate over the raw logs and unpacked data for PunkOffered events raised by the Store contract.
type StorePunkOfferedIterator struct {
	Event *StorePunkOffered // Event containing the contract specifics and raw log

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
func (it *StorePunkOfferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorePunkOffered)
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
		it.Event = new(StorePunkOffered)
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
func (it *StorePunkOfferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorePunkOfferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorePunkOffered represents a PunkOffered event raised by the Store contract.
type StorePunkOffered struct {
	PunkIndex *big.Int
	MinValue  *big.Int
	ToAddress common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPunkOffered is a free log retrieval operation binding the contract event 0x3c7b682d5da98001a9b8cbda6c647d2c63d698a4184fd1d55e2ce7b66f5d21eb.
//
// Solidity: event PunkOffered(uint256 indexed punkIndex, uint256 minValue, address indexed toAddress)
func (_Store *StoreFilterer) FilterPunkOffered(opts *bind.FilterOpts, punkIndex []*big.Int, toAddress []common.Address) (*StorePunkOfferedIterator, error) {

	var punkIndexRule []interface{}
	for _, punkIndexItem := range punkIndex {
		punkIndexRule = append(punkIndexRule, punkIndexItem)
	}

	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "PunkOffered", punkIndexRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return &StorePunkOfferedIterator{contract: _Store.contract, event: "PunkOffered", logs: logs, sub: sub}, nil
}

// WatchPunkOffered is a free log subscription operation binding the contract event 0x3c7b682d5da98001a9b8cbda6c647d2c63d698a4184fd1d55e2ce7b66f5d21eb.
//
// Solidity: event PunkOffered(uint256 indexed punkIndex, uint256 minValue, address indexed toAddress)
func (_Store *StoreFilterer) WatchPunkOffered(opts *bind.WatchOpts, sink chan<- *StorePunkOffered, punkIndex []*big.Int, toAddress []common.Address) (event.Subscription, error) {

	var punkIndexRule []interface{}
	for _, punkIndexItem := range punkIndex {
		punkIndexRule = append(punkIndexRule, punkIndexItem)
	}

	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "PunkOffered", punkIndexRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorePunkOffered)
				if err := _Store.contract.UnpackLog(event, "PunkOffered", log); err != nil {
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
func (_Store *StoreFilterer) ParsePunkOffered(log types.Log) (*StorePunkOffered, error) {
	event := new(StorePunkOffered)
	if err := _Store.contract.UnpackLog(event, "PunkOffered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorePunkTransferIterator is returned from FilterPunkTransfer and is used to iterate over the raw logs and unpacked data for PunkTransfer events raised by the Store contract.
type StorePunkTransferIterator struct {
	Event *StorePunkTransfer // Event containing the contract specifics and raw log

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
func (it *StorePunkTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorePunkTransfer)
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
		it.Event = new(StorePunkTransfer)
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
func (it *StorePunkTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorePunkTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorePunkTransfer represents a PunkTransfer event raised by the Store contract.
type StorePunkTransfer struct {
	From      common.Address
	To        common.Address
	PunkIndex *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPunkTransfer is a free log retrieval operation binding the contract event 0x05af636b70da6819000c49f85b21fa82081c632069bb626f30932034099107d8.
//
// Solidity: event PunkTransfer(address indexed from, address indexed to, uint256 punkIndex)
func (_Store *StoreFilterer) FilterPunkTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StorePunkTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "PunkTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StorePunkTransferIterator{contract: _Store.contract, event: "PunkTransfer", logs: logs, sub: sub}, nil
}

// WatchPunkTransfer is a free log subscription operation binding the contract event 0x05af636b70da6819000c49f85b21fa82081c632069bb626f30932034099107d8.
//
// Solidity: event PunkTransfer(address indexed from, address indexed to, uint256 punkIndex)
func (_Store *StoreFilterer) WatchPunkTransfer(opts *bind.WatchOpts, sink chan<- *StorePunkTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "PunkTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorePunkTransfer)
				if err := _Store.contract.UnpackLog(event, "PunkTransfer", log); err != nil {
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
func (_Store *StoreFilterer) ParsePunkTransfer(log types.Log) (*StorePunkTransfer, error) {
	event := new(StorePunkTransfer)
	if err := _Store.contract.UnpackLog(event, "PunkTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Store contract.
type StoreTransferIterator struct {
	Event *StoreTransfer // Event containing the contract specifics and raw log

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
func (it *StoreTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreTransfer)
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
		it.Event = new(StoreTransfer)
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
func (it *StoreTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreTransfer represents a Transfer event raised by the Store contract.
type StoreTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Store *StoreFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StoreTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StoreTransferIterator{contract: _Store.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Store *StoreFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StoreTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreTransfer)
				if err := _Store.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Store *StoreFilterer) ParseTransfer(log types.Log) (*StoreTransfer, error) {
	event := new(StoreTransfer)
	if err := _Store.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
