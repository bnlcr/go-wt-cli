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

// Hotel_InterfaceABI is the input ABI used to generate the binding from.
const Hotel_InterfaceABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"zip\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_description\",\"type\":\"string\"}],\"name\":\"editInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"url\",\"type\":\"string\"}],\"name\":\"addImage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"unitAddress\",\"type\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"callUnit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"publicCallData\",\"type\":\"bytes\"},{\"name\":\"privateData\",\"type\":\"bytes\"}],\"name\":\"beginCall\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"unitTypeNames\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lineOne\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"unitTypes\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getUnitTypeNames\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"created\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pendingCalls\",\"outputs\":[{\"name\":\"callData\",\"type\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"},{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"unitType\",\"type\":\"bytes32\"},{\"name\":\"newAddr\",\"type\":\"address\"}],\"name\":\"changeUnitType\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"msgDataHash\",\"type\":\"bytes32\"}],\"name\":\"continueCall\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_waitConfirmation\",\"type\":\"bool\"}],\"name\":\"changeConfirmation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"latitude\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"longitude\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"unitsIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getImagesLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"unitAddress\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"fromDay\",\"type\":\"uint256\"},{\"name\":\"daysAmount\",\"type\":\"uint256\"}],\"name\":\"bookWithLif\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"unitType\",\"type\":\"bytes32\"}],\"name\":\"getUnitType\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_lineOne\",\"type\":\"string\"},{\"name\":\"_lineTwo\",\"type\":\"string\"},{\"name\":\"_zip\",\"type\":\"string\"},{\"name\":\"_country\",\"type\":\"string\"}],\"name\":\"editAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"images\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeImage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timezone\",\"type\":\"uint256\"},{\"name\":\"_longitude\",\"type\":\"uint256\"},{\"name\":\"_latitude\",\"type\":\"uint256\"}],\"name\":\"editLocation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"unit\",\"type\":\"address\"}],\"name\":\"addUnit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"waitConfirmation\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"unitAddress\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"fromDay\",\"type\":\"uint256\"},{\"name\":\"daysAmount\",\"type\":\"uint256\"}],\"name\":\"book\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"timezone\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractType\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"unitType\",\"type\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeUnitType\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"country\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"unitType\",\"type\":\"bytes32\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"callUnitType\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addUnitType\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"unit\",\"type\":\"address\"}],\"name\":\"removeUnit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"units\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lineTwo\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"unit\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"fromDay\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"daysAmount\",\"type\":\"uint256\"}],\"name\":\"Book\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"CallStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"CallFinish\",\"type\":\"event\"}]"

// Hotel_InterfaceBin is the compiled bytecode used for deploying new contracts.
const Hotel_InterfaceBin = `0x`

// DeployHotel_Interface deploys a new Ethereum contract, binding an instance of Hotel_Interface to it.
func DeployHotel_Interface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Hotel_Interface, error) {
	parsed, err := abi.JSON(strings.NewReader(Hotel_InterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Hotel_InterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Hotel_Interface{Hotel_InterfaceCaller: Hotel_InterfaceCaller{contract: contract}, Hotel_InterfaceTransactor: Hotel_InterfaceTransactor{contract: contract}}, nil
}

// Hotel_Interface is an auto generated Go binding around an Ethereum contract.
type Hotel_Interface struct {
	Hotel_InterfaceCaller     // Read-only binding to the contract
	Hotel_InterfaceTransactor // Write-only binding to the contract
}

// Hotel_InterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type Hotel_InterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Hotel_InterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Hotel_InterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Hotel_InterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Hotel_InterfaceSession struct {
	Contract     *Hotel_Interface  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Hotel_InterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Hotel_InterfaceCallerSession struct {
	Contract *Hotel_InterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// Hotel_InterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Hotel_InterfaceTransactorSession struct {
	Contract     *Hotel_InterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// Hotel_InterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type Hotel_InterfaceRaw struct {
	Contract *Hotel_Interface // Generic contract binding to access the raw methods on
}

// Hotel_InterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Hotel_InterfaceCallerRaw struct {
	Contract *Hotel_InterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// Hotel_InterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Hotel_InterfaceTransactorRaw struct {
	Contract *Hotel_InterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHotel_Interface creates a new instance of Hotel_Interface, bound to a specific deployed contract.
func NewHotel_Interface(address common.Address, backend bind.ContractBackend) (*Hotel_Interface, error) {
	contract, err := bindHotel_Interface(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hotel_Interface{Hotel_InterfaceCaller: Hotel_InterfaceCaller{contract: contract}, Hotel_InterfaceTransactor: Hotel_InterfaceTransactor{contract: contract}}, nil
}

// NewHotel_InterfaceCaller creates a new read-only instance of Hotel_Interface, bound to a specific deployed contract.
func NewHotel_InterfaceCaller(address common.Address, caller bind.ContractCaller) (*Hotel_InterfaceCaller, error) {
	contract, err := bindHotel_Interface(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &Hotel_InterfaceCaller{contract: contract}, nil
}

// NewHotel_InterfaceTransactor creates a new write-only instance of Hotel_Interface, bound to a specific deployed contract.
func NewHotel_InterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*Hotel_InterfaceTransactor, error) {
	contract, err := bindHotel_Interface(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &Hotel_InterfaceTransactor{contract: contract}, nil
}

// bindHotel_Interface binds a generic wrapper to an already deployed contract.
func bindHotel_Interface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Hotel_InterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hotel_Interface *Hotel_InterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Hotel_Interface.Contract.Hotel_InterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hotel_Interface *Hotel_InterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.Hotel_InterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hotel_Interface *Hotel_InterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.Hotel_InterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hotel_Interface *Hotel_InterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Hotel_Interface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hotel_Interface *Hotel_InterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hotel_Interface *Hotel_InterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.contract.Transact(opts, method, params...)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() constant returns(bytes32)
func (_Hotel_Interface *Hotel_InterfaceCaller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Hotel_Interface.contract.Call(opts, out, "contractType")
	return *ret0, err
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() constant returns(bytes32)
func (_Hotel_Interface *Hotel_InterfaceSession) ContractType() ([32]byte, error) {
	return _Hotel_Interface.Contract.ContractType(&_Hotel_Interface.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() constant returns(bytes32)
func (_Hotel_Interface *Hotel_InterfaceCallerSession) ContractType() ([32]byte, error) {
	return _Hotel_Interface.Contract.ContractType(&_Hotel_Interface.CallOpts)
}

// Country is a free data retrieval call binding the contract method 0xd8b0b499.
//
// Solidity: function country() constant returns(string)
func (_Hotel_Interface *Hotel_InterfaceCaller) Country(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Hotel_Interface.contract.Call(opts, out, "country")
	return *ret0, err
}

// Country is a free data retrieval call binding the contract method 0xd8b0b499.
//
// Solidity: function country() constant returns(string)
func (_Hotel_Interface *Hotel_InterfaceSession) Country() (string, error) {
	return _Hotel_Interface.Contract.Country(&_Hotel_Interface.CallOpts)
}

// Country is a free data retrieval call binding the contract method 0xd8b0b499.
//
// Solidity: function country() constant returns(string)
func (_Hotel_Interface *Hotel_InterfaceCallerSession) Country() (string, error) {
	return _Hotel_Interface.Contract.Country(&_Hotel_Interface.CallOpts)
}

// Created is a free data retrieval call binding the contract method 0x325a19f1.
//
// Solidity: function created() constant returns(uint256)
func (_Hotel_Interface *Hotel_InterfaceCaller) Created(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Hotel_Interface.contract.Call(opts, out, "created")
	return *ret0, err
}

// Created is a free data retrieval call binding the contract method 0x325a19f1.
//
// Solidity: function created() constant returns(uint256)
func (_Hotel_Interface *Hotel_InterfaceSession) Created() (*big.Int, error) {
	return _Hotel_Interface.Contract.Created(&_Hotel_Interface.CallOpts)
}

// Created is a free data retrieval call binding the contract method 0x325a19f1.
//
// Solidity: function created() constant returns(uint256)
func (_Hotel_Interface *Hotel_InterfaceCallerSession) Created() (*big.Int, error) {
	return _Hotel_Interface.Contract.Created(&_Hotel_Interface.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() constant returns(string)
func (_Hotel_Interface *Hotel_InterfaceCaller) Description(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Hotel_Interface.contract.Call(opts, out, "description")
	return *ret0, err
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() constant returns(string)
func (_Hotel_Interface *Hotel_InterfaceSession) Description() (string, error) {
	return _Hotel_Interface.Contract.Description(&_Hotel_Interface.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() constant returns(string)
func (_Hotel_Interface *Hotel_InterfaceCallerSession) Description() (string, error) {
	return _Hotel_Interface.Contract.Description(&_Hotel_Interface.CallOpts)
}

// GetImagesLength is a free data retrieval call binding the contract method 0x60839bd8.
//
// Solidity: function getImagesLength() constant returns(uint256)
func (_Hotel_Interface *Hotel_InterfaceCaller) GetImagesLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Hotel_Interface.contract.Call(opts, out, "getImagesLength")
	return *ret0, err
}

// GetImagesLength is a free data retrieval call binding the contract method 0x60839bd8.
//
// Solidity: function getImagesLength() constant returns(uint256)
func (_Hotel_Interface *Hotel_InterfaceSession) GetImagesLength() (*big.Int, error) {
	return _Hotel_Interface.Contract.GetImagesLength(&_Hotel_Interface.CallOpts)
}

// GetImagesLength is a free data retrieval call binding the contract method 0x60839bd8.
//
// Solidity: function getImagesLength() constant returns(uint256)
func (_Hotel_Interface *Hotel_InterfaceCallerSession) GetImagesLength() (*big.Int, error) {
	return _Hotel_Interface.Contract.GetImagesLength(&_Hotel_Interface.CallOpts)
}

// GetUnitType is a free data retrieval call binding the contract method 0x7e1d63c3.
//
// Solidity: function getUnitType(unitType bytes32) constant returns(address)
func (_Hotel_Interface *Hotel_InterfaceCaller) GetUnitType(opts *bind.CallOpts, unitType [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Hotel_Interface.contract.Call(opts, out, "getUnitType", unitType)
	return *ret0, err
}

// GetUnitType is a free data retrieval call binding the contract method 0x7e1d63c3.
//
// Solidity: function getUnitType(unitType bytes32) constant returns(address)
func (_Hotel_Interface *Hotel_InterfaceSession) GetUnitType(unitType [32]byte) (common.Address, error) {
	return _Hotel_Interface.Contract.GetUnitType(&_Hotel_Interface.CallOpts, unitType)
}

// GetUnitType is a free data retrieval call binding the contract method 0x7e1d63c3.
//
// Solidity: function getUnitType(unitType bytes32) constant returns(address)
func (_Hotel_Interface *Hotel_InterfaceCallerSession) GetUnitType(unitType [32]byte) (common.Address, error) {
	return _Hotel_Interface.Contract.GetUnitType(&_Hotel_Interface.CallOpts, unitType)
}

// GetUnitTypeNames is a free data retrieval call binding the contract method 0x2db25a38.
//
// Solidity: function getUnitTypeNames() constant returns(bytes32[])
func (_Hotel_Interface *Hotel_InterfaceCaller) GetUnitTypeNames(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _Hotel_Interface.contract.Call(opts, out, "getUnitTypeNames")
	return *ret0, err
}

// GetUnitTypeNames is a free data retrieval call binding the contract method 0x2db25a38.
//
// Solidity: function getUnitTypeNames() constant returns(bytes32[])
func (_Hotel_Interface *Hotel_InterfaceSession) GetUnitTypeNames() ([][32]byte, error) {
	return _Hotel_Interface.Contract.GetUnitTypeNames(&_Hotel_Interface.CallOpts)
}

// GetUnitTypeNames is a free data retrieval call binding the contract method 0x2db25a38.
//
// Solidity: function getUnitTypeNames() constant returns(bytes32[])
func (_Hotel_Interface *Hotel_InterfaceCallerSession) GetUnitTypeNames() ([][32]byte, error) {
	return _Hotel_Interface.Contract.GetUnitTypeNames(&_Hotel_Interface.CallOpts)
}

// Images is a free data retrieval call binding the contract method 0x84856482.
//
// Solidity: function images( uint256) constant returns(string)
func (_Hotel_Interface *Hotel_InterfaceCaller) Images(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Hotel_Interface.contract.Call(opts, out, "images", arg0)
	return *ret0, err
}

// Images is a free data retrieval call binding the contract method 0x84856482.
//
// Solidity: function images( uint256) constant returns(string)
func (_Hotel_Interface *Hotel_InterfaceSession) Images(arg0 *big.Int) (string, error) {
	return _Hotel_Interface.Contract.Images(&_Hotel_Interface.CallOpts, arg0)
}

// Images is a free data retrieval call binding the contract method 0x84856482.
//
// Solidity: function images( uint256) constant returns(string)
func (_Hotel_Interface *Hotel_InterfaceCallerSession) Images(arg0 *big.Int) (string, error) {
	return _Hotel_Interface.Contract.Images(&_Hotel_Interface.CallOpts, arg0)
}

// Latitude is a free data retrieval call binding the contract method 0x4fd7d76a.
//
// Solidity: function latitude() constant returns(uint256)
func (_Hotel_Interface *Hotel_InterfaceCaller) Latitude(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Hotel_Interface.contract.Call(opts, out, "latitude")
	return *ret0, err
}

// Latitude is a free data retrieval call binding the contract method 0x4fd7d76a.
//
// Solidity: function latitude() constant returns(uint256)
func (_Hotel_Interface *Hotel_InterfaceSession) Latitude() (*big.Int, error) {
	return _Hotel_Interface.Contract.Latitude(&_Hotel_Interface.CallOpts)
}

// Latitude is a free data retrieval call binding the contract method 0x4fd7d76a.
//
// Solidity: function latitude() constant returns(uint256)
func (_Hotel_Interface *Hotel_InterfaceCallerSession) Latitude() (*big.Int, error) {
	return _Hotel_Interface.Contract.Latitude(&_Hotel_Interface.CallOpts)
}

// LineOne is a free data retrieval call binding the contract method 0x2a323572.
//
// Solidity: function lineOne() constant returns(string)
func (_Hotel_Interface *Hotel_InterfaceCaller) LineOne(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Hotel_Interface.contract.Call(opts, out, "lineOne")
	return *ret0, err
}

// LineOne is a free data retrieval call binding the contract method 0x2a323572.
//
// Solidity: function lineOne() constant returns(string)
func (_Hotel_Interface *Hotel_InterfaceSession) LineOne() (string, error) {
	return _Hotel_Interface.Contract.LineOne(&_Hotel_Interface.CallOpts)
}

// LineOne is a free data retrieval call binding the contract method 0x2a323572.
//
// Solidity: function lineOne() constant returns(string)
func (_Hotel_Interface *Hotel_InterfaceCallerSession) LineOne() (string, error) {
	return _Hotel_Interface.Contract.LineOne(&_Hotel_Interface.CallOpts)
}

// LineTwo is a free data retrieval call binding the contract method 0xed7fa4f1.
//
// Solidity: function lineTwo() constant returns(string)
func (_Hotel_Interface *Hotel_InterfaceCaller) LineTwo(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Hotel_Interface.contract.Call(opts, out, "lineTwo")
	return *ret0, err
}

// LineTwo is a free data retrieval call binding the contract method 0xed7fa4f1.
//
// Solidity: function lineTwo() constant returns(string)
func (_Hotel_Interface *Hotel_InterfaceSession) LineTwo() (string, error) {
	return _Hotel_Interface.Contract.LineTwo(&_Hotel_Interface.CallOpts)
}

// LineTwo is a free data retrieval call binding the contract method 0xed7fa4f1.
//
// Solidity: function lineTwo() constant returns(string)
func (_Hotel_Interface *Hotel_InterfaceCallerSession) LineTwo() (string, error) {
	return _Hotel_Interface.Contract.LineTwo(&_Hotel_Interface.CallOpts)
}

// Longitude is a free data retrieval call binding the contract method 0x589af69c.
//
// Solidity: function longitude() constant returns(uint256)
func (_Hotel_Interface *Hotel_InterfaceCaller) Longitude(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Hotel_Interface.contract.Call(opts, out, "longitude")
	return *ret0, err
}

// Longitude is a free data retrieval call binding the contract method 0x589af69c.
//
// Solidity: function longitude() constant returns(uint256)
func (_Hotel_Interface *Hotel_InterfaceSession) Longitude() (*big.Int, error) {
	return _Hotel_Interface.Contract.Longitude(&_Hotel_Interface.CallOpts)
}

// Longitude is a free data retrieval call binding the contract method 0x589af69c.
//
// Solidity: function longitude() constant returns(uint256)
func (_Hotel_Interface *Hotel_InterfaceCallerSession) Longitude() (*big.Int, error) {
	return _Hotel_Interface.Contract.Longitude(&_Hotel_Interface.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_Hotel_Interface *Hotel_InterfaceCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Hotel_Interface.contract.Call(opts, out, "manager")
	return *ret0, err
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_Hotel_Interface *Hotel_InterfaceSession) Manager() (common.Address, error) {
	return _Hotel_Interface.Contract.Manager(&_Hotel_Interface.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_Hotel_Interface *Hotel_InterfaceCallerSession) Manager() (common.Address, error) {
	return _Hotel_Interface.Contract.Manager(&_Hotel_Interface.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Hotel_Interface *Hotel_InterfaceCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Hotel_Interface.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Hotel_Interface *Hotel_InterfaceSession) Name() (string, error) {
	return _Hotel_Interface.Contract.Name(&_Hotel_Interface.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Hotel_Interface *Hotel_InterfaceCallerSession) Name() (string, error) {
	return _Hotel_Interface.Contract.Name(&_Hotel_Interface.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Hotel_Interface *Hotel_InterfaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Hotel_Interface.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Hotel_Interface *Hotel_InterfaceSession) Owner() (common.Address, error) {
	return _Hotel_Interface.Contract.Owner(&_Hotel_Interface.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Hotel_Interface *Hotel_InterfaceCallerSession) Owner() (common.Address, error) {
	return _Hotel_Interface.Contract.Owner(&_Hotel_Interface.CallOpts)
}

// PendingCalls is a free data retrieval call binding the contract method 0x32fdd45c.
//
// Solidity: function pendingCalls( bytes32) constant returns(callData bytes, sender address, approved bool, success bool)
func (_Hotel_Interface *Hotel_InterfaceCaller) PendingCalls(opts *bind.CallOpts, arg0 [32]byte) (struct {
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
	err := _Hotel_Interface.contract.Call(opts, out, "pendingCalls", arg0)
	return *ret, err
}

// PendingCalls is a free data retrieval call binding the contract method 0x32fdd45c.
//
// Solidity: function pendingCalls( bytes32) constant returns(callData bytes, sender address, approved bool, success bool)
func (_Hotel_Interface *Hotel_InterfaceSession) PendingCalls(arg0 [32]byte) (struct {
	CallData []byte
	Sender   common.Address
	Approved bool
	Success  bool
}, error) {
	return _Hotel_Interface.Contract.PendingCalls(&_Hotel_Interface.CallOpts, arg0)
}

// PendingCalls is a free data retrieval call binding the contract method 0x32fdd45c.
//
// Solidity: function pendingCalls( bytes32) constant returns(callData bytes, sender address, approved bool, success bool)
func (_Hotel_Interface *Hotel_InterfaceCallerSession) PendingCalls(arg0 [32]byte) (struct {
	CallData []byte
	Sender   common.Address
	Approved bool
	Success  bool
}, error) {
	return _Hotel_Interface.Contract.PendingCalls(&_Hotel_Interface.CallOpts, arg0)
}

// Timezone is a free data retrieval call binding the contract method 0xc4148fe5.
//
// Solidity: function timezone() constant returns(uint256)
func (_Hotel_Interface *Hotel_InterfaceCaller) Timezone(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Hotel_Interface.contract.Call(opts, out, "timezone")
	return *ret0, err
}

// Timezone is a free data retrieval call binding the contract method 0xc4148fe5.
//
// Solidity: function timezone() constant returns(uint256)
func (_Hotel_Interface *Hotel_InterfaceSession) Timezone() (*big.Int, error) {
	return _Hotel_Interface.Contract.Timezone(&_Hotel_Interface.CallOpts)
}

// Timezone is a free data retrieval call binding the contract method 0xc4148fe5.
//
// Solidity: function timezone() constant returns(uint256)
func (_Hotel_Interface *Hotel_InterfaceCallerSession) Timezone() (*big.Int, error) {
	return _Hotel_Interface.Contract.Timezone(&_Hotel_Interface.CallOpts)
}

// UnitTypeNames is a free data retrieval call binding the contract method 0x28ef6b93.
//
// Solidity: function unitTypeNames( uint256) constant returns(bytes32)
func (_Hotel_Interface *Hotel_InterfaceCaller) UnitTypeNames(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Hotel_Interface.contract.Call(opts, out, "unitTypeNames", arg0)
	return *ret0, err
}

// UnitTypeNames is a free data retrieval call binding the contract method 0x28ef6b93.
//
// Solidity: function unitTypeNames( uint256) constant returns(bytes32)
func (_Hotel_Interface *Hotel_InterfaceSession) UnitTypeNames(arg0 *big.Int) ([32]byte, error) {
	return _Hotel_Interface.Contract.UnitTypeNames(&_Hotel_Interface.CallOpts, arg0)
}

// UnitTypeNames is a free data retrieval call binding the contract method 0x28ef6b93.
//
// Solidity: function unitTypeNames( uint256) constant returns(bytes32)
func (_Hotel_Interface *Hotel_InterfaceCallerSession) UnitTypeNames(arg0 *big.Int) ([32]byte, error) {
	return _Hotel_Interface.Contract.UnitTypeNames(&_Hotel_Interface.CallOpts, arg0)
}

// UnitTypes is a free data retrieval call binding the contract method 0x2bb03f6e.
//
// Solidity: function unitTypes( bytes32) constant returns(address)
func (_Hotel_Interface *Hotel_InterfaceCaller) UnitTypes(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Hotel_Interface.contract.Call(opts, out, "unitTypes", arg0)
	return *ret0, err
}

// UnitTypes is a free data retrieval call binding the contract method 0x2bb03f6e.
//
// Solidity: function unitTypes( bytes32) constant returns(address)
func (_Hotel_Interface *Hotel_InterfaceSession) UnitTypes(arg0 [32]byte) (common.Address, error) {
	return _Hotel_Interface.Contract.UnitTypes(&_Hotel_Interface.CallOpts, arg0)
}

// UnitTypes is a free data retrieval call binding the contract method 0x2bb03f6e.
//
// Solidity: function unitTypes( bytes32) constant returns(address)
func (_Hotel_Interface *Hotel_InterfaceCallerSession) UnitTypes(arg0 [32]byte) (common.Address, error) {
	return _Hotel_Interface.Contract.UnitTypes(&_Hotel_Interface.CallOpts, arg0)
}

// Units is a free data retrieval call binding the contract method 0xe5fba6cc.
//
// Solidity: function units( uint256) constant returns(address)
func (_Hotel_Interface *Hotel_InterfaceCaller) Units(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Hotel_Interface.contract.Call(opts, out, "units", arg0)
	return *ret0, err
}

// Units is a free data retrieval call binding the contract method 0xe5fba6cc.
//
// Solidity: function units( uint256) constant returns(address)
func (_Hotel_Interface *Hotel_InterfaceSession) Units(arg0 *big.Int) (common.Address, error) {
	return _Hotel_Interface.Contract.Units(&_Hotel_Interface.CallOpts, arg0)
}

// Units is a free data retrieval call binding the contract method 0xe5fba6cc.
//
// Solidity: function units( uint256) constant returns(address)
func (_Hotel_Interface *Hotel_InterfaceCallerSession) Units(arg0 *big.Int) (common.Address, error) {
	return _Hotel_Interface.Contract.Units(&_Hotel_Interface.CallOpts, arg0)
}

// UnitsIndex is a free data retrieval call binding the contract method 0x5edad458.
//
// Solidity: function unitsIndex( address) constant returns(uint256)
func (_Hotel_Interface *Hotel_InterfaceCaller) UnitsIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Hotel_Interface.contract.Call(opts, out, "unitsIndex", arg0)
	return *ret0, err
}

// UnitsIndex is a free data retrieval call binding the contract method 0x5edad458.
//
// Solidity: function unitsIndex( address) constant returns(uint256)
func (_Hotel_Interface *Hotel_InterfaceSession) UnitsIndex(arg0 common.Address) (*big.Int, error) {
	return _Hotel_Interface.Contract.UnitsIndex(&_Hotel_Interface.CallOpts, arg0)
}

// UnitsIndex is a free data retrieval call binding the contract method 0x5edad458.
//
// Solidity: function unitsIndex( address) constant returns(uint256)
func (_Hotel_Interface *Hotel_InterfaceCallerSession) UnitsIndex(arg0 common.Address) (*big.Int, error) {
	return _Hotel_Interface.Contract.UnitsIndex(&_Hotel_Interface.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Hotel_Interface *Hotel_InterfaceCaller) Version(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Hotel_Interface.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Hotel_Interface *Hotel_InterfaceSession) Version() ([32]byte, error) {
	return _Hotel_Interface.Contract.Version(&_Hotel_Interface.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Hotel_Interface *Hotel_InterfaceCallerSession) Version() ([32]byte, error) {
	return _Hotel_Interface.Contract.Version(&_Hotel_Interface.CallOpts)
}

// WaitConfirmation is a free data retrieval call binding the contract method 0xac72bb98.
//
// Solidity: function waitConfirmation() constant returns(bool)
func (_Hotel_Interface *Hotel_InterfaceCaller) WaitConfirmation(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Hotel_Interface.contract.Call(opts, out, "waitConfirmation")
	return *ret0, err
}

// WaitConfirmation is a free data retrieval call binding the contract method 0xac72bb98.
//
// Solidity: function waitConfirmation() constant returns(bool)
func (_Hotel_Interface *Hotel_InterfaceSession) WaitConfirmation() (bool, error) {
	return _Hotel_Interface.Contract.WaitConfirmation(&_Hotel_Interface.CallOpts)
}

// WaitConfirmation is a free data retrieval call binding the contract method 0xac72bb98.
//
// Solidity: function waitConfirmation() constant returns(bool)
func (_Hotel_Interface *Hotel_InterfaceCallerSession) WaitConfirmation() (bool, error) {
	return _Hotel_Interface.Contract.WaitConfirmation(&_Hotel_Interface.CallOpts)
}

// Zip is a free data retrieval call binding the contract method 0x00919055.
//
// Solidity: function zip() constant returns(string)
func (_Hotel_Interface *Hotel_InterfaceCaller) Zip(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Hotel_Interface.contract.Call(opts, out, "zip")
	return *ret0, err
}

// Zip is a free data retrieval call binding the contract method 0x00919055.
//
// Solidity: function zip() constant returns(string)
func (_Hotel_Interface *Hotel_InterfaceSession) Zip() (string, error) {
	return _Hotel_Interface.Contract.Zip(&_Hotel_Interface.CallOpts)
}

// Zip is a free data retrieval call binding the contract method 0x00919055.
//
// Solidity: function zip() constant returns(string)
func (_Hotel_Interface *Hotel_InterfaceCallerSession) Zip() (string, error) {
	return _Hotel_Interface.Contract.Zip(&_Hotel_Interface.CallOpts)
}

// AddImage is a paid mutator transaction binding the contract method 0x166f1779.
//
// Solidity: function addImage(url string) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactor) AddImage(opts *bind.TransactOpts, url string) (*types.Transaction, error) {
	return _Hotel_Interface.contract.Transact(opts, "addImage", url)
}

// AddImage is a paid mutator transaction binding the contract method 0x166f1779.
//
// Solidity: function addImage(url string) returns()
func (_Hotel_Interface *Hotel_InterfaceSession) AddImage(url string) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.AddImage(&_Hotel_Interface.TransactOpts, url)
}

// AddImage is a paid mutator transaction binding the contract method 0x166f1779.
//
// Solidity: function addImage(url string) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactorSession) AddImage(url string) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.AddImage(&_Hotel_Interface.TransactOpts, url)
}

// AddUnit is a paid mutator transaction binding the contract method 0xa9f7f247.
//
// Solidity: function addUnit(unit address) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactor) AddUnit(opts *bind.TransactOpts, unit common.Address) (*types.Transaction, error) {
	return _Hotel_Interface.contract.Transact(opts, "addUnit", unit)
}

// AddUnit is a paid mutator transaction binding the contract method 0xa9f7f247.
//
// Solidity: function addUnit(unit address) returns()
func (_Hotel_Interface *Hotel_InterfaceSession) AddUnit(unit common.Address) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.AddUnit(&_Hotel_Interface.TransactOpts, unit)
}

// AddUnit is a paid mutator transaction binding the contract method 0xa9f7f247.
//
// Solidity: function addUnit(unit address) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactorSession) AddUnit(unit common.Address) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.AddUnit(&_Hotel_Interface.TransactOpts, unit)
}

// AddUnitType is a paid mutator transaction binding the contract method 0xdc8f86e5.
//
// Solidity: function addUnitType(addr address) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactor) AddUnitType(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Hotel_Interface.contract.Transact(opts, "addUnitType", addr)
}

// AddUnitType is a paid mutator transaction binding the contract method 0xdc8f86e5.
//
// Solidity: function addUnitType(addr address) returns()
func (_Hotel_Interface *Hotel_InterfaceSession) AddUnitType(addr common.Address) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.AddUnitType(&_Hotel_Interface.TransactOpts, addr)
}

// AddUnitType is a paid mutator transaction binding the contract method 0xdc8f86e5.
//
// Solidity: function addUnitType(addr address) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactorSession) AddUnitType(addr common.Address) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.AddUnitType(&_Hotel_Interface.TransactOpts, addr)
}

// BeginCall is a paid mutator transaction binding the contract method 0x203eaf27.
//
// Solidity: function beginCall(publicCallData bytes, privateData bytes) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactor) BeginCall(opts *bind.TransactOpts, publicCallData []byte, privateData []byte) (*types.Transaction, error) {
	return _Hotel_Interface.contract.Transact(opts, "beginCall", publicCallData, privateData)
}

// BeginCall is a paid mutator transaction binding the contract method 0x203eaf27.
//
// Solidity: function beginCall(publicCallData bytes, privateData bytes) returns()
func (_Hotel_Interface *Hotel_InterfaceSession) BeginCall(publicCallData []byte, privateData []byte) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.BeginCall(&_Hotel_Interface.TransactOpts, publicCallData, privateData)
}

// BeginCall is a paid mutator transaction binding the contract method 0x203eaf27.
//
// Solidity: function beginCall(publicCallData bytes, privateData bytes) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactorSession) BeginCall(publicCallData []byte, privateData []byte) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.BeginCall(&_Hotel_Interface.TransactOpts, publicCallData, privateData)
}

// Book is a paid mutator transaction binding the contract method 0xb9a527b4.
//
// Solidity: function book(unitAddress address, from address, fromDay uint256, daysAmount uint256) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactor) Book(opts *bind.TransactOpts, unitAddress common.Address, from common.Address, fromDay *big.Int, daysAmount *big.Int) (*types.Transaction, error) {
	return _Hotel_Interface.contract.Transact(opts, "book", unitAddress, from, fromDay, daysAmount)
}

// Book is a paid mutator transaction binding the contract method 0xb9a527b4.
//
// Solidity: function book(unitAddress address, from address, fromDay uint256, daysAmount uint256) returns()
func (_Hotel_Interface *Hotel_InterfaceSession) Book(unitAddress common.Address, from common.Address, fromDay *big.Int, daysAmount *big.Int) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.Book(&_Hotel_Interface.TransactOpts, unitAddress, from, fromDay, daysAmount)
}

// Book is a paid mutator transaction binding the contract method 0xb9a527b4.
//
// Solidity: function book(unitAddress address, from address, fromDay uint256, daysAmount uint256) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactorSession) Book(unitAddress common.Address, from common.Address, fromDay *big.Int, daysAmount *big.Int) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.Book(&_Hotel_Interface.TransactOpts, unitAddress, from, fromDay, daysAmount)
}

// BookWithLif is a paid mutator transaction binding the contract method 0x676e0079.
//
// Solidity: function bookWithLif(unitAddress address, from address, fromDay uint256, daysAmount uint256) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactor) BookWithLif(opts *bind.TransactOpts, unitAddress common.Address, from common.Address, fromDay *big.Int, daysAmount *big.Int) (*types.Transaction, error) {
	return _Hotel_Interface.contract.Transact(opts, "bookWithLif", unitAddress, from, fromDay, daysAmount)
}

// BookWithLif is a paid mutator transaction binding the contract method 0x676e0079.
//
// Solidity: function bookWithLif(unitAddress address, from address, fromDay uint256, daysAmount uint256) returns()
func (_Hotel_Interface *Hotel_InterfaceSession) BookWithLif(unitAddress common.Address, from common.Address, fromDay *big.Int, daysAmount *big.Int) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.BookWithLif(&_Hotel_Interface.TransactOpts, unitAddress, from, fromDay, daysAmount)
}

// BookWithLif is a paid mutator transaction binding the contract method 0x676e0079.
//
// Solidity: function bookWithLif(unitAddress address, from address, fromDay uint256, daysAmount uint256) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactorSession) BookWithLif(unitAddress common.Address, from common.Address, fromDay *big.Int, daysAmount *big.Int) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.BookWithLif(&_Hotel_Interface.TransactOpts, unitAddress, from, fromDay, daysAmount)
}

// CallUnit is a paid mutator transaction binding the contract method 0x1c4e4d46.
//
// Solidity: function callUnit(unitAddress address, data bytes) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactor) CallUnit(opts *bind.TransactOpts, unitAddress common.Address, data []byte) (*types.Transaction, error) {
	return _Hotel_Interface.contract.Transact(opts, "callUnit", unitAddress, data)
}

// CallUnit is a paid mutator transaction binding the contract method 0x1c4e4d46.
//
// Solidity: function callUnit(unitAddress address, data bytes) returns()
func (_Hotel_Interface *Hotel_InterfaceSession) CallUnit(unitAddress common.Address, data []byte) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.CallUnit(&_Hotel_Interface.TransactOpts, unitAddress, data)
}

