// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package yamv3

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

// Yamv3ABI is the input ABI used to generate the binding from.
const Yamv3ABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"initTotalSupply_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"implementation_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"becomeImplementationData\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromDelegate\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toDelegate\",\"type\":\"address\"}],\"name\":\"DelegateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"DelegateVotesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldGov\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newGov\",\"type\":\"address\"}],\"name\":\"NewGov\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldImplementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"NewImplementation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldIncentivizer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newIncentivizer\",\"type\":\"address\"}],\"name\":\"NewIncentivizer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldMigrator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newMigrator\",\"type\":\"address\"}],\"name\":\"NewMigrator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPendingGov\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingGov\",\"type\":\"address\"}],\"name\":\"NewPendingGov\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldRebaser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRebaser\",\"type\":\"address\"}],\"name\":\"NewRebaser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevYamsScalingFactor\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newYamsScalingFactor\",\"type\":\"uint256\"}],\"name\":\"Rebase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DELEGATION_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"checkpoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"fromBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"incentivizer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"internalDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"migrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"numCheckpoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingGov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rebaser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"yamsScalingFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowResign\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"becomeImplementationData\",\"type\":\"bytes\"}],\"name\":\"_setImplementation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxScalingFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indexDelta\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"positive\",\"type\":\"bool\"}],\"name\":\"rebase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"delegates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOfUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingGov\",\"type\":\"address\"}],\"name\":\"_setPendingGov\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rebaser_\",\"type\":\"address\"}],\"name\":\"_setRebaser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"incentivizer_\",\"type\":\"address\"}],\"name\":\"_setIncentivizer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"migrator_\",\"type\":\"address\"}],\"name\":\"_setMigrator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"_acceptGov\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getPriorVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"delegateBySig\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getCurrentVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"yam\",\"type\":\"uint256\"}],\"name\":\"yamToFragment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"fragmentToYam\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"delegateToImplementation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"delegateToViewImplementation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Yamv3 is an auto generated Go binding around an Ethereum contract.
type Yamv3 struct {
	Yamv3Caller     // Read-only binding to the contract
	Yamv3Transactor // Write-only binding to the contract
	Yamv3Filterer   // Log filterer for contract events
}

