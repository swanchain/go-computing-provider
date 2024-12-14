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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"frozenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"operation\",\"type\":\"string\"}],\"name\":\"CollateralAdjusted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"CollateralLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"CollateralSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumSwanCreditCollateral.CollateralStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"CollateralStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"CollateralUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fundingWallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fundingWallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"}],\"name\":\"DepositLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"DisputeProof\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SlashedFundsIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cpAccountAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cpOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"WithdrawConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"WithdrawRequestCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"WithdrawRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralContratOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashfund\",\"type\":\"uint256\"}],\"name\":\"WithdrawSlash\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedToWithdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"availableBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"cpAccounts\",\"type\":\"address[]\"},{\"internalType\":\"int256[]\",\"name\":\"amounts\",\"type\":\"int256[]\"}],\"name\":\"batchAddLockedCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"cpAccounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"}],\"name\":\"cancelWithdrawRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"slashedFunds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashRatio\",\"type\":\"uint256\"}],\"internalType\":\"structSwanCreditCollateral.ContractInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"}],\"name\":\"confirmWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAddress\",\"type\":\"address\"}],\"name\":\"cpInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"availableBalance\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"lockedBalance\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"}],\"internalType\":\"structSwanCreditCollateral.CPInfoWithLockedBalance\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cpStatus\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"getTaskInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"cpList\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"enumSwanCreditCollateral.CollateralStatus\",\"name\":\"collateralStatus\",\"type\":\"uint8\"}],\"internalType\":\"structSwanCreditCollateral.Task\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWithdrawDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"isSignatureUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lockCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockedBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockedCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"requestWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseCollateral\",\"type\":\"uint256\"}],\"name\":\"setBaseCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"setCollateralToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slashRatio\",\"type\":\"uint256\"}],\"name\":\"setSlashRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawDelay\",\"type\":\"uint256\"}],\"name\":\"setWithdrawDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slashAmount\",\"type\":\"uint256\"}],\"name\":\"slashCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slashAmount\",\"type\":\"uint256\"}],\"name\":\"slashCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashedFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"enumSwanCreditCollateral.CollateralStatus\",\"name\":\"collateralStatus\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"taskCollateral\",\"type\":\"uint256\"}],\"name\":\"unlockCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"unlockRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPending\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"}],\"name\":\"viewWithdrawRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requestTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPending\",\"type\":\"bool\"}],\"internalType\":\"structSwanCreditCollateral.UnlockRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawSlashedFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801561004357600080fd5b5061005261005760201b60201c565b6101c1565b600061006761015b60201b60201c565b90508060000160089054906101000a900460ff16156100b2576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff80168160000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff16146101585767ffffffffffffffff8160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d267ffffffffffffffff60405161014f91906101a6565b60405180910390a15b50565b60007ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b600067ffffffffffffffff82169050919050565b6101a081610183565b82525050565b60006020820190506101bb6000830184610197565b92915050565b60805161629c6101ea60003960008181613c9a01528181613cef0152613eaa015261629c6000f3fe6080604052600436106102675760003560e01c8063715018a611610144578063b4eae1cb116100b6578063d2bfc1c71161007a578063d2bfc1c714610940578063d763581514610969578063f2fde38b14610992578063f3fef3a3146109bb578063f97348e1146109e4578063fe3300d014610a2357610267565b8063b4eae1cb1461085b578063b587b82c14610886578063bede6e31146108af578063ce3518aa146108ec578063d27ca89b1461091557610267565b806392bdf9ba1161010857806392bdf9ba146107115780639ae697bf1461074e5780639b5ddf091461078b578063a0821be3146107b6578063a664c216146107f3578063ad3cb1cc1461083057610267565b8063715018a61461066657806372f0cb301461067d5780637f58a7e5146106a65780638129fc1c146106cf5780638da5cb5b146106e657610267565b806347a7d107116101dd57806353ad8720116101a157806353ad87201461055757806358709cf2146105825780635d2cd2a7146105c0578063666181a9146105e95780636f99f15c14610612578063704802751461063d57610267565b806347a7d1071461049357806347e7ef24146104bc5780634f1ef286146104e557806352d1902d14610501578063536f60701461052c57610267565b806324d7806c1161022f57806324d7806c146103615780632894493f1461039e5780632d291cad146103c757806330a9073614610404578063397a1b281461042d5780633fe651771461045657610267565b80631150f0f31461026c5780631785f53c146102a957806317f17498146102d25780631b209463146102fb5780631d47a62d14610338575b600080fd5b34801561027857600080fd5b50610293600480360381019061028e91906146ae565b610a4e565b6040516102a09190614712565b60405180910390f35b3480156102b557600080fd5b506102d060048036038101906102cb919061478b565b610a84565b005b3480156102de57600080fd5b506102f960048036038101906102f491906147ee565b610ae7565b005b34801561030757600080fd5b50610322600480360381019061031d91906148cf565b610d16565b60405161032f9190614aac565b60405180910390f35b34801561034457600080fd5b5061035f600480360381019061035a91906147ee565b610e21565b005b34801561036d57600080fd5b506103886004803603810190610383919061478b565b6113ec565b6040516103959190614712565b60405180910390f35b3480156103aa57600080fd5b506103c560048036038101906103c09190614ace565b61140c565b005b3480156103d357600080fd5b506103ee60048036038101906103e9919061478b565b611563565b6040516103fb9190614712565b60405180910390f35b34801561041057600080fd5b5061042b60048036038101906104269190614c86565b611583565b005b34801561043957600080fd5b50610454600480360381019061044f91906147ee565b611836565b005b34801561046257600080fd5b5061047d6004803603810190610478919061478b565b611b45565b60405161048a9190614d7d565b60405180910390f35b34801561049f57600080fd5b506104ba60048036038101906104b591906147ee565b611be5565b005b3480156104c857600080fd5b506104e360048036038101906104de91906147ee565b611e3d565b005b6104ff60048036038101906104fa9190614d9f565b611fa8565b005b34801561050d57600080fd5b50610516611fc7565b6040516105239190614e14565b60405180910390f35b34801561053857600080fd5b50610541611ffa565b60405161054e9190614e3e565b60405180910390f35b34801561056357600080fd5b5061056c612000565b6040516105799190614eae565b60405180910390f35b34801561058e57600080fd5b506105a960048036038101906105a491906148cf565b612036565b6040516105b7929190614ed8565b60405180910390f35b3480156105cc57600080fd5b506105e760048036038101906105e2919061478b565b61207d565b005b3480156105f557600080fd5b50610610600480360381019061060b919061478b565b61236a565b005b34801561061e57600080fd5b506106276123b6565b6040516106349190614e3e565b60405180910390f35b34801561064957600080fd5b50610664600480360381019061065f919061478b565b6123bc565b005b34801561067257600080fd5b5061067b61241f565b005b34801561068957600080fd5b506106a4600480360381019061069f9190614ace565b612433565b005b3480156106b257600080fd5b506106cd60048036038101906106c89190614ace565b612445565b005b3480156106db57600080fd5b506106e46124db565b005b3480156106f257600080fd5b506106fb6126e8565b6040516107089190614f10565b60405180910390f35b34801561071d57600080fd5b506107386004803603810190610733919061478b565b612720565b6040516107459190614e3e565b60405180910390f35b34801561075a57600080fd5b506107756004803603810190610770919061478b565b612738565b6040516107829190614f44565b60405180910390f35b34801561079757600080fd5b506107a0612750565b6040516107ad9190614e3e565b60405180910390f35b3480156107c257600080fd5b506107dd60048036038101906107d8919061478b565b612756565b6040516107ea9190614f44565b60405180910390f35b3480156107ff57600080fd5b5061081a6004803603810190610815919061478b565b61276e565b604051610827919061501b565b60405180910390f35b34801561083c57600080fd5b506108456128f9565b6040516108529190614d7d565b60405180910390f35b34801561086757600080fd5b50610870612932565b60405161087d9190614e3e565b60405180910390f35b34801561089257600080fd5b506108ad60048036038101906108a8919061503d565b612938565b005b3480156108bb57600080fd5b506108d660048036038101906108d1919061478b565b612fa4565b6040516108e391906150fd565b60405180910390f35b3480156108f857600080fd5b50610913600480360381019061090e9190614ace565b61302c565b005b34801561092157600080fd5b5061092a61303e565b6040516109379190614e3e565b60405180910390f35b34801561094c57600080fd5b506109676004803603810190610962919061478b565b613044565b005b34801561097557600080fd5b50610990600480360381019061098b9190615207565b613513565b005b34801561099e57600080fd5b506109b960048036038101906109b4919061478b565b61365c565b005b3480156109c757600080fd5b506109e260048036038101906109dd91906147ee565b6136e2565b005b3480156109f057600080fd5b50610a0b6004803603810190610a06919061478b565b613a6e565b604051610a1a9392919061527f565b60405180910390f35b348015610a2f57600080fd5b50610a38613aa5565b604051610a459190614e3e565b60405180910390f35b600b818051602081018201805184825260208301602085012081835280955050505050506000915054906101000a900460ff1681565b610a8c613aaf565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610b73576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b6a90615328565b60405180910390fd5b6000600d60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060008113610bfa576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bf190615394565b60405180910390fd5b6000818313610c095782610c0b565b815b905080600d60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610c5c91906153e3565b9250508190555080600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610cb29190615426565b92505081905550610cc284613b36565b8373ffffffffffffffffffffffffffffffffffffffff167fb4eaf47ecd4bc76248f192433e8067c96cb3e17aced42fbc47a512742fb7421682604051610d089190614e3e565b60405180910390a250505050565b610d1e614498565b600582604051610d2e91906154a6565b908152602001604051809103902060405180606001604052908160008201805480602002602001604051908101604052809291908181526020018280548015610dcc57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610d82575b50505050508152602001600182015481526020016002820160009054906101000a900460ff166002811115610e0457610e036149e5565b5b6002811115610e1657610e156149e5565b5b815250509050919050565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610ead576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ea490615328565b60405180910390fd5b600d60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548113610f9d577fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b81604051610f229190614e3e565b60405180910390a180600080828254610f3b91906154bd565b9250508190555080600d60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610f9191906153e3565b92505081905550611391565b60008190506000600d60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205413156110d457600d60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205460008082825461103a91906154bd565b92505081905550600d60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548161108c91906153e3565b90506000600d60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b80600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054126111c45780600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461116991906153e3565b925050819055508060008082825461118191906154bd565b925050819055507fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b826040516111b79190614e3e565b60405180910390a161138f565b6000600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205413156112f657600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548161125791906153e3565b9050600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546000808282546112a991906154bd565b925050819055506000600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b7fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b818361132391906154f1565b6040516113309190614e3e565b60405180910390a180600d60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461138791906153e3565b925050819055505b505b61139a82613b36565b8173ffffffffffffffffffffffffffffffffffffffff167f403feb2cd85cc25c910fe59289105b583f08fe9e4335ebbb08c1851f47ff2d34826040516113e0919061554b565b60405180910390a25050565b60046020528060005260406000206000915054906101000a900460ff1681565b611414613aaf565b806000541015611459576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611450906155eb565b60405180910390fd5b8060008082825461146a91906154f1565b92505081905550600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b81526004016114ce92919061560b565b6020604051808303816000875af11580156114ed573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115119190615660565b503373ffffffffffffffffffffffffffffffffffffffff167fbfd9c82485e2178fcfd5c40379d8e66fe60acc50aa1ef4c50966431eb1e381cd826040516115589190614e3e565b60405180910390a250565b600a6020528060005260406000206000915054906101000a900460ff1681565b80518251146115c7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115be906156ff565b60405180910390fd5b6000805b825181101561174f578281815181106115e7576115e661571f565b5b6020026020010151826115fa91906154bd565b915082818151811061160f5761160e61571f565b5b60200260200101516006600086848151811061162e5761162d61571f565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461167f9190615426565b925050819055506116a984828151811061169c5761169b61571f565b5b6020026020010151613b36565b8381815181106116bc576116bb61571f565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f628584815181106117255761172461571f565b5b602002602001015160405161173a9190614e3e565b60405180910390a380806001019150506115cb565b50600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b81526004016117af9392919061574e565b6020604051808303816000875af11580156117ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117f29190615660565b611831576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611828906157d1565b60405180910390fd5b505050565b816000808273ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527f893d20e8000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516118e19190615838565b6000604051808303816000865af19150503d806000811461191e576040519150601f19603f3d011682016040523d82523d6000602084013e611923565b606091505b509150915081611968576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161195f906158c1565b60405180910390fd5b60008180602001905181019061197e919061591f565b90508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146119ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119e590615998565b60405180910390fd5b84600d60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541215611a70576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a6790615a2a565b60405180910390fd5b604051806060016040528043815260200186815260200160011515815250600c60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548160ff0219169083151502179055509050507ff7774b688d56120b783560a913ee60792a73dfd511812b7be5eccf10d08c66898686604051611b3592919061560b565b60405180910390a1505050505050565b60086020528060005260406000206000915090508054611b6490615a79565b80601f0160208091040260200160405190810160405280929190818152602001828054611b9090615a79565b8015611bdd5780601f10611bb257610100808354040283529160200191611bdd565b820191906000526020600020905b815481529060010190602001808311611bc057829003601f168201915b505050505081565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611c71576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c6890615328565b60405180910390fd5b60008111611cb4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611cab90615af6565b60405180910390fd5b600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054811315611d36576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d2d90615b62565b60405180910390fd5b80600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611d8591906153e3565b9250508190555080600d60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611ddb9190615426565b92505081905550611deb82613b36565b8173ffffffffffffffffffffffffffffffffffffffff167f5f3d004cf9164b95ed5dbf47d1f04018a4eabcb20b4320fe229ed92236ace63482604051611e31919061554b565b60405180910390a25050565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b8152600401611e9c9392919061574e565b6020604051808303816000875af1158015611ebb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611edf9190615660565b5080600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611f2f9190615426565b925050819055508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f6283604051611f939190614e3e565b60405180910390a3611fa482613b36565b5050565b611fb0613c98565b611fb982613d7e565b611fc38282613d89565b5050565b6000611fd1613ea8565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b905090565b600e5481565b6120086144cb565b6040518060800160405280600054815260200160015481526020016002548152602001600354815250905090565b6005818051602081018201805184825260208301602085012081835280955050505050506000915090508060010154908060020160009054906101000a900460ff16905082565b806000808273ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527f893d20e8000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516121289190615838565b6000604051808303816000865af19150503d8060008114612165576040519150601f19603f3d011682016040523d82523d6000602084013e61216a565b606091505b5091509150816121af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121a6906158c1565b60405180910390fd5b6000818060200190518101906121c5919061591f565b90508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612235576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161222c90615998565b60405180910390fd5b6000600c60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060008160010154116122bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122b690615bce565b60405180910390fd5b7fa3895d397a34e928a95d593331e293e2fc281d9d8996df5cc6c57f1cef629d428682600101546040516122f492919061560b565b60405180910390a1600c60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008082016000905560018201600090556002820160006101000a81549060ff02191690555050505050505050565b612372613aaf565b80600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60005481565b6123c4613aaf565b6001600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b612427613aaf565b6124316000613f2f565b565b61243b613aaf565b80600e8190555050565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166124d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016124c890615328565b60405180910390fd5b8060018190555050565b60006124e5614006565b905060008160000160089054906101000a900460ff1615905060008260000160009054906101000a900467ffffffffffffffff1690506000808267ffffffffffffffff161480156125335750825b9050600060018367ffffffffffffffff16148015612568575060003073ffffffffffffffffffffffffffffffffffffffff163b145b905081158015612576575080155b156125ad576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018560000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555083156125fd5760018560000160086101000a81548160ff0219169083151502179055505b6126063361402e565b61260e614042565b6001600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600281905550600160038190555067016345785d8a000060018190555083156126e15760008560000160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d260016040516126d89190615c47565b60405180910390a15b5050505050565b6000806126f361404c565b90508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505090565b60076020528060005260406000206000915090505481565b600d6020528060005260406000206000915090505481565b60015481565b60066020528060005260406000206000915090505481565b6127766144f3565b60405180608001604052808373ffffffffffffffffffffffffffffffffffffffff168152602001600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548152602001600d60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548152602001600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805461287190615a79565b80601f016020809104026020016040519081016040528092919081815260200182805461289d90615a79565b80156128ea5780601f106128bf576101008083540402835291602001916128ea565b820191906000526020600020905b8154815290600101906020018083116128cd57829003601f168201915b50505050508152509050919050565b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b60025481565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166129c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016129bb90615328565b60405180910390fd5b60006005846040516129d691906154a6565b90815260200160405180910390209050600d60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548213612ad6577fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b82604051612a5b9190614e3e565b60405180910390a181600080828254612a7491906154bd565b9250508190555081600d60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254612aca91906153e3565b92505081905550612eca565b60008290506000600d60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541315612c0d57600d60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600080828254612b7391906154bd565b92505081905550600d60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481612bc591906153e3565b90506000600d60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b80600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205412612cfd5780600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254612ca291906153e3565b9250508190555080600080828254612cba91906154bd565b925050819055507fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b83604051612cf09190614e3e565b60405180910390a1612ec8565b6000600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541315612e2f57600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481612d9091906153e3565b9050600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600080828254612de291906154bd565b925050819055506000600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b7fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b8184612e5c91906154f1565b604051612e699190614e3e565b60405180910390a180600d60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254612ec091906153e3565b925050819055505b505b60028160020160006101000a81548160ff02191690836002811115612ef257612ef16149e5565b5b0217905550612f0083613b36565b8273ffffffffffffffffffffffffffffffffffffffff167f403feb2cd85cc25c910fe59289105b583f08fe9e4335ebbb08c1851f47ff2d348386604051612f48929190615c62565b60405180910390a283604051612f5e91906154a6565b60405180910390207f4a2ced9ada462e244851a86e998eb0b5bf558c2c9c6923b7f970ed2b19b073b06002604051612f969190615c92565b60405180910390a250505050565b612fac614531565b600c60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060405180606001604052908160008201548152602001600182015481526020016002820160009054906101000a900460ff1615151515815250509050919050565b613034613aaf565b8060038190555050565b60035481565b806000808273ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527f893d20e8000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516130ef9190615838565b6000604051808303816000865af19150503d806000811461312c576040519150601f19603f3d011682016040523d82523d6000602084013e613131565b606091505b509150915081613176576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161316d906158c1565b60405180910390fd5b60008180602001905181019061318c919061591f565b90508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146131fc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016131f390615998565b60405180910390fd5b6000600c60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000816001015411613286576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161327d90615bce565b60405180910390fd5b600e54816000015461329891906154bd565b4310156132da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016132d190615cf9565b60405180910390fd5b8060010154600d60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541215613360576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161335790615a2a565b60405180910390fd5b60008160010154905080600d60008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546133b891906153e3565b92505081905550600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b815260040161341c92919061560b565b6020604051808303816000875af115801561343b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061345f9190615660565b507fe6fc6292b9fd5ba92a34a05a92d07b066535578023d71f0c6bf83a2622ca4e548733836040516134939392919061574e565b60405180910390a1600c60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008082016000905560018201600090556002820160006101000a81549060ff0219169055505061350a87613b36565b50505050505050565b61351b613aaf565b60005b81518110156136575781818151811061353a5761353961571f565b5b6020026020010151600d60008584815181106135595761355861571f565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546135aa9190615426565b925050819055508281815181106135c4576135c361571f565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f18b5d63f2a4f63f9d724d087c9580fdb3c4501bf641ce600afe4b97282e738e284848151811061362d5761362c61571f565b5b60200260200101516040516136429190614e3e565b60405180910390a3808060010191505061351e565b505050565b613664613aaf565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036136d65760006040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016136cd9190614f10565b60405180910390fd5b6136df81613f2f565b50565b60008111613725576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161371c90615d65565b60405180910390fd5b6000600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054136137a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161379e90615dd1565b60405180910390fd5b600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054811115613829576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161382090615e63565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1663893d20e86040518163ffffffff1660e01b8152600401602060405180830381865afa15801561388b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138af9190615e98565b73ffffffffffffffffffffffffffffffffffffffff1614613905576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016138fc90615f37565b60405180910390fd5b80600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461395491906153e3565b92505081905550600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b81526004016139b892919061560b565b6020604051808303816000875af11580156139d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139fb9190615660565b50613a0582613b36565b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb83604051613a629190614e3e565b60405180910390a35050565b600c6020528060005260406000206000915090508060000154908060010154908060020160009054906101000a900460ff16905083565b6000600e54905090565b613ab7614074565b73ffffffffffffffffffffffffffffffffffffffff16613ad56126e8565b73ffffffffffffffffffffffffffffffffffffffff1614613b3457613af8614074565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401613b2b9190614f10565b60405180910390fd5b565b600154600254613b469190615f57565b600d60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205412613c12576040518060400160405280600981526020017f7a6b41756374696f6e0000000000000000000000000000000000000000000000815250600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209081613c0c919061613b565b50613c95565b6040518060400160405280600381526020017f4e53430000000000000000000000000000000000000000000000000000000000815250600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209081613c93919061613b565b505b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161480613d4557507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16613d2c61407c565b73ffffffffffffffffffffffffffffffffffffffff1614155b15613d7c576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b613d86613aaf565b50565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015613df157506040513d601f19601f82011682018060405250810190613dee9190616239565b60015b613e3257816040517f4c9c8ce3000000000000000000000000000000000000000000000000000000008152600401613e299190614f10565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b8114613e9957806040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600401613e909190614e14565b60405180910390fd5b613ea383836140d3565b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614613f2d576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6000613f3961404c565b905060008160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050828260000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3505050565b60007ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b614036614146565b61403f81614186565b50565b61404a614146565b565b60007f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300905090565b600033905090565b60006140aa7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b61420c565b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6140dc82614216565b8173ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a26000815111156141395761413382826142e3565b50614142565b614141614367565b5b5050565b61414e6143a4565b614184576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b61418e614146565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036142005760006040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016141f79190614f10565b60405180910390fd5b61420981613f2f565b50565b6000819050919050565b60008173ffffffffffffffffffffffffffffffffffffffff163b0361427257806040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526004016142699190614f10565b60405180910390fd5b8061429f7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b61420c565b60000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60606000808473ffffffffffffffffffffffffffffffffffffffff168460405161430d9190615838565b600060405180830381855af49150503d8060008114614348576040519150601f19603f3d011682016040523d82523d6000602084013e61434d565b606091505b509150915061435d8583836143c4565b9250505092915050565b60003411156143a2576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60006143ae614006565b60000160089054906101000a900460ff16905090565b6060826143d9576143d482614453565b61444b565b60008251148015614401575060008473ffffffffffffffffffffffffffffffffffffffff163b145b1561444357836040517f9996b31500000000000000000000000000000000000000000000000000000000815260040161443a9190614f10565b60405180910390fd5b81905061444c565b5b9392505050565b6000815111156144665780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60405180606001604052806060815260200160008152602001600060028111156144c5576144c46149e5565b5b81525090565b6040518060800160405280600081526020016000815260200160008152602001600081525090565b6040518060800160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001606081525090565b604051806060016040528060008152602001600081526020016000151581525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6145bb82614572565b810181811067ffffffffffffffff821117156145da576145d9614583565b5b80604052505050565b60006145ed614554565b90506145f982826145b2565b919050565b600067ffffffffffffffff82111561461957614618614583565b5b61462282614572565b9050602081019050919050565b82818337600083830152505050565b600061465161464c846145fe565b6145e3565b90508281526020810184848401111561466d5761466c61456d565b5b61467884828561462f565b509392505050565b600082601f83011261469557614694614568565b5b81356146a584826020860161463e565b91505092915050565b6000602082840312156146c4576146c361455e565b5b600082013567ffffffffffffffff8111156146e2576146e1614563565b5b6146ee84828501614680565b91505092915050565b60008115159050919050565b61470c816146f7565b82525050565b60006020820190506147276000830184614703565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006147588261472d565b9050919050565b6147688161474d565b811461477357600080fd5b50565b6000813590506147858161475f565b92915050565b6000602082840312156147a1576147a061455e565b5b60006147af84828501614776565b91505092915050565b6000819050919050565b6147cb816147b8565b81146147d657600080fd5b50565b6000813590506147e8816147c2565b92915050565b600080604083850312156148055761480461455e565b5b600061481385828601614776565b9250506020614824858286016147d9565b9150509250929050565b600067ffffffffffffffff82111561484957614848614583565b5b61485282614572565b9050602081019050919050565b600061487261486d8461482e565b6145e3565b90508281526020810184848401111561488e5761488d61456d565b5b61489984828561462f565b509392505050565b600082601f8301126148b6576148b5614568565b5b81356148c684826020860161485f565b91505092915050565b6000602082840312156148e5576148e461455e565b5b600082013567ffffffffffffffff81111561490357614902614563565b5b61490f848285016148a1565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61494d8161474d565b82525050565b600061495f8383614944565b60208301905092915050565b6000602082019050919050565b600061498382614918565b61498d8185614923565b935061499883614934565b8060005b838110156149c95781516149b08882614953565b97506149bb8361496b565b92505060018101905061499c565b5085935050505092915050565b6149df816147b8565b82525050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60038110614a2557614a246149e5565b5b50565b6000819050614a3682614a14565b919050565b6000614a4682614a28565b9050919050565b614a5681614a3b565b82525050565b60006060830160008301518482036000860152614a798282614978565b9150506020830151614a8e60208601826149d6565b506040830151614aa16040860182614a4d565b508091505092915050565b60006020820190508181036000830152614ac68184614a5c565b905092915050565b600060208284031215614ae457614ae361455e565b5b6000614af2848285016147d9565b91505092915050565b600067ffffffffffffffff821115614b1657614b15614583565b5b602082029050602081019050919050565b600080fd5b6000614b3f614b3a84614afb565b6145e3565b90508083825260208201905060208402830185811115614b6257614b61614b27565b5b835b81811015614b8b5780614b778882614776565b845260208401935050602081019050614b64565b5050509392505050565b600082601f830112614baa57614ba9614568565b5b8135614bba848260208601614b2c565b91505092915050565b600067ffffffffffffffff821115614bde57614bdd614583565b5b602082029050602081019050919050565b6000614c02614bfd84614bc3565b6145e3565b90508083825260208201905060208402830185811115614c2557614c24614b27565b5b835b81811015614c4e5780614c3a88826147d9565b845260208401935050602081019050614c27565b5050509392505050565b600082601f830112614c6d57614c6c614568565b5b8135614c7d848260208601614bef565b91505092915050565b60008060408385031215614c9d57614c9c61455e565b5b600083013567ffffffffffffffff811115614cbb57614cba614563565b5b614cc785828601614b95565b925050602083013567ffffffffffffffff811115614ce857614ce7614563565b5b614cf485828601614c58565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b60005b83811015614d38578082015181840152602081019050614d1d565b60008484015250505050565b6000614d4f82614cfe565b614d598185614d09565b9350614d69818560208601614d1a565b614d7281614572565b840191505092915050565b60006020820190508181036000830152614d978184614d44565b905092915050565b60008060408385031215614db657614db561455e565b5b6000614dc485828601614776565b925050602083013567ffffffffffffffff811115614de557614de4614563565b5b614df185828601614680565b9150509250929050565b6000819050919050565b614e0e81614dfb565b82525050565b6000602082019050614e296000830184614e05565b92915050565b614e38816147b8565b82525050565b6000602082019050614e536000830184614e2f565b92915050565b608082016000820151614e6f60008501826149d6565b506020820151614e8260208501826149d6565b506040820151614e9560408501826149d6565b506060820151614ea860608501826149d6565b50505050565b6000608082019050614ec36000830184614e59565b92915050565b614ed281614a3b565b82525050565b6000604082019050614eed6000830185614e2f565b614efa6020830184614ec9565b9392505050565b614f0a8161474d565b82525050565b6000602082019050614f256000830184614f01565b92915050565b6000819050919050565b614f3e81614f2b565b82525050565b6000602082019050614f596000830184614f35565b92915050565b614f6881614f2b565b82525050565b600082825260208201905092915050565b6000614f8a82614cfe565b614f948185614f6e565b9350614fa4818560208601614d1a565b614fad81614572565b840191505092915050565b6000608083016000830151614fd06000860182614944565b506020830151614fe36020860182614f5f565b506040830151614ff66040860182614f5f565b506060830151848203606086015261500e8282614f7f565b9150508091505092915050565b600060208201905081810360008301526150358184614fb8565b905092915050565b6000806000606084860312156150565761505561455e565b5b600084013567ffffffffffffffff81111561507457615073614563565b5b615080868287016148a1565b935050602061509186828701614776565b92505060406150a2868287016147d9565b9150509250925092565b6150b5816146f7565b82525050565b6060820160008201516150d160008501826149d6565b5060208201516150e460208501826149d6565b5060408201516150f760408501826150ac565b50505050565b600060608201905061511260008301846150bb565b92915050565b600067ffffffffffffffff82111561513357615132614583565b5b602082029050602081019050919050565b61514d81614f2b565b811461515857600080fd5b50565b60008135905061516a81615144565b92915050565b600061518361517e84615118565b6145e3565b905080838252602082019050602084028301858111156151a6576151a5614b27565b5b835b818110156151cf57806151bb888261515b565b8452602084019350506020810190506151a8565b5050509392505050565b600082601f8301126151ee576151ed614568565b5b81356151fe848260208601615170565b91505092915050565b6000806040838503121561521e5761521d61455e565b5b600083013567ffffffffffffffff81111561523c5761523b614563565b5b61524885828601614b95565b925050602083013567ffffffffffffffff81111561526957615268614563565b5b615275858286016151d9565b9150509250929050565b60006060820190506152946000830186614e2f565b6152a16020830185614e2f565b6152ae6040830184614703565b949350505050565b7f4f6e6c79207468652061646d696e2063616e2063616c6c20746869732066756e60008201527f6374696f6e2e0000000000000000000000000000000000000000000000000000602082015250565b6000615312602683614d09565b915061531d826152b6565b604082019050919050565b6000602082019050818103600083015261534181615305565b9050919050565b7f6e6f7420656e6f756768206c6f636b65642062616c616e636500000000000000600082015250565b600061537e601983614d09565b915061538982615348565b602082019050919050565b600060208201905081810360008301526153ad81615371565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006153ee82614f2b565b91506153f983614f2b565b92508282039050818112600084121682821360008512151617156154205761541f6153b4565b5b92915050565b600061543182614f2b565b915061543c83614f2b565b925082820190508281121560008312168382126000841215161715615464576154636153b4565b5b92915050565b600081905092915050565b600061548082614cfe565b61548a818561546a565b935061549a818560208601614d1a565b80840191505092915050565b60006154b28284615475565b915081905092915050565b60006154c8826147b8565b91506154d3836147b8565b92508282019050808211156154eb576154ea6153b4565b5b92915050565b60006154fc826147b8565b9150615507836147b8565b925082820390508181111561551f5761551e6153b4565b5b92915050565b50565b6000615535600083614d09565b915061554082615525565b600082019050919050565b60006040820190506155606000830184614e2f565b818103602083015261557181615528565b905092915050565b7f576974686472617720616d6f756e7420616d6f756e742065786365656473207360008201527f6c617368656446756e6473000000000000000000000000000000000000000000602082015250565b60006155d5602b83614d09565b91506155e082615579565b604082019050919050565b60006020820190508181036000830152615604816155c8565b9050919050565b60006040820190506156206000830185614f01565b61562d6020830184614e2f565b9392505050565b61563d816146f7565b811461564857600080fd5b50565b60008151905061565a81615634565b92915050565b6000602082840312156156765761567561455e565b5b60006156848482850161564b565b91505092915050565b7f6370206163636f756e747320616e6420616d6f756e747320617265206469666660008201527f6572656e74206c656e6774687300000000000000000000000000000000000000602082015250565b60006156e9602d83614d09565b91506156f48261568d565b604082019050919050565b60006020820190508181036000830152615718816156dc565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006060820190506157636000830186614f01565b6157706020830185614f01565b61577d6040830184614e2f565b949350505050565b7f7472616e73666572206661696c65640000000000000000000000000000000000600082015250565b60006157bb600f83614d09565b91506157c682615785565b602082019050919050565b600060208201905081810360008301526157ea816157ae565b9050919050565b600081519050919050565b600081905092915050565b6000615812826157f1565b61581c81856157fc565b935061582c818560208601614d1a565b80840191505092915050565b60006158448284615807565b915081905092915050565b7f4661696c656420746f2063616c6c206765744f776e65722066756e6374696f6e60008201527f206f662043504163636f756e7400000000000000000000000000000000000000602082015250565b60006158ab602d83614d09565b91506158b68261584f565b604082019050919050565b600060208201905081810360008301526158da8161589e565b9050919050565b60006158ec8261472d565b9050919050565b6158fc816158e1565b811461590757600080fd5b50565b600081519050615919816158f3565b92915050565b6000602082840312156159355761593461455e565b5b60006159438482850161590a565b91505092915050565b7f63616c6c6572206973206e6f74206370206f776e657200000000000000000000600082015250565b6000615982601683614d09565b915061598d8261594c565b602082019050919050565b600060208201905081810360008301526159b181615975565b9050919050565b7f4e6f7420656e6f7567682066726f7a656e2062616c616e636520746f2077697460008201527f6864726177000000000000000000000000000000000000000000000000000000602082015250565b6000615a14602583614d09565b9150615a1f826159b8565b604082019050919050565b60006020820190508181036000830152615a4381615a07565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680615a9157607f821691505b602082108103615aa457615aa3615a4a565b5b50919050565b7f616d6f756e74206d7573742062652067726561746572207468616e2030000000600082015250565b6000615ae0601d83614d09565b9150615aeb82615aaa565b602082019050919050565b60006020820190508181036000830152615b0f81615ad3565b9050919050565b7f6e6f7420656e6f7567682062616c616e63650000000000000000000000000000600082015250565b6000615b4c601283614d09565b9150615b5782615b16565b602082019050919050565b60006020820190508181036000830152615b7b81615b3f565b9050919050565b7f4e6f2070656e64696e6720776974686472617720726571756573740000000000600082015250565b6000615bb8601b83614d09565b9150615bc382615b82565b602082019050919050565b60006020820190508181036000830152615be781615bab565b9050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b6000819050919050565b6000615c31615c2c615c2784615bee565b615c0c565b615bf8565b9050919050565b615c4181615c16565b82525050565b6000602082019050615c5c6000830184615c38565b92915050565b6000604082019050615c776000830185614e2f565b8181036020830152615c898184614d44565b90509392505050565b6000602082019050615ca76000830184614ec9565b92915050565b7f57697468647261772064656c6179206e6f742070617373656400000000000000600082015250565b6000615ce3601983614d09565b9150615cee82615cad565b602082019050919050565b60006020820190508181036000830152615d1281615cd6565b9050919050565b7f776974686472617720616d6f756e74206d75737420626520706f736974697665600082015250565b6000615d4f602083614d09565b9150615d5a82615d19565b602082019050919050565b60006020820190508181036000830152615d7e81615d42565b9050919050565b7f6e6f7468696e6720746f20776974686472617700000000000000000000000000600082015250565b6000615dbb601383614d09565b9150615dc682615d85565b602082019050919050565b60006020820190508181036000830152615dea81615dae565b9050919050565b7f776974686472617720616d6f756e742067726561746572207468616e2061766160008201527f696c61626c652062616c616e6365000000000000000000000000000000000000602082015250565b6000615e4d602e83614d09565b9150615e5882615df1565b604082019050919050565b60006020820190508181036000830152615e7c81615e40565b9050919050565b600081519050615e928161475f565b92915050565b600060208284031215615eae57615ead61455e565b5b6000615ebc84828501615e83565b91505092915050565b7f4f6e6c792043504163636f756e74206f776e65722063616e207769746864726160008201527f772074686520636f6c6c61746572616c2066756e647300000000000000000000602082015250565b6000615f21603683614d09565b9150615f2c82615ec5565b604082019050919050565b60006020820190508181036000830152615f5081615f14565b9050919050565b6000615f62826147b8565b9150615f6d836147b8565b9250828202615f7b816147b8565b91508282048414831517615f9257615f916153b4565b5b5092915050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302615ffb7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82615fbe565b6160058683615fbe565b95508019841693508086168417925050509392505050565b600061603861603361602e846147b8565b615c0c565b6147b8565b9050919050565b6000819050919050565b6160528361601d565b61606661605e8261603f565b848454615fcb565b825550505050565b600090565b61607b61606e565b616086818484616049565b505050565b5b818110156160aa5761609f600082616073565b60018101905061608c565b5050565b601f8211156160ef576160c081615f99565b6160c984615fae565b810160208510156160d8578190505b6160ec6160e485615fae565b83018261608b565b50505b505050565b600082821c905092915050565b6000616112600019846008026160f4565b1980831691505092915050565b600061612b8383616101565b9150826002028217905092915050565b61614482614cfe565b67ffffffffffffffff81111561615d5761615c614583565b5b6161678254615a79565b6161728282856160ae565b600060209050601f8311600181146161a55760008415616193578287015190505b61619d858261611f565b865550616205565b601f1984166161b386615f99565b60005b828110156161db578489015182556001820191506020850194506020810190506161b6565b868310156161f857848901516161f4601f891682616101565b8355505b6001600288020188555050505b505050505050565b61621681614dfb565b811461622157600080fd5b50565b6000815190506162338161620d565b92915050565b60006020828403121561624f5761624e61455e565b5b600061625d84828501616224565b9150509291505056fea2646970667358221220998e8c0c6b85d753bf3374b2a51526a5976d443852acc8f84336a55b250dffe664736f6c63430008190033",
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
