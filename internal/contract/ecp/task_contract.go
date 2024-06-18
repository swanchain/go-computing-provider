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

// ECPTaskTaskInfo is an auto generated low-level Go binding around an user-defined struct.
type ECPTaskTaskInfo struct {
	TaskID               *big.Int
	TaskType             *big.Int
	ResourceType         string
	InputParam           string
	VerifyParam          string
	CpAccount            common.Address
	Proof                string
	Deadline             *big.Int
	TaskRegistryContract common.Address
	CheckCode            string
	Owner                common.Address
}

// TaskMetaData contains all meta data concerning the Task contract.
var TaskMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_taskID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_taskType\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_resourceType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_inputParam\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_verifyParam\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_cpAccount\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_proof\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_taskRegistryContract\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_checkCode\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taskContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"RegisteredToTaskRegistry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"inputParam\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"checkCode\",\"type\":\"string\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getTaskInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"taskType\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"resourceType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"inputParam\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"verifyParam\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"proof\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taskRegistryContract\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"checkCode\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structECPTask.TaskInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"taskType\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"resourceType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"inputParam\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"verifyParam\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"proof\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taskRegistryContract\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"checkCode\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620020c5380380620020c5833981810160405281019062000037919062000965565b60008a116200007d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000749062000b7d565b60405180910390fd5b60008911620000c3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000ba9062000c15565b60405180910390fd5b60008851116200010a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001019062000c87565b60405180910390fd5b600087511162000151576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001489062000cf9565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1603620001c3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001ba9062000d91565b60405180910390fd5b42831162000208576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001ff9062000e03565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036200027a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002719062000e9b565b60405180910390fd5b6000815111620002c1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002b89062000f0d565b60405180910390fd5b6040518061016001604052808b81526020018a81526020018981526020018881526020018781526020018673ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018481526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020013373ffffffffffffffffffffffffffffffffffffffff1681525060008082015181600001556020820151816001015560408201518160020190816200037b919062001170565b50606082015181600301908162000393919062001170565b506080820151816004019081620003ab919062001170565b5060a08201518160050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c08201518160060190816200040a919062001170565b5060e082015181600701556101008201518160080160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061012082015181600901908162000475919062001170565b5061014082015181600a0160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050507f7494b64152547ee0befdc148256d2de656f0aa9a5d6bf982acdf988ba8ff202e8a86898685604051620004fa959493929190620012ba565b60405180910390a1620005126200052260201b60201c565b5050505050505050505062001450565b60008060080160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16306000600a0160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040516024016200059992919062001325565b6040516020818303038152906040527fcaa29fc2000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516200062591906200139f565b6000604051808303816000865af19150503d806000811462000664576040519150601f19603f3d011682016040523d82523d6000602084013e62000669565b606091505b5050905080620006b0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620006a7906200142e565b60405180910390fd5b6000600a0160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff167f411a83d4fcdb8a204895aa1c817c68da89892ae8a277620988dcd6ea44650b7f60405160405180910390a350565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b6200075b8162000746565b81146200076757600080fd5b50565b6000815190506200077b8162000750565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620007d6826200078b565b810181811067ffffffffffffffff82111715620007f857620007f76200079c565b5b80604052505050565b60006200080d62000732565b90506200081b8282620007cb565b919050565b600067ffffffffffffffff8211156200083e576200083d6200079c565b5b62000849826200078b565b9050602081019050919050565b60005b838110156200087657808201518184015260208101905062000859565b60008484015250505050565b600062000899620008938462000820565b62000801565b905082815260208101848484011115620008b857620008b762000786565b5b620008c584828562000856565b509392505050565b600082601f830112620008e557620008e462000781565b5b8151620008f784826020860162000882565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200092d8262000900565b9050919050565b6200093f8162000920565b81146200094b57600080fd5b50565b6000815190506200095f8162000934565b92915050565b6000806000806000806000806000806101408b8d0312156200098c576200098b6200073c565b5b60006200099c8d828e016200076a565b9a50506020620009af8d828e016200076a565b99505060408b015167ffffffffffffffff811115620009d357620009d262000741565b5b620009e18d828e01620008cd565b98505060608b015167ffffffffffffffff81111562000a055762000a0462000741565b5b62000a138d828e01620008cd565b97505060808b015167ffffffffffffffff81111562000a375762000a3662000741565b5b62000a458d828e01620008cd565b96505060a062000a588d828e016200094e565b95505060c08b015167ffffffffffffffff81111562000a7c5762000a7b62000741565b5b62000a8a8d828e01620008cd565b94505060e062000a9d8d828e016200076a565b93505061010062000ab18d828e016200094e565b9250506101208b015167ffffffffffffffff81111562000ad65762000ad562000741565b5b62000ae48d828e01620008cd565b9150509295989b9194979a5092959850565b600082825260208201905092915050565b7f5461736b204944206d7573742062652067726561746572207468616e207a657260008201527f6f00000000000000000000000000000000000000000000000000000000000000602082015250565b600062000b6560218362000af6565b915062000b728262000b07565b604082019050919050565b6000602082019050818103600083015262000b988162000b56565b9050919050565b7f5461736b2074797065206d7573742062652067726561746572207468616e207a60008201527f65726f0000000000000000000000000000000000000000000000000000000000602082015250565b600062000bfd60238362000af6565b915062000c0a8262000b9f565b604082019050919050565b6000602082019050818103600083015262000c308162000bee565b9050919050565b7f5265736f757263652074797065206d7573742062652070726f76696465640000600082015250565b600062000c6f601e8362000af6565b915062000c7c8262000c37565b602082019050919050565b6000602082019050818103600083015262000ca28162000c60565b9050919050565b7f496e70757420706172616d206d7573742062652070726f766964656400000000600082015250565b600062000ce1601c8362000af6565b915062000cee8262000ca9565b602082019050919050565b6000602082019050818103600083015262000d148162000cd2565b9050919050565b7f4350206163636f756e742061646472657373206d7573742062652070726f766960008201527f6465640000000000000000000000000000000000000000000000000000000000602082015250565b600062000d7960238362000af6565b915062000d868262000d1b565b604082019050919050565b6000602082019050818103600083015262000dac8162000d6a565b9050919050565b7f446561646c696e65206d75737420626520696e20746865206675747572650000600082015250565b600062000deb601e8362000af6565b915062000df88262000db3565b602082019050919050565b6000602082019050818103600083015262000e1e8162000ddc565b9050919050565b7f5461736b20726567697374727920636f6e74726163742061646472657373206d60008201527f7573742062652070726f76696465640000000000000000000000000000000000602082015250565b600062000e83602f8362000af6565b915062000e908262000e25565b604082019050919050565b6000602082019050818103600083015262000eb68162000e74565b9050919050565b7f436865636b20636f6465206d7573742062652070726f76696465640000000000600082015250565b600062000ef5601b8362000af6565b915062000f028262000ebd565b602082019050919050565b6000602082019050818103600083015262000f288162000ee6565b9050919050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168062000f8257607f821691505b60208210810362000f985762000f9762000f3a565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620010027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000fc3565b6200100e868362000fc3565b95508019841693508086168417925050509392505050565b6000819050919050565b6000620010516200104b620010458462000746565b62001026565b62000746565b9050919050565b6000819050919050565b6200106d8362001030565b620010856200107c8262001058565b84845462000fd0565b825550505050565b600090565b6200109c6200108d565b620010a981848462001062565b505050565b5b81811015620010d157620010c560008262001092565b600181019050620010af565b5050565b601f8211156200112057620010ea8162000f9e565b620010f58462000fb3565b8101602085101562001105578190505b6200111d620011148562000fb3565b830182620010ae565b50505b505050565b600082821c905092915050565b6000620011456000198460080262001125565b1980831691505092915050565b600062001160838362001132565b9150826002028217905092915050565b6200117b8262000f2f565b67ffffffffffffffff8111156200119757620011966200079c565b5b620011a3825462000f69565b620011b0828285620010d5565b600060209050601f831160018114620011e85760008415620011d3578287015190505b620011df858262001152565b8655506200124f565b601f198416620011f88662000f9e565b60005b828110156200122257848901518255600182019150602085019450602081019050620011fb565b868310156200124257848901516200123e601f89168262001132565b8355505b6001600288020188555050505b505050505050565b620012628162000746565b82525050565b620012738162000920565b82525050565b6000620012868262000f2f565b62001292818562000af6565b9350620012a481856020860162000856565b620012af816200078b565b840191505092915050565b600060a082019050620012d1600083018862001257565b620012e0602083018762001268565b8181036040830152620012f4818662001279565b905062001305606083018562001257565b818103608083015262001319818462001279565b90509695505050505050565b60006040820190506200133c600083018562001268565b6200134b602083018462001268565b9392505050565b600081519050919050565b600081905092915050565b6000620013758262001352565b6200138181856200135d565b93506200139381856020860162000856565b80840191505092915050565b6000620013ad828462001368565b915081905092915050565b7f4661696c656420746f207265676973746572207461736b20636f6e747261637460008201527f20746f205461736b526567697374727900000000000000000000000000000000602082015250565b60006200141660308362000af6565b91506200142382620013b8565b604082019050919050565b60006020820190508181036000830152620014498162001407565b9050919050565b610c6580620014606000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063ac7225511461003b578063bac64a2e14610063575b600080fd5b610043610081565b60405161005a9b9a9998979695949392919061096c565b60405180910390f35b61006b6103d1565b6040516100789190610bad565b60405180910390f35b600080600001549080600101549080600201805461009e90610bfe565b80601f01602080910402602001604051908101604052809291908181526020018280546100ca90610bfe565b80156101175780601f106100ec57610100808354040283529160200191610117565b820191906000526020600020905b8154815290600101906020018083116100fa57829003601f168201915b50505050509080600301805461012c90610bfe565b80601f016020809104026020016040519081016040528092919081815260200182805461015890610bfe565b80156101a55780601f1061017a576101008083540402835291602001916101a5565b820191906000526020600020905b81548152906001019060200180831161018857829003601f168201915b5050505050908060040180546101ba90610bfe565b80601f01602080910402602001604051908101604052809291908181526020018280546101e690610bfe565b80156102335780601f1061020857610100808354040283529160200191610233565b820191906000526020600020905b81548152906001019060200180831161021657829003601f168201915b5050505050908060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600601805461026e90610bfe565b80601f016020809104026020016040519081016040528092919081815260200182805461029a90610bfe565b80156102e75780601f106102bc576101008083540402835291602001916102e7565b820191906000526020600020905b8154815290600101906020018083116102ca57829003601f168201915b5050505050908060070154908060080160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600901805461032890610bfe565b80601f016020809104026020016040519081016040528092919081815260200182805461035490610bfe565b80156103a15780601f10610376576101008083540402835291602001916103a1565b820191906000526020600020905b81548152906001019060200180831161038457829003601f168201915b50505050509080600a0160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508b565b6103d96107e6565b600060405180610160016040529081600082015481526020016001820154815260200160028201805461040b90610bfe565b80601f016020809104026020016040519081016040528092919081815260200182805461043790610bfe565b80156104845780601f1061045957610100808354040283529160200191610484565b820191906000526020600020905b81548152906001019060200180831161046757829003601f168201915b5050505050815260200160038201805461049d90610bfe565b80601f01602080910402602001604051908101604052809291908181526020018280546104c990610bfe565b80156105165780601f106104eb57610100808354040283529160200191610516565b820191906000526020600020905b8154815290600101906020018083116104f957829003601f168201915b5050505050815260200160048201805461052f90610bfe565b80601f016020809104026020016040519081016040528092919081815260200182805461055b90610bfe565b80156105a85780601f1061057d576101008083540402835291602001916105a8565b820191906000526020600020905b81548152906001019060200180831161058b57829003601f168201915b505050505081526020016005820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160068201805461061790610bfe565b80601f016020809104026020016040519081016040528092919081815260200182805461064390610bfe565b80156106905780601f1061066557610100808354040283529160200191610690565b820191906000526020600020905b81548152906001019060200180831161067357829003601f168201915b50505050508152602001600782015481526020016008820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160098201805461070990610bfe565b80601f016020809104026020016040519081016040528092919081815260200182805461073590610bfe565b80156107825780601f1061075757610100808354040283529160200191610782565b820191906000526020600020905b81548152906001019060200180831161076557829003601f168201915b50505050508152602001600a820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681525050905090565b6040518061016001604052806000815260200160008152602001606081526020016060815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016060815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff1681525090565b6000819050919050565b61089581610882565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b838110156108d55780820151818401526020810190506108ba565b60008484015250505050565b6000601f19601f8301169050919050565b60006108fd8261089b565b61090781856108a6565b93506109178185602086016108b7565b610920816108e1565b840191505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006109568261092b565b9050919050565b6109668161094b565b82525050565b600061016082019050610982600083018e61088c565b61098f602083018d61088c565b81810360408301526109a1818c6108f2565b905081810360608301526109b5818b6108f2565b905081810360808301526109c9818a6108f2565b90506109d860a083018961095d565b81810360c08301526109ea81886108f2565b90506109f960e083018761088c565b610a0761010083018661095d565b818103610120830152610a1a81856108f2565b9050610a2a61014083018461095d565b9c9b505050505050505050505050565b610a4381610882565b82525050565b600082825260208201905092915050565b6000610a658261089b565b610a6f8185610a49565b9350610a7f8185602086016108b7565b610a88816108e1565b840191505092915050565b610a9c8161094b565b82525050565b600061016083016000830151610abb6000860182610a3a565b506020830151610ace6020860182610a3a565b5060408301518482036040860152610ae68282610a5a565b91505060608301518482036060860152610b008282610a5a565b91505060808301518482036080860152610b1a8282610a5a565b91505060a0830151610b2f60a0860182610a93565b5060c083015184820360c0860152610b478282610a5a565b91505060e0830151610b5c60e0860182610a3a565b50610100830151610b71610100860182610a93565b50610120830151848203610120860152610b8b8282610a5a565b915050610140830151610ba2610140860182610a93565b508091505092915050565b60006020820190508181036000830152610bc78184610aa2565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610c1657607f821691505b602082108103610c2957610c28610bcf565b5b5091905056fea2646970667358221220864b8f72361118d22c765e83cda9b21f8a8fa82bbd4470ae4b5f1d6886203cc964736f6c63430008140033",
}

