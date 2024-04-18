// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package account

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

// AccountMetaData contains all meta data concerning the Account contract.
var AccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nodeId\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_multiAddresses\",\"type\":\"string[]\"},{\"internalType\":\"uint8\",\"name\":\"_ubiFlag\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_beneficiaryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quota\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"name\":\"BeneficiaryChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"newMultiaddrs\",\"type\":\"string[]\"}],\"name\":\"MultiaddrsChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"ubiFlag\",\"type\":\"uint8\"}],\"name\":\"UBIFlagChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"taskId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"taskType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"zkType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"proof\",\"type\":\"string\"}],\"name\":\"UBIProofSubmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quota\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBeneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newQuota\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newExpiration\",\"type\":\"uint256\"}],\"name\":\"changeBeneficiary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"newMultiaddrs\",\"type\":\"string[]\"}],\"name\":\"changeMultiaddrs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwnerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newUbiFlag\",\"type\":\"uint8\"}],\"name\":\"changeUbiFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"multiAddresses\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_taskId\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_taskType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_zkType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_proof\",\"type\":\"string\"}],\"name\":\"submitUBIProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"taskId\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"taskType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"zkType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proof\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isSubmitted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ubiFlag\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200267338038062002673833981810160405281019062000037919062000579565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550836001908162000088919062000874565b508260029080519060200190620000a19291906200015a565b5081600360006101000a81548160ff021916908360ff16021790555060405180606001604052808273ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016000815250600460008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020155905050505050506200095b565b828054828255906000526020600020908101928215620001a7579160200282015b82811115620001a657825182908162000195919062000874565b50916020019190600101906200017b565b5b509050620001b69190620001ba565b5090565b5b80821115620001de5760008181620001d49190620001e2565b50600101620001bb565b5090565b508054620001f09062000663565b6000825580601f1062000204575062000225565b601f01602090049060005260206000209081019062000224919062000228565b5b50565b5b808211156200024357600081600090555060010162000229565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620002b08262000265565b810181811067ffffffffffffffff82111715620002d257620002d162000276565b5b80604052505050565b6000620002e762000247565b9050620002f58282620002a5565b919050565b600067ffffffffffffffff82111562000318576200031762000276565b5b620003238262000265565b9050602081019050919050565b60005b838110156200035057808201518184015260208101905062000333565b60008484015250505050565b6000620003736200036d84620002fa565b620002db565b90508281526020810184848401111562000392576200039162000260565b5b6200039f84828562000330565b509392505050565b600082601f830112620003bf57620003be6200025b565b5b8151620003d18482602086016200035c565b91505092915050565b600067ffffffffffffffff821115620003f857620003f762000276565b5b602082029050602081019050919050565b600080fd5b6000620004256200041f84620003da565b620002db565b905080838252602082019050602084028301858111156200044b576200044a62000409565b5b835b818110156200049957805167ffffffffffffffff8111156200047457620004736200025b565b5b808601620004838982620003a7565b855260208501945050506020810190506200044d565b5050509392505050565b600082601f830112620004bb57620004ba6200025b565b5b8151620004cd8482602086016200040e565b91505092915050565b600060ff82169050919050565b620004ee81620004d6565b8114620004fa57600080fd5b50565b6000815190506200050e81620004e3565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620005418262000514565b9050919050565b620005538162000534565b81146200055f57600080fd5b50565b600081519050620005738162000548565b92915050565b6000806000806080858703121562000596576200059562000251565b5b600085015167ffffffffffffffff811115620005b757620005b662000256565b5b620005c587828801620003a7565b945050602085015167ffffffffffffffff811115620005e957620005e862000256565b5b620005f787828801620004a3565b93505060406200060a87828801620004fd565b92505060606200061d8782880162000562565b91505092959194509250565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200067c57607f821691505b60208210810362000692576200069162000634565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620006fc7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620006bd565b620007088683620006bd565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620007556200074f620007498462000720565b6200072a565b62000720565b9050919050565b6000819050919050565b620007718362000734565b6200078962000780826200075c565b848454620006ca565b825550505050565b600090565b620007a062000791565b620007ad81848462000766565b505050565b5b81811015620007d557620007c960008262000796565b600181019050620007b3565b5050565b601f8211156200082457620007ee8162000698565b620007f984620006ad565b8101602085101562000809578190505b620008216200081885620006ad565b830182620007b2565b50505b505050565b600082821c905092915050565b6000620008496000198460080262000829565b1980831691505092915050565b600062000864838362000836565b9150826002028217905092915050565b6200087f8262000629565b67ffffffffffffffff8111156200089b576200089a62000276565b5b620008a7825462000663565b620008b4828285620007d9565b600060209050601f831160018114620008ec5760008415620008d7578287015190505b620008e3858262000856565b86555062000953565b601f198416620008fc8662000698565b60005b828110156200092657848901518255600182019150602085019450602081019050620008ff565b8683101562000946578489015162000942601f89168262000836565b8355505b6001600288020188555050505b505050505050565b611d08806200096b6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c806385eac05f1161008c57806394f219381161006657806394f2193814610206578063db613e8114610222578063e8deed9714610246578063ee2c5e8914610264576100cf565b806385eac05f146101ae578063893d20e8146101ca5780638da5cb5b146101e8576100cf565b8063139d7fed146100d457806338af3eed146100f25780633d1333a51461011257806358709cf2146101425780636553b6ef1461017657806377ecf4df14610192575b600080fd5b6100dc610280565b6040516100e99190611016565b60405180910390f35b6100fa61030e565b60405161010993929190611092565b60405180910390f35b61012c60048036038101906101279190611109565b610346565b6040516101399190611016565b60405180910390f35b61015c6004803603810190610157919061126b565b6103f2565b60405161016d9594939291906112eb565b60405180910390f35b610190600480360381019061018b919061137f565b6105f0565b005b6101ac60048036038101906101a791906113ac565b6106d3565b005b6101c860048036038101906101c39190611493565b6108fb565b005b6101d2610a26565b6040516101df91906114c0565b60405180910390f35b6101f0610a4f565b6040516101fd91906114c0565b60405180910390f35b610220600480360381019061021b91906115c1565b610a73565b005b61022a610b52565b60405161023d9796959493929190611716565b60405180910390f35b61024e610d3a565b60405161025b9190611793565b60405180910390f35b61027e600480360381019061027991906117ae565b610d4d565b005b6001805461028d90611830565b80601f01602080910402602001604051908101604052809291908181526020018280546102b990611830565b80156103065780601f106102db57610100808354040283529160200191610306565b820191906000526020600020905b8154815290600101906020018083116102e957829003601f168201915b505050505081565b60048060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154905083565b6002818154811061035657600080fd5b90600052602060002001600091509050805461037190611830565b80601f016020809104026020016040519081016040528092919081815260200182805461039d90611830565b80156103ea5780601f106103bf576101008083540402835291602001916103ea565b820191906000526020600020905b8154815290600101906020018083116103cd57829003601f168201915b505050505081565b60078180516020810182018051848252602083016020850120818352809550505050505060009150905080600001805461042b90611830565b80601f016020809104026020016040519081016040528092919081815260200182805461045790611830565b80156104a45780601f10610479576101008083540402835291602001916104a4565b820191906000526020600020905b81548152906001019060200180831161048757829003601f168201915b5050505050908060010160009054906101000a900460ff16908060020180546104cc90611830565b80601f01602080910402602001604051908101604052809291908181526020018280546104f890611830565b80156105455780601f1061051a57610100808354040283529160200191610545565b820191906000526020600020905b81548152906001019060200180831161052857829003601f168201915b50505050509080600301805461055a90611830565b80601f016020809104026020016040519081016040528092919081815260200182805461058690611830565b80156105d35780601f106105a8576101008083540402835291602001916105d3565b820191906000526020600020905b8154815290600101906020018083116105b657829003601f168201915b5050505050908060040160009054906101000a900460ff16905085565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461067e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610675906118d3565b60405180910390fd5b80600360006101000a81548160ff021916908360ff1602179055507f2504aabda145c6cbc7a70eee13f3abeb510de8e65904c53b824d58ff9fc5b84a816040516106c89190611793565b60405180910390a150565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610761576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610758906118d3565b60405180910390fd5b600784604051610771919061192f565b908152602001604051809103902060040160009054906101000a900460ff16156107d0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107c7906119b8565b60405180910390fd5b6040518060a001604052808581526020018460ff1681526020018381526020018281526020016001151581525060078560405161080d919061192f565b908152602001604051809103902060008201518160000190816108309190611b84565b5060208201518160010160006101000a81548160ff021916908360ff16021790555060408201518160020190816108679190611b84565b50606082015181600301908161087d9190611b84565b5060808201518160040160006101000a81548160ff0219169083151502179055509050503373ffffffffffffffffffffffffffffffffffffffff167f18a95158362c8c4a1cacd8dc12f81e1fd419696c958630a3c06e3466f4daad00858585856040516108ed9493929190611c56565b60405180910390a250505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610989576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610980906118d3565b60405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a350565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b01576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610af8906118d3565b60405180910390fd5b8060029080519060200190610b17929190610eac565b507f34f0cbd777d8d394caa00ae618bb83ee053819e2943bbd10c6c61019d4c5758281604051610b479190611cb0565b60405180910390a150565b600060608060008060008060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660016002600360009054906101000a900460ff16600460000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600460010154600460020154858054610bd090611830565b80601f0160208091040260200160405190810160405280929190818152602001828054610bfc90611830565b8015610c495780601f10610c1e57610100808354040283529160200191610c49565b820191906000526020600020905b815481529060010190602001808311610c2c57829003601f168201915b5050505050955084805480602002602001604051908101604052809291908181526020016000905b82821015610d1d578382906000526020600020018054610c9090611830565b80601f0160208091040260200160405190810160405280929190818152602001828054610cbc90611830565b8015610d095780601f10610cde57610100808354040283529160200191610d09565b820191906000526020600020905b815481529060010190602001808311610cec57829003601f168201915b505050505081526020019060010190610c71565b505050509450965096509650965096509650965090919293949596565b600360009054906101000a900460ff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ddb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dd2906118d3565b60405180910390fd5b60405180606001604052808473ffffffffffffffffffffffffffffffffffffffff16815260200183815260200182815250600460008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155604082015181600201559050507f21c5f4c4149658f3fb981d0f3b26cb82ad95e6aab98e5a3dc73f9a2d6ab6b311838383604051610e9f93929190611092565b60405180910390a1505050565b828054828255906000526020600020908101928215610ef4579160200282015b82811115610ef3578251829081610ee39190611b84565b5091602001919060010190610ecc565b5b509050610f019190610f05565b5090565b5b80821115610f255760008181610f1c9190610f29565b50600101610f06565b5090565b508054610f3590611830565b6000825580601f10610f475750610f66565b601f016020900490600052602060002090810190610f659190610f69565b5b50565b5b80821115610f82576000816000905550600101610f6a565b5090565b600081519050919050565b600082825260208201905092915050565b60005b83811015610fc0578082015181840152602081019050610fa5565b60008484015250505050565b6000601f19601f8301169050919050565b6000610fe882610f86565b610ff28185610f91565b9350611002818560208601610fa2565b61100b81610fcc565b840191505092915050565b600060208201905081810360008301526110308184610fdd565b905092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061106382611038565b9050919050565b61107381611058565b82525050565b6000819050919050565b61108c81611079565b82525050565b60006060820190506110a7600083018661106a565b6110b46020830185611083565b6110c16040830184611083565b949350505050565b6000604051905090565b600080fd5b600080fd5b6110e681611079565b81146110f157600080fd5b50565b600081359050611103816110dd565b92915050565b60006020828403121561111f5761111e6110d3565b5b600061112d848285016110f4565b91505092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61117882610fcc565b810181811067ffffffffffffffff8211171561119757611196611140565b5b80604052505050565b60006111aa6110c9565b90506111b6828261116f565b919050565b600067ffffffffffffffff8211156111d6576111d5611140565b5b6111df82610fcc565b9050602081019050919050565b82818337600083830152505050565b600061120e611209846111bb565b6111a0565b90508281526020810184848401111561122a5761122961113b565b5b6112358482856111ec565b509392505050565b600082601f83011261125257611251611136565b5b81356112628482602086016111fb565b91505092915050565b600060208284031215611281576112806110d3565b5b600082013567ffffffffffffffff81111561129f5761129e6110d8565b5b6112ab8482850161123d565b91505092915050565b600060ff82169050919050565b6112ca816112b4565b82525050565b60008115159050919050565b6112e5816112d0565b82525050565b600060a08201905081810360008301526113058188610fdd565b905061131460208301876112c1565b81810360408301526113268186610fdd565b9050818103606083015261133a8185610fdd565b905061134960808301846112dc565b9695505050505050565b61135c816112b4565b811461136757600080fd5b50565b60008135905061137981611353565b92915050565b600060208284031215611395576113946110d3565b5b60006113a38482850161136a565b91505092915050565b600080600080608085870312156113c6576113c56110d3565b5b600085013567ffffffffffffffff8111156113e4576113e36110d8565b5b6113f08782880161123d565b94505060206114018782880161136a565b935050604085013567ffffffffffffffff811115611422576114216110d8565b5b61142e8782880161123d565b925050606085013567ffffffffffffffff81111561144f5761144e6110d8565b5b61145b8782880161123d565b91505092959194509250565b61147081611058565b811461147b57600080fd5b50565b60008135905061148d81611467565b92915050565b6000602082840312156114a9576114a86110d3565b5b60006114b78482850161147e565b91505092915050565b60006020820190506114d5600083018461106a565b92915050565b600067ffffffffffffffff8211156114f6576114f5611140565b5b602082029050602081019050919050565b600080fd5b600061151f61151a846114db565b6111a0565b9050808382526020820190506020840283018581111561154257611541611507565b5b835b8181101561158957803567ffffffffffffffff81111561156757611566611136565b5b808601611574898261123d565b85526020850194505050602081019050611544565b5050509392505050565b600082601f8301126115a8576115a7611136565b5b81356115b884826020860161150c565b91505092915050565b6000602082840312156115d7576115d66110d3565b5b600082013567ffffffffffffffff8111156115f5576115f46110d8565b5b61160184828501611593565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600082825260208201905092915050565b600061165282610f86565b61165c8185611636565b935061166c818560208601610fa2565b61167581610fcc565b840191505092915050565b600061168c8383611647565b905092915050565b6000602082019050919050565b60006116ac8261160a565b6116b68185611615565b9350836020820285016116c885611626565b8060005b8581101561170457848403895281516116e58582611680565b94506116f083611694565b925060208a019950506001810190506116cc565b50829750879550505050505092915050565b600060e08201905061172b600083018a61106a565b818103602083015261173d8189610fdd565b9050818103604083015261175181886116a1565b905061176060608301876112c1565b61176d608083018661106a565b61177a60a0830185611083565b61178760c0830184611083565b98975050505050505050565b60006020820190506117a860008301846112c1565b92915050565b6000806000606084860312156117c7576117c66110d3565b5b60006117d58682870161147e565b93505060206117e6868287016110f4565b92505060406117f7868287016110f4565b9150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061184857607f821691505b60208210810361185b5761185a611801565b5b50919050565b7f4f6e6c79206f776e65722063616e2063616c6c20746869732066756e6374696f60008201527f6e2e000000000000000000000000000000000000000000000000000000000000602082015250565b60006118bd602283610f91565b91506118c882611861565b604082019050919050565b600060208201905081810360008301526118ec816118b0565b9050919050565b600081905092915050565b600061190982610f86565b61191381856118f3565b9350611923818560208601610fa2565b80840191505092915050565b600061193b82846118fe565b915081905092915050565b7f50726f6f6620666f722074686973207461736b20697320616c7265616479207360008201527f75626d69747465642e0000000000000000000000000000000000000000000000602082015250565b60006119a2602983610f91565b91506119ad82611946565b604082019050919050565b600060208201905081810360008301526119d181611995565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302611a3a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826119fd565b611a4486836119fd565b95508019841693508086168417925050509392505050565b6000819050919050565b6000611a81611a7c611a7784611079565b611a5c565b611079565b9050919050565b6000819050919050565b611a9b83611a66565b611aaf611aa782611a88565b848454611a0a565b825550505050565b600090565b611ac4611ab7565b611acf818484611a92565b505050565b5b81811015611af357611ae8600082611abc565b600181019050611ad5565b5050565b601f821115611b3857611b09816119d8565b611b12846119ed565b81016020851015611b21578190505b611b35611b2d856119ed565b830182611ad4565b50505b505050565b600082821c905092915050565b6000611b5b60001984600802611b3d565b1980831691505092915050565b6000611b748383611b4a565b9150826002028217905092915050565b611b8d82610f86565b67ffffffffffffffff811115611ba657611ba5611140565b5b611bb08254611830565b611bbb828285611af7565b600060209050601f831160018114611bee5760008415611bdc578287015190505b611be68582611b68565b865550611c4e565b601f198416611bfc866119d8565b60005b82811015611c2457848901518255600182019150602085019450602081019050611bff565b86831015611c415784890151611c3d601f891682611b4a565b8355505b6001600288020188555050505b505050505050565b60006080820190508181036000830152611c708187610fdd565b9050611c7f60208301866112c1565b8181036040830152611c918185610fdd565b90508181036060830152611ca58184610fdd565b905095945050505050565b60006020820190508181036000830152611cca81846116a1565b90509291505056fea26469706673582212207ed44ce0b16f8d7f6e87ed9136f2c1b9da10b7d3ba7415e99495720d00a9913664736f6c63430008140033",
}

