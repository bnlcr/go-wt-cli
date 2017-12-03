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

// UnitABI is the input ABI used to generate the binding from.
const UnitABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"active\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_currencyCode\",\"type\":\"bytes8\"}],\"name\":\"setCurrencyCode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"fromDay\",\"type\":\"uint256\"},{\"name\":\"daysAmount\",\"type\":\"uint256\"}],\"name\":\"setSpecialPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"defaultLifPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setDefaultPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unitType\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"fromDay\",\"type\":\"uint256\"},{\"name\":\"daysAmount\",\"type\":\"uint256\"}],\"name\":\"getLifCost\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"fromDay\",\"type\":\"uint256\"},{\"name\":\"daysAmount\",\"type\":\"uint256\"}],\"name\":\"getCost\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"setActive\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setDefaultLifPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"fromDay\",\"type\":\"uint256\"},{\"name\":\"daysAmount\",\"type\":\"uint256\"}],\"name\":\"book\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractType\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"day\",\"type\":\"uint256\"}],\"name\":\"getReservation\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currencyCode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"defaultPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"fromDay\",\"type\":\"uint256\"},{\"name\":\"daysAmount\",\"type\":\"uint256\"}],\"name\":\"setSpecialLifPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_unitType\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// UnitBin is the compiled bytecode used for deploying new contracts.
const UnitBin = `0x60606040527f302e302e312d616c7068610000000000000000000000000000000000000000006001906000191690557f756e697400000000000000000000000000000000000000000000000000000000600290600019169055341561006357600080fd5b604051604080610e7083398101604052808051906020019091908051906020019091905050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600381600019169055506001600460006101000a81548160ff0219169083151502179055505050610d318061013f6000396000f300606060405260043610610107576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806302fb0c5e1461010c5780632914fc0a1461013957806346a8e678146101775780634e922d6d146101ac57806354fd4d50146101d55780636d3c7ec51461020657806385ad4f90146102295780638da5cb5b1461025a5780639757876b146102af578063a7c18c14146102ef578063acec338a1461032f578063c55b2ee714610354578063c78e687c14610377578063cb2ef6f7146103da578063d57763c31461040b578063e102e5e31461047c578063e69e04b3146104db578063f2fde38b14610504578063fd480e5f1461053d575b600080fd5b341561011757600080fd5b61011f610572565b604051808215151515815260200191505060405180910390f35b341561014457600080fd5b610175600480803577ffffffffffffffffffffffffffffffffffffffffffffffff1916906020019091905050610585565b005b341561018257600080fd5b6101aa600480803590602001909190803590602001909190803590602001909190505061061e565b005b34156101b757600080fd5b6101bf6106bc565b6040518082815260200191505060405180910390f35b34156101e057600080fd5b6101e86106c2565b60405180826000191660001916815260200191505060405180910390f35b341561021157600080fd5b61022760048080359060200190919050506106c8565b005b341561023457600080fd5b61023c61072d565b60405180826000191660001916815260200191505060405180910390f35b341561026557600080fd5b61026d610733565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156102ba57600080fd5b6102d96004808035906020019091908035906020019091905050610758565b6040518082815260200191505060405180910390f35b34156102fa57600080fd5b61031960048080359060200190919080359060200190919050506107d3565b6040518082815260200191505060405180910390f35b341561033a57600080fd5b6103526004808035151590602001909190505061084e565b005b341561035f57600080fd5b61037560048080359060200190919050506108c6565b005b341561038257600080fd5b6103c0600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190803590602001909190505061092b565b604051808215151515815260200191505060405180910390f35b34156103e557600080fd5b6103ed610ac9565b60405180826000191660001916815260200191505060405180910390f35b341561041657600080fd5b61042c6004808035906020019091905050610acf565b604051808481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001935050505060405180910390f35b341561048757600080fd5b61048f610b46565b604051808277ffffffffffffffffffffffffffffffffffffffffffffffff191677ffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b34156104e657600080fd5b6104ee610b71565b6040518082815260200191505060405180910390f35b341561050f57600080fd5b61053b600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610b77565b005b341561054857600080fd5b6105706004808035906020019091908035906020019091908035906020019091905050610c4c565b005b600460009054906101000a900460ff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156105e057600080fd5b80600660006101000a81548167ffffffffffffffff021916908378010000000000000000000000000000000000000000000000009004021790555050565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561067c57600080fd5b82840191508390505b818110156106b5578460086000838152602001908152602001600020600001819055508080600101915050610685565b5050505050565b60055481565b60015481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561072357600080fd5b8060078190555050565b60035481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000806000808486019250600091508590505b828110156107c7576000600860008381526020019081526020016000206001015411156107b2576008600082815260200190815260200160002060010154820191506107ba565b600554820191505b808060010191505061076b565b81935050505092915050565b6000806000808486019250600091508590505b828110156108425760006008600083815260200190815260200160002060000154111561082d57600860008281526020019081526020016000206000015482019150610835565b600754820191505b80806001019150506107e6565b81935050505092915050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156108a957600080fd5b80600460006101000a81548160ff02191690831515021790555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561092157600080fd5b8060058190555050565b60008060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561098b57600080fd5b61099485610cea565b151561099f57600080fd5b600460009054906101000a900460ff1615156109ba57600080fd5b83850191508490505b81811015610a4d57600073ffffffffffffffffffffffffffffffffffffffff166008600083815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515610a405760009250610ac0565b80806001019150506109c3565b8490505b81811015610abb57856008600083815260200190815260200160002060020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508080600101915050610a51565b600192505b50509392505050565b60025481565b6000806000600860008581526020019081526020016000206000015460086000868152602001908152602001600020600101546008600087815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169250925092509193909250565b600660009054906101000a900478010000000000000000000000000000000000000000000000000281565b60075481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610bd257600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515610c4957806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610caa57600080fd5b82840191508390505b81811015610ce3578460086000838152602001908152602001600020600101819055508080600101915050610cb3565b5050505050565b6000816201518042811515610cfb57fe5b04111590509190505600a165627a7a723058205f9d43d0c00b51040d0a271a0d2c2ee0d4125d21c3c89413de0737add3c0a40b0029`

