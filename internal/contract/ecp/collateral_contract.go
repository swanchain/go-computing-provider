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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cps\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"BatchFrozenDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"frozenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"operation\",\"type\":\"string\"}],\"name\":\"CollateralAdjusted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"CollateralLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CollateralSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"CollateralUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fundingWallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taskContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"}],\"name\":\"DisputeProof\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cpOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawRequestCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralContratOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashfund\",\"type\":\"uint256\"}],\"name\":\"WithdrawSlash\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"cps\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchFrozenDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"cps\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"taskCollaterals\",\"type\":\"uint256[]\"}],\"name\":\"batchLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"cps\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"slashAmounts\",\"type\":\"uint256[]\"}],\"name\":\"batchSlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"cps\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"taskCollaterals\",\"type\":\"uint256[]\"}],\"name\":\"batchUnlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"}],\"name\":\"cancelWithdrawRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"}],\"name\":\"confirmWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"}],\"name\":\"cpInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"balance\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"frozenBalance\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"}],\"internalType\":\"structECPCollateral.CPInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cpStatus\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"taskContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"}],\"name\":\"disputeProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"frozenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getECPCollateralInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slashedFunds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawDelay\",\"type\":\"uint256\"}],\"internalType\":\"structECPCollateral.ContractInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"taskCollateral\",\"type\":\"uint256\"}],\"name\":\"lockCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"requestWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseCollateral\",\"type\":\"uint256\"}],\"name\":\"setBaseCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_collateralRatio\",\"type\":\"uint256\"}],\"name\":\"setCollateralRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"setCollateralToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slashRatio\",\"type\":\"uint256\"}],\"name\":\"setSlashRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawDelay\",\"type\":\"uint256\"}],\"name\":\"setWithdrawDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slashAmount\",\"type\":\"uint256\"}],\"name\":\"slashCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashedFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"taskCollateral\",\"type\":\"uint256\"}],\"name\":\"unlockCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"}],\"name\":\"viewWithdrawRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestBlock\",\"type\":\"uint256\"}],\"internalType\":\"structECPCollateral.WithdrawRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slashfund\",\"type\":\"uint256\"}],\"name\":\"withdrawSlashedFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526201d88060065534801561001757600080fd5b5033600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361008b5760006040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401610082919061021c565b60405180910390fd5b61009a8161011760201b60201c565b506100aa3361011760201b60201c565b6001600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060056004819055506002600581905550610237565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610206826101db565b9050919050565b610216816101fb565b82525050565b6000602082019050610231600083018461020d565b92915050565b61530d806102466000396000f3fe608060405234801561001057600080fd5b506004361061023d5760003560e01c8063704802751161013b578063a664c216116100b8578063ce3518aa1161007c578063ce3518aa14610685578063d27ca89b146106a1578063d2bfc1c7146106bf578063f2fde38b146106db578063f3fef3a3146106f75761023d565b8063a664c216146105cd578063b2016bd4146105fd578063b4eae1cb1461061b578063bede6e3114610639578063c6ff4555146106695761023d565b80638331f8e5116100ff5780638331f8e51461053b5780638da5cb5b146105575780639939cd1814610575578063998fc5cc146105935780639b5ddf09146105af5761023d565b806370480275146104bf57806370b72944146104db578063715018a6146104f957806372f0cb30146105035780637f58a7e51461051f5761023d565b80633fe65177116101c95780635f7d0e841161018d5780635f7d0e84146104315780636060663e1461044d57806363215bb714610469578063666181a9146104855780636f99f15c146104a15761023d565b80633fe651771461037c57806347a7d107146103ac57806347e7ef24146103c857806352df49ec146103e45780635d2cd2a7146104155761023d565b806324d7806c1161021057806324d7806c146102b4578063266565a9146102e457806327e235e3146103145780632894493f14610344578063397a1b28146103605761023d565b80630288a39c146102425780631785f53c1461026057806317f174981461027c5780631d47a62d14610298575b600080fd5b61024a610713565b60405161025791906139cc565b60405180910390f35b61027a60048036038101906102759190613a4f565b610719565b005b61029660048036038101906102919190613aa8565b61077c565b005b6102b260048036038101906102ad9190613aa8565b6109d7565b005b6102ce60048036038101906102c99190613a4f565b610db6565b6040516102db9190613b03565b60405180910390f35b6102fe60048036038101906102f99190613a4f565b610dd6565b60405161030b91906139cc565b60405180910390f35b61032e60048036038101906103299190613a4f565b610dee565b60405161033b9190613b37565b60405180910390f35b61035e60048036038101906103599190613b52565b610e06565b005b61037a60048036038101906103759190613aa8565b610f5e565b005b61039660048036038101906103919190613a4f565b6112c6565b6040516103a39190613c0f565b60405180910390f35b6103c660048036038101906103c19190613aa8565b611366565b005b6103e260048036038101906103dd9190613aa8565b6115ea565b005b6103fe60048036038101906103f99190613a4f565b611807565b60405161040c929190613c31565b60405180910390f35b61042f600480360381019061042a9190613a4f565b61182b565b005b61044b60048036038101906104469190613d15565b611b86565b005b61046760048036038101906104629190613b52565b611d5c565b005b610483600480360381019061047e9190613d15565b611d6e565b005b61049f600480360381019061049a9190613a4f565b611f44565b005b6104a9611fff565b6040516104b691906139cc565b60405180910390f35b6104d960048036038101906104d49190613a4f565b612005565b005b6104e3612068565b6040516104f091906139cc565b60405180910390f35b610501612072565b005b61051d60048036038101906105189190613b52565b612086565b005b61053960048036038101906105349190613b52565b612098565b005b61055560048036038101906105509190613d96565b61212e565b005b61055f612278565b60405161056c9190613df8565b60405180910390f35b61057d6122a1565b60405161058a9190613eac565b60405180910390f35b6105ad60048036038101906105a89190613d15565b61231d565b005b6105b76126e9565b6040516105c491906139cc565b60405180910390f35b6105e760048036038101906105e29190613a4f565b6126ef565b6040516105f49190613f83565b60405180910390f35b6106056128e9565b6040516106129190614004565b60405180910390f35b61062361290f565b60405161063091906139cc565b60405180910390f35b610653600480360381019061064e9190613a4f565b612915565b604051610660919061404e565b60405180910390f35b610683600480360381019061067e9190613d15565b6129f1565b005b61069f600480360381019061069a9190613b52565b612bc7565b005b6106a9612bd9565b6040516106b691906139cc565b60405180910390f35b6106d960048036038101906106d49190613a4f565b612bdf565b005b6106f560048036038101906106f09190613a4f565b613113565b005b610711600480360381019061070c9190613aa8565b613199565b005b60065481565b6107216135ad565b6000600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610808576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107ff906140db565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610877576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161086e9061416d565b60405180910390fd5b6000600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060008183116108ca57826108cc565b815b905080600960008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461091d91906141bc565b9250508190555080600860008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461097391906141f0565b9250508190555061098384613634565b8373ffffffffffffffffffffffffffffffffffffffff167fb4eaf47ecd4bc76248f192433e8067c96cb3e17aced42fbc47a512742fb74216826040516109c991906139cc565b60405180910390a250505050565b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610a63576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a5a906140db565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610ad2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ac99061416d565b60405180910390fd5b60008111610b15576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b0c906142a6565b60405180910390fd5b6000600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050600080600860008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205413610ba8576000610be9565b600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020545b90506000828411610bfa5783610bfc565b825b90506000818511610c0e576000610c1b565b8185610c1a91906141bc565b5b90506000838211610c2c5781610c2e565b835b905060008184610c3e91906142c6565b905083600960008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610c8f91906141bc565b9250508190555082600860008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610ce591906142fa565b925050819055508060026000828254610cfe91906142c6565b92505081905550610d0e88613634565b8773ffffffffffffffffffffffffffffffffffffffff167f5138f522ae83cccdefee151fa33feeb62b6bbe619fdeb8f83cd1c6c3f8bdf92182604051610d5491906139cc565b60405180910390a28773ffffffffffffffffffffffffffffffffffffffff167f42f1a8a1aee108e84e4eabfaa8d51d7baaa1a02d482295297883a44b2debd3e48584604051610da4929190614389565b60405180910390a25050505050505050565b60076020528060005260406000206000915054906101000a900460ff1681565b60096020528060005260406000206000915090505481565b60086020528060005260406000206000915090505481565b610e0e6135ad565b806002541015610e53576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e4a90614437565b60405180910390fd5b8060026000828254610e6591906141bc565b92505081905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b8152600401610ec9929190614457565b6020604051808303816000875af1158015610ee8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f0c91906144ac565b503373ffffffffffffffffffffffffffffffffffffffff167fbfd9c82485e2178fcfd5c40379d8e66fe60acc50aa1ef4c50966431eb1e381cd82604051610f5391906139cc565b60405180910390a250565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610fcd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fc49061454b565b60405180910390fd5b6000808373ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527f893d20e8000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505060405161107791906145b2565b6000604051808303816000865af19150503d80600081146110b4576040519150601f19603f3d011682016040523d82523d6000602084013e6110b9565b606091505b5091509150816110fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110f59061463b565b60405180910390fd5b6000818060200190518101906111149190614699565b90508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611184576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161117b90614738565b60405180910390fd5b83600960008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015611206576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111fd906147ca565b60405180910390fd5b604051806040016040528085815260200143815250600b60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000155602082015181600101559050508473ffffffffffffffffffffffffffffffffffffffff167ff7774b688d56120b783560a913ee60792a73dfd511812b7be5eccf10d08c6689856040516112b791906139cc565b60405180910390a25050505050565b600a60205280600052604060002060009150905080546112e590614819565b80601f016020809104026020016040519081016040528092919081815260200182805461131190614819565b801561135e5780601f106113335761010080835404028352916020019161135e565b820191906000526020600020905b81548152906001019060200180831161134157829003601f168201915b505050505081565b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166113f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113e9906140db565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611461576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114589061416d565b60405180910390fd5b80600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205412156114e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114da906148bc565b60405180910390fd5b80600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461153291906142fa565b9250508190555080600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461158891906142c6565b9250508190555061159882613634565b8173ffffffffffffffffffffffffffffffffffffffff167f2251f6a4ed7fe619e9e8ce557d05a63dd484284f9c95c9ab334f6a7707cd0800826040516115de91906139cc565b60405180910390a25050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611659576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116509061454b565b60405180910390fd5b6000811161169c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116939061494e565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b81526004016116fb9392919061496e565b6020604051808303816000875af115801561171a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061173e91906144ac565b5080600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461178e91906141f0565b925050819055508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62836040516117f291906139cc565b60405180910390a361180382613634565b5050565b600b6020528060005260406000206000915090508060000154908060010154905082565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361189a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118919061454b565b60405180910390fd5b6000808273ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527f893d20e8000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505060405161194491906145b2565b6000604051808303816000865af19150503d8060008114611981576040519150601f19603f3d011682016040523d82523d6000602084013e611986565b606091505b5091509150816119cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119c29061463b565b60405180910390fd5b6000818060200190518101906119e19190614699565b90508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611a51576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a4890614a17565b60405180910390fd5b6000600b60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000816000015411611adb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ad290614a83565b60405180910390fd5b8473ffffffffffffffffffffffffffffffffffffffff167fa3895d397a34e928a95d593331e293e2fc281d9d8996df5cc6c57f1cef629d428260000154604051611b2591906139cc565b60405180910390a2600b60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600080820160009055600182016000905550505050505050565b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611c12576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c09906140db565b60405180910390fd5b818190508484905014611c5a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c5190614aef565b60405180910390fd5b60005b84849050811015611d5557600073ffffffffffffffffffffffffffffffffffffffff16858583818110611c9357611c92614b0f565b5b9050602002016020810190611ca89190613a4f565b73ffffffffffffffffffffffffffffffffffffffff1603611cfe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611cf59061416d565b60405180910390fd5b611d48858583818110611d1457611d13614b0f565b5b9050602002016020810190611d299190613a4f565b848484818110611d3c57611d3b614b0f565b5b9050602002013561077c565b8080600101915050611c5d565b5050505050565b611d646135ad565b8060048190555050565b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611dfa576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611df1906140db565b60405180910390fd5b818190508484905014611e42576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e3990614aef565b60405180910390fd5b60005b84849050811015611f3d57600073ffffffffffffffffffffffffffffffffffffffff16858583818110611e7b57611e7a614b0f565b5b9050602002016020810190611e909190613a4f565b73ffffffffffffffffffffffffffffffffffffffff1603611ee6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611edd9061416d565b60405180910390fd5b611f30858583818110611efc57611efb614b0f565b5b9050602002016020810190611f119190613a4f565b848484818110611f2457611f23614b0f565b5b90506020020135611366565b8080600101915050611e45565b5050505050565b611f4c6135ad565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611fbb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611fb290614bb0565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60025481565b61200d6135ad565b6001600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b6000600354905090565b61207a6135ad565b6120846000613843565b565b61208e6135ad565b8060068190555050565b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16612124576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161211b906140db565b60405180910390fd5b8060038190555050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361219d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121949061454b565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361220c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161220390614c42565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f6010bb1c19b181c59c29bde1a47441eae0c5e2e587b409d5a7ac30f01e8dcf3c848460405161226b929190614457565b60405180910390a3505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6122a961390f565b6040518060c00160405280600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002548152602001600354815260200160045481526020016005548152602001600654815250905090565b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166123a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016123a0906140db565b60405180910390fd5b8181905084849050146123f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016123e890614aef565b60405180910390fd5b6000805b838390508110156124345783838281811061241357612412614b0f565b5b905060200201358261242591906142c6565b915080806001019150506123f5565b50600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b81526004016124949392919061496e565b6020604051808303816000875af11580156124b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124d791906144ac565b612516576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161250d90614cae565b60405180910390fd5b60005b8585905081101561268d57600073ffffffffffffffffffffffffffffffffffffffff1686868381811061254f5761254e614b0f565b5b90506020020160208101906125649190613a4f565b73ffffffffffffffffffffffffffffffffffffffff16036125ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125b19061416d565b60405180910390fd5b8383828181106125cd576125cc614b0f565b5b90506020020135600960008888858181106125eb576125ea614b0f565b5b90506020020160208101906126009190613a4f565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461264991906142c6565b9250508190555061268086868381811061266657612665614b0f565b5b905060200201602081019061267b9190613a4f565b613634565b8080600101915050612519565b503373ffffffffffffffffffffffffffffffffffffffff167ff6bc9c339713a32d9c780cc55ab0d0b62a32de9986165175992e78e6973e5b1b868686866040516126da9493929190614dfd565b60405180910390a25050505050565b60035481565b6126f761395b565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603612766576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161275d9061454b565b60405180910390fd5b60405180608001604052808373ffffffffffffffffffffffffffffffffffffffff168152602001600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548152602001600960008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548152602001600a60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805461286190614819565b80601f016020809104026020016040519081016040528092919081815260200182805461288d90614819565b80156128da5780601f106128af576101008083540402835291602001916128da565b820191906000526020600020905b8154815290600101906020018083116128bd57829003601f168201915b50505050508152509050919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60045481565b61291d613999565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361298c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016129839061454b565b60405180910390fd5b600b60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806040016040529081600082015481526020016001820154815250509050919050565b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16612a7d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a74906140db565b60405180910390fd5b818190508484905014612ac5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612abc90614aef565b60405180910390fd5b60005b84849050811015612bc057600073ffffffffffffffffffffffffffffffffffffffff16858583818110612afe57612afd614b0f565b5b9050602002016020810190612b139190613a4f565b73ffffffffffffffffffffffffffffffffffffffff1603612b69576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b609061416d565b60405180910390fd5b612bb3858583818110612b7f57612b7e614b0f565b5b9050602002016020810190612b949190613a4f565b848484818110612ba757612ba6614b0f565b5b905060200201356109d7565b8080600101915050612ac8565b5050505050565b612bcf6135ad565b8060058190555050565b60055481565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612c4e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c459061454b565b60405180910390fd5b6000808273ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527f893d20e8000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051612cf891906145b2565b6000604051808303816000865af19150503d8060008114612d35576040519150601f19603f3d011682016040523d82523d6000602084013e612d3a565b606091505b509150915081612d7f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d769061463b565b60405180910390fd5b600081806020019051810190612d959190614699565b90508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612e05576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612dfc90614eaa565b60405180910390fd5b6000600b60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000816000015411612e8f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e8690614a83565b60405180910390fd5b6006548160010154612ea191906142c6565b431015612ee3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612eda90614f16565b60405180910390fd5b8060000154600960008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015612f69576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612f60906147ca565b60405180910390fd5b60008160000154905080600960008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254612fc191906141bc565b92505081905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84836040518363ffffffff1660e01b8152600401613025929190614457565b6020604051808303816000875af1158015613044573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061306891906144ac565b508573ffffffffffffffffffffffffffffffffffffffff167fe6fc6292b9fd5ba92a34a05a92d07b066535578023d71f0c6bf83a2622ca4e5484836040516130b1929190614457565b60405180910390a2600b60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008082016000905560018201600090555050505050505050565b61311b6135ad565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361318d5760006040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016131849190613df8565b60405180910390fd5b61319681613843565b50565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603613208576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016131ff9061454b565b60405180910390fd5b6000808373ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527f893d20e8000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516132b291906145b2565b6000604051808303816000865af19150503d80600081146132ef576040519150601f19603f3d011682016040523d82523d6000602084013e6132f4565b606091505b509150915081613339576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133309061463b565b60405180910390fd5b60008180602001905181019061334f9190614699565b905083600860008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205412156133d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133ca90614f82565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614613441576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161343890615014565b60405180910390fd5b83600860008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461349091906142fa565b92505081905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33866040518363ffffffff1660e01b81526004016134f4929190614457565b6020604051808303816000875af1158015613513573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061353791906144ac565b5061354185613634565b8473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb8660405161359e91906139cc565b60405180910390a35050505050565b6135b5613907565b73ffffffffffffffffffffffffffffffffffffffff166135d3612278565b73ffffffffffffffffffffffffffffffffffffffff1614613632576135f6613907565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016136299190613df8565b60405180910390fd5b565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036136a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161369a9061454b565b60405180910390fd5b6000600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414801561373157506000600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054145b156137bd576040518060400160405280600381526020017f4e53430000000000000000000000000000000000000000000000000000000000815250600a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090816137b79190615205565b50613840565b6040518060400160405280600681526020017f4163746976650000000000000000000000000000000000000000000000000000815250600a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020908161383e9190615205565b505b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b6040518060c00160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081526020016000815260200160008152602001600081525090565b6040518060800160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001606081525090565b604051806040016040528060008152602001600081525090565b6000819050919050565b6139c6816139b3565b82525050565b60006020820190506139e160008301846139bd565b92915050565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000613a1c826139f1565b9050919050565b613a2c81613a11565b8114613a3757600080fd5b50565b600081359050613a4981613a23565b92915050565b600060208284031215613a6557613a646139e7565b5b6000613a7384828501613a3a565b91505092915050565b613a85816139b3565b8114613a9057600080fd5b50565b600081359050613aa281613a7c565b92915050565b60008060408385031215613abf57613abe6139e7565b5b6000613acd85828601613a3a565b9250506020613ade85828601613a93565b9150509250929050565b60008115159050919050565b613afd81613ae8565b82525050565b6000602082019050613b186000830184613af4565b92915050565b6000819050919050565b613b3181613b1e565b82525050565b6000602082019050613b4c6000830184613b28565b92915050565b600060208284031215613b6857613b676139e7565b5b6000613b7684828501613a93565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015613bb9578082015181840152602081019050613b9e565b60008484015250505050565b6000601f19601f8301169050919050565b6000613be182613b7f565b613beb8185613b8a565b9350613bfb818560208601613b9b565b613c0481613bc5565b840191505092915050565b60006020820190508181036000830152613c298184613bd6565b905092915050565b6000604082019050613c4660008301856139bd565b613c5360208301846139bd565b9392505050565b600080fd5b600080fd5b600080fd5b60008083601f840112613c7f57613c7e613c5a565b5b8235905067ffffffffffffffff811115613c9c57613c9b613c5f565b5b602083019150836020820283011115613cb857613cb7613c64565b5b9250929050565b60008083601f840112613cd557613cd4613c5a565b5b8235905067ffffffffffffffff811115613cf257613cf1613c5f565b5b602083019150836020820283011115613d0e57613d0d613c64565b5b9250929050565b60008060008060408587031215613d2f57613d2e6139e7565b5b600085013567ffffffffffffffff811115613d4d57613d4c6139ec565b5b613d5987828801613c69565b9450945050602085013567ffffffffffffffff811115613d7c57613d7b6139ec565b5b613d8887828801613cbf565b925092505092959194509250565b600080600060608486031215613daf57613dae6139e7565b5b6000613dbd86828701613a3a565b9350506020613dce86828701613a3a565b9250506040613ddf86828701613a93565b9150509250925092565b613df281613a11565b82525050565b6000602082019050613e0d6000830184613de9565b92915050565b613e1c81613a11565b82525050565b613e2b816139b3565b82525050565b60c082016000820151613e476000850182613e13565b506020820151613e5a6020850182613e22565b506040820151613e6d6040850182613e22565b506060820151613e806060850182613e22565b506080820151613e936080850182613e22565b5060a0820151613ea660a0850182613e22565b50505050565b600060c082019050613ec16000830184613e31565b92915050565b613ed081613b1e565b82525050565b600082825260208201905092915050565b6000613ef282613b7f565b613efc8185613ed6565b9350613f0c818560208601613b9b565b613f1581613bc5565b840191505092915050565b6000608083016000830151613f386000860182613e13565b506020830151613f4b6020860182613ec7565b506040830151613f5e6040860182613e22565b5060608301518482036060860152613f768282613ee7565b9150508091505092915050565b60006020820190508181036000830152613f9d8184613f20565b905092915050565b6000819050919050565b6000613fca613fc5613fc0846139f1565b613fa5565b6139f1565b9050919050565b6000613fdc82613faf565b9050919050565b6000613fee82613fd1565b9050919050565b613ffe81613fe3565b82525050565b60006020820190506140196000830184613ff5565b92915050565b6040820160008201516140356000850182613e22565b5060208201516140486020850182613e22565b50505050565b6000604082019050614063600083018461401f565b92915050565b7f4f6e6c79207468652061646d696e2063616e2063616c6c20746869732066756e60008201527f6374696f6e2e0000000000000000000000000000000000000000000000000000602082015250565b60006140c5602683613b8a565b91506140d082614069565b604082019050919050565b600060208201905081810360008301526140f4816140b8565b9050919050565b7f496e76616c696420616464726573733a2063702063616e6e6f7420626520746860008201527f65207a65726f2061646472657373000000000000000000000000000000000000602082015250565b6000614157602e83613b8a565b9150614162826140fb565b604082019050919050565b600060208201905081810360008301526141868161414a565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006141c7826139b3565b91506141d2836139b3565b92508282039050818111156141ea576141e961418d565b5b92915050565b60006141fb82613b1e565b915061420683613b1e565b92508282019050828112156000831216838212600084121516171561422e5761422d61418d565b5b92915050565b7f496e76616c696420616d6f756e743a20736c617368416d6f756e74206d75737460008201527f2062652067726561746572207468616e207a65726f0000000000000000000000602082015250565b6000614290603583613b8a565b915061429b82614234565b604082019050919050565b600060208201905081810360008301526142bf81614283565b9050919050565b60006142d1826139b3565b91506142dc836139b3565b92508282019050808211156142f4576142f361418d565b5b92915050565b600061430582613b1e565b915061431083613b1e565b92508282039050818112600084121682821360008512151617156143375761433661418d565b5b92915050565b7f536c617368656400000000000000000000000000000000000000000000000000600082015250565b6000614373600783613b8a565b915061437e8261433d565b602082019050919050565b600060608201905061439e60008301856139bd565b6143ab60208301846139bd565b81810360408301526143bc81614366565b90509392505050565b7f576974686472617720736c61736866756e6420616d6f756e742065786365656460008201527f7320736c617368656446756e6473000000000000000000000000000000000000602082015250565b6000614421602e83613b8a565b915061442c826143c5565b604082019050919050565b6000602082019050818103600083015261445081614414565b9050919050565b600060408201905061446c6000830185613de9565b61447960208301846139bd565b9392505050565b61448981613ae8565b811461449457600080fd5b50565b6000815190506144a681614480565b92915050565b6000602082840312156144c2576144c16139e7565b5b60006144d084828501614497565b91505092915050565b7f496e76616c696420616464726573733a2063704163636f756e742063616e6e6f60008201527f7420626520746865207a65726f20616464726573730000000000000000000000602082015250565b6000614535603583613b8a565b9150614540826144d9565b604082019050919050565b6000602082019050818103600083015261456481614528565b9050919050565b600081519050919050565b600081905092915050565b600061458c8261456b565b6145968185614576565b93506145a6818560208601613b9b565b80840191505092915050565b60006145be8284614581565b915081905092915050565b7f4661696c656420746f2063616c6c206765744f776e65722066756e6374696f6e60008201527f206f662043504163636f756e7400000000000000000000000000000000000000602082015250565b6000614625602d83613b8a565b9150614630826145c9565b604082019050919050565b6000602082019050818103600083015261465481614618565b9050919050565b6000614666826139f1565b9050919050565b6146768161465b565b811461468157600080fd5b50565b6000815190506146938161466d565b92915050565b6000602082840312156146af576146ae6139e7565b5b60006146bd84828501614684565b91505092915050565b7f4f6e6c792043502773206f776e65722063616e2072657175657374207769746860008201527f64726177616c0000000000000000000000000000000000000000000000000000602082015250565b6000614722602683613b8a565b915061472d826146c6565b604082019050919050565b6000602082019050818103600083015261475181614715565b9050919050565b7f4e6f7420656e6f7567682066726f7a656e2062616c616e636520746f2077697460008201527f6864726177000000000000000000000000000000000000000000000000000000602082015250565b60006147b4602583613b8a565b91506147bf82614758565b604082019050919050565b600060208201905081810360008301526147e3816147a7565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061483157607f821691505b602082108103614844576148436147ea565b5b50919050565b7f4e6f7420656e6f7567682062616c616e636520666f7220636f6c6c617465726160008201527f6c00000000000000000000000000000000000000000000000000000000000000602082015250565b60006148a6602183613b8a565b91506148b18261484a565b604082019050919050565b600060208201905081810360008301526148d581614899565b9050919050565b7f496e76616c696420616d6f756e743a206d75737420626520677265617465722060008201527f7468616e207a65726f0000000000000000000000000000000000000000000000602082015250565b6000614938602983613b8a565b9150614943826148dc565b604082019050919050565b600060208201905081810360008301526149678161492b565b9050919050565b60006060820190506149836000830186613de9565b6149906020830185613de9565b61499d60408301846139bd565b949350505050565b7f4f6e6c792043502773206f776e65722063616e2063616e63656c20776974686460008201527f7261772072657175657374000000000000000000000000000000000000000000602082015250565b6000614a01602b83613b8a565b9150614a0c826149a5565b604082019050919050565b60006020820190508181036000830152614a30816149f4565b9050919050565b7f4e6f2070656e64696e6720776974686472617720726571756573740000000000600082015250565b6000614a6d601b83613b8a565b9150614a7882614a37565b602082019050919050565b60006020820190508181036000830152614a9c81614a60565b9050919050565b7f4172726179206c656e67746873206d757374206d617463680000000000000000600082015250565b6000614ad9601883613b8a565b9150614ae482614aa3565b602082019050919050565b60006020820190508181036000830152614b0881614acc565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f496e76616c696420616464726573733a20746f6b656e4164647265737320636160008201527f6e6e6f7420626520746865207a65726f20616464726573730000000000000000602082015250565b6000614b9a603883613b8a565b9150614ba582614b3e565b604082019050919050565b60006020820190508181036000830152614bc981614b8d565b9050919050565b7f496e76616c696420616464726573733a207461736b436f6e747261637420636160008201527f6e6e6f7420626520746865207a65726f20616464726573730000000000000000602082015250565b6000614c2c603883613b8a565b9150614c3782614bd0565b604082019050919050565b60006020820190508181036000830152614c5b81614c1f565b9050919050565b7f546f6b656e207472616e73666572206661696c65640000000000000000000000600082015250565b6000614c98601583613b8a565b9150614ca382614c62565b602082019050919050565b60006020820190508181036000830152614cc781614c8b565b9050919050565b600082825260208201905092915050565b6000819050919050565b6000614cf58383613e13565b60208301905092915050565b6000614d106020840184613a3a565b905092915050565b6000602082019050919050565b6000614d318385614cce565b9350614d3c82614cdf565b8060005b85811015614d7557614d528284614d01565b614d5c8882614ce9565b9750614d6783614d18565b925050600181019050614d40565b5085925050509392505050565b600082825260208201905092915050565b600080fd5b82818337505050565b6000614dad8385614d82565b93507f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831115614de057614ddf614d93565b5b602083029250614df1838584614d98565b82840190509392505050565b60006040820190508181036000830152614e18818688614d25565b90508181036020830152614e2d818486614da1565b905095945050505050565b7f4f6e6c792043502773206f776e65722063616e20636f6e6669726d207769746860008201527f64726177616c0000000000000000000000000000000000000000000000000000602082015250565b6000614e94602683613b8a565b9150614e9f82614e38565b604082019050919050565b60006020820190508181036000830152614ec381614e87565b9050919050565b7f57697468647261772064656c6179206e6f742070617373656400000000000000600082015250565b6000614f00601983613b8a565b9150614f0b82614eca565b602082019050919050565b60006020820190508181036000830152614f2f81614ef3565b9050919050565b7f576974686472617720616d6f756e7420657863656564732062616c616e636500600082015250565b6000614f6c601f83613b8a565b9150614f7782614f36565b602082019050919050565b60006020820190508181036000830152614f9b81614f5f565b9050919050565b7f4f6e6c792043502773206f776e65722063616e2077697468647261772074686560008201527f20636f6c6c61746572616c2066756e6473000000000000000000000000000000602082015250565b6000614ffe603183613b8a565b915061500982614fa2565b604082019050919050565b6000602082019050818103600083015261502d81614ff1565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026150c57fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82615088565b6150cf8683615088565b95508019841693508086168417925050509392505050565b60006151026150fd6150f8846139b3565b613fa5565b6139b3565b9050919050565b6000819050919050565b61511c836150e7565b61513061512882615109565b848454615095565b825550505050565b600090565b615145615138565b615150818484615113565b505050565b5b818110156151745761516960008261513d565b600181019050615156565b5050565b601f8211156151b95761518a81615063565b61519384615078565b810160208510156151a2578190505b6151b66151ae85615078565b830182615155565b50505b505050565b600082821c905092915050565b60006151dc600019846008026151be565b1980831691505092915050565b60006151f583836151cb565b9150826002028217905092915050565b61520e82613b7f565b67ffffffffffffffff81111561522757615226615034565b5b6152318254614819565b61523c828285615178565b600060209050601f83116001811461526f576000841561525d578287015190505b61526785826151e9565b8655506152cf565b601f19841661527d86615063565b60005b828110156152a557848901518255600182019150602085019450602081019050615280565b868310156152c257848901516152be601f8916826151cb565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220db90eb976a69a2029673371828e62340c82b62ef4f84d84f1679aa19386ff8b864736f6c63430008190033",
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

