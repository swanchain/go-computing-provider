// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ecp

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

// AggregatedTaskMetaData contains all meta data concerning the AggregatedTask contract.
var AggregatedTaskMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_taskBlobCID\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_taskRegistryContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taskContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"RegisteredToTaskRegistry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"taskBlobCID\",\"type\":\"string\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"taskBlobCID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051610c32380380610c328339818101604052810190610031919061047a565b5f825111610074576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161006b9061052e565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036100e2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100d9906105bc565b60405180910390fd5b815f90816100f091906107e7565b507fabf948a1223daba3f04a6050c5ee3bfb2b135790471a6bbf2d2a93bc29faaf9f8260405161012091906108ee565b60405180910390a16101378161013e60201b60201c565b5050610a2c565b5f8173ffffffffffffffffffffffffffffffffffffffff16303360405160240161016992919061091d565b6040516020818303038152906040527fcaa29fc2000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516101f39190610988565b5f604051808303815f865af19150503d805f811461022c576040519150601f19603f3d011682016040523d82523d5f602084013e610231565b606091505b5050905080610275576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026c90610a0e565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff167f411a83d4fcdb8a204895aa1c817c68da89892ae8a277620988dcd6ea44650b7f60405160405180910390a35050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610332826102ec565b810181811067ffffffffffffffff82111715610351576103506102fc565b5b80604052505050565b5f6103636102d3565b905061036f8282610329565b919050565b5f67ffffffffffffffff82111561038e5761038d6102fc565b5b610397826102ec565b9050602081019050919050565b8281835e5f83830152505050565b5f6103c46103bf84610374565b61035a565b9050828152602081018484840111156103e0576103df6102e8565b5b6103eb8482856103a4565b509392505050565b5f82601f830112610407576104066102e4565b5b81516104178482602086016103b2565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61044982610420565b9050919050565b6104598161043f565b8114610463575f80fd5b50565b5f8151905061047481610450565b92915050565b5f80604083850312156104905761048f6102dc565b5b5f83015167ffffffffffffffff8111156104ad576104ac6102e0565b5b6104b9858286016103f3565b92505060206104ca85828601610466565b9150509250929050565b5f82825260208201905092915050565b7f5461736b20426c6f6220434944206d757374206e6f7420626520656d707479005f82015250565b5f610518601f836104d4565b9150610523826104e4565b602082019050919050565b5f6020820190508181035f8301526105458161050c565b9050919050565b7f5461736b20726567697374727920636f6e74726163742061646472657373206d5f8201527f7573742062652070726f76696465640000000000000000000000000000000000602082015250565b5f6105a6602f836104d4565b91506105b18261054c565b604082019050919050565b5f6020820190508181035f8301526105d38161059a565b9050919050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061062857607f821691505b60208210810361063b5761063a6105e4565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261069d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610662565b6106a78683610662565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6106eb6106e66106e1846106bf565b6106c8565b6106bf565b9050919050565b5f819050919050565b610704836106d1565b610718610710826106f2565b84845461066e565b825550505050565b5f90565b61072c610720565b6107378184846106fb565b505050565b5b8181101561075a5761074f5f82610724565b60018101905061073d565b5050565b601f82111561079f5761077081610641565b61077984610653565b81016020851015610788578190505b61079c61079485610653565b83018261073c565b50505b505050565b5f82821c905092915050565b5f6107bf5f19846008026107a4565b1980831691505092915050565b5f6107d783836107b0565b9150826002028217905092915050565b6107f0826105da565b67ffffffffffffffff811115610809576108086102fc565b5b6108138254610611565b61081e82828561075e565b5f60209050601f83116001811461084f575f841561083d578287015190505b61084785826107cc565b8655506108ae565b601f19841661085d86610641565b5f5b828110156108845784890151825560018201915060208501945060208101905061085f565b868310156108a1578489015161089d601f8916826107b0565b8355505b6001600288020188555050505b505050505050565b5f6108c0826105da565b6108ca81856104d4565b93506108da8185602086016103a4565b6108e3816102ec565b840191505092915050565b5f6020820190508181035f83015261090681846108b6565b905092915050565b6109178161043f565b82525050565b5f6040820190506109305f83018561090e565b61093d602083018461090e565b9392505050565b5f81519050919050565b5f81905092915050565b5f61096282610944565b61096c818561094e565b935061097c8185602086016103a4565b80840191505092915050565b5f6109938284610958565b915081905092915050565b7f4661696c656420746f207265676973746572207461736b20636f6e74726163745f8201527f20746f205461736b526567697374727900000000000000000000000000000000602082015250565b5f6109f86030836104d4565b9150610a038261099e565b604082019050919050565b5f6020820190508181035f830152610a25816109ec565b9050919050565b6101f980610a395f395ff3fe608060405234801561000f575f80fd5b5060043610610029575f3560e01c8063cf9dc5571461002d575b5f80fd5b61003561004b565b6040516100429190610146565b60405180910390f35b5f805461005790610193565b80601f016020809104026020016040519081016040528092919081815260200182805461008390610193565b80156100ce5780601f106100a5576101008083540402835291602001916100ce565b820191905f5260205f20905b8154815290600101906020018083116100b157829003601f168201915b505050505081565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610118826100d6565b61012281856100e0565b93506101328185602086016100f0565b61013b816100fe565b840191505092915050565b5f6020820190508181035f83015261015e818461010e565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806101aa57607f821691505b6020821081036101bd576101bc610166565b5b5091905056fea264697066735822122087b7f47902850ce902706d3b8cbd2011c4815a02b83efdbe67cac9eb954dbaa064736f6c63430008190033",
}

