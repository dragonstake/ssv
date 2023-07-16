// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// Authorization is an auto generated low-level Go binding around an user-defined struct.
type Authorization struct {
	RegisterOperator  bool
	RegisterValidator bool
}

// ISSVNetworkCoreCluster is an auto generated low-level Go binding around an user-defined struct.
type ISSVNetworkCoreCluster struct {
	ValidatorCount  uint32
	NetworkFeeIndex uint64
	Index           uint64
	Active          bool
	Balance         *big.Int
}

// SSVNetworkMetaData contains all meta data concerning the SSVNetwork contract.
var SSVNetworkMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ApprovalNotWithinTimeframe\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClusterAlreadyEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClusterDoesNotExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClusterIsLiquidated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClusterNotLiquidatable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedValidatorLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeExceedsIncreaseLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeIncreaseNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectClusterState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectValidatorState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOperatorIdsLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKeyLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewBlockPeriodIsBelowMinimum\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeeDeclared\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorsListNotUnique\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameFeeChangeNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TargetModuleDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsortedOperatorsList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorDoesNotExist\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ClusterDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ClusterLiquidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ClusterReactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ClusterWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"DeclareOperatorFeePeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"ExecuteOperatorFeePeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"}],\"name\":\"FeeRecipientAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"LiquidationThresholdPeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MinimumLiquidationCollateralUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"NetworkEarningsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"NetworkFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"OperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"OperatorFeeCancellationDeclared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"OperatorFeeDeclared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"OperatorFeeExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"OperatorFeeIncreaseLimitUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"OperatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"whitelisted\",\"type\":\"address\"}],\"name\":\"OperatorWhitelistUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"OperatorWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"shares\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ValidatorRemoved\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"cancelDeclaredOperatorFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"declareOperatorFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"executeOperatorFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"getRegisterAuth\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"registerOperator\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"registerValidator\",\"type\":\"bool\"}],\"internalType\":\"structAuthorization\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"contractIFnSSVOperators\",\"name\":\"ssvOperators_\",\"type\":\"address\"},{\"internalType\":\"contractIFnSSVClusters\",\"name\":\"ssvClusters_\",\"type\":\"address\"},{\"internalType\":\"contractIFnSSVDAO\",\"name\":\"ssvDAO_\",\"type\":\"address\"},{\"internalType\":\"contractIFnSSVViews\",\"name\":\"ssvViews_\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"minimumBlocksBeforeLiquidation_\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"minimumLiquidationCollateral_\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"validatorsPerOperatorLimit_\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"declareOperatorFeePeriod_\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executeOperatorFeePeriod_\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"operatorMaxFeeIncrease_\",\"type\":\"uint64\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"reactivate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"reduceOperatorFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"registerOperator\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes\",\"name\":\"sharesData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"registerValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"}],\"name\":\"setFeeRecipientAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"whitelisted\",\"type\":\"address\"}],\"name\":\"setOperatorWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"registerOperator\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"registerValidator\",\"type\":\"bool\"}],\"internalType\":\"structAuthorization\",\"name\":\"auth\",\"type\":\"tuple\"}],\"name\":\"setRegisterAuth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"timeInSeconds\",\"type\":\"uint64\"}],\"name\":\"updateDeclareOperatorFeePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"timeInSeconds\",\"type\":\"uint64\"}],\"name\":\"updateExecuteOperatorFeePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"blocks\",\"type\":\"uint64\"}],\"name\":\"updateLiquidationThresholdPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"updateMinimumLiquidationCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"updateNetworkFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"percentage\",\"type\":\"uint64\"}],\"name\":\"updateOperatorFeeIncreaseLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSSVModules\",\"name\":\"moduleId\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"moduleAddress\",\"type\":\"address\"}],\"name\":\"upgradeModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawNetworkEarnings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawOperatorEarnings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"withdrawOperatorEarnings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SSVNetworkABI is the input ABI used to generate the binding from.
// Deprecated: Use SSVNetworkMetaData.ABI instead.
var SSVNetworkABI = SSVNetworkMetaData.ABI

// SSVNetwork is an auto generated Go binding around an Ethereum contract.
type SSVNetwork struct {
	SSVNetworkCaller     // Read-only binding to the contract
	SSVNetworkTransactor // Write-only binding to the contract
	SSVNetworkFilterer   // Log filterer for contract events
}

// SSVNetworkCaller is an auto generated read-only Go binding around an Ethereum contract.
type SSVNetworkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SSVNetworkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SSVNetworkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SSVNetworkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SSVNetworkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SSVNetworkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SSVNetworkSession struct {
	Contract     *SSVNetwork       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SSVNetworkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SSVNetworkCallerSession struct {
	Contract *SSVNetworkCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SSVNetworkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SSVNetworkTransactorSession struct {
	Contract     *SSVNetworkTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SSVNetworkRaw is an auto generated low-level Go binding around an Ethereum contract.
type SSVNetworkRaw struct {
	Contract *SSVNetwork // Generic contract binding to access the raw methods on
}

// SSVNetworkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SSVNetworkCallerRaw struct {
	Contract *SSVNetworkCaller // Generic read-only contract binding to access the raw methods on
}

// SSVNetworkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SSVNetworkTransactorRaw struct {
	Contract *SSVNetworkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSSVNetwork creates a new instance of SSVNetwork, bound to a specific deployed contract.
func NewSSVNetwork(address common.Address, backend bind.ContractBackend) (*SSVNetwork, error) {
	contract, err := bindSSVNetwork(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SSVNetwork{SSVNetworkCaller: SSVNetworkCaller{contract: contract}, SSVNetworkTransactor: SSVNetworkTransactor{contract: contract}, SSVNetworkFilterer: SSVNetworkFilterer{contract: contract}}, nil
}

// NewSSVNetworkCaller creates a new read-only instance of SSVNetwork, bound to a specific deployed contract.
func NewSSVNetworkCaller(address common.Address, caller bind.ContractCaller) (*SSVNetworkCaller, error) {
	contract, err := bindSSVNetwork(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SSVNetworkCaller{contract: contract}, nil
}

// NewSSVNetworkTransactor creates a new write-only instance of SSVNetwork, bound to a specific deployed contract.
func NewSSVNetworkTransactor(address common.Address, transactor bind.ContractTransactor) (*SSVNetworkTransactor, error) {
	contract, err := bindSSVNetwork(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SSVNetworkTransactor{contract: contract}, nil
}

// NewSSVNetworkFilterer creates a new log filterer instance of SSVNetwork, bound to a specific deployed contract.
func NewSSVNetworkFilterer(address common.Address, filterer bind.ContractFilterer) (*SSVNetworkFilterer, error) {
	contract, err := bindSSVNetwork(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SSVNetworkFilterer{contract: contract}, nil
}

// bindSSVNetwork binds a generic wrapper to an already deployed contract.
func bindSSVNetwork(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SSVNetworkMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SSVNetwork *SSVNetworkRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SSVNetwork.Contract.SSVNetworkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SSVNetwork *SSVNetworkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SSVNetwork.Contract.SSVNetworkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SSVNetwork *SSVNetworkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SSVNetwork.Contract.SSVNetworkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SSVNetwork *SSVNetworkCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SSVNetwork.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SSVNetwork *SSVNetworkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SSVNetwork.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SSVNetwork *SSVNetworkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SSVNetwork.Contract.contract.Transact(opts, method, params...)
}

// GetRegisterAuth is a free data retrieval call binding the contract method 0x7398ca6c.
//
// Solidity: function getRegisterAuth(address userAddress) view returns((bool,bool))
func (_SSVNetwork *SSVNetworkCaller) GetRegisterAuth(opts *bind.CallOpts, userAddress common.Address) (Authorization, error) {
	var out []interface{}
	err := _SSVNetwork.contract.Call(opts, &out, "getRegisterAuth", userAddress)

	if err != nil {
		return *new(Authorization), err
	}

	out0 := *abi.ConvertType(out[0], new(Authorization)).(*Authorization)

	return out0, err

}

// GetRegisterAuth is a free data retrieval call binding the contract method 0x7398ca6c.
//
// Solidity: function getRegisterAuth(address userAddress) view returns((bool,bool))
func (_SSVNetwork *SSVNetworkSession) GetRegisterAuth(userAddress common.Address) (Authorization, error) {
	return _SSVNetwork.Contract.GetRegisterAuth(&_SSVNetwork.CallOpts, userAddress)
}

// GetRegisterAuth is a free data retrieval call binding the contract method 0x7398ca6c.
//
// Solidity: function getRegisterAuth(address userAddress) view returns((bool,bool))
func (_SSVNetwork *SSVNetworkCallerSession) GetRegisterAuth(userAddress common.Address) (Authorization, error) {
	return _SSVNetwork.Contract.GetRegisterAuth(&_SSVNetwork.CallOpts, userAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SSVNetwork *SSVNetworkCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SSVNetwork.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SSVNetwork *SSVNetworkSession) Owner() (common.Address, error) {
	return _SSVNetwork.Contract.Owner(&_SSVNetwork.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SSVNetwork *SSVNetworkCallerSession) Owner() (common.Address, error) {
	return _SSVNetwork.Contract.Owner(&_SSVNetwork.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_SSVNetwork *SSVNetworkCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SSVNetwork.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_SSVNetwork *SSVNetworkSession) PendingOwner() (common.Address, error) {
	return _SSVNetwork.Contract.PendingOwner(&_SSVNetwork.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_SSVNetwork *SSVNetworkCallerSession) PendingOwner() (common.Address, error) {
	return _SSVNetwork.Contract.PendingOwner(&_SSVNetwork.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SSVNetwork *SSVNetworkCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SSVNetwork.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SSVNetwork *SSVNetworkSession) ProxiableUUID() ([32]byte, error) {
	return _SSVNetwork.Contract.ProxiableUUID(&_SSVNetwork.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SSVNetwork *SSVNetworkCallerSession) ProxiableUUID() ([32]byte, error) {
	return _SSVNetwork.Contract.ProxiableUUID(&_SSVNetwork.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SSVNetwork *SSVNetworkTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SSVNetwork *SSVNetworkSession) AcceptOwnership() (*types.Transaction, error) {
	return _SSVNetwork.Contract.AcceptOwnership(&_SSVNetwork.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SSVNetwork *SSVNetworkTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _SSVNetwork.Contract.AcceptOwnership(&_SSVNetwork.TransactOpts)
}

// CancelDeclaredOperatorFee is a paid mutator transaction binding the contract method 0x23d68a6d.
//
// Solidity: function cancelDeclaredOperatorFee(uint64 operatorId) returns()
func (_SSVNetwork *SSVNetworkTransactor) CancelDeclaredOperatorFee(opts *bind.TransactOpts, operatorId uint64) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "cancelDeclaredOperatorFee", operatorId)
}

// CancelDeclaredOperatorFee is a paid mutator transaction binding the contract method 0x23d68a6d.
//
// Solidity: function cancelDeclaredOperatorFee(uint64 operatorId) returns()
func (_SSVNetwork *SSVNetworkSession) CancelDeclaredOperatorFee(operatorId uint64) (*types.Transaction, error) {
	return _SSVNetwork.Contract.CancelDeclaredOperatorFee(&_SSVNetwork.TransactOpts, operatorId)
}

// CancelDeclaredOperatorFee is a paid mutator transaction binding the contract method 0x23d68a6d.
//
// Solidity: function cancelDeclaredOperatorFee(uint64 operatorId) returns()
func (_SSVNetwork *SSVNetworkTransactorSession) CancelDeclaredOperatorFee(operatorId uint64) (*types.Transaction, error) {
	return _SSVNetwork.Contract.CancelDeclaredOperatorFee(&_SSVNetwork.TransactOpts, operatorId)
}

// DeclareOperatorFee is a paid mutator transaction binding the contract method 0xb317c35f.
//
// Solidity: function declareOperatorFee(uint64 operatorId, uint256 fee) returns()
func (_SSVNetwork *SSVNetworkTransactor) DeclareOperatorFee(opts *bind.TransactOpts, operatorId uint64, fee *big.Int) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "declareOperatorFee", operatorId, fee)
}

// DeclareOperatorFee is a paid mutator transaction binding the contract method 0xb317c35f.
//
// Solidity: function declareOperatorFee(uint64 operatorId, uint256 fee) returns()
func (_SSVNetwork *SSVNetworkSession) DeclareOperatorFee(operatorId uint64, fee *big.Int) (*types.Transaction, error) {
	return _SSVNetwork.Contract.DeclareOperatorFee(&_SSVNetwork.TransactOpts, operatorId, fee)
}

// DeclareOperatorFee is a paid mutator transaction binding the contract method 0xb317c35f.
//
// Solidity: function declareOperatorFee(uint64 operatorId, uint256 fee) returns()
func (_SSVNetwork *SSVNetworkTransactorSession) DeclareOperatorFee(operatorId uint64, fee *big.Int) (*types.Transaction, error) {
	return _SSVNetwork.Contract.DeclareOperatorFee(&_SSVNetwork.TransactOpts, operatorId, fee)
}

// Deposit is a paid mutator transaction binding the contract method 0xbc26e7e5.
//
// Solidity: function deposit(address owner, uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SSVNetwork *SSVNetworkTransactor) Deposit(opts *bind.TransactOpts, owner common.Address, operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "deposit", owner, operatorIds, amount, cluster)
}

// Deposit is a paid mutator transaction binding the contract method 0xbc26e7e5.
//
// Solidity: function deposit(address owner, uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SSVNetwork *SSVNetworkSession) Deposit(owner common.Address, operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SSVNetwork.Contract.Deposit(&_SSVNetwork.TransactOpts, owner, operatorIds, amount, cluster)
}

// Deposit is a paid mutator transaction binding the contract method 0xbc26e7e5.
//
// Solidity: function deposit(address owner, uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SSVNetwork *SSVNetworkTransactorSession) Deposit(owner common.Address, operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SSVNetwork.Contract.Deposit(&_SSVNetwork.TransactOpts, owner, operatorIds, amount, cluster)
}

// ExecuteOperatorFee is a paid mutator transaction binding the contract method 0x8932cee0.
//
// Solidity: function executeOperatorFee(uint64 operatorId) returns()
func (_SSVNetwork *SSVNetworkTransactor) ExecuteOperatorFee(opts *bind.TransactOpts, operatorId uint64) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "executeOperatorFee", operatorId)
}

// ExecuteOperatorFee is a paid mutator transaction binding the contract method 0x8932cee0.
//
// Solidity: function executeOperatorFee(uint64 operatorId) returns()
func (_SSVNetwork *SSVNetworkSession) ExecuteOperatorFee(operatorId uint64) (*types.Transaction, error) {
	return _SSVNetwork.Contract.ExecuteOperatorFee(&_SSVNetwork.TransactOpts, operatorId)
}

// ExecuteOperatorFee is a paid mutator transaction binding the contract method 0x8932cee0.
//
// Solidity: function executeOperatorFee(uint64 operatorId) returns()
func (_SSVNetwork *SSVNetworkTransactorSession) ExecuteOperatorFee(operatorId uint64) (*types.Transaction, error) {
	return _SSVNetwork.Contract.ExecuteOperatorFee(&_SSVNetwork.TransactOpts, operatorId)
}

// Initialize is a paid mutator transaction binding the contract method 0xc626c3c6.
//
// Solidity: function initialize(address token_, address ssvOperators_, address ssvClusters_, address ssvDAO_, address ssvViews_, uint64 minimumBlocksBeforeLiquidation_, uint256 minimumLiquidationCollateral_, uint32 validatorsPerOperatorLimit_, uint64 declareOperatorFeePeriod_, uint64 executeOperatorFeePeriod_, uint64 operatorMaxFeeIncrease_) returns()
func (_SSVNetwork *SSVNetworkTransactor) Initialize(opts *bind.TransactOpts, token_ common.Address, ssvOperators_ common.Address, ssvClusters_ common.Address, ssvDAO_ common.Address, ssvViews_ common.Address, minimumBlocksBeforeLiquidation_ uint64, minimumLiquidationCollateral_ *big.Int, validatorsPerOperatorLimit_ uint32, declareOperatorFeePeriod_ uint64, executeOperatorFeePeriod_ uint64, operatorMaxFeeIncrease_ uint64) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "initialize", token_, ssvOperators_, ssvClusters_, ssvDAO_, ssvViews_, minimumBlocksBeforeLiquidation_, minimumLiquidationCollateral_, validatorsPerOperatorLimit_, declareOperatorFeePeriod_, executeOperatorFeePeriod_, operatorMaxFeeIncrease_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc626c3c6.
//
// Solidity: function initialize(address token_, address ssvOperators_, address ssvClusters_, address ssvDAO_, address ssvViews_, uint64 minimumBlocksBeforeLiquidation_, uint256 minimumLiquidationCollateral_, uint32 validatorsPerOperatorLimit_, uint64 declareOperatorFeePeriod_, uint64 executeOperatorFeePeriod_, uint64 operatorMaxFeeIncrease_) returns()
func (_SSVNetwork *SSVNetworkSession) Initialize(token_ common.Address, ssvOperators_ common.Address, ssvClusters_ common.Address, ssvDAO_ common.Address, ssvViews_ common.Address, minimumBlocksBeforeLiquidation_ uint64, minimumLiquidationCollateral_ *big.Int, validatorsPerOperatorLimit_ uint32, declareOperatorFeePeriod_ uint64, executeOperatorFeePeriod_ uint64, operatorMaxFeeIncrease_ uint64) (*types.Transaction, error) {
	return _SSVNetwork.Contract.Initialize(&_SSVNetwork.TransactOpts, token_, ssvOperators_, ssvClusters_, ssvDAO_, ssvViews_, minimumBlocksBeforeLiquidation_, minimumLiquidationCollateral_, validatorsPerOperatorLimit_, declareOperatorFeePeriod_, executeOperatorFeePeriod_, operatorMaxFeeIncrease_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc626c3c6.
//
// Solidity: function initialize(address token_, address ssvOperators_, address ssvClusters_, address ssvDAO_, address ssvViews_, uint64 minimumBlocksBeforeLiquidation_, uint256 minimumLiquidationCollateral_, uint32 validatorsPerOperatorLimit_, uint64 declareOperatorFeePeriod_, uint64 executeOperatorFeePeriod_, uint64 operatorMaxFeeIncrease_) returns()
func (_SSVNetwork *SSVNetworkTransactorSession) Initialize(token_ common.Address, ssvOperators_ common.Address, ssvClusters_ common.Address, ssvDAO_ common.Address, ssvViews_ common.Address, minimumBlocksBeforeLiquidation_ uint64, minimumLiquidationCollateral_ *big.Int, validatorsPerOperatorLimit_ uint32, declareOperatorFeePeriod_ uint64, executeOperatorFeePeriod_ uint64, operatorMaxFeeIncrease_ uint64) (*types.Transaction, error) {
	return _SSVNetwork.Contract.Initialize(&_SSVNetwork.TransactOpts, token_, ssvOperators_, ssvClusters_, ssvDAO_, ssvViews_, minimumBlocksBeforeLiquidation_, minimumLiquidationCollateral_, validatorsPerOperatorLimit_, declareOperatorFeePeriod_, executeOperatorFeePeriod_, operatorMaxFeeIncrease_)
}

// Liquidate is a paid mutator transaction binding the contract method 0xbf0f2fb2.
//
// Solidity: function liquidate(address owner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SSVNetwork *SSVNetworkTransactor) Liquidate(opts *bind.TransactOpts, owner common.Address, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "liquidate", owner, operatorIds, cluster)
}

// Liquidate is a paid mutator transaction binding the contract method 0xbf0f2fb2.
//
// Solidity: function liquidate(address owner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SSVNetwork *SSVNetworkSession) Liquidate(owner common.Address, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SSVNetwork.Contract.Liquidate(&_SSVNetwork.TransactOpts, owner, operatorIds, cluster)
}

// Liquidate is a paid mutator transaction binding the contract method 0xbf0f2fb2.
//
// Solidity: function liquidate(address owner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SSVNetwork *SSVNetworkTransactorSession) Liquidate(owner common.Address, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SSVNetwork.Contract.Liquidate(&_SSVNetwork.TransactOpts, owner, operatorIds, cluster)
}

// Reactivate is a paid mutator transaction binding the contract method 0x5fec6dd0.
//
// Solidity: function reactivate(uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SSVNetwork *SSVNetworkTransactor) Reactivate(opts *bind.TransactOpts, operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "reactivate", operatorIds, amount, cluster)
}

// Reactivate is a paid mutator transaction binding the contract method 0x5fec6dd0.
//
// Solidity: function reactivate(uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SSVNetwork *SSVNetworkSession) Reactivate(operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SSVNetwork.Contract.Reactivate(&_SSVNetwork.TransactOpts, operatorIds, amount, cluster)
}

// Reactivate is a paid mutator transaction binding the contract method 0x5fec6dd0.
//
// Solidity: function reactivate(uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SSVNetwork *SSVNetworkTransactorSession) Reactivate(operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SSVNetwork.Contract.Reactivate(&_SSVNetwork.TransactOpts, operatorIds, amount, cluster)
}

// ReduceOperatorFee is a paid mutator transaction binding the contract method 0x190d82e4.
//
// Solidity: function reduceOperatorFee(uint64 operatorId, uint256 fee) returns()
func (_SSVNetwork *SSVNetworkTransactor) ReduceOperatorFee(opts *bind.TransactOpts, operatorId uint64, fee *big.Int) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "reduceOperatorFee", operatorId, fee)
}

