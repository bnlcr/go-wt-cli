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

// VestedPaymentABI is the input ABI used to generate the binding from.
const VestedPaymentABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"secondsPerPeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokens\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cliffDuration\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAvailableTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"claimed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalPeriods\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_startTimestamp\",\"type\":\"uint256\"},{\"name\":\"_secondsPerPeriod\",\"type\":\"uint256\"},{\"name\":\"_totalPeriods\",\"type\":\"uint256\"},{\"name\":\"_cliffDuration\",\"type\":\"uint256\"},{\"name\":\"_tokens\",\"type\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// VestedPaymentBin is the compiled bytecode used for deploying new contracts.
const VestedPaymentBin = `0x6060604052341561000f57600080fd5b60405160c0806108bc83398101604052808051906020019091908051906020019091908051906020019091908051906020019091908051906020019091908051906020019091905050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055504286101515156100a757600080fd5b6000851115156100b657600080fd5b6000841115156100c557600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561010157600080fd5b838310151561010f57600080fd5b60008211151561011e57600080fd5b856001819055508460028190555083600381905550826007819055508160048190555080600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050505050610725806101976000396000f3006060604052600436106100af576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063407f8001146100b457806346e04a2f146100dd5780638da5cb5b146101005780639d63848a14610155578063d85349f71461017e578063e35568cb146101a7578063e6fd48bc146101d0578063e834a834146101f9578063f2fde38b14610222578063fc0c546a1461025b578063fea708f6146102b0575b600080fd5b34156100bf57600080fd5b6100c76102d9565b6040518082815260200191505060405180910390f35b34156100e857600080fd5b6100fe60048080359060200190919050506102df565b005b341561010b57600080fd5b610113610475565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561016057600080fd5b61016861049a565b6040518082815260200191505060405180910390f35b341561018957600080fd5b6101916104a0565b6040518082815260200191505060405180910390f35b34156101b257600080fd5b6101ba6104a6565b6040518082815260200191505060405180910390f35b34156101db57600080fd5b6101e3610567565b6040518082815260200191505060405180910390f35b341561020457600080fd5b61020c61056d565b6040518082815260200191505060405180910390f35b341561022d57600080fd5b610259600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610573565b005b341561026657600080fd5b61026e610648565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156102bb57600080fd5b6102c361066e565b6040518082815260200191505060405180910390f35b60025481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561033a57600080fd5b806103436104a6565b1015151561034d57fe5b6103628160055461067490919063ffffffff16565b600581905550600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836000604051602001526040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b151561045657600080fd5b6102c65a03f1151561046757600080fd5b505050604051805190505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60045481565b60075481565b6000806104d26002546104c46001544261069290919063ffffffff16565b6106ab90919063ffffffff16565b90506007548110156104e75760009150610563565b600354811015156105105761050960055460045461069290919063ffffffff16565b9150610563565b61056060055461055260035461054461053360018761067490919063ffffffff16565b6004546106c690919063ffffffff16565b6106ab90919063ffffffff16565b61069290919063ffffffff16565b91505b5090565b60015481565b60055481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156105ce57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151561064557806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60035481565b600080828401905083811015151561068857fe5b8091505092915050565b60008282111515156106a057fe5b818303905092915050565b60008082848115156106b957fe5b0490508091505092915050565b600080828402905060008414806106e757508284828115156106e457fe5b04145b15156106ef57fe5b80915050929150505600a165627a7a72305820fe948aae26901650967a331d5c6ff5879973e4f281ec779d144345b151f70fdb0029`

