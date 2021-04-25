// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package emp

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FixedPointUnsigned is an auto generated low-level Go binding around an user-defined struct.
type FixedPointUnsigned struct {
	RawValue *big.Int
}

// LiquidatableLiquidationData is an auto generated low-level Go binding around an user-defined struct.
type LiquidatableLiquidationData struct {
	Sponsor              common.Address
	Liquidator           common.Address
	State                uint8
	LiquidationTime      *big.Int
	TokensOutstanding    FixedPointUnsigned
	LockedCollateral     FixedPointUnsigned
	LiquidatedCollateral FixedPointUnsigned
	RawUnitCollateral    FixedPointUnsigned
	Disputer             common.Address
	SettlementPrice      FixedPointUnsigned
	FinalFee             FixedPointUnsigned
}

// LiquidatableRewardsData is an auto generated low-level Go binding around an user-defined struct.
type LiquidatableRewardsData struct {
	PayToSponsor     FixedPointUnsigned
	PayToLiquidator  FixedPointUnsigned
	PayToDisputer    FixedPointUnsigned
	PaidToSponsor    FixedPointUnsigned
	PaidToLiquidator FixedPointUnsigned
	PaidToDisputer   FixedPointUnsigned
}

// EmpABI is the input ABI used to generate the binding from.
const EmpABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"expirationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalLiveness\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collateralAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"finderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"timerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"financialProductLibraryAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"priceFeedIdentifier\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"minSponsorTokens\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"liquidationLiveness\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"collateralRequirement\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"disputeBondPercentage\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"sponsorDisputeRewardPercentage\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"disputerDisputeRewardPercentage\",\"type\":\"tuple\"}],\"internalType\":\"structLiquidatable.ConstructorParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"ContractExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"disputer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidationId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"disputeSucceeded\",\"type\":\"bool\"}],\"name\":\"DisputeSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originalExpirationTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shutdownTimestamp\",\"type\":\"uint256\"}],\"name\":\"EmergencyShutdown\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"}],\"name\":\"EndedSponsorPosition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FinalFeesPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"liquidationId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensOutstanding\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockedCollateral\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidatedCollateral\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidationTime\",\"type\":\"uint256\"}],\"name\":\"LiquidationCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"disputer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidationId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"disputeBondAmount\",\"type\":\"uint256\"}],\"name\":\"LiquidationDisputed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paidToLiquidator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paidToDisputer\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paidToSponsor\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumLiquidatable.Status\",\"name\":\"liquidationStatus\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"settlementPrice\",\"type\":\"uint256\"}],\"name\":\"LiquidationWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"}],\"name\":\"NewSponsor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"PositionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"regularFee\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"lateFee\",\"type\":\"uint256\"}],\"name\":\"RegularFeesPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"numTokensRepaid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newTokenCount\",\"type\":\"uint256\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldSponsor\",\"type\":\"address\"}],\"name\":\"RequestTransferPosition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldSponsor\",\"type\":\"address\"}],\"name\":\"RequestTransferPositionCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldSponsor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newSponsor\",\"type\":\"address\"}],\"name\":\"RequestTransferPositionExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"RequestWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"RequestWithdrawalCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"RequestWithdrawalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"collateralReturned\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokensBurned\",\"type\":\"uint256\"}],\"name\":\"SettleExpiredPosition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collateralAddress\",\"type\":\"address\"}],\"name\":\"_getSyntheticDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelTransferPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralCurrency\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralRequirement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractState\",\"outputs\":[{\"internalType\":\"enumPricelessPositionManager.ContractState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"collateralAmount\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"numTokens\",\"type\":\"tuple\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"minCollateralPerToken\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"maxCollateralPerToken\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"maxTokensToLiquidate\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"createLiquidation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidationId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"tokensLiquidated\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"finalFeeBond\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cumulativeFeeMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"collateralAmount\",\"type\":\"tuple\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"collateralAmount\",\"type\":\"tuple\"}],\"name\":\"depositTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidationId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"}],\"name\":\"dispute\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"totalPaid\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disputeBondPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disputerDisputeRewardPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyShutdown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expirationTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expire\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expiryPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"financialProductLibrary\",\"outputs\":[{\"internalType\":\"contractFinancialProductLibrary\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finder\",\"outputs\":[{\"internalType\":\"contractFinderInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"}],\"name\":\"getCollateral\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"}],\"name\":\"getLiquidations\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"enumLiquidatable.Status\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"liquidationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"tokensOutstanding\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"lockedCollateral\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"liquidatedCollateral\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"rawUnitCollateral\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"disputer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"settlementPrice\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"finalFee\",\"type\":\"tuple\"}],\"internalType\":\"structLiquidatable.LiquidationData[]\",\"name\":\"liquidationData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"getOutstandingRegularFees\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"regularFee\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"latePenalty\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"totalPaid\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gulp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidationLiveness\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"liquidations\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"enumLiquidatable.Status\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"liquidationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"tokensOutstanding\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"lockedCollateral\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"liquidatedCollateral\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"rawUnitCollateral\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"disputer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"settlementPrice\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"finalFee\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minSponsorTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payRegularFees\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pfc\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"positions\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"tokensOutstanding\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalRequestPassTimestamp\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"withdrawalRequestAmount\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"rawCollateral\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"transferPositionRequestPassTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceIdentifier\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rawLiquidationCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rawTotalPositionCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"numTokens\",\"type\":\"tuple\"}],\"name\":\"redeem\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"amountWithdrawn\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remargin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"numTokens\",\"type\":\"tuple\"}],\"name\":\"repay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestTransferPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"collateralAmount\",\"type\":\"tuple\"}],\"name\":\"requestWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"setCurrentTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settleExpired\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"amountWithdrawn\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sponsorDisputeRewardPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenCurrency\",\"outputs\":[{\"internalType\":\"contractExpandedIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPositionCollateral\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalTokensOutstanding\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSponsorAddress\",\"type\":\"address\"}],\"name\":\"transferPositionPassedRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"price\",\"type\":\"tuple\"}],\"name\":\"transformCollateralRequirement\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"price\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"requestTime\",\"type\":\"uint256\"}],\"name\":\"transformPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestTime\",\"type\":\"uint256\"}],\"name\":\"transformPriceIdentifier\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"collateralAmount\",\"type\":\"tuple\"}],\"name\":\"withdraw\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"amountWithdrawn\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidationId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"}],\"name\":\"withdrawLiquidation\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"payToSponsor\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"payToLiquidator\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"payToDisputer\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"paidToSponsor\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"paidToLiquidator\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"paidToDisputer\",\"type\":\"tuple\"}],\"internalType\":\"structLiquidatable.RewardsData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawPassedRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rawValue\",\"type\":\"uint256\"}],\"internalType\":\"structFixedPoint.Unsigned\",\"name\":\"amountWithdrawn\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalLiveness\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Emp is an auto generated Go binding around an Ethereum contract.
type Emp struct {
	EmpCaller     // Read-only binding to the contract
	EmpTransactor // Write-only binding to the contract
	EmpFilterer   // Log filterer for contract events
}

// EmpCaller is an auto generated read-only Go binding around an Ethereum contract.
type EmpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EmpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EmpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EmpSession struct {
	Contract     *Emp              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EmpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EmpCallerSession struct {
	Contract *EmpCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EmpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EmpTransactorSession struct {
	Contract     *EmpTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EmpRaw is an auto generated low-level Go binding around an Ethereum contract.
type EmpRaw struct {
	Contract *Emp // Generic contract binding to access the raw methods on
}

// EmpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EmpCallerRaw struct {
	Contract *EmpCaller // Generic read-only contract binding to access the raw methods on
}

// EmpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EmpTransactorRaw struct {
	Contract *EmpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEmp creates a new instance of Emp, bound to a specific deployed contract.
func NewEmp(address common.Address, backend bind.ContractBackend) (*Emp, error) {
	contract, err := bindEmp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Emp{EmpCaller: EmpCaller{contract: contract}, EmpTransactor: EmpTransactor{contract: contract}, EmpFilterer: EmpFilterer{contract: contract}}, nil
}

// NewEmpCaller creates a new read-only instance of Emp, bound to a specific deployed contract.
func NewEmpCaller(address common.Address, caller bind.ContractCaller) (*EmpCaller, error) {
	contract, err := bindEmp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EmpCaller{contract: contract}, nil
}

// NewEmpTransactor creates a new write-only instance of Emp, bound to a specific deployed contract.
func NewEmpTransactor(address common.Address, transactor bind.ContractTransactor) (*EmpTransactor, error) {
	contract, err := bindEmp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EmpTransactor{contract: contract}, nil
}

// NewEmpFilterer creates a new log filterer instance of Emp, bound to a specific deployed contract.
func NewEmpFilterer(address common.Address, filterer bind.ContractFilterer) (*EmpFilterer, error) {
	contract, err := bindEmp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EmpFilterer{contract: contract}, nil
}

// bindEmp binds a generic wrapper to an already deployed contract.
func bindEmp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EmpABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Emp *EmpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Emp.Contract.EmpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Emp *EmpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Emp.Contract.EmpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Emp *EmpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Emp.Contract.EmpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Emp *EmpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Emp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Emp *EmpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Emp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Emp *EmpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Emp.Contract.contract.Transact(opts, method, params...)
}

// GetSyntheticDecimals is a free data retrieval call binding the contract method 0x4ead6e51.
//
// Solidity: function _getSyntheticDecimals(address _collateralAddress) view returns(uint8 decimals)
func (_Emp *EmpCaller) GetSyntheticDecimals(opts *bind.CallOpts, _collateralAddress common.Address) (uint8, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "_getSyntheticDecimals", _collateralAddress)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetSyntheticDecimals is a free data retrieval call binding the contract method 0x4ead6e51.
//
// Solidity: function _getSyntheticDecimals(address _collateralAddress) view returns(uint8 decimals)
func (_Emp *EmpSession) GetSyntheticDecimals(_collateralAddress common.Address) (uint8, error) {
	return _Emp.Contract.GetSyntheticDecimals(&_Emp.CallOpts, _collateralAddress)
}

// GetSyntheticDecimals is a free data retrieval call binding the contract method 0x4ead6e51.
//
// Solidity: function _getSyntheticDecimals(address _collateralAddress) view returns(uint8 decimals)
func (_Emp *EmpCallerSession) GetSyntheticDecimals(_collateralAddress common.Address) (uint8, error) {
	return _Emp.Contract.GetSyntheticDecimals(&_Emp.CallOpts, _collateralAddress)
}

// CollateralCurrency is a free data retrieval call binding the contract method 0x0de15fd9.
//
// Solidity: function collateralCurrency() view returns(address)
func (_Emp *EmpCaller) CollateralCurrency(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "collateralCurrency")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralCurrency is a free data retrieval call binding the contract method 0x0de15fd9.
//
// Solidity: function collateralCurrency() view returns(address)
func (_Emp *EmpSession) CollateralCurrency() (common.Address, error) {
	return _Emp.Contract.CollateralCurrency(&_Emp.CallOpts)
}

// CollateralCurrency is a free data retrieval call binding the contract method 0x0de15fd9.
//
// Solidity: function collateralCurrency() view returns(address)
func (_Emp *EmpCallerSession) CollateralCurrency() (common.Address, error) {
	return _Emp.Contract.CollateralCurrency(&_Emp.CallOpts)
}

// CollateralRequirement is a free data retrieval call binding the contract method 0x48e30c3f.
//
// Solidity: function collateralRequirement() view returns(uint256 rawValue)
func (_Emp *EmpCaller) CollateralRequirement(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "collateralRequirement")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollateralRequirement is a free data retrieval call binding the contract method 0x48e30c3f.
//
// Solidity: function collateralRequirement() view returns(uint256 rawValue)
func (_Emp *EmpSession) CollateralRequirement() (*big.Int, error) {
	return _Emp.Contract.CollateralRequirement(&_Emp.CallOpts)
}

// CollateralRequirement is a free data retrieval call binding the contract method 0x48e30c3f.
//
// Solidity: function collateralRequirement() view returns(uint256 rawValue)
func (_Emp *EmpCallerSession) CollateralRequirement() (*big.Int, error) {
	return _Emp.Contract.CollateralRequirement(&_Emp.CallOpts)
}

// ContractState is a free data retrieval call binding the contract method 0x85209ee0.
//
// Solidity: function contractState() view returns(uint8)
func (_Emp *EmpCaller) ContractState(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "contractState")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ContractState is a free data retrieval call binding the contract method 0x85209ee0.
//
// Solidity: function contractState() view returns(uint8)
func (_Emp *EmpSession) ContractState() (uint8, error) {
	return _Emp.Contract.ContractState(&_Emp.CallOpts)
}

// ContractState is a free data retrieval call binding the contract method 0x85209ee0.
//
// Solidity: function contractState() view returns(uint8)
func (_Emp *EmpCallerSession) ContractState() (uint8, error) {
	return _Emp.Contract.ContractState(&_Emp.CallOpts)
}

// CumulativeFeeMultiplier is a free data retrieval call binding the contract method 0xdd0eef3d.
//
// Solidity: function cumulativeFeeMultiplier() view returns(uint256 rawValue)
func (_Emp *EmpCaller) CumulativeFeeMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "cumulativeFeeMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeFeeMultiplier is a free data retrieval call binding the contract method 0xdd0eef3d.
//
// Solidity: function cumulativeFeeMultiplier() view returns(uint256 rawValue)
func (_Emp *EmpSession) CumulativeFeeMultiplier() (*big.Int, error) {
	return _Emp.Contract.CumulativeFeeMultiplier(&_Emp.CallOpts)
}

// CumulativeFeeMultiplier is a free data retrieval call binding the contract method 0xdd0eef3d.
//
// Solidity: function cumulativeFeeMultiplier() view returns(uint256 rawValue)
func (_Emp *EmpCallerSession) CumulativeFeeMultiplier() (*big.Int, error) {
	return _Emp.Contract.CumulativeFeeMultiplier(&_Emp.CallOpts)
}

// DisputeBondPercentage is a free data retrieval call binding the contract method 0x2e154f2e.
//
// Solidity: function disputeBondPercentage() view returns(uint256 rawValue)
func (_Emp *EmpCaller) DisputeBondPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "disputeBondPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DisputeBondPercentage is a free data retrieval call binding the contract method 0x2e154f2e.
//
// Solidity: function disputeBondPercentage() view returns(uint256 rawValue)
func (_Emp *EmpSession) DisputeBondPercentage() (*big.Int, error) {
	return _Emp.Contract.DisputeBondPercentage(&_Emp.CallOpts)
}

// DisputeBondPercentage is a free data retrieval call binding the contract method 0x2e154f2e.
//
// Solidity: function disputeBondPercentage() view returns(uint256 rawValue)
func (_Emp *EmpCallerSession) DisputeBondPercentage() (*big.Int, error) {
	return _Emp.Contract.DisputeBondPercentage(&_Emp.CallOpts)
}

// DisputerDisputeRewardPercentage is a free data retrieval call binding the contract method 0x7e398c22.
//
// Solidity: function disputerDisputeRewardPercentage() view returns(uint256 rawValue)
func (_Emp *EmpCaller) DisputerDisputeRewardPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "disputerDisputeRewardPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DisputerDisputeRewardPercentage is a free data retrieval call binding the contract method 0x7e398c22.
//
// Solidity: function disputerDisputeRewardPercentage() view returns(uint256 rawValue)
func (_Emp *EmpSession) DisputerDisputeRewardPercentage() (*big.Int, error) {
	return _Emp.Contract.DisputerDisputeRewardPercentage(&_Emp.CallOpts)
}

// DisputerDisputeRewardPercentage is a free data retrieval call binding the contract method 0x7e398c22.
//
// Solidity: function disputerDisputeRewardPercentage() view returns(uint256 rawValue)
func (_Emp *EmpCallerSession) DisputerDisputeRewardPercentage() (*big.Int, error) {
	return _Emp.Contract.DisputerDisputeRewardPercentage(&_Emp.CallOpts)
}

// ExpirationTimestamp is a free data retrieval call binding the contract method 0x9f43ddd2.
//
// Solidity: function expirationTimestamp() view returns(uint256)
func (_Emp *EmpCaller) ExpirationTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "expirationTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpirationTimestamp is a free data retrieval call binding the contract method 0x9f43ddd2.
//
// Solidity: function expirationTimestamp() view returns(uint256)
func (_Emp *EmpSession) ExpirationTimestamp() (*big.Int, error) {
	return _Emp.Contract.ExpirationTimestamp(&_Emp.CallOpts)
}

// ExpirationTimestamp is a free data retrieval call binding the contract method 0x9f43ddd2.
//
// Solidity: function expirationTimestamp() view returns(uint256)
func (_Emp *EmpCallerSession) ExpirationTimestamp() (*big.Int, error) {
	return _Emp.Contract.ExpirationTimestamp(&_Emp.CallOpts)
}

// ExpiryPrice is a free data retrieval call binding the contract method 0xedfa9a9b.
//
// Solidity: function expiryPrice() view returns(uint256 rawValue)
func (_Emp *EmpCaller) ExpiryPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "expiryPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpiryPrice is a free data retrieval call binding the contract method 0xedfa9a9b.
//
// Solidity: function expiryPrice() view returns(uint256 rawValue)
func (_Emp *EmpSession) ExpiryPrice() (*big.Int, error) {
	return _Emp.Contract.ExpiryPrice(&_Emp.CallOpts)
}

// ExpiryPrice is a free data retrieval call binding the contract method 0xedfa9a9b.
//
// Solidity: function expiryPrice() view returns(uint256 rawValue)
func (_Emp *EmpCallerSession) ExpiryPrice() (*big.Int, error) {
	return _Emp.Contract.ExpiryPrice(&_Emp.CallOpts)
}

// FinancialProductLibrary is a free data retrieval call binding the contract method 0x9375f0e9.
//
// Solidity: function financialProductLibrary() view returns(address)
func (_Emp *EmpCaller) FinancialProductLibrary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "financialProductLibrary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FinancialProductLibrary is a free data retrieval call binding the contract method 0x9375f0e9.
//
// Solidity: function financialProductLibrary() view returns(address)
func (_Emp *EmpSession) FinancialProductLibrary() (common.Address, error) {
	return _Emp.Contract.FinancialProductLibrary(&_Emp.CallOpts)
}

// FinancialProductLibrary is a free data retrieval call binding the contract method 0x9375f0e9.
//
// Solidity: function financialProductLibrary() view returns(address)
func (_Emp *EmpCallerSession) FinancialProductLibrary() (common.Address, error) {
	return _Emp.Contract.FinancialProductLibrary(&_Emp.CallOpts)
}

// Finder is a free data retrieval call binding the contract method 0xb9a3c84c.
//
// Solidity: function finder() view returns(address)
func (_Emp *EmpCaller) Finder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "finder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Finder is a free data retrieval call binding the contract method 0xb9a3c84c.
//
// Solidity: function finder() view returns(address)
func (_Emp *EmpSession) Finder() (common.Address, error) {
	return _Emp.Contract.Finder(&_Emp.CallOpts)
}

// Finder is a free data retrieval call binding the contract method 0xb9a3c84c.
//
// Solidity: function finder() view returns(address)
func (_Emp *EmpCallerSession) Finder() (common.Address, error) {
	return _Emp.Contract.Finder(&_Emp.CallOpts)
}

// GetCollateral is a free data retrieval call binding the contract method 0x9b56d6c9.
//
// Solidity: function getCollateral(address sponsor) view returns((uint256))
func (_Emp *EmpCaller) GetCollateral(opts *bind.CallOpts, sponsor common.Address) (FixedPointUnsigned, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "getCollateral", sponsor)

	if err != nil {
		return *new(FixedPointUnsigned), err
	}

	out0 := *abi.ConvertType(out[0], new(FixedPointUnsigned)).(*FixedPointUnsigned)

	return out0, err

}

