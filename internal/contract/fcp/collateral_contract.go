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

// SwanCreditCollateralCPInfoWithLockedBalance is an auto generated low-level Go binding around an user-defined struct.
type SwanCreditCollateralCPInfoWithLockedBalance struct {
	CpAccount        common.Address
	AvailableBalance *big.Int
	LockedBalance    *big.Int
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

// SwanCreditCollateralUnlockRequest is an auto generated low-level Go binding around an user-defined struct.
type SwanCreditCollateralUnlockRequest struct {
	RequestTimestamp *big.Int
	UnlockAmount     *big.Int
	IsPending        bool
}

// SwanCreditCollateralMetaData contains all meta data concerning the SwanCreditCollateral contract.
var SwanCreditCollateralMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"frozenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"operation\",\"type\":\"string\"}],\"name\":\"CollateralAdjusted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"CollateralLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"CollateralSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumSwanCreditCollateral.CollateralStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"CollateralStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"CollateralUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fundingWallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fundingWallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"}],\"name\":\"DepositLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"DisputeProof\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SlashedFundsIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cpAccountAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cpOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"WithdrawConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"WithdrawRequestCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"WithdrawRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralContratOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashfund\",\"type\":\"uint256\"}],\"name\":\"WithdrawSlash\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedToWithdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"availableBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"cpAccounts\",\"type\":\"address[]\"},{\"internalType\":\"int256[]\",\"name\":\"amounts\",\"type\":\"int256[]\"}],\"name\":\"batchAddLockedCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"cpAccounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"}],\"name\":\"cancelWithdrawRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"slashedFunds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashRatio\",\"type\":\"uint256\"}],\"internalType\":\"structSwanCreditCollateral.ContractInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"}],\"name\":\"confirmWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAddress\",\"type\":\"address\"}],\"name\":\"cpInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"availableBalance\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"lockedBalance\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"}],\"internalType\":\"structSwanCreditCollateral.CPInfoWithLockedBalance\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cpStatus\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"escrowBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"getTaskInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"cpList\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"enumSwanCreditCollateral.CollateralStatus\",\"name\":\"collateralStatus\",\"type\":\"uint8\"}],\"internalType\":\"structSwanCreditCollateral.Task\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWithdrawDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"isSignatureUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lockCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockedBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockedCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"requestWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseCollateral\",\"type\":\"uint256\"}],\"name\":\"setBaseCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"setCollateralToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slashRatio\",\"type\":\"uint256\"}],\"name\":\"setSlashRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawDelay\",\"type\":\"uint256\"}],\"name\":\"setWithdrawDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slashAmount\",\"type\":\"uint256\"}],\"name\":\"slashCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slashAmount\",\"type\":\"uint256\"}],\"name\":\"slashCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashedFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"enumSwanCreditCollateral.CollateralStatus\",\"name\":\"collateralStatus\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"taskCollateral\",\"type\":\"uint256\"}],\"name\":\"unlockCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"unlockRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPending\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"}],\"name\":\"viewWithdrawRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requestTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPending\",\"type\":\"bool\"}],\"internalType\":\"structSwanCreditCollateral.UnlockRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawSlashedFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801561004357600080fd5b5061005261005760201b60201c565b6101c1565b600061006761015b60201b60201c565b90508060000160089054906101000a900460ff16156100b2576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff80168160000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff16146101585767ffffffffffffffff8160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d267ffffffffffffffff60405161014f91906101a6565b60405180910390a15b50565b60007ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b600067ffffffffffffffff82169050919050565b6101a081610183565b82525050565b60006020820190506101bb6000830184610197565b92915050565b6080516162fc6101ea60003960008181613cfa01528181613d4f0152613f0a01526162fc6000f3fe6080604052600436106102725760003560e01c8063704802751161014f578063ad3cb1cc116100c1578063d2bfc1c71161007a578063d2bfc1c714610988578063d7635815146109b1578063f2fde38b146109da578063f3fef3a314610a03578063f97348e114610a2c578063fe3300d014610a6b57610272565b8063ad3cb1cc14610878578063b4eae1cb146108a3578063b587b82c146108ce578063bede6e31146108f7578063ce3518aa14610934578063d27ca89b1461095d57610272565b80638da5cb5b116101135780638da5cb5b1461072e57806392bdf9ba146107595780639ae697bf146107965780639b5ddf09146107d3578063a0821be3146107fe578063a664c2161461083b57610272565b80637048027514610685578063715018a6146106ae57806372f0cb30146106c55780637f58a7e5146106ee5780638129fc1c1461071757610272565b806347a7d107116101e857806353ad8720116101ac57806353ad87201461056257806355af63531461058d57806358709cf2146105ca5780635d2cd2a714610608578063666181a9146106315780636f99f15c1461065a57610272565b806347a7d1071461049e57806347e7ef24146104c75780634f1ef286146104f057806352d1902d1461050c578063536f60701461053757610272565b806324d7806c1161023a57806324d7806c1461036c5780632894493f146103a95780632d291cad146103d257806330a907361461040f578063397a1b28146104385780633fe651771461046157610272565b80631150f0f3146102775780631785f53c146102b457806317f17498146102dd5780631b209463146103065780631d47a62d14610343575b600080fd5b34801561028357600080fd5b5061029e6004803603810190610299919061470e565b610a96565b6040516102ab9190614772565b60405180910390f35b3480156102c057600080fd5b506102db60048036038101906102d691906147eb565b610acc565b005b3480156102e957600080fd5b5061030460048036038101906102ff919061484e565b610b2f565b005b34801561031257600080fd5b5061032d6004803603810190610328919061492f565b610d5e565b60405161033a9190614b0c565b60405180910390f35b34801561034f57600080fd5b5061036a6004803603810190610365919061484e565b610e69565b005b34801561037857600080fd5b50610393600480360381019061038e91906147eb565b611434565b6040516103a09190614772565b60405180910390f35b3480156103b557600080fd5b506103d060048036038101906103cb9190614b2e565b611454565b005b3480156103de57600080fd5b506103f960048036038101906103f491906147eb565b6115ab565b6040516104069190614772565b60405180910390f35b34801561041b57600080fd5b5061043660048036038101906104319190614ce6565b6115cb565b005b34801561044457600080fd5b5061045f600480360381019061045a919061484e565b61187e565b005b34801561046d57600080fd5b50610488600480360381019061048391906147eb565b611b8d565b6040516104959190614ddd565b60405180910390f35b3480156104aa57600080fd5b506104c560048036038101906104c0919061484e565b611c2d565b005b3480156104d357600080fd5b506104ee60048036038101906104e9919061484e565b611e85565b005b61050a60048036038101906105059190614dff565b611ff0565b005b34801561051857600080fd5b5061052161200f565b60405161052e9190614e74565b60405180910390f35b34801561054357600080fd5b5061054c612042565b6040516105599190614e9e565b60405180910390f35b34801561056e57600080fd5b50610577612048565b6040516105849190614f0e565b60405180910390f35b34801561059957600080fd5b506105b460048036038101906105af91906147eb565b61207e565b6040516105c19190614f42565b60405180910390f35b3480156105d657600080fd5b506105f160048036038101906105ec919061492f565b612096565b6040516105ff929190614f6c565b60405180910390f35b34801561061457600080fd5b5061062f600480360381019061062a91906147eb565b6120dd565b005b34801561063d57600080fd5b50610658600480360381019061065391906147eb565b6123ca565b005b34801561066657600080fd5b5061066f612416565b60405161067c9190614e9e565b60405180910390f35b34801561069157600080fd5b506106ac60048036038101906106a791906147eb565b61241c565b005b3480156106ba57600080fd5b506106c361247f565b005b3480156106d157600080fd5b506106ec60048036038101906106e79190614b2e565b612493565b005b3480156106fa57600080fd5b5061071560048036038101906107109190614b2e565b6124a5565b005b34801561072357600080fd5b5061072c61253b565b005b34801561073a57600080fd5b50610743612748565b6040516107509190614fa4565b60405180910390f35b34801561076557600080fd5b50610780600480360381019061077b91906147eb565b612780565b60405161078d9190614e9e565b60405180910390f35b3480156107a257600080fd5b506107bd60048036038101906107b891906147eb565b612798565b6040516107ca9190614f42565b60405180910390f35b3480156107df57600080fd5b506107e86127b0565b6040516107f59190614e9e565b60405180910390f35b34801561080a57600080fd5b50610825600480360381019061082091906147eb565b6127b6565b6040516108329190614f42565b60405180910390f35b34801561084757600080fd5b50610862600480360381019061085d91906147eb565b6127ce565b60405161086f919061507b565b60405180910390f35b34801561088457600080fd5b5061088d612959565b60405161089a9190614ddd565b60405180910390f35b3480156108af57600080fd5b506108b8612992565b6040516108c59190614e9e565b60405180910390f35b3480156108da57600080fd5b506108f560048036038101906108f0919061509d565b612998565b005b34801561090357600080fd5b5061091e600480360381019061091991906147eb565b613004565b60405161092b919061515d565b60405180910390f35b34801561094057600080fd5b5061095b60048036038101906109569190614b2e565b61308c565b005b34801561096957600080fd5b5061097261309e565b60405161097f9190614e9e565b60405180910390f35b34801561099457600080fd5b506109af60048036038101906109aa91906147eb565b6130a4565b005b3480156109bd57600080fd5b506109d860048036038101906109d39190615267565b613573565b005b3480156109e657600080fd5b50610a0160048036038101906109fc91906147eb565b6136bc565b005b348015610a0f57600080fd5b50610a2a6004803603810190610a25919061484e565b613742565b005b348015610a3857600080fd5b50610a536004803603810190610a4e91906147eb565b613ace565b604051610a62939291906152df565b60405180910390f35b348015610a7757600080fd5b50610a80613b05565b604051610a8d9190614e9e565b60405180910390f35b600b818051602081018201805184825260208301602085012081835280955050505050506000915054906101000a900460ff1681565b610ad4613b0f565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610bbb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bb290615388565b60405180910390fd5b6000601060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060008113610c42576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c39906153f4565b60405180910390fd5b6000818313610c515782610c53565b815b905080601060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610ca49190615443565b9250508190555080600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610cfa9190615486565b92505081905550610d0a84613b96565b8373ffffffffffffffffffffffffffffffffffffffff167fb4eaf47ecd4bc76248f192433e8067c96cb3e17aced42fbc47a512742fb7421682604051610d509190614e9e565b60405180910390a250505050565b610d666144f8565b600582604051610d769190615506565b908152602001604051809103902060405180606001604052908160008201805480602002602001604051908101604052809291908181526020018280548015610e1457602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610dca575b50505050508152602001600182015481526020016002820160009054906101000a900460ff166002811115610e4c57610e4b614a45565b5b6002811115610e5e57610e5d614a45565b5b815250509050919050565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610ef5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610eec90615388565b60405180910390fd5b601060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548113610fe5577fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b81604051610f6a9190614e9e565b60405180910390a180600080828254610f83919061551d565b9250508190555080601060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610fd99190615443565b925050819055506113d9565b60008190506000601060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054131561111c57601060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600080828254611082919061551d565b92505081905550601060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054816110d49190615443565b90506000601060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b80600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541261120c5780600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546111b19190615443565b92505081905550806000808282546111c9919061551d565b925050819055507fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b826040516111ff9190614e9e565b60405180910390a16113d7565b6000600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054131561133e57600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548161129f9190615443565b9050600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546000808282546112f1919061551d565b925050819055506000600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b7fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b818361136b9190615551565b6040516113789190614e9e565b60405180910390a180601060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546113cf9190615443565b925050819055505b505b6113e282613b96565b8173ffffffffffffffffffffffffffffffffffffffff167f403feb2cd85cc25c910fe59289105b583f08fe9e4335ebbb08c1851f47ff2d348260405161142891906155ab565b60405180910390a25050565b60046020528060005260406000206000915054906101000a900460ff1681565b61145c613b0f565b8060005410156114a1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114989061564b565b60405180910390fd5b806000808282546114b29190615551565b92505081905550600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b815260040161151692919061566b565b6020604051808303816000875af1158015611535573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061155991906156c0565b503373ffffffffffffffffffffffffffffffffffffffff167fbfd9c82485e2178fcfd5c40379d8e66fe60acc50aa1ef4c50966431eb1e381cd826040516115a09190614e9e565b60405180910390a250565b600a6020528060005260406000206000915054906101000a900460ff1681565b805182511461160f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116069061575f565b60405180910390fd5b6000805b82518110156117975782818151811061162f5761162e61577f565b5b602002602001015182611642919061551d565b91508281815181106116575761165661577f565b5b6020026020010151600660008684815181106116765761167561577f565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546116c79190615486565b925050819055506116f18482815181106116e4576116e361577f565b5b6020026020010151613b96565b8381815181106117045761170361577f565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f6285848151811061176d5761176c61577f565b5b60200260200101516040516117829190614e9e565b60405180910390a38080600101915050611613565b50600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b81526004016117f7939291906157ae565b6020604051808303816000875af1158015611816573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061183a91906156c0565b611879576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161187090615831565b60405180910390fd5b505050565b816000808273ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527f893d20e8000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516119299190615898565b6000604051808303816000865af19150503d8060008114611966576040519150601f19603f3d011682016040523d82523d6000602084013e61196b565b606091505b5091509150816119b0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119a790615921565b60405180910390fd5b6000818060200190518101906119c6919061597f565b90508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611a36576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a2d906159f8565b60405180910390fd5b84601060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541215611ab8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611aaf90615a8a565b60405180910390fd5b604051806060016040528043815260200186815260200160011515815250600d60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548160ff0219169083151502179055509050507ff7774b688d56120b783560a913ee60792a73dfd511812b7be5eccf10d08c66898686604051611b7d92919061566b565b60405180910390a1505050505050565b60086020528060005260406000206000915090508054611bac90615ad9565b80601f0160208091040260200160405190810160405280929190818152602001828054611bd890615ad9565b8015611c255780601f10611bfa57610100808354040283529160200191611c25565b820191906000526020600020905b815481529060010190602001808311611c0857829003601f168201915b505050505081565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611cb9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611cb090615388565b60405180910390fd5b60008111611cfc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611cf390615b56565b60405180910390fd5b600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054811315611d7e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d7590615bc2565b60405180910390fd5b80600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611dcd9190615443565b9250508190555080601060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611e239190615486565b92505081905550611e3382613b96565b8173ffffffffffffffffffffffffffffffffffffffff167f5f3d004cf9164b95ed5dbf47d1f04018a4eabcb20b4320fe229ed92236ace63482604051611e7991906155ab565b60405180910390a25050565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b8152600401611ee4939291906157ae565b6020604051808303816000875af1158015611f03573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f2791906156c0565b5080600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611f779190615486565b925050819055508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f6283604051611fdb9190614e9e565b60405180910390a3611fec82613b96565b5050565b611ff8613cf8565b61200182613dde565b61200b8282613de9565b5050565b6000612019613f08565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b905090565b600e5481565b61205061452b565b6040518060800160405280600054815260200160015481526020016002548152602001600354815250905090565b600c6020528060005260406000206000915090505481565b6005818051602081018201805184825260208301602085012081835280955050505050506000915090508060010154908060020160009054906101000a900460ff16905082565b806000808273ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527f893d20e8000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516121889190615898565b6000604051808303816000865af19150503d80600081146121c5576040519150601f19603f3d011682016040523d82523d6000602084013e6121ca565b606091505b50915091508161220f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161220690615921565b60405180910390fd5b600081806020019051810190612225919061597f565b90508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612295576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161228c906159f8565b60405180910390fd5b6000600d60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600081600101541161231f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161231690615c2e565b60405180910390fd5b7fa3895d397a34e928a95d593331e293e2fc281d9d8996df5cc6c57f1cef629d4286826001015460405161235492919061566b565b60405180910390a1600d60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008082016000905560018201600090556002820160006101000a81549060ff02191690555050505050505050565b6123d2613b0f565b80600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60005481565b612424613b0f565b6001600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b612487613b0f565b6124916000613f8f565b565b61249b613b0f565b80600e8190555050565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16612531576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161252890615388565b60405180910390fd5b8060018190555050565b6000612545614066565b905060008160000160089054906101000a900460ff1615905060008260000160009054906101000a900467ffffffffffffffff1690506000808267ffffffffffffffff161480156125935750825b9050600060018367ffffffffffffffff161480156125c8575060003073ffffffffffffffffffffffffffffffffffffffff163b145b9050811580156125d6575080155b1561260d576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018560000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550831561265d5760018560000160086101000a81548160ff0219169083151502179055505b6126663361408e565b61266e6140a2565b6001600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600281905550600160038190555067016345785d8a000060018190555083156127415760008560000160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d260016040516127389190615ca7565b60405180910390a15b5050505050565b6000806127536140ac565b90508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505090565b60076020528060005260406000206000915090505481565b60106020528060005260406000206000915090505481565b60015481565b60066020528060005260406000206000915090505481565b6127d6614553565b60405180608001604052808373ffffffffffffffffffffffffffffffffffffffff168152602001600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548152602001601060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548152602001600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080546128d190615ad9565b80601f01602080910402602001604051908101604052809291908181526020018280546128fd90615ad9565b801561294a5780601f1061291f5761010080835404028352916020019161294a565b820191906000526020600020905b81548152906001019060200180831161292d57829003601f168201915b50505050508152509050919050565b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b60025481565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16612a24576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a1b90615388565b60405180910390fd5b6000600584604051612a369190615506565b90815260200160405180910390209050601060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548213612b36577fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b82604051612abb9190614e9e565b60405180910390a181600080828254612ad4919061551d565b9250508190555081601060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254612b2a9190615443565b92505081905550612f2a565b60008290506000601060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541315612c6d57601060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600080828254612bd3919061551d565b92505081905550601060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481612c259190615443565b90506000601060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b80600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205412612d5d5780600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254612d029190615443565b9250508190555080600080828254612d1a919061551d565b925050819055507fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b83604051612d509190614e9e565b60405180910390a1612f28565b6000600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541315612e8f57600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481612df09190615443565b9050600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600080828254612e42919061551d565b925050819055506000600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b7fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b8184612ebc9190615551565b604051612ec99190614e9e565b60405180910390a180601060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254612f209190615443565b925050819055505b505b60028160020160006101000a81548160ff02191690836002811115612f5257612f51614a45565b5b0217905550612f6083613b96565b8273ffffffffffffffffffffffffffffffffffffffff167f403feb2cd85cc25c910fe59289105b583f08fe9e4335ebbb08c1851f47ff2d348386604051612fa8929190615cc2565b60405180910390a283604051612fbe9190615506565b60405180910390207f4a2ced9ada462e244851a86e998eb0b5bf558c2c9c6923b7f970ed2b19b073b06002604051612ff69190615cf2565b60405180910390a250505050565b61300c614591565b600d60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060405180606001604052908160008201548152602001600182015481526020016002820160009054906101000a900460ff1615151515815250509050919050565b613094613b0f565b8060038190555050565b60035481565b806000808273ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527f893d20e8000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505060405161314f9190615898565b6000604051808303816000865af19150503d806000811461318c576040519150601f19603f3d011682016040523d82523d6000602084013e613191565b606091505b5091509150816131d6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016131cd90615921565b60405180910390fd5b6000818060200190518101906131ec919061597f565b90508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461325c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613253906159f8565b60405180910390fd5b6000600d60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060008160010154116132e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016132dd90615c2e565b60405180910390fd5b600e5481600001546132f8919061551d565b43101561333a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161333190615d59565b60405180910390fd5b8060010154601060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205412156133c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133b790615a8a565b60405180910390fd5b60008160010154905080601060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546134189190615443565b92505081905550600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b815260040161347c92919061566b565b6020604051808303816000875af115801561349b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134bf91906156c0565b507fe6fc6292b9fd5ba92a34a05a92d07b066535578023d71f0c6bf83a2622ca4e548733836040516134f3939291906157ae565b60405180910390a1600d60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008082016000905560018201600090556002820160006101000a81549060ff0219169055505061356a87613b96565b50505050505050565b61357b613b0f565b60005b81518110156136b75781818151811061359a5761359961577f565b5b6020026020010151601060008584815181106135b9576135b861577f565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461360a9190615486565b925050819055508281815181106136245761362361577f565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f18b5d63f2a4f63f9d724d087c9580fdb3c4501bf641ce600afe4b97282e738e284848151811061368d5761368c61577f565b5b60200260200101516040516136a29190614e9e565b60405180910390a3808060010191505061357e565b505050565b6136c4613b0f565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036137365760006040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161372d9190614fa4565b60405180910390fd5b61373f81613f8f565b50565b60008111613785576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161377c90615dc5565b60405180910390fd5b6000600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205413613807576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016137fe90615e31565b60405180910390fd5b600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054811115613889576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161388090615ec3565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1663893d20e86040518163ffffffff1660e01b8152600401602060405180830381865afa1580156138eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061390f9190615ef8565b73ffffffffffffffffffffffffffffffffffffffff1614613965576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161395c90615f97565b60405180910390fd5b80600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546139b49190615443565b92505081905550600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b8152600401613a1892919061566b565b6020604051808303816000875af1158015613a37573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613a5b91906156c0565b50613a6582613b96565b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb83604051613ac29190614e9e565b60405180910390a35050565b600d6020528060005260406000206000915090508060000154908060010154908060020160009054906101000a900460ff16905083565b6000600e54905090565b613b176140d4565b73ffffffffffffffffffffffffffffffffffffffff16613b35612748565b73ffffffffffffffffffffffffffffffffffffffff1614613b9457613b586140d4565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401613b8b9190614fa4565b60405180910390fd5b565b600154600254613ba69190615fb7565b601060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205412613c72576040518060400160405280600981526020017f7a6b41756374696f6e0000000000000000000000000000000000000000000000815250600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209081613c6c919061619b565b50613cf5565b6040518060400160405280600381526020017f4e53430000000000000000000000000000000000000000000000000000000000815250600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209081613cf3919061619b565b505b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161480613da557507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16613d8c6140dc565b73ffffffffffffffffffffffffffffffffffffffff1614155b15613ddc576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b613de6613b0f565b50565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015613e5157506040513d601f19601f82011682018060405250810190613e4e9190616299565b60015b613e9257816040517f4c9c8ce3000000000000000000000000000000000000000000000000000000008152600401613e899190614fa4565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b8114613ef957806040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600401613ef09190614e74565b60405180910390fd5b613f038383614133565b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614613f8d576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6000613f996140ac565b905060008160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050828260000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3505050565b60007ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b6140966141a6565b61409f816141e6565b50565b6140aa6141a6565b565b60007f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300905090565b600033905090565b600061410a7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b61426c565b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61413c82614276565b8173ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a2600081511115614199576141938282614343565b506141a2565b6141a16143c7565b5b5050565b6141ae614404565b6141e4576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6141ee6141a6565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036142605760006040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016142579190614fa4565b60405180910390fd5b61426981613f8f565b50565b6000819050919050565b60008173ffffffffffffffffffffffffffffffffffffffff163b036142d257806040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526004016142c99190614fa4565b60405180910390fd5b806142ff7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b61426c565b60000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60606000808473ffffffffffffffffffffffffffffffffffffffff168460405161436d9190615898565b600060405180830381855af49150503d80600081146143a8576040519150601f19603f3d011682016040523d82523d6000602084013e6143ad565b606091505b50915091506143bd858383614424565b9250505092915050565b6000341115614402576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b600061440e614066565b60000160089054906101000a900460ff16905090565b60608261443957614434826144b3565b6144ab565b60008251148015614461575060008473ffffffffffffffffffffffffffffffffffffffff163b145b156144a357836040517f9996b31500000000000000000000000000000000000000000000000000000000815260040161449a9190614fa4565b60405180910390fd5b8190506144ac565b5b9392505050565b6000815111156144c65780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604051806060016040528060608152602001600081526020016000600281111561452557614524614a45565b5b81525090565b6040518060800160405280600081526020016000815260200160008152602001600081525090565b6040518060800160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001606081525090565b604051806060016040528060008152602001600081526020016000151581525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61461b826145d2565b810181811067ffffffffffffffff8211171561463a576146396145e3565b5b80604052505050565b600061464d6145b4565b90506146598282614612565b919050565b600067ffffffffffffffff821115614679576146786145e3565b5b614682826145d2565b9050602081019050919050565b82818337600083830152505050565b60006146b16146ac8461465e565b614643565b9050828152602081018484840111156146cd576146cc6145cd565b5b6146d884828561468f565b509392505050565b600082601f8301126146f5576146f46145c8565b5b813561470584826020860161469e565b91505092915050565b600060208284031215614724576147236145be565b5b600082013567ffffffffffffffff811115614742576147416145c3565b5b61474e848285016146e0565b91505092915050565b60008115159050919050565b61476c81614757565b82525050565b60006020820190506147876000830184614763565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006147b88261478d565b9050919050565b6147c8816147ad565b81146147d357600080fd5b50565b6000813590506147e5816147bf565b92915050565b600060208284031215614801576148006145be565b5b600061480f848285016147d6565b91505092915050565b6000819050919050565b61482b81614818565b811461483657600080fd5b50565b60008135905061484881614822565b92915050565b60008060408385031215614865576148646145be565b5b6000614873858286016147d6565b925050602061488485828601614839565b9150509250929050565b600067ffffffffffffffff8211156148a9576148a86145e3565b5b6148b2826145d2565b9050602081019050919050565b60006148d26148cd8461488e565b614643565b9050828152602081018484840111156148ee576148ed6145cd565b5b6148f984828561468f565b509392505050565b600082601f830112614916576149156145c8565b5b81356149268482602086016148bf565b91505092915050565b600060208284031215614945576149446145be565b5b600082013567ffffffffffffffff811115614963576149626145c3565b5b61496f84828501614901565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6149ad816147ad565b82525050565b60006149bf83836149a4565b60208301905092915050565b6000602082019050919050565b60006149e382614978565b6149ed8185614983565b93506149f883614994565b8060005b83811015614a29578151614a1088826149b3565b9750614a1b836149cb565b9250506001810190506149fc565b5085935050505092915050565b614a3f81614818565b82525050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60038110614a8557614a84614a45565b5b50565b6000819050614a9682614a74565b919050565b6000614aa682614a88565b9050919050565b614ab681614a9b565b82525050565b60006060830160008301518482036000860152614ad982826149d8565b9150506020830151614aee6020860182614a36565b506040830151614b016040860182614aad565b508091505092915050565b60006020820190508181036000830152614b268184614abc565b905092915050565b600060208284031215614b4457614b436145be565b5b6000614b5284828501614839565b91505092915050565b600067ffffffffffffffff821115614b7657614b756145e3565b5b602082029050602081019050919050565b600080fd5b6000614b9f614b9a84614b5b565b614643565b90508083825260208201905060208402830185811115614bc257614bc1614b87565b5b835b81811015614beb5780614bd788826147d6565b845260208401935050602081019050614bc4565b5050509392505050565b600082601f830112614c0a57614c096145c8565b5b8135614c1a848260208601614b8c565b91505092915050565b600067ffffffffffffffff821115614c3e57614c3d6145e3565b5b602082029050602081019050919050565b6000614c62614c5d84614c23565b614643565b90508083825260208201905060208402830185811115614c8557614c84614b87565b5b835b81811015614cae5780614c9a8882614839565b845260208401935050602081019050614c87565b5050509392505050565b600082601f830112614ccd57614ccc6145c8565b5b8135614cdd848260208601614c4f565b91505092915050565b60008060408385031215614cfd57614cfc6145be565b5b600083013567ffffffffffffffff811115614d1b57614d1a6145c3565b5b614d2785828601614bf5565b925050602083013567ffffffffffffffff811115614d4857614d476145c3565b5b614d5485828601614cb8565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b60005b83811015614d98578082015181840152602081019050614d7d565b60008484015250505050565b6000614daf82614d5e565b614db98185614d69565b9350614dc9818560208601614d7a565b614dd2816145d2565b840191505092915050565b60006020820190508181036000830152614df78184614da4565b905092915050565b60008060408385031215614e1657614e156145be565b5b6000614e24858286016147d6565b925050602083013567ffffffffffffffff811115614e4557614e446145c3565b5b614e51858286016146e0565b9150509250929050565b6000819050919050565b614e6e81614e5b565b82525050565b6000602082019050614e896000830184614e65565b92915050565b614e9881614818565b82525050565b6000602082019050614eb36000830184614e8f565b92915050565b608082016000820151614ecf6000850182614a36565b506020820151614ee26020850182614a36565b506040820151614ef56040850182614a36565b506060820151614f086060850182614a36565b50505050565b6000608082019050614f236000830184614eb9565b92915050565b6000819050919050565b614f3c81614f29565b82525050565b6000602082019050614f576000830184614f33565b92915050565b614f6681614a9b565b82525050565b6000604082019050614f816000830185614e8f565b614f8e6020830184614f5d565b9392505050565b614f9e816147ad565b82525050565b6000602082019050614fb96000830184614f95565b92915050565b614fc881614f29565b82525050565b600082825260208201905092915050565b6000614fea82614d5e565b614ff48185614fce565b9350615004818560208601614d7a565b61500d816145d2565b840191505092915050565b600060808301600083015161503060008601826149a4565b5060208301516150436020860182614fbf565b5060408301516150566040860182614fbf565b506060830151848203606086015261506e8282614fdf565b9150508091505092915050565b600060208201905081810360008301526150958184615018565b905092915050565b6000806000606084860312156150b6576150b56145be565b5b600084013567ffffffffffffffff8111156150d4576150d36145c3565b5b6150e086828701614901565b93505060206150f1868287016147d6565b925050604061510286828701614839565b9150509250925092565b61511581614757565b82525050565b6060820160008201516151316000850182614a36565b5060208201516151446020850182614a36565b506040820151615157604085018261510c565b50505050565b6000606082019050615172600083018461511b565b92915050565b600067ffffffffffffffff821115615193576151926145e3565b5b602082029050602081019050919050565b6151ad81614f29565b81146151b857600080fd5b50565b6000813590506151ca816151a4565b92915050565b60006151e36151de84615178565b614643565b9050808382526020820190506020840283018581111561520657615205614b87565b5b835b8181101561522f578061521b88826151bb565b845260208401935050602081019050615208565b5050509392505050565b600082601f83011261524e5761524d6145c8565b5b813561525e8482602086016151d0565b91505092915050565b6000806040838503121561527e5761527d6145be565b5b600083013567ffffffffffffffff81111561529c5761529b6145c3565b5b6152a885828601614bf5565b925050602083013567ffffffffffffffff8111156152c9576152c86145c3565b5b6152d585828601615239565b9150509250929050565b60006060820190506152f46000830186614e8f565b6153016020830185614e8f565b61530e6040830184614763565b949350505050565b7f4f6e6c79207468652061646d696e2063616e2063616c6c20746869732066756e60008201527f6374696f6e2e0000000000000000000000000000000000000000000000000000602082015250565b6000615372602683614d69565b915061537d82615316565b604082019050919050565b600060208201905081810360008301526153a181615365565b9050919050565b7f6e6f7420656e6f756768206c6f636b65642062616c616e636500000000000000600082015250565b60006153de601983614d69565b91506153e9826153a8565b602082019050919050565b6000602082019050818103600083015261540d816153d1565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061544e82614f29565b915061545983614f29565b92508282039050818112600084121682821360008512151617156154805761547f615414565b5b92915050565b600061549182614f29565b915061549c83614f29565b9250828201905082811215600083121683821260008412151617156154c4576154c3615414565b5b92915050565b600081905092915050565b60006154e082614d5e565b6154ea81856154ca565b93506154fa818560208601614d7a565b80840191505092915050565b600061551282846154d5565b915081905092915050565b600061552882614818565b915061553383614818565b925082820190508082111561554b5761554a615414565b5b92915050565b600061555c82614818565b915061556783614818565b925082820390508181111561557f5761557e615414565b5b92915050565b50565b6000615595600083614d69565b91506155a082615585565b600082019050919050565b60006040820190506155c06000830184614e8f565b81810360208301526155d181615588565b905092915050565b7f576974686472617720616d6f756e7420616d6f756e742065786365656473207360008201527f6c617368656446756e6473000000000000000000000000000000000000000000602082015250565b6000615635602b83614d69565b9150615640826155d9565b604082019050919050565b6000602082019050818103600083015261566481615628565b9050919050565b60006040820190506156806000830185614f95565b61568d6020830184614e8f565b9392505050565b61569d81614757565b81146156a857600080fd5b50565b6000815190506156ba81615694565b92915050565b6000602082840312156156d6576156d56145be565b5b60006156e4848285016156ab565b91505092915050565b7f6370206163636f756e747320616e6420616d6f756e747320617265206469666660008201527f6572656e74206c656e6774687300000000000000000000000000000000000000602082015250565b6000615749602d83614d69565b9150615754826156ed565b604082019050919050565b600060208201905081810360008301526157788161573c565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006060820190506157c36000830186614f95565b6157d06020830185614f95565b6157dd6040830184614e8f565b949350505050565b7f7472616e73666572206661696c65640000000000000000000000000000000000600082015250565b600061581b600f83614d69565b9150615826826157e5565b602082019050919050565b6000602082019050818103600083015261584a8161580e565b9050919050565b600081519050919050565b600081905092915050565b600061587282615851565b61587c818561585c565b935061588c818560208601614d7a565b80840191505092915050565b60006158a48284615867565b915081905092915050565b7f4661696c656420746f2063616c6c206765744f776e65722066756e6374696f6e60008201527f206f662043504163636f756e7400000000000000000000000000000000000000602082015250565b600061590b602d83614d69565b9150615916826158af565b604082019050919050565b6000602082019050818103600083015261593a816158fe565b9050919050565b600061594c8261478d565b9050919050565b61595c81615941565b811461596757600080fd5b50565b60008151905061597981615953565b92915050565b600060208284031215615995576159946145be565b5b60006159a38482850161596a565b91505092915050565b7f63616c6c6572206973206e6f74206370206f776e657200000000000000000000600082015250565b60006159e2601683614d69565b91506159ed826159ac565b602082019050919050565b60006020820190508181036000830152615a11816159d5565b9050919050565b7f4e6f7420656e6f7567682066726f7a656e2062616c616e636520746f2077697460008201527f6864726177000000000000000000000000000000000000000000000000000000602082015250565b6000615a74602583614d69565b9150615a7f82615a18565b604082019050919050565b60006020820190508181036000830152615aa381615a67565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680615af157607f821691505b602082108103615b0457615b03615aaa565b5b50919050565b7f616d6f756e74206d7573742062652067726561746572207468616e2030000000600082015250565b6000615b40601d83614d69565b9150615b4b82615b0a565b602082019050919050565b60006020820190508181036000830152615b6f81615b33565b9050919050565b7f6e6f7420656e6f7567682062616c616e63650000000000000000000000000000600082015250565b6000615bac601283614d69565b9150615bb782615b76565b602082019050919050565b60006020820190508181036000830152615bdb81615b9f565b9050919050565b7f4e6f2070656e64696e6720776974686472617720726571756573740000000000600082015250565b6000615c18601b83614d69565b9150615c2382615be2565b602082019050919050565b60006020820190508181036000830152615c4781615c0b565b9050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b6000819050919050565b6000615c91615c8c615c8784615c4e565b615c6c565b615c58565b9050919050565b615ca181615c76565b82525050565b6000602082019050615cbc6000830184615c98565b92915050565b6000604082019050615cd76000830185614e8f565b8181036020830152615ce98184614da4565b90509392505050565b6000602082019050615d076000830184614f5d565b92915050565b7f57697468647261772064656c6179206e6f742070617373656400000000000000600082015250565b6000615d43601983614d69565b9150615d4e82615d0d565b602082019050919050565b60006020820190508181036000830152615d7281615d36565b9050919050565b7f776974686472617720616d6f756e74206d75737420626520706f736974697665600082015250565b6000615daf602083614d69565b9150615dba82615d79565b602082019050919050565b60006020820190508181036000830152615dde81615da2565b9050919050565b7f6e6f7468696e6720746f20776974686472617700000000000000000000000000600082015250565b6000615e1b601383614d69565b9150615e2682615de5565b602082019050919050565b60006020820190508181036000830152615e4a81615e0e565b9050919050565b7f776974686472617720616d6f756e742067726561746572207468616e2061766160008201527f696c61626c652062616c616e6365000000000000000000000000000000000000602082015250565b6000615ead602e83614d69565b9150615eb882615e51565b604082019050919050565b60006020820190508181036000830152615edc81615ea0565b9050919050565b600081519050615ef2816147bf565b92915050565b600060208284031215615f0e57615f0d6145be565b5b6000615f1c84828501615ee3565b91505092915050565b7f4f6e6c792043504163636f756e74206f776e65722063616e207769746864726160008201527f772074686520636f6c6c61746572616c2066756e647300000000000000000000602082015250565b6000615f81603683614d69565b9150615f8c82615f25565b604082019050919050565b60006020820190508181036000830152615fb081615f74565b9050919050565b6000615fc282614818565b9150615fcd83614818565b9250828202615fdb81614818565b91508282048414831517615ff257615ff1615414565b5b5092915050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b60006008830261605b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261601e565b616065868361601e565b95508019841693508086168417925050509392505050565b600061609861609361608e84614818565b615c6c565b614818565b9050919050565b6000819050919050565b6160b28361607d565b6160c66160be8261609f565b84845461602b565b825550505050565b600090565b6160db6160ce565b6160e68184846160a9565b505050565b5b8181101561610a576160ff6000826160d3565b6001810190506160ec565b5050565b601f82111561614f5761612081615ff9565b6161298461600e565b81016020851015616138578190505b61614c6161448561600e565b8301826160eb565b50505b505050565b600082821c905092915050565b600061617260001984600802616154565b1980831691505092915050565b600061618b8383616161565b9150826002028217905092915050565b6161a482614d5e565b67ffffffffffffffff8111156161bd576161bc6145e3565b5b6161c78254615ad9565b6161d282828561610e565b600060209050601f83116001811461620557600084156161f3578287015190505b6161fd858261617f565b865550616265565b601f19841661621386615ff9565b60005b8281101561623b57848901518255600182019150602085019450602081019050616216565b868310156162585784890151616254601f891682616161565b8355505b6001600288020188555050505b505050505050565b61627681614e5b565b811461628157600080fd5b50565b6000815190506162938161626d565b92915050565b6000602082840312156162af576162ae6145be565b5b60006162bd84828501616284565b9150509291505056fea264697066735822122030318a4fdd25f908efd7bcb1f546473301acf3a42aa0bff53e638642218a599364736f6c63430008190033",
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
// Solidity: function cpInfo(address cpAddress) view returns((address,int256,int256,string))
func (_SwanCreditCollateral *SwanCreditCollateralCaller) CpInfo(opts *bind.CallOpts, cpAddress common.Address) (SwanCreditCollateralCPInfoWithLockedBalance, error) {
	var out []interface{}
	err := _SwanCreditCollateral.contract.Call(opts, &out, "cpInfo", cpAddress)

	if err != nil {
		return *new(SwanCreditCollateralCPInfoWithLockedBalance), err
	}

	out0 := *abi.ConvertType(out[0], new(SwanCreditCollateralCPInfoWithLockedBalance)).(*SwanCreditCollateralCPInfoWithLockedBalance)

	return out0, err

}