// DeployVestedPayment deploys a new Ethereum contract, binding an instance of VestedPayment to it.
func DeployVestedPayment(auth *bind.TransactOpts, backend bind.ContractBackend, _startTimestamp *big.Int, _secondsPerPeriod *big.Int, _totalPeriods *big.Int, _cliffDuration *big.Int, _tokens *big.Int, tokenAddress common.Address) (common.Address, *types.Transaction, *VestedPayment, error) {
	parsed, err := abi.JSON(strings.NewReader(VestedPaymentABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VestedPaymentBin), backend, _startTimestamp, _secondsPerPeriod, _totalPeriods, _cliffDuration, _tokens, tokenAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VestedPayment{VestedPaymentCaller: VestedPaymentCaller{contract: contract}, VestedPaymentTransactor: VestedPaymentTransactor{contract: contract}}, nil
}

// VestedPayment is an auto generated Go binding around an Ethereum contract.
type VestedPayment struct {
	VestedPaymentCaller     // Read-only binding to the contract
	VestedPaymentTransactor // Write-only binding to the contract
}

// VestedPaymentCaller is an auto generated read-only Go binding around an Ethereum contract.
type VestedPaymentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VestedPaymentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VestedPaymentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VestedPaymentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VestedPaymentSession struct {
	Contract     *VestedPayment    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VestedPaymentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VestedPaymentCallerSession struct {
	Contract *VestedPaymentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// VestedPaymentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VestedPaymentTransactorSession struct {
	Contract     *VestedPaymentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// VestedPaymentRaw is an auto generated low-level Go binding around an Ethereum contract.
type VestedPaymentRaw struct {
	Contract *VestedPayment // Generic contract binding to access the raw methods on
}

// VestedPaymentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VestedPaymentCallerRaw struct {
	Contract *VestedPaymentCaller // Generic read-only contract binding to access the raw methods on
}

// VestedPaymentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VestedPaymentTransactorRaw struct {
	Contract *VestedPaymentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVestedPayment creates a new instance of VestedPayment, bound to a specific deployed contract.
func NewVestedPayment(address common.Address, backend bind.ContractBackend) (*VestedPayment, error) {
	contract, err := bindVestedPayment(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VestedPayment{VestedPaymentCaller: VestedPaymentCaller{contract: contract}, VestedPaymentTransactor: VestedPaymentTransactor{contract: contract}}, nil
}

// NewVestedPaymentCaller creates a new read-only instance of VestedPayment, bound to a specific deployed contract.
func NewVestedPaymentCaller(address common.Address, caller bind.ContractCaller) (*VestedPaymentCaller, error) {
	contract, err := bindVestedPayment(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &VestedPaymentCaller{contract: contract}, nil
}

// NewVestedPaymentTransactor creates a new write-only instance of VestedPayment, bound to a specific deployed contract.
func NewVestedPaymentTransactor(address common.Address, transactor bind.ContractTransactor) (*VestedPaymentTransactor, error) {
	contract, err := bindVestedPayment(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &VestedPaymentTransactor{contract: contract}, nil
}

// bindVestedPayment binds a generic wrapper to an already deployed contract.
func bindVestedPayment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VestedPaymentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VestedPayment *VestedPaymentRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VestedPayment.Contract.VestedPaymentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VestedPayment *VestedPaymentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VestedPayment.Contract.VestedPaymentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VestedPayment *VestedPaymentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VestedPayment.Contract.VestedPaymentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VestedPayment *VestedPaymentCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VestedPayment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VestedPayment *VestedPaymentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VestedPayment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VestedPayment *VestedPaymentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VestedPayment.Contract.contract.Transact(opts, method, params...)
}

// Claimed is a free data retrieval call binding the contract method 0xe834a834.
//
// Solidity: function claimed() constant returns(uint256)
func (_VestedPayment *VestedPaymentCaller) Claimed(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VestedPayment.contract.Call(opts, out, "claimed")
	return *ret0, err
}

// Claimed is a free data retrieval call binding the contract method 0xe834a834.
//
// Solidity: function claimed() constant returns(uint256)
func (_VestedPayment *VestedPaymentSession) Claimed() (*big.Int, error) {
	return _VestedPayment.Contract.Claimed(&_VestedPayment.CallOpts)
}

// Claimed is a free data retrieval call binding the contract method 0xe834a834.
//
// Solidity: function claimed() constant returns(uint256)
func (_VestedPayment *VestedPaymentCallerSession) Claimed() (*big.Int, error) {
	return _VestedPayment.Contract.Claimed(&_VestedPayment.CallOpts)
}

// CliffDuration is a free data retrieval call binding the contract method 0xd85349f7.
//
// Solidity: function cliffDuration() constant returns(uint256)
func (_VestedPayment *VestedPaymentCaller) CliffDuration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VestedPayment.contract.Call(opts, out, "cliffDuration")
	return *ret0, err
}

// CliffDuration is a free data retrieval call binding the contract method 0xd85349f7.
//
// Solidity: function cliffDuration() constant returns(uint256)
func (_VestedPayment *VestedPaymentSession) CliffDuration() (*big.Int, error) {
	return _VestedPayment.Contract.CliffDuration(&_VestedPayment.CallOpts)
}

// CliffDuration is a free data retrieval call binding the contract method 0xd85349f7.
//
// Solidity: function cliffDuration() constant returns(uint256)
func (_VestedPayment *VestedPaymentCallerSession) CliffDuration() (*big.Int, error) {
	return _VestedPayment.Contract.CliffDuration(&_VestedPayment.CallOpts)
}

// GetAvailableTokens is a free data retrieval call binding the contract method 0xe35568cb.
//
// Solidity: function getAvailableTokens() constant returns(uint256)
func (_VestedPayment *VestedPaymentCaller) GetAvailableTokens(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VestedPayment.contract.Call(opts, out, "getAvailableTokens")
	return *ret0, err
}

// GetAvailableTokens is a free data retrieval call binding the contract method 0xe35568cb.
//
// Solidity: function getAvailableTokens() constant returns(uint256)
func (_VestedPayment *VestedPaymentSession) GetAvailableTokens() (*big.Int, error) {
	return _VestedPayment.Contract.GetAvailableTokens(&_VestedPayment.CallOpts)
}

// GetAvailableTokens is a free data retrieval call binding the contract method 0xe35568cb.
//
// Solidity: function getAvailableTokens() constant returns(uint256)
func (_VestedPayment *VestedPaymentCallerSession) GetAvailableTokens() (*big.Int, error) {
	return _VestedPayment.Contract.GetAvailableTokens(&_VestedPayment.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_VestedPayment *VestedPaymentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VestedPayment.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_VestedPayment *VestedPaymentSession) Owner() (common.Address, error) {
	return _VestedPayment.Contract.Owner(&_VestedPayment.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_VestedPayment *VestedPaymentCallerSession) Owner() (common.Address, error) {
	return _VestedPayment.Contract.Owner(&_VestedPayment.CallOpts)
}

// SecondsPerPeriod is a free data retrieval call binding the contract method 0x407f8001.
//
// Solidity: function secondsPerPeriod() constant returns(uint256)
func (_VestedPayment *VestedPaymentCaller) SecondsPerPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VestedPayment.contract.Call(opts, out, "secondsPerPeriod")
	return *ret0, err
}

// SecondsPerPeriod is a free data retrieval call binding the contract method 0x407f8001.
//
// Solidity: function secondsPerPeriod() constant returns(uint256)
func (_VestedPayment *VestedPaymentSession) SecondsPerPeriod() (*big.Int, error) {
	return _VestedPayment.Contract.SecondsPerPeriod(&_VestedPayment.CallOpts)
}

// SecondsPerPeriod is a free data retrieval call binding the contract method 0x407f8001.
//
// Solidity: function secondsPerPeriod() constant returns(uint256)
func (_VestedPayment *VestedPaymentCallerSession) SecondsPerPeriod() (*big.Int, error) {
	return _VestedPayment.Contract.SecondsPerPeriod(&_VestedPayment.CallOpts)
}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() constant returns(uint256)
func (_VestedPayment *VestedPaymentCaller) StartTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VestedPayment.contract.Call(opts, out, "startTimestamp")
	return *ret0, err
}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() constant returns(uint256)
func (_VestedPayment *VestedPaymentSession) StartTimestamp() (*big.Int, error) {
	return _VestedPayment.Contract.StartTimestamp(&_VestedPayment.CallOpts)
}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() constant returns(uint256)
func (_VestedPayment *VestedPaymentCallerSession) StartTimestamp() (*big.Int, error) {
	return _VestedPayment.Contract.StartTimestamp(&_VestedPayment.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_VestedPayment *VestedPaymentCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VestedPayment.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_VestedPayment *VestedPaymentSession) Token() (common.Address, error) {
	return _VestedPayment.Contract.Token(&_VestedPayment.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_VestedPayment *VestedPaymentCallerSession) Token() (common.Address, error) {
	return _VestedPayment.Contract.Token(&_VestedPayment.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(uint256)
func (_VestedPayment *VestedPaymentCaller) Tokens(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VestedPayment.contract.Call(opts, out, "tokens")
	return *ret0, err
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(uint256)
func (_VestedPayment *VestedPaymentSession) Tokens() (*big.Int, error) {
	return _VestedPayment.Contract.Tokens(&_VestedPayment.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(uint256)
func (_VestedPayment *VestedPaymentCallerSession) Tokens() (*big.Int, error) {
	return _VestedPayment.Contract.Tokens(&_VestedPayment.CallOpts)
}

// TotalPeriods is a free data retrieval call binding the contract method 0xfea708f6.
//
// Solidity: function totalPeriods() constant returns(uint256)
func (_VestedPayment *VestedPaymentCaller) TotalPeriods(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VestedPayment.contract.Call(opts, out, "totalPeriods")
	return *ret0, err
}

// TotalPeriods is a free data retrieval call binding the contract method 0xfea708f6.
//
// Solidity: function totalPeriods() constant returns(uint256)
func (_VestedPayment *VestedPaymentSession) TotalPeriods() (*big.Int, error) {
	return _VestedPayment.Contract.TotalPeriods(&_VestedPayment.CallOpts)
}

// TotalPeriods is a free data retrieval call binding the contract method 0xfea708f6.
//
// Solidity: function totalPeriods() constant returns(uint256)
func (_VestedPayment *VestedPaymentCallerSession) TotalPeriods() (*big.Int, error) {
	return _VestedPayment.Contract.TotalPeriods(&_VestedPayment.CallOpts)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x46e04a2f.
//
// Solidity: function claimTokens(amount uint256) returns()
func (_VestedPayment *VestedPaymentTransactor) ClaimTokens(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _VestedPayment.contract.Transact(opts, "claimTokens", amount)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x46e04a2f.
//
// Solidity: function claimTokens(amount uint256) returns()
func (_VestedPayment *VestedPaymentSession) ClaimTokens(amount *big.Int) (*types.Transaction, error) {
	return _VestedPayment.Contract.ClaimTokens(&_VestedPayment.TransactOpts, amount)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x46e04a2f.
//
// Solidity: function claimTokens(amount uint256) returns()
func (_VestedPayment *VestedPaymentTransactorSession) ClaimTokens(amount *big.Int) (*types.Transaction, error) {
	return _VestedPayment.Contract.ClaimTokens(&_VestedPayment.TransactOpts, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_VestedPayment *VestedPaymentTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VestedPayment.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_VestedPayment *VestedPaymentSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VestedPayment.Contract.TransferOwnership(&_VestedPayment.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_VestedPayment *VestedPaymentTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VestedPayment.Contract.TransferOwnership(&_VestedPayment.TransactOpts, newOwner)
}
