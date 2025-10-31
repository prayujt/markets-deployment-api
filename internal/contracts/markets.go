// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// AncillaryDataUpdate is an auto generated low-level Go binding around an user-defined struct.
type AncillaryDataUpdate struct {
	Timestamp *big.Int
	Update    []byte
}

// QuestionData is an auto generated low-level Go binding around an user-defined struct.
type QuestionData struct {
	RequestTimestamp          *big.Int
	Reward                    *big.Int
	ProposalBond              *big.Int
	Liveness                  *big.Int
	ManualResolutionTimestamp *big.Int
	Resolved                  bool
	Paused                    bool
	Reset                     bool
	Refund                    bool
	RewardToken               common.Address
	Creator                   common.Address
	AncillaryData             []byte
}

// MarketMetaData contains all meta data concerning the Market contract.
var MarketMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_ctf\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_finder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_oo\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MAX_ANCILLARY_DATA\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SAFETY_PERIOD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"YES_OR_NO_IDENTIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addAdmin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"admins\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collateralWhitelist\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAddressWhitelist\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ctf\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIConditionalTokens\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"flag\",\"inputs\":[{\"name\":\"questionID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getExpectedPayouts\",\"inputs\":[{\"name\":\"questionID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestUpdate\",\"inputs\":[{\"name\":\"questionID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structAncillaryDataUpdate\",\"components\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"update\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuestion\",\"inputs\":[{\"name\":\"questionID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structQuestionData\",\"components\":[{\"name\":\"requestTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reward\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposalBond\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"liveness\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"manualResolutionTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resolved\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"paused\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"reset\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"refund\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"rewardToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"ancillaryData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUpdates\",\"inputs\":[{\"name\":\"questionID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structAncillaryDataUpdate[]\",\"components\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"update\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"ancillaryData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rewardToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reward\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposalBond\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"liveness\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"questionID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isAdmin\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isFlagged\",\"inputs\":[{\"name\":\"questionID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isInitialized\",\"inputs\":[{\"name\":\"questionID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"optimisticOracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOptimisticOracleV2\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"questionID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"postUpdate\",\"inputs\":[{\"name\":\"questionID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"update\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"priceDisputed\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ancillaryData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"questions\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"requestTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reward\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposalBond\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"liveness\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"manualResolutionTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resolved\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"paused\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"reset\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"refund\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"rewardToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"ancillaryData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ready\",\"inputs\":[{\"name\":\"questionID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeAdmin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceAdmin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reset\",\"inputs\":[{\"name\":\"questionID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resolve\",\"inputs\":[{\"name\":\"questionID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resolveManually\",\"inputs\":[{\"name\":\"questionID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"payouts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unflag\",\"inputs\":[{\"name\":\"questionID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"questionID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updates\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"update\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AncillaryDataUpdated\",\"inputs\":[{\"name\":\"questionID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"update\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewAdmin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAdminAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuestionFlagged\",\"inputs\":[{\"name\":\"questionID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuestionInitialized\",\"inputs\":[{\"name\":\"questionID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"requestTimestamp\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"ancillaryData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"rewardToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"reward\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"proposalBond\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuestionManuallyResolved\",\"inputs\":[{\"name\":\"questionID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"payouts\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuestionPaused\",\"inputs\":[{\"name\":\"questionID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuestionReset\",\"inputs\":[{\"name\":\"questionID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuestionResolved\",\"inputs\":[{\"name\":\"questionID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"settledPrice\",\"type\":\"int256\",\"indexed\":true,\"internalType\":\"int256\"},{\"name\":\"payouts\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuestionUnflagged\",\"inputs\":[{\"name\":\"questionID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuestionUnpaused\",\"inputs\":[{\"name\":\"questionID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemovedAdmin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"removedAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"Flagged\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Initialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAncillaryData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOOPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPayouts\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotFlagged\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOptimisticOracle\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotReadyToResolve\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Paused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PriceNotAvailable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Resolved\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafetyPeriodNotPassed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafetyPeriodPassed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[]}]",
}

// MarketABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketMetaData.ABI instead.
var MarketABI = MarketMetaData.ABI

// Market is an auto generated Go binding around an Ethereum contract.
type Market struct {
	MarketCaller     // Read-only binding to the contract
	MarketTransactor // Write-only binding to the contract
	MarketFilterer   // Log filterer for contract events
}

// MarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketSession struct {
	Contract     *Market           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketCallerSession struct {
	Contract *MarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketTransactorSession struct {
	Contract     *MarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketRaw struct {
	Contract *Market // Generic contract binding to access the raw methods on
}

// MarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketCallerRaw struct {
	Contract *MarketCaller // Generic read-only contract binding to access the raw methods on
}

// MarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketTransactorRaw struct {
	Contract *MarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarket creates a new instance of Market, bound to a specific deployed contract.
func NewMarket(address common.Address, backend bind.ContractBackend) (*Market, error) {
	contract, err := bindMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Market{MarketCaller: MarketCaller{contract: contract}, MarketTransactor: MarketTransactor{contract: contract}, MarketFilterer: MarketFilterer{contract: contract}}, nil
}

// NewMarketCaller creates a new read-only instance of Market, bound to a specific deployed contract.
func NewMarketCaller(address common.Address, caller bind.ContractCaller) (*MarketCaller, error) {
	contract, err := bindMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketCaller{contract: contract}, nil
}

// NewMarketTransactor creates a new write-only instance of Market, bound to a specific deployed contract.
func NewMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketTransactor, error) {
	contract, err := bindMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketTransactor{contract: contract}, nil
}

// NewMarketFilterer creates a new log filterer instance of Market, bound to a specific deployed contract.
func NewMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketFilterer, error) {
	contract, err := bindMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketFilterer{contract: contract}, nil
}

// bindMarket binds a generic wrapper to an already deployed contract.
func bindMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MarketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Market *MarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Market.Contract.MarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Market *MarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.Contract.MarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Market *MarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Market.Contract.MarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Market *MarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Market.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Market *MarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Market *MarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Market.Contract.contract.Transact(opts, method, params...)
}

// MAXANCILLARYDATA is a free data retrieval call binding the contract method 0x27f8feac.
//
// Solidity: function MAX_ANCILLARY_DATA() view returns(uint256)
func (_Market *MarketCaller) MAXANCILLARYDATA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "MAX_ANCILLARY_DATA")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXANCILLARYDATA is a free data retrieval call binding the contract method 0x27f8feac.
//
// Solidity: function MAX_ANCILLARY_DATA() view returns(uint256)
func (_Market *MarketSession) MAXANCILLARYDATA() (*big.Int, error) {
	return _Market.Contract.MAXANCILLARYDATA(&_Market.CallOpts)
}

// MAXANCILLARYDATA is a free data retrieval call binding the contract method 0x27f8feac.
//
// Solidity: function MAX_ANCILLARY_DATA() view returns(uint256)
func (_Market *MarketCallerSession) MAXANCILLARYDATA() (*big.Int, error) {
	return _Market.Contract.MAXANCILLARYDATA(&_Market.CallOpts)
}

// SAFETYPERIOD is a free data retrieval call binding the contract method 0xd1dfb2e9.
//
// Solidity: function SAFETY_PERIOD() view returns(uint256)
func (_Market *MarketCaller) SAFETYPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "SAFETY_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SAFETYPERIOD is a free data retrieval call binding the contract method 0xd1dfb2e9.
//
// Solidity: function SAFETY_PERIOD() view returns(uint256)
func (_Market *MarketSession) SAFETYPERIOD() (*big.Int, error) {
	return _Market.Contract.SAFETYPERIOD(&_Market.CallOpts)
}

// SAFETYPERIOD is a free data retrieval call binding the contract method 0xd1dfb2e9.
//
// Solidity: function SAFETY_PERIOD() view returns(uint256)
func (_Market *MarketCallerSession) SAFETYPERIOD() (*big.Int, error) {
	return _Market.Contract.SAFETYPERIOD(&_Market.CallOpts)
}