// TaskABI is the input ABI used to generate the binding from.
// Deprecated: Use TaskMetaData.ABI instead.
var TaskABI = TaskMetaData.ABI

// TaskBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TaskMetaData.Bin instead.
var TaskBin = TaskMetaData.Bin

// DeployTask deploys a new Ethereum contract, binding an instance of Task to it.
func DeployTask(auth *bind.TransactOpts, backend bind.ContractBackend, _taskID *big.Int, _taskType *big.Int, _resourceType string, _inputParam string, _verifyParam string, _cpAccount common.Address, _proof string, _deadline *big.Int, _taskRegistryContract common.Address, _checkCode string) (common.Address, *types.Transaction, *Task, error) {
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

// GetTaskInfo is a free data retrieval call binding the contract method 0xbac64a2e.
//
// Solidity: function getTaskInfo() view returns((uint256,uint256,string,string,string,address,string,uint256,address,string,address))
func (_Task *TaskCaller) GetTaskInfo(opts *bind.CallOpts) (ECPTaskTaskInfo, error) {
	var out []interface{}
	err := _Task.contract.Call(opts, &out, "getTaskInfo")

	if err != nil {
		return *new(ECPTaskTaskInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ECPTaskTaskInfo)).(*ECPTaskTaskInfo)

	return out0, err

}

// GetTaskInfo is a free data retrieval call binding the contract method 0xbac64a2e.
//
// Solidity: function getTaskInfo() view returns((uint256,uint256,string,string,string,address,string,uint256,address,string,address))
func (_Task *TaskSession) GetTaskInfo() (ECPTaskTaskInfo, error) {
	return _Task.Contract.GetTaskInfo(&_Task.CallOpts)
}

// GetTaskInfo is a free data retrieval call binding the contract method 0xbac64a2e.
//
// Solidity: function getTaskInfo() view returns((uint256,uint256,string,string,string,address,string,uint256,address,string,address))
func (_Task *TaskCallerSession) GetTaskInfo() (ECPTaskTaskInfo, error) {
	return _Task.Contract.GetTaskInfo(&_Task.CallOpts)
}

// TaskInfo is a free data retrieval call binding the contract method 0xac722551.
//
// Solidity: function taskInfo() view returns(uint256 taskID, uint256 taskType, string resourceType, string inputParam, string verifyParam, address cpAccount, string proof, uint256 deadline, address taskRegistryContract, string checkCode, address owner)
func (_Task *TaskCaller) TaskInfo(opts *bind.CallOpts) (struct {
	TaskID               *big.Int
	TaskType             *big.Int
	ResourceType         string
	InputParam           string
	VerifyParam          string
	CpAccount            common.Address
	Proof                string
	Deadline             *big.Int
	TaskRegistryContract common.Address
	CheckCode            string
	Owner                common.Address
}, error) {
	var out []interface{}
	err := _Task.contract.Call(opts, &out, "taskInfo")

	outstruct := new(struct {
		TaskID               *big.Int
		TaskType             *big.Int
		ResourceType         string
		InputParam           string
		VerifyParam          string
		CpAccount            common.Address
		Proof                string
		Deadline             *big.Int
		TaskRegistryContract common.Address
		CheckCode            string
		Owner                common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TaskID = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TaskType = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ResourceType = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.InputParam = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.VerifyParam = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.CpAccount = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.Proof = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.Deadline = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.TaskRegistryContract = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)
	outstruct.CheckCode = *abi.ConvertType(out[9], new(string)).(*string)
	outstruct.Owner = *abi.ConvertType(out[10], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// TaskInfo is a free data retrieval call binding the contract method 0xac722551.
//
// Solidity: function taskInfo() view returns(uint256 taskID, uint256 taskType, string resourceType, string inputParam, string verifyParam, address cpAccount, string proof, uint256 deadline, address taskRegistryContract, string checkCode, address owner)
func (_Task *TaskSession) TaskInfo() (struct {
	TaskID               *big.Int
	TaskType             *big.Int
	ResourceType         string
	InputParam           string
	VerifyParam          string
	CpAccount            common.Address
	Proof                string
	Deadline             *big.Int
	TaskRegistryContract common.Address
	CheckCode            string
	Owner                common.Address
}, error) {
	return _Task.Contract.TaskInfo(&_Task.CallOpts)
}

// TaskInfo is a free data retrieval call binding the contract method 0xac722551.
//
// Solidity: function taskInfo() view returns(uint256 taskID, uint256 taskType, string resourceType, string inputParam, string verifyParam, address cpAccount, string proof, uint256 deadline, address taskRegistryContract, string checkCode, address owner)
func (_Task *TaskCallerSession) TaskInfo() (struct {
	TaskID               *big.Int
	TaskType             *big.Int
	ResourceType         string
	InputParam           string
	VerifyParam          string
	CpAccount            common.Address
	Proof                string
	Deadline             *big.Int
	TaskRegistryContract common.Address
	CheckCode            string
	Owner                common.Address
}, error) {
	return _Task.Contract.TaskInfo(&_Task.CallOpts)
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
