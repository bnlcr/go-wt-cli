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

// Unit_InterfaceABI is the input ABI used to generate the binding from.
const Unit_InterfaceABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"active\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_currencyCode\",\"type\":\"bytes8\"}],\"name\":\"setCurrencyCode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"fromDay\",\"type\":\"uint256\"},{\"name\":\"daysAmount\",\"type\":\"uint256\"}],\"name\":\"setSpecialPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setDefaultPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"defaultLifTokenPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unitType\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"fromDay\",\"type\":\"uint256\"},{\"name\":\"daysAmount\",\"type\":\"uint256\"}],\"name\":\"getLifCost\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"fromDay\",\"type\":\"uint256\"},{\"name\":\"daysAmount\",\"type\":\"uint256\"}],\"name\":\"getCost\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"setActive\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setDefaultLifPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"fromDay\",\"type\":\"uint256\"},{\"name\":\"daysAmount\",\"type\":\"uint256\"}],\"name\":\"book\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"day\",\"type\":\"uint256\"}],\"name\":\"getReservation\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"fromDay\",\"type\":\"uint256\"},{\"name\":\"daysAmount\",\"type\":\"uint256\"}],\"name\":\"setSpecialLifPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"fromDay\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"daysAmount\",\"type\":\"uint256\"}],\"name\":\"Book\",\"type\":\"event\"}]"

// Unit_InterfaceBin is the compiled bytecode used for deploying new contracts.
const Unit_InterfaceBin = `0x`