// YESORNOIDENTIFIER is a free data retrieval call binding the contract method 0x6c66f07d.
//
// Solidity: function YES_OR_NO_IDENTIFIER() view returns(bytes32)
func (_Market *MarketCaller) YESORNOIDENTIFIER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "YES_OR_NO_IDENTIFIER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// YESORNOIDENTIFIER is a free data retrieval call binding the contract method 0x6c66f07d.
//
// Solidity: function YES_OR_NO_IDENTIFIER() view returns(bytes32)
func (_Market *MarketSession) YESORNOIDENTIFIER() ([32]byte, error) {
	return _Market.Contract.YESORNOIDENTIFIER(&_Market.CallOpts)
}

// YESORNOIDENTIFIER is a free data retrieval call binding the contract method 0x6c66f07d.
//
// Solidity: function YES_OR_NO_IDENTIFIER() view returns(bytes32)
func (_Market *MarketCallerSession) YESORNOIDENTIFIER() ([32]byte, error) {
	return _Market.Contract.YESORNOIDENTIFIER(&_Market.CallOpts)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_Market *MarketCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_Market *MarketSession) Admins(arg0 common.Address) (*big.Int, error) {
	return _Market.Contract.Admins(&_Market.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_Market *MarketCallerSession) Admins(arg0 common.Address) (*big.Int, error) {
	return _Market.Contract.Admins(&_Market.CallOpts, arg0)
}

// CollateralWhitelist is a free data retrieval call binding the contract method 0xe4ee614a.
//
// Solidity: function collateralWhitelist() view returns(address)
func (_Market *MarketCaller) CollateralWhitelist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "collateralWhitelist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralWhitelist is a free data retrieval call binding the contract method 0xe4ee614a.
//
// Solidity: function collateralWhitelist() view returns(address)
func (_Market *MarketSession) CollateralWhitelist() (common.Address, error) {
	return _Market.Contract.CollateralWhitelist(&_Market.CallOpts)
}

// CollateralWhitelist is a free data retrieval call binding the contract method 0xe4ee614a.
//
// Solidity: function collateralWhitelist() view returns(address)
func (_Market *MarketCallerSession) CollateralWhitelist() (common.Address, error) {
	return _Market.Contract.CollateralWhitelist(&_Market.CallOpts)
}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_Market *MarketCaller) Ctf(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "ctf")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_Market *MarketSession) Ctf() (common.Address, error) {
	return _Market.Contract.Ctf(&_Market.CallOpts)
}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_Market *MarketCallerSession) Ctf() (common.Address, error) {
	return _Market.Contract.Ctf(&_Market.CallOpts)
}

