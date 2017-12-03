// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gen

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// SmartTokenABI is the input ABI used to generate the binding from.
const SmartTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"approveData\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"transferData\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"transferDataFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// SmartTokenBin is the compiled bytecode used for deploying new contracts.
const SmartTokenBin = `0x6060604052341561000f57600080fd5b610e6b8061001e6000396000f300606060405260043610610099576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063095ea7b31461009e57806318160ddd146100f857806323b872dd146101215780636ef3ef7e1461019a57806370a0823114610237578063a9059cbb14610284578063c0e37b15146102de578063dd62ed3e1461037b578063efef445b146103e7575b600080fd5b34156100a957600080fd5b6100de600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919080359060200190919050506104a3565b604051808215151515815260200191505060405180910390f35b341561010357600080fd5b61010b61062a565b6040518082815260200191505060405180910390f35b341561012c57600080fd5b610180600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff16906020019091908035906020019091905050610630565b604051808215151515815260200191505060405180910390f35b34156101a557600080fd5b61021d600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919050506108e0565b604051808215151515815260200191505060405180910390f35b341561024257600080fd5b61026e600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506109c9565b6040518082815260200191505060405180910390f35b341561028f57600080fd5b6102c4600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091908035906020019091905050610a12565b604051808215151515815260200191505060405180910390f35b34156102e957600080fd5b610361600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610bad565b604051808215151515815260200191505060405180910390f35b341561038657600080fd5b6103d1600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610c96565b6040518082815260200191505060405180910390f35b34156103f257600080fd5b610489600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610d1d565b604051808215151515815260200191505060405180910390f35b60008082148061052f57506000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054145b151561053a57600080fd5b81600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a36001905092915050565b60005481565b600080600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905061070483600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610e0890919063ffffffff16565b600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061079983600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610e2690919063ffffffff16565b600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506107ef8382610e2690919063ffffffff16565b600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef856040518082815260200191505060405180910390a360019150509392505050565b60003073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415151561091d57600080fd5b61092784846104a3565b508373ffffffffffffffffffffffffffffffffffffffff168260405180828051906020019080838360005b8381101561096d578082015181840152602081019050610952565b50505050905090810190601f16801561099a5780820380516001836020036101000a031916815260200191505b5091505060006040518083038160008661646e5a03f191505015156109be57600080fd5b600190509392505050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000610a6682600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610e2690919063ffffffff16565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610afb82600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610e0890919063ffffffff16565b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a36001905092915050565b60003073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614151515610bea57600080fd5b8373ffffffffffffffffffffffffffffffffffffffff168260405180828051906020019080838360005b83811015610c2f578082015181840152602081019050610c14565b50505050905090810190601f168015610c5c5780820380516001836020036101000a031916815260200191505b5091505060006040518083038160008661646e5a03f19150501515610c8057600080fd5b610c8a8484610a12565b50600190509392505050565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b60003073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614151515610d5a57600080fd5b8373ffffffffffffffffffffffffffffffffffffffff168260405180828051906020019080838360005b83811015610d9f578082015181840152602081019050610d84565b50505050905090810190601f168015610dcc5780820380516001836020036101000a031916815260200191505b5091505060006040518083038160008661646e5a03f19150501515610df057600080fd5b610dfb858585610630565b5060019050949350505050565b6000808284019050838110151515610e1c57fe5b8091505092915050565b6000828211151515610e3457fe5b8183039050929150505600a165627a7a723058203830048504390fab9fdb2cb64fc9269b6a1f04b3d133949bd82a2062dc9faa7d0029`