// GetCollateral is a free data retrieval call binding the contract method 0x9b56d6c9.
//
// Solidity: function getCollateral(address sponsor) view returns((uint256))
func (_Emp *EmpSession) GetCollateral(sponsor common.Address) (FixedPointUnsigned, error) {
	return _Emp.Contract.GetCollateral(&_Emp.CallOpts, sponsor)
}

// GetCollateral is a free data retrieval call binding the contract method 0x9b56d6c9.
//
// Solidity: function getCollateral(address sponsor) view returns((uint256))
func (_Emp *EmpCallerSession) GetCollateral(sponsor common.Address) (FixedPointUnsigned, error) {
	return _Emp.Contract.GetCollateral(&_Emp.CallOpts, sponsor)
}

// GetCurrentTime is a free data retrieval call binding the contract method 0x29cb924d.
//
// Solidity: function getCurrentTime() view returns(uint256)
func (_Emp *EmpCaller) GetCurrentTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "getCurrentTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentTime is a free data retrieval call binding the contract method 0x29cb924d.
//
// Solidity: function getCurrentTime() view returns(uint256)
func (_Emp *EmpSession) GetCurrentTime() (*big.Int, error) {
	return _Emp.Contract.GetCurrentTime(&_Emp.CallOpts)
}

// GetCurrentTime is a free data retrieval call binding the contract method 0x29cb924d.
//
// Solidity: function getCurrentTime() view returns(uint256)
func (_Emp *EmpCallerSession) GetCurrentTime() (*big.Int, error) {
	return _Emp.Contract.GetCurrentTime(&_Emp.CallOpts)
}

// GetLiquidations is a free data retrieval call binding the contract method 0xa1c4d1e7.
//
// Solidity: function getLiquidations(address sponsor) view returns((address,address,uint8,uint256,(uint256),(uint256),(uint256),(uint256),address,(uint256),(uint256))[] liquidationData)
func (_Emp *EmpCaller) GetLiquidations(opts *bind.CallOpts, sponsor common.Address) ([]LiquidatableLiquidationData, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "getLiquidations", sponsor)

	if err != nil {
		return *new([]LiquidatableLiquidationData), err
	}

	out0 := *abi.ConvertType(out[0], new([]LiquidatableLiquidationData)).(*[]LiquidatableLiquidationData)

	return out0, err

}

// GetLiquidations is a free data retrieval call binding the contract method 0xa1c4d1e7.
//
// Solidity: function getLiquidations(address sponsor) view returns((address,address,uint8,uint256,(uint256),(uint256),(uint256),(uint256),address,(uint256),(uint256))[] liquidationData)
func (_Emp *EmpSession) GetLiquidations(sponsor common.Address) ([]LiquidatableLiquidationData, error) {
	return _Emp.Contract.GetLiquidations(&_Emp.CallOpts, sponsor)
}

// GetLiquidations is a free data retrieval call binding the contract method 0xa1c4d1e7.
//
// Solidity: function getLiquidations(address sponsor) view returns((address,address,uint8,uint256,(uint256),(uint256),(uint256),(uint256),address,(uint256),(uint256))[] liquidationData)
func (_Emp *EmpCallerSession) GetLiquidations(sponsor common.Address) ([]LiquidatableLiquidationData, error) {
	return _Emp.Contract.GetLiquidations(&_Emp.CallOpts, sponsor)
}

// GetOutstandingRegularFees is a free data retrieval call binding the contract method 0x9e4efaa0.
//
// Solidity: function getOutstandingRegularFees(uint256 time) view returns((uint256) regularFee, (uint256) latePenalty, (uint256) totalPaid)
func (_Emp *EmpCaller) GetOutstandingRegularFees(opts *bind.CallOpts, time *big.Int) (struct {
	RegularFee  FixedPointUnsigned
	LatePenalty FixedPointUnsigned
	TotalPaid   FixedPointUnsigned
}, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "getOutstandingRegularFees", time)

	outstruct := new(struct {
		RegularFee  FixedPointUnsigned
		LatePenalty FixedPointUnsigned
		TotalPaid   FixedPointUnsigned
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RegularFee = *abi.ConvertType(out[0], new(FixedPointUnsigned)).(*FixedPointUnsigned)
	outstruct.LatePenalty = *abi.ConvertType(out[1], new(FixedPointUnsigned)).(*FixedPointUnsigned)
	outstruct.TotalPaid = *abi.ConvertType(out[2], new(FixedPointUnsigned)).(*FixedPointUnsigned)

	return *outstruct, err

}

// GetOutstandingRegularFees is a free data retrieval call binding the contract method 0x9e4efaa0.
//
// Solidity: function getOutstandingRegularFees(uint256 time) view returns((uint256) regularFee, (uint256) latePenalty, (uint256) totalPaid)
func (_Emp *EmpSession) GetOutstandingRegularFees(time *big.Int) (struct {
	RegularFee  FixedPointUnsigned
	LatePenalty FixedPointUnsigned
	TotalPaid   FixedPointUnsigned
}, error) {
	return _Emp.Contract.GetOutstandingRegularFees(&_Emp.CallOpts, time)
}

// GetOutstandingRegularFees is a free data retrieval call binding the contract method 0x9e4efaa0.
//
// Solidity: function getOutstandingRegularFees(uint256 time) view returns((uint256) regularFee, (uint256) latePenalty, (uint256) totalPaid)
func (_Emp *EmpCallerSession) GetOutstandingRegularFees(time *big.Int) (struct {
	RegularFee  FixedPointUnsigned
	LatePenalty FixedPointUnsigned
	TotalPaid   FixedPointUnsigned
}, error) {
	return _Emp.Contract.GetOutstandingRegularFees(&_Emp.CallOpts, time)
}

// LiquidationLiveness is a free data retrieval call binding the contract method 0x2d5436cf.
//
// Solidity: function liquidationLiveness() view returns(uint256)
func (_Emp *EmpCaller) LiquidationLiveness(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "liquidationLiveness")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationLiveness is a free data retrieval call binding the contract method 0x2d5436cf.
//
// Solidity: function liquidationLiveness() view returns(uint256)
func (_Emp *EmpSession) LiquidationLiveness() (*big.Int, error) {
	return _Emp.Contract.LiquidationLiveness(&_Emp.CallOpts)
}

// LiquidationLiveness is a free data retrieval call binding the contract method 0x2d5436cf.
//
// Solidity: function liquidationLiveness() view returns(uint256)
func (_Emp *EmpCallerSession) LiquidationLiveness() (*big.Int, error) {
	return _Emp.Contract.LiquidationLiveness(&_Emp.CallOpts)
}

// Liquidations is a free data retrieval call binding the contract method 0x4f8c4847.
//
// Solidity: function liquidations(address , uint256 ) view returns(address sponsor, address liquidator, uint8 state, uint256 liquidationTime, (uint256) tokensOutstanding, (uint256) lockedCollateral, (uint256) liquidatedCollateral, (uint256) rawUnitCollateral, address disputer, (uint256) settlementPrice, (uint256) finalFee)
func (_Emp *EmpCaller) Liquidations(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Sponsor              common.Address
	Liquidator           common.Address
	State                uint8
	LiquidationTime      *big.Int
	TokensOutstanding    FixedPointUnsigned
	LockedCollateral     FixedPointUnsigned
	LiquidatedCollateral FixedPointUnsigned
	RawUnitCollateral    FixedPointUnsigned
	Disputer             common.Address
	SettlementPrice      FixedPointUnsigned
	FinalFee             FixedPointUnsigned
}, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "liquidations", arg0, arg1)

	outstruct := new(struct {
		Sponsor              common.Address
		Liquidator           common.Address
		State                uint8
		LiquidationTime      *big.Int
		TokensOutstanding    FixedPointUnsigned
		LockedCollateral     FixedPointUnsigned
		LiquidatedCollateral FixedPointUnsigned
		RawUnitCollateral    FixedPointUnsigned
		Disputer             common.Address
		SettlementPrice      FixedPointUnsigned
		FinalFee             FixedPointUnsigned
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Sponsor = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Liquidator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.State = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.LiquidationTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TokensOutstanding = *abi.ConvertType(out[4], new(FixedPointUnsigned)).(*FixedPointUnsigned)
	outstruct.LockedCollateral = *abi.ConvertType(out[5], new(FixedPointUnsigned)).(*FixedPointUnsigned)
	outstruct.LiquidatedCollateral = *abi.ConvertType(out[6], new(FixedPointUnsigned)).(*FixedPointUnsigned)
	outstruct.RawUnitCollateral = *abi.ConvertType(out[7], new(FixedPointUnsigned)).(*FixedPointUnsigned)
	outstruct.Disputer = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)
	outstruct.SettlementPrice = *abi.ConvertType(out[9], new(FixedPointUnsigned)).(*FixedPointUnsigned)
	outstruct.FinalFee = *abi.ConvertType(out[10], new(FixedPointUnsigned)).(*FixedPointUnsigned)

	return *outstruct, err

}

// Liquidations is a free data retrieval call binding the contract method 0x4f8c4847.
//
// Solidity: function liquidations(address , uint256 ) view returns(address sponsor, address liquidator, uint8 state, uint256 liquidationTime, (uint256) tokensOutstanding, (uint256) lockedCollateral, (uint256) liquidatedCollateral, (uint256) rawUnitCollateral, address disputer, (uint256) settlementPrice, (uint256) finalFee)
func (_Emp *EmpSession) Liquidations(arg0 common.Address, arg1 *big.Int) (struct {
	Sponsor              common.Address
	Liquidator           common.Address
	State                uint8
	LiquidationTime      *big.Int
	TokensOutstanding    FixedPointUnsigned
	LockedCollateral     FixedPointUnsigned
	LiquidatedCollateral FixedPointUnsigned
	RawUnitCollateral    FixedPointUnsigned
	Disputer             common.Address
	SettlementPrice      FixedPointUnsigned
	FinalFee             FixedPointUnsigned
}, error) {
	return _Emp.Contract.Liquidations(&_Emp.CallOpts, arg0, arg1)
}

// Liquidations is a free data retrieval call binding the contract method 0x4f8c4847.
//
// Solidity: function liquidations(address , uint256 ) view returns(address sponsor, address liquidator, uint8 state, uint256 liquidationTime, (uint256) tokensOutstanding, (uint256) lockedCollateral, (uint256) liquidatedCollateral, (uint256) rawUnitCollateral, address disputer, (uint256) settlementPrice, (uint256) finalFee)
func (_Emp *EmpCallerSession) Liquidations(arg0 common.Address, arg1 *big.Int) (struct {
	Sponsor              common.Address
	Liquidator           common.Address
	State                uint8
	LiquidationTime      *big.Int
	TokensOutstanding    FixedPointUnsigned
	LockedCollateral     FixedPointUnsigned
	LiquidatedCollateral FixedPointUnsigned
	RawUnitCollateral    FixedPointUnsigned
	Disputer             common.Address
	SettlementPrice      FixedPointUnsigned
	FinalFee             FixedPointUnsigned
}, error) {
	return _Emp.Contract.Liquidations(&_Emp.CallOpts, arg0, arg1)
}

// MinSponsorTokens is a free data retrieval call binding the contract method 0x92120aec.
//
// Solidity: function minSponsorTokens() view returns(uint256 rawValue)
func (_Emp *EmpCaller) MinSponsorTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "minSponsorTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSponsorTokens is a free data retrieval call binding the contract method 0x92120aec.
//
// Solidity: function minSponsorTokens() view returns(uint256 rawValue)
func (_Emp *EmpSession) MinSponsorTokens() (*big.Int, error) {
	return _Emp.Contract.MinSponsorTokens(&_Emp.CallOpts)
}

// MinSponsorTokens is a free data retrieval call binding the contract method 0x92120aec.
//
// Solidity: function minSponsorTokens() view returns(uint256 rawValue)
func (_Emp *EmpCallerSession) MinSponsorTokens() (*big.Int, error) {
	return _Emp.Contract.MinSponsorTokens(&_Emp.CallOpts)
}

// Pfc is a free data retrieval call binding the contract method 0x81a10ae1.
//
// Solidity: function pfc() view returns((uint256))
func (_Emp *EmpCaller) Pfc(opts *bind.CallOpts) (FixedPointUnsigned, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "pfc")

	if err != nil {
		return *new(FixedPointUnsigned), err
	}

	out0 := *abi.ConvertType(out[0], new(FixedPointUnsigned)).(*FixedPointUnsigned)

	return out0, err

}

// Pfc is a free data retrieval call binding the contract method 0x81a10ae1.
//
// Solidity: function pfc() view returns((uint256))
func (_Emp *EmpSession) Pfc() (FixedPointUnsigned, error) {
	return _Emp.Contract.Pfc(&_Emp.CallOpts)
}

// Pfc is a free data retrieval call binding the contract method 0x81a10ae1.
//
// Solidity: function pfc() view returns((uint256))
func (_Emp *EmpCallerSession) Pfc() (FixedPointUnsigned, error) {
	return _Emp.Contract.Pfc(&_Emp.CallOpts)
}