// CallUnit is a paid mutator transaction binding the contract method 0x1c4e4d46.
//
// Solidity: function callUnit(unitAddress address, data bytes) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactorSession) CallUnit(unitAddress common.Address, data []byte) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.CallUnit(&_Hotel_Interface.TransactOpts, unitAddress, data)
}

// CallUnitType is a paid mutator transaction binding the contract method 0xdbc6c290.
//
// Solidity: function callUnitType(unitType bytes32, data bytes) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactor) CallUnitType(opts *bind.TransactOpts, unitType [32]byte, data []byte) (*types.Transaction, error) {
	return _Hotel_Interface.contract.Transact(opts, "callUnitType", unitType, data)
}

// CallUnitType is a paid mutator transaction binding the contract method 0xdbc6c290.
//
// Solidity: function callUnitType(unitType bytes32, data bytes) returns()
func (_Hotel_Interface *Hotel_InterfaceSession) CallUnitType(unitType [32]byte, data []byte) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.CallUnitType(&_Hotel_Interface.TransactOpts, unitType, data)
}

// CallUnitType is a paid mutator transaction binding the contract method 0xdbc6c290.
//
// Solidity: function callUnitType(unitType bytes32, data bytes) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactorSession) CallUnitType(unitType [32]byte, data []byte) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.CallUnitType(&_Hotel_Interface.TransactOpts, unitType, data)
}