// GetExpectedPayouts is a free data retrieval call binding the contract method 0x34e5e28e.
//
// Solidity: function getExpectedPayouts(bytes32 questionID) view returns(uint256[])
func (_Market *MarketCaller) GetExpectedPayouts(opts *bind.CallOpts, questionID [32]byte) ([]*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getExpectedPayouts", questionID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetExpectedPayouts is a free data retrieval call binding the contract method 0x34e5e28e.
//
// Solidity: function getExpectedPayouts(bytes32 questionID) view returns(uint256[])
func (_Market *MarketSession) GetExpectedPayouts(questionID [32]byte) ([]*big.Int, error) {
	return _Market.Contract.GetExpectedPayouts(&_Market.CallOpts, questionID)
}

// GetExpectedPayouts is a free data retrieval call binding the contract method 0x34e5e28e.
//
// Solidity: function getExpectedPayouts(bytes32 questionID) view returns(uint256[])
func (_Market *MarketCallerSession) GetExpectedPayouts(questionID [32]byte) ([]*big.Int, error) {
	return _Market.Contract.GetExpectedPayouts(&_Market.CallOpts, questionID)
}

// GetLatestUpdate is a free data retrieval call binding the contract method 0xc0cab0a2.
//
// Solidity: function getLatestUpdate(bytes32 questionID, address owner) view returns((uint256,bytes))
func (_Market *MarketCaller) GetLatestUpdate(opts *bind.CallOpts, questionID [32]byte, owner common.Address) (AncillaryDataUpdate, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getLatestUpdate", questionID, owner)

	if err != nil {
		return *new(AncillaryDataUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(AncillaryDataUpdate)).(*AncillaryDataUpdate)

	return out0, err

}

// GetLatestUpdate is a free data retrieval call binding the contract method 0xc0cab0a2.
//
// Solidity: function getLatestUpdate(bytes32 questionID, address owner) view returns((uint256,bytes))
func (_Market *MarketSession) GetLatestUpdate(questionID [32]byte, owner common.Address) (AncillaryDataUpdate, error) {
	return _Market.Contract.GetLatestUpdate(&_Market.CallOpts, questionID, owner)
}

// GetLatestUpdate is a free data retrieval call binding the contract method 0xc0cab0a2.
//
// Solidity: function getLatestUpdate(bytes32 questionID, address owner) view returns((uint256,bytes))
func (_Market *MarketCallerSession) GetLatestUpdate(questionID [32]byte, owner common.Address) (AncillaryDataUpdate, error) {
	return _Market.Contract.GetLatestUpdate(&_Market.CallOpts, questionID, owner)
}

// GetQuestion is a free data retrieval call binding the contract method 0x58c039cd.
//
// Solidity: function getQuestion(bytes32 questionID) view returns((uint256,uint256,uint256,uint256,uint256,bool,bool,bool,bool,address,address,bytes))
func (_Market *MarketCaller) GetQuestion(opts *bind.CallOpts, questionID [32]byte) (QuestionData, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getQuestion", questionID)

	if err != nil {
		return *new(QuestionData), err
	}

	out0 := *abi.ConvertType(out[0], new(QuestionData)).(*QuestionData)

	return out0, err

}

// GetQuestion is a free data retrieval call binding the contract method 0x58c039cd.
//
// Solidity: function getQuestion(bytes32 questionID) view returns((uint256,uint256,uint256,uint256,uint256,bool,bool,bool,bool,address,address,bytes))
func (_Market *MarketSession) GetQuestion(questionID [32]byte) (QuestionData, error) {
	return _Market.Contract.GetQuestion(&_Market.CallOpts, questionID)
}

// GetQuestion is a free data retrieval call binding the contract method 0x58c039cd.
//
// Solidity: function getQuestion(bytes32 questionID) view returns((uint256,uint256,uint256,uint256,uint256,bool,bool,bool,bool,address,address,bytes))
func (_Market *MarketCallerSession) GetQuestion(questionID [32]byte) (QuestionData, error) {
	return _Market.Contract.GetQuestion(&_Market.CallOpts, questionID)
}

// GetUpdates is a free data retrieval call binding the contract method 0x555c56fc.
//
// Solidity: function getUpdates(bytes32 questionID, address owner) view returns((uint256,bytes)[])
func (_Market *MarketCaller) GetUpdates(opts *bind.CallOpts, questionID [32]byte, owner common.Address) ([]AncillaryDataUpdate, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getUpdates", questionID, owner)

	if err != nil {
		return *new([]AncillaryDataUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new([]AncillaryDataUpdate)).(*[]AncillaryDataUpdate)

	return out0, err

}

// GetUpdates is a free data retrieval call binding the contract method 0x555c56fc.
//
// Solidity: function getUpdates(bytes32 questionID, address owner) view returns((uint256,bytes)[])
func (_Market *MarketSession) GetUpdates(questionID [32]byte, owner common.Address) ([]AncillaryDataUpdate, error) {
	return _Market.Contract.GetUpdates(&_Market.CallOpts, questionID, owner)
}

// GetUpdates is a free data retrieval call binding the contract method 0x555c56fc.
//
// Solidity: function getUpdates(bytes32 questionID, address owner) view returns((uint256,bytes)[])
func (_Market *MarketCallerSession) GetUpdates(questionID [32]byte, owner common.Address) ([]AncillaryDataUpdate, error) {
	return _Market.Contract.GetUpdates(&_Market.CallOpts, questionID, owner)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address addr) view returns(bool)
func (_Market *MarketCaller) IsAdmin(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "isAdmin", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address addr) view returns(bool)
func (_Market *MarketSession) IsAdmin(addr common.Address) (bool, error) {
	return _Market.Contract.IsAdmin(&_Market.CallOpts, addr)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address addr) view returns(bool)
func (_Market *MarketCallerSession) IsAdmin(addr common.Address) (bool, error) {
	return _Market.Contract.IsAdmin(&_Market.CallOpts, addr)
}

// IsFlagged is a free data retrieval call binding the contract method 0xbf2dde38.
//
// Solidity: function isFlagged(bytes32 questionID) view returns(bool)
func (_Market *MarketCaller) IsFlagged(opts *bind.CallOpts, questionID [32]byte) (bool, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "isFlagged", questionID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFlagged is a free data retrieval call binding the contract method 0xbf2dde38.
//
// Solidity: function isFlagged(bytes32 questionID) view returns(bool)
func (_Market *MarketSession) IsFlagged(questionID [32]byte) (bool, error) {
	return _Market.Contract.IsFlagged(&_Market.CallOpts, questionID)
}

// IsFlagged is a free data retrieval call binding the contract method 0xbf2dde38.
//
// Solidity: function isFlagged(bytes32 questionID) view returns(bool)
func (_Market *MarketCallerSession) IsFlagged(questionID [32]byte) (bool, error) {
	return _Market.Contract.IsFlagged(&_Market.CallOpts, questionID)
}

// IsInitialized is a free data retrieval call binding the contract method 0xf7b637bb.
//
// Solidity: function isInitialized(bytes32 questionID) view returns(bool)
func (_Market *MarketCaller) IsInitialized(opts *bind.CallOpts, questionID [32]byte) (bool, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "isInitialized", questionID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0xf7b637bb.
//
// Solidity: function isInitialized(bytes32 questionID) view returns(bool)
func (_Market *MarketSession) IsInitialized(questionID [32]byte) (bool, error) {
	return _Market.Contract.IsInitialized(&_Market.CallOpts, questionID)
}

// IsInitialized is a free data retrieval call binding the contract method 0xf7b637bb.
//
// Solidity: function isInitialized(bytes32 questionID) view returns(bool)
func (_Market *MarketCallerSession) IsInitialized(questionID [32]byte) (bool, error) {
	return _Market.Contract.IsInitialized(&_Market.CallOpts, questionID)
}

// OptimisticOracle is a free data retrieval call binding the contract method 0x22302922.
//
// Solidity: function optimisticOracle() view returns(address)
func (_Market *MarketCaller) OptimisticOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "optimisticOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OptimisticOracle is a free data retrieval call binding the contract method 0x22302922.
//
// Solidity: function optimisticOracle() view returns(address)
func (_Market *MarketSession) OptimisticOracle() (common.Address, error) {
	return _Market.Contract.OptimisticOracle(&_Market.CallOpts)
}

// OptimisticOracle is a free data retrieval call binding the contract method 0x22302922.
//
// Solidity: function optimisticOracle() view returns(address)
func (_Market *MarketCallerSession) OptimisticOracle() (common.Address, error) {
	return _Market.Contract.OptimisticOracle(&_Market.CallOpts)
}

// Questions is a free data retrieval call binding the contract method 0x95addb90.
//
// Solidity: function questions(bytes32 ) view returns(uint256 requestTimestamp, uint256 reward, uint256 proposalBond, uint256 liveness, uint256 manualResolutionTimestamp, bool resolved, bool paused, bool reset, bool refund, address rewardToken, address creator, bytes ancillaryData)
func (_Market *MarketCaller) Questions(opts *bind.CallOpts, arg0 [32]byte) (struct {
	RequestTimestamp          *big.Int
	Reward                    *big.Int
	ProposalBond              *big.Int
	Liveness                  *big.Int
	ManualResolutionTimestamp *big.Int
	Resolved                  bool
	Paused                    bool
	Reset                     bool
	Refund                    bool
	RewardToken               common.Address
	Creator                   common.Address
	AncillaryData             []byte
}, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "questions", arg0)

	outstruct := new(struct {
		RequestTimestamp          *big.Int
		Reward                    *big.Int
		ProposalBond              *big.Int
		Liveness                  *big.Int
		ManualResolutionTimestamp *big.Int
		Resolved                  bool
		Paused                    bool
		Reset                     bool
		Refund                    bool
		RewardToken               common.Address
		Creator                   common.Address
		AncillaryData             []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequestTimestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ProposalBond = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Liveness = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ManualResolutionTimestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Resolved = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.Paused = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.Reset = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.Refund = *abi.ConvertType(out[8], new(bool)).(*bool)
	outstruct.RewardToken = *abi.ConvertType(out[9], new(common.Address)).(*common.Address)
	outstruct.Creator = *abi.ConvertType(out[10], new(common.Address)).(*common.Address)
	outstruct.AncillaryData = *abi.ConvertType(out[11], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Questions is a free data retrieval call binding the contract method 0x95addb90.
//
// Solidity: function questions(bytes32 ) view returns(uint256 requestTimestamp, uint256 reward, uint256 proposalBond, uint256 liveness, uint256 manualResolutionTimestamp, bool resolved, bool paused, bool reset, bool refund, address rewardToken, address creator, bytes ancillaryData)
func (_Market *MarketSession) Questions(arg0 [32]byte) (struct {
	RequestTimestamp          *big.Int
	Reward                    *big.Int
	ProposalBond              *big.Int
	Liveness                  *big.Int
	ManualResolutionTimestamp *big.Int
	Resolved                  bool
	Paused                    bool
	Reset                     bool
	Refund                    bool
	RewardToken               common.Address
	Creator                   common.Address
	AncillaryData             []byte
}, error) {
	return _Market.Contract.Questions(&_Market.CallOpts, arg0)
}

// Questions is a free data retrieval call binding the contract method 0x95addb90.
//
// Solidity: function questions(bytes32 ) view returns(uint256 requestTimestamp, uint256 reward, uint256 proposalBond, uint256 liveness, uint256 manualResolutionTimestamp, bool resolved, bool paused, bool reset, bool refund, address rewardToken, address creator, bytes ancillaryData)
func (_Market *MarketCallerSession) Questions(arg0 [32]byte) (struct {
	RequestTimestamp          *big.Int
	Reward                    *big.Int
	ProposalBond              *big.Int
	Liveness                  *big.Int
	ManualResolutionTimestamp *big.Int
	Resolved                  bool
	Paused                    bool
	Reset                     bool
	Refund                    bool
	RewardToken               common.Address
	Creator                   common.Address
	AncillaryData             []byte
}, error) {
	return _Market.Contract.Questions(&_Market.CallOpts, arg0)
}

// Ready is a free data retrieval call binding the contract method 0xfcac49a2.
//
// Solidity: function ready(bytes32 questionID) view returns(bool)
func (_Market *MarketCaller) Ready(opts *bind.CallOpts, questionID [32]byte) (bool, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "ready", questionID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Ready is a free data retrieval call binding the contract method 0xfcac49a2.
//
// Solidity: function ready(bytes32 questionID) view returns(bool)
func (_Market *MarketSession) Ready(questionID [32]byte) (bool, error) {
	return _Market.Contract.Ready(&_Market.CallOpts, questionID)
}

// Ready is a free data retrieval call binding the contract method 0xfcac49a2.
//
// Solidity: function ready(bytes32 questionID) view returns(bool)
func (_Market *MarketCallerSession) Ready(questionID [32]byte) (bool, error) {
	return _Market.Contract.Ready(&_Market.CallOpts, questionID)
}

// Updates is a free data retrieval call binding the contract method 0x89ab0871.
//
// Solidity: function updates(bytes32 , uint256 ) view returns(uint256 timestamp, bytes update)
func (_Market *MarketCaller) Updates(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (struct {
	Timestamp *big.Int
	Update    []byte
}, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "updates", arg0, arg1)

	outstruct := new(struct {
		Timestamp *big.Int
		Update    []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Update = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Updates is a free data retrieval call binding the contract method 0x89ab0871.
//
// Solidity: function updates(bytes32 , uint256 ) view returns(uint256 timestamp, bytes update)
func (_Market *MarketSession) Updates(arg0 [32]byte, arg1 *big.Int) (struct {
	Timestamp *big.Int
	Update    []byte
}, error) {
	return _Market.Contract.Updates(&_Market.CallOpts, arg0, arg1)
}

// Updates is a free data retrieval call binding the contract method 0x89ab0871.
//
// Solidity: function updates(bytes32 , uint256 ) view returns(uint256 timestamp, bytes update)
func (_Market *MarketCallerSession) Updates(arg0 [32]byte, arg1 *big.Int) (struct {
	Timestamp *big.Int
	Update    []byte
}, error) {
	return _Market.Contract.Updates(&_Market.CallOpts, arg0, arg1)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_Market *MarketTransactor) AddAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "addAdmin", admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_Market *MarketSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _Market.Contract.AddAdmin(&_Market.TransactOpts, admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_Market *MarketTransactorSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _Market.Contract.AddAdmin(&_Market.TransactOpts, admin)
}

// Flag is a paid mutator transaction binding the contract method 0x78165a48.
//
// Solidity: function flag(bytes32 questionID) returns()
func (_Market *MarketTransactor) Flag(opts *bind.TransactOpts, questionID [32]byte) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "flag", questionID)
}

// Flag is a paid mutator transaction binding the contract method 0x78165a48.
//
// Solidity: function flag(bytes32 questionID) returns()
func (_Market *MarketSession) Flag(questionID [32]byte) (*types.Transaction, error) {
	return _Market.Contract.Flag(&_Market.TransactOpts, questionID)
}

// Flag is a paid mutator transaction binding the contract method 0x78165a48.
//
// Solidity: function flag(bytes32 questionID) returns()
func (_Market *MarketTransactorSession) Flag(questionID [32]byte) (*types.Transaction, error) {
	return _Market.Contract.Flag(&_Market.TransactOpts, questionID)
}

// Initialize is a paid mutator transaction binding the contract method 0x185d1646.
//
// Solidity: function initialize(bytes ancillaryData, address rewardToken, uint256 reward, uint256 proposalBond, uint256 liveness) returns(bytes32 questionID)
func (_Market *MarketTransactor) Initialize(opts *bind.TransactOpts, ancillaryData []byte, rewardToken common.Address, reward *big.Int, proposalBond *big.Int, liveness *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "initialize", ancillaryData, rewardToken, reward, proposalBond, liveness)
}

// Initialize is a paid mutator transaction binding the contract method 0x185d1646.
//
// Solidity: function initialize(bytes ancillaryData, address rewardToken, uint256 reward, uint256 proposalBond, uint256 liveness) returns(bytes32 questionID)
func (_Market *MarketSession) Initialize(ancillaryData []byte, rewardToken common.Address, reward *big.Int, proposalBond *big.Int, liveness *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Initialize(&_Market.TransactOpts, ancillaryData, rewardToken, reward, proposalBond, liveness)
}

// Initialize is a paid mutator transaction binding the contract method 0x185d1646.
//
// Solidity: function initialize(bytes ancillaryData, address rewardToken, uint256 reward, uint256 proposalBond, uint256 liveness) returns(bytes32 questionID)
func (_Market *MarketTransactorSession) Initialize(ancillaryData []byte, rewardToken common.Address, reward *big.Int, proposalBond *big.Int, liveness *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Initialize(&_Market.TransactOpts, ancillaryData, rewardToken, reward, proposalBond, liveness)
}

// Pause is a paid mutator transaction binding the contract method 0xed56531a.
//
// Solidity: function pause(bytes32 questionID) returns()
func (_Market *MarketTransactor) Pause(opts *bind.TransactOpts, questionID [32]byte) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "pause", questionID)
}

// Pause is a paid mutator transaction binding the contract method 0xed56531a.
//
// Solidity: function pause(bytes32 questionID) returns()
func (_Market *MarketSession) Pause(questionID [32]byte) (*types.Transaction, error) {
	return _Market.Contract.Pause(&_Market.TransactOpts, questionID)
}

// Pause is a paid mutator transaction binding the contract method 0xed56531a.
//
// Solidity: function pause(bytes32 questionID) returns()
func (_Market *MarketTransactorSession) Pause(questionID [32]byte) (*types.Transaction, error) {
	return _Market.Contract.Pause(&_Market.TransactOpts, questionID)
}

// PostUpdate is a paid mutator transaction binding the contract method 0x072d1259.
//
// Solidity: function postUpdate(bytes32 questionID, bytes update) returns()
func (_Market *MarketTransactor) PostUpdate(opts *bind.TransactOpts, questionID [32]byte, update []byte) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "postUpdate", questionID, update)
}

// PostUpdate is a paid mutator transaction binding the contract method 0x072d1259.
//
// Solidity: function postUpdate(bytes32 questionID, bytes update) returns()
func (_Market *MarketSession) PostUpdate(questionID [32]byte, update []byte) (*types.Transaction, error) {
	return _Market.Contract.PostUpdate(&_Market.TransactOpts, questionID, update)
}

// PostUpdate is a paid mutator transaction binding the contract method 0x072d1259.
//
// Solidity: function postUpdate(bytes32 questionID, bytes update) returns()
func (_Market *MarketTransactorSession) PostUpdate(questionID [32]byte, update []byte) (*types.Transaction, error) {
	return _Market.Contract.PostUpdate(&_Market.TransactOpts, questionID, update)
}

// PriceDisputed is a paid mutator transaction binding the contract method 0x0d8f2372.
//
// Solidity: function priceDisputed(bytes32 , uint256 , bytes ancillaryData, uint256 ) returns()
func (_Market *MarketTransactor) PriceDisputed(opts *bind.TransactOpts, arg0 [32]byte, arg1 *big.Int, ancillaryData []byte, arg3 *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "priceDisputed", arg0, arg1, ancillaryData, arg3)
}

// PriceDisputed is a paid mutator transaction binding the contract method 0x0d8f2372.
//
// Solidity: function priceDisputed(bytes32 , uint256 , bytes ancillaryData, uint256 ) returns()
func (_Market *MarketSession) PriceDisputed(arg0 [32]byte, arg1 *big.Int, ancillaryData []byte, arg3 *big.Int) (*types.Transaction, error) {
	return _Market.Contract.PriceDisputed(&_Market.TransactOpts, arg0, arg1, ancillaryData, arg3)
}

// PriceDisputed is a paid mutator transaction binding the contract method 0x0d8f2372.
//
// Solidity: function priceDisputed(bytes32 , uint256 , bytes ancillaryData, uint256 ) returns()
func (_Market *MarketTransactorSession) PriceDisputed(arg0 [32]byte, arg1 *big.Int, ancillaryData []byte, arg3 *big.Int) (*types.Transaction, error) {
	return _Market.Contract.PriceDisputed(&_Market.TransactOpts, arg0, arg1, ancillaryData, arg3)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_Market *MarketTransactor) RemoveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "removeAdmin", admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_Market *MarketSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _Market.Contract.RemoveAdmin(&_Market.TransactOpts, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_Market *MarketTransactorSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _Market.Contract.RemoveAdmin(&_Market.TransactOpts, admin)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_Market *MarketTransactor) RenounceAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "renounceAdmin")
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_Market *MarketSession) RenounceAdmin() (*types.Transaction, error) {
	return _Market.Contract.RenounceAdmin(&_Market.TransactOpts)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_Market *MarketTransactorSession) RenounceAdmin() (*types.Transaction, error) {
	return _Market.Contract.RenounceAdmin(&_Market.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0xed3c7d40.
//
// Solidity: function reset(bytes32 questionID) returns()
func (_Market *MarketTransactor) Reset(opts *bind.TransactOpts, questionID [32]byte) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "reset", questionID)
}

// Reset is a paid mutator transaction binding the contract method 0xed3c7d40.
//
// Solidity: function reset(bytes32 questionID) returns()
func (_Market *MarketSession) Reset(questionID [32]byte) (*types.Transaction, error) {
	return _Market.Contract.Reset(&_Market.TransactOpts, questionID)
}

// Reset is a paid mutator transaction binding the contract method 0xed3c7d40.
//
// Solidity: function reset(bytes32 questionID) returns()
func (_Market *MarketTransactorSession) Reset(questionID [32]byte) (*types.Transaction, error) {
	return _Market.Contract.Reset(&_Market.TransactOpts, questionID)
}

// Resolve is a paid mutator transaction binding the contract method 0x5c23bdf5.
//
// Solidity: function resolve(bytes32 questionID) returns()
func (_Market *MarketTransactor) Resolve(opts *bind.TransactOpts, questionID [32]byte) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "resolve", questionID)
}