// CpInfo is a free data retrieval call binding the contract method 0xa664c216.
//
// Solidity: function cpInfo(address cpAddress) view returns((address,int256,int256,string))
func (_SwanCreditCollateral *SwanCreditCollateralSession) CpInfo(cpAddress common.Address) (SwanCreditCollateralCPInfoWithLockedBalance, error) {
	return _SwanCreditCollateral.Contract.CpInfo(&_SwanCreditCollateral.CallOpts, cpAddress)
}

// CpInfo is a free data retrieval call binding the contract method 0xa664c216.
//
// Solidity: function cpInfo(address cpAddress) view returns((address,int256,int256,string))
func (_SwanCreditCollateral *SwanCreditCollateralCallerSession) CpInfo(cpAddress common.Address) (SwanCreditCollateralCPInfoWithLockedBalance, error) {
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

// GetWithdrawDelay is a free data retrieval call binding the contract method 0xfe3300d0.
//
// Solidity: function getWithdrawDelay() view returns(uint256)
func (_SwanCreditCollateral *SwanCreditCollateralCaller) GetWithdrawDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwanCreditCollateral.contract.Call(opts, &out, "getWithdrawDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawDelay is a free data retrieval call binding the contract method 0xfe3300d0.
//
// Solidity: function getWithdrawDelay() view returns(uint256)
func (_SwanCreditCollateral *SwanCreditCollateralSession) GetWithdrawDelay() (*big.Int, error) {
	return _SwanCreditCollateral.Contract.GetWithdrawDelay(&_SwanCreditCollateral.CallOpts)
}

// GetWithdrawDelay is a free data retrieval call binding the contract method 0xfe3300d0.
//
// Solidity: function getWithdrawDelay() view returns(uint256)
func (_SwanCreditCollateral *SwanCreditCollateralCallerSession) GetWithdrawDelay() (*big.Int, error) {
	return _SwanCreditCollateral.Contract.GetWithdrawDelay(&_SwanCreditCollateral.CallOpts)
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

// ViewWithdrawRequest is a free data retrieval call binding the contract method 0xbede6e31.
//
// Solidity: function viewWithdrawRequest(address cpAccount) view returns((uint256,uint256,bool))
func (_SwanCreditCollateral *SwanCreditCollateralCaller) ViewWithdrawRequest(opts *bind.CallOpts, cpAccount common.Address) (SwanCreditCollateralUnlockRequest, error) {
	var out []interface{}
	err := _SwanCreditCollateral.contract.Call(opts, &out, "viewWithdrawRequest", cpAccount)

	if err != nil {
		return *new(SwanCreditCollateralUnlockRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(SwanCreditCollateralUnlockRequest)).(*SwanCreditCollateralUnlockRequest)

	return out0, err

}

// ViewWithdrawRequest is a free data retrieval call binding the contract method 0xbede6e31.
//
// Solidity: function viewWithdrawRequest(address cpAccount) view returns((uint256,uint256,bool))
func (_SwanCreditCollateral *SwanCreditCollateralSession) ViewWithdrawRequest(cpAccount common.Address) (SwanCreditCollateralUnlockRequest, error) {
	return _SwanCreditCollateral.Contract.ViewWithdrawRequest(&_SwanCreditCollateral.CallOpts, cpAccount)
}

// ViewWithdrawRequest is a free data retrieval call binding the contract method 0xbede6e31.
//
// Solidity: function viewWithdrawRequest(address cpAccount) view returns((uint256,uint256,bool))
func (_SwanCreditCollateral *SwanCreditCollateralCallerSession) ViewWithdrawRequest(cpAccount common.Address) (SwanCreditCollateralUnlockRequest, error) {
	return _SwanCreditCollateral.Contract.ViewWithdrawRequest(&_SwanCreditCollateral.CallOpts, cpAccount)
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

// BatchAddLockedCollateral is a paid mutator transaction binding the contract method 0xd7635815.
//
// Solidity: function batchAddLockedCollateral(address[] cpAccounts, int256[] amounts) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactor) BatchAddLockedCollateral(opts *bind.TransactOpts, cpAccounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.contract.Transact(opts, "batchAddLockedCollateral", cpAccounts, amounts)
}

// BatchAddLockedCollateral is a paid mutator transaction binding the contract method 0xd7635815.
//
// Solidity: function batchAddLockedCollateral(address[] cpAccounts, int256[] amounts) returns()
func (_SwanCreditCollateral *SwanCreditCollateralSession) BatchAddLockedCollateral(cpAccounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.BatchAddLockedCollateral(&_SwanCreditCollateral.TransactOpts, cpAccounts, amounts)
}

// BatchAddLockedCollateral is a paid mutator transaction binding the contract method 0xd7635815.
//
// Solidity: function batchAddLockedCollateral(address[] cpAccounts, int256[] amounts) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactorSession) BatchAddLockedCollateral(cpAccounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.BatchAddLockedCollateral(&_SwanCreditCollateral.TransactOpts, cpAccounts, amounts)
}

// BatchDeposit is a paid mutator transaction binding the contract method 0x30a90736.
//
// Solidity: function batchDeposit(address[] cpAccounts, uint256[] amounts) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactor) BatchDeposit(opts *bind.TransactOpts, cpAccounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.contract.Transact(opts, "batchDeposit", cpAccounts, amounts)
}

// BatchDeposit is a paid mutator transaction binding the contract method 0x30a90736.
//
// Solidity: function batchDeposit(address[] cpAccounts, uint256[] amounts) returns()
func (_SwanCreditCollateral *SwanCreditCollateralSession) BatchDeposit(cpAccounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.BatchDeposit(&_SwanCreditCollateral.TransactOpts, cpAccounts, amounts)
}

// BatchDeposit is a paid mutator transaction binding the contract method 0x30a90736.
//
// Solidity: function batchDeposit(address[] cpAccounts, uint256[] amounts) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactorSession) BatchDeposit(cpAccounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.BatchDeposit(&_SwanCreditCollateral.TransactOpts, cpAccounts, amounts)
}

// CancelWithdrawRequest is a paid mutator transaction binding the contract method 0x5d2cd2a7.
//
// Solidity: function cancelWithdrawRequest(address cpAccount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactor) CancelWithdrawRequest(opts *bind.TransactOpts, cpAccount common.Address) (*types.Transaction, error) {
	return _SwanCreditCollateral.contract.Transact(opts, "cancelWithdrawRequest", cpAccount)
}