// AccountABI is the input ABI used to generate the binding from.
// Deprecated: Use AccountMetaData.ABI instead.
var AccountABI = AccountMetaData.ABI

// AccountBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AccountMetaData.Bin instead.
var AccountBin = AccountMetaData.Bin

// DeployAccount deploys a new Ethereum contract, binding an instance of Account to it.
func DeployAccount(auth *bind.TransactOpts, backend bind.ContractBackend, _nodeId string, _multiAddresses []string, _ubiFlag uint8, _beneficiaryAddress common.Address) (common.Address, *types.Transaction, *Account, error) {
	parsed, err := AccountMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AccountBin), backend, _nodeId, _multiAddresses, _ubiFlag, _beneficiaryAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Account{AccountCaller: AccountCaller{contract: contract}, AccountTransactor: AccountTransactor{contract: contract}, AccountFilterer: AccountFilterer{contract: contract}}, nil
}

// Account is an auto generated Go binding around an Ethereum contract.
type Account struct {
	AccountCaller     // Read-only binding to the contract
	AccountTransactor // Write-only binding to the contract
	AccountFilterer   // Log filterer for contract events
}

// AccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountSession struct {
	Contract     *Account          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountCallerSession struct {
	Contract *AccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountTransactorSession struct {
	Contract     *AccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountRaw struct {
	Contract *Account // Generic contract binding to access the raw methods on
}

// AccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountCallerRaw struct {
	Contract *AccountCaller // Generic read-only contract binding to access the raw methods on
}

// AccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountTransactorRaw struct {
	Contract *AccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccount creates a new instance of Account, bound to a specific deployed contract.
func NewAccount(address common.Address, backend bind.ContractBackend) (*Account, error) {
	contract, err := bindAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Account{AccountCaller: AccountCaller{contract: contract}, AccountTransactor: AccountTransactor{contract: contract}, AccountFilterer: AccountFilterer{contract: contract}}, nil
}

// NewAccountCaller creates a new read-only instance of Account, bound to a specific deployed contract.
func NewAccountCaller(address common.Address, caller bind.ContractCaller) (*AccountCaller, error) {
	contract, err := bindAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountCaller{contract: contract}, nil
}

// NewAccountTransactor creates a new write-only instance of Account, bound to a specific deployed contract.
func NewAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountTransactor, error) {
	contract, err := bindAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountTransactor{contract: contract}, nil
}