// BatchFrozenDeposit is a paid mutator transaction binding the contract method 0x998fc5cc.
//
// Solidity: function batchFrozenDeposit(address[] cps, uint256[] amounts) returns()
func (_EcpCollateral *EcpCollateralTransactor) BatchFrozenDeposit(opts *bind.TransactOpts, cps []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _EcpCollateral.contract.Transact(opts, "batchFrozenDeposit", cps, amounts)
}

// BatchFrozenDeposit is a paid mutator transaction binding the contract method 0x998fc5cc.
//
// Solidity: function batchFrozenDeposit(address[] cps, uint256[] amounts) returns()
func (_EcpCollateral *EcpCollateralSession) BatchFrozenDeposit(cps []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.BatchFrozenDeposit(&_EcpCollateral.TransactOpts, cps, amounts)
}

// BatchFrozenDeposit is a paid mutator transaction binding the contract method 0x998fc5cc.
//
// Solidity: function batchFrozenDeposit(address[] cps, uint256[] amounts) returns()
func (_EcpCollateral *EcpCollateralTransactorSession) BatchFrozenDeposit(cps []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _EcpCollateral.Contract.BatchFrozenDeposit(&_EcpCollateral.TransactOpts, cps, amounts)
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

// EcpCollateralBatchFrozenDepositIterator is returned from FilterBatchFrozenDeposit and is used to iterate over the raw logs and unpacked data for BatchFrozenDeposit events raised by the EcpCollateral contract.
type EcpCollateralBatchFrozenDepositIterator struct {
	Event *EcpCollateralBatchFrozenDeposit // Event containing the contract specifics and raw log

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
func (it *EcpCollateralBatchFrozenDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpCollateralBatchFrozenDeposit)
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
		it.Event = new(EcpCollateralBatchFrozenDeposit)
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
func (it *EcpCollateralBatchFrozenDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpCollateralBatchFrozenDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpCollateralBatchFrozenDeposit represents a BatchFrozenDeposit event raised by the EcpCollateral contract.
type EcpCollateralBatchFrozenDeposit struct {
	Admin   common.Address
	Cps     []common.Address
	Amounts []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBatchFrozenDeposit is a free log retrieval operation binding the contract event 0xf6bc9c339713a32d9c780cc55ab0d0b62a32de9986165175992e78e6973e5b1b.
//
// Solidity: event BatchFrozenDeposit(address indexed admin, address[] cps, uint256[] amounts)
func (_EcpCollateral *EcpCollateralFilterer) FilterBatchFrozenDeposit(opts *bind.FilterOpts, admin []common.Address) (*EcpCollateralBatchFrozenDepositIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _EcpCollateral.contract.FilterLogs(opts, "BatchFrozenDeposit", adminRule)
	if err != nil {
		return nil, err
	}
	return &EcpCollateralBatchFrozenDepositIterator{contract: _EcpCollateral.contract, event: "BatchFrozenDeposit", logs: logs, sub: sub}, nil
}

// WatchBatchFrozenDeposit is a free log subscription operation binding the contract event 0xf6bc9c339713a32d9c780cc55ab0d0b62a32de9986165175992e78e6973e5b1b.
//
// Solidity: event BatchFrozenDeposit(address indexed admin, address[] cps, uint256[] amounts)
func (_EcpCollateral *EcpCollateralFilterer) WatchBatchFrozenDeposit(opts *bind.WatchOpts, sink chan<- *EcpCollateralBatchFrozenDeposit, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _EcpCollateral.contract.WatchLogs(opts, "BatchFrozenDeposit", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpCollateralBatchFrozenDeposit)
				if err := _EcpCollateral.contract.UnpackLog(event, "BatchFrozenDeposit", log); err != nil {
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

// ParseBatchFrozenDeposit is a log parse operation binding the contract event 0xf6bc9c339713a32d9c780cc55ab0d0b62a32de9986165175992e78e6973e5b1b.
//
// Solidity: event BatchFrozenDeposit(address indexed admin, address[] cps, uint256[] amounts)
func (_EcpCollateral *EcpCollateralFilterer) ParseBatchFrozenDeposit(log types.Log) (*EcpCollateralBatchFrozenDeposit, error) {
	event := new(EcpCollateralBatchFrozenDeposit)
	if err := _EcpCollateral.contract.UnpackLog(event, "BatchFrozenDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	Cp      common.Address
	CpOwner common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawConfirmed is a free log retrieval operation binding the contract event 0xe6fc6292b9fd5ba92a34a05a92d07b066535578023d71f0c6bf83a2622ca4e54.
//
// Solidity: event WithdrawConfirmed(address indexed cp, address cpOwner, uint256 amount)
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

// WatchWithdrawConfirmed is a free log subscription operation binding the contract event 0xe6fc6292b9fd5ba92a34a05a92d07b066535578023d71f0c6bf83a2622ca4e54.
//
// Solidity: event WithdrawConfirmed(address indexed cp, address cpOwner, uint256 amount)
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

// ParseWithdrawConfirmed is a log parse operation binding the contract event 0xe6fc6292b9fd5ba92a34a05a92d07b066535578023d71f0c6bf83a2622ca4e54.
//
// Solidity: event WithdrawConfirmed(address indexed cp, address cpOwner, uint256 amount)
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