// CancelWithdrawRequest is a paid mutator transaction binding the contract method 0x5d2cd2a7.
//
// Solidity: function cancelWithdrawRequest(address cpAccount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralSession) CancelWithdrawRequest(cpAccount common.Address) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.CancelWithdrawRequest(&_SwanCreditCollateral.TransactOpts, cpAccount)
}

// CancelWithdrawRequest is a paid mutator transaction binding the contract method 0x5d2cd2a7.
//
// Solidity: function cancelWithdrawRequest(address cpAccount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactorSession) CancelWithdrawRequest(cpAccount common.Address) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.CancelWithdrawRequest(&_SwanCreditCollateral.TransactOpts, cpAccount)
}

// ConfirmWithdraw is a paid mutator transaction binding the contract method 0xd2bfc1c7.
//
// Solidity: function confirmWithdraw(address cpAccount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactor) ConfirmWithdraw(opts *bind.TransactOpts, cpAccount common.Address) (*types.Transaction, error) {
	return _SwanCreditCollateral.contract.Transact(opts, "confirmWithdraw", cpAccount)
}

// ConfirmWithdraw is a paid mutator transaction binding the contract method 0xd2bfc1c7.
//
// Solidity: function confirmWithdraw(address cpAccount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralSession) ConfirmWithdraw(cpAccount common.Address) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.ConfirmWithdraw(&_SwanCreditCollateral.TransactOpts, cpAccount)
}