// Resolve is a paid mutator transaction binding the contract method 0x5c23bdf5.
//
// Solidity: function resolve(bytes32 questionID) returns()
func (_Market *MarketSession) Resolve(questionID [32]byte) (*types.Transaction, error) {
	return _Market.Contract.Resolve(&_Market.TransactOpts, questionID)
}

// Resolve is a paid mutator transaction binding the contract method 0x5c23bdf5.
//
// Solidity: function resolve(bytes32 questionID) returns()
func (_Market *MarketTransactorSession) Resolve(questionID [32]byte) (*types.Transaction, error) {
	return _Market.Contract.Resolve(&_Market.TransactOpts, questionID)
}

// ResolveManually is a paid mutator transaction binding the contract method 0x80696d85.
//
// Solidity: function resolveManually(bytes32 questionID, uint256[] payouts) returns()
func (_Market *MarketTransactor) ResolveManually(opts *bind.TransactOpts, questionID [32]byte, payouts []*big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "resolveManually", questionID, payouts)
}

// ResolveManually is a paid mutator transaction binding the contract method 0x80696d85.
//
// Solidity: function resolveManually(bytes32 questionID, uint256[] payouts) returns()
func (_Market *MarketSession) ResolveManually(questionID [32]byte, payouts []*big.Int) (*types.Transaction, error) {
	return _Market.Contract.ResolveManually(&_Market.TransactOpts, questionID, payouts)
}

