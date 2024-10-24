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

// SwanCreditCollateralCPInfo is an auto generated low-level Go binding around an user-defined struct.
type SwanCreditCollateralCPInfo struct {
	CpAccount        common.Address
	AvailableBalance *big.Int
	LockedCollateral *big.Int
	Status           string
}

// SwanCreditCollateralContractInfo is an auto generated low-level Go binding around an user-defined struct.
type SwanCreditCollateralContractInfo struct {
	SlashedFunds    *big.Int
	BaseCollateral  *big.Int
	CollateralRatio *big.Int
	SlashRatio      *big.Int
}

// SwanCreditCollateralTask is an auto generated low-level Go binding around an user-defined struct.
type SwanCreditCollateralTask struct {
	CpList           []common.Address
	Collateral       *big.Int
	CollateralStatus uint8
}

// SwanCreditCollateralMetaData contains all meta data concerning the SwanCreditCollateral contract.
var SwanCreditCollateralMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"frozenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"operation\",\"type\":\"string\"}],\"name\":\"CollateralAdjusted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"CollateralLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"CollateralSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumSwanCreditCollateral.CollateralStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"CollateralStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"CollateralUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fundingWallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"DisputeProof\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SlashedFundsIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cpAccountAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralContratOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashfund\",\"type\":\"uint256\"}],\"name\":\"WithdrawSlash\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedToWithdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"availableBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"slashedFunds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashRatio\",\"type\":\"uint256\"}],\"internalType\":\"structSwanCreditCollateral.ContractInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAddress\",\"type\":\"address\"}],\"name\":\"cpInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"availableBalance\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"lockedCollateral\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"}],\"internalType\":\"structSwanCreditCollateral.CPInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cpStatus\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"disputeProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"escrowBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"getTaskInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"cpList\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"enumSwanCreditCollateral.CollateralStatus\",\"name\":\"collateralStatus\",\"type\":\"uint8\"}],\"internalType\":\"structSwanCreditCollateral.Task\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"isSignatureUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockedBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockedCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseCollateral\",\"type\":\"uint256\"}],\"name\":\"setBaseCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"setCollateralToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slashRatio\",\"type\":\"uint256\"}],\"name\":\"setSlashRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slashAmount\",\"type\":\"uint256\"}],\"name\":\"slashCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slashAmount\",\"type\":\"uint256\"}],\"name\":\"slashCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashedFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"enumSwanCreditCollateral.CollateralStatus\",\"name\":\"collateralStatus\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"unlockRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPending\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawSlashedFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801561004357600080fd5b5061005261005760201b60201c565b6101c1565b600061006761015b60201b60201c565b90508060000160089054906101000a900460ff16156100b2576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff80168160000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff16146101585767ffffffffffffffff8160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d267ffffffffffffffff60405161014f91906101a6565b60405180910390a15b50565b60007ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b600067ffffffffffffffff82169050919050565b6101a081610183565b82525050565b60006020820190506101bb6000830184610197565b92915050565b608051613e806101ea600039600081816123f90152818161244e01526126090152613e806000f3fe60806040526004361061020f5760003560e01c80637048027511610118578063a664c216116100a0578063ce3518aa1161006f578063ce3518aa146107c9578063d27ca89b146107f2578063e80d9dec1461081d578063f2fde38b14610846578063f97348e11461086f5761020f565b8063a664c2161461070d578063ad3cb1cc1461074a578063b4eae1cb14610775578063b587b82c146107a05761020f565b80638da5cb5b116100e75780638da5cb5b1461060057806392bdf9ba1461062b5780639ae697bf146106685780639b5ddf09146106a5578063a0821be3146106d05761020f565b80637048027514610580578063715018a6146105a95780637f58a7e5146105c05780638129fc1c146105e95761020f565b80634f1ef2861161019b57806354fd4d501161016a57806354fd4d501461048657806355af6353146104b157806358709cf2146104ee578063666181a91461052c5780636f99f15c146105555761020f565b80634f1ef286146103e957806352d1902d14610405578063536f60701461043057806353ad87201461045b5761020f565b806324d7806c116101e257806324d7806c146102e05780632894493f1461031d5780632d291cad146103465780633fe651771461038357806347e7ef24146103c05761020f565b80631150f0f3146102145780631785f53c146102515780631b2094631461027a5780631d47a62d146102b7575b600080fd5b34801561022057600080fd5b5061023b60048036038101906102369190612dea565b6108ae565b6040516102489190612e4e565b60405180910390f35b34801561025d57600080fd5b5061027860048036038101906102739190612ec7565b6108e4565b005b34801561028657600080fd5b506102a1600480360381019061029c9190612f95565b610947565b6040516102ae919061317c565b60405180910390f35b3480156102c357600080fd5b506102de60048036038101906102d991906131ca565b610a52565b005b3480156102ec57600080fd5b5061030760048036038101906103029190612ec7565b61101d565b6040516103149190612e4e565b60405180910390f35b34801561032957600080fd5b50610344600480360381019061033f919061320a565b61103d565b005b34801561035257600080fd5b5061036d60048036038101906103689190612ec7565b611194565b60405161037a9190612e4e565b60405180910390f35b34801561038f57600080fd5b506103aa60048036038101906103a59190612ec7565b6111b4565b6040516103b791906132b6565b60405180910390f35b3480156103cc57600080fd5b506103e760048036038101906103e291906131ca565b611254565b005b61040360048036038101906103fe91906132d8565b6113bf565b005b34801561041157600080fd5b5061041a6113de565b604051610427919061334d565b60405180910390f35b34801561043c57600080fd5b50610445611411565b6040516104529190613377565b60405180910390f35b34801561046757600080fd5b50610470611417565b60405161047d91906133e7565b60405180910390f35b34801561049257600080fd5b5061049b61144d565b6040516104a89190613377565b60405180910390f35b3480156104bd57600080fd5b506104d860048036038101906104d39190612ec7565b611456565b6040516104e5919061341b565b60405180910390f35b3480156104fa57600080fd5b5061051560048036038101906105109190612f95565b61146e565b604051610523929190613445565b60405180910390f35b34801561053857600080fd5b50610553600480360381019061054e9190612ec7565b6114b5565b005b34801561056157600080fd5b5061056a611501565b6040516105779190613377565b60405180910390f35b34801561058c57600080fd5b506105a760048036038101906105a29190612ec7565b611507565b005b3480156105b557600080fd5b506105be61156a565b005b3480156105cc57600080fd5b506105e760048036038101906105e2919061320a565b61157e565b005b3480156105f557600080fd5b506105fe611614565b005b34801561060c57600080fd5b50610615611821565b604051610622919061347d565b60405180910390f35b34801561063757600080fd5b50610652600480360381019061064d9190612ec7565b611859565b60405161065f9190613377565b60405180910390f35b34801561067457600080fd5b5061068f600480360381019061068a9190612ec7565b611871565b60405161069c919061341b565b60405180910390f35b3480156106b157600080fd5b506106ba611889565b6040516106c79190613377565b60405180910390f35b3480156106dc57600080fd5b506106f760048036038101906106f29190612ec7565b61188f565b604051610704919061341b565b60405180910390f35b34801561071957600080fd5b50610734600480360381019061072f9190612ec7565b6118a7565b6040516107419190613554565b60405180910390f35b34801561075657600080fd5b5061075f611a32565b60405161076c91906132b6565b60405180910390f35b34801561078157600080fd5b5061078a611a6b565b6040516107979190613377565b60405180910390f35b3480156107ac57600080fd5b506107c760048036038101906107c29190613576565b611a71565b005b3480156107d557600080fd5b506107f060048036038101906107eb919061320a565b6120dd565b005b3480156107fe57600080fd5b506108076120ef565b6040516108149190613377565b60405180910390f35b34801561082957600080fd5b50610844600480360381019061083f9190612f95565b6120f5565b005b34801561085257600080fd5b5061086d60048036038101906108689190612ec7565b612151565b005b34801561087b57600080fd5b5061089660048036038101906108919190612ec7565b6121d7565b6040516108a5939291906135e5565b60405180910390f35b600b818051602081018201805184825260208301602085012081835280955050505050506000915054906101000a900460ff1681565b6108ec61220e565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b61094f612bf7565b60058260405161095f9190613658565b9081526020016040518091039020604051806060016040529081600082018054806020026020016040519081016040528092919081815260200182805480156109fd57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116109b3575b50505050508152602001600182015481526020016002820160009054906101000a900460ff166002811115610a3557610a346130b5565b5b6002811115610a4757610a466130b5565b5b815250509050919050565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610ade576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ad5906136e1565b60405180910390fd5b601060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548113610bce577fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b81604051610b539190613377565b60405180910390a180600080828254610b6c9190613730565b9250508190555080601060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610bc29190613764565b92505081905550610fc2565b60008190506000601060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541315610d0557601060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600080828254610c6b9190613730565b92505081905550601060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481610cbd9190613764565b90506000601060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b80600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205412610df55780600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610d9a9190613764565b9250508190555080600080828254610db29190613730565b925050819055507fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b82604051610de89190613377565b60405180910390a1610fc0565b6000600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541315610f2757600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481610e889190613764565b9050600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600080828254610eda9190613730565b925050819055506000600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b7fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b8183610f5491906137a7565b604051610f619190613377565b60405180910390a180601060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610fb89190613764565b925050819055505b505b610fcb82612295565b8173ffffffffffffffffffffffffffffffffffffffff167f403feb2cd85cc25c910fe59289105b583f08fe9e4335ebbb08c1851f47ff2d34826040516110119190613801565b60405180910390a25050565b60046020528060005260406000206000915054906101000a900460ff1681565b61104561220e565b80600054101561108a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611081906138a1565b60405180910390fd5b8060008082825461109b91906137a7565b92505081905550600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b81526004016110ff9291906138c1565b6020604051808303816000875af115801561111e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111429190613916565b503373ffffffffffffffffffffffffffffffffffffffff167fbfd9c82485e2178fcfd5c40379d8e66fe60acc50aa1ef4c50966431eb1e381cd826040516111899190613377565b60405180910390a250565b600a6020528060005260406000206000915054906101000a900460ff1681565b600860205280600052604060002060009150905080546111d390613972565b80601f01602080910402602001604051908101604052809291908181526020018280546111ff90613972565b801561124c5780601f106112215761010080835404028352916020019161124c565b820191906000526020600020905b81548152906001019060200180831161122f57829003601f168201915b505050505081565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b81526004016112b3939291906139a3565b6020604051808303816000875af11580156112d2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112f69190613916565b5080601060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461134691906139da565b925050819055508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62836040516113aa9190613377565b60405180910390a36113bb82612295565b5050565b6113c76123f7565b6113d0826124dd565b6113da82826124e8565b5050565b60006113e8612607565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b905090565b600e5481565b61141f612c2a565b6040518060800160405280600054815260200160015481526020016002548152602001600354815250905090565b60006002905090565b600c6020528060005260406000206000915090505481565b6005818051602081018201805184825260208301602085012081835280955050505050506000915090508060010154908060020160009054906101000a900460ff16905082565b6114bd61220e565b80600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60005481565b61150f61220e565b6001600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b61157261220e565b61157c600061268e565b565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661160a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611601906136e1565b60405180910390fd5b8060018190555050565b600061161e612765565b905060008160000160089054906101000a900460ff1615905060008260000160009054906101000a900467ffffffffffffffff1690506000808267ffffffffffffffff1614801561166c5750825b9050600060018367ffffffffffffffff161480156116a1575060003073ffffffffffffffffffffffffffffffffffffffff163b145b9050811580156116af575080155b156116e6576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018560000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555083156117365760018560000160086101000a81548160ff0219169083151502179055505b61173f3361278d565b6117476127a1565b6001600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600281905550600160038190555067016345785d8a0000600181905550831561181a5760008560000160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d260016040516118119190613a77565b60405180910390a15b5050505050565b60008061182c6127ab565b90508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505090565b60076020528060005260406000206000915090505481565b60106020528060005260406000206000915090505481565b60015481565b60066020528060005260406000206000915090505481565b6118af612c52565b60405180608001604052808373ffffffffffffffffffffffffffffffffffffffff168152602001600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548152602001600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548152602001600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080546119aa90613972565b80601f01602080910402602001604051908101604052809291908181526020018280546119d690613972565b8015611a235780601f106119f857610100808354040283529160200191611a23565b820191906000526020600020905b815481529060010190602001808311611a0657829003601f168201915b50505050508152509050919050565b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b60025481565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611afd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611af4906136e1565b60405180910390fd5b6000600584604051611b0f9190613658565b90815260200160405180910390209050601060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548213611c0f577fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b82604051611b949190613377565b60405180910390a181600080828254611bad9190613730565b9250508190555081601060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611c039190613764565b92505081905550612003565b60008290506000601060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541315611d4657601060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600080828254611cac9190613730565b92505081905550601060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481611cfe9190613764565b90506000601060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b80600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205412611e365780600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611ddb9190613764565b9250508190555080600080828254611df39190613730565b925050819055507fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b83604051611e299190613377565b60405180910390a1612001565b6000600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541315611f6857600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481611ec99190613764565b9050600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600080828254611f1b9190613730565b925050819055506000600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b7fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b8184611f9591906137a7565b604051611fa29190613377565b60405180910390a180601060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611ff99190613764565b925050819055505b505b60028160020160006101000a81548160ff0219169083600281111561202b5761202a6130b5565b5b021790555061203983612295565b8273ffffffffffffffffffffffffffffffffffffffff167f403feb2cd85cc25c910fe59289105b583f08fe9e4335ebbb08c1851f47ff2d348386604051612081929190613a92565b60405180910390a2836040516120979190613658565b60405180910390207f4a2ced9ada462e244851a86e998eb0b5bf558c2c9c6923b7f970ed2b19b073b060026040516120cf9190613ac2565b60405180910390a250505050565b6120e561220e565b8060038190555050565b60035481565b806040516121039190613658565b60405180910390203373ffffffffffffffffffffffffffffffffffffffff167faec1d412a3c1e4a13fc2a2e19ac38a5af192a9cf17b074fca8146a2d0655e0c360405160405180910390a350565b61215961220e565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036121cb5760006040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016121c2919061347d565b60405180910390fd5b6121d48161268e565b50565b600d6020528060005260406000206000915090508060000154908060010154908060020160009054906101000a900460ff16905083565b6122166127d3565b73ffffffffffffffffffffffffffffffffffffffff16612234611821565b73ffffffffffffffffffffffffffffffffffffffff1614612293576122576127d3565b6040517f118cdaa700000000000000000000000000000000000000000000000000000000815260040161228a919061347d565b60405180910390fd5b565b6001546002546122a59190613add565b601060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205412612371576040518060400160405280600981526020017f7a6b41756374696f6e0000000000000000000000000000000000000000000000815250600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020908161236b9190613cc1565b506123f4565b6040518060400160405280600381526020017f4e53430000000000000000000000000000000000000000000000000000000000815250600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090816123f29190613cc1565b505b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614806124a457507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661248b6127db565b73ffffffffffffffffffffffffffffffffffffffff1614155b156124db576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6124e561220e565b50565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561255057506040513d601f19601f8201168201806040525081019061254d9190613dbf565b60015b61259157816040517f4c9c8ce3000000000000000000000000000000000000000000000000000000008152600401612588919061347d565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b81146125f857806040517faa1d49a40000000000000000000000000000000000000000000000000000000081526004016125ef919061334d565b60405180910390fd5b6126028383612832565b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161461268c576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60006126986127ab565b905060008160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050828260000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3505050565b60007ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b6127956128a5565b61279e816128e5565b50565b6127a96128a5565b565b60007f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300905090565b600033905090565b60006128097f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b61296b565b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61283b82612975565b8173ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a2600081511115612898576128928282612a42565b506128a1565b6128a0612ac6565b5b5050565b6128ad612b03565b6128e3576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6128ed6128a5565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361295f5760006040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401612956919061347d565b60405180910390fd5b6129688161268e565b50565b6000819050919050565b60008173ffffffffffffffffffffffffffffffffffffffff163b036129d157806040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526004016129c8919061347d565b60405180910390fd5b806129fe7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b61296b565b60000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60606000808473ffffffffffffffffffffffffffffffffffffffff1684604051612a6c9190613e33565b600060405180830381855af49150503d8060008114612aa7576040519150601f19603f3d011682016040523d82523d6000602084013e612aac565b606091505b5091509150612abc858383612b23565b9250505092915050565b6000341115612b01576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6000612b0d612765565b60000160089054906101000a900460ff16905090565b606082612b3857612b3382612bb2565b612baa565b60008251148015612b60575060008473ffffffffffffffffffffffffffffffffffffffff163b145b15612ba257836040517f9996b315000000000000000000000000000000000000000000000000000000008152600401612b99919061347d565b60405180910390fd5b819050612bab565b5b9392505050565b600081511115612bc55780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040518060600160405280606081526020016000815260200160006002811115612c2457612c236130b5565b5b81525090565b6040518060800160405280600081526020016000815260200160008152602001600081525090565b6040518060800160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001606081525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b612cf782612cae565b810181811067ffffffffffffffff82111715612d1657612d15612cbf565b5b80604052505050565b6000612d29612c90565b9050612d358282612cee565b919050565b600067ffffffffffffffff821115612d5557612d54612cbf565b5b612d5e82612cae565b9050602081019050919050565b82818337600083830152505050565b6000612d8d612d8884612d3a565b612d1f565b905082815260208101848484011115612da957612da8612ca9565b5b612db4848285612d6b565b509392505050565b600082601f830112612dd157612dd0612ca4565b5b8135612de1848260208601612d7a565b91505092915050565b600060208284031215612e0057612dff612c9a565b5b600082013567ffffffffffffffff811115612e1e57612e1d612c9f565b5b612e2a84828501612dbc565b91505092915050565b60008115159050919050565b612e4881612e33565b82525050565b6000602082019050612e636000830184612e3f565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000612e9482612e69565b9050919050565b612ea481612e89565b8114612eaf57600080fd5b50565b600081359050612ec181612e9b565b92915050565b600060208284031215612edd57612edc612c9a565b5b6000612eeb84828501612eb2565b91505092915050565b600067ffffffffffffffff821115612f0f57612f0e612cbf565b5b612f1882612cae565b9050602081019050919050565b6000612f38612f3384612ef4565b612d1f565b905082815260208101848484011115612f5457612f53612ca9565b5b612f5f848285612d6b565b509392505050565b600082601f830112612f7c57612f7b612ca4565b5b8135612f8c848260208601612f25565b91505092915050565b600060208284031215612fab57612faa612c9a565b5b600082013567ffffffffffffffff811115612fc957612fc8612c9f565b5b612fd584828501612f67565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61301381612e89565b82525050565b6000613025838361300a565b60208301905092915050565b6000602082019050919050565b600061304982612fde565b6130538185612fe9565b935061305e83612ffa565b8060005b8381101561308f5781516130768882613019565b975061308183613031565b925050600181019050613062565b5085935050505092915050565b6000819050919050565b6130af8161309c565b82525050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600381106130f5576130f46130b5565b5b50565b6000819050613106826130e4565b919050565b6000613116826130f8565b9050919050565b6131268161310b565b82525050565b60006060830160008301518482036000860152613149828261303e565b915050602083015161315e60208601826130a6565b506040830151613171604086018261311d565b508091505092915050565b60006020820190508181036000830152613196818461312c565b905092915050565b6131a78161309c565b81146131b257600080fd5b50565b6000813590506131c48161319e565b92915050565b600080604083850312156131e1576131e0612c9a565b5b60006131ef85828601612eb2565b9250506020613200858286016131b5565b9150509250929050565b6000602082840312156132205761321f612c9a565b5b600061322e848285016131b5565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015613271578082015181840152602081019050613256565b60008484015250505050565b600061328882613237565b6132928185613242565b93506132a2818560208601613253565b6132ab81612cae565b840191505092915050565b600060208201905081810360008301526132d0818461327d565b905092915050565b600080604083850312156132ef576132ee612c9a565b5b60006132fd85828601612eb2565b925050602083013567ffffffffffffffff81111561331e5761331d612c9f565b5b61332a85828601612dbc565b9150509250929050565b6000819050919050565b61334781613334565b82525050565b6000602082019050613362600083018461333e565b92915050565b6133718161309c565b82525050565b600060208201905061338c6000830184613368565b92915050565b6080820160008201516133a860008501826130a6565b5060208201516133bb60208501826130a6565b5060408201516133ce60408501826130a6565b5060608201516133e160608501826130a6565b50505050565b60006080820190506133fc6000830184613392565b92915050565b6000819050919050565b61341581613402565b82525050565b6000602082019050613430600083018461340c565b92915050565b61343f8161310b565b82525050565b600060408201905061345a6000830185613368565b6134676020830184613436565b9392505050565b61347781612e89565b82525050565b6000602082019050613492600083018461346e565b92915050565b6134a181613402565b82525050565b600082825260208201905092915050565b60006134c382613237565b6134cd81856134a7565b93506134dd818560208601613253565b6134e681612cae565b840191505092915050565b6000608083016000830151613509600086018261300a565b50602083015161351c6020860182613498565b50604083015161352f60408601826130a6565b506060830151848203606086015261354782826134b8565b9150508091505092915050565b6000602082019050818103600083015261356e81846134f1565b905092915050565b60008060006060848603121561358f5761358e612c9a565b5b600084013567ffffffffffffffff8111156135ad576135ac612c9f565b5b6135b986828701612f67565b93505060206135ca86828701612eb2565b92505060406135db868287016131b5565b9150509250925092565b60006060820190506135fa6000830186613368565b6136076020830185613368565b6136146040830184612e3f565b949350505050565b600081905092915050565b600061363282613237565b61363c818561361c565b935061364c818560208601613253565b80840191505092915050565b60006136648284613627565b915081905092915050565b7f4f6e6c79207468652061646d696e2063616e2063616c6c20746869732066756e60008201527f6374696f6e2e0000000000000000000000000000000000000000000000000000602082015250565b60006136cb602683613242565b91506136d68261366f565b604082019050919050565b600060208201905081810360008301526136fa816136be565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061373b8261309c565b91506137468361309c565b925082820190508082111561375e5761375d613701565b5b92915050565b600061376f82613402565b915061377a83613402565b92508282039050818112600084121682821360008512151617156137a1576137a0613701565b5b92915050565b60006137b28261309c565b91506137bd8361309c565b92508282039050818111156137d5576137d4613701565b5b92915050565b50565b60006137eb600083613242565b91506137f6826137db565b600082019050919050565b60006040820190506138166000830184613368565b8181036020830152613827816137de565b905092915050565b7f576974686472617720616d6f756e7420616d6f756e742065786365656473207360008201527f6c617368656446756e6473000000000000000000000000000000000000000000602082015250565b600061388b602b83613242565b91506138968261382f565b604082019050919050565b600060208201905081810360008301526138ba8161387e565b9050919050565b60006040820190506138d6600083018561346e565b6138e36020830184613368565b9392505050565b6138f381612e33565b81146138fe57600080fd5b50565b600081519050613910816138ea565b92915050565b60006020828403121561392c5761392b612c9a565b5b600061393a84828501613901565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061398a57607f821691505b60208210810361399d5761399c613943565b5b50919050565b60006060820190506139b8600083018661346e565b6139c5602083018561346e565b6139d26040830184613368565b949350505050565b60006139e582613402565b91506139f083613402565b925082820190508281121560008312168382126000841215161715613a1857613a17613701565b5b92915050565b6000819050919050565b600067ffffffffffffffff82169050919050565b6000819050919050565b6000613a61613a5c613a5784613a1e565b613a3c565b613a28565b9050919050565b613a7181613a46565b82525050565b6000602082019050613a8c6000830184613a68565b92915050565b6000604082019050613aa76000830185613368565b8181036020830152613ab9818461327d565b90509392505050565b6000602082019050613ad76000830184613436565b92915050565b6000613ae88261309c565b9150613af38361309c565b9250828202613b018161309c565b91508282048414831517613b1857613b17613701565b5b5092915050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302613b817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82613b44565b613b8b8683613b44565b95508019841693508086168417925050509392505050565b6000613bbe613bb9613bb48461309c565b613a3c565b61309c565b9050919050565b6000819050919050565b613bd883613ba3565b613bec613be482613bc5565b848454613b51565b825550505050565b600090565b613c01613bf4565b613c0c818484613bcf565b505050565b5b81811015613c3057613c25600082613bf9565b600181019050613c12565b5050565b601f821115613c7557613c4681613b1f565b613c4f84613b34565b81016020851015613c5e578190505b613c72613c6a85613b34565b830182613c11565b50505b505050565b600082821c905092915050565b6000613c9860001984600802613c7a565b1980831691505092915050565b6000613cb18383613c87565b9150826002028217905092915050565b613cca82613237565b67ffffffffffffffff811115613ce357613ce2612cbf565b5b613ced8254613972565b613cf8828285613c34565b600060209050601f831160018114613d2b5760008415613d19578287015190505b613d238582613ca5565b865550613d8b565b601f198416613d3986613b1f565b60005b82811015613d6157848901518255600182019150602085019450602081019050613d3c565b86831015613d7e5784890151613d7a601f891682613c87565b8355505b6001600288020188555050505b505050505050565b613d9c81613334565b8114613da757600080fd5b50565b600081519050613db981613d93565b92915050565b600060208284031215613dd557613dd4612c9a565b5b6000613de384828501613daa565b91505092915050565b600081519050919050565b600081905092915050565b6000613e0d82613dec565b613e178185613df7565b9350613e27818560208601613253565b80840191505092915050565b6000613e3f8284613e02565b91508190509291505056fea26469706673582212208dfa31df8d36f9fd2f05861942855db2e6bdbde3f5d010092711c027f985122764736f6c63430008190033",
}

