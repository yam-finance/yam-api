// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IndexStakingRewards

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

// IndexStakingRewardsABI is the input ABI used to generate the binding from.
const IndexStakingRewardsABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsDistribution\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardForDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodFinish\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerTokenStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsDistribution\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IndexStakingRewards is an auto generated Go binding around an Ethereum contract.
type IndexStakingRewards struct {
	IndexStakingRewardsCaller     // Read-only binding to the contract
	IndexStakingRewardsTransactor // Write-only binding to the contract
	IndexStakingRewardsFilterer   // Log filterer for contract events
}

// IndexStakingRewardsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IndexStakingRewardsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IndexStakingRewardsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IndexStakingRewardsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IndexStakingRewardsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IndexStakingRewardsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IndexStakingRewardsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IndexStakingRewardsSession struct {
	Contract     *IndexStakingRewards // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IndexStakingRewardsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IndexStakingRewardsCallerSession struct {
	Contract *IndexStakingRewardsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IndexStakingRewardsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IndexStakingRewardsTransactorSession struct {
	Contract     *IndexStakingRewardsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IndexStakingRewardsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IndexStakingRewardsRaw struct {
	Contract *IndexStakingRewards // Generic contract binding to access the raw methods on
}

// IndexStakingRewardsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IndexStakingRewardsCallerRaw struct {
	Contract *IndexStakingRewardsCaller // Generic read-only contract binding to access the raw methods on
}

// IndexStakingRewardsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IndexStakingRewardsTransactorRaw struct {
	Contract *IndexStakingRewardsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIndexStakingRewards creates a new instance of IndexStakingRewards, bound to a specific deployed contract.