// ReduceOperatorFee is a paid mutator transaction binding the contract method 0x190d82e4.
//
// Solidity: function reduceOperatorFee(uint64 operatorId, uint256 fee) returns()
func (_SSVNetwork *SSVNetworkSession) ReduceOperatorFee(operatorId uint64, fee *big.Int) (*types.Transaction, error) {
	return _SSVNetwork.Contract.ReduceOperatorFee(&_SSVNetwork.TransactOpts, operatorId, fee)
}

// ReduceOperatorFee is a paid mutator transaction binding the contract method 0x190d82e4.
//
// Solidity: function reduceOperatorFee(uint64 operatorId, uint256 fee) returns()
func (_SSVNetwork *SSVNetworkTransactorSession) ReduceOperatorFee(operatorId uint64, fee *big.Int) (*types.Transaction, error) {
	return _SSVNetwork.Contract.ReduceOperatorFee(&_SSVNetwork.TransactOpts, operatorId, fee)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xff212c5c.
//
// Solidity: function registerOperator(bytes publicKey, uint256 fee) returns(uint64 id)
func (_SSVNetwork *SSVNetworkTransactor) RegisterOperator(opts *bind.TransactOpts, publicKey []byte, fee *big.Int) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "registerOperator", publicKey, fee)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xff212c5c.
//
// Solidity: function registerOperator(bytes publicKey, uint256 fee) returns(uint64 id)
func (_SSVNetwork *SSVNetworkSession) RegisterOperator(publicKey []byte, fee *big.Int) (*types.Transaction, error) {
	return _SSVNetwork.Contract.RegisterOperator(&_SSVNetwork.TransactOpts, publicKey, fee)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xff212c5c.
//
// Solidity: function registerOperator(bytes publicKey, uint256 fee) returns(uint64 id)
func (_SSVNetwork *SSVNetworkTransactorSession) RegisterOperator(publicKey []byte, fee *big.Int) (*types.Transaction, error) {
	return _SSVNetwork.Contract.RegisterOperator(&_SSVNetwork.TransactOpts, publicKey, fee)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x06e8fb9c.
//
// Solidity: function registerValidator(bytes publicKey, uint64[] operatorIds, bytes sharesData, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SSVNetwork *SSVNetworkTransactor) RegisterValidator(opts *bind.TransactOpts, publicKey []byte, operatorIds []uint64, sharesData []byte, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "registerValidator", publicKey, operatorIds, sharesData, amount, cluster)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x06e8fb9c.
//
// Solidity: function registerValidator(bytes publicKey, uint64[] operatorIds, bytes sharesData, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SSVNetwork *SSVNetworkSession) RegisterValidator(publicKey []byte, operatorIds []uint64, sharesData []byte, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SSVNetwork.Contract.RegisterValidator(&_SSVNetwork.TransactOpts, publicKey, operatorIds, sharesData, amount, cluster)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x06e8fb9c.
//
// Solidity: function registerValidator(bytes publicKey, uint64[] operatorIds, bytes sharesData, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SSVNetwork *SSVNetworkTransactorSession) RegisterValidator(publicKey []byte, operatorIds []uint64, sharesData []byte, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SSVNetwork.Contract.RegisterValidator(&_SSVNetwork.TransactOpts, publicKey, operatorIds, sharesData, amount, cluster)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x2e168e0e.
//
// Solidity: function removeOperator(uint64 operatorId) returns()
func (_SSVNetwork *SSVNetworkTransactor) RemoveOperator(opts *bind.TransactOpts, operatorId uint64) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "removeOperator", operatorId)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x2e168e0e.
//
// Solidity: function removeOperator(uint64 operatorId) returns()
func (_SSVNetwork *SSVNetworkSession) RemoveOperator(operatorId uint64) (*types.Transaction, error) {
	return _SSVNetwork.Contract.RemoveOperator(&_SSVNetwork.TransactOpts, operatorId)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x2e168e0e.
//
// Solidity: function removeOperator(uint64 operatorId) returns()
func (_SSVNetwork *SSVNetworkTransactorSession) RemoveOperator(operatorId uint64) (*types.Transaction, error) {
	return _SSVNetwork.Contract.RemoveOperator(&_SSVNetwork.TransactOpts, operatorId)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x12b3fc19.
//
// Solidity: function removeValidator(bytes publicKey, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SSVNetwork *SSVNetworkTransactor) RemoveValidator(opts *bind.TransactOpts, publicKey []byte, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "removeValidator", publicKey, operatorIds, cluster)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x12b3fc19.
//
// Solidity: function removeValidator(bytes publicKey, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SSVNetwork *SSVNetworkSession) RemoveValidator(publicKey []byte, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SSVNetwork.Contract.RemoveValidator(&_SSVNetwork.TransactOpts, publicKey, operatorIds, cluster)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x12b3fc19.
//
// Solidity: function removeValidator(bytes publicKey, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SSVNetwork *SSVNetworkTransactorSession) RemoveValidator(publicKey []byte, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SSVNetwork.Contract.RemoveValidator(&_SSVNetwork.TransactOpts, publicKey, operatorIds, cluster)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SSVNetwork *SSVNetworkTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SSVNetwork *SSVNetworkSession) RenounceOwnership() (*types.Transaction, error) {
	return _SSVNetwork.Contract.RenounceOwnership(&_SSVNetwork.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SSVNetwork *SSVNetworkTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SSVNetwork.Contract.RenounceOwnership(&_SSVNetwork.TransactOpts)
}

// SetFeeRecipientAddress is a paid mutator transaction binding the contract method 0xdbcdc2cc.
//
// Solidity: function setFeeRecipientAddress(address recipientAddress) returns()
func (_SSVNetwork *SSVNetworkTransactor) SetFeeRecipientAddress(opts *bind.TransactOpts, recipientAddress common.Address) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "setFeeRecipientAddress", recipientAddress)
}

// SetFeeRecipientAddress is a paid mutator transaction binding the contract method 0xdbcdc2cc.
//
// Solidity: function setFeeRecipientAddress(address recipientAddress) returns()
func (_SSVNetwork *SSVNetworkSession) SetFeeRecipientAddress(recipientAddress common.Address) (*types.Transaction, error) {
	return _SSVNetwork.Contract.SetFeeRecipientAddress(&_SSVNetwork.TransactOpts, recipientAddress)
}

// SetFeeRecipientAddress is a paid mutator transaction binding the contract method 0xdbcdc2cc.
//
// Solidity: function setFeeRecipientAddress(address recipientAddress) returns()
func (_SSVNetwork *SSVNetworkTransactorSession) SetFeeRecipientAddress(recipientAddress common.Address) (*types.Transaction, error) {
	return _SSVNetwork.Contract.SetFeeRecipientAddress(&_SSVNetwork.TransactOpts, recipientAddress)
}

// SetOperatorWhitelist is a paid mutator transaction binding the contract method 0xc90a7eab.
//
// Solidity: function setOperatorWhitelist(uint64 operatorId, address whitelisted) returns()
func (_SSVNetwork *SSVNetworkTransactor) SetOperatorWhitelist(opts *bind.TransactOpts, operatorId uint64, whitelisted common.Address) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "setOperatorWhitelist", operatorId, whitelisted)
}