// SwanCreditCollateralABI is the input ABI used to generate the binding from.
// Deprecated: Use SwanCreditCollateralMetaData.ABI instead.
var SwanCreditCollateralABI = SwanCreditCollateralMetaData.ABI

// SwanCreditCollateralBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SwanCreditCollateralMetaData.Bin instead.
var SwanCreditCollateralBin = SwanCreditCollateralMetaData.Bin

// DeploySwanCreditCollateral deploys a new Ethereum contract, binding an instance of SwanCreditCollateral to it.
func DeploySwanCreditCollateral(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SwanCreditCollateral, error) {
	parsed, err := SwanCreditCollateralMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SwanCreditCollateralBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SwanCreditCollateral{SwanCreditCollateralCaller: SwanCreditCollateralCaller{contract: contract}, SwanCreditCollateralTransactor: SwanCreditCollateralTransactor{contract: contract}, SwanCreditCollateralFilterer: SwanCreditCollateralFilterer{contract: contract}}, nil
}

// SwanCreditCollateral is an auto generated Go binding around an Ethereum contract.
type SwanCreditCollateral struct {
	SwanCreditCollateralCaller     // Read-only binding to the contract
	SwanCreditCollateralTransactor // Write-only binding to the contract
	SwanCreditCollateralFilterer   // Log filterer for contract events
}

// SwanCreditCollateralCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwanCreditCollateralCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwanCreditCollateralTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwanCreditCollateralTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwanCreditCollateralFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwanCreditCollateralFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwanCreditCollateralSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwanCreditCollateralSession struct {
	Contract     *SwanCreditCollateral // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SwanCreditCollateralCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwanCreditCollateralCallerSession struct {
	Contract *SwanCreditCollateralCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// SwanCreditCollateralTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwanCreditCollateralTransactorSession struct {
	Contract     *SwanCreditCollateralTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// SwanCreditCollateralRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwanCreditCollateralRaw struct {
	Contract *SwanCreditCollateral // Generic contract binding to access the raw methods on
}

// SwanCreditCollateralCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwanCreditCollateralCallerRaw struct {
	Contract *SwanCreditCollateralCaller // Generic read-only contract binding to access the raw methods on
}

// SwanCreditCollateralTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwanCreditCollateralTransactorRaw struct {
	Contract *SwanCreditCollateralTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwanCreditCollateral creates a new instance of SwanCreditCollateral, bound to a specific deployed contract.
func NewSwanCreditCollateral(address common.Address, backend bind.ContractBackend) (*SwanCreditCollateral, error) {
	contract, err := bindSwanCreditCollateral(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwanCreditCollateral{SwanCreditCollateralCaller: SwanCreditCollateralCaller{contract: contract}, SwanCreditCollateralTransactor: SwanCreditCollateralTransactor{contract: contract}, SwanCreditCollateralFilterer: SwanCreditCollateralFilterer{contract: contract}}, nil
}

// NewSwanCreditCollateralCaller creates a new read-only instance of SwanCreditCollateral, bound to a specific deployed contract.
func NewSwanCreditCollateralCaller(address common.Address, caller bind.ContractCaller) (*SwanCreditCollateralCaller, error) {
	contract, err := bindSwanCreditCollateral(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwanCreditCollateralCaller{contract: contract}, nil
}

// NewSwanCreditCollateralTransactor creates a new write-only instance of SwanCreditCollateral, bound to a specific deployed contract.
func NewSwanCreditCollateralTransactor(address common.Address, transactor bind.ContractTransactor) (*SwanCreditCollateralTransactor, error) {
	contract, err := bindSwanCreditCollateral(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwanCreditCollateralTransactor{contract: contract}, nil
}

// NewSwanCreditCollateralFilterer creates a new log filterer instance of SwanCreditCollateral, bound to a specific deployed contract.
func NewSwanCreditCollateralFilterer(address common.Address, filterer bind.ContractFilterer) (*SwanCreditCollateralFilterer, error) {
	contract, err := bindSwanCreditCollateral(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwanCreditCollateralFilterer{contract: contract}, nil
}

// bindSwanCreditCollateral binds a generic wrapper to an already deployed contract.
func bindSwanCreditCollateral(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SwanCreditCollateralMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwanCreditCollateral *SwanCreditCollateralRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwanCreditCollateral.Contract.SwanCreditCollateralCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwanCreditCollateral *SwanCreditCollateralRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.SwanCreditCollateralTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwanCreditCollateral *SwanCreditCollateralRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.SwanCreditCollateralTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwanCreditCollateral *SwanCreditCollateralCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwanCreditCollateral.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwanCreditCollateral *SwanCreditCollateralTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwanCreditCollateral *SwanCreditCollateralTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_SwanCreditCollateral *SwanCreditCollateralCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SwanCreditCollateral.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_SwanCreditCollateral *SwanCreditCollateralSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _SwanCreditCollateral.Contract.UPGRADEINTERFACEVERSION(&_SwanCreditCollateral.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_SwanCreditCollateral *SwanCreditCollateralCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _SwanCreditCollateral.Contract.UPGRADEINTERFACEVERSION(&_SwanCreditCollateral.CallOpts)
}

// AllowedToWithdraw is a free data retrieval call binding the contract method 0x2d291cad.
//
// Solidity: function allowedToWithdraw(address ) view returns(bool)
func (_SwanCreditCollateral *SwanCreditCollateralCaller) AllowedToWithdraw(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _SwanCreditCollateral.contract.Call(opts, &out, "allowedToWithdraw", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedToWithdraw is a free data retrieval call binding the contract method 0x2d291cad.
//
// Solidity: function allowedToWithdraw(address ) view returns(bool)
func (_SwanCreditCollateral *SwanCreditCollateralSession) AllowedToWithdraw(arg0 common.Address) (bool, error) {
	return _SwanCreditCollateral.Contract.AllowedToWithdraw(&_SwanCreditCollateral.CallOpts, arg0)
}

// AllowedToWithdraw is a free data retrieval call binding the contract method 0x2d291cad.
//
// Solidity: function allowedToWithdraw(address ) view returns(bool)
func (_SwanCreditCollateral *SwanCreditCollateralCallerSession) AllowedToWithdraw(arg0 common.Address) (bool, error) {
	return _SwanCreditCollateral.Contract.AllowedToWithdraw(&_SwanCreditCollateral.CallOpts, arg0)
}

// AvailableBalance is a free data retrieval call binding the contract method 0xa0821be3.
//
// Solidity: function availableBalance(address ) view returns(int256)
func (_SwanCreditCollateral *SwanCreditCollateralCaller) AvailableBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SwanCreditCollateral.contract.Call(opts, &out, "availableBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableBalance is a free data retrieval call binding the contract method 0xa0821be3.
//
// Solidity: function availableBalance(address ) view returns(int256)
func (_SwanCreditCollateral *SwanCreditCollateralSession) AvailableBalance(arg0 common.Address) (*big.Int, error) {
	return _SwanCreditCollateral.Contract.AvailableBalance(&_SwanCreditCollateral.CallOpts, arg0)
}

// AvailableBalance is a free data retrieval call binding the contract method 0xa0821be3.
//
// Solidity: function availableBalance(address ) view returns(int256)
func (_SwanCreditCollateral *SwanCreditCollateralCallerSession) AvailableBalance(arg0 common.Address) (*big.Int, error) {
	return _SwanCreditCollateral.Contract.AvailableBalance(&_SwanCreditCollateral.CallOpts, arg0)
}

// BaseCollateral is a free data retrieval call binding the contract method 0x9b5ddf09.
//
// Solidity: function baseCollateral() view returns(uint256)
func (_SwanCreditCollateral *SwanCreditCollateralCaller) BaseCollateral(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwanCreditCollateral.contract.Call(opts, &out, "baseCollateral")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseCollateral is a free data retrieval call binding the contract method 0x9b5ddf09.
//
// Solidity: function baseCollateral() view returns(uint256)
func (_SwanCreditCollateral *SwanCreditCollateralSession) BaseCollateral() (*big.Int, error) {
	return _SwanCreditCollateral.Contract.BaseCollateral(&_SwanCreditCollateral.CallOpts)
}

// BaseCollateral is a free data retrieval call binding the contract method 0x9b5ddf09.
//
// Solidity: function baseCollateral() view returns(uint256)
func (_SwanCreditCollateral *SwanCreditCollateralCallerSession) BaseCollateral() (*big.Int, error) {
	return _SwanCreditCollateral.Contract.BaseCollateral(&_SwanCreditCollateral.CallOpts)
}

// CollateralInfo is a free data retrieval call binding the contract method 0x53ad8720.
//
// Solidity: function collateralInfo() view returns((uint256,uint256,uint256,uint256))
func (_SwanCreditCollateral *SwanCreditCollateralCaller) CollateralInfo(opts *bind.CallOpts) (SwanCreditCollateralContractInfo, error) {
	var out []interface{}
	err := _SwanCreditCollateral.contract.Call(opts, &out, "collateralInfo")

	if err != nil {
		return *new(SwanCreditCollateralContractInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(SwanCreditCollateralContractInfo)).(*SwanCreditCollateralContractInfo)

	return out0, err

}

// CollateralInfo is a free data retrieval call binding the contract method 0x53ad8720.
//
// Solidity: function collateralInfo() view returns((uint256,uint256,uint256,uint256))
func (_SwanCreditCollateral *SwanCreditCollateralSession) CollateralInfo() (SwanCreditCollateralContractInfo, error) {
	return _SwanCreditCollateral.Contract.CollateralInfo(&_SwanCreditCollateral.CallOpts)
}

// CollateralInfo is a free data retrieval call binding the contract method 0x53ad8720.
//
// Solidity: function collateralInfo() view returns((uint256,uint256,uint256,uint256))
func (_SwanCreditCollateral *SwanCreditCollateralCallerSession) CollateralInfo() (SwanCreditCollateralContractInfo, error) {
	return _SwanCreditCollateral.Contract.CollateralInfo(&_SwanCreditCollateral.CallOpts)
}

// CollateralRatio is a free data retrieval call binding the contract method 0xb4eae1cb.
//
// Solidity: function collateralRatio() view returns(uint256)
func (_SwanCreditCollateral *SwanCreditCollateralCaller) CollateralRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwanCreditCollateral.contract.Call(opts, &out, "collateralRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollateralRatio is a free data retrieval call binding the contract method 0xb4eae1cb.
//
// Solidity: function collateralRatio() view returns(uint256)
func (_SwanCreditCollateral *SwanCreditCollateralSession) CollateralRatio() (*big.Int, error) {
	return _SwanCreditCollateral.Contract.CollateralRatio(&_SwanCreditCollateral.CallOpts)
}

// CollateralRatio is a free data retrieval call binding the contract method 0xb4eae1cb.
//
// Solidity: function collateralRatio() view returns(uint256)
func (_SwanCreditCollateral *SwanCreditCollateralCallerSession) CollateralRatio() (*big.Int, error) {
	return _SwanCreditCollateral.Contract.CollateralRatio(&_SwanCreditCollateral.CallOpts)
}

// CpInfo is a free data retrieval call binding the contract method 0xa664c216.
//
// Solidity: function cpInfo(address cpAddress) view returns((address,int256,uint256,string))
func (_SwanCreditCollateral *SwanCreditCollateralCaller) CpInfo(opts *bind.CallOpts, cpAddress common.Address) (SwanCreditCollateralCPInfo, error) {
	var out []interface{}
	err := _SwanCreditCollateral.contract.Call(opts, &out, "cpInfo", cpAddress)

	if err != nil {
		return *new(SwanCreditCollateralCPInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(SwanCreditCollateralCPInfo)).(*SwanCreditCollateralCPInfo)

	return out0, err

}

// CpInfo is a free data retrieval call binding the contract method 0xa664c216.
//
// Solidity: function cpInfo(address cpAddress) view returns((address,int256,uint256,string))
func (_SwanCreditCollateral *SwanCreditCollateralSession) CpInfo(cpAddress common.Address) (SwanCreditCollateralCPInfo, error) {
	return _SwanCreditCollateral.Contract.CpInfo(&_SwanCreditCollateral.CallOpts, cpAddress)
}

// CpInfo is a free data retrieval call binding the contract method 0xa664c216.
//
// Solidity: function cpInfo(address cpAddress) view returns((address,int256,uint256,string))
func (_SwanCreditCollateral *SwanCreditCollateralCallerSession) CpInfo(cpAddress common.Address) (SwanCreditCollateralCPInfo, error) {
	return _SwanCreditCollateral.Contract.CpInfo(&_SwanCreditCollateral.CallOpts, cpAddress)
}

// CpStatus is a free data retrieval call binding the contract method 0x3fe65177.
//
// Solidity: function cpStatus(address ) view returns(string)
func (_SwanCreditCollateral *SwanCreditCollateralCaller) CpStatus(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _SwanCreditCollateral.contract.Call(opts, &out, "cpStatus", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CpStatus is a free data retrieval call binding the contract method 0x3fe65177.
//
// Solidity: function cpStatus(address ) view returns(string)
func (_SwanCreditCollateral *SwanCreditCollateralSession) CpStatus(arg0 common.Address) (string, error) {
	return _SwanCreditCollateral.Contract.CpStatus(&_SwanCreditCollateral.CallOpts, arg0)
}

// CpStatus is a free data retrieval call binding the contract method 0x3fe65177.
//
// Solidity: function cpStatus(address ) view returns(string)
func (_SwanCreditCollateral *SwanCreditCollateralCallerSession) CpStatus(arg0 common.Address) (string, error) {
	return _SwanCreditCollateral.Contract.CpStatus(&_SwanCreditCollateral.CallOpts, arg0)
}

// EscrowBalance is a free data retrieval call binding the contract method 0x55af6353.
//
// Solidity: function escrowBalance(address ) view returns(int256)
func (_SwanCreditCollateral *SwanCreditCollateralCaller) EscrowBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SwanCreditCollateral.contract.Call(opts, &out, "escrowBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EscrowBalance is a free data retrieval call binding the contract method 0x55af6353.
//
// Solidity: function escrowBalance(address ) view returns(int256)
func (_SwanCreditCollateral *SwanCreditCollateralSession) EscrowBalance(arg0 common.Address) (*big.Int, error) {
	return _SwanCreditCollateral.Contract.EscrowBalance(&_SwanCreditCollateral.CallOpts, arg0)
}

// EscrowBalance is a free data retrieval call binding the contract method 0x55af6353.
//
// Solidity: function escrowBalance(address ) view returns(int256)
func (_SwanCreditCollateral *SwanCreditCollateralCallerSession) EscrowBalance(arg0 common.Address) (*big.Int, error) {
	return _SwanCreditCollateral.Contract.EscrowBalance(&_SwanCreditCollateral.CallOpts, arg0)
}

// GetTaskInfo is a free data retrieval call binding the contract method 0x1b209463.
//
// Solidity: function getTaskInfo(string taskUid) view returns((address[],uint256,uint8))
func (_SwanCreditCollateral *SwanCreditCollateralCaller) GetTaskInfo(opts *bind.CallOpts, taskUid string) (SwanCreditCollateralTask, error) {
	var out []interface{}
	err := _SwanCreditCollateral.contract.Call(opts, &out, "getTaskInfo", taskUid)

	if err != nil {
		return *new(SwanCreditCollateralTask), err
	}

	out0 := *abi.ConvertType(out[0], new(SwanCreditCollateralTask)).(*SwanCreditCollateralTask)

	return out0, err

}

// GetTaskInfo is a free data retrieval call binding the contract method 0x1b209463.
//
// Solidity: function getTaskInfo(string taskUid) view returns((address[],uint256,uint8))
func (_SwanCreditCollateral *SwanCreditCollateralSession) GetTaskInfo(taskUid string) (SwanCreditCollateralTask, error) {
	return _SwanCreditCollateral.Contract.GetTaskInfo(&_SwanCreditCollateral.CallOpts, taskUid)
}

// GetTaskInfo is a free data retrieval call binding the contract method 0x1b209463.
//
// Solidity: function getTaskInfo(string taskUid) view returns((address[],uint256,uint8))
func (_SwanCreditCollateral *SwanCreditCollateralCallerSession) GetTaskInfo(taskUid string) (SwanCreditCollateralTask, error) {
	return _SwanCreditCollateral.Contract.GetTaskInfo(&_SwanCreditCollateral.CallOpts, taskUid)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_SwanCreditCollateral *SwanCreditCollateralCaller) IsAdmin(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _SwanCreditCollateral.contract.Call(opts, &out, "isAdmin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_SwanCreditCollateral *SwanCreditCollateralSession) IsAdmin(arg0 common.Address) (bool, error) {
	return _SwanCreditCollateral.Contract.IsAdmin(&_SwanCreditCollateral.CallOpts, arg0)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_SwanCreditCollateral *SwanCreditCollateralCallerSession) IsAdmin(arg0 common.Address) (bool, error) {
	return _SwanCreditCollateral.Contract.IsAdmin(&_SwanCreditCollateral.CallOpts, arg0)
}

// IsSignatureUsed is a free data retrieval call binding the contract method 0x1150f0f3.
//
// Solidity: function isSignatureUsed(bytes ) view returns(bool)
func (_SwanCreditCollateral *SwanCreditCollateralCaller) IsSignatureUsed(opts *bind.CallOpts, arg0 []byte) (bool, error) {
	var out []interface{}
	err := _SwanCreditCollateral.contract.Call(opts, &out, "isSignatureUsed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSignatureUsed is a free data retrieval call binding the contract method 0x1150f0f3.
//
// Solidity: function isSignatureUsed(bytes ) view returns(bool)
func (_SwanCreditCollateral *SwanCreditCollateralSession) IsSignatureUsed(arg0 []byte) (bool, error) {
	return _SwanCreditCollateral.Contract.IsSignatureUsed(&_SwanCreditCollateral.CallOpts, arg0)
}

// IsSignatureUsed is a free data retrieval call binding the contract method 0x1150f0f3.
//
// Solidity: function isSignatureUsed(bytes ) view returns(bool)
func (_SwanCreditCollateral *SwanCreditCollateralCallerSession) IsSignatureUsed(arg0 []byte) (bool, error) {
	return _SwanCreditCollateral.Contract.IsSignatureUsed(&_SwanCreditCollateral.CallOpts, arg0)
}

// LockedBalance is a free data retrieval call binding the contract method 0x9ae697bf.
//
// Solidity: function lockedBalance(address ) view returns(int256)
func (_SwanCreditCollateral *SwanCreditCollateralCaller) LockedBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SwanCreditCollateral.contract.Call(opts, &out, "lockedBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedBalance is a free data retrieval call binding the contract method 0x9ae697bf.
//
// Solidity: function lockedBalance(address ) view returns(int256)
func (_SwanCreditCollateral *SwanCreditCollateralSession) LockedBalance(arg0 common.Address) (*big.Int, error) {
	return _SwanCreditCollateral.Contract.LockedBalance(&_SwanCreditCollateral.CallOpts, arg0)
}

// LockedBalance is a free data retrieval call binding the contract method 0x9ae697bf.
//
// Solidity: function lockedBalance(address ) view returns(int256)
func (_SwanCreditCollateral *SwanCreditCollateralCallerSession) LockedBalance(arg0 common.Address) (*big.Int, error) {
	return _SwanCreditCollateral.Contract.LockedBalance(&_SwanCreditCollateral.CallOpts, arg0)
}

// LockedCollateral is a free data retrieval call binding the contract method 0x92bdf9ba.
//
// Solidity: function lockedCollateral(address ) view returns(uint256)
func (_SwanCreditCollateral *SwanCreditCollateralCaller) LockedCollateral(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SwanCreditCollateral.contract.Call(opts, &out, "lockedCollateral", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedCollateral is a free data retrieval call binding the contract method 0x92bdf9ba.
//
// Solidity: function lockedCollateral(address ) view returns(uint256)
func (_SwanCreditCollateral *SwanCreditCollateralSession) LockedCollateral(arg0 common.Address) (*big.Int, error) {
	return _SwanCreditCollateral.Contract.LockedCollateral(&_SwanCreditCollateral.CallOpts, arg0)
}

// LockedCollateral is a free data retrieval call binding the contract method 0x92bdf9ba.
//
// Solidity: function lockedCollateral(address ) view returns(uint256)
func (_SwanCreditCollateral *SwanCreditCollateralCallerSession) LockedCollateral(arg0 common.Address) (*big.Int, error) {
	return _SwanCreditCollateral.Contract.LockedCollateral(&_SwanCreditCollateral.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwanCreditCollateral *SwanCreditCollateralCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwanCreditCollateral.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwanCreditCollateral *SwanCreditCollateralSession) Owner() (common.Address, error) {
	return _SwanCreditCollateral.Contract.Owner(&_SwanCreditCollateral.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwanCreditCollateral *SwanCreditCollateralCallerSession) Owner() (common.Address, error) {
	return _SwanCreditCollateral.Contract.Owner(&_SwanCreditCollateral.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SwanCreditCollateral *SwanCreditCollateralCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SwanCreditCollateral.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SwanCreditCollateral *SwanCreditCollateralSession) ProxiableUUID() ([32]byte, error) {
	return _SwanCreditCollateral.Contract.ProxiableUUID(&_SwanCreditCollateral.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SwanCreditCollateral *SwanCreditCollateralCallerSession) ProxiableUUID() ([32]byte, error) {
	return _SwanCreditCollateral.Contract.ProxiableUUID(&_SwanCreditCollateral.CallOpts)
}

// RequestInterval is a free data retrieval call binding the contract method 0x536f6070.
//
// Solidity: function requestInterval() view returns(uint256)
func (_SwanCreditCollateral *SwanCreditCollateralCaller) RequestInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwanCreditCollateral.contract.Call(opts, &out, "requestInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestInterval is a free data retrieval call binding the contract method 0x536f6070.
//
// Solidity: function requestInterval() view returns(uint256)
func (_SwanCreditCollateral *SwanCreditCollateralSession) RequestInterval() (*big.Int, error) {
	return _SwanCreditCollateral.Contract.RequestInterval(&_SwanCreditCollateral.CallOpts)
}

// RequestInterval is a free data retrieval call binding the contract method 0x536f6070.
//
// Solidity: function requestInterval() view returns(uint256)
func (_SwanCreditCollateral *SwanCreditCollateralCallerSession) RequestInterval() (*big.Int, error) {
	return _SwanCreditCollateral.Contract.RequestInterval(&_SwanCreditCollateral.CallOpts)
}

// SlashRatio is a free data retrieval call binding the contract method 0xd27ca89b.
//
// Solidity: function slashRatio() view returns(uint256)
func (_SwanCreditCollateral *SwanCreditCollateralCaller) SlashRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwanCreditCollateral.contract.Call(opts, &out, "slashRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashRatio is a free data retrieval call binding the contract method 0xd27ca89b.
//
// Solidity: function slashRatio() view returns(uint256)
func (_SwanCreditCollateral *SwanCreditCollateralSession) SlashRatio() (*big.Int, error) {
	return _SwanCreditCollateral.Contract.SlashRatio(&_SwanCreditCollateral.CallOpts)
}

// SlashRatio is a free data retrieval call binding the contract method 0xd27ca89b.
//
// Solidity: function slashRatio() view returns(uint256)
func (_SwanCreditCollateral *SwanCreditCollateralCallerSession) SlashRatio() (*big.Int, error) {
	return _SwanCreditCollateral.Contract.SlashRatio(&_SwanCreditCollateral.CallOpts)
}

// SlashedFunds is a free data retrieval call binding the contract method 0x6f99f15c.
//
// Solidity: function slashedFunds() view returns(uint256)
func (_SwanCreditCollateral *SwanCreditCollateralCaller) SlashedFunds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwanCreditCollateral.contract.Call(opts, &out, "slashedFunds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashedFunds is a free data retrieval call binding the contract method 0x6f99f15c.
//
// Solidity: function slashedFunds() view returns(uint256)
func (_SwanCreditCollateral *SwanCreditCollateralSession) SlashedFunds() (*big.Int, error) {
	return _SwanCreditCollateral.Contract.SlashedFunds(&_SwanCreditCollateral.CallOpts)
}

// SlashedFunds is a free data retrieval call binding the contract method 0x6f99f15c.
//
// Solidity: function slashedFunds() view returns(uint256)
func (_SwanCreditCollateral *SwanCreditCollateralCallerSession) SlashedFunds() (*big.Int, error) {
	return _SwanCreditCollateral.Contract.SlashedFunds(&_SwanCreditCollateral.CallOpts)
}

// Tasks is a free data retrieval call binding the contract method 0x58709cf2.
//
// Solidity: function tasks(string ) view returns(uint256 collateral, uint8 collateralStatus)
func (_SwanCreditCollateral *SwanCreditCollateralCaller) Tasks(opts *bind.CallOpts, arg0 string) (struct {
	Collateral       *big.Int
	CollateralStatus uint8
}, error) {
	var out []interface{}
	err := _SwanCreditCollateral.contract.Call(opts, &out, "tasks", arg0)

	outstruct := new(struct {
		Collateral       *big.Int
		CollateralStatus uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Collateral = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CollateralStatus = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// Tasks is a free data retrieval call binding the contract method 0x58709cf2.
//
// Solidity: function tasks(string ) view returns(uint256 collateral, uint8 collateralStatus)
func (_SwanCreditCollateral *SwanCreditCollateralSession) Tasks(arg0 string) (struct {
	Collateral       *big.Int
	CollateralStatus uint8
}, error) {
	return _SwanCreditCollateral.Contract.Tasks(&_SwanCreditCollateral.CallOpts, arg0)
}

// Tasks is a free data retrieval call binding the contract method 0x58709cf2.
//
// Solidity: function tasks(string ) view returns(uint256 collateral, uint8 collateralStatus)
func (_SwanCreditCollateral *SwanCreditCollateralCallerSession) Tasks(arg0 string) (struct {
	Collateral       *big.Int
	CollateralStatus uint8
}, error) {
	return _SwanCreditCollateral.Contract.Tasks(&_SwanCreditCollateral.CallOpts, arg0)
}

// UnlockRequest is a free data retrieval call binding the contract method 0xf97348e1.
//
// Solidity: function unlockRequest(address ) view returns(uint256 requestTimestamp, uint256 unlockAmount, bool isPending)
func (_SwanCreditCollateral *SwanCreditCollateralCaller) UnlockRequest(opts *bind.CallOpts, arg0 common.Address) (struct {
	RequestTimestamp *big.Int
	UnlockAmount     *big.Int
	IsPending        bool
}, error) {
	var out []interface{}
	err := _SwanCreditCollateral.contract.Call(opts, &out, "unlockRequest", arg0)

	outstruct := new(struct {
		RequestTimestamp *big.Int
		UnlockAmount     *big.Int
		IsPending        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequestTimestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UnlockAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsPending = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// UnlockRequest is a free data retrieval call binding the contract method 0xf97348e1.
//
// Solidity: function unlockRequest(address ) view returns(uint256 requestTimestamp, uint256 unlockAmount, bool isPending)
func (_SwanCreditCollateral *SwanCreditCollateralSession) UnlockRequest(arg0 common.Address) (struct {
	RequestTimestamp *big.Int
	UnlockAmount     *big.Int
	IsPending        bool
}, error) {
	return _SwanCreditCollateral.Contract.UnlockRequest(&_SwanCreditCollateral.CallOpts, arg0)
}

// UnlockRequest is a free data retrieval call binding the contract method 0xf97348e1.
//
// Solidity: function unlockRequest(address ) view returns(uint256 requestTimestamp, uint256 unlockAmount, bool isPending)
func (_SwanCreditCollateral *SwanCreditCollateralCallerSession) UnlockRequest(arg0 common.Address) (struct {
	RequestTimestamp *big.Int
	UnlockAmount     *big.Int
	IsPending        bool
}, error) {
	return _SwanCreditCollateral.Contract.UnlockRequest(&_SwanCreditCollateral.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_SwanCreditCollateral *SwanCreditCollateralCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwanCreditCollateral.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_SwanCreditCollateral *SwanCreditCollateralSession) Version() (*big.Int, error) {
	return _SwanCreditCollateral.Contract.Version(&_SwanCreditCollateral.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_SwanCreditCollateral *SwanCreditCollateralCallerSession) Version() (*big.Int, error) {
	return _SwanCreditCollateral.Contract.Version(&_SwanCreditCollateral.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address newAdmin) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactor) AddAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _SwanCreditCollateral.contract.Transact(opts, "addAdmin", newAdmin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address newAdmin) returns()
func (_SwanCreditCollateral *SwanCreditCollateralSession) AddAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.AddAdmin(&_SwanCreditCollateral.TransactOpts, newAdmin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address newAdmin) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactorSession) AddAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.AddAdmin(&_SwanCreditCollateral.TransactOpts, newAdmin)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address cpAccount, uint256 amount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactor) Deposit(opts *bind.TransactOpts, cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.contract.Transact(opts, "deposit", cpAccount, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address cpAccount, uint256 amount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralSession) Deposit(cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.Deposit(&_SwanCreditCollateral.TransactOpts, cpAccount, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address cpAccount, uint256 amount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactorSession) Deposit(cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.Deposit(&_SwanCreditCollateral.TransactOpts, cpAccount, amount)
}

// DisputeProof is a paid mutator transaction binding the contract method 0xe80d9dec.
//
// Solidity: function disputeProof(string taskUid) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactor) DisputeProof(opts *bind.TransactOpts, taskUid string) (*types.Transaction, error) {
	return _SwanCreditCollateral.contract.Transact(opts, "disputeProof", taskUid)
}

// DisputeProof is a paid mutator transaction binding the contract method 0xe80d9dec.
//
// Solidity: function disputeProof(string taskUid) returns()
func (_SwanCreditCollateral *SwanCreditCollateralSession) DisputeProof(taskUid string) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.DisputeProof(&_SwanCreditCollateral.TransactOpts, taskUid)
}

// DisputeProof is a paid mutator transaction binding the contract method 0xe80d9dec.
//
// Solidity: function disputeProof(string taskUid) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactorSession) DisputeProof(taskUid string) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.DisputeProof(&_SwanCreditCollateral.TransactOpts, taskUid)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwanCreditCollateral.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_SwanCreditCollateral *SwanCreditCollateralSession) Initialize() (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.Initialize(&_SwanCreditCollateral.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactorSession) Initialize() (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.Initialize(&_SwanCreditCollateral.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactor) RemoveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _SwanCreditCollateral.contract.Transact(opts, "removeAdmin", admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_SwanCreditCollateral *SwanCreditCollateralSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.RemoveAdmin(&_SwanCreditCollateral.TransactOpts, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactorSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.RemoveAdmin(&_SwanCreditCollateral.TransactOpts, admin)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwanCreditCollateral.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SwanCreditCollateral *SwanCreditCollateralSession) RenounceOwnership() (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.RenounceOwnership(&_SwanCreditCollateral.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.RenounceOwnership(&_SwanCreditCollateral.TransactOpts)
}

// SetBaseCollateral is a paid mutator transaction binding the contract method 0x7f58a7e5.
//
// Solidity: function setBaseCollateral(uint256 _baseCollateral) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactor) SetBaseCollateral(opts *bind.TransactOpts, _baseCollateral *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.contract.Transact(opts, "setBaseCollateral", _baseCollateral)
}

// SetBaseCollateral is a paid mutator transaction binding the contract method 0x7f58a7e5.
//
// Solidity: function setBaseCollateral(uint256 _baseCollateral) returns()
func (_SwanCreditCollateral *SwanCreditCollateralSession) SetBaseCollateral(_baseCollateral *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.SetBaseCollateral(&_SwanCreditCollateral.TransactOpts, _baseCollateral)
}

// SetBaseCollateral is a paid mutator transaction binding the contract method 0x7f58a7e5.
//
// Solidity: function setBaseCollateral(uint256 _baseCollateral) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactorSession) SetBaseCollateral(_baseCollateral *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.SetBaseCollateral(&_SwanCreditCollateral.TransactOpts, _baseCollateral)
}

// SetCollateralToken is a paid mutator transaction binding the contract method 0x666181a9.
//
// Solidity: function setCollateralToken(address tokenAddress) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactor) SetCollateralToken(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _SwanCreditCollateral.contract.Transact(opts, "setCollateralToken", tokenAddress)
}

// SetCollateralToken is a paid mutator transaction binding the contract method 0x666181a9.
//
// Solidity: function setCollateralToken(address tokenAddress) returns()
func (_SwanCreditCollateral *SwanCreditCollateralSession) SetCollateralToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.SetCollateralToken(&_SwanCreditCollateral.TransactOpts, tokenAddress)
}

// SetCollateralToken is a paid mutator transaction binding the contract method 0x666181a9.
//
// Solidity: function setCollateralToken(address tokenAddress) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactorSession) SetCollateralToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.SetCollateralToken(&_SwanCreditCollateral.TransactOpts, tokenAddress)
}

// SetSlashRatio is a paid mutator transaction binding the contract method 0xce3518aa.
//
// Solidity: function setSlashRatio(uint256 _slashRatio) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactor) SetSlashRatio(opts *bind.TransactOpts, _slashRatio *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.contract.Transact(opts, "setSlashRatio", _slashRatio)
}

// SetSlashRatio is a paid mutator transaction binding the contract method 0xce3518aa.
//
// Solidity: function setSlashRatio(uint256 _slashRatio) returns()
func (_SwanCreditCollateral *SwanCreditCollateralSession) SetSlashRatio(_slashRatio *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.SetSlashRatio(&_SwanCreditCollateral.TransactOpts, _slashRatio)
}

// SetSlashRatio is a paid mutator transaction binding the contract method 0xce3518aa.
//
// Solidity: function setSlashRatio(uint256 _slashRatio) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactorSession) SetSlashRatio(_slashRatio *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.SetSlashRatio(&_SwanCreditCollateral.TransactOpts, _slashRatio)
}

// SlashCollateral is a paid mutator transaction binding the contract method 0x1d47a62d.
//
// Solidity: function slashCollateral(address cpAccount, uint256 slashAmount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactor) SlashCollateral(opts *bind.TransactOpts, cpAccount common.Address, slashAmount *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.contract.Transact(opts, "slashCollateral", cpAccount, slashAmount)
}

// SlashCollateral is a paid mutator transaction binding the contract method 0x1d47a62d.
//
// Solidity: function slashCollateral(address cpAccount, uint256 slashAmount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralSession) SlashCollateral(cpAccount common.Address, slashAmount *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.SlashCollateral(&_SwanCreditCollateral.TransactOpts, cpAccount, slashAmount)
}

// SlashCollateral is a paid mutator transaction binding the contract method 0x1d47a62d.
//
// Solidity: function slashCollateral(address cpAccount, uint256 slashAmount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactorSession) SlashCollateral(cpAccount common.Address, slashAmount *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.SlashCollateral(&_SwanCreditCollateral.TransactOpts, cpAccount, slashAmount)
}

// SlashCollateral0 is a paid mutator transaction binding the contract method 0xb587b82c.
//
// Solidity: function slashCollateral(string taskUid, address cpAccount, uint256 slashAmount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactor) SlashCollateral0(opts *bind.TransactOpts, taskUid string, cpAccount common.Address, slashAmount *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.contract.Transact(opts, "slashCollateral0", taskUid, cpAccount, slashAmount)
}

// SlashCollateral0 is a paid mutator transaction binding the contract method 0xb587b82c.
//
// Solidity: function slashCollateral(string taskUid, address cpAccount, uint256 slashAmount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralSession) SlashCollateral0(taskUid string, cpAccount common.Address, slashAmount *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.SlashCollateral0(&_SwanCreditCollateral.TransactOpts, taskUid, cpAccount, slashAmount)
}

// SlashCollateral0 is a paid mutator transaction binding the contract method 0xb587b82c.
//
// Solidity: function slashCollateral(string taskUid, address cpAccount, uint256 slashAmount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactorSession) SlashCollateral0(taskUid string, cpAccount common.Address, slashAmount *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.SlashCollateral0(&_SwanCreditCollateral.TransactOpts, taskUid, cpAccount, slashAmount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SwanCreditCollateral.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SwanCreditCollateral *SwanCreditCollateralSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.TransferOwnership(&_SwanCreditCollateral.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.TransferOwnership(&_SwanCreditCollateral.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SwanCreditCollateral.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SwanCreditCollateral *SwanCreditCollateralSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.UpgradeToAndCall(&_SwanCreditCollateral.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.UpgradeToAndCall(&_SwanCreditCollateral.TransactOpts, newImplementation, data)
}

// WithdrawSlashedFunds is a paid mutator transaction binding the contract method 0x2894493f.
//
// Solidity: function withdrawSlashedFunds(uint256 amount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactor) WithdrawSlashedFunds(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.contract.Transact(opts, "withdrawSlashedFunds", amount)
}

// WithdrawSlashedFunds is a paid mutator transaction binding the contract method 0x2894493f.
//
// Solidity: function withdrawSlashedFunds(uint256 amount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralSession) WithdrawSlashedFunds(amount *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.WithdrawSlashedFunds(&_SwanCreditCollateral.TransactOpts, amount)
}

// WithdrawSlashedFunds is a paid mutator transaction binding the contract method 0x2894493f.
//
// Solidity: function withdrawSlashedFunds(uint256 amount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactorSession) WithdrawSlashedFunds(amount *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.WithdrawSlashedFunds(&_SwanCreditCollateral.TransactOpts, amount)
}

// SwanCreditCollateralCollateralAdjustedIterator is returned from FilterCollateralAdjusted and is used to iterate over the raw logs and unpacked data for CollateralAdjusted events raised by the SwanCreditCollateral contract.
type SwanCreditCollateralCollateralAdjustedIterator struct {
	Event *SwanCreditCollateralCollateralAdjusted // Event containing the contract specifics and raw log

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
func (it *SwanCreditCollateralCollateralAdjustedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwanCreditCollateralCollateralAdjusted)
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
		it.Event = new(SwanCreditCollateralCollateralAdjusted)
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
func (it *SwanCreditCollateralCollateralAdjustedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwanCreditCollateralCollateralAdjustedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwanCreditCollateralCollateralAdjusted represents a CollateralAdjusted event raised by the SwanCreditCollateral contract.
type SwanCreditCollateralCollateralAdjusted struct {
	Cp            common.Address
	FrozenAmount  *big.Int
	BalanceAmount *big.Int
	Operation     string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCollateralAdjusted is a free log retrieval operation binding the contract event 0x42f1a8a1aee108e84e4eabfaa8d51d7baaa1a02d482295297883a44b2debd3e4.
//
// Solidity: event CollateralAdjusted(address indexed cp, uint256 frozenAmount, uint256 balanceAmount, string operation)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) FilterCollateralAdjusted(opts *bind.FilterOpts, cp []common.Address) (*SwanCreditCollateralCollateralAdjustedIterator, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _SwanCreditCollateral.contract.FilterLogs(opts, "CollateralAdjusted", cpRule)
	if err != nil {
		return nil, err
	}
	return &SwanCreditCollateralCollateralAdjustedIterator{contract: _SwanCreditCollateral.contract, event: "CollateralAdjusted", logs: logs, sub: sub}, nil
}

// WatchCollateralAdjusted is a free log subscription operation binding the contract event 0x42f1a8a1aee108e84e4eabfaa8d51d7baaa1a02d482295297883a44b2debd3e4.
//
// Solidity: event CollateralAdjusted(address indexed cp, uint256 frozenAmount, uint256 balanceAmount, string operation)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) WatchCollateralAdjusted(opts *bind.WatchOpts, sink chan<- *SwanCreditCollateralCollateralAdjusted, cp []common.Address) (event.Subscription, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _SwanCreditCollateral.contract.WatchLogs(opts, "CollateralAdjusted", cpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwanCreditCollateralCollateralAdjusted)
				if err := _SwanCreditCollateral.contract.UnpackLog(event, "CollateralAdjusted", log); err != nil {
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

// ParseCollateralAdjusted is a log parse operation binding the contract event 0x42f1a8a1aee108e84e4eabfaa8d51d7baaa1a02d482295297883a44b2debd3e4.
//
// Solidity: event CollateralAdjusted(address indexed cp, uint256 frozenAmount, uint256 balanceAmount, string operation)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) ParseCollateralAdjusted(log types.Log) (*SwanCreditCollateralCollateralAdjusted, error) {
	event := new(SwanCreditCollateralCollateralAdjusted)
	if err := _SwanCreditCollateral.contract.UnpackLog(event, "CollateralAdjusted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwanCreditCollateralCollateralLockedIterator is returned from FilterCollateralLocked and is used to iterate over the raw logs and unpacked data for CollateralLocked events raised by the SwanCreditCollateral contract.
type SwanCreditCollateralCollateralLockedIterator struct {
	Event *SwanCreditCollateralCollateralLocked // Event containing the contract specifics and raw log

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
func (it *SwanCreditCollateralCollateralLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwanCreditCollateralCollateralLocked)
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
		it.Event = new(SwanCreditCollateralCollateralLocked)
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
func (it *SwanCreditCollateralCollateralLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwanCreditCollateralCollateralLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwanCreditCollateralCollateralLocked represents a CollateralLocked event raised by the SwanCreditCollateral contract.
type SwanCreditCollateralCollateralLocked struct {
	Cp               common.Address
	CollateralAmount *big.Int
	TaskUid          string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCollateralLocked is a free log retrieval operation binding the contract event 0x5f3d004cf9164b95ed5dbf47d1f04018a4eabcb20b4320fe229ed92236ace634.
//
// Solidity: event CollateralLocked(address indexed cp, uint256 collateralAmount, string taskUid)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) FilterCollateralLocked(opts *bind.FilterOpts, cp []common.Address) (*SwanCreditCollateralCollateralLockedIterator, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _SwanCreditCollateral.contract.FilterLogs(opts, "CollateralLocked", cpRule)
	if err != nil {
		return nil, err
	}
	return &SwanCreditCollateralCollateralLockedIterator{contract: _SwanCreditCollateral.contract, event: "CollateralLocked", logs: logs, sub: sub}, nil
}

// WatchCollateralLocked is a free log subscription operation binding the contract event 0x5f3d004cf9164b95ed5dbf47d1f04018a4eabcb20b4320fe229ed92236ace634.
//
// Solidity: event CollateralLocked(address indexed cp, uint256 collateralAmount, string taskUid)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) WatchCollateralLocked(opts *bind.WatchOpts, sink chan<- *SwanCreditCollateralCollateralLocked, cp []common.Address) (event.Subscription, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _SwanCreditCollateral.contract.WatchLogs(opts, "CollateralLocked", cpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwanCreditCollateralCollateralLocked)
				if err := _SwanCreditCollateral.contract.UnpackLog(event, "CollateralLocked", log); err != nil {
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

// ParseCollateralLocked is a log parse operation binding the contract event 0x5f3d004cf9164b95ed5dbf47d1f04018a4eabcb20b4320fe229ed92236ace634.
//
// Solidity: event CollateralLocked(address indexed cp, uint256 collateralAmount, string taskUid)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) ParseCollateralLocked(log types.Log) (*SwanCreditCollateralCollateralLocked, error) {
	event := new(SwanCreditCollateralCollateralLocked)
	if err := _SwanCreditCollateral.contract.UnpackLog(event, "CollateralLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwanCreditCollateralCollateralSlashedIterator is returned from FilterCollateralSlashed and is used to iterate over the raw logs and unpacked data for CollateralSlashed events raised by the SwanCreditCollateral contract.
type SwanCreditCollateralCollateralSlashedIterator struct {
	Event *SwanCreditCollateralCollateralSlashed // Event containing the contract specifics and raw log

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
func (it *SwanCreditCollateralCollateralSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwanCreditCollateralCollateralSlashed)
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
		it.Event = new(SwanCreditCollateralCollateralSlashed)
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
func (it *SwanCreditCollateralCollateralSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwanCreditCollateralCollateralSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwanCreditCollateralCollateralSlashed represents a CollateralSlashed event raised by the SwanCreditCollateral contract.
type SwanCreditCollateralCollateralSlashed struct {
	Cp      common.Address
	Amount  *big.Int
	TaskUid string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCollateralSlashed is a free log retrieval operation binding the contract event 0x403feb2cd85cc25c910fe59289105b583f08fe9e4335ebbb08c1851f47ff2d34.
//
// Solidity: event CollateralSlashed(address indexed cp, uint256 amount, string taskUid)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) FilterCollateralSlashed(opts *bind.FilterOpts, cp []common.Address) (*SwanCreditCollateralCollateralSlashedIterator, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _SwanCreditCollateral.contract.FilterLogs(opts, "CollateralSlashed", cpRule)
	if err != nil {
		return nil, err
	}
	return &SwanCreditCollateralCollateralSlashedIterator{contract: _SwanCreditCollateral.contract, event: "CollateralSlashed", logs: logs, sub: sub}, nil
}

// WatchCollateralSlashed is a free log subscription operation binding the contract event 0x403feb2cd85cc25c910fe59289105b583f08fe9e4335ebbb08c1851f47ff2d34.
//
// Solidity: event CollateralSlashed(address indexed cp, uint256 amount, string taskUid)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) WatchCollateralSlashed(opts *bind.WatchOpts, sink chan<- *SwanCreditCollateralCollateralSlashed, cp []common.Address) (event.Subscription, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _SwanCreditCollateral.contract.WatchLogs(opts, "CollateralSlashed", cpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwanCreditCollateralCollateralSlashed)
				if err := _SwanCreditCollateral.contract.UnpackLog(event, "CollateralSlashed", log); err != nil {
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

// ParseCollateralSlashed is a log parse operation binding the contract event 0x403feb2cd85cc25c910fe59289105b583f08fe9e4335ebbb08c1851f47ff2d34.
//
// Solidity: event CollateralSlashed(address indexed cp, uint256 amount, string taskUid)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) ParseCollateralSlashed(log types.Log) (*SwanCreditCollateralCollateralSlashed, error) {
	event := new(SwanCreditCollateralCollateralSlashed)
	if err := _SwanCreditCollateral.contract.UnpackLog(event, "CollateralSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwanCreditCollateralCollateralStatusChangedIterator is returned from FilterCollateralStatusChanged and is used to iterate over the raw logs and unpacked data for CollateralStatusChanged events raised by the SwanCreditCollateral contract.
type SwanCreditCollateralCollateralStatusChangedIterator struct {
	Event *SwanCreditCollateralCollateralStatusChanged // Event containing the contract specifics and raw log

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
func (it *SwanCreditCollateralCollateralStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwanCreditCollateralCollateralStatusChanged)
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
		it.Event = new(SwanCreditCollateralCollateralStatusChanged)
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
func (it *SwanCreditCollateralCollateralStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwanCreditCollateralCollateralStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwanCreditCollateralCollateralStatusChanged represents a CollateralStatusChanged event raised by the SwanCreditCollateral contract.
type SwanCreditCollateralCollateralStatusChanged struct {
	TaskUid   common.Hash
	NewStatus uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollateralStatusChanged is a free log retrieval operation binding the contract event 0x4a2ced9ada462e244851a86e998eb0b5bf558c2c9c6923b7f970ed2b19b073b0.
//
// Solidity: event CollateralStatusChanged(string indexed taskUid, uint8 newStatus)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) FilterCollateralStatusChanged(opts *bind.FilterOpts, taskUid []string) (*SwanCreditCollateralCollateralStatusChangedIterator, error) {

	var taskUidRule []interface{}
	for _, taskUidItem := range taskUid {
		taskUidRule = append(taskUidRule, taskUidItem)
	}

	logs, sub, err := _SwanCreditCollateral.contract.FilterLogs(opts, "CollateralStatusChanged", taskUidRule)
	if err != nil {
		return nil, err
	}
	return &SwanCreditCollateralCollateralStatusChangedIterator{contract: _SwanCreditCollateral.contract, event: "CollateralStatusChanged", logs: logs, sub: sub}, nil
}

// WatchCollateralStatusChanged is a free log subscription operation binding the contract event 0x4a2ced9ada462e244851a86e998eb0b5bf558c2c9c6923b7f970ed2b19b073b0.
//
// Solidity: event CollateralStatusChanged(string indexed taskUid, uint8 newStatus)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) WatchCollateralStatusChanged(opts *bind.WatchOpts, sink chan<- *SwanCreditCollateralCollateralStatusChanged, taskUid []string) (event.Subscription, error) {

	var taskUidRule []interface{}
	for _, taskUidItem := range taskUid {
		taskUidRule = append(taskUidRule, taskUidItem)
	}

	logs, sub, err := _SwanCreditCollateral.contract.WatchLogs(opts, "CollateralStatusChanged", taskUidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwanCreditCollateralCollateralStatusChanged)
				if err := _SwanCreditCollateral.contract.UnpackLog(event, "CollateralStatusChanged", log); err != nil {
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

// ParseCollateralStatusChanged is a log parse operation binding the contract event 0x4a2ced9ada462e244851a86e998eb0b5bf558c2c9c6923b7f970ed2b19b073b0.
//
// Solidity: event CollateralStatusChanged(string indexed taskUid, uint8 newStatus)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) ParseCollateralStatusChanged(log types.Log) (*SwanCreditCollateralCollateralStatusChanged, error) {
	event := new(SwanCreditCollateralCollateralStatusChanged)
	if err := _SwanCreditCollateral.contract.UnpackLog(event, "CollateralStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwanCreditCollateralCollateralUnlockedIterator is returned from FilterCollateralUnlocked and is used to iterate over the raw logs and unpacked data for CollateralUnlocked events raised by the SwanCreditCollateral contract.
type SwanCreditCollateralCollateralUnlockedIterator struct {
	Event *SwanCreditCollateralCollateralUnlocked // Event containing the contract specifics and raw log

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
func (it *SwanCreditCollateralCollateralUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwanCreditCollateralCollateralUnlocked)
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
		it.Event = new(SwanCreditCollateralCollateralUnlocked)
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
func (it *SwanCreditCollateralCollateralUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwanCreditCollateralCollateralUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwanCreditCollateralCollateralUnlocked represents a CollateralUnlocked event raised by the SwanCreditCollateral contract.
type SwanCreditCollateralCollateralUnlocked struct {
	Cp               common.Address
	CollateralAmount *big.Int
	TaskUid          string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCollateralUnlocked is a free log retrieval operation binding the contract event 0x0f2440b3ca071b7d18e917a25289e7d7e7de8a491546d45acc2efbec7b3e1ae8.
//
// Solidity: event CollateralUnlocked(address indexed cp, uint256 collateralAmount, string taskUid)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) FilterCollateralUnlocked(opts *bind.FilterOpts, cp []common.Address) (*SwanCreditCollateralCollateralUnlockedIterator, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _SwanCreditCollateral.contract.FilterLogs(opts, "CollateralUnlocked", cpRule)
	if err != nil {
		return nil, err
	}
	return &SwanCreditCollateralCollateralUnlockedIterator{contract: _SwanCreditCollateral.contract, event: "CollateralUnlocked", logs: logs, sub: sub}, nil
}

// WatchCollateralUnlocked is a free log subscription operation binding the contract event 0x0f2440b3ca071b7d18e917a25289e7d7e7de8a491546d45acc2efbec7b3e1ae8.
//
// Solidity: event CollateralUnlocked(address indexed cp, uint256 collateralAmount, string taskUid)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) WatchCollateralUnlocked(opts *bind.WatchOpts, sink chan<- *SwanCreditCollateralCollateralUnlocked, cp []common.Address) (event.Subscription, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _SwanCreditCollateral.contract.WatchLogs(opts, "CollateralUnlocked", cpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwanCreditCollateralCollateralUnlocked)
				if err := _SwanCreditCollateral.contract.UnpackLog(event, "CollateralUnlocked", log); err != nil {
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

// ParseCollateralUnlocked is a log parse operation binding the contract event 0x0f2440b3ca071b7d18e917a25289e7d7e7de8a491546d45acc2efbec7b3e1ae8.
//
// Solidity: event CollateralUnlocked(address indexed cp, uint256 collateralAmount, string taskUid)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) ParseCollateralUnlocked(log types.Log) (*SwanCreditCollateralCollateralUnlocked, error) {
	event := new(SwanCreditCollateralCollateralUnlocked)
	if err := _SwanCreditCollateral.contract.UnpackLog(event, "CollateralUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwanCreditCollateralDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the SwanCreditCollateral contract.
type SwanCreditCollateralDepositIterator struct {
	Event *SwanCreditCollateralDeposit // Event containing the contract specifics and raw log

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
func (it *SwanCreditCollateralDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwanCreditCollateralDeposit)
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
		it.Event = new(SwanCreditCollateralDeposit)
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
func (it *SwanCreditCollateralDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwanCreditCollateralDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwanCreditCollateralDeposit represents a Deposit event raised by the SwanCreditCollateral contract.
type SwanCreditCollateralDeposit struct {
	FundingWallet common.Address
	CpAccount     common.Address
	DepositAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed fundingWallet, address indexed cpAccount, uint256 depositAmount)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) FilterDeposit(opts *bind.FilterOpts, fundingWallet []common.Address, cpAccount []common.Address) (*SwanCreditCollateralDepositIterator, error) {

	var fundingWalletRule []interface{}
	for _, fundingWalletItem := range fundingWallet {
		fundingWalletRule = append(fundingWalletRule, fundingWalletItem)
	}
	var cpAccountRule []interface{}
	for _, cpAccountItem := range cpAccount {
		cpAccountRule = append(cpAccountRule, cpAccountItem)
	}

	logs, sub, err := _SwanCreditCollateral.contract.FilterLogs(opts, "Deposit", fundingWalletRule, cpAccountRule)
	if err != nil {
		return nil, err
	}
	return &SwanCreditCollateralDepositIterator{contract: _SwanCreditCollateral.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed fundingWallet, address indexed cpAccount, uint256 depositAmount)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *SwanCreditCollateralDeposit, fundingWallet []common.Address, cpAccount []common.Address) (event.Subscription, error) {

	var fundingWalletRule []interface{}
	for _, fundingWalletItem := range fundingWallet {
		fundingWalletRule = append(fundingWalletRule, fundingWalletItem)
	}
	var cpAccountRule []interface{}
	for _, cpAccountItem := range cpAccount {
		cpAccountRule = append(cpAccountRule, cpAccountItem)
	}

	logs, sub, err := _SwanCreditCollateral.contract.WatchLogs(opts, "Deposit", fundingWalletRule, cpAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwanCreditCollateralDeposit)
				if err := _SwanCreditCollateral.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed fundingWallet, address indexed cpAccount, uint256 depositAmount)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) ParseDeposit(log types.Log) (*SwanCreditCollateralDeposit, error) {
	event := new(SwanCreditCollateralDeposit)
	if err := _SwanCreditCollateral.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwanCreditCollateralDisputeProofIterator is returned from FilterDisputeProof and is used to iterate over the raw logs and unpacked data for DisputeProof events raised by the SwanCreditCollateral contract.
type SwanCreditCollateralDisputeProofIterator struct {
	Event *SwanCreditCollateralDisputeProof // Event containing the contract specifics and raw log

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
func (it *SwanCreditCollateralDisputeProofIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwanCreditCollateralDisputeProof)
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
		it.Event = new(SwanCreditCollateralDisputeProof)
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
func (it *SwanCreditCollateralDisputeProofIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwanCreditCollateralDisputeProofIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwanCreditCollateralDisputeProof represents a DisputeProof event raised by the SwanCreditCollateral contract.
type SwanCreditCollateralDisputeProof struct {
	Challenger common.Address
	TaskUid    common.Hash
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDisputeProof is a free log retrieval operation binding the contract event 0xaec1d412a3c1e4a13fc2a2e19ac38a5af192a9cf17b074fca8146a2d0655e0c3.
//
// Solidity: event DisputeProof(address indexed challenger, string indexed taskUid)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) FilterDisputeProof(opts *bind.FilterOpts, challenger []common.Address, taskUid []string) (*SwanCreditCollateralDisputeProofIterator, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}
	var taskUidRule []interface{}
	for _, taskUidItem := range taskUid {
		taskUidRule = append(taskUidRule, taskUidItem)
	}

	logs, sub, err := _SwanCreditCollateral.contract.FilterLogs(opts, "DisputeProof", challengerRule, taskUidRule)
	if err != nil {
		return nil, err
	}
	return &SwanCreditCollateralDisputeProofIterator{contract: _SwanCreditCollateral.contract, event: "DisputeProof", logs: logs, sub: sub}, nil
}

// WatchDisputeProof is a free log subscription operation binding the contract event 0xaec1d412a3c1e4a13fc2a2e19ac38a5af192a9cf17b074fca8146a2d0655e0c3.
//
// Solidity: event DisputeProof(address indexed challenger, string indexed taskUid)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) WatchDisputeProof(opts *bind.WatchOpts, sink chan<- *SwanCreditCollateralDisputeProof, challenger []common.Address, taskUid []string) (event.Subscription, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}
	var taskUidRule []interface{}
	for _, taskUidItem := range taskUid {
		taskUidRule = append(taskUidRule, taskUidItem)
	}

	logs, sub, err := _SwanCreditCollateral.contract.WatchLogs(opts, "DisputeProof", challengerRule, taskUidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwanCreditCollateralDisputeProof)
				if err := _SwanCreditCollateral.contract.UnpackLog(event, "DisputeProof", log); err != nil {
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

// ParseDisputeProof is a log parse operation binding the contract event 0xaec1d412a3c1e4a13fc2a2e19ac38a5af192a9cf17b074fca8146a2d0655e0c3.
//
// Solidity: event DisputeProof(address indexed challenger, string indexed taskUid)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) ParseDisputeProof(log types.Log) (*SwanCreditCollateralDisputeProof, error) {
	event := new(SwanCreditCollateralDisputeProof)
	if err := _SwanCreditCollateral.contract.UnpackLog(event, "DisputeProof", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwanCreditCollateralInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SwanCreditCollateral contract.
type SwanCreditCollateralInitializedIterator struct {
	Event *SwanCreditCollateralInitialized // Event containing the contract specifics and raw log

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
func (it *SwanCreditCollateralInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwanCreditCollateralInitialized)
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
		it.Event = new(SwanCreditCollateralInitialized)
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
func (it *SwanCreditCollateralInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwanCreditCollateralInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwanCreditCollateralInitialized represents a Initialized event raised by the SwanCreditCollateral contract.
type SwanCreditCollateralInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) FilterInitialized(opts *bind.FilterOpts) (*SwanCreditCollateralInitializedIterator, error) {

	logs, sub, err := _SwanCreditCollateral.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SwanCreditCollateralInitializedIterator{contract: _SwanCreditCollateral.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SwanCreditCollateralInitialized) (event.Subscription, error) {

	logs, sub, err := _SwanCreditCollateral.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwanCreditCollateralInitialized)
				if err := _SwanCreditCollateral.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) ParseInitialized(log types.Log) (*SwanCreditCollateralInitialized, error) {
	event := new(SwanCreditCollateralInitialized)
	if err := _SwanCreditCollateral.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwanCreditCollateralOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SwanCreditCollateral contract.
type SwanCreditCollateralOwnershipTransferredIterator struct {
	Event *SwanCreditCollateralOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SwanCreditCollateralOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwanCreditCollateralOwnershipTransferred)
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
		it.Event = new(SwanCreditCollateralOwnershipTransferred)
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
func (it *SwanCreditCollateralOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwanCreditCollateralOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwanCreditCollateralOwnershipTransferred represents a OwnershipTransferred event raised by the SwanCreditCollateral contract.
type SwanCreditCollateralOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SwanCreditCollateralOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SwanCreditCollateral.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SwanCreditCollateralOwnershipTransferredIterator{contract: _SwanCreditCollateral.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SwanCreditCollateralOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SwanCreditCollateral.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwanCreditCollateralOwnershipTransferred)
				if err := _SwanCreditCollateral.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) ParseOwnershipTransferred(log types.Log) (*SwanCreditCollateralOwnershipTransferred, error) {
	event := new(SwanCreditCollateralOwnershipTransferred)
	if err := _SwanCreditCollateral.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwanCreditCollateralSlashedFundsIncreasedIterator is returned from FilterSlashedFundsIncreased and is used to iterate over the raw logs and unpacked data for SlashedFundsIncreased events raised by the SwanCreditCollateral contract.
type SwanCreditCollateralSlashedFundsIncreasedIterator struct {
	Event *SwanCreditCollateralSlashedFundsIncreased // Event containing the contract specifics and raw log

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
func (it *SwanCreditCollateralSlashedFundsIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwanCreditCollateralSlashedFundsIncreased)
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
		it.Event = new(SwanCreditCollateralSlashedFundsIncreased)
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
func (it *SwanCreditCollateralSlashedFundsIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwanCreditCollateralSlashedFundsIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwanCreditCollateralSlashedFundsIncreased represents a SlashedFundsIncreased event raised by the SwanCreditCollateral contract.
type SwanCreditCollateralSlashedFundsIncreased struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSlashedFundsIncreased is a free log retrieval operation binding the contract event 0xe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b.
//
// Solidity: event SlashedFundsIncreased(uint256 amount)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) FilterSlashedFundsIncreased(opts *bind.FilterOpts) (*SwanCreditCollateralSlashedFundsIncreasedIterator, error) {

	logs, sub, err := _SwanCreditCollateral.contract.FilterLogs(opts, "SlashedFundsIncreased")
	if err != nil {
		return nil, err
	}
	return &SwanCreditCollateralSlashedFundsIncreasedIterator{contract: _SwanCreditCollateral.contract, event: "SlashedFundsIncreased", logs: logs, sub: sub}, nil
}

// WatchSlashedFundsIncreased is a free log subscription operation binding the contract event 0xe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b.
//
// Solidity: event SlashedFundsIncreased(uint256 amount)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) WatchSlashedFundsIncreased(opts *bind.WatchOpts, sink chan<- *SwanCreditCollateralSlashedFundsIncreased) (event.Subscription, error) {

	logs, sub, err := _SwanCreditCollateral.contract.WatchLogs(opts, "SlashedFundsIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwanCreditCollateralSlashedFundsIncreased)
				if err := _SwanCreditCollateral.contract.UnpackLog(event, "SlashedFundsIncreased", log); err != nil {
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

// ParseSlashedFundsIncreased is a log parse operation binding the contract event 0xe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b.
//
// Solidity: event SlashedFundsIncreased(uint256 amount)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) ParseSlashedFundsIncreased(log types.Log) (*SwanCreditCollateralSlashedFundsIncreased, error) {
	event := new(SwanCreditCollateralSlashedFundsIncreased)
	if err := _SwanCreditCollateral.contract.UnpackLog(event, "SlashedFundsIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwanCreditCollateralTaskCreatedIterator is returned from FilterTaskCreated and is used to iterate over the raw logs and unpacked data for TaskCreated events raised by the SwanCreditCollateral contract.
type SwanCreditCollateralTaskCreatedIterator struct {
	Event *SwanCreditCollateralTaskCreated // Event containing the contract specifics and raw log

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
func (it *SwanCreditCollateralTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwanCreditCollateralTaskCreated)
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
		it.Event = new(SwanCreditCollateralTaskCreated)
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
func (it *SwanCreditCollateralTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwanCreditCollateralTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwanCreditCollateralTaskCreated represents a TaskCreated event raised by the SwanCreditCollateral contract.
type SwanCreditCollateralTaskCreated struct {
	TaskUid          common.Hash
	CpAccountAddress common.Address
	Collateral       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTaskCreated is a free log retrieval operation binding the contract event 0x5bebc56a5428fd7b8cf43ed525f03223f8363907fbe44665b7a3426d1de96800.
//
// Solidity: event TaskCreated(string indexed taskUid, address cpAccountAddress, uint256 collateral)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) FilterTaskCreated(opts *bind.FilterOpts, taskUid []string) (*SwanCreditCollateralTaskCreatedIterator, error) {

	var taskUidRule []interface{}
	for _, taskUidItem := range taskUid {
		taskUidRule = append(taskUidRule, taskUidItem)
	}

	logs, sub, err := _SwanCreditCollateral.contract.FilterLogs(opts, "TaskCreated", taskUidRule)
	if err != nil {
		return nil, err
	}
	return &SwanCreditCollateralTaskCreatedIterator{contract: _SwanCreditCollateral.contract, event: "TaskCreated", logs: logs, sub: sub}, nil
}

// WatchTaskCreated is a free log subscription operation binding the contract event 0x5bebc56a5428fd7b8cf43ed525f03223f8363907fbe44665b7a3426d1de96800.
//
// Solidity: event TaskCreated(string indexed taskUid, address cpAccountAddress, uint256 collateral)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) WatchTaskCreated(opts *bind.WatchOpts, sink chan<- *SwanCreditCollateralTaskCreated, taskUid []string) (event.Subscription, error) {

	var taskUidRule []interface{}
	for _, taskUidItem := range taskUid {
		taskUidRule = append(taskUidRule, taskUidItem)
	}

	logs, sub, err := _SwanCreditCollateral.contract.WatchLogs(opts, "TaskCreated", taskUidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwanCreditCollateralTaskCreated)
				if err := _SwanCreditCollateral.contract.UnpackLog(event, "TaskCreated", log); err != nil {
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

// ParseTaskCreated is a log parse operation binding the contract event 0x5bebc56a5428fd7b8cf43ed525f03223f8363907fbe44665b7a3426d1de96800.
//
// Solidity: event TaskCreated(string indexed taskUid, address cpAccountAddress, uint256 collateral)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) ParseTaskCreated(log types.Log) (*SwanCreditCollateralTaskCreated, error) {
	event := new(SwanCreditCollateralTaskCreated)
	if err := _SwanCreditCollateral.contract.UnpackLog(event, "TaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwanCreditCollateralUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the SwanCreditCollateral contract.
type SwanCreditCollateralUpgradedIterator struct {
	Event *SwanCreditCollateralUpgraded // Event containing the contract specifics and raw log

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
func (it *SwanCreditCollateralUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwanCreditCollateralUpgraded)
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
		it.Event = new(SwanCreditCollateralUpgraded)
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
func (it *SwanCreditCollateralUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwanCreditCollateralUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwanCreditCollateralUpgraded represents a Upgraded event raised by the SwanCreditCollateral contract.
type SwanCreditCollateralUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*SwanCreditCollateralUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SwanCreditCollateral.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &SwanCreditCollateralUpgradedIterator{contract: _SwanCreditCollateral.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *SwanCreditCollateralUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SwanCreditCollateral.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwanCreditCollateralUpgraded)
				if err := _SwanCreditCollateral.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) ParseUpgraded(log types.Log) (*SwanCreditCollateralUpgraded, error) {
	event := new(SwanCreditCollateralUpgraded)
	if err := _SwanCreditCollateral.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwanCreditCollateralWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the SwanCreditCollateral contract.
type SwanCreditCollateralWithdrawIterator struct {
	Event *SwanCreditCollateralWithdraw // Event containing the contract specifics and raw log

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
func (it *SwanCreditCollateralWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwanCreditCollateralWithdraw)
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
		it.Event = new(SwanCreditCollateralWithdraw)
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
func (it *SwanCreditCollateralWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwanCreditCollateralWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwanCreditCollateralWithdraw represents a Withdraw event raised by the SwanCreditCollateral contract.
type SwanCreditCollateralWithdraw struct {
	CpOwner        common.Address
	CpAccount      common.Address
	WithdrawAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed cpOwner, address indexed cpAccount, uint256 withdrawAmount)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) FilterWithdraw(opts *bind.FilterOpts, cpOwner []common.Address, cpAccount []common.Address) (*SwanCreditCollateralWithdrawIterator, error) {

	var cpOwnerRule []interface{}
	for _, cpOwnerItem := range cpOwner {
		cpOwnerRule = append(cpOwnerRule, cpOwnerItem)
	}
	var cpAccountRule []interface{}
	for _, cpAccountItem := range cpAccount {
		cpAccountRule = append(cpAccountRule, cpAccountItem)
	}

	logs, sub, err := _SwanCreditCollateral.contract.FilterLogs(opts, "Withdraw", cpOwnerRule, cpAccountRule)
	if err != nil {
		return nil, err
	}
	return &SwanCreditCollateralWithdrawIterator{contract: _SwanCreditCollateral.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed cpOwner, address indexed cpAccount, uint256 withdrawAmount)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *SwanCreditCollateralWithdraw, cpOwner []common.Address, cpAccount []common.Address) (event.Subscription, error) {

	var cpOwnerRule []interface{}
	for _, cpOwnerItem := range cpOwner {
		cpOwnerRule = append(cpOwnerRule, cpOwnerItem)
	}
	var cpAccountRule []interface{}
	for _, cpAccountItem := range cpAccount {
		cpAccountRule = append(cpAccountRule, cpAccountItem)
	}

	logs, sub, err := _SwanCreditCollateral.contract.WatchLogs(opts, "Withdraw", cpOwnerRule, cpAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwanCreditCollateralWithdraw)
				if err := _SwanCreditCollateral.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed cpOwner, address indexed cpAccount, uint256 withdrawAmount)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) ParseWithdraw(log types.Log) (*SwanCreditCollateralWithdraw, error) {
	event := new(SwanCreditCollateralWithdraw)
	if err := _SwanCreditCollateral.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwanCreditCollateralWithdrawSlashIterator is returned from FilterWithdrawSlash and is used to iterate over the raw logs and unpacked data for WithdrawSlash events raised by the SwanCreditCollateral contract.
type SwanCreditCollateralWithdrawSlashIterator struct {
	Event *SwanCreditCollateralWithdrawSlash // Event containing the contract specifics and raw log

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
func (it *SwanCreditCollateralWithdrawSlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwanCreditCollateralWithdrawSlash)
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
		it.Event = new(SwanCreditCollateralWithdrawSlash)
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
func (it *SwanCreditCollateralWithdrawSlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwanCreditCollateralWithdrawSlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwanCreditCollateralWithdrawSlash represents a WithdrawSlash event raised by the SwanCreditCollateral contract.
type SwanCreditCollateralWithdrawSlash struct {
	CollateralContratOwner common.Address
	Slashfund              *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawSlash is a free log retrieval operation binding the contract event 0xbfd9c82485e2178fcfd5c40379d8e66fe60acc50aa1ef4c50966431eb1e381cd.
//
// Solidity: event WithdrawSlash(address indexed collateralContratOwner, uint256 slashfund)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) FilterWithdrawSlash(opts *bind.FilterOpts, collateralContratOwner []common.Address) (*SwanCreditCollateralWithdrawSlashIterator, error) {

	var collateralContratOwnerRule []interface{}
	for _, collateralContratOwnerItem := range collateralContratOwner {
		collateralContratOwnerRule = append(collateralContratOwnerRule, collateralContratOwnerItem)
	}

	logs, sub, err := _SwanCreditCollateral.contract.FilterLogs(opts, "WithdrawSlash", collateralContratOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SwanCreditCollateralWithdrawSlashIterator{contract: _SwanCreditCollateral.contract, event: "WithdrawSlash", logs: logs, sub: sub}, nil
}

// WatchWithdrawSlash is a free log subscription operation binding the contract event 0xbfd9c82485e2178fcfd5c40379d8e66fe60acc50aa1ef4c50966431eb1e381cd.
//
// Solidity: event WithdrawSlash(address indexed collateralContratOwner, uint256 slashfund)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) WatchWithdrawSlash(opts *bind.WatchOpts, sink chan<- *SwanCreditCollateralWithdrawSlash, collateralContratOwner []common.Address) (event.Subscription, error) {

	var collateralContratOwnerRule []interface{}
	for _, collateralContratOwnerItem := range collateralContratOwner {
		collateralContratOwnerRule = append(collateralContratOwnerRule, collateralContratOwnerItem)
	}

	logs, sub, err := _SwanCreditCollateral.contract.WatchLogs(opts, "WithdrawSlash", collateralContratOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwanCreditCollateralWithdrawSlash)
				if err := _SwanCreditCollateral.contract.UnpackLog(event, "WithdrawSlash", log); err != nil {
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

// ParseWithdrawSlash is a log parse operation binding the contract event 0xbfd9c82485e2178fcfd5c40379d8e66fe60acc50aa1ef4c50966431eb1e381cd.
//
// Solidity: event WithdrawSlash(address indexed collateralContratOwner, uint256 slashfund)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) ParseWithdrawSlash(log types.Log) (*SwanCreditCollateralWithdrawSlash, error) {
	event := new(SwanCreditCollateralWithdrawSlash)
	if err := _SwanCreditCollateral.contract.UnpackLog(event, "WithdrawSlash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