// ConfirmWithdraw is a paid mutator transaction binding the contract method 0xd2bfc1c7.
//
// Solidity: function confirmWithdraw(address cpAccount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactorSession) ConfirmWithdraw(cpAccount common.Address) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.ConfirmWithdraw(&_SwanCreditCollateral.TransactOpts, cpAccount)
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

// LockCollateral is a paid mutator transaction binding the contract method 0x47a7d107.
//
// Solidity: function lockCollateral(address cpAccount, uint256 amount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactor) LockCollateral(opts *bind.TransactOpts, cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.contract.Transact(opts, "lockCollateral", cpAccount, amount)
}

// LockCollateral is a paid mutator transaction binding the contract method 0x47a7d107.
//
// Solidity: function lockCollateral(address cpAccount, uint256 amount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralSession) LockCollateral(cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.LockCollateral(&_SwanCreditCollateral.TransactOpts, cpAccount, amount)
}

// LockCollateral is a paid mutator transaction binding the contract method 0x47a7d107.
//
// Solidity: function lockCollateral(address cpAccount, uint256 amount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactorSession) LockCollateral(cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.LockCollateral(&_SwanCreditCollateral.TransactOpts, cpAccount, amount)
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

// RequestWithdraw is a paid mutator transaction binding the contract method 0x397a1b28.
//
// Solidity: function requestWithdraw(address cpAccount, uint256 amount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactor) RequestWithdraw(opts *bind.TransactOpts, cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.contract.Transact(opts, "requestWithdraw", cpAccount, amount)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x397a1b28.
//
// Solidity: function requestWithdraw(address cpAccount, uint256 amount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralSession) RequestWithdraw(cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.RequestWithdraw(&_SwanCreditCollateral.TransactOpts, cpAccount, amount)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x397a1b28.
//
// Solidity: function requestWithdraw(address cpAccount, uint256 amount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactorSession) RequestWithdraw(cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.RequestWithdraw(&_SwanCreditCollateral.TransactOpts, cpAccount, amount)
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