// SetOperatorWhitelist is a paid mutator transaction binding the contract method 0xc90a7eab.
//
// Solidity: function setOperatorWhitelist(uint64 operatorId, address whitelisted) returns()
func (_SSVNetwork *SSVNetworkSession) SetOperatorWhitelist(operatorId uint64, whitelisted common.Address) (*types.Transaction, error) {
	return _SSVNetwork.Contract.SetOperatorWhitelist(&_SSVNetwork.TransactOpts, operatorId, whitelisted)
}

// SetOperatorWhitelist is a paid mutator transaction binding the contract method 0xc90a7eab.
//
// Solidity: function setOperatorWhitelist(uint64 operatorId, address whitelisted) returns()
func (_SSVNetwork *SSVNetworkTransactorSession) SetOperatorWhitelist(operatorId uint64, whitelisted common.Address) (*types.Transaction, error) {
	return _SSVNetwork.Contract.SetOperatorWhitelist(&_SSVNetwork.TransactOpts, operatorId, whitelisted)
}

// SetRegisterAuth is a paid mutator transaction binding the contract method 0x066a4aa2.
//
// Solidity: function setRegisterAuth(address userAddress, (bool,bool) auth) returns()
func (_SSVNetwork *SSVNetworkTransactor) SetRegisterAuth(opts *bind.TransactOpts, userAddress common.Address, auth Authorization) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "setRegisterAuth", userAddress, auth)
}

// SetRegisterAuth is a paid mutator transaction binding the contract method 0x066a4aa2.
//
// Solidity: function setRegisterAuth(address userAddress, (bool,bool) auth) returns()
func (_SSVNetwork *SSVNetworkSession) SetRegisterAuth(userAddress common.Address, auth Authorization) (*types.Transaction, error) {
	return _SSVNetwork.Contract.SetRegisterAuth(&_SSVNetwork.TransactOpts, userAddress, auth)
}

// SetRegisterAuth is a paid mutator transaction binding the contract method 0x066a4aa2.
//
// Solidity: function setRegisterAuth(address userAddress, (bool,bool) auth) returns()
func (_SSVNetwork *SSVNetworkTransactorSession) SetRegisterAuth(userAddress common.Address, auth Authorization) (*types.Transaction, error) {
	return _SSVNetwork.Contract.SetRegisterAuth(&_SSVNetwork.TransactOpts, userAddress, auth)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SSVNetwork *SSVNetworkTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SSVNetwork *SSVNetworkSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SSVNetwork.Contract.TransferOwnership(&_SSVNetwork.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SSVNetwork *SSVNetworkTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SSVNetwork.Contract.TransferOwnership(&_SSVNetwork.TransactOpts, newOwner)
}

// UpdateDeclareOperatorFeePeriod is a paid mutator transaction binding the contract method 0x79e3e4e4.
//
// Solidity: function updateDeclareOperatorFeePeriod(uint64 timeInSeconds) returns()
func (_SSVNetwork *SSVNetworkTransactor) UpdateDeclareOperatorFeePeriod(opts *bind.TransactOpts, timeInSeconds uint64) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "updateDeclareOperatorFeePeriod", timeInSeconds)
}

// UpdateDeclareOperatorFeePeriod is a paid mutator transaction binding the contract method 0x79e3e4e4.
//
// Solidity: function updateDeclareOperatorFeePeriod(uint64 timeInSeconds) returns()
func (_SSVNetwork *SSVNetworkSession) UpdateDeclareOperatorFeePeriod(timeInSeconds uint64) (*types.Transaction, error) {
	return _SSVNetwork.Contract.UpdateDeclareOperatorFeePeriod(&_SSVNetwork.TransactOpts, timeInSeconds)
}

// UpdateDeclareOperatorFeePeriod is a paid mutator transaction binding the contract method 0x79e3e4e4.
//
// Solidity: function updateDeclareOperatorFeePeriod(uint64 timeInSeconds) returns()
func (_SSVNetwork *SSVNetworkTransactorSession) UpdateDeclareOperatorFeePeriod(timeInSeconds uint64) (*types.Transaction, error) {
	return _SSVNetwork.Contract.UpdateDeclareOperatorFeePeriod(&_SSVNetwork.TransactOpts, timeInSeconds)
}

// UpdateExecuteOperatorFeePeriod is a paid mutator transaction binding the contract method 0xeb608022.
//
// Solidity: function updateExecuteOperatorFeePeriod(uint64 timeInSeconds) returns()
func (_SSVNetwork *SSVNetworkTransactor) UpdateExecuteOperatorFeePeriod(opts *bind.TransactOpts, timeInSeconds uint64) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "updateExecuteOperatorFeePeriod", timeInSeconds)
}

// UpdateExecuteOperatorFeePeriod is a paid mutator transaction binding the contract method 0xeb608022.
//
// Solidity: function updateExecuteOperatorFeePeriod(uint64 timeInSeconds) returns()
func (_SSVNetwork *SSVNetworkSession) UpdateExecuteOperatorFeePeriod(timeInSeconds uint64) (*types.Transaction, error) {
	return _SSVNetwork.Contract.UpdateExecuteOperatorFeePeriod(&_SSVNetwork.TransactOpts, timeInSeconds)
}

// UpdateExecuteOperatorFeePeriod is a paid mutator transaction binding the contract method 0xeb608022.
//
// Solidity: function updateExecuteOperatorFeePeriod(uint64 timeInSeconds) returns()
func (_SSVNetwork *SSVNetworkTransactorSession) UpdateExecuteOperatorFeePeriod(timeInSeconds uint64) (*types.Transaction, error) {
	return _SSVNetwork.Contract.UpdateExecuteOperatorFeePeriod(&_SSVNetwork.TransactOpts, timeInSeconds)
}

// UpdateLiquidationThresholdPeriod is a paid mutator transaction binding the contract method 0x6512447d.
//
// Solidity: function updateLiquidationThresholdPeriod(uint64 blocks) returns()
func (_SSVNetwork *SSVNetworkTransactor) UpdateLiquidationThresholdPeriod(opts *bind.TransactOpts, blocks uint64) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "updateLiquidationThresholdPeriod", blocks)
}

// UpdateLiquidationThresholdPeriod is a paid mutator transaction binding the contract method 0x6512447d.
//
// Solidity: function updateLiquidationThresholdPeriod(uint64 blocks) returns()
func (_SSVNetwork *SSVNetworkSession) UpdateLiquidationThresholdPeriod(blocks uint64) (*types.Transaction, error) {
	return _SSVNetwork.Contract.UpdateLiquidationThresholdPeriod(&_SSVNetwork.TransactOpts, blocks)
}

// UpdateLiquidationThresholdPeriod is a paid mutator transaction binding the contract method 0x6512447d.
//
// Solidity: function updateLiquidationThresholdPeriod(uint64 blocks) returns()
func (_SSVNetwork *SSVNetworkTransactorSession) UpdateLiquidationThresholdPeriod(blocks uint64) (*types.Transaction, error) {
	return _SSVNetwork.Contract.UpdateLiquidationThresholdPeriod(&_SSVNetwork.TransactOpts, blocks)
}

// UpdateMinimumLiquidationCollateral is a paid mutator transaction binding the contract method 0xb4c9c408.
//
// Solidity: function updateMinimumLiquidationCollateral(uint256 amount) returns()
func (_SSVNetwork *SSVNetworkTransactor) UpdateMinimumLiquidationCollateral(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "updateMinimumLiquidationCollateral", amount)
}

// UpdateMinimumLiquidationCollateral is a paid mutator transaction binding the contract method 0xb4c9c408.
//
// Solidity: function updateMinimumLiquidationCollateral(uint256 amount) returns()
func (_SSVNetwork *SSVNetworkSession) UpdateMinimumLiquidationCollateral(amount *big.Int) (*types.Transaction, error) {
	return _SSVNetwork.Contract.UpdateMinimumLiquidationCollateral(&_SSVNetwork.TransactOpts, amount)
}

// UpdateMinimumLiquidationCollateral is a paid mutator transaction binding the contract method 0xb4c9c408.
//
// Solidity: function updateMinimumLiquidationCollateral(uint256 amount) returns()
func (_SSVNetwork *SSVNetworkTransactorSession) UpdateMinimumLiquidationCollateral(amount *big.Int) (*types.Transaction, error) {
	return _SSVNetwork.Contract.UpdateMinimumLiquidationCollateral(&_SSVNetwork.TransactOpts, amount)
}

// UpdateNetworkFee is a paid mutator transaction binding the contract method 0x1f1f9fd5.
//
// Solidity: function updateNetworkFee(uint256 fee) returns()
func (_SSVNetwork *SSVNetworkTransactor) UpdateNetworkFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "updateNetworkFee", fee)
}

// UpdateNetworkFee is a paid mutator transaction binding the contract method 0x1f1f9fd5.
//
// Solidity: function updateNetworkFee(uint256 fee) returns()
func (_SSVNetwork *SSVNetworkSession) UpdateNetworkFee(fee *big.Int) (*types.Transaction, error) {
	return _SSVNetwork.Contract.UpdateNetworkFee(&_SSVNetwork.TransactOpts, fee)
}

// UpdateNetworkFee is a paid mutator transaction binding the contract method 0x1f1f9fd5.
//
// Solidity: function updateNetworkFee(uint256 fee) returns()
func (_SSVNetwork *SSVNetworkTransactorSession) UpdateNetworkFee(fee *big.Int) (*types.Transaction, error) {
	return _SSVNetwork.Contract.UpdateNetworkFee(&_SSVNetwork.TransactOpts, fee)
}

// UpdateOperatorFeeIncreaseLimit is a paid mutator transaction binding the contract method 0x3631983f.
//
// Solidity: function updateOperatorFeeIncreaseLimit(uint64 percentage) returns()
func (_SSVNetwork *SSVNetworkTransactor) UpdateOperatorFeeIncreaseLimit(opts *bind.TransactOpts, percentage uint64) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "updateOperatorFeeIncreaseLimit", percentage)
}

// UpdateOperatorFeeIncreaseLimit is a paid mutator transaction binding the contract method 0x3631983f.
//
// Solidity: function updateOperatorFeeIncreaseLimit(uint64 percentage) returns()
func (_SSVNetwork *SSVNetworkSession) UpdateOperatorFeeIncreaseLimit(percentage uint64) (*types.Transaction, error) {
	return _SSVNetwork.Contract.UpdateOperatorFeeIncreaseLimit(&_SSVNetwork.TransactOpts, percentage)
}

// UpdateOperatorFeeIncreaseLimit is a paid mutator transaction binding the contract method 0x3631983f.
//
// Solidity: function updateOperatorFeeIncreaseLimit(uint64 percentage) returns()
func (_SSVNetwork *SSVNetworkTransactorSession) UpdateOperatorFeeIncreaseLimit(percentage uint64) (*types.Transaction, error) {
	return _SSVNetwork.Contract.UpdateOperatorFeeIncreaseLimit(&_SSVNetwork.TransactOpts, percentage)
}

// UpgradeModule is a paid mutator transaction binding the contract method 0x541fcf60.
//
// Solidity: function upgradeModule(uint8 moduleId, address moduleAddress) returns()
func (_SSVNetwork *SSVNetworkTransactor) UpgradeModule(opts *bind.TransactOpts, moduleId uint8, moduleAddress common.Address) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "upgradeModule", moduleId, moduleAddress)
}

// UpgradeModule is a paid mutator transaction binding the contract method 0x541fcf60.
//
// Solidity: function upgradeModule(uint8 moduleId, address moduleAddress) returns()
func (_SSVNetwork *SSVNetworkSession) UpgradeModule(moduleId uint8, moduleAddress common.Address) (*types.Transaction, error) {
	return _SSVNetwork.Contract.UpgradeModule(&_SSVNetwork.TransactOpts, moduleId, moduleAddress)
}

// UpgradeModule is a paid mutator transaction binding the contract method 0x541fcf60.
//
// Solidity: function upgradeModule(uint8 moduleId, address moduleAddress) returns()
func (_SSVNetwork *SSVNetworkTransactorSession) UpgradeModule(moduleId uint8, moduleAddress common.Address) (*types.Transaction, error) {
	return _SSVNetwork.Contract.UpgradeModule(&_SSVNetwork.TransactOpts, moduleId, moduleAddress)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SSVNetwork *SSVNetworkTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SSVNetwork *SSVNetworkSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _SSVNetwork.Contract.UpgradeTo(&_SSVNetwork.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SSVNetwork *SSVNetworkTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _SSVNetwork.Contract.UpgradeTo(&_SSVNetwork.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SSVNetwork *SSVNetworkTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SSVNetwork *SSVNetworkSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SSVNetwork.Contract.UpgradeToAndCall(&_SSVNetwork.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SSVNetwork *SSVNetworkTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SSVNetwork.Contract.UpgradeToAndCall(&_SSVNetwork.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x686e682c.
//
// Solidity: function withdraw(uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SSVNetwork *SSVNetworkTransactor) Withdraw(opts *bind.TransactOpts, operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "withdraw", operatorIds, amount, cluster)
}

// Withdraw is a paid mutator transaction binding the contract method 0x686e682c.
//
// Solidity: function withdraw(uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SSVNetwork *SSVNetworkSession) Withdraw(operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SSVNetwork.Contract.Withdraw(&_SSVNetwork.TransactOpts, operatorIds, amount, cluster)
}

// Withdraw is a paid mutator transaction binding the contract method 0x686e682c.
//
// Solidity: function withdraw(uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_SSVNetwork *SSVNetworkTransactorSession) Withdraw(operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _SSVNetwork.Contract.Withdraw(&_SSVNetwork.TransactOpts, operatorIds, amount, cluster)
}

// WithdrawNetworkEarnings is a paid mutator transaction binding the contract method 0xd2231741.
//
// Solidity: function withdrawNetworkEarnings(uint256 amount) returns()
func (_SSVNetwork *SSVNetworkTransactor) WithdrawNetworkEarnings(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "withdrawNetworkEarnings", amount)
}

// WithdrawNetworkEarnings is a paid mutator transaction binding the contract method 0xd2231741.
//
// Solidity: function withdrawNetworkEarnings(uint256 amount) returns()
func (_SSVNetwork *SSVNetworkSession) WithdrawNetworkEarnings(amount *big.Int) (*types.Transaction, error) {
	return _SSVNetwork.Contract.WithdrawNetworkEarnings(&_SSVNetwork.TransactOpts, amount)
}

// WithdrawNetworkEarnings is a paid mutator transaction binding the contract method 0xd2231741.
//
// Solidity: function withdrawNetworkEarnings(uint256 amount) returns()
func (_SSVNetwork *SSVNetworkTransactorSession) WithdrawNetworkEarnings(amount *big.Int) (*types.Transaction, error) {
	return _SSVNetwork.Contract.WithdrawNetworkEarnings(&_SSVNetwork.TransactOpts, amount)
}

// WithdrawOperatorEarnings is a paid mutator transaction binding the contract method 0x35f63767.
//
// Solidity: function withdrawOperatorEarnings(uint64 operatorId, uint256 amount) returns()
func (_SSVNetwork *SSVNetworkTransactor) WithdrawOperatorEarnings(opts *bind.TransactOpts, operatorId uint64, amount *big.Int) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "withdrawOperatorEarnings", operatorId, amount)
}