// ChangeConfirmation is a paid mutator transaction binding the contract method 0x44f5484a.
//
// Solidity: function changeConfirmation(_waitConfirmation bool) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactor) ChangeConfirmation(opts *bind.TransactOpts, _waitConfirmation bool) (*types.Transaction, error) {
	return _Hotel_Interface.contract.Transact(opts, "changeConfirmation", _waitConfirmation)
}

// ChangeConfirmation is a paid mutator transaction binding the contract method 0x44f5484a.
//
// Solidity: function changeConfirmation(_waitConfirmation bool) returns()
func (_Hotel_Interface *Hotel_InterfaceSession) ChangeConfirmation(_waitConfirmation bool) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.ChangeConfirmation(&_Hotel_Interface.TransactOpts, _waitConfirmation)
}

// ChangeConfirmation is a paid mutator transaction binding the contract method 0x44f5484a.
//
// Solidity: function changeConfirmation(_waitConfirmation bool) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactorSession) ChangeConfirmation(_waitConfirmation bool) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.ChangeConfirmation(&_Hotel_Interface.TransactOpts, _waitConfirmation)
}

// ChangeUnitType is a paid mutator transaction binding the contract method 0x37844767.
//
// Solidity: function changeUnitType(unitType bytes32, newAddr address) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactor) ChangeUnitType(opts *bind.TransactOpts, unitType [32]byte, newAddr common.Address) (*types.Transaction, error) {
	return _Hotel_Interface.contract.Transact(opts, "changeUnitType", unitType, newAddr)
}

