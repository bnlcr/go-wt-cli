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

// Index_InterfaceABI is the input ABI used to generate the binding from.
const Index_InterfaceABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"LifToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Index_InterfaceBin is the compiled bytecode used for deploying new contracts.
const Index_InterfaceBin = `0x6060604052341561000f57600080fd5b60e78061001d6000396000f300606060405260043610603f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063554d8b37146044575b600080fd5b3415604e57600080fd5b60546096565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a72305820e014a6290b607e592286e41a4d873bcea76273117b1a0ac3f62ddd221b1e318a0029`

// DeployIndex_Interface deploys a new Ethereum contract, binding an instance of Index_Interface to it.
func DeployIndex_Interface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Index_Interface, error) {
	parsed, err := abi.JSON(strings.NewReader(Index_InterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Index_InterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Index_Interface{Index_InterfaceCaller: Index_InterfaceCaller{contract: contract}, Index_InterfaceTransactor: Index_InterfaceTransactor{contract: contract}}, nil
}

// Index_Interface is an auto generated Go binding around an Ethereum contract.
type Index_Interface struct {
	Index_InterfaceCaller     // Read-only binding to the contract
	Index_InterfaceTransactor // Write-only binding to the contract
}

// Index_InterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type Index_InterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Index_InterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Index_InterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Index_InterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Index_InterfaceSession struct {
	Contract     *Index_Interface  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Index_InterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Index_InterfaceCallerSession struct {
	Contract *Index_InterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// Index_InterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Index_InterfaceTransactorSession struct {
	Contract     *Index_InterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// Index_InterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type Index_InterfaceRaw struct {
	Contract *Index_Interface // Generic contract binding to access the raw methods on
}

// Index_InterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Index_InterfaceCallerRaw struct {
	Contract *Index_InterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// Index_InterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Index_InterfaceTransactorRaw struct {
	Contract *Index_InterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIndex_Interface creates a new instance of Index_Interface, bound to a specific deployed contract.
func NewIndex_Interface(address common.Address, backend bind.ContractBackend) (*Index_Interface, error) {
	contract, err := bindIndex_Interface(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Index_Interface{Index_InterfaceCaller: Index_InterfaceCaller{contract: contract}, Index_InterfaceTransactor: Index_InterfaceTransactor{contract: contract}}, nil
}

// NewIndex_InterfaceCaller creates a new read-only instance of Index_Interface, bound to a specific deployed contract.
func NewIndex_InterfaceCaller(address common.Address, caller bind.ContractCaller) (*Index_InterfaceCaller, error) {
	contract, err := bindIndex_Interface(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &Index_InterfaceCaller{contract: contract}, nil
}

// NewIndex_InterfaceTransactor creates a new write-only instance of Index_Interface, bound to a specific deployed contract.
func NewIndex_InterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*Index_InterfaceTransactor, error) {
	contract, err := bindIndex_Interface(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &Index_InterfaceTransactor{contract: contract}, nil
}

// bindIndex_Interface binds a generic wrapper to an already deployed contract.
func bindIndex_Interface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Index_InterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Index_Interface *Index_InterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Index_Interface.Contract.Index_InterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Index_Interface *Index_InterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Index_Interface.Contract.Index_InterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Index_Interface *Index_InterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Index_Interface.Contract.Index_InterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Index_Interface *Index_InterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Index_Interface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Index_Interface *Index_InterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Index_Interface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Index_Interface *Index_InterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Index_Interface.Contract.contract.Transact(opts, method, params...)
}

// LifToken is a free data retrieval call binding the contract method 0x554d8b37.
//
// Solidity: function LifToken() constant returns(address)
func (_Index_Interface *Index_InterfaceCaller) LifToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Index_Interface.contract.Call(opts, out, "LifToken")
	return *ret0, err
}

// LifToken is a free data retrieval call binding the contract method 0x554d8b37.
//
// Solidity: function LifToken() constant returns(address)
func (_Index_Interface *Index_InterfaceSession) LifToken() (common.Address, error) {
	return _Index_Interface.Contract.LifToken(&_Index_Interface.CallOpts)
}

// LifToken is a free data retrieval call binding the contract method 0x554d8b37.
//
// Solidity: function LifToken() constant returns(address)
func (_Index_Interface *Index_InterfaceCallerSession) LifToken() (common.Address, error) {
	return _Index_Interface.Contract.LifToken(&_Index_Interface.CallOpts)
}