// DeployUnit deploys a new Ethereum contract, binding an instance of Unit to it.
func DeployUnit(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _unitType [32]byte) (common.Address, *types.Transaction, *Unit, error) {
	parsed, err := abi.JSON(strings.NewReader(UnitABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UnitBin), backend, _owner, _unitType)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Unit{UnitCaller: UnitCaller{contract: contract}, UnitTransactor: UnitTransactor{contract: contract}}, nil
}

// Unit is an auto generated Go binding around an Ethereum contract.
type Unit struct {
	UnitCaller     // Read-only binding to the contract
	UnitTransactor // Write-only binding to the contract
}

// UnitCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnitSession struct {
	Contract     *Unit             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UnitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnitCallerSession struct {
	Contract *UnitCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UnitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnitTransactorSession struct {
	Contract     *UnitTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UnitRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnitRaw struct {
	Contract *Unit // Generic contract binding to access the raw methods on
}

// UnitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnitCallerRaw struct {
	Contract *UnitCaller // Generic read-only contract binding to access the raw methods on
}

// UnitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnitTransactorRaw struct {
	Contract *UnitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnit creates a new instance of Unit, bound to a specific deployed contract.
func NewUnit(address common.Address, backend bind.ContractBackend) (*Unit, error) {
	contract, err := bindUnit(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Unit{UnitCaller: UnitCaller{contract: contract}, UnitTransactor: UnitTransactor{contract: contract}}, nil
}

// NewUnitCaller creates a new read-only instance of Unit, bound to a specific deployed contract.
func NewUnitCaller(address common.Address, caller bind.ContractCaller) (*UnitCaller, error) {
	contract, err := bindUnit(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &UnitCaller{contract: contract}, nil
}

// NewUnitTransactor creates a new write-only instance of Unit, bound to a specific deployed contract.
func NewUnitTransactor(address common.Address, transactor bind.ContractTransactor) (*UnitTransactor, error) {
	contract, err := bindUnit(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &UnitTransactor{contract: contract}, nil
}

// bindUnit binds a generic wrapper to an already deployed contract.
func bindUnit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UnitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Unit *UnitRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Unit.Contract.UnitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Unit *UnitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Unit.Contract.UnitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Unit *UnitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Unit.Contract.UnitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Unit *UnitCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Unit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Unit *UnitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Unit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Unit *UnitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Unit.Contract.contract.Transact(opts, method, params...)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() constant returns(bool)
func (_Unit *UnitCaller) Active(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Unit.contract.Call(opts, out, "active")
	return *ret0, err
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() constant returns(bool)
func (_Unit *UnitSession) Active() (bool, error) {
	return _Unit.Contract.Active(&_Unit.CallOpts)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() constant returns(bool)
func (_Unit *UnitCallerSession) Active() (bool, error) {
	return _Unit.Contract.Active(&_Unit.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() constant returns(bytes32)
func (_Unit *UnitCaller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Unit.contract.Call(opts, out, "contractType")
	return *ret0, err
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() constant returns(bytes32)
func (_Unit *UnitSession) ContractType() ([32]byte, error) {
	return _Unit.Contract.ContractType(&_Unit.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() constant returns(bytes32)
func (_Unit *UnitCallerSession) ContractType() ([32]byte, error) {
	return _Unit.Contract.ContractType(&_Unit.CallOpts)
}

// CurrencyCode is a free data retrieval call binding the contract method 0xe102e5e3.
//
// Solidity: function currencyCode() constant returns(bytes8)
func (_Unit *UnitCaller) CurrencyCode(opts *bind.CallOpts) ([8]byte, error) {
	var (
		ret0 = new([8]byte)
	)
	out := ret0
	err := _Unit.contract.Call(opts, out, "currencyCode")
	return *ret0, err
}

// CurrencyCode is a free data retrieval call binding the contract method 0xe102e5e3.
//
// Solidity: function currencyCode() constant returns(bytes8)
func (_Unit *UnitSession) CurrencyCode() ([8]byte, error) {
	return _Unit.Contract.CurrencyCode(&_Unit.CallOpts)
}

// CurrencyCode is a free data retrieval call binding the contract method 0xe102e5e3.
//
// Solidity: function currencyCode() constant returns(bytes8)
func (_Unit *UnitCallerSession) CurrencyCode() ([8]byte, error) {
	return _Unit.Contract.CurrencyCode(&_Unit.CallOpts)
}

// DefaultLifPrice is a free data retrieval call binding the contract method 0x4e922d6d.
//
// Solidity: function defaultLifPrice() constant returns(uint256)
func (_Unit *UnitCaller) DefaultLifPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Unit.contract.Call(opts, out, "defaultLifPrice")
	return *ret0, err
}

// DefaultLifPrice is a free data retrieval call binding the contract method 0x4e922d6d.
//
// Solidity: function defaultLifPrice() constant returns(uint256)
func (_Unit *UnitSession) DefaultLifPrice() (*big.Int, error) {
	return _Unit.Contract.DefaultLifPrice(&_Unit.CallOpts)
}

// DefaultLifPrice is a free data retrieval call binding the contract method 0x4e922d6d.
//
// Solidity: function defaultLifPrice() constant returns(uint256)
func (_Unit *UnitCallerSession) DefaultLifPrice() (*big.Int, error) {
	return _Unit.Contract.DefaultLifPrice(&_Unit.CallOpts)
}

// DefaultPrice is a free data retrieval call binding the contract method 0xe69e04b3.
//
// Solidity: function defaultPrice() constant returns(uint256)
func (_Unit *UnitCaller) DefaultPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Unit.contract.Call(opts, out, "defaultPrice")
	return *ret0, err
}

// DefaultPrice is a free data retrieval call binding the contract method 0xe69e04b3.
//
// Solidity: function defaultPrice() constant returns(uint256)
func (_Unit *UnitSession) DefaultPrice() (*big.Int, error) {
	return _Unit.Contract.DefaultPrice(&_Unit.CallOpts)
}

// DefaultPrice is a free data retrieval call binding the contract method 0xe69e04b3.
//
// Solidity: function defaultPrice() constant returns(uint256)
func (_Unit *UnitCallerSession) DefaultPrice() (*big.Int, error) {
	return _Unit.Contract.DefaultPrice(&_Unit.CallOpts)
}

// GetCost is a free data retrieval call binding the contract method 0xa7c18c14.
//
// Solidity: function getCost(fromDay uint256, daysAmount uint256) constant returns(uint256)
func (_Unit *UnitCaller) GetCost(opts *bind.CallOpts, fromDay *big.Int, daysAmount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Unit.contract.Call(opts, out, "getCost", fromDay, daysAmount)
	return *ret0, err
}

// GetCost is a free data retrieval call binding the contract method 0xa7c18c14.
//
// Solidity: function getCost(fromDay uint256, daysAmount uint256) constant returns(uint256)
func (_Unit *UnitSession) GetCost(fromDay *big.Int, daysAmount *big.Int) (*big.Int, error) {
	return _Unit.Contract.GetCost(&_Unit.CallOpts, fromDay, daysAmount)
}

// GetCost is a free data retrieval call binding the contract method 0xa7c18c14.
//
// Solidity: function getCost(fromDay uint256, daysAmount uint256) constant returns(uint256)
func (_Unit *UnitCallerSession) GetCost(fromDay *big.Int, daysAmount *big.Int) (*big.Int, error) {
	return _Unit.Contract.GetCost(&_Unit.CallOpts, fromDay, daysAmount)
}

// GetLifCost is a free data retrieval call binding the contract method 0x9757876b.
//
// Solidity: function getLifCost(fromDay uint256, daysAmount uint256) constant returns(uint256)
func (_Unit *UnitCaller) GetLifCost(opts *bind.CallOpts, fromDay *big.Int, daysAmount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Unit.contract.Call(opts, out, "getLifCost", fromDay, daysAmount)
	return *ret0, err
}

// GetLifCost is a free data retrieval call binding the contract method 0x9757876b.
//
// Solidity: function getLifCost(fromDay uint256, daysAmount uint256) constant returns(uint256)
func (_Unit *UnitSession) GetLifCost(fromDay *big.Int, daysAmount *big.Int) (*big.Int, error) {
	return _Unit.Contract.GetLifCost(&_Unit.CallOpts, fromDay, daysAmount)
}

// GetLifCost is a free data retrieval call binding the contract method 0x9757876b.
//
// Solidity: function getLifCost(fromDay uint256, daysAmount uint256) constant returns(uint256)
func (_Unit *UnitCallerSession) GetLifCost(fromDay *big.Int, daysAmount *big.Int) (*big.Int, error) {
	return _Unit.Contract.GetLifCost(&_Unit.CallOpts, fromDay, daysAmount)
}

// GetReservation is a free data retrieval call binding the contract method 0xd57763c3.
//
// Solidity: function getReservation(day uint256) constant returns(uint256, uint256, address)
func (_Unit *UnitCaller) GetReservation(opts *bind.CallOpts, day *big.Int) (*big.Int, *big.Int, common.Address, error) {
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
	err := _Unit.contract.Call(opts, out, "getReservation", day)
	return *ret0, *ret1, *ret2, err
}

// GetReservation is a free data retrieval call binding the contract method 0xd57763c3.
//
// Solidity: function getReservation(day uint256) constant returns(uint256, uint256, address)
func (_Unit *UnitSession) GetReservation(day *big.Int) (*big.Int, *big.Int, common.Address, error) {
	return _Unit.Contract.GetReservation(&_Unit.CallOpts, day)
}

// GetReservation is a free data retrieval call binding the contract method 0xd57763c3.
//
// Solidity: function getReservation(day uint256) constant returns(uint256, uint256, address)
func (_Unit *UnitCallerSession) GetReservation(day *big.Int) (*big.Int, *big.Int, common.Address, error) {
	return _Unit.Contract.GetReservation(&_Unit.CallOpts, day)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Unit *UnitCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Unit.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Unit *UnitSession) Owner() (common.Address, error) {
	return _Unit.Contract.Owner(&_Unit.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Unit *UnitCallerSession) Owner() (common.Address, error) {
	return _Unit.Contract.Owner(&_Unit.CallOpts)
}

// UnitType is a free data retrieval call binding the contract method 0x85ad4f90.
//
// Solidity: function unitType() constant returns(bytes32)
func (_Unit *UnitCaller) UnitType(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Unit.contract.Call(opts, out, "unitType")
	return *ret0, err
}

// UnitType is a free data retrieval call binding the contract method 0x85ad4f90.
//
// Solidity: function unitType() constant returns(bytes32)
func (_Unit *UnitSession) UnitType() ([32]byte, error) {
	return _Unit.Contract.UnitType(&_Unit.CallOpts)
}

// UnitType is a free data retrieval call binding the contract method 0x85ad4f90.
//
// Solidity: function unitType() constant returns(bytes32)
func (_Unit *UnitCallerSession) UnitType() ([32]byte, error) {
	return _Unit.Contract.UnitType(&_Unit.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Unit *UnitCaller) Version(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Unit.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Unit *UnitSession) Version() ([32]byte, error) {
	return _Unit.Contract.Version(&_Unit.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Unit *UnitCallerSession) Version() ([32]byte, error) {
	return _Unit.Contract.Version(&_Unit.CallOpts)
}

// Book is a paid mutator transaction binding the contract method 0xc78e687c.
//
// Solidity: function book(from address, fromDay uint256, daysAmount uint256) returns(bool)
func (_Unit *UnitTransactor) Book(opts *bind.TransactOpts, from common.Address, fromDay *big.Int, daysAmount *big.Int) (*types.Transaction, error) {
	return _Unit.contract.Transact(opts, "book", from, fromDay, daysAmount)
}

// Book is a paid mutator transaction binding the contract method 0xc78e687c.
//
// Solidity: function book(from address, fromDay uint256, daysAmount uint256) returns(bool)
func (_Unit *UnitSession) Book(from common.Address, fromDay *big.Int, daysAmount *big.Int) (*types.Transaction, error) {
	return _Unit.Contract.Book(&_Unit.TransactOpts, from, fromDay, daysAmount)
}

// Book is a paid mutator transaction binding the contract method 0xc78e687c.
//
// Solidity: function book(from address, fromDay uint256, daysAmount uint256) returns(bool)
func (_Unit *UnitTransactorSession) Book(from common.Address, fromDay *big.Int, daysAmount *big.Int) (*types.Transaction, error) {
	return _Unit.Contract.Book(&_Unit.TransactOpts, from, fromDay, daysAmount)
}

// SetActive is a paid mutator transaction binding the contract method 0xacec338a.
//
// Solidity: function setActive(_active bool) returns()
func (_Unit *UnitTransactor) SetActive(opts *bind.TransactOpts, _active bool) (*types.Transaction, error) {
	return _Unit.contract.Transact(opts, "setActive", _active)
}

// SetActive is a paid mutator transaction binding the contract method 0xacec338a.
//
// Solidity: function setActive(_active bool) returns()
func (_Unit *UnitSession) SetActive(_active bool) (*types.Transaction, error) {
	return _Unit.Contract.SetActive(&_Unit.TransactOpts, _active)
}

// SetActive is a paid mutator transaction binding the contract method 0xacec338a.
//
// Solidity: function setActive(_active bool) returns()
func (_Unit *UnitTransactorSession) SetActive(_active bool) (*types.Transaction, error) {
	return _Unit.Contract.SetActive(&_Unit.TransactOpts, _active)
}

// SetCurrencyCode is a paid mutator transaction binding the contract method 0x2914fc0a.
//
// Solidity: function setCurrencyCode(_currencyCode bytes8) returns()
func (_Unit *UnitTransactor) SetCurrencyCode(opts *bind.TransactOpts, _currencyCode [8]byte) (*types.Transaction, error) {
	return _Unit.contract.Transact(opts, "setCurrencyCode", _currencyCode)
}

// SetCurrencyCode is a paid mutator transaction binding the contract method 0x2914fc0a.
//
// Solidity: function setCurrencyCode(_currencyCode bytes8) returns()
func (_Unit *UnitSession) SetCurrencyCode(_currencyCode [8]byte) (*types.Transaction, error) {
	return _Unit.Contract.SetCurrencyCode(&_Unit.TransactOpts, _currencyCode)
}

// SetCurrencyCode is a paid mutator transaction binding the contract method 0x2914fc0a.
//
// Solidity: function setCurrencyCode(_currencyCode bytes8) returns()
func (_Unit *UnitTransactorSession) SetCurrencyCode(_currencyCode [8]byte) (*types.Transaction, error) {
	return _Unit.Contract.SetCurrencyCode(&_Unit.TransactOpts, _currencyCode)
}

// SetDefaultLifPrice is a paid mutator transaction binding the contract method 0xc55b2ee7.
//
// Solidity: function setDefaultLifPrice(price uint256) returns()
func (_Unit *UnitTransactor) SetDefaultLifPrice(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _Unit.contract.Transact(opts, "setDefaultLifPrice", price)
}

// SetDefaultLifPrice is a paid mutator transaction binding the contract method 0xc55b2ee7.
//
// Solidity: function setDefaultLifPrice(price uint256) returns()
func (_Unit *UnitSession) SetDefaultLifPrice(price *big.Int) (*types.Transaction, error) {
	return _Unit.Contract.SetDefaultLifPrice(&_Unit.TransactOpts, price)
}

// SetDefaultLifPrice is a paid mutator transaction binding the contract method 0xc55b2ee7.
//
// Solidity: function setDefaultLifPrice(price uint256) returns()
func (_Unit *UnitTransactorSession) SetDefaultLifPrice(price *big.Int) (*types.Transaction, error) {
	return _Unit.Contract.SetDefaultLifPrice(&_Unit.TransactOpts, price)
}

// SetDefaultPrice is a paid mutator transaction binding the contract method 0x6d3c7ec5.
//
// Solidity: function setDefaultPrice(price uint256) returns()
func (_Unit *UnitTransactor) SetDefaultPrice(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _Unit.contract.Transact(opts, "setDefaultPrice", price)
}

// SetDefaultPrice is a paid mutator transaction binding the contract method 0x6d3c7ec5.
//
// Solidity: function setDefaultPrice(price uint256) returns()
func (_Unit *UnitSession) SetDefaultPrice(price *big.Int) (*types.Transaction, error) {
	return _Unit.Contract.SetDefaultPrice(&_Unit.TransactOpts, price)
}

// SetDefaultPrice is a paid mutator transaction binding the contract method 0x6d3c7ec5.
//
// Solidity: function setDefaultPrice(price uint256) returns()
func (_Unit *UnitTransactorSession) SetDefaultPrice(price *big.Int) (*types.Transaction, error) {
	return _Unit.Contract.SetDefaultPrice(&_Unit.TransactOpts, price)
}

// SetSpecialLifPrice is a paid mutator transaction binding the contract method 0xfd480e5f.
//
// Solidity: function setSpecialLifPrice(price uint256, fromDay uint256, daysAmount uint256) returns()
func (_Unit *UnitTransactor) SetSpecialLifPrice(opts *bind.TransactOpts, price *big.Int, fromDay *big.Int, daysAmount *big.Int) (*types.Transaction, error) {
	return _Unit.contract.Transact(opts, "setSpecialLifPrice", price, fromDay, daysAmount)
}

// SetSpecialLifPrice is a paid mutator transaction binding the contract method 0xfd480e5f.
//
// Solidity: function setSpecialLifPrice(price uint256, fromDay uint256, daysAmount uint256) returns()
func (_Unit *UnitSession) SetSpecialLifPrice(price *big.Int, fromDay *big.Int, daysAmount *big.Int) (*types.Transaction, error) {
	return _Unit.Contract.SetSpecialLifPrice(&_Unit.TransactOpts, price, fromDay, daysAmount)
}

// SetSpecialLifPrice is a paid mutator transaction binding the contract method 0xfd480e5f.
//
// Solidity: function setSpecialLifPrice(price uint256, fromDay uint256, daysAmount uint256) returns()
func (_Unit *UnitTransactorSession) SetSpecialLifPrice(price *big.Int, fromDay *big.Int, daysAmount *big.Int) (*types.Transaction, error) {
	return _Unit.Contract.SetSpecialLifPrice(&_Unit.TransactOpts, price, fromDay, daysAmount)
}

// SetSpecialPrice is a paid mutator transaction binding the contract method 0x46a8e678.
//
// Solidity: function setSpecialPrice(price uint256, fromDay uint256, daysAmount uint256) returns()
func (_Unit *UnitTransactor) SetSpecialPrice(opts *bind.TransactOpts, price *big.Int, fromDay *big.Int, daysAmount *big.Int) (*types.Transaction, error) {
	return _Unit.contract.Transact(opts, "setSpecialPrice", price, fromDay, daysAmount)
}

// SetSpecialPrice is a paid mutator transaction binding the contract method 0x46a8e678.
//
// Solidity: function setSpecialPrice(price uint256, fromDay uint256, daysAmount uint256) returns()
func (_Unit *UnitSession) SetSpecialPrice(price *big.Int, fromDay *big.Int, daysAmount *big.Int) (*types.Transaction, error) {
	return _Unit.Contract.SetSpecialPrice(&_Unit.TransactOpts, price, fromDay, daysAmount)
}

// SetSpecialPrice is a paid mutator transaction binding the contract method 0x46a8e678.
//
// Solidity: function setSpecialPrice(price uint256, fromDay uint256, daysAmount uint256) returns()
func (_Unit *UnitTransactorSession) SetSpecialPrice(price *big.Int, fromDay *big.Int, daysAmount *big.Int) (*types.Transaction, error) {
	return _Unit.Contract.SetSpecialPrice(&_Unit.TransactOpts, price, fromDay, daysAmount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Unit *UnitTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Unit.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Unit *UnitSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Unit.Contract.TransferOwnership(&_Unit.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Unit *UnitTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Unit.Contract.TransferOwnership(&_Unit.TransactOpts, newOwner)
}