// ChangeUnitType is a paid mutator transaction binding the contract method 0x37844767.
//
// Solidity: function changeUnitType(unitType bytes32, newAddr address) returns()
func (_Hotel_Interface *Hotel_InterfaceSession) ChangeUnitType(unitType [32]byte, newAddr common.Address) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.ChangeUnitType(&_Hotel_Interface.TransactOpts, unitType, newAddr)
}

// ChangeUnitType is a paid mutator transaction binding the contract method 0x37844767.
//
// Solidity: function changeUnitType(unitType bytes32, newAddr address) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactorSession) ChangeUnitType(unitType [32]byte, newAddr common.Address) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.ChangeUnitType(&_Hotel_Interface.TransactOpts, unitType, newAddr)
}

// ContinueCall is a paid mutator transaction binding the contract method 0x411f2351.
//
// Solidity: function continueCall(msgDataHash bytes32) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactor) ContinueCall(opts *bind.TransactOpts, msgDataHash [32]byte) (*types.Transaction, error) {
	return _Hotel_Interface.contract.Transact(opts, "continueCall", msgDataHash)
}

// ContinueCall is a paid mutator transaction binding the contract method 0x411f2351.
//
// Solidity: function continueCall(msgDataHash bytes32) returns()
func (_Hotel_Interface *Hotel_InterfaceSession) ContinueCall(msgDataHash [32]byte) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.ContinueCall(&_Hotel_Interface.TransactOpts, msgDataHash)
}

