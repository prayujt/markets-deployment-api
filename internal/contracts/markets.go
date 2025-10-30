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

// MarketsMarket is an auto generated low-level Go binding around an user-defined struct.
type MarketsMarket struct {
	QuestionId    string
	ConditionId   string
	PositionIdYes string
	PositionIdNo  string
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"questionId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"conditionId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"positionIdYes\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"positionIdNo\",\"type\":\"string\"}],\"name\":\"MarketDeployed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"questionId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"conditionId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"positionIdYes\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"positionIdNo\",\"type\":\"string\"}],\"name\":\"deployMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMarkets\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"questionId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"conditionId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"positionIdYes\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"positionIdNo\",\"type\":\"string\"}],\"internalType\":\"structMarkets.Market[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// GetMarkets is a free data retrieval call binding the contract method 0xec2c9016.
//
// Solidity: function getMarkets() view returns((string,string,string,string)[])
func (_Contracts *ContractsCaller) GetMarkets(opts *bind.CallOpts) ([]MarketsMarket, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getMarkets")

	if err != nil {
		return *new([]MarketsMarket), err
	}

	out0 := *abi.ConvertType(out[0], new([]MarketsMarket)).(*[]MarketsMarket)

	return out0, err

}

// GetMarkets is a free data retrieval call binding the contract method 0xec2c9016.
//
// Solidity: function getMarkets() view returns((string,string,string,string)[])
func (_Contracts *ContractsSession) GetMarkets() ([]MarketsMarket, error) {
	return _Contracts.Contract.GetMarkets(&_Contracts.CallOpts)
}

// GetMarkets is a free data retrieval call binding the contract method 0xec2c9016.
//
// Solidity: function getMarkets() view returns((string,string,string,string)[])
func (_Contracts *ContractsCallerSession) GetMarkets() ([]MarketsMarket, error) {
	return _Contracts.Contract.GetMarkets(&_Contracts.CallOpts)
}

// MarketsCount is a free data retrieval call binding the contract method 0x1e8b5708.
//
// Solidity: function marketsCount() view returns(uint256)
func (_Contracts *ContractsCaller) MarketsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "marketsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketsCount is a free data retrieval call binding the contract method 0x1e8b5708.
//
// Solidity: function marketsCount() view returns(uint256)
func (_Contracts *ContractsSession) MarketsCount() (*big.Int, error) {
	return _Contracts.Contract.MarketsCount(&_Contracts.CallOpts)
}

// MarketsCount is a free data retrieval call binding the contract method 0x1e8b5708.
//
// Solidity: function marketsCount() view returns(uint256)
func (_Contracts *ContractsCallerSession) MarketsCount() (*big.Int, error) {
	return _Contracts.Contract.MarketsCount(&_Contracts.CallOpts)
}

// DeployMarket is a paid mutator transaction binding the contract method 0xb494960b.
//
// Solidity: function deployMarket(string questionId, string conditionId, string positionIdYes, string positionIdNo) returns(uint256 index)
func (_Contracts *ContractsTransactor) DeployMarket(opts *bind.TransactOpts, questionId string, conditionId string, positionIdYes string, positionIdNo string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "deployMarket", questionId, conditionId, positionIdYes, positionIdNo)
}

// DeployMarket is a paid mutator transaction binding the contract method 0xb494960b.
//
// Solidity: function deployMarket(string questionId, string conditionId, string positionIdYes, string positionIdNo) returns(uint256 index)
func (_Contracts *ContractsSession) DeployMarket(questionId string, conditionId string, positionIdYes string, positionIdNo string) (*types.Transaction, error) {
	return _Contracts.Contract.DeployMarket(&_Contracts.TransactOpts, questionId, conditionId, positionIdYes, positionIdNo)
}

// DeployMarket is a paid mutator transaction binding the contract method 0xb494960b.
//
// Solidity: function deployMarket(string questionId, string conditionId, string positionIdYes, string positionIdNo) returns(uint256 index)
func (_Contracts *ContractsTransactorSession) DeployMarket(questionId string, conditionId string, positionIdYes string, positionIdNo string) (*types.Transaction, error) {
	return _Contracts.Contract.DeployMarket(&_Contracts.TransactOpts, questionId, conditionId, positionIdYes, positionIdNo)
}

// ContractsMarketDeployedIterator is returned from FilterMarketDeployed and is used to iterate over the raw logs and unpacked data for MarketDeployed events raised by the Contracts contract.
type ContractsMarketDeployedIterator struct {
	Event *ContractsMarketDeployed // Event containing the contract specifics and raw log

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
func (it *ContractsMarketDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsMarketDeployed)
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
		it.Event = new(ContractsMarketDeployed)
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
func (it *ContractsMarketDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsMarketDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsMarketDeployed represents a MarketDeployed event raised by the Contracts contract.
type ContractsMarketDeployed struct {
	Index         *big.Int
	QuestionId    string
	ConditionId   string
	PositionIdYes string
	PositionIdNo  string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMarketDeployed is a free log retrieval operation binding the contract event 0x34fd3c72798e638687b037641bb6d6a098005dc94d3e35582ae52ab44e9aadd8.
//
// Solidity: event MarketDeployed(uint256 indexed index, string questionId, string conditionId, string positionIdYes, string positionIdNo)
func (_Contracts *ContractsFilterer) FilterMarketDeployed(opts *bind.FilterOpts, index []*big.Int) (*ContractsMarketDeployedIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "MarketDeployed", indexRule)
	if err != nil {
		return nil, err
	}
	return &ContractsMarketDeployedIterator{contract: _Contracts.contract, event: "MarketDeployed", logs: logs, sub: sub}, nil
}

// WatchMarketDeployed is a free log subscription operation binding the contract event 0x34fd3c72798e638687b037641bb6d6a098005dc94d3e35582ae52ab44e9aadd8.
//
// Solidity: event MarketDeployed(uint256 indexed index, string questionId, string conditionId, string positionIdYes, string positionIdNo)
func (_Contracts *ContractsFilterer) WatchMarketDeployed(opts *bind.WatchOpts, sink chan<- *ContractsMarketDeployed, index []*big.Int) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "MarketDeployed", indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsMarketDeployed)
				if err := _Contracts.contract.UnpackLog(event, "MarketDeployed", log); err != nil {
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

// ParseMarketDeployed is a log parse operation binding the contract event 0x34fd3c72798e638687b037641bb6d6a098005dc94d3e35582ae52ab44e9aadd8.
//
// Solidity: event MarketDeployed(uint256 indexed index, string questionId, string conditionId, string positionIdYes, string positionIdNo)
func (_Contracts *ContractsFilterer) ParseMarketDeployed(log types.Log) (*ContractsMarketDeployed, error) {
	event := new(ContractsMarketDeployed)
	if err := _Contracts.contract.UnpackLog(event, "MarketDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