// Positions is a free data retrieval call binding the contract method 0x55f57510.
//
// Solidity: function positions(address ) view returns((uint256) tokensOutstanding, uint256 withdrawalRequestPassTimestamp, (uint256) withdrawalRequestAmount, (uint256) rawCollateral, uint256 transferPositionRequestPassTimestamp)
func (_Emp *EmpCaller) Positions(opts *bind.CallOpts, arg0 common.Address) (struct {
	TokensOutstanding                    FixedPointUnsigned
	WithdrawalRequestPassTimestamp       *big.Int
	WithdrawalRequestAmount              FixedPointUnsigned
	RawCollateral                        FixedPointUnsigned
	TransferPositionRequestPassTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "positions", arg0)

	outstruct := new(struct {
		TokensOutstanding                    FixedPointUnsigned
		WithdrawalRequestPassTimestamp       *big.Int
		WithdrawalRequestAmount              FixedPointUnsigned
		RawCollateral                        FixedPointUnsigned
		TransferPositionRequestPassTimestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokensOutstanding = *abi.ConvertType(out[0], new(FixedPointUnsigned)).(*FixedPointUnsigned)
	outstruct.WithdrawalRequestPassTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.WithdrawalRequestAmount = *abi.ConvertType(out[2], new(FixedPointUnsigned)).(*FixedPointUnsigned)
	outstruct.RawCollateral = *abi.ConvertType(out[3], new(FixedPointUnsigned)).(*FixedPointUnsigned)
	outstruct.TransferPositionRequestPassTimestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Positions is a free data retrieval call binding the contract method 0x55f57510.
//
// Solidity: function positions(address ) view returns((uint256) tokensOutstanding, uint256 withdrawalRequestPassTimestamp, (uint256) withdrawalRequestAmount, (uint256) rawCollateral, uint256 transferPositionRequestPassTimestamp)
func (_Emp *EmpSession) Positions(arg0 common.Address) (struct {
	TokensOutstanding                    FixedPointUnsigned
	WithdrawalRequestPassTimestamp       *big.Int
	WithdrawalRequestAmount              FixedPointUnsigned
	RawCollateral                        FixedPointUnsigned
	TransferPositionRequestPassTimestamp *big.Int
}, error) {
	return _Emp.Contract.Positions(&_Emp.CallOpts, arg0)
}

// Positions is a free data retrieval call binding the contract method 0x55f57510.
//
// Solidity: function positions(address ) view returns((uint256) tokensOutstanding, uint256 withdrawalRequestPassTimestamp, (uint256) withdrawalRequestAmount, (uint256) rawCollateral, uint256 transferPositionRequestPassTimestamp)
func (_Emp *EmpCallerSession) Positions(arg0 common.Address) (struct {
	TokensOutstanding                    FixedPointUnsigned
	WithdrawalRequestPassTimestamp       *big.Int
	WithdrawalRequestAmount              FixedPointUnsigned
	RawCollateral                        FixedPointUnsigned
	TransferPositionRequestPassTimestamp *big.Int
}, error) {
	return _Emp.Contract.Positions(&_Emp.CallOpts, arg0)
}

// PriceIdentifier is a free data retrieval call binding the contract method 0x97523661.
//
// Solidity: function priceIdentifier() view returns(bytes32)
func (_Emp *EmpCaller) PriceIdentifier(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "priceIdentifier")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PriceIdentifier is a free data retrieval call binding the contract method 0x97523661.
//
// Solidity: function priceIdentifier() view returns(bytes32)
func (_Emp *EmpSession) PriceIdentifier() ([32]byte, error) {
	return _Emp.Contract.PriceIdentifier(&_Emp.CallOpts)
}

// PriceIdentifier is a free data retrieval call binding the contract method 0x97523661.
//
// Solidity: function priceIdentifier() view returns(bytes32)
func (_Emp *EmpCallerSession) PriceIdentifier() ([32]byte, error) {
	return _Emp.Contract.PriceIdentifier(&_Emp.CallOpts)
}

// RawLiquidationCollateral is a free data retrieval call binding the contract method 0x50f49846.
//
// Solidity: function rawLiquidationCollateral() view returns(uint256 rawValue)
func (_Emp *EmpCaller) RawLiquidationCollateral(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "rawLiquidationCollateral")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RawLiquidationCollateral is a free data retrieval call binding the contract method 0x50f49846.
//
// Solidity: function rawLiquidationCollateral() view returns(uint256 rawValue)
func (_Emp *EmpSession) RawLiquidationCollateral() (*big.Int, error) {
	return _Emp.Contract.RawLiquidationCollateral(&_Emp.CallOpts)
}

// RawLiquidationCollateral is a free data retrieval call binding the contract method 0x50f49846.
//
// Solidity: function rawLiquidationCollateral() view returns(uint256 rawValue)
func (_Emp *EmpCallerSession) RawLiquidationCollateral() (*big.Int, error) {
	return _Emp.Contract.RawLiquidationCollateral(&_Emp.CallOpts)
}

// RawTotalPositionCollateral is a free data retrieval call binding the contract method 0x8c382eb2.
//
// Solidity: function rawTotalPositionCollateral() view returns(uint256 rawValue)
func (_Emp *EmpCaller) RawTotalPositionCollateral(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "rawTotalPositionCollateral")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RawTotalPositionCollateral is a free data retrieval call binding the contract method 0x8c382eb2.
//
// Solidity: function rawTotalPositionCollateral() view returns(uint256 rawValue)
func (_Emp *EmpSession) RawTotalPositionCollateral() (*big.Int, error) {
	return _Emp.Contract.RawTotalPositionCollateral(&_Emp.CallOpts)
}

// RawTotalPositionCollateral is a free data retrieval call binding the contract method 0x8c382eb2.
//
// Solidity: function rawTotalPositionCollateral() view returns(uint256 rawValue)
func (_Emp *EmpCallerSession) RawTotalPositionCollateral() (*big.Int, error) {
	return _Emp.Contract.RawTotalPositionCollateral(&_Emp.CallOpts)
}

// SponsorDisputeRewardPercentage is a free data retrieval call binding the contract method 0x081b314e.
//
// Solidity: function sponsorDisputeRewardPercentage() view returns(uint256 rawValue)
func (_Emp *EmpCaller) SponsorDisputeRewardPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "sponsorDisputeRewardPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SponsorDisputeRewardPercentage is a free data retrieval call binding the contract method 0x081b314e.
//
// Solidity: function sponsorDisputeRewardPercentage() view returns(uint256 rawValue)
func (_Emp *EmpSession) SponsorDisputeRewardPercentage() (*big.Int, error) {
	return _Emp.Contract.SponsorDisputeRewardPercentage(&_Emp.CallOpts)
}

// SponsorDisputeRewardPercentage is a free data retrieval call binding the contract method 0x081b314e.
//
// Solidity: function sponsorDisputeRewardPercentage() view returns(uint256 rawValue)
func (_Emp *EmpCallerSession) SponsorDisputeRewardPercentage() (*big.Int, error) {
	return _Emp.Contract.SponsorDisputeRewardPercentage(&_Emp.CallOpts)
}

// TimerAddress is a free data retrieval call binding the contract method 0x1c39c38d.
//
// Solidity: function timerAddress() view returns(address)
func (_Emp *EmpCaller) TimerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "timerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TimerAddress is a free data retrieval call binding the contract method 0x1c39c38d.
//
// Solidity: function timerAddress() view returns(address)
func (_Emp *EmpSession) TimerAddress() (common.Address, error) {
	return _Emp.Contract.TimerAddress(&_Emp.CallOpts)
}

// TimerAddress is a free data retrieval call binding the contract method 0x1c39c38d.
//
// Solidity: function timerAddress() view returns(address)
func (_Emp *EmpCallerSession) TimerAddress() (common.Address, error) {
	return _Emp.Contract.TimerAddress(&_Emp.CallOpts)
}

// TokenCurrency is a free data retrieval call binding the contract method 0x7048594b.
//
// Solidity: function tokenCurrency() view returns(address)
func (_Emp *EmpCaller) TokenCurrency(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "tokenCurrency")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenCurrency is a free data retrieval call binding the contract method 0x7048594b.
//
// Solidity: function tokenCurrency() view returns(address)
func (_Emp *EmpSession) TokenCurrency() (common.Address, error) {
	return _Emp.Contract.TokenCurrency(&_Emp.CallOpts)
}

// TokenCurrency is a free data retrieval call binding the contract method 0x7048594b.
//
// Solidity: function tokenCurrency() view returns(address)
func (_Emp *EmpCallerSession) TokenCurrency() (common.Address, error) {
	return _Emp.Contract.TokenCurrency(&_Emp.CallOpts)
}

// TotalPositionCollateral is a free data retrieval call binding the contract method 0x43e4771b.
//
// Solidity: function totalPositionCollateral() view returns((uint256))
func (_Emp *EmpCaller) TotalPositionCollateral(opts *bind.CallOpts) (FixedPointUnsigned, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "totalPositionCollateral")

	if err != nil {
		return *new(FixedPointUnsigned), err
	}

	out0 := *abi.ConvertType(out[0], new(FixedPointUnsigned)).(*FixedPointUnsigned)

	return out0, err

}

// TotalPositionCollateral is a free data retrieval call binding the contract method 0x43e4771b.
//
// Solidity: function totalPositionCollateral() view returns((uint256))
func (_Emp *EmpSession) TotalPositionCollateral() (FixedPointUnsigned, error) {
	return _Emp.Contract.TotalPositionCollateral(&_Emp.CallOpts)
}

// TotalPositionCollateral is a free data retrieval call binding the contract method 0x43e4771b.
//
// Solidity: function totalPositionCollateral() view returns((uint256))
func (_Emp *EmpCallerSession) TotalPositionCollateral() (FixedPointUnsigned, error) {
	return _Emp.Contract.TotalPositionCollateral(&_Emp.CallOpts)
}

// TotalTokensOutstanding is a free data retrieval call binding the contract method 0x0c9229ca.
//
// Solidity: function totalTokensOutstanding() view returns(uint256 rawValue)
func (_Emp *EmpCaller) TotalTokensOutstanding(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "totalTokensOutstanding")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalTokensOutstanding is a free data retrieval call binding the contract method 0x0c9229ca.
//
// Solidity: function totalTokensOutstanding() view returns(uint256 rawValue)
func (_Emp *EmpSession) TotalTokensOutstanding() (*big.Int, error) {
	return _Emp.Contract.TotalTokensOutstanding(&_Emp.CallOpts)
}

// TotalTokensOutstanding is a free data retrieval call binding the contract method 0x0c9229ca.
//
// Solidity: function totalTokensOutstanding() view returns(uint256 rawValue)
func (_Emp *EmpCallerSession) TotalTokensOutstanding() (*big.Int, error) {
	return _Emp.Contract.TotalTokensOutstanding(&_Emp.CallOpts)
}

// TransformCollateralRequirement is a free data retrieval call binding the contract method 0x197f7848.
//
// Solidity: function transformCollateralRequirement((uint256) price) view returns((uint256))
func (_Emp *EmpCaller) TransformCollateralRequirement(opts *bind.CallOpts, price FixedPointUnsigned) (FixedPointUnsigned, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "transformCollateralRequirement", price)

	if err != nil {
		return *new(FixedPointUnsigned), err
	}

	out0 := *abi.ConvertType(out[0], new(FixedPointUnsigned)).(*FixedPointUnsigned)

	return out0, err

}

// TransformCollateralRequirement is a free data retrieval call binding the contract method 0x197f7848.
//
// Solidity: function transformCollateralRequirement((uint256) price) view returns((uint256))
func (_Emp *EmpSession) TransformCollateralRequirement(price FixedPointUnsigned) (FixedPointUnsigned, error) {
	return _Emp.Contract.TransformCollateralRequirement(&_Emp.CallOpts, price)
}

// TransformCollateralRequirement is a free data retrieval call binding the contract method 0x197f7848.
//
// Solidity: function transformCollateralRequirement((uint256) price) view returns((uint256))
func (_Emp *EmpCallerSession) TransformCollateralRequirement(price FixedPointUnsigned) (FixedPointUnsigned, error) {
	return _Emp.Contract.TransformCollateralRequirement(&_Emp.CallOpts, price)
}

// TransformPrice is a free data retrieval call binding the contract method 0x0ff49b90.
//
// Solidity: function transformPrice((uint256) price, uint256 requestTime) view returns((uint256))
func (_Emp *EmpCaller) TransformPrice(opts *bind.CallOpts, price FixedPointUnsigned, requestTime *big.Int) (FixedPointUnsigned, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "transformPrice", price, requestTime)

	if err != nil {
		return *new(FixedPointUnsigned), err
	}

	out0 := *abi.ConvertType(out[0], new(FixedPointUnsigned)).(*FixedPointUnsigned)

	return out0, err

}

// TransformPrice is a free data retrieval call binding the contract method 0x0ff49b90.
//
// Solidity: function transformPrice((uint256) price, uint256 requestTime) view returns((uint256))
func (_Emp *EmpSession) TransformPrice(price FixedPointUnsigned, requestTime *big.Int) (FixedPointUnsigned, error) {
	return _Emp.Contract.TransformPrice(&_Emp.CallOpts, price, requestTime)
}

// TransformPrice is a free data retrieval call binding the contract method 0x0ff49b90.
//
// Solidity: function transformPrice((uint256) price, uint256 requestTime) view returns((uint256))
func (_Emp *EmpCallerSession) TransformPrice(price FixedPointUnsigned, requestTime *big.Int) (FixedPointUnsigned, error) {
	return _Emp.Contract.TransformPrice(&_Emp.CallOpts, price, requestTime)
}

// TransformPriceIdentifier is a free data retrieval call binding the contract method 0x62b5f7f5.
//
// Solidity: function transformPriceIdentifier(uint256 requestTime) view returns(bytes32)
func (_Emp *EmpCaller) TransformPriceIdentifier(opts *bind.CallOpts, requestTime *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "transformPriceIdentifier", requestTime)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TransformPriceIdentifier is a free data retrieval call binding the contract method 0x62b5f7f5.
//
// Solidity: function transformPriceIdentifier(uint256 requestTime) view returns(bytes32)
func (_Emp *EmpSession) TransformPriceIdentifier(requestTime *big.Int) ([32]byte, error) {
	return _Emp.Contract.TransformPriceIdentifier(&_Emp.CallOpts, requestTime)
}

// TransformPriceIdentifier is a free data retrieval call binding the contract method 0x62b5f7f5.
//
// Solidity: function transformPriceIdentifier(uint256 requestTime) view returns(bytes32)
func (_Emp *EmpCallerSession) TransformPriceIdentifier(requestTime *big.Int) ([32]byte, error) {
	return _Emp.Contract.TransformPriceIdentifier(&_Emp.CallOpts, requestTime)
}

// WithdrawalLiveness is a free data retrieval call binding the contract method 0x9ff4dea8.
//
// Solidity: function withdrawalLiveness() view returns(uint256)
func (_Emp *EmpCaller) WithdrawalLiveness(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Emp.contract.Call(opts, &out, "withdrawalLiveness")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalLiveness is a free data retrieval call binding the contract method 0x9ff4dea8.
//
// Solidity: function withdrawalLiveness() view returns(uint256)
func (_Emp *EmpSession) WithdrawalLiveness() (*big.Int, error) {
	return _Emp.Contract.WithdrawalLiveness(&_Emp.CallOpts)
}

// WithdrawalLiveness is a free data retrieval call binding the contract method 0x9ff4dea8.
//
// Solidity: function withdrawalLiveness() view returns(uint256)
func (_Emp *EmpCallerSession) WithdrawalLiveness() (*big.Int, error) {
	return _Emp.Contract.WithdrawalLiveness(&_Emp.CallOpts)
}

// CancelTransferPosition is a paid mutator transaction binding the contract method 0xb795f0d4.
//
// Solidity: function cancelTransferPosition() returns()
func (_Emp *EmpTransactor) CancelTransferPosition(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Emp.contract.Transact(opts, "cancelTransferPosition")
}

// CancelTransferPosition is a paid mutator transaction binding the contract method 0xb795f0d4.
//
// Solidity: function cancelTransferPosition() returns()
func (_Emp *EmpSession) CancelTransferPosition() (*types.Transaction, error) {
	return _Emp.Contract.CancelTransferPosition(&_Emp.TransactOpts)
}

// CancelTransferPosition is a paid mutator transaction binding the contract method 0xb795f0d4.
//
// Solidity: function cancelTransferPosition() returns()
func (_Emp *EmpTransactorSession) CancelTransferPosition() (*types.Transaction, error) {
	return _Emp.Contract.CancelTransferPosition(&_Emp.TransactOpts)
}

// CancelWithdrawal is a paid mutator transaction binding the contract method 0x22611280.
//
// Solidity: function cancelWithdrawal() returns()
func (_Emp *EmpTransactor) CancelWithdrawal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Emp.contract.Transact(opts, "cancelWithdrawal")
}

// CancelWithdrawal is a paid mutator transaction binding the contract method 0x22611280.
//
// Solidity: function cancelWithdrawal() returns()
func (_Emp *EmpSession) CancelWithdrawal() (*types.Transaction, error) {
	return _Emp.Contract.CancelWithdrawal(&_Emp.TransactOpts)
}

// CancelWithdrawal is a paid mutator transaction binding the contract method 0x22611280.
//
// Solidity: function cancelWithdrawal() returns()
func (_Emp *EmpTransactorSession) CancelWithdrawal() (*types.Transaction, error) {
	return _Emp.Contract.CancelWithdrawal(&_Emp.TransactOpts)
}

// Create is a paid mutator transaction binding the contract method 0x6ba2f992.
//
// Solidity: function create((uint256) collateralAmount, (uint256) numTokens) returns()
func (_Emp *EmpTransactor) Create(opts *bind.TransactOpts, collateralAmount FixedPointUnsigned, numTokens FixedPointUnsigned) (*types.Transaction, error) {
	return _Emp.contract.Transact(opts, "create", collateralAmount, numTokens)
}

// Create is a paid mutator transaction binding the contract method 0x6ba2f992.
//
// Solidity: function create((uint256) collateralAmount, (uint256) numTokens) returns()
func (_Emp *EmpSession) Create(collateralAmount FixedPointUnsigned, numTokens FixedPointUnsigned) (*types.Transaction, error) {
	return _Emp.Contract.Create(&_Emp.TransactOpts, collateralAmount, numTokens)
}

// Create is a paid mutator transaction binding the contract method 0x6ba2f992.
//
// Solidity: function create((uint256) collateralAmount, (uint256) numTokens) returns()
func (_Emp *EmpTransactorSession) Create(collateralAmount FixedPointUnsigned, numTokens FixedPointUnsigned) (*types.Transaction, error) {
	return _Emp.Contract.Create(&_Emp.TransactOpts, collateralAmount, numTokens)
}

// CreateLiquidation is a paid mutator transaction binding the contract method 0x25ed4dd8.
//
// Solidity: function createLiquidation(address sponsor, (uint256) minCollateralPerToken, (uint256) maxCollateralPerToken, (uint256) maxTokensToLiquidate, uint256 deadline) returns(uint256 liquidationId, (uint256) tokensLiquidated, (uint256) finalFeeBond)
func (_Emp *EmpTransactor) CreateLiquidation(opts *bind.TransactOpts, sponsor common.Address, minCollateralPerToken FixedPointUnsigned, maxCollateralPerToken FixedPointUnsigned, maxTokensToLiquidate FixedPointUnsigned, deadline *big.Int) (*types.Transaction, error) {
	return _Emp.contract.Transact(opts, "createLiquidation", sponsor, minCollateralPerToken, maxCollateralPerToken, maxTokensToLiquidate, deadline)
}