// ContinueCall is a paid mutator transaction binding the contract method 0x411f2351.
//
// Solidity: function continueCall(msgDataHash bytes32) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactorSession) ContinueCall(msgDataHash [32]byte) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.ContinueCall(&_Hotel_Interface.TransactOpts, msgDataHash)
}

// EditAddress is a paid mutator transaction binding the contract method 0x7e46ce2f.
//
// Solidity: function editAddress(_lineOne string, _lineTwo string, _zip string, _country string) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactor) EditAddress(opts *bind.TransactOpts, _lineOne string, _lineTwo string, _zip string, _country string) (*types.Transaction, error) {
	return _Hotel_Interface.contract.Transact(opts, "editAddress", _lineOne, _lineTwo, _zip, _country)
}

// EditAddress is a paid mutator transaction binding the contract method 0x7e46ce2f.
//
// Solidity: function editAddress(_lineOne string, _lineTwo string, _zip string, _country string) returns()
func (_Hotel_Interface *Hotel_InterfaceSession) EditAddress(_lineOne string, _lineTwo string, _zip string, _country string) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.EditAddress(&_Hotel_Interface.TransactOpts, _lineOne, _lineTwo, _zip, _country)
}

// EditAddress is a paid mutator transaction binding the contract method 0x7e46ce2f.
//
// Solidity: function editAddress(_lineOne string, _lineTwo string, _zip string, _country string) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactorSession) EditAddress(_lineOne string, _lineTwo string, _zip string, _country string) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.EditAddress(&_Hotel_Interface.TransactOpts, _lineOne, _lineTwo, _zip, _country)
}