// ResolveManually is a paid mutator transaction binding the contract method 0x80696d85.
//
// Solidity: function resolveManually(bytes32 questionID, uint256[] payouts) returns()
func (_Market *MarketTransactorSession) ResolveManually(questionID [32]byte, payouts []*big.Int) (*types.Transaction, error) {
	return _Market.Contract.ResolveManually(&_Market.TransactOpts, questionID, payouts)
}

// Unflag is a paid mutator transaction binding the contract method 0x88697de4.
//
// Solidity: function unflag(bytes32 questionID) returns()
func (_Market *MarketTransactor) Unflag(opts *bind.TransactOpts, questionID [32]byte) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "unflag", questionID)
}

// Unflag is a paid mutator transaction binding the contract method 0x88697de4.
//
// Solidity: function unflag(bytes32 questionID) returns()
func (_Market *MarketSession) Unflag(questionID [32]byte) (*types.Transaction, error) {
	return _Market.Contract.Unflag(&_Market.TransactOpts, questionID)
}

// Unflag is a paid mutator transaction binding the contract method 0x88697de4.
//
// Solidity: function unflag(bytes32 questionID) returns()
func (_Market *MarketTransactorSession) Unflag(questionID [32]byte) (*types.Transaction, error) {
	return _Market.Contract.Unflag(&_Market.TransactOpts, questionID)
}

// Unpause is a paid mutator transaction binding the contract method 0x2f4dae9f.
//
// Solidity: function unpause(bytes32 questionID) returns()
func (_Market *MarketTransactor) Unpause(opts *bind.TransactOpts, questionID [32]byte) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "unpause", questionID)
}

// Unpause is a paid mutator transaction binding the contract method 0x2f4dae9f.
//
// Solidity: function unpause(bytes32 questionID) returns()
func (_Market *MarketSession) Unpause(questionID [32]byte) (*types.Transaction, error) {
	return _Market.Contract.Unpause(&_Market.TransactOpts, questionID)
}

// Unpause is a paid mutator transaction binding the contract method 0x2f4dae9f.
//
// Solidity: function unpause(bytes32 questionID) returns()
func (_Market *MarketTransactorSession) Unpause(questionID [32]byte) (*types.Transaction, error) {
	return _Market.Contract.Unpause(&_Market.TransactOpts, questionID)
}

