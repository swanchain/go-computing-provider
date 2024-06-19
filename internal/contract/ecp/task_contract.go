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

// TaskMetaData contains all meta data concerning the Task contract.
var TaskMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_taskID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_taskType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_resourceType\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_inputParam\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_verifyParam\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_cpAccount\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_proof\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_taskRegistryContract\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_checkCode\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taskContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"RegisteredToTaskRegistry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"inputParam\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"checkCode\",\"type\":\"string\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"taskInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"taskType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resourceType\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"inputParam\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"verifyParam\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"proof\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taskRegistryContract\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"checkCode\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200146538038062001465833981810160405281019062000037919062000721565b6040518061018001604052808b81526020018a81526020018981526020018881526020018781526020018673ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018481526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020013373ffffffffffffffffffffffffffffffffffffffff1681526020016040518060400160405280600381526020017f332e30000000000000000000000000000000000000000000000000000000000081525081525060008082015181600001556020820151816001015560408201518160020155606082015181600301908162000136919062000ad4565b5060808201518160040190816200014e919062000ad4565b5060a08201518160050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c0820151816006019081620001ad919062000ad4565b5060e082015181600701556101008201518160080160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061012082015181600901908162000218919062000ad4565b5061014082015181600a0160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061016082015181600b01908162000279919062000ad4565b509050507f7494b64152547ee0befdc148256d2de656f0aa9a5d6bf982acdf988ba8ff202e8a86898685604051620002b695949392919062000c2f565b60405180910390a1620002ce620002de60201b60201c565b5050505050505050505062000dc5565b60008060080160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16306000600a0160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040516024016200035592919062000c9a565b6040516020818303038152906040527fcaa29fc2000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051620003e1919062000d14565b6000604051808303816000865af19150503d806000811462000420576040519150601f19603f3d011682016040523d82523d6000602084013e62000425565b606091505b50509050806200046c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620004639062000da3565b60405180910390fd5b6000600a0160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff167f411a83d4fcdb8a204895aa1c817c68da89892ae8a277620988dcd6ea44650b7f60405160405180910390a350565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b620005178162000502565b81146200052357600080fd5b50565b60008151905062000537816200050c565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620005928262000547565b810181811067ffffffffffffffff82111715620005b457620005b362000558565b5b80604052505050565b6000620005c9620004ee565b9050620005d7828262000587565b919050565b600067ffffffffffffffff821115620005fa57620005f962000558565b5b620006058262000547565b9050602081019050919050565b60005b838110156200063257808201518184015260208101905062000615565b60008484015250505050565b6000620006556200064f84620005dc565b620005bd565b90508281526020810184848401111562000674576200067362000542565b5b6200068184828562000612565b509392505050565b600082601f830112620006a157620006a06200053d565b5b8151620006b38482602086016200063e565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620006e982620006bc565b9050919050565b620006fb81620006dc565b81146200070757600080fd5b50565b6000815190506200071b81620006f0565b92915050565b6000806000806000806000806000806101408b8d031215620007485762000747620004f8565b5b6000620007588d828e0162000526565b9a505060206200076b8d828e0162000526565b99505060406200077e8d828e0162000526565b98505060608b015167ffffffffffffffff811115620007a257620007a1620004fd565b5b620007b08d828e0162000689565b97505060808b015167ffffffffffffffff811115620007d457620007d3620004fd565b5b620007e28d828e0162000689565b96505060a0620007f58d828e016200070a565b95505060c08b015167ffffffffffffffff811115620008195762000818620004fd565b5b620008278d828e0162000689565b94505060e06200083a8d828e0162000526565b9350506101006200084e8d828e016200070a565b9250506101208b015167ffffffffffffffff811115620008735762000872620004fd565b5b620008818d828e0162000689565b9150509295989b9194979a5092959850565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620008e657607f821691505b602082108103620008fc57620008fb6200089e565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620009667fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000927565b62000972868362000927565b95508019841693508086168417925050509392505050565b6000819050919050565b6000620009b5620009af620009a98462000502565b6200098a565b62000502565b9050919050565b6000819050919050565b620009d18362000994565b620009e9620009e082620009bc565b84845462000934565b825550505050565b600090565b62000a00620009f1565b62000a0d818484620009c6565b505050565b5b8181101562000a355762000a29600082620009f6565b60018101905062000a13565b5050565b601f82111562000a845762000a4e8162000902565b62000a598462000917565b8101602085101562000a69578190505b62000a8162000a788562000917565b83018262000a12565b50505b505050565b600082821c905092915050565b600062000aa96000198460080262000a89565b1980831691505092915050565b600062000ac4838362000a96565b9150826002028217905092915050565b62000adf8262000893565b67ffffffffffffffff81111562000afb5762000afa62000558565b5b62000b078254620008cd565b62000b1482828562000a39565b600060209050601f83116001811462000b4c576000841562000b37578287015190505b62000b43858262000ab6565b86555062000bb3565b601f19841662000b5c8662000902565b60005b8281101562000b865784890151825560018201915060208501945060208101905062000b5f565b8683101562000ba6578489015162000ba2601f89168262000a96565b8355505b6001600288020188555050505b505050505050565b62000bc68162000502565b82525050565b62000bd781620006dc565b82525050565b600082825260208201905092915050565b600062000bfb8262000893565b62000c07818562000bdd565b935062000c1981856020860162000612565b62000c248162000547565b840191505092915050565b600060a08201905062000c46600083018862000bbb565b62000c55602083018762000bcc565b818103604083015262000c69818662000bee565b905062000c7a606083018562000bbb565b818103608083015262000c8e818462000bee565b90509695505050505050565b600060408201905062000cb1600083018562000bcc565b62000cc0602083018462000bcc565b9392505050565b600081519050919050565b600081905092915050565b600062000cea8262000cc7565b62000cf6818562000cd2565b935062000d0881856020860162000612565b80840191505092915050565b600062000d22828462000cdd565b915081905092915050565b7f4661696c656420746f207265676973746572207461736b20636f6e747261637460008201527f20746f205461736b526567697374727900000000000000000000000000000000602082015250565b600062000d8b60308362000bdd565b915062000d988262000d2d565b604082019050919050565b6000602082019050818103600083015262000dbe8162000d7c565b9050919050565b6106908062000dd56000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806354fd4d501461003b578063ac72255114610059575b600080fd5b610043610082565b60405161005091906104a1565b60405180910390f35b6100616100bb565b6040516100799c9b9a9998979695949392919061051d565b60405180910390f35b6040518060400160405280600381526020017f332e30000000000000000000000000000000000000000000000000000000000081525081565b60008060000154908060010154908060020154908060030180546100de90610629565b80601f016020809104026020016040519081016040528092919081815260200182805461010a90610629565b80156101575780601f1061012c57610100808354040283529160200191610157565b820191906000526020600020905b81548152906001019060200180831161013a57829003601f168201915b50505050509080600401805461016c90610629565b80601f016020809104026020016040519081016040528092919081815260200182805461019890610629565b80156101e55780601f106101ba576101008083540402835291602001916101e5565b820191906000526020600020905b8154815290600101906020018083116101c857829003601f168201915b5050505050908060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600601805461022090610629565b80601f016020809104026020016040519081016040528092919081815260200182805461024c90610629565b80156102995780601f1061026e57610100808354040283529160200191610299565b820191906000526020600020905b81548152906001019060200180831161027c57829003601f168201915b5050505050908060070154908060080160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060090180546102da90610629565b80601f016020809104026020016040519081016040528092919081815260200182805461030690610629565b80156103535780601f1061032857610100808354040283529160200191610353565b820191906000526020600020905b81548152906001019060200180831161033657829003601f168201915b50505050509080600a0160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600b01805461038e90610629565b80601f01602080910402602001604051908101604052809291908181526020018280546103ba90610629565b80156104075780601f106103dc57610100808354040283529160200191610407565b820191906000526020600020905b8154815290600101906020018083116103ea57829003601f168201915b505050505090508c565b600081519050919050565b600082825260208201905092915050565b60005b8381101561044b578082015181840152602081019050610430565b60008484015250505050565b6000601f19601f8301169050919050565b600061047382610411565b61047d818561041c565b935061048d81856020860161042d565b61049681610457565b840191505092915050565b600060208201905081810360008301526104bb8184610468565b905092915050565b6000819050919050565b6104d6816104c3565b82525050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610507826104dc565b9050919050565b610517816104fc565b82525050565b600061018082019050610533600083018f6104cd565b610540602083018e6104cd565b61054d604083018d6104cd565b818103606083015261055f818c610468565b90508181036080830152610573818b610468565b905061058260a083018a61050e565b81810360c08301526105948189610468565b90506105a360e08301886104cd565b6105b161010083018761050e565b8181036101208301526105c48186610468565b90506105d461014083018561050e565b8181036101608301526105e78184610468565b90509d9c50505050505050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061064157607f821691505b602082108103610654576106536105fa565b5b5091905056fea2646970667358221220b20a66f838453262d486eb6c5e21d5c4fc5375763d941839b37ac39e0efb5c7f64736f6c63430008140033",
}