// EditInfo is a paid mutator transaction binding the contract method 0x0b0d3ca2.
//
// Solidity: function editInfo(_name string, _description string) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactor) EditInfo(opts *bind.TransactOpts, _name string, _description string) (*types.Transaction, error) {
	return _Hotel_Interface.contract.Transact(opts, "editInfo", _name, _description)
}

// EditInfo is a paid mutator transaction binding the contract method 0x0b0d3ca2.
//
// Solidity: function editInfo(_name string, _description string) returns()
func (_Hotel_Interface *Hotel_InterfaceSession) EditInfo(_name string, _description string) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.EditInfo(&_Hotel_Interface.TransactOpts, _name, _description)
}

// EditInfo is a paid mutator transaction binding the contract method 0x0b0d3ca2.
//
// Solidity: function editInfo(_name string, _description string) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactorSession) EditInfo(_name string, _description string) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.EditInfo(&_Hotel_Interface.TransactOpts, _name, _description)
}

// EditLocation is a paid mutator transaction binding the contract method 0x979a1936.
//
// Solidity: function editLocation(_timezone uint256, _longitude uint256, _latitude uint256) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactor) EditLocation(opts *bind.TransactOpts, _timezone *big.Int, _longitude *big.Int, _latitude *big.Int) (*types.Transaction, error) {
	return _Hotel_Interface.contract.Transact(opts, "editLocation", _timezone, _longitude, _latitude)
}