// WithdrawOperatorEarnings is a paid mutator transaction binding the contract method 0x35f63767.
//
// Solidity: function withdrawOperatorEarnings(uint64 operatorId, uint256 amount) returns()
func (_SSVNetwork *SSVNetworkSession) WithdrawOperatorEarnings(operatorId uint64, amount *big.Int) (*types.Transaction, error) {
	return _SSVNetwork.Contract.WithdrawOperatorEarnings(&_SSVNetwork.TransactOpts, operatorId, amount)
}

// WithdrawOperatorEarnings is a paid mutator transaction binding the contract method 0x35f63767.
//
// Solidity: function withdrawOperatorEarnings(uint64 operatorId, uint256 amount) returns()
func (_SSVNetwork *SSVNetworkTransactorSession) WithdrawOperatorEarnings(operatorId uint64, amount *big.Int) (*types.Transaction, error) {
	return _SSVNetwork.Contract.WithdrawOperatorEarnings(&_SSVNetwork.TransactOpts, operatorId, amount)
}

// WithdrawOperatorEarnings0 is a paid mutator transaction binding the contract method 0xc8c876c6.
//
// Solidity: function withdrawOperatorEarnings(uint64 operatorId) returns()
func (_SSVNetwork *SSVNetworkTransactor) WithdrawOperatorEarnings0(opts *bind.TransactOpts, operatorId uint64) (*types.Transaction, error) {
	return _SSVNetwork.contract.Transact(opts, "withdrawOperatorEarnings0", operatorId)
}

// WithdrawOperatorEarnings0 is a paid mutator transaction binding the contract method 0xc8c876c6.
//
// Solidity: function withdrawOperatorEarnings(uint64 operatorId) returns()
func (_SSVNetwork *SSVNetworkSession) WithdrawOperatorEarnings0(operatorId uint64) (*types.Transaction, error) {
	return _SSVNetwork.Contract.WithdrawOperatorEarnings0(&_SSVNetwork.TransactOpts, operatorId)
}