// DeploySmartToken deploys a new Ethereum contract, binding an instance of SmartToken to it.
func DeploySmartToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SmartToken, error) {
	parsed, err := abi.JSON(strings.NewReader(SmartTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SmartTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SmartToken{SmartTokenCaller: SmartTokenCaller{contract: contract}, SmartTokenTransactor: SmartTokenTransactor{contract: contract}}, nil
}

// SmartToken is an auto generated Go binding around an Ethereum contract.
type SmartToken struct {
	SmartTokenCaller     // Read-only binding to the contract
	SmartTokenTransactor // Write-only binding to the contract
}

// SmartTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type SmartTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SmartTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SmartTokenSession struct {
	Contract     *SmartToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SmartTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SmartTokenCallerSession struct {
	Contract *SmartTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SmartTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SmartTokenTransactorSession struct {
	Contract     *SmartTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SmartTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type SmartTokenRaw struct {
	Contract *SmartToken // Generic contract binding to access the raw methods on
}

// SmartTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SmartTokenCallerRaw struct {
	Contract *SmartTokenCaller // Generic read-only contract binding to access the raw methods on
}

// SmartTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SmartTokenTransactorRaw struct {
	Contract *SmartTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSmartToken creates a new instance of SmartToken, bound to a specific deployed contract.
func NewSmartToken(address common.Address, backend bind.ContractBackend) (*SmartToken, error) {
	contract, err := bindSmartToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SmartToken{SmartTokenCaller: SmartTokenCaller{contract: contract}, SmartTokenTransactor: SmartTokenTransactor{contract: contract}}, nil
}

// NewSmartTokenCaller creates a new read-only instance of SmartToken, bound to a specific deployed contract.
func NewSmartTokenCaller(address common.Address, caller bind.ContractCaller) (*SmartTokenCaller, error) {
	contract, err := bindSmartToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SmartTokenCaller{contract: contract}, nil
}

// NewSmartTokenTransactor creates a new write-only instance of SmartToken, bound to a specific deployed contract.
func NewSmartTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*SmartTokenTransactor, error) {
	contract, err := bindSmartToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SmartTokenTransactor{contract: contract}, nil
}

// bindSmartToken binds a generic wrapper to an already deployed contract.
func bindSmartToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SmartTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmartToken *SmartTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SmartToken.Contract.SmartTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmartToken *SmartTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartToken.Contract.SmartTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmartToken *SmartTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmartToken.Contract.SmartTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmartToken *SmartTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SmartToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmartToken *SmartTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmartToken *SmartTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmartToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_SmartToken *SmartTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SmartToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_SmartToken *SmartTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _SmartToken.Contract.Allowance(&_SmartToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_SmartToken *SmartTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _SmartToken.Contract.Allowance(&_SmartToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_SmartToken *SmartTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SmartToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_SmartToken *SmartTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _SmartToken.Contract.BalanceOf(&_SmartToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_SmartToken *SmartTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _SmartToken.Contract.BalanceOf(&_SmartToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SmartToken *SmartTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SmartToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SmartToken *SmartTokenSession) TotalSupply() (*big.Int, error) {
	return _SmartToken.Contract.TotalSupply(&_SmartToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SmartToken *SmartTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _SmartToken.Contract.TotalSupply(&_SmartToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_SmartToken *SmartTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SmartToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_SmartToken *SmartTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SmartToken.Contract.Approve(&_SmartToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_SmartToken *SmartTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SmartToken.Contract.Approve(&_SmartToken.TransactOpts, _spender, _value)
}

// ApproveData is a paid mutator transaction binding the contract method 0x6ef3ef7e.
//
// Solidity: function approveData(_spender address, _value uint256, _data bytes) returns(bool)
func (_SmartToken *SmartTokenTransactor) ApproveData(opts *bind.TransactOpts, _spender common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _SmartToken.contract.Transact(opts, "approveData", _spender, _value, _data)
}

// ApproveData is a paid mutator transaction binding the contract method 0x6ef3ef7e.
//
// Solidity: function approveData(_spender address, _value uint256, _data bytes) returns(bool)
func (_SmartToken *SmartTokenSession) ApproveData(_spender common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _SmartToken.Contract.ApproveData(&_SmartToken.TransactOpts, _spender, _value, _data)
}

// ApproveData is a paid mutator transaction binding the contract method 0x6ef3ef7e.
//
// Solidity: function approveData(_spender address, _value uint256, _data bytes) returns(bool)
func (_SmartToken *SmartTokenTransactorSession) ApproveData(_spender common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _SmartToken.Contract.ApproveData(&_SmartToken.TransactOpts, _spender, _value, _data)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_SmartToken *SmartTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SmartToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_SmartToken *SmartTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SmartToken.Contract.Transfer(&_SmartToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_SmartToken *SmartTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SmartToken.Contract.Transfer(&_SmartToken.TransactOpts, _to, _value)
}

// TransferData is a paid mutator transaction binding the contract method 0xc0e37b15.
//
// Solidity: function transferData(_to address, _value uint256, _data bytes) returns(bool)
func (_SmartToken *SmartTokenTransactor) TransferData(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _SmartToken.contract.Transact(opts, "transferData", _to, _value, _data)
}

// TransferData is a paid mutator transaction binding the contract method 0xc0e37b15.
//
// Solidity: function transferData(_to address, _value uint256, _data bytes) returns(bool)
func (_SmartToken *SmartTokenSession) TransferData(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _SmartToken.Contract.TransferData(&_SmartToken.TransactOpts, _to, _value, _data)
}

// TransferData is a paid mutator transaction binding the contract method 0xc0e37b15.
//
// Solidity: function transferData(_to address, _value uint256, _data bytes) returns(bool)
func (_SmartToken *SmartTokenTransactorSession) TransferData(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _SmartToken.Contract.TransferData(&_SmartToken.TransactOpts, _to, _value, _data)
}

// TransferDataFrom is a paid mutator transaction binding the contract method 0xefef445b.
//
// Solidity: function transferDataFrom(_from address, _to address, _value uint256, _data bytes) returns(bool)
func (_SmartToken *SmartTokenTransactor) TransferDataFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _SmartToken.contract.Transact(opts, "transferDataFrom", _from, _to, _value, _data)
}

// TransferDataFrom is a paid mutator transaction binding the contract method 0xefef445b.
//
// Solidity: function transferDataFrom(_from address, _to address, _value uint256, _data bytes) returns(bool)
func (_SmartToken *SmartTokenSession) TransferDataFrom(_from common.Address, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _SmartToken.Contract.TransferDataFrom(&_SmartToken.TransactOpts, _from, _to, _value, _data)
}

// TransferDataFrom is a paid mutator transaction binding the contract method 0xefef445b.
//
// Solidity: function transferDataFrom(_from address, _to address, _value uint256, _data bytes) returns(bool)
func (_SmartToken *SmartTokenTransactorSession) TransferDataFrom(_from common.Address, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _SmartToken.Contract.TransferDataFrom(&_SmartToken.TransactOpts, _from, _to, _value, _data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_SmartToken *SmartTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SmartToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_SmartToken *SmartTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SmartToken.Contract.TransferFrom(&_SmartToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_SmartToken *SmartTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SmartToken.Contract.TransferFrom(&_SmartToken.TransactOpts, _from, _to, _value)
}