// CreateLiquidation is a paid mutator transaction binding the contract method 0x25ed4dd8.
//
// Solidity: function createLiquidation(address sponsor, (uint256) minCollateralPerToken, (uint256) maxCollateralPerToken, (uint256) maxTokensToLiquidate, uint256 deadline) returns(uint256 liquidationId, (uint256) tokensLiquidated, (uint256) finalFeeBond)
func (_Emp *EmpSession) CreateLiquidation(sponsor common.Address, minCollateralPerToken FixedPointUnsigned, maxCollateralPerToken FixedPointUnsigned, maxTokensToLiquidate FixedPointUnsigned, deadline *big.Int) (*types.Transaction, error) {
	return _Emp.Contract.CreateLiquidation(&_Emp.TransactOpts, sponsor, minCollateralPerToken, maxCollateralPerToken, maxTokensToLiquidate, deadline)
}

// CreateLiquidation is a paid mutator transaction binding the contract method 0x25ed4dd8.
//
// Solidity: function createLiquidation(address sponsor, (uint256) minCollateralPerToken, (uint256) maxCollateralPerToken, (uint256) maxTokensToLiquidate, uint256 deadline) returns(uint256 liquidationId, (uint256) tokensLiquidated, (uint256) finalFeeBond)
func (_Emp *EmpTransactorSession) CreateLiquidation(sponsor common.Address, minCollateralPerToken FixedPointUnsigned, maxCollateralPerToken FixedPointUnsigned, maxTokensToLiquidate FixedPointUnsigned, deadline *big.Int) (*types.Transaction, error) {
	return _Emp.Contract.CreateLiquidation(&_Emp.TransactOpts, sponsor, minCollateralPerToken, maxCollateralPerToken, maxTokensToLiquidate, deadline)
}

// Deposit is a paid mutator transaction binding the contract method 0xd1e92c11.
//
// Solidity: function deposit((uint256) collateralAmount) returns()
func (_Emp *EmpTransactor) Deposit(opts *bind.TransactOpts, collateralAmount FixedPointUnsigned) (*types.Transaction, error) {
	return _Emp.contract.Transact(opts, "deposit", collateralAmount)
}

// Deposit is a paid mutator transaction binding the contract method 0xd1e92c11.
//
// Solidity: function deposit((uint256) collateralAmount) returns()
func (_Emp *EmpSession) Deposit(collateralAmount FixedPointUnsigned) (*types.Transaction, error) {
	return _Emp.Contract.Deposit(&_Emp.TransactOpts, collateralAmount)
}

// Deposit is a paid mutator transaction binding the contract method 0xd1e92c11.
//
// Solidity: function deposit((uint256) collateralAmount) returns()
func (_Emp *EmpTransactorSession) Deposit(collateralAmount FixedPointUnsigned) (*types.Transaction, error) {
	return _Emp.Contract.Deposit(&_Emp.TransactOpts, collateralAmount)
}

// DepositTo is a paid mutator transaction binding the contract method 0x18928a0c.
//
// Solidity: function depositTo(address sponsor, (uint256) collateralAmount) returns()
func (_Emp *EmpTransactor) DepositTo(opts *bind.TransactOpts, sponsor common.Address, collateralAmount FixedPointUnsigned) (*types.Transaction, error) {
	return _Emp.contract.Transact(opts, "depositTo", sponsor, collateralAmount)
}

// DepositTo is a paid mutator transaction binding the contract method 0x18928a0c.
//
// Solidity: function depositTo(address sponsor, (uint256) collateralAmount) returns()
func (_Emp *EmpSession) DepositTo(sponsor common.Address, collateralAmount FixedPointUnsigned) (*types.Transaction, error) {
	return _Emp.Contract.DepositTo(&_Emp.TransactOpts, sponsor, collateralAmount)
}

// DepositTo is a paid mutator transaction binding the contract method 0x18928a0c.
//
// Solidity: function depositTo(address sponsor, (uint256) collateralAmount) returns()
func (_Emp *EmpTransactorSession) DepositTo(sponsor common.Address, collateralAmount FixedPointUnsigned) (*types.Transaction, error) {
	return _Emp.Contract.DepositTo(&_Emp.TransactOpts, sponsor, collateralAmount)
}

// Dispute is a paid mutator transaction binding the contract method 0xa765fbea.
//
// Solidity: function dispute(uint256 liquidationId, address sponsor) returns((uint256) totalPaid)
func (_Emp *EmpTransactor) Dispute(opts *bind.TransactOpts, liquidationId *big.Int, sponsor common.Address) (*types.Transaction, error) {
	return _Emp.contract.Transact(opts, "dispute", liquidationId, sponsor)
}

// Dispute is a paid mutator transaction binding the contract method 0xa765fbea.
//
// Solidity: function dispute(uint256 liquidationId, address sponsor) returns((uint256) totalPaid)
func (_Emp *EmpSession) Dispute(liquidationId *big.Int, sponsor common.Address) (*types.Transaction, error) {
	return _Emp.Contract.Dispute(&_Emp.TransactOpts, liquidationId, sponsor)
}

// Dispute is a paid mutator transaction binding the contract method 0xa765fbea.
//
// Solidity: function dispute(uint256 liquidationId, address sponsor) returns((uint256) totalPaid)
func (_Emp *EmpTransactorSession) Dispute(liquidationId *big.Int, sponsor common.Address) (*types.Transaction, error) {
	return _Emp.Contract.Dispute(&_Emp.TransactOpts, liquidationId, sponsor)
}

// EmergencyShutdown is a paid mutator transaction binding the contract method 0x3403c2fc.
//
// Solidity: function emergencyShutdown() returns()
func (_Emp *EmpTransactor) EmergencyShutdown(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Emp.contract.Transact(opts, "emergencyShutdown")
}

// EmergencyShutdown is a paid mutator transaction binding the contract method 0x3403c2fc.
//
// Solidity: function emergencyShutdown() returns()
func (_Emp *EmpSession) EmergencyShutdown() (*types.Transaction, error) {
	return _Emp.Contract.EmergencyShutdown(&_Emp.TransactOpts)
}

// EmergencyShutdown is a paid mutator transaction binding the contract method 0x3403c2fc.
//
// Solidity: function emergencyShutdown() returns()
func (_Emp *EmpTransactorSession) EmergencyShutdown() (*types.Transaction, error) {
	return _Emp.Contract.EmergencyShutdown(&_Emp.TransactOpts)
}

// Expire is a paid mutator transaction binding the contract method 0x79599f96.
//
// Solidity: function expire() returns()
func (_Emp *EmpTransactor) Expire(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Emp.contract.Transact(opts, "expire")
}

// Expire is a paid mutator transaction binding the contract method 0x79599f96.
//
// Solidity: function expire() returns()
func (_Emp *EmpSession) Expire() (*types.Transaction, error) {
	return _Emp.Contract.Expire(&_Emp.TransactOpts)
}

// Expire is a paid mutator transaction binding the contract method 0x79599f96.
//
// Solidity: function expire() returns()
func (_Emp *EmpTransactorSession) Expire() (*types.Transaction, error) {
	return _Emp.Contract.Expire(&_Emp.TransactOpts)
}

// Gulp is a paid mutator transaction binding the contract method 0x94909e62.
//
// Solidity: function gulp() returns()
func (_Emp *EmpTransactor) Gulp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Emp.contract.Transact(opts, "gulp")
}

// Gulp is a paid mutator transaction binding the contract method 0x94909e62.
//
// Solidity: function gulp() returns()
func (_Emp *EmpSession) Gulp() (*types.Transaction, error) {
	return _Emp.Contract.Gulp(&_Emp.TransactOpts)
}

// Gulp is a paid mutator transaction binding the contract method 0x94909e62.
//
// Solidity: function gulp() returns()
func (_Emp *EmpTransactorSession) Gulp() (*types.Transaction, error) {
	return _Emp.Contract.Gulp(&_Emp.TransactOpts)
}

// PayRegularFees is a paid mutator transaction binding the contract method 0x3cb6ce83.
//
// Solidity: function payRegularFees() returns((uint256))
func (_Emp *EmpTransactor) PayRegularFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Emp.contract.Transact(opts, "payRegularFees")
}

// PayRegularFees is a paid mutator transaction binding the contract method 0x3cb6ce83.
//
// Solidity: function payRegularFees() returns((uint256))
func (_Emp *EmpSession) PayRegularFees() (*types.Transaction, error) {
	return _Emp.Contract.PayRegularFees(&_Emp.TransactOpts)
}

