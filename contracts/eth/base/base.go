// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package base

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

// BaseMetaData contains all meta data concerning the Base contract.
var BaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"callerChain\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structBaseContract.Message\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"onReceive\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"code\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610321806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80637ffc3e7714610030575b600080fd5b61004a6004803603810190610045919061015f565b610060565b60405161005791906101c4565b60405180910390f35b60008073ffffffffffffffffffffffffffffffffffffffff1682600001602081019061008c919061023d565b73ffffffffffffffffffffffffffffffffffffffff16036100ac57600080fd5b8160000160208101906100bf919061023d565b600160008054815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600080815480929190610123906102a3565b919050555060039050919050565b600080fd5b600080fd5b600080fd5b6000608082840312156101565761015561013b565b5b81905092915050565b60006020828403121561017557610174610131565b5b600082013567ffffffffffffffff81111561019357610192610136565b5b61019f84828501610140565b91505092915050565b600060ff82169050919050565b6101be816101a8565b82525050565b60006020820190506101d960008301846101b5565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061020a826101df565b9050919050565b61021a816101ff565b811461022557600080fd5b50565b60008135905061023781610211565b92915050565b60006020828403121561025357610252610131565b5b600061026184828501610228565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000819050919050565b60006102ae82610299565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036102e0576102df61026a565b5b60018201905091905056fea264697066735822122071cead7dd5af26dfbe7de8bc7ebc8cd591fb9b8b7fe6a41ed6885c9f3f5c2cda64736f6c63430008130033",
}

// BaseABI is the input ABI used to generate the binding from.
// Deprecated: Use BaseMetaData.ABI instead.
var BaseABI = BaseMetaData.ABI

// BaseBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BaseMetaData.Bin instead.
var BaseBin = BaseMetaData.Bin

// DeployBase deploys a new Ethereum contract, binding an instance of Base to it.
func DeployBase(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Base, error) {
	parsed, err := BaseMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BaseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Base{BaseCaller: BaseCaller{contract: contract}, BaseTransactor: BaseTransactor{contract: contract}, BaseFilterer: BaseFilterer{contract: contract}}, nil
}

// Base is an auto generated Go binding around an Ethereum contract.
type Base struct {
	BaseCaller     // Read-only binding to the contract
	BaseTransactor // Write-only binding to the contract
	BaseFilterer   // Log filterer for contract events
}

// BaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseSession struct {
	Contract     *Base             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseCallerSession struct {
	Contract *BaseCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseTransactorSession struct {
	Contract     *BaseTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseRaw struct {
	Contract *Base // Generic contract binding to access the raw methods on
}

// BaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseCallerRaw struct {
	Contract *BaseCaller // Generic read-only contract binding to access the raw methods on
}

// BaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseTransactorRaw struct {
	Contract *BaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBase creates a new instance of Base, bound to a specific deployed contract.
func NewBase(address common.Address, backend bind.ContractBackend) (*Base, error) {
	contract, err := bindBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Base{BaseCaller: BaseCaller{contract: contract}, BaseTransactor: BaseTransactor{contract: contract}, BaseFilterer: BaseFilterer{contract: contract}}, nil
}

// NewBaseCaller creates a new read-only instance of Base, bound to a specific deployed contract.
func NewBaseCaller(address common.Address, caller bind.ContractCaller) (*BaseCaller, error) {
	contract, err := bindBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseCaller{contract: contract}, nil
}

// NewBaseTransactor creates a new write-only instance of Base, bound to a specific deployed contract.
func NewBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseTransactor, error) {
	contract, err := bindBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseTransactor{contract: contract}, nil
}

// NewBaseFilterer creates a new log filterer instance of Base, bound to a specific deployed contract.
func NewBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseFilterer, error) {
	contract, err := bindBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseFilterer{contract: contract}, nil
}

// bindBase binds a generic wrapper to an already deployed contract.
func bindBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Base *BaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Base.Contract.BaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Base *BaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Base.Contract.BaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Base *BaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Base.Contract.BaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Base *BaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Base.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Base *BaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Base.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Base *BaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Base.Contract.contract.Transact(opts, method, params...)
}

// OnReceive is a paid mutator transaction binding the contract method 0x7ffc3e77.
//
// Solidity: function onReceive((address,uint256,address,bytes) input) returns(uint8 code)
func (_Base *BaseTransactor) OnReceive(opts *bind.TransactOpts, input BaseContractMessage) (*types.Transaction, error) {
	return _Base.contract.Transact(opts, "onReceive", input)
}

// OnReceive is a paid mutator transaction binding the contract method 0x7ffc3e77.
//
// Solidity: function onReceive((address,uint256,address,bytes) input) returns(uint8 code)
func (_Base *BaseSession) OnReceive(input BaseContractMessage) (*types.Transaction, error) {
	return _Base.Contract.OnReceive(&_Base.TransactOpts, input)
}

// OnReceive is a paid mutator transaction binding the contract method 0x7ffc3e77.
//
// Solidity: function onReceive((address,uint256,address,bytes) input) returns(uint8 code)
func (_Base *BaseTransactorSession) OnReceive(input BaseContractMessage) (*types.Transaction, error) {
	return _Base.Contract.OnReceive(&_Base.TransactOpts, input)
}