// AggregatedTaskABI is the input ABI used to generate the binding from.
// Deprecated: Use AggregatedTaskMetaData.ABI instead.
var AggregatedTaskABI = AggregatedTaskMetaData.ABI

// AggregatedTaskBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AggregatedTaskMetaData.Bin instead.
var AggregatedTaskBin = AggregatedTaskMetaData.Bin

// DeployAggregatedTask deploys a new Ethereum contract, binding an instance of AggregatedTask to it.
func DeployAggregatedTask(auth *bind.TransactOpts, backend bind.ContractBackend, _taskBlobCID string, _taskRegistryContract common.Address) (common.Address, *types.Transaction, *AggregatedTask, error) {
	parsed, err := AggregatedTaskMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AggregatedTaskBin), backend, _taskBlobCID, _taskRegistryContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AggregatedTask{AggregatedTaskCaller: AggregatedTaskCaller{contract: contract}, AggregatedTaskTransactor: AggregatedTaskTransactor{contract: contract}, AggregatedTaskFilterer: AggregatedTaskFilterer{contract: contract}}, nil
}

// AggregatedTask is an auto generated Go binding around an Ethereum contract.
type AggregatedTask struct {
	AggregatedTaskCaller     // Read-only binding to the contract
	AggregatedTaskTransactor // Write-only binding to the contract
	AggregatedTaskFilterer   // Log filterer for contract events
}

