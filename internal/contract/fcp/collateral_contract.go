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

// FcpCollateralMetaData contains all meta data concerning the FcpCollateral contract.
var FcpCollateralMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"frozenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"operation\",\"type\":\"string\"}],\"name\":\"CollateralAdjusted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"CollateralLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"CollateralSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumSwanCreditCollateral.CollateralStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"CollateralStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"CollateralUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fundingWallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"DisputeProof\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SlashedFundsIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cpAccountAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralContratOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashfund\",\"type\":\"uint256\"}],\"name\":\"WithdrawSlash\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedToWithdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"}],\"name\":\"approveWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"availableBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"cpList\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchUnlockCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"slashedFunds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashRatio\",\"type\":\"uint256\"}],\"internalType\":\"structSwanCreditCollateral.ContractInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAddress\",\"type\":\"address\"}],\"name\":\"cpInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"availableBalance\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"lockedCollateral\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"}],\"internalType\":\"structSwanCreditCollateral.CPInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cpStatus\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"disputeProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"cpList\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"emitCollateralUnlockedEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"}],\"name\":\"getTaskInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"cpList\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"enumSwanCreditCollateral.CollateralStatus\",\"name\":\"collateralStatus\",\"type\":\"uint8\"}],\"internalType\":\"structSwanCreditCollateral.Task\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"isSignatureUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"cpList\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"taskCollateral\",\"type\":\"uint256\"}],\"name\":\"lockCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockedCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"}],\"name\":\"requestWithdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseCollateral\",\"type\":\"uint256\"}],\"name\":\"setBaseCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"setCollateralToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slashRatio\",\"type\":\"uint256\"}],\"name\":\"setSlashRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slashAmount\",\"type\":\"uint256\"}],\"name\":\"slashCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slashAmount\",\"type\":\"uint256\"}],\"name\":\"slashCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashedFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"subtractBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"enumSwanCreditCollateral.CollateralStatus\",\"name\":\"collateralStatus\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"taskUid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unlockCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawSlashedFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801561004357600080fd5b5061005261005760201b60201c565b6101c1565b600061006761015b60201b60201c565b90508060000160089054906101000a900460ff16156100b2576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff80168160000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff16146101585767ffffffffffffffff8160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d267ffffffffffffffff60405161014f91906101a6565b60405180910390a15b50565b60007ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b600067ffffffffffffffff82169050919050565b6101a081610183565b82525050565b60006020820190506101bb6000830184610197565b92915050565b608051614e096101ea60003960008181612caf01528181612d040152612ebf0152614e096000f3fe6080604052600436106102465760003560e01c80637048027511610139578063a664c216116100b6578063cfd4f18a1161007a578063cfd4f18a1461086d578063d27ca89b14610896578063e02a63b5146108c1578063e80d9dec146108ea578063f2fde38b14610913578063f3fef3a31461093c57610246565b8063a664c21614610788578063ad3cb1cc146107c5578063b4eae1cb146107f0578063b587b82c1461081b578063ce3518aa1461084457610246565b80638da5cb5b116100fd5780638da5cb5b1461069c57806392bdf9ba146106c75780639b5ddf0914610704578063a0821be31461072f578063a35a36e91461076c57610246565b806370480275146105f3578063715018a61461061c5780637f58a7e5146106335780638129fc1c1461065c578063853911b21461067357610246565b80634b9f0bea116101c757806358709cf21161018b57806358709cf21461050f57806364a197f31461054d5780636579e35c14610576578063666181a91461059f5780636f99f15c146105c857610246565b80634b9f0bea146104495780634f1ef2861461047257806352d1902d1461048e57806353ad8720146104b957806354fd4d50146104e457610246565b80632894493f1161020e5780632894493f146103545780632d291cad1461037d5780632d45ac91146103ba5780633fe65177146103e357806347e7ef241461042057610246565b80631150f0f31461024b5780631785f53c146102885780631b209463146102b15780631d47a62d146102ee57806324d7806c14610317575b600080fd5b34801561025757600080fd5b50610272600480360381019061026d9190613747565b610965565b60405161027f91906137ab565b60405180910390f35b34801561029457600080fd5b506102af60048036038101906102aa9190613824565b61099b565b005b3480156102bd57600080fd5b506102d860048036038101906102d391906138f2565b6109fe565b6040516102e59190613ad9565b60405180910390f35b3480156102fa57600080fd5b5061031560048036038101906103109190613b27565b610b09565b005b34801561032357600080fd5b5061033e60048036038101906103399190613824565b610daf565b60405161034b91906137ab565b60405180910390f35b34801561036057600080fd5b5061037b60048036038101906103769190613b67565b610dcf565b005b34801561038957600080fd5b506103a4600480360381019061039f9190613824565b610f26565b6040516103b191906137ab565b60405180910390f35b3480156103c657600080fd5b506103e160048036038101906103dc9190613d1f565b610f46565b005b3480156103ef57600080fd5b5061040a60048036038101906104059190613824565b611125565b6040516104179190613e16565b60405180910390f35b34801561042c57600080fd5b5061044760048036038101906104429190613b27565b6111c5565b005b34801561045557600080fd5b50610470600480360381019061046b9190613e38565b611330565b005b61048c60048036038101906104879190613e94565b6116bc565b005b34801561049a57600080fd5b506104a36116db565b6040516104b09190613f09565b60405180910390f35b3480156104c557600080fd5b506104ce61170e565b6040516104db9190613f79565b60405180910390f35b3480156104f057600080fd5b506104f9611744565b6040516105069190613fa3565b60405180910390f35b34801561051b57600080fd5b50610536600480360381019061053191906138f2565b61174d565b604051610544929190613fcd565b60405180910390f35b34801561055957600080fd5b50610574600480360381019061056f9190613b27565b611794565b005b34801561058257600080fd5b5061059d60048036038101906105989190613824565b6117e7565b005b3480156105ab57600080fd5b506105c660048036038101906105c19190613824565b6118ce565b005b3480156105d457600080fd5b506105dd61191a565b6040516105ea9190613fa3565b60405180910390f35b3480156105ff57600080fd5b5061061a60048036038101906106159190613824565b611920565b005b34801561062857600080fd5b50610631611983565b005b34801561063f57600080fd5b5061065a60048036038101906106559190613b67565b611997565b005b34801561066857600080fd5b50610671611a2d565b005b34801561067f57600080fd5b5061069a60048036038101906106959190613ff6565b611c3a565b005b3480156106a857600080fd5b506106b1611fb2565b6040516106be9190614090565b60405180910390f35b3480156106d357600080fd5b506106ee60048036038101906106e99190613824565b611fea565b6040516106fb9190613fa3565b60405180910390f35b34801561071057600080fd5b50610719612002565b6040516107269190613fa3565b60405180910390f35b34801561073b57600080fd5b5061075660048036038101906107519190613824565b612008565b60405161076391906140c4565b60405180910390f35b61078660048036038101906107819190613824565b612020565b005b34801561079457600080fd5b506107af60048036038101906107aa9190613824565b612023565b6040516107bc919061419b565b60405180910390f35b3480156107d157600080fd5b506107da6121ae565b6040516107e79190613e16565b60405180910390f35b3480156107fc57600080fd5b506108056121e7565b6040516108129190613fa3565b60405180910390f35b34801561082757600080fd5b50610842600480360381019061083d91906141bd565b6121ed565b005b34801561085057600080fd5b5061086b60048036038101906108669190613b67565b612534565b005b34801561087957600080fd5b50610894600480360381019061088f9190613b27565b612546565b005b3480156108a257600080fd5b506108ab6125a8565b6040516108b89190613fa3565b60405180910390f35b3480156108cd57600080fd5b506108e860048036038101906108e39190613d1f565b6125ae565b005b3480156108f657600080fd5b50610911600480360381019061090c91906138f2565b612656565b005b34801561091f57600080fd5b5061093a60048036038101906109359190613824565b6126b2565b005b34801561094857600080fd5b50610963600480360381019061095e9190613b27565b612738565b005b600b818051602081018201805184825260208301602085012081835280955050505050506000915054906101000a900460ff1681565b6109a3612ac4565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b610a066134ad565b600582604051610a169190614268565b908152602001604051809103902060405180606001604052908160008201805480602002602001604051908101604052809291908181526020018280548015610ab457602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610a6a575b50505050508152602001600182015481526020016002820160009054906101000a900460ff166002811115610aec57610aeb613a12565b5b6002811115610afe57610afd613a12565b5b815250509050919050565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610b95576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b8c906142f1565b60405180910390fd5b600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054811315610cae577fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054604051610c4a9190613fa3565b60405180910390a1600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600080828254610ca29190614340565b92505081905550610cfe565b7fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b81604051610cdd9190613fa3565b60405180910390a180600080828254610cf69190614340565b925050819055505b80600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610d4d9190614374565b92505081905550610d5d82612b4b565b8173ffffffffffffffffffffffffffffffffffffffff167f403feb2cd85cc25c910fe59289105b583f08fe9e4335ebbb08c1851f47ff2d3482604051610da391906143dd565b60405180910390a25050565b60046020528060005260406000206000915054906101000a900460ff1681565b610dd7612ac4565b806000541015610e1c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e139061447d565b60405180910390fd5b80600080828254610e2d919061449d565b92505081905550600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b8152600401610e919291906144d1565b6020604051808303816000875af1158015610eb0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ed49190614526565b503373ffffffffffffffffffffffffffffffffffffffff167fbfd9c82485e2178fcfd5c40379d8e66fe60acc50aa1ef4c50966431eb1e381cd82604051610f1b9190613fa3565b60405180910390a250565b600a6020528060005260406000206000915054906101000a900460ff1681565b610f4e612ac4565b60005b825181101561112057818181518110610f6d57610f6c614553565b5b602002602001015160076000858481518110610f8c57610f8b614553565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610fdd919061449d565b92505081905550818181518110610ff757610ff6614553565b5b60200260200101516006600085848151811061101657611015614553565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546110679190614582565b9250508190555082818151811061108157611080614553565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff167f0f2440b3ca071b7d18e917a25289e7d7e7de8a491546d45acc2efbec7b3e1ae88383815181106110d3576110d2614553565b5b60200260200101516040516110e89190614612565b60405180910390a261111383828151811061110657611105614553565b5b6020026020010151612b4b565b8080600101915050610f51565b505050565b600860205280600052604060002060009150905080546111449061466f565b80601f01602080910402602001604051908101604052809291908181526020018280546111709061466f565b80156111bd5780601f10611192576101008083540402835291602001916111bd565b820191906000526020600020905b8154815290600101906020018083116111a057829003601f168201915b505050505081565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b8152600401611224939291906146a0565b6020604051808303816000875af1158015611243573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112679190614526565b5080600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546112b79190614582565b925050819055508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f628360405161131b9190613fa3565b60405180910390a361132c82612b4b565b5050565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166113bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113b3906142f1565b60405180910390fd5b60006005836040516113ce9190614268565b9081526020016040518091039020905060008160010154905060008184116113f657836113f8565b815b90508083600101600082825461140e919061449d565b9250508190555060005b836000018054905081101561163957816007600086600001848154811061144257611441614553565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546114b6919061449d565b9250508190555081600660008660000184815481106114d8576114d7614553565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461154c9190614582565b9250508190555083600001818154811061156957611568614553565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f0f2440b3ca071b7d18e917a25289e7d7e7de8a491546d45acc2efbec7b3e1ae883886040516115db9291906146d7565b60405180910390a261162c8460000182815481106115fc576115fb614553565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16612b4b565b8080600101915050611418565b5060018360020160006101000a81548160ff0219169083600281111561166257611661613a12565b5b0217905550846040516116759190614268565b60405180910390207f4a2ced9ada462e244851a86e998eb0b5bf558c2c9c6923b7f970ed2b19b073b060016040516116ad9190614707565b60405180910390a25050505050565b6116c4612cad565b6116cd82612d93565b6116d78282612d9e565b5050565b60006116e5612ebd565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b905090565b6117166134e0565b6040518060800160405280600054815260200160015481526020016002548152602001600354815250905090565b60006002905090565b6005818051602081018201805184825260208301602085012081835280955050505050506000915090508060010154908060020160009054906101000a900460ff16905082565b61179c612ac4565b8173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156117e2573d6000803e3d6000fd5b505050565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611873576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161186a906142f1565b60405180910390fd5b6001600a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b6118d6612ac4565b80600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60005481565b611928612ac4565b6001600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b61198b612ac4565b6119956000612f44565b565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611a23576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a1a906142f1565b60405180910390fd5b8060018190555050565b6000611a3761301b565b905060008160000160089054906101000a900460ff1615905060008260000160009054906101000a900467ffffffffffffffff1690506000808267ffffffffffffffff16148015611a855750825b9050600060018367ffffffffffffffff16148015611aba575060003073ffffffffffffffffffffffffffffffffffffffff163b145b905081158015611ac8575080155b15611aff576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018560000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508315611b4f5760018560000160086101000a81548160ff0219169083151502179055505b611b5833613043565b611b60613057565b6001600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600281905550600160038190555067016345785d8a00006001819055508315611c335760008560000160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d26001604051611c2a919061477b565b60405180910390a15b5050505050565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611cc6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611cbd906142f1565b60405180910390fd5b818160005b8251811015611d7d578160066000858481518110611cec57611ceb614553565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541215611d70576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d6790614808565b60405180910390fd5b8080600101915050611ccb565b5060005b8451811015611f04578360066000878481518110611da257611da1614553565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611df39190614374565b925050819055508360076000878481518110611e1257611e11614553565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611e639190614340565b92505081905550848181518110611e7d57611e7c614553565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff167f5f3d004cf9164b95ed5dbf47d1f04018a4eabcb20b4320fe229ed92236ace6348588604051611ecc9291906146d7565b60405180910390a2611ef7858281518110611eea57611ee9614553565b5b6020026020010151612b4b565b8080600101915050611d81565b50604051806060016040528085815260200184815260200160006002811115611f3057611f2f613a12565b5b815250600586604051611f439190614268565b90815260200160405180910390206000820151816000019080519060200190611f6d929190613508565b506020820151816001015560408201518160020160006101000a81548160ff02191690836002811115611fa357611fa2613a12565b5b02179055509050505050505050565b600080611fbd613061565b90508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505090565b60076020528060005260406000206000915090505481565b60015481565b60066020528060005260406000206000915090505481565b50565b61202b613592565b60405180608001604052808373ffffffffffffffffffffffffffffffffffffffff168152602001600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548152602001600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548152602001600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080546121269061466f565b80601f01602080910402602001604051908101604052809291908181526020018280546121529061466f565b801561219f5780601f106121745761010080835404028352916020019161219f565b820191906000526020600020905b81548152906001019060200180831161218257829003601f168201915b50505050508152509050919050565b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b60025481565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16612279576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612270906142f1565b60405180910390fd5b600060058460405161228b9190614268565b90815260200160405180910390209050600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548213156123b4577fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546040516123509190613fa3565b60405180910390a1600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546000808282546123a89190614340565b92505081905550612404565b7fe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b826040516123e39190613fa3565b60405180910390a1816000808282546123fc9190614340565b925050819055505b60028160020160006101000a81548160ff0219169083600281111561242c5761242b613a12565b5b021790555081600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546124809190614374565b9250508190555061249083612b4b565b8273ffffffffffffffffffffffffffffffffffffffff167f403feb2cd85cc25c910fe59289105b583f08fe9e4335ebbb08c1851f47ff2d3483866040516124d89291906146d7565b60405180910390a2836040516124ee9190614268565b60405180910390207f4a2ced9ada462e244851a86e998eb0b5bf558c2c9c6923b7f970ed2b19b073b060026040516125269190614707565b60405180910390a250505050565b61253c612ac4565b8060038190555050565b61254e612ac4565b80600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461259d9190614374565b925050819055505050565b60035481565b6125b6612ac4565b60005b8251811015612651578281815181106125d5576125d4614553565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff167f0f2440b3ca071b7d18e917a25289e7d7e7de8a491546d45acc2efbec7b3e1ae883838151811061262757612626614553565b5b602002602001015160405161263c9190614612565b60405180910390a280806001019150506125b9565b505050565b806040516126649190614268565b60405180910390203373ffffffffffffffffffffffffffffffffffffffff167faec1d412a3c1e4a13fc2a2e19ac38a5af192a9cf17b074fca8146a2d0655e0c360405160405180910390a350565b6126ba612ac4565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361272c5760006040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016127239190614090565b60405180910390fd5b61273581612f44565b50565b6000811161277b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161277290614874565b60405180910390fd5b6000600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054136127fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016127f4906148e0565b60405180910390fd5b600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481111561287f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161287690614972565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1663893d20e86040518163ffffffff1660e01b8152600401602060405180830381865afa1580156128e1573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061290591906149a7565b73ffffffffffffffffffffffffffffffffffffffff161461295b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161295290614a46565b60405180910390fd5b80600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546129aa9190614374565b92505081905550600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b8152600401612a0e9291906144d1565b6020604051808303816000875af1158015612a2d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a519190614526565b50612a5b82612b4b565b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb83604051612ab89190613fa3565b60405180910390a35050565b612acc613089565b73ffffffffffffffffffffffffffffffffffffffff16612aea611fb2565b73ffffffffffffffffffffffffffffffffffffffff1614612b4957612b0d613089565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401612b409190614090565b60405180910390fd5b565b600154600254612b5b9190614a66565b600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205412612c27576040518060400160405280600981526020017f7a6b41756374696f6e0000000000000000000000000000000000000000000000815250600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209081612c219190614c4a565b50612caa565b6040518060400160405280600381526020017f4e53430000000000000000000000000000000000000000000000000000000000815250600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209081612ca89190614c4a565b505b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161480612d5a57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16612d41613091565b73ffffffffffffffffffffffffffffffffffffffff1614155b15612d91576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b612d9b612ac4565b50565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612e0657506040513d601f19601f82011682018060405250810190612e039190614d48565b60015b612e4757816040517f4c9c8ce3000000000000000000000000000000000000000000000000000000008152600401612e3e9190614090565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b8114612eae57806040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600401612ea59190613f09565b60405180910390fd5b612eb883836130e8565b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614612f42576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6000612f4e613061565b905060008160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050828260000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3505050565b60007ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b61304b61315b565b6130548161319b565b50565b61305f61315b565b565b60007f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300905090565b600033905090565b60006130bf7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b613221565b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6130f18261322b565b8173ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a260008151111561314e5761314882826132f8565b50613157565b61315661337c565b5b5050565b6131636133b9565b613199576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6131a361315b565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036132155760006040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161320c9190614090565b60405180910390fd5b61321e81612f44565b50565b6000819050919050565b60008173ffffffffffffffffffffffffffffffffffffffff163b0361328757806040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815260040161327e9190614090565b60405180910390fd5b806132b47f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b613221565b60000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60606000808473ffffffffffffffffffffffffffffffffffffffff16846040516133229190614dbc565b600060405180830381855af49150503d806000811461335d576040519150601f19603f3d011682016040523d82523d6000602084013e613362565b606091505b50915091506133728583836133d9565b9250505092915050565b60003411156133b7576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60006133c361301b565b60000160089054906101000a900460ff16905090565b6060826133ee576133e982613468565b613460565b60008251148015613416575060008473ffffffffffffffffffffffffffffffffffffffff163b145b1561345857836040517f9996b31500000000000000000000000000000000000000000000000000000000815260040161344f9190614090565b60405180910390fd5b819050613461565b5b9392505050565b60008151111561347b5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60405180606001604052806060815260200160008152602001600060028111156134da576134d9613a12565b5b81525090565b6040518060800160405280600081526020016000815260200160008152602001600081525090565b828054828255906000526020600020908101928215613581579160200282015b828111156135805782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190613528565b5b50905061358e91906135d0565b5090565b6040518060800160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001606081525090565b5b808211156135e95760008160009055506001016135d1565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6136548261360b565b810181811067ffffffffffffffff821117156136735761367261361c565b5b80604052505050565b60006136866135ed565b9050613692828261364b565b919050565b600067ffffffffffffffff8211156136b2576136b161361c565b5b6136bb8261360b565b9050602081019050919050565b82818337600083830152505050565b60006136ea6136e584613697565b61367c565b90508281526020810184848401111561370657613705613606565b5b6137118482856136c8565b509392505050565b600082601f83011261372e5761372d613601565b5b813561373e8482602086016136d7565b91505092915050565b60006020828403121561375d5761375c6135f7565b5b600082013567ffffffffffffffff81111561377b5761377a6135fc565b5b61378784828501613719565b91505092915050565b60008115159050919050565b6137a581613790565b82525050565b60006020820190506137c0600083018461379c565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006137f1826137c6565b9050919050565b613801816137e6565b811461380c57600080fd5b50565b60008135905061381e816137f8565b92915050565b60006020828403121561383a576138396135f7565b5b60006138488482850161380f565b91505092915050565b600067ffffffffffffffff82111561386c5761386b61361c565b5b6138758261360b565b9050602081019050919050565b600061389561389084613851565b61367c565b9050828152602081018484840111156138b1576138b0613606565b5b6138bc8482856136c8565b509392505050565b600082601f8301126138d9576138d8613601565b5b81356138e9848260208601613882565b91505092915050565b600060208284031215613908576139076135f7565b5b600082013567ffffffffffffffff811115613926576139256135fc565b5b613932848285016138c4565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b613970816137e6565b82525050565b60006139828383613967565b60208301905092915050565b6000602082019050919050565b60006139a68261393b565b6139b08185613946565b93506139bb83613957565b8060005b838110156139ec5781516139d38882613976565b97506139de8361398e565b9250506001810190506139bf565b5085935050505092915050565b6000819050919050565b613a0c816139f9565b82525050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60038110613a5257613a51613a12565b5b50565b6000819050613a6382613a41565b919050565b6000613a7382613a55565b9050919050565b613a8381613a68565b82525050565b60006060830160008301518482036000860152613aa6828261399b565b9150506020830151613abb6020860182613a03565b506040830151613ace6040860182613a7a565b508091505092915050565b60006020820190508181036000830152613af38184613a89565b905092915050565b613b04816139f9565b8114613b0f57600080fd5b50565b600081359050613b2181613afb565b92915050565b60008060408385031215613b3e57613b3d6135f7565b5b6000613b4c8582860161380f565b9250506020613b5d85828601613b12565b9150509250929050565b600060208284031215613b7d57613b7c6135f7565b5b6000613b8b84828501613b12565b91505092915050565b600067ffffffffffffffff821115613baf57613bae61361c565b5b602082029050602081019050919050565b600080fd5b6000613bd8613bd384613b94565b61367c565b90508083825260208201905060208402830185811115613bfb57613bfa613bc0565b5b835b81811015613c245780613c10888261380f565b845260208401935050602081019050613bfd565b5050509392505050565b600082601f830112613c4357613c42613601565b5b8135613c53848260208601613bc5565b91505092915050565b600067ffffffffffffffff821115613c7757613c7661361c565b5b602082029050602081019050919050565b6000613c9b613c9684613c5c565b61367c565b90508083825260208201905060208402830185811115613cbe57613cbd613bc0565b5b835b81811015613ce75780613cd38882613b12565b845260208401935050602081019050613cc0565b5050509392505050565b600082601f830112613d0657613d05613601565b5b8135613d16848260208601613c88565b91505092915050565b60008060408385031215613d3657613d356135f7565b5b600083013567ffffffffffffffff811115613d5457613d536135fc565b5b613d6085828601613c2e565b925050602083013567ffffffffffffffff811115613d8157613d806135fc565b5b613d8d85828601613cf1565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b60005b83811015613dd1578082015181840152602081019050613db6565b60008484015250505050565b6000613de882613d97565b613df28185613da2565b9350613e02818560208601613db3565b613e0b8161360b565b840191505092915050565b60006020820190508181036000830152613e308184613ddd565b905092915050565b60008060408385031215613e4f57613e4e6135f7565b5b600083013567ffffffffffffffff811115613e6d57613e6c6135fc565b5b613e79858286016138c4565b9250506020613e8a85828601613b12565b9150509250929050565b60008060408385031215613eab57613eaa6135f7565b5b6000613eb98582860161380f565b925050602083013567ffffffffffffffff811115613eda57613ed96135fc565b5b613ee685828601613719565b9150509250929050565b6000819050919050565b613f0381613ef0565b82525050565b6000602082019050613f1e6000830184613efa565b92915050565b608082016000820151613f3a6000850182613a03565b506020820151613f4d6020850182613a03565b506040820151613f606040850182613a03565b506060820151613f736060850182613a03565b50505050565b6000608082019050613f8e6000830184613f24565b92915050565b613f9d816139f9565b82525050565b6000602082019050613fb86000830184613f94565b92915050565b613fc781613a68565b82525050565b6000604082019050613fe26000830185613f94565b613fef6020830184613fbe565b9392505050565b60008060006060848603121561400f5761400e6135f7565b5b600084013567ffffffffffffffff81111561402d5761402c6135fc565b5b614039868287016138c4565b935050602084013567ffffffffffffffff81111561405a576140596135fc565b5b61406686828701613c2e565b925050604061407786828701613b12565b9150509250925092565b61408a816137e6565b82525050565b60006020820190506140a56000830184614081565b92915050565b6000819050919050565b6140be816140ab565b82525050565b60006020820190506140d960008301846140b5565b92915050565b6140e8816140ab565b82525050565b600082825260208201905092915050565b600061410a82613d97565b61411481856140ee565b9350614124818560208601613db3565b61412d8161360b565b840191505092915050565b60006080830160008301516141506000860182613967565b50602083015161416360208601826140df565b5060408301516141766040860182613a03565b506060830151848203606086015261418e82826140ff565b9150508091505092915050565b600060208201905081810360008301526141b58184614138565b905092915050565b6000806000606084860312156141d6576141d56135f7565b5b600084013567ffffffffffffffff8111156141f4576141f36135fc565b5b614200868287016138c4565b93505060206142118682870161380f565b925050604061422286828701613b12565b9150509250925092565b600081905092915050565b600061424282613d97565b61424c818561422c565b935061425c818560208601613db3565b80840191505092915050565b60006142748284614237565b915081905092915050565b7f4f6e6c79207468652061646d696e2063616e2063616c6c20746869732066756e60008201527f6374696f6e2e0000000000000000000000000000000000000000000000000000602082015250565b60006142db602683613da2565b91506142e68261427f565b604082019050919050565b6000602082019050818103600083015261430a816142ce565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061434b826139f9565b9150614356836139f9565b925082820190508082111561436e5761436d614311565b5b92915050565b600061437f826140ab565b915061438a836140ab565b92508282039050818112600084121682821360008512151617156143b1576143b0614311565b5b92915050565b50565b60006143c7600083613da2565b91506143d2826143b7565b600082019050919050565b60006040820190506143f26000830184613f94565b8181036020830152614403816143ba565b905092915050565b7f576974686472617720616d6f756e7420616d6f756e742065786365656473207360008201527f6c617368656446756e6473000000000000000000000000000000000000000000602082015250565b6000614467602b83613da2565b91506144728261440b565b604082019050919050565b600060208201905081810360008301526144968161445a565b9050919050565b60006144a8826139f9565b91506144b3836139f9565b92508282039050818111156144cb576144ca614311565b5b92915050565b60006040820190506144e66000830185614081565b6144f36020830184613f94565b9392505050565b61450381613790565b811461450e57600080fd5b50565b600081519050614520816144fa565b92915050565b60006020828403121561453c5761453b6135f7565b5b600061454a84828501614511565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600061458d826140ab565b9150614598836140ab565b9250828201905082811215600083121683821260008412151617156145c0576145bf614311565b5b92915050565b7f6261746368656400000000000000000000000000000000000000000000000000600082015250565b60006145fc600783613da2565b9150614607826145c6565b602082019050919050565b60006040820190506146276000830184613f94565b8181036020830152614638816145ef565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061468757607f821691505b60208210810361469a57614699614640565b5b50919050565b60006060820190506146b56000830186614081565b6146c26020830185614081565b6146cf6040830184613f94565b949350505050565b60006040820190506146ec6000830185613f94565b81810360208301526146fe8184613ddd565b90509392505050565b600060208201905061471c6000830184613fbe565b92915050565b6000819050919050565b600067ffffffffffffffff82169050919050565b6000819050919050565b600061476561476061475b84614722565b614740565b61472c565b9050919050565b6147758161474a565b82525050565b6000602082019050614790600083018461476c565b92915050565b7f4e6f7420616c6c20435073206861766520656e6f75676820617661696c61626c60008201527f652062616c616e63650000000000000000000000000000000000000000000000602082015250565b60006147f2602983613da2565b91506147fd82614796565b604082019050919050565b60006020820190508181036000830152614821816147e5565b9050919050565b7f776974686472617720616d6f756e74206d75737420626520706f736974697665600082015250565b600061485e602083613da2565b915061486982614828565b602082019050919050565b6000602082019050818103600083015261488d81614851565b9050919050565b7f6e6f7468696e6720746f20776974686472617700000000000000000000000000600082015250565b60006148ca601383613da2565b91506148d582614894565b602082019050919050565b600060208201905081810360008301526148f9816148bd565b9050919050565b7f776974686472617720616d6f756e742067726561746572207468616e2061766160008201527f696c61626c652062616c616e6365000000000000000000000000000000000000602082015250565b600061495c602e83613da2565b915061496782614900565b604082019050919050565b6000602082019050818103600083015261498b8161494f565b9050919050565b6000815190506149a1816137f8565b92915050565b6000602082840312156149bd576149bc6135f7565b5b60006149cb84828501614992565b91505092915050565b7f4f6e6c792043504163636f756e74206f776e65722063616e207769746864726160008201527f772074686520636f6c6c61746572616c2066756e647300000000000000000000602082015250565b6000614a30603683613da2565b9150614a3b826149d4565b604082019050919050565b60006020820190508181036000830152614a5f81614a23565b9050919050565b6000614a71826139f9565b9150614a7c836139f9565b9250828202614a8a816139f9565b91508282048414831517614aa157614aa0614311565b5b5092915050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302614b0a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82614acd565b614b148683614acd565b95508019841693508086168417925050509392505050565b6000614b47614b42614b3d846139f9565b614740565b6139f9565b9050919050565b6000819050919050565b614b6183614b2c565b614b75614b6d82614b4e565b848454614ada565b825550505050565b600090565b614b8a614b7d565b614b95818484614b58565b505050565b5b81811015614bb957614bae600082614b82565b600181019050614b9b565b5050565b601f821115614bfe57614bcf81614aa8565b614bd884614abd565b81016020851015614be7578190505b614bfb614bf385614abd565b830182614b9a565b50505b505050565b600082821c905092915050565b6000614c2160001984600802614c03565b1980831691505092915050565b6000614c3a8383614c10565b9150826002028217905092915050565b614c5382613d97565b67ffffffffffffffff811115614c6c57614c6b61361c565b5b614c76825461466f565b614c81828285614bbd565b600060209050601f831160018114614cb45760008415614ca2578287015190505b614cac8582614c2e565b865550614d14565b601f198416614cc286614aa8565b60005b82811015614cea57848901518255600182019150602085019450602081019050614cc5565b86831015614d075784890151614d03601f891682614c10565b8355505b6001600288020188555050505b505050505050565b614d2581613ef0565b8114614d3057600080fd5b50565b600081519050614d4281614d1c565b92915050565b600060208284031215614d5e57614d5d6135f7565b5b6000614d6c84828501614d33565b91505092915050565b600081519050919050565b600081905092915050565b6000614d9682614d75565b614da08185614d80565b9350614db0818560208601613db3565b80840191505092915050565b6000614dc88284614d8b565b91508190509291505056fea26469706673582212207a7554bd17d652a834b6ba04313c0c517d8bb75107dfd0f0ecbf0e6998c39dab64736f6c63430008190033",
}

