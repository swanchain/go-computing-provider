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
}

// TaskMetaData contains all meta data concerning the Task contract.
var TaskMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_taskID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_taskType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_resourceType\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_inputParam\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_verifyParam\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_cpAccount\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_proof\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_taskRegistryContract\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_checkCode\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taskContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"RegisteredToTaskRegistry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"inputParam\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"checkCode\",\"type\":\"string\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getTaskInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"taskType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resourceType\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"inputParam\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"verifyParam\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"proof\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taskRegistryContract\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"checkCode\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"internalType\":\"structECPTask.TaskInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"taskType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resourceType\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"inputParam\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"verifyParam\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"proof\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taskRegistryContract\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"checkCode\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620021ab380380620021ab8339818101604052810190620000379190620009aa565b60008a116200007d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000749062000ba3565b60405180910390fd5b60008911620000c3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000ba9062000c3b565b60405180910390fd5b6000881162000109576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001009062000cad565b60405180910390fd5b600087511162000150576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001479062000d1f565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1603620001c2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001b99062000db7565b60405180910390fd5b43831162000207576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001fe9062000e29565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160362000279576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002709062000ec1565b60405180910390fd5b6000815111620002c0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002b79062000f33565b60405180910390fd5b6040518061018001604052808b81526020018a81526020018981526020018881526020018781526020018673ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018481526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020013373ffffffffffffffffffffffffffffffffffffffff1681526020016040518060400160405280600381526020017f332e300000000000000000000000000000000000000000000000000000000000815250815250600080820151816000015560208201518160010155604082015181600201556060820151816003019081620003bf919062001196565b506080820151816004019081620003d7919062001196565b5060a08201518160050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c082015181600601908162000436919062001196565b5060e082015181600701556101008201518160080160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610120820151816009019081620004a1919062001196565b5061014082015181600a0160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061016082015181600b01908162000502919062001196565b509050507f7494b64152547ee0befdc148256d2de656f0aa9a5d6bf982acdf988ba8ff202e8a868986856040516200053f959493929190620012e0565b60405180910390a1620005576200056760201b60201c565b5050505050505050505062001476565b60008060080160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16306000600a0160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051602401620005de9291906200134b565b6040516020818303038152906040527fcaa29fc2000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516200066a9190620013c5565b6000604051808303816000865af19150503d8060008114620006a9576040519150601f19603f3d011682016040523d82523d6000602084013e620006ae565b606091505b5050905080620006f5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620006ec9062001454565b60405180910390fd5b6000600a0160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff167f411a83d4fcdb8a204895aa1c817c68da89892ae8a277620988dcd6ea44650b7f60405160405180910390a350565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b620007a0816200078b565b8114620007ac57600080fd5b50565b600081519050620007c08162000795565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6200081b82620007d0565b810181811067ffffffffffffffff821117156200083d576200083c620007e1565b5b80604052505050565b60006200085262000777565b905062000860828262000810565b919050565b600067ffffffffffffffff821115620008835762000882620007e1565b5b6200088e82620007d0565b9050602081019050919050565b60005b83811015620008bb5780820151818401526020810190506200089e565b60008484015250505050565b6000620008de620008d88462000865565b62000846565b905082815260208101848484011115620008fd57620008fc620007cb565b5b6200090a8482856200089b565b509392505050565b600082601f8301126200092a5762000929620007c6565b5b81516200093c848260208601620008c7565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620009728262000945565b9050919050565b620009848162000965565b81146200099057600080fd5b50565b600081519050620009a48162000979565b92915050565b6000806000806000806000806000806101408b8d031215620009d157620009d062000781565b5b6000620009e18d828e01620007af565b9a50506020620009f48d828e01620007af565b995050604062000a078d828e01620007af565b98505060608b015167ffffffffffffffff81111562000a2b5762000a2a62000786565b5b62000a398d828e0162000912565b97505060808b015167ffffffffffffffff81111562000a5d5762000a5c62000786565b5b62000a6b8d828e0162000912565b96505060a062000a7e8d828e0162000993565b95505060c08b015167ffffffffffffffff81111562000aa25762000aa162000786565b5b62000ab08d828e0162000912565b94505060e062000ac38d828e01620007af565b93505061010062000ad78d828e0162000993565b9250506101208b015167ffffffffffffffff81111562000afc5762000afb62000786565b5b62000b0a8d828e0162000912565b9150509295989b9194979a5092959850565b600082825260208201905092915050565b7f5461736b204944206d7573742062652067726561746572207468616e207a657260008201527f6f00000000000000000000000000000000000000000000000000000000000000602082015250565b600062000b8b60218362000b1c565b915062000b988262000b2d565b604082019050919050565b6000602082019050818103600083015262000bbe8162000b7c565b9050919050565b7f5461736b2074797065206d7573742062652067726561746572207468616e207a60008201527f65726f0000000000000000000000000000000000000000000000000000000000602082015250565b600062000c2360238362000b1c565b915062000c308262000bc5565b604082019050919050565b6000602082019050818103600083015262000c568162000c14565b9050919050565b7f5265736f757263652074797065206d7573742062652070726f76696465640000600082015250565b600062000c95601e8362000b1c565b915062000ca28262000c5d565b602082019050919050565b6000602082019050818103600083015262000cc88162000c86565b9050919050565b7f496e70757420706172616d206d7573742062652070726f766964656400000000600082015250565b600062000d07601c8362000b1c565b915062000d148262000ccf565b602082019050919050565b6000602082019050818103600083015262000d3a8162000cf8565b9050919050565b7f4350206163636f756e742061646472657373206d7573742062652070726f766960008201527f6465640000000000000000000000000000000000000000000000000000000000602082015250565b600062000d9f60238362000b1c565b915062000dac8262000d41565b604082019050919050565b6000602082019050818103600083015262000dd28162000d90565b9050919050565b7f446561646c696e65206d75737420626520696e20746865206675747572650000600082015250565b600062000e11601e8362000b1c565b915062000e1e8262000dd9565b602082019050919050565b6000602082019050818103600083015262000e448162000e02565b9050919050565b7f5461736b20726567697374727920636f6e74726163742061646472657373206d60008201527f7573742062652070726f76696465640000000000000000000000000000000000602082015250565b600062000ea9602f8362000b1c565b915062000eb68262000e4b565b604082019050919050565b6000602082019050818103600083015262000edc8162000e9a565b9050919050565b7f436865636b20636f6465206d7573742062652070726f76696465640000000000600082015250565b600062000f1b601b8362000b1c565b915062000f288262000ee3565b602082019050919050565b6000602082019050818103600083015262000f4e8162000f0c565b9050919050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168062000fa857607f821691505b60208210810362000fbe5762000fbd62000f60565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620010287fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000fe9565b62001034868362000fe9565b95508019841693508086168417925050509392505050565b6000819050919050565b600062001077620010716200106b846200078b565b6200104c565b6200078b565b9050919050565b6000819050919050565b620010938362001056565b620010ab620010a2826200107e565b84845462000ff6565b825550505050565b600090565b620010c2620010b3565b620010cf81848462001088565b505050565b5b81811015620010f757620010eb600082620010b8565b600181019050620010d5565b5050565b601f8211156200114657620011108162000fc4565b6200111b8462000fd9565b810160208510156200112b578190505b620011436200113a8562000fd9565b830182620010d4565b50505b505050565b600082821c905092915050565b60006200116b600019846008026200114b565b1980831691505092915050565b600062001186838362001158565b9150826002028217905092915050565b620011a18262000f55565b67ffffffffffffffff811115620011bd57620011bc620007e1565b5b620011c9825462000f8f565b620011d6828285620010fb565b600060209050601f8311600181146200120e5760008415620011f9578287015190505b62001205858262001178565b86555062001275565b601f1984166200121e8662000fc4565b60005b82811015620012485784890151825560018201915060208501945060208101905062001221565b8683101562001268578489015162001264601f89168262001158565b8355505b6001600288020188555050505b505050505050565b62001288816200078b565b82525050565b620012998162000965565b82525050565b6000620012ac8262000f55565b620012b8818562000b1c565b9350620012ca8185602086016200089b565b620012d581620007d0565b840191505092915050565b600060a082019050620012f760008301886200127d565b6200130660208301876200128e565b81810360408301526200131a81866200129f565b90506200132b60608301856200127d565b81810360808301526200133f81846200129f565b90509695505050505050565b60006040820190506200136260008301856200128e565b6200137160208301846200128e565b9392505050565b600081519050919050565b600081905092915050565b60006200139b8262001378565b620013a7818562001383565b9350620013b98185602086016200089b565b80840191505092915050565b6000620013d382846200138e565b915081905092915050565b7f4661696c656420746f207265676973746572207461736b20636f6e747261637460008201527f20746f205461736b526567697374727900000000000000000000000000000000602082015250565b60006200143c60308362000b1c565b91506200144982620013de565b604082019050919050565b600060208201905081810360008301526200146f816200142d565b9050919050565b610d2580620014866000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806354fd4d5014610046578063ac72255114610064578063bac64a2e1461008d575b600080fd5b61004e6100ab565b60405161005b919061098c565b60405180910390f35b61006c6100e4565b6040516100849c9b9a99989796959493929190610a08565b60405180910390f35b61009561043a565b6040516100a29190610c6d565b60405180910390f35b6040518060400160405280600381526020017f332e30000000000000000000000000000000000000000000000000000000000081525081565b600080600001549080600101549080600201549080600301805461010790610cbe565b80601f016020809104026020016040519081016040528092919081815260200182805461013390610cbe565b80156101805780601f1061015557610100808354040283529160200191610180565b820191906000526020600020905b81548152906001019060200180831161016357829003601f168201915b50505050509080600401805461019590610cbe565b80601f01602080910402602001604051908101604052809291908181526020018280546101c190610cbe565b801561020e5780601f106101e35761010080835404028352916020019161020e565b820191906000526020600020905b8154815290600101906020018083116101f157829003601f168201915b5050505050908060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600601805461024990610cbe565b80601f016020809104026020016040519081016040528092919081815260200182805461027590610cbe565b80156102c25780601f10610297576101008083540402835291602001916102c2565b820191906000526020600020905b8154815290600101906020018083116102a557829003601f168201915b5050505050908060070154908060080160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600901805461030390610cbe565b80601f016020809104026020016040519081016040528092919081815260200182805461032f90610cbe565b801561037c5780601f106103515761010080835404028352916020019161037c565b820191906000526020600020905b81548152906001019060200180831161035f57829003601f168201915b50505050509080600a0160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600b0180546103b790610cbe565b80601f01602080910402602001604051908101604052809291908181526020018280546103e390610cbe565b80156104305780601f1061040557610100808354040283529160200191610430565b820191906000526020600020905b81548152906001019060200180831161041357829003601f168201915b505050505090508c565b610442610859565b60006040518061018001604052908160008201548152602001600182015481526020016002820154815260200160038201805461047e90610cbe565b80601f01602080910402602001604051908101604052809291908181526020018280546104aa90610cbe565b80156104f75780601f106104cc576101008083540402835291602001916104f7565b820191906000526020600020905b8154815290600101906020018083116104da57829003601f168201915b5050505050815260200160048201805461051090610cbe565b80601f016020809104026020016040519081016040528092919081815260200182805461053c90610cbe565b80156105895780601f1061055e57610100808354040283529160200191610589565b820191906000526020600020905b81548152906001019060200180831161056c57829003601f168201915b505050505081526020016005820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016006820180546105f890610cbe565b80601f016020809104026020016040519081016040528092919081815260200182805461062490610cbe565b80156106715780601f1061064657610100808354040283529160200191610671565b820191906000526020600020905b81548152906001019060200180831161065457829003601f168201915b50505050508152602001600782015481526020016008820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016009820180546106ea90610cbe565b80601f016020809104026020016040519081016040528092919081815260200182805461071690610cbe565b80156107635780601f1061073857610100808354040283529160200191610763565b820191906000526020600020905b81548152906001019060200180831161074657829003601f168201915b50505050508152602001600a820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600b820180546107d290610cbe565b80601f01602080910402602001604051908101604052809291908181526020018280546107fe90610cbe565b801561084b5780601f106108205761010080835404028352916020019161084b565b820191906000526020600020905b81548152906001019060200180831161082e57829003601f168201915b505050505081525050905090565b6040518061018001604052806000815260200160008152602001600081526020016060815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016060815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b600081519050919050565b600082825260208201905092915050565b60005b8381101561093657808201518184015260208101905061091b565b60008484015250505050565b6000601f19601f8301169050919050565b600061095e826108fc565b6109688185610907565b9350610978818560208601610918565b61098181610942565b840191505092915050565b600060208201905081810360008301526109a68184610953565b905092915050565b6000819050919050565b6109c1816109ae565b82525050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006109f2826109c7565b9050919050565b610a02816109e7565b82525050565b600061018082019050610a1e600083018f6109b8565b610a2b602083018e6109b8565b610a38604083018d6109b8565b8181036060830152610a4a818c610953565b90508181036080830152610a5e818b610953565b9050610a6d60a083018a6109f9565b81810360c0830152610a7f8189610953565b9050610a8e60e08301886109b8565b610a9c6101008301876109f9565b818103610120830152610aaf8186610953565b9050610abf6101408301856109f9565b818103610160830152610ad28184610953565b90509d9c50505050505050505050505050565b610aee816109ae565b82525050565b600082825260208201905092915050565b6000610b10826108fc565b610b1a8185610af4565b9350610b2a818560208601610918565b610b3381610942565b840191505092915050565b610b47816109e7565b82525050565b600061018083016000830151610b666000860182610ae5565b506020830151610b796020860182610ae5565b506040830151610b8c6040860182610ae5565b5060608301518482036060860152610ba48282610b05565b91505060808301518482036080860152610bbe8282610b05565b91505060a0830151610bd360a0860182610b3e565b5060c083015184820360c0860152610beb8282610b05565b91505060e0830151610c0060e0860182610ae5565b50610100830151610c15610100860182610b3e565b50610120830151848203610120860152610c2f8282610b05565b915050610140830151610c46610140860182610b3e565b50610160830151848203610160860152610c608282610b05565b9150508091505092915050565b60006020820190508181036000830152610c878184610b4d565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610cd657607f821691505b602082108103610ce957610ce8610c8f565b5b5091905056fea2646970667358221220c2d2bdcaa2f3fcb178f6fda86a9f809d7d61dbf6e79803482060b1daf029a80364736f6c63430008140033",
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

// GetTaskInfo is a free data retrieval call binding the contract method 0xbac64a2e.
//
// Solidity: function getTaskInfo() view returns((uint256,uint256,uint256,string,string,address,string,uint256,address,string,address,string))
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
// Solidity: function getTaskInfo() view returns((uint256,uint256,uint256,string,string,address,string,uint256,address,string,address,string))
func (_Task *TaskSession) GetTaskInfo() (ECPTaskTaskInfo, error) {
	return _Task.Contract.GetTaskInfo(&_Task.CallOpts)
}

// GetTaskInfo is a free data retrieval call binding the contract method 0xbac64a2e.
//
// Solidity: function getTaskInfo() view returns((uint256,uint256,uint256,string,string,address,string,uint256,address,string,address,string))
func (_Task *TaskCallerSession) GetTaskInfo() (ECPTaskTaskInfo, error) {
	return _Task.Contract.GetTaskInfo(&_Task.CallOpts)
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