// Yamv3Caller is an auto generated read-only Go binding around an Ethereum contract.
type Yamv3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Yamv3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Yamv3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Yamv3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Yamv3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Yamv3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Yamv3Session struct {
	Contract     *Yamv3            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Yamv3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Yamv3CallerSession struct {
	Contract *Yamv3Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Yamv3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Yamv3TransactorSession struct {
	Contract     *Yamv3Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Yamv3Raw is an auto generated low-level Go binding around an Ethereum contract.
type Yamv3Raw struct {
	Contract *Yamv3 // Generic contract binding to access the raw methods on
}

// Yamv3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Yamv3CallerRaw struct {
	Contract *Yamv3Caller // Generic read-only contract binding to access the raw methods on
}

// Yamv3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Yamv3TransactorRaw struct {
	Contract *Yamv3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewYamv3 creates a new instance of Yamv3, bound to a specific deployed contract.
func NewYamv3(address common.Address, backend bind.ContractBackend) (*Yamv3, error) {
	contract, err := bindYamv3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Yamv3{Yamv3Caller: Yamv3Caller{contract: contract}, Yamv3Transactor: Yamv3Transactor{contract: contract}, Yamv3Filterer: Yamv3Filterer{contract: contract}}, nil
}

// NewYamv3Caller creates a new read-only instance of Yamv3, bound to a specific deployed contract.
func NewYamv3Caller(address common.Address, caller bind.ContractCaller) (*Yamv3Caller, error) {
	contract, err := bindYamv3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Yamv3Caller{contract: contract}, nil
}

// NewYamv3Transactor creates a new write-only instance of Yamv3, bound to a specific deployed contract.
func NewYamv3Transactor(address common.Address, transactor bind.ContractTransactor) (*Yamv3Transactor, error) {
	contract, err := bindYamv3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Yamv3Transactor{contract: contract}, nil
}

// NewYamv3Filterer creates a new log filterer instance of Yamv3, bound to a specific deployed contract.
func NewYamv3Filterer(address common.Address, filterer bind.ContractFilterer) (*Yamv3Filterer, error) {
	contract, err := bindYamv3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Yamv3Filterer{contract: contract}, nil
}

// bindYamv3 binds a generic wrapper to an already deployed contract.
func bindYamv3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Yamv3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yamv3 *Yamv3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yamv3.Contract.Yamv3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yamv3 *Yamv3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yamv3.Contract.Yamv3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yamv3 *Yamv3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yamv3.Contract.Yamv3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yamv3 *Yamv3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yamv3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yamv3 *Yamv3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yamv3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yamv3 *Yamv3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yamv3.Contract.contract.Transact(opts, method, params...)
}

// BASE is a free data retrieval call binding the contract method 0xec342ad0.
//
// Solidity: function BASE() view returns(uint256)
func (_Yamv3 *Yamv3Caller) BASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASE is a free data retrieval call binding the contract method 0xec342ad0.
//
// Solidity: function BASE() view returns(uint256)
func (_Yamv3 *Yamv3Session) BASE() (*big.Int, error) {
	return _Yamv3.Contract.BASE(&_Yamv3.CallOpts)
}

// BASE is a free data retrieval call binding the contract method 0xec342ad0.
//
// Solidity: function BASE() view returns(uint256)
func (_Yamv3 *Yamv3CallerSession) BASE() (*big.Int, error) {
	return _Yamv3.Contract.BASE(&_Yamv3.CallOpts)
}

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_Yamv3 *Yamv3Caller) DELEGATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "DELEGATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_Yamv3 *Yamv3Session) DELEGATIONTYPEHASH() ([32]byte, error) {
	return _Yamv3.Contract.DELEGATIONTYPEHASH(&_Yamv3.CallOpts)
}

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_Yamv3 *Yamv3CallerSession) DELEGATIONTYPEHASH() ([32]byte, error) {
	return _Yamv3.Contract.DELEGATIONTYPEHASH(&_Yamv3.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Yamv3 *Yamv3Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Yamv3 *Yamv3Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _Yamv3.Contract.DOMAINSEPARATOR(&_Yamv3.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Yamv3 *Yamv3CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Yamv3.Contract.DOMAINSEPARATOR(&_Yamv3.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Yamv3 *Yamv3Caller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Yamv3 *Yamv3Session) DOMAINTYPEHASH() ([32]byte, error) {
	return _Yamv3.Contract.DOMAINTYPEHASH(&_Yamv3.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Yamv3 *Yamv3CallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _Yamv3.Contract.DOMAINTYPEHASH(&_Yamv3.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Yamv3 *Yamv3Caller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Yamv3 *Yamv3Session) PERMITTYPEHASH() ([32]byte, error) {
	return _Yamv3.Contract.PERMITTYPEHASH(&_Yamv3.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Yamv3 *Yamv3CallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Yamv3.Contract.PERMITTYPEHASH(&_Yamv3.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Yamv3 *Yamv3Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Yamv3 *Yamv3Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Yamv3.Contract.Allowance(&_Yamv3.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Yamv3 *Yamv3CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Yamv3.Contract.Allowance(&_Yamv3.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Yamv3 *Yamv3Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Yamv3 *Yamv3Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Yamv3.Contract.BalanceOf(&_Yamv3.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Yamv3 *Yamv3CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Yamv3.Contract.BalanceOf(&_Yamv3.CallOpts, owner)
}

// BalanceOfUnderlying is a free data retrieval call binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) view returns(uint256)
func (_Yamv3 *Yamv3Caller) BalanceOfUnderlying(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "balanceOfUnderlying", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfUnderlying is a free data retrieval call binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) view returns(uint256)
func (_Yamv3 *Yamv3Session) BalanceOfUnderlying(owner common.Address) (*big.Int, error) {
	return _Yamv3.Contract.BalanceOfUnderlying(&_Yamv3.CallOpts, owner)
}

// BalanceOfUnderlying is a free data retrieval call binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) view returns(uint256)
func (_Yamv3 *Yamv3CallerSession) BalanceOfUnderlying(owner common.Address) (*big.Int, error) {
	return _Yamv3.Contract.BalanceOfUnderlying(&_Yamv3.CallOpts, owner)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 fromBlock, uint256 votes)
func (_Yamv3 *Yamv3Caller) Checkpoints(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (struct {
	FromBlock uint32
	Votes     *big.Int
}, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "checkpoints", arg0, arg1)

	outstruct := new(struct {
		FromBlock uint32
		Votes     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FromBlock = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.Votes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 fromBlock, uint256 votes)
func (_Yamv3 *Yamv3Session) Checkpoints(arg0 common.Address, arg1 uint32) (struct {
	FromBlock uint32
	Votes     *big.Int
}, error) {
	return _Yamv3.Contract.Checkpoints(&_Yamv3.CallOpts, arg0, arg1)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 fromBlock, uint256 votes)
func (_Yamv3 *Yamv3CallerSession) Checkpoints(arg0 common.Address, arg1 uint32) (struct {
	FromBlock uint32
	Votes     *big.Int
}, error) {
	return _Yamv3.Contract.Checkpoints(&_Yamv3.CallOpts, arg0, arg1)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Yamv3 *Yamv3Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Yamv3 *Yamv3Session) Decimals() (uint8, error) {
	return _Yamv3.Contract.Decimals(&_Yamv3.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Yamv3 *Yamv3CallerSession) Decimals() (uint8, error) {
	return _Yamv3.Contract.Decimals(&_Yamv3.CallOpts)
}

// DelegateToViewImplementation is a free data retrieval call binding the contract method 0x4487152f.
//
// Solidity: function delegateToViewImplementation(bytes data) view returns(bytes)
func (_Yamv3 *Yamv3Caller) DelegateToViewImplementation(opts *bind.CallOpts, data []byte) ([]byte, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "delegateToViewImplementation", data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DelegateToViewImplementation is a free data retrieval call binding the contract method 0x4487152f.
//
// Solidity: function delegateToViewImplementation(bytes data) view returns(bytes)
func (_Yamv3 *Yamv3Session) DelegateToViewImplementation(data []byte) ([]byte, error) {
	return _Yamv3.Contract.DelegateToViewImplementation(&_Yamv3.CallOpts, data)
}

// DelegateToViewImplementation is a free data retrieval call binding the contract method 0x4487152f.
//
// Solidity: function delegateToViewImplementation(bytes data) view returns(bytes)
func (_Yamv3 *Yamv3CallerSession) DelegateToViewImplementation(data []byte) ([]byte, error) {
	return _Yamv3.Contract.DelegateToViewImplementation(&_Yamv3.CallOpts, data)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address delegator) view returns(address)
func (_Yamv3 *Yamv3Caller) Delegates(opts *bind.CallOpts, delegator common.Address) (common.Address, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "delegates", delegator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address delegator) view returns(address)
func (_Yamv3 *Yamv3Session) Delegates(delegator common.Address) (common.Address, error) {
	return _Yamv3.Contract.Delegates(&_Yamv3.CallOpts, delegator)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address delegator) view returns(address)
func (_Yamv3 *Yamv3CallerSession) Delegates(delegator common.Address) (common.Address, error) {
	return _Yamv3.Contract.Delegates(&_Yamv3.CallOpts, delegator)
}

// FragmentToYam is a free data retrieval call binding the contract method 0x09c86403.
//
// Solidity: function fragmentToYam(uint256 value) view returns(uint256)
func (_Yamv3 *Yamv3Caller) FragmentToYam(opts *bind.CallOpts, value *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "fragmentToYam", value)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FragmentToYam is a free data retrieval call binding the contract method 0x09c86403.
//
// Solidity: function fragmentToYam(uint256 value) view returns(uint256)
func (_Yamv3 *Yamv3Session) FragmentToYam(value *big.Int) (*big.Int, error) {
	return _Yamv3.Contract.FragmentToYam(&_Yamv3.CallOpts, value)
}

// FragmentToYam is a free data retrieval call binding the contract method 0x09c86403.
//
// Solidity: function fragmentToYam(uint256 value) view returns(uint256)
func (_Yamv3 *Yamv3CallerSession) FragmentToYam(value *big.Int) (*big.Int, error) {
	return _Yamv3.Contract.FragmentToYam(&_Yamv3.CallOpts, value)
}

// GetCurrentVotes is a free data retrieval call binding the contract method 0xb4b5ea57.
//
// Solidity: function getCurrentVotes(address account) view returns(uint256)
func (_Yamv3 *Yamv3Caller) GetCurrentVotes(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "getCurrentVotes", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentVotes is a free data retrieval call binding the contract method 0xb4b5ea57.
//
// Solidity: function getCurrentVotes(address account) view returns(uint256)
func (_Yamv3 *Yamv3Session) GetCurrentVotes(account common.Address) (*big.Int, error) {
	return _Yamv3.Contract.GetCurrentVotes(&_Yamv3.CallOpts, account)
}

// GetCurrentVotes is a free data retrieval call binding the contract method 0xb4b5ea57.
//
// Solidity: function getCurrentVotes(address account) view returns(uint256)
func (_Yamv3 *Yamv3CallerSession) GetCurrentVotes(account common.Address) (*big.Int, error) {
	return _Yamv3.Contract.GetCurrentVotes(&_Yamv3.CallOpts, account)
}

// GetPriorVotes is a free data retrieval call binding the contract method 0x782d6fe1.
//
// Solidity: function getPriorVotes(address account, uint256 blockNumber) view returns(uint256)
func (_Yamv3 *Yamv3Caller) GetPriorVotes(opts *bind.CallOpts, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "getPriorVotes", account, blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriorVotes is a free data retrieval call binding the contract method 0x782d6fe1.
//
// Solidity: function getPriorVotes(address account, uint256 blockNumber) view returns(uint256)
func (_Yamv3 *Yamv3Session) GetPriorVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _Yamv3.Contract.GetPriorVotes(&_Yamv3.CallOpts, account, blockNumber)
}

// GetPriorVotes is a free data retrieval call binding the contract method 0x782d6fe1.
//
// Solidity: function getPriorVotes(address account, uint256 blockNumber) view returns(uint256)
func (_Yamv3 *Yamv3CallerSession) GetPriorVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _Yamv3.Contract.GetPriorVotes(&_Yamv3.CallOpts, account, blockNumber)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Yamv3 *Yamv3Caller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Yamv3 *Yamv3Session) Gov() (common.Address, error) {
	return _Yamv3.Contract.Gov(&_Yamv3.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Yamv3 *Yamv3CallerSession) Gov() (common.Address, error) {
	return _Yamv3.Contract.Gov(&_Yamv3.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Yamv3 *Yamv3Caller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Yamv3 *Yamv3Session) Implementation() (common.Address, error) {
	return _Yamv3.Contract.Implementation(&_Yamv3.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Yamv3 *Yamv3CallerSession) Implementation() (common.Address, error) {
	return _Yamv3.Contract.Implementation(&_Yamv3.CallOpts)
}

// Incentivizer is a free data retrieval call binding the contract method 0x6fc6407c.
//
// Solidity: function incentivizer() view returns(address)
func (_Yamv3 *Yamv3Caller) Incentivizer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "incentivizer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Incentivizer is a free data retrieval call binding the contract method 0x6fc6407c.
//
// Solidity: function incentivizer() view returns(address)
func (_Yamv3 *Yamv3Session) Incentivizer() (common.Address, error) {
	return _Yamv3.Contract.Incentivizer(&_Yamv3.CallOpts)
}

// Incentivizer is a free data retrieval call binding the contract method 0x6fc6407c.
//
// Solidity: function incentivizer() view returns(address)
func (_Yamv3 *Yamv3CallerSession) Incentivizer() (common.Address, error) {
	return _Yamv3.Contract.Incentivizer(&_Yamv3.CallOpts)
}

// InitSupply is a free data retrieval call binding the contract method 0x97d63f93.
//
// Solidity: function initSupply() view returns(uint256)
func (_Yamv3 *Yamv3Caller) InitSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "initSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitSupply is a free data retrieval call binding the contract method 0x97d63f93.
//
// Solidity: function initSupply() view returns(uint256)
func (_Yamv3 *Yamv3Session) InitSupply() (*big.Int, error) {
	return _Yamv3.Contract.InitSupply(&_Yamv3.CallOpts)
}

// InitSupply is a free data retrieval call binding the contract method 0x97d63f93.
//
// Solidity: function initSupply() view returns(uint256)
func (_Yamv3 *Yamv3CallerSession) InitSupply() (*big.Int, error) {
	return _Yamv3.Contract.InitSupply(&_Yamv3.CallOpts)
}

// InternalDecimals is a free data retrieval call binding the contract method 0x64dd48f5.
//
// Solidity: function internalDecimals() view returns(uint256)
func (_Yamv3 *Yamv3Caller) InternalDecimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "internalDecimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InternalDecimals is a free data retrieval call binding the contract method 0x64dd48f5.
//
// Solidity: function internalDecimals() view returns(uint256)
func (_Yamv3 *Yamv3Session) InternalDecimals() (*big.Int, error) {
	return _Yamv3.Contract.InternalDecimals(&_Yamv3.CallOpts)
}

// InternalDecimals is a free data retrieval call binding the contract method 0x64dd48f5.
//
// Solidity: function internalDecimals() view returns(uint256)
func (_Yamv3 *Yamv3CallerSession) InternalDecimals() (*big.Int, error) {
	return _Yamv3.Contract.InternalDecimals(&_Yamv3.CallOpts)
}

// MaxScalingFactor is a free data retrieval call binding the contract method 0x11d3e6c4.
//
// Solidity: function maxScalingFactor() view returns(uint256)
func (_Yamv3 *Yamv3Caller) MaxScalingFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "maxScalingFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxScalingFactor is a free data retrieval call binding the contract method 0x11d3e6c4.
//
// Solidity: function maxScalingFactor() view returns(uint256)
func (_Yamv3 *Yamv3Session) MaxScalingFactor() (*big.Int, error) {
	return _Yamv3.Contract.MaxScalingFactor(&_Yamv3.CallOpts)
}

// MaxScalingFactor is a free data retrieval call binding the contract method 0x11d3e6c4.
//
// Solidity: function maxScalingFactor() view returns(uint256)
func (_Yamv3 *Yamv3CallerSession) MaxScalingFactor() (*big.Int, error) {
	return _Yamv3.Contract.MaxScalingFactor(&_Yamv3.CallOpts)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_Yamv3 *Yamv3Caller) Migrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "migrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_Yamv3 *Yamv3Session) Migrator() (common.Address, error) {
	return _Yamv3.Contract.Migrator(&_Yamv3.CallOpts)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_Yamv3 *Yamv3CallerSession) Migrator() (common.Address, error) {
	return _Yamv3.Contract.Migrator(&_Yamv3.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Yamv3 *Yamv3Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Yamv3 *Yamv3Session) Name() (string, error) {
	return _Yamv3.Contract.Name(&_Yamv3.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Yamv3 *Yamv3CallerSession) Name() (string, error) {
	return _Yamv3.Contract.Name(&_Yamv3.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Yamv3 *Yamv3Caller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Yamv3 *Yamv3Session) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Yamv3.Contract.Nonces(&_Yamv3.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Yamv3 *Yamv3CallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Yamv3.Contract.Nonces(&_Yamv3.CallOpts, arg0)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_Yamv3 *Yamv3Caller) NumCheckpoints(opts *bind.CallOpts, arg0 common.Address) (uint32, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "numCheckpoints", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_Yamv3 *Yamv3Session) NumCheckpoints(arg0 common.Address) (uint32, error) {
	return _Yamv3.Contract.NumCheckpoints(&_Yamv3.CallOpts, arg0)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_Yamv3 *Yamv3CallerSession) NumCheckpoints(arg0 common.Address) (uint32, error) {
	return _Yamv3.Contract.NumCheckpoints(&_Yamv3.CallOpts, arg0)
}

// PendingGov is a free data retrieval call binding the contract method 0x25240810.
//
// Solidity: function pendingGov() view returns(address)
func (_Yamv3 *Yamv3Caller) PendingGov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "pendingGov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingGov is a free data retrieval call binding the contract method 0x25240810.
//
// Solidity: function pendingGov() view returns(address)
func (_Yamv3 *Yamv3Session) PendingGov() (common.Address, error) {
	return _Yamv3.Contract.PendingGov(&_Yamv3.CallOpts)
}

// PendingGov is a free data retrieval call binding the contract method 0x25240810.
//
// Solidity: function pendingGov() view returns(address)
func (_Yamv3 *Yamv3CallerSession) PendingGov() (common.Address, error) {
	return _Yamv3.Contract.PendingGov(&_Yamv3.CallOpts)
}

// Rebaser is a free data retrieval call binding the contract method 0x11fd8a83.
//
// Solidity: function rebaser() view returns(address)
func (_Yamv3 *Yamv3Caller) Rebaser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "rebaser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rebaser is a free data retrieval call binding the contract method 0x11fd8a83.
//
// Solidity: function rebaser() view returns(address)
func (_Yamv3 *Yamv3Session) Rebaser() (common.Address, error) {
	return _Yamv3.Contract.Rebaser(&_Yamv3.CallOpts)
}

// Rebaser is a free data retrieval call binding the contract method 0x11fd8a83.
//
// Solidity: function rebaser() view returns(address)
func (_Yamv3 *Yamv3CallerSession) Rebaser() (common.Address, error) {
	return _Yamv3.Contract.Rebaser(&_Yamv3.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Yamv3 *Yamv3Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Yamv3 *Yamv3Session) Symbol() (string, error) {
	return _Yamv3.Contract.Symbol(&_Yamv3.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Yamv3 *Yamv3CallerSession) Symbol() (string, error) {
	return _Yamv3.Contract.Symbol(&_Yamv3.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Yamv3 *Yamv3Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Yamv3 *Yamv3Session) TotalSupply() (*big.Int, error) {
	return _Yamv3.Contract.TotalSupply(&_Yamv3.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Yamv3 *Yamv3CallerSession) TotalSupply() (*big.Int, error) {
	return _Yamv3.Contract.TotalSupply(&_Yamv3.CallOpts)
}

// YamToFragment is a free data retrieval call binding the contract method 0xf18d9b63.
//
// Solidity: function yamToFragment(uint256 yam) view returns(uint256)
func (_Yamv3 *Yamv3Caller) YamToFragment(opts *bind.CallOpts, yam *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "yamToFragment", yam)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// YamToFragment is a free data retrieval call binding the contract method 0xf18d9b63.
//
// Solidity: function yamToFragment(uint256 yam) view returns(uint256)
func (_Yamv3 *Yamv3Session) YamToFragment(yam *big.Int) (*big.Int, error) {
	return _Yamv3.Contract.YamToFragment(&_Yamv3.CallOpts, yam)
}

// YamToFragment is a free data retrieval call binding the contract method 0xf18d9b63.
//
// Solidity: function yamToFragment(uint256 yam) view returns(uint256)
func (_Yamv3 *Yamv3CallerSession) YamToFragment(yam *big.Int) (*big.Int, error) {
	return _Yamv3.Contract.YamToFragment(&_Yamv3.CallOpts, yam)
}

// YamsScalingFactor is a free data retrieval call binding the contract method 0xb6fa8576.
//
// Solidity: function yamsScalingFactor() view returns(uint256)
func (_Yamv3 *Yamv3Caller) YamsScalingFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yamv3.contract.Call(opts, &out, "yamsScalingFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// YamsScalingFactor is a free data retrieval call binding the contract method 0xb6fa8576.
//
// Solidity: function yamsScalingFactor() view returns(uint256)
func (_Yamv3 *Yamv3Session) YamsScalingFactor() (*big.Int, error) {
	return _Yamv3.Contract.YamsScalingFactor(&_Yamv3.CallOpts)
}

// YamsScalingFactor is a free data retrieval call binding the contract method 0xb6fa8576.
//
// Solidity: function yamsScalingFactor() view returns(uint256)
func (_Yamv3 *Yamv3CallerSession) YamsScalingFactor() (*big.Int, error) {
	return _Yamv3.Contract.YamsScalingFactor(&_Yamv3.CallOpts)
}

// AcceptGov is a paid mutator transaction binding the contract method 0x4bda2e20.
//
// Solidity: function _acceptGov() returns()
func (_Yamv3 *Yamv3Transactor) AcceptGov(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yamv3.contract.Transact(opts, "_acceptGov")
}

// AcceptGov is a paid mutator transaction binding the contract method 0x4bda2e20.
//
// Solidity: function _acceptGov() returns()
func (_Yamv3 *Yamv3Session) AcceptGov() (*types.Transaction, error) {
	return _Yamv3.Contract.AcceptGov(&_Yamv3.TransactOpts)
}

// AcceptGov is a paid mutator transaction binding the contract method 0x4bda2e20.
//
// Solidity: function _acceptGov() returns()
func (_Yamv3 *Yamv3TransactorSession) AcceptGov() (*types.Transaction, error) {
	return _Yamv3.Contract.AcceptGov(&_Yamv3.TransactOpts)
}

// SetImplementation is a paid mutator transaction binding the contract method 0x555bcc40.
//
// Solidity: function _setImplementation(address implementation_, bool allowResign, bytes becomeImplementationData) returns()
func (_Yamv3 *Yamv3Transactor) SetImplementation(opts *bind.TransactOpts, implementation_ common.Address, allowResign bool, becomeImplementationData []byte) (*types.Transaction, error) {
	return _Yamv3.contract.Transact(opts, "_setImplementation", implementation_, allowResign, becomeImplementationData)
}

// SetImplementation is a paid mutator transaction binding the contract method 0x555bcc40.
//
// Solidity: function _setImplementation(address implementation_, bool allowResign, bytes becomeImplementationData) returns()
func (_Yamv3 *Yamv3Session) SetImplementation(implementation_ common.Address, allowResign bool, becomeImplementationData []byte) (*types.Transaction, error) {
	return _Yamv3.Contract.SetImplementation(&_Yamv3.TransactOpts, implementation_, allowResign, becomeImplementationData)
}

// SetImplementation is a paid mutator transaction binding the contract method 0x555bcc40.
//
// Solidity: function _setImplementation(address implementation_, bool allowResign, bytes becomeImplementationData) returns()
func (_Yamv3 *Yamv3TransactorSession) SetImplementation(implementation_ common.Address, allowResign bool, becomeImplementationData []byte) (*types.Transaction, error) {
	return _Yamv3.Contract.SetImplementation(&_Yamv3.TransactOpts, implementation_, allowResign, becomeImplementationData)
}

// SetIncentivizer is a paid mutator transaction binding the contract method 0x98dca210.
//
// Solidity: function _setIncentivizer(address incentivizer_) returns()
func (_Yamv3 *Yamv3Transactor) SetIncentivizer(opts *bind.TransactOpts, incentivizer_ common.Address) (*types.Transaction, error) {
	return _Yamv3.contract.Transact(opts, "_setIncentivizer", incentivizer_)
}

// SetIncentivizer is a paid mutator transaction binding the contract method 0x98dca210.
//
// Solidity: function _setIncentivizer(address incentivizer_) returns()
func (_Yamv3 *Yamv3Session) SetIncentivizer(incentivizer_ common.Address) (*types.Transaction, error) {
	return _Yamv3.Contract.SetIncentivizer(&_Yamv3.TransactOpts, incentivizer_)
}

// SetIncentivizer is a paid mutator transaction binding the contract method 0x98dca210.
//
// Solidity: function _setIncentivizer(address incentivizer_) returns()
func (_Yamv3 *Yamv3TransactorSession) SetIncentivizer(incentivizer_ common.Address) (*types.Transaction, error) {
	return _Yamv3.Contract.SetIncentivizer(&_Yamv3.TransactOpts, incentivizer_)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x47dfe70d.
//
// Solidity: function _setMigrator(address migrator_) returns()
func (_Yamv3 *Yamv3Transactor) SetMigrator(opts *bind.TransactOpts, migrator_ common.Address) (*types.Transaction, error) {
	return _Yamv3.contract.Transact(opts, "_setMigrator", migrator_)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x47dfe70d.
//
// Solidity: function _setMigrator(address migrator_) returns()
func (_Yamv3 *Yamv3Session) SetMigrator(migrator_ common.Address) (*types.Transaction, error) {
	return _Yamv3.Contract.SetMigrator(&_Yamv3.TransactOpts, migrator_)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x47dfe70d.
//
// Solidity: function _setMigrator(address migrator_) returns()
func (_Yamv3 *Yamv3TransactorSession) SetMigrator(migrator_ common.Address) (*types.Transaction, error) {
	return _Yamv3.Contract.SetMigrator(&_Yamv3.TransactOpts, migrator_)
}

// SetPendingGov is a paid mutator transaction binding the contract method 0x73f03dff.
//
// Solidity: function _setPendingGov(address newPendingGov) returns()
func (_Yamv3 *Yamv3Transactor) SetPendingGov(opts *bind.TransactOpts, newPendingGov common.Address) (*types.Transaction, error) {
	return _Yamv3.contract.Transact(opts, "_setPendingGov", newPendingGov)
}

// SetPendingGov is a paid mutator transaction binding the contract method 0x73f03dff.
//
// Solidity: function _setPendingGov(address newPendingGov) returns()
func (_Yamv3 *Yamv3Session) SetPendingGov(newPendingGov common.Address) (*types.Transaction, error) {
	return _Yamv3.Contract.SetPendingGov(&_Yamv3.TransactOpts, newPendingGov)
}

// SetPendingGov is a paid mutator transaction binding the contract method 0x73f03dff.
//
// Solidity: function _setPendingGov(address newPendingGov) returns()
func (_Yamv3 *Yamv3TransactorSession) SetPendingGov(newPendingGov common.Address) (*types.Transaction, error) {
	return _Yamv3.Contract.SetPendingGov(&_Yamv3.TransactOpts, newPendingGov)
}

// SetRebaser is a paid mutator transaction binding the contract method 0xfa8f3455.
//
// Solidity: function _setRebaser(address rebaser_) returns()
func (_Yamv3 *Yamv3Transactor) SetRebaser(opts *bind.TransactOpts, rebaser_ common.Address) (*types.Transaction, error) {
	return _Yamv3.contract.Transact(opts, "_setRebaser", rebaser_)
}

// SetRebaser is a paid mutator transaction binding the contract method 0xfa8f3455.
//
// Solidity: function _setRebaser(address rebaser_) returns()
func (_Yamv3 *Yamv3Session) SetRebaser(rebaser_ common.Address) (*types.Transaction, error) {
	return _Yamv3.Contract.SetRebaser(&_Yamv3.TransactOpts, rebaser_)
}

// SetRebaser is a paid mutator transaction binding the contract method 0xfa8f3455.
//
// Solidity: function _setRebaser(address rebaser_) returns()
func (_Yamv3 *Yamv3TransactorSession) SetRebaser(rebaser_ common.Address) (*types.Transaction, error) {
	return _Yamv3.Contract.SetRebaser(&_Yamv3.TransactOpts, rebaser_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Yamv3 *Yamv3Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yamv3.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Yamv3 *Yamv3Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yamv3.Contract.Approve(&_Yamv3.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Yamv3 *Yamv3TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yamv3.Contract.Approve(&_Yamv3.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Yamv3 *Yamv3Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Yamv3.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Yamv3 *Yamv3Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Yamv3.Contract.DecreaseAllowance(&_Yamv3.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Yamv3 *Yamv3TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Yamv3.Contract.DecreaseAllowance(&_Yamv3.TransactOpts, spender, subtractedValue)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_Yamv3 *Yamv3Transactor) Delegate(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _Yamv3.contract.Transact(opts, "delegate", delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_Yamv3 *Yamv3Session) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _Yamv3.Contract.Delegate(&_Yamv3.TransactOpts, delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_Yamv3 *Yamv3TransactorSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _Yamv3.Contract.Delegate(&_Yamv3.TransactOpts, delegatee)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Yamv3 *Yamv3Transactor) DelegateBySig(opts *bind.TransactOpts, delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Yamv3.contract.Transact(opts, "delegateBySig", delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Yamv3 *Yamv3Session) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Yamv3.Contract.DelegateBySig(&_Yamv3.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Yamv3 *Yamv3TransactorSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Yamv3.Contract.DelegateBySig(&_Yamv3.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// DelegateToImplementation is a paid mutator transaction binding the contract method 0x0933c1ed.
//
// Solidity: function delegateToImplementation(bytes data) returns(bytes)
func (_Yamv3 *Yamv3Transactor) DelegateToImplementation(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Yamv3.contract.Transact(opts, "delegateToImplementation", data)
}

// DelegateToImplementation is a paid mutator transaction binding the contract method 0x0933c1ed.
//
// Solidity: function delegateToImplementation(bytes data) returns(bytes)
func (_Yamv3 *Yamv3Session) DelegateToImplementation(data []byte) (*types.Transaction, error) {
	return _Yamv3.Contract.DelegateToImplementation(&_Yamv3.TransactOpts, data)
}

// DelegateToImplementation is a paid mutator transaction binding the contract method 0x0933c1ed.
//
// Solidity: function delegateToImplementation(bytes data) returns(bytes)
func (_Yamv3 *Yamv3TransactorSession) DelegateToImplementation(data []byte) (*types.Transaction, error) {
	return _Yamv3.Contract.DelegateToImplementation(&_Yamv3.TransactOpts, data)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Yamv3 *Yamv3Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Yamv3.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Yamv3 *Yamv3Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Yamv3.Contract.IncreaseAllowance(&_Yamv3.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Yamv3 *Yamv3TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Yamv3.Contract.IncreaseAllowance(&_Yamv3.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 mintAmount) returns(bool)
func (_Yamv3 *Yamv3Transactor) Mint(opts *bind.TransactOpts, to common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Yamv3.contract.Transact(opts, "mint", to, mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 mintAmount) returns(bool)
func (_Yamv3 *Yamv3Session) Mint(to common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Yamv3.Contract.Mint(&_Yamv3.TransactOpts, to, mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 mintAmount) returns(bool)
func (_Yamv3 *Yamv3TransactorSession) Mint(to common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Yamv3.Contract.Mint(&_Yamv3.TransactOpts, to, mintAmount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Yamv3 *Yamv3Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Yamv3.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Yamv3 *Yamv3Session) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Yamv3.Contract.Permit(&_Yamv3.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Yamv3 *Yamv3TransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Yamv3.Contract.Permit(&_Yamv3.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Rebase is a paid mutator transaction binding the contract method 0x7af548c1.
//
// Solidity: function rebase(uint256 epoch, uint256 indexDelta, bool positive) returns(uint256)
func (_Yamv3 *Yamv3Transactor) Rebase(opts *bind.TransactOpts, epoch *big.Int, indexDelta *big.Int, positive bool) (*types.Transaction, error) {
	return _Yamv3.contract.Transact(opts, "rebase", epoch, indexDelta, positive)
}

// Rebase is a paid mutator transaction binding the contract method 0x7af548c1.
//
// Solidity: function rebase(uint256 epoch, uint256 indexDelta, bool positive) returns(uint256)
func (_Yamv3 *Yamv3Session) Rebase(epoch *big.Int, indexDelta *big.Int, positive bool) (*types.Transaction, error) {
	return _Yamv3.Contract.Rebase(&_Yamv3.TransactOpts, epoch, indexDelta, positive)
}

// Rebase is a paid mutator transaction binding the contract method 0x7af548c1.
//
// Solidity: function rebase(uint256 epoch, uint256 indexDelta, bool positive) returns(uint256)
func (_Yamv3 *Yamv3TransactorSession) Rebase(epoch *big.Int, indexDelta *big.Int, positive bool) (*types.Transaction, error) {
	return _Yamv3.Contract.Rebase(&_Yamv3.TransactOpts, epoch, indexDelta, positive)
}

// RescueTokens is a paid mutator transaction binding the contract method 0xcea9d26f.
//
// Solidity: function rescueTokens(address token, address to, uint256 amount) returns(bool)
func (_Yamv3 *Yamv3Transactor) RescueTokens(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yamv3.contract.Transact(opts, "rescueTokens", token, to, amount)
}

// RescueTokens is a paid mutator transaction binding the contract method 0xcea9d26f.
//
// Solidity: function rescueTokens(address token, address to, uint256 amount) returns(bool)
func (_Yamv3 *Yamv3Session) RescueTokens(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yamv3.Contract.RescueTokens(&_Yamv3.TransactOpts, token, to, amount)
}

// RescueTokens is a paid mutator transaction binding the contract method 0xcea9d26f.
//
// Solidity: function rescueTokens(address token, address to, uint256 amount) returns(bool)
func (_Yamv3 *Yamv3TransactorSession) RescueTokens(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yamv3.Contract.RescueTokens(&_Yamv3.TransactOpts, token, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Yamv3 *Yamv3Transactor) Transfer(opts *bind.TransactOpts, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yamv3.contract.Transact(opts, "transfer", dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Yamv3 *Yamv3Session) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yamv3.Contract.Transfer(&_Yamv3.TransactOpts, dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Yamv3 *Yamv3TransactorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yamv3.Contract.Transfer(&_Yamv3.TransactOpts, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Yamv3 *Yamv3Transactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yamv3.contract.Transact(opts, "transferFrom", src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Yamv3 *Yamv3Session) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yamv3.Contract.TransferFrom(&_Yamv3.TransactOpts, src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Yamv3 *Yamv3TransactorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yamv3.Contract.TransferFrom(&_Yamv3.TransactOpts, src, dst, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Yamv3 *Yamv3Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Yamv3.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Yamv3 *Yamv3Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Yamv3.Contract.Fallback(&_Yamv3.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Yamv3 *Yamv3TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Yamv3.Contract.Fallback(&_Yamv3.TransactOpts, calldata)
}

// Yamv3ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Yamv3 contract.
type Yamv3ApprovalIterator struct {
	Event *Yamv3Approval // Event containing the contract specifics and raw log

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
func (it *Yamv3ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yamv3Approval)
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
		it.Event = new(Yamv3Approval)
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
func (it *Yamv3ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yamv3ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yamv3Approval represents a Approval event raised by the Yamv3 contract.
type Yamv3Approval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Yamv3 *Yamv3Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Yamv3ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Yamv3.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Yamv3ApprovalIterator{contract: _Yamv3.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Yamv3 *Yamv3Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Yamv3Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Yamv3.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yamv3Approval)
				if err := _Yamv3.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Yamv3 *Yamv3Filterer) ParseApproval(log types.Log) (*Yamv3Approval, error) {
	event := new(Yamv3Approval)
	if err := _Yamv3.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yamv3DelegateChangedIterator is returned from FilterDelegateChanged and is used to iterate over the raw logs and unpacked data for DelegateChanged events raised by the Yamv3 contract.
type Yamv3DelegateChangedIterator struct {
	Event *Yamv3DelegateChanged // Event containing the contract specifics and raw log

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
func (it *Yamv3DelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yamv3DelegateChanged)
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
		it.Event = new(Yamv3DelegateChanged)
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
func (it *Yamv3DelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yamv3DelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yamv3DelegateChanged represents a DelegateChanged event raised by the Yamv3 contract.
type Yamv3DelegateChanged struct {
	Delegator    common.Address
	FromDelegate common.Address
	ToDelegate   common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegateChanged is a free log retrieval operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_Yamv3 *Yamv3Filterer) FilterDelegateChanged(opts *bind.FilterOpts, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (*Yamv3DelegateChangedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _Yamv3.contract.FilterLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return &Yamv3DelegateChangedIterator{contract: _Yamv3.contract, event: "DelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateChanged is a free log subscription operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_Yamv3 *Yamv3Filterer) WatchDelegateChanged(opts *bind.WatchOpts, sink chan<- *Yamv3DelegateChanged, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _Yamv3.contract.WatchLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yamv3DelegateChanged)
				if err := _Yamv3.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
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

// ParseDelegateChanged is a log parse operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_Yamv3 *Yamv3Filterer) ParseDelegateChanged(log types.Log) (*Yamv3DelegateChanged, error) {
	event := new(Yamv3DelegateChanged)
	if err := _Yamv3.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yamv3DelegateVotesChangedIterator is returned from FilterDelegateVotesChanged and is used to iterate over the raw logs and unpacked data for DelegateVotesChanged events raised by the Yamv3 contract.
type Yamv3DelegateVotesChangedIterator struct {
	Event *Yamv3DelegateVotesChanged // Event containing the contract specifics and raw log

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
func (it *Yamv3DelegateVotesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yamv3DelegateVotesChanged)
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
		it.Event = new(Yamv3DelegateVotesChanged)
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
func (it *Yamv3DelegateVotesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yamv3DelegateVotesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yamv3DelegateVotesChanged represents a DelegateVotesChanged event raised by the Yamv3 contract.
type Yamv3DelegateVotesChanged struct {
	Delegate        common.Address
	PreviousBalance *big.Int
	NewBalance      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelegateVotesChanged is a free log retrieval operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_Yamv3 *Yamv3Filterer) FilterDelegateVotesChanged(opts *bind.FilterOpts, delegate []common.Address) (*Yamv3DelegateVotesChangedIterator, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _Yamv3.contract.FilterLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return &Yamv3DelegateVotesChangedIterator{contract: _Yamv3.contract, event: "DelegateVotesChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateVotesChanged is a free log subscription operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_Yamv3 *Yamv3Filterer) WatchDelegateVotesChanged(opts *bind.WatchOpts, sink chan<- *Yamv3DelegateVotesChanged, delegate []common.Address) (event.Subscription, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _Yamv3.contract.WatchLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yamv3DelegateVotesChanged)
				if err := _Yamv3.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
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

// ParseDelegateVotesChanged is a log parse operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_Yamv3 *Yamv3Filterer) ParseDelegateVotesChanged(log types.Log) (*Yamv3DelegateVotesChanged, error) {
	event := new(Yamv3DelegateVotesChanged)
	if err := _Yamv3.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yamv3MintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Yamv3 contract.
type Yamv3MintIterator struct {
	Event *Yamv3Mint // Event containing the contract specifics and raw log

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
func (it *Yamv3MintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yamv3Mint)
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
		it.Event = new(Yamv3Mint)
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
func (it *Yamv3MintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yamv3MintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yamv3Mint represents a Mint event raised by the Yamv3 contract.
type Yamv3Mint struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address to, uint256 amount)
func (_Yamv3 *Yamv3Filterer) FilterMint(opts *bind.FilterOpts) (*Yamv3MintIterator, error) {

	logs, sub, err := _Yamv3.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &Yamv3MintIterator{contract: _Yamv3.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address to, uint256 amount)
func (_Yamv3 *Yamv3Filterer) WatchMint(opts *bind.WatchOpts, sink chan<- *Yamv3Mint) (event.Subscription, error) {

	logs, sub, err := _Yamv3.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yamv3Mint)
				if err := _Yamv3.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address to, uint256 amount)
func (_Yamv3 *Yamv3Filterer) ParseMint(log types.Log) (*Yamv3Mint, error) {
	event := new(Yamv3Mint)
	if err := _Yamv3.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yamv3NewGovIterator is returned from FilterNewGov and is used to iterate over the raw logs and unpacked data for NewGov events raised by the Yamv3 contract.
type Yamv3NewGovIterator struct {
	Event *Yamv3NewGov // Event containing the contract specifics and raw log

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
func (it *Yamv3NewGovIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yamv3NewGov)
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
		it.Event = new(Yamv3NewGov)
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
func (it *Yamv3NewGovIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yamv3NewGovIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yamv3NewGov represents a NewGov event raised by the Yamv3 contract.
type Yamv3NewGov struct {
	OldGov common.Address
	NewGov common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewGov is a free log retrieval operation binding the contract event 0x1f14cfc03e486d23acee577b07bc0b3b23f4888c91fcdba5e0fef5a2549d5523.
//
// Solidity: event NewGov(address oldGov, address newGov)
func (_Yamv3 *Yamv3Filterer) FilterNewGov(opts *bind.FilterOpts) (*Yamv3NewGovIterator, error) {

	logs, sub, err := _Yamv3.contract.FilterLogs(opts, "NewGov")
	if err != nil {
		return nil, err
	}
	return &Yamv3NewGovIterator{contract: _Yamv3.contract, event: "NewGov", logs: logs, sub: sub}, nil
}

// WatchNewGov is a free log subscription operation binding the contract event 0x1f14cfc03e486d23acee577b07bc0b3b23f4888c91fcdba5e0fef5a2549d5523.
//
// Solidity: event NewGov(address oldGov, address newGov)
func (_Yamv3 *Yamv3Filterer) WatchNewGov(opts *bind.WatchOpts, sink chan<- *Yamv3NewGov) (event.Subscription, error) {

	logs, sub, err := _Yamv3.contract.WatchLogs(opts, "NewGov")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yamv3NewGov)
				if err := _Yamv3.contract.UnpackLog(event, "NewGov", log); err != nil {
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

// ParseNewGov is a log parse operation binding the contract event 0x1f14cfc03e486d23acee577b07bc0b3b23f4888c91fcdba5e0fef5a2549d5523.
//
// Solidity: event NewGov(address oldGov, address newGov)
func (_Yamv3 *Yamv3Filterer) ParseNewGov(log types.Log) (*Yamv3NewGov, error) {
	event := new(Yamv3NewGov)
	if err := _Yamv3.contract.UnpackLog(event, "NewGov", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yamv3NewImplementationIterator is returned from FilterNewImplementation and is used to iterate over the raw logs and unpacked data for NewImplementation events raised by the Yamv3 contract.
type Yamv3NewImplementationIterator struct {
	Event *Yamv3NewImplementation // Event containing the contract specifics and raw log

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
func (it *Yamv3NewImplementationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yamv3NewImplementation)
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
		it.Event = new(Yamv3NewImplementation)
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
func (it *Yamv3NewImplementationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yamv3NewImplementationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yamv3NewImplementation represents a NewImplementation event raised by the Yamv3 contract.
type Yamv3NewImplementation struct {
	OldImplementation common.Address
	NewImplementation common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewImplementation is a free log retrieval operation binding the contract event 0xd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a.
//
// Solidity: event NewImplementation(address oldImplementation, address newImplementation)
func (_Yamv3 *Yamv3Filterer) FilterNewImplementation(opts *bind.FilterOpts) (*Yamv3NewImplementationIterator, error) {

	logs, sub, err := _Yamv3.contract.FilterLogs(opts, "NewImplementation")
	if err != nil {
		return nil, err
	}
	return &Yamv3NewImplementationIterator{contract: _Yamv3.contract, event: "NewImplementation", logs: logs, sub: sub}, nil
}

// WatchNewImplementation is a free log subscription operation binding the contract event 0xd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a.
//
// Solidity: event NewImplementation(address oldImplementation, address newImplementation)
func (_Yamv3 *Yamv3Filterer) WatchNewImplementation(opts *bind.WatchOpts, sink chan<- *Yamv3NewImplementation) (event.Subscription, error) {

	logs, sub, err := _Yamv3.contract.WatchLogs(opts, "NewImplementation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yamv3NewImplementation)
				if err := _Yamv3.contract.UnpackLog(event, "NewImplementation", log); err != nil {
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

// ParseNewImplementation is a log parse operation binding the contract event 0xd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a.
//
// Solidity: event NewImplementation(address oldImplementation, address newImplementation)
func (_Yamv3 *Yamv3Filterer) ParseNewImplementation(log types.Log) (*Yamv3NewImplementation, error) {
	event := new(Yamv3NewImplementation)
	if err := _Yamv3.contract.UnpackLog(event, "NewImplementation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yamv3NewIncentivizerIterator is returned from FilterNewIncentivizer and is used to iterate over the raw logs and unpacked data for NewIncentivizer events raised by the Yamv3 contract.
type Yamv3NewIncentivizerIterator struct {
	Event *Yamv3NewIncentivizer // Event containing the contract specifics and raw log

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
func (it *Yamv3NewIncentivizerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yamv3NewIncentivizer)
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
		it.Event = new(Yamv3NewIncentivizer)
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
func (it *Yamv3NewIncentivizerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yamv3NewIncentivizerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yamv3NewIncentivizer represents a NewIncentivizer event raised by the Yamv3 contract.
type Yamv3NewIncentivizer struct {
	OldIncentivizer common.Address
	NewIncentivizer common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewIncentivizer is a free log retrieval operation binding the contract event 0x2ee668ca7d17a9122dc00c0bfd73b684f2f7d681f17733cc02b294f69f1b3896.
//
// Solidity: event NewIncentivizer(address oldIncentivizer, address newIncentivizer)
func (_Yamv3 *Yamv3Filterer) FilterNewIncentivizer(opts *bind.FilterOpts) (*Yamv3NewIncentivizerIterator, error) {

	logs, sub, err := _Yamv3.contract.FilterLogs(opts, "NewIncentivizer")
	if err != nil {
		return nil, err
	}
	return &Yamv3NewIncentivizerIterator{contract: _Yamv3.contract, event: "NewIncentivizer", logs: logs, sub: sub}, nil
}

// WatchNewIncentivizer is a free log subscription operation binding the contract event 0x2ee668ca7d17a9122dc00c0bfd73b684f2f7d681f17733cc02b294f69f1b3896.
//
// Solidity: event NewIncentivizer(address oldIncentivizer, address newIncentivizer)
func (_Yamv3 *Yamv3Filterer) WatchNewIncentivizer(opts *bind.WatchOpts, sink chan<- *Yamv3NewIncentivizer) (event.Subscription, error) {

	logs, sub, err := _Yamv3.contract.WatchLogs(opts, "NewIncentivizer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yamv3NewIncentivizer)
				if err := _Yamv3.contract.UnpackLog(event, "NewIncentivizer", log); err != nil {
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

// ParseNewIncentivizer is a log parse operation binding the contract event 0x2ee668ca7d17a9122dc00c0bfd73b684f2f7d681f17733cc02b294f69f1b3896.
//
// Solidity: event NewIncentivizer(address oldIncentivizer, address newIncentivizer)
func (_Yamv3 *Yamv3Filterer) ParseNewIncentivizer(log types.Log) (*Yamv3NewIncentivizer, error) {
	event := new(Yamv3NewIncentivizer)
	if err := _Yamv3.contract.UnpackLog(event, "NewIncentivizer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yamv3NewMigratorIterator is returned from FilterNewMigrator and is used to iterate over the raw logs and unpacked data for NewMigrator events raised by the Yamv3 contract.
type Yamv3NewMigratorIterator struct {
	Event *Yamv3NewMigrator // Event containing the contract specifics and raw log

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
func (it *Yamv3NewMigratorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yamv3NewMigrator)
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
		it.Event = new(Yamv3NewMigrator)
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
func (it *Yamv3NewMigratorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yamv3NewMigratorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yamv3NewMigrator represents a NewMigrator event raised by the Yamv3 contract.
type Yamv3NewMigrator struct {
	OldMigrator common.Address
	NewMigrator common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewMigrator is a free log retrieval operation binding the contract event 0x99b2b7456799067566d467831e63363500739eac62c12ccb8cf9745078f06d2a.
//
// Solidity: event NewMigrator(address oldMigrator, address newMigrator)
func (_Yamv3 *Yamv3Filterer) FilterNewMigrator(opts *bind.FilterOpts) (*Yamv3NewMigratorIterator, error) {

	logs, sub, err := _Yamv3.contract.FilterLogs(opts, "NewMigrator")
	if err != nil {
		return nil, err
	}
	return &Yamv3NewMigratorIterator{contract: _Yamv3.contract, event: "NewMigrator", logs: logs, sub: sub}, nil
}

// WatchNewMigrator is a free log subscription operation binding the contract event 0x99b2b7456799067566d467831e63363500739eac62c12ccb8cf9745078f06d2a.
//
// Solidity: event NewMigrator(address oldMigrator, address newMigrator)
func (_Yamv3 *Yamv3Filterer) WatchNewMigrator(opts *bind.WatchOpts, sink chan<- *Yamv3NewMigrator) (event.Subscription, error) {

	logs, sub, err := _Yamv3.contract.WatchLogs(opts, "NewMigrator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yamv3NewMigrator)
				if err := _Yamv3.contract.UnpackLog(event, "NewMigrator", log); err != nil {
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

// ParseNewMigrator is a log parse operation binding the contract event 0x99b2b7456799067566d467831e63363500739eac62c12ccb8cf9745078f06d2a.
//
// Solidity: event NewMigrator(address oldMigrator, address newMigrator)
func (_Yamv3 *Yamv3Filterer) ParseNewMigrator(log types.Log) (*Yamv3NewMigrator, error) {
	event := new(Yamv3NewMigrator)
	if err := _Yamv3.contract.UnpackLog(event, "NewMigrator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yamv3NewPendingGovIterator is returned from FilterNewPendingGov and is used to iterate over the raw logs and unpacked data for NewPendingGov events raised by the Yamv3 contract.
type Yamv3NewPendingGovIterator struct {
	Event *Yamv3NewPendingGov // Event containing the contract specifics and raw log

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
func (it *Yamv3NewPendingGovIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yamv3NewPendingGov)
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
		it.Event = new(Yamv3NewPendingGov)
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
func (it *Yamv3NewPendingGovIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yamv3NewPendingGovIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yamv3NewPendingGov represents a NewPendingGov event raised by the Yamv3 contract.
type Yamv3NewPendingGov struct {
	OldPendingGov common.Address
	NewPendingGov common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewPendingGov is a free log retrieval operation binding the contract event 0x6163d5b9efd962645dd649e6e48a61bcb0f9df00997a2398b80d135a9ab0c61e.
//
// Solidity: event NewPendingGov(address oldPendingGov, address newPendingGov)
func (_Yamv3 *Yamv3Filterer) FilterNewPendingGov(opts *bind.FilterOpts) (*Yamv3NewPendingGovIterator, error) {

	logs, sub, err := _Yamv3.contract.FilterLogs(opts, "NewPendingGov")
	if err != nil {
		return nil, err
	}
	return &Yamv3NewPendingGovIterator{contract: _Yamv3.contract, event: "NewPendingGov", logs: logs, sub: sub}, nil
}

// WatchNewPendingGov is a free log subscription operation binding the contract event 0x6163d5b9efd962645dd649e6e48a61bcb0f9df00997a2398b80d135a9ab0c61e.
//
// Solidity: event NewPendingGov(address oldPendingGov, address newPendingGov)
func (_Yamv3 *Yamv3Filterer) WatchNewPendingGov(opts *bind.WatchOpts, sink chan<- *Yamv3NewPendingGov) (event.Subscription, error) {

	logs, sub, err := _Yamv3.contract.WatchLogs(opts, "NewPendingGov")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yamv3NewPendingGov)
				if err := _Yamv3.contract.UnpackLog(event, "NewPendingGov", log); err != nil {
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

// ParseNewPendingGov is a log parse operation binding the contract event 0x6163d5b9efd962645dd649e6e48a61bcb0f9df00997a2398b80d135a9ab0c61e.
//
// Solidity: event NewPendingGov(address oldPendingGov, address newPendingGov)
func (_Yamv3 *Yamv3Filterer) ParseNewPendingGov(log types.Log) (*Yamv3NewPendingGov, error) {
	event := new(Yamv3NewPendingGov)
	if err := _Yamv3.contract.UnpackLog(event, "NewPendingGov", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yamv3NewRebaserIterator is returned from FilterNewRebaser and is used to iterate over the raw logs and unpacked data for NewRebaser events raised by the Yamv3 contract.
type Yamv3NewRebaserIterator struct {
	Event *Yamv3NewRebaser // Event containing the contract specifics and raw log

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
func (it *Yamv3NewRebaserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yamv3NewRebaser)
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
		it.Event = new(Yamv3NewRebaser)
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
func (it *Yamv3NewRebaserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yamv3NewRebaserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yamv3NewRebaser represents a NewRebaser event raised by the Yamv3 contract.
type Yamv3NewRebaser struct {
	OldRebaser common.Address
	NewRebaser common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewRebaser is a free log retrieval operation binding the contract event 0x51f448520e2183de499e224808a409ee01a1f380edb2e8497572320c15030545.
//
// Solidity: event NewRebaser(address oldRebaser, address newRebaser)
func (_Yamv3 *Yamv3Filterer) FilterNewRebaser(opts *bind.FilterOpts) (*Yamv3NewRebaserIterator, error) {

	logs, sub, err := _Yamv3.contract.FilterLogs(opts, "NewRebaser")
	if err != nil {
		return nil, err
	}
	return &Yamv3NewRebaserIterator{contract: _Yamv3.contract, event: "NewRebaser", logs: logs, sub: sub}, nil
}

// WatchNewRebaser is a free log subscription operation binding the contract event 0x51f448520e2183de499e224808a409ee01a1f380edb2e8497572320c15030545.
//
// Solidity: event NewRebaser(address oldRebaser, address newRebaser)
func (_Yamv3 *Yamv3Filterer) WatchNewRebaser(opts *bind.WatchOpts, sink chan<- *Yamv3NewRebaser) (event.Subscription, error) {

	logs, sub, err := _Yamv3.contract.WatchLogs(opts, "NewRebaser")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yamv3NewRebaser)
				if err := _Yamv3.contract.UnpackLog(event, "NewRebaser", log); err != nil {
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

// ParseNewRebaser is a log parse operation binding the contract event 0x51f448520e2183de499e224808a409ee01a1f380edb2e8497572320c15030545.
//
// Solidity: event NewRebaser(address oldRebaser, address newRebaser)
func (_Yamv3 *Yamv3Filterer) ParseNewRebaser(log types.Log) (*Yamv3NewRebaser, error) {
	event := new(Yamv3NewRebaser)
	if err := _Yamv3.contract.UnpackLog(event, "NewRebaser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yamv3RebaseIterator is returned from FilterRebase and is used to iterate over the raw logs and unpacked data for Rebase events raised by the Yamv3 contract.
type Yamv3RebaseIterator struct {
	Event *Yamv3Rebase // Event containing the contract specifics and raw log

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
func (it *Yamv3RebaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yamv3Rebase)
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
		it.Event = new(Yamv3Rebase)
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
func (it *Yamv3RebaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yamv3RebaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yamv3Rebase represents a Rebase event raised by the Yamv3 contract.
type Yamv3Rebase struct {
	Epoch                 *big.Int
	PrevYamsScalingFactor *big.Int
	NewYamsScalingFactor  *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRebase is a free log retrieval operation binding the contract event 0xc6642d24d84e7f3d36ca39f5cce10e75639d9b158d5193aa350e2f900653e4c0.
//
// Solidity: event Rebase(uint256 epoch, uint256 prevYamsScalingFactor, uint256 newYamsScalingFactor)
func (_Yamv3 *Yamv3Filterer) FilterRebase(opts *bind.FilterOpts) (*Yamv3RebaseIterator, error) {

	logs, sub, err := _Yamv3.contract.FilterLogs(opts, "Rebase")
	if err != nil {
		return nil, err
	}
	return &Yamv3RebaseIterator{contract: _Yamv3.contract, event: "Rebase", logs: logs, sub: sub}, nil
}

// WatchRebase is a free log subscription operation binding the contract event 0xc6642d24d84e7f3d36ca39f5cce10e75639d9b158d5193aa350e2f900653e4c0.
//
// Solidity: event Rebase(uint256 epoch, uint256 prevYamsScalingFactor, uint256 newYamsScalingFactor)
func (_Yamv3 *Yamv3Filterer) WatchRebase(opts *bind.WatchOpts, sink chan<- *Yamv3Rebase) (event.Subscription, error) {

	logs, sub, err := _Yamv3.contract.WatchLogs(opts, "Rebase")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yamv3Rebase)
				if err := _Yamv3.contract.UnpackLog(event, "Rebase", log); err != nil {
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

// ParseRebase is a log parse operation binding the contract event 0xc6642d24d84e7f3d36ca39f5cce10e75639d9b158d5193aa350e2f900653e4c0.
//
// Solidity: event Rebase(uint256 epoch, uint256 prevYamsScalingFactor, uint256 newYamsScalingFactor)
func (_Yamv3 *Yamv3Filterer) ParseRebase(log types.Log) (*Yamv3Rebase, error) {
	event := new(Yamv3Rebase)
	if err := _Yamv3.contract.UnpackLog(event, "Rebase", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yamv3TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Yamv3 contract.
type Yamv3TransferIterator struct {
	Event *Yamv3Transfer // Event containing the contract specifics and raw log

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
func (it *Yamv3TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yamv3Transfer)
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
		it.Event = new(Yamv3Transfer)
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
func (it *Yamv3TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yamv3TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yamv3Transfer represents a Transfer event raised by the Yamv3 contract.
type Yamv3Transfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Yamv3 *Yamv3Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Yamv3TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Yamv3.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Yamv3TransferIterator{contract: _Yamv3.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Yamv3 *Yamv3Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Yamv3Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Yamv3.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yamv3Transfer)
				if err := _Yamv3.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Yamv3 *Yamv3Filterer) ParseTransfer(log types.Log) (*Yamv3Transfer, error) {
	event := new(Yamv3Transfer)
	if err := _Yamv3.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
