// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sushibar

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

// SushibarABI is the input ABI used to generate the binding from.
const SushibarABI = "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_sushi\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"enter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_share\",\"type\":\"uint256\"}],\"name\":\"leave\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sushi\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Sushibar is an auto generated Go binding around an Ethereum contract.
type Sushibar struct {
	SushibarCaller     // Read-only binding to the contract
	SushibarTransactor // Write-only binding to the contract
	SushibarFilterer   // Log filterer for contract events
}

// SushibarCaller is an auto generated read-only Go binding around an Ethereum contract.
type SushibarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SushibarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SushibarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SushibarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SushibarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SushibarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SushibarSession struct {
	Contract     *Sushibar         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SushibarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SushibarCallerSession struct {
	Contract *SushibarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SushibarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SushibarTransactorSession struct {
	Contract     *SushibarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SushibarRaw is an auto generated low-level Go binding around an Ethereum contract.
type SushibarRaw struct {
	Contract *Sushibar // Generic contract binding to access the raw methods on
}

// SushibarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SushibarCallerRaw struct {
	Contract *SushibarCaller // Generic read-only contract binding to access the raw methods on
}

// SushibarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SushibarTransactorRaw struct {
	Contract *SushibarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSushibar creates a new instance of Sushibar, bound to a specific deployed contract.
func NewSushibar(address common.Address, backend bind.ContractBackend) (*Sushibar, error) {
	contract, err := bindSushibar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sushibar{SushibarCaller: SushibarCaller{contract: contract}, SushibarTransactor: SushibarTransactor{contract: contract}, SushibarFilterer: SushibarFilterer{contract: contract}}, nil
}

// NewSushibarCaller creates a new read-only instance of Sushibar, bound to a specific deployed contract.
func NewSushibarCaller(address common.Address, caller bind.ContractCaller) (*SushibarCaller, error) {
	contract, err := bindSushibar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SushibarCaller{contract: contract}, nil
}

// NewSushibarTransactor creates a new write-only instance of Sushibar, bound to a specific deployed contract.
func NewSushibarTransactor(address common.Address, transactor bind.ContractTransactor) (*SushibarTransactor, error) {
	contract, err := bindSushibar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SushibarTransactor{contract: contract}, nil
}

// NewSushibarFilterer creates a new log filterer instance of Sushibar, bound to a specific deployed contract.
func NewSushibarFilterer(address common.Address, filterer bind.ContractFilterer) (*SushibarFilterer, error) {
	contract, err := bindSushibar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SushibarFilterer{contract: contract}, nil
}

// bindSushibar binds a generic wrapper to an already deployed contract.
func bindSushibar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SushibarABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sushibar *SushibarRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sushibar.Contract.SushibarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sushibar *SushibarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sushibar.Contract.SushibarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sushibar *SushibarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sushibar.Contract.SushibarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sushibar *SushibarCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sushibar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sushibar *SushibarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sushibar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sushibar *SushibarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sushibar.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Sushibar *SushibarCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Sushibar.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Sushibar *SushibarSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Sushibar.Contract.Allowance(&_Sushibar.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Sushibar *SushibarCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Sushibar.Contract.Allowance(&_Sushibar.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Sushibar *SushibarCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Sushibar.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Sushibar *SushibarSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Sushibar.Contract.BalanceOf(&_Sushibar.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Sushibar *SushibarCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Sushibar.Contract.BalanceOf(&_Sushibar.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Sushibar *SushibarCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Sushibar.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Sushibar *SushibarSession) Decimals() (uint8, error) {
	return _Sushibar.Contract.Decimals(&_Sushibar.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Sushibar *SushibarCallerSession) Decimals() (uint8, error) {
	return _Sushibar.Contract.Decimals(&_Sushibar.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Sushibar *SushibarCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Sushibar.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Sushibar *SushibarSession) Name() (string, error) {
	return _Sushibar.Contract.Name(&_Sushibar.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Sushibar *SushibarCallerSession) Name() (string, error) {
	return _Sushibar.Contract.Name(&_Sushibar.CallOpts)
}

// Sushi is a free data retrieval call binding the contract method 0x0a087903.
//
// Solidity: function sushi() view returns(address)
func (_Sushibar *SushibarCaller) Sushi(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sushibar.contract.Call(opts, &out, "sushi")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sushi is a free data retrieval call binding the contract method 0x0a087903.
//
// Solidity: function sushi() view returns(address)
func (_Sushibar *SushibarSession) Sushi() (common.Address, error) {
	return _Sushibar.Contract.Sushi(&_Sushibar.CallOpts)
}

// Sushi is a free data retrieval call binding the contract method 0x0a087903.
//
// Solidity: function sushi() view returns(address)
func (_Sushibar *SushibarCallerSession) Sushi() (common.Address, error) {
	return _Sushibar.Contract.Sushi(&_Sushibar.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Sushibar *SushibarCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Sushibar.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Sushibar *SushibarSession) Symbol() (string, error) {
	return _Sushibar.Contract.Symbol(&_Sushibar.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Sushibar *SushibarCallerSession) Symbol() (string, error) {
	return _Sushibar.Contract.Symbol(&_Sushibar.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Sushibar *SushibarCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sushibar.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Sushibar *SushibarSession) TotalSupply() (*big.Int, error) {
	return _Sushibar.Contract.TotalSupply(&_Sushibar.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Sushibar *SushibarCallerSession) TotalSupply() (*big.Int, error) {
	return _Sushibar.Contract.TotalSupply(&_Sushibar.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Sushibar *SushibarTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Sushibar.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Sushibar *SushibarSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Sushibar.Contract.Approve(&_Sushibar.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Sushibar *SushibarTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Sushibar.Contract.Approve(&_Sushibar.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Sushibar *SushibarTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Sushibar.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Sushibar *SushibarSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Sushibar.Contract.DecreaseAllowance(&_Sushibar.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Sushibar *SushibarTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Sushibar.Contract.DecreaseAllowance(&_Sushibar.TransactOpts, spender, subtractedValue)
}

// Enter is a paid mutator transaction binding the contract method 0xa59f3e0c.
//
// Solidity: function enter(uint256 _amount) returns()
func (_Sushibar *SushibarTransactor) Enter(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Sushibar.contract.Transact(opts, "enter", _amount)
}

// Enter is a paid mutator transaction binding the contract method 0xa59f3e0c.
//
// Solidity: function enter(uint256 _amount) returns()
func (_Sushibar *SushibarSession) Enter(_amount *big.Int) (*types.Transaction, error) {
	return _Sushibar.Contract.Enter(&_Sushibar.TransactOpts, _amount)
}

// Enter is a paid mutator transaction binding the contract method 0xa59f3e0c.
//
// Solidity: function enter(uint256 _amount) returns()
func (_Sushibar *SushibarTransactorSession) Enter(_amount *big.Int) (*types.Transaction, error) {
	return _Sushibar.Contract.Enter(&_Sushibar.TransactOpts, _amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Sushibar *SushibarTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Sushibar.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Sushibar *SushibarSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Sushibar.Contract.IncreaseAllowance(&_Sushibar.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Sushibar *SushibarTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Sushibar.Contract.IncreaseAllowance(&_Sushibar.TransactOpts, spender, addedValue)
}

// Leave is a paid mutator transaction binding the contract method 0x67dfd4c9.
//
// Solidity: function leave(uint256 _share) returns()
func (_Sushibar *SushibarTransactor) Leave(opts *bind.TransactOpts, _share *big.Int) (*types.Transaction, error) {
	return _Sushibar.contract.Transact(opts, "leave", _share)
}

// Leave is a paid mutator transaction binding the contract method 0x67dfd4c9.
//
// Solidity: function leave(uint256 _share) returns()
func (_Sushibar *SushibarSession) Leave(_share *big.Int) (*types.Transaction, error) {
	return _Sushibar.Contract.Leave(&_Sushibar.TransactOpts, _share)
}

// Leave is a paid mutator transaction binding the contract method 0x67dfd4c9.
//
// Solidity: function leave(uint256 _share) returns()
func (_Sushibar *SushibarTransactorSession) Leave(_share *big.Int) (*types.Transaction, error) {
	return _Sushibar.Contract.Leave(&_Sushibar.TransactOpts, _share)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Sushibar *SushibarTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Sushibar.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Sushibar *SushibarSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Sushibar.Contract.Transfer(&_Sushibar.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Sushibar *SushibarTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Sushibar.Contract.Transfer(&_Sushibar.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Sushibar *SushibarTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Sushibar.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Sushibar *SushibarSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Sushibar.Contract.TransferFrom(&_Sushibar.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Sushibar *SushibarTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Sushibar.Contract.TransferFrom(&_Sushibar.TransactOpts, sender, recipient, amount)
}

// SushibarApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Sushibar contract.
type SushibarApprovalIterator struct {
	Event *SushibarApproval // Event containing the contract specifics and raw log

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
func (it *SushibarApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SushibarApproval)
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
		it.Event = new(SushibarApproval)
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
func (it *SushibarApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SushibarApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SushibarApproval represents a Approval event raised by the Sushibar contract.
type SushibarApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Sushibar *SushibarFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SushibarApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Sushibar.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SushibarApprovalIterator{contract: _Sushibar.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Sushibar *SushibarFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SushibarApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Sushibar.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SushibarApproval)
				if err := _Sushibar.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Sushibar *SushibarFilterer) ParseApproval(log types.Log) (*SushibarApproval, error) {
	event := new(SushibarApproval)
	if err := _Sushibar.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SushibarTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Sushibar contract.
type SushibarTransferIterator struct {
	Event *SushibarTransfer // Event containing the contract specifics and raw log

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
func (it *SushibarTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SushibarTransfer)
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
		it.Event = new(SushibarTransfer)
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
func (it *SushibarTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SushibarTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SushibarTransfer represents a Transfer event raised by the Sushibar contract.
type SushibarTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Sushibar *SushibarFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SushibarTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Sushibar.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SushibarTransferIterator{contract: _Sushibar.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Sushibar *SushibarFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SushibarTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Sushibar.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SushibarTransfer)
				if err := _Sushibar.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Sushibar *SushibarFilterer) ParseTransfer(log types.Log) (*SushibarTransfer, error) {
	event := new(SushibarTransfer)
	if err := _Sushibar.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
