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

// SwanCreditCollateralMetaData contains all meta data concerning the SwanCreditCollateral contract.
var SwanCreditCollateralMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"frozenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"operation\",\"type\":\"string\"}],\"name\":\"CollateralAdjusted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"CollateralLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"CollateralSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumSwanCreditCollateral.CollateralStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"CollateralStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"CollateralUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fundingWallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fundingWallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"}],\"name\":\"DepositLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"DisputeProof\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SlashedFundsIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cpAccountAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralContratOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashfund\",\"type\":\"uint256\"}],\"name\":\"WithdrawSlash\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedToWithdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"availableBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"slashedFunds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashRatio\",\"type\":\"uint256\"}],\"internalType\":\"structSwanCreditCollateral.ContractInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAddress\",\"type\":\"address\"}],\"name\":\"cpInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"availableBalance\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"lockedBalance\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"}],\"internalType\":\"structSwanCreditCollateral.CPInfoWithLockedBalance\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cpStatus\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositLockedBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"disputeProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"escrowBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"getTaskInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"cpList\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"enumSwanCreditCollateral.CollateralStatus\",\"name\":\"collateralStatus\",\"type\":\"uint8\"}],\"internalType\":\"structSwanCreditCollateral.Task\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"isSignatureUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lockCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockedBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockedCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseCollateral\",\"type\":\"uint256\"}],\"name\":\"setBaseCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"setCollateralToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slashRatio\",\"type\":\"uint256\"}],\"name\":\"setSlashRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slashAmount\",\"type\":\"uint256\"}],\"name\":\"slashCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slashAmount\",\"type\":\"uint256\"}],\"name\":\"slashCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashedFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"enumSwanCreditCollateral.CollateralStatus\",\"name\":\"collateralStatus\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"unlockRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPending\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawSlashedFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801561004357600080fd5b5061005261005760201b60201c565b6101c1565b600061006761015b60201b60201c565b90508060000160089054906101000a900460ff16156100b2576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff80168160000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff16146101585767ffffffffffffffff8160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d267ffffffffffffffff60405161014f91906101a6565b60405180910390a15b50565b60007ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b600067ffffffffffffffff82169050919050565b6101a081610183565b82525050565b60006020820190506101bb6000830184610197565b92915050565b6080516149786101ea60003960008181612bdb01528181612c300152612deb01526149786000f3fe6080604052600436106102305760003560e01c8063704802751161012e578063ad3cb1cc116100ab578063e80d9dec1161006f578063e80d9dec14610867578063e9a9ee8314610890578063f2fde38b146108b9578063f3fef3a3146108e2578063f97348e11461090b57610230565b8063ad3cb1cc14610794578063b4eae1cb146107bf578063b587b82c146107ea578063ce3518aa14610813578063d27ca89b1461083c57610230565b806392bdf9ba116100f257806392bdf9ba146106755780639ae697bf146106b25780639b5ddf09146106ef578063a0821be31461071a578063a664c2161461075757610230565b806370480275146105ca578063715018a6146105f35780637f58a7e51461060a5780638129fc1c146106335780638da5cb5b1461064a57610230565b806347e7ef24116101bc57806354fd4d501161018057806354fd4d50146104d057806355af6353146104fb57806358709cf214610538578063666181a9146105765780636f99f15c1461059f57610230565b806347e7ef241461040a5780634f1ef2861461043357806352d1902d1461044f578063536f60701461047a57806353ad8720146104a557610230565b806324d7806c1161020357806324d7806c146103015780632894493f1461033e5780632d291cad146103675780633fe65177146103a457806347a7d107146103e157610230565b80631150f0f3146102355780631785f53c146102725780631b2094631461029b5780631d47a62d146102d8575b600080fd5b34801561024157600080fd5b5061025c600480360381019061025791906135cc565b61094a565b6040516102699190613630565b60405180910390f35b34801561027e57600080fd5b50610299600480360381019061029491906136a9565b610980565b005b3480156102a757600080fd5b506102c260048036038101906102bd9190613777565b6109e3565b6040516102cf919061395e565b60405180910390f35b3480156102e457600080fd5b506102ff60048036038101906102fa91906139ac565b610aee565b005b34801561030d57600080fd5b50610328600480360381019061032391906136a9565b6110b9565b6040516103359190613630565b60405180910390f35b34801561034a57600080fd5b50610365600480360381019061036091906139ec565b6110d9565b005b34801561037357600080fd5b5061038e600480360381019061038991906136a9565b611230565b60405161039b9190613630565b60405180910390f35b3480156103b057600080fd5b506103cb60048036038101906103c691906136a9565b611250565b6040516103d89190613a98565b60405180910390f35b3480156103ed57600080fd5b50610408600480360381019061040391906139ac565b6112f0565b005b34801561041657600080fd5b50610431600480360381019061042c91906139ac565b611548565b005b61044d60048036038101906104489190613aba565b6116b3565b005b34801561045b57600080fd5b506104646116d2565b6040516104719190613b2f565b60405180910390f35b34801561048657600080fd5b5061048f611705565b60405161049c9190613b59565b60405180910390f35b3480156104b157600080fd5b506104ba61170b565b6040516104c79190613bc9565b60405180910390f35b3480156104dc57600080fd5b506104e5611741565b6040516104f29190613b59565b60405180910390f35b34801561050757600080fd5b50610522600480360381019061051d91906136a9565b61174a565b60405161052f9190613bfd565b60405180910390f35b34801561054457600080fd5b5061055f600480360381019061055a9190613777565b611762565b60405161056d929190613c27565b60405180910390f35b34801561058257600080fd5b5061059d600480360381019061059891906136a9565b6117a9565b005b3480156105ab57600080fd5b506105b46117f5565b6040516105c19190613b59565b60405180910390f35b3480156105d657600080fd5b506105f160048036038101906105ec91906136a9565b6117fb565b005b3480156105ff57600080fd5b5061060861185e565b005b34801561061657600080fd5b50610631600480360381019061062c91906139ec565b611872565b005b34801561063f57600080fd5b50610648611908565b005b34801561065657600080fd5b5061065f611b15565b60405161066c9190613c5f565b60405180910390f35b34801561068157600080fd5b5061069c600480360381019061069791906136a9565b611b4d565b6040516106a99190613b59565b60405180910390f35b3480156106be57600080fd5b506106d960048036038101906106d491906136a9565b611b65565b6040516106e69190613bfd565b60405180910390f35b3480156106fb57600080fd5b50610704611b7d565b6040516107119190613b59565b60405180910390f35b34801561072657600080fd5b50610741600480360381019061073c91906136a9565b611b83565b60405161074e9190613bfd565b60405180910390f35b34801561076357600080fd5b5061077e600480360381019061077991906136a9565b611b9b565b60405161078b9190613d36565b60405180910390f35b3480156107a057600080fd5b506107a9611d26565b6040516107b69190613a98565b60405180910390f35b3480156107cb57600080fd5b506107d4611d5f565b6040516107e19190613b59565b60405180910390f35b3480156107f657600080fd5b50610811600480360381019061080c9190613d58565b611d65565b005b34801561081f57600080fd5b5061083a600480360381019061083591906139ec565b6123d1565b005b34801561084857600080fd5b506108516123e3565b60405161085e9190613b59565b60405180910390f35b34801561087357600080fd5b5061088e60048036038101906108899190613777565b6123e9565b005b34801561089c57600080fd5b506108b760048036038101906108b291906139ac565b612445565b005b3480156108c557600080fd5b506108e060048036038101906108db91906136a9565b6125a7565b005b3480156108ee57600080fd5b50610909600480360381019061090491906139ac565b61262d565b005b34801561091757600080fd5b50610932600480360381019061092d91906136a9565b6129b9565b60405161094193929190613dc7565b60405180910390f35b600b818051602081018201805184825260208301602085012081835280955050505050506000915054906101000a900460ff1681565b6109886129f0565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b6109eb6133d9565b6005826040516109fb9190613e3a565b908152602001604051809103902060405180606001604052908160008201805480602002602001604051908101604052809291908181526020018280548015610a9957602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610a4f575b50505050508152602001600182015481526020016002820160009054906101000a900460ff166002811115610ad157610ad0613897565b5b6002811115610ae357610ae2613897565b5b815250509050919050565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610b7a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b7190613ec3565b60405180910390fd5b601060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548113610c6a577fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b81604051610bef9190613b59565b60405180910390a180600080828254610c089190613f12565b9250508190555080601060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610c5e9190613f46565b9250508190555061105e565b60008190506000601060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541315610da157601060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600080828254610d079190613f12565b92505081905550601060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481610d599190613f46565b90506000601060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b80600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205412610e915780600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610e369190613f46565b9250508190555080600080828254610e4e9190613f12565b925050819055507fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b82604051610e849190613b59565b60405180910390a161105c565b6000600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541315610fc357600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481610f249190613f46565b9050600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600080828254610f769190613f12565b925050819055506000600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b7fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b8183610ff09190613f89565b604051610ffd9190613b59565b60405180910390a180601060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546110549190613f46565b925050819055505b505b61106782612a77565b8173ffffffffffffffffffffffffffffffffffffffff167f403feb2cd85cc25c910fe59289105b583f08fe9e4335ebbb08c1851f47ff2d34826040516110ad9190613fe3565b60405180910390a25050565b60046020528060005260406000206000915054906101000a900460ff1681565b6110e16129f0565b806000541015611126576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161111d90614083565b60405180910390fd5b806000808282546111379190613f89565b92505081905550600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b815260040161119b9291906140a3565b6020604051808303816000875af11580156111ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111de91906140f8565b503373ffffffffffffffffffffffffffffffffffffffff167fbfd9c82485e2178fcfd5c40379d8e66fe60acc50aa1ef4c50966431eb1e381cd826040516112259190613b59565b60405180910390a250565b600a6020528060005260406000206000915054906101000a900460ff1681565b6008602052806000526040600020600091509050805461126f90614154565b80601f016020809104026020016040519081016040528092919081815260200182805461129b90614154565b80156112e85780601f106112bd576101008083540402835291602001916112e8565b820191906000526020600020905b8154815290600101906020018083116112cb57829003601f168201915b505050505081565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661137c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161137390613ec3565b60405180910390fd5b600081116113bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113b6906141d1565b60405180910390fd5b600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054811315611441576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114389061423d565b60405180910390fd5b80600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546114909190613f46565b9250508190555080601060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546114e6919061425d565b925050819055506114f682612a77565b8173ffffffffffffffffffffffffffffffffffffffff167f5f3d004cf9164b95ed5dbf47d1f04018a4eabcb20b4320fe229ed92236ace6348260405161153c9190613fe3565b60405180910390a25050565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b81526004016115a7939291906142a1565b6020604051808303816000875af11580156115c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115ea91906140f8565b5080600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461163a919061425d565b925050819055508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f628360405161169e9190613b59565b60405180910390a36116af82612a77565b5050565b6116bb612bd9565b6116c482612cbf565b6116ce8282612cca565b5050565b60006116dc612de9565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b905090565b600e5481565b61171361340c565b6040518060800160405280600054815260200160015481526020016002548152602001600354815250905090565b60006002905090565b600c6020528060005260406000206000915090505481565b6005818051602081018201805184825260208301602085012081835280955050505050506000915090508060010154908060020160009054906101000a900460ff16905082565b6117b16129f0565b80600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60005481565b6118036129f0565b6001600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b6118666129f0565b6118706000612e70565b565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166118fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118f590613ec3565b60405180910390fd5b8060018190555050565b6000611912612f47565b905060008160000160089054906101000a900460ff1615905060008260000160009054906101000a900467ffffffffffffffff1690506000808267ffffffffffffffff161480156119605750825b9050600060018367ffffffffffffffff16148015611995575060003073ffffffffffffffffffffffffffffffffffffffff163b145b9050811580156119a3575080155b156119da576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018560000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508315611a2a5760018560000160086101000a81548160ff0219169083151502179055505b611a3333612f6f565b611a3b612f83565b6001600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600281905550600160038190555067016345785d8a00006001819055508315611b0e5760008560000160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d26001604051611b059190614331565b60405180910390a15b5050505050565b600080611b20612f8d565b90508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505090565b60076020528060005260406000206000915090505481565b60106020528060005260406000206000915090505481565b60015481565b60066020528060005260406000206000915090505481565b611ba3613434565b60405180608001604052808373ffffffffffffffffffffffffffffffffffffffff168152602001600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548152602001601060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548152602001600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054611c9e90614154565b80601f0160208091040260200160405190810160405280929190818152602001828054611cca90614154565b8015611d175780601f10611cec57610100808354040283529160200191611d17565b820191906000526020600020905b815481529060010190602001808311611cfa57829003601f168201915b50505050508152509050919050565b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b60025481565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611df1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611de890613ec3565b60405180910390fd5b6000600584604051611e039190613e3a565b90815260200160405180910390209050601060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548213611f03577fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b82604051611e889190613b59565b60405180910390a181600080828254611ea19190613f12565b9250508190555081601060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611ef79190613f46565b925050819055506122f7565b60008290506000601060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054131561203a57601060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600080828254611fa09190613f12565b92505081905550601060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481611ff29190613f46565b90506000601060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b80600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541261212a5780600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546120cf9190613f46565b92505081905550806000808282546120e79190613f12565b925050819055507fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b8360405161211d9190613b59565b60405180910390a16122f5565b6000600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054131561225c57600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054816121bd9190613f46565b9050600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205460008082825461220f9190613f12565b925050819055506000600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b7fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b81846122899190613f89565b6040516122969190613b59565b60405180910390a180601060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546122ed9190613f46565b925050819055505b505b60028160020160006101000a81548160ff0219169083600281111561231f5761231e613897565b5b021790555061232d83612a77565b8273ffffffffffffffffffffffffffffffffffffffff167f403feb2cd85cc25c910fe59289105b583f08fe9e4335ebbb08c1851f47ff2d34838660405161237592919061434c565b60405180910390a28360405161238b9190613e3a565b60405180910390207f4a2ced9ada462e244851a86e998eb0b5bf558c2c9c6923b7f970ed2b19b073b060026040516123c3919061437c565b60405180910390a250505050565b6123d96129f0565b8060038190555050565b60035481565b806040516123f79190613e3a565b60405180910390203373ffffffffffffffffffffffffffffffffffffffff167faec1d412a3c1e4a13fc2a2e19ac38a5af192a9cf17b074fca8146a2d0655e0c360405160405180910390a350565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b81526004016124a4939291906142a1565b6020604051808303816000875af11580156124c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124e791906140f8565b5080601060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254612537919061425d565b925050819055508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f18b5d63f2a4f63f9d724d087c9580fdb3c4501bf641ce600afe4b97282e738e28360405161259b9190613b59565b60405180910390a35050565b6125af6129f0565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036126215760006040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016126189190613c5f565b60405180910390fd5b61262a81612e70565b50565b60008111612670576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612667906143e3565b60405180910390fd5b6000600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054136126f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016126e99061444f565b60405180910390fd5b600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054811115612774576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161276b906144e1565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1663893d20e86040518163ffffffff1660e01b8152600401602060405180830381865afa1580156127d6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127fa9190614516565b73ffffffffffffffffffffffffffffffffffffffff1614612850576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612847906145b5565b60405180910390fd5b80600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461289f9190613f46565b92505081905550600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b81526004016129039291906140a3565b6020604051808303816000875af1158015612922573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061294691906140f8565b5061295082612a77565b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb836040516129ad9190613b59565b60405180910390a35050565b600d6020528060005260406000206000915090508060000154908060010154908060020160009054906101000a900460ff16905083565b6129f8612fb5565b73ffffffffffffffffffffffffffffffffffffffff16612a16611b15565b73ffffffffffffffffffffffffffffffffffffffff1614612a7557612a39612fb5565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401612a6c9190613c5f565b60405180910390fd5b565b600154600254612a8791906145d5565b601060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205412612b53576040518060400160405280600981526020017f7a6b41756374696f6e0000000000000000000000000000000000000000000000815250600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209081612b4d91906147b9565b50612bd6565b6040518060400160405280600381526020017f4e53430000000000000000000000000000000000000000000000000000000000815250600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209081612bd491906147b9565b505b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161480612c8657507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16612c6d612fbd565b73ffffffffffffffffffffffffffffffffffffffff1614155b15612cbd576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b612cc76129f0565b50565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612d3257506040513d601f19601f82011682018060405250810190612d2f91906148b7565b60015b612d7357816040517f4c9c8ce3000000000000000000000000000000000000000000000000000000008152600401612d6a9190613c5f565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b8114612dda57806040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600401612dd19190613b2f565b60405180910390fd5b612de48383613014565b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614612e6e576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6000612e7a612f8d565b905060008160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050828260000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3505050565b60007ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b612f77613087565b612f80816130c7565b50565b612f8b613087565b565b60007f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300905090565b600033905090565b6000612feb7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b61314d565b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61301d82613157565b8173ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a260008151111561307a576130748282613224565b50613083565b6130826132a8565b5b5050565b61308f6132e5565b6130c5576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6130cf613087565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036131415760006040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016131389190613c5f565b60405180910390fd5b61314a81612e70565b50565b6000819050919050565b60008173ffffffffffffffffffffffffffffffffffffffff163b036131b357806040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526004016131aa9190613c5f565b60405180910390fd5b806131e07f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b61314d565b60000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60606000808473ffffffffffffffffffffffffffffffffffffffff168460405161324e919061492b565b600060405180830381855af49150503d8060008114613289576040519150601f19603f3d011682016040523d82523d6000602084013e61328e565b606091505b509150915061329e858383613305565b9250505092915050565b60003411156132e3576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60006132ef612f47565b60000160089054906101000a900460ff16905090565b60608261331a5761331582613394565b61338c565b60008251148015613342575060008473ffffffffffffffffffffffffffffffffffffffff163b145b1561338457836040517f9996b31500000000000000000000000000000000000000000000000000000000815260040161337b9190613c5f565b60405180910390fd5b81905061338d565b5b9392505050565b6000815111156133a75780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604051806060016040528060608152602001600081526020016000600281111561340657613405613897565b5b81525090565b6040518060800160405280600081526020016000815260200160008152602001600081525090565b6040518060800160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001606081525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6134d982613490565b810181811067ffffffffffffffff821117156134f8576134f76134a1565b5b80604052505050565b600061350b613472565b905061351782826134d0565b919050565b600067ffffffffffffffff821115613537576135366134a1565b5b61354082613490565b9050602081019050919050565b82818337600083830152505050565b600061356f61356a8461351c565b613501565b90508281526020810184848401111561358b5761358a61348b565b5b61359684828561354d565b509392505050565b600082601f8301126135b3576135b2613486565b5b81356135c384826020860161355c565b91505092915050565b6000602082840312156135e2576135e161347c565b5b600082013567ffffffffffffffff811115613600576135ff613481565b5b61360c8482850161359e565b91505092915050565b60008115159050919050565b61362a81613615565b82525050565b60006020820190506136456000830184613621565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006136768261364b565b9050919050565b6136868161366b565b811461369157600080fd5b50565b6000813590506136a38161367d565b92915050565b6000602082840312156136bf576136be61347c565b5b60006136cd84828501613694565b91505092915050565b600067ffffffffffffffff8211156136f1576136f06134a1565b5b6136fa82613490565b9050602081019050919050565b600061371a613715846136d6565b613501565b9050828152602081018484840111156137365761373561348b565b5b61374184828561354d565b509392505050565b600082601f83011261375e5761375d613486565b5b813561376e848260208601613707565b91505092915050565b60006020828403121561378d5761378c61347c565b5b600082013567ffffffffffffffff8111156137ab576137aa613481565b5b6137b784828501613749565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6137f58161366b565b82525050565b600061380783836137ec565b60208301905092915050565b6000602082019050919050565b600061382b826137c0565b61383581856137cb565b9350613840836137dc565b8060005b8381101561387157815161385888826137fb565b975061386383613813565b925050600181019050613844565b5085935050505092915050565b6000819050919050565b6138918161387e565b82525050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600381106138d7576138d6613897565b5b50565b60008190506138e8826138c6565b919050565b60006138f8826138da565b9050919050565b613908816138ed565b82525050565b6000606083016000830151848203600086015261392b8282613820565b91505060208301516139406020860182613888565b50604083015161395360408601826138ff565b508091505092915050565b60006020820190508181036000830152613978818461390e565b905092915050565b6139898161387e565b811461399457600080fd5b50565b6000813590506139a681613980565b92915050565b600080604083850312156139c3576139c261347c565b5b60006139d185828601613694565b92505060206139e285828601613997565b9150509250929050565b600060208284031215613a0257613a0161347c565b5b6000613a1084828501613997565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015613a53578082015181840152602081019050613a38565b60008484015250505050565b6000613a6a82613a19565b613a748185613a24565b9350613a84818560208601613a35565b613a8d81613490565b840191505092915050565b60006020820190508181036000830152613ab28184613a5f565b905092915050565b60008060408385031215613ad157613ad061347c565b5b6000613adf85828601613694565b925050602083013567ffffffffffffffff811115613b0057613aff613481565b5b613b0c8582860161359e565b9150509250929050565b6000819050919050565b613b2981613b16565b82525050565b6000602082019050613b446000830184613b20565b92915050565b613b538161387e565b82525050565b6000602082019050613b6e6000830184613b4a565b92915050565b608082016000820151613b8a6000850182613888565b506020820151613b9d6020850182613888565b506040820151613bb06040850182613888565b506060820151613bc36060850182613888565b50505050565b6000608082019050613bde6000830184613b74565b92915050565b6000819050919050565b613bf781613be4565b82525050565b6000602082019050613c126000830184613bee565b92915050565b613c21816138ed565b82525050565b6000604082019050613c3c6000830185613b4a565b613c496020830184613c18565b9392505050565b613c598161366b565b82525050565b6000602082019050613c746000830184613c50565b92915050565b613c8381613be4565b82525050565b600082825260208201905092915050565b6000613ca582613a19565b613caf8185613c89565b9350613cbf818560208601613a35565b613cc881613490565b840191505092915050565b6000608083016000830151613ceb60008601826137ec565b506020830151613cfe6020860182613c7a565b506040830151613d116040860182613c7a565b5060608301518482036060860152613d298282613c9a565b9150508091505092915050565b60006020820190508181036000830152613d508184613cd3565b905092915050565b600080600060608486031215613d7157613d7061347c565b5b600084013567ffffffffffffffff811115613d8f57613d8e613481565b5b613d9b86828701613749565b9350506020613dac86828701613694565b9250506040613dbd86828701613997565b9150509250925092565b6000606082019050613ddc6000830186613b4a565b613de96020830185613b4a565b613df66040830184613621565b949350505050565b600081905092915050565b6000613e1482613a19565b613e1e8185613dfe565b9350613e2e818560208601613a35565b80840191505092915050565b6000613e468284613e09565b915081905092915050565b7f4f6e6c79207468652061646d696e2063616e2063616c6c20746869732066756e60008201527f6374696f6e2e0000000000000000000000000000000000000000000000000000602082015250565b6000613ead602683613a24565b9150613eb882613e51565b604082019050919050565b60006020820190508181036000830152613edc81613ea0565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000613f1d8261387e565b9150613f288361387e565b9250828201905080821115613f4057613f3f613ee3565b5b92915050565b6000613f5182613be4565b9150613f5c83613be4565b9250828203905081811260008412168282136000851215161715613f8357613f82613ee3565b5b92915050565b6000613f948261387e565b9150613f9f8361387e565b9250828203905081811115613fb757613fb6613ee3565b5b92915050565b50565b6000613fcd600083613a24565b9150613fd882613fbd565b600082019050919050565b6000604082019050613ff86000830184613b4a565b818103602083015261400981613fc0565b905092915050565b7f576974686472617720616d6f756e7420616d6f756e742065786365656473207360008201527f6c617368656446756e6473000000000000000000000000000000000000000000602082015250565b600061406d602b83613a24565b915061407882614011565b604082019050919050565b6000602082019050818103600083015261409c81614060565b9050919050565b60006040820190506140b86000830185613c50565b6140c56020830184613b4a565b9392505050565b6140d581613615565b81146140e057600080fd5b50565b6000815190506140f2816140cc565b92915050565b60006020828403121561410e5761410d61347c565b5b600061411c848285016140e3565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061416c57607f821691505b60208210810361417f5761417e614125565b5b50919050565b7f616d6f756e74206d7573742062652067726561746572207468616e2030000000600082015250565b60006141bb601d83613a24565b91506141c682614185565b602082019050919050565b600060208201905081810360008301526141ea816141ae565b9050919050565b7f6e6f7420656e6f7567682062616c616e63650000000000000000000000000000600082015250565b6000614227601283613a24565b9150614232826141f1565b602082019050919050565b600060208201905081810360008301526142568161421a565b9050919050565b600061426882613be4565b915061427383613be4565b92508282019050828112156000831216838212600084121516171561429b5761429a613ee3565b5b92915050565b60006060820190506142b66000830186613c50565b6142c36020830185613c50565b6142d06040830184613b4a565b949350505050565b6000819050919050565b600067ffffffffffffffff82169050919050565b6000819050919050565b600061431b614316614311846142d8565b6142f6565b6142e2565b9050919050565b61432b81614300565b82525050565b60006020820190506143466000830184614322565b92915050565b60006040820190506143616000830185613b4a565b81810360208301526143738184613a5f565b90509392505050565b60006020820190506143916000830184613c18565b92915050565b7f776974686472617720616d6f756e74206d75737420626520706f736974697665600082015250565b60006143cd602083613a24565b91506143d882614397565b602082019050919050565b600060208201905081810360008301526143fc816143c0565b9050919050565b7f6e6f7468696e6720746f20776974686472617700000000000000000000000000600082015250565b6000614439601383613a24565b915061444482614403565b602082019050919050565b600060208201905081810360008301526144688161442c565b9050919050565b7f776974686472617720616d6f756e742067726561746572207468616e2061766160008201527f696c61626c652062616c616e6365000000000000000000000000000000000000602082015250565b60006144cb602e83613a24565b91506144d68261446f565b604082019050919050565b600060208201905081810360008301526144fa816144be565b9050919050565b6000815190506145108161367d565b92915050565b60006020828403121561452c5761452b61347c565b5b600061453a84828501614501565b91505092915050565b7f4f6e6c792043504163636f756e74206f776e65722063616e207769746864726160008201527f772074686520636f6c6c61746572616c2066756e647300000000000000000000602082015250565b600061459f603683613a24565b91506145aa82614543565b604082019050919050565b600060208201905081810360008301526145ce81614592565b9050919050565b60006145e08261387e565b91506145eb8361387e565b92508282026145f98161387e565b915082820484148315176146105761460f613ee3565b5b5092915050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026146797fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261463c565b614683868361463c565b95508019841693508086168417925050509392505050565b60006146b66146b16146ac8461387e565b6142f6565b61387e565b9050919050565b6000819050919050565b6146d08361469b565b6146e46146dc826146bd565b848454614649565b825550505050565b600090565b6146f96146ec565b6147048184846146c7565b505050565b5b818110156147285761471d6000826146f1565b60018101905061470a565b5050565b601f82111561476d5761473e81614617565b6147478461462c565b81016020851015614756578190505b61476a6147628561462c565b830182614709565b50505b505050565b600082821c905092915050565b600061479060001984600802614772565b1980831691505092915050565b60006147a9838361477f565b9150826002028217905092915050565b6147c282613a19565b67ffffffffffffffff8111156147db576147da6134a1565b5b6147e58254614154565b6147f082828561472c565b600060209050601f8311600181146148235760008415614811578287015190505b61481b858261479d565b865550614883565b601f19841661483186614617565b60005b8281101561485957848901518255600182019150602085019450602081019050614834565b868310156148765784890151614872601f89168261477f565b8355505b6001600288020188555050505b505050505050565b61489481613b16565b811461489f57600080fd5b50565b6000815190506148b18161488b565b92915050565b6000602082840312156148cd576148cc61347c565b5b60006148db848285016148a2565b91505092915050565b600081519050919050565b600081905092915050565b6000614905826148e4565b61490f81856148ef565b935061491f818560208601613a35565b80840191505092915050565b600061493782846148fa565b91508190509291505056fea264697066735822122021707fe08f2db41ed9dd5e268c1cf57c4f65de646a5cb848b078e21a1fb3de8f64736f6c63430008190033",
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

// DepositLockedBalance is a paid mutator transaction binding the contract method 0xe9a9ee83.
//
// Solidity: function depositLockedBalance(address cpAccount, uint256 amount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactor) DepositLockedBalance(opts *bind.TransactOpts, cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.contract.Transact(opts, "depositLockedBalance", cpAccount, amount)
}

// DepositLockedBalance is a paid mutator transaction binding the contract method 0xe9a9ee83.
//
// Solidity: function depositLockedBalance(address cpAccount, uint256 amount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralSession) DepositLockedBalance(cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.DepositLockedBalance(&_SwanCreditCollateral.TransactOpts, cpAccount, amount)
}

// DepositLockedBalance is a paid mutator transaction binding the contract method 0xe9a9ee83.
//
// Solidity: function depositLockedBalance(address cpAccount, uint256 amount) returns()
func (_SwanCreditCollateral *SwanCreditCollateralTransactorSession) DepositLockedBalance(cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SwanCreditCollateral.Contract.DepositLockedBalance(&_SwanCreditCollateral.TransactOpts, cpAccount, amount)
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
