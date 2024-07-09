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

// ECPCollateralCPInfo is an auto generated low-level Go binding around an user-defined struct.
type ECPCollateralCPInfo struct {
	Cp            common.Address
	Balance       *big.Int
	FrozenBalance *big.Int
	Status        string
}

// ECPCollateralContractInfo is an auto generated low-level Go binding around an user-defined struct.
type ECPCollateralContractInfo struct {
	CollateralToken common.Address
	SlashedFunds    *big.Int
	BaseCollateral  *big.Int
	CollateralRatio *big.Int
	SlashRatio      *big.Int
	WithdrawDelay   *big.Int
}

// ECPCollateralWithdrawRequest is an auto generated low-level Go binding around an user-defined struct.
type ECPCollateralWithdrawRequest struct {
	Amount       *big.Int
	RequestBlock *big.Int
}

// EcpCollateralMetaData contains all meta data concerning the EcpCollateral contract.
var EcpCollateralMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"frozenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"operation\",\"type\":\"string\"}],\"name\":\"CollateralAdjusted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"CollateralLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CollateralSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"CollateralUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fundingWallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taskContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"}],\"name\":\"DisputeProof\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawRequestCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralContratOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashfund\",\"type\":\"uint256\"}],\"name\":\"WithdrawSlash\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"cps\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"taskCollaterals\",\"type\":\"uint256[]\"}],\"name\":\"batchLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"cps\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"slashAmounts\",\"type\":\"uint256[]\"}],\"name\":\"batchSlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"cps\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"taskCollaterals\",\"type\":\"uint256[]\"}],\"name\":\"batchUnlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"}],\"name\":\"cancelWithdrawRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"}],\"name\":\"confirmWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"}],\"name\":\"cpInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"balance\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"frozenBalance\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"}],\"internalType\":\"structECPCollateral.CPInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cpStatus\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"taskContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"}],\"name\":\"disputeProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"frozenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getECPCollateralInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slashedFunds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawDelay\",\"type\":\"uint256\"}],\"internalType\":\"structECPCollateral.ContractInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"taskCollateral\",\"type\":\"uint256\"}],\"name\":\"lockCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"requestWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseCollateral\",\"type\":\"uint256\"}],\"name\":\"setBaseCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_collateralRatio\",\"type\":\"uint256\"}],\"name\":\"setCollateralRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"setCollateralToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slashRatio\",\"type\":\"uint256\"}],\"name\":\"setSlashRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawDelay\",\"type\":\"uint256\"}],\"name\":\"setWithdrawDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slashAmount\",\"type\":\"uint256\"}],\"name\":\"slashCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashedFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"taskCollateral\",\"type\":\"uint256\"}],\"name\":\"unlockCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"}],\"name\":\"viewWithdrawRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestBlock\",\"type\":\"uint256\"}],\"internalType\":\"structECPCollateral.WithdrawRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slashfund\",\"type\":\"uint256\"}],\"name\":\"withdrawSlashedFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405261870060065534801561001657600080fd5b506100336100286100af60201b60201c565b6100b760201b60201c565b610042336100b760201b60201c565b6001600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506005600481905550600260058190555061017b565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6141b68061018a6000396000f3fe608060405234801561001057600080fd5b50600436106102325760003560e01c80637048027511610130578063a664c216116100b8578063ce3518aa1161007c578063ce3518aa1461065e578063d27ca89b1461067a578063d2bfc1c714610698578063f2fde38b146106b4578063f3fef3a3146106d057610232565b8063a664c216146105a6578063b2016bd4146105d6578063b4eae1cb146105f4578063bede6e3114610612578063c6ff45551461064257610232565b80637f58a7e5116100ff5780637f58a7e5146105145780638331f8e5146105305780638da5cb5b1461054c5780639939cd181461056a5780639b5ddf091461058857610232565b806370480275146104b457806370b72944146104d0578063715018a6146104ee57806372f0cb30146104f857610232565b80633fe65177116101be5780635f7d0e84116101825780635f7d0e84146104265780636060663e1461044257806363215bb71461045e578063666181a91461047a5780636f99f15c1461049657610232565b80633fe651771461037157806347a7d107146103a157806347e7ef24146103bd57806352df49ec146103d95780635d2cd2a71461040a57610232565b806324d7806c1161020557806324d7806c146102a9578063266565a9146102d957806327e235e3146103095780632894493f14610339578063397a1b281461035557610232565b80630288a39c146102375780631785f53c1461025557806317f17498146102715780631d47a62d1461028d575b600080fd5b61023f6106ec565b60405161024c9190612cb9565b60405180910390f35b61026f600480360381019061026a9190612d3c565b6106f2565b005b61028b60048036038101906102869190612d95565b610755565b005b6102a760048036038101906102a29190612d95565b610941565b005b6102c360048036038101906102be9190612d3c565b610bb6565b6040516102d09190612df0565b60405180910390f35b6102f360048036038101906102ee9190612d3c565b610bd6565b6040516103009190612cb9565b60405180910390f35b610323600480360381019061031e9190612d3c565b610bee565b6040516103309190612e24565b60405180910390f35b610353600480360381019061034e9190612e3f565b610c06565b005b61036f600480360381019061036a9190612d95565b610d5e565b005b61038b60048036038101906103869190612d3c565b611057565b6040516103989190612efc565b60405180910390f35b6103bb60048036038101906103b69190612d95565b6110f7565b005b6103d760048036038101906103d29190612d95565b61130c565b005b6103f360048036038101906103ee9190612d3c565b611477565b604051610401929190612f1e565b60405180910390f35b610424600480360381019061041f9190612d3c565b61149b565b005b610440600480360381019061043b9190613002565b611787565b005b61045c60048036038101906104579190612e3f565b6118c7565b005b61047860048036038101906104739190613002565b6118d9565b005b610494600480360381019061048f9190612d3c565b611a19565b005b61049e611a65565b6040516104ab9190612cb9565b60405180910390f35b6104ce60048036038101906104c99190612d3c565b611a6b565b005b6104d8611ace565b6040516104e59190612cb9565b60405180910390f35b6104f6611ad8565b005b610512600480360381019061050d9190612e3f565b611aec565b005b61052e60048036038101906105299190612e3f565b611afe565b005b61054a60048036038101906105459190613083565b611b94565b005b610554611c00565b60405161056191906130e5565b60405180910390f35b610572611c29565b60405161057f9190613199565b60405180910390f35b610590611ca5565b60405161059d9190612cb9565b60405180910390f35b6105c060048036038101906105bb9190612d3c565b611cab565b6040516105cd9190613270565b60405180910390f35b6105de611e36565b6040516105eb91906132f1565b60405180910390f35b6105fc611e5c565b6040516106099190612cb9565b60405180910390f35b61062c60048036038101906106279190612d3c565b611e62565b604051610639919061333b565b60405180910390f35b61065c60048036038101906106579190613002565b611ecf565b005b61067860048036038101906106739190612e3f565b61200f565b005b610682612021565b60405161068f9190612cb9565b60405180910390f35b6106b260048036038101906106ad9190612d3c565b612027565b005b6106ce60048036038101906106c99190612d3c565b6124ea565b005b6106ea60048036038101906106e59190612d95565b61256d565b005b60065481565b6106fa612912565b6000600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166107e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107d8906133c8565b60405180910390fd5b6000600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060008183116108345782610836565b815b905080600960008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546108879190613417565b9250508190555080600860008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546108dd919061344b565b925050819055506108ed84612990565b8373ffffffffffffffffffffffffffffffffffffffff167fb4eaf47ecd4bc76248f192433e8067c96cb3e17aced42fbc47a512742fb74216826040516109339190612cb9565b60405180910390a250505050565b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166109cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109c4906133c8565b60405180910390fd5b6000600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000818311610a205782610a22565b815b90506000818411610a34576000610a41565b8184610a409190613417565b5b905081600960008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610a929190613417565b9250508190555080600860008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610ae8919061348f565b925050819055508360026000828254610b0191906134d2565b92505081905550610b1185612990565b8473ffffffffffffffffffffffffffffffffffffffff167f5138f522ae83cccdefee151fa33feeb62b6bbe619fdeb8f83cd1c6c3f8bdf92185604051610b579190612cb9565b60405180910390a28473ffffffffffffffffffffffffffffffffffffffff167f42f1a8a1aee108e84e4eabfaa8d51d7baaa1a02d482295297883a44b2debd3e48383604051610ba7929190613552565b60405180910390a25050505050565b60076020528060005260406000206000915054906101000a900460ff1681565b60096020528060005260406000206000915090505481565b60086020528060005260406000206000915090505481565b610c0e612912565b806002541015610c53576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c4a90613600565b60405180910390fd5b8060026000828254610c659190613417565b92505081905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b8152600401610cc9929190613620565b6020604051808303816000875af1158015610ce8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d0c9190613675565b503373ffffffffffffffffffffffffffffffffffffffff167fbfd9c82485e2178fcfd5c40379d8e66fe60acc50aa1ef4c50966431eb1e381cd82604051610d539190612cb9565b60405180910390a250565b6000808373ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527f893d20e8000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051610e0891906136e9565b6000604051808303816000865af19150503d8060008114610e45576040519150601f19603f3d011682016040523d82523d6000602084013e610e4a565b606091505b509150915081610e8f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e8690613772565b60405180910390fd5b600081806020019051810190610ea591906137d0565b90508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f15576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f0c9061386f565b60405180910390fd5b83600960008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610f97576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f8e90613901565b60405180910390fd5b604051806040016040528085815260200143815250600b60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000155602082015181600101559050508473ffffffffffffffffffffffffffffffffffffffff167ff7774b688d56120b783560a913ee60792a73dfd511812b7be5eccf10d08c6689856040516110489190612cb9565b60405180910390a25050505050565b600a602052806000526040600020600091509050805461107690613950565b80601f01602080910402602001604051908101604052809291908181526020018280546110a290613950565b80156110ef5780601f106110c4576101008083540402835291602001916110ef565b820191906000526020600020905b8154815290600101906020018083116110d257829003601f168201915b505050505081565b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611183576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161117a906133c8565b60405180910390fd5b80600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541215611205576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111fc906139f3565b60405180910390fd5b80600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611254919061348f565b9250508190555080600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546112aa91906134d2565b925050819055506112ba82612990565b8173ffffffffffffffffffffffffffffffffffffffff167f2251f6a4ed7fe619e9e8ce557d05a63dd484284f9c95c9ab334f6a7707cd0800826040516113009190612cb9565b60405180910390a25050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b815260040161136b93929190613a13565b6020604051808303816000875af115801561138a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113ae9190613675565b5080600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546113fe919061344b565b925050819055508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62836040516114629190612cb9565b60405180910390a361147382612990565b5050565b600b6020528060005260406000206000915090508060000154908060010154905082565b6000808273ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527f893d20e8000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505060405161154591906136e9565b6000604051808303816000865af19150503d8060008114611582576040519150601f19603f3d011682016040523d82523d6000602084013e611587565b606091505b5091509150816115cc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115c390613772565b60405180910390fd5b6000818060200190518101906115e291906137d0565b90508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611652576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161164990613abc565b60405180910390fd5b6000600b60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060008160000154116116dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116d390613b28565b60405180910390fd5b8473ffffffffffffffffffffffffffffffffffffffff167fa3895d397a34e928a95d593331e293e2fc281d9d8996df5cc6c57f1cef629d4282600001546040516117269190612cb9565b60405180910390a2600b60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600080820160009055600182016000905550505050505050565b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611813576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161180a906133c8565b60405180910390fd5b81819050848490501461185b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161185290613b94565b60405180910390fd5b60005b848490508110156118c0576118b385858381811061187f5761187e613bb4565b5b90506020020160208101906118949190612d3c565b8484848181106118a7576118a6613bb4565b5b90506020020135610755565b808060010191505061185e565b5050505050565b6118cf612912565b8060048190555050565b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611965576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161195c906133c8565b60405180910390fd5b8181905084849050146119ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119a490613b94565b60405180910390fd5b60005b84849050811015611a1257611a058585838181106119d1576119d0613bb4565b5b90506020020160208101906119e69190612d3c565b8484848181106119f9576119f8613bb4565b5b905060200201356110f7565b80806001019150506119b0565b5050505050565b611a21612912565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60025481565b611a73612912565b6001600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b6000600354905090565b611ae0612912565b611aea6000612b30565b565b611af4612912565b8060068190555050565b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611b8a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b81906133c8565b60405180910390fd5b8060038190555050565b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f6010bb1c19b181c59c29bde1a47441eae0c5e2e587b409d5a7ac30f01e8dcf3c8484604051611bf3929190613620565b60405180910390a3505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b611c31612bfc565b6040518060c00160405280600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002548152602001600354815260200160045481526020016005548152602001600654815250905090565b60035481565b611cb3612c48565b60405180608001604052808373ffffffffffffffffffffffffffffffffffffffff168152602001600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548152602001600960008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548152602001600a60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054611dae90613950565b80601f0160208091040260200160405190810160405280929190818152602001828054611dda90613950565b8015611e275780601f10611dfc57610100808354040283529160200191611e27565b820191906000526020600020905b815481529060010190602001808311611e0a57829003601f168201915b50505050508152509050919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60045481565b611e6a612c86565b600b60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806040016040529081600082015481526020016001820154815250509050919050565b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611f5b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f52906133c8565b60405180910390fd5b818190508484905014611fa3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f9a90613b94565b60405180910390fd5b60005b8484905081101561200857611ffb858583818110611fc757611fc6613bb4565b5b9050602002016020810190611fdc9190612d3c565b848484818110611fef57611fee613bb4565b5b90506020020135610941565b8080600101915050611fa6565b5050505050565b612017612912565b8060058190555050565b60055481565b6000808273ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527f893d20e8000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516120d191906136e9565b6000604051808303816000865af19150503d806000811461210e576040519150601f19603f3d011682016040523d82523d6000602084013e612113565b606091505b509150915081612158576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161214f90613772565b60405180910390fd5b60008180602001905181019061216e91906137d0565b90508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146121de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121d590613c55565b60405180910390fd5b6000600b60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000816000015411612268576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161225f90613b28565b60405180910390fd5b600654816001015461227a91906134d2565b4310156122bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122b390613cc1565b60405180910390fd5b8060000154600960008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015612342576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161233990613901565b60405180910390fd5b60008160000154905080600960008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461239a9190613417565b92505081905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb87836040518363ffffffff1660e01b81526004016123fe929190613620565b6020604051808303816000875af115801561241d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124419190613675565b508573ffffffffffffffffffffffffffffffffffffffff167f1a98aba99d2d38026b07feddaca8e333649ae8a5f5a238687f91ce7791ee998e826040516124889190612cb9565b60405180910390a2600b60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008082016000905560018201600090555050505050505050565b6124f2612912565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612561576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161255890613d53565b60405180910390fd5b61256a81612b30565b50565b6000808373ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527f893d20e8000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505060405161261791906136e9565b6000604051808303816000865af19150503d8060008114612654576040519150601f19603f3d011682016040523d82523d6000602084013e612659565b606091505b50915091508161269e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161269590613772565b60405180910390fd5b6000818060200190518101906126b491906137d0565b905083600860008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541215612738576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161272f90613dbf565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146127a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161279d90613e51565b60405180910390fd5b83600860008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546127f5919061348f565b92505081905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33866040518363ffffffff1660e01b8152600401612859929190613620565b6020604051808303816000875af1158015612878573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061289c9190613675565b506128a685612990565b8473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb866040516129039190612cb9565b60405180910390a35050505050565b61291a612bf4565b73ffffffffffffffffffffffffffffffffffffffff16612938611c00565b73ffffffffffffffffffffffffffffffffffffffff161461298e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161298590613ebd565b60405180910390fd5b565b6000600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054148015612a1e57506000600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054145b15612aaa576040518060400160405280600381526020017f4e53430000000000000000000000000000000000000000000000000000000000815250600a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209081612aa491906140ae565b50612b2d565b6040518060400160405280600681526020017f4163746976650000000000000000000000000000000000000000000000000000815250600a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209081612b2b91906140ae565b505b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b6040518060c00160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081526020016000815260200160008152602001600081525090565b6040518060800160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001606081525090565b604051806040016040528060008152602001600081525090565b6000819050919050565b612cb381612ca0565b82525050565b6000602082019050612cce6000830184612caa565b92915050565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000612d0982612cde565b9050919050565b612d1981612cfe565b8114612d2457600080fd5b50565b600081359050612d3681612d10565b92915050565b600060208284031215612d5257612d51612cd4565b5b6000612d6084828501612d27565b91505092915050565b612d7281612ca0565b8114612d7d57600080fd5b50565b600081359050612d8f81612d69565b92915050565b60008060408385031215612dac57612dab612cd4565b5b6000612dba85828601612d27565b9250506020612dcb85828601612d80565b9150509250929050565b60008115159050919050565b612dea81612dd5565b82525050565b6000602082019050612e056000830184612de1565b92915050565b6000819050919050565b612e1e81612e0b565b82525050565b6000602082019050612e396000830184612e15565b92915050565b600060208284031215612e5557612e54612cd4565b5b6000612e6384828501612d80565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015612ea6578082015181840152602081019050612e8b565b60008484015250505050565b6000601f19601f8301169050919050565b6000612ece82612e6c565b612ed88185612e77565b9350612ee8818560208601612e88565b612ef181612eb2565b840191505092915050565b60006020820190508181036000830152612f168184612ec3565b905092915050565b6000604082019050612f336000830185612caa565b612f406020830184612caa565b9392505050565b600080fd5b600080fd5b600080fd5b60008083601f840112612f6c57612f6b612f47565b5b8235905067ffffffffffffffff811115612f8957612f88612f4c565b5b602083019150836020820283011115612fa557612fa4612f51565b5b9250929050565b60008083601f840112612fc257612fc1612f47565b5b8235905067ffffffffffffffff811115612fdf57612fde612f4c565b5b602083019150836020820283011115612ffb57612ffa612f51565b5b9250929050565b6000806000806040858703121561301c5761301b612cd4565b5b600085013567ffffffffffffffff81111561303a57613039612cd9565b5b61304687828801612f56565b9450945050602085013567ffffffffffffffff81111561306957613068612cd9565b5b61307587828801612fac565b925092505092959194509250565b60008060006060848603121561309c5761309b612cd4565b5b60006130aa86828701612d27565b93505060206130bb86828701612d27565b92505060406130cc86828701612d80565b9150509250925092565b6130df81612cfe565b82525050565b60006020820190506130fa60008301846130d6565b92915050565b61310981612cfe565b82525050565b61311881612ca0565b82525050565b60c0820160008201516131346000850182613100565b506020820151613147602085018261310f565b50604082015161315a604085018261310f565b50606082015161316d606085018261310f565b506080820151613180608085018261310f565b5060a082015161319360a085018261310f565b50505050565b600060c0820190506131ae600083018461311e565b92915050565b6131bd81612e0b565b82525050565b600082825260208201905092915050565b60006131df82612e6c565b6131e981856131c3565b93506131f9818560208601612e88565b61320281612eb2565b840191505092915050565b60006080830160008301516132256000860182613100565b50602083015161323860208601826131b4565b50604083015161324b604086018261310f565b506060830151848203606086015261326382826131d4565b9150508091505092915050565b6000602082019050818103600083015261328a818461320d565b905092915050565b6000819050919050565b60006132b76132b26132ad84612cde565b613292565b612cde565b9050919050565b60006132c98261329c565b9050919050565b60006132db826132be565b9050919050565b6132eb816132d0565b82525050565b600060208201905061330660008301846132e2565b92915050565b604082016000820151613322600085018261310f565b506020820151613335602085018261310f565b50505050565b6000604082019050613350600083018461330c565b92915050565b7f4f6e6c79207468652061646d696e2063616e2063616c6c20746869732066756e60008201527f6374696f6e2e0000000000000000000000000000000000000000000000000000602082015250565b60006133b2602683612e77565b91506133bd82613356565b604082019050919050565b600060208201905081810360008301526133e1816133a5565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061342282612ca0565b915061342d83612ca0565b9250828203905081811115613445576134446133e8565b5b92915050565b600061345682612e0b565b915061346183612e0b565b925082820190508281121560008312168382126000841215161715613489576134886133e8565b5b92915050565b600061349a82612e0b565b91506134a583612e0b565b92508282039050818112600084121682821360008512151617156134cc576134cb6133e8565b5b92915050565b60006134dd82612ca0565b91506134e883612ca0565b9250828201905080821115613500576134ff6133e8565b5b92915050565b7f536c617368656400000000000000000000000000000000000000000000000000600082015250565b600061353c600783612e77565b915061354782613506565b602082019050919050565b60006060820190506135676000830185612caa565b6135746020830184612caa565b81810360408301526135858161352f565b90509392505050565b7f576974686472617720736c61736866756e6420616d6f756e742065786365656460008201527f7320736c617368656446756e6473000000000000000000000000000000000000602082015250565b60006135ea602e83612e77565b91506135f58261358e565b604082019050919050565b60006020820190508181036000830152613619816135dd565b9050919050565b600060408201905061363560008301856130d6565b6136426020830184612caa565b9392505050565b61365281612dd5565b811461365d57600080fd5b50565b60008151905061366f81613649565b92915050565b60006020828403121561368b5761368a612cd4565b5b600061369984828501613660565b91505092915050565b600081519050919050565b600081905092915050565b60006136c3826136a2565b6136cd81856136ad565b93506136dd818560208601612e88565b80840191505092915050565b60006136f582846136b8565b915081905092915050565b7f4661696c656420746f2063616c6c206765744f776e65722066756e6374696f6e60008201527f206f662043504163636f756e7400000000000000000000000000000000000000602082015250565b600061375c602d83612e77565b915061376782613700565b604082019050919050565b6000602082019050818103600083015261378b8161374f565b9050919050565b600061379d82612cde565b9050919050565b6137ad81613792565b81146137b857600080fd5b50565b6000815190506137ca816137a4565b92915050565b6000602082840312156137e6576137e5612cd4565b5b60006137f4848285016137bb565b91505092915050565b7f4f6e6c792043502773206f776e65722063616e2072657175657374207769746860008201527f64726177616c0000000000000000000000000000000000000000000000000000602082015250565b6000613859602683612e77565b9150613864826137fd565b604082019050919050565b600060208201905081810360008301526138888161384c565b9050919050565b7f4e6f7420656e6f7567682066726f7a656e2062616c616e636520746f2077697460008201527f6864726177000000000000000000000000000000000000000000000000000000602082015250565b60006138eb602583612e77565b91506138f68261388f565b604082019050919050565b6000602082019050818103600083015261391a816138de565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061396857607f821691505b60208210810361397b5761397a613921565b5b50919050565b7f4e6f7420656e6f7567682062616c616e636520666f7220636f6c6c617465726160008201527f6c00000000000000000000000000000000000000000000000000000000000000602082015250565b60006139dd602183612e77565b91506139e882613981565b604082019050919050565b60006020820190508181036000830152613a0c816139d0565b9050919050565b6000606082019050613a2860008301866130d6565b613a3560208301856130d6565b613a426040830184612caa565b949350505050565b7f4f6e6c792043502773206f776e65722063616e2063616e63656c20776974686460008201527f7261772072657175657374000000000000000000000000000000000000000000602082015250565b6000613aa6602b83612e77565b9150613ab182613a4a565b604082019050919050565b60006020820190508181036000830152613ad581613a99565b9050919050565b7f4e6f2070656e64696e6720776974686472617720726571756573740000000000600082015250565b6000613b12601b83612e77565b9150613b1d82613adc565b602082019050919050565b60006020820190508181036000830152613b4181613b05565b9050919050565b7f4172726179206c656e67746873206d757374206d617463680000000000000000600082015250565b6000613b7e601883612e77565b9150613b8982613b48565b602082019050919050565b60006020820190508181036000830152613bad81613b71565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4f6e6c792043502773206f776e65722063616e20636f6e6669726d207769746860008201527f64726177616c0000000000000000000000000000000000000000000000000000602082015250565b6000613c3f602683612e77565b9150613c4a82613be3565b604082019050919050565b60006020820190508181036000830152613c6e81613c32565b9050919050565b7f57697468647261772064656c6179206e6f742070617373656400000000000000600082015250565b6000613cab601983612e77565b9150613cb682613c75565b602082019050919050565b60006020820190508181036000830152613cda81613c9e565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000613d3d602683612e77565b9150613d4882613ce1565b604082019050919050565b60006020820190508181036000830152613d6c81613d30565b9050919050565b7f576974686472617720616d6f756e7420657863656564732062616c616e636500600082015250565b6000613da9601f83612e77565b9150613db482613d73565b602082019050919050565b60006020820190508181036000830152613dd881613d9c565b9050919050565b7f4f6e6c792043502773206f776e65722063616e2077697468647261772074686560008201527f20636f6c6c61746572616c2066756e6473000000000000000000000000000000602082015250565b6000613e3b603183612e77565b9150613e4682613ddf565b604082019050919050565b60006020820190508181036000830152613e6a81613e2e565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000613ea7602083612e77565b9150613eb282613e71565b602082019050919050565b60006020820190508181036000830152613ed681613e9a565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302613f6e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82613f31565b613f788683613f31565b95508019841693508086168417925050509392505050565b6000613fab613fa6613fa184612ca0565b613292565b612ca0565b9050919050565b6000819050919050565b613fc583613f90565b613fd9613fd182613fb2565b848454613f3e565b825550505050565b600090565b613fee613fe1565b613ff9818484613fbc565b505050565b5b8181101561401d57614012600082613fe6565b600181019050613fff565b5050565b601f8211156140625761403381613f0c565b61403c84613f21565b8101602085101561404b578190505b61405f61405785613f21565b830182613ffe565b50505b505050565b600082821c905092915050565b600061408560001984600802614067565b1980831691505092915050565b600061409e8383614074565b9150826002028217905092915050565b6140b782612e6c565b67ffffffffffffffff8111156140d0576140cf613edd565b5b6140da8254613950565b6140e5828285614021565b600060209050601f8311600181146141185760008415614106578287015190505b6141108582614092565b865550614178565b601f19841661412686613f0c565b60005b8281101561414e57848901518255600182019150602085019450602081019050614129565b8683101561416b5784890151614167601f891682614074565b8355505b6001600288020188555050505b50505050505056fea264697066735822122069681e6f3a50fd12b1f4e038331eec4f5945f50e953d602bfab8262aab72608f64736f6c63430008190033",
}

