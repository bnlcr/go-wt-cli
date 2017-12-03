// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gen

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Base_InterfaceABI is the input ABI used to generate the binding from.
const Base_InterfaceABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractType\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Base_InterfaceBin is the compiled bytecode used for deploying new contracts.
const Base_InterfaceBin = `0x6060604052341561000f57600080fd5b60e28061001d6000396000f3006060604052600436106049576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806354fd4d5014604e578063cb2ef6f714607c575b600080fd5b3415605857600080fd5b605e60aa565b60405180826000191660001916815260200191505060405180910390f35b3415608657600080fd5b608c60b0565b60405180826000191660001916815260200191505060405180910390f35b60005481565b600154815600a165627a7a72305820fde0f14a801f5c45c08b6a75ce44f2449305e1b4d5923932bce5105564d2f1cb0029`

// DeployBase_Interface deploys a new Ethereum contract, binding an instance of Base_Interface to it.
func DeployBase_Interface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Base_Interface, error) {
	parsed, err := abi.JSON(strings.NewReader(Base_InterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Base_InterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Base_Interface{Base_InterfaceCaller: Base_InterfaceCaller{contract: contract}, Base_InterfaceTransactor: Base_InterfaceTransactor{contract: contract}}, nil
}

// Base_Interface is an auto generated Go binding around an Ethereum contract.
type Base_Interface struct {
	Base_InterfaceCaller     // Read-only binding to the contract
	Base_InterfaceTransactor // Write-only binding to the contract
}

// Base_InterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type Base_InterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Base_InterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Base_InterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Base_InterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Base_InterfaceSession struct {
	Contract     *Base_Interface   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Base_InterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Base_InterfaceCallerSession struct {
	Contract *Base_InterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// Base_InterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Base_InterfaceTransactorSession struct {
	Contract     *Base_InterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// Base_InterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type Base_InterfaceRaw struct {
	Contract *Base_Interface // Generic contract binding to access the raw methods on
}

// Base_InterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Base_InterfaceCallerRaw struct {
	Contract *Base_InterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// Base_InterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Base_InterfaceTransactorRaw struct {
	Contract *Base_InterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBase_Interface creates a new instance of Base_Interface, bound to a specific deployed contract.
func NewBase_Interface(address common.Address, backend bind.ContractBackend) (*Base_Interface, error) {
	contract, err := bindBase_Interface(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Base_Interface{Base_InterfaceCaller: Base_InterfaceCaller{contract: contract}, Base_InterfaceTransactor: Base_InterfaceTransactor{contract: contract}}, nil
}

// NewBase_InterfaceCaller creates a new read-only instance of Base_Interface, bound to a specific deployed contract.
func NewBase_InterfaceCaller(address common.Address, caller bind.ContractCaller) (*Base_InterfaceCaller, error) {
	contract, err := bindBase_Interface(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &Base_InterfaceCaller{contract: contract}, nil
}

// NewBase_InterfaceTransactor creates a new write-only instance of Base_Interface, bound to a specific deployed contract.
func NewBase_InterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*Base_InterfaceTransactor, error) {
	contract, err := bindBase_Interface(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &Base_InterfaceTransactor{contract: contract}, nil
}

// bindBase_Interface binds a generic wrapper to an already deployed contract.
func bindBase_Interface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Base_InterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Base_Interface *Base_InterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Base_Interface.Contract.Base_InterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Base_Interface *Base_InterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Base_Interface.Contract.Base_InterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Base_Interface *Base_InterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Base_Interface.Contract.Base_InterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Base_Interface *Base_InterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Base_Interface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Base_Interface *Base_InterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Base_Interface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Base_Interface *Base_InterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Base_Interface.Contract.contract.Transact(opts, method, params...)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() constant returns(bytes32)
func (_Base_Interface *Base_InterfaceCaller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Base_Interface.contract.Call(opts, out, "contractType")
	return *ret0, err
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() constant returns(bytes32)
func (_Base_Interface *Base_InterfaceSession) ContractType() ([32]byte, error) {
	return _Base_Interface.Contract.ContractType(&_Base_Interface.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() constant returns(bytes32)
func (_Base_Interface *Base_InterfaceCallerSession) ContractType() ([32]byte, error) {
	return _Base_Interface.Contract.ContractType(&_Base_Interface.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Base_Interface *Base_InterfaceCaller) Version(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Base_Interface.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Base_Interface *Base_InterfaceSession) Version() ([32]byte, error) {
	return _Base_Interface.Contract.Version(&_Base_Interface.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Base_Interface *Base_InterfaceCallerSession) Version() ([32]byte, error) {
	return _Base_Interface.Contract.Version(&_Base_Interface.CallOpts)
}