// NewAccountFilterer creates a new log filterer instance of Account, bound to a specific deployed contract.
func NewAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountFilterer, error) {
	contract, err := bindAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountFilterer{contract: contract}, nil
}

// bindAccount binds a generic wrapper to an already deployed contract.
func bindAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Account *AccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Account.Contract.AccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Account *AccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Account.Contract.AccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Account *AccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Account.Contract.AccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Account *AccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Account.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Account *AccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Account.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Account *AccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Account.Contract.contract.Transact(opts, method, params...)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address beneficiaryAddress, uint256 quota, uint256 expiration)
func (_Account *AccountCaller) Beneficiary(opts *bind.CallOpts) (struct {
	BeneficiaryAddress common.Address
	Quota              *big.Int
	Expiration         *big.Int
}, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "beneficiary")

	outstruct := new(struct {
		BeneficiaryAddress common.Address
		Quota              *big.Int
		Expiration         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BeneficiaryAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Quota = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Expiration = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address beneficiaryAddress, uint256 quota, uint256 expiration)
func (_Account *AccountSession) Beneficiary() (struct {
	BeneficiaryAddress common.Address
	Quota              *big.Int
	Expiration         *big.Int
}, error) {
	return _Account.Contract.Beneficiary(&_Account.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address beneficiaryAddress, uint256 quota, uint256 expiration)
func (_Account *AccountCallerSession) Beneficiary() (struct {
	BeneficiaryAddress common.Address
	Quota              *big.Int
	Expiration         *big.Int
}, error) {
	return _Account.Contract.Beneficiary(&_Account.CallOpts)
}

// GetAccount is a free data retrieval call binding the contract method 0xdb613e81.
//
// Solidity: function getAccount() view returns(address, string, string[], uint8, address, uint256, uint256)
func (_Account *AccountCaller) GetAccount(opts *bind.CallOpts) (common.Address, string, []string, uint8, common.Address, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "getAccount")

	if err != nil {
		return *new(common.Address), *new(string), *new([]string), *new(uint8), *new(common.Address), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new([]string)).(*[]string)
	out3 := *abi.ConvertType(out[3], new(uint8)).(*uint8)
	out4 := *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, err

}

// GetAccount is a free data retrieval call binding the contract method 0xdb613e81.
//
// Solidity: function getAccount() view returns(address, string, string[], uint8, address, uint256, uint256)
func (_Account *AccountSession) GetAccount() (common.Address, string, []string, uint8, common.Address, *big.Int, *big.Int, error) {
	return _Account.Contract.GetAccount(&_Account.CallOpts)
}

// GetAccount is a free data retrieval call binding the contract method 0xdb613e81.
//
// Solidity: function getAccount() view returns(address, string, string[], uint8, address, uint256, uint256)
func (_Account *AccountCallerSession) GetAccount() (common.Address, string, []string, uint8, common.Address, *big.Int, *big.Int, error) {
	return _Account.Contract.GetAccount(&_Account.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Account *AccountCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Account *AccountSession) GetOwner() (common.Address, error) {
	return _Account.Contract.GetOwner(&_Account.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Account *AccountCallerSession) GetOwner() (common.Address, error) {
	return _Account.Contract.GetOwner(&_Account.CallOpts)
}

// MultiAddresses is a free data retrieval call binding the contract method 0x3d1333a5.
//
// Solidity: function multiAddresses(uint256 ) view returns(string)
func (_Account *AccountCaller) MultiAddresses(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "multiAddresses", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MultiAddresses is a free data retrieval call binding the contract method 0x3d1333a5.
//
// Solidity: function multiAddresses(uint256 ) view returns(string)
func (_Account *AccountSession) MultiAddresses(arg0 *big.Int) (string, error) {
	return _Account.Contract.MultiAddresses(&_Account.CallOpts, arg0)
}

// MultiAddresses is a free data retrieval call binding the contract method 0x3d1333a5.
//
// Solidity: function multiAddresses(uint256 ) view returns(string)
func (_Account *AccountCallerSession) MultiAddresses(arg0 *big.Int) (string, error) {
	return _Account.Contract.MultiAddresses(&_Account.CallOpts, arg0)
}

// NodeId is a free data retrieval call binding the contract method 0x139d7fed.
//
// Solidity: function nodeId() view returns(string)
func (_Account *AccountCaller) NodeId(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "nodeId")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NodeId is a free data retrieval call binding the contract method 0x139d7fed.
//
// Solidity: function nodeId() view returns(string)
func (_Account *AccountSession) NodeId() (string, error) {
	return _Account.Contract.NodeId(&_Account.CallOpts)
}

// NodeId is a free data retrieval call binding the contract method 0x139d7fed.
//
// Solidity: function nodeId() view returns(string)
func (_Account *AccountCallerSession) NodeId() (string, error) {
	return _Account.Contract.NodeId(&_Account.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Account *AccountCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Account *AccountSession) Owner() (common.Address, error) {
	return _Account.Contract.Owner(&_Account.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Account *AccountCallerSession) Owner() (common.Address, error) {
	return _Account.Contract.Owner(&_Account.CallOpts)
}

// Tasks is a free data retrieval call binding the contract method 0x58709cf2.
//
// Solidity: function tasks(string ) view returns(string taskId, uint8 taskType, string zkType, string proof, bool isSubmitted)
func (_Account *AccountCaller) Tasks(opts *bind.CallOpts, arg0 string) (struct {
	TaskId      string
	TaskType    uint8
	ZkType      string
	Proof       string
	IsSubmitted bool
}, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "tasks", arg0)

	outstruct := new(struct {
		TaskId      string
		TaskType    uint8
		ZkType      string
		Proof       string
		IsSubmitted bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TaskId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.TaskType = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.ZkType = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Proof = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.IsSubmitted = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Tasks is a free data retrieval call binding the contract method 0x58709cf2.
//
// Solidity: function tasks(string ) view returns(string taskId, uint8 taskType, string zkType, string proof, bool isSubmitted)
func (_Account *AccountSession) Tasks(arg0 string) (struct {
	TaskId      string
	TaskType    uint8
	ZkType      string
	Proof       string
	IsSubmitted bool
}, error) {
	return _Account.Contract.Tasks(&_Account.CallOpts, arg0)
}

// Tasks is a free data retrieval call binding the contract method 0x58709cf2.
//
// Solidity: function tasks(string ) view returns(string taskId, uint8 taskType, string zkType, string proof, bool isSubmitted)
func (_Account *AccountCallerSession) Tasks(arg0 string) (struct {
	TaskId      string
	TaskType    uint8
	ZkType      string
	Proof       string
	IsSubmitted bool
}, error) {
	return _Account.Contract.Tasks(&_Account.CallOpts, arg0)
}

// UbiFlag is a free data retrieval call binding the contract method 0xe8deed97.
//
// Solidity: function ubiFlag() view returns(uint8)
func (_Account *AccountCaller) UbiFlag(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "ubiFlag")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// UbiFlag is a free data retrieval call binding the contract method 0xe8deed97.
//
// Solidity: function ubiFlag() view returns(uint8)
func (_Account *AccountSession) UbiFlag() (uint8, error) {
	return _Account.Contract.UbiFlag(&_Account.CallOpts)
}

// UbiFlag is a free data retrieval call binding the contract method 0xe8deed97.
//
// Solidity: function ubiFlag() view returns(uint8)
func (_Account *AccountCallerSession) UbiFlag() (uint8, error) {
	return _Account.Contract.UbiFlag(&_Account.CallOpts)
}

// ChangeBeneficiary is a paid mutator transaction binding the contract method 0xee2c5e89.
//
// Solidity: function changeBeneficiary(address newBeneficiary, uint256 newQuota, uint256 newExpiration) returns()
func (_Account *AccountTransactor) ChangeBeneficiary(opts *bind.TransactOpts, newBeneficiary common.Address, newQuota *big.Int, newExpiration *big.Int) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "changeBeneficiary", newBeneficiary, newQuota, newExpiration)
}

// ChangeBeneficiary is a paid mutator transaction binding the contract method 0xee2c5e89.
//
// Solidity: function changeBeneficiary(address newBeneficiary, uint256 newQuota, uint256 newExpiration) returns()
func (_Account *AccountSession) ChangeBeneficiary(newBeneficiary common.Address, newQuota *big.Int, newExpiration *big.Int) (*types.Transaction, error) {
	return _Account.Contract.ChangeBeneficiary(&_Account.TransactOpts, newBeneficiary, newQuota, newExpiration)
}

// ChangeBeneficiary is a paid mutator transaction binding the contract method 0xee2c5e89.
//
// Solidity: function changeBeneficiary(address newBeneficiary, uint256 newQuota, uint256 newExpiration) returns()
func (_Account *AccountTransactorSession) ChangeBeneficiary(newBeneficiary common.Address, newQuota *big.Int, newExpiration *big.Int) (*types.Transaction, error) {
	return _Account.Contract.ChangeBeneficiary(&_Account.TransactOpts, newBeneficiary, newQuota, newExpiration)
}

// ChangeMultiaddrs is a paid mutator transaction binding the contract method 0x94f21938.
//
// Solidity: function changeMultiaddrs(string[] newMultiaddrs) returns()
func (_Account *AccountTransactor) ChangeMultiaddrs(opts *bind.TransactOpts, newMultiaddrs []string) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "changeMultiaddrs", newMultiaddrs)
}

// ChangeMultiaddrs is a paid mutator transaction binding the contract method 0x94f21938.
//
// Solidity: function changeMultiaddrs(string[] newMultiaddrs) returns()
func (_Account *AccountSession) ChangeMultiaddrs(newMultiaddrs []string) (*types.Transaction, error) {
	return _Account.Contract.ChangeMultiaddrs(&_Account.TransactOpts, newMultiaddrs)
}

// ChangeMultiaddrs is a paid mutator transaction binding the contract method 0x94f21938.
//
// Solidity: function changeMultiaddrs(string[] newMultiaddrs) returns()
func (_Account *AccountTransactorSession) ChangeMultiaddrs(newMultiaddrs []string) (*types.Transaction, error) {
	return _Account.Contract.ChangeMultiaddrs(&_Account.TransactOpts, newMultiaddrs)
}

// ChangeOwnerAddress is a paid mutator transaction binding the contract method 0x85eac05f.
//
// Solidity: function changeOwnerAddress(address newOwner) returns()
func (_Account *AccountTransactor) ChangeOwnerAddress(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "changeOwnerAddress", newOwner)
}

// ChangeOwnerAddress is a paid mutator transaction binding the contract method 0x85eac05f.
//
// Solidity: function changeOwnerAddress(address newOwner) returns()
func (_Account *AccountSession) ChangeOwnerAddress(newOwner common.Address) (*types.Transaction, error) {
	return _Account.Contract.ChangeOwnerAddress(&_Account.TransactOpts, newOwner)
}

// ChangeOwnerAddress is a paid mutator transaction binding the contract method 0x85eac05f.
//
// Solidity: function changeOwnerAddress(address newOwner) returns()
func (_Account *AccountTransactorSession) ChangeOwnerAddress(newOwner common.Address) (*types.Transaction, error) {
	return _Account.Contract.ChangeOwnerAddress(&_Account.TransactOpts, newOwner)
}

// ChangeUbiFlag is a paid mutator transaction binding the contract method 0x6553b6ef.
//
// Solidity: function changeUbiFlag(uint8 newUbiFlag) returns()
func (_Account *AccountTransactor) ChangeUbiFlag(opts *bind.TransactOpts, newUbiFlag uint8) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "changeUbiFlag", newUbiFlag)
}

// ChangeUbiFlag is a paid mutator transaction binding the contract method 0x6553b6ef.
//
// Solidity: function changeUbiFlag(uint8 newUbiFlag) returns()
func (_Account *AccountSession) ChangeUbiFlag(newUbiFlag uint8) (*types.Transaction, error) {
	return _Account.Contract.ChangeUbiFlag(&_Account.TransactOpts, newUbiFlag)
}

// ChangeUbiFlag is a paid mutator transaction binding the contract method 0x6553b6ef.
//
// Solidity: function changeUbiFlag(uint8 newUbiFlag) returns()
func (_Account *AccountTransactorSession) ChangeUbiFlag(newUbiFlag uint8) (*types.Transaction, error) {
	return _Account.Contract.ChangeUbiFlag(&_Account.TransactOpts, newUbiFlag)
}

// SubmitUBIProof is a paid mutator transaction binding the contract method 0x77ecf4df.
//
// Solidity: function submitUBIProof(string _taskId, uint8 _taskType, string _zkType, string _proof) returns()
func (_Account *AccountTransactor) SubmitUBIProof(opts *bind.TransactOpts, _taskId string, _taskType uint8, _zkType string, _proof string) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "submitUBIProof", _taskId, _taskType, _zkType, _proof)
}