// EcpCollateralABI is the input ABI used to generate the binding from.
// Deprecated: Use EcpCollateralMetaData.ABI instead.
var EcpCollateralABI = EcpCollateralMetaData.ABI

// EcpCollateralBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EcpCollateralMetaData.Bin instead.
var EcpCollateralBin = EcpCollateralMetaData.Bin

// DeployEcpCollateral deploys a new Ethereum contract, binding an instance of EcpCollateral to it.
func DeployEcpCollateral(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EcpCollateral, error) {
	parsed, err := EcpCollateralMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EcpCollateralBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EcpCollateral{EcpCollateralCaller: EcpCollateralCaller{contract: contract}, EcpCollateralTransactor: EcpCollateralTransactor{contract: contract}, EcpCollateralFilterer: EcpCollateralFilterer{contract: contract}}, nil
}

// EcpCollateral is an auto generated Go binding around an Ethereum contract.
type EcpCollateral struct {
	EcpCollateralCaller     // Read-only binding to the contract
	EcpCollateralTransactor // Write-only binding to the contract
	EcpCollateralFilterer   // Log filterer for contract events
}

// EcpCollateralCaller is an auto generated read-only Go binding around an Ethereum contract.
type EcpCollateralCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EcpCollateralTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EcpCollateralTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EcpCollateralFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EcpCollateralFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EcpCollateralSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EcpCollateralSession struct {
	Contract     *EcpCollateral    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EcpCollateralCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EcpCollateralCallerSession struct {
	Contract *EcpCollateralCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EcpCollateralTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EcpCollateralTransactorSession struct {
	Contract     *EcpCollateralTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EcpCollateralRaw is an auto generated low-level Go binding around an Ethereum contract.
type EcpCollateralRaw struct {
	Contract *EcpCollateral // Generic contract binding to access the raw methods on
}

// EcpCollateralCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EcpCollateralCallerRaw struct {
	Contract *EcpCollateralCaller // Generic read-only contract binding to access the raw methods on
}

// EcpCollateralTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EcpCollateralTransactorRaw struct {
	Contract *EcpCollateralTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEcpCollateral creates a new instance of EcpCollateral, bound to a specific deployed contract.
func NewEcpCollateral(address common.Address, backend bind.ContractBackend) (*EcpCollateral, error) {
	contract, err := bindEcpCollateral(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EcpCollateral{EcpCollateralCaller: EcpCollateralCaller{contract: contract}, EcpCollateralTransactor: EcpCollateralTransactor{contract: contract}, EcpCollateralFilterer: EcpCollateralFilterer{contract: contract}}, nil
}

// NewEcpCollateralCaller creates a new read-only instance of EcpCollateral, bound to a specific deployed contract.
func NewEcpCollateralCaller(address common.Address, caller bind.ContractCaller) (*EcpCollateralCaller, error) {
	contract, err := bindEcpCollateral(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EcpCollateralCaller{contract: contract}, nil
}

// NewEcpCollateralTransactor creates a new write-only instance of EcpCollateral, bound to a specific deployed contract.
func NewEcpCollateralTransactor(address common.Address, transactor bind.ContractTransactor) (*EcpCollateralTransactor, error) {
	contract, err := bindEcpCollateral(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EcpCollateralTransactor{contract: contract}, nil
}

// NewEcpCollateralFilterer creates a new log filterer instance of EcpCollateral, bound to a specific deployed contract.
func NewEcpCollateralFilterer(address common.Address, filterer bind.ContractFilterer) (*EcpCollateralFilterer, error) {
	contract, err := bindEcpCollateral(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EcpCollateralFilterer{contract: contract}, nil
}

// bindEcpCollateral binds a generic wrapper to an already deployed contract.
func bindEcpCollateral(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EcpCollateralMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EcpCollateral *EcpCollateralRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EcpCollateral.Contract.EcpCollateralCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EcpCollateral *EcpCollateralRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EcpCollateral.Contract.EcpCollateralTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EcpCollateral *EcpCollateralRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EcpCollateral.Contract.EcpCollateralTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EcpCollateral *EcpCollateralCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EcpCollateral.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EcpCollateral *EcpCollateralTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EcpCollateral.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EcpCollateral *EcpCollateralTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EcpCollateral.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(int256)
func (_EcpCollateral *EcpCollateralCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EcpCollateral.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(int256)
func (_EcpCollateral *EcpCollateralSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _EcpCollateral.Contract.Balances(&_EcpCollateral.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(int256)
func (_EcpCollateral *EcpCollateralCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _EcpCollateral.Contract.Balances(&_EcpCollateral.CallOpts, arg0)
}

// BaseCollateral is a free data retrieval call binding the contract method 0x9b5ddf09.
//
// Solidity: function baseCollateral() view returns(uint256)
func (_EcpCollateral *EcpCollateralCaller) BaseCollateral(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EcpCollateral.contract.Call(opts, &out, "baseCollateral")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseCollateral is a free data retrieval call binding the contract method 0x9b5ddf09.
//
// Solidity: function baseCollateral() view returns(uint256)
func (_EcpCollateral *EcpCollateralSession) BaseCollateral() (*big.Int, error) {
	return _EcpCollateral.Contract.BaseCollateral(&_EcpCollateral.CallOpts)
}

// BaseCollateral is a free data retrieval call binding the contract method 0x9b5ddf09.
//
// Solidity: function baseCollateral() view returns(uint256)
func (_EcpCollateral *EcpCollateralCallerSession) BaseCollateral() (*big.Int, error) {
	return _EcpCollateral.Contract.BaseCollateral(&_EcpCollateral.CallOpts)
}

// CollateralRatio is a free data retrieval call binding the contract method 0xb4eae1cb.
//
// Solidity: function collateralRatio() view returns(uint256)
func (_EcpCollateral *EcpCollateralCaller) CollateralRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EcpCollateral.contract.Call(opts, &out, "collateralRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollateralRatio is a free data retrieval call binding the contract method 0xb4eae1cb.
//
// Solidity: function collateralRatio() view returns(uint256)
func (_EcpCollateral *EcpCollateralSession) CollateralRatio() (*big.Int, error) {
	return _EcpCollateral.Contract.CollateralRatio(&_EcpCollateral.CallOpts)
}

// CollateralRatio is a free data retrieval call binding the contract method 0xb4eae1cb.
//
// Solidity: function collateralRatio() view returns(uint256)
func (_EcpCollateral *EcpCollateralCallerSession) CollateralRatio() (*big.Int, error) {
	return _EcpCollateral.Contract.CollateralRatio(&_EcpCollateral.CallOpts)
}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_EcpCollateral *EcpCollateralCaller) CollateralToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EcpCollateral.contract.Call(opts, &out, "collateralToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_EcpCollateral *EcpCollateralSession) CollateralToken() (common.Address, error) {
	return _EcpCollateral.Contract.CollateralToken(&_EcpCollateral.CallOpts)
}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_EcpCollateral *EcpCollateralCallerSession) CollateralToken() (common.Address, error) {
	return _EcpCollateral.Contract.CollateralToken(&_EcpCollateral.CallOpts)
}

// CpInfo is a free data retrieval call binding the contract method 0xa664c216.
//
// Solidity: function cpInfo(address cpAccount) view returns((address,int256,uint256,string))
func (_EcpCollateral *EcpCollateralCaller) CpInfo(opts *bind.CallOpts, cpAccount common.Address) (ECPCollateralCPInfo, error) {
	var out []interface{}
	err := _EcpCollateral.contract.Call(opts, &out, "cpInfo", cpAccount)

	if err != nil {
		return *new(ECPCollateralCPInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ECPCollateralCPInfo)).(*ECPCollateralCPInfo)

	return out0, err

}

// CpInfo is a free data retrieval call binding the contract method 0xa664c216.
//
// Solidity: function cpInfo(address cpAccount) view returns((address,int256,uint256,string))
func (_EcpCollateral *EcpCollateralSession) CpInfo(cpAccount common.Address) (ECPCollateralCPInfo, error) {
	return _EcpCollateral.Contract.CpInfo(&_EcpCollateral.CallOpts, cpAccount)
}

// CpInfo is a free data retrieval call binding the contract method 0xa664c216.
//
// Solidity: function cpInfo(address cpAccount) view returns((address,int256,uint256,string))
func (_EcpCollateral *EcpCollateralCallerSession) CpInfo(cpAccount common.Address) (ECPCollateralCPInfo, error) {
	return _EcpCollateral.Contract.CpInfo(&_EcpCollateral.CallOpts, cpAccount)
}

// CpStatus is a free data retrieval call binding the contract method 0x3fe65177.
//
// Solidity: function cpStatus(address ) view returns(string)
func (_EcpCollateral *EcpCollateralCaller) CpStatus(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _EcpCollateral.contract.Call(opts, &out, "cpStatus", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CpStatus is a free data retrieval call binding the contract method 0x3fe65177.
//
// Solidity: function cpStatus(address ) view returns(string)
func (_EcpCollateral *EcpCollateralSession) CpStatus(arg0 common.Address) (string, error) {
	return _EcpCollateral.Contract.CpStatus(&_EcpCollateral.CallOpts, arg0)
}

// CpStatus is a free data retrieval call binding the contract method 0x3fe65177.
//
// Solidity: function cpStatus(address ) view returns(string)
func (_EcpCollateral *EcpCollateralCallerSession) CpStatus(arg0 common.Address) (string, error) {
	return _EcpCollateral.Contract.CpStatus(&_EcpCollateral.CallOpts, arg0)
}

// FrozenBalance is a free data retrieval call binding the contract method 0x266565a9.
//
// Solidity: function frozenBalance(address ) view returns(uint256)
func (_EcpCollateral *EcpCollateralCaller) FrozenBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EcpCollateral.contract.Call(opts, &out, "frozenBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FrozenBalance is a free data retrieval call binding the contract method 0x266565a9.
//
// Solidity: function frozenBalance(address ) view returns(uint256)
func (_EcpCollateral *EcpCollateralSession) FrozenBalance(arg0 common.Address) (*big.Int, error) {
	return _EcpCollateral.Contract.FrozenBalance(&_EcpCollateral.CallOpts, arg0)
}

// FrozenBalance is a free data retrieval call binding the contract method 0x266565a9.
//
// Solidity: function frozenBalance(address ) view returns(uint256)
func (_EcpCollateral *EcpCollateralCallerSession) FrozenBalance(arg0 common.Address) (*big.Int, error) {
	return _EcpCollateral.Contract.FrozenBalance(&_EcpCollateral.CallOpts, arg0)
}

// GetBaseCollateral is a free data retrieval call binding the contract method 0x70b72944.
//
// Solidity: function getBaseCollateral() view returns(uint256)
func (_EcpCollateral *EcpCollateralCaller) GetBaseCollateral(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EcpCollateral.contract.Call(opts, &out, "getBaseCollateral")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBaseCollateral is a free data retrieval call binding the contract method 0x70b72944.
//
// Solidity: function getBaseCollateral() view returns(uint256)
func (_EcpCollateral *EcpCollateralSession) GetBaseCollateral() (*big.Int, error) {
	return _EcpCollateral.Contract.GetBaseCollateral(&_EcpCollateral.CallOpts)
}

// GetBaseCollateral is a free data retrieval call binding the contract method 0x70b72944.
//
// Solidity: function getBaseCollateral() view returns(uint256)
func (_EcpCollateral *EcpCollateralCallerSession) GetBaseCollateral() (*big.Int, error) {
	return _EcpCollateral.Contract.GetBaseCollateral(&_EcpCollateral.CallOpts)
}

// GetECPCollateralInfo is a free data retrieval call binding the contract method 0x9939cd18.
//
// Solidity: function getECPCollateralInfo() view returns((address,uint256,uint256,uint256,uint256,uint256))
func (_EcpCollateral *EcpCollateralCaller) GetECPCollateralInfo(opts *bind.CallOpts) (ECPCollateralContractInfo, error) {
	var out []interface{}
	err := _EcpCollateral.contract.Call(opts, &out, "getECPCollateralInfo")

	if err != nil {
		return *new(ECPCollateralContractInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ECPCollateralContractInfo)).(*ECPCollateralContractInfo)

	return out0, err

}

// GetECPCollateralInfo is a free data retrieval call binding the contract method 0x9939cd18.
//
// Solidity: function getECPCollateralInfo() view returns((address,uint256,uint256,uint256,uint256,uint256))
func (_EcpCollateral *EcpCollateralSession) GetECPCollateralInfo() (ECPCollateralContractInfo, error) {
	return _EcpCollateral.Contract.GetECPCollateralInfo(&_EcpCollateral.CallOpts)
}

// GetECPCollateralInfo is a free data retrieval call binding the contract method 0x9939cd18.
//
// Solidity: function getECPCollateralInfo() view returns((address,uint256,uint256,uint256,uint256,uint256))
func (_EcpCollateral *EcpCollateralCallerSession) GetECPCollateralInfo() (ECPCollateralContractInfo, error) {
	return _EcpCollateral.Contract.GetECPCollateralInfo(&_EcpCollateral.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_EcpCollateral *EcpCollateralCaller) IsAdmin(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _EcpCollateral.contract.Call(opts, &out, "isAdmin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_EcpCollateral *EcpCollateralSession) IsAdmin(arg0 common.Address) (bool, error) {
	return _EcpCollateral.Contract.IsAdmin(&_EcpCollateral.CallOpts, arg0)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_EcpCollateral *EcpCollateralCallerSession) IsAdmin(arg0 common.Address) (bool, error) {
	return _EcpCollateral.Contract.IsAdmin(&_EcpCollateral.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EcpCollateral *EcpCollateralCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EcpCollateral.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EcpCollateral *EcpCollateralSession) Owner() (common.Address, error) {
	return _EcpCollateral.Contract.Owner(&_EcpCollateral.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EcpCollateral *EcpCollateralCallerSession) Owner() (common.Address, error) {
	return _EcpCollateral.Contract.Owner(&_EcpCollateral.CallOpts)
}

// SlashRatio is a free data retrieval call binding the contract method 0xd27ca89b.
//
// Solidity: function slashRatio() view returns(uint256)
func (_EcpCollateral *EcpCollateralCaller) SlashRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EcpCollateral.contract.Call(opts, &out, "slashRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashRatio is a free data retrieval call binding the contract method 0xd27ca89b.
//
// Solidity: function slashRatio() view returns(uint256)
func (_EcpCollateral *EcpCollateralSession) SlashRatio() (*big.Int, error) {
	return _EcpCollateral.Contract.SlashRatio(&_EcpCollateral.CallOpts)
}

// SlashRatio is a free data retrieval call binding the contract method 0xd27ca89b.
//
// Solidity: function slashRatio() view returns(uint256)
func (_EcpCollateral *EcpCollateralCallerSession) SlashRatio() (*big.Int, error) {
	return _EcpCollateral.Contract.SlashRatio(&_EcpCollateral.CallOpts)
}

// SlashedFunds is a free data retrieval call binding the contract method 0x6f99f15c.
//
// Solidity: function slashedFunds() view returns(uint256)
func (_EcpCollateral *EcpCollateralCaller) SlashedFunds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EcpCollateral.contract.Call(opts, &out, "slashedFunds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashedFunds is a free data retrieval call binding the contract method 0x6f99f15c.
//
// Solidity: function slashedFunds() view returns(uint256)
func (_EcpCollateral *EcpCollateralSession) SlashedFunds() (*big.Int, error) {
	return _EcpCollateral.Contract.SlashedFunds(&_EcpCollateral.CallOpts)
}

// SlashedFunds is a free data retrieval call binding the contract method 0x6f99f15c.
//
// Solidity: function slashedFunds() view returns(uint256)
func (_EcpCollateral *EcpCollateralCallerSession) SlashedFunds() (*big.Int, error) {
	return _EcpCollateral.Contract.SlashedFunds(&_EcpCollateral.CallOpts)
}

// ViewWithdrawRequest is a free data retrieval call binding the contract method 0xbede6e31.
//
// Solidity: function viewWithdrawRequest(address cpAccount) view returns((uint256,uint256))
func (_EcpCollateral *EcpCollateralCaller) ViewWithdrawRequest(opts *bind.CallOpts, cpAccount common.Address) (ECPCollateralWithdrawRequest, error) {
	var out []interface{}
	err := _EcpCollateral.contract.Call(opts, &out, "viewWithdrawRequest", cpAccount)

	if err != nil {
		return *new(ECPCollateralWithdrawRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(ECPCollateralWithdrawRequest)).(*ECPCollateralWithdrawRequest)

	return out0, err

}

// ViewWithdrawRequest is a free data retrieval call binding the contract method 0xbede6e31.
//
// Solidity: function viewWithdrawRequest(address cpAccount) view returns((uint256,uint256))
func (_EcpCollateral *EcpCollateralSession) ViewWithdrawRequest(cpAccount common.Address) (ECPCollateralWithdrawRequest, error) {
	return _EcpCollateral.Contract.ViewWithdrawRequest(&_EcpCollateral.CallOpts, cpAccount)
}

// ViewWithdrawRequest is a free data retrieval call binding the contract method 0xbede6e31.
//
// Solidity: function viewWithdrawRequest(address cpAccount) view returns((uint256,uint256))
func (_EcpCollateral *EcpCollateralCallerSession) ViewWithdrawRequest(cpAccount common.Address) (ECPCollateralWithdrawRequest, error) {
	return _EcpCollateral.Contract.ViewWithdrawRequest(&_EcpCollateral.CallOpts, cpAccount)
}

// WithdrawDelay is a free data retrieval call binding the contract method 0x0288a39c.
//
// Solidity: function withdrawDelay() view returns(uint256)
func (_EcpCollateral *EcpCollateralCaller) WithdrawDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EcpCollateral.contract.Call(opts, &out, "withdrawDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawDelay is a free data retrieval call binding the contract method 0x0288a39c.
//
// Solidity: function withdrawDelay() view returns(uint256)
func (_EcpCollateral *EcpCollateralSession) WithdrawDelay() (*big.Int, error) {
	return _EcpCollateral.Contract.WithdrawDelay(&_EcpCollateral.CallOpts)
}

// WithdrawDelay is a free data retrieval call binding the contract method 0x0288a39c.
//
// Solidity: function withdrawDelay() view returns(uint256)
func (_EcpCollateral *EcpCollateralCallerSession) WithdrawDelay() (*big.Int, error) {
	return _EcpCollateral.Contract.WithdrawDelay(&_EcpCollateral.CallOpts)
}

// WithdrawRequests is a free data retrieval call binding the contract method 0x52df49ec.
//
// Solidity: function withdrawRequests(address ) view returns(uint256 amount, uint256 requestBlock)
func (_EcpCollateral *EcpCollateralCaller) WithdrawRequests(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount       *big.Int
	RequestBlock *big.Int
}, error) {
	var out []interface{}
	err := _EcpCollateral.contract.Call(opts, &out, "withdrawRequests", arg0)

	outstruct := new(struct {
		Amount       *big.Int
		RequestBlock *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RequestBlock = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// WithdrawRequests is a free data retrieval call binding the contract method 0x52df49ec.
//
// Solidity: function withdrawRequests(address ) view returns(uint256 amount, uint256 requestBlock)
func (_EcpCollateral *EcpCollateralSession) WithdrawRequests(arg0 common.Address) (struct {
	Amount       *big.Int
	RequestBlock *big.Int
}, error) {
	return _EcpCollateral.Contract.WithdrawRequests(&_EcpCollateral.CallOpts, arg0)
}

// WithdrawRequests is a free data retrieval call binding the contract method 0x52df49ec.
//
// Solidity: function withdrawRequests(address ) view returns(uint256 amount, uint256 requestBlock)
func (_EcpCollateral *EcpCollateralCallerSession) WithdrawRequests(arg0 common.Address) (struct {
	Amount       *big.Int
	RequestBlock *big.Int
}, error) {
	return _EcpCollateral.Contract.WithdrawRequests(&_EcpCollateral.CallOpts, arg0)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address newAdmin) returns()
func (_EcpCollateral *EcpCollateralTransactor) AddAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _EcpCollateral.contract.Transact(opts, "addAdmin", newAdmin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address newAdmin) returns()
func (_EcpCollateral *EcpCollateralSession) AddAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EcpCollateral.Contract.AddAdmin(&_EcpCollateral.TransactOpts, newAdmin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address newAdmin) returns()
func (_EcpCollateral *EcpCollateralTransactorSession) AddAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EcpCollateral.Contract.AddAdmin(&_EcpCollateral.TransactOpts, newAdmin)
}

// BatchLock is a paid mutator transaction binding the contract method 0x63215bb7.
//
// Solidity: function batchLock(address[] cps, uint256[] taskCollaterals) returns()
func (_EcpCollateral *EcpCollateralTransactor) BatchLock(opts *bind.TransactOpts, cps []common.Address, taskCollaterals []*big.Int) (*types.Transaction, error) {
	return _EcpCollateral.contract.Transact(opts, "batchLock", cps, taskCollaterals)
}

// BatchLock is a paid mutator transaction binding the contract method 0x63215bb7.
//
// Solidity: function batchLock(address[] cps, uint256[] taskCollaterals) returns()
func (_EcpCollateral *EcpCollateralSession) BatchLock(cps []common.Address, taskCollaterals []*big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.BatchLock(&_EcpCollateral.TransactOpts, cps, taskCollaterals)
}

// BatchLock is a paid mutator transaction binding the contract method 0x63215bb7.
//
// Solidity: function batchLock(address[] cps, uint256[] taskCollaterals) returns()
func (_EcpCollateral *EcpCollateralTransactorSession) BatchLock(cps []common.Address, taskCollaterals []*big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.BatchLock(&_EcpCollateral.TransactOpts, cps, taskCollaterals)
}

// BatchSlash is a paid mutator transaction binding the contract method 0xc6ff4555.
//
// Solidity: function batchSlash(address[] cps, uint256[] slashAmounts) returns()
func (_EcpCollateral *EcpCollateralTransactor) BatchSlash(opts *bind.TransactOpts, cps []common.Address, slashAmounts []*big.Int) (*types.Transaction, error) {
	return _EcpCollateral.contract.Transact(opts, "batchSlash", cps, slashAmounts)
}

// BatchSlash is a paid mutator transaction binding the contract method 0xc6ff4555.
//
// Solidity: function batchSlash(address[] cps, uint256[] slashAmounts) returns()
func (_EcpCollateral *EcpCollateralSession) BatchSlash(cps []common.Address, slashAmounts []*big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.BatchSlash(&_EcpCollateral.TransactOpts, cps, slashAmounts)
}

// BatchSlash is a paid mutator transaction binding the contract method 0xc6ff4555.
//
// Solidity: function batchSlash(address[] cps, uint256[] slashAmounts) returns()
func (_EcpCollateral *EcpCollateralTransactorSession) BatchSlash(cps []common.Address, slashAmounts []*big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.BatchSlash(&_EcpCollateral.TransactOpts, cps, slashAmounts)
}

// BatchUnlock is a paid mutator transaction binding the contract method 0x5f7d0e84.
//
// Solidity: function batchUnlock(address[] cps, uint256[] taskCollaterals) returns()
func (_EcpCollateral *EcpCollateralTransactor) BatchUnlock(opts *bind.TransactOpts, cps []common.Address, taskCollaterals []*big.Int) (*types.Transaction, error) {
	return _EcpCollateral.contract.Transact(opts, "batchUnlock", cps, taskCollaterals)
}

// BatchUnlock is a paid mutator transaction binding the contract method 0x5f7d0e84.
//
// Solidity: function batchUnlock(address[] cps, uint256[] taskCollaterals) returns()
func (_EcpCollateral *EcpCollateralSession) BatchUnlock(cps []common.Address, taskCollaterals []*big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.BatchUnlock(&_EcpCollateral.TransactOpts, cps, taskCollaterals)
}

// BatchUnlock is a paid mutator transaction binding the contract method 0x5f7d0e84.
//
// Solidity: function batchUnlock(address[] cps, uint256[] taskCollaterals) returns()
func (_EcpCollateral *EcpCollateralTransactorSession) BatchUnlock(cps []common.Address, taskCollaterals []*big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.BatchUnlock(&_EcpCollateral.TransactOpts, cps, taskCollaterals)
}

// CancelWithdrawRequest is a paid mutator transaction binding the contract method 0x5d2cd2a7.
//
// Solidity: function cancelWithdrawRequest(address cpAccount) returns()
func (_EcpCollateral *EcpCollateralTransactor) CancelWithdrawRequest(opts *bind.TransactOpts, cpAccount common.Address) (*types.Transaction, error) {
	return _EcpCollateral.contract.Transact(opts, "cancelWithdrawRequest", cpAccount)
}

// CancelWithdrawRequest is a paid mutator transaction binding the contract method 0x5d2cd2a7.
//
// Solidity: function cancelWithdrawRequest(address cpAccount) returns()
func (_EcpCollateral *EcpCollateralSession) CancelWithdrawRequest(cpAccount common.Address) (*types.Transaction, error) {
	return _EcpCollateral.Contract.CancelWithdrawRequest(&_EcpCollateral.TransactOpts, cpAccount)
}

// CancelWithdrawRequest is a paid mutator transaction binding the contract method 0x5d2cd2a7.
//
// Solidity: function cancelWithdrawRequest(address cpAccount) returns()
func (_EcpCollateral *EcpCollateralTransactorSession) CancelWithdrawRequest(cpAccount common.Address) (*types.Transaction, error) {
	return _EcpCollateral.Contract.CancelWithdrawRequest(&_EcpCollateral.TransactOpts, cpAccount)
}

// ConfirmWithdraw is a paid mutator transaction binding the contract method 0xd2bfc1c7.
//
// Solidity: function confirmWithdraw(address cpAccount) returns()
func (_EcpCollateral *EcpCollateralTransactor) ConfirmWithdraw(opts *bind.TransactOpts, cpAccount common.Address) (*types.Transaction, error) {
	return _EcpCollateral.contract.Transact(opts, "confirmWithdraw", cpAccount)
}

// ConfirmWithdraw is a paid mutator transaction binding the contract method 0xd2bfc1c7.
//
// Solidity: function confirmWithdraw(address cpAccount) returns()
func (_EcpCollateral *EcpCollateralSession) ConfirmWithdraw(cpAccount common.Address) (*types.Transaction, error) {
	return _EcpCollateral.Contract.ConfirmWithdraw(&_EcpCollateral.TransactOpts, cpAccount)
}

// ConfirmWithdraw is a paid mutator transaction binding the contract method 0xd2bfc1c7.
//
// Solidity: function confirmWithdraw(address cpAccount) returns()
func (_EcpCollateral *EcpCollateralTransactorSession) ConfirmWithdraw(cpAccount common.Address) (*types.Transaction, error) {
	return _EcpCollateral.Contract.ConfirmWithdraw(&_EcpCollateral.TransactOpts, cpAccount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address cpAccount, uint256 amount) returns()
func (_EcpCollateral *EcpCollateralTransactor) Deposit(opts *bind.TransactOpts, cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.contract.Transact(opts, "deposit", cpAccount, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address cpAccount, uint256 amount) returns()
func (_EcpCollateral *EcpCollateralSession) Deposit(cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.Deposit(&_EcpCollateral.TransactOpts, cpAccount, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address cpAccount, uint256 amount) returns()
func (_EcpCollateral *EcpCollateralTransactorSession) Deposit(cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.Deposit(&_EcpCollateral.TransactOpts, cpAccount, amount)
}

// DisputeProof is a paid mutator transaction binding the contract method 0x8331f8e5.
//
// Solidity: function disputeProof(address taskContractAddress, address cpAccount, uint256 taskID) returns()
func (_EcpCollateral *EcpCollateralTransactor) DisputeProof(opts *bind.TransactOpts, taskContractAddress common.Address, cpAccount common.Address, taskID *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.contract.Transact(opts, "disputeProof", taskContractAddress, cpAccount, taskID)
}

// DisputeProof is a paid mutator transaction binding the contract method 0x8331f8e5.
//
// Solidity: function disputeProof(address taskContractAddress, address cpAccount, uint256 taskID) returns()
func (_EcpCollateral *EcpCollateralSession) DisputeProof(taskContractAddress common.Address, cpAccount common.Address, taskID *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.DisputeProof(&_EcpCollateral.TransactOpts, taskContractAddress, cpAccount, taskID)
}

// DisputeProof is a paid mutator transaction binding the contract method 0x8331f8e5.
//
// Solidity: function disputeProof(address taskContractAddress, address cpAccount, uint256 taskID) returns()
func (_EcpCollateral *EcpCollateralTransactorSession) DisputeProof(taskContractAddress common.Address, cpAccount common.Address, taskID *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.DisputeProof(&_EcpCollateral.TransactOpts, taskContractAddress, cpAccount, taskID)
}

// LockCollateral is a paid mutator transaction binding the contract method 0x47a7d107.
//
// Solidity: function lockCollateral(address cp, uint256 taskCollateral) returns()
func (_EcpCollateral *EcpCollateralTransactor) LockCollateral(opts *bind.TransactOpts, cp common.Address, taskCollateral *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.contract.Transact(opts, "lockCollateral", cp, taskCollateral)
}

// LockCollateral is a paid mutator transaction binding the contract method 0x47a7d107.
//
// Solidity: function lockCollateral(address cp, uint256 taskCollateral) returns()
func (_EcpCollateral *EcpCollateralSession) LockCollateral(cp common.Address, taskCollateral *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.LockCollateral(&_EcpCollateral.TransactOpts, cp, taskCollateral)
}

// LockCollateral is a paid mutator transaction binding the contract method 0x47a7d107.
//
// Solidity: function lockCollateral(address cp, uint256 taskCollateral) returns()
func (_EcpCollateral *EcpCollateralTransactorSession) LockCollateral(cp common.Address, taskCollateral *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.LockCollateral(&_EcpCollateral.TransactOpts, cp, taskCollateral)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_EcpCollateral *EcpCollateralTransactor) RemoveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _EcpCollateral.contract.Transact(opts, "removeAdmin", admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_EcpCollateral *EcpCollateralSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _EcpCollateral.Contract.RemoveAdmin(&_EcpCollateral.TransactOpts, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_EcpCollateral *EcpCollateralTransactorSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _EcpCollateral.Contract.RemoveAdmin(&_EcpCollateral.TransactOpts, admin)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EcpCollateral *EcpCollateralTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EcpCollateral.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EcpCollateral *EcpCollateralSession) RenounceOwnership() (*types.Transaction, error) {
	return _EcpCollateral.Contract.RenounceOwnership(&_EcpCollateral.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EcpCollateral *EcpCollateralTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EcpCollateral.Contract.RenounceOwnership(&_EcpCollateral.TransactOpts)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x397a1b28.
//
// Solidity: function requestWithdraw(address cpAccount, uint256 amount) returns()
func (_EcpCollateral *EcpCollateralTransactor) RequestWithdraw(opts *bind.TransactOpts, cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.contract.Transact(opts, "requestWithdraw", cpAccount, amount)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x397a1b28.
//
// Solidity: function requestWithdraw(address cpAccount, uint256 amount) returns()
func (_EcpCollateral *EcpCollateralSession) RequestWithdraw(cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.RequestWithdraw(&_EcpCollateral.TransactOpts, cpAccount, amount)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x397a1b28.
//
// Solidity: function requestWithdraw(address cpAccount, uint256 amount) returns()
func (_EcpCollateral *EcpCollateralTransactorSession) RequestWithdraw(cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.RequestWithdraw(&_EcpCollateral.TransactOpts, cpAccount, amount)
}

// SetBaseCollateral is a paid mutator transaction binding the contract method 0x7f58a7e5.
//
// Solidity: function setBaseCollateral(uint256 _baseCollateral) returns()
func (_EcpCollateral *EcpCollateralTransactor) SetBaseCollateral(opts *bind.TransactOpts, _baseCollateral *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.contract.Transact(opts, "setBaseCollateral", _baseCollateral)
}

// SetBaseCollateral is a paid mutator transaction binding the contract method 0x7f58a7e5.
//
// Solidity: function setBaseCollateral(uint256 _baseCollateral) returns()
func (_EcpCollateral *EcpCollateralSession) SetBaseCollateral(_baseCollateral *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.SetBaseCollateral(&_EcpCollateral.TransactOpts, _baseCollateral)
}

// SetBaseCollateral is a paid mutator transaction binding the contract method 0x7f58a7e5.
//
// Solidity: function setBaseCollateral(uint256 _baseCollateral) returns()
func (_EcpCollateral *EcpCollateralTransactorSession) SetBaseCollateral(_baseCollateral *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.SetBaseCollateral(&_EcpCollateral.TransactOpts, _baseCollateral)
}

// SetCollateralRatio is a paid mutator transaction binding the contract method 0x6060663e.
//
// Solidity: function setCollateralRatio(uint256 _collateralRatio) returns()
func (_EcpCollateral *EcpCollateralTransactor) SetCollateralRatio(opts *bind.TransactOpts, _collateralRatio *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.contract.Transact(opts, "setCollateralRatio", _collateralRatio)
}

// SetCollateralRatio is a paid mutator transaction binding the contract method 0x6060663e.
//
// Solidity: function setCollateralRatio(uint256 _collateralRatio) returns()
func (_EcpCollateral *EcpCollateralSession) SetCollateralRatio(_collateralRatio *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.SetCollateralRatio(&_EcpCollateral.TransactOpts, _collateralRatio)
}

// SetCollateralRatio is a paid mutator transaction binding the contract method 0x6060663e.
//
// Solidity: function setCollateralRatio(uint256 _collateralRatio) returns()
func (_EcpCollateral *EcpCollateralTransactorSession) SetCollateralRatio(_collateralRatio *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.SetCollateralRatio(&_EcpCollateral.TransactOpts, _collateralRatio)
}

// SetCollateralToken is a paid mutator transaction binding the contract method 0x666181a9.
//
// Solidity: function setCollateralToken(address tokenAddress) returns()
func (_EcpCollateral *EcpCollateralTransactor) SetCollateralToken(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _EcpCollateral.contract.Transact(opts, "setCollateralToken", tokenAddress)
}

// SetCollateralToken is a paid mutator transaction binding the contract method 0x666181a9.
//
// Solidity: function setCollateralToken(address tokenAddress) returns()
func (_EcpCollateral *EcpCollateralSession) SetCollateralToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _EcpCollateral.Contract.SetCollateralToken(&_EcpCollateral.TransactOpts, tokenAddress)
}

// SetCollateralToken is a paid mutator transaction binding the contract method 0x666181a9.
//
// Solidity: function setCollateralToken(address tokenAddress) returns()
func (_EcpCollateral *EcpCollateralTransactorSession) SetCollateralToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _EcpCollateral.Contract.SetCollateralToken(&_EcpCollateral.TransactOpts, tokenAddress)
}

// SetSlashRatio is a paid mutator transaction binding the contract method 0xce3518aa.
//
// Solidity: function setSlashRatio(uint256 _slashRatio) returns()
func (_EcpCollateral *EcpCollateralTransactor) SetSlashRatio(opts *bind.TransactOpts, _slashRatio *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.contract.Transact(opts, "setSlashRatio", _slashRatio)
}

// SetSlashRatio is a paid mutator transaction binding the contract method 0xce3518aa.
//
// Solidity: function setSlashRatio(uint256 _slashRatio) returns()
func (_EcpCollateral *EcpCollateralSession) SetSlashRatio(_slashRatio *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.SetSlashRatio(&_EcpCollateral.TransactOpts, _slashRatio)
}

// SetSlashRatio is a paid mutator transaction binding the contract method 0xce3518aa.
//
// Solidity: function setSlashRatio(uint256 _slashRatio) returns()
func (_EcpCollateral *EcpCollateralTransactorSession) SetSlashRatio(_slashRatio *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.SetSlashRatio(&_EcpCollateral.TransactOpts, _slashRatio)
}

// SetWithdrawDelay is a paid mutator transaction binding the contract method 0x72f0cb30.
//
// Solidity: function setWithdrawDelay(uint256 _withdrawDelay) returns()
func (_EcpCollateral *EcpCollateralTransactor) SetWithdrawDelay(opts *bind.TransactOpts, _withdrawDelay *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.contract.Transact(opts, "setWithdrawDelay", _withdrawDelay)
}

// SetWithdrawDelay is a paid mutator transaction binding the contract method 0x72f0cb30.
//
// Solidity: function setWithdrawDelay(uint256 _withdrawDelay) returns()
func (_EcpCollateral *EcpCollateralSession) SetWithdrawDelay(_withdrawDelay *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.SetWithdrawDelay(&_EcpCollateral.TransactOpts, _withdrawDelay)
}

// SetWithdrawDelay is a paid mutator transaction binding the contract method 0x72f0cb30.
//
// Solidity: function setWithdrawDelay(uint256 _withdrawDelay) returns()
func (_EcpCollateral *EcpCollateralTransactorSession) SetWithdrawDelay(_withdrawDelay *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.SetWithdrawDelay(&_EcpCollateral.TransactOpts, _withdrawDelay)
}

// SlashCollateral is a paid mutator transaction binding the contract method 0x1d47a62d.
//
// Solidity: function slashCollateral(address cp, uint256 slashAmount) returns()
func (_EcpCollateral *EcpCollateralTransactor) SlashCollateral(opts *bind.TransactOpts, cp common.Address, slashAmount *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.contract.Transact(opts, "slashCollateral", cp, slashAmount)
}

// SlashCollateral is a paid mutator transaction binding the contract method 0x1d47a62d.
//
// Solidity: function slashCollateral(address cp, uint256 slashAmount) returns()
func (_EcpCollateral *EcpCollateralSession) SlashCollateral(cp common.Address, slashAmount *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.SlashCollateral(&_EcpCollateral.TransactOpts, cp, slashAmount)
}

// SlashCollateral is a paid mutator transaction binding the contract method 0x1d47a62d.
//
// Solidity: function slashCollateral(address cp, uint256 slashAmount) returns()
func (_EcpCollateral *EcpCollateralTransactorSession) SlashCollateral(cp common.Address, slashAmount *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.SlashCollateral(&_EcpCollateral.TransactOpts, cp, slashAmount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EcpCollateral *EcpCollateralTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EcpCollateral.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EcpCollateral *EcpCollateralSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EcpCollateral.Contract.TransferOwnership(&_EcpCollateral.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EcpCollateral *EcpCollateralTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EcpCollateral.Contract.TransferOwnership(&_EcpCollateral.TransactOpts, newOwner)
}

// UnlockCollateral is a paid mutator transaction binding the contract method 0x17f17498.
//
// Solidity: function unlockCollateral(address cp, uint256 taskCollateral) returns()
func (_EcpCollateral *EcpCollateralTransactor) UnlockCollateral(opts *bind.TransactOpts, cp common.Address, taskCollateral *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.contract.Transact(opts, "unlockCollateral", cp, taskCollateral)
}

// UnlockCollateral is a paid mutator transaction binding the contract method 0x17f17498.
//
// Solidity: function unlockCollateral(address cp, uint256 taskCollateral) returns()
func (_EcpCollateral *EcpCollateralSession) UnlockCollateral(cp common.Address, taskCollateral *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.UnlockCollateral(&_EcpCollateral.TransactOpts, cp, taskCollateral)
}

// UnlockCollateral is a paid mutator transaction binding the contract method 0x17f17498.
//
// Solidity: function unlockCollateral(address cp, uint256 taskCollateral) returns()
func (_EcpCollateral *EcpCollateralTransactorSession) UnlockCollateral(cp common.Address, taskCollateral *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.UnlockCollateral(&_EcpCollateral.TransactOpts, cp, taskCollateral)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address cpAccount, uint256 amount) returns()
func (_EcpCollateral *EcpCollateralTransactor) Withdraw(opts *bind.TransactOpts, cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.contract.Transact(opts, "withdraw", cpAccount, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address cpAccount, uint256 amount) returns()
func (_EcpCollateral *EcpCollateralSession) Withdraw(cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.Withdraw(&_EcpCollateral.TransactOpts, cpAccount, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address cpAccount, uint256 amount) returns()
func (_EcpCollateral *EcpCollateralTransactorSession) Withdraw(cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.Withdraw(&_EcpCollateral.TransactOpts, cpAccount, amount)
}

// WithdrawSlashedFunds is a paid mutator transaction binding the contract method 0x2894493f.
//
// Solidity: function withdrawSlashedFunds(uint256 slashfund) returns()
func (_EcpCollateral *EcpCollateralTransactor) WithdrawSlashedFunds(opts *bind.TransactOpts, slashfund *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.contract.Transact(opts, "withdrawSlashedFunds", slashfund)
}

// WithdrawSlashedFunds is a paid mutator transaction binding the contract method 0x2894493f.
//
// Solidity: function withdrawSlashedFunds(uint256 slashfund) returns()
func (_EcpCollateral *EcpCollateralSession) WithdrawSlashedFunds(slashfund *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.WithdrawSlashedFunds(&_EcpCollateral.TransactOpts, slashfund)
}

// WithdrawSlashedFunds is a paid mutator transaction binding the contract method 0x2894493f.
//
// Solidity: function withdrawSlashedFunds(uint256 slashfund) returns()
func (_EcpCollateral *EcpCollateralTransactorSession) WithdrawSlashedFunds(slashfund *big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.WithdrawSlashedFunds(&_EcpCollateral.TransactOpts, slashfund)
}

// EcpCollateralCollateralAdjustedIterator is returned from FilterCollateralAdjusted and is used to iterate over the raw logs and unpacked data for CollateralAdjusted events raised by the EcpCollateral contract.
type EcpCollateralCollateralAdjustedIterator struct {
	Event *EcpCollateralCollateralAdjusted // Event containing the contract specifics and raw log

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
func (it *EcpCollateralCollateralAdjustedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpCollateralCollateralAdjusted)
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
		it.Event = new(EcpCollateralCollateralAdjusted)
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
func (it *EcpCollateralCollateralAdjustedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpCollateralCollateralAdjustedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpCollateralCollateralAdjusted represents a CollateralAdjusted event raised by the EcpCollateral contract.
type EcpCollateralCollateralAdjusted struct {
	Cp            common.Address
	FrozenAmount  *big.Int
	BalanceAmount *big.Int
	Operation     string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCollateralAdjusted is a free log retrieval operation binding the contract event 0x42f1a8a1aee108e84e4eabfaa8d51d7baaa1a02d482295297883a44b2debd3e4.
//
// Solidity: event CollateralAdjusted(address indexed cp, uint256 frozenAmount, uint256 balanceAmount, string operation)
func (_EcpCollateral *EcpCollateralFilterer) FilterCollateralAdjusted(opts *bind.FilterOpts, cp []common.Address) (*EcpCollateralCollateralAdjustedIterator, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _EcpCollateral.contract.FilterLogs(opts, "CollateralAdjusted", cpRule)
	if err != nil {
		return nil, err
	}
	return &EcpCollateralCollateralAdjustedIterator{contract: _EcpCollateral.contract, event: "CollateralAdjusted", logs: logs, sub: sub}, nil
}

// WatchCollateralAdjusted is a free log subscription operation binding the contract event 0x42f1a8a1aee108e84e4eabfaa8d51d7baaa1a02d482295297883a44b2debd3e4.
//
// Solidity: event CollateralAdjusted(address indexed cp, uint256 frozenAmount, uint256 balanceAmount, string operation)
func (_EcpCollateral *EcpCollateralFilterer) WatchCollateralAdjusted(opts *bind.WatchOpts, sink chan<- *EcpCollateralCollateralAdjusted, cp []common.Address) (event.Subscription, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _EcpCollateral.contract.WatchLogs(opts, "CollateralAdjusted", cpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpCollateralCollateralAdjusted)
				if err := _EcpCollateral.contract.UnpackLog(event, "CollateralAdjusted", log); err != nil {
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
func (_EcpCollateral *EcpCollateralFilterer) ParseCollateralAdjusted(log types.Log) (*EcpCollateralCollateralAdjusted, error) {
	event := new(EcpCollateralCollateralAdjusted)
	if err := _EcpCollateral.contract.UnpackLog(event, "CollateralAdjusted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcpCollateralCollateralLockedIterator is returned from FilterCollateralLocked and is used to iterate over the raw logs and unpacked data for CollateralLocked events raised by the EcpCollateral contract.
type EcpCollateralCollateralLockedIterator struct {
	Event *EcpCollateralCollateralLocked // Event containing the contract specifics and raw log

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
func (it *EcpCollateralCollateralLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpCollateralCollateralLocked)
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
		it.Event = new(EcpCollateralCollateralLocked)
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
func (it *EcpCollateralCollateralLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpCollateralCollateralLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpCollateralCollateralLocked represents a CollateralLocked event raised by the EcpCollateral contract.
type EcpCollateralCollateralLocked struct {
	Cp               common.Address
	CollateralAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCollateralLocked is a free log retrieval operation binding the contract event 0x2251f6a4ed7fe619e9e8ce557d05a63dd484284f9c95c9ab334f6a7707cd0800.
//
// Solidity: event CollateralLocked(address indexed cp, uint256 collateralAmount)
func (_EcpCollateral *EcpCollateralFilterer) FilterCollateralLocked(opts *bind.FilterOpts, cp []common.Address) (*EcpCollateralCollateralLockedIterator, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _EcpCollateral.contract.FilterLogs(opts, "CollateralLocked", cpRule)
	if err != nil {
		return nil, err
	}
	return &EcpCollateralCollateralLockedIterator{contract: _EcpCollateral.contract, event: "CollateralLocked", logs: logs, sub: sub}, nil
}

// WatchCollateralLocked is a free log subscription operation binding the contract event 0x2251f6a4ed7fe619e9e8ce557d05a63dd484284f9c95c9ab334f6a7707cd0800.
//
// Solidity: event CollateralLocked(address indexed cp, uint256 collateralAmount)
func (_EcpCollateral *EcpCollateralFilterer) WatchCollateralLocked(opts *bind.WatchOpts, sink chan<- *EcpCollateralCollateralLocked, cp []common.Address) (event.Subscription, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _EcpCollateral.contract.WatchLogs(opts, "CollateralLocked", cpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpCollateralCollateralLocked)
				if err := _EcpCollateral.contract.UnpackLog(event, "CollateralLocked", log); err != nil {
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

// ParseCollateralLocked is a log parse operation binding the contract event 0x2251f6a4ed7fe619e9e8ce557d05a63dd484284f9c95c9ab334f6a7707cd0800.
//
// Solidity: event CollateralLocked(address indexed cp, uint256 collateralAmount)
func (_EcpCollateral *EcpCollateralFilterer) ParseCollateralLocked(log types.Log) (*EcpCollateralCollateralLocked, error) {
	event := new(EcpCollateralCollateralLocked)
	if err := _EcpCollateral.contract.UnpackLog(event, "CollateralLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcpCollateralCollateralSlashedIterator is returned from FilterCollateralSlashed and is used to iterate over the raw logs and unpacked data for CollateralSlashed events raised by the EcpCollateral contract.
type EcpCollateralCollateralSlashedIterator struct {
	Event *EcpCollateralCollateralSlashed // Event containing the contract specifics and raw log

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
func (it *EcpCollateralCollateralSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpCollateralCollateralSlashed)
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
		it.Event = new(EcpCollateralCollateralSlashed)
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
func (it *EcpCollateralCollateralSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpCollateralCollateralSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpCollateralCollateralSlashed represents a CollateralSlashed event raised by the EcpCollateral contract.
type EcpCollateralCollateralSlashed struct {
	Cp     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCollateralSlashed is a free log retrieval operation binding the contract event 0x5138f522ae83cccdefee151fa33feeb62b6bbe619fdeb8f83cd1c6c3f8bdf921.
//
// Solidity: event CollateralSlashed(address indexed cp, uint256 amount)
func (_EcpCollateral *EcpCollateralFilterer) FilterCollateralSlashed(opts *bind.FilterOpts, cp []common.Address) (*EcpCollateralCollateralSlashedIterator, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _EcpCollateral.contract.FilterLogs(opts, "CollateralSlashed", cpRule)
	if err != nil {
		return nil, err
	}
	return &EcpCollateralCollateralSlashedIterator{contract: _EcpCollateral.contract, event: "CollateralSlashed", logs: logs, sub: sub}, nil
}

// WatchCollateralSlashed is a free log subscription operation binding the contract event 0x5138f522ae83cccdefee151fa33feeb62b6bbe619fdeb8f83cd1c6c3f8bdf921.
//
// Solidity: event CollateralSlashed(address indexed cp, uint256 amount)
func (_EcpCollateral *EcpCollateralFilterer) WatchCollateralSlashed(opts *bind.WatchOpts, sink chan<- *EcpCollateralCollateralSlashed, cp []common.Address) (event.Subscription, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _EcpCollateral.contract.WatchLogs(opts, "CollateralSlashed", cpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpCollateralCollateralSlashed)
				if err := _EcpCollateral.contract.UnpackLog(event, "CollateralSlashed", log); err != nil {
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

// ParseCollateralSlashed is a log parse operation binding the contract event 0x5138f522ae83cccdefee151fa33feeb62b6bbe619fdeb8f83cd1c6c3f8bdf921.
//
// Solidity: event CollateralSlashed(address indexed cp, uint256 amount)
func (_EcpCollateral *EcpCollateralFilterer) ParseCollateralSlashed(log types.Log) (*EcpCollateralCollateralSlashed, error) {
	event := new(EcpCollateralCollateralSlashed)
	if err := _EcpCollateral.contract.UnpackLog(event, "CollateralSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcpCollateralCollateralUnlockedIterator is returned from FilterCollateralUnlocked and is used to iterate over the raw logs and unpacked data for CollateralUnlocked events raised by the EcpCollateral contract.
type EcpCollateralCollateralUnlockedIterator struct {
	Event *EcpCollateralCollateralUnlocked // Event containing the contract specifics and raw log

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
func (it *EcpCollateralCollateralUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpCollateralCollateralUnlocked)
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
		it.Event = new(EcpCollateralCollateralUnlocked)
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
func (it *EcpCollateralCollateralUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpCollateralCollateralUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpCollateralCollateralUnlocked represents a CollateralUnlocked event raised by the EcpCollateral contract.
type EcpCollateralCollateralUnlocked struct {
	Cp               common.Address
	CollateralAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCollateralUnlocked is a free log retrieval operation binding the contract event 0xb4eaf47ecd4bc76248f192433e8067c96cb3e17aced42fbc47a512742fb74216.
//
// Solidity: event CollateralUnlocked(address indexed cp, uint256 collateralAmount)
func (_EcpCollateral *EcpCollateralFilterer) FilterCollateralUnlocked(opts *bind.FilterOpts, cp []common.Address) (*EcpCollateralCollateralUnlockedIterator, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _EcpCollateral.contract.FilterLogs(opts, "CollateralUnlocked", cpRule)
	if err != nil {
		return nil, err
	}
	return &EcpCollateralCollateralUnlockedIterator{contract: _EcpCollateral.contract, event: "CollateralUnlocked", logs: logs, sub: sub}, nil
}

// WatchCollateralUnlocked is a free log subscription operation binding the contract event 0xb4eaf47ecd4bc76248f192433e8067c96cb3e17aced42fbc47a512742fb74216.
//
// Solidity: event CollateralUnlocked(address indexed cp, uint256 collateralAmount)
func (_EcpCollateral *EcpCollateralFilterer) WatchCollateralUnlocked(opts *bind.WatchOpts, sink chan<- *EcpCollateralCollateralUnlocked, cp []common.Address) (event.Subscription, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _EcpCollateral.contract.WatchLogs(opts, "CollateralUnlocked", cpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpCollateralCollateralUnlocked)
				if err := _EcpCollateral.contract.UnpackLog(event, "CollateralUnlocked", log); err != nil {
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
func (_EcpCollateral *EcpCollateralFilterer) ParseCollateralUnlocked(log types.Log) (*EcpCollateralCollateralUnlocked, error) {
	event := new(EcpCollateralCollateralUnlocked)
	if err := _EcpCollateral.contract.UnpackLog(event, "CollateralUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcpCollateralDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the EcpCollateral contract.
type EcpCollateralDepositIterator struct {
	Event *EcpCollateralDeposit // Event containing the contract specifics and raw log

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
func (it *EcpCollateralDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpCollateralDeposit)
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
		it.Event = new(EcpCollateralDeposit)
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
func (it *EcpCollateralDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpCollateralDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpCollateralDeposit represents a Deposit event raised by the EcpCollateral contract.
type EcpCollateralDeposit struct {
	FundingWallet common.Address
	CpAccount     common.Address
	DepositAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed fundingWallet, address indexed cpAccount, uint256 depositAmount)
func (_EcpCollateral *EcpCollateralFilterer) FilterDeposit(opts *bind.FilterOpts, fundingWallet []common.Address, cpAccount []common.Address) (*EcpCollateralDepositIterator, error) {

	var fundingWalletRule []interface{}
	for _, fundingWalletItem := range fundingWallet {
		fundingWalletRule = append(fundingWalletRule, fundingWalletItem)
	}
	var cpAccountRule []interface{}
	for _, cpAccountItem := range cpAccount {
		cpAccountRule = append(cpAccountRule, cpAccountItem)
	}

	logs, sub, err := _EcpCollateral.contract.FilterLogs(opts, "Deposit", fundingWalletRule, cpAccountRule)
	if err != nil {
		return nil, err
	}
	return &EcpCollateralDepositIterator{contract: _EcpCollateral.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed fundingWallet, address indexed cpAccount, uint256 depositAmount)
func (_EcpCollateral *EcpCollateralFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *EcpCollateralDeposit, fundingWallet []common.Address, cpAccount []common.Address) (event.Subscription, error) {

	var fundingWalletRule []interface{}
	for _, fundingWalletItem := range fundingWallet {
		fundingWalletRule = append(fundingWalletRule, fundingWalletItem)
	}
	var cpAccountRule []interface{}
	for _, cpAccountItem := range cpAccount {
		cpAccountRule = append(cpAccountRule, cpAccountItem)
	}

	logs, sub, err := _EcpCollateral.contract.WatchLogs(opts, "Deposit", fundingWalletRule, cpAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpCollateralDeposit)
				if err := _EcpCollateral.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_EcpCollateral *EcpCollateralFilterer) ParseDeposit(log types.Log) (*EcpCollateralDeposit, error) {
	event := new(EcpCollateralDeposit)
	if err := _EcpCollateral.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcpCollateralDisputeProofIterator is returned from FilterDisputeProof and is used to iterate over the raw logs and unpacked data for DisputeProof events raised by the EcpCollateral contract.
type EcpCollateralDisputeProofIterator struct {
	Event *EcpCollateralDisputeProof // Event containing the contract specifics and raw log

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
func (it *EcpCollateralDisputeProofIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpCollateralDisputeProof)
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
		it.Event = new(EcpCollateralDisputeProof)
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
func (it *EcpCollateralDisputeProofIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpCollateralDisputeProofIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpCollateralDisputeProof represents a DisputeProof event raised by the EcpCollateral contract.
type EcpCollateralDisputeProof struct {
	Challenger          common.Address
	TaskContractAddress common.Address
	CpAccount           common.Address
	TaskID              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDisputeProof is a free log retrieval operation binding the contract event 0x6010bb1c19b181c59c29bde1a47441eae0c5e2e587b409d5a7ac30f01e8dcf3c.
//
// Solidity: event DisputeProof(address indexed challenger, address indexed taskContractAddress, address cpAccount, uint256 taskID)
func (_EcpCollateral *EcpCollateralFilterer) FilterDisputeProof(opts *bind.FilterOpts, challenger []common.Address, taskContractAddress []common.Address) (*EcpCollateralDisputeProofIterator, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}
	var taskContractAddressRule []interface{}
	for _, taskContractAddressItem := range taskContractAddress {
		taskContractAddressRule = append(taskContractAddressRule, taskContractAddressItem)
	}

	logs, sub, err := _EcpCollateral.contract.FilterLogs(opts, "DisputeProof", challengerRule, taskContractAddressRule)
	if err != nil {
		return nil, err
	}
	return &EcpCollateralDisputeProofIterator{contract: _EcpCollateral.contract, event: "DisputeProof", logs: logs, sub: sub}, nil
}

// WatchDisputeProof is a free log subscription operation binding the contract event 0x6010bb1c19b181c59c29bde1a47441eae0c5e2e587b409d5a7ac30f01e8dcf3c.
//
// Solidity: event DisputeProof(address indexed challenger, address indexed taskContractAddress, address cpAccount, uint256 taskID)
func (_EcpCollateral *EcpCollateralFilterer) WatchDisputeProof(opts *bind.WatchOpts, sink chan<- *EcpCollateralDisputeProof, challenger []common.Address, taskContractAddress []common.Address) (event.Subscription, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}
	var taskContractAddressRule []interface{}
	for _, taskContractAddressItem := range taskContractAddress {
		taskContractAddressRule = append(taskContractAddressRule, taskContractAddressItem)
	}

	logs, sub, err := _EcpCollateral.contract.WatchLogs(opts, "DisputeProof", challengerRule, taskContractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpCollateralDisputeProof)
				if err := _EcpCollateral.contract.UnpackLog(event, "DisputeProof", log); err != nil {
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

// ParseDisputeProof is a log parse operation binding the contract event 0x6010bb1c19b181c59c29bde1a47441eae0c5e2e587b409d5a7ac30f01e8dcf3c.
//
// Solidity: event DisputeProof(address indexed challenger, address indexed taskContractAddress, address cpAccount, uint256 taskID)
func (_EcpCollateral *EcpCollateralFilterer) ParseDisputeProof(log types.Log) (*EcpCollateralDisputeProof, error) {
	event := new(EcpCollateralDisputeProof)
	if err := _EcpCollateral.contract.UnpackLog(event, "DisputeProof", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcpCollateralOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EcpCollateral contract.
type EcpCollateralOwnershipTransferredIterator struct {
	Event *EcpCollateralOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EcpCollateralOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpCollateralOwnershipTransferred)
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
		it.Event = new(EcpCollateralOwnershipTransferred)
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
func (it *EcpCollateralOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpCollateralOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpCollateralOwnershipTransferred represents a OwnershipTransferred event raised by the EcpCollateral contract.
type EcpCollateralOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EcpCollateral *EcpCollateralFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EcpCollateralOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EcpCollateral.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EcpCollateralOwnershipTransferredIterator{contract: _EcpCollateral.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EcpCollateral *EcpCollateralFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EcpCollateralOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EcpCollateral.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpCollateralOwnershipTransferred)
				if err := _EcpCollateral.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EcpCollateral *EcpCollateralFilterer) ParseOwnershipTransferred(log types.Log) (*EcpCollateralOwnershipTransferred, error) {
	event := new(EcpCollateralOwnershipTransferred)
	if err := _EcpCollateral.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcpCollateralWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the EcpCollateral contract.
type EcpCollateralWithdrawIterator struct {
	Event *EcpCollateralWithdraw // Event containing the contract specifics and raw log

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
func (it *EcpCollateralWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpCollateralWithdraw)
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
		it.Event = new(EcpCollateralWithdraw)
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
func (it *EcpCollateralWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpCollateralWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpCollateralWithdraw represents a Withdraw event raised by the EcpCollateral contract.
type EcpCollateralWithdraw struct {
	CpOwner        common.Address
	CpAccount      common.Address
	WithdrawAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed cpOwner, address indexed cpAccount, uint256 withdrawAmount)
func (_EcpCollateral *EcpCollateralFilterer) FilterWithdraw(opts *bind.FilterOpts, cpOwner []common.Address, cpAccount []common.Address) (*EcpCollateralWithdrawIterator, error) {

	var cpOwnerRule []interface{}
	for _, cpOwnerItem := range cpOwner {
		cpOwnerRule = append(cpOwnerRule, cpOwnerItem)
	}
	var cpAccountRule []interface{}
	for _, cpAccountItem := range cpAccount {
		cpAccountRule = append(cpAccountRule, cpAccountItem)
	}

	logs, sub, err := _EcpCollateral.contract.FilterLogs(opts, "Withdraw", cpOwnerRule, cpAccountRule)
	if err != nil {
		return nil, err
	}
	return &EcpCollateralWithdrawIterator{contract: _EcpCollateral.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed cpOwner, address indexed cpAccount, uint256 withdrawAmount)
func (_EcpCollateral *EcpCollateralFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *EcpCollateralWithdraw, cpOwner []common.Address, cpAccount []common.Address) (event.Subscription, error) {

	var cpOwnerRule []interface{}
	for _, cpOwnerItem := range cpOwner {
		cpOwnerRule = append(cpOwnerRule, cpOwnerItem)
	}
	var cpAccountRule []interface{}
	for _, cpAccountItem := range cpAccount {
		cpAccountRule = append(cpAccountRule, cpAccountItem)
	}

	logs, sub, err := _EcpCollateral.contract.WatchLogs(opts, "Withdraw", cpOwnerRule, cpAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpCollateralWithdraw)
				if err := _EcpCollateral.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_EcpCollateral *EcpCollateralFilterer) ParseWithdraw(log types.Log) (*EcpCollateralWithdraw, error) {
	event := new(EcpCollateralWithdraw)
	if err := _EcpCollateral.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcpCollateralWithdrawConfirmedIterator is returned from FilterWithdrawConfirmed and is used to iterate over the raw logs and unpacked data for WithdrawConfirmed events raised by the EcpCollateral contract.
type EcpCollateralWithdrawConfirmedIterator struct {
	Event *EcpCollateralWithdrawConfirmed // Event containing the contract specifics and raw log

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
func (it *EcpCollateralWithdrawConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpCollateralWithdrawConfirmed)
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
		it.Event = new(EcpCollateralWithdrawConfirmed)
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
func (it *EcpCollateralWithdrawConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpCollateralWithdrawConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpCollateralWithdrawConfirmed represents a WithdrawConfirmed event raised by the EcpCollateral contract.
type EcpCollateralWithdrawConfirmed struct {
	Cp     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawConfirmed is a free log retrieval operation binding the contract event 0x1a98aba99d2d38026b07feddaca8e333649ae8a5f5a238687f91ce7791ee998e.
//
// Solidity: event WithdrawConfirmed(address indexed cp, uint256 amount)
func (_EcpCollateral *EcpCollateralFilterer) FilterWithdrawConfirmed(opts *bind.FilterOpts, cp []common.Address) (*EcpCollateralWithdrawConfirmedIterator, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _EcpCollateral.contract.FilterLogs(opts, "WithdrawConfirmed", cpRule)
	if err != nil {
		return nil, err
	}
	return &EcpCollateralWithdrawConfirmedIterator{contract: _EcpCollateral.contract, event: "WithdrawConfirmed", logs: logs, sub: sub}, nil
}

// WatchWithdrawConfirmed is a free log subscription operation binding the contract event 0x1a98aba99d2d38026b07feddaca8e333649ae8a5f5a238687f91ce7791ee998e.
//
// Solidity: event WithdrawConfirmed(address indexed cp, uint256 amount)
func (_EcpCollateral *EcpCollateralFilterer) WatchWithdrawConfirmed(opts *bind.WatchOpts, sink chan<- *EcpCollateralWithdrawConfirmed, cp []common.Address) (event.Subscription, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _EcpCollateral.contract.WatchLogs(opts, "WithdrawConfirmed", cpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpCollateralWithdrawConfirmed)
				if err := _EcpCollateral.contract.UnpackLog(event, "WithdrawConfirmed", log); err != nil {
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

// ParseWithdrawConfirmed is a log parse operation binding the contract event 0x1a98aba99d2d38026b07feddaca8e333649ae8a5f5a238687f91ce7791ee998e.
//
// Solidity: event WithdrawConfirmed(address indexed cp, uint256 amount)
func (_EcpCollateral *EcpCollateralFilterer) ParseWithdrawConfirmed(log types.Log) (*EcpCollateralWithdrawConfirmed, error) {
	event := new(EcpCollateralWithdrawConfirmed)
	if err := _EcpCollateral.contract.UnpackLog(event, "WithdrawConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcpCollateralWithdrawRequestCanceledIterator is returned from FilterWithdrawRequestCanceled and is used to iterate over the raw logs and unpacked data for WithdrawRequestCanceled events raised by the EcpCollateral contract.
type EcpCollateralWithdrawRequestCanceledIterator struct {
	Event *EcpCollateralWithdrawRequestCanceled // Event containing the contract specifics and raw log

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
func (it *EcpCollateralWithdrawRequestCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpCollateralWithdrawRequestCanceled)
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
		it.Event = new(EcpCollateralWithdrawRequestCanceled)
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
func (it *EcpCollateralWithdrawRequestCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpCollateralWithdrawRequestCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpCollateralWithdrawRequestCanceled represents a WithdrawRequestCanceled event raised by the EcpCollateral contract.
type EcpCollateralWithdrawRequestCanceled struct {
	Cp     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawRequestCanceled is a free log retrieval operation binding the contract event 0xa3895d397a34e928a95d593331e293e2fc281d9d8996df5cc6c57f1cef629d42.
//
// Solidity: event WithdrawRequestCanceled(address indexed cp, uint256 amount)
func (_EcpCollateral *EcpCollateralFilterer) FilterWithdrawRequestCanceled(opts *bind.FilterOpts, cp []common.Address) (*EcpCollateralWithdrawRequestCanceledIterator, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _EcpCollateral.contract.FilterLogs(opts, "WithdrawRequestCanceled", cpRule)
	if err != nil {
		return nil, err
	}
	return &EcpCollateralWithdrawRequestCanceledIterator{contract: _EcpCollateral.contract, event: "WithdrawRequestCanceled", logs: logs, sub: sub}, nil
}

// WatchWithdrawRequestCanceled is a free log subscription operation binding the contract event 0xa3895d397a34e928a95d593331e293e2fc281d9d8996df5cc6c57f1cef629d42.
//
// Solidity: event WithdrawRequestCanceled(address indexed cp, uint256 amount)
func (_EcpCollateral *EcpCollateralFilterer) WatchWithdrawRequestCanceled(opts *bind.WatchOpts, sink chan<- *EcpCollateralWithdrawRequestCanceled, cp []common.Address) (event.Subscription, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _EcpCollateral.contract.WatchLogs(opts, "WithdrawRequestCanceled", cpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpCollateralWithdrawRequestCanceled)
				if err := _EcpCollateral.contract.UnpackLog(event, "WithdrawRequestCanceled", log); err != nil {
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
// Solidity: event WithdrawRequestCanceled(address indexed cp, uint256 amount)
func (_EcpCollateral *EcpCollateralFilterer) ParseWithdrawRequestCanceled(log types.Log) (*EcpCollateralWithdrawRequestCanceled, error) {
	event := new(EcpCollateralWithdrawRequestCanceled)
	if err := _EcpCollateral.contract.UnpackLog(event, "WithdrawRequestCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcpCollateralWithdrawRequestedIterator is returned from FilterWithdrawRequested and is used to iterate over the raw logs and unpacked data for WithdrawRequested events raised by the EcpCollateral contract.
type EcpCollateralWithdrawRequestedIterator struct {
	Event *EcpCollateralWithdrawRequested // Event containing the contract specifics and raw log

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
func (it *EcpCollateralWithdrawRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpCollateralWithdrawRequested)
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
		it.Event = new(EcpCollateralWithdrawRequested)
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
func (it *EcpCollateralWithdrawRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpCollateralWithdrawRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpCollateralWithdrawRequested represents a WithdrawRequested event raised by the EcpCollateral contract.
type EcpCollateralWithdrawRequested struct {
	Cp     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawRequested is a free log retrieval operation binding the contract event 0xf7774b688d56120b783560a913ee60792a73dfd511812b7be5eccf10d08c6689.
//
// Solidity: event WithdrawRequested(address indexed cp, uint256 amount)
func (_EcpCollateral *EcpCollateralFilterer) FilterWithdrawRequested(opts *bind.FilterOpts, cp []common.Address) (*EcpCollateralWithdrawRequestedIterator, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _EcpCollateral.contract.FilterLogs(opts, "WithdrawRequested", cpRule)
	if err != nil {
		return nil, err
	}
	return &EcpCollateralWithdrawRequestedIterator{contract: _EcpCollateral.contract, event: "WithdrawRequested", logs: logs, sub: sub}, nil
}

// WatchWithdrawRequested is a free log subscription operation binding the contract event 0xf7774b688d56120b783560a913ee60792a73dfd511812b7be5eccf10d08c6689.
//
// Solidity: event WithdrawRequested(address indexed cp, uint256 amount)
func (_EcpCollateral *EcpCollateralFilterer) WatchWithdrawRequested(opts *bind.WatchOpts, sink chan<- *EcpCollateralWithdrawRequested, cp []common.Address) (event.Subscription, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _EcpCollateral.contract.WatchLogs(opts, "WithdrawRequested", cpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpCollateralWithdrawRequested)
				if err := _EcpCollateral.contract.UnpackLog(event, "WithdrawRequested", log); err != nil {
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
// Solidity: event WithdrawRequested(address indexed cp, uint256 amount)
func (_EcpCollateral *EcpCollateralFilterer) ParseWithdrawRequested(log types.Log) (*EcpCollateralWithdrawRequested, error) {
	event := new(EcpCollateralWithdrawRequested)
	if err := _EcpCollateral.contract.UnpackLog(event, "WithdrawRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcpCollateralWithdrawSlashIterator is returned from FilterWithdrawSlash and is used to iterate over the raw logs and unpacked data for WithdrawSlash events raised by the EcpCollateral contract.
type EcpCollateralWithdrawSlashIterator struct {
	Event *EcpCollateralWithdrawSlash // Event containing the contract specifics and raw log

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
func (it *EcpCollateralWithdrawSlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpCollateralWithdrawSlash)
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
		it.Event = new(EcpCollateralWithdrawSlash)
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
func (it *EcpCollateralWithdrawSlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpCollateralWithdrawSlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpCollateralWithdrawSlash represents a WithdrawSlash event raised by the EcpCollateral contract.
type EcpCollateralWithdrawSlash struct {
	CollateralContratOwner common.Address
	Slashfund              *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawSlash is a free log retrieval operation binding the contract event 0xbfd9c82485e2178fcfd5c40379d8e66fe60acc50aa1ef4c50966431eb1e381cd.
//
// Solidity: event WithdrawSlash(address indexed collateralContratOwner, uint256 slashfund)
func (_EcpCollateral *EcpCollateralFilterer) FilterWithdrawSlash(opts *bind.FilterOpts, collateralContratOwner []common.Address) (*EcpCollateralWithdrawSlashIterator, error) {

	var collateralContratOwnerRule []interface{}
	for _, collateralContratOwnerItem := range collateralContratOwner {
		collateralContratOwnerRule = append(collateralContratOwnerRule, collateralContratOwnerItem)
	}

	logs, sub, err := _EcpCollateral.contract.FilterLogs(opts, "WithdrawSlash", collateralContratOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EcpCollateralWithdrawSlashIterator{contract: _EcpCollateral.contract, event: "WithdrawSlash", logs: logs, sub: sub}, nil
}

// WatchWithdrawSlash is a free log subscription operation binding the contract event 0xbfd9c82485e2178fcfd5c40379d8e66fe60acc50aa1ef4c50966431eb1e381cd.
//
// Solidity: event WithdrawSlash(address indexed collateralContratOwner, uint256 slashfund)
func (_EcpCollateral *EcpCollateralFilterer) WatchWithdrawSlash(opts *bind.WatchOpts, sink chan<- *EcpCollateralWithdrawSlash, collateralContratOwner []common.Address) (event.Subscription, error) {

	var collateralContratOwnerRule []interface{}
	for _, collateralContratOwnerItem := range collateralContratOwner {
		collateralContratOwnerRule = append(collateralContratOwnerRule, collateralContratOwnerItem)
	}

	logs, sub, err := _EcpCollateral.contract.WatchLogs(opts, "WithdrawSlash", collateralContratOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpCollateralWithdrawSlash)
				if err := _EcpCollateral.contract.UnpackLog(event, "WithdrawSlash", log); err != nil {
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
func (_EcpCollateral *EcpCollateralFilterer) ParseWithdrawSlash(log types.Log) (*EcpCollateralWithdrawSlash, error) {
	event := new(EcpCollateralWithdrawSlash)
	if err := _EcpCollateral.contract.UnpackLog(event, "WithdrawSlash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