// AggregatedTaskCaller is an auto generated read-only Go binding around an Ethereum contract.
type AggregatedTaskCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatedTaskTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AggregatedTaskTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatedTaskFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AggregatedTaskFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatedTaskSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AggregatedTaskSession struct {
	Contract     *AggregatedTask   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AggregatedTaskCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AggregatedTaskCallerSession struct {
	Contract *AggregatedTaskCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AggregatedTaskTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AggregatedTaskTransactorSession struct {
	Contract     *AggregatedTaskTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AggregatedTaskRaw is an auto generated low-level Go binding around an Ethereum contract.
type AggregatedTaskRaw struct {
	Contract *AggregatedTask // Generic contract binding to access the raw methods on
}

// AggregatedTaskCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AggregatedTaskCallerRaw struct {
	Contract *AggregatedTaskCaller // Generic read-only contract binding to access the raw methods on
}

// AggregatedTaskTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AggregatedTaskTransactorRaw struct {
	Contract *AggregatedTaskTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAggregatedTask creates a new instance of AggregatedTask, bound to a specific deployed contract.
func NewAggregatedTask(address common.Address, backend bind.ContractBackend) (*AggregatedTask, error) {
	contract, err := bindAggregatedTask(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AggregatedTask{AggregatedTaskCaller: AggregatedTaskCaller{contract: contract}, AggregatedTaskTransactor: AggregatedTaskTransactor{contract: contract}, AggregatedTaskFilterer: AggregatedTaskFilterer{contract: contract}}, nil
}

// NewAggregatedTaskCaller creates a new read-only instance of AggregatedTask, bound to a specific deployed contract.
func NewAggregatedTaskCaller(address common.Address, caller bind.ContractCaller) (*AggregatedTaskCaller, error) {
	contract, err := bindAggregatedTask(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatedTaskCaller{contract: contract}, nil
}

// NewAggregatedTaskTransactor creates a new write-only instance of AggregatedTask, bound to a specific deployed contract.
func NewAggregatedTaskTransactor(address common.Address, transactor bind.ContractTransactor) (*AggregatedTaskTransactor, error) {
	contract, err := bindAggregatedTask(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatedTaskTransactor{contract: contract}, nil
}

// NewAggregatedTaskFilterer creates a new log filterer instance of AggregatedTask, bound to a specific deployed contract.
func NewAggregatedTaskFilterer(address common.Address, filterer bind.ContractFilterer) (*AggregatedTaskFilterer, error) {
	contract, err := bindAggregatedTask(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AggregatedTaskFilterer{contract: contract}, nil
}

// bindAggregatedTask binds a generic wrapper to an already deployed contract.
func bindAggregatedTask(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AggregatedTaskMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AggregatedTask *AggregatedTaskRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregatedTask.Contract.AggregatedTaskCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AggregatedTask *AggregatedTaskRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregatedTask.Contract.AggregatedTaskTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AggregatedTask *AggregatedTaskRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregatedTask.Contract.AggregatedTaskTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AggregatedTask *AggregatedTaskCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregatedTask.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AggregatedTask *AggregatedTaskTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregatedTask.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AggregatedTask *AggregatedTaskTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregatedTask.Contract.contract.Transact(opts, method, params...)
}

// TaskBlobCID is a free data retrieval call binding the contract method 0xcf9dc557.
//
// Solidity: function taskBlobCID() view returns(string)
func (_AggregatedTask *AggregatedTaskCaller) TaskBlobCID(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AggregatedTask.contract.Call(opts, &out, "taskBlobCID")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TaskBlobCID is a free data retrieval call binding the contract method 0xcf9dc557.
//
// Solidity: function taskBlobCID() view returns(string)
func (_AggregatedTask *AggregatedTaskSession) TaskBlobCID() (string, error) {
	return _AggregatedTask.Contract.TaskBlobCID(&_AggregatedTask.CallOpts)
}

// TaskBlobCID is a free data retrieval call binding the contract method 0xcf9dc557.
//
// Solidity: function taskBlobCID() view returns(string)
func (_AggregatedTask *AggregatedTaskCallerSession) TaskBlobCID() (string, error) {
	return _AggregatedTask.Contract.TaskBlobCID(&_AggregatedTask.CallOpts)
}

// AggregatedTaskRegisteredToTaskRegistryIterator is returned from FilterRegisteredToTaskRegistry and is used to iterate over the raw logs and unpacked data for RegisteredToTaskRegistry events raised by the AggregatedTask contract.
type AggregatedTaskRegisteredToTaskRegistryIterator struct {
	Event *AggregatedTaskRegisteredToTaskRegistry // Event containing the contract specifics and raw log

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
func (it *AggregatedTaskRegisteredToTaskRegistryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatedTaskRegisteredToTaskRegistry)
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
		it.Event = new(AggregatedTaskRegisteredToTaskRegistry)
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
func (it *AggregatedTaskRegisteredToTaskRegistryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatedTaskRegisteredToTaskRegistryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatedTaskRegisteredToTaskRegistry represents a RegisteredToTaskRegistry event raised by the AggregatedTask contract.
type AggregatedTaskRegisteredToTaskRegistry struct {
	TaskContract common.Address
	Owner        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRegisteredToTaskRegistry is a free log retrieval operation binding the contract event 0x411a83d4fcdb8a204895aa1c817c68da89892ae8a277620988dcd6ea44650b7f.
//
// Solidity: event RegisteredToTaskRegistry(address indexed taskContract, address indexed owner)
func (_AggregatedTask *AggregatedTaskFilterer) FilterRegisteredToTaskRegistry(opts *bind.FilterOpts, taskContract []common.Address, owner []common.Address) (*AggregatedTaskRegisteredToTaskRegistryIterator, error) {

	var taskContractRule []interface{}
	for _, taskContractItem := range taskContract {
		taskContractRule = append(taskContractRule, taskContractItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _AggregatedTask.contract.FilterLogs(opts, "RegisteredToTaskRegistry", taskContractRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &AggregatedTaskRegisteredToTaskRegistryIterator{contract: _AggregatedTask.contract, event: "RegisteredToTaskRegistry", logs: logs, sub: sub}, nil
}

// WatchRegisteredToTaskRegistry is a free log subscription operation binding the contract event 0x411a83d4fcdb8a204895aa1c817c68da89892ae8a277620988dcd6ea44650b7f.
//
// Solidity: event RegisteredToTaskRegistry(address indexed taskContract, address indexed owner)
func (_AggregatedTask *AggregatedTaskFilterer) WatchRegisteredToTaskRegistry(opts *bind.WatchOpts, sink chan<- *AggregatedTaskRegisteredToTaskRegistry, taskContract []common.Address, owner []common.Address) (event.Subscription, error) {

	var taskContractRule []interface{}
	for _, taskContractItem := range taskContract {
		taskContractRule = append(taskContractRule, taskContractItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _AggregatedTask.contract.WatchLogs(opts, "RegisteredToTaskRegistry", taskContractRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatedTaskRegisteredToTaskRegistry)
				if err := _AggregatedTask.contract.UnpackLog(event, "RegisteredToTaskRegistry", log); err != nil {
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

// ParseRegisteredToTaskRegistry is a log parse operation binding the contract event 0x411a83d4fcdb8a204895aa1c817c68da89892ae8a277620988dcd6ea44650b7f.
//
// Solidity: event RegisteredToTaskRegistry(address indexed taskContract, address indexed owner)
func (_AggregatedTask *AggregatedTaskFilterer) ParseRegisteredToTaskRegistry(log types.Log) (*AggregatedTaskRegisteredToTaskRegistry, error) {
	event := new(AggregatedTaskRegisteredToTaskRegistry)
	if err := _AggregatedTask.contract.UnpackLog(event, "RegisteredToTaskRegistry", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatedTaskTaskCreatedIterator is returned from FilterTaskCreated and is used to iterate over the raw logs and unpacked data for TaskCreated events raised by the AggregatedTask contract.
type AggregatedTaskTaskCreatedIterator struct {
	Event *AggregatedTaskTaskCreated // Event containing the contract specifics and raw log

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
func (it *AggregatedTaskTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatedTaskTaskCreated)
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
		it.Event = new(AggregatedTaskTaskCreated)
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
func (it *AggregatedTaskTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatedTaskTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatedTaskTaskCreated represents a TaskCreated event raised by the AggregatedTask contract.
type AggregatedTaskTaskCreated struct {
	TaskBlobCID string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTaskCreated is a free log retrieval operation binding the contract event 0xabf948a1223daba3f04a6050c5ee3bfb2b135790471a6bbf2d2a93bc29faaf9f.
//
// Solidity: event TaskCreated(string taskBlobCID)
func (_AggregatedTask *AggregatedTaskFilterer) FilterTaskCreated(opts *bind.FilterOpts) (*AggregatedTaskTaskCreatedIterator, error) {

	logs, sub, err := _AggregatedTask.contract.FilterLogs(opts, "TaskCreated")
	if err != nil {
		return nil, err
	}
	return &AggregatedTaskTaskCreatedIterator{contract: _AggregatedTask.contract, event: "TaskCreated", logs: logs, sub: sub}, nil
}

// WatchTaskCreated is a free log subscription operation binding the contract event 0xabf948a1223daba3f04a6050c5ee3bfb2b135790471a6bbf2d2a93bc29faaf9f.
//
// Solidity: event TaskCreated(string taskBlobCID)
func (_AggregatedTask *AggregatedTaskFilterer) WatchTaskCreated(opts *bind.WatchOpts, sink chan<- *AggregatedTaskTaskCreated) (event.Subscription, error) {

	logs, sub, err := _AggregatedTask.contract.WatchLogs(opts, "TaskCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatedTaskTaskCreated)
				if err := _AggregatedTask.contract.UnpackLog(event, "TaskCreated", log); err != nil {
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

// ParseTaskCreated is a log parse operation binding the contract event 0xabf948a1223daba3f04a6050c5ee3bfb2b135790471a6bbf2d2a93bc29faaf9f.
//
// Solidity: event TaskCreated(string taskBlobCID)
func (_AggregatedTask *AggregatedTaskFilterer) ParseTaskCreated(log types.Log) (*AggregatedTaskTaskCreated, error) {
	event := new(AggregatedTaskTaskCreated)
	if err := _AggregatedTask.contract.UnpackLog(event, "TaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