// SubmitUBIProof is a paid mutator transaction binding the contract method 0x77ecf4df.
//
// Solidity: function submitUBIProof(string _taskId, uint8 _taskType, string _zkType, string _proof) returns()
func (_Account *AccountSession) SubmitUBIProof(_taskId string, _taskType uint8, _zkType string, _proof string) (*types.Transaction, error) {
	return _Account.Contract.SubmitUBIProof(&_Account.TransactOpts, _taskId, _taskType, _zkType, _proof)
}

// SubmitUBIProof is a paid mutator transaction binding the contract method 0x77ecf4df.
//
// Solidity: function submitUBIProof(string _taskId, uint8 _taskType, string _zkType, string _proof) returns()
func (_Account *AccountTransactorSession) SubmitUBIProof(_taskId string, _taskType uint8, _zkType string, _proof string) (*types.Transaction, error) {
	return _Account.Contract.SubmitUBIProof(&_Account.TransactOpts, _taskId, _taskType, _zkType, _proof)
}

// AccountBeneficiaryChangedIterator is returned from FilterBeneficiaryChanged and is used to iterate over the raw logs and unpacked data for BeneficiaryChanged events raised by the Account contract.
type AccountBeneficiaryChangedIterator struct {
	Event *AccountBeneficiaryChanged // Event containing the contract specifics and raw log

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
func (it *AccountBeneficiaryChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountBeneficiaryChanged)
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
		it.Event = new(AccountBeneficiaryChanged)
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
func (it *AccountBeneficiaryChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountBeneficiaryChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountBeneficiaryChanged represents a BeneficiaryChanged event raised by the Account contract.
type AccountBeneficiaryChanged struct {
	Beneficiary common.Address
	Quota       *big.Int
	Expiration  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBeneficiaryChanged is a free log retrieval operation binding the contract event 0x21c5f4c4149658f3fb981d0f3b26cb82ad95e6aab98e5a3dc73f9a2d6ab6b311.
//
// Solidity: event BeneficiaryChanged(address beneficiary, uint256 quota, uint256 expiration)
func (_Account *AccountFilterer) FilterBeneficiaryChanged(opts *bind.FilterOpts) (*AccountBeneficiaryChangedIterator, error) {

	logs, sub, err := _Account.contract.FilterLogs(opts, "BeneficiaryChanged")
	if err != nil {
		return nil, err
	}
	return &AccountBeneficiaryChangedIterator{contract: _Account.contract, event: "BeneficiaryChanged", logs: logs, sub: sub}, nil
}

// WatchBeneficiaryChanged is a free log subscription operation binding the contract event 0x21c5f4c4149658f3fb981d0f3b26cb82ad95e6aab98e5a3dc73f9a2d6ab6b311.
//
// Solidity: event BeneficiaryChanged(address beneficiary, uint256 quota, uint256 expiration)
func (_Account *AccountFilterer) WatchBeneficiaryChanged(opts *bind.WatchOpts, sink chan<- *AccountBeneficiaryChanged) (event.Subscription, error) {

	logs, sub, err := _Account.contract.WatchLogs(opts, "BeneficiaryChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountBeneficiaryChanged)
				if err := _Account.contract.UnpackLog(event, "BeneficiaryChanged", log); err != nil {
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

// ParseBeneficiaryChanged is a log parse operation binding the contract event 0x21c5f4c4149658f3fb981d0f3b26cb82ad95e6aab98e5a3dc73f9a2d6ab6b311.
//
// Solidity: event BeneficiaryChanged(address beneficiary, uint256 quota, uint256 expiration)
func (_Account *AccountFilterer) ParseBeneficiaryChanged(log types.Log) (*AccountBeneficiaryChanged, error) {
	event := new(AccountBeneficiaryChanged)
	if err := _Account.contract.UnpackLog(event, "BeneficiaryChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountMultiaddrsChangedIterator is returned from FilterMultiaddrsChanged and is used to iterate over the raw logs and unpacked data for MultiaddrsChanged events raised by the Account contract.
type AccountMultiaddrsChangedIterator struct {
	Event *AccountMultiaddrsChanged // Event containing the contract specifics and raw log

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
func (it *AccountMultiaddrsChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountMultiaddrsChanged)
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
		it.Event = new(AccountMultiaddrsChanged)
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
func (it *AccountMultiaddrsChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountMultiaddrsChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountMultiaddrsChanged represents a MultiaddrsChanged event raised by the Account contract.
type AccountMultiaddrsChanged struct {
	NewMultiaddrs []string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMultiaddrsChanged is a free log retrieval operation binding the contract event 0x34f0cbd777d8d394caa00ae618bb83ee053819e2943bbd10c6c61019d4c57582.
//
// Solidity: event MultiaddrsChanged(string[] newMultiaddrs)
func (_Account *AccountFilterer) FilterMultiaddrsChanged(opts *bind.FilterOpts) (*AccountMultiaddrsChangedIterator, error) {

	logs, sub, err := _Account.contract.FilterLogs(opts, "MultiaddrsChanged")
	if err != nil {
		return nil, err
	}
	return &AccountMultiaddrsChangedIterator{contract: _Account.contract, event: "MultiaddrsChanged", logs: logs, sub: sub}, nil
}

// WatchMultiaddrsChanged is a free log subscription operation binding the contract event 0x34f0cbd777d8d394caa00ae618bb83ee053819e2943bbd10c6c61019d4c57582.
//
// Solidity: event MultiaddrsChanged(string[] newMultiaddrs)
func (_Account *AccountFilterer) WatchMultiaddrsChanged(opts *bind.WatchOpts, sink chan<- *AccountMultiaddrsChanged) (event.Subscription, error) {

	logs, sub, err := _Account.contract.WatchLogs(opts, "MultiaddrsChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountMultiaddrsChanged)
				if err := _Account.contract.UnpackLog(event, "MultiaddrsChanged", log); err != nil {
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

// ParseMultiaddrsChanged is a log parse operation binding the contract event 0x34f0cbd777d8d394caa00ae618bb83ee053819e2943bbd10c6c61019d4c57582.
//
// Solidity: event MultiaddrsChanged(string[] newMultiaddrs)
func (_Account *AccountFilterer) ParseMultiaddrsChanged(log types.Log) (*AccountMultiaddrsChanged, error) {
	event := new(AccountMultiaddrsChanged)
	if err := _Account.contract.UnpackLog(event, "MultiaddrsChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Account contract.
type AccountOwnershipTransferredIterator struct {
	Event *AccountOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AccountOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountOwnershipTransferred)
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
		it.Event = new(AccountOwnershipTransferred)
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
func (it *AccountOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountOwnershipTransferred represents a OwnershipTransferred event raised by the Account contract.
type AccountOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Account *AccountFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AccountOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Account.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AccountOwnershipTransferredIterator{contract: _Account.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Account *AccountFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AccountOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Account.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountOwnershipTransferred)
				if err := _Account.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Account *AccountFilterer) ParseOwnershipTransferred(log types.Log) (*AccountOwnershipTransferred, error) {
	event := new(AccountOwnershipTransferred)
	if err := _Account.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountUBIFlagChangedIterator is returned from FilterUBIFlagChanged and is used to iterate over the raw logs and unpacked data for UBIFlagChanged events raised by the Account contract.
type AccountUBIFlagChangedIterator struct {
	Event *AccountUBIFlagChanged // Event containing the contract specifics and raw log

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
func (it *AccountUBIFlagChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountUBIFlagChanged)
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
		it.Event = new(AccountUBIFlagChanged)
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
func (it *AccountUBIFlagChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountUBIFlagChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountUBIFlagChanged represents a UBIFlagChanged event raised by the Account contract.
type AccountUBIFlagChanged struct {
	UbiFlag uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUBIFlagChanged is a free log retrieval operation binding the contract event 0x2504aabda145c6cbc7a70eee13f3abeb510de8e65904c53b824d58ff9fc5b84a.
//
// Solidity: event UBIFlagChanged(uint8 ubiFlag)
func (_Account *AccountFilterer) FilterUBIFlagChanged(opts *bind.FilterOpts) (*AccountUBIFlagChangedIterator, error) {

	logs, sub, err := _Account.contract.FilterLogs(opts, "UBIFlagChanged")
	if err != nil {
		return nil, err
	}
	return &AccountUBIFlagChangedIterator{contract: _Account.contract, event: "UBIFlagChanged", logs: logs, sub: sub}, nil
}

// WatchUBIFlagChanged is a free log subscription operation binding the contract event 0x2504aabda145c6cbc7a70eee13f3abeb510de8e65904c53b824d58ff9fc5b84a.
//
// Solidity: event UBIFlagChanged(uint8 ubiFlag)
func (_Account *AccountFilterer) WatchUBIFlagChanged(opts *bind.WatchOpts, sink chan<- *AccountUBIFlagChanged) (event.Subscription, error) {

	logs, sub, err := _Account.contract.WatchLogs(opts, "UBIFlagChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountUBIFlagChanged)
				if err := _Account.contract.UnpackLog(event, "UBIFlagChanged", log); err != nil {
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

// ParseUBIFlagChanged is a log parse operation binding the contract event 0x2504aabda145c6cbc7a70eee13f3abeb510de8e65904c53b824d58ff9fc5b84a.
//
// Solidity: event UBIFlagChanged(uint8 ubiFlag)
func (_Account *AccountFilterer) ParseUBIFlagChanged(log types.Log) (*AccountUBIFlagChanged, error) {
	event := new(AccountUBIFlagChanged)
	if err := _Account.contract.UnpackLog(event, "UBIFlagChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountUBIProofSubmittedIterator is returned from FilterUBIProofSubmitted and is used to iterate over the raw logs and unpacked data for UBIProofSubmitted events raised by the Account contract.
type AccountUBIProofSubmittedIterator struct {
	Event *AccountUBIProofSubmitted // Event containing the contract specifics and raw log

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
func (it *AccountUBIProofSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountUBIProofSubmitted)
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
		it.Event = new(AccountUBIProofSubmitted)
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
func (it *AccountUBIProofSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountUBIProofSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountUBIProofSubmitted represents a UBIProofSubmitted event raised by the Account contract.
type AccountUBIProofSubmitted struct {
	Submitter common.Address
	TaskId    string
	TaskType  uint8
	ZkType    string
	Proof     string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUBIProofSubmitted is a free log retrieval operation binding the contract event 0x18a95158362c8c4a1cacd8dc12f81e1fd419696c958630a3c06e3466f4daad00.
//
// Solidity: event UBIProofSubmitted(address indexed submitter, string taskId, uint8 taskType, string zkType, string proof)
func (_Account *AccountFilterer) FilterUBIProofSubmitted(opts *bind.FilterOpts, submitter []common.Address) (*AccountUBIProofSubmittedIterator, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _Account.contract.FilterLogs(opts, "UBIProofSubmitted", submitterRule)
	if err != nil {
		return nil, err
	}
	return &AccountUBIProofSubmittedIterator{contract: _Account.contract, event: "UBIProofSubmitted", logs: logs, sub: sub}, nil
}

// WatchUBIProofSubmitted is a free log subscription operation binding the contract event 0x18a95158362c8c4a1cacd8dc12f81e1fd419696c958630a3c06e3466f4daad00.
//
// Solidity: event UBIProofSubmitted(address indexed submitter, string taskId, uint8 taskType, string zkType, string proof)
func (_Account *AccountFilterer) WatchUBIProofSubmitted(opts *bind.WatchOpts, sink chan<- *AccountUBIProofSubmitted, submitter []common.Address) (event.Subscription, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _Account.contract.WatchLogs(opts, "UBIProofSubmitted", submitterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountUBIProofSubmitted)
				if err := _Account.contract.UnpackLog(event, "UBIProofSubmitted", log); err != nil {
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

// ParseUBIProofSubmitted is a log parse operation binding the contract event 0x18a95158362c8c4a1cacd8dc12f81e1fd419696c958630a3c06e3466f4daad00.
//
// Solidity: event UBIProofSubmitted(address indexed submitter, string taskId, uint8 taskType, string zkType, string proof)
func (_Account *AccountFilterer) ParseUBIProofSubmitted(log types.Log) (*AccountUBIProofSubmitted, error) {
	event := new(AccountUBIProofSubmitted)
	if err := _Account.contract.UnpackLog(event, "UBIProofSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
