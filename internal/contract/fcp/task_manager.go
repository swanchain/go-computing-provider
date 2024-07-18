// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fcp

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
	_ = abi.ConvertType
)

// TaskManagerFullTaskInfo is an auto generated low-level Go binding around an user-defined struct.
type TaskManagerFullTaskInfo struct {
	TaskUid            string
	CpList             []common.Address
	User               common.Address
	Reward             *big.Int
	Collateral         *big.Int
	StartTimestamp     *big.Int
	TerminateTimestamp *big.Int
	Duration           *big.Int
	TaskStatus         uint8
	CollateralStatus   uint8
}

// EcpTaskManagerMetaData contains all meta data concerning the EcpTaskManager contract.
var EcpTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"name\":\"RewardReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"TaskCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cpList\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"addedDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDuration\",\"type\":\"uint256\"}],\"name\":\"TaskExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"terminateTimestamp\",\"type\":\"uint256\"}],\"name\":\"TaskTerminated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"URISubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"apWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"arWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"cpAccountList\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"assignTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralContract\",\"outputs\":[{\"internalType\":\"contractCollateralV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"completeTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"extraDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"extraReward\",\"type\":\"uint256\"}],\"name\":\"extendTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"getTaskInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"cpList\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"terminateTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"enumTaskManager.TaskStatus\",\"name\":\"taskStatus\",\"type\":\"uint8\"},{\"internalType\":\"enumCollateralV2.CollateralStatus\",\"name\":\"collateralStatus\",\"type\":\"uint8\"}],\"internalType\":\"structTaskManager.FullTaskInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ar\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ap\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardTokenAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"releaseReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCollateralContract\",\"type\":\"address\"}],\"name\":\"setCollateralContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"resultUri\",\"type\":\"string\"}],\"name\":\"submitResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"terminateTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"enumTaskManager.TaskStatus\",\"name\":\"taskStatus\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"terminateTime\",\"type\":\"uint256\"}],\"name\":\"terminateTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// EcpTaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use EcpTaskManagerMetaData.ABI instead.
var EcpTaskManagerABI = EcpTaskManagerMetaData.ABI

// EcpTaskManager is an auto generated Go binding around an Ethereum contract.
type EcpTaskManager struct {
	EcpTaskManagerCaller     // Read-only binding to the contract
	EcpTaskManagerTransactor // Write-only binding to the contract
	EcpTaskManagerFilterer   // Log filterer for contract events
}

// EcpTaskManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type EcpTaskManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EcpTaskManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EcpTaskManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EcpTaskManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EcpTaskManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EcpTaskManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EcpTaskManagerSession struct {
	Contract     *EcpTaskManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EcpTaskManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EcpTaskManagerCallerSession struct {
	Contract *EcpTaskManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// EcpTaskManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EcpTaskManagerTransactorSession struct {
	Contract     *EcpTaskManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// EcpTaskManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type EcpTaskManagerRaw struct {
	Contract *EcpTaskManager // Generic contract binding to access the raw methods on
}

// EcpTaskManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EcpTaskManagerCallerRaw struct {
	Contract *EcpTaskManagerCaller // Generic read-only contract binding to access the raw methods on
}

// EcpTaskManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EcpTaskManagerTransactorRaw struct {
	Contract *EcpTaskManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEcpTaskManager creates a new instance of EcpTaskManager, bound to a specific deployed contract.
func NewEcpTaskManager(address common.Address, backend bind.ContractBackend) (*EcpTaskManager, error) {
	contract, err := bindEcpTaskManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EcpTaskManager{EcpTaskManagerCaller: EcpTaskManagerCaller{contract: contract}, EcpTaskManagerTransactor: EcpTaskManagerTransactor{contract: contract}, EcpTaskManagerFilterer: EcpTaskManagerFilterer{contract: contract}}, nil
}

// NewEcpTaskManagerCaller creates a new read-only instance of EcpTaskManager, bound to a specific deployed contract.
func NewEcpTaskManagerCaller(address common.Address, caller bind.ContractCaller) (*EcpTaskManagerCaller, error) {
	contract, err := bindEcpTaskManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EcpTaskManagerCaller{contract: contract}, nil
}

// NewEcpTaskManagerTransactor creates a new write-only instance of EcpTaskManager, bound to a specific deployed contract.
func NewEcpTaskManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*EcpTaskManagerTransactor, error) {
	contract, err := bindEcpTaskManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EcpTaskManagerTransactor{contract: contract}, nil
}

// NewEcpTaskManagerFilterer creates a new log filterer instance of EcpTaskManager, bound to a specific deployed contract.
func NewEcpTaskManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*EcpTaskManagerFilterer, error) {
	contract, err := bindEcpTaskManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EcpTaskManagerFilterer{contract: contract}, nil
}

// bindEcpTaskManager binds a generic wrapper to an already deployed contract.
func bindEcpTaskManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EcpTaskManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EcpTaskManager *EcpTaskManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EcpTaskManager.Contract.EcpTaskManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EcpTaskManager *EcpTaskManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.EcpTaskManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EcpTaskManager *EcpTaskManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.EcpTaskManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EcpTaskManager *EcpTaskManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EcpTaskManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EcpTaskManager *EcpTaskManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EcpTaskManager *EcpTaskManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.contract.Transact(opts, method, params...)
}