// SetWithdrawDelay is a paid mutator transaction binding the contract method 0x72f0cb30.
//
// Solidity: function setWithdrawDelay(uint256 _withdrawDelay) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactor) SetWithdrawDelay(opts *bind.TransactOpts, _withdrawDelay *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.contract.Transact(opts, "setWithdrawDelay", _withdrawDelay)
}

// SetWithdrawDelay is a paid mutator transaction binding the contract method 0x72f0cb30.
//
// Solidity: function setWithdrawDelay(uint256 _withdrawDelay) returns()
func (_SwanCreditCollateral *SwanCreditCollateralSession) SetWithdrawDelay(_withdrawDelay *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.SetWithdrawDelay(&_SwanCreditCollateral.TransactOpts, _withdrawDelay)
}

// SetWithdrawDelay is a paid mutator transaction binding the contract method 0x72f0cb30.
//
// Solidity: function setWithdrawDelay(uint256 _withdrawDelay) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactorSession) SetWithdrawDelay(_withdrawDelay *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.SetWithdrawDelay(&_SwanCreditCollateral.TransactOpts, _withdrawDelay)
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

// UnlockCollateral is a paid mutator transaction binding the contract method 0x17f17498.
//
// Solidity: function unlockCollateral(address cp, uint256 taskCollateral) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactor) UnlockCollateral(opts *bind.TransactOpts, cp common.Address, taskCollateral *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.contract.Transact(opts, "unlockCollateral", cp, taskCollateral)
}