// EditLocation is a paid mutator transaction binding the contract method 0x979a1936.
//
// Solidity: function editLocation(_timezone uint256, _longitude uint256, _latitude uint256) returns()
func (_Hotel_Interface *Hotel_InterfaceSession) EditLocation(_timezone *big.Int, _longitude *big.Int, _latitude *big.Int) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.EditLocation(&_Hotel_Interface.TransactOpts, _timezone, _longitude, _latitude)
}

// EditLocation is a paid mutator transaction binding the contract method 0x979a1936.
//
// Solidity: function editLocation(_timezone uint256, _longitude uint256, _latitude uint256) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactorSession) EditLocation(_timezone *big.Int, _longitude *big.Int, _latitude *big.Int) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.EditLocation(&_Hotel_Interface.TransactOpts, _timezone, _longitude, _latitude)
}

// RemoveImage is a paid mutator transaction binding the contract method 0x8b25261f.
//
// Solidity: function removeImage(index uint256) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactor) RemoveImage(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _Hotel_Interface.contract.Transact(opts, "removeImage", index)
}

// RemoveImage is a paid mutator transaction binding the contract method 0x8b25261f.
//
// Solidity: function removeImage(index uint256) returns()
func (_Hotel_Interface *Hotel_InterfaceSession) RemoveImage(index *big.Int) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.RemoveImage(&_Hotel_Interface.TransactOpts, index)
}