func NewIndexStakingRewards(address common.Address, backend bind.ContractBackend) (*IndexStakingRewards, error) {
	contract, err := bindIndexStakingRewards(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IndexStakingRewards{IndexStakingRewardsCaller: IndexStakingRewardsCaller{contract: contract}, IndexStakingRewardsTransactor: IndexStakingRewardsTransactor{contract: contract}, IndexStakingRewardsFilterer: IndexStakingRewardsFilterer{contract: contract}}, nil
}

// NewIndexStakingRewardsCaller creates a new read-only instance of IndexStakingRewards, bound to a specific deployed contract.
func NewIndexStakingRewardsCaller(address common.Address, caller bind.ContractCaller) (*IndexStakingRewardsCaller, error) {
	contract, err := bindIndexStakingRewards(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IndexStakingRewardsCaller{contract: contract}, nil
}

// NewIndexStakingRewardsTransactor creates a new write-only instance of IndexStakingRewards, bound to a specific deployed contract.
func NewIndexStakingRewardsTransactor(address common.Address, transactor bind.ContractTransactor) (*IndexStakingRewardsTransactor, error) {
	contract, err := bindIndexStakingRewards(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IndexStakingRewardsTransactor{contract: contract}, nil
}

// NewIndexStakingRewardsFilterer creates a new log filterer instance of IndexStakingRewards, bound to a specific deployed contract.
func NewIndexStakingRewardsFilterer(address common.Address, filterer bind.ContractFilterer) (*IndexStakingRewardsFilterer, error) {
	contract, err := bindIndexStakingRewards(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IndexStakingRewardsFilterer{contract: contract}, nil
}

// bindIndexStakingRewards binds a generic wrapper to an already deployed contract.
func bindIndexStakingRewards(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IndexStakingRewardsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IndexStakingRewards *IndexStakingRewardsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IndexStakingRewards.Contract.IndexStakingRewardsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IndexStakingRewards *IndexStakingRewardsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IndexStakingRewards.Contract.IndexStakingRewardsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IndexStakingRewards *IndexStakingRewardsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IndexStakingRewards.Contract.IndexStakingRewardsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IndexStakingRewards *IndexStakingRewardsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IndexStakingRewards.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IndexStakingRewards *IndexStakingRewardsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IndexStakingRewards.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IndexStakingRewards *IndexStakingRewardsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IndexStakingRewards.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IndexStakingRewards.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IndexStakingRewards.Contract.BalanceOf(&_IndexStakingRewards.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IndexStakingRewards.Contract.BalanceOf(&_IndexStakingRewards.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsCaller) Earned(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IndexStakingRewards.contract.Call(opts, &out, "earned", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsSession) Earned(account common.Address) (*big.Int, error) {
	return _IndexStakingRewards.Contract.Earned(&_IndexStakingRewards.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsCallerSession) Earned(account common.Address) (*big.Int, error) {
	return _IndexStakingRewards.Contract.Earned(&_IndexStakingRewards.CallOpts, account)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsCaller) GetRewardForDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IndexStakingRewards.contract.Call(opts, &out, "getRewardForDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsSession) GetRewardForDuration() (*big.Int, error) {
	return _IndexStakingRewards.Contract.GetRewardForDuration(&_IndexStakingRewards.CallOpts)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsCallerSession) GetRewardForDuration() (*big.Int, error) {
	return _IndexStakingRewards.Contract.GetRewardForDuration(&_IndexStakingRewards.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsCaller) LastTimeRewardApplicable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IndexStakingRewards.contract.Call(opts, &out, "lastTimeRewardApplicable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _IndexStakingRewards.Contract.LastTimeRewardApplicable(&_IndexStakingRewards.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsCallerSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _IndexStakingRewards.Contract.LastTimeRewardApplicable(&_IndexStakingRewards.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsCaller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IndexStakingRewards.contract.Call(opts, &out, "lastUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsSession) LastUpdateTime() (*big.Int, error) {
	return _IndexStakingRewards.Contract.LastUpdateTime(&_IndexStakingRewards.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsCallerSession) LastUpdateTime() (*big.Int, error) {
	return _IndexStakingRewards.Contract.LastUpdateTime(&_IndexStakingRewards.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsCaller) PeriodFinish(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IndexStakingRewards.contract.Call(opts, &out, "periodFinish")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsSession) PeriodFinish() (*big.Int, error) {
	return _IndexStakingRewards.Contract.PeriodFinish(&_IndexStakingRewards.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsCallerSession) PeriodFinish() (*big.Int, error) {
	return _IndexStakingRewards.Contract.PeriodFinish(&_IndexStakingRewards.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsCaller) RewardPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IndexStakingRewards.contract.Call(opts, &out, "rewardPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsSession) RewardPerToken() (*big.Int, error) {
	return _IndexStakingRewards.Contract.RewardPerToken(&_IndexStakingRewards.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsCallerSession) RewardPerToken() (*big.Int, error) {
	return _IndexStakingRewards.Contract.RewardPerToken(&_IndexStakingRewards.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsCaller) RewardPerTokenStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IndexStakingRewards.contract.Call(opts, &out, "rewardPerTokenStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsSession) RewardPerTokenStored() (*big.Int, error) {
	return _IndexStakingRewards.Contract.RewardPerTokenStored(&_IndexStakingRewards.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsCallerSession) RewardPerTokenStored() (*big.Int, error) {
	return _IndexStakingRewards.Contract.RewardPerTokenStored(&_IndexStakingRewards.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IndexStakingRewards.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsSession) RewardRate() (*big.Int, error) {
	return _IndexStakingRewards.Contract.RewardRate(&_IndexStakingRewards.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsCallerSession) RewardRate() (*big.Int, error) {
	return _IndexStakingRewards.Contract.RewardRate(&_IndexStakingRewards.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IndexStakingRewards.contract.Call(opts, &out, "rewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _IndexStakingRewards.Contract.Rewards(&_IndexStakingRewards.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsCallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _IndexStakingRewards.Contract.Rewards(&_IndexStakingRewards.CallOpts, arg0)
}

// RewardsDistribution is a free data retrieval call binding the contract method 0x3fc6df6e.
//
// Solidity: function rewardsDistribution() view returns(address)
func (_IndexStakingRewards *IndexStakingRewardsCaller) RewardsDistribution(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IndexStakingRewards.contract.Call(opts, &out, "rewardsDistribution")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsDistribution is a free data retrieval call binding the contract method 0x3fc6df6e.
//
// Solidity: function rewardsDistribution() view returns(address)
func (_IndexStakingRewards *IndexStakingRewardsSession) RewardsDistribution() (common.Address, error) {
	return _IndexStakingRewards.Contract.RewardsDistribution(&_IndexStakingRewards.CallOpts)
}

// RewardsDistribution is a free data retrieval call binding the contract method 0x3fc6df6e.
//
// Solidity: function rewardsDistribution() view returns(address)
func (_IndexStakingRewards *IndexStakingRewardsCallerSession) RewardsDistribution() (common.Address, error) {
	return _IndexStakingRewards.Contract.RewardsDistribution(&_IndexStakingRewards.CallOpts)
}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsCaller) RewardsDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IndexStakingRewards.contract.Call(opts, &out, "rewardsDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsSession) RewardsDuration() (*big.Int, error) {
	return _IndexStakingRewards.Contract.RewardsDuration(&_IndexStakingRewards.CallOpts)
}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsCallerSession) RewardsDuration() (*big.Int, error) {
	return _IndexStakingRewards.Contract.RewardsDuration(&_IndexStakingRewards.CallOpts)
}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_IndexStakingRewards *IndexStakingRewardsCaller) RewardsToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IndexStakingRewards.contract.Call(opts, &out, "rewardsToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_IndexStakingRewards *IndexStakingRewardsSession) RewardsToken() (common.Address, error) {
	return _IndexStakingRewards.Contract.RewardsToken(&_IndexStakingRewards.CallOpts)
}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_IndexStakingRewards *IndexStakingRewardsCallerSession) RewardsToken() (common.Address, error) {
	return _IndexStakingRewards.Contract.RewardsToken(&_IndexStakingRewards.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_IndexStakingRewards *IndexStakingRewardsCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IndexStakingRewards.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_IndexStakingRewards *IndexStakingRewardsSession) StakingToken() (common.Address, error) {
	return _IndexStakingRewards.Contract.StakingToken(&_IndexStakingRewards.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_IndexStakingRewards *IndexStakingRewardsCallerSession) StakingToken() (common.Address, error) {
	return _IndexStakingRewards.Contract.StakingToken(&_IndexStakingRewards.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IndexStakingRewards.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsSession) TotalSupply() (*big.Int, error) {
	return _IndexStakingRewards.Contract.TotalSupply(&_IndexStakingRewards.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsCallerSession) TotalSupply() (*big.Int, error) {
	return _IndexStakingRewards.Contract.TotalSupply(&_IndexStakingRewards.CallOpts)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsCaller) UserRewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IndexStakingRewards.contract.Call(opts, &out, "userRewardPerTokenPaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _IndexStakingRewards.Contract.UserRewardPerTokenPaid(&_IndexStakingRewards.CallOpts, arg0)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_IndexStakingRewards *IndexStakingRewardsCallerSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _IndexStakingRewards.Contract.UserRewardPerTokenPaid(&_IndexStakingRewards.CallOpts, arg0)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_IndexStakingRewards *IndexStakingRewardsTransactor) Exit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IndexStakingRewards.contract.Transact(opts, "exit")
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_IndexStakingRewards *IndexStakingRewardsSession) Exit() (*types.Transaction, error) {
	return _IndexStakingRewards.Contract.Exit(&_IndexStakingRewards.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_IndexStakingRewards *IndexStakingRewardsTransactorSession) Exit() (*types.Transaction, error) {
	return _IndexStakingRewards.Contract.Exit(&_IndexStakingRewards.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_IndexStakingRewards *IndexStakingRewardsTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IndexStakingRewards.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_IndexStakingRewards *IndexStakingRewardsSession) GetReward() (*types.Transaction, error) {
	return _IndexStakingRewards.Contract.GetReward(&_IndexStakingRewards.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_IndexStakingRewards *IndexStakingRewardsTransactorSession) GetReward() (*types.Transaction, error) {
	return _IndexStakingRewards.Contract.GetReward(&_IndexStakingRewards.TransactOpts)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_IndexStakingRewards *IndexStakingRewardsTransactor) NotifyRewardAmount(opts *bind.TransactOpts, reward *big.Int) (*types.Transaction, error) {
	return _IndexStakingRewards.contract.Transact(opts, "notifyRewardAmount", reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_IndexStakingRewards *IndexStakingRewardsSession) NotifyRewardAmount(reward *big.Int) (*types.Transaction, error) {
	return _IndexStakingRewards.Contract.NotifyRewardAmount(&_IndexStakingRewards.TransactOpts, reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_IndexStakingRewards *IndexStakingRewardsTransactorSession) NotifyRewardAmount(reward *big.Int) (*types.Transaction, error) {
	return _IndexStakingRewards.Contract.NotifyRewardAmount(&_IndexStakingRewards.TransactOpts, reward)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_IndexStakingRewards *IndexStakingRewardsTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IndexStakingRewards.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_IndexStakingRewards *IndexStakingRewardsSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _IndexStakingRewards.Contract.Stake(&_IndexStakingRewards.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_IndexStakingRewards *IndexStakingRewardsTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _IndexStakingRewards.Contract.Stake(&_IndexStakingRewards.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_IndexStakingRewards *IndexStakingRewardsTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IndexStakingRewards.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_IndexStakingRewards *IndexStakingRewardsSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _IndexStakingRewards.Contract.Withdraw(&_IndexStakingRewards.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_IndexStakingRewards *IndexStakingRewardsTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _IndexStakingRewards.Contract.Withdraw(&_IndexStakingRewards.TransactOpts, amount)
}

// IndexStakingRewardsRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the IndexStakingRewards contract.
type IndexStakingRewardsRewardAddedIterator struct {
	Event *IndexStakingRewardsRewardAdded // Event containing the contract specifics and raw log

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
func (it *IndexStakingRewardsRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IndexStakingRewardsRewardAdded)
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
		it.Event = new(IndexStakingRewardsRewardAdded)
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
func (it *IndexStakingRewardsRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IndexStakingRewardsRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IndexStakingRewardsRewardAdded represents a RewardAdded event raised by the IndexStakingRewards contract.
type IndexStakingRewardsRewardAdded struct {
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_IndexStakingRewards *IndexStakingRewardsFilterer) FilterRewardAdded(opts *bind.FilterOpts) (*IndexStakingRewardsRewardAddedIterator, error) {

	logs, sub, err := _IndexStakingRewards.contract.FilterLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return &IndexStakingRewardsRewardAddedIterator{contract: _IndexStakingRewards.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_IndexStakingRewards *IndexStakingRewardsFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *IndexStakingRewardsRewardAdded) (event.Subscription, error) {

	logs, sub, err := _IndexStakingRewards.contract.WatchLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IndexStakingRewardsRewardAdded)
				if err := _IndexStakingRewards.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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

// ParseRewardAdded is a log parse operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_IndexStakingRewards *IndexStakingRewardsFilterer) ParseRewardAdded(log types.Log) (*IndexStakingRewardsRewardAdded, error) {
	event := new(IndexStakingRewardsRewardAdded)
	if err := _IndexStakingRewards.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IndexStakingRewardsRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the IndexStakingRewards contract.
type IndexStakingRewardsRewardPaidIterator struct {
	Event *IndexStakingRewardsRewardPaid // Event containing the contract specifics and raw log

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
func (it *IndexStakingRewardsRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IndexStakingRewardsRewardPaid)
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
		it.Event = new(IndexStakingRewardsRewardPaid)
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
func (it *IndexStakingRewardsRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IndexStakingRewardsRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IndexStakingRewardsRewardPaid represents a RewardPaid event raised by the IndexStakingRewards contract.
type IndexStakingRewardsRewardPaid struct {
	User   common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_IndexStakingRewards *IndexStakingRewardsFilterer) FilterRewardPaid(opts *bind.FilterOpts, user []common.Address) (*IndexStakingRewardsRewardPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _IndexStakingRewards.contract.FilterLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return &IndexStakingRewardsRewardPaidIterator{contract: _IndexStakingRewards.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_IndexStakingRewards *IndexStakingRewardsFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *IndexStakingRewardsRewardPaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _IndexStakingRewards.contract.WatchLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IndexStakingRewardsRewardPaid)
				if err := _IndexStakingRewards.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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

// ParseRewardPaid is a log parse operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_IndexStakingRewards *IndexStakingRewardsFilterer) ParseRewardPaid(log types.Log) (*IndexStakingRewardsRewardPaid, error) {
	event := new(IndexStakingRewardsRewardPaid)
	if err := _IndexStakingRewards.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IndexStakingRewardsStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the IndexStakingRewards contract.
type IndexStakingRewardsStakedIterator struct {
	Event *IndexStakingRewardsStaked // Event containing the contract specifics and raw log

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
func (it *IndexStakingRewardsStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IndexStakingRewardsStaked)
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
		it.Event = new(IndexStakingRewardsStaked)
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
func (it *IndexStakingRewardsStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IndexStakingRewardsStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IndexStakingRewardsStaked represents a Staked event raised by the IndexStakingRewards contract.
type IndexStakingRewardsStaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_IndexStakingRewards *IndexStakingRewardsFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*IndexStakingRewardsStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _IndexStakingRewards.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &IndexStakingRewardsStakedIterator{contract: _IndexStakingRewards.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_IndexStakingRewards *IndexStakingRewardsFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *IndexStakingRewardsStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _IndexStakingRewards.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IndexStakingRewardsStaked)
				if err := _IndexStakingRewards.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_IndexStakingRewards *IndexStakingRewardsFilterer) ParseStaked(log types.Log) (*IndexStakingRewardsStaked, error) {
	event := new(IndexStakingRewardsStaked)
	if err := _IndexStakingRewards.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IndexStakingRewardsWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the IndexStakingRewards contract.
type IndexStakingRewardsWithdrawnIterator struct {
	Event *IndexStakingRewardsWithdrawn // Event containing the contract specifics and raw log

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
func (it *IndexStakingRewardsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IndexStakingRewardsWithdrawn)
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
		it.Event = new(IndexStakingRewardsWithdrawn)
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
func (it *IndexStakingRewardsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IndexStakingRewardsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IndexStakingRewardsWithdrawn represents a Withdrawn event raised by the IndexStakingRewards contract.
type IndexStakingRewardsWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_IndexStakingRewards *IndexStakingRewardsFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*IndexStakingRewardsWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _IndexStakingRewards.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &IndexStakingRewardsWithdrawnIterator{contract: _IndexStakingRewards.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_IndexStakingRewards *IndexStakingRewardsFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *IndexStakingRewardsWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _IndexStakingRewards.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IndexStakingRewardsWithdrawn)
				if err := _IndexStakingRewards.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_IndexStakingRewards *IndexStakingRewardsFilterer) ParseWithdrawn(log types.Log) (*IndexStakingRewardsWithdrawn, error) {
	event := new(IndexStakingRewardsWithdrawn)
	if err := _IndexStakingRewards.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
