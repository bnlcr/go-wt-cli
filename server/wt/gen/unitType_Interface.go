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

// UnitType_InterfaceABI is the input ABI used to generate the binding from.
const UnitType_InterfaceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"amenity\",\"type\":\"uint256\"}],\"name\":\"removeAmenity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"url\",\"type\":\"string\"}],\"name\":\"addImage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"increaseUnits\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getImagesLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalUnits\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"decreaseUnits\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"images\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unitType\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeImage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"description\",\"type\":\"string\"},{\"name\":\"minGuests\",\"type\":\"uint256\"},{\"name\":\"maxGuests\",\"type\":\"uint256\"},{\"name\":\"price\",\"type\":\"string\"}],\"name\":\"edit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractType\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAmenities\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amenity\",\"type\":\"uint256\"}],\"name\":\"addAmenity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"unitIndex\",\"type\":\"uint256\"}],\"name\":\"removeUnit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// UnitType_InterfaceBin is the compiled bytecode used for deploying new contracts.
const UnitType_InterfaceBin = `0x`

// DeployUnitType_Interface deploys a new Ethereum contract, binding an instance of UnitType_Interface to it.
func DeployUnitType_Interface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UnitType_Interface, error) {
	parsed, err := abi.JSON(strings.NewReader(UnitType_InterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UnitType_InterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UnitType_Interface{UnitType_InterfaceCaller: UnitType_InterfaceCaller{contract: contract}, UnitType_InterfaceTransactor: UnitType_InterfaceTransactor{contract: contract}}, nil
}

// UnitType_Interface is an auto generated Go binding around an Ethereum contract.
type UnitType_Interface struct {
	UnitType_InterfaceCaller     // Read-only binding to the contract
	UnitType_InterfaceTransactor // Write-only binding to the contract
}

// UnitType_InterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnitType_InterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnitType_InterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnitType_InterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnitType_InterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnitType_InterfaceSession struct {
	Contract     *UnitType_Interface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// UnitType_InterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnitType_InterfaceCallerSession struct {
	Contract *UnitType_InterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// UnitType_InterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnitType_InterfaceTransactorSession struct {
	Contract     *UnitType_InterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// UnitType_InterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnitType_InterfaceRaw struct {
	Contract *UnitType_Interface // Generic contract binding to access the raw methods on
}

// UnitType_InterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnitType_InterfaceCallerRaw struct {
	Contract *UnitType_InterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// UnitType_InterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnitType_InterfaceTransactorRaw struct {
	Contract *UnitType_InterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnitType_Interface creates a new instance of UnitType_Interface, bound to a specific deployed contract.
func NewUnitType_Interface(address common.Address, backend bind.ContractBackend) (*UnitType_Interface, error) {
	contract, err := bindUnitType_Interface(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UnitType_Interface{UnitType_InterfaceCaller: UnitType_InterfaceCaller{contract: contract}, UnitType_InterfaceTransactor: UnitType_InterfaceTransactor{contract: contract}}, nil
}

// NewUnitType_InterfaceCaller creates a new read-only instance of UnitType_Interface, bound to a specific deployed contract.
func NewUnitType_InterfaceCaller(address common.Address, caller bind.ContractCaller) (*UnitType_InterfaceCaller, error) {
	contract, err := bindUnitType_Interface(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &UnitType_InterfaceCaller{contract: contract}, nil
}

// NewUnitType_InterfaceTransactor creates a new write-only instance of UnitType_Interface, bound to a specific deployed contract.
func NewUnitType_InterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*UnitType_InterfaceTransactor, error) {
	contract, err := bindUnitType_Interface(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &UnitType_InterfaceTransactor{contract: contract}, nil
}

// bindUnitType_Interface binds a generic wrapper to an already deployed contract.
func bindUnitType_Interface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UnitType_InterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnitType_Interface *UnitType_InterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UnitType_Interface.Contract.UnitType_InterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnitType_Interface *UnitType_InterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnitType_Interface.Contract.UnitType_InterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnitType_Interface *UnitType_InterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnitType_Interface.Contract.UnitType_InterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnitType_Interface *UnitType_InterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UnitType_Interface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnitType_Interface *UnitType_InterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnitType_Interface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnitType_Interface *UnitType_InterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnitType_Interface.Contract.contract.Transact(opts, method, params...)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() constant returns(bytes32)
func (_UnitType_Interface *UnitType_InterfaceCaller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _UnitType_Interface.contract.Call(opts, out, "contractType")
	return *ret0, err
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() constant returns(bytes32)
func (_UnitType_Interface *UnitType_InterfaceSession) ContractType() ([32]byte, error) {
	return _UnitType_Interface.Contract.ContractType(&_UnitType_Interface.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() constant returns(bytes32)
func (_UnitType_Interface *UnitType_InterfaceCallerSession) ContractType() ([32]byte, error) {
	return _UnitType_Interface.Contract.ContractType(&_UnitType_Interface.CallOpts)
}

// GetAmenities is a free data retrieval call binding the contract method 0xd0ceaae8.
//
// Solidity: function getAmenities() constant returns(uint256[])
func (_UnitType_Interface *UnitType_InterfaceCaller) GetAmenities(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _UnitType_Interface.contract.Call(opts, out, "getAmenities")
	return *ret0, err
}

// GetAmenities is a free data retrieval call binding the contract method 0xd0ceaae8.
//
// Solidity: function getAmenities() constant returns(uint256[])
func (_UnitType_Interface *UnitType_InterfaceSession) GetAmenities() ([]*big.Int, error) {
	return _UnitType_Interface.Contract.GetAmenities(&_UnitType_Interface.CallOpts)
}

// GetAmenities is a free data retrieval call binding the contract method 0xd0ceaae8.
//
// Solidity: function getAmenities() constant returns(uint256[])
func (_UnitType_Interface *UnitType_InterfaceCallerSession) GetAmenities() ([]*big.Int, error) {
	return _UnitType_Interface.Contract.GetAmenities(&_UnitType_Interface.CallOpts)
}

// GetImagesLength is a free data retrieval call binding the contract method 0x60839bd8.
//
// Solidity: function getImagesLength() constant returns(uint256)
func (_UnitType_Interface *UnitType_InterfaceCaller) GetImagesLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UnitType_Interface.contract.Call(opts, out, "getImagesLength")
	return *ret0, err
}

// GetImagesLength is a free data retrieval call binding the contract method 0x60839bd8.
//
// Solidity: function getImagesLength() constant returns(uint256)
func (_UnitType_Interface *UnitType_InterfaceSession) GetImagesLength() (*big.Int, error) {
	return _UnitType_Interface.Contract.GetImagesLength(&_UnitType_Interface.CallOpts)
}

// GetImagesLength is a free data retrieval call binding the contract method 0x60839bd8.
//
// Solidity: function getImagesLength() constant returns(uint256)
func (_UnitType_Interface *UnitType_InterfaceCallerSession) GetImagesLength() (*big.Int, error) {
	return _UnitType_Interface.Contract.GetImagesLength(&_UnitType_Interface.CallOpts)
}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() constant returns(string, uint256, uint256, string, bool)
func (_UnitType_Interface *UnitType_InterfaceCaller) GetInfo(opts *bind.CallOpts) (string, *big.Int, *big.Int, string, bool, error) {
	var (
		ret0 = new(string)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(string)
		ret4 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _UnitType_Interface.contract.Call(opts, out, "getInfo")
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() constant returns(string, uint256, uint256, string, bool)
func (_UnitType_Interface *UnitType_InterfaceSession) GetInfo() (string, *big.Int, *big.Int, string, bool, error) {
	return _UnitType_Interface.Contract.GetInfo(&_UnitType_Interface.CallOpts)
}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() constant returns(string, uint256, uint256, string, bool)
func (_UnitType_Interface *UnitType_InterfaceCallerSession) GetInfo() (string, *big.Int, *big.Int, string, bool, error) {
	return _UnitType_Interface.Contract.GetInfo(&_UnitType_Interface.CallOpts)
}

// Images is a free data retrieval call binding the contract method 0x84856482.
//
// Solidity: function images( uint256) constant returns(string)
func (_UnitType_Interface *UnitType_InterfaceCaller) Images(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _UnitType_Interface.contract.Call(opts, out, "images", arg0)
	return *ret0, err
}

// Images is a free data retrieval call binding the contract method 0x84856482.
//
// Solidity: function images( uint256) constant returns(string)
func (_UnitType_Interface *UnitType_InterfaceSession) Images(arg0 *big.Int) (string, error) {
	return _UnitType_Interface.Contract.Images(&_UnitType_Interface.CallOpts, arg0)
}

// Images is a free data retrieval call binding the contract method 0x84856482.
//
// Solidity: function images( uint256) constant returns(string)
func (_UnitType_Interface *UnitType_InterfaceCallerSession) Images(arg0 *big.Int) (string, error) {
	return _UnitType_Interface.Contract.Images(&_UnitType_Interface.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UnitType_Interface *UnitType_InterfaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UnitType_Interface.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UnitType_Interface *UnitType_InterfaceSession) Owner() (common.Address, error) {
	return _UnitType_Interface.Contract.Owner(&_UnitType_Interface.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UnitType_Interface *UnitType_InterfaceCallerSession) Owner() (common.Address, error) {
	return _UnitType_Interface.Contract.Owner(&_UnitType_Interface.CallOpts)
}

// TotalUnits is a free data retrieval call binding the contract method 0x6d86acc4.
//
// Solidity: function totalUnits() constant returns(uint256)
func (_UnitType_Interface *UnitType_InterfaceCaller) TotalUnits(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UnitType_Interface.contract.Call(opts, out, "totalUnits")
	return *ret0, err
}

// TotalUnits is a free data retrieval call binding the contract method 0x6d86acc4.
//
// Solidity: function totalUnits() constant returns(uint256)
func (_UnitType_Interface *UnitType_InterfaceSession) TotalUnits() (*big.Int, error) {
	return _UnitType_Interface.Contract.TotalUnits(&_UnitType_Interface.CallOpts)
}

// TotalUnits is a free data retrieval call binding the contract method 0x6d86acc4.
//
// Solidity: function totalUnits() constant returns(uint256)
func (_UnitType_Interface *UnitType_InterfaceCallerSession) TotalUnits() (*big.Int, error) {
	return _UnitType_Interface.Contract.TotalUnits(&_UnitType_Interface.CallOpts)
}

// UnitType is a free data retrieval call binding the contract method 0x85ad4f90.
//
// Solidity: function unitType() constant returns(bytes32)
func (_UnitType_Interface *UnitType_InterfaceCaller) UnitType(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _UnitType_Interface.contract.Call(opts, out, "unitType")
	return *ret0, err
}

// UnitType is a free data retrieval call binding the contract method 0x85ad4f90.
//
// Solidity: function unitType() constant returns(bytes32)
func (_UnitType_Interface *UnitType_InterfaceSession) UnitType() ([32]byte, error) {
	return _UnitType_Interface.Contract.UnitType(&_UnitType_Interface.CallOpts)
}

// UnitType is a free data retrieval call binding the contract method 0x85ad4f90.
//
// Solidity: function unitType() constant returns(bytes32)
func (_UnitType_Interface *UnitType_InterfaceCallerSession) UnitType() ([32]byte, error) {
	return _UnitType_Interface.Contract.UnitType(&_UnitType_Interface.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_UnitType_Interface *UnitType_InterfaceCaller) Version(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _UnitType_Interface.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_UnitType_Interface *UnitType_InterfaceSession) Version() ([32]byte, error) {
	return _UnitType_Interface.Contract.Version(&_UnitType_Interface.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_UnitType_Interface *UnitType_InterfaceCallerSession) Version() ([32]byte, error) {
	return _UnitType_Interface.Contract.Version(&_UnitType_Interface.CallOpts)
}

// AddAmenity is a paid mutator transaction binding the contract method 0xf240180c.
//
// Solidity: function addAmenity(amenity uint256) returns()
func (_UnitType_Interface *UnitType_InterfaceTransactor) AddAmenity(opts *bind.TransactOpts, amenity *big.Int) (*types.Transaction, error) {
	return _UnitType_Interface.contract.Transact(opts, "addAmenity", amenity)
}

// AddAmenity is a paid mutator transaction binding the contract method 0xf240180c.
//
// Solidity: function addAmenity(amenity uint256) returns()
func (_UnitType_Interface *UnitType_InterfaceSession) AddAmenity(amenity *big.Int) (*types.Transaction, error) {
	return _UnitType_Interface.Contract.AddAmenity(&_UnitType_Interface.TransactOpts, amenity)
}

// AddAmenity is a paid mutator transaction binding the contract method 0xf240180c.
//
// Solidity: function addAmenity(amenity uint256) returns()
func (_UnitType_Interface *UnitType_InterfaceTransactorSession) AddAmenity(amenity *big.Int) (*types.Transaction, error) {
	return _UnitType_Interface.Contract.AddAmenity(&_UnitType_Interface.TransactOpts, amenity)
}

// AddImage is a paid mutator transaction binding the contract method 0x166f1779.
//
// Solidity: function addImage(url string) returns()
func (_UnitType_Interface *UnitType_InterfaceTransactor) AddImage(opts *bind.TransactOpts, url string) (*types.Transaction, error) {
	return _UnitType_Interface.contract.Transact(opts, "addImage", url)
}

// AddImage is a paid mutator transaction binding the contract method 0x166f1779.
//
// Solidity: function addImage(url string) returns()
func (_UnitType_Interface *UnitType_InterfaceSession) AddImage(url string) (*types.Transaction, error) {
	return _UnitType_Interface.Contract.AddImage(&_UnitType_Interface.TransactOpts, url)
}

// AddImage is a paid mutator transaction binding the contract method 0x166f1779.
//
// Solidity: function addImage(url string) returns()
func (_UnitType_Interface *UnitType_InterfaceTransactorSession) AddImage(url string) (*types.Transaction, error) {
	return _UnitType_Interface.Contract.AddImage(&_UnitType_Interface.TransactOpts, url)
}

// DecreaseUnits is a paid mutator transaction binding the contract method 0x72730c01.
//
// Solidity: function decreaseUnits() returns()
func (_UnitType_Interface *UnitType_InterfaceTransactor) DecreaseUnits(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnitType_Interface.contract.Transact(opts, "decreaseUnits")
}

// DecreaseUnits is a paid mutator transaction binding the contract method 0x72730c01.
//
// Solidity: function decreaseUnits() returns()
func (_UnitType_Interface *UnitType_InterfaceSession) DecreaseUnits() (*types.Transaction, error) {
	return _UnitType_Interface.Contract.DecreaseUnits(&_UnitType_Interface.TransactOpts)
}

// DecreaseUnits is a paid mutator transaction binding the contract method 0x72730c01.
//
// Solidity: function decreaseUnits() returns()
func (_UnitType_Interface *UnitType_InterfaceTransactorSession) DecreaseUnits() (*types.Transaction, error) {
	return _UnitType_Interface.Contract.DecreaseUnits(&_UnitType_Interface.TransactOpts)
}

// Edit is a paid mutator transaction binding the contract method 0x91e31183.
//
// Solidity: function edit(description string, minGuests uint256, maxGuests uint256, price string) returns()
func (_UnitType_Interface *UnitType_InterfaceTransactor) Edit(opts *bind.TransactOpts, description string, minGuests *big.Int, maxGuests *big.Int, price string) (*types.Transaction, error) {
	return _UnitType_Interface.contract.Transact(opts, "edit", description, minGuests, maxGuests, price)
}

// Edit is a paid mutator transaction binding the contract method 0x91e31183.
//
// Solidity: function edit(description string, minGuests uint256, maxGuests uint256, price string) returns()
func (_UnitType_Interface *UnitType_InterfaceSession) Edit(description string, minGuests *big.Int, maxGuests *big.Int, price string) (*types.Transaction, error) {
	return _UnitType_Interface.Contract.Edit(&_UnitType_Interface.TransactOpts, description, minGuests, maxGuests, price)
}

// Edit is a paid mutator transaction binding the contract method 0x91e31183.
//
// Solidity: function edit(description string, minGuests uint256, maxGuests uint256, price string) returns()
func (_UnitType_Interface *UnitType_InterfaceTransactorSession) Edit(description string, minGuests *big.Int, maxGuests *big.Int, price string) (*types.Transaction, error) {
	return _UnitType_Interface.Contract.Edit(&_UnitType_Interface.TransactOpts, description, minGuests, maxGuests, price)
}

// IncreaseUnits is a paid mutator transaction binding the contract method 0x362db291.
//
// Solidity: function increaseUnits() returns()
func (_UnitType_Interface *UnitType_InterfaceTransactor) IncreaseUnits(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnitType_Interface.contract.Transact(opts, "increaseUnits")
}

// IncreaseUnits is a paid mutator transaction binding the contract method 0x362db291.
//
// Solidity: function increaseUnits() returns()
func (_UnitType_Interface *UnitType_InterfaceSession) IncreaseUnits() (*types.Transaction, error) {
	return _UnitType_Interface.Contract.IncreaseUnits(&_UnitType_Interface.TransactOpts)
}

// IncreaseUnits is a paid mutator transaction binding the contract method 0x362db291.
//
// Solidity: function increaseUnits() returns()
func (_UnitType_Interface *UnitType_InterfaceTransactorSession) IncreaseUnits() (*types.Transaction, error) {
	return _UnitType_Interface.Contract.IncreaseUnits(&_UnitType_Interface.TransactOpts)
}

// RemoveAmenity is a paid mutator transaction binding the contract method 0x0ad23923.
//
// Solidity: function removeAmenity(amenity uint256) returns()
func (_UnitType_Interface *UnitType_InterfaceTransactor) RemoveAmenity(opts *bind.TransactOpts, amenity *big.Int) (*types.Transaction, error) {
	return _UnitType_Interface.contract.Transact(opts, "removeAmenity", amenity)
}

// RemoveAmenity is a paid mutator transaction binding the contract method 0x0ad23923.
//
// Solidity: function removeAmenity(amenity uint256) returns()
func (_UnitType_Interface *UnitType_InterfaceSession) RemoveAmenity(amenity *big.Int) (*types.Transaction, error) {
	return _UnitType_Interface.Contract.RemoveAmenity(&_UnitType_Interface.TransactOpts, amenity)
}

// RemoveAmenity is a paid mutator transaction binding the contract method 0x0ad23923.
//
// Solidity: function removeAmenity(amenity uint256) returns()
func (_UnitType_Interface *UnitType_InterfaceTransactorSession) RemoveAmenity(amenity *big.Int) (*types.Transaction, error) {
	return _UnitType_Interface.Contract.RemoveAmenity(&_UnitType_Interface.TransactOpts, amenity)
}

// RemoveImage is a paid mutator transaction binding the contract method 0x8b25261f.
//
// Solidity: function removeImage(index uint256) returns()
func (_UnitType_Interface *UnitType_InterfaceTransactor) RemoveImage(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _UnitType_Interface.contract.Transact(opts, "removeImage", index)
}

// RemoveImage is a paid mutator transaction binding the contract method 0x8b25261f.
//
// Solidity: function removeImage(index uint256) returns()
func (_UnitType_Interface *UnitType_InterfaceSession) RemoveImage(index *big.Int) (*types.Transaction, error) {
	return _UnitType_Interface.Contract.RemoveImage(&_UnitType_Interface.TransactOpts, index)
}

// RemoveImage is a paid mutator transaction binding the contract method 0x8b25261f.
//
// Solidity: function removeImage(index uint256) returns()
func (_UnitType_Interface *UnitType_InterfaceTransactorSession) RemoveImage(index *big.Int) (*types.Transaction, error) {
	return _UnitType_Interface.Contract.RemoveImage(&_UnitType_Interface.TransactOpts, index)
}

// RemoveUnit is a paid mutator transaction binding the contract method 0xff3c6039.
//
// Solidity: function removeUnit(unitIndex uint256) returns()
func (_UnitType_Interface *UnitType_InterfaceTransactor) RemoveUnit(opts *bind.TransactOpts, unitIndex *big.Int) (*types.Transaction, error) {
	return _UnitType_Interface.contract.Transact(opts, "removeUnit", unitIndex)
}

// RemoveUnit is a paid mutator transaction binding the contract method 0xff3c6039.
//
// Solidity: function removeUnit(unitIndex uint256) returns()
func (_UnitType_Interface *UnitType_InterfaceSession) RemoveUnit(unitIndex *big.Int) (*types.Transaction, error) {
	return _UnitType_Interface.Contract.RemoveUnit(&_UnitType_Interface.TransactOpts, unitIndex)
}

// RemoveUnit is a paid mutator transaction binding the contract method 0xff3c6039.
//
// Solidity: function removeUnit(unitIndex uint256) returns()
func (_UnitType_Interface *UnitType_InterfaceTransactorSession) RemoveUnit(unitIndex *big.Int) (*types.Transaction, error) {
	return _UnitType_Interface.Contract.RemoveUnit(&_UnitType_Interface.TransactOpts, unitIndex)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_UnitType_Interface *UnitType_InterfaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UnitType_Interface.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_UnitType_Interface *UnitType_InterfaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UnitType_Interface.Contract.TransferOwnership(&_UnitType_Interface.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_UnitType_Interface *UnitType_InterfaceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UnitType_Interface.Contract.TransferOwnership(&_UnitType_Interface.TransactOpts, newOwner)
}