// RemoveImage is a paid mutator transaction binding the contract method 0x8b25261f.
//
// Solidity: function removeImage(index uint256) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactorSession) RemoveImage(index *big.Int) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.RemoveImage(&_Hotel_Interface.TransactOpts, index)
}

// RemoveUnit is a paid mutator transaction binding the contract method 0xdf155165.
//
// Solidity: function removeUnit(unit address) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactor) RemoveUnit(opts *bind.TransactOpts, unit common.Address) (*types.Transaction, error) {
	return _Hotel_Interface.contract.Transact(opts, "removeUnit", unit)
}

// RemoveUnit is a paid mutator transaction binding the contract method 0xdf155165.
//
// Solidity: function removeUnit(unit address) returns()
func (_Hotel_Interface *Hotel_InterfaceSession) RemoveUnit(unit common.Address) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.RemoveUnit(&_Hotel_Interface.TransactOpts, unit)
}

// RemoveUnit is a paid mutator transaction binding the contract method 0xdf155165.
//
// Solidity: function removeUnit(unit address) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactorSession) RemoveUnit(unit common.Address) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.RemoveUnit(&_Hotel_Interface.TransactOpts, unit)
}

// RemoveUnitType is a paid mutator transaction binding the contract method 0xd5268e8c.
//
// Solidity: function removeUnitType(unitType bytes32, index uint256) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactor) RemoveUnitType(opts *bind.TransactOpts, unitType [32]byte, index *big.Int) (*types.Transaction, error) {
	return _Hotel_Interface.contract.Transact(opts, "removeUnitType", unitType, index)
}

// RemoveUnitType is a paid mutator transaction binding the contract method 0xd5268e8c.
//
// Solidity: function removeUnitType(unitType bytes32, index uint256) returns()
func (_Hotel_Interface *Hotel_InterfaceSession) RemoveUnitType(unitType [32]byte, index *big.Int) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.RemoveUnitType(&_Hotel_Interface.TransactOpts, unitType, index)
}

// RemoveUnitType is a paid mutator transaction binding the contract method 0xd5268e8c.
//
// Solidity: function removeUnitType(unitType bytes32, index uint256) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactorSession) RemoveUnitType(unitType [32]byte, index *big.Int) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.RemoveUnitType(&_Hotel_Interface.TransactOpts, unitType, index)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Hotel_Interface.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Hotel_Interface *Hotel_InterfaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.TransferOwnership(&_Hotel_Interface.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Hotel_Interface *Hotel_InterfaceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Hotel_Interface.Contract.TransferOwnership(&_Hotel_Interface.TransactOpts, newOwner)
}