// TaskABI is the input ABI used to generate the binding from.
// Deprecated: Use TaskMetaData.ABI instead.
var TaskABI = TaskMetaData.ABI

// TaskBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TaskMetaData.Bin instead.
var TaskBin = TaskMetaData.Bin

// DeployTask deploys a new Ethereum contract, binding an instance of Task to it.
func DeployTask(auth *bind.TransactOpts, backend bind.ContractBackend, _taskID *big.Int, _taskType *big.Int, _resourceType *big.Int, _inputParam string, _verifyParam string, _cpAccount common.Address, _proof string, _deadline *big.Int, _taskRegistryContract common.Address, _checkCode string) (common.Address, *types.Transaction, *Task, error) {
	parsed, err := TaskMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TaskBin), backend, _taskID, _taskType, _resourceType, _inputParam, _verifyParam, _cpAccount, _proof, _deadline, _taskRegistryContract, _checkCode)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Task{TaskCaller: TaskCaller{contract: contract}, TaskTransactor: TaskTransactor{contract: contract}, TaskFilterer: TaskFilterer{contract: contract}}, nil
}

// Task is an auto generated Go binding around an Ethereum contract.
type Task struct {
	TaskCaller     // Read-only binding to the contract
	TaskTransactor // Write-only binding to the contract
	TaskFilterer   // Log filterer for contract events
}

// TaskCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaskCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaskTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaskFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaskSession struct {
	Contract     *Task             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaskCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaskCallerSession struct {
	Contract *TaskCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TaskTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaskTransactorSession struct {
	Contract     *TaskTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaskRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaskRaw struct {
	Contract *Task // Generic contract binding to access the raw methods on
}

// TaskCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaskCallerRaw struct {
	Contract *TaskCaller // Generic read-only contract binding to access the raw methods on
}

// TaskTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaskTransactorRaw struct {
	Contract *TaskTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTask creates a new instance of Task, bound to a specific deployed contract.
func NewTask(address common.Address, backend bind.ContractBackend) (*Task, error) {
	contract, err := bindTask(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Task{TaskCaller: TaskCaller{contract: contract}, TaskTransactor: TaskTransactor{contract: contract}, TaskFilterer: TaskFilterer{contract: contract}}, nil
}

// NewTaskCaller creates a new read-only instance of Task, bound to a specific deployed contract.
func NewTaskCaller(address common.Address, caller bind.ContractCaller) (*TaskCaller, error) {
	contract, err := bindTask(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaskCaller{contract: contract}, nil
}

// NewTaskTransactor creates a new write-only instance of Task, bound to a specific deployed contract.
func NewTaskTransactor(address common.Address, transactor bind.ContractTransactor) (*TaskTransactor, error) {
	contract, err := bindTask(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaskTransactor{contract: contract}, nil
}

// NewTaskFilterer creates a new log filterer instance of Task, bound to a specific deployed contract.
func NewTaskFilterer(address common.Address, filterer bind.ContractFilterer) (*TaskFilterer, error) {
	contract, err := bindTask(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaskFilterer{contract: contract}, nil
}

// bindTask binds a generic wrapper to an already deployed contract.
func bindTask(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TaskMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Task *TaskRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Task.Contract.TaskCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Task *TaskRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Task.Contract.TaskTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Task *TaskRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Task.Contract.TaskTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Task *TaskCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Task.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Task *TaskTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Task.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Task *TaskTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Task.Contract.contract.Transact(opts, method, params...)
}

// TaskInfo is a free data retrieval call binding the contract method 0xac722551.
//
// Solidity: function taskInfo() view returns(uint256 taskID, uint256 taskType, uint256 resourceType, string inputParam, string verifyParam, address cpAccount, string proof, uint256 deadline, address taskRegistryContract, string checkCode, address owner, string version)
func (_Task *TaskCaller) TaskInfo(opts *bind.CallOpts) (struct {
	TaskID               *big.Int
	TaskType             *big.Int
	ResourceType         *big.Int
	InputParam           string
	VerifyParam          string
	CpAccount            common.Address
	Proof                string
	Deadline             *big.Int
	TaskRegistryContract common.Address
	CheckCode            string
	Owner                common.Address
	Version              string
}, error) {
	var out []interface{}
	err := _Task.contract.Call(opts, &out, "taskInfo")

	outstruct := new(struct {
		TaskID               *big.Int
		TaskType             *big.Int
		ResourceType         *big.Int
		InputParam           string
		VerifyParam          string
		CpAccount            common.Address
		Proof                string
		Deadline             *big.Int
		TaskRegistryContract common.Address
		CheckCode            string
		Owner                common.Address
		Version              string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TaskID = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TaskType = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ResourceType = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.InputParam = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.VerifyParam = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.CpAccount = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.Proof = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.Deadline = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.TaskRegistryContract = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)
	outstruct.CheckCode = *abi.ConvertType(out[9], new(string)).(*string)
	outstruct.Owner = *abi.ConvertType(out[10], new(common.Address)).(*common.Address)
	outstruct.Version = *abi.ConvertType(out[11], new(string)).(*string)

	return *outstruct, err

}

// TaskInfo is a free data retrieval call binding the contract method 0xac722551.
//
// Solidity: function taskInfo() view returns(uint256 taskID, uint256 taskType, uint256 resourceType, string inputParam, string verifyParam, address cpAccount, string proof, uint256 deadline, address taskRegistryContract, string checkCode, address owner, string version)
func (_Task *TaskSession) TaskInfo() (struct {
	TaskID               *big.Int
	TaskType             *big.Int
	ResourceType         *big.Int
	InputParam           string
	VerifyParam          string
	CpAccount            common.Address
	Proof                string
	Deadline             *big.Int
	TaskRegistryContract common.Address
	CheckCode            string
	Owner                common.Address
	Version              string
}, error) {
	return _Task.Contract.TaskInfo(&_Task.CallOpts)
}

// TaskInfo is a free data retrieval call binding the contract method 0xac722551.
//
// Solidity: function taskInfo() view returns(uint256 taskID, uint256 taskType, uint256 resourceType, string inputParam, string verifyParam, address cpAccount, string proof, uint256 deadline, address taskRegistryContract, string checkCode, address owner, string version)
func (_Task *TaskCallerSession) TaskInfo() (struct {
	TaskID               *big.Int
	TaskType             *big.Int
	ResourceType         *big.Int
	InputParam           string
	VerifyParam          string
	CpAccount            common.Address
	Proof                string
	Deadline             *big.Int
	TaskRegistryContract common.Address
	CheckCode            string
	Owner                common.Address
	Version              string
}, error) {
	return _Task.Contract.TaskInfo(&_Task.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Task *TaskCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Task.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Task *TaskSession) Version() (string, error) {
	return _Task.Contract.Version(&_Task.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Task *TaskCallerSession) Version() (string, error) {
	return _Task.Contract.Version(&_Task.CallOpts)
}

// TaskRegisteredToTaskRegistryIterator is returned from FilterRegisteredToTaskRegistry and is used to iterate over the raw logs and unpacked data for RegisteredToTaskRegistry events raised by the Task contract.
type TaskRegisteredToTaskRegistryIterator struct {
	Event *TaskRegisteredToTaskRegistry // Event containing the contract specifics and raw log

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
func (it *TaskRegisteredToTaskRegistryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskRegisteredToTaskRegistry)
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
		it.Event = new(TaskRegisteredToTaskRegistry)
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
func (it *TaskRegisteredToTaskRegistryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskRegisteredToTaskRegistryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskRegisteredToTaskRegistry represents a RegisteredToTaskRegistry event raised by the Task contract.
type TaskRegisteredToTaskRegistry struct {
	TaskContract common.Address
	Owner        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRegisteredToTaskRegistry is a free log retrieval operation binding the contract event 0x411a83d4fcdb8a204895aa1c817c68da89892ae8a277620988dcd6ea44650b7f.
//
// Solidity: event RegisteredToTaskRegistry(address indexed taskContract, address indexed owner)
func (_Task *TaskFilterer) FilterRegisteredToTaskRegistry(opts *bind.FilterOpts, taskContract []common.Address, owner []common.Address) (*TaskRegisteredToTaskRegistryIterator, error) {

	var taskContractRule []interface{}
	for _, taskContractItem := range taskContract {
		taskContractRule = append(taskContractRule, taskContractItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Task.contract.FilterLogs(opts, "RegisteredToTaskRegistry", taskContractRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &TaskRegisteredToTaskRegistryIterator{contract: _Task.contract, event: "RegisteredToTaskRegistry", logs: logs, sub: sub}, nil
}

// WatchRegisteredToTaskRegistry is a free log subscription operation binding the contract event 0x411a83d4fcdb8a204895aa1c817c68da89892ae8a277620988dcd6ea44650b7f.
//
// Solidity: event RegisteredToTaskRegistry(address indexed taskContract, address indexed owner)
func (_Task *TaskFilterer) WatchRegisteredToTaskRegistry(opts *bind.WatchOpts, sink chan<- *TaskRegisteredToTaskRegistry, taskContract []common.Address, owner []common.Address) (event.Subscription, error) {

	var taskContractRule []interface{}
	for _, taskContractItem := range taskContract {
		taskContractRule = append(taskContractRule, taskContractItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Task.contract.WatchLogs(opts, "RegisteredToTaskRegistry", taskContractRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskRegisteredToTaskRegistry)
				if err := _Task.contract.UnpackLog(event, "RegisteredToTaskRegistry", log); err != nil {
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
func (_Task *TaskFilterer) ParseRegisteredToTaskRegistry(log types.Log) (*TaskRegisteredToTaskRegistry, error) {
	event := new(TaskRegisteredToTaskRegistry)
	if err := _Task.contract.UnpackLog(event, "RegisteredToTaskRegistry", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskTaskCreatedIterator is returned from FilterTaskCreated and is used to iterate over the raw logs and unpacked data for TaskCreated events raised by the Task contract.
type TaskTaskCreatedIterator struct {
	Event *TaskTaskCreated // Event containing the contract specifics and raw log

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
func (it *TaskTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskTaskCreated)
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
		it.Event = new(TaskTaskCreated)
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
func (it *TaskTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskTaskCreated represents a TaskCreated event raised by the Task contract.
type TaskTaskCreated struct {
	TaskID     *big.Int
	CpAccount  common.Address
	InputParam string
	Deadline   *big.Int
	CheckCode  string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskCreated is a free log retrieval operation binding the contract event 0x7494b64152547ee0befdc148256d2de656f0aa9a5d6bf982acdf988ba8ff202e.
//
// Solidity: event TaskCreated(uint256 taskID, address cpAccount, string inputParam, uint256 deadline, string checkCode)
func (_Task *TaskFilterer) FilterTaskCreated(opts *bind.FilterOpts) (*TaskTaskCreatedIterator, error) {

	logs, sub, err := _Task.contract.FilterLogs(opts, "TaskCreated")
	if err != nil {
		return nil, err
	}
	return &TaskTaskCreatedIterator{contract: _Task.contract, event: "TaskCreated", logs: logs, sub: sub}, nil
}

// WatchTaskCreated is a free log subscription operation binding the contract event 0x7494b64152547ee0befdc148256d2de656f0aa9a5d6bf982acdf988ba8ff202e.
//
// Solidity: event TaskCreated(uint256 taskID, address cpAccount, string inputParam, uint256 deadline, string checkCode)
func (_Task *TaskFilterer) WatchTaskCreated(opts *bind.WatchOpts, sink chan<- *TaskTaskCreated) (event.Subscription, error) {

	logs, sub, err := _Task.contract.WatchLogs(opts, "TaskCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskTaskCreated)
				if err := _Task.contract.UnpackLog(event, "TaskCreated", log); err != nil {
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

// ParseTaskCreated is a log parse operation binding the contract event 0x7494b64152547ee0befdc148256d2de656f0aa9a5d6bf982acdf988ba8ff202e.
//
// Solidity: event TaskCreated(uint256 taskID, address cpAccount, string inputParam, uint256 deadline, string checkCode)
func (_Task *TaskFilterer) ParseTaskCreated(log types.Log) (*TaskTaskCreated, error) {
	event := new(TaskTaskCreated)
	if err := _Task.contract.UnpackLog(event, "TaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