// MarketAncillaryDataUpdatedIterator is returned from FilterAncillaryDataUpdated and is used to iterate over the raw logs and unpacked data for AncillaryDataUpdated events raised by the Market contract.
type MarketAncillaryDataUpdatedIterator struct {
	Event *MarketAncillaryDataUpdated // Event containing the contract specifics and raw log

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
func (it *MarketAncillaryDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketAncillaryDataUpdated)
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
		it.Event = new(MarketAncillaryDataUpdated)
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
func (it *MarketAncillaryDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketAncillaryDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketAncillaryDataUpdated represents a AncillaryDataUpdated event raised by the Market contract.
type MarketAncillaryDataUpdated struct {
	QuestionID [32]byte
	Owner      common.Address
	Update     []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAncillaryDataUpdated is a free log retrieval operation binding the contract event 0x0059e11815211969c0c4aaf3f498b52b6c2f2d14f286275d0862d70de22a836b.
//
// Solidity: event AncillaryDataUpdated(bytes32 indexed questionID, address indexed owner, bytes update)
func (_Market *MarketFilterer) FilterAncillaryDataUpdated(opts *bind.FilterOpts, questionID [][32]byte, owner []common.Address) (*MarketAncillaryDataUpdatedIterator, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "AncillaryDataUpdated", questionIDRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &MarketAncillaryDataUpdatedIterator{contract: _Market.contract, event: "AncillaryDataUpdated", logs: logs, sub: sub}, nil
}

// WatchAncillaryDataUpdated is a free log subscription operation binding the contract event 0x0059e11815211969c0c4aaf3f498b52b6c2f2d14f286275d0862d70de22a836b.
//
// Solidity: event AncillaryDataUpdated(bytes32 indexed questionID, address indexed owner, bytes update)
func (_Market *MarketFilterer) WatchAncillaryDataUpdated(opts *bind.WatchOpts, sink chan<- *MarketAncillaryDataUpdated, questionID [][32]byte, owner []common.Address) (event.Subscription, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "AncillaryDataUpdated", questionIDRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketAncillaryDataUpdated)
				if err := _Market.contract.UnpackLog(event, "AncillaryDataUpdated", log); err != nil {
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

// ParseAncillaryDataUpdated is a log parse operation binding the contract event 0x0059e11815211969c0c4aaf3f498b52b6c2f2d14f286275d0862d70de22a836b.
//
// Solidity: event AncillaryDataUpdated(bytes32 indexed questionID, address indexed owner, bytes update)
func (_Market *MarketFilterer) ParseAncillaryDataUpdated(log types.Log) (*MarketAncillaryDataUpdated, error) {
	event := new(MarketAncillaryDataUpdated)
	if err := _Market.contract.UnpackLog(event, "AncillaryDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the Market contract.
type MarketNewAdminIterator struct {
	Event *MarketNewAdmin // Event containing the contract specifics and raw log

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
func (it *MarketNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketNewAdmin)
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
		it.Event = new(MarketNewAdmin)
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
func (it *MarketNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketNewAdmin represents a NewAdmin event raised by the Market contract.
type MarketNewAdmin struct {
	Admin           common.Address
	NewAdminAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed admin, address indexed newAdminAddress)
func (_Market *MarketFilterer) FilterNewAdmin(opts *bind.FilterOpts, admin []common.Address, newAdminAddress []common.Address) (*MarketNewAdminIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var newAdminAddressRule []interface{}
	for _, newAdminAddressItem := range newAdminAddress {
		newAdminAddressRule = append(newAdminAddressRule, newAdminAddressItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "NewAdmin", adminRule, newAdminAddressRule)
	if err != nil {
		return nil, err
	}
	return &MarketNewAdminIterator{contract: _Market.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed admin, address indexed newAdminAddress)
func (_Market *MarketFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *MarketNewAdmin, admin []common.Address, newAdminAddress []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var newAdminAddressRule []interface{}
	for _, newAdminAddressItem := range newAdminAddress {
		newAdminAddressRule = append(newAdminAddressRule, newAdminAddressItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "NewAdmin", adminRule, newAdminAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketNewAdmin)
				if err := _Market.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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

// ParseNewAdmin is a log parse operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed admin, address indexed newAdminAddress)
func (_Market *MarketFilterer) ParseNewAdmin(log types.Log) (*MarketNewAdmin, error) {
	event := new(MarketNewAdmin)
	if err := _Market.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketQuestionFlaggedIterator is returned from FilterQuestionFlagged and is used to iterate over the raw logs and unpacked data for QuestionFlagged events raised by the Market contract.
type MarketQuestionFlaggedIterator struct {
	Event *MarketQuestionFlagged // Event containing the contract specifics and raw log

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
func (it *MarketQuestionFlaggedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketQuestionFlagged)
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
		it.Event = new(MarketQuestionFlagged)
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
func (it *MarketQuestionFlaggedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketQuestionFlaggedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketQuestionFlagged represents a QuestionFlagged event raised by the Market contract.
type MarketQuestionFlagged struct {
	QuestionID [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterQuestionFlagged is a free log retrieval operation binding the contract event 0x2435a0347185933b12027c6f394a5fd9c03646dba233e956f50658719dfc0b35.
//
// Solidity: event QuestionFlagged(bytes32 indexed questionID)
func (_Market *MarketFilterer) FilterQuestionFlagged(opts *bind.FilterOpts, questionID [][32]byte) (*MarketQuestionFlaggedIterator, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "QuestionFlagged", questionIDRule)
	if err != nil {
		return nil, err
	}
	return &MarketQuestionFlaggedIterator{contract: _Market.contract, event: "QuestionFlagged", logs: logs, sub: sub}, nil
}

// WatchQuestionFlagged is a free log subscription operation binding the contract event 0x2435a0347185933b12027c6f394a5fd9c03646dba233e956f50658719dfc0b35.
//
// Solidity: event QuestionFlagged(bytes32 indexed questionID)
func (_Market *MarketFilterer) WatchQuestionFlagged(opts *bind.WatchOpts, sink chan<- *MarketQuestionFlagged, questionID [][32]byte) (event.Subscription, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "QuestionFlagged", questionIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketQuestionFlagged)
				if err := _Market.contract.UnpackLog(event, "QuestionFlagged", log); err != nil {
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

// ParseQuestionFlagged is a log parse operation binding the contract event 0x2435a0347185933b12027c6f394a5fd9c03646dba233e956f50658719dfc0b35.
//
// Solidity: event QuestionFlagged(bytes32 indexed questionID)
func (_Market *MarketFilterer) ParseQuestionFlagged(log types.Log) (*MarketQuestionFlagged, error) {
	event := new(MarketQuestionFlagged)
	if err := _Market.contract.UnpackLog(event, "QuestionFlagged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketQuestionInitializedIterator is returned from FilterQuestionInitialized and is used to iterate over the raw logs and unpacked data for QuestionInitialized events raised by the Market contract.
type MarketQuestionInitializedIterator struct {
	Event *MarketQuestionInitialized // Event containing the contract specifics and raw log

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
func (it *MarketQuestionInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketQuestionInitialized)
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
		it.Event = new(MarketQuestionInitialized)
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
func (it *MarketQuestionInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketQuestionInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketQuestionInitialized represents a QuestionInitialized event raised by the Market contract.
type MarketQuestionInitialized struct {
	QuestionID       [32]byte
	RequestTimestamp *big.Int
	Creator          common.Address
	AncillaryData    []byte
	RewardToken      common.Address
	Reward           *big.Int
	ProposalBond     *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterQuestionInitialized is a free log retrieval operation binding the contract event 0xeee0897acd6893adcaf2ba5158191b3601098ab6bece35c5d57874340b64c5b7.
//
// Solidity: event QuestionInitialized(bytes32 indexed questionID, uint256 indexed requestTimestamp, address indexed creator, bytes ancillaryData, address rewardToken, uint256 reward, uint256 proposalBond)
func (_Market *MarketFilterer) FilterQuestionInitialized(opts *bind.FilterOpts, questionID [][32]byte, requestTimestamp []*big.Int, creator []common.Address) (*MarketQuestionInitializedIterator, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}
	var requestTimestampRule []interface{}
	for _, requestTimestampItem := range requestTimestamp {
		requestTimestampRule = append(requestTimestampRule, requestTimestampItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "QuestionInitialized", questionIDRule, requestTimestampRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &MarketQuestionInitializedIterator{contract: _Market.contract, event: "QuestionInitialized", logs: logs, sub: sub}, nil
}

// WatchQuestionInitialized is a free log subscription operation binding the contract event 0xeee0897acd6893adcaf2ba5158191b3601098ab6bece35c5d57874340b64c5b7.
//
// Solidity: event QuestionInitialized(bytes32 indexed questionID, uint256 indexed requestTimestamp, address indexed creator, bytes ancillaryData, address rewardToken, uint256 reward, uint256 proposalBond)
func (_Market *MarketFilterer) WatchQuestionInitialized(opts *bind.WatchOpts, sink chan<- *MarketQuestionInitialized, questionID [][32]byte, requestTimestamp []*big.Int, creator []common.Address) (event.Subscription, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}
	var requestTimestampRule []interface{}
	for _, requestTimestampItem := range requestTimestamp {
		requestTimestampRule = append(requestTimestampRule, requestTimestampItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "QuestionInitialized", questionIDRule, requestTimestampRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketQuestionInitialized)
				if err := _Market.contract.UnpackLog(event, "QuestionInitialized", log); err != nil {
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

// ParseQuestionInitialized is a log parse operation binding the contract event 0xeee0897acd6893adcaf2ba5158191b3601098ab6bece35c5d57874340b64c5b7.
//
// Solidity: event QuestionInitialized(bytes32 indexed questionID, uint256 indexed requestTimestamp, address indexed creator, bytes ancillaryData, address rewardToken, uint256 reward, uint256 proposalBond)
func (_Market *MarketFilterer) ParseQuestionInitialized(log types.Log) (*MarketQuestionInitialized, error) {
	event := new(MarketQuestionInitialized)
	if err := _Market.contract.UnpackLog(event, "QuestionInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketQuestionManuallyResolvedIterator is returned from FilterQuestionManuallyResolved and is used to iterate over the raw logs and unpacked data for QuestionManuallyResolved events raised by the Market contract.
type MarketQuestionManuallyResolvedIterator struct {
	Event *MarketQuestionManuallyResolved // Event containing the contract specifics and raw log

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
func (it *MarketQuestionManuallyResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketQuestionManuallyResolved)
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
		it.Event = new(MarketQuestionManuallyResolved)
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
func (it *MarketQuestionManuallyResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketQuestionManuallyResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketQuestionManuallyResolved represents a QuestionManuallyResolved event raised by the Market contract.
type MarketQuestionManuallyResolved struct {
	QuestionID [32]byte
	Payouts    []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterQuestionManuallyResolved is a free log retrieval operation binding the contract event 0x5909815fe7fe0a550d5fcb95fbf33821b580521d3c97089c6ce12808d1cd0566.
//
// Solidity: event QuestionManuallyResolved(bytes32 indexed questionID, uint256[] payouts)
func (_Market *MarketFilterer) FilterQuestionManuallyResolved(opts *bind.FilterOpts, questionID [][32]byte) (*MarketQuestionManuallyResolvedIterator, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "QuestionManuallyResolved", questionIDRule)
	if err != nil {
		return nil, err
	}
	return &MarketQuestionManuallyResolvedIterator{contract: _Market.contract, event: "QuestionManuallyResolved", logs: logs, sub: sub}, nil
}

// WatchQuestionManuallyResolved is a free log subscription operation binding the contract event 0x5909815fe7fe0a550d5fcb95fbf33821b580521d3c97089c6ce12808d1cd0566.
//
// Solidity: event QuestionManuallyResolved(bytes32 indexed questionID, uint256[] payouts)
func (_Market *MarketFilterer) WatchQuestionManuallyResolved(opts *bind.WatchOpts, sink chan<- *MarketQuestionManuallyResolved, questionID [][32]byte) (event.Subscription, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "QuestionManuallyResolved", questionIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketQuestionManuallyResolved)
				if err := _Market.contract.UnpackLog(event, "QuestionManuallyResolved", log); err != nil {
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

// ParseQuestionManuallyResolved is a log parse operation binding the contract event 0x5909815fe7fe0a550d5fcb95fbf33821b580521d3c97089c6ce12808d1cd0566.
//
// Solidity: event QuestionManuallyResolved(bytes32 indexed questionID, uint256[] payouts)
func (_Market *MarketFilterer) ParseQuestionManuallyResolved(log types.Log) (*MarketQuestionManuallyResolved, error) {
	event := new(MarketQuestionManuallyResolved)
	if err := _Market.contract.UnpackLog(event, "QuestionManuallyResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketQuestionPausedIterator is returned from FilterQuestionPaused and is used to iterate over the raw logs and unpacked data for QuestionPaused events raised by the Market contract.
type MarketQuestionPausedIterator struct {
	Event *MarketQuestionPaused // Event containing the contract specifics and raw log

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
func (it *MarketQuestionPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketQuestionPaused)
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
		it.Event = new(MarketQuestionPaused)
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
func (it *MarketQuestionPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketQuestionPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketQuestionPaused represents a QuestionPaused event raised by the Market contract.
type MarketQuestionPaused struct {
	QuestionID [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterQuestionPaused is a free log retrieval operation binding the contract event 0x6ded7250a9d5f79aef5add44600fc20a74a0af6f4730baa4fc4ab87bf484b812.
//
// Solidity: event QuestionPaused(bytes32 indexed questionID)
func (_Market *MarketFilterer) FilterQuestionPaused(opts *bind.FilterOpts, questionID [][32]byte) (*MarketQuestionPausedIterator, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "QuestionPaused", questionIDRule)
	if err != nil {
		return nil, err
	}
	return &MarketQuestionPausedIterator{contract: _Market.contract, event: "QuestionPaused", logs: logs, sub: sub}, nil
}

// WatchQuestionPaused is a free log subscription operation binding the contract event 0x6ded7250a9d5f79aef5add44600fc20a74a0af6f4730baa4fc4ab87bf484b812.
//
// Solidity: event QuestionPaused(bytes32 indexed questionID)
func (_Market *MarketFilterer) WatchQuestionPaused(opts *bind.WatchOpts, sink chan<- *MarketQuestionPaused, questionID [][32]byte) (event.Subscription, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "QuestionPaused", questionIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketQuestionPaused)
				if err := _Market.contract.UnpackLog(event, "QuestionPaused", log); err != nil {
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

// ParseQuestionPaused is a log parse operation binding the contract event 0x6ded7250a9d5f79aef5add44600fc20a74a0af6f4730baa4fc4ab87bf484b812.
//
// Solidity: event QuestionPaused(bytes32 indexed questionID)
func (_Market *MarketFilterer) ParseQuestionPaused(log types.Log) (*MarketQuestionPaused, error) {
	event := new(MarketQuestionPaused)
	if err := _Market.contract.UnpackLog(event, "QuestionPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketQuestionResetIterator is returned from FilterQuestionReset and is used to iterate over the raw logs and unpacked data for QuestionReset events raised by the Market contract.
type MarketQuestionResetIterator struct {
	Event *MarketQuestionReset // Event containing the contract specifics and raw log

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
func (it *MarketQuestionResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketQuestionReset)
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
		it.Event = new(MarketQuestionReset)
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
func (it *MarketQuestionResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketQuestionResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketQuestionReset represents a QuestionReset event raised by the Market contract.
type MarketQuestionReset struct {
	QuestionID [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterQuestionReset is a free log retrieval operation binding the contract event 0x7981b5832932948db4e32a4a16a0f44b2ce7ff088574afb9364b313f70f82e8f.
//
// Solidity: event QuestionReset(bytes32 indexed questionID)
func (_Market *MarketFilterer) FilterQuestionReset(opts *bind.FilterOpts, questionID [][32]byte) (*MarketQuestionResetIterator, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "QuestionReset", questionIDRule)
	if err != nil {
		return nil, err
	}
	return &MarketQuestionResetIterator{contract: _Market.contract, event: "QuestionReset", logs: logs, sub: sub}, nil
}

// WatchQuestionReset is a free log subscription operation binding the contract event 0x7981b5832932948db4e32a4a16a0f44b2ce7ff088574afb9364b313f70f82e8f.
//
// Solidity: event QuestionReset(bytes32 indexed questionID)
func (_Market *MarketFilterer) WatchQuestionReset(opts *bind.WatchOpts, sink chan<- *MarketQuestionReset, questionID [][32]byte) (event.Subscription, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "QuestionReset", questionIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketQuestionReset)
				if err := _Market.contract.UnpackLog(event, "QuestionReset", log); err != nil {
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

// ParseQuestionReset is a log parse operation binding the contract event 0x7981b5832932948db4e32a4a16a0f44b2ce7ff088574afb9364b313f70f82e8f.
//
// Solidity: event QuestionReset(bytes32 indexed questionID)
func (_Market *MarketFilterer) ParseQuestionReset(log types.Log) (*MarketQuestionReset, error) {
	event := new(MarketQuestionReset)
	if err := _Market.contract.UnpackLog(event, "QuestionReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketQuestionResolvedIterator is returned from FilterQuestionResolved and is used to iterate over the raw logs and unpacked data for QuestionResolved events raised by the Market contract.
type MarketQuestionResolvedIterator struct {
	Event *MarketQuestionResolved // Event containing the contract specifics and raw log

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
func (it *MarketQuestionResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketQuestionResolved)
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
		it.Event = new(MarketQuestionResolved)
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
func (it *MarketQuestionResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketQuestionResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketQuestionResolved represents a QuestionResolved event raised by the Market contract.
type MarketQuestionResolved struct {
	QuestionID   [32]byte
	SettledPrice *big.Int
	Payouts      []*big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterQuestionResolved is a free log retrieval operation binding the contract event 0x566c3fbdd12dd86bb341787f6d531f79fd7ad4ce7e3ae2d15ac0ca1b601af9df.
//
// Solidity: event QuestionResolved(bytes32 indexed questionID, int256 indexed settledPrice, uint256[] payouts)
func (_Market *MarketFilterer) FilterQuestionResolved(opts *bind.FilterOpts, questionID [][32]byte, settledPrice []*big.Int) (*MarketQuestionResolvedIterator, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}
	var settledPriceRule []interface{}
	for _, settledPriceItem := range settledPrice {
		settledPriceRule = append(settledPriceRule, settledPriceItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "QuestionResolved", questionIDRule, settledPriceRule)
	if err != nil {
		return nil, err
	}
	return &MarketQuestionResolvedIterator{contract: _Market.contract, event: "QuestionResolved", logs: logs, sub: sub}, nil
}

// WatchQuestionResolved is a free log subscription operation binding the contract event 0x566c3fbdd12dd86bb341787f6d531f79fd7ad4ce7e3ae2d15ac0ca1b601af9df.
//
// Solidity: event QuestionResolved(bytes32 indexed questionID, int256 indexed settledPrice, uint256[] payouts)
func (_Market *MarketFilterer) WatchQuestionResolved(opts *bind.WatchOpts, sink chan<- *MarketQuestionResolved, questionID [][32]byte, settledPrice []*big.Int) (event.Subscription, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}
	var settledPriceRule []interface{}
	for _, settledPriceItem := range settledPrice {
		settledPriceRule = append(settledPriceRule, settledPriceItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "QuestionResolved", questionIDRule, settledPriceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketQuestionResolved)
				if err := _Market.contract.UnpackLog(event, "QuestionResolved", log); err != nil {
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

// ParseQuestionResolved is a log parse operation binding the contract event 0x566c3fbdd12dd86bb341787f6d531f79fd7ad4ce7e3ae2d15ac0ca1b601af9df.
//
// Solidity: event QuestionResolved(bytes32 indexed questionID, int256 indexed settledPrice, uint256[] payouts)
func (_Market *MarketFilterer) ParseQuestionResolved(log types.Log) (*MarketQuestionResolved, error) {
	event := new(MarketQuestionResolved)
	if err := _Market.contract.UnpackLog(event, "QuestionResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketQuestionUnflaggedIterator is returned from FilterQuestionUnflagged and is used to iterate over the raw logs and unpacked data for QuestionUnflagged events raised by the Market contract.
type MarketQuestionUnflaggedIterator struct {
	Event *MarketQuestionUnflagged // Event containing the contract specifics and raw log

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
func (it *MarketQuestionUnflaggedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketQuestionUnflagged)
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
		it.Event = new(MarketQuestionUnflagged)
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
func (it *MarketQuestionUnflaggedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketQuestionUnflaggedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketQuestionUnflagged represents a QuestionUnflagged event raised by the Market contract.
type MarketQuestionUnflagged struct {
	QuestionID [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterQuestionUnflagged is a free log retrieval operation binding the contract event 0x052435bc04fc49113a7bfd9198a92c0852ca622a621800f6da66d4b29b786c05.
//
// Solidity: event QuestionUnflagged(bytes32 indexed questionID)
func (_Market *MarketFilterer) FilterQuestionUnflagged(opts *bind.FilterOpts, questionID [][32]byte) (*MarketQuestionUnflaggedIterator, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "QuestionUnflagged", questionIDRule)
	if err != nil {
		return nil, err
	}
	return &MarketQuestionUnflaggedIterator{contract: _Market.contract, event: "QuestionUnflagged", logs: logs, sub: sub}, nil
}

// WatchQuestionUnflagged is a free log subscription operation binding the contract event 0x052435bc04fc49113a7bfd9198a92c0852ca622a621800f6da66d4b29b786c05.
//
// Solidity: event QuestionUnflagged(bytes32 indexed questionID)
func (_Market *MarketFilterer) WatchQuestionUnflagged(opts *bind.WatchOpts, sink chan<- *MarketQuestionUnflagged, questionID [][32]byte) (event.Subscription, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "QuestionUnflagged", questionIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketQuestionUnflagged)
				if err := _Market.contract.UnpackLog(event, "QuestionUnflagged", log); err != nil {
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

// ParseQuestionUnflagged is a log parse operation binding the contract event 0x052435bc04fc49113a7bfd9198a92c0852ca622a621800f6da66d4b29b786c05.
//
// Solidity: event QuestionUnflagged(bytes32 indexed questionID)
func (_Market *MarketFilterer) ParseQuestionUnflagged(log types.Log) (*MarketQuestionUnflagged, error) {
	event := new(MarketQuestionUnflagged)
	if err := _Market.contract.UnpackLog(event, "QuestionUnflagged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketQuestionUnpausedIterator is returned from FilterQuestionUnpaused and is used to iterate over the raw logs and unpacked data for QuestionUnpaused events raised by the Market contract.
type MarketQuestionUnpausedIterator struct {
	Event *MarketQuestionUnpaused // Event containing the contract specifics and raw log

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
func (it *MarketQuestionUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketQuestionUnpaused)
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
		it.Event = new(MarketQuestionUnpaused)
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
func (it *MarketQuestionUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketQuestionUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketQuestionUnpaused represents a QuestionUnpaused event raised by the Market contract.
type MarketQuestionUnpaused struct {
	QuestionID [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterQuestionUnpaused is a free log retrieval operation binding the contract event 0x92d28918c5574e7fc0f4f948c39502682c81cfb4089b07b83f95b3264e5e5e06.
//
// Solidity: event QuestionUnpaused(bytes32 indexed questionID)
func (_Market *MarketFilterer) FilterQuestionUnpaused(opts *bind.FilterOpts, questionID [][32]byte) (*MarketQuestionUnpausedIterator, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "QuestionUnpaused", questionIDRule)
	if err != nil {
		return nil, err
	}
	return &MarketQuestionUnpausedIterator{contract: _Market.contract, event: "QuestionUnpaused", logs: logs, sub: sub}, nil
}

// WatchQuestionUnpaused is a free log subscription operation binding the contract event 0x92d28918c5574e7fc0f4f948c39502682c81cfb4089b07b83f95b3264e5e5e06.
//
// Solidity: event QuestionUnpaused(bytes32 indexed questionID)
func (_Market *MarketFilterer) WatchQuestionUnpaused(opts *bind.WatchOpts, sink chan<- *MarketQuestionUnpaused, questionID [][32]byte) (event.Subscription, error) {

	var questionIDRule []interface{}
	for _, questionIDItem := range questionID {
		questionIDRule = append(questionIDRule, questionIDItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "QuestionUnpaused", questionIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketQuestionUnpaused)
				if err := _Market.contract.UnpackLog(event, "QuestionUnpaused", log); err != nil {
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

// ParseQuestionUnpaused is a log parse operation binding the contract event 0x92d28918c5574e7fc0f4f948c39502682c81cfb4089b07b83f95b3264e5e5e06.
//
// Solidity: event QuestionUnpaused(bytes32 indexed questionID)
func (_Market *MarketFilterer) ParseQuestionUnpaused(log types.Log) (*MarketQuestionUnpaused, error) {
	event := new(MarketQuestionUnpaused)
	if err := _Market.contract.UnpackLog(event, "QuestionUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketRemovedAdminIterator is returned from FilterRemovedAdmin and is used to iterate over the raw logs and unpacked data for RemovedAdmin events raised by the Market contract.
type MarketRemovedAdminIterator struct {
	Event *MarketRemovedAdmin // Event containing the contract specifics and raw log

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
func (it *MarketRemovedAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketRemovedAdmin)
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
		it.Event = new(MarketRemovedAdmin)
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
func (it *MarketRemovedAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketRemovedAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketRemovedAdmin represents a RemovedAdmin event raised by the Market contract.
type MarketRemovedAdmin struct {
	Admin        common.Address
	RemovedAdmin common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemovedAdmin is a free log retrieval operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed admin, address indexed removedAdmin)
func (_Market *MarketFilterer) FilterRemovedAdmin(opts *bind.FilterOpts, admin []common.Address, removedAdmin []common.Address) (*MarketRemovedAdminIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var removedAdminRule []interface{}
	for _, removedAdminItem := range removedAdmin {
		removedAdminRule = append(removedAdminRule, removedAdminItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "RemovedAdmin", adminRule, removedAdminRule)
	if err != nil {
		return nil, err
	}
	return &MarketRemovedAdminIterator{contract: _Market.contract, event: "RemovedAdmin", logs: logs, sub: sub}, nil
}

// WatchRemovedAdmin is a free log subscription operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed admin, address indexed removedAdmin)
func (_Market *MarketFilterer) WatchRemovedAdmin(opts *bind.WatchOpts, sink chan<- *MarketRemovedAdmin, admin []common.Address, removedAdmin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var removedAdminRule []interface{}
	for _, removedAdminItem := range removedAdmin {
		removedAdminRule = append(removedAdminRule, removedAdminItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "RemovedAdmin", adminRule, removedAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketRemovedAdmin)
				if err := _Market.contract.UnpackLog(event, "RemovedAdmin", log); err != nil {
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

// ParseRemovedAdmin is a log parse operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed admin, address indexed removedAdmin)
func (_Market *MarketFilterer) ParseRemovedAdmin(log types.Log) (*MarketRemovedAdmin, error) {
	event := new(MarketRemovedAdmin)
	if err := _Market.contract.UnpackLog(event, "RemovedAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
