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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"frozenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"operation\",\"type\":\"string\"}],\"name\":\"CollateralAdjusted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"CollateralLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"CollateralSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumSwanCreditCollateral.CollateralStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"CollateralStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"CollateralUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fundingWallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fundingWallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"}],\"name\":\"DepositLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"DisputeProof\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SlashedFundsIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cpAccountAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cpOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"WithdrawConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"WithdrawRequestCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"WithdrawRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralContratOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashfund\",\"type\":\"uint256\"}],\"name\":\"WithdrawSlash\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedToWithdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"availableBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"cpAccounts\",\"type\":\"address[]\"},{\"internalType\":\"int256[]\",\"name\":\"amounts\",\"type\":\"int256[]\"}],\"name\":\"batchAddLockedCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"cpAccounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"}],\"name\":\"cancelWithdrawRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"slashedFunds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashRatio\",\"type\":\"uint256\"}],\"internalType\":\"structSwanCreditCollateral.ContractInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"}],\"name\":\"confirmWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAddress\",\"type\":\"address\"}],\"name\":\"cpInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"availableBalance\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"lockedBalance\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"}],\"internalType\":\"structSwanCreditCollateral.CPInfoWithLockedBalance\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cpStatus\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"escrowBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"getTaskInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"cpList\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"enumSwanCreditCollateral.CollateralStatus\",\"name\":\"collateralStatus\",\"type\":\"uint8\"}],\"internalType\":\"structSwanCreditCollateral.Task\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWithdrawDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"isSignatureUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lockCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockedBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockedCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"requestWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseCollateral\",\"type\":\"uint256\"}],\"name\":\"setBaseCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"setCollateralToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slashRatio\",\"type\":\"uint256\"}],\"name\":\"setSlashRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawDelay\",\"type\":\"uint256\"}],\"name\":\"setWithdrawDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slashAmount\",\"type\":\"uint256\"}],\"name\":\"slashCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slashAmount\",\"type\":\"uint256\"}],\"name\":\"slashCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashedFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"enumSwanCreditCollateral.CollateralStatus\",\"name\":\"collateralStatus\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"unlockRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPending\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"}],\"name\":\"viewWithdrawRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requestTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPending\",\"type\":\"bool\"}],\"internalType\":\"structSwanCreditCollateral.UnlockRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawSlashedFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801561004357600080fd5b5061005261005760201b60201c565b6101c1565b600061006761015b60201b60201c565b90508060000160089054906101000a900460ff16156100b2576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff80168160000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff16146101585767ffffffffffffffff8160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d267ffffffffffffffff60405161014f91906101a6565b60405180910390a15b50565b60007ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b600067ffffffffffffffff82169050919050565b6101a081610183565b82525050565b60006020820190506101bb6000830184610197565b92915050565b60805161602d6101ea60003960008181613a9701528181613aec0152613ca7015261602d6000f3fe6080604052600436106102675760003560e01c8063715018a611610144578063b4eae1cb116100b6578063d2bfc1c71161007a578063d2bfc1c714610954578063d76358151461097d578063f2fde38b146109a6578063f3fef3a3146109cf578063f97348e1146109f8578063fe3300d014610a3757610267565b8063b4eae1cb1461086f578063b587b82c1461089a578063bede6e31146108c3578063ce3518aa14610900578063d27ca89b1461092957610267565b806392bdf9ba1161010857806392bdf9ba146107255780639ae697bf146107625780639b5ddf091461079f578063a0821be3146107ca578063a664c21614610807578063ad3cb1cc1461084457610267565b8063715018a61461067a57806372f0cb30146106915780637f58a7e5146106ba5780638129fc1c146106e35780638da5cb5b146106fa57610267565b806347e7ef24116101dd57806355af6353116101a157806355af63531461055957806358709cf2146105965780635d2cd2a7146105d4578063666181a9146105fd5780636f99f15c14610626578063704802751461065157610267565b806347e7ef24146104935780634f1ef286146104bc57806352d1902d146104d8578063536f60701461050357806353ad87201461052e57610267565b80632894493f1161022f5780632894493f146103755780632d291cad1461039e57806330a90736146103db578063397a1b28146104045780633fe651771461042d57806347a7d1071461046a57610267565b80631150f0f31461026c5780631785f53c146102a95780631b209463146102d25780631d47a62d1461030f57806324d7806c14610338575b600080fd5b34801561027857600080fd5b50610293600480360381019061028e91906144ab565b610a62565b6040516102a0919061450f565b60405180910390f35b3480156102b557600080fd5b506102d060048036038101906102cb9190614588565b610a98565b005b3480156102de57600080fd5b506102f960048036038101906102f49190614656565b610afb565b604051610306919061483d565b60405180910390f35b34801561031b57600080fd5b506103366004803603810190610331919061488b565b610c06565b005b34801561034457600080fd5b5061035f600480360381019061035a9190614588565b6111d1565b60405161036c919061450f565b60405180910390f35b34801561038157600080fd5b5061039c600480360381019061039791906148cb565b6111f1565b005b3480156103aa57600080fd5b506103c560048036038101906103c09190614588565b611348565b6040516103d2919061450f565b60405180910390f35b3480156103e757600080fd5b5061040260048036038101906103fd9190614a83565b611368565b005b34801561041057600080fd5b5061042b6004803603810190610426919061488b565b61161b565b005b34801561043957600080fd5b50610454600480360381019061044f9190614588565b61192a565b6040516104619190614b7a565b60405180910390f35b34801561047657600080fd5b50610491600480360381019061048c919061488b565b6119ca565b005b34801561049f57600080fd5b506104ba60048036038101906104b5919061488b565b611c22565b005b6104d660048036038101906104d19190614b9c565b611d8d565b005b3480156104e457600080fd5b506104ed611dac565b6040516104fa9190614c11565b60405180910390f35b34801561050f57600080fd5b50610518611ddf565b6040516105259190614c3b565b60405180910390f35b34801561053a57600080fd5b50610543611de5565b6040516105509190614cab565b60405180910390f35b34801561056557600080fd5b50610580600480360381019061057b9190614588565b611e1b565b60405161058d9190614cdf565b60405180910390f35b3480156105a257600080fd5b506105bd60048036038101906105b89190614656565b611e33565b6040516105cb929190614d09565b60405180910390f35b3480156105e057600080fd5b506105fb60048036038101906105f69190614588565b611e7a565b005b34801561060957600080fd5b50610624600480360381019061061f9190614588565b612167565b005b34801561063257600080fd5b5061063b6121b3565b6040516106489190614c3b565b60405180910390f35b34801561065d57600080fd5b5061067860048036038101906106739190614588565b6121b9565b005b34801561068657600080fd5b5061068f61221c565b005b34801561069d57600080fd5b506106b860048036038101906106b391906148cb565b612230565b005b3480156106c657600080fd5b506106e160048036038101906106dc91906148cb565b612242565b005b3480156106ef57600080fd5b506106f86122d8565b005b34801561070657600080fd5b5061070f6124e5565b60405161071c9190614d41565b60405180910390f35b34801561073157600080fd5b5061074c60048036038101906107479190614588565b61251d565b6040516107599190614c3b565b60405180910390f35b34801561076e57600080fd5b5061078960048036038101906107849190614588565b612535565b6040516107969190614cdf565b60405180910390f35b3480156107ab57600080fd5b506107b461254d565b6040516107c19190614c3b565b60405180910390f35b3480156107d657600080fd5b506107f160048036038101906107ec9190614588565b612553565b6040516107fe9190614cdf565b60405180910390f35b34801561081357600080fd5b5061082e60048036038101906108299190614588565b61256b565b60405161083b9190614e18565b60405180910390f35b34801561085057600080fd5b506108596126f6565b6040516108669190614b7a565b60405180910390f35b34801561087b57600080fd5b5061088461272f565b6040516108919190614c3b565b60405180910390f35b3480156108a657600080fd5b506108c160048036038101906108bc9190614e3a565b612735565b005b3480156108cf57600080fd5b506108ea60048036038101906108e59190614588565b612da1565b6040516108f79190614efa565b60405180910390f35b34801561090c57600080fd5b50610927600480360381019061092291906148cb565b612e29565b005b34801561093557600080fd5b5061093e612e3b565b60405161094b9190614c3b565b60405180910390f35b34801561096057600080fd5b5061097b60048036038101906109769190614588565b612e41565b005b34801561098957600080fd5b506109a4600480360381019061099f9190615004565b613310565b005b3480156109b257600080fd5b506109cd60048036038101906109c89190614588565b613459565b005b3480156109db57600080fd5b506109f660048036038101906109f1919061488b565b6134df565b005b348015610a0457600080fd5b50610a1f6004803603810190610a1a9190614588565b61386b565b604051610a2e9392919061507c565b60405180910390f35b348015610a4357600080fd5b50610a4c6138a2565b604051610a599190614c3b565b60405180910390f35b600b818051602081018201805184825260208301602085012081835280955050505050506000915054906101000a900460ff1681565b610aa06138ac565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b610b03614295565b600582604051610b1391906150ef565b908152602001604051809103902060405180606001604052908160008201805480602002602001604051908101604052809291908181526020018280548015610bb157602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610b67575b50505050508152602001600182015481526020016002820160009054906101000a900460ff166002811115610be957610be8614776565b5b6002811115610bfb57610bfa614776565b5b815250509050919050565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610c92576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c8990615178565b60405180910390fd5b601060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548113610d82577fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b81604051610d079190614c3b565b60405180910390a180600080828254610d2091906151c7565b9250508190555080601060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610d7691906151fb565b92505081905550611176565b60008190506000601060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541315610eb957601060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600080828254610e1f91906151c7565b92505081905550601060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481610e7191906151fb565b90506000601060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b80600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205412610fa95780600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610f4e91906151fb565b9250508190555080600080828254610f6691906151c7565b925050819055507fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b82604051610f9c9190614c3b565b60405180910390a1611174565b6000600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205413156110db57600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548161103c91906151fb565b9050600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205460008082825461108e91906151c7565b925050819055506000600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b7fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b8183611108919061523e565b6040516111159190614c3b565b60405180910390a180601060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461116c91906151fb565b925050819055505b505b61117f82613933565b8173ffffffffffffffffffffffffffffffffffffffff167f403feb2cd85cc25c910fe59289105b583f08fe9e4335ebbb08c1851f47ff2d34826040516111c59190615298565b60405180910390a25050565b60046020528060005260406000206000915054906101000a900460ff1681565b6111f96138ac565b80600054101561123e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161123590615338565b60405180910390fd5b8060008082825461124f919061523e565b92505081905550600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b81526004016112b3929190615358565b6020604051808303816000875af11580156112d2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112f691906153ad565b503373ffffffffffffffffffffffffffffffffffffffff167fbfd9c82485e2178fcfd5c40379d8e66fe60acc50aa1ef4c50966431eb1e381cd8260405161133d9190614c3b565b60405180910390a250565b600a6020528060005260406000206000915054906101000a900460ff1681565b80518251146113ac576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113a39061544c565b60405180910390fd5b6000805b8251811015611534578281815181106113cc576113cb61546c565b5b6020026020010151826113df91906151c7565b91508281815181106113f4576113f361546c565b5b6020026020010151600660008684815181106114135761141261546c565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611464919061549b565b9250508190555061148e8482815181106114815761148061546c565b5b6020026020010151613933565b8381815181106114a1576114a061546c565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f6285848151811061150a5761150961546c565b5b602002602001015160405161151f9190614c3b565b60405180910390a380806001019150506113b0565b50600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b8152600401611594939291906154df565b6020604051808303816000875af11580156115b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115d791906153ad565b611616576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161160d90615562565b60405180910390fd5b505050565b816000808273ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527f893d20e8000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516116c691906155c9565b6000604051808303816000865af19150503d8060008114611703576040519150601f19603f3d011682016040523d82523d6000602084013e611708565b606091505b50915091508161174d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161174490615652565b60405180910390fd5b60008180602001905181019061176391906156b0565b90508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146117d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117ca90615729565b60405180910390fd5b84601060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541215611855576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161184c906157bb565b60405180910390fd5b604051806060016040528043815260200186815260200160011515815250600d60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548160ff0219169083151502179055509050507ff7774b688d56120b783560a913ee60792a73dfd511812b7be5eccf10d08c6689868660405161191a929190615358565b60405180910390a1505050505050565b600860205280600052604060002060009150905080546119499061580a565b80601f01602080910402602001604051908101604052809291908181526020018280546119759061580a565b80156119c25780601f10611997576101008083540402835291602001916119c2565b820191906000526020600020905b8154815290600101906020018083116119a557829003601f168201915b505050505081565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611a56576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a4d90615178565b60405180910390fd5b60008111611a99576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a9090615887565b60405180910390fd5b600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054811315611b1b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b12906158f3565b60405180910390fd5b80600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611b6a91906151fb565b9250508190555080601060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611bc0919061549b565b92505081905550611bd082613933565b8173ffffffffffffffffffffffffffffffffffffffff167f5f3d004cf9164b95ed5dbf47d1f04018a4eabcb20b4320fe229ed92236ace63482604051611c169190615298565b60405180910390a25050565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b8152600401611c81939291906154df565b6020604051808303816000875af1158015611ca0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cc491906153ad565b5080600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611d14919061549b565b925050819055508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f6283604051611d789190614c3b565b60405180910390a3611d8982613933565b5050565b611d95613a95565b611d9e82613b7b565b611da88282613b86565b5050565b6000611db6613ca5565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b905090565b600e5481565b611ded6142c8565b6040518060800160405280600054815260200160015481526020016002548152602001600354815250905090565b600c6020528060005260406000206000915090505481565b6005818051602081018201805184825260208301602085012081835280955050505050506000915090508060010154908060020160009054906101000a900460ff16905082565b806000808273ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527f893d20e8000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051611f2591906155c9565b6000604051808303816000865af19150503d8060008114611f62576040519150601f19603f3d011682016040523d82523d6000602084013e611f67565b606091505b509150915081611fac576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611fa390615652565b60405180910390fd5b600081806020019051810190611fc291906156b0565b90508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612032576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161202990615729565b60405180910390fd5b6000600d60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060008160010154116120bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120b39061595f565b60405180910390fd5b7fa3895d397a34e928a95d593331e293e2fc281d9d8996df5cc6c57f1cef629d428682600101546040516120f1929190615358565b60405180910390a1600d60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008082016000905560018201600090556002820160006101000a81549060ff02191690555050505050505050565b61216f6138ac565b80600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60005481565b6121c16138ac565b6001600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b6122246138ac565b61222e6000613d2c565b565b6122386138ac565b80600e8190555050565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166122ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122c590615178565b60405180910390fd5b8060018190555050565b60006122e2613e03565b905060008160000160089054906101000a900460ff1615905060008260000160009054906101000a900467ffffffffffffffff1690506000808267ffffffffffffffff161480156123305750825b9050600060018367ffffffffffffffff16148015612365575060003073ffffffffffffffffffffffffffffffffffffffff163b145b905081158015612373575080155b156123aa576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018560000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555083156123fa5760018560000160086101000a81548160ff0219169083151502179055505b61240333613e2b565b61240b613e3f565b6001600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600281905550600160038190555067016345785d8a000060018190555083156124de5760008560000160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d260016040516124d591906159d8565b60405180910390a15b5050505050565b6000806124f0613e49565b90508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505090565b60076020528060005260406000206000915090505481565b60106020528060005260406000206000915090505481565b60015481565b60066020528060005260406000206000915090505481565b6125736142f0565b60405180608001604052808373ffffffffffffffffffffffffffffffffffffffff168152602001600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548152602001601060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548152602001600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805461266e9061580a565b80601f016020809104026020016040519081016040528092919081815260200182805461269a9061580a565b80156126e75780601f106126bc576101008083540402835291602001916126e7565b820191906000526020600020905b8154815290600101906020018083116126ca57829003601f168201915b50505050508152509050919050565b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b60025481565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166127c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016127b890615178565b60405180910390fd5b60006005846040516127d391906150ef565b90815260200160405180910390209050601060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482136128d3577fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b826040516128589190614c3b565b60405180910390a18160008082825461287191906151c7565b9250508190555081601060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546128c791906151fb565b92505081905550612cc7565b60008290506000601060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541315612a0a57601060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205460008082825461297091906151c7565b92505081905550601060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054816129c291906151fb565b90506000601060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b80600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205412612afa5780600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254612a9f91906151fb565b9250508190555080600080828254612ab791906151c7565b925050819055507fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b83604051612aed9190614c3b565b60405180910390a1612cc5565b6000600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541315612c2c57600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481612b8d91906151fb565b9050600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600080828254612bdf91906151c7565b925050819055506000600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b7fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b8184612c59919061523e565b604051612c669190614c3b565b60405180910390a180601060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254612cbd91906151fb565b925050819055505b505b60028160020160006101000a81548160ff02191690836002811115612cef57612cee614776565b5b0217905550612cfd83613933565b8273ffffffffffffffffffffffffffffffffffffffff167f403feb2cd85cc25c910fe59289105b583f08fe9e4335ebbb08c1851f47ff2d348386604051612d459291906159f3565b60405180910390a283604051612d5b91906150ef565b60405180910390207f4a2ced9ada462e244851a86e998eb0b5bf558c2c9c6923b7f970ed2b19b073b06002604051612d939190615a23565b60405180910390a250505050565b612da961432e565b600d60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060405180606001604052908160008201548152602001600182015481526020016002820160009054906101000a900460ff1615151515815250509050919050565b612e316138ac565b8060038190555050565b60035481565b806000808273ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527f893d20e8000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051612eec91906155c9565b6000604051808303816000865af19150503d8060008114612f29576040519150601f19603f3d011682016040523d82523d6000602084013e612f2e565b606091505b509150915081612f73576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612f6a90615652565b60405180910390fd5b600081806020019051810190612f8991906156b0565b90508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612ff9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612ff090615729565b60405180910390fd5b6000600d60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000816001015411613083576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161307a9061595f565b60405180910390fd5b600e54816000015461309591906151c7565b4310156130d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016130ce90615a8a565b60405180910390fd5b8060010154601060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054121561315d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613154906157bb565b60405180910390fd5b60008160010154905080601060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546131b591906151fb565b92505081905550600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b8152600401613219929190615358565b6020604051808303816000875af1158015613238573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061325c91906153ad565b507fe6fc6292b9fd5ba92a34a05a92d07b066535578023d71f0c6bf83a2622ca4e54873383604051613290939291906154df565b60405180910390a1600d60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008082016000905560018201600090556002820160006101000a81549060ff0219169055505061330787613933565b50505050505050565b6133186138ac565b60005b8151811015613454578181815181106133375761333661546c565b5b6020026020010151601060008584815181106133565761335561546c565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546133a7919061549b565b925050819055508281815181106133c1576133c061546c565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f18b5d63f2a4f63f9d724d087c9580fdb3c4501bf641ce600afe4b97282e738e284848151811061342a5761342961546c565b5b602002602001015160405161343f9190614c3b565b60405180910390a3808060010191505061331b565b505050565b6134616138ac565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036134d35760006040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016134ca9190614d41565b60405180910390fd5b6134dc81613d2c565b50565b60008111613522576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161351990615af6565b60405180910390fd5b6000600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054136135a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161359b90615b62565b60405180910390fd5b600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054811115613626576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161361d90615bf4565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1663893d20e86040518163ffffffff1660e01b8152600401602060405180830381865afa158015613688573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136ac9190615c29565b73ffffffffffffffffffffffffffffffffffffffff1614613702576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016136f990615cc8565b60405180910390fd5b80600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461375191906151fb565b92505081905550600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b81526004016137b5929190615358565b6020604051808303816000875af11580156137d4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137f891906153ad565b5061380282613933565b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb8360405161385f9190614c3b565b60405180910390a35050565b600d6020528060005260406000206000915090508060000154908060010154908060020160009054906101000a900460ff16905083565b6000600e54905090565b6138b4613e71565b73ffffffffffffffffffffffffffffffffffffffff166138d26124e5565b73ffffffffffffffffffffffffffffffffffffffff1614613931576138f5613e71565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016139289190614d41565b60405180910390fd5b565b6001546002546139439190615ce8565b601060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205412613a0f576040518060400160405280600981526020017f7a6b41756374696f6e0000000000000000000000000000000000000000000000815250600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209081613a099190615ecc565b50613a92565b6040518060400160405280600381526020017f4e53430000000000000000000000000000000000000000000000000000000000815250600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209081613a909190615ecc565b505b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161480613b4257507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16613b29613e79565b73ffffffffffffffffffffffffffffffffffffffff1614155b15613b79576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b613b836138ac565b50565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015613bee57506040513d601f19601f82011682018060405250810190613beb9190615fca565b60015b613c2f57816040517f4c9c8ce3000000000000000000000000000000000000000000000000000000008152600401613c269190614d41565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b8114613c9657806040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600401613c8d9190614c11565b60405180910390fd5b613ca08383613ed0565b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614613d2a576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6000613d36613e49565b905060008160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050828260000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3505050565b60007ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b613e33613f43565b613e3c81613f83565b50565b613e47613f43565b565b60007f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300905090565b600033905090565b6000613ea77f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b614009565b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b613ed982614013565b8173ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a2600081511115613f3657613f3082826140e0565b50613f3f565b613f3e614164565b5b5050565b613f4b6141a1565b613f81576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b613f8b613f43565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603613ffd5760006040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401613ff49190614d41565b60405180910390fd5b61400681613d2c565b50565b6000819050919050565b60008173ffffffffffffffffffffffffffffffffffffffff163b0361406f57806040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526004016140669190614d41565b60405180910390fd5b8061409c7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b614009565b60000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60606000808473ffffffffffffffffffffffffffffffffffffffff168460405161410a91906155c9565b600060405180830381855af49150503d8060008114614145576040519150601f19603f3d011682016040523d82523d6000602084013e61414a565b606091505b509150915061415a8583836141c1565b9250505092915050565b600034111561419f576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60006141ab613e03565b60000160089054906101000a900460ff16905090565b6060826141d6576141d182614250565b614248565b600082511480156141fe575060008473ffffffffffffffffffffffffffffffffffffffff163b145b1561424057836040517f9996b3150000000000000000000000000000000000000000000000000000000081526004016142379190614d41565b60405180910390fd5b819050614249565b5b9392505050565b6000815111156142635780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60405180606001604052806060815260200160008152602001600060028111156142c2576142c1614776565b5b81525090565b6040518060800160405280600081526020016000815260200160008152602001600081525090565b6040518060800160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001606081525090565b604051806060016040528060008152602001600081526020016000151581525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6143b88261436f565b810181811067ffffffffffffffff821117156143d7576143d6614380565b5b80604052505050565b60006143ea614351565b90506143f682826143af565b919050565b600067ffffffffffffffff82111561441657614415614380565b5b61441f8261436f565b9050602081019050919050565b82818337600083830152505050565b600061444e614449846143fb565b6143e0565b90508281526020810184848401111561446a5761446961436a565b5b61447584828561442c565b509392505050565b600082601f83011261449257614491614365565b5b81356144a284826020860161443b565b91505092915050565b6000602082840312156144c1576144c061435b565b5b600082013567ffffffffffffffff8111156144df576144de614360565b5b6144eb8482850161447d565b91505092915050565b60008115159050919050565b614509816144f4565b82525050565b60006020820190506145246000830184614500565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006145558261452a565b9050919050565b6145658161454a565b811461457057600080fd5b50565b6000813590506145828161455c565b92915050565b60006020828403121561459e5761459d61435b565b5b60006145ac84828501614573565b91505092915050565b600067ffffffffffffffff8211156145d0576145cf614380565b5b6145d98261436f565b9050602081019050919050565b60006145f96145f4846145b5565b6143e0565b9050828152602081018484840111156146155761461461436a565b5b61462084828561442c565b509392505050565b600082601f83011261463d5761463c614365565b5b813561464d8482602086016145e6565b91505092915050565b60006020828403121561466c5761466b61435b565b5b600082013567ffffffffffffffff81111561468a57614689614360565b5b61469684828501614628565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6146d48161454a565b82525050565b60006146e683836146cb565b60208301905092915050565b6000602082019050919050565b600061470a8261469f565b61471481856146aa565b935061471f836146bb565b8060005b8381101561475057815161473788826146da565b9750614742836146f2565b925050600181019050614723565b5085935050505092915050565b6000819050919050565b6147708161475d565b82525050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600381106147b6576147b5614776565b5b50565b60008190506147c7826147a5565b919050565b60006147d7826147b9565b9050919050565b6147e7816147cc565b82525050565b6000606083016000830151848203600086015261480a82826146ff565b915050602083015161481f6020860182614767565b50604083015161483260408601826147de565b508091505092915050565b6000602082019050818103600083015261485781846147ed565b905092915050565b6148688161475d565b811461487357600080fd5b50565b6000813590506148858161485f565b92915050565b600080604083850312156148a2576148a161435b565b5b60006148b085828601614573565b92505060206148c185828601614876565b9150509250929050565b6000602082840312156148e1576148e061435b565b5b60006148ef84828501614876565b91505092915050565b600067ffffffffffffffff82111561491357614912614380565b5b602082029050602081019050919050565b600080fd5b600061493c614937846148f8565b6143e0565b9050808382526020820190506020840283018581111561495f5761495e614924565b5b835b8181101561498857806149748882614573565b845260208401935050602081019050614961565b5050509392505050565b600082601f8301126149a7576149a6614365565b5b81356149b7848260208601614929565b91505092915050565b600067ffffffffffffffff8211156149db576149da614380565b5b602082029050602081019050919050565b60006149ff6149fa846149c0565b6143e0565b90508083825260208201905060208402830185811115614a2257614a21614924565b5b835b81811015614a4b5780614a378882614876565b845260208401935050602081019050614a24565b5050509392505050565b600082601f830112614a6a57614a69614365565b5b8135614a7a8482602086016149ec565b91505092915050565b60008060408385031215614a9a57614a9961435b565b5b600083013567ffffffffffffffff811115614ab857614ab7614360565b5b614ac485828601614992565b925050602083013567ffffffffffffffff811115614ae557614ae4614360565b5b614af185828601614a55565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b60005b83811015614b35578082015181840152602081019050614b1a565b60008484015250505050565b6000614b4c82614afb565b614b568185614b06565b9350614b66818560208601614b17565b614b6f8161436f565b840191505092915050565b60006020820190508181036000830152614b948184614b41565b905092915050565b60008060408385031215614bb357614bb261435b565b5b6000614bc185828601614573565b925050602083013567ffffffffffffffff811115614be257614be1614360565b5b614bee8582860161447d565b9150509250929050565b6000819050919050565b614c0b81614bf8565b82525050565b6000602082019050614c266000830184614c02565b92915050565b614c358161475d565b82525050565b6000602082019050614c506000830184614c2c565b92915050565b608082016000820151614c6c6000850182614767565b506020820151614c7f6020850182614767565b506040820151614c926040850182614767565b506060820151614ca56060850182614767565b50505050565b6000608082019050614cc06000830184614c56565b92915050565b6000819050919050565b614cd981614cc6565b82525050565b6000602082019050614cf46000830184614cd0565b92915050565b614d03816147cc565b82525050565b6000604082019050614d1e6000830185614c2c565b614d2b6020830184614cfa565b9392505050565b614d3b8161454a565b82525050565b6000602082019050614d566000830184614d32565b92915050565b614d6581614cc6565b82525050565b600082825260208201905092915050565b6000614d8782614afb565b614d918185614d6b565b9350614da1818560208601614b17565b614daa8161436f565b840191505092915050565b6000608083016000830151614dcd60008601826146cb565b506020830151614de06020860182614d5c565b506040830151614df36040860182614d5c565b5060608301518482036060860152614e0b8282614d7c565b9150508091505092915050565b60006020820190508181036000830152614e328184614db5565b905092915050565b600080600060608486031215614e5357614e5261435b565b5b600084013567ffffffffffffffff811115614e7157614e70614360565b5b614e7d86828701614628565b9350506020614e8e86828701614573565b9250506040614e9f86828701614876565b9150509250925092565b614eb2816144f4565b82525050565b606082016000820151614ece6000850182614767565b506020820151614ee16020850182614767565b506040820151614ef46040850182614ea9565b50505050565b6000606082019050614f0f6000830184614eb8565b92915050565b600067ffffffffffffffff821115614f3057614f2f614380565b5b602082029050602081019050919050565b614f4a81614cc6565b8114614f5557600080fd5b50565b600081359050614f6781614f41565b92915050565b6000614f80614f7b84614f15565b6143e0565b90508083825260208201905060208402830185811115614fa357614fa2614924565b5b835b81811015614fcc5780614fb88882614f58565b845260208401935050602081019050614fa5565b5050509392505050565b600082601f830112614feb57614fea614365565b5b8135614ffb848260208601614f6d565b91505092915050565b6000806040838503121561501b5761501a61435b565b5b600083013567ffffffffffffffff81111561503957615038614360565b5b61504585828601614992565b925050602083013567ffffffffffffffff81111561506657615065614360565b5b61507285828601614fd6565b9150509250929050565b60006060820190506150916000830186614c2c565b61509e6020830185614c2c565b6150ab6040830184614500565b949350505050565b600081905092915050565b60006150c982614afb565b6150d381856150b3565b93506150e3818560208601614b17565b80840191505092915050565b60006150fb82846150be565b915081905092915050565b7f4f6e6c79207468652061646d696e2063616e2063616c6c20746869732066756e60008201527f6374696f6e2e0000000000000000000000000000000000000000000000000000602082015250565b6000615162602683614b06565b915061516d82615106565b604082019050919050565b6000602082019050818103600083015261519181615155565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006151d28261475d565b91506151dd8361475d565b92508282019050808211156151f5576151f4615198565b5b92915050565b600061520682614cc6565b915061521183614cc6565b925082820390508181126000841216828213600085121516171561523857615237615198565b5b92915050565b60006152498261475d565b91506152548361475d565b925082820390508181111561526c5761526b615198565b5b92915050565b50565b6000615282600083614b06565b915061528d82615272565b600082019050919050565b60006040820190506152ad6000830184614c2c565b81810360208301526152be81615275565b905092915050565b7f576974686472617720616d6f756e7420616d6f756e742065786365656473207360008201527f6c617368656446756e6473000000000000000000000000000000000000000000602082015250565b6000615322602b83614b06565b915061532d826152c6565b604082019050919050565b6000602082019050818103600083015261535181615315565b9050919050565b600060408201905061536d6000830185614d32565b61537a6020830184614c2c565b9392505050565b61538a816144f4565b811461539557600080fd5b50565b6000815190506153a781615381565b92915050565b6000602082840312156153c3576153c261435b565b5b60006153d184828501615398565b91505092915050565b7f6370206163636f756e747320616e6420616d6f756e747320617265206469666660008201527f6572656e74206c656e6774687300000000000000000000000000000000000000602082015250565b6000615436602d83614b06565b9150615441826153da565b604082019050919050565b6000602082019050818103600083015261546581615429565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006154a682614cc6565b91506154b183614cc6565b9250828201905082811215600083121683821260008412151617156154d9576154d8615198565b5b92915050565b60006060820190506154f46000830186614d32565b6155016020830185614d32565b61550e6040830184614c2c565b949350505050565b7f7472616e73666572206661696c65640000000000000000000000000000000000600082015250565b600061554c600f83614b06565b915061555782615516565b602082019050919050565b6000602082019050818103600083015261557b8161553f565b9050919050565b600081519050919050565b600081905092915050565b60006155a382615582565b6155ad818561558d565b93506155bd818560208601614b17565b80840191505092915050565b60006155d58284615598565b915081905092915050565b7f4661696c656420746f2063616c6c206765744f776e65722066756e6374696f6e60008201527f206f662043504163636f756e7400000000000000000000000000000000000000602082015250565b600061563c602d83614b06565b9150615647826155e0565b604082019050919050565b6000602082019050818103600083015261566b8161562f565b9050919050565b600061567d8261452a565b9050919050565b61568d81615672565b811461569857600080fd5b50565b6000815190506156aa81615684565b92915050565b6000602082840312156156c6576156c561435b565b5b60006156d48482850161569b565b91505092915050565b7f63616c6c6572206973206e6f74206370206f776e657200000000000000000000600082015250565b6000615713601683614b06565b915061571e826156dd565b602082019050919050565b6000602082019050818103600083015261574281615706565b9050919050565b7f4e6f7420656e6f7567682066726f7a656e2062616c616e636520746f2077697460008201527f6864726177000000000000000000000000000000000000000000000000000000602082015250565b60006157a5602583614b06565b91506157b082615749565b604082019050919050565b600060208201905081810360008301526157d481615798565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061582257607f821691505b602082108103615835576158346157db565b5b50919050565b7f616d6f756e74206d7573742062652067726561746572207468616e2030000000600082015250565b6000615871601d83614b06565b915061587c8261583b565b602082019050919050565b600060208201905081810360008301526158a081615864565b9050919050565b7f6e6f7420656e6f7567682062616c616e63650000000000000000000000000000600082015250565b60006158dd601283614b06565b91506158e8826158a7565b602082019050919050565b6000602082019050818103600083015261590c816158d0565b9050919050565b7f4e6f2070656e64696e6720776974686472617720726571756573740000000000600082015250565b6000615949601b83614b06565b915061595482615913565b602082019050919050565b600060208201905081810360008301526159788161593c565b9050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b6000819050919050565b60006159c26159bd6159b88461597f565b61599d565b615989565b9050919050565b6159d2816159a7565b82525050565b60006020820190506159ed60008301846159c9565b92915050565b6000604082019050615a086000830185614c2c565b8181036020830152615a1a8184614b41565b90509392505050565b6000602082019050615a386000830184614cfa565b92915050565b7f57697468647261772064656c6179206e6f742070617373656400000000000000600082015250565b6000615a74601983614b06565b9150615a7f82615a3e565b602082019050919050565b60006020820190508181036000830152615aa381615a67565b9050919050565b7f776974686472617720616d6f756e74206d75737420626520706f736974697665600082015250565b6000615ae0602083614b06565b9150615aeb82615aaa565b602082019050919050565b60006020820190508181036000830152615b0f81615ad3565b9050919050565b7f6e6f7468696e6720746f20776974686472617700000000000000000000000000600082015250565b6000615b4c601383614b06565b9150615b5782615b16565b602082019050919050565b60006020820190508181036000830152615b7b81615b3f565b9050919050565b7f776974686472617720616d6f756e742067726561746572207468616e2061766160008201527f696c61626c652062616c616e6365000000000000000000000000000000000000602082015250565b6000615bde602e83614b06565b9150615be982615b82565b604082019050919050565b60006020820190508181036000830152615c0d81615bd1565b9050919050565b600081519050615c238161455c565b92915050565b600060208284031215615c3f57615c3e61435b565b5b6000615c4d84828501615c14565b91505092915050565b7f4f6e6c792043504163636f756e74206f776e65722063616e207769746864726160008201527f772074686520636f6c6c61746572616c2066756e647300000000000000000000602082015250565b6000615cb2603683614b06565b9150615cbd82615c56565b604082019050919050565b60006020820190508181036000830152615ce181615ca5565b9050919050565b6000615cf38261475d565b9150615cfe8361475d565b9250828202615d0c8161475d565b91508282048414831517615d2357615d22615198565b5b5092915050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302615d8c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82615d4f565b615d968683615d4f565b95508019841693508086168417925050509392505050565b6000615dc9615dc4615dbf8461475d565b61599d565b61475d565b9050919050565b6000819050919050565b615de383615dae565b615df7615def82615dd0565b848454615d5c565b825550505050565b600090565b615e0c615dff565b615e17818484615dda565b505050565b5b81811015615e3b57615e30600082615e04565b600181019050615e1d565b5050565b601f821115615e8057615e5181615d2a565b615e5a84615d3f565b81016020851015615e69578190505b615e7d615e7585615d3f565b830182615e1c565b50505b505050565b600082821c905092915050565b6000615ea360001984600802615e85565b1980831691505092915050565b6000615ebc8383615e92565b9150826002028217905092915050565b615ed582614afb565b67ffffffffffffffff811115615eee57615eed614380565b5b615ef8825461580a565b615f03828285615e3f565b600060209050601f831160018114615f365760008415615f24578287015190505b615f2e8582615eb0565b865550615f96565b601f198416615f4486615d2a565b60005b82811015615f6c57848901518255600182019150602085019450602081019050615f47565b86831015615f895784890151615f85601f891682615e92565b8355505b6001600288020188555050505b505050505050565b615fa781614bf8565b8114615fb257600080fd5b50565b600081519050615fc481615f9e565b92915050565b600060208284031215615fe057615fdf61435b565b5b6000615fee84828501615fb5565b9150509291505056fea26469706673582212200c4076ac6208f10be312f1202acd508a41015693a6643d5c3b0f182386fa32d364736f6c63430008190033",
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