// PayRegularFees is a paid mutator transaction binding the contract method 0x3cb6ce83.
//
// Solidity: function payRegularFees() returns((uint256))
func (_Emp *EmpTransactorSession) PayRegularFees() (*types.Transaction, error) {
	return _Emp.Contract.PayRegularFees(&_Emp.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0x5f1af1ca.
//
// Solidity: function redeem((uint256) numTokens) returns((uint256) amountWithdrawn)
func (_Emp *EmpTransactor) Redeem(opts *bind.TransactOpts, numTokens FixedPointUnsigned) (*types.Transaction, error) {
	return _Emp.contract.Transact(opts, "redeem", numTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0x5f1af1ca.
//
// Solidity: function redeem((uint256) numTokens) returns((uint256) amountWithdrawn)
func (_Emp *EmpSession) Redeem(numTokens FixedPointUnsigned) (*types.Transaction, error) {
	return _Emp.Contract.Redeem(&_Emp.TransactOpts, numTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0x5f1af1ca.
//
// Solidity: function redeem((uint256) numTokens) returns((uint256) amountWithdrawn)
func (_Emp *EmpTransactorSession) Redeem(numTokens FixedPointUnsigned) (*types.Transaction, error) {
	return _Emp.Contract.Redeem(&_Emp.TransactOpts, numTokens)
}

// Remargin is a paid mutator transaction binding the contract method 0xbda02e77.
//
// Solidity: function remargin() returns()
func (_Emp *EmpTransactor) Remargin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Emp.contract.Transact(opts, "remargin")
}

// Remargin is a paid mutator transaction binding the contract method 0xbda02e77.
//
// Solidity: function remargin() returns()
func (_Emp *EmpSession) Remargin() (*types.Transaction, error) {
	return _Emp.Contract.Remargin(&_Emp.TransactOpts)
}

// Remargin is a paid mutator transaction binding the contract method 0xbda02e77.
//
// Solidity: function remargin() returns()
func (_Emp *EmpTransactorSession) Remargin() (*types.Transaction, error) {
	return _Emp.Contract.Remargin(&_Emp.TransactOpts)
}

// Repay is a paid mutator transaction binding the contract method 0x5aa266c9.
//
// Solidity: function repay((uint256) numTokens) returns()
func (_Emp *EmpTransactor) Repay(opts *bind.TransactOpts, numTokens FixedPointUnsigned) (*types.Transaction, error) {
	return _Emp.contract.Transact(opts, "repay", numTokens)
}

// Repay is a paid mutator transaction binding the contract method 0x5aa266c9.
//
// Solidity: function repay((uint256) numTokens) returns()
func (_Emp *EmpSession) Repay(numTokens FixedPointUnsigned) (*types.Transaction, error) {
	return _Emp.Contract.Repay(&_Emp.TransactOpts, numTokens)
}

// Repay is a paid mutator transaction binding the contract method 0x5aa266c9.
//
// Solidity: function repay((uint256) numTokens) returns()
func (_Emp *EmpTransactorSession) Repay(numTokens FixedPointUnsigned) (*types.Transaction, error) {
	return _Emp.Contract.Repay(&_Emp.TransactOpts, numTokens)
}

// RequestTransferPosition is a paid mutator transaction binding the contract method 0x36980f58.
//
// Solidity: function requestTransferPosition() returns()
func (_Emp *EmpTransactor) RequestTransferPosition(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Emp.contract.Transact(opts, "requestTransferPosition")
}

// RequestTransferPosition is a paid mutator transaction binding the contract method 0x36980f58.
//
// Solidity: function requestTransferPosition() returns()
func (_Emp *EmpSession) RequestTransferPosition() (*types.Transaction, error) {
	return _Emp.Contract.RequestTransferPosition(&_Emp.TransactOpts)
}

// RequestTransferPosition is a paid mutator transaction binding the contract method 0x36980f58.
//
// Solidity: function requestTransferPosition() returns()
func (_Emp *EmpTransactorSession) RequestTransferPosition() (*types.Transaction, error) {
	return _Emp.Contract.RequestTransferPosition(&_Emp.TransactOpts)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0xbc121630.
//
// Solidity: function requestWithdrawal((uint256) collateralAmount) returns()
func (_Emp *EmpTransactor) RequestWithdrawal(opts *bind.TransactOpts, collateralAmount FixedPointUnsigned) (*types.Transaction, error) {
	return _Emp.contract.Transact(opts, "requestWithdrawal", collateralAmount)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0xbc121630.
//
// Solidity: function requestWithdrawal((uint256) collateralAmount) returns()
func (_Emp *EmpSession) RequestWithdrawal(collateralAmount FixedPointUnsigned) (*types.Transaction, error) {
	return _Emp.Contract.RequestWithdrawal(&_Emp.TransactOpts, collateralAmount)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0xbc121630.
//
// Solidity: function requestWithdrawal((uint256) collateralAmount) returns()
func (_Emp *EmpTransactorSession) RequestWithdrawal(collateralAmount FixedPointUnsigned) (*types.Transaction, error) {
	return _Emp.Contract.RequestWithdrawal(&_Emp.TransactOpts, collateralAmount)
}

// SetCurrentTime is a paid mutator transaction binding the contract method 0x22f8e566.
//
// Solidity: function setCurrentTime(uint256 time) returns()
func (_Emp *EmpTransactor) SetCurrentTime(opts *bind.TransactOpts, time *big.Int) (*types.Transaction, error) {
	return _Emp.contract.Transact(opts, "setCurrentTime", time)
}

// SetCurrentTime is a paid mutator transaction binding the contract method 0x22f8e566.
//
// Solidity: function setCurrentTime(uint256 time) returns()
func (_Emp *EmpSession) SetCurrentTime(time *big.Int) (*types.Transaction, error) {
	return _Emp.Contract.SetCurrentTime(&_Emp.TransactOpts, time)
}

// SetCurrentTime is a paid mutator transaction binding the contract method 0x22f8e566.
//
// Solidity: function setCurrentTime(uint256 time) returns()
func (_Emp *EmpTransactorSession) SetCurrentTime(time *big.Int) (*types.Transaction, error) {
	return _Emp.Contract.SetCurrentTime(&_Emp.TransactOpts, time)
}

// SettleExpired is a paid mutator transaction binding the contract method 0xfcccedc7.
//
// Solidity: function settleExpired() returns((uint256) amountWithdrawn)
func (_Emp *EmpTransactor) SettleExpired(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Emp.contract.Transact(opts, "settleExpired")
}

// SettleExpired is a paid mutator transaction binding the contract method 0xfcccedc7.
//
// Solidity: function settleExpired() returns((uint256) amountWithdrawn)
func (_Emp *EmpSession) SettleExpired() (*types.Transaction, error) {
	return _Emp.Contract.SettleExpired(&_Emp.TransactOpts)
}

// SettleExpired is a paid mutator transaction binding the contract method 0xfcccedc7.
//
// Solidity: function settleExpired() returns((uint256) amountWithdrawn)
func (_Emp *EmpTransactorSession) SettleExpired() (*types.Transaction, error) {
	return _Emp.Contract.SettleExpired(&_Emp.TransactOpts)
}

// TransferPositionPassedRequest is a paid mutator transaction binding the contract method 0x5617151c.
//
// Solidity: function transferPositionPassedRequest(address newSponsorAddress) returns()
func (_Emp *EmpTransactor) TransferPositionPassedRequest(opts *bind.TransactOpts, newSponsorAddress common.Address) (*types.Transaction, error) {
	return _Emp.contract.Transact(opts, "transferPositionPassedRequest", newSponsorAddress)
}

// TransferPositionPassedRequest is a paid mutator transaction binding the contract method 0x5617151c.
//
// Solidity: function transferPositionPassedRequest(address newSponsorAddress) returns()
func (_Emp *EmpSession) TransferPositionPassedRequest(newSponsorAddress common.Address) (*types.Transaction, error) {
	return _Emp.Contract.TransferPositionPassedRequest(&_Emp.TransactOpts, newSponsorAddress)
}

// TransferPositionPassedRequest is a paid mutator transaction binding the contract method 0x5617151c.
//
// Solidity: function transferPositionPassedRequest(address newSponsorAddress) returns()
func (_Emp *EmpTransactorSession) TransferPositionPassedRequest(newSponsorAddress common.Address) (*types.Transaction, error) {
	return _Emp.Contract.TransferPositionPassedRequest(&_Emp.TransactOpts, newSponsorAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ee7a5ce.
//
// Solidity: function withdraw((uint256) collateralAmount) returns((uint256) amountWithdrawn)
func (_Emp *EmpTransactor) Withdraw(opts *bind.TransactOpts, collateralAmount FixedPointUnsigned) (*types.Transaction, error) {
	return _Emp.contract.Transact(opts, "withdraw", collateralAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ee7a5ce.
//
// Solidity: function withdraw((uint256) collateralAmount) returns((uint256) amountWithdrawn)
func (_Emp *EmpSession) Withdraw(collateralAmount FixedPointUnsigned) (*types.Transaction, error) {
	return _Emp.Contract.Withdraw(&_Emp.TransactOpts, collateralAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ee7a5ce.
//
// Solidity: function withdraw((uint256) collateralAmount) returns((uint256) amountWithdrawn)
func (_Emp *EmpTransactorSession) Withdraw(collateralAmount FixedPointUnsigned) (*types.Transaction, error) {
	return _Emp.Contract.Withdraw(&_Emp.TransactOpts, collateralAmount)
}

// WithdrawLiquidation is a paid mutator transaction binding the contract method 0x360598e1.
//
// Solidity: function withdrawLiquidation(uint256 liquidationId, address sponsor) returns(((uint256),(uint256),(uint256),(uint256),(uint256),(uint256)))
func (_Emp *EmpTransactor) WithdrawLiquidation(opts *bind.TransactOpts, liquidationId *big.Int, sponsor common.Address) (*types.Transaction, error) {
	return _Emp.contract.Transact(opts, "withdrawLiquidation", liquidationId, sponsor)
}

// WithdrawLiquidation is a paid mutator transaction binding the contract method 0x360598e1.
//
// Solidity: function withdrawLiquidation(uint256 liquidationId, address sponsor) returns(((uint256),(uint256),(uint256),(uint256),(uint256),(uint256)))
func (_Emp *EmpSession) WithdrawLiquidation(liquidationId *big.Int, sponsor common.Address) (*types.Transaction, error) {
	return _Emp.Contract.WithdrawLiquidation(&_Emp.TransactOpts, liquidationId, sponsor)
}

// WithdrawLiquidation is a paid mutator transaction binding the contract method 0x360598e1.
//
// Solidity: function withdrawLiquidation(uint256 liquidationId, address sponsor) returns(((uint256),(uint256),(uint256),(uint256),(uint256),(uint256)))
func (_Emp *EmpTransactorSession) WithdrawLiquidation(liquidationId *big.Int, sponsor common.Address) (*types.Transaction, error) {
	return _Emp.Contract.WithdrawLiquidation(&_Emp.TransactOpts, liquidationId, sponsor)
}

// WithdrawPassedRequest is a paid mutator transaction binding the contract method 0x33a46ca2.
//
// Solidity: function withdrawPassedRequest() returns((uint256) amountWithdrawn)
func (_Emp *EmpTransactor) WithdrawPassedRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Emp.contract.Transact(opts, "withdrawPassedRequest")
}

// WithdrawPassedRequest is a paid mutator transaction binding the contract method 0x33a46ca2.
//
// Solidity: function withdrawPassedRequest() returns((uint256) amountWithdrawn)
func (_Emp *EmpSession) WithdrawPassedRequest() (*types.Transaction, error) {
	return _Emp.Contract.WithdrawPassedRequest(&_Emp.TransactOpts)
}

// WithdrawPassedRequest is a paid mutator transaction binding the contract method 0x33a46ca2.
//
// Solidity: function withdrawPassedRequest() returns((uint256) amountWithdrawn)
func (_Emp *EmpTransactorSession) WithdrawPassedRequest() (*types.Transaction, error) {
	return _Emp.Contract.WithdrawPassedRequest(&_Emp.TransactOpts)
}

// EmpContractExpiredIterator is returned from FilterContractExpired and is used to iterate over the raw logs and unpacked data for ContractExpired events raised by the Emp contract.
type EmpContractExpiredIterator struct {
	Event *EmpContractExpired // Event containing the contract specifics and raw log

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
func (it *EmpContractExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmpContractExpired)
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
		it.Event = new(EmpContractExpired)
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
func (it *EmpContractExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmpContractExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmpContractExpired represents a ContractExpired event raised by the Emp contract.
type EmpContractExpired struct {
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterContractExpired is a free log retrieval operation binding the contract event 0x18600820405d6cf356e3556301762ca32395e72d8c81494fa344835c9da3633d.
//
// Solidity: event ContractExpired(address indexed caller)
func (_Emp *EmpFilterer) FilterContractExpired(opts *bind.FilterOpts, caller []common.Address) (*EmpContractExpiredIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Emp.contract.FilterLogs(opts, "ContractExpired", callerRule)
	if err != nil {
		return nil, err
	}
	return &EmpContractExpiredIterator{contract: _Emp.contract, event: "ContractExpired", logs: logs, sub: sub}, nil
}

// WatchContractExpired is a free log subscription operation binding the contract event 0x18600820405d6cf356e3556301762ca32395e72d8c81494fa344835c9da3633d.
//
// Solidity: event ContractExpired(address indexed caller)
func (_Emp *EmpFilterer) WatchContractExpired(opts *bind.WatchOpts, sink chan<- *EmpContractExpired, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Emp.contract.WatchLogs(opts, "ContractExpired", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmpContractExpired)
				if err := _Emp.contract.UnpackLog(event, "ContractExpired", log); err != nil {
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

// ParseContractExpired is a log parse operation binding the contract event 0x18600820405d6cf356e3556301762ca32395e72d8c81494fa344835c9da3633d.
//
// Solidity: event ContractExpired(address indexed caller)
func (_Emp *EmpFilterer) ParseContractExpired(log types.Log) (*EmpContractExpired, error) {
	event := new(EmpContractExpired)
	if err := _Emp.contract.UnpackLog(event, "ContractExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmpDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Emp contract.
type EmpDepositIterator struct {
	Event *EmpDeposit // Event containing the contract specifics and raw log

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
func (it *EmpDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmpDeposit)
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
		it.Event = new(EmpDeposit)
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
func (it *EmpDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmpDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmpDeposit represents a Deposit event raised by the Emp contract.
type EmpDeposit struct {
	Sponsor          common.Address
	CollateralAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sponsor, uint256 indexed collateralAmount)
func (_Emp *EmpFilterer) FilterDeposit(opts *bind.FilterOpts, sponsor []common.Address, collateralAmount []*big.Int) (*EmpDepositIterator, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var collateralAmountRule []interface{}
	for _, collateralAmountItem := range collateralAmount {
		collateralAmountRule = append(collateralAmountRule, collateralAmountItem)
	}

	logs, sub, err := _Emp.contract.FilterLogs(opts, "Deposit", sponsorRule, collateralAmountRule)
	if err != nil {
		return nil, err
	}
	return &EmpDepositIterator{contract: _Emp.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sponsor, uint256 indexed collateralAmount)
func (_Emp *EmpFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *EmpDeposit, sponsor []common.Address, collateralAmount []*big.Int) (event.Subscription, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var collateralAmountRule []interface{}
	for _, collateralAmountItem := range collateralAmount {
		collateralAmountRule = append(collateralAmountRule, collateralAmountItem)
	}

	logs, sub, err := _Emp.contract.WatchLogs(opts, "Deposit", sponsorRule, collateralAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmpDeposit)
				if err := _Emp.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sponsor, uint256 indexed collateralAmount)
func (_Emp *EmpFilterer) ParseDeposit(log types.Log) (*EmpDeposit, error) {
	event := new(EmpDeposit)
	if err := _Emp.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmpDisputeSettledIterator is returned from FilterDisputeSettled and is used to iterate over the raw logs and unpacked data for DisputeSettled events raised by the Emp contract.
type EmpDisputeSettledIterator struct {
	Event *EmpDisputeSettled // Event containing the contract specifics and raw log

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
func (it *EmpDisputeSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmpDisputeSettled)
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
		it.Event = new(EmpDisputeSettled)
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
func (it *EmpDisputeSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmpDisputeSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmpDisputeSettled represents a DisputeSettled event raised by the Emp contract.
type EmpDisputeSettled struct {
	Caller           common.Address
	Sponsor          common.Address
	Liquidator       common.Address
	Disputer         common.Address
	LiquidationId    *big.Int
	DisputeSucceeded bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDisputeSettled is a free log retrieval operation binding the contract event 0x6c5582199868fabbe697f9ea10abe481bacf53ac78c02a965b34dff82fd20e3b.
//
// Solidity: event DisputeSettled(address indexed caller, address indexed sponsor, address indexed liquidator, address disputer, uint256 liquidationId, bool disputeSucceeded)
func (_Emp *EmpFilterer) FilterDisputeSettled(opts *bind.FilterOpts, caller []common.Address, sponsor []common.Address, liquidator []common.Address) (*EmpDisputeSettledIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _Emp.contract.FilterLogs(opts, "DisputeSettled", callerRule, sponsorRule, liquidatorRule)
	if err != nil {
		return nil, err
	}
	return &EmpDisputeSettledIterator{contract: _Emp.contract, event: "DisputeSettled", logs: logs, sub: sub}, nil
}

// WatchDisputeSettled is a free log subscription operation binding the contract event 0x6c5582199868fabbe697f9ea10abe481bacf53ac78c02a965b34dff82fd20e3b.
//
// Solidity: event DisputeSettled(address indexed caller, address indexed sponsor, address indexed liquidator, address disputer, uint256 liquidationId, bool disputeSucceeded)
func (_Emp *EmpFilterer) WatchDisputeSettled(opts *bind.WatchOpts, sink chan<- *EmpDisputeSettled, caller []common.Address, sponsor []common.Address, liquidator []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _Emp.contract.WatchLogs(opts, "DisputeSettled", callerRule, sponsorRule, liquidatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmpDisputeSettled)
				if err := _Emp.contract.UnpackLog(event, "DisputeSettled", log); err != nil {
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

// ParseDisputeSettled is a log parse operation binding the contract event 0x6c5582199868fabbe697f9ea10abe481bacf53ac78c02a965b34dff82fd20e3b.
//
// Solidity: event DisputeSettled(address indexed caller, address indexed sponsor, address indexed liquidator, address disputer, uint256 liquidationId, bool disputeSucceeded)
func (_Emp *EmpFilterer) ParseDisputeSettled(log types.Log) (*EmpDisputeSettled, error) {
	event := new(EmpDisputeSettled)
	if err := _Emp.contract.UnpackLog(event, "DisputeSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmpEmergencyShutdownIterator is returned from FilterEmergencyShutdown and is used to iterate over the raw logs and unpacked data for EmergencyShutdown events raised by the Emp contract.
type EmpEmergencyShutdownIterator struct {
	Event *EmpEmergencyShutdown // Event containing the contract specifics and raw log

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
func (it *EmpEmergencyShutdownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmpEmergencyShutdown)
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
		it.Event = new(EmpEmergencyShutdown)
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
func (it *EmpEmergencyShutdownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmpEmergencyShutdownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmpEmergencyShutdown represents a EmergencyShutdown event raised by the Emp contract.
type EmpEmergencyShutdown struct {
	Caller                      common.Address
	OriginalExpirationTimestamp *big.Int
	ShutdownTimestamp           *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterEmergencyShutdown is a free log retrieval operation binding the contract event 0xd39eeb7157d9c446579a0893ecf9ecd87d1f466cdb270c6a189cf38ca1e30f48.
//
// Solidity: event EmergencyShutdown(address indexed caller, uint256 originalExpirationTimestamp, uint256 shutdownTimestamp)
func (_Emp *EmpFilterer) FilterEmergencyShutdown(opts *bind.FilterOpts, caller []common.Address) (*EmpEmergencyShutdownIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Emp.contract.FilterLogs(opts, "EmergencyShutdown", callerRule)
	if err != nil {
		return nil, err
	}
	return &EmpEmergencyShutdownIterator{contract: _Emp.contract, event: "EmergencyShutdown", logs: logs, sub: sub}, nil
}

// WatchEmergencyShutdown is a free log subscription operation binding the contract event 0xd39eeb7157d9c446579a0893ecf9ecd87d1f466cdb270c6a189cf38ca1e30f48.
//
// Solidity: event EmergencyShutdown(address indexed caller, uint256 originalExpirationTimestamp, uint256 shutdownTimestamp)
func (_Emp *EmpFilterer) WatchEmergencyShutdown(opts *bind.WatchOpts, sink chan<- *EmpEmergencyShutdown, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Emp.contract.WatchLogs(opts, "EmergencyShutdown", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmpEmergencyShutdown)
				if err := _Emp.contract.UnpackLog(event, "EmergencyShutdown", log); err != nil {
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

// ParseEmergencyShutdown is a log parse operation binding the contract event 0xd39eeb7157d9c446579a0893ecf9ecd87d1f466cdb270c6a189cf38ca1e30f48.
//
// Solidity: event EmergencyShutdown(address indexed caller, uint256 originalExpirationTimestamp, uint256 shutdownTimestamp)
func (_Emp *EmpFilterer) ParseEmergencyShutdown(log types.Log) (*EmpEmergencyShutdown, error) {
	event := new(EmpEmergencyShutdown)
	if err := _Emp.contract.UnpackLog(event, "EmergencyShutdown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmpEndedSponsorPositionIterator is returned from FilterEndedSponsorPosition and is used to iterate over the raw logs and unpacked data for EndedSponsorPosition events raised by the Emp contract.
type EmpEndedSponsorPositionIterator struct {
	Event *EmpEndedSponsorPosition // Event containing the contract specifics and raw log

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
func (it *EmpEndedSponsorPositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmpEndedSponsorPosition)
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
		it.Event = new(EmpEndedSponsorPosition)
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
func (it *EmpEndedSponsorPositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmpEndedSponsorPositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmpEndedSponsorPosition represents a EndedSponsorPosition event raised by the Emp contract.
type EmpEndedSponsorPosition struct {
	Sponsor common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEndedSponsorPosition is a free log retrieval operation binding the contract event 0xcad20625296d189a6fc6e5b39d0d544e5bd99dbda0c8f2f0ecffef3e0fbcc282.
//
// Solidity: event EndedSponsorPosition(address indexed sponsor)
func (_Emp *EmpFilterer) FilterEndedSponsorPosition(opts *bind.FilterOpts, sponsor []common.Address) (*EmpEndedSponsorPositionIterator, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}

	logs, sub, err := _Emp.contract.FilterLogs(opts, "EndedSponsorPosition", sponsorRule)
	if err != nil {
		return nil, err
	}
	return &EmpEndedSponsorPositionIterator{contract: _Emp.contract, event: "EndedSponsorPosition", logs: logs, sub: sub}, nil
}

// WatchEndedSponsorPosition is a free log subscription operation binding the contract event 0xcad20625296d189a6fc6e5b39d0d544e5bd99dbda0c8f2f0ecffef3e0fbcc282.
//
// Solidity: event EndedSponsorPosition(address indexed sponsor)
func (_Emp *EmpFilterer) WatchEndedSponsorPosition(opts *bind.WatchOpts, sink chan<- *EmpEndedSponsorPosition, sponsor []common.Address) (event.Subscription, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}

	logs, sub, err := _Emp.contract.WatchLogs(opts, "EndedSponsorPosition", sponsorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmpEndedSponsorPosition)
				if err := _Emp.contract.UnpackLog(event, "EndedSponsorPosition", log); err != nil {
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

// ParseEndedSponsorPosition is a log parse operation binding the contract event 0xcad20625296d189a6fc6e5b39d0d544e5bd99dbda0c8f2f0ecffef3e0fbcc282.
//
// Solidity: event EndedSponsorPosition(address indexed sponsor)
func (_Emp *EmpFilterer) ParseEndedSponsorPosition(log types.Log) (*EmpEndedSponsorPosition, error) {
	event := new(EmpEndedSponsorPosition)
	if err := _Emp.contract.UnpackLog(event, "EndedSponsorPosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmpFinalFeesPaidIterator is returned from FilterFinalFeesPaid and is used to iterate over the raw logs and unpacked data for FinalFeesPaid events raised by the Emp contract.
type EmpFinalFeesPaidIterator struct {
	Event *EmpFinalFeesPaid // Event containing the contract specifics and raw log

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
func (it *EmpFinalFeesPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmpFinalFeesPaid)
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
		it.Event = new(EmpFinalFeesPaid)
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
func (it *EmpFinalFeesPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmpFinalFeesPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmpFinalFeesPaid represents a FinalFeesPaid event raised by the Emp contract.
type EmpFinalFeesPaid struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFinalFeesPaid is a free log retrieval operation binding the contract event 0x4f9bf7e8cd0f2456f9c43d2597bedcf1446c9c64544053f1ece6423ae9a07e52.
//
// Solidity: event FinalFeesPaid(uint256 indexed amount)
func (_Emp *EmpFilterer) FilterFinalFeesPaid(opts *bind.FilterOpts, amount []*big.Int) (*EmpFinalFeesPaidIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Emp.contract.FilterLogs(opts, "FinalFeesPaid", amountRule)
	if err != nil {
		return nil, err
	}
	return &EmpFinalFeesPaidIterator{contract: _Emp.contract, event: "FinalFeesPaid", logs: logs, sub: sub}, nil
}

// WatchFinalFeesPaid is a free log subscription operation binding the contract event 0x4f9bf7e8cd0f2456f9c43d2597bedcf1446c9c64544053f1ece6423ae9a07e52.
//
// Solidity: event FinalFeesPaid(uint256 indexed amount)
func (_Emp *EmpFilterer) WatchFinalFeesPaid(opts *bind.WatchOpts, sink chan<- *EmpFinalFeesPaid, amount []*big.Int) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Emp.contract.WatchLogs(opts, "FinalFeesPaid", amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmpFinalFeesPaid)
				if err := _Emp.contract.UnpackLog(event, "FinalFeesPaid", log); err != nil {
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

// ParseFinalFeesPaid is a log parse operation binding the contract event 0x4f9bf7e8cd0f2456f9c43d2597bedcf1446c9c64544053f1ece6423ae9a07e52.
//
// Solidity: event FinalFeesPaid(uint256 indexed amount)
func (_Emp *EmpFilterer) ParseFinalFeesPaid(log types.Log) (*EmpFinalFeesPaid, error) {
	event := new(EmpFinalFeesPaid)
	if err := _Emp.contract.UnpackLog(event, "FinalFeesPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmpLiquidationCreatedIterator is returned from FilterLiquidationCreated and is used to iterate over the raw logs and unpacked data for LiquidationCreated events raised by the Emp contract.
type EmpLiquidationCreatedIterator struct {
	Event *EmpLiquidationCreated // Event containing the contract specifics and raw log

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
func (it *EmpLiquidationCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmpLiquidationCreated)
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
		it.Event = new(EmpLiquidationCreated)
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
func (it *EmpLiquidationCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmpLiquidationCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmpLiquidationCreated represents a LiquidationCreated event raised by the Emp contract.
type EmpLiquidationCreated struct {
	Sponsor              common.Address
	Liquidator           common.Address
	LiquidationId        *big.Int
	TokensOutstanding    *big.Int
	LockedCollateral     *big.Int
	LiquidatedCollateral *big.Int
	LiquidationTime      *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterLiquidationCreated is a free log retrieval operation binding the contract event 0x39b4371645b4132767fd76a1aad3108ff95c20d7b687b24d171555f5459a7597.
//
// Solidity: event LiquidationCreated(address indexed sponsor, address indexed liquidator, uint256 indexed liquidationId, uint256 tokensOutstanding, uint256 lockedCollateral, uint256 liquidatedCollateral, uint256 liquidationTime)
func (_Emp *EmpFilterer) FilterLiquidationCreated(opts *bind.FilterOpts, sponsor []common.Address, liquidator []common.Address, liquidationId []*big.Int) (*EmpLiquidationCreatedIterator, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var liquidationIdRule []interface{}
	for _, liquidationIdItem := range liquidationId {
		liquidationIdRule = append(liquidationIdRule, liquidationIdItem)
	}

	logs, sub, err := _Emp.contract.FilterLogs(opts, "LiquidationCreated", sponsorRule, liquidatorRule, liquidationIdRule)
	if err != nil {
		return nil, err
	}
	return &EmpLiquidationCreatedIterator{contract: _Emp.contract, event: "LiquidationCreated", logs: logs, sub: sub}, nil
}

// WatchLiquidationCreated is a free log subscription operation binding the contract event 0x39b4371645b4132767fd76a1aad3108ff95c20d7b687b24d171555f5459a7597.
//
// Solidity: event LiquidationCreated(address indexed sponsor, address indexed liquidator, uint256 indexed liquidationId, uint256 tokensOutstanding, uint256 lockedCollateral, uint256 liquidatedCollateral, uint256 liquidationTime)
func (_Emp *EmpFilterer) WatchLiquidationCreated(opts *bind.WatchOpts, sink chan<- *EmpLiquidationCreated, sponsor []common.Address, liquidator []common.Address, liquidationId []*big.Int) (event.Subscription, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var liquidationIdRule []interface{}
	for _, liquidationIdItem := range liquidationId {
		liquidationIdRule = append(liquidationIdRule, liquidationIdItem)
	}

	logs, sub, err := _Emp.contract.WatchLogs(opts, "LiquidationCreated", sponsorRule, liquidatorRule, liquidationIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmpLiquidationCreated)
				if err := _Emp.contract.UnpackLog(event, "LiquidationCreated", log); err != nil {
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

// ParseLiquidationCreated is a log parse operation binding the contract event 0x39b4371645b4132767fd76a1aad3108ff95c20d7b687b24d171555f5459a7597.
//
// Solidity: event LiquidationCreated(address indexed sponsor, address indexed liquidator, uint256 indexed liquidationId, uint256 tokensOutstanding, uint256 lockedCollateral, uint256 liquidatedCollateral, uint256 liquidationTime)
func (_Emp *EmpFilterer) ParseLiquidationCreated(log types.Log) (*EmpLiquidationCreated, error) {
	event := new(EmpLiquidationCreated)
	if err := _Emp.contract.UnpackLog(event, "LiquidationCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmpLiquidationDisputedIterator is returned from FilterLiquidationDisputed and is used to iterate over the raw logs and unpacked data for LiquidationDisputed events raised by the Emp contract.
type EmpLiquidationDisputedIterator struct {
	Event *EmpLiquidationDisputed // Event containing the contract specifics and raw log

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
func (it *EmpLiquidationDisputedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmpLiquidationDisputed)
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
		it.Event = new(EmpLiquidationDisputed)
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
func (it *EmpLiquidationDisputedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmpLiquidationDisputedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmpLiquidationDisputed represents a LiquidationDisputed event raised by the Emp contract.
type EmpLiquidationDisputed struct {
	Sponsor           common.Address
	Liquidator        common.Address
	Disputer          common.Address
	LiquidationId     *big.Int
	DisputeBondAmount *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLiquidationDisputed is a free log retrieval operation binding the contract event 0xcaca181ccad7979cf36ed4fc921e496001ab5264608f0fac7007ae1b43d36102.
//
// Solidity: event LiquidationDisputed(address indexed sponsor, address indexed liquidator, address indexed disputer, uint256 liquidationId, uint256 disputeBondAmount)
func (_Emp *EmpFilterer) FilterLiquidationDisputed(opts *bind.FilterOpts, sponsor []common.Address, liquidator []common.Address, disputer []common.Address) (*EmpLiquidationDisputedIterator, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var disputerRule []interface{}
	for _, disputerItem := range disputer {
		disputerRule = append(disputerRule, disputerItem)
	}

	logs, sub, err := _Emp.contract.FilterLogs(opts, "LiquidationDisputed", sponsorRule, liquidatorRule, disputerRule)
	if err != nil {
		return nil, err
	}
	return &EmpLiquidationDisputedIterator{contract: _Emp.contract, event: "LiquidationDisputed", logs: logs, sub: sub}, nil
}

// WatchLiquidationDisputed is a free log subscription operation binding the contract event 0xcaca181ccad7979cf36ed4fc921e496001ab5264608f0fac7007ae1b43d36102.
//
// Solidity: event LiquidationDisputed(address indexed sponsor, address indexed liquidator, address indexed disputer, uint256 liquidationId, uint256 disputeBondAmount)
func (_Emp *EmpFilterer) WatchLiquidationDisputed(opts *bind.WatchOpts, sink chan<- *EmpLiquidationDisputed, sponsor []common.Address, liquidator []common.Address, disputer []common.Address) (event.Subscription, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var disputerRule []interface{}
	for _, disputerItem := range disputer {
		disputerRule = append(disputerRule, disputerItem)
	}

	logs, sub, err := _Emp.contract.WatchLogs(opts, "LiquidationDisputed", sponsorRule, liquidatorRule, disputerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmpLiquidationDisputed)
				if err := _Emp.contract.UnpackLog(event, "LiquidationDisputed", log); err != nil {
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

// ParseLiquidationDisputed is a log parse operation binding the contract event 0xcaca181ccad7979cf36ed4fc921e496001ab5264608f0fac7007ae1b43d36102.
//
// Solidity: event LiquidationDisputed(address indexed sponsor, address indexed liquidator, address indexed disputer, uint256 liquidationId, uint256 disputeBondAmount)
func (_Emp *EmpFilterer) ParseLiquidationDisputed(log types.Log) (*EmpLiquidationDisputed, error) {
	event := new(EmpLiquidationDisputed)
	if err := _Emp.contract.UnpackLog(event, "LiquidationDisputed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmpLiquidationWithdrawnIterator is returned from FilterLiquidationWithdrawn and is used to iterate over the raw logs and unpacked data for LiquidationWithdrawn events raised by the Emp contract.
type EmpLiquidationWithdrawnIterator struct {
	Event *EmpLiquidationWithdrawn // Event containing the contract specifics and raw log

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
func (it *EmpLiquidationWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmpLiquidationWithdrawn)
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
		it.Event = new(EmpLiquidationWithdrawn)
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
func (it *EmpLiquidationWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmpLiquidationWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmpLiquidationWithdrawn represents a LiquidationWithdrawn event raised by the Emp contract.
type EmpLiquidationWithdrawn struct {
	Caller            common.Address
	PaidToLiquidator  *big.Int
	PaidToDisputer    *big.Int
	PaidToSponsor     *big.Int
	LiquidationStatus uint8
	SettlementPrice   *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLiquidationWithdrawn is a free log retrieval operation binding the contract event 0xb479588a37dc7f6bac1c91587fcfc539cac4949cf26bb536ad9c8d061f00f50d.
//
// Solidity: event LiquidationWithdrawn(address indexed caller, uint256 paidToLiquidator, uint256 paidToDisputer, uint256 paidToSponsor, uint8 indexed liquidationStatus, uint256 settlementPrice)
func (_Emp *EmpFilterer) FilterLiquidationWithdrawn(opts *bind.FilterOpts, caller []common.Address, liquidationStatus []uint8) (*EmpLiquidationWithdrawnIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	var liquidationStatusRule []interface{}
	for _, liquidationStatusItem := range liquidationStatus {
		liquidationStatusRule = append(liquidationStatusRule, liquidationStatusItem)
	}

	logs, sub, err := _Emp.contract.FilterLogs(opts, "LiquidationWithdrawn", callerRule, liquidationStatusRule)
	if err != nil {
		return nil, err
	}
	return &EmpLiquidationWithdrawnIterator{contract: _Emp.contract, event: "LiquidationWithdrawn", logs: logs, sub: sub}, nil
}

// WatchLiquidationWithdrawn is a free log subscription operation binding the contract event 0xb479588a37dc7f6bac1c91587fcfc539cac4949cf26bb536ad9c8d061f00f50d.
//
// Solidity: event LiquidationWithdrawn(address indexed caller, uint256 paidToLiquidator, uint256 paidToDisputer, uint256 paidToSponsor, uint8 indexed liquidationStatus, uint256 settlementPrice)
func (_Emp *EmpFilterer) WatchLiquidationWithdrawn(opts *bind.WatchOpts, sink chan<- *EmpLiquidationWithdrawn, caller []common.Address, liquidationStatus []uint8) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	var liquidationStatusRule []interface{}
	for _, liquidationStatusItem := range liquidationStatus {
		liquidationStatusRule = append(liquidationStatusRule, liquidationStatusItem)
	}

	logs, sub, err := _Emp.contract.WatchLogs(opts, "LiquidationWithdrawn", callerRule, liquidationStatusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmpLiquidationWithdrawn)
				if err := _Emp.contract.UnpackLog(event, "LiquidationWithdrawn", log); err != nil {
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

// ParseLiquidationWithdrawn is a log parse operation binding the contract event 0xb479588a37dc7f6bac1c91587fcfc539cac4949cf26bb536ad9c8d061f00f50d.
//
// Solidity: event LiquidationWithdrawn(address indexed caller, uint256 paidToLiquidator, uint256 paidToDisputer, uint256 paidToSponsor, uint8 indexed liquidationStatus, uint256 settlementPrice)
func (_Emp *EmpFilterer) ParseLiquidationWithdrawn(log types.Log) (*EmpLiquidationWithdrawn, error) {
	event := new(EmpLiquidationWithdrawn)
	if err := _Emp.contract.UnpackLog(event, "LiquidationWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmpNewSponsorIterator is returned from FilterNewSponsor and is used to iterate over the raw logs and unpacked data for NewSponsor events raised by the Emp contract.
type EmpNewSponsorIterator struct {
	Event *EmpNewSponsor // Event containing the contract specifics and raw log

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
func (it *EmpNewSponsorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmpNewSponsor)
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
		it.Event = new(EmpNewSponsor)
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
func (it *EmpNewSponsorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmpNewSponsorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmpNewSponsor represents a NewSponsor event raised by the Emp contract.
type EmpNewSponsor struct {
	Sponsor common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewSponsor is a free log retrieval operation binding the contract event 0xf60993fa76f94c9e0a803526ee6e1314814ed4d2b0d223febf1436b36897fb37.
//
// Solidity: event NewSponsor(address indexed sponsor)
func (_Emp *EmpFilterer) FilterNewSponsor(opts *bind.FilterOpts, sponsor []common.Address) (*EmpNewSponsorIterator, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}

	logs, sub, err := _Emp.contract.FilterLogs(opts, "NewSponsor", sponsorRule)
	if err != nil {
		return nil, err
	}
	return &EmpNewSponsorIterator{contract: _Emp.contract, event: "NewSponsor", logs: logs, sub: sub}, nil
}

// WatchNewSponsor is a free log subscription operation binding the contract event 0xf60993fa76f94c9e0a803526ee6e1314814ed4d2b0d223febf1436b36897fb37.
//
// Solidity: event NewSponsor(address indexed sponsor)
func (_Emp *EmpFilterer) WatchNewSponsor(opts *bind.WatchOpts, sink chan<- *EmpNewSponsor, sponsor []common.Address) (event.Subscription, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}

	logs, sub, err := _Emp.contract.WatchLogs(opts, "NewSponsor", sponsorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmpNewSponsor)
				if err := _Emp.contract.UnpackLog(event, "NewSponsor", log); err != nil {
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

// ParseNewSponsor is a log parse operation binding the contract event 0xf60993fa76f94c9e0a803526ee6e1314814ed4d2b0d223febf1436b36897fb37.
//
// Solidity: event NewSponsor(address indexed sponsor)
func (_Emp *EmpFilterer) ParseNewSponsor(log types.Log) (*EmpNewSponsor, error) {
	event := new(EmpNewSponsor)
	if err := _Emp.contract.UnpackLog(event, "NewSponsor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmpPositionCreatedIterator is returned from FilterPositionCreated and is used to iterate over the raw logs and unpacked data for PositionCreated events raised by the Emp contract.
type EmpPositionCreatedIterator struct {
	Event *EmpPositionCreated // Event containing the contract specifics and raw log

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
func (it *EmpPositionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmpPositionCreated)
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
		it.Event = new(EmpPositionCreated)
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
func (it *EmpPositionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmpPositionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmpPositionCreated represents a PositionCreated event raised by the Emp contract.
type EmpPositionCreated struct {
	Sponsor          common.Address
	CollateralAmount *big.Int
	TokenAmount      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPositionCreated is a free log retrieval operation binding the contract event 0x4b82aa16e071a61de1a6b9aeec9edab0356331f8122c78683b469ac8e685dabc.
//
// Solidity: event PositionCreated(address indexed sponsor, uint256 indexed collateralAmount, uint256 indexed tokenAmount)
func (_Emp *EmpFilterer) FilterPositionCreated(opts *bind.FilterOpts, sponsor []common.Address, collateralAmount []*big.Int, tokenAmount []*big.Int) (*EmpPositionCreatedIterator, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var collateralAmountRule []interface{}
	for _, collateralAmountItem := range collateralAmount {
		collateralAmountRule = append(collateralAmountRule, collateralAmountItem)
	}
	var tokenAmountRule []interface{}
	for _, tokenAmountItem := range tokenAmount {
		tokenAmountRule = append(tokenAmountRule, tokenAmountItem)
	}

	logs, sub, err := _Emp.contract.FilterLogs(opts, "PositionCreated", sponsorRule, collateralAmountRule, tokenAmountRule)
	if err != nil {
		return nil, err
	}
	return &EmpPositionCreatedIterator{contract: _Emp.contract, event: "PositionCreated", logs: logs, sub: sub}, nil
}

// WatchPositionCreated is a free log subscription operation binding the contract event 0x4b82aa16e071a61de1a6b9aeec9edab0356331f8122c78683b469ac8e685dabc.
//
// Solidity: event PositionCreated(address indexed sponsor, uint256 indexed collateralAmount, uint256 indexed tokenAmount)
func (_Emp *EmpFilterer) WatchPositionCreated(opts *bind.WatchOpts, sink chan<- *EmpPositionCreated, sponsor []common.Address, collateralAmount []*big.Int, tokenAmount []*big.Int) (event.Subscription, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var collateralAmountRule []interface{}
	for _, collateralAmountItem := range collateralAmount {
		collateralAmountRule = append(collateralAmountRule, collateralAmountItem)
	}
	var tokenAmountRule []interface{}
	for _, tokenAmountItem := range tokenAmount {
		tokenAmountRule = append(tokenAmountRule, tokenAmountItem)
	}

	logs, sub, err := _Emp.contract.WatchLogs(opts, "PositionCreated", sponsorRule, collateralAmountRule, tokenAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmpPositionCreated)
				if err := _Emp.contract.UnpackLog(event, "PositionCreated", log); err != nil {
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

// ParsePositionCreated is a log parse operation binding the contract event 0x4b82aa16e071a61de1a6b9aeec9edab0356331f8122c78683b469ac8e685dabc.
//
// Solidity: event PositionCreated(address indexed sponsor, uint256 indexed collateralAmount, uint256 indexed tokenAmount)
func (_Emp *EmpFilterer) ParsePositionCreated(log types.Log) (*EmpPositionCreated, error) {
	event := new(EmpPositionCreated)
	if err := _Emp.contract.UnpackLog(event, "PositionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmpRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the Emp contract.
type EmpRedeemIterator struct {
	Event *EmpRedeem // Event containing the contract specifics and raw log

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
func (it *EmpRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmpRedeem)
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
		it.Event = new(EmpRedeem)
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
func (it *EmpRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmpRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmpRedeem represents a Redeem event raised by the Emp contract.
type EmpRedeem struct {
	Sponsor          common.Address
	CollateralAmount *big.Int
	TokenAmount      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address indexed sponsor, uint256 indexed collateralAmount, uint256 indexed tokenAmount)
func (_Emp *EmpFilterer) FilterRedeem(opts *bind.FilterOpts, sponsor []common.Address, collateralAmount []*big.Int, tokenAmount []*big.Int) (*EmpRedeemIterator, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var collateralAmountRule []interface{}
	for _, collateralAmountItem := range collateralAmount {
		collateralAmountRule = append(collateralAmountRule, collateralAmountItem)
	}
	var tokenAmountRule []interface{}
	for _, tokenAmountItem := range tokenAmount {
		tokenAmountRule = append(tokenAmountRule, tokenAmountItem)
	}

	logs, sub, err := _Emp.contract.FilterLogs(opts, "Redeem", sponsorRule, collateralAmountRule, tokenAmountRule)
	if err != nil {
		return nil, err
	}
	return &EmpRedeemIterator{contract: _Emp.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address indexed sponsor, uint256 indexed collateralAmount, uint256 indexed tokenAmount)
func (_Emp *EmpFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *EmpRedeem, sponsor []common.Address, collateralAmount []*big.Int, tokenAmount []*big.Int) (event.Subscription, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var collateralAmountRule []interface{}
	for _, collateralAmountItem := range collateralAmount {
		collateralAmountRule = append(collateralAmountRule, collateralAmountItem)
	}
	var tokenAmountRule []interface{}
	for _, tokenAmountItem := range tokenAmount {
		tokenAmountRule = append(tokenAmountRule, tokenAmountItem)
	}

	logs, sub, err := _Emp.contract.WatchLogs(opts, "Redeem", sponsorRule, collateralAmountRule, tokenAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmpRedeem)
				if err := _Emp.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address indexed sponsor, uint256 indexed collateralAmount, uint256 indexed tokenAmount)
func (_Emp *EmpFilterer) ParseRedeem(log types.Log) (*EmpRedeem, error) {
	event := new(EmpRedeem)
	if err := _Emp.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmpRegularFeesPaidIterator is returned from FilterRegularFeesPaid and is used to iterate over the raw logs and unpacked data for RegularFeesPaid events raised by the Emp contract.
type EmpRegularFeesPaidIterator struct {
	Event *EmpRegularFeesPaid // Event containing the contract specifics and raw log

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
func (it *EmpRegularFeesPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmpRegularFeesPaid)
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
		it.Event = new(EmpRegularFeesPaid)
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
func (it *EmpRegularFeesPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmpRegularFeesPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmpRegularFeesPaid represents a RegularFeesPaid event raised by the Emp contract.
type EmpRegularFeesPaid struct {
	RegularFee *big.Int
	LateFee    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRegularFeesPaid is a free log retrieval operation binding the contract event 0x19b92e73d08d517d71ec46136266e4f5d526a8cd4f8501d73713cebfe4f335ef.
//
// Solidity: event RegularFeesPaid(uint256 indexed regularFee, uint256 indexed lateFee)
func (_Emp *EmpFilterer) FilterRegularFeesPaid(opts *bind.FilterOpts, regularFee []*big.Int, lateFee []*big.Int) (*EmpRegularFeesPaidIterator, error) {

	var regularFeeRule []interface{}
	for _, regularFeeItem := range regularFee {
		regularFeeRule = append(regularFeeRule, regularFeeItem)
	}
	var lateFeeRule []interface{}
	for _, lateFeeItem := range lateFee {
		lateFeeRule = append(lateFeeRule, lateFeeItem)
	}

	logs, sub, err := _Emp.contract.FilterLogs(opts, "RegularFeesPaid", regularFeeRule, lateFeeRule)
	if err != nil {
		return nil, err
	}
	return &EmpRegularFeesPaidIterator{contract: _Emp.contract, event: "RegularFeesPaid", logs: logs, sub: sub}, nil
}

// WatchRegularFeesPaid is a free log subscription operation binding the contract event 0x19b92e73d08d517d71ec46136266e4f5d526a8cd4f8501d73713cebfe4f335ef.
//
// Solidity: event RegularFeesPaid(uint256 indexed regularFee, uint256 indexed lateFee)
func (_Emp *EmpFilterer) WatchRegularFeesPaid(opts *bind.WatchOpts, sink chan<- *EmpRegularFeesPaid, regularFee []*big.Int, lateFee []*big.Int) (event.Subscription, error) {

	var regularFeeRule []interface{}
	for _, regularFeeItem := range regularFee {
		regularFeeRule = append(regularFeeRule, regularFeeItem)
	}
	var lateFeeRule []interface{}
	for _, lateFeeItem := range lateFee {
		lateFeeRule = append(lateFeeRule, lateFeeItem)
	}

	logs, sub, err := _Emp.contract.WatchLogs(opts, "RegularFeesPaid", regularFeeRule, lateFeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmpRegularFeesPaid)
				if err := _Emp.contract.UnpackLog(event, "RegularFeesPaid", log); err != nil {
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

// ParseRegularFeesPaid is a log parse operation binding the contract event 0x19b92e73d08d517d71ec46136266e4f5d526a8cd4f8501d73713cebfe4f335ef.
//
// Solidity: event RegularFeesPaid(uint256 indexed regularFee, uint256 indexed lateFee)
func (_Emp *EmpFilterer) ParseRegularFeesPaid(log types.Log) (*EmpRegularFeesPaid, error) {
	event := new(EmpRegularFeesPaid)
	if err := _Emp.contract.UnpackLog(event, "RegularFeesPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmpRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the Emp contract.
type EmpRepayIterator struct {
	Event *EmpRepay // Event containing the contract specifics and raw log

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
func (it *EmpRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmpRepay)
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
		it.Event = new(EmpRepay)
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
func (it *EmpRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmpRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmpRepay represents a Repay event raised by the Emp contract.
type EmpRepay struct {
	Sponsor         common.Address
	NumTokensRepaid *big.Int
	NewTokenCount   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0x77c6871227e5d2dec8dadd5354f78453203e22e669cd0ec4c19d9a8c5edb31d0.
//
// Solidity: event Repay(address indexed sponsor, uint256 indexed numTokensRepaid, uint256 indexed newTokenCount)
func (_Emp *EmpFilterer) FilterRepay(opts *bind.FilterOpts, sponsor []common.Address, numTokensRepaid []*big.Int, newTokenCount []*big.Int) (*EmpRepayIterator, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var numTokensRepaidRule []interface{}
	for _, numTokensRepaidItem := range numTokensRepaid {
		numTokensRepaidRule = append(numTokensRepaidRule, numTokensRepaidItem)
	}
	var newTokenCountRule []interface{}
	for _, newTokenCountItem := range newTokenCount {
		newTokenCountRule = append(newTokenCountRule, newTokenCountItem)
	}

	logs, sub, err := _Emp.contract.FilterLogs(opts, "Repay", sponsorRule, numTokensRepaidRule, newTokenCountRule)
	if err != nil {
		return nil, err
	}
	return &EmpRepayIterator{contract: _Emp.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0x77c6871227e5d2dec8dadd5354f78453203e22e669cd0ec4c19d9a8c5edb31d0.
//
// Solidity: event Repay(address indexed sponsor, uint256 indexed numTokensRepaid, uint256 indexed newTokenCount)
func (_Emp *EmpFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *EmpRepay, sponsor []common.Address, numTokensRepaid []*big.Int, newTokenCount []*big.Int) (event.Subscription, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var numTokensRepaidRule []interface{}
	for _, numTokensRepaidItem := range numTokensRepaid {
		numTokensRepaidRule = append(numTokensRepaidRule, numTokensRepaidItem)
	}
	var newTokenCountRule []interface{}
	for _, newTokenCountItem := range newTokenCount {
		newTokenCountRule = append(newTokenCountRule, newTokenCountItem)
	}

	logs, sub, err := _Emp.contract.WatchLogs(opts, "Repay", sponsorRule, numTokensRepaidRule, newTokenCountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmpRepay)
				if err := _Emp.contract.UnpackLog(event, "Repay", log); err != nil {
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

// ParseRepay is a log parse operation binding the contract event 0x77c6871227e5d2dec8dadd5354f78453203e22e669cd0ec4c19d9a8c5edb31d0.
//
// Solidity: event Repay(address indexed sponsor, uint256 indexed numTokensRepaid, uint256 indexed newTokenCount)
func (_Emp *EmpFilterer) ParseRepay(log types.Log) (*EmpRepay, error) {
	event := new(EmpRepay)
	if err := _Emp.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmpRequestTransferPositionIterator is returned from FilterRequestTransferPosition and is used to iterate over the raw logs and unpacked data for RequestTransferPosition events raised by the Emp contract.
type EmpRequestTransferPositionIterator struct {
	Event *EmpRequestTransferPosition // Event containing the contract specifics and raw log

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
func (it *EmpRequestTransferPositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmpRequestTransferPosition)
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
		it.Event = new(EmpRequestTransferPosition)
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
func (it *EmpRequestTransferPositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmpRequestTransferPositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmpRequestTransferPosition represents a RequestTransferPosition event raised by the Emp contract.
type EmpRequestTransferPosition struct {
	OldSponsor common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRequestTransferPosition is a free log retrieval operation binding the contract event 0xbf457c80c8bf299d5c48272c4c1168bf87b33d83b13f0ab9aac332ce1161ed1e.
//
// Solidity: event RequestTransferPosition(address indexed oldSponsor)
func (_Emp *EmpFilterer) FilterRequestTransferPosition(opts *bind.FilterOpts, oldSponsor []common.Address) (*EmpRequestTransferPositionIterator, error) {

	var oldSponsorRule []interface{}
	for _, oldSponsorItem := range oldSponsor {
		oldSponsorRule = append(oldSponsorRule, oldSponsorItem)
	}

	logs, sub, err := _Emp.contract.FilterLogs(opts, "RequestTransferPosition", oldSponsorRule)
	if err != nil {
		return nil, err
	}
	return &EmpRequestTransferPositionIterator{contract: _Emp.contract, event: "RequestTransferPosition", logs: logs, sub: sub}, nil
}

// WatchRequestTransferPosition is a free log subscription operation binding the contract event 0xbf457c80c8bf299d5c48272c4c1168bf87b33d83b13f0ab9aac332ce1161ed1e.
//
// Solidity: event RequestTransferPosition(address indexed oldSponsor)
func (_Emp *EmpFilterer) WatchRequestTransferPosition(opts *bind.WatchOpts, sink chan<- *EmpRequestTransferPosition, oldSponsor []common.Address) (event.Subscription, error) {

	var oldSponsorRule []interface{}
	for _, oldSponsorItem := range oldSponsor {
		oldSponsorRule = append(oldSponsorRule, oldSponsorItem)
	}

	logs, sub, err := _Emp.contract.WatchLogs(opts, "RequestTransferPosition", oldSponsorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmpRequestTransferPosition)
				if err := _Emp.contract.UnpackLog(event, "RequestTransferPosition", log); err != nil {
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

// ParseRequestTransferPosition is a log parse operation binding the contract event 0xbf457c80c8bf299d5c48272c4c1168bf87b33d83b13f0ab9aac332ce1161ed1e.
//
// Solidity: event RequestTransferPosition(address indexed oldSponsor)
func (_Emp *EmpFilterer) ParseRequestTransferPosition(log types.Log) (*EmpRequestTransferPosition, error) {
	event := new(EmpRequestTransferPosition)
	if err := _Emp.contract.UnpackLog(event, "RequestTransferPosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmpRequestTransferPositionCanceledIterator is returned from FilterRequestTransferPositionCanceled and is used to iterate over the raw logs and unpacked data for RequestTransferPositionCanceled events raised by the Emp contract.
type EmpRequestTransferPositionCanceledIterator struct {
	Event *EmpRequestTransferPositionCanceled // Event containing the contract specifics and raw log

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
func (it *EmpRequestTransferPositionCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmpRequestTransferPositionCanceled)
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
		it.Event = new(EmpRequestTransferPositionCanceled)
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
func (it *EmpRequestTransferPositionCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmpRequestTransferPositionCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmpRequestTransferPositionCanceled represents a RequestTransferPositionCanceled event raised by the Emp contract.
type EmpRequestTransferPositionCanceled struct {
	OldSponsor common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRequestTransferPositionCanceled is a free log retrieval operation binding the contract event 0x2e5702420c76e041698ad7ba57a9ff5cadccf647ea8d96e6007a40b5b2662f56.
//
// Solidity: event RequestTransferPositionCanceled(address indexed oldSponsor)
func (_Emp *EmpFilterer) FilterRequestTransferPositionCanceled(opts *bind.FilterOpts, oldSponsor []common.Address) (*EmpRequestTransferPositionCanceledIterator, error) {

	var oldSponsorRule []interface{}
	for _, oldSponsorItem := range oldSponsor {
		oldSponsorRule = append(oldSponsorRule, oldSponsorItem)
	}

	logs, sub, err := _Emp.contract.FilterLogs(opts, "RequestTransferPositionCanceled", oldSponsorRule)
	if err != nil {
		return nil, err
	}
	return &EmpRequestTransferPositionCanceledIterator{contract: _Emp.contract, event: "RequestTransferPositionCanceled", logs: logs, sub: sub}, nil
}

// WatchRequestTransferPositionCanceled is a free log subscription operation binding the contract event 0x2e5702420c76e041698ad7ba57a9ff5cadccf647ea8d96e6007a40b5b2662f56.
//
// Solidity: event RequestTransferPositionCanceled(address indexed oldSponsor)
func (_Emp *EmpFilterer) WatchRequestTransferPositionCanceled(opts *bind.WatchOpts, sink chan<- *EmpRequestTransferPositionCanceled, oldSponsor []common.Address) (event.Subscription, error) {

	var oldSponsorRule []interface{}
	for _, oldSponsorItem := range oldSponsor {
		oldSponsorRule = append(oldSponsorRule, oldSponsorItem)
	}

	logs, sub, err := _Emp.contract.WatchLogs(opts, "RequestTransferPositionCanceled", oldSponsorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmpRequestTransferPositionCanceled)
				if err := _Emp.contract.UnpackLog(event, "RequestTransferPositionCanceled", log); err != nil {
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

// ParseRequestTransferPositionCanceled is a log parse operation binding the contract event 0x2e5702420c76e041698ad7ba57a9ff5cadccf647ea8d96e6007a40b5b2662f56.
//
// Solidity: event RequestTransferPositionCanceled(address indexed oldSponsor)
func (_Emp *EmpFilterer) ParseRequestTransferPositionCanceled(log types.Log) (*EmpRequestTransferPositionCanceled, error) {
	event := new(EmpRequestTransferPositionCanceled)
	if err := _Emp.contract.UnpackLog(event, "RequestTransferPositionCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmpRequestTransferPositionExecutedIterator is returned from FilterRequestTransferPositionExecuted and is used to iterate over the raw logs and unpacked data for RequestTransferPositionExecuted events raised by the Emp contract.
type EmpRequestTransferPositionExecutedIterator struct {
	Event *EmpRequestTransferPositionExecuted // Event containing the contract specifics and raw log

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
func (it *EmpRequestTransferPositionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmpRequestTransferPositionExecuted)
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
		it.Event = new(EmpRequestTransferPositionExecuted)
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
func (it *EmpRequestTransferPositionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmpRequestTransferPositionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmpRequestTransferPositionExecuted represents a RequestTransferPositionExecuted event raised by the Emp contract.
type EmpRequestTransferPositionExecuted struct {
	OldSponsor common.Address
	NewSponsor common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRequestTransferPositionExecuted is a free log retrieval operation binding the contract event 0xf1a2dcf23621f1a96185c79d39a5776b5ba3dadbea70c5aa86d84c17c7e9418e.
//
// Solidity: event RequestTransferPositionExecuted(address indexed oldSponsor, address indexed newSponsor)
func (_Emp *EmpFilterer) FilterRequestTransferPositionExecuted(opts *bind.FilterOpts, oldSponsor []common.Address, newSponsor []common.Address) (*EmpRequestTransferPositionExecutedIterator, error) {

	var oldSponsorRule []interface{}
	for _, oldSponsorItem := range oldSponsor {
		oldSponsorRule = append(oldSponsorRule, oldSponsorItem)
	}
	var newSponsorRule []interface{}
	for _, newSponsorItem := range newSponsor {
		newSponsorRule = append(newSponsorRule, newSponsorItem)
	}

	logs, sub, err := _Emp.contract.FilterLogs(opts, "RequestTransferPositionExecuted", oldSponsorRule, newSponsorRule)
	if err != nil {
		return nil, err
	}
	return &EmpRequestTransferPositionExecutedIterator{contract: _Emp.contract, event: "RequestTransferPositionExecuted", logs: logs, sub: sub}, nil
}

// WatchRequestTransferPositionExecuted is a free log subscription operation binding the contract event 0xf1a2dcf23621f1a96185c79d39a5776b5ba3dadbea70c5aa86d84c17c7e9418e.
//
// Solidity: event RequestTransferPositionExecuted(address indexed oldSponsor, address indexed newSponsor)
func (_Emp *EmpFilterer) WatchRequestTransferPositionExecuted(opts *bind.WatchOpts, sink chan<- *EmpRequestTransferPositionExecuted, oldSponsor []common.Address, newSponsor []common.Address) (event.Subscription, error) {

	var oldSponsorRule []interface{}
	for _, oldSponsorItem := range oldSponsor {
		oldSponsorRule = append(oldSponsorRule, oldSponsorItem)
	}
	var newSponsorRule []interface{}
	for _, newSponsorItem := range newSponsor {
		newSponsorRule = append(newSponsorRule, newSponsorItem)
	}

	logs, sub, err := _Emp.contract.WatchLogs(opts, "RequestTransferPositionExecuted", oldSponsorRule, newSponsorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmpRequestTransferPositionExecuted)
				if err := _Emp.contract.UnpackLog(event, "RequestTransferPositionExecuted", log); err != nil {
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

// ParseRequestTransferPositionExecuted is a log parse operation binding the contract event 0xf1a2dcf23621f1a96185c79d39a5776b5ba3dadbea70c5aa86d84c17c7e9418e.
//
// Solidity: event RequestTransferPositionExecuted(address indexed oldSponsor, address indexed newSponsor)
func (_Emp *EmpFilterer) ParseRequestTransferPositionExecuted(log types.Log) (*EmpRequestTransferPositionExecuted, error) {
	event := new(EmpRequestTransferPositionExecuted)
	if err := _Emp.contract.UnpackLog(event, "RequestTransferPositionExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmpRequestWithdrawalIterator is returned from FilterRequestWithdrawal and is used to iterate over the raw logs and unpacked data for RequestWithdrawal events raised by the Emp contract.
type EmpRequestWithdrawalIterator struct {
	Event *EmpRequestWithdrawal // Event containing the contract specifics and raw log

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
func (it *EmpRequestWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmpRequestWithdrawal)
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
		it.Event = new(EmpRequestWithdrawal)
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
func (it *EmpRequestWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmpRequestWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmpRequestWithdrawal represents a RequestWithdrawal event raised by the Emp contract.
type EmpRequestWithdrawal struct {
	Sponsor          common.Address
	CollateralAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRequestWithdrawal is a free log retrieval operation binding the contract event 0xd33b726e11d2c5d38e6702b16613df0160a07f7ba5185455ee3c45d0494fab11.
//
// Solidity: event RequestWithdrawal(address indexed sponsor, uint256 indexed collateralAmount)
func (_Emp *EmpFilterer) FilterRequestWithdrawal(opts *bind.FilterOpts, sponsor []common.Address, collateralAmount []*big.Int) (*EmpRequestWithdrawalIterator, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var collateralAmountRule []interface{}
	for _, collateralAmountItem := range collateralAmount {
		collateralAmountRule = append(collateralAmountRule, collateralAmountItem)
	}

	logs, sub, err := _Emp.contract.FilterLogs(opts, "RequestWithdrawal", sponsorRule, collateralAmountRule)
	if err != nil {
		return nil, err
	}
	return &EmpRequestWithdrawalIterator{contract: _Emp.contract, event: "RequestWithdrawal", logs: logs, sub: sub}, nil
}

// WatchRequestWithdrawal is a free log subscription operation binding the contract event 0xd33b726e11d2c5d38e6702b16613df0160a07f7ba5185455ee3c45d0494fab11.
//
// Solidity: event RequestWithdrawal(address indexed sponsor, uint256 indexed collateralAmount)
func (_Emp *EmpFilterer) WatchRequestWithdrawal(opts *bind.WatchOpts, sink chan<- *EmpRequestWithdrawal, sponsor []common.Address, collateralAmount []*big.Int) (event.Subscription, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var collateralAmountRule []interface{}
	for _, collateralAmountItem := range collateralAmount {
		collateralAmountRule = append(collateralAmountRule, collateralAmountItem)
	}

	logs, sub, err := _Emp.contract.WatchLogs(opts, "RequestWithdrawal", sponsorRule, collateralAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmpRequestWithdrawal)
				if err := _Emp.contract.UnpackLog(event, "RequestWithdrawal", log); err != nil {
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

// ParseRequestWithdrawal is a log parse operation binding the contract event 0xd33b726e11d2c5d38e6702b16613df0160a07f7ba5185455ee3c45d0494fab11.
//
// Solidity: event RequestWithdrawal(address indexed sponsor, uint256 indexed collateralAmount)
func (_Emp *EmpFilterer) ParseRequestWithdrawal(log types.Log) (*EmpRequestWithdrawal, error) {
	event := new(EmpRequestWithdrawal)
	if err := _Emp.contract.UnpackLog(event, "RequestWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmpRequestWithdrawalCanceledIterator is returned from FilterRequestWithdrawalCanceled and is used to iterate over the raw logs and unpacked data for RequestWithdrawalCanceled events raised by the Emp contract.
type EmpRequestWithdrawalCanceledIterator struct {
	Event *EmpRequestWithdrawalCanceled // Event containing the contract specifics and raw log

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
func (it *EmpRequestWithdrawalCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmpRequestWithdrawalCanceled)
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
		it.Event = new(EmpRequestWithdrawalCanceled)
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
func (it *EmpRequestWithdrawalCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmpRequestWithdrawalCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmpRequestWithdrawalCanceled represents a RequestWithdrawalCanceled event raised by the Emp contract.
type EmpRequestWithdrawalCanceled struct {
	Sponsor          common.Address
	CollateralAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRequestWithdrawalCanceled is a free log retrieval operation binding the contract event 0x74d8a3658feb89d1a5c335229bbbfc3bbcfaf492769feb7aa4cd2d92efeaf691.
//
// Solidity: event RequestWithdrawalCanceled(address indexed sponsor, uint256 indexed collateralAmount)
func (_Emp *EmpFilterer) FilterRequestWithdrawalCanceled(opts *bind.FilterOpts, sponsor []common.Address, collateralAmount []*big.Int) (*EmpRequestWithdrawalCanceledIterator, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var collateralAmountRule []interface{}
	for _, collateralAmountItem := range collateralAmount {
		collateralAmountRule = append(collateralAmountRule, collateralAmountItem)
	}

	logs, sub, err := _Emp.contract.FilterLogs(opts, "RequestWithdrawalCanceled", sponsorRule, collateralAmountRule)
	if err != nil {
		return nil, err
	}
	return &EmpRequestWithdrawalCanceledIterator{contract: _Emp.contract, event: "RequestWithdrawalCanceled", logs: logs, sub: sub}, nil
}

// WatchRequestWithdrawalCanceled is a free log subscription operation binding the contract event 0x74d8a3658feb89d1a5c335229bbbfc3bbcfaf492769feb7aa4cd2d92efeaf691.
//
// Solidity: event RequestWithdrawalCanceled(address indexed sponsor, uint256 indexed collateralAmount)
func (_Emp *EmpFilterer) WatchRequestWithdrawalCanceled(opts *bind.WatchOpts, sink chan<- *EmpRequestWithdrawalCanceled, sponsor []common.Address, collateralAmount []*big.Int) (event.Subscription, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var collateralAmountRule []interface{}
	for _, collateralAmountItem := range collateralAmount {
		collateralAmountRule = append(collateralAmountRule, collateralAmountItem)
	}

	logs, sub, err := _Emp.contract.WatchLogs(opts, "RequestWithdrawalCanceled", sponsorRule, collateralAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmpRequestWithdrawalCanceled)
				if err := _Emp.contract.UnpackLog(event, "RequestWithdrawalCanceled", log); err != nil {
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

// ParseRequestWithdrawalCanceled is a log parse operation binding the contract event 0x74d8a3658feb89d1a5c335229bbbfc3bbcfaf492769feb7aa4cd2d92efeaf691.
//
// Solidity: event RequestWithdrawalCanceled(address indexed sponsor, uint256 indexed collateralAmount)
func (_Emp *EmpFilterer) ParseRequestWithdrawalCanceled(log types.Log) (*EmpRequestWithdrawalCanceled, error) {
	event := new(EmpRequestWithdrawalCanceled)
	if err := _Emp.contract.UnpackLog(event, "RequestWithdrawalCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmpRequestWithdrawalExecutedIterator is returned from FilterRequestWithdrawalExecuted and is used to iterate over the raw logs and unpacked data for RequestWithdrawalExecuted events raised by the Emp contract.
type EmpRequestWithdrawalExecutedIterator struct {
	Event *EmpRequestWithdrawalExecuted // Event containing the contract specifics and raw log

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
func (it *EmpRequestWithdrawalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmpRequestWithdrawalExecuted)
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
		it.Event = new(EmpRequestWithdrawalExecuted)
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
func (it *EmpRequestWithdrawalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmpRequestWithdrawalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmpRequestWithdrawalExecuted represents a RequestWithdrawalExecuted event raised by the Emp contract.
type EmpRequestWithdrawalExecuted struct {
	Sponsor          common.Address
	CollateralAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRequestWithdrawalExecuted is a free log retrieval operation binding the contract event 0xc86c3298cb79f486674dca87d9247e88b76146160e7d412cc59b26b14c358a68.
//
// Solidity: event RequestWithdrawalExecuted(address indexed sponsor, uint256 indexed collateralAmount)
func (_Emp *EmpFilterer) FilterRequestWithdrawalExecuted(opts *bind.FilterOpts, sponsor []common.Address, collateralAmount []*big.Int) (*EmpRequestWithdrawalExecutedIterator, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var collateralAmountRule []interface{}
	for _, collateralAmountItem := range collateralAmount {
		collateralAmountRule = append(collateralAmountRule, collateralAmountItem)
	}

	logs, sub, err := _Emp.contract.FilterLogs(opts, "RequestWithdrawalExecuted", sponsorRule, collateralAmountRule)
	if err != nil {
		return nil, err
	}
	return &EmpRequestWithdrawalExecutedIterator{contract: _Emp.contract, event: "RequestWithdrawalExecuted", logs: logs, sub: sub}, nil
}

// WatchRequestWithdrawalExecuted is a free log subscription operation binding the contract event 0xc86c3298cb79f486674dca87d9247e88b76146160e7d412cc59b26b14c358a68.
//
// Solidity: event RequestWithdrawalExecuted(address indexed sponsor, uint256 indexed collateralAmount)
func (_Emp *EmpFilterer) WatchRequestWithdrawalExecuted(opts *bind.WatchOpts, sink chan<- *EmpRequestWithdrawalExecuted, sponsor []common.Address, collateralAmount []*big.Int) (event.Subscription, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var collateralAmountRule []interface{}
	for _, collateralAmountItem := range collateralAmount {
		collateralAmountRule = append(collateralAmountRule, collateralAmountItem)
	}

	logs, sub, err := _Emp.contract.WatchLogs(opts, "RequestWithdrawalExecuted", sponsorRule, collateralAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmpRequestWithdrawalExecuted)
				if err := _Emp.contract.UnpackLog(event, "RequestWithdrawalExecuted", log); err != nil {
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

// ParseRequestWithdrawalExecuted is a log parse operation binding the contract event 0xc86c3298cb79f486674dca87d9247e88b76146160e7d412cc59b26b14c358a68.
//
// Solidity: event RequestWithdrawalExecuted(address indexed sponsor, uint256 indexed collateralAmount)
func (_Emp *EmpFilterer) ParseRequestWithdrawalExecuted(log types.Log) (*EmpRequestWithdrawalExecuted, error) {
	event := new(EmpRequestWithdrawalExecuted)
	if err := _Emp.contract.UnpackLog(event, "RequestWithdrawalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmpSettleExpiredPositionIterator is returned from FilterSettleExpiredPosition and is used to iterate over the raw logs and unpacked data for SettleExpiredPosition events raised by the Emp contract.
type EmpSettleExpiredPositionIterator struct {
	Event *EmpSettleExpiredPosition // Event containing the contract specifics and raw log

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
func (it *EmpSettleExpiredPositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmpSettleExpiredPosition)
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
		it.Event = new(EmpSettleExpiredPosition)
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
func (it *EmpSettleExpiredPositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmpSettleExpiredPositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmpSettleExpiredPosition represents a SettleExpiredPosition event raised by the Emp contract.
type EmpSettleExpiredPosition struct {
	Caller             common.Address
	CollateralReturned *big.Int
	TokensBurned       *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSettleExpiredPosition is a free log retrieval operation binding the contract event 0x9d349c102bec959fb7f20f9a3621e015819d3ae4ed6e9afd1f56a69d58456006.
//
// Solidity: event SettleExpiredPosition(address indexed caller, uint256 indexed collateralReturned, uint256 indexed tokensBurned)
func (_Emp *EmpFilterer) FilterSettleExpiredPosition(opts *bind.FilterOpts, caller []common.Address, collateralReturned []*big.Int, tokensBurned []*big.Int) (*EmpSettleExpiredPositionIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var collateralReturnedRule []interface{}
	for _, collateralReturnedItem := range collateralReturned {
		collateralReturnedRule = append(collateralReturnedRule, collateralReturnedItem)
	}
	var tokensBurnedRule []interface{}
	for _, tokensBurnedItem := range tokensBurned {
		tokensBurnedRule = append(tokensBurnedRule, tokensBurnedItem)
	}

	logs, sub, err := _Emp.contract.FilterLogs(opts, "SettleExpiredPosition", callerRule, collateralReturnedRule, tokensBurnedRule)
	if err != nil {
		return nil, err
	}
	return &EmpSettleExpiredPositionIterator{contract: _Emp.contract, event: "SettleExpiredPosition", logs: logs, sub: sub}, nil
}

// WatchSettleExpiredPosition is a free log subscription operation binding the contract event 0x9d349c102bec959fb7f20f9a3621e015819d3ae4ed6e9afd1f56a69d58456006.
//
// Solidity: event SettleExpiredPosition(address indexed caller, uint256 indexed collateralReturned, uint256 indexed tokensBurned)
func (_Emp *EmpFilterer) WatchSettleExpiredPosition(opts *bind.WatchOpts, sink chan<- *EmpSettleExpiredPosition, caller []common.Address, collateralReturned []*big.Int, tokensBurned []*big.Int) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var collateralReturnedRule []interface{}
	for _, collateralReturnedItem := range collateralReturned {
		collateralReturnedRule = append(collateralReturnedRule, collateralReturnedItem)
	}
	var tokensBurnedRule []interface{}
	for _, tokensBurnedItem := range tokensBurned {
		tokensBurnedRule = append(tokensBurnedRule, tokensBurnedItem)
	}

	logs, sub, err := _Emp.contract.WatchLogs(opts, "SettleExpiredPosition", callerRule, collateralReturnedRule, tokensBurnedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmpSettleExpiredPosition)
				if err := _Emp.contract.UnpackLog(event, "SettleExpiredPosition", log); err != nil {
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

// ParseSettleExpiredPosition is a log parse operation binding the contract event 0x9d349c102bec959fb7f20f9a3621e015819d3ae4ed6e9afd1f56a69d58456006.
//
// Solidity: event SettleExpiredPosition(address indexed caller, uint256 indexed collateralReturned, uint256 indexed tokensBurned)
func (_Emp *EmpFilterer) ParseSettleExpiredPosition(log types.Log) (*EmpSettleExpiredPosition, error) {
	event := new(EmpSettleExpiredPosition)
	if err := _Emp.contract.UnpackLog(event, "SettleExpiredPosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmpWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the Emp contract.
type EmpWithdrawalIterator struct {
	Event *EmpWithdrawal // Event containing the contract specifics and raw log

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
func (it *EmpWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmpWithdrawal)
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
		it.Event = new(EmpWithdrawal)
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
func (it *EmpWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmpWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmpWithdrawal represents a Withdrawal event raised by the Emp contract.
type EmpWithdrawal struct {
	Sponsor          common.Address
	CollateralAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sponsor, uint256 indexed collateralAmount)
func (_Emp *EmpFilterer) FilterWithdrawal(opts *bind.FilterOpts, sponsor []common.Address, collateralAmount []*big.Int) (*EmpWithdrawalIterator, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var collateralAmountRule []interface{}
	for _, collateralAmountItem := range collateralAmount {
		collateralAmountRule = append(collateralAmountRule, collateralAmountItem)
	}

	logs, sub, err := _Emp.contract.FilterLogs(opts, "Withdrawal", sponsorRule, collateralAmountRule)
	if err != nil {
		return nil, err
	}
	return &EmpWithdrawalIterator{contract: _Emp.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sponsor, uint256 indexed collateralAmount)
func (_Emp *EmpFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *EmpWithdrawal, sponsor []common.Address, collateralAmount []*big.Int) (event.Subscription, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var collateralAmountRule []interface{}
	for _, collateralAmountItem := range collateralAmount {
		collateralAmountRule = append(collateralAmountRule, collateralAmountItem)
	}

	logs, sub, err := _Emp.contract.WatchLogs(opts, "Withdrawal", sponsorRule, collateralAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmpWithdrawal)
				if err := _Emp.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sponsor, uint256 indexed collateralAmount)
func (_Emp *EmpFilterer) ParseWithdrawal(log types.Log) (*EmpWithdrawal, error) {
	event := new(EmpWithdrawal)
	if err := _Emp.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
