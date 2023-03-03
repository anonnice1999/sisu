// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package example

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

// BaseContractMessage is an auto generated low-level Go binding around an user-defined struct.
type BaseContractMessage struct {
	Vault       common.Address
	CallerChain *big.Int
	Caller      common.Address
	Message     []byte
}

// ExampleMetaData contains all meta data concerning the Example contract.
var ExampleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"MessageReceived\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"callerChain\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structBaseContract.Message\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"onReceive\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"code\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506104e9806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80637ffc3e7714610030575b600080fd5b61004a60048036038101906100459190610199565b610060565b60405161005791906101fe565b60405180910390f35b60008073ffffffffffffffffffffffffffffffffffffffff1682600001602081019061008c9190610277565b73ffffffffffffffffffffffffffffffffffffffff16036100ac57600080fd5b6000808154809291906100be906102dd565b919050555060018280606001906100d59190610334565b905011610117576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161010e906103f4565b60405180910390fd5b7f6c3543130e60b15dd1f1f886e61a65ec337da34b356e078fcfe585b5578f39ef60005483806060019061014b9190610334565b60405161015a93929190610481565b60405180910390a160009050919050565b600080fd5b600080fd5b600080fd5b6000608082840312156101905761018f610175565b5b81905092915050565b6000602082840312156101af576101ae61016b565b5b600082013567ffffffffffffffff8111156101cd576101cc610170565b5b6101d98482850161017a565b91505092915050565b600060ff82169050919050565b6101f8816101e2565b82525050565b600060208201905061021360008301846101ef565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061024482610219565b9050919050565b61025481610239565b811461025f57600080fd5b50565b6000813590506102718161024b565b92915050565b60006020828403121561028d5761028c61016b565b5b600061029b84828501610262565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000819050919050565b60006102e8826102d3565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361031a576103196102a4565b5b600182019050919050565b600080fd5b600080fd5b600080fd5b6000808335600160200384360303811261035157610350610325565b5b80840192508235915067ffffffffffffffff8211156103735761037261032a565b5b60208301925060018202360383131561038f5761038e61032f565b5b509250929050565b600082825260208201905092915050565b7f696e76616c6964206d6573736167650000000000000000000000000000000000600082015250565b60006103de600f83610397565b91506103e9826103a8565b602082019050919050565b6000602082019050818103600083015261040d816103d1565b9050919050565b61041d816102d3565b82525050565b600082825260208201905092915050565b82818337600083830152505050565b6000601f19601f8301169050919050565b60006104608385610423565b935061046d838584610434565b61047683610443565b840190509392505050565b60006040820190506104966000830186610414565b81810360208301526104a9818486610454565b905094935050505056fea2646970667358221220404eb65938b0c6fdf78c8b901066a75c0de5f527877693a698b0b82d2ea3f2d064736f6c63430008130033",
}

// ExampleABI is the input ABI used to generate the binding from.
// Deprecated: Use ExampleMetaData.ABI instead.
var ExampleABI = ExampleMetaData.ABI

// ExampleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExampleMetaData.Bin instead.
var ExampleBin = ExampleMetaData.Bin

// DeployExample deploys a new Ethereum contract, binding an instance of Example to it.
func DeployExample(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Example, error) {
	parsed, err := ExampleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExampleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Example{ExampleCaller: ExampleCaller{contract: contract}, ExampleTransactor: ExampleTransactor{contract: contract}, ExampleFilterer: ExampleFilterer{contract: contract}}, nil
}

// Example is an auto generated Go binding around an Ethereum contract.
type Example struct {
	ExampleCaller     // Read-only binding to the contract
	ExampleTransactor // Write-only binding to the contract
	ExampleFilterer   // Log filterer for contract events
}

// ExampleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExampleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExampleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExampleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExampleSession struct {
	Contract     *Example          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExampleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExampleCallerSession struct {
	Contract *ExampleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ExampleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExampleTransactorSession struct {
	Contract     *ExampleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ExampleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExampleRaw struct {
	Contract *Example // Generic contract binding to access the raw methods on
}

// ExampleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExampleCallerRaw struct {
	Contract *ExampleCaller // Generic read-only contract binding to access the raw methods on
}

// ExampleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExampleTransactorRaw struct {
	Contract *ExampleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExample creates a new instance of Example, bound to a specific deployed contract.
func NewExample(address common.Address, backend bind.ContractBackend) (*Example, error) {
	contract, err := bindExample(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Example{ExampleCaller: ExampleCaller{contract: contract}, ExampleTransactor: ExampleTransactor{contract: contract}, ExampleFilterer: ExampleFilterer{contract: contract}}, nil
}

// NewExampleCaller creates a new read-only instance of Example, bound to a specific deployed contract.
func NewExampleCaller(address common.Address, caller bind.ContractCaller) (*ExampleCaller, error) {
	contract, err := bindExample(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleCaller{contract: contract}, nil
}

// NewExampleTransactor creates a new write-only instance of Example, bound to a specific deployed contract.
func NewExampleTransactor(address common.Address, transactor bind.ContractTransactor) (*ExampleTransactor, error) {
	contract, err := bindExample(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleTransactor{contract: contract}, nil
}

// NewExampleFilterer creates a new log filterer instance of Example, bound to a specific deployed contract.
func NewExampleFilterer(address common.Address, filterer bind.ContractFilterer) (*ExampleFilterer, error) {
	contract, err := bindExample(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExampleFilterer{contract: contract}, nil
}

// bindExample binds a generic wrapper to an already deployed contract.
func bindExample(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExampleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Example *ExampleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Example.Contract.ExampleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Example *ExampleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Example.Contract.ExampleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Example *ExampleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Example.Contract.ExampleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Example *ExampleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Example.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Example *ExampleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Example.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Example *ExampleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Example.Contract.contract.Transact(opts, method, params...)
}

// OnReceive is a paid mutator transaction binding the contract method 0x7ffc3e77.
//
// Solidity: function onReceive((address,uint256,address,bytes) input) returns(uint8 code)
func (_Example *ExampleTransactor) OnReceive(opts *bind.TransactOpts, input BaseContractMessage) (*types.Transaction, error) {
	return _Example.contract.Transact(opts, "onReceive", input)
}

// OnReceive is a paid mutator transaction binding the contract method 0x7ffc3e77.
//
// Solidity: function onReceive((address,uint256,address,bytes) input) returns(uint8 code)
func (_Example *ExampleSession) OnReceive(input BaseContractMessage) (*types.Transaction, error) {
	return _Example.Contract.OnReceive(&_Example.TransactOpts, input)
}

// OnReceive is a paid mutator transaction binding the contract method 0x7ffc3e77.
//
// Solidity: function onReceive((address,uint256,address,bytes) input) returns(uint8 code)
func (_Example *ExampleTransactorSession) OnReceive(input BaseContractMessage) (*types.Transaction, error) {
	return _Example.Contract.OnReceive(&_Example.TransactOpts, input)
}

// ExampleMessageReceivedIterator is returned from FilterMessageReceived and is used to iterate over the raw logs and unpacked data for MessageReceived events raised by the Example contract.
type ExampleMessageReceivedIterator struct {
	Event *ExampleMessageReceived // Event containing the contract specifics and raw log

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
func (it *ExampleMessageReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleMessageReceived)
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
		it.Event = new(ExampleMessageReceived)
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
func (it *ExampleMessageReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleMessageReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleMessageReceived represents a MessageReceived event raised by the Example contract.
type ExampleMessageReceived struct {
	Id      *big.Int
	Message []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMessageReceived is a free log retrieval operation binding the contract event 0x6c3543130e60b15dd1f1f886e61a65ec337da34b356e078fcfe585b5578f39ef.
//
// Solidity: event MessageReceived(uint256 id, bytes message)
func (_Example *ExampleFilterer) FilterMessageReceived(opts *bind.FilterOpts) (*ExampleMessageReceivedIterator, error) {

	logs, sub, err := _Example.contract.FilterLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return &ExampleMessageReceivedIterator{contract: _Example.contract, event: "MessageReceived", logs: logs, sub: sub}, nil
}

// WatchMessageReceived is a free log subscription operation binding the contract event 0x6c3543130e60b15dd1f1f886e61a65ec337da34b356e078fcfe585b5578f39ef.
//
// Solidity: event MessageReceived(uint256 id, bytes message)
func (_Example *ExampleFilterer) WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *ExampleMessageReceived) (event.Subscription, error) {

	logs, sub, err := _Example.contract.WatchLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleMessageReceived)
				if err := _Example.contract.UnpackLog(event, "MessageReceived", log); err != nil {
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

// ParseMessageReceived is a log parse operation binding the contract event 0x6c3543130e60b15dd1f1f886e61a65ec337da34b356e078fcfe585b5578f39ef.
//
// Solidity: event MessageReceived(uint256 id, bytes message)
func (_Example *ExampleFilterer) ParseMessageReceived(log types.Log) (*ExampleMessageReceived, error) {
	event := new(ExampleMessageReceived)
	if err := _Example.contract.UnpackLog(event, "MessageReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