// FcpCollateralABI is the input ABI used to generate the binding from.
// Deprecated: Use FcpCollateralMetaData.ABI instead.
var FcpCollateralABI = FcpCollateralMetaData.ABI

// FcpCollateralBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FcpCollateralMetaData.Bin instead.
var FcpCollateralBin = FcpCollateralMetaData.Bin

// DeployFcpCollateral deploys a new Ethereum contract, binding an instance of FcpCollateral to it.
func DeployFcpCollateral(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FcpCollateral, error) {
	parsed, err := FcpCollateralMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FcpCollateralBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FcpCollateral{FcpCollateralCaller: FcpCollateralCaller{contract: contract}, FcpCollateralTransactor: FcpCollateralTransactor{contract: contract}, FcpCollateralFilterer: FcpCollateralFilterer{contract: contract}}, nil
}

// FcpCollateral is an auto generated Go binding around an Ethereum contract.
type FcpCollateral struct {
	FcpCollateralCaller     // Read-only binding to the contract
	FcpCollateralTransactor // Write-only binding to the contract
	FcpCollateralFilterer   // Log filterer for contract events
}

// FcpCollateralCaller is an auto generated read-only Go binding around an Ethereum contract.
type FcpCollateralCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FcpCollateralTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FcpCollateralTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FcpCollateralFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FcpCollateralFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FcpCollateralSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FcpCollateralSession struct {
	Contract     *FcpCollateral    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FcpCollateralCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FcpCollateralCallerSession struct {
	Contract *FcpCollateralCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// FcpCollateralTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FcpCollateralTransactorSession struct {
	Contract     *FcpCollateralTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// FcpCollateralRaw is an auto generated low-level Go binding around an Ethereum contract.
type FcpCollateralRaw struct {
	Contract *FcpCollateral // Generic contract binding to access the raw methods on
}

// FcpCollateralCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FcpCollateralCallerRaw struct {
	Contract *FcpCollateralCaller // Generic read-only contract binding to access the raw methods on
}

// FcpCollateralTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FcpCollateralTransactorRaw struct {
	Contract *FcpCollateralTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFcpCollateral creates a new instance of FcpCollateral, bound to a specific deployed contract.
func NewFcpCollateral(address common.Address, backend bind.ContractBackend) (*FcpCollateral, error) {
	contract, err := bindFcpCollateral(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FcpCollateral{FcpCollateralCaller: FcpCollateralCaller{contract: contract}, FcpCollateralTransactor: FcpCollateralTransactor{contract: contract}, FcpCollateralFilterer: FcpCollateralFilterer{contract: contract}}, nil
}

// NewFcpCollateralCaller creates a new read-only instance of FcpCollateral, bound to a specific deployed contract.
func NewFcpCollateralCaller(address common.Address, caller bind.ContractCaller) (*FcpCollateralCaller, error) {
	contract, err := bindFcpCollateral(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FcpCollateralCaller{contract: contract}, nil
}

// NewFcpCollateralTransactor creates a new write-only instance of FcpCollateral, bound to a specific deployed contract.
func NewFcpCollateralTransactor(address common.Address, transactor bind.ContractTransactor) (*FcpCollateralTransactor, error) {
	contract, err := bindFcpCollateral(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FcpCollateralTransactor{contract: contract}, nil
}

// NewFcpCollateralFilterer creates a new log filterer instance of FcpCollateral, bound to a specific deployed contract.
func NewFcpCollateralFilterer(address common.Address, filterer bind.ContractFilterer) (*FcpCollateralFilterer, error) {
	contract, err := bindFcpCollateral(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FcpCollateralFilterer{contract: contract}, nil
}

// bindFcpCollateral binds a generic wrapper to an already deployed contract.
func bindFcpCollateral(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FcpCollateralMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FcpCollateral *FcpCollateralRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FcpCollateral.Contract.FcpCollateralCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FcpCollateral *FcpCollateralRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FcpCollateral.Contract.FcpCollateralTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FcpCollateral *FcpCollateralRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FcpCollateral.Contract.FcpCollateralTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FcpCollateral *FcpCollateralCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FcpCollateral.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FcpCollateral *FcpCollateralTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FcpCollateral.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FcpCollateral *FcpCollateralTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FcpCollateral.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_FcpCollateral *FcpCollateralCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FcpCollateral.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_FcpCollateral *FcpCollateralSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _FcpCollateral.Contract.UPGRADEINTERFACEVERSION(&_FcpCollateral.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_FcpCollateral *FcpCollateralCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _FcpCollateral.Contract.UPGRADEINTERFACEVERSION(&_FcpCollateral.CallOpts)
}

// AllowedToWithdraw is a free data retrieval call binding the contract method 0x2d291cad.
//
// Solidity: function allowedToWithdraw(address ) view returns(bool)
func (_FcpCollateral *FcpCollateralCaller) AllowedToWithdraw(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _FcpCollateral.contract.Call(opts, &out, "allowedToWithdraw", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedToWithdraw is a free data retrieval call binding the contract method 0x2d291cad.
//
// Solidity: function allowedToWithdraw(address ) view returns(bool)
func (_FcpCollateral *FcpCollateralSession) AllowedToWithdraw(arg0 common.Address) (bool, error) {
	return _FcpCollateral.Contract.AllowedToWithdraw(&_FcpCollateral.CallOpts, arg0)
}

// AllowedToWithdraw is a free data retrieval call binding the contract method 0x2d291cad.
//
// Solidity: function allowedToWithdraw(address ) view returns(bool)
func (_FcpCollateral *FcpCollateralCallerSession) AllowedToWithdraw(arg0 common.Address) (bool, error) {
	return _FcpCollateral.Contract.AllowedToWithdraw(&_FcpCollateral.CallOpts, arg0)
}

// AvailableBalance is a free data retrieval call binding the contract method 0xa0821be3.
//
// Solidity: function availableBalance(address ) view returns(int256)
func (_FcpCollateral *FcpCollateralCaller) AvailableBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FcpCollateral.contract.Call(opts, &out, "availableBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableBalance is a free data retrieval call binding the contract method 0xa0821be3.
//
// Solidity: function availableBalance(address ) view returns(int256)
func (_FcpCollateral *FcpCollateralSession) AvailableBalance(arg0 common.Address) (*big.Int, error) {
	return _FcpCollateral.Contract.AvailableBalance(&_FcpCollateral.CallOpts, arg0)
}

// AvailableBalance is a free data retrieval call binding the contract method 0xa0821be3.
//
// Solidity: function availableBalance(address ) view returns(int256)
func (_FcpCollateral *FcpCollateralCallerSession) AvailableBalance(arg0 common.Address) (*big.Int, error) {
	return _FcpCollateral.Contract.AvailableBalance(&_FcpCollateral.CallOpts, arg0)
}

// BaseCollateral is a free data retrieval call binding the contract method 0x9b5ddf09.
//
// Solidity: function baseCollateral() view returns(uint256)
func (_FcpCollateral *FcpCollateralCaller) BaseCollateral(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FcpCollateral.contract.Call(opts, &out, "baseCollateral")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseCollateral is a free data retrieval call binding the contract method 0x9b5ddf09.
//
// Solidity: function baseCollateral() view returns(uint256)
func (_FcpCollateral *FcpCollateralSession) BaseCollateral() (*big.Int, error) {
	return _FcpCollateral.Contract.BaseCollateral(&_FcpCollateral.CallOpts)
}

// BaseCollateral is a free data retrieval call binding the contract method 0x9b5ddf09.
//
// Solidity: function baseCollateral() view returns(uint256)
func (_FcpCollateral *FcpCollateralCallerSession) BaseCollateral() (*big.Int, error) {
	return _FcpCollateral.Contract.BaseCollateral(&_FcpCollateral.CallOpts)
}

// CollateralInfo is a free data retrieval call binding the contract method 0x53ad8720.
//
// Solidity: function collateralInfo() view returns((uint256,uint256,uint256,uint256))
func (_FcpCollateral *FcpCollateralCaller) CollateralInfo(opts *bind.CallOpts) (SwanCreditCollateralContractInfo, error) {
	var out []interface{}
	err := _FcpCollateral.contract.Call(opts, &out, "collateralInfo")

	if err != nil {
		return *new(SwanCreditCollateralContractInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(SwanCreditCollateralContractInfo)).(*SwanCreditCollateralContractInfo)

	return out0, err

}

// CollateralInfo is a free data retrieval call binding the contract method 0x53ad8720.
//
// Solidity: function collateralInfo() view returns((uint256,uint256,uint256,uint256))
func (_FcpCollateral *FcpCollateralSession) CollateralInfo() (SwanCreditCollateralContractInfo, error) {
	return _FcpCollateral.Contract.CollateralInfo(&_FcpCollateral.CallOpts)
}

// CollateralInfo is a free data retrieval call binding the contract method 0x53ad8720.
//
// Solidity: function collateralInfo() view returns((uint256,uint256,uint256,uint256))
func (_FcpCollateral *FcpCollateralCallerSession) CollateralInfo() (SwanCreditCollateralContractInfo, error) {
	return _FcpCollateral.Contract.CollateralInfo(&_FcpCollateral.CallOpts)
}

// CollateralRatio is a free data retrieval call binding the contract method 0xb4eae1cb.
//
// Solidity: function collateralRatio() view returns(uint256)
func (_FcpCollateral *FcpCollateralCaller) CollateralRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FcpCollateral.contract.Call(opts, &out, "collateralRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollateralRatio is a free data retrieval call binding the contract method 0xb4eae1cb.
//
// Solidity: function collateralRatio() view returns(uint256)
func (_FcpCollateral *FcpCollateralSession) CollateralRatio() (*big.Int, error) {
	return _FcpCollateral.Contract.CollateralRatio(&_FcpCollateral.CallOpts)
}

// CollateralRatio is a free data retrieval call binding the contract method 0xb4eae1cb.
//
// Solidity: function collateralRatio() view returns(uint256)
func (_FcpCollateral *FcpCollateralCallerSession) CollateralRatio() (*big.Int, error) {
	return _FcpCollateral.Contract.CollateralRatio(&_FcpCollateral.CallOpts)
}

// CpInfo is a free data retrieval call binding the contract method 0xa664c216.
//
// Solidity: function cpInfo(address cpAddress) view returns((address,int256,uint256,string))
func (_FcpCollateral *FcpCollateralCaller) CpInfo(opts *bind.CallOpts, cpAddress common.Address) (SwanCreditCollateralCPInfo, error) {
	var out []interface{}
	err := _FcpCollateral.contract.Call(opts, &out, "cpInfo", cpAddress)

	if err != nil {
		return *new(SwanCreditCollateralCPInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(SwanCreditCollateralCPInfo)).(*SwanCreditCollateralCPInfo)

	return out0, err

}

// CpInfo is a free data retrieval call binding the contract method 0xa664c216.
//
// Solidity: function cpInfo(address cpAddress) view returns((address,int256,uint256,string))
func (_FcpCollateral *FcpCollateralSession) CpInfo(cpAddress common.Address) (SwanCreditCollateralCPInfo, error) {
	return _FcpCollateral.Contract.CpInfo(&_FcpCollateral.CallOpts, cpAddress)
}

// CpInfo is a free data retrieval call binding the contract method 0xa664c216.
//
// Solidity: function cpInfo(address cpAddress) view returns((address,int256,uint256,string))
func (_FcpCollateral *FcpCollateralCallerSession) CpInfo(cpAddress common.Address) (SwanCreditCollateralCPInfo, error) {
	return _FcpCollateral.Contract.CpInfo(&_FcpCollateral.CallOpts, cpAddress)
}

// CpStatus is a free data retrieval call binding the contract method 0x3fe65177.
//
// Solidity: function cpStatus(address ) view returns(string)
func (_FcpCollateral *FcpCollateralCaller) CpStatus(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _FcpCollateral.contract.Call(opts, &out, "cpStatus", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CpStatus is a free data retrieval call binding the contract method 0x3fe65177.
//
// Solidity: function cpStatus(address ) view returns(string)
func (_FcpCollateral *FcpCollateralSession) CpStatus(arg0 common.Address) (string, error) {
	return _FcpCollateral.Contract.CpStatus(&_FcpCollateral.CallOpts, arg0)
}

// CpStatus is a free data retrieval call binding the contract method 0x3fe65177.
//
// Solidity: function cpStatus(address ) view returns(string)
func (_FcpCollateral *FcpCollateralCallerSession) CpStatus(arg0 common.Address) (string, error) {
	return _FcpCollateral.Contract.CpStatus(&_FcpCollateral.CallOpts, arg0)
}

// GetTaskInfo is a free data retrieval call binding the contract method 0x1b209463.
//
// Solidity: function getTaskInfo(string taskUid) view returns((address[],uint256,uint8))
func (_FcpCollateral *FcpCollateralCaller) GetTaskInfo(opts *bind.CallOpts, taskUid string) (SwanCreditCollateralTask, error) {
	var out []interface{}
	err := _FcpCollateral.contract.Call(opts, &out, "getTaskInfo", taskUid)

	if err != nil {
		return *new(SwanCreditCollateralTask), err
	}

	out0 := *abi.ConvertType(out[0], new(SwanCreditCollateralTask)).(*SwanCreditCollateralTask)

	return out0, err

}

// GetTaskInfo is a free data retrieval call binding the contract method 0x1b209463.
//
// Solidity: function getTaskInfo(string taskUid) view returns((address[],uint256,uint8))
func (_FcpCollateral *FcpCollateralSession) GetTaskInfo(taskUid string) (SwanCreditCollateralTask, error) {
	return _FcpCollateral.Contract.GetTaskInfo(&_FcpCollateral.CallOpts, taskUid)
}

// GetTaskInfo is a free data retrieval call binding the contract method 0x1b209463.
//
// Solidity: function getTaskInfo(string taskUid) view returns((address[],uint256,uint8))
func (_FcpCollateral *FcpCollateralCallerSession) GetTaskInfo(taskUid string) (SwanCreditCollateralTask, error) {
	return _FcpCollateral.Contract.GetTaskInfo(&_FcpCollateral.CallOpts, taskUid)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_FcpCollateral *FcpCollateralCaller) IsAdmin(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _FcpCollateral.contract.Call(opts, &out, "isAdmin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_FcpCollateral *FcpCollateralSession) IsAdmin(arg0 common.Address) (bool, error) {
	return _FcpCollateral.Contract.IsAdmin(&_FcpCollateral.CallOpts, arg0)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_FcpCollateral *FcpCollateralCallerSession) IsAdmin(arg0 common.Address) (bool, error) {
	return _FcpCollateral.Contract.IsAdmin(&_FcpCollateral.CallOpts, arg0)
}

// IsSignatureUsed is a free data retrieval call binding the contract method 0x1150f0f3.
//
// Solidity: function isSignatureUsed(bytes ) view returns(bool)
func (_FcpCollateral *FcpCollateralCaller) IsSignatureUsed(opts *bind.CallOpts, arg0 []byte) (bool, error) {
	var out []interface{}
	err := _FcpCollateral.contract.Call(opts, &out, "isSignatureUsed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSignatureUsed is a free data retrieval call binding the contract method 0x1150f0f3.
//
// Solidity: function isSignatureUsed(bytes ) view returns(bool)
func (_FcpCollateral *FcpCollateralSession) IsSignatureUsed(arg0 []byte) (bool, error) {
	return _FcpCollateral.Contract.IsSignatureUsed(&_FcpCollateral.CallOpts, arg0)
}

// IsSignatureUsed is a free data retrieval call binding the contract method 0x1150f0f3.
//
// Solidity: function isSignatureUsed(bytes ) view returns(bool)
func (_FcpCollateral *FcpCollateralCallerSession) IsSignatureUsed(arg0 []byte) (bool, error) {
	return _FcpCollateral.Contract.IsSignatureUsed(&_FcpCollateral.CallOpts, arg0)
}

// LockedCollateral is a free data retrieval call binding the contract method 0x92bdf9ba.
//
// Solidity: function lockedCollateral(address ) view returns(uint256)
func (_FcpCollateral *FcpCollateralCaller) LockedCollateral(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FcpCollateral.contract.Call(opts, &out, "lockedCollateral", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedCollateral is a free data retrieval call binding the contract method 0x92bdf9ba.
//
// Solidity: function lockedCollateral(address ) view returns(uint256)
func (_FcpCollateral *FcpCollateralSession) LockedCollateral(arg0 common.Address) (*big.Int, error) {
	return _FcpCollateral.Contract.LockedCollateral(&_FcpCollateral.CallOpts, arg0)
}

// LockedCollateral is a free data retrieval call binding the contract method 0x92bdf9ba.
//
// Solidity: function lockedCollateral(address ) view returns(uint256)
func (_FcpCollateral *FcpCollateralCallerSession) LockedCollateral(arg0 common.Address) (*big.Int, error) {
	return _FcpCollateral.Contract.LockedCollateral(&_FcpCollateral.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FcpCollateral *FcpCollateralCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FcpCollateral.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FcpCollateral *FcpCollateralSession) Owner() (common.Address, error) {
	return _FcpCollateral.Contract.Owner(&_FcpCollateral.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FcpCollateral *FcpCollateralCallerSession) Owner() (common.Address, error) {
	return _FcpCollateral.Contract.Owner(&_FcpCollateral.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FcpCollateral *FcpCollateralCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FcpCollateral.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FcpCollateral *FcpCollateralSession) ProxiableUUID() ([32]byte, error) {
	return _FcpCollateral.Contract.ProxiableUUID(&_FcpCollateral.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FcpCollateral *FcpCollateralCallerSession) ProxiableUUID() ([32]byte, error) {
	return _FcpCollateral.Contract.ProxiableUUID(&_FcpCollateral.CallOpts)
}

// SlashRatio is a free data retrieval call binding the contract method 0xd27ca89b.
//
// Solidity: function slashRatio() view returns(uint256)
func (_FcpCollateral *FcpCollateralCaller) SlashRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FcpCollateral.contract.Call(opts, &out, "slashRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashRatio is a free data retrieval call binding the contract method 0xd27ca89b.
//
// Solidity: function slashRatio() view returns(uint256)
func (_FcpCollateral *FcpCollateralSession) SlashRatio() (*big.Int, error) {
	return _FcpCollateral.Contract.SlashRatio(&_FcpCollateral.CallOpts)
}

// SlashRatio is a free data retrieval call binding the contract method 0xd27ca89b.
//
// Solidity: function slashRatio() view returns(uint256)
func (_FcpCollateral *FcpCollateralCallerSession) SlashRatio() (*big.Int, error) {
	return _FcpCollateral.Contract.SlashRatio(&_FcpCollateral.CallOpts)
}

// SlashedFunds is a free data retrieval call binding the contract method 0x6f99f15c.
//
// Solidity: function slashedFunds() view returns(uint256)
func (_FcpCollateral *FcpCollateralCaller) SlashedFunds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FcpCollateral.contract.Call(opts, &out, "slashedFunds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashedFunds is a free data retrieval call binding the contract method 0x6f99f15c.
//
// Solidity: function slashedFunds() view returns(uint256)
func (_FcpCollateral *FcpCollateralSession) SlashedFunds() (*big.Int, error) {
	return _FcpCollateral.Contract.SlashedFunds(&_FcpCollateral.CallOpts)
}

// SlashedFunds is a free data retrieval call binding the contract method 0x6f99f15c.
//
// Solidity: function slashedFunds() view returns(uint256)
func (_FcpCollateral *FcpCollateralCallerSession) SlashedFunds() (*big.Int, error) {
	return _FcpCollateral.Contract.SlashedFunds(&_FcpCollateral.CallOpts)
}

// Tasks is a free data retrieval call binding the contract method 0x58709cf2.
//
// Solidity: function tasks(string ) view returns(uint256 collateral, uint8 collateralStatus)
func (_FcpCollateral *FcpCollateralCaller) Tasks(opts *bind.CallOpts, arg0 string) (struct {
	Collateral       *big.Int
	CollateralStatus uint8
}, error) {
	var out []interface{}
	err := _FcpCollateral.contract.Call(opts, &out, "tasks", arg0)

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
func (_FcpCollateral *FcpCollateralSession) Tasks(arg0 string) (struct {
	Collateral       *big.Int
	CollateralStatus uint8
}, error) {
	return _FcpCollateral.Contract.Tasks(&_FcpCollateral.CallOpts, arg0)
}

// Tasks is a free data retrieval call binding the contract method 0x58709cf2.
//
// Solidity: function tasks(string ) view returns(uint256 collateral, uint8 collateralStatus)
func (_FcpCollateral *FcpCollateralCallerSession) Tasks(arg0 string) (struct {
	Collateral       *big.Int
	CollateralStatus uint8
}, error) {
	return _FcpCollateral.Contract.Tasks(&_FcpCollateral.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_FcpCollateral *FcpCollateralCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FcpCollateral.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_FcpCollateral *FcpCollateralSession) Version() (*big.Int, error) {
	return _FcpCollateral.Contract.Version(&_FcpCollateral.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_FcpCollateral *FcpCollateralCallerSession) Version() (*big.Int, error) {
	return _FcpCollateral.Contract.Version(&_FcpCollateral.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address newAdmin) returns()
func (_FcpCollateral *FcpCollateralTransactor) AddAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _FcpCollateral.contract.Transact(opts, "addAdmin", newAdmin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address newAdmin) returns()
func (_FcpCollateral *FcpCollateralSession) AddAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _FcpCollateral.Contract.AddAdmin(&_FcpCollateral.TransactOpts, newAdmin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address newAdmin) returns()
func (_FcpCollateral *FcpCollateralTransactorSession) AddAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _FcpCollateral.Contract.AddAdmin(&_FcpCollateral.TransactOpts, newAdmin)
}

// ApproveWithdraw is a paid mutator transaction binding the contract method 0x6579e35c.
//
// Solidity: function approveWithdraw(address cpAccount) returns()
func (_FcpCollateral *FcpCollateralTransactor) ApproveWithdraw(opts *bind.TransactOpts, cpAccount common.Address) (*types.Transaction, error) {
	return _FcpCollateral.contract.Transact(opts, "approveWithdraw", cpAccount)
}

// ApproveWithdraw is a paid mutator transaction binding the contract method 0x6579e35c.
//
// Solidity: function approveWithdraw(address cpAccount) returns()
func (_FcpCollateral *FcpCollateralSession) ApproveWithdraw(cpAccount common.Address) (*types.Transaction, error) {
	return _FcpCollateral.Contract.ApproveWithdraw(&_FcpCollateral.TransactOpts, cpAccount)
}

// ApproveWithdraw is a paid mutator transaction binding the contract method 0x6579e35c.
//
// Solidity: function approveWithdraw(address cpAccount) returns()
func (_FcpCollateral *FcpCollateralTransactorSession) ApproveWithdraw(cpAccount common.Address) (*types.Transaction, error) {
	return _FcpCollateral.Contract.ApproveWithdraw(&_FcpCollateral.TransactOpts, cpAccount)
}

// BatchUnlockCollateral is a paid mutator transaction binding the contract method 0x2d45ac91.
//
// Solidity: function batchUnlockCollateral(address[] cpList, uint256[] amounts) returns()
func (_FcpCollateral *FcpCollateralTransactor) BatchUnlockCollateral(opts *bind.TransactOpts, cpList []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _FcpCollateral.contract.Transact(opts, "batchUnlockCollateral", cpList, amounts)
}

// BatchUnlockCollateral is a paid mutator transaction binding the contract method 0x2d45ac91.
//
// Solidity: function batchUnlockCollateral(address[] cpList, uint256[] amounts) returns()
func (_FcpCollateral *FcpCollateralSession) BatchUnlockCollateral(cpList []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _FcpCollateral.Contract.BatchUnlockCollateral(&_FcpCollateral.TransactOpts, cpList, amounts)
}

// BatchUnlockCollateral is a paid mutator transaction binding the contract method 0x2d45ac91.
//
// Solidity: function batchUnlockCollateral(address[] cpList, uint256[] amounts) returns()
func (_FcpCollateral *FcpCollateralTransactorSession) BatchUnlockCollateral(cpList []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _FcpCollateral.Contract.BatchUnlockCollateral(&_FcpCollateral.TransactOpts, cpList, amounts)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address cpAccount, uint256 amount) returns()
func (_FcpCollateral *FcpCollateralTransactor) Deposit(opts *bind.TransactOpts, cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.contract.Transact(opts, "deposit", cpAccount, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address cpAccount, uint256 amount) returns()
func (_FcpCollateral *FcpCollateralSession) Deposit(cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.Contract.Deposit(&_FcpCollateral.TransactOpts, cpAccount, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address cpAccount, uint256 amount) returns()
func (_FcpCollateral *FcpCollateralTransactorSession) Deposit(cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.Contract.Deposit(&_FcpCollateral.TransactOpts, cpAccount, amount)
}

// DisputeProof is a paid mutator transaction binding the contract method 0xe80d9dec.
//
// Solidity: function disputeProof(string taskUid) returns()
func (_FcpCollateral *FcpCollateralTransactor) DisputeProof(opts *bind.TransactOpts, taskUid string) (*types.Transaction, error) {
	return _FcpCollateral.contract.Transact(opts, "disputeProof", taskUid)
}

// DisputeProof is a paid mutator transaction binding the contract method 0xe80d9dec.
//
// Solidity: function disputeProof(string taskUid) returns()
func (_FcpCollateral *FcpCollateralSession) DisputeProof(taskUid string) (*types.Transaction, error) {
	return _FcpCollateral.Contract.DisputeProof(&_FcpCollateral.TransactOpts, taskUid)
}

// DisputeProof is a paid mutator transaction binding the contract method 0xe80d9dec.
//
// Solidity: function disputeProof(string taskUid) returns()
func (_FcpCollateral *FcpCollateralTransactorSession) DisputeProof(taskUid string) (*types.Transaction, error) {
	return _FcpCollateral.Contract.DisputeProof(&_FcpCollateral.TransactOpts, taskUid)
}

// EmitCollateralUnlockedEvent is a paid mutator transaction binding the contract method 0xe02a63b5.
//
// Solidity: function emitCollateralUnlockedEvent(address[] cpList, uint256[] amounts) returns()
func (_FcpCollateral *FcpCollateralTransactor) EmitCollateralUnlockedEvent(opts *bind.TransactOpts, cpList []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _FcpCollateral.contract.Transact(opts, "emitCollateralUnlockedEvent", cpList, amounts)
}

// EmitCollateralUnlockedEvent is a paid mutator transaction binding the contract method 0xe02a63b5.
//
// Solidity: function emitCollateralUnlockedEvent(address[] cpList, uint256[] amounts) returns()
func (_FcpCollateral *FcpCollateralSession) EmitCollateralUnlockedEvent(cpList []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _FcpCollateral.Contract.EmitCollateralUnlockedEvent(&_FcpCollateral.TransactOpts, cpList, amounts)
}

// EmitCollateralUnlockedEvent is a paid mutator transaction binding the contract method 0xe02a63b5.
//
// Solidity: function emitCollateralUnlockedEvent(address[] cpList, uint256[] amounts) returns()
func (_FcpCollateral *FcpCollateralTransactorSession) EmitCollateralUnlockedEvent(cpList []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _FcpCollateral.Contract.EmitCollateralUnlockedEvent(&_FcpCollateral.TransactOpts, cpList, amounts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_FcpCollateral *FcpCollateralTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FcpCollateral.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_FcpCollateral *FcpCollateralSession) Initialize() (*types.Transaction, error) {
	return _FcpCollateral.Contract.Initialize(&_FcpCollateral.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_FcpCollateral *FcpCollateralTransactorSession) Initialize() (*types.Transaction, error) {
	return _FcpCollateral.Contract.Initialize(&_FcpCollateral.TransactOpts)
}

// LockCollateral is a paid mutator transaction binding the contract method 0x853911b2.
//
// Solidity: function lockCollateral(string taskUid, address[] cpList, uint256 taskCollateral) returns()
func (_FcpCollateral *FcpCollateralTransactor) LockCollateral(opts *bind.TransactOpts, taskUid string, cpList []common.Address, taskCollateral *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.contract.Transact(opts, "lockCollateral", taskUid, cpList, taskCollateral)
}

// LockCollateral is a paid mutator transaction binding the contract method 0x853911b2.
//
// Solidity: function lockCollateral(string taskUid, address[] cpList, uint256 taskCollateral) returns()
func (_FcpCollateral *FcpCollateralSession) LockCollateral(taskUid string, cpList []common.Address, taskCollateral *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.Contract.LockCollateral(&_FcpCollateral.TransactOpts, taskUid, cpList, taskCollateral)
}

// LockCollateral is a paid mutator transaction binding the contract method 0x853911b2.
//
// Solidity: function lockCollateral(string taskUid, address[] cpList, uint256 taskCollateral) returns()
func (_FcpCollateral *FcpCollateralTransactorSession) LockCollateral(taskUid string, cpList []common.Address, taskCollateral *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.Contract.LockCollateral(&_FcpCollateral.TransactOpts, taskUid, cpList, taskCollateral)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_FcpCollateral *FcpCollateralTransactor) RemoveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _FcpCollateral.contract.Transact(opts, "removeAdmin", admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_FcpCollateral *FcpCollateralSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _FcpCollateral.Contract.RemoveAdmin(&_FcpCollateral.TransactOpts, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_FcpCollateral *FcpCollateralTransactorSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _FcpCollateral.Contract.RemoveAdmin(&_FcpCollateral.TransactOpts, admin)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FcpCollateral *FcpCollateralTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FcpCollateral.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FcpCollateral *FcpCollateralSession) RenounceOwnership() (*types.Transaction, error) {
	return _FcpCollateral.Contract.RenounceOwnership(&_FcpCollateral.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FcpCollateral *FcpCollateralTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FcpCollateral.Contract.RenounceOwnership(&_FcpCollateral.TransactOpts)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0xa35a36e9.
//
// Solidity: function requestWithdraw(address cpAccount) payable returns()
func (_FcpCollateral *FcpCollateralTransactor) RequestWithdraw(opts *bind.TransactOpts, cpAccount common.Address) (*types.Transaction, error) {
	return _FcpCollateral.contract.Transact(opts, "requestWithdraw", cpAccount)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0xa35a36e9.
//
// Solidity: function requestWithdraw(address cpAccount) payable returns()
func (_FcpCollateral *FcpCollateralSession) RequestWithdraw(cpAccount common.Address) (*types.Transaction, error) {
	return _FcpCollateral.Contract.RequestWithdraw(&_FcpCollateral.TransactOpts, cpAccount)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0xa35a36e9.
//
// Solidity: function requestWithdraw(address cpAccount) payable returns()
func (_FcpCollateral *FcpCollateralTransactorSession) RequestWithdraw(cpAccount common.Address) (*types.Transaction, error) {
	return _FcpCollateral.Contract.RequestWithdraw(&_FcpCollateral.TransactOpts, cpAccount)
}

// SendETH is a paid mutator transaction binding the contract method 0x64a197f3.
//
// Solidity: function sendETH(address recipient, uint256 amount) returns()
func (_FcpCollateral *FcpCollateralTransactor) SendETH(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.contract.Transact(opts, "sendETH", recipient, amount)
}

// SendETH is a paid mutator transaction binding the contract method 0x64a197f3.
//
// Solidity: function sendETH(address recipient, uint256 amount) returns()
func (_FcpCollateral *FcpCollateralSession) SendETH(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.Contract.SendETH(&_FcpCollateral.TransactOpts, recipient, amount)
}

// SendETH is a paid mutator transaction binding the contract method 0x64a197f3.
//
// Solidity: function sendETH(address recipient, uint256 amount) returns()
func (_FcpCollateral *FcpCollateralTransactorSession) SendETH(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.Contract.SendETH(&_FcpCollateral.TransactOpts, recipient, amount)
}

// SetBaseCollateral is a paid mutator transaction binding the contract method 0x7f58a7e5.
//
// Solidity: function setBaseCollateral(uint256 _baseCollateral) returns()
func (_FcpCollateral *FcpCollateralTransactor) SetBaseCollateral(opts *bind.TransactOpts, _baseCollateral *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.contract.Transact(opts, "setBaseCollateral", _baseCollateral)
}

// SetBaseCollateral is a paid mutator transaction binding the contract method 0x7f58a7e5.
//
// Solidity: function setBaseCollateral(uint256 _baseCollateral) returns()
func (_FcpCollateral *FcpCollateralSession) SetBaseCollateral(_baseCollateral *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.Contract.SetBaseCollateral(&_FcpCollateral.TransactOpts, _baseCollateral)
}

// SetBaseCollateral is a paid mutator transaction binding the contract method 0x7f58a7e5.
//
// Solidity: function setBaseCollateral(uint256 _baseCollateral) returns()
func (_FcpCollateral *FcpCollateralTransactorSession) SetBaseCollateral(_baseCollateral *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.Contract.SetBaseCollateral(&_FcpCollateral.TransactOpts, _baseCollateral)
}

// SetCollateralToken is a paid mutator transaction binding the contract method 0x666181a9.
//
// Solidity: function setCollateralToken(address tokenAddress) returns()
func (_FcpCollateral *FcpCollateralTransactor) SetCollateralToken(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _FcpCollateral.contract.Transact(opts, "setCollateralToken", tokenAddress)
}

// SetCollateralToken is a paid mutator transaction binding the contract method 0x666181a9.
//
// Solidity: function setCollateralToken(address tokenAddress) returns()
func (_FcpCollateral *FcpCollateralSession) SetCollateralToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _FcpCollateral.Contract.SetCollateralToken(&_FcpCollateral.TransactOpts, tokenAddress)
}

// SetCollateralToken is a paid mutator transaction binding the contract method 0x666181a9.
//
// Solidity: function setCollateralToken(address tokenAddress) returns()
func (_FcpCollateral *FcpCollateralTransactorSession) SetCollateralToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _FcpCollateral.Contract.SetCollateralToken(&_FcpCollateral.TransactOpts, tokenAddress)
}

// SetSlashRatio is a paid mutator transaction binding the contract method 0xce3518aa.
//
// Solidity: function setSlashRatio(uint256 _slashRatio) returns()
func (_FcpCollateral *FcpCollateralTransactor) SetSlashRatio(opts *bind.TransactOpts, _slashRatio *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.contract.Transact(opts, "setSlashRatio", _slashRatio)
}

// SetSlashRatio is a paid mutator transaction binding the contract method 0xce3518aa.
//
// Solidity: function setSlashRatio(uint256 _slashRatio) returns()
func (_FcpCollateral *FcpCollateralSession) SetSlashRatio(_slashRatio *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.Contract.SetSlashRatio(&_FcpCollateral.TransactOpts, _slashRatio)
}

// SetSlashRatio is a paid mutator transaction binding the contract method 0xce3518aa.
//
// Solidity: function setSlashRatio(uint256 _slashRatio) returns()
func (_FcpCollateral *FcpCollateralTransactorSession) SetSlashRatio(_slashRatio *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.Contract.SetSlashRatio(&_FcpCollateral.TransactOpts, _slashRatio)
}

// SlashCollateral is a paid mutator transaction binding the contract method 0x1d47a62d.
//
// Solidity: function slashCollateral(address cpAccount, uint256 slashAmount) returns()
func (_FcpCollateral *FcpCollateralTransactor) SlashCollateral(opts *bind.TransactOpts, cpAccount common.Address, slashAmount *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.contract.Transact(opts, "slashCollateral", cpAccount, slashAmount)
}

// SlashCollateral is a paid mutator transaction binding the contract method 0x1d47a62d.
//
// Solidity: function slashCollateral(address cpAccount, uint256 slashAmount) returns()
func (_FcpCollateral *FcpCollateralSession) SlashCollateral(cpAccount common.Address, slashAmount *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.Contract.SlashCollateral(&_FcpCollateral.TransactOpts, cpAccount, slashAmount)
}

// SlashCollateral is a paid mutator transaction binding the contract method 0x1d47a62d.
//
// Solidity: function slashCollateral(address cpAccount, uint256 slashAmount) returns()
func (_FcpCollateral *FcpCollateralTransactorSession) SlashCollateral(cpAccount common.Address, slashAmount *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.Contract.SlashCollateral(&_FcpCollateral.TransactOpts, cpAccount, slashAmount)
}

// SlashCollateral0 is a paid mutator transaction binding the contract method 0xb587b82c.
//
// Solidity: function slashCollateral(string taskUid, address cpAccount, uint256 slashAmount) returns()
func (_FcpCollateral *FcpCollateralTransactor) SlashCollateral0(opts *bind.TransactOpts, taskUid string, cpAccount common.Address, slashAmount *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.contract.Transact(opts, "slashCollateral0", taskUid, cpAccount, slashAmount)
}

// SlashCollateral0 is a paid mutator transaction binding the contract method 0xb587b82c.
//
// Solidity: function slashCollateral(string taskUid, address cpAccount, uint256 slashAmount) returns()
func (_FcpCollateral *FcpCollateralSession) SlashCollateral0(taskUid string, cpAccount common.Address, slashAmount *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.Contract.SlashCollateral0(&_FcpCollateral.TransactOpts, taskUid, cpAccount, slashAmount)
}

// SlashCollateral0 is a paid mutator transaction binding the contract method 0xb587b82c.
//
// Solidity: function slashCollateral(string taskUid, address cpAccount, uint256 slashAmount) returns()
func (_FcpCollateral *FcpCollateralTransactorSession) SlashCollateral0(taskUid string, cpAccount common.Address, slashAmount *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.Contract.SlashCollateral0(&_FcpCollateral.TransactOpts, taskUid, cpAccount, slashAmount)
}

// SubtractBalance is a paid mutator transaction binding the contract method 0xcfd4f18a.
//
// Solidity: function subtractBalance(address recipient, uint256 amount) returns()
func (_FcpCollateral *FcpCollateralTransactor) SubtractBalance(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.contract.Transact(opts, "subtractBalance", recipient, amount)
}

// SubtractBalance is a paid mutator transaction binding the contract method 0xcfd4f18a.
//
// Solidity: function subtractBalance(address recipient, uint256 amount) returns()
func (_FcpCollateral *FcpCollateralSession) SubtractBalance(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.Contract.SubtractBalance(&_FcpCollateral.TransactOpts, recipient, amount)
}

// SubtractBalance is a paid mutator transaction binding the contract method 0xcfd4f18a.
//
// Solidity: function subtractBalance(address recipient, uint256 amount) returns()
func (_FcpCollateral *FcpCollateralTransactorSession) SubtractBalance(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.Contract.SubtractBalance(&_FcpCollateral.TransactOpts, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FcpCollateral *FcpCollateralTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FcpCollateral.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FcpCollateral *FcpCollateralSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FcpCollateral.Contract.TransferOwnership(&_FcpCollateral.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FcpCollateral *FcpCollateralTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FcpCollateral.Contract.TransferOwnership(&_FcpCollateral.TransactOpts, newOwner)
}

// UnlockCollateral is a paid mutator transaction binding the contract method 0x4b9f0bea.
//
// Solidity: function unlockCollateral(string taskUid, uint256 amount) returns()
func (_FcpCollateral *FcpCollateralTransactor) UnlockCollateral(opts *bind.TransactOpts, taskUid string, amount *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.contract.Transact(opts, "unlockCollateral", taskUid, amount)
}

// UnlockCollateral is a paid mutator transaction binding the contract method 0x4b9f0bea.
//
// Solidity: function unlockCollateral(string taskUid, uint256 amount) returns()
func (_FcpCollateral *FcpCollateralSession) UnlockCollateral(taskUid string, amount *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.Contract.UnlockCollateral(&_FcpCollateral.TransactOpts, taskUid, amount)
}

// UnlockCollateral is a paid mutator transaction binding the contract method 0x4b9f0bea.
//
// Solidity: function unlockCollateral(string taskUid, uint256 amount) returns()
func (_FcpCollateral *FcpCollateralTransactorSession) UnlockCollateral(taskUid string, amount *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.Contract.UnlockCollateral(&_FcpCollateral.TransactOpts, taskUid, amount)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FcpCollateral *FcpCollateralTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FcpCollateral.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FcpCollateral *FcpCollateralSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FcpCollateral.Contract.UpgradeToAndCall(&_FcpCollateral.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FcpCollateral *FcpCollateralTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FcpCollateral.Contract.UpgradeToAndCall(&_FcpCollateral.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address cpAccount, uint256 withdrawAmount) returns()
func (_FcpCollateral *FcpCollateralTransactor) Withdraw(opts *bind.TransactOpts, cpAccount common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.contract.Transact(opts, "withdraw", cpAccount, withdrawAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address cpAccount, uint256 withdrawAmount) returns()
func (_FcpCollateral *FcpCollateralSession) Withdraw(cpAccount common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.Contract.Withdraw(&_FcpCollateral.TransactOpts, cpAccount, withdrawAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address cpAccount, uint256 withdrawAmount) returns()
func (_FcpCollateral *FcpCollateralTransactorSession) Withdraw(cpAccount common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.Contract.Withdraw(&_FcpCollateral.TransactOpts, cpAccount, withdrawAmount)
}

// WithdrawSlashedFunds is a paid mutator transaction binding the contract method 0x2894493f.
//
// Solidity: function withdrawSlashedFunds(uint256 amount) returns()
func (_FcpCollateral *FcpCollateralTransactor) WithdrawSlashedFunds(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.contract.Transact(opts, "withdrawSlashedFunds", amount)
}

// WithdrawSlashedFunds is a paid mutator transaction binding the contract method 0x2894493f.
//
// Solidity: function withdrawSlashedFunds(uint256 amount) returns()
func (_FcpCollateral *FcpCollateralSession) WithdrawSlashedFunds(amount *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.Contract.WithdrawSlashedFunds(&_FcpCollateral.TransactOpts, amount)
}

// WithdrawSlashedFunds is a paid mutator transaction binding the contract method 0x2894493f.
//
// Solidity: function withdrawSlashedFunds(uint256 amount) returns()
func (_FcpCollateral *FcpCollateralTransactorSession) WithdrawSlashedFunds(amount *big.Int) (*types.Transaction, error) {
	return _FcpCollateral.Contract.WithdrawSlashedFunds(&_FcpCollateral.TransactOpts, amount)
}

// FcpCollateralCollateralAdjustedIterator is returned from FilterCollateralAdjusted and is used to iterate over the raw logs and unpacked data for CollateralAdjusted events raised by the FcpCollateral contract.
type FcpCollateralCollateralAdjustedIterator struct {
	Event *FcpCollateralCollateralAdjusted // Event containing the contract specifics and raw log

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
func (it *FcpCollateralCollateralAdjustedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcpCollateralCollateralAdjusted)
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
		it.Event = new(FcpCollateralCollateralAdjusted)
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
func (it *FcpCollateralCollateralAdjustedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcpCollateralCollateralAdjustedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcpCollateralCollateralAdjusted represents a CollateralAdjusted event raised by the FcpCollateral contract.
type FcpCollateralCollateralAdjusted struct {
	Cp            common.Address
	FrozenAmount  *big.Int
	BalanceAmount *big.Int
	Operation     string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCollateralAdjusted is a free log retrieval operation binding the contract event 0x42f1a8a1aee108e84e4eabfaa8d51d7baaa1a02d482295297883a44b2debd3e4.
//
// Solidity: event CollateralAdjusted(address indexed cp, uint256 frozenAmount, uint256 balanceAmount, string operation)
func (_FcpCollateral *FcpCollateralFilterer) FilterCollateralAdjusted(opts *bind.FilterOpts, cp []common.Address) (*FcpCollateralCollateralAdjustedIterator, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _FcpCollateral.contract.FilterLogs(opts, "CollateralAdjusted", cpRule)
	if err != nil {
		return nil, err
	}
	return &FcpCollateralCollateralAdjustedIterator{contract: _FcpCollateral.contract, event: "CollateralAdjusted", logs: logs, sub: sub}, nil
}

// WatchCollateralAdjusted is a free log subscription operation binding the contract event 0x42f1a8a1aee108e84e4eabfaa8d51d7baaa1a02d482295297883a44b2debd3e4.
//
// Solidity: event CollateralAdjusted(address indexed cp, uint256 frozenAmount, uint256 balanceAmount, string operation)
func (_FcpCollateral *FcpCollateralFilterer) WatchCollateralAdjusted(opts *bind.WatchOpts, sink chan<- *FcpCollateralCollateralAdjusted, cp []common.Address) (event.Subscription, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _FcpCollateral.contract.WatchLogs(opts, "CollateralAdjusted", cpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcpCollateralCollateralAdjusted)
				if err := _FcpCollateral.contract.UnpackLog(event, "CollateralAdjusted", log); err != nil {
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
func (_FcpCollateral *FcpCollateralFilterer) ParseCollateralAdjusted(log types.Log) (*FcpCollateralCollateralAdjusted, error) {
	event := new(FcpCollateralCollateralAdjusted)
	if err := _FcpCollateral.contract.UnpackLog(event, "CollateralAdjusted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcpCollateralCollateralLockedIterator is returned from FilterCollateralLocked and is used to iterate over the raw logs and unpacked data for CollateralLocked events raised by the FcpCollateral contract.
type FcpCollateralCollateralLockedIterator struct {
	Event *FcpCollateralCollateralLocked // Event containing the contract specifics and raw log

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
func (it *FcpCollateralCollateralLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcpCollateralCollateralLocked)
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
		it.Event = new(FcpCollateralCollateralLocked)
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
func (it *FcpCollateralCollateralLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcpCollateralCollateralLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcpCollateralCollateralLocked represents a CollateralLocked event raised by the FcpCollateral contract.
type FcpCollateralCollateralLocked struct {
	Cp               common.Address
	CollateralAmount *big.Int
	TaskUid          string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCollateralLocked is a free log retrieval operation binding the contract event 0x5f3d004cf9164b95ed5dbf47d1f04018a4eabcb20b4320fe229ed92236ace634.
//
// Solidity: event CollateralLocked(address indexed cp, uint256 collateralAmount, string taskUid)
func (_FcpCollateral *FcpCollateralFilterer) FilterCollateralLocked(opts *bind.FilterOpts, cp []common.Address) (*FcpCollateralCollateralLockedIterator, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _FcpCollateral.contract.FilterLogs(opts, "CollateralLocked", cpRule)
	if err != nil {
		return nil, err
	}
	return &FcpCollateralCollateralLockedIterator{contract: _FcpCollateral.contract, event: "CollateralLocked", logs: logs, sub: sub}, nil
}

// WatchCollateralLocked is a free log subscription operation binding the contract event 0x5f3d004cf9164b95ed5dbf47d1f04018a4eabcb20b4320fe229ed92236ace634.
//
// Solidity: event CollateralLocked(address indexed cp, uint256 collateralAmount, string taskUid)
func (_FcpCollateral *FcpCollateralFilterer) WatchCollateralLocked(opts *bind.WatchOpts, sink chan<- *FcpCollateralCollateralLocked, cp []common.Address) (event.Subscription, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _FcpCollateral.contract.WatchLogs(opts, "CollateralLocked", cpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcpCollateralCollateralLocked)
				if err := _FcpCollateral.contract.UnpackLog(event, "CollateralLocked", log); err != nil {
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
func (_FcpCollateral *FcpCollateralFilterer) ParseCollateralLocked(log types.Log) (*FcpCollateralCollateralLocked, error) {
	event := new(FcpCollateralCollateralLocked)
	if err := _FcpCollateral.contract.UnpackLog(event, "CollateralLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcpCollateralCollateralSlashedIterator is returned from FilterCollateralSlashed and is used to iterate over the raw logs and unpacked data for CollateralSlashed events raised by the FcpCollateral contract.
type FcpCollateralCollateralSlashedIterator struct {
	Event *FcpCollateralCollateralSlashed // Event containing the contract specifics and raw log

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
func (it *FcpCollateralCollateralSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcpCollateralCollateralSlashed)
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
		it.Event = new(FcpCollateralCollateralSlashed)
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
func (it *FcpCollateralCollateralSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcpCollateralCollateralSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcpCollateralCollateralSlashed represents a CollateralSlashed event raised by the FcpCollateral contract.
type FcpCollateralCollateralSlashed struct {
	Cp      common.Address
	Amount  *big.Int
	TaskUid string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCollateralSlashed is a free log retrieval operation binding the contract event 0x403feb2cd85cc25c910fe59289105b583f08fe9e4335ebbb08c1851f47ff2d34.
//
// Solidity: event CollateralSlashed(address indexed cp, uint256 amount, string taskUid)
func (_FcpCollateral *FcpCollateralFilterer) FilterCollateralSlashed(opts *bind.FilterOpts, cp []common.Address) (*FcpCollateralCollateralSlashedIterator, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _FcpCollateral.contract.FilterLogs(opts, "CollateralSlashed", cpRule)
	if err != nil {
		return nil, err
	}
	return &FcpCollateralCollateralSlashedIterator{contract: _FcpCollateral.contract, event: "CollateralSlashed", logs: logs, sub: sub}, nil
}

// WatchCollateralSlashed is a free log subscription operation binding the contract event 0x403feb2cd85cc25c910fe59289105b583f08fe9e4335ebbb08c1851f47ff2d34.
//
// Solidity: event CollateralSlashed(address indexed cp, uint256 amount, string taskUid)
func (_FcpCollateral *FcpCollateralFilterer) WatchCollateralSlashed(opts *bind.WatchOpts, sink chan<- *FcpCollateralCollateralSlashed, cp []common.Address) (event.Subscription, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _FcpCollateral.contract.WatchLogs(opts, "CollateralSlashed", cpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcpCollateralCollateralSlashed)
				if err := _FcpCollateral.contract.UnpackLog(event, "CollateralSlashed", log); err != nil {
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
func (_FcpCollateral *FcpCollateralFilterer) ParseCollateralSlashed(log types.Log) (*FcpCollateralCollateralSlashed, error) {
	event := new(FcpCollateralCollateralSlashed)
	if err := _FcpCollateral.contract.UnpackLog(event, "CollateralSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcpCollateralCollateralStatusChangedIterator is returned from FilterCollateralStatusChanged and is used to iterate over the raw logs and unpacked data for CollateralStatusChanged events raised by the FcpCollateral contract.
type FcpCollateralCollateralStatusChangedIterator struct {
	Event *FcpCollateralCollateralStatusChanged // Event containing the contract specifics and raw log

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
func (it *FcpCollateralCollateralStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcpCollateralCollateralStatusChanged)
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
		it.Event = new(FcpCollateralCollateralStatusChanged)
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
func (it *FcpCollateralCollateralStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcpCollateralCollateralStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcpCollateralCollateralStatusChanged represents a CollateralStatusChanged event raised by the FcpCollateral contract.
type FcpCollateralCollateralStatusChanged struct {
	TaskUid   common.Hash
	NewStatus uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollateralStatusChanged is a free log retrieval operation binding the contract event 0x4a2ced9ada462e244851a86e998eb0b5bf558c2c9c6923b7f970ed2b19b073b0.
//
// Solidity: event CollateralStatusChanged(string indexed taskUid, uint8 newStatus)
func (_FcpCollateral *FcpCollateralFilterer) FilterCollateralStatusChanged(opts *bind.FilterOpts, taskUid []string) (*FcpCollateralCollateralStatusChangedIterator, error) {

	var taskUidRule []interface{}
	for _, taskUidItem := range taskUid {
		taskUidRule = append(taskUidRule, taskUidItem)
	}

	logs, sub, err := _FcpCollateral.contract.FilterLogs(opts, "CollateralStatusChanged", taskUidRule)
	if err != nil {
		return nil, err
	}
	return &FcpCollateralCollateralStatusChangedIterator{contract: _FcpCollateral.contract, event: "CollateralStatusChanged", logs: logs, sub: sub}, nil
}

// WatchCollateralStatusChanged is a free log subscription operation binding the contract event 0x4a2ced9ada462e244851a86e998eb0b5bf558c2c9c6923b7f970ed2b19b073b0.
//
// Solidity: event CollateralStatusChanged(string indexed taskUid, uint8 newStatus)
func (_FcpCollateral *FcpCollateralFilterer) WatchCollateralStatusChanged(opts *bind.WatchOpts, sink chan<- *FcpCollateralCollateralStatusChanged, taskUid []string) (event.Subscription, error) {

	var taskUidRule []interface{}
	for _, taskUidItem := range taskUid {
		taskUidRule = append(taskUidRule, taskUidItem)
	}

	logs, sub, err := _FcpCollateral.contract.WatchLogs(opts, "CollateralStatusChanged", taskUidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcpCollateralCollateralStatusChanged)
				if err := _FcpCollateral.contract.UnpackLog(event, "CollateralStatusChanged", log); err != nil {
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
func (_FcpCollateral *FcpCollateralFilterer) ParseCollateralStatusChanged(log types.Log) (*FcpCollateralCollateralStatusChanged, error) {
	event := new(FcpCollateralCollateralStatusChanged)
	if err := _FcpCollateral.contract.UnpackLog(event, "CollateralStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcpCollateralCollateralUnlockedIterator is returned from FilterCollateralUnlocked and is used to iterate over the raw logs and unpacked data for CollateralUnlocked events raised by the FcpCollateral contract.
type FcpCollateralCollateralUnlockedIterator struct {
	Event *FcpCollateralCollateralUnlocked // Event containing the contract specifics and raw log

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
func (it *FcpCollateralCollateralUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcpCollateralCollateralUnlocked)
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
		it.Event = new(FcpCollateralCollateralUnlocked)
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
func (it *FcpCollateralCollateralUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcpCollateralCollateralUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcpCollateralCollateralUnlocked represents a CollateralUnlocked event raised by the FcpCollateral contract.
type FcpCollateralCollateralUnlocked struct {
	Cp               common.Address
	CollateralAmount *big.Int
	TaskUid          string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCollateralUnlocked is a free log retrieval operation binding the contract event 0x0f2440b3ca071b7d18e917a25289e7d7e7de8a491546d45acc2efbec7b3e1ae8.
//
// Solidity: event CollateralUnlocked(address indexed cp, uint256 collateralAmount, string taskUid)
func (_FcpCollateral *FcpCollateralFilterer) FilterCollateralUnlocked(opts *bind.FilterOpts, cp []common.Address) (*FcpCollateralCollateralUnlockedIterator, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _FcpCollateral.contract.FilterLogs(opts, "CollateralUnlocked", cpRule)
	if err != nil {
		return nil, err
	}
	return &FcpCollateralCollateralUnlockedIterator{contract: _FcpCollateral.contract, event: "CollateralUnlocked", logs: logs, sub: sub}, nil
}

// WatchCollateralUnlocked is a free log subscription operation binding the contract event 0x0f2440b3ca071b7d18e917a25289e7d7e7de8a491546d45acc2efbec7b3e1ae8.
//
// Solidity: event CollateralUnlocked(address indexed cp, uint256 collateralAmount, string taskUid)
func (_FcpCollateral *FcpCollateralFilterer) WatchCollateralUnlocked(opts *bind.WatchOpts, sink chan<- *FcpCollateralCollateralUnlocked, cp []common.Address) (event.Subscription, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _FcpCollateral.contract.WatchLogs(opts, "CollateralUnlocked", cpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcpCollateralCollateralUnlocked)
				if err := _FcpCollateral.contract.UnpackLog(event, "CollateralUnlocked", log); err != nil {
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
func (_FcpCollateral *FcpCollateralFilterer) ParseCollateralUnlocked(log types.Log) (*FcpCollateralCollateralUnlocked, error) {
	event := new(FcpCollateralCollateralUnlocked)
	if err := _FcpCollateral.contract.UnpackLog(event, "CollateralUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcpCollateralDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the FcpCollateral contract.
type FcpCollateralDepositIterator struct {
	Event *FcpCollateralDeposit // Event containing the contract specifics and raw log

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
func (it *FcpCollateralDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcpCollateralDeposit)
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
		it.Event = new(FcpCollateralDeposit)
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
func (it *FcpCollateralDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcpCollateralDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcpCollateralDeposit represents a Deposit event raised by the FcpCollateral contract.
type FcpCollateralDeposit struct {
	FundingWallet common.Address
	CpAccount     common.Address
	DepositAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed fundingWallet, address indexed cpAccount, uint256 depositAmount)
func (_FcpCollateral *FcpCollateralFilterer) FilterDeposit(opts *bind.FilterOpts, fundingWallet []common.Address, cpAccount []common.Address) (*FcpCollateralDepositIterator, error) {

	var fundingWalletRule []interface{}
	for _, fundingWalletItem := range fundingWallet {
		fundingWalletRule = append(fundingWalletRule, fundingWalletItem)
	}
	var cpAccountRule []interface{}
	for _, cpAccountItem := range cpAccount {
		cpAccountRule = append(cpAccountRule, cpAccountItem)
	}

	logs, sub, err := _FcpCollateral.contract.FilterLogs(opts, "Deposit", fundingWalletRule, cpAccountRule)
	if err != nil {
		return nil, err
	}
	return &FcpCollateralDepositIterator{contract: _FcpCollateral.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed fundingWallet, address indexed cpAccount, uint256 depositAmount)
func (_FcpCollateral *FcpCollateralFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *FcpCollateralDeposit, fundingWallet []common.Address, cpAccount []common.Address) (event.Subscription, error) {

	var fundingWalletRule []interface{}
	for _, fundingWalletItem := range fundingWallet {
		fundingWalletRule = append(fundingWalletRule, fundingWalletItem)
	}
	var cpAccountRule []interface{}
	for _, cpAccountItem := range cpAccount {
		cpAccountRule = append(cpAccountRule, cpAccountItem)
	}

	logs, sub, err := _FcpCollateral.contract.WatchLogs(opts, "Deposit", fundingWalletRule, cpAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcpCollateralDeposit)
				if err := _FcpCollateral.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_FcpCollateral *FcpCollateralFilterer) ParseDeposit(log types.Log) (*FcpCollateralDeposit, error) {
	event := new(FcpCollateralDeposit)
	if err := _FcpCollateral.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcpCollateralDisputeProofIterator is returned from FilterDisputeProof and is used to iterate over the raw logs and unpacked data for DisputeProof events raised by the FcpCollateral contract.
type FcpCollateralDisputeProofIterator struct {
	Event *FcpCollateralDisputeProof // Event containing the contract specifics and raw log

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
func (it *FcpCollateralDisputeProofIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcpCollateralDisputeProof)
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
		it.Event = new(FcpCollateralDisputeProof)
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
func (it *FcpCollateralDisputeProofIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcpCollateralDisputeProofIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcpCollateralDisputeProof represents a DisputeProof event raised by the FcpCollateral contract.
type FcpCollateralDisputeProof struct {
	Challenger common.Address
	TaskUid    common.Hash
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDisputeProof is a free log retrieval operation binding the contract event 0xaec1d412a3c1e4a13fc2a2e19ac38a5af192a9cf17b074fca8146a2d0655e0c3.
//
// Solidity: event DisputeProof(address indexed challenger, string indexed taskUid)
func (_FcpCollateral *FcpCollateralFilterer) FilterDisputeProof(opts *bind.FilterOpts, challenger []common.Address, taskUid []string) (*FcpCollateralDisputeProofIterator, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}
	var taskUidRule []interface{}
	for _, taskUidItem := range taskUid {
		taskUidRule = append(taskUidRule, taskUidItem)
	}

	logs, sub, err := _FcpCollateral.contract.FilterLogs(opts, "DisputeProof", challengerRule, taskUidRule)
	if err != nil {
		return nil, err
	}
	return &FcpCollateralDisputeProofIterator{contract: _FcpCollateral.contract, event: "DisputeProof", logs: logs, sub: sub}, nil
}

// WatchDisputeProof is a free log subscription operation binding the contract event 0xaec1d412a3c1e4a13fc2a2e19ac38a5af192a9cf17b074fca8146a2d0655e0c3.
//
// Solidity: event DisputeProof(address indexed challenger, string indexed taskUid)
func (_FcpCollateral *FcpCollateralFilterer) WatchDisputeProof(opts *bind.WatchOpts, sink chan<- *FcpCollateralDisputeProof, challenger []common.Address, taskUid []string) (event.Subscription, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}
	var taskUidRule []interface{}
	for _, taskUidItem := range taskUid {
		taskUidRule = append(taskUidRule, taskUidItem)
	}

	logs, sub, err := _FcpCollateral.contract.WatchLogs(opts, "DisputeProof", challengerRule, taskUidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcpCollateralDisputeProof)
				if err := _FcpCollateral.contract.UnpackLog(event, "DisputeProof", log); err != nil {
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
func (_FcpCollateral *FcpCollateralFilterer) ParseDisputeProof(log types.Log) (*FcpCollateralDisputeProof, error) {
	event := new(FcpCollateralDisputeProof)
	if err := _FcpCollateral.contract.UnpackLog(event, "DisputeProof", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcpCollateralInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the FcpCollateral contract.
type FcpCollateralInitializedIterator struct {
	Event *FcpCollateralInitialized // Event containing the contract specifics and raw log

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
func (it *FcpCollateralInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcpCollateralInitialized)
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
		it.Event = new(FcpCollateralInitialized)
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
func (it *FcpCollateralInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcpCollateralInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcpCollateralInitialized represents a Initialized event raised by the FcpCollateral contract.
type FcpCollateralInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FcpCollateral *FcpCollateralFilterer) FilterInitialized(opts *bind.FilterOpts) (*FcpCollateralInitializedIterator, error) {

	logs, sub, err := _FcpCollateral.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &FcpCollateralInitializedIterator{contract: _FcpCollateral.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FcpCollateral *FcpCollateralFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *FcpCollateralInitialized) (event.Subscription, error) {

	logs, sub, err := _FcpCollateral.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcpCollateralInitialized)
				if err := _FcpCollateral.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_FcpCollateral *FcpCollateralFilterer) ParseInitialized(log types.Log) (*FcpCollateralInitialized, error) {
	event := new(FcpCollateralInitialized)
	if err := _FcpCollateral.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcpCollateralOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FcpCollateral contract.
type FcpCollateralOwnershipTransferredIterator struct {
	Event *FcpCollateralOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FcpCollateralOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcpCollateralOwnershipTransferred)
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
		it.Event = new(FcpCollateralOwnershipTransferred)
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
func (it *FcpCollateralOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcpCollateralOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcpCollateralOwnershipTransferred represents a OwnershipTransferred event raised by the FcpCollateral contract.
type FcpCollateralOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FcpCollateral *FcpCollateralFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FcpCollateralOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FcpCollateral.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FcpCollateralOwnershipTransferredIterator{contract: _FcpCollateral.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FcpCollateral *FcpCollateralFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FcpCollateralOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FcpCollateral.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcpCollateralOwnershipTransferred)
				if err := _FcpCollateral.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FcpCollateral *FcpCollateralFilterer) ParseOwnershipTransferred(log types.Log) (*FcpCollateralOwnershipTransferred, error) {
	event := new(FcpCollateralOwnershipTransferred)
	if err := _FcpCollateral.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcpCollateralSlashedFundsIncreasedIterator is returned from FilterSlashedFundsIncreased and is used to iterate over the raw logs and unpacked data for SlashedFundsIncreased events raised by the FcpCollateral contract.
type FcpCollateralSlashedFundsIncreasedIterator struct {
	Event *FcpCollateralSlashedFundsIncreased // Event containing the contract specifics and raw log

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
func (it *FcpCollateralSlashedFundsIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcpCollateralSlashedFundsIncreased)
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
		it.Event = new(FcpCollateralSlashedFundsIncreased)
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
func (it *FcpCollateralSlashedFundsIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcpCollateralSlashedFundsIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcpCollateralSlashedFundsIncreased represents a SlashedFundsIncreased event raised by the FcpCollateral contract.
type FcpCollateralSlashedFundsIncreased struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSlashedFundsIncreased is a free log retrieval operation binding the contract event 0xe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b.
//
// Solidity: event SlashedFundsIncreased(uint256 amount)
func (_FcpCollateral *FcpCollateralFilterer) FilterSlashedFundsIncreased(opts *bind.FilterOpts) (*FcpCollateralSlashedFundsIncreasedIterator, error) {

	logs, sub, err := _FcpCollateral.contract.FilterLogs(opts, "SlashedFundsIncreased")
	if err != nil {
		return nil, err
	}
	return &FcpCollateralSlashedFundsIncreasedIterator{contract: _FcpCollateral.contract, event: "SlashedFundsIncreased", logs: logs, sub: sub}, nil
}

// WatchSlashedFundsIncreased is a free log subscription operation binding the contract event 0xe69f9e72017aaa026e6e6b9186b7b1e197ec951c6c7df9805190316cb8e6f34b.
//
// Solidity: event SlashedFundsIncreased(uint256 amount)
func (_FcpCollateral *FcpCollateralFilterer) WatchSlashedFundsIncreased(opts *bind.WatchOpts, sink chan<- *FcpCollateralSlashedFundsIncreased) (event.Subscription, error) {

	logs, sub, err := _FcpCollateral.contract.WatchLogs(opts, "SlashedFundsIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcpCollateralSlashedFundsIncreased)
				if err := _FcpCollateral.contract.UnpackLog(event, "SlashedFundsIncreased", log); err != nil {
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
func (_FcpCollateral *FcpCollateralFilterer) ParseSlashedFundsIncreased(log types.Log) (*FcpCollateralSlashedFundsIncreased, error) {
	event := new(FcpCollateralSlashedFundsIncreased)
	if err := _FcpCollateral.contract.UnpackLog(event, "SlashedFundsIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcpCollateralTaskCreatedIterator is returned from FilterTaskCreated and is used to iterate over the raw logs and unpacked data for TaskCreated events raised by the FcpCollateral contract.
type FcpCollateralTaskCreatedIterator struct {
	Event *FcpCollateralTaskCreated // Event containing the contract specifics and raw log

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
func (it *FcpCollateralTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcpCollateralTaskCreated)
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
		it.Event = new(FcpCollateralTaskCreated)
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
func (it *FcpCollateralTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcpCollateralTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcpCollateralTaskCreated represents a TaskCreated event raised by the FcpCollateral contract.
type FcpCollateralTaskCreated struct {
	TaskUid          common.Hash
	CpAccountAddress common.Address
	Collateral       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTaskCreated is a free log retrieval operation binding the contract event 0x5bebc56a5428fd7b8cf43ed525f03223f8363907fbe44665b7a3426d1de96800.
//
// Solidity: event TaskCreated(string indexed taskUid, address cpAccountAddress, uint256 collateral)
func (_FcpCollateral *FcpCollateralFilterer) FilterTaskCreated(opts *bind.FilterOpts, taskUid []string) (*FcpCollateralTaskCreatedIterator, error) {

	var taskUidRule []interface{}
	for _, taskUidItem := range taskUid {
		taskUidRule = append(taskUidRule, taskUidItem)
	}

	logs, sub, err := _FcpCollateral.contract.FilterLogs(opts, "TaskCreated", taskUidRule)
	if err != nil {
		return nil, err
	}
	return &FcpCollateralTaskCreatedIterator{contract: _FcpCollateral.contract, event: "TaskCreated", logs: logs, sub: sub}, nil
}

// WatchTaskCreated is a free log subscription operation binding the contract event 0x5bebc56a5428fd7b8cf43ed525f03223f8363907fbe44665b7a3426d1de96800.
//
// Solidity: event TaskCreated(string indexed taskUid, address cpAccountAddress, uint256 collateral)
func (_FcpCollateral *FcpCollateralFilterer) WatchTaskCreated(opts *bind.WatchOpts, sink chan<- *FcpCollateralTaskCreated, taskUid []string) (event.Subscription, error) {

	var taskUidRule []interface{}
	for _, taskUidItem := range taskUid {
		taskUidRule = append(taskUidRule, taskUidItem)
	}

	logs, sub, err := _FcpCollateral.contract.WatchLogs(opts, "TaskCreated", taskUidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcpCollateralTaskCreated)
				if err := _FcpCollateral.contract.UnpackLog(event, "TaskCreated", log); err != nil {
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
func (_FcpCollateral *FcpCollateralFilterer) ParseTaskCreated(log types.Log) (*FcpCollateralTaskCreated, error) {
	event := new(FcpCollateralTaskCreated)
	if err := _FcpCollateral.contract.UnpackLog(event, "TaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcpCollateralUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the FcpCollateral contract.
type FcpCollateralUpgradedIterator struct {
	Event *FcpCollateralUpgraded // Event containing the contract specifics and raw log

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
func (it *FcpCollateralUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcpCollateralUpgraded)
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
		it.Event = new(FcpCollateralUpgraded)
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
func (it *FcpCollateralUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcpCollateralUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcpCollateralUpgraded represents a Upgraded event raised by the FcpCollateral contract.
type FcpCollateralUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FcpCollateral *FcpCollateralFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*FcpCollateralUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _FcpCollateral.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &FcpCollateralUpgradedIterator{contract: _FcpCollateral.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FcpCollateral *FcpCollateralFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *FcpCollateralUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _FcpCollateral.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcpCollateralUpgraded)
				if err := _FcpCollateral.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_FcpCollateral *FcpCollateralFilterer) ParseUpgraded(log types.Log) (*FcpCollateralUpgraded, error) {
	event := new(FcpCollateralUpgraded)
	if err := _FcpCollateral.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcpCollateralWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the FcpCollateral contract.
type FcpCollateralWithdrawIterator struct {
	Event *FcpCollateralWithdraw // Event containing the contract specifics and raw log

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
func (it *FcpCollateralWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcpCollateralWithdraw)
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
		it.Event = new(FcpCollateralWithdraw)
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
func (it *FcpCollateralWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcpCollateralWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcpCollateralWithdraw represents a Withdraw event raised by the FcpCollateral contract.
type FcpCollateralWithdraw struct {
	CpOwner        common.Address
	CpAccount      common.Address
	WithdrawAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed cpOwner, address indexed cpAccount, uint256 withdrawAmount)
func (_FcpCollateral *FcpCollateralFilterer) FilterWithdraw(opts *bind.FilterOpts, cpOwner []common.Address, cpAccount []common.Address) (*FcpCollateralWithdrawIterator, error) {

	var cpOwnerRule []interface{}
	for _, cpOwnerItem := range cpOwner {
		cpOwnerRule = append(cpOwnerRule, cpOwnerItem)
	}
	var cpAccountRule []interface{}
	for _, cpAccountItem := range cpAccount {
		cpAccountRule = append(cpAccountRule, cpAccountItem)
	}

	logs, sub, err := _FcpCollateral.contract.FilterLogs(opts, "Withdraw", cpOwnerRule, cpAccountRule)
	if err != nil {
		return nil, err
	}
	return &FcpCollateralWithdrawIterator{contract: _FcpCollateral.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed cpOwner, address indexed cpAccount, uint256 withdrawAmount)
func (_FcpCollateral *FcpCollateralFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *FcpCollateralWithdraw, cpOwner []common.Address, cpAccount []common.Address) (event.Subscription, error) {

	var cpOwnerRule []interface{}
	for _, cpOwnerItem := range cpOwner {
		cpOwnerRule = append(cpOwnerRule, cpOwnerItem)
	}
	var cpAccountRule []interface{}
	for _, cpAccountItem := range cpAccount {
		cpAccountRule = append(cpAccountRule, cpAccountItem)
	}

	logs, sub, err := _FcpCollateral.contract.WatchLogs(opts, "Withdraw", cpOwnerRule, cpAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcpCollateralWithdraw)
				if err := _FcpCollateral.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_FcpCollateral *FcpCollateralFilterer) ParseWithdraw(log types.Log) (*FcpCollateralWithdraw, error) {
	event := new(FcpCollateralWithdraw)
	if err := _FcpCollateral.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcpCollateralWithdrawSlashIterator is returned from FilterWithdrawSlash and is used to iterate over the raw logs and unpacked data for WithdrawSlash events raised by the FcpCollateral contract.
type FcpCollateralWithdrawSlashIterator struct {
	Event *FcpCollateralWithdrawSlash // Event containing the contract specifics and raw log

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
func (it *FcpCollateralWithdrawSlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcpCollateralWithdrawSlash)
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
		it.Event = new(FcpCollateralWithdrawSlash)
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
func (it *FcpCollateralWithdrawSlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcpCollateralWithdrawSlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcpCollateralWithdrawSlash represents a WithdrawSlash event raised by the FcpCollateral contract.
type FcpCollateralWithdrawSlash struct {
	CollateralContratOwner common.Address
	Slashfund              *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawSlash is a free log retrieval operation binding the contract event 0xbfd9c82485e2178fcfd5c40379d8e66fe60acc50aa1ef4c50966431eb1e381cd.
//
// Solidity: event WithdrawSlash(address indexed collateralContratOwner, uint256 slashfund)
func (_FcpCollateral *FcpCollateralFilterer) FilterWithdrawSlash(opts *bind.FilterOpts, collateralContratOwner []common.Address) (*FcpCollateralWithdrawSlashIterator, error) {

	var collateralContratOwnerRule []interface{}
	for _, collateralContratOwnerItem := range collateralContratOwner {
		collateralContratOwnerRule = append(collateralContratOwnerRule, collateralContratOwnerItem)
	}

	logs, sub, err := _FcpCollateral.contract.FilterLogs(opts, "WithdrawSlash", collateralContratOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FcpCollateralWithdrawSlashIterator{contract: _FcpCollateral.contract, event: "WithdrawSlash", logs: logs, sub: sub}, nil
}

// WatchWithdrawSlash is a free log subscription operation binding the contract event 0xbfd9c82485e2178fcfd5c40379d8e66fe60acc50aa1ef4c50966431eb1e381cd.
//
// Solidity: event WithdrawSlash(address indexed collateralContratOwner, uint256 slashfund)
func (_FcpCollateral *FcpCollateralFilterer) WatchWithdrawSlash(opts *bind.WatchOpts, sink chan<- *FcpCollateralWithdrawSlash, collateralContratOwner []common.Address) (event.Subscription, error) {

	var collateralContratOwnerRule []interface{}
	for _, collateralContratOwnerItem := range collateralContratOwner {
		collateralContratOwnerRule = append(collateralContratOwnerRule, collateralContratOwnerItem)
	}

	logs, sub, err := _FcpCollateral.contract.WatchLogs(opts, "WithdrawSlash", collateralContratOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcpCollateralWithdrawSlash)
				if err := _FcpCollateral.contract.UnpackLog(event, "WithdrawSlash", log); err != nil {
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
func (_FcpCollateral *FcpCollateralFilterer) ParseWithdrawSlash(log types.Log) (*FcpCollateralWithdrawSlash, error) {
	event := new(FcpCollateralWithdrawSlash)
	if err := _FcpCollateral.contract.UnpackLog(event, "WithdrawSlash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