// WithdrawOperatorEarnings0 is a paid mutator transaction binding the contract method 0xc8c876c6.
//
// Solidity: function withdrawOperatorEarnings(uint64 operatorId) returns()
func (_SSVNetwork *SSVNetworkTransactorSession) WithdrawOperatorEarnings0(operatorId uint64) (*types.Transaction, error) {
	return _SSVNetwork.Contract.WithdrawOperatorEarnings0(&_SSVNetwork.TransactOpts, operatorId)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_SSVNetwork *SSVNetworkTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _SSVNetwork.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_SSVNetwork *SSVNetworkSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SSVNetwork.Contract.Fallback(&_SSVNetwork.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_SSVNetwork *SSVNetworkTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SSVNetwork.Contract.Fallback(&_SSVNetwork.TransactOpts, calldata)
}

// SSVNetworkAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the SSVNetwork contract.
type SSVNetworkAdminChangedIterator struct {
	Event *SSVNetworkAdminChanged // Event containing the contract specifics and raw log

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
func (it *SSVNetworkAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SSVNetworkAdminChanged)
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
		it.Event = new(SSVNetworkAdminChanged)
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
func (it *SSVNetworkAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SSVNetworkAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SSVNetworkAdminChanged represents a AdminChanged event raised by the SSVNetwork contract.
type SSVNetworkAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_SSVNetwork *SSVNetworkFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*SSVNetworkAdminChangedIterator, error) {

	logs, sub, err := _SSVNetwork.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &SSVNetworkAdminChangedIterator{contract: _SSVNetwork.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_SSVNetwork *SSVNetworkFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *SSVNetworkAdminChanged) (event.Subscription, error) {

	logs, sub, err := _SSVNetwork.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SSVNetworkAdminChanged)
				if err := _SSVNetwork.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_SSVNetwork *SSVNetworkFilterer) ParseAdminChanged(log types.Log) (*SSVNetworkAdminChanged, error) {
	event := new(SSVNetworkAdminChanged)
	if err := _SSVNetwork.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SSVNetworkBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the SSVNetwork contract.
type SSVNetworkBeaconUpgradedIterator struct {
	Event *SSVNetworkBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *SSVNetworkBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SSVNetworkBeaconUpgraded)
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
		it.Event = new(SSVNetworkBeaconUpgraded)
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
func (it *SSVNetworkBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SSVNetworkBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SSVNetworkBeaconUpgraded represents a BeaconUpgraded event raised by the SSVNetwork contract.
type SSVNetworkBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_SSVNetwork *SSVNetworkFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*SSVNetworkBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _SSVNetwork.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &SSVNetworkBeaconUpgradedIterator{contract: _SSVNetwork.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_SSVNetwork *SSVNetworkFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *SSVNetworkBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _SSVNetwork.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SSVNetworkBeaconUpgraded)
				if err := _SSVNetwork.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_SSVNetwork *SSVNetworkFilterer) ParseBeaconUpgraded(log types.Log) (*SSVNetworkBeaconUpgraded, error) {
	event := new(SSVNetworkBeaconUpgraded)
	if err := _SSVNetwork.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SSVNetworkClusterDepositedIterator is returned from FilterClusterDeposited and is used to iterate over the raw logs and unpacked data for ClusterDeposited events raised by the SSVNetwork contract.
type SSVNetworkClusterDepositedIterator struct {
	Event *SSVNetworkClusterDeposited // Event containing the contract specifics and raw log

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
func (it *SSVNetworkClusterDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SSVNetworkClusterDeposited)
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
		it.Event = new(SSVNetworkClusterDeposited)
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
func (it *SSVNetworkClusterDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SSVNetworkClusterDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SSVNetworkClusterDeposited represents a ClusterDeposited event raised by the SSVNetwork contract.
type SSVNetworkClusterDeposited struct {
	Owner       common.Address
	OperatorIds []uint64
	Value       *big.Int
	Cluster     ISSVNetworkCoreCluster
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClusterDeposited is a free log retrieval operation binding the contract event 0x2bac1912f2481d12f0df08647c06bee174967c62d3a03cbc078eb215dc1bd9a2.
//
// Solidity: event ClusterDeposited(address indexed owner, uint64[] operatorIds, uint256 value, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SSVNetwork *SSVNetworkFilterer) FilterClusterDeposited(opts *bind.FilterOpts, owner []common.Address) (*SSVNetworkClusterDepositedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SSVNetwork.contract.FilterLogs(opts, "ClusterDeposited", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SSVNetworkClusterDepositedIterator{contract: _SSVNetwork.contract, event: "ClusterDeposited", logs: logs, sub: sub}, nil
}

// WatchClusterDeposited is a free log subscription operation binding the contract event 0x2bac1912f2481d12f0df08647c06bee174967c62d3a03cbc078eb215dc1bd9a2.
//
// Solidity: event ClusterDeposited(address indexed owner, uint64[] operatorIds, uint256 value, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SSVNetwork *SSVNetworkFilterer) WatchClusterDeposited(opts *bind.WatchOpts, sink chan<- *SSVNetworkClusterDeposited, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SSVNetwork.contract.WatchLogs(opts, "ClusterDeposited", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SSVNetworkClusterDeposited)
				if err := _SSVNetwork.contract.UnpackLog(event, "ClusterDeposited", log); err != nil {
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

// ParseClusterDeposited is a log parse operation binding the contract event 0x2bac1912f2481d12f0df08647c06bee174967c62d3a03cbc078eb215dc1bd9a2.
//
// Solidity: event ClusterDeposited(address indexed owner, uint64[] operatorIds, uint256 value, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SSVNetwork *SSVNetworkFilterer) ParseClusterDeposited(log types.Log) (*SSVNetworkClusterDeposited, error) {
	event := new(SSVNetworkClusterDeposited)
	if err := _SSVNetwork.contract.UnpackLog(event, "ClusterDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SSVNetworkClusterLiquidatedIterator is returned from FilterClusterLiquidated and is used to iterate over the raw logs and unpacked data for ClusterLiquidated events raised by the SSVNetwork contract.
type SSVNetworkClusterLiquidatedIterator struct {
	Event *SSVNetworkClusterLiquidated // Event containing the contract specifics and raw log

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
func (it *SSVNetworkClusterLiquidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SSVNetworkClusterLiquidated)
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
		it.Event = new(SSVNetworkClusterLiquidated)
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
func (it *SSVNetworkClusterLiquidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SSVNetworkClusterLiquidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SSVNetworkClusterLiquidated represents a ClusterLiquidated event raised by the SSVNetwork contract.
type SSVNetworkClusterLiquidated struct {
	Owner       common.Address
	OperatorIds []uint64
	Cluster     ISSVNetworkCoreCluster
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClusterLiquidated is a free log retrieval operation binding the contract event 0x1fce24c373e07f89214e9187598635036111dbb363e99f4ce498488cdc66e688.
//
// Solidity: event ClusterLiquidated(address indexed owner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SSVNetwork *SSVNetworkFilterer) FilterClusterLiquidated(opts *bind.FilterOpts, owner []common.Address) (*SSVNetworkClusterLiquidatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SSVNetwork.contract.FilterLogs(opts, "ClusterLiquidated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SSVNetworkClusterLiquidatedIterator{contract: _SSVNetwork.contract, event: "ClusterLiquidated", logs: logs, sub: sub}, nil
}

// WatchClusterLiquidated is a free log subscription operation binding the contract event 0x1fce24c373e07f89214e9187598635036111dbb363e99f4ce498488cdc66e688.
//
// Solidity: event ClusterLiquidated(address indexed owner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SSVNetwork *SSVNetworkFilterer) WatchClusterLiquidated(opts *bind.WatchOpts, sink chan<- *SSVNetworkClusterLiquidated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SSVNetwork.contract.WatchLogs(opts, "ClusterLiquidated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SSVNetworkClusterLiquidated)
				if err := _SSVNetwork.contract.UnpackLog(event, "ClusterLiquidated", log); err != nil {
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

// ParseClusterLiquidated is a log parse operation binding the contract event 0x1fce24c373e07f89214e9187598635036111dbb363e99f4ce498488cdc66e688.
//
// Solidity: event ClusterLiquidated(address indexed owner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SSVNetwork *SSVNetworkFilterer) ParseClusterLiquidated(log types.Log) (*SSVNetworkClusterLiquidated, error) {
	event := new(SSVNetworkClusterLiquidated)
	if err := _SSVNetwork.contract.UnpackLog(event, "ClusterLiquidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SSVNetworkClusterReactivatedIterator is returned from FilterClusterReactivated and is used to iterate over the raw logs and unpacked data for ClusterReactivated events raised by the SSVNetwork contract.
type SSVNetworkClusterReactivatedIterator struct {
	Event *SSVNetworkClusterReactivated // Event containing the contract specifics and raw log

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
func (it *SSVNetworkClusterReactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SSVNetworkClusterReactivated)
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
		it.Event = new(SSVNetworkClusterReactivated)
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
func (it *SSVNetworkClusterReactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SSVNetworkClusterReactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SSVNetworkClusterReactivated represents a ClusterReactivated event raised by the SSVNetwork contract.
type SSVNetworkClusterReactivated struct {
	Owner       common.Address
	OperatorIds []uint64
	Cluster     ISSVNetworkCoreCluster
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClusterReactivated is a free log retrieval operation binding the contract event 0xc803f8c01343fcdaf32068f4c283951623ef2b3fa0c547551931356f456b6859.
//
// Solidity: event ClusterReactivated(address indexed owner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SSVNetwork *SSVNetworkFilterer) FilterClusterReactivated(opts *bind.FilterOpts, owner []common.Address) (*SSVNetworkClusterReactivatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SSVNetwork.contract.FilterLogs(opts, "ClusterReactivated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SSVNetworkClusterReactivatedIterator{contract: _SSVNetwork.contract, event: "ClusterReactivated", logs: logs, sub: sub}, nil
}

// WatchClusterReactivated is a free log subscription operation binding the contract event 0xc803f8c01343fcdaf32068f4c283951623ef2b3fa0c547551931356f456b6859.
//
// Solidity: event ClusterReactivated(address indexed owner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SSVNetwork *SSVNetworkFilterer) WatchClusterReactivated(opts *bind.WatchOpts, sink chan<- *SSVNetworkClusterReactivated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SSVNetwork.contract.WatchLogs(opts, "ClusterReactivated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SSVNetworkClusterReactivated)
				if err := _SSVNetwork.contract.UnpackLog(event, "ClusterReactivated", log); err != nil {
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

// ParseClusterReactivated is a log parse operation binding the contract event 0xc803f8c01343fcdaf32068f4c283951623ef2b3fa0c547551931356f456b6859.
//
// Solidity: event ClusterReactivated(address indexed owner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SSVNetwork *SSVNetworkFilterer) ParseClusterReactivated(log types.Log) (*SSVNetworkClusterReactivated, error) {
	event := new(SSVNetworkClusterReactivated)
	if err := _SSVNetwork.contract.UnpackLog(event, "ClusterReactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SSVNetworkClusterWithdrawnIterator is returned from FilterClusterWithdrawn and is used to iterate over the raw logs and unpacked data for ClusterWithdrawn events raised by the SSVNetwork contract.
type SSVNetworkClusterWithdrawnIterator struct {
	Event *SSVNetworkClusterWithdrawn // Event containing the contract specifics and raw log

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
func (it *SSVNetworkClusterWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SSVNetworkClusterWithdrawn)
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
		it.Event = new(SSVNetworkClusterWithdrawn)
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
func (it *SSVNetworkClusterWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SSVNetworkClusterWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SSVNetworkClusterWithdrawn represents a ClusterWithdrawn event raised by the SSVNetwork contract.
type SSVNetworkClusterWithdrawn struct {
	Owner       common.Address
	OperatorIds []uint64
	Value       *big.Int
	Cluster     ISSVNetworkCoreCluster
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClusterWithdrawn is a free log retrieval operation binding the contract event 0x39d1320bbda24947e77f3560661323384aa0a1cb9d5e040e617e5cbf50b6dbe0.
//
// Solidity: event ClusterWithdrawn(address indexed owner, uint64[] operatorIds, uint256 value, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SSVNetwork *SSVNetworkFilterer) FilterClusterWithdrawn(opts *bind.FilterOpts, owner []common.Address) (*SSVNetworkClusterWithdrawnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SSVNetwork.contract.FilterLogs(opts, "ClusterWithdrawn", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SSVNetworkClusterWithdrawnIterator{contract: _SSVNetwork.contract, event: "ClusterWithdrawn", logs: logs, sub: sub}, nil
}

// WatchClusterWithdrawn is a free log subscription operation binding the contract event 0x39d1320bbda24947e77f3560661323384aa0a1cb9d5e040e617e5cbf50b6dbe0.
//
// Solidity: event ClusterWithdrawn(address indexed owner, uint64[] operatorIds, uint256 value, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SSVNetwork *SSVNetworkFilterer) WatchClusterWithdrawn(opts *bind.WatchOpts, sink chan<- *SSVNetworkClusterWithdrawn, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SSVNetwork.contract.WatchLogs(opts, "ClusterWithdrawn", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SSVNetworkClusterWithdrawn)
				if err := _SSVNetwork.contract.UnpackLog(event, "ClusterWithdrawn", log); err != nil {
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

// ParseClusterWithdrawn is a log parse operation binding the contract event 0x39d1320bbda24947e77f3560661323384aa0a1cb9d5e040e617e5cbf50b6dbe0.
//
// Solidity: event ClusterWithdrawn(address indexed owner, uint64[] operatorIds, uint256 value, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SSVNetwork *SSVNetworkFilterer) ParseClusterWithdrawn(log types.Log) (*SSVNetworkClusterWithdrawn, error) {
	event := new(SSVNetworkClusterWithdrawn)
	if err := _SSVNetwork.contract.UnpackLog(event, "ClusterWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SSVNetworkDeclareOperatorFeePeriodUpdatedIterator is returned from FilterDeclareOperatorFeePeriodUpdated and is used to iterate over the raw logs and unpacked data for DeclareOperatorFeePeriodUpdated events raised by the SSVNetwork contract.
type SSVNetworkDeclareOperatorFeePeriodUpdatedIterator struct {
	Event *SSVNetworkDeclareOperatorFeePeriodUpdated // Event containing the contract specifics and raw log

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
func (it *SSVNetworkDeclareOperatorFeePeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SSVNetworkDeclareOperatorFeePeriodUpdated)
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
		it.Event = new(SSVNetworkDeclareOperatorFeePeriodUpdated)
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
func (it *SSVNetworkDeclareOperatorFeePeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SSVNetworkDeclareOperatorFeePeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SSVNetworkDeclareOperatorFeePeriodUpdated represents a DeclareOperatorFeePeriodUpdated event raised by the SSVNetwork contract.
type SSVNetworkDeclareOperatorFeePeriodUpdated struct {
	Value uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDeclareOperatorFeePeriodUpdated is a free log retrieval operation binding the contract event 0x5fbd75d987b37490f91aa1909db948e7ff14c6ffb495b2f8e0b2334da9b192f1.
//
// Solidity: event DeclareOperatorFeePeriodUpdated(uint64 value)
func (_SSVNetwork *SSVNetworkFilterer) FilterDeclareOperatorFeePeriodUpdated(opts *bind.FilterOpts) (*SSVNetworkDeclareOperatorFeePeriodUpdatedIterator, error) {

	logs, sub, err := _SSVNetwork.contract.FilterLogs(opts, "DeclareOperatorFeePeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &SSVNetworkDeclareOperatorFeePeriodUpdatedIterator{contract: _SSVNetwork.contract, event: "DeclareOperatorFeePeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchDeclareOperatorFeePeriodUpdated is a free log subscription operation binding the contract event 0x5fbd75d987b37490f91aa1909db948e7ff14c6ffb495b2f8e0b2334da9b192f1.
//
// Solidity: event DeclareOperatorFeePeriodUpdated(uint64 value)
func (_SSVNetwork *SSVNetworkFilterer) WatchDeclareOperatorFeePeriodUpdated(opts *bind.WatchOpts, sink chan<- *SSVNetworkDeclareOperatorFeePeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _SSVNetwork.contract.WatchLogs(opts, "DeclareOperatorFeePeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SSVNetworkDeclareOperatorFeePeriodUpdated)
				if err := _SSVNetwork.contract.UnpackLog(event, "DeclareOperatorFeePeriodUpdated", log); err != nil {
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

// ParseDeclareOperatorFeePeriodUpdated is a log parse operation binding the contract event 0x5fbd75d987b37490f91aa1909db948e7ff14c6ffb495b2f8e0b2334da9b192f1.
//
// Solidity: event DeclareOperatorFeePeriodUpdated(uint64 value)
func (_SSVNetwork *SSVNetworkFilterer) ParseDeclareOperatorFeePeriodUpdated(log types.Log) (*SSVNetworkDeclareOperatorFeePeriodUpdated, error) {
	event := new(SSVNetworkDeclareOperatorFeePeriodUpdated)
	if err := _SSVNetwork.contract.UnpackLog(event, "DeclareOperatorFeePeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SSVNetworkExecuteOperatorFeePeriodUpdatedIterator is returned from FilterExecuteOperatorFeePeriodUpdated and is used to iterate over the raw logs and unpacked data for ExecuteOperatorFeePeriodUpdated events raised by the SSVNetwork contract.
type SSVNetworkExecuteOperatorFeePeriodUpdatedIterator struct {
	Event *SSVNetworkExecuteOperatorFeePeriodUpdated // Event containing the contract specifics and raw log

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
func (it *SSVNetworkExecuteOperatorFeePeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SSVNetworkExecuteOperatorFeePeriodUpdated)
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
		it.Event = new(SSVNetworkExecuteOperatorFeePeriodUpdated)
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
func (it *SSVNetworkExecuteOperatorFeePeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SSVNetworkExecuteOperatorFeePeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SSVNetworkExecuteOperatorFeePeriodUpdated represents a ExecuteOperatorFeePeriodUpdated event raised by the SSVNetwork contract.
type SSVNetworkExecuteOperatorFeePeriodUpdated struct {
	Value uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterExecuteOperatorFeePeriodUpdated is a free log retrieval operation binding the contract event 0xf6b8a2b45d0a60381de51a7b980c4660d9e5b82db6e07a4d342bfc17a6ff96bf.
//
// Solidity: event ExecuteOperatorFeePeriodUpdated(uint64 value)
func (_SSVNetwork *SSVNetworkFilterer) FilterExecuteOperatorFeePeriodUpdated(opts *bind.FilterOpts) (*SSVNetworkExecuteOperatorFeePeriodUpdatedIterator, error) {

	logs, sub, err := _SSVNetwork.contract.FilterLogs(opts, "ExecuteOperatorFeePeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &SSVNetworkExecuteOperatorFeePeriodUpdatedIterator{contract: _SSVNetwork.contract, event: "ExecuteOperatorFeePeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchExecuteOperatorFeePeriodUpdated is a free log subscription operation binding the contract event 0xf6b8a2b45d0a60381de51a7b980c4660d9e5b82db6e07a4d342bfc17a6ff96bf.
//
// Solidity: event ExecuteOperatorFeePeriodUpdated(uint64 value)
func (_SSVNetwork *SSVNetworkFilterer) WatchExecuteOperatorFeePeriodUpdated(opts *bind.WatchOpts, sink chan<- *SSVNetworkExecuteOperatorFeePeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _SSVNetwork.contract.WatchLogs(opts, "ExecuteOperatorFeePeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SSVNetworkExecuteOperatorFeePeriodUpdated)
				if err := _SSVNetwork.contract.UnpackLog(event, "ExecuteOperatorFeePeriodUpdated", log); err != nil {
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

// ParseExecuteOperatorFeePeriodUpdated is a log parse operation binding the contract event 0xf6b8a2b45d0a60381de51a7b980c4660d9e5b82db6e07a4d342bfc17a6ff96bf.
//
// Solidity: event ExecuteOperatorFeePeriodUpdated(uint64 value)
func (_SSVNetwork *SSVNetworkFilterer) ParseExecuteOperatorFeePeriodUpdated(log types.Log) (*SSVNetworkExecuteOperatorFeePeriodUpdated, error) {
	event := new(SSVNetworkExecuteOperatorFeePeriodUpdated)
	if err := _SSVNetwork.contract.UnpackLog(event, "ExecuteOperatorFeePeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SSVNetworkFeeRecipientAddressUpdatedIterator is returned from FilterFeeRecipientAddressUpdated and is used to iterate over the raw logs and unpacked data for FeeRecipientAddressUpdated events raised by the SSVNetwork contract.
type SSVNetworkFeeRecipientAddressUpdatedIterator struct {
	Event *SSVNetworkFeeRecipientAddressUpdated // Event containing the contract specifics and raw log

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
func (it *SSVNetworkFeeRecipientAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SSVNetworkFeeRecipientAddressUpdated)
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
		it.Event = new(SSVNetworkFeeRecipientAddressUpdated)
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
func (it *SSVNetworkFeeRecipientAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SSVNetworkFeeRecipientAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SSVNetworkFeeRecipientAddressUpdated represents a FeeRecipientAddressUpdated event raised by the SSVNetwork contract.
type SSVNetworkFeeRecipientAddressUpdated struct {
	Owner            common.Address
	RecipientAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFeeRecipientAddressUpdated is a free log retrieval operation binding the contract event 0x259235c230d57def1521657e7c7951d3b385e76193378bc87ef6b56bc2ec3548.
//
// Solidity: event FeeRecipientAddressUpdated(address indexed owner, address recipientAddress)
func (_SSVNetwork *SSVNetworkFilterer) FilterFeeRecipientAddressUpdated(opts *bind.FilterOpts, owner []common.Address) (*SSVNetworkFeeRecipientAddressUpdatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SSVNetwork.contract.FilterLogs(opts, "FeeRecipientAddressUpdated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SSVNetworkFeeRecipientAddressUpdatedIterator{contract: _SSVNetwork.contract, event: "FeeRecipientAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeRecipientAddressUpdated is a free log subscription operation binding the contract event 0x259235c230d57def1521657e7c7951d3b385e76193378bc87ef6b56bc2ec3548.
//
// Solidity: event FeeRecipientAddressUpdated(address indexed owner, address recipientAddress)
func (_SSVNetwork *SSVNetworkFilterer) WatchFeeRecipientAddressUpdated(opts *bind.WatchOpts, sink chan<- *SSVNetworkFeeRecipientAddressUpdated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SSVNetwork.contract.WatchLogs(opts, "FeeRecipientAddressUpdated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SSVNetworkFeeRecipientAddressUpdated)
				if err := _SSVNetwork.contract.UnpackLog(event, "FeeRecipientAddressUpdated", log); err != nil {
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

// ParseFeeRecipientAddressUpdated is a log parse operation binding the contract event 0x259235c230d57def1521657e7c7951d3b385e76193378bc87ef6b56bc2ec3548.
//
// Solidity: event FeeRecipientAddressUpdated(address indexed owner, address recipientAddress)
func (_SSVNetwork *SSVNetworkFilterer) ParseFeeRecipientAddressUpdated(log types.Log) (*SSVNetworkFeeRecipientAddressUpdated, error) {
	event := new(SSVNetworkFeeRecipientAddressUpdated)
	if err := _SSVNetwork.contract.UnpackLog(event, "FeeRecipientAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SSVNetworkInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SSVNetwork contract.
type SSVNetworkInitializedIterator struct {
	Event *SSVNetworkInitialized // Event containing the contract specifics and raw log

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
func (it *SSVNetworkInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SSVNetworkInitialized)
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
		it.Event = new(SSVNetworkInitialized)
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
func (it *SSVNetworkInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SSVNetworkInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SSVNetworkInitialized represents a Initialized event raised by the SSVNetwork contract.
type SSVNetworkInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SSVNetwork *SSVNetworkFilterer) FilterInitialized(opts *bind.FilterOpts) (*SSVNetworkInitializedIterator, error) {

	logs, sub, err := _SSVNetwork.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SSVNetworkInitializedIterator{contract: _SSVNetwork.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SSVNetwork *SSVNetworkFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SSVNetworkInitialized) (event.Subscription, error) {

	logs, sub, err := _SSVNetwork.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SSVNetworkInitialized)
				if err := _SSVNetwork.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SSVNetwork *SSVNetworkFilterer) ParseInitialized(log types.Log) (*SSVNetworkInitialized, error) {
	event := new(SSVNetworkInitialized)
	if err := _SSVNetwork.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SSVNetworkLiquidationThresholdPeriodUpdatedIterator is returned from FilterLiquidationThresholdPeriodUpdated and is used to iterate over the raw logs and unpacked data for LiquidationThresholdPeriodUpdated events raised by the SSVNetwork contract.
type SSVNetworkLiquidationThresholdPeriodUpdatedIterator struct {
	Event *SSVNetworkLiquidationThresholdPeriodUpdated // Event containing the contract specifics and raw log

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
func (it *SSVNetworkLiquidationThresholdPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SSVNetworkLiquidationThresholdPeriodUpdated)
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
		it.Event = new(SSVNetworkLiquidationThresholdPeriodUpdated)
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
func (it *SSVNetworkLiquidationThresholdPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SSVNetworkLiquidationThresholdPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SSVNetworkLiquidationThresholdPeriodUpdated represents a LiquidationThresholdPeriodUpdated event raised by the SSVNetwork contract.
type SSVNetworkLiquidationThresholdPeriodUpdated struct {
	Value uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLiquidationThresholdPeriodUpdated is a free log retrieval operation binding the contract event 0x42af14411036d7a50e5e92daf825781450fc8fac8fb65cbdb04720ff08efb84f.
//
// Solidity: event LiquidationThresholdPeriodUpdated(uint64 value)
func (_SSVNetwork *SSVNetworkFilterer) FilterLiquidationThresholdPeriodUpdated(opts *bind.FilterOpts) (*SSVNetworkLiquidationThresholdPeriodUpdatedIterator, error) {

	logs, sub, err := _SSVNetwork.contract.FilterLogs(opts, "LiquidationThresholdPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &SSVNetworkLiquidationThresholdPeriodUpdatedIterator{contract: _SSVNetwork.contract, event: "LiquidationThresholdPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchLiquidationThresholdPeriodUpdated is a free log subscription operation binding the contract event 0x42af14411036d7a50e5e92daf825781450fc8fac8fb65cbdb04720ff08efb84f.
//
// Solidity: event LiquidationThresholdPeriodUpdated(uint64 value)
func (_SSVNetwork *SSVNetworkFilterer) WatchLiquidationThresholdPeriodUpdated(opts *bind.WatchOpts, sink chan<- *SSVNetworkLiquidationThresholdPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _SSVNetwork.contract.WatchLogs(opts, "LiquidationThresholdPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SSVNetworkLiquidationThresholdPeriodUpdated)
				if err := _SSVNetwork.contract.UnpackLog(event, "LiquidationThresholdPeriodUpdated", log); err != nil {
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

// ParseLiquidationThresholdPeriodUpdated is a log parse operation binding the contract event 0x42af14411036d7a50e5e92daf825781450fc8fac8fb65cbdb04720ff08efb84f.
//
// Solidity: event LiquidationThresholdPeriodUpdated(uint64 value)
func (_SSVNetwork *SSVNetworkFilterer) ParseLiquidationThresholdPeriodUpdated(log types.Log) (*SSVNetworkLiquidationThresholdPeriodUpdated, error) {
	event := new(SSVNetworkLiquidationThresholdPeriodUpdated)
	if err := _SSVNetwork.contract.UnpackLog(event, "LiquidationThresholdPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SSVNetworkMinimumLiquidationCollateralUpdatedIterator is returned from FilterMinimumLiquidationCollateralUpdated and is used to iterate over the raw logs and unpacked data for MinimumLiquidationCollateralUpdated events raised by the SSVNetwork contract.
type SSVNetworkMinimumLiquidationCollateralUpdatedIterator struct {
	Event *SSVNetworkMinimumLiquidationCollateralUpdated // Event containing the contract specifics and raw log

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
func (it *SSVNetworkMinimumLiquidationCollateralUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SSVNetworkMinimumLiquidationCollateralUpdated)
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
		it.Event = new(SSVNetworkMinimumLiquidationCollateralUpdated)
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
func (it *SSVNetworkMinimumLiquidationCollateralUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SSVNetworkMinimumLiquidationCollateralUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SSVNetworkMinimumLiquidationCollateralUpdated represents a MinimumLiquidationCollateralUpdated event raised by the SSVNetwork contract.
type SSVNetworkMinimumLiquidationCollateralUpdated struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinimumLiquidationCollateralUpdated is a free log retrieval operation binding the contract event 0xd363ab4392efaf967a89d8616cba1ff0c6f05a04c2f214671be365f0fab05960.
//
// Solidity: event MinimumLiquidationCollateralUpdated(uint256 value)
func (_SSVNetwork *SSVNetworkFilterer) FilterMinimumLiquidationCollateralUpdated(opts *bind.FilterOpts) (*SSVNetworkMinimumLiquidationCollateralUpdatedIterator, error) {

	logs, sub, err := _SSVNetwork.contract.FilterLogs(opts, "MinimumLiquidationCollateralUpdated")
	if err != nil {
		return nil, err
	}
	return &SSVNetworkMinimumLiquidationCollateralUpdatedIterator{contract: _SSVNetwork.contract, event: "MinimumLiquidationCollateralUpdated", logs: logs, sub: sub}, nil
}

// WatchMinimumLiquidationCollateralUpdated is a free log subscription operation binding the contract event 0xd363ab4392efaf967a89d8616cba1ff0c6f05a04c2f214671be365f0fab05960.
//
// Solidity: event MinimumLiquidationCollateralUpdated(uint256 value)
func (_SSVNetwork *SSVNetworkFilterer) WatchMinimumLiquidationCollateralUpdated(opts *bind.WatchOpts, sink chan<- *SSVNetworkMinimumLiquidationCollateralUpdated) (event.Subscription, error) {

	logs, sub, err := _SSVNetwork.contract.WatchLogs(opts, "MinimumLiquidationCollateralUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SSVNetworkMinimumLiquidationCollateralUpdated)
				if err := _SSVNetwork.contract.UnpackLog(event, "MinimumLiquidationCollateralUpdated", log); err != nil {
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

// ParseMinimumLiquidationCollateralUpdated is a log parse operation binding the contract event 0xd363ab4392efaf967a89d8616cba1ff0c6f05a04c2f214671be365f0fab05960.
//
// Solidity: event MinimumLiquidationCollateralUpdated(uint256 value)
func (_SSVNetwork *SSVNetworkFilterer) ParseMinimumLiquidationCollateralUpdated(log types.Log) (*SSVNetworkMinimumLiquidationCollateralUpdated, error) {
	event := new(SSVNetworkMinimumLiquidationCollateralUpdated)
	if err := _SSVNetwork.contract.UnpackLog(event, "MinimumLiquidationCollateralUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SSVNetworkNetworkEarningsWithdrawnIterator is returned from FilterNetworkEarningsWithdrawn and is used to iterate over the raw logs and unpacked data for NetworkEarningsWithdrawn events raised by the SSVNetwork contract.
type SSVNetworkNetworkEarningsWithdrawnIterator struct {
	Event *SSVNetworkNetworkEarningsWithdrawn // Event containing the contract specifics and raw log

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
func (it *SSVNetworkNetworkEarningsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SSVNetworkNetworkEarningsWithdrawn)
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
		it.Event = new(SSVNetworkNetworkEarningsWithdrawn)
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
func (it *SSVNetworkNetworkEarningsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SSVNetworkNetworkEarningsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SSVNetworkNetworkEarningsWithdrawn represents a NetworkEarningsWithdrawn event raised by the SSVNetwork contract.
type SSVNetworkNetworkEarningsWithdrawn struct {
	Value     *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNetworkEarningsWithdrawn is a free log retrieval operation binding the contract event 0x370342c3bb9245e20bffe6dced02ba2fceca979701f881d5adc72d838e83f1c5.
//
// Solidity: event NetworkEarningsWithdrawn(uint256 value, address recipient)
func (_SSVNetwork *SSVNetworkFilterer) FilterNetworkEarningsWithdrawn(opts *bind.FilterOpts) (*SSVNetworkNetworkEarningsWithdrawnIterator, error) {

	logs, sub, err := _SSVNetwork.contract.FilterLogs(opts, "NetworkEarningsWithdrawn")
	if err != nil {
		return nil, err
	}
	return &SSVNetworkNetworkEarningsWithdrawnIterator{contract: _SSVNetwork.contract, event: "NetworkEarningsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchNetworkEarningsWithdrawn is a free log subscription operation binding the contract event 0x370342c3bb9245e20bffe6dced02ba2fceca979701f881d5adc72d838e83f1c5.
//
// Solidity: event NetworkEarningsWithdrawn(uint256 value, address recipient)
func (_SSVNetwork *SSVNetworkFilterer) WatchNetworkEarningsWithdrawn(opts *bind.WatchOpts, sink chan<- *SSVNetworkNetworkEarningsWithdrawn) (event.Subscription, error) {

	logs, sub, err := _SSVNetwork.contract.WatchLogs(opts, "NetworkEarningsWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SSVNetworkNetworkEarningsWithdrawn)
				if err := _SSVNetwork.contract.UnpackLog(event, "NetworkEarningsWithdrawn", log); err != nil {
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

// ParseNetworkEarningsWithdrawn is a log parse operation binding the contract event 0x370342c3bb9245e20bffe6dced02ba2fceca979701f881d5adc72d838e83f1c5.
//
// Solidity: event NetworkEarningsWithdrawn(uint256 value, address recipient)
func (_SSVNetwork *SSVNetworkFilterer) ParseNetworkEarningsWithdrawn(log types.Log) (*SSVNetworkNetworkEarningsWithdrawn, error) {
	event := new(SSVNetworkNetworkEarningsWithdrawn)
	if err := _SSVNetwork.contract.UnpackLog(event, "NetworkEarningsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SSVNetworkNetworkFeeUpdatedIterator is returned from FilterNetworkFeeUpdated and is used to iterate over the raw logs and unpacked data for NetworkFeeUpdated events raised by the SSVNetwork contract.
type SSVNetworkNetworkFeeUpdatedIterator struct {
	Event *SSVNetworkNetworkFeeUpdated // Event containing the contract specifics and raw log

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
func (it *SSVNetworkNetworkFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SSVNetworkNetworkFeeUpdated)
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
		it.Event = new(SSVNetworkNetworkFeeUpdated)
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
func (it *SSVNetworkNetworkFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SSVNetworkNetworkFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SSVNetworkNetworkFeeUpdated represents a NetworkFeeUpdated event raised by the SSVNetwork contract.
type SSVNetworkNetworkFeeUpdated struct {
	OldFee *big.Int
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNetworkFeeUpdated is a free log retrieval operation binding the contract event 0x8f49a76c5d617bd72673d92d3a019ff8f04f204536aae7a3d10e7ca85603f3cc.
//
// Solidity: event NetworkFeeUpdated(uint256 oldFee, uint256 newFee)
func (_SSVNetwork *SSVNetworkFilterer) FilterNetworkFeeUpdated(opts *bind.FilterOpts) (*SSVNetworkNetworkFeeUpdatedIterator, error) {

	logs, sub, err := _SSVNetwork.contract.FilterLogs(opts, "NetworkFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &SSVNetworkNetworkFeeUpdatedIterator{contract: _SSVNetwork.contract, event: "NetworkFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchNetworkFeeUpdated is a free log subscription operation binding the contract event 0x8f49a76c5d617bd72673d92d3a019ff8f04f204536aae7a3d10e7ca85603f3cc.
//
// Solidity: event NetworkFeeUpdated(uint256 oldFee, uint256 newFee)
func (_SSVNetwork *SSVNetworkFilterer) WatchNetworkFeeUpdated(opts *bind.WatchOpts, sink chan<- *SSVNetworkNetworkFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _SSVNetwork.contract.WatchLogs(opts, "NetworkFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SSVNetworkNetworkFeeUpdated)
				if err := _SSVNetwork.contract.UnpackLog(event, "NetworkFeeUpdated", log); err != nil {
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

// ParseNetworkFeeUpdated is a log parse operation binding the contract event 0x8f49a76c5d617bd72673d92d3a019ff8f04f204536aae7a3d10e7ca85603f3cc.
//
// Solidity: event NetworkFeeUpdated(uint256 oldFee, uint256 newFee)
func (_SSVNetwork *SSVNetworkFilterer) ParseNetworkFeeUpdated(log types.Log) (*SSVNetworkNetworkFeeUpdated, error) {
	event := new(SSVNetworkNetworkFeeUpdated)
	if err := _SSVNetwork.contract.UnpackLog(event, "NetworkFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SSVNetworkOperatorAddedIterator is returned from FilterOperatorAdded and is used to iterate over the raw logs and unpacked data for OperatorAdded events raised by the SSVNetwork contract.
type SSVNetworkOperatorAddedIterator struct {
	Event *SSVNetworkOperatorAdded // Event containing the contract specifics and raw log

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
func (it *SSVNetworkOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SSVNetworkOperatorAdded)
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
		it.Event = new(SSVNetworkOperatorAdded)
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
func (it *SSVNetworkOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SSVNetworkOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SSVNetworkOperatorAdded represents a OperatorAdded event raised by the SSVNetwork contract.
type SSVNetworkOperatorAdded struct {
	OperatorId uint64
	Owner      common.Address
	PublicKey  []byte
	Fee        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorAdded is a free log retrieval operation binding the contract event 0xd839f31c14bd632f424e307b36abff63ca33684f77f28e35dc13718ef338f7f4.
//
// Solidity: event OperatorAdded(uint64 indexed operatorId, address indexed owner, bytes publicKey, uint256 fee)
func (_SSVNetwork *SSVNetworkFilterer) FilterOperatorAdded(opts *bind.FilterOpts, operatorId []uint64, owner []common.Address) (*SSVNetworkOperatorAddedIterator, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SSVNetwork.contract.FilterLogs(opts, "OperatorAdded", operatorIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &SSVNetworkOperatorAddedIterator{contract: _SSVNetwork.contract, event: "OperatorAdded", logs: logs, sub: sub}, nil
}

// WatchOperatorAdded is a free log subscription operation binding the contract event 0xd839f31c14bd632f424e307b36abff63ca33684f77f28e35dc13718ef338f7f4.
//
// Solidity: event OperatorAdded(uint64 indexed operatorId, address indexed owner, bytes publicKey, uint256 fee)
func (_SSVNetwork *SSVNetworkFilterer) WatchOperatorAdded(opts *bind.WatchOpts, sink chan<- *SSVNetworkOperatorAdded, operatorId []uint64, owner []common.Address) (event.Subscription, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SSVNetwork.contract.WatchLogs(opts, "OperatorAdded", operatorIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SSVNetworkOperatorAdded)
				if err := _SSVNetwork.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
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

// ParseOperatorAdded is a log parse operation binding the contract event 0xd839f31c14bd632f424e307b36abff63ca33684f77f28e35dc13718ef338f7f4.
//
// Solidity: event OperatorAdded(uint64 indexed operatorId, address indexed owner, bytes publicKey, uint256 fee)
func (_SSVNetwork *SSVNetworkFilterer) ParseOperatorAdded(log types.Log) (*SSVNetworkOperatorAdded, error) {
	event := new(SSVNetworkOperatorAdded)
	if err := _SSVNetwork.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SSVNetworkOperatorFeeCancellationDeclaredIterator is returned from FilterOperatorFeeCancellationDeclared and is used to iterate over the raw logs and unpacked data for OperatorFeeCancellationDeclared events raised by the SSVNetwork contract.
type SSVNetworkOperatorFeeCancellationDeclaredIterator struct {
	Event *SSVNetworkOperatorFeeCancellationDeclared // Event containing the contract specifics and raw log

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
func (it *SSVNetworkOperatorFeeCancellationDeclaredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SSVNetworkOperatorFeeCancellationDeclared)
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
		it.Event = new(SSVNetworkOperatorFeeCancellationDeclared)
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
func (it *SSVNetworkOperatorFeeCancellationDeclaredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SSVNetworkOperatorFeeCancellationDeclaredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SSVNetworkOperatorFeeCancellationDeclared represents a OperatorFeeCancellationDeclared event raised by the SSVNetwork contract.
type SSVNetworkOperatorFeeCancellationDeclared struct {
	Owner      common.Address
	OperatorId uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorFeeCancellationDeclared is a free log retrieval operation binding the contract event 0x269e5b30b1f6e64e296fff57d56c77a4eb77fda7421146f1497f567539ba8906.
//
// Solidity: event OperatorFeeCancellationDeclared(address indexed owner, uint64 indexed operatorId)
func (_SSVNetwork *SSVNetworkFilterer) FilterOperatorFeeCancellationDeclared(opts *bind.FilterOpts, owner []common.Address, operatorId []uint64) (*SSVNetworkOperatorFeeCancellationDeclaredIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _SSVNetwork.contract.FilterLogs(opts, "OperatorFeeCancellationDeclared", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &SSVNetworkOperatorFeeCancellationDeclaredIterator{contract: _SSVNetwork.contract, event: "OperatorFeeCancellationDeclared", logs: logs, sub: sub}, nil
}

// WatchOperatorFeeCancellationDeclared is a free log subscription operation binding the contract event 0x269e5b30b1f6e64e296fff57d56c77a4eb77fda7421146f1497f567539ba8906.
//
// Solidity: event OperatorFeeCancellationDeclared(address indexed owner, uint64 indexed operatorId)
func (_SSVNetwork *SSVNetworkFilterer) WatchOperatorFeeCancellationDeclared(opts *bind.WatchOpts, sink chan<- *SSVNetworkOperatorFeeCancellationDeclared, owner []common.Address, operatorId []uint64) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _SSVNetwork.contract.WatchLogs(opts, "OperatorFeeCancellationDeclared", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SSVNetworkOperatorFeeCancellationDeclared)
				if err := _SSVNetwork.contract.UnpackLog(event, "OperatorFeeCancellationDeclared", log); err != nil {
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

// ParseOperatorFeeCancellationDeclared is a log parse operation binding the contract event 0x269e5b30b1f6e64e296fff57d56c77a4eb77fda7421146f1497f567539ba8906.
//
// Solidity: event OperatorFeeCancellationDeclared(address indexed owner, uint64 indexed operatorId)
func (_SSVNetwork *SSVNetworkFilterer) ParseOperatorFeeCancellationDeclared(log types.Log) (*SSVNetworkOperatorFeeCancellationDeclared, error) {
	event := new(SSVNetworkOperatorFeeCancellationDeclared)
	if err := _SSVNetwork.contract.UnpackLog(event, "OperatorFeeCancellationDeclared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SSVNetworkOperatorFeeDeclaredIterator is returned from FilterOperatorFeeDeclared and is used to iterate over the raw logs and unpacked data for OperatorFeeDeclared events raised by the SSVNetwork contract.
type SSVNetworkOperatorFeeDeclaredIterator struct {
	Event *SSVNetworkOperatorFeeDeclared // Event containing the contract specifics and raw log

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
func (it *SSVNetworkOperatorFeeDeclaredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SSVNetworkOperatorFeeDeclared)
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
		it.Event = new(SSVNetworkOperatorFeeDeclared)
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
func (it *SSVNetworkOperatorFeeDeclaredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SSVNetworkOperatorFeeDeclaredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SSVNetworkOperatorFeeDeclared represents a OperatorFeeDeclared event raised by the SSVNetwork contract.
type SSVNetworkOperatorFeeDeclared struct {
	Owner       common.Address
	OperatorId  uint64
	BlockNumber *big.Int
	Fee         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorFeeDeclared is a free log retrieval operation binding the contract event 0x796204296f2eb56d7432fa85961e9750d0cb21741873ebf7077e28263e327358.
//
// Solidity: event OperatorFeeDeclared(address indexed owner, uint64 indexed operatorId, uint256 blockNumber, uint256 fee)
func (_SSVNetwork *SSVNetworkFilterer) FilterOperatorFeeDeclared(opts *bind.FilterOpts, owner []common.Address, operatorId []uint64) (*SSVNetworkOperatorFeeDeclaredIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _SSVNetwork.contract.FilterLogs(opts, "OperatorFeeDeclared", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &SSVNetworkOperatorFeeDeclaredIterator{contract: _SSVNetwork.contract, event: "OperatorFeeDeclared", logs: logs, sub: sub}, nil
}

// WatchOperatorFeeDeclared is a free log subscription operation binding the contract event 0x796204296f2eb56d7432fa85961e9750d0cb21741873ebf7077e28263e327358.
//
// Solidity: event OperatorFeeDeclared(address indexed owner, uint64 indexed operatorId, uint256 blockNumber, uint256 fee)
func (_SSVNetwork *SSVNetworkFilterer) WatchOperatorFeeDeclared(opts *bind.WatchOpts, sink chan<- *SSVNetworkOperatorFeeDeclared, owner []common.Address, operatorId []uint64) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _SSVNetwork.contract.WatchLogs(opts, "OperatorFeeDeclared", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SSVNetworkOperatorFeeDeclared)
				if err := _SSVNetwork.contract.UnpackLog(event, "OperatorFeeDeclared", log); err != nil {
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

// ParseOperatorFeeDeclared is a log parse operation binding the contract event 0x796204296f2eb56d7432fa85961e9750d0cb21741873ebf7077e28263e327358.
//
// Solidity: event OperatorFeeDeclared(address indexed owner, uint64 indexed operatorId, uint256 blockNumber, uint256 fee)
func (_SSVNetwork *SSVNetworkFilterer) ParseOperatorFeeDeclared(log types.Log) (*SSVNetworkOperatorFeeDeclared, error) {
	event := new(SSVNetworkOperatorFeeDeclared)
	if err := _SSVNetwork.contract.UnpackLog(event, "OperatorFeeDeclared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SSVNetworkOperatorFeeExecutedIterator is returned from FilterOperatorFeeExecuted and is used to iterate over the raw logs and unpacked data for OperatorFeeExecuted events raised by the SSVNetwork contract.
type SSVNetworkOperatorFeeExecutedIterator struct {
	Event *SSVNetworkOperatorFeeExecuted // Event containing the contract specifics and raw log

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
func (it *SSVNetworkOperatorFeeExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SSVNetworkOperatorFeeExecuted)
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
		it.Event = new(SSVNetworkOperatorFeeExecuted)
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
func (it *SSVNetworkOperatorFeeExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SSVNetworkOperatorFeeExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SSVNetworkOperatorFeeExecuted represents a OperatorFeeExecuted event raised by the SSVNetwork contract.
type SSVNetworkOperatorFeeExecuted struct {
	Owner       common.Address
	OperatorId  uint64
	BlockNumber *big.Int
	Fee         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorFeeExecuted is a free log retrieval operation binding the contract event 0x513e931ff778ed01e676d55880d8db185c29b0094546ff2b3e9f5b6920d16bef.
//
// Solidity: event OperatorFeeExecuted(address indexed owner, uint64 indexed operatorId, uint256 blockNumber, uint256 fee)
func (_SSVNetwork *SSVNetworkFilterer) FilterOperatorFeeExecuted(opts *bind.FilterOpts, owner []common.Address, operatorId []uint64) (*SSVNetworkOperatorFeeExecutedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _SSVNetwork.contract.FilterLogs(opts, "OperatorFeeExecuted", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &SSVNetworkOperatorFeeExecutedIterator{contract: _SSVNetwork.contract, event: "OperatorFeeExecuted", logs: logs, sub: sub}, nil
}

// WatchOperatorFeeExecuted is a free log subscription operation binding the contract event 0x513e931ff778ed01e676d55880d8db185c29b0094546ff2b3e9f5b6920d16bef.
//
// Solidity: event OperatorFeeExecuted(address indexed owner, uint64 indexed operatorId, uint256 blockNumber, uint256 fee)
func (_SSVNetwork *SSVNetworkFilterer) WatchOperatorFeeExecuted(opts *bind.WatchOpts, sink chan<- *SSVNetworkOperatorFeeExecuted, owner []common.Address, operatorId []uint64) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _SSVNetwork.contract.WatchLogs(opts, "OperatorFeeExecuted", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SSVNetworkOperatorFeeExecuted)
				if err := _SSVNetwork.contract.UnpackLog(event, "OperatorFeeExecuted", log); err != nil {
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

// ParseOperatorFeeExecuted is a log parse operation binding the contract event 0x513e931ff778ed01e676d55880d8db185c29b0094546ff2b3e9f5b6920d16bef.
//
// Solidity: event OperatorFeeExecuted(address indexed owner, uint64 indexed operatorId, uint256 blockNumber, uint256 fee)
func (_SSVNetwork *SSVNetworkFilterer) ParseOperatorFeeExecuted(log types.Log) (*SSVNetworkOperatorFeeExecuted, error) {
	event := new(SSVNetworkOperatorFeeExecuted)
	if err := _SSVNetwork.contract.UnpackLog(event, "OperatorFeeExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SSVNetworkOperatorFeeIncreaseLimitUpdatedIterator is returned from FilterOperatorFeeIncreaseLimitUpdated and is used to iterate over the raw logs and unpacked data for OperatorFeeIncreaseLimitUpdated events raised by the SSVNetwork contract.
type SSVNetworkOperatorFeeIncreaseLimitUpdatedIterator struct {
	Event *SSVNetworkOperatorFeeIncreaseLimitUpdated // Event containing the contract specifics and raw log

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
func (it *SSVNetworkOperatorFeeIncreaseLimitUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SSVNetworkOperatorFeeIncreaseLimitUpdated)
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
		it.Event = new(SSVNetworkOperatorFeeIncreaseLimitUpdated)
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
func (it *SSVNetworkOperatorFeeIncreaseLimitUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SSVNetworkOperatorFeeIncreaseLimitUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SSVNetworkOperatorFeeIncreaseLimitUpdated represents a OperatorFeeIncreaseLimitUpdated event raised by the SSVNetwork contract.
type SSVNetworkOperatorFeeIncreaseLimitUpdated struct {
	Value uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOperatorFeeIncreaseLimitUpdated is a free log retrieval operation binding the contract event 0x2fff7e5a48a4befc2c2be4d77e141f6d97907798977ce452429ec55c2658a342.
//
// Solidity: event OperatorFeeIncreaseLimitUpdated(uint64 value)
func (_SSVNetwork *SSVNetworkFilterer) FilterOperatorFeeIncreaseLimitUpdated(opts *bind.FilterOpts) (*SSVNetworkOperatorFeeIncreaseLimitUpdatedIterator, error) {

	logs, sub, err := _SSVNetwork.contract.FilterLogs(opts, "OperatorFeeIncreaseLimitUpdated")
	if err != nil {
		return nil, err
	}
	return &SSVNetworkOperatorFeeIncreaseLimitUpdatedIterator{contract: _SSVNetwork.contract, event: "OperatorFeeIncreaseLimitUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorFeeIncreaseLimitUpdated is a free log subscription operation binding the contract event 0x2fff7e5a48a4befc2c2be4d77e141f6d97907798977ce452429ec55c2658a342.
//
// Solidity: event OperatorFeeIncreaseLimitUpdated(uint64 value)
func (_SSVNetwork *SSVNetworkFilterer) WatchOperatorFeeIncreaseLimitUpdated(opts *bind.WatchOpts, sink chan<- *SSVNetworkOperatorFeeIncreaseLimitUpdated) (event.Subscription, error) {

	logs, sub, err := _SSVNetwork.contract.WatchLogs(opts, "OperatorFeeIncreaseLimitUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SSVNetworkOperatorFeeIncreaseLimitUpdated)
				if err := _SSVNetwork.contract.UnpackLog(event, "OperatorFeeIncreaseLimitUpdated", log); err != nil {
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

// ParseOperatorFeeIncreaseLimitUpdated is a log parse operation binding the contract event 0x2fff7e5a48a4befc2c2be4d77e141f6d97907798977ce452429ec55c2658a342.
//
// Solidity: event OperatorFeeIncreaseLimitUpdated(uint64 value)
func (_SSVNetwork *SSVNetworkFilterer) ParseOperatorFeeIncreaseLimitUpdated(log types.Log) (*SSVNetworkOperatorFeeIncreaseLimitUpdated, error) {
	event := new(SSVNetworkOperatorFeeIncreaseLimitUpdated)
	if err := _SSVNetwork.contract.UnpackLog(event, "OperatorFeeIncreaseLimitUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SSVNetworkOperatorRemovedIterator is returned from FilterOperatorRemoved and is used to iterate over the raw logs and unpacked data for OperatorRemoved events raised by the SSVNetwork contract.
type SSVNetworkOperatorRemovedIterator struct {
	Event *SSVNetworkOperatorRemoved // Event containing the contract specifics and raw log

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
func (it *SSVNetworkOperatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SSVNetworkOperatorRemoved)
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
		it.Event = new(SSVNetworkOperatorRemoved)
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
func (it *SSVNetworkOperatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SSVNetworkOperatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SSVNetworkOperatorRemoved represents a OperatorRemoved event raised by the SSVNetwork contract.
type SSVNetworkOperatorRemoved struct {
	OperatorId uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemoved is a free log retrieval operation binding the contract event 0x0e0ba6c2b04de36d6d509ec5bd155c43a9fe862f8052096dd54f3902a74cca3e.
//
// Solidity: event OperatorRemoved(uint64 indexed operatorId)
func (_SSVNetwork *SSVNetworkFilterer) FilterOperatorRemoved(opts *bind.FilterOpts, operatorId []uint64) (*SSVNetworkOperatorRemovedIterator, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _SSVNetwork.contract.FilterLogs(opts, "OperatorRemoved", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &SSVNetworkOperatorRemovedIterator{contract: _SSVNetwork.contract, event: "OperatorRemoved", logs: logs, sub: sub}, nil
}

// WatchOperatorRemoved is a free log subscription operation binding the contract event 0x0e0ba6c2b04de36d6d509ec5bd155c43a9fe862f8052096dd54f3902a74cca3e.
//
// Solidity: event OperatorRemoved(uint64 indexed operatorId)
func (_SSVNetwork *SSVNetworkFilterer) WatchOperatorRemoved(opts *bind.WatchOpts, sink chan<- *SSVNetworkOperatorRemoved, operatorId []uint64) (event.Subscription, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _SSVNetwork.contract.WatchLogs(opts, "OperatorRemoved", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SSVNetworkOperatorRemoved)
				if err := _SSVNetwork.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
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

// ParseOperatorRemoved is a log parse operation binding the contract event 0x0e0ba6c2b04de36d6d509ec5bd155c43a9fe862f8052096dd54f3902a74cca3e.
//
// Solidity: event OperatorRemoved(uint64 indexed operatorId)
func (_SSVNetwork *SSVNetworkFilterer) ParseOperatorRemoved(log types.Log) (*SSVNetworkOperatorRemoved, error) {
	event := new(SSVNetworkOperatorRemoved)
	if err := _SSVNetwork.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SSVNetworkOperatorWhitelistUpdatedIterator is returned from FilterOperatorWhitelistUpdated and is used to iterate over the raw logs and unpacked data for OperatorWhitelistUpdated events raised by the SSVNetwork contract.
type SSVNetworkOperatorWhitelistUpdatedIterator struct {
	Event *SSVNetworkOperatorWhitelistUpdated // Event containing the contract specifics and raw log

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
func (it *SSVNetworkOperatorWhitelistUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SSVNetworkOperatorWhitelistUpdated)
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
		it.Event = new(SSVNetworkOperatorWhitelistUpdated)
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
func (it *SSVNetworkOperatorWhitelistUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SSVNetworkOperatorWhitelistUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SSVNetworkOperatorWhitelistUpdated represents a OperatorWhitelistUpdated event raised by the SSVNetwork contract.
type SSVNetworkOperatorWhitelistUpdated struct {
	OperatorId  uint64
	Whitelisted common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorWhitelistUpdated is a free log retrieval operation binding the contract event 0x29f72634ccb172103f8c542da23de7f6cf9bce724c5bb91bd6f3a516b14c63fe.
//
// Solidity: event OperatorWhitelistUpdated(uint64 indexed operatorId, address whitelisted)
func (_SSVNetwork *SSVNetworkFilterer) FilterOperatorWhitelistUpdated(opts *bind.FilterOpts, operatorId []uint64) (*SSVNetworkOperatorWhitelistUpdatedIterator, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _SSVNetwork.contract.FilterLogs(opts, "OperatorWhitelistUpdated", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &SSVNetworkOperatorWhitelistUpdatedIterator{contract: _SSVNetwork.contract, event: "OperatorWhitelistUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorWhitelistUpdated is a free log subscription operation binding the contract event 0x29f72634ccb172103f8c542da23de7f6cf9bce724c5bb91bd6f3a516b14c63fe.
//
// Solidity: event OperatorWhitelistUpdated(uint64 indexed operatorId, address whitelisted)
func (_SSVNetwork *SSVNetworkFilterer) WatchOperatorWhitelistUpdated(opts *bind.WatchOpts, sink chan<- *SSVNetworkOperatorWhitelistUpdated, operatorId []uint64) (event.Subscription, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _SSVNetwork.contract.WatchLogs(opts, "OperatorWhitelistUpdated", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SSVNetworkOperatorWhitelistUpdated)
				if err := _SSVNetwork.contract.UnpackLog(event, "OperatorWhitelistUpdated", log); err != nil {
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

// ParseOperatorWhitelistUpdated is a log parse operation binding the contract event 0x29f72634ccb172103f8c542da23de7f6cf9bce724c5bb91bd6f3a516b14c63fe.
//
// Solidity: event OperatorWhitelistUpdated(uint64 indexed operatorId, address whitelisted)
func (_SSVNetwork *SSVNetworkFilterer) ParseOperatorWhitelistUpdated(log types.Log) (*SSVNetworkOperatorWhitelistUpdated, error) {
	event := new(SSVNetworkOperatorWhitelistUpdated)
	if err := _SSVNetwork.contract.UnpackLog(event, "OperatorWhitelistUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SSVNetworkOperatorWithdrawnIterator is returned from FilterOperatorWithdrawn and is used to iterate over the raw logs and unpacked data for OperatorWithdrawn events raised by the SSVNetwork contract.
type SSVNetworkOperatorWithdrawnIterator struct {
	Event *SSVNetworkOperatorWithdrawn // Event containing the contract specifics and raw log

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
func (it *SSVNetworkOperatorWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SSVNetworkOperatorWithdrawn)
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
		it.Event = new(SSVNetworkOperatorWithdrawn)
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
func (it *SSVNetworkOperatorWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SSVNetworkOperatorWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SSVNetworkOperatorWithdrawn represents a OperatorWithdrawn event raised by the SSVNetwork contract.
type SSVNetworkOperatorWithdrawn struct {
	Owner      common.Address
	OperatorId uint64
	Value      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorWithdrawn is a free log retrieval operation binding the contract event 0x178bf78bdd8914b8483d640b4a4f84e20943b5eb6b639b7474286364c7651d60.
//
// Solidity: event OperatorWithdrawn(address indexed owner, uint64 indexed operatorId, uint256 value)
func (_SSVNetwork *SSVNetworkFilterer) FilterOperatorWithdrawn(opts *bind.FilterOpts, owner []common.Address, operatorId []uint64) (*SSVNetworkOperatorWithdrawnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _SSVNetwork.contract.FilterLogs(opts, "OperatorWithdrawn", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &SSVNetworkOperatorWithdrawnIterator{contract: _SSVNetwork.contract, event: "OperatorWithdrawn", logs: logs, sub: sub}, nil
}

// WatchOperatorWithdrawn is a free log subscription operation binding the contract event 0x178bf78bdd8914b8483d640b4a4f84e20943b5eb6b639b7474286364c7651d60.
//
// Solidity: event OperatorWithdrawn(address indexed owner, uint64 indexed operatorId, uint256 value)
func (_SSVNetwork *SSVNetworkFilterer) WatchOperatorWithdrawn(opts *bind.WatchOpts, sink chan<- *SSVNetworkOperatorWithdrawn, owner []common.Address, operatorId []uint64) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _SSVNetwork.contract.WatchLogs(opts, "OperatorWithdrawn", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SSVNetworkOperatorWithdrawn)
				if err := _SSVNetwork.contract.UnpackLog(event, "OperatorWithdrawn", log); err != nil {
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

// ParseOperatorWithdrawn is a log parse operation binding the contract event 0x178bf78bdd8914b8483d640b4a4f84e20943b5eb6b639b7474286364c7651d60.
//
// Solidity: event OperatorWithdrawn(address indexed owner, uint64 indexed operatorId, uint256 value)
func (_SSVNetwork *SSVNetworkFilterer) ParseOperatorWithdrawn(log types.Log) (*SSVNetworkOperatorWithdrawn, error) {
	event := new(SSVNetworkOperatorWithdrawn)
	if err := _SSVNetwork.contract.UnpackLog(event, "OperatorWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SSVNetworkOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the SSVNetwork contract.
type SSVNetworkOwnershipTransferStartedIterator struct {
	Event *SSVNetworkOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *SSVNetworkOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SSVNetworkOwnershipTransferStarted)
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
		it.Event = new(SSVNetworkOwnershipTransferStarted)
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
func (it *SSVNetworkOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SSVNetworkOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SSVNetworkOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the SSVNetwork contract.
type SSVNetworkOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_SSVNetwork *SSVNetworkFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SSVNetworkOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SSVNetwork.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SSVNetworkOwnershipTransferStartedIterator{contract: _SSVNetwork.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_SSVNetwork *SSVNetworkFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *SSVNetworkOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SSVNetwork.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SSVNetworkOwnershipTransferStarted)
				if err := _SSVNetwork.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_SSVNetwork *SSVNetworkFilterer) ParseOwnershipTransferStarted(log types.Log) (*SSVNetworkOwnershipTransferStarted, error) {
	event := new(SSVNetworkOwnershipTransferStarted)
	if err := _SSVNetwork.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SSVNetworkOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SSVNetwork contract.
type SSVNetworkOwnershipTransferredIterator struct {
	Event *SSVNetworkOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SSVNetworkOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SSVNetworkOwnershipTransferred)
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
		it.Event = new(SSVNetworkOwnershipTransferred)
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
func (it *SSVNetworkOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SSVNetworkOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SSVNetworkOwnershipTransferred represents a OwnershipTransferred event raised by the SSVNetwork contract.
type SSVNetworkOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SSVNetwork *SSVNetworkFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SSVNetworkOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SSVNetwork.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SSVNetworkOwnershipTransferredIterator{contract: _SSVNetwork.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SSVNetwork *SSVNetworkFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SSVNetworkOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SSVNetwork.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SSVNetworkOwnershipTransferred)
				if err := _SSVNetwork.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SSVNetwork *SSVNetworkFilterer) ParseOwnershipTransferred(log types.Log) (*SSVNetworkOwnershipTransferred, error) {
	event := new(SSVNetworkOwnershipTransferred)
	if err := _SSVNetwork.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SSVNetworkUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the SSVNetwork contract.
type SSVNetworkUpgradedIterator struct {
	Event *SSVNetworkUpgraded // Event containing the contract specifics and raw log

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
func (it *SSVNetworkUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SSVNetworkUpgraded)
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
		it.Event = new(SSVNetworkUpgraded)
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
func (it *SSVNetworkUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SSVNetworkUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SSVNetworkUpgraded represents a Upgraded event raised by the SSVNetwork contract.
type SSVNetworkUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SSVNetwork *SSVNetworkFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*SSVNetworkUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SSVNetwork.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &SSVNetworkUpgradedIterator{contract: _SSVNetwork.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SSVNetwork *SSVNetworkFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *SSVNetworkUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SSVNetwork.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SSVNetworkUpgraded)
				if err := _SSVNetwork.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_SSVNetwork *SSVNetworkFilterer) ParseUpgraded(log types.Log) (*SSVNetworkUpgraded, error) {
	event := new(SSVNetworkUpgraded)
	if err := _SSVNetwork.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SSVNetworkValidatorAddedIterator is returned from FilterValidatorAdded and is used to iterate over the raw logs and unpacked data for ValidatorAdded events raised by the SSVNetwork contract.
type SSVNetworkValidatorAddedIterator struct {
	Event *SSVNetworkValidatorAdded // Event containing the contract specifics and raw log

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
func (it *SSVNetworkValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SSVNetworkValidatorAdded)
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
		it.Event = new(SSVNetworkValidatorAdded)
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
func (it *SSVNetworkValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SSVNetworkValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SSVNetworkValidatorAdded represents a ValidatorAdded event raised by the SSVNetwork contract.
type SSVNetworkValidatorAdded struct {
	Owner       common.Address
	OperatorIds []uint64
	PublicKey   []byte
	Shares      []byte
	Cluster     ISSVNetworkCoreCluster
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterValidatorAdded is a free log retrieval operation binding the contract event 0x48a3ea0796746043948f6341d17ff8200937b99262a0b48c2663b951ed7114e5.
//
// Solidity: event ValidatorAdded(address indexed owner, uint64[] operatorIds, bytes publicKey, bytes shares, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SSVNetwork *SSVNetworkFilterer) FilterValidatorAdded(opts *bind.FilterOpts, owner []common.Address) (*SSVNetworkValidatorAddedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SSVNetwork.contract.FilterLogs(opts, "ValidatorAdded", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SSVNetworkValidatorAddedIterator{contract: _SSVNetwork.contract, event: "ValidatorAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorAdded is a free log subscription operation binding the contract event 0x48a3ea0796746043948f6341d17ff8200937b99262a0b48c2663b951ed7114e5.
//
// Solidity: event ValidatorAdded(address indexed owner, uint64[] operatorIds, bytes publicKey, bytes shares, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SSVNetwork *SSVNetworkFilterer) WatchValidatorAdded(opts *bind.WatchOpts, sink chan<- *SSVNetworkValidatorAdded, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SSVNetwork.contract.WatchLogs(opts, "ValidatorAdded", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SSVNetworkValidatorAdded)
				if err := _SSVNetwork.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
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

// ParseValidatorAdded is a log parse operation binding the contract event 0x48a3ea0796746043948f6341d17ff8200937b99262a0b48c2663b951ed7114e5.
//
// Solidity: event ValidatorAdded(address indexed owner, uint64[] operatorIds, bytes publicKey, bytes shares, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SSVNetwork *SSVNetworkFilterer) ParseValidatorAdded(log types.Log) (*SSVNetworkValidatorAdded, error) {
	event := new(SSVNetworkValidatorAdded)
	if err := _SSVNetwork.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SSVNetworkValidatorRemovedIterator is returned from FilterValidatorRemoved and is used to iterate over the raw logs and unpacked data for ValidatorRemoved events raised by the SSVNetwork contract.
type SSVNetworkValidatorRemovedIterator struct {
	Event *SSVNetworkValidatorRemoved // Event containing the contract specifics and raw log

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
func (it *SSVNetworkValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SSVNetworkValidatorRemoved)
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
		it.Event = new(SSVNetworkValidatorRemoved)
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
func (it *SSVNetworkValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SSVNetworkValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SSVNetworkValidatorRemoved represents a ValidatorRemoved event raised by the SSVNetwork contract.
type SSVNetworkValidatorRemoved struct {
	Owner       common.Address
	OperatorIds []uint64
	PublicKey   []byte
	Cluster     ISSVNetworkCoreCluster
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemoved is a free log retrieval operation binding the contract event 0xccf4370403e5fbbde0cd3f13426479dcd8a5916b05db424b7a2c04978cf8ce6e.
//
// Solidity: event ValidatorRemoved(address indexed owner, uint64[] operatorIds, bytes publicKey, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SSVNetwork *SSVNetworkFilterer) FilterValidatorRemoved(opts *bind.FilterOpts, owner []common.Address) (*SSVNetworkValidatorRemovedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SSVNetwork.contract.FilterLogs(opts, "ValidatorRemoved", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SSVNetworkValidatorRemovedIterator{contract: _SSVNetwork.contract, event: "ValidatorRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorRemoved is a free log subscription operation binding the contract event 0xccf4370403e5fbbde0cd3f13426479dcd8a5916b05db424b7a2c04978cf8ce6e.
//
// Solidity: event ValidatorRemoved(address indexed owner, uint64[] operatorIds, bytes publicKey, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SSVNetwork *SSVNetworkFilterer) WatchValidatorRemoved(opts *bind.WatchOpts, sink chan<- *SSVNetworkValidatorRemoved, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SSVNetwork.contract.WatchLogs(opts, "ValidatorRemoved", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SSVNetworkValidatorRemoved)
				if err := _SSVNetwork.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
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

// ParseValidatorRemoved is a log parse operation binding the contract event 0xccf4370403e5fbbde0cd3f13426479dcd8a5916b05db424b7a2c04978cf8ce6e.
//
// Solidity: event ValidatorRemoved(address indexed owner, uint64[] operatorIds, bytes publicKey, (uint32,uint64,uint64,bool,uint256) cluster)
func (_SSVNetwork *SSVNetworkFilterer) ParseValidatorRemoved(log types.Log) (*SSVNetworkValidatorRemoved, error) {
	event := new(SSVNetworkValidatorRemoved)
	if err := _SSVNetwork.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