// DeployUnit_Interface deploys a new Ethereum contract, binding an instance of Unit_Interface to it.
func DeployUnit_Interface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Unit_Interface, error) {
	parsed, err := abi.JSON(strings.NewReader(Unit_InterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Unit_InterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Unit_Interface{Unit_InterfaceCaller: Unit_InterfaceCaller{contract: contract}, Unit_InterfaceTransactor: Unit_InterfaceTransactor{contract: contract}}, nil
}

// Unit_Interface is an auto generated Go binding around an Ethereum contract.
type Unit_Interface struct {
	Unit_InterfaceCaller     // Read-only binding to the contract
	Unit_InterfaceTransactor // Write-only binding to the contract
}

// Unit_InterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type Unit_InterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Unit_InterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Unit_InterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Unit_InterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Unit_InterfaceSession struct {
	Contract     *Unit_Interface   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Unit_InterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Unit_InterfaceCallerSession struct {
	Contract *Unit_InterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// Unit_InterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Unit_InterfaceTransactorSession struct {
	Contract     *Unit_InterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// Unit_InterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type Unit_InterfaceRaw struct {
	Contract *Unit_Interface // Generic contract binding to access the raw methods on
}

// Unit_InterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Unit_InterfaceCallerRaw struct {
	Contract *Unit_InterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// Unit_InterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Unit_InterfaceTransactorRaw struct {
	Contract *Unit_InterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnit_Interface creates a new instance of Unit_Interface, bound to a specific deployed contract.
func NewUnit_Interface(address common.Address, backend bind.ContractBackend) (*Unit_Interface, error) {
	contract, err := bindUnit_Interface(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Unit_Interface{Unit_InterfaceCaller: Unit_InterfaceCaller{contract: contract}, Unit_InterfaceTransactor: Unit_InterfaceTransactor{contract: contract}}, nil
}

// NewUnit_InterfaceCaller creates a new read-only instance of Unit_Interface, bound to a specific deployed contract.
func NewUnit_InterfaceCaller(address common.Address, caller bind.ContractCaller) (*Unit_InterfaceCaller, error) {
	contract, err := bindUnit_Interface(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &Unit_InterfaceCaller{contract: contract}, nil
}

// NewUnit_InterfaceTransactor creates a new write-only instance of Unit_Interface, bound to a specific deployed contract.
func NewUnit_InterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*Unit_InterfaceTransactor, error) {
	contract, err := bindUnit_Interface(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &Unit_InterfaceTransactor{contract: contract}, nil
}

// bindUnit_Interface binds a generic wrapper to an already deployed contract.
func bindUnit_Interface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Unit_InterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Unit_Interface *Unit_InterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Unit_Interface.Contract.Unit_InterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Unit_Interface *Unit_InterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Unit_Interface.Contract.Unit_InterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Unit_Interface *Unit_InterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Unit_Interface.Contract.Unit_InterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Unit_Interface *Unit_InterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Unit_Interface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Unit_Interface *Unit_InterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Unit_Interface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Unit_Interface *Unit_InterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Unit_Interface.Contract.contract.Transact(opts, method, params...)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() constant returns(bool)
func (_Unit_Interface *Unit_InterfaceCaller) Active(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Unit_Interface.contract.Call(opts, out, "active")
	return *ret0, err
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() constant returns(bool)
func (_Unit_Interface *Unit_InterfaceSession) Active() (bool, error) {
	return _Unit_Interface.Contract.Active(&_Unit_Interface.CallOpts)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() constant returns(bool)
func (_Unit_Interface *Unit_InterfaceCallerSession) Active() (bool, error) {
	return _Unit_Interface.Contract.Active(&_Unit_Interface.CallOpts)
}

// DefaultLifTokenPrice is a free data retrieval call binding the contract method 0x7d40192d.
//
// Solidity: function defaultLifTokenPrice() constant returns(uint256)
func (_Unit_Interface *Unit_InterfaceCaller) DefaultLifTokenPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Unit_Interface.contract.Call(opts, out, "defaultLifTokenPrice")
	return *ret0, err
}

// DefaultLifTokenPrice is a free data retrieval call binding the contract method 0x7d40192d.
//
// Solidity: function defaultLifTokenPrice() constant returns(uint256)
func (_Unit_Interface *Unit_InterfaceSession) DefaultLifTokenPrice() (*big.Int, error) {
	return _Unit_Interface.Contract.DefaultLifTokenPrice(&_Unit_Interface.CallOpts)
}

// DefaultLifTokenPrice is a free data retrieval call binding the contract method 0x7d40192d.
//
// Solidity: function defaultLifTokenPrice() constant returns(uint256)
func (_Unit_Interface *Unit_InterfaceCallerSession) DefaultLifTokenPrice() (*big.Int, error) {
	return _Unit_Interface.Contract.DefaultLifTokenPrice(&_Unit_Interface.CallOpts)
}

// GetCost is a free data retrieval call binding the contract method 0xa7c18c14.
//
// Solidity: function getCost(fromDay uint256, daysAmount uint256) constant returns(uint256)
func (_Unit_Interface *Unit_InterfaceCaller) GetCost(opts *bind.CallOpts, fromDay *big.Int, daysAmount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Unit_Interface.contract.Call(opts, out, "getCost", fromDay, daysAmount)
	return *ret0, err
}

// GetCost is a free data retrieval call binding the contract method 0xa7c18c14.
//
// Solidity: function getCost(fromDay uint256, daysAmount uint256) constant returns(uint256)
func (_Unit_Interface *Unit_InterfaceSession) GetCost(fromDay *big.Int, daysAmount *big.Int) (*big.Int, error) {
	return _Unit_Interface.Contract.GetCost(&_Unit_Interface.CallOpts, fromDay, daysAmount)
}

// GetCost is a free data retrieval call binding the contract method 0xa7c18c14.
//
// Solidity: function getCost(fromDay uint256, daysAmount uint256) constant returns(uint256)
func (_Unit_Interface *Unit_InterfaceCallerSession) GetCost(fromDay *big.Int, daysAmount *big.Int) (*big.Int, error) {
	return _Unit_Interface.Contract.GetCost(&_Unit_Interface.CallOpts, fromDay, daysAmount)
}

// GetLifCost is a free data retrieval call binding the contract method 0x9757876b.
//
// Solidity: function getLifCost(fromDay uint256, daysAmount uint256) constant returns(uint256)
func (_Unit_Interface *Unit_InterfaceCaller) GetLifCost(opts *bind.CallOpts, fromDay *big.Int, daysAmount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Unit_Interface.contract.Call(opts, out, "getLifCost", fromDay, daysAmount)
	return *ret0, err
}

// GetLifCost is a free data retrieval call binding the contract method 0x9757876b.
//
// Solidity: function getLifCost(fromDay uint256, daysAmount uint256) constant returns(uint256)
func (_Unit_Interface *Unit_InterfaceSession) GetLifCost(fromDay *big.Int, daysAmount *big.Int) (*big.Int, error) {
	return _Unit_Interface.Contract.GetLifCost(&_Unit_Interface.CallOpts, fromDay, daysAmount)
}

// GetLifCost is a free data retrieval call binding the contract method 0x9757876b.
//
// Solidity: function getLifCost(fromDay uint256, daysAmount uint256) constant returns(uint256)
func (_Unit_Interface *Unit_InterfaceCallerSession) GetLifCost(fromDay *big.Int, daysAmount *big.Int) (*big.Int, error) {
	return _Unit_Interface.Contract.GetLifCost(&_Unit_Interface.CallOpts, fromDay, daysAmount)
}

// GetReservation is a free data retrieval call binding the contract method 0xd57763c3.
//
// Solidity: function getReservation(day uint256) constant returns(uint256, uint256, address)
func (_Unit_Interface *Unit_InterfaceCaller) GetReservation(opts *bind.CallOpts, day *big.Int) (*big.Int, *big.Int, common.Address, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Unit_Interface.contract.Call(opts, out, "getReservation", day)
	return *ret0, *ret1, *ret2, err
}

// GetReservation is a free data retrieval call binding the contract method 0xd57763c3.
//
// Solidity: function getReservation(day uint256) constant returns(uint256, uint256, address)
func (_Unit_Interface *Unit_InterfaceSession) GetReservation(day *big.Int) (*big.Int, *big.Int, common.Address, error) {
	return _Unit_Interface.Contract.GetReservation(&_Unit_Interface.CallOpts, day)
}

// GetReservation is a free data retrieval call binding the contract method 0xd57763c3.
//
// Solidity: function getReservation(day uint256) constant returns(uint256, uint256, address)
func (_Unit_Interface *Unit_InterfaceCallerSession) GetReservation(day *big.Int) (*big.Int, *big.Int, common.Address, error) {
	return _Unit_Interface.Contract.GetReservation(&_Unit_Interface.CallOpts, day)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Unit_Interface *Unit_InterfaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Unit_Interface.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Unit_Interface *Unit_InterfaceSession) Owner() (common.Address, error) {
	return _Unit_Interface.Contract.Owner(&_Unit_Interface.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Unit_Interface *Unit_InterfaceCallerSession) Owner() (common.Address, error) {
	return _Unit_Interface.Contract.Owner(&_Unit_Interface.CallOpts)
}

// UnitType is a free data retrieval call binding the contract method 0x85ad4f90.
//
// Solidity: function unitType() constant returns(bytes32)
func (_Unit_Interface *Unit_InterfaceCaller) UnitType(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Unit_Interface.contract.Call(opts, out, "unitType")
	return *ret0, err
}

// UnitType is a free data retrieval call binding the contract method 0x85ad4f90.
//
// Solidity: function unitType() constant returns(bytes32)
func (_Unit_Interface *Unit_InterfaceSession) UnitType() ([32]byte, error) {
	return _Unit_Interface.Contract.UnitType(&_Unit_Interface.CallOpts)
}

// UnitType is a free data retrieval call binding the contract method 0x85ad4f90.
//
// Solidity: function unitType() constant returns(bytes32)
func (_Unit_Interface *Unit_InterfaceCallerSession) UnitType() ([32]byte, error) {
	return _Unit_Interface.Contract.UnitType(&_Unit_Interface.CallOpts)
}

// Book is a paid mutator transaction binding the contract method 0xc78e687c.
//
// Solidity: function book(from address, fromDay uint256, daysAmount uint256) returns(bool)
func (_Unit_Interface *Unit_InterfaceTransactor) Book(opts *bind.TransactOpts, from common.Address, fromDay *big.Int, daysAmount *big.Int) (*types.Transaction, error) {
	return _Unit_Interface.contract.Transact(opts, "book", from, fromDay, daysAmount)
}

// Book is a paid mutator transaction binding the contract method 0xc78e687c.
//
// Solidity: function book(from address, fromDay uint256, daysAmount uint256) returns(bool)
func (_Unit_Interface *Unit_InterfaceSession) Book(from common.Address, fromDay *big.Int, daysAmount *big.Int) (*types.Transaction, error) {
	return _Unit_Interface.Contract.Book(&_Unit_Interface.TransactOpts, from, fromDay, daysAmount)
}

// Book is a paid mutator transaction binding the contract method 0xc78e687c.
//
// Solidity: function book(from address, fromDay uint256, daysAmount uint256) returns(bool)
func (_Unit_Interface *Unit_InterfaceTransactorSession) Book(from common.Address, fromDay *big.Int, daysAmount *big.Int) (*types.Transaction, error) {
	return _Unit_Interface.Contract.Book(&_Unit_Interface.TransactOpts, from, fromDay, daysAmount)
}

// SetActive is a paid mutator transaction binding the contract method 0xacec338a.
//
// Solidity: function setActive(_active bool) returns()
func (_Unit_Interface *Unit_InterfaceTransactor) SetActive(opts *bind.TransactOpts, _active bool) (*types.Transaction, error) {
	return _Unit_Interface.contract.Transact(opts, "setActive", _active)
}

// SetActive is a paid mutator transaction binding the contract method 0xacec338a.
//
// Solidity: function setActive(_active bool) returns()
func (_Unit_Interface *Unit_InterfaceSession) SetActive(_active bool) (*types.Transaction, error) {
	return _Unit_Interface.Contract.SetActive(&_Unit_Interface.TransactOpts, _active)
}

// SetActive is a paid mutator transaction binding the contract method 0xacec338a.
//
// Solidity: function setActive(_active bool) returns()
func (_Unit_Interface *Unit_InterfaceTransactorSession) SetActive(_active bool) (*types.Transaction, error) {
	return _Unit_Interface.Contract.SetActive(&_Unit_Interface.TransactOpts, _active)
}

// SetCurrencyCode is a paid mutator transaction binding the contract method 0x2914fc0a.
//
// Solidity: function setCurrencyCode(_currencyCode bytes8) returns()
func (_Unit_Interface *Unit_InterfaceTransactor) SetCurrencyCode(opts *bind.TransactOpts, _currencyCode [8]byte) (*types.Transaction, error) {
	return _Unit_Interface.contract.Transact(opts, "setCurrencyCode", _currencyCode)
}

// SetCurrencyCode is a paid mutator transaction binding the contract method 0x2914fc0a.
//
// Solidity: function setCurrencyCode(_currencyCode bytes8) returns()
func (_Unit_Interface *Unit_InterfaceSession) SetCurrencyCode(_currencyCode [8]byte) (*types.Transaction, error) {
	return _Unit_Interface.Contract.SetCurrencyCode(&_Unit_Interface.TransactOpts, _currencyCode)
}

// SetCurrencyCode is a paid mutator transaction binding the contract method 0x2914fc0a.
//
// Solidity: function setCurrencyCode(_currencyCode bytes8) returns()
func (_Unit_Interface *Unit_InterfaceTransactorSession) SetCurrencyCode(_currencyCode [8]byte) (*types.Transaction, error) {
	return _Unit_Interface.Contract.SetCurrencyCode(&_Unit_Interface.TransactOpts, _currencyCode)
}

// SetDefaultLifPrice is a paid mutator transaction binding the contract method 0xc55b2ee7.
//
// Solidity: function setDefaultLifPrice(price uint256) returns()
func (_Unit_Interface *Unit_InterfaceTransactor) SetDefaultLifPrice(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _Unit_Interface.contract.Transact(opts, "setDefaultLifPrice", price)
}

// SetDefaultLifPrice is a paid mutator transaction binding the contract method 0xc55b2ee7.
//
// Solidity: function setDefaultLifPrice(price uint256) returns()
func (_Unit_Interface *Unit_InterfaceSession) SetDefaultLifPrice(price *big.Int) (*types.Transaction, error) {
	return _Unit_Interface.Contract.SetDefaultLifPrice(&_Unit_Interface.TransactOpts, price)
}

// SetDefaultLifPrice is a paid mutator transaction binding the contract method 0xc55b2ee7.
//
// Solidity: function setDefaultLifPrice(price uint256) returns()
func (_Unit_Interface *Unit_InterfaceTransactorSession) SetDefaultLifPrice(price *big.Int) (*types.Transaction, error) {
	return _Unit_Interface.Contract.SetDefaultLifPrice(&_Unit_Interface.TransactOpts, price)
}

// SetDefaultPrice is a paid mutator transaction binding the contract method 0x6d3c7ec5.
//
// Solidity: function setDefaultPrice(price uint256) returns()
func (_Unit_Interface *Unit_InterfaceTransactor) SetDefaultPrice(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _Unit_Interface.contract.Transact(opts, "setDefaultPrice", price)
}

// SetDefaultPrice is a paid mutator transaction binding the contract method 0x6d3c7ec5.
//
// Solidity: function setDefaultPrice(price uint256) returns()
func (_Unit_Interface *Unit_InterfaceSession) SetDefaultPrice(price *big.Int) (*types.Transaction, error) {
	return _Unit_Interface.Contract.SetDefaultPrice(&_Unit_Interface.TransactOpts, price)
}

// SetDefaultPrice is a paid mutator transaction binding the contract method 0x6d3c7ec5.
//
// Solidity: function setDefaultPrice(price uint256) returns()
func (_Unit_Interface *Unit_InterfaceTransactorSession) SetDefaultPrice(price *big.Int) (*types.Transaction, error) {
	return _Unit_Interface.Contract.SetDefaultPrice(&_Unit_Interface.TransactOpts, price)
}

// SetSpecialLifPrice is a paid mutator transaction binding the contract method 0xfd480e5f.
//
// Solidity: function setSpecialLifPrice(price uint256, fromDay uint256, daysAmount uint256) returns()
func (_Unit_Interface *Unit_InterfaceTransactor) SetSpecialLifPrice(opts *bind.TransactOpts, price *big.Int, fromDay *big.Int, daysAmount *big.Int) (*types.Transaction, error) {
	return _Unit_Interface.contract.Transact(opts, "setSpecialLifPrice", price, fromDay, daysAmount)
}

// SetSpecialLifPrice is a paid mutator transaction binding the contract method 0xfd480e5f.
//
// Solidity: function setSpecialLifPrice(price uint256, fromDay uint256, daysAmount uint256) returns()
func (_Unit_Interface *Unit_InterfaceSession) SetSpecialLifPrice(price *big.Int, fromDay *big.Int, daysAmount *big.Int) (*types.Transaction, error) {
	return _Unit_Interface.Contract.SetSpecialLifPrice(&_Unit_Interface.TransactOpts, price, fromDay, daysAmount)
}

// SetSpecialLifPrice is a paid mutator transaction binding the contract method 0xfd480e5f.
//
// Solidity: function setSpecialLifPrice(price uint256, fromDay uint256, daysAmount uint256) returns()
func (_Unit_Interface *Unit_InterfaceTransactorSession) SetSpecialLifPrice(price *big.Int, fromDay *big.Int, daysAmount *big.Int) (*types.Transaction, error) {
	return _Unit_Interface.Contract.SetSpecialLifPrice(&_Unit_Interface.TransactOpts, price, fromDay, daysAmount)
}

// SetSpecialPrice is a paid mutator transaction binding the contract method 0x46a8e678.
//
// Solidity: function setSpecialPrice(price uint256, fromDay uint256, daysAmount uint256) returns()
func (_Unit_Interface *Unit_InterfaceTransactor) SetSpecialPrice(opts *bind.TransactOpts, price *big.Int, fromDay *big.Int, daysAmount *big.Int) (*types.Transaction, error) {
	return _Unit_Interface.contract.Transact(opts, "setSpecialPrice", price, fromDay, daysAmount)
}

// SetSpecialPrice is a paid mutator transaction binding the contract method 0x46a8e678.
//
// Solidity: function setSpecialPrice(price uint256, fromDay uint256, daysAmount uint256) returns()
func (_Unit_Interface *Unit_InterfaceSession) SetSpecialPrice(price *big.Int, fromDay *big.Int, daysAmount *big.Int) (*types.Transaction, error) {
	return _Unit_Interface.Contract.SetSpecialPrice(&_Unit_Interface.TransactOpts, price, fromDay, daysAmount)
}

// SetSpecialPrice is a paid mutator transaction binding the contract method 0x46a8e678.
//
// Solidity: function setSpecialPrice(price uint256, fromDay uint256, daysAmount uint256) returns()
func (_Unit_Interface *Unit_InterfaceTransactorSession) SetSpecialPrice(price *big.Int, fromDay *big.Int, daysAmount *big.Int) (*types.Transaction, error) {
	return _Unit_Interface.Contract.SetSpecialPrice(&_Unit_Interface.TransactOpts, price, fromDay, daysAmount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Unit_Interface *Unit_InterfaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Unit_Interface.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Unit_Interface *Unit_InterfaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Unit_Interface.Contract.TransferOwnership(&_Unit_Interface.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Unit_Interface *Unit_InterfaceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Unit_Interface.Contract.TransferOwnership(&_Unit_Interface.TransactOpts, newOwner)
}