// ApWallet is a free data retrieval call binding the contract method 0x4dc74d22.
//
// Solidity: function apWallet() view returns(address)
func (_EcpTaskManager *EcpTaskManagerCaller) ApWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EcpTaskManager.contract.Call(opts, &out, "apWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ApWallet is a free data retrieval call binding the contract method 0x4dc74d22.
//
// Solidity: function apWallet() view returns(address)
func (_EcpTaskManager *EcpTaskManagerSession) ApWallet() (common.Address, error) {
	return _EcpTaskManager.Contract.ApWallet(&_EcpTaskManager.CallOpts)
}

// ApWallet is a free data retrieval call binding the contract method 0x4dc74d22.
//
// Solidity: function apWallet() view returns(address)
func (_EcpTaskManager *EcpTaskManagerCallerSession) ApWallet() (common.Address, error) {
	return _EcpTaskManager.Contract.ApWallet(&_EcpTaskManager.CallOpts)
}

// ArWallet is a free data retrieval call binding the contract method 0x69ab84f2.
//
// Solidity: function arWallet() view returns(address)
func (_EcpTaskManager *EcpTaskManagerCaller) ArWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EcpTaskManager.contract.Call(opts, &out, "arWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ArWallet is a free data retrieval call binding the contract method 0x69ab84f2.
//
// Solidity: function arWallet() view returns(address)
func (_EcpTaskManager *EcpTaskManagerSession) ArWallet() (common.Address, error) {
	return _EcpTaskManager.Contract.ArWallet(&_EcpTaskManager.CallOpts)
}

// ArWallet is a free data retrieval call binding the contract method 0x69ab84f2.
//
// Solidity: function arWallet() view returns(address)
func (_EcpTaskManager *EcpTaskManagerCallerSession) ArWallet() (common.Address, error) {
	return _EcpTaskManager.Contract.ArWallet(&_EcpTaskManager.CallOpts)
}

// CollateralContract is a free data retrieval call binding the contract method 0xc6e1c7c9.
//
// Solidity: function collateralContract() view returns(address)
func (_EcpTaskManager *EcpTaskManagerCaller) CollateralContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EcpTaskManager.contract.Call(opts, &out, "collateralContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralContract is a free data retrieval call binding the contract method 0xc6e1c7c9.
//
// Solidity: function collateralContract() view returns(address)
func (_EcpTaskManager *EcpTaskManagerSession) CollateralContract() (common.Address, error) {
	return _EcpTaskManager.Contract.CollateralContract(&_EcpTaskManager.CallOpts)
}

// CollateralContract is a free data retrieval call binding the contract method 0xc6e1c7c9.
//
// Solidity: function collateralContract() view returns(address)
func (_EcpTaskManager *EcpTaskManagerCallerSession) CollateralContract() (common.Address, error) {
	return _EcpTaskManager.Contract.CollateralContract(&_EcpTaskManager.CallOpts)
}

// GetTaskInfo is a free data retrieval call binding the contract method 0x1b209463.
//
// Solidity: function getTaskInfo(string taskUid) view returns((string,address[],address,uint256,uint256,uint256,uint256,uint256,uint8,uint8))
func (_EcpTaskManager *EcpTaskManagerCaller) GetTaskInfo(opts *bind.CallOpts, taskUid string) (TaskManagerFullTaskInfo, error) {
	var out []interface{}
	err := _EcpTaskManager.contract.Call(opts, &out, "getTaskInfo", taskUid)

	if err != nil {
		return *new(TaskManagerFullTaskInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(TaskManagerFullTaskInfo)).(*TaskManagerFullTaskInfo)

	return out0, err

}

// GetTaskInfo is a free data retrieval call binding the contract method 0x1b209463.
//
// Solidity: function getTaskInfo(string taskUid) view returns((string,address[],address,uint256,uint256,uint256,uint256,uint256,uint8,uint8))
func (_EcpTaskManager *EcpTaskManagerSession) GetTaskInfo(taskUid string) (TaskManagerFullTaskInfo, error) {
	return _EcpTaskManager.Contract.GetTaskInfo(&_EcpTaskManager.CallOpts, taskUid)
}

// GetTaskInfo is a free data retrieval call binding the contract method 0x1b209463.
//
// Solidity: function getTaskInfo(string taskUid) view returns((string,address[],address,uint256,uint256,uint256,uint256,uint256,uint8,uint8))
func (_EcpTaskManager *EcpTaskManagerCallerSession) GetTaskInfo(taskUid string) (TaskManagerFullTaskInfo, error) {
	return _EcpTaskManager.Contract.GetTaskInfo(&_EcpTaskManager.CallOpts, taskUid)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_EcpTaskManager *EcpTaskManagerCaller) IsAdmin(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _EcpTaskManager.contract.Call(opts, &out, "isAdmin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_EcpTaskManager *EcpTaskManagerSession) IsAdmin(arg0 common.Address) (bool, error) {
	return _EcpTaskManager.Contract.IsAdmin(&_EcpTaskManager.CallOpts, arg0)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_EcpTaskManager *EcpTaskManagerCallerSession) IsAdmin(arg0 common.Address) (bool, error) {
	return _EcpTaskManager.Contract.IsAdmin(&_EcpTaskManager.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EcpTaskManager *EcpTaskManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EcpTaskManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EcpTaskManager *EcpTaskManagerSession) Owner() (common.Address, error) {
	return _EcpTaskManager.Contract.Owner(&_EcpTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EcpTaskManager *EcpTaskManagerCallerSession) Owner() (common.Address, error) {
	return _EcpTaskManager.Contract.Owner(&_EcpTaskManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EcpTaskManager *EcpTaskManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EcpTaskManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EcpTaskManager *EcpTaskManagerSession) ProxiableUUID() ([32]byte, error) {
	return _EcpTaskManager.Contract.ProxiableUUID(&_EcpTaskManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EcpTaskManager *EcpTaskManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _EcpTaskManager.Contract.ProxiableUUID(&_EcpTaskManager.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_EcpTaskManager *EcpTaskManagerCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EcpTaskManager.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_EcpTaskManager *EcpTaskManagerSession) RewardToken() (common.Address, error) {
	return _EcpTaskManager.Contract.RewardToken(&_EcpTaskManager.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_EcpTaskManager *EcpTaskManagerCallerSession) RewardToken() (common.Address, error) {
	return _EcpTaskManager.Contract.RewardToken(&_EcpTaskManager.CallOpts)
}

// Tasks is a free data retrieval call binding the contract method 0x58709cf2.
//
// Solidity: function tasks(string ) view returns(string taskUid, address user, uint256 reward, uint256 startTimestamp, uint256 terminateTimestamp, uint256 duration, uint8 taskStatus)
func (_EcpTaskManager *EcpTaskManagerCaller) Tasks(opts *bind.CallOpts, arg0 string) (struct {
	TaskUid            string
	User               common.Address
	Reward             *big.Int
	StartTimestamp     *big.Int
	TerminateTimestamp *big.Int
	Duration           *big.Int
	TaskStatus         uint8
}, error) {
	var out []interface{}
	err := _EcpTaskManager.contract.Call(opts, &out, "tasks", arg0)

	outstruct := new(struct {
		TaskUid            string
		User               common.Address
		Reward             *big.Int
		StartTimestamp     *big.Int
		TerminateTimestamp *big.Int
		Duration           *big.Int
		TaskStatus         uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TaskUid = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.User = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Reward = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartTimestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TerminateTimestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TaskStatus = *abi.ConvertType(out[6], new(uint8)).(*uint8)

	return *outstruct, err

}

// Tasks is a free data retrieval call binding the contract method 0x58709cf2.
//
// Solidity: function tasks(string ) view returns(string taskUid, address user, uint256 reward, uint256 startTimestamp, uint256 terminateTimestamp, uint256 duration, uint8 taskStatus)
func (_EcpTaskManager *EcpTaskManagerSession) Tasks(arg0 string) (struct {
	TaskUid            string
	User               common.Address
	Reward             *big.Int
	StartTimestamp     *big.Int
	TerminateTimestamp *big.Int
	Duration           *big.Int
	TaskStatus         uint8
}, error) {
	return _EcpTaskManager.Contract.Tasks(&_EcpTaskManager.CallOpts, arg0)
}

// Tasks is a free data retrieval call binding the contract method 0x58709cf2.
//
// Solidity: function tasks(string ) view returns(string taskUid, address user, uint256 reward, uint256 startTimestamp, uint256 terminateTimestamp, uint256 duration, uint8 taskStatus)
func (_EcpTaskManager *EcpTaskManagerCallerSession) Tasks(arg0 string) (struct {
	TaskUid            string
	User               common.Address
	Reward             *big.Int
	StartTimestamp     *big.Int
	TerminateTimestamp *big.Int
	Duration           *big.Int
	TaskStatus         uint8
}, error) {
	return _EcpTaskManager.Contract.Tasks(&_EcpTaskManager.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_EcpTaskManager *EcpTaskManagerCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EcpTaskManager.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_EcpTaskManager *EcpTaskManagerSession) Version() (*big.Int, error) {
	return _EcpTaskManager.Contract.Version(&_EcpTaskManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_EcpTaskManager *EcpTaskManagerCallerSession) Version() (*big.Int, error) {
	return _EcpTaskManager.Contract.Version(&_EcpTaskManager.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address newAdmin) returns()
func (_EcpTaskManager *EcpTaskManagerTransactor) AddAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _EcpTaskManager.contract.Transact(opts, "addAdmin", newAdmin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address newAdmin) returns()
func (_EcpTaskManager *EcpTaskManagerSession) AddAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.AddAdmin(&_EcpTaskManager.TransactOpts, newAdmin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address newAdmin) returns()
func (_EcpTaskManager *EcpTaskManagerTransactorSession) AddAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.AddAdmin(&_EcpTaskManager.TransactOpts, newAdmin)
}

// AssignTask is a paid mutator transaction binding the contract method 0x9a8afe9c.
//
// Solidity: function assignTask(string taskUid, address[] cpAccountList, address user, uint256 reward, uint256 collateral, uint256 duration) returns()
func (_EcpTaskManager *EcpTaskManagerTransactor) AssignTask(opts *bind.TransactOpts, taskUid string, cpAccountList []common.Address, user common.Address, reward *big.Int, collateral *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _EcpTaskManager.contract.Transact(opts, "assignTask", taskUid, cpAccountList, user, reward, collateral, duration)
}

// AssignTask is a paid mutator transaction binding the contract method 0x9a8afe9c.
//
// Solidity: function assignTask(string taskUid, address[] cpAccountList, address user, uint256 reward, uint256 collateral, uint256 duration) returns()
func (_EcpTaskManager *EcpTaskManagerSession) AssignTask(taskUid string, cpAccountList []common.Address, user common.Address, reward *big.Int, collateral *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.AssignTask(&_EcpTaskManager.TransactOpts, taskUid, cpAccountList, user, reward, collateral, duration)
}

// AssignTask is a paid mutator transaction binding the contract method 0x9a8afe9c.
//
// Solidity: function assignTask(string taskUid, address[] cpAccountList, address user, uint256 reward, uint256 collateral, uint256 duration) returns()
func (_EcpTaskManager *EcpTaskManagerTransactorSession) AssignTask(taskUid string, cpAccountList []common.Address, user common.Address, reward *big.Int, collateral *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.AssignTask(&_EcpTaskManager.TransactOpts, taskUid, cpAccountList, user, reward, collateral, duration)
}

// CompleteTask is a paid mutator transaction binding the contract method 0x394c244b.
//
// Solidity: function completeTask(string taskUid) returns()
func (_EcpTaskManager *EcpTaskManagerTransactor) CompleteTask(opts *bind.TransactOpts, taskUid string) (*types.Transaction, error) {
	return _EcpTaskManager.contract.Transact(opts, "completeTask", taskUid)
}

// CompleteTask is a paid mutator transaction binding the contract method 0x394c244b.
//
// Solidity: function completeTask(string taskUid) returns()
func (_EcpTaskManager *EcpTaskManagerSession) CompleteTask(taskUid string) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.CompleteTask(&_EcpTaskManager.TransactOpts, taskUid)
}

// CompleteTask is a paid mutator transaction binding the contract method 0x394c244b.
//
// Solidity: function completeTask(string taskUid) returns()
func (_EcpTaskManager *EcpTaskManagerTransactorSession) CompleteTask(taskUid string) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.CompleteTask(&_EcpTaskManager.TransactOpts, taskUid)
}

// ExtendTask is a paid mutator transaction binding the contract method 0x3b97f5b3.
//
// Solidity: function extendTask(string taskUid, uint256 extraDuration, uint256 extraReward) returns()
func (_EcpTaskManager *EcpTaskManagerTransactor) ExtendTask(opts *bind.TransactOpts, taskUid string, extraDuration *big.Int, extraReward *big.Int) (*types.Transaction, error) {
	return _EcpTaskManager.contract.Transact(opts, "extendTask", taskUid, extraDuration, extraReward)
}

// ExtendTask is a paid mutator transaction binding the contract method 0x3b97f5b3.
//
// Solidity: function extendTask(string taskUid, uint256 extraDuration, uint256 extraReward) returns()
func (_EcpTaskManager *EcpTaskManagerSession) ExtendTask(taskUid string, extraDuration *big.Int, extraReward *big.Int) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.ExtendTask(&_EcpTaskManager.TransactOpts, taskUid, extraDuration, extraReward)
}

// ExtendTask is a paid mutator transaction binding the contract method 0x3b97f5b3.
//
// Solidity: function extendTask(string taskUid, uint256 extraDuration, uint256 extraReward) returns()
func (_EcpTaskManager *EcpTaskManagerTransactorSession) ExtendTask(taskUid string, extraDuration *big.Int, extraReward *big.Int) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.ExtendTask(&_EcpTaskManager.TransactOpts, taskUid, extraDuration, extraReward)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address ar, address ap, address collateralContractAddress, address rewardTokenAddress) returns()
func (_EcpTaskManager *EcpTaskManagerTransactor) Initialize(opts *bind.TransactOpts, ar common.Address, ap common.Address, collateralContractAddress common.Address, rewardTokenAddress common.Address) (*types.Transaction, error) {
	return _EcpTaskManager.contract.Transact(opts, "initialize", ar, ap, collateralContractAddress, rewardTokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address ar, address ap, address collateralContractAddress, address rewardTokenAddress) returns()
func (_EcpTaskManager *EcpTaskManagerSession) Initialize(ar common.Address, ap common.Address, collateralContractAddress common.Address, rewardTokenAddress common.Address) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.Initialize(&_EcpTaskManager.TransactOpts, ar, ap, collateralContractAddress, rewardTokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address ar, address ap, address collateralContractAddress, address rewardTokenAddress) returns()
func (_EcpTaskManager *EcpTaskManagerTransactorSession) Initialize(ar common.Address, ap common.Address, collateralContractAddress common.Address, rewardTokenAddress common.Address) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.Initialize(&_EcpTaskManager.TransactOpts, ar, ap, collateralContractAddress, rewardTokenAddress)
}

// ReleaseReward is a paid mutator transaction binding the contract method 0xfe8d3188.
//
// Solidity: function releaseReward(string taskUid) returns()
func (_EcpTaskManager *EcpTaskManagerTransactor) ReleaseReward(opts *bind.TransactOpts, taskUid string) (*types.Transaction, error) {
	return _EcpTaskManager.contract.Transact(opts, "releaseReward", taskUid)
}

// ReleaseReward is a paid mutator transaction binding the contract method 0xfe8d3188.
//
// Solidity: function releaseReward(string taskUid) returns()
func (_EcpTaskManager *EcpTaskManagerSession) ReleaseReward(taskUid string) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.ReleaseReward(&_EcpTaskManager.TransactOpts, taskUid)
}

// ReleaseReward is a paid mutator transaction binding the contract method 0xfe8d3188.
//
// Solidity: function releaseReward(string taskUid) returns()
func (_EcpTaskManager *EcpTaskManagerTransactorSession) ReleaseReward(taskUid string) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.ReleaseReward(&_EcpTaskManager.TransactOpts, taskUid)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_EcpTaskManager *EcpTaskManagerTransactor) RemoveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _EcpTaskManager.contract.Transact(opts, "removeAdmin", admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_EcpTaskManager *EcpTaskManagerSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.RemoveAdmin(&_EcpTaskManager.TransactOpts, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_EcpTaskManager *EcpTaskManagerTransactorSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.RemoveAdmin(&_EcpTaskManager.TransactOpts, admin)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EcpTaskManager *EcpTaskManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EcpTaskManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EcpTaskManager *EcpTaskManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _EcpTaskManager.Contract.RenounceOwnership(&_EcpTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EcpTaskManager *EcpTaskManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EcpTaskManager.Contract.RenounceOwnership(&_EcpTaskManager.TransactOpts)
}

// SetCollateralContract is a paid mutator transaction binding the contract method 0xe941cd06.
//
// Solidity: function setCollateralContract(address newCollateralContract) returns()
func (_EcpTaskManager *EcpTaskManagerTransactor) SetCollateralContract(opts *bind.TransactOpts, newCollateralContract common.Address) (*types.Transaction, error) {
	return _EcpTaskManager.contract.Transact(opts, "setCollateralContract", newCollateralContract)
}

// SetCollateralContract is a paid mutator transaction binding the contract method 0xe941cd06.
//
// Solidity: function setCollateralContract(address newCollateralContract) returns()
func (_EcpTaskManager *EcpTaskManagerSession) SetCollateralContract(newCollateralContract common.Address) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.SetCollateralContract(&_EcpTaskManager.TransactOpts, newCollateralContract)
}

// SetCollateralContract is a paid mutator transaction binding the contract method 0xe941cd06.
//
// Solidity: function setCollateralContract(address newCollateralContract) returns()
func (_EcpTaskManager *EcpTaskManagerTransactorSession) SetCollateralContract(newCollateralContract common.Address) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.SetCollateralContract(&_EcpTaskManager.TransactOpts, newCollateralContract)
}

// SubmitResult is a paid mutator transaction binding the contract method 0x1b622fa8.
//
// Solidity: function submitResult(string taskUid, string resultUri) returns()
func (_EcpTaskManager *EcpTaskManagerTransactor) SubmitResult(opts *bind.TransactOpts, taskUid string, resultUri string) (*types.Transaction, error) {
	return _EcpTaskManager.contract.Transact(opts, "submitResult", taskUid, resultUri)
}

// SubmitResult is a paid mutator transaction binding the contract method 0x1b622fa8.
//
// Solidity: function submitResult(string taskUid, string resultUri) returns()
func (_EcpTaskManager *EcpTaskManagerSession) SubmitResult(taskUid string, resultUri string) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.SubmitResult(&_EcpTaskManager.TransactOpts, taskUid, resultUri)
}

// SubmitResult is a paid mutator transaction binding the contract method 0x1b622fa8.
//
// Solidity: function submitResult(string taskUid, string resultUri) returns()
func (_EcpTaskManager *EcpTaskManagerTransactorSession) SubmitResult(taskUid string, resultUri string) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.SubmitResult(&_EcpTaskManager.TransactOpts, taskUid, resultUri)
}

// TerminateTask is a paid mutator transaction binding the contract method 0xed7bb9d1.
//
// Solidity: function terminateTask(string taskUid, uint256 terminateTime) returns()
func (_EcpTaskManager *EcpTaskManagerTransactor) TerminateTask(opts *bind.TransactOpts, taskUid string, terminateTime *big.Int) (*types.Transaction, error) {
	return _EcpTaskManager.contract.Transact(opts, "terminateTask", taskUid, terminateTime)
}

// TerminateTask is a paid mutator transaction binding the contract method 0xed7bb9d1.
//
// Solidity: function terminateTask(string taskUid, uint256 terminateTime) returns()
func (_EcpTaskManager *EcpTaskManagerSession) TerminateTask(taskUid string, terminateTime *big.Int) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.TerminateTask(&_EcpTaskManager.TransactOpts, taskUid, terminateTime)
}

// TerminateTask is a paid mutator transaction binding the contract method 0xed7bb9d1.
//
// Solidity: function terminateTask(string taskUid, uint256 terminateTime) returns()
func (_EcpTaskManager *EcpTaskManagerTransactorSession) TerminateTask(taskUid string, terminateTime *big.Int) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.TerminateTask(&_EcpTaskManager.TransactOpts, taskUid, terminateTime)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EcpTaskManager *EcpTaskManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EcpTaskManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EcpTaskManager *EcpTaskManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.TransferOwnership(&_EcpTaskManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EcpTaskManager *EcpTaskManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.TransferOwnership(&_EcpTaskManager.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_EcpTaskManager *EcpTaskManagerTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _EcpTaskManager.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_EcpTaskManager *EcpTaskManagerSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.UpgradeTo(&_EcpTaskManager.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_EcpTaskManager *EcpTaskManagerTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.UpgradeTo(&_EcpTaskManager.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EcpTaskManager *EcpTaskManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EcpTaskManager.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EcpTaskManager *EcpTaskManagerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.UpgradeToAndCall(&_EcpTaskManager.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EcpTaskManager *EcpTaskManagerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EcpTaskManager.Contract.UpgradeToAndCall(&_EcpTaskManager.TransactOpts, newImplementation, data)
}

// EcpTaskManagerAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the EcpTaskManager contract.
type EcpTaskManagerAdminChangedIterator struct {
	Event *EcpTaskManagerAdminChanged // Event containing the contract specifics and raw log

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
func (it *EcpTaskManagerAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpTaskManagerAdminChanged)
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
		it.Event = new(EcpTaskManagerAdminChanged)
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
func (it *EcpTaskManagerAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpTaskManagerAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpTaskManagerAdminChanged represents a AdminChanged event raised by the EcpTaskManager contract.
type EcpTaskManagerAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_EcpTaskManager *EcpTaskManagerFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*EcpTaskManagerAdminChangedIterator, error) {

	logs, sub, err := _EcpTaskManager.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &EcpTaskManagerAdminChangedIterator{contract: _EcpTaskManager.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_EcpTaskManager *EcpTaskManagerFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *EcpTaskManagerAdminChanged) (event.Subscription, error) {

	logs, sub, err := _EcpTaskManager.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpTaskManagerAdminChanged)
				if err := _EcpTaskManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_EcpTaskManager *EcpTaskManagerFilterer) ParseAdminChanged(log types.Log) (*EcpTaskManagerAdminChanged, error) {
	event := new(EcpTaskManagerAdminChanged)
	if err := _EcpTaskManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcpTaskManagerBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the EcpTaskManager contract.
type EcpTaskManagerBeaconUpgradedIterator struct {
	Event *EcpTaskManagerBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *EcpTaskManagerBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpTaskManagerBeaconUpgraded)
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
		it.Event = new(EcpTaskManagerBeaconUpgraded)
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
func (it *EcpTaskManagerBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpTaskManagerBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpTaskManagerBeaconUpgraded represents a BeaconUpgraded event raised by the EcpTaskManager contract.
type EcpTaskManagerBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_EcpTaskManager *EcpTaskManagerFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*EcpTaskManagerBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _EcpTaskManager.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &EcpTaskManagerBeaconUpgradedIterator{contract: _EcpTaskManager.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_EcpTaskManager *EcpTaskManagerFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *EcpTaskManagerBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _EcpTaskManager.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpTaskManagerBeaconUpgraded)
				if err := _EcpTaskManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_EcpTaskManager *EcpTaskManagerFilterer) ParseBeaconUpgraded(log types.Log) (*EcpTaskManagerBeaconUpgraded, error) {
	event := new(EcpTaskManagerBeaconUpgraded)
	if err := _EcpTaskManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcpTaskManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the EcpTaskManager contract.
type EcpTaskManagerInitializedIterator struct {
	Event *EcpTaskManagerInitialized // Event containing the contract specifics and raw log

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
func (it *EcpTaskManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpTaskManagerInitialized)
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
		it.Event = new(EcpTaskManagerInitialized)
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
func (it *EcpTaskManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpTaskManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpTaskManagerInitialized represents a Initialized event raised by the EcpTaskManager contract.
type EcpTaskManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EcpTaskManager *EcpTaskManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*EcpTaskManagerInitializedIterator, error) {

	logs, sub, err := _EcpTaskManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EcpTaskManagerInitializedIterator{contract: _EcpTaskManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EcpTaskManager *EcpTaskManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EcpTaskManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _EcpTaskManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpTaskManagerInitialized)
				if err := _EcpTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_EcpTaskManager *EcpTaskManagerFilterer) ParseInitialized(log types.Log) (*EcpTaskManagerInitialized, error) {
	event := new(EcpTaskManagerInitialized)
	if err := _EcpTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcpTaskManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EcpTaskManager contract.
type EcpTaskManagerOwnershipTransferredIterator struct {
	Event *EcpTaskManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EcpTaskManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpTaskManagerOwnershipTransferred)
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
		it.Event = new(EcpTaskManagerOwnershipTransferred)
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
func (it *EcpTaskManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpTaskManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpTaskManagerOwnershipTransferred represents a OwnershipTransferred event raised by the EcpTaskManager contract.
type EcpTaskManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EcpTaskManager *EcpTaskManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EcpTaskManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EcpTaskManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EcpTaskManagerOwnershipTransferredIterator{contract: _EcpTaskManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EcpTaskManager *EcpTaskManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EcpTaskManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EcpTaskManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpTaskManagerOwnershipTransferred)
				if err := _EcpTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EcpTaskManager *EcpTaskManagerFilterer) ParseOwnershipTransferred(log types.Log) (*EcpTaskManagerOwnershipTransferred, error) {
	event := new(EcpTaskManagerOwnershipTransferred)
	if err := _EcpTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcpTaskManagerRewardReleasedIterator is returned from FilterRewardReleased and is used to iterate over the raw logs and unpacked data for RewardReleased events raised by the EcpTaskManager contract.
type EcpTaskManagerRewardReleasedIterator struct {
	Event *EcpTaskManagerRewardReleased // Event containing the contract specifics and raw log

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
func (it *EcpTaskManagerRewardReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpTaskManagerRewardReleased)
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
		it.Event = new(EcpTaskManagerRewardReleased)
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
func (it *EcpTaskManagerRewardReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpTaskManagerRewardReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpTaskManagerRewardReleased represents a RewardReleased event raised by the EcpTaskManager contract.
type EcpTaskManagerRewardReleased struct {
	TaskUid      common.Hash
	CpAccount    common.Address
	Beneficiary  common.Address
	RewardAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRewardReleased is a free log retrieval operation binding the contract event 0x6504a330caacb9bebf5a13c7a7356d83a2ff093b07185671301738c1761475d2.
//
// Solidity: event RewardReleased(string indexed taskUid, address indexed cpAccount, address beneficiary, uint256 rewardAmount)
func (_EcpTaskManager *EcpTaskManagerFilterer) FilterRewardReleased(opts *bind.FilterOpts, taskUid []string, cpAccount []common.Address) (*EcpTaskManagerRewardReleasedIterator, error) {

	var taskUidRule []interface{}
	for _, taskUidItem := range taskUid {
		taskUidRule = append(taskUidRule, taskUidItem)
	}
	var cpAccountRule []interface{}
	for _, cpAccountItem := range cpAccount {
		cpAccountRule = append(cpAccountRule, cpAccountItem)
	}

	logs, sub, err := _EcpTaskManager.contract.FilterLogs(opts, "RewardReleased", taskUidRule, cpAccountRule)
	if err != nil {
		return nil, err
	}
	return &EcpTaskManagerRewardReleasedIterator{contract: _EcpTaskManager.contract, event: "RewardReleased", logs: logs, sub: sub}, nil
}

// WatchRewardReleased is a free log subscription operation binding the contract event 0x6504a330caacb9bebf5a13c7a7356d83a2ff093b07185671301738c1761475d2.
//
// Solidity: event RewardReleased(string indexed taskUid, address indexed cpAccount, address beneficiary, uint256 rewardAmount)
func (_EcpTaskManager *EcpTaskManagerFilterer) WatchRewardReleased(opts *bind.WatchOpts, sink chan<- *EcpTaskManagerRewardReleased, taskUid []string, cpAccount []common.Address) (event.Subscription, error) {

	var taskUidRule []interface{}
	for _, taskUidItem := range taskUid {
		taskUidRule = append(taskUidRule, taskUidItem)
	}
	var cpAccountRule []interface{}
	for _, cpAccountItem := range cpAccount {
		cpAccountRule = append(cpAccountRule, cpAccountItem)
	}

	logs, sub, err := _EcpTaskManager.contract.WatchLogs(opts, "RewardReleased", taskUidRule, cpAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpTaskManagerRewardReleased)
				if err := _EcpTaskManager.contract.UnpackLog(event, "RewardReleased", log); err != nil {
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

// ParseRewardReleased is a log parse operation binding the contract event 0x6504a330caacb9bebf5a13c7a7356d83a2ff093b07185671301738c1761475d2.
//
// Solidity: event RewardReleased(string indexed taskUid, address indexed cpAccount, address beneficiary, uint256 rewardAmount)
func (_EcpTaskManager *EcpTaskManagerFilterer) ParseRewardReleased(log types.Log) (*EcpTaskManagerRewardReleased, error) {
	event := new(EcpTaskManagerRewardReleased)
	if err := _EcpTaskManager.contract.UnpackLog(event, "RewardReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcpTaskManagerTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the EcpTaskManager contract.
type EcpTaskManagerTaskCompletedIterator struct {
	Event *EcpTaskManagerTaskCompleted // Event containing the contract specifics and raw log

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
func (it *EcpTaskManagerTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpTaskManagerTaskCompleted)
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
		it.Event = new(EcpTaskManagerTaskCompleted)
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
func (it *EcpTaskManagerTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpTaskManagerTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpTaskManagerTaskCompleted represents a TaskCompleted event raised by the EcpTaskManager contract.
type EcpTaskManagerTaskCompleted struct {
	TaskUid common.Hash
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x2ea2764b7d80d0107c1d2c0210ead4a6fdee0f86235483a041d9faffbf875880.
//
// Solidity: event TaskCompleted(string indexed taskUid)
func (_EcpTaskManager *EcpTaskManagerFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskUid []string) (*EcpTaskManagerTaskCompletedIterator, error) {

	var taskUidRule []interface{}
	for _, taskUidItem := range taskUid {
		taskUidRule = append(taskUidRule, taskUidItem)
	}

	logs, sub, err := _EcpTaskManager.contract.FilterLogs(opts, "TaskCompleted", taskUidRule)
	if err != nil {
		return nil, err
	}
	return &EcpTaskManagerTaskCompletedIterator{contract: _EcpTaskManager.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x2ea2764b7d80d0107c1d2c0210ead4a6fdee0f86235483a041d9faffbf875880.
//
// Solidity: event TaskCompleted(string indexed taskUid)
func (_EcpTaskManager *EcpTaskManagerFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *EcpTaskManagerTaskCompleted, taskUid []string) (event.Subscription, error) {

	var taskUidRule []interface{}
	for _, taskUidItem := range taskUid {
		taskUidRule = append(taskUidRule, taskUidItem)
	}

	logs, sub, err := _EcpTaskManager.contract.WatchLogs(opts, "TaskCompleted", taskUidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpTaskManagerTaskCompleted)
				if err := _EcpTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
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

// ParseTaskCompleted is a log parse operation binding the contract event 0x2ea2764b7d80d0107c1d2c0210ead4a6fdee0f86235483a041d9faffbf875880.
//
// Solidity: event TaskCompleted(string indexed taskUid)
func (_EcpTaskManager *EcpTaskManagerFilterer) ParseTaskCompleted(log types.Log) (*EcpTaskManagerTaskCompleted, error) {
	event := new(EcpTaskManagerTaskCompleted)
	if err := _EcpTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcpTaskManagerTaskCreatedIterator is returned from FilterTaskCreated and is used to iterate over the raw logs and unpacked data for TaskCreated events raised by the EcpTaskManager contract.
type EcpTaskManagerTaskCreatedIterator struct {
	Event *EcpTaskManagerTaskCreated // Event containing the contract specifics and raw log

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
func (it *EcpTaskManagerTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpTaskManagerTaskCreated)
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
		it.Event = new(EcpTaskManagerTaskCreated)
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
func (it *EcpTaskManagerTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpTaskManagerTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpTaskManagerTaskCreated represents a TaskCreated event raised by the EcpTaskManager contract.
type EcpTaskManagerTaskCreated struct {
	TaskUid  common.Hash
	CpList   []common.Address
	User     common.Address
	Reward   *big.Int
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTaskCreated is a free log retrieval operation binding the contract event 0x1455e6e35b1b89916dd3024a995f41593b0b27f0e7cb61b253a69ba2c2d37951.
//
// Solidity: event TaskCreated(string indexed taskUid, address[] cpList, address user, uint256 reward, uint256 duration)
func (_EcpTaskManager *EcpTaskManagerFilterer) FilterTaskCreated(opts *bind.FilterOpts, taskUid []string) (*EcpTaskManagerTaskCreatedIterator, error) {

	var taskUidRule []interface{}
	for _, taskUidItem := range taskUid {
		taskUidRule = append(taskUidRule, taskUidItem)
	}

	logs, sub, err := _EcpTaskManager.contract.FilterLogs(opts, "TaskCreated", taskUidRule)
	if err != nil {
		return nil, err
	}
	return &EcpTaskManagerTaskCreatedIterator{contract: _EcpTaskManager.contract, event: "TaskCreated", logs: logs, sub: sub}, nil
}

// WatchTaskCreated is a free log subscription operation binding the contract event 0x1455e6e35b1b89916dd3024a995f41593b0b27f0e7cb61b253a69ba2c2d37951.
//
// Solidity: event TaskCreated(string indexed taskUid, address[] cpList, address user, uint256 reward, uint256 duration)
func (_EcpTaskManager *EcpTaskManagerFilterer) WatchTaskCreated(opts *bind.WatchOpts, sink chan<- *EcpTaskManagerTaskCreated, taskUid []string) (event.Subscription, error) {

	var taskUidRule []interface{}
	for _, taskUidItem := range taskUid {
		taskUidRule = append(taskUidRule, taskUidItem)
	}

	logs, sub, err := _EcpTaskManager.contract.WatchLogs(opts, "TaskCreated", taskUidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpTaskManagerTaskCreated)
				if err := _EcpTaskManager.contract.UnpackLog(event, "TaskCreated", log); err != nil {
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

// ParseTaskCreated is a log parse operation binding the contract event 0x1455e6e35b1b89916dd3024a995f41593b0b27f0e7cb61b253a69ba2c2d37951.
//
// Solidity: event TaskCreated(string indexed taskUid, address[] cpList, address user, uint256 reward, uint256 duration)
func (_EcpTaskManager *EcpTaskManagerFilterer) ParseTaskCreated(log types.Log) (*EcpTaskManagerTaskCreated, error) {
	event := new(EcpTaskManagerTaskCreated)
	if err := _EcpTaskManager.contract.UnpackLog(event, "TaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcpTaskManagerTaskExtendedIterator is returned from FilterTaskExtended and is used to iterate over the raw logs and unpacked data for TaskExtended events raised by the EcpTaskManager contract.
type EcpTaskManagerTaskExtendedIterator struct {
	Event *EcpTaskManagerTaskExtended // Event containing the contract specifics and raw log

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
func (it *EcpTaskManagerTaskExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpTaskManagerTaskExtended)
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
		it.Event = new(EcpTaskManagerTaskExtended)
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
func (it *EcpTaskManagerTaskExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpTaskManagerTaskExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpTaskManagerTaskExtended represents a TaskExtended event raised by the EcpTaskManager contract.
type EcpTaskManagerTaskExtended struct {
	TaskUid       common.Hash
	AddedDuration *big.Int
	TotalDuration *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTaskExtended is a free log retrieval operation binding the contract event 0xd1fa78153d824892404081883d06a6a4d9dd70c117aff5a8bf61c8961ffc941c.
//
// Solidity: event TaskExtended(string indexed taskUid, uint256 addedDuration, uint256 totalDuration)
func (_EcpTaskManager *EcpTaskManagerFilterer) FilterTaskExtended(opts *bind.FilterOpts, taskUid []string) (*EcpTaskManagerTaskExtendedIterator, error) {

	var taskUidRule []interface{}
	for _, taskUidItem := range taskUid {
		taskUidRule = append(taskUidRule, taskUidItem)
	}

	logs, sub, err := _EcpTaskManager.contract.FilterLogs(opts, "TaskExtended", taskUidRule)
	if err != nil {
		return nil, err
	}
	return &EcpTaskManagerTaskExtendedIterator{contract: _EcpTaskManager.contract, event: "TaskExtended", logs: logs, sub: sub}, nil
}

// WatchTaskExtended is a free log subscription operation binding the contract event 0xd1fa78153d824892404081883d06a6a4d9dd70c117aff5a8bf61c8961ffc941c.
//
// Solidity: event TaskExtended(string indexed taskUid, uint256 addedDuration, uint256 totalDuration)
func (_EcpTaskManager *EcpTaskManagerFilterer) WatchTaskExtended(opts *bind.WatchOpts, sink chan<- *EcpTaskManagerTaskExtended, taskUid []string) (event.Subscription, error) {

	var taskUidRule []interface{}
	for _, taskUidItem := range taskUid {
		taskUidRule = append(taskUidRule, taskUidItem)
	}

	logs, sub, err := _EcpTaskManager.contract.WatchLogs(opts, "TaskExtended", taskUidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpTaskManagerTaskExtended)
				if err := _EcpTaskManager.contract.UnpackLog(event, "TaskExtended", log); err != nil {
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

// ParseTaskExtended is a log parse operation binding the contract event 0xd1fa78153d824892404081883d06a6a4d9dd70c117aff5a8bf61c8961ffc941c.
//
// Solidity: event TaskExtended(string indexed taskUid, uint256 addedDuration, uint256 totalDuration)
func (_EcpTaskManager *EcpTaskManagerFilterer) ParseTaskExtended(log types.Log) (*EcpTaskManagerTaskExtended, error) {
	event := new(EcpTaskManagerTaskExtended)
	if err := _EcpTaskManager.contract.UnpackLog(event, "TaskExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcpTaskManagerTaskTerminatedIterator is returned from FilterTaskTerminated and is used to iterate over the raw logs and unpacked data for TaskTerminated events raised by the EcpTaskManager contract.
type EcpTaskManagerTaskTerminatedIterator struct {
	Event *EcpTaskManagerTaskTerminated // Event containing the contract specifics and raw log

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
func (it *EcpTaskManagerTaskTerminatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpTaskManagerTaskTerminated)
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
		it.Event = new(EcpTaskManagerTaskTerminated)
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
func (it *EcpTaskManagerTaskTerminatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpTaskManagerTaskTerminatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpTaskManagerTaskTerminated represents a TaskTerminated event raised by the EcpTaskManager contract.
type EcpTaskManagerTaskTerminated struct {
	TaskUid            common.Hash
	TerminateTimestamp *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTaskTerminated is a free log retrieval operation binding the contract event 0x58136d68c79c10fb55282098b850a5a1c041b28af29cb65ea657a824f5bd6a36.
//
// Solidity: event TaskTerminated(string indexed taskUid, uint256 terminateTimestamp)
func (_EcpTaskManager *EcpTaskManagerFilterer) FilterTaskTerminated(opts *bind.FilterOpts, taskUid []string) (*EcpTaskManagerTaskTerminatedIterator, error) {

	var taskUidRule []interface{}
	for _, taskUidItem := range taskUid {
		taskUidRule = append(taskUidRule, taskUidItem)
	}

	logs, sub, err := _EcpTaskManager.contract.FilterLogs(opts, "TaskTerminated", taskUidRule)
	if err != nil {
		return nil, err
	}
	return &EcpTaskManagerTaskTerminatedIterator{contract: _EcpTaskManager.contract, event: "TaskTerminated", logs: logs, sub: sub}, nil
}

// WatchTaskTerminated is a free log subscription operation binding the contract event 0x58136d68c79c10fb55282098b850a5a1c041b28af29cb65ea657a824f5bd6a36.
//
// Solidity: event TaskTerminated(string indexed taskUid, uint256 terminateTimestamp)
func (_EcpTaskManager *EcpTaskManagerFilterer) WatchTaskTerminated(opts *bind.WatchOpts, sink chan<- *EcpTaskManagerTaskTerminated, taskUid []string) (event.Subscription, error) {

	var taskUidRule []interface{}
	for _, taskUidItem := range taskUid {
		taskUidRule = append(taskUidRule, taskUidItem)
	}

	logs, sub, err := _EcpTaskManager.contract.WatchLogs(opts, "TaskTerminated", taskUidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpTaskManagerTaskTerminated)
				if err := _EcpTaskManager.contract.UnpackLog(event, "TaskTerminated", log); err != nil {
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

// ParseTaskTerminated is a log parse operation binding the contract event 0x58136d68c79c10fb55282098b850a5a1c041b28af29cb65ea657a824f5bd6a36.
//
// Solidity: event TaskTerminated(string indexed taskUid, uint256 terminateTimestamp)
func (_EcpTaskManager *EcpTaskManagerFilterer) ParseTaskTerminated(log types.Log) (*EcpTaskManagerTaskTerminated, error) {
	event := new(EcpTaskManagerTaskTerminated)
	if err := _EcpTaskManager.contract.UnpackLog(event, "TaskTerminated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcpTaskManagerURISubmittedIterator is returned from FilterURISubmitted and is used to iterate over the raw logs and unpacked data for URISubmitted events raised by the EcpTaskManager contract.
type EcpTaskManagerURISubmittedIterator struct {
	Event *EcpTaskManagerURISubmitted // Event containing the contract specifics and raw log

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
func (it *EcpTaskManagerURISubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpTaskManagerURISubmitted)
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
		it.Event = new(EcpTaskManagerURISubmitted)
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
func (it *EcpTaskManagerURISubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpTaskManagerURISubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpTaskManagerURISubmitted represents a URISubmitted event raised by the EcpTaskManager contract.
type EcpTaskManagerURISubmitted struct {
	TaskUid   common.Hash
	Submitter common.Address
	Uri       string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterURISubmitted is a free log retrieval operation binding the contract event 0x2b9a4b6cdefbc34e0fe2cf2e7c540e2382e53f5a2d4e7dfc6f26dc1c3ed3fe66.
//
// Solidity: event URISubmitted(string indexed taskUid, address submitter, string uri)
func (_EcpTaskManager *EcpTaskManagerFilterer) FilterURISubmitted(opts *bind.FilterOpts, taskUid []string) (*EcpTaskManagerURISubmittedIterator, error) {

	var taskUidRule []interface{}
	for _, taskUidItem := range taskUid {
		taskUidRule = append(taskUidRule, taskUidItem)
	}

	logs, sub, err := _EcpTaskManager.contract.FilterLogs(opts, "URISubmitted", taskUidRule)
	if err != nil {
		return nil, err
	}
	return &EcpTaskManagerURISubmittedIterator{contract: _EcpTaskManager.contract, event: "URISubmitted", logs: logs, sub: sub}, nil
}

// WatchURISubmitted is a free log subscription operation binding the contract event 0x2b9a4b6cdefbc34e0fe2cf2e7c540e2382e53f5a2d4e7dfc6f26dc1c3ed3fe66.
//
// Solidity: event URISubmitted(string indexed taskUid, address submitter, string uri)
func (_EcpTaskManager *EcpTaskManagerFilterer) WatchURISubmitted(opts *bind.WatchOpts, sink chan<- *EcpTaskManagerURISubmitted, taskUid []string) (event.Subscription, error) {

	var taskUidRule []interface{}
	for _, taskUidItem := range taskUid {
		taskUidRule = append(taskUidRule, taskUidItem)
	}

	logs, sub, err := _EcpTaskManager.contract.WatchLogs(opts, "URISubmitted", taskUidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpTaskManagerURISubmitted)
				if err := _EcpTaskManager.contract.UnpackLog(event, "URISubmitted", log); err != nil {
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

// ParseURISubmitted is a log parse operation binding the contract event 0x2b9a4b6cdefbc34e0fe2cf2e7c540e2382e53f5a2d4e7dfc6f26dc1c3ed3fe66.
//
// Solidity: event URISubmitted(string indexed taskUid, address submitter, string uri)
func (_EcpTaskManager *EcpTaskManagerFilterer) ParseURISubmitted(log types.Log) (*EcpTaskManagerURISubmitted, error) {
	event := new(EcpTaskManagerURISubmitted)
	if err := _EcpTaskManager.contract.UnpackLog(event, "URISubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcpTaskManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the EcpTaskManager contract.
type EcpTaskManagerUpgradedIterator struct {
	Event *EcpTaskManagerUpgraded // Event containing the contract specifics and raw log

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
func (it *EcpTaskManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpTaskManagerUpgraded)
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
		it.Event = new(EcpTaskManagerUpgraded)
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
func (it *EcpTaskManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpTaskManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpTaskManagerUpgraded represents a Upgraded event raised by the EcpTaskManager contract.
type EcpTaskManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EcpTaskManager *EcpTaskManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*EcpTaskManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EcpTaskManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &EcpTaskManagerUpgradedIterator{contract: _EcpTaskManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EcpTaskManager *EcpTaskManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *EcpTaskManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EcpTaskManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpTaskManagerUpgraded)
				if err := _EcpTaskManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_EcpTaskManager *EcpTaskManagerFilterer) ParseUpgraded(log types.Log) (*EcpTaskManagerUpgraded, error) {
	event := new(EcpTaskManagerUpgraded)
	if err := _EcpTaskManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
