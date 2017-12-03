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

// AsyncCall_InterfaceABI is the input ABI used to generate the binding from.
const AsyncCall_InterfaceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"publicCallData\",\"type\":\"bytes\"},{\"name\":\"privateData\",\"type\":\"bytes\"}],\"name\":\"beginCall\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pendingCalls\",\"outputs\":[{\"name\":\"callData\",\"type\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"},{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"msgDataHash\",\"type\":\"bytes32\"}],\"name\":\"continueCall\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_waitConfirmation\",\"type\":\"bool\"}],\"name\":\"changeConfirmation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"waitConfirmation\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"CallStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"CallFinish\",\"type\":\"event\"}]"

// AsyncCall_InterfaceBin is the compiled bytecode used for deploying new contracts.
const AsyncCall_InterfaceBin = `0x`

// DeployAsyncCall_Interface deploys a new Ethereum contract, binding an instance of AsyncCall_Interface to it.
func DeployAsyncCall_Interface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AsyncCall_Interface, error) {
	parsed, err := abi.JSON(strings.NewReader(AsyncCall_InterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AsyncCall_InterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AsyncCall_Interface{AsyncCall_InterfaceCaller: AsyncCall_InterfaceCaller{contract: contract}, AsyncCall_InterfaceTransactor: AsyncCall_InterfaceTransactor{contract: contract}}, nil
}

// AsyncCall_Interface is an auto generated Go binding around an Ethereum contract.
type AsyncCall_Interface struct {
	AsyncCall_InterfaceCaller     // Read-only binding to the contract
	AsyncCall_InterfaceTransactor // Write-only binding to the contract
}

// AsyncCall_InterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type AsyncCall_InterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AsyncCall_InterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AsyncCall_InterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AsyncCall_InterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AsyncCall_InterfaceSession struct {
	Contract     *AsyncCall_Interface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AsyncCall_InterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AsyncCall_InterfaceCallerSession struct {
	Contract *AsyncCall_InterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// AsyncCall_InterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AsyncCall_InterfaceTransactorSession struct {
	Contract     *AsyncCall_InterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// AsyncCall_InterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type AsyncCall_InterfaceRaw struct {
	Contract *AsyncCall_Interface // Generic contract binding to access the raw methods on
}

// AsyncCall_InterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AsyncCall_InterfaceCallerRaw struct {
	Contract *AsyncCall_InterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// AsyncCall_InterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AsyncCall_InterfaceTransactorRaw struct {
	Contract *AsyncCall_InterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAsyncCall_Interface creates a new instance of AsyncCall_Interface, bound to a specific deployed contract.
func NewAsyncCall_Interface(address common.Address, backend bind.ContractBackend) (*AsyncCall_Interface, error) {
	contract, err := bindAsyncCall_Interface(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AsyncCall_Interface{AsyncCall_InterfaceCaller: AsyncCall_InterfaceCaller{contract: contract}, AsyncCall_InterfaceTransactor: AsyncCall_InterfaceTransactor{contract: contract}}, nil
}

// NewAsyncCall_InterfaceCaller creates a new read-only instance of AsyncCall_Interface, bound to a specific deployed contract.
func NewAsyncCall_InterfaceCaller(address common.Address, caller bind.ContractCaller) (*AsyncCall_InterfaceCaller, error) {
	contract, err := bindAsyncCall_Interface(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &AsyncCall_InterfaceCaller{contract: contract}, nil
}

// NewAsyncCall_InterfaceTransactor creates a new write-only instance of AsyncCall_Interface, bound to a specific deployed contract.
func NewAsyncCall_InterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*AsyncCall_InterfaceTransactor, error) {
	contract, err := bindAsyncCall_Interface(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &AsyncCall_InterfaceTransactor{contract: contract}, nil
}

// bindAsyncCall_Interface binds a generic wrapper to an already deployed contract.
func bindAsyncCall_Interface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AsyncCall_InterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AsyncCall_Interface *AsyncCall_InterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AsyncCall_Interface.Contract.AsyncCall_InterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AsyncCall_Interface *AsyncCall_InterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AsyncCall_Interface.Contract.AsyncCall_InterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AsyncCall_Interface *AsyncCall_InterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AsyncCall_Interface.Contract.AsyncCall_InterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AsyncCall_Interface *AsyncCall_InterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AsyncCall_Interface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AsyncCall_Interface *AsyncCall_InterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AsyncCall_Interface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AsyncCall_Interface *AsyncCall_InterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AsyncCall_Interface.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AsyncCall_Interface *AsyncCall_InterfaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AsyncCall_Interface.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AsyncCall_Interface *AsyncCall_InterfaceSession) Owner() (common.Address, error) {
	return _AsyncCall_Interface.Contract.Owner(&_AsyncCall_Interface.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AsyncCall_Interface *AsyncCall_InterfaceCallerSession) Owner() (common.Address, error) {
	return _AsyncCall_Interface.Contract.Owner(&_AsyncCall_Interface.CallOpts)
}

// PendingCalls is a free data retrieval call binding the contract method 0x32fdd45c.
//
// Solidity: function pendingCalls( bytes32) constant returns(callData bytes, sender address, approved bool, success bool)
func (_AsyncCall_Interface *AsyncCall_InterfaceCaller) PendingCalls(opts *bind.CallOpts, arg0 [32]byte) (struct {
	CallData []byte
	Sender   common.Address
	Approved bool
	Success  bool
}, error) {
	ret := new(struct {
		CallData []byte
		Sender   common.Address
		Approved bool
		Success  bool
	})
	out := ret
	err := _AsyncCall_Interface.contract.Call(opts, out, "pendingCalls", arg0)
	return *ret, err
}

// PendingCalls is a free data retrieval call binding the contract method 0x32fdd45c.
//
// Solidity: function pendingCalls( bytes32) constant returns(callData bytes, sender address, approved bool, success bool)
func (_AsyncCall_Interface *AsyncCall_InterfaceSession) PendingCalls(arg0 [32]byte) (struct {
	CallData []byte
	Sender   common.Address
	Approved bool
	Success  bool
}, error) {
	return _AsyncCall_Interface.Contract.PendingCalls(&_AsyncCall_Interface.CallOpts, arg0)
}

// PendingCalls is a free data retrieval call binding the contract method 0x32fdd45c.
//
// Solidity: function pendingCalls( bytes32) constant returns(callData bytes, sender address, approved bool, success bool)
func (_AsyncCall_Interface *AsyncCall_InterfaceCallerSession) PendingCalls(arg0 [32]byte) (struct {
	CallData []byte
	Sender   common.Address
	Approved bool
	Success  bool
}, error) {
	return _AsyncCall_Interface.Contract.PendingCalls(&_AsyncCall_Interface.CallOpts, arg0)
}

// WaitConfirmation is a free data retrieval call binding the contract method 0xac72bb98.
//
// Solidity: function waitConfirmation() constant returns(bool)
func (_AsyncCall_Interface *AsyncCall_InterfaceCaller) WaitConfirmation(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AsyncCall_Interface.contract.Call(opts, out, "waitConfirmation")
	return *ret0, err
}

// WaitConfirmation is a free data retrieval call binding the contract method 0xac72bb98.
//
// Solidity: function waitConfirmation() constant returns(bool)
func (_AsyncCall_Interface *AsyncCall_InterfaceSession) WaitConfirmation() (bool, error) {
	return _AsyncCall_Interface.Contract.WaitConfirmation(&_AsyncCall_Interface.CallOpts)
}

// WaitConfirmation is a free data retrieval call binding the contract method 0xac72bb98.
//
// Solidity: function waitConfirmation() constant returns(bool)
func (_AsyncCall_Interface *AsyncCall_InterfaceCallerSession) WaitConfirmation() (bool, error) {
	return _AsyncCall_Interface.Contract.WaitConfirmation(&_AsyncCall_Interface.CallOpts)
}

// BeginCall is a paid mutator transaction binding the contract method 0x203eaf27.
//
// Solidity: function beginCall(publicCallData bytes, privateData bytes) returns()
func (_AsyncCall_Interface *AsyncCall_InterfaceTransactor) BeginCall(opts *bind.TransactOpts, publicCallData []byte, privateData []byte) (*types.Transaction, error) {
	return _AsyncCall_Interface.contract.Transact(opts, "beginCall", publicCallData, privateData)
}

// BeginCall is a paid mutator transaction binding the contract method 0x203eaf27.
//
// Solidity: function beginCall(publicCallData bytes, privateData bytes) returns()
func (_AsyncCall_Interface *AsyncCall_InterfaceSession) BeginCall(publicCallData []byte, privateData []byte) (*types.Transaction, error) {
	return _AsyncCall_Interface.Contract.BeginCall(&_AsyncCall_Interface.TransactOpts, publicCallData, privateData)
}

// BeginCall is a paid mutator transaction binding the contract method 0x203eaf27.
//
// Solidity: function beginCall(publicCallData bytes, privateData bytes) returns()
func (_AsyncCall_Interface *AsyncCall_InterfaceTransactorSession) BeginCall(publicCallData []byte, privateData []byte) (*types.Transaction, error) {
	return _AsyncCall_Interface.Contract.BeginCall(&_AsyncCall_Interface.TransactOpts, publicCallData, privateData)
}

// ChangeConfirmation is a paid mutator transaction binding the contract method 0x44f5484a.
//
// Solidity: function changeConfirmation(_waitConfirmation bool) returns()
func (_AsyncCall_Interface *AsyncCall_InterfaceTransactor) ChangeConfirmation(opts *bind.TransactOpts, _waitConfirmation bool) (*types.Transaction, error) {
	return _AsyncCall_Interface.contract.Transact(opts, "changeConfirmation", _waitConfirmation)
}

// ChangeConfirmation is a paid mutator transaction binding the contract method 0x44f5484a.
//
// Solidity: function changeConfirmation(_waitConfirmation bool) returns()
func (_AsyncCall_Interface *AsyncCall_InterfaceSession) ChangeConfirmation(_waitConfirmation bool) (*types.Transaction, error) {
	return _AsyncCall_Interface.Contract.ChangeConfirmation(&_AsyncCall_Interface.TransactOpts, _waitConfirmation)
}

// ChangeConfirmation is a paid mutator transaction binding the contract method 0x44f5484a.
//
// Solidity: function changeConfirmation(_waitConfirmation bool) returns()
func (_AsyncCall_Interface *AsyncCall_InterfaceTransactorSession) ChangeConfirmation(_waitConfirmation bool) (*types.Transaction, error) {
	return _AsyncCall_Interface.Contract.ChangeConfirmation(&_AsyncCall_Interface.TransactOpts, _waitConfirmation)
}

// ContinueCall is a paid mutator transaction binding the contract method 0x411f2351.
//
// Solidity: function continueCall(msgDataHash bytes32) returns()
func (_AsyncCall_Interface *AsyncCall_InterfaceTransactor) ContinueCall(opts *bind.TransactOpts, msgDataHash [32]byte) (*types.Transaction, error) {
	return _AsyncCall_Interface.contract.Transact(opts, "continueCall", msgDataHash)
}

// ContinueCall is a paid mutator transaction binding the contract method 0x411f2351.
//
// Solidity: function continueCall(msgDataHash bytes32) returns()
func (_AsyncCall_Interface *AsyncCall_InterfaceSession) ContinueCall(msgDataHash [32]byte) (*types.Transaction, error) {
	return _AsyncCall_Interface.Contract.ContinueCall(&_AsyncCall_Interface.TransactOpts, msgDataHash)
}

// ContinueCall is a paid mutator transaction binding the contract method 0x411f2351.
//
// Solidity: function continueCall(msgDataHash bytes32) returns()
func (_AsyncCall_Interface *AsyncCall_InterfaceTransactorSession) ContinueCall(msgDataHash [32]byte) (*types.Transaction, error) {
	return _AsyncCall_Interface.Contract.ContinueCall(&_AsyncCall_Interface.TransactOpts, msgDataHash)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_AsyncCall_Interface *AsyncCall_InterfaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AsyncCall_Interface.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_AsyncCall_Interface *AsyncCall_InterfaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AsyncCall_Interface.Contract.TransferOwnership(&_AsyncCall_Interface.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_AsyncCall_Interface *AsyncCall_InterfaceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AsyncCall_Interface.Contract.TransferOwnership(&_AsyncCall_Interface.TransactOpts, newOwner)
}