// UnlockCollateral is a paid mutator transaction binding the contract method 0x17f17498.
//
// Solidity: function unlockCollateral(address cp, uint256 taskCollateral) returns()
func (_SwanCreditCollateral *SwanCreditCollateralSession) UnlockCollateral(cp common.Address, taskCollateral *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.UnlockCollateral(&_SwanCreditCollateral.TransactOpts, cp, taskCollateral)
}

// UnlockCollateral is a paid mutator transaction binding the contract method 0x17f17498.
//
// Solidity: function unlockCollateral(address cp, uint256 taskCollateral) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactorSession) UnlockCollateral(cp common.Address, taskCollateral *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.UnlockCollateral(&_SwanCreditCollateral.TransactOpts, cp, taskCollateral)
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

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address cpAccount, uint256 withdrawAmount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactor) Withdraw(opts *bind.TransactOpts, cpAccount common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.contract.Transact(opts, "withdraw", cpAccount, withdrawAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address cpAccount, uint256 withdrawAmount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralSession) Withdraw(cpAccount common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.Withdraw(&_SwanCreditCollateral.TransactOpts, cpAccount, withdrawAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address cpAccount, uint256 withdrawAmount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactorSession) Withdraw(cpAccount common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.Withdraw(&_SwanCreditCollateral.TransactOpts, cpAccount, withdrawAmount)
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
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCollateralUnlocked is a free log retrieval operation binding the contract event 0xb4eaf47ecd4bc76248f192433e8067c96cb3e17aced42fbc47a512742fb74216.
//
// Solidity: event CollateralUnlocked(address indexed cp, uint256 collateralAmount)
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

// WatchCollateralUnlocked is a free log subscription operation binding the contract event 0xb4eaf47ecd4bc76248f192433e8067c96cb3e17aced42fbc47a512742fb74216.
//
// Solidity: event CollateralUnlocked(address indexed cp, uint256 collateralAmount)
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

// ParseCollateralUnlocked is a log parse operation binding the contract event 0xb4eaf47ecd4bc76248f192433e8067c96cb3e17aced42fbc47a512742fb74216.
//
// Solidity: event CollateralUnlocked(address indexed cp, uint256 collateralAmount)
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

// SwanCreditCollateralDepositLockedIterator is returned from FilterDepositLocked and is used to iterate over the raw logs and unpacked data for DepositLocked events raised by the SwanCreditCollateral contract.
type SwanCreditCollateralDepositLockedIterator struct {
	Event *SwanCreditCollateralDepositLocked // Event containing the contract specifics and raw log

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
func (it *SwanCreditCollateralDepositLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwanCreditCollateralDepositLocked)
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
		it.Event = new(SwanCreditCollateralDepositLocked)
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
func (it *SwanCreditCollateralDepositLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwanCreditCollateralDepositLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwanCreditCollateralDepositLocked represents a DepositLocked event raised by the SwanCreditCollateral contract.
type SwanCreditCollateralDepositLocked struct {
	FundingWallet common.Address
	CpAccount     common.Address
	DepositAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDepositLocked is a free log retrieval operation binding the contract event 0x18b5d63f2a4f63f9d724d087c9580fdb3c4501bf641ce600afe4b97282e738e2.
//
// Solidity: event DepositLocked(address indexed fundingWallet, address indexed cpAccount, uint256 depositAmount)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) FilterDepositLocked(opts *bind.FilterOpts, fundingWallet []common.Address, cpAccount []common.Address) (*SwanCreditCollateralDepositLockedIterator, error) {

	var fundingWalletRule []interface{}
	for _, fundingWalletItem := range fundingWallet {
		fundingWalletRule = append(fundingWalletRule, fundingWalletItem)
	}
	var cpAccountRule []interface{}
	for _, cpAccountItem := range cpAccount {
		cpAccountRule = append(cpAccountRule, cpAccountItem)
	}

	logs, sub, err := _SwanCreditCollateral.contract.FilterLogs(opts, "DepositLocked", fundingWalletRule, cpAccountRule)
	if err != nil {
		return nil, err
	}
	return &SwanCreditCollateralDepositLockedIterator{contract: _SwanCreditCollateral.contract, event: "DepositLocked", logs: logs, sub: sub}, nil
}

// WatchDepositLocked is a free log subscription operation binding the contract event 0x18b5d63f2a4f63f9d724d087c9580fdb3c4501bf641ce600afe4b97282e738e2.
//
// Solidity: event DepositLocked(address indexed fundingWallet, address indexed cpAccount, uint256 depositAmount)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) WatchDepositLocked(opts *bind.WatchOpts, sink chan<- *SwanCreditCollateralDepositLocked, fundingWallet []common.Address, cpAccount []common.Address) (event.Subscription, error) {

	var fundingWalletRule []interface{}
	for _, fundingWalletItem := range fundingWallet {
		fundingWalletRule = append(fundingWalletRule, fundingWalletItem)
	}
	var cpAccountRule []interface{}
	for _, cpAccountItem := range cpAccount {
		cpAccountRule = append(cpAccountRule, cpAccountItem)
	}

	logs, sub, err := _SwanCreditCollateral.contract.WatchLogs(opts, "DepositLocked", fundingWalletRule, cpAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwanCreditCollateralDepositLocked)
				if err := _SwanCreditCollateral.contract.UnpackLog(event, "DepositLocked", log); err != nil {
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

// ParseDepositLocked is a log parse operation binding the contract event 0x18b5d63f2a4f63f9d724d087c9580fdb3c4501bf641ce600afe4b97282e738e2.
//
// Solidity: event DepositLocked(address indexed fundingWallet, address indexed cpAccount, uint256 depositAmount)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) ParseDepositLocked(log types.Log) (*SwanCreditCollateralDepositLocked, error) {
	event := new(SwanCreditCollateralDepositLocked)
	if err := _SwanCreditCollateral.contract.UnpackLog(event, "DepositLocked", log); err != nil {
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

// SwanCreditCollateralWithdrawConfirmedIterator is returned from FilterWithdrawConfirmed and is used to iterate over the raw logs and unpacked data for WithdrawConfirmed events raised by the SwanCreditCollateral contract.
type SwanCreditCollateralWithdrawConfirmedIterator struct {
	Event *SwanCreditCollateralWithdrawConfirmed // Event containing the contract specifics and raw log

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
func (it *SwanCreditCollateralWithdrawConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwanCreditCollateralWithdrawConfirmed)
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
		it.Event = new(SwanCreditCollateralWithdrawConfirmed)
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
func (it *SwanCreditCollateralWithdrawConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwanCreditCollateralWithdrawConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwanCreditCollateralWithdrawConfirmed represents a WithdrawConfirmed event raised by the SwanCreditCollateral contract.
type SwanCreditCollateralWithdrawConfirmed struct {
	CpAccount      common.Address
	CpOwner        common.Address
	WithdrawAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawConfirmed is a free log retrieval operation binding the contract event 0xe6fc6292b9fd5ba92a34a05a92d07b066535578023d71f0c6bf83a2622ca4e54.
//
// Solidity: event WithdrawConfirmed(address cpAccount, address cpOwner, uint256 withdrawAmount)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) FilterWithdrawConfirmed(opts *bind.FilterOpts) (*SwanCreditCollateralWithdrawConfirmedIterator, error) {

	logs, sub, err := _SwanCreditCollateral.contract.FilterLogs(opts, "WithdrawConfirmed")
	if err != nil {
		return nil, err
	}
	return &SwanCreditCollateralWithdrawConfirmedIterator{contract: _SwanCreditCollateral.contract, event: "WithdrawConfirmed", logs: logs, sub: sub}, nil
}

// WatchWithdrawConfirmed is a free log subscription operation binding the contract event 0xe6fc6292b9fd5ba92a34a05a92d07b066535578023d71f0c6bf83a2622ca4e54.
//
// Solidity: event WithdrawConfirmed(address cpAccount, address cpOwner, uint256 withdrawAmount)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) WatchWithdrawConfirmed(opts *bind.WatchOpts, sink chan<- *SwanCreditCollateralWithdrawConfirmed) (event.Subscription, error) {

	logs, sub, err := _SwanCreditCollateral.contract.WatchLogs(opts, "WithdrawConfirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwanCreditCollateralWithdrawConfirmed)
				if err := _SwanCreditCollateral.contract.UnpackLog(event, "WithdrawConfirmed", log); err != nil {
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

// ParseWithdrawConfirmed is a log parse operation binding the contract event 0xe6fc6292b9fd5ba92a34a05a92d07b066535578023d71f0c6bf83a2622ca4e54.
//
// Solidity: event WithdrawConfirmed(address cpAccount, address cpOwner, uint256 withdrawAmount)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) ParseWithdrawConfirmed(log types.Log) (*SwanCreditCollateralWithdrawConfirmed, error) {
	event := new(SwanCreditCollateralWithdrawConfirmed)
	if err := _SwanCreditCollateral.contract.UnpackLog(event, "WithdrawConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwanCreditCollateralWithdrawRequestCanceledIterator is returned from FilterWithdrawRequestCanceled and is used to iterate over the raw logs and unpacked data for WithdrawRequestCanceled events raised by the SwanCreditCollateral contract.
type SwanCreditCollateralWithdrawRequestCanceledIterator struct {
	Event *SwanCreditCollateralWithdrawRequestCanceled // Event containing the contract specifics and raw log

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
func (it *SwanCreditCollateralWithdrawRequestCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwanCreditCollateralWithdrawRequestCanceled)
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
		it.Event = new(SwanCreditCollateralWithdrawRequestCanceled)
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
func (it *SwanCreditCollateralWithdrawRequestCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwanCreditCollateralWithdrawRequestCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwanCreditCollateralWithdrawRequestCanceled represents a WithdrawRequestCanceled event raised by the SwanCreditCollateral contract.
type SwanCreditCollateralWithdrawRequestCanceled struct {
	CpAccount      common.Address
	WithdrawAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawRequestCanceled is a free log retrieval operation binding the contract event 0xa3895d397a34e928a95d593331e293e2fc281d9d8996df5cc6c57f1cef629d42.
//
// Solidity: event WithdrawRequestCanceled(address cpAccount, uint256 withdrawAmount)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) FilterWithdrawRequestCanceled(opts *bind.FilterOpts) (*SwanCreditCollateralWithdrawRequestCanceledIterator, error) {

	logs, sub, err := _SwanCreditCollateral.contract.FilterLogs(opts, "WithdrawRequestCanceled")
	if err != nil {
		return nil, err
	}
	return &SwanCreditCollateralWithdrawRequestCanceledIterator{contract: _SwanCreditCollateral.contract, event: "WithdrawRequestCanceled", logs: logs, sub: sub}, nil
}

// WatchWithdrawRequestCanceled is a free log subscription operation binding the contract event 0xa3895d397a34e928a95d593331e293e2fc281d9d8996df5cc6c57f1cef629d42.
//
// Solidity: event WithdrawRequestCanceled(address cpAccount, uint256 withdrawAmount)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) WatchWithdrawRequestCanceled(opts *bind.WatchOpts, sink chan<- *SwanCreditCollateralWithdrawRequestCanceled) (event.Subscription, error) {

	logs, sub, err := _SwanCreditCollateral.contract.WatchLogs(opts, "WithdrawRequestCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwanCreditCollateralWithdrawRequestCanceled)
				if err := _SwanCreditCollateral.contract.UnpackLog(event, "WithdrawRequestCanceled", log); err != nil {
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

// ParseWithdrawRequestCanceled is a log parse operation binding the contract event 0xa3895d397a34e928a95d593331e293e2fc281d9d8996df5cc6c57f1cef629d42.
//
// Solidity: event WithdrawRequestCanceled(address cpAccount, uint256 withdrawAmount)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) ParseWithdrawRequestCanceled(log types.Log) (*SwanCreditCollateralWithdrawRequestCanceled, error) {
	event := new(SwanCreditCollateralWithdrawRequestCanceled)
	if err := _SwanCreditCollateral.contract.UnpackLog(event, "WithdrawRequestCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwanCreditCollateralWithdrawRequestedIterator is returned from FilterWithdrawRequested and is used to iterate over the raw logs and unpacked data for WithdrawRequested events raised by the SwanCreditCollateral contract.
type SwanCreditCollateralWithdrawRequestedIterator struct {
	Event *SwanCreditCollateralWithdrawRequested // Event containing the contract specifics and raw log

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
func (it *SwanCreditCollateralWithdrawRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwanCreditCollateralWithdrawRequested)
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
		it.Event = new(SwanCreditCollateralWithdrawRequested)
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
func (it *SwanCreditCollateralWithdrawRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwanCreditCollateralWithdrawRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwanCreditCollateralWithdrawRequested represents a WithdrawRequested event raised by the SwanCreditCollateral contract.
type SwanCreditCollateralWithdrawRequested struct {
	CpAccount      common.Address
	WithdrawAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawRequested is a free log retrieval operation binding the contract event 0xf7774b688d56120b783560a913ee60792a73dfd511812b7be5eccf10d08c6689.
//
// Solidity: event WithdrawRequested(address cpAccount, uint256 withdrawAmount)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) FilterWithdrawRequested(opts *bind.FilterOpts) (*SwanCreditCollateralWithdrawRequestedIterator, error) {

	logs, sub, err := _SwanCreditCollateral.contract.FilterLogs(opts, "WithdrawRequested")
	if err != nil {
		return nil, err
	}
	return &SwanCreditCollateralWithdrawRequestedIterator{contract: _SwanCreditCollateral.contract, event: "WithdrawRequested", logs: logs, sub: sub}, nil
}

// WatchWithdrawRequested is a free log subscription operation binding the contract event 0xf7774b688d56120b783560a913ee60792a73dfd511812b7be5eccf10d08c6689.
//
// Solidity: event WithdrawRequested(address cpAccount, uint256 withdrawAmount)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) WatchWithdrawRequested(opts *bind.WatchOpts, sink chan<- *SwanCreditCollateralWithdrawRequested) (event.Subscription, error) {

	logs, sub, err := _SwanCreditCollateral.contract.WatchLogs(opts, "WithdrawRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwanCreditCollateralWithdrawRequested)
				if err := _SwanCreditCollateral.contract.UnpackLog(event, "WithdrawRequested", log); err != nil {
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

// ParseWithdrawRequested is a log parse operation binding the contract event 0xf7774b688d56120b783560a913ee60792a73dfd511812b7be5eccf10d08c6689.
//
// Solidity: event WithdrawRequested(address cpAccount, uint256 withdrawAmount)
func (_SwanCreditCollateral *SwanCreditCollateralFilterer) ParseWithdrawRequested(log types.Log) (*SwanCreditCollateralWithdrawRequested, error) {
	event := new(SwanCreditCollateralWithdrawRequested)
	if err := _SwanCreditCollateral.contract.UnpackLog(event, "WithdrawRequested", log); err != nil {
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
