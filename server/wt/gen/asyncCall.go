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

// AsyncCallABI is the input ABI used to generate the binding from.
const AsyncCallABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"publicCallData\",\"type\":\"bytes\"},{\"name\":\"privateData\",\"type\":\"bytes\"}],\"name\":\"beginCall\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pendingCalls\",\"outputs\":[{\"name\":\"callData\",\"type\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"},{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"msgDataHash\",\"type\":\"bytes32\"}],\"name\":\"continueCall\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_waitConfirmation\",\"type\":\"bool\"}],\"name\":\"changeConfirmation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"waitConfirmation\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractType\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"CallStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"CallFinish\",\"type\":\"event\"}]"

// AsyncCallBin is the compiled bytecode used for deploying new contracts.
const AsyncCallBin = `0x60606040527f302e302e312d616c7068610000000000000000000000000000000000000000006001906000191690557f7072697661746563616c6c000000000000000000000000000000000000000000600290600019169055336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610d0d806100a76000396000f300606060405260043610610099576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063203eaf271461009e57806332fdd45c1461013e578063411f23511461024457806344f5484a1461026b57806354fd4d50146102905780638da5cb5b146102c1578063ac72bb9814610316578063cb2ef6f714610343578063f2fde38b14610374575b600080fd5b34156100a957600080fd5b61013c600480803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919050506103ad565b005b341561014957600080fd5b61016360048080356000191690602001909190505061078c565b60405180806020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200184151515158152602001831515151581526020018281038252868181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156102325780601f1061020757610100808354040283529160200191610232565b820191906000526020600020905b81548152906001019060200180831161021557829003601f168201915b50509550505050505060405180910390f35b341561024f57600080fd5b6102696004808035600019169060200190919050506107f5565b005b341561027657600080fd5b61028e60048080351515906020019091905050610aab565b005b341561029b57600080fd5b6102a3610b23565b60405180826000191660001916815260200191505060405180910390f35b34156102cc57600080fd5b6102d4610b29565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561032157600080fd5b610329610b4e565b604051808215151515815260200191505060405180910390f35b341561034e57600080fd5b610356610b61565b60405180826000191660001916815260200191505060405180910390f35b341561037f57600080fd5b6103ab600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610b67565b005b600080366040518083838082843782019150509250505060405180910390209050600073ffffffffffffffffffffffffffffffffffffffff1660036000836000191660001916815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614151561044757600080fd5b6080604051908101604052808481526020013273ffffffffffffffffffffffffffffffffffffffff168152602001600460009054906101000a900460ff1615151581526020016000151581525060036000836000191660001916815260200190815260200160002060008201518160000190805190602001906104cb929190610c3c565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160146101000a81548160ff02191690831515021790555060608201518160010160156101000a81548160ff0219169083151502179055509050507f5ec33322685848c709cbf116179bd3646c27f9e2d88123679b424506db62f70b3282604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182600019166000191681526020019250505060405180910390a1600460009054906101000a900460ff161515610787573073ffffffffffffffffffffffffffffffffffffffff1660036000836000191660001916815260200190815260200160002060000160405180828054600181600116156101000203166002900480156106795780601f1061064e57610100808354040283529160200191610679565b820191906000526020600020905b81548152906001019060200180831161065c57829003601f168201915b505091505060006040518083038160008661646e5a03f1915050151561069e57600080fd5b600160036000836000191660001916815260200190815260200160002060010160156101000a81548160ff0219169083151502179055507fb2be3874ffa3a459ad88116e697afbf466ac94ca363d9f32b0701777983b51f760036000836000191660001916815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182600019166000191681526020019250505060405180910390a15b505050565b600360205280600052604060002060009150905080600001908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160149054906101000a900460ff16908060010160159054906101000a900460ff16905084565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561085057600080fd5b600073ffffffffffffffffffffffffffffffffffffffff1660036000836000191660001916815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515156108ca57600080fd5b600160036000836000191660001916815260200190815260200160002060010160146101000a81548160ff0219169083151502179055503073ffffffffffffffffffffffffffffffffffffffff16600360008360001916600019168152602001908152602001600020600001604051808280546001816001161561010002031660029004801561099b5780601f106109705761010080835404028352916020019161099b565b820191906000526020600020905b81548152906001019060200180831161097e57829003601f168201915b505091505060006040518083038160008661646e5a03f191505015156109c057600080fd5b600160036000836000191660001916815260200190815260200160002060010160156101000a81548160ff0219169083151502179055507fb2be3874ffa3a459ad88116e697afbf466ac94ca363d9f32b0701777983b51f760036000836000191660001916815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182600019166000191681526020019250505060405180910390a150565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610b0657600080fd5b80600460006101000a81548160ff02191690831515021790555050565b60015481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600460009054906101000a900460ff1681565b60025481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610bc257600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515610c3957806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610c7d57805160ff1916838001178555610cab565b82800160010185558215610cab579182015b82811115610caa578251825591602001919060010190610c8f565b5b509050610cb89190610cbc565b5090565b610cde91905b80821115610cda576000816000905550600101610cc2565b5090565b905600a165627a7a723058200d5ab9f39cc51fb4fa44f25eece138294abffe6cfe0b8ec531a2b091707b93eb0029`

// DeployAsyncCall deploys a new Ethereum contract, binding an instance of AsyncCall to it.
func DeployAsyncCall(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AsyncCall, error) {
	parsed, err := abi.JSON(strings.NewReader(AsyncCallABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AsyncCallBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AsyncCall{AsyncCallCaller: AsyncCallCaller{contract: contract}, AsyncCallTransactor: AsyncCallTransactor{contract: contract}}, nil
}

// AsyncCall is an auto generated Go binding around an Ethereum contract.
type AsyncCall struct {
	AsyncCallCaller     // Read-only binding to the contract
	AsyncCallTransactor // Write-only binding to the contract
}

// AsyncCallCaller is an auto generated read-only Go binding around an Ethereum contract.
type AsyncCallCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AsyncCallTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AsyncCallTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AsyncCallSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AsyncCallSession struct {
	Contract     *AsyncCall        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AsyncCallCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AsyncCallCallerSession struct {
	Contract *AsyncCallCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// AsyncCallTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AsyncCallTransactorSession struct {
	Contract     *AsyncCallTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AsyncCallRaw is an auto generated low-level Go binding around an Ethereum contract.
type AsyncCallRaw struct {
	Contract *AsyncCall // Generic contract binding to access the raw methods on
}

// AsyncCallCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AsyncCallCallerRaw struct {
	Contract *AsyncCallCaller // Generic read-only contract binding to access the raw methods on
}

// AsyncCallTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AsyncCallTransactorRaw struct {
	Contract *AsyncCallTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAsyncCall creates a new instance of AsyncCall, bound to a specific deployed contract.
func NewAsyncCall(address common.Address, backend bind.ContractBackend) (*AsyncCall, error) {
	contract, err := bindAsyncCall(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AsyncCall{AsyncCallCaller: AsyncCallCaller{contract: contract}, AsyncCallTransactor: AsyncCallTransactor{contract: contract}}, nil
}

// NewAsyncCallCaller creates a new read-only instance of AsyncCall, bound to a specific deployed contract.
func NewAsyncCallCaller(address common.Address, caller bind.ContractCaller) (*AsyncCallCaller, error) {
	contract, err := bindAsyncCall(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &AsyncCallCaller{contract: contract}, nil
}

// NewAsyncCallTransactor creates a new write-only instance of AsyncCall, bound to a specific deployed contract.
func NewAsyncCallTransactor(address common.Address, transactor bind.ContractTransactor) (*AsyncCallTransactor, error) {
	contract, err := bindAsyncCall(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &AsyncCallTransactor{contract: contract}, nil
}

// bindAsyncCall binds a generic wrapper to an already deployed contract.
func bindAsyncCall(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AsyncCallABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AsyncCall *AsyncCallRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AsyncCall.Contract.AsyncCallCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AsyncCall *AsyncCallRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AsyncCall.Contract.AsyncCallTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AsyncCall *AsyncCallRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AsyncCall.Contract.AsyncCallTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AsyncCall *AsyncCallCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AsyncCall.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AsyncCall *AsyncCallTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AsyncCall.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AsyncCall *AsyncCallTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AsyncCall.Contract.contract.Transact(opts, method, params...)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() constant returns(bytes32)
func (_AsyncCall *AsyncCallCaller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _AsyncCall.contract.Call(opts, out, "contractType")
	return *ret0, err
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() constant returns(bytes32)
func (_AsyncCall *AsyncCallSession) ContractType() ([32]byte, error) {
	return _AsyncCall.Contract.ContractType(&_AsyncCall.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() constant returns(bytes32)
func (_AsyncCall *AsyncCallCallerSession) ContractType() ([32]byte, error) {
	return _AsyncCall.Contract.ContractType(&_AsyncCall.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AsyncCall *AsyncCallCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AsyncCall.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AsyncCall *AsyncCallSession) Owner() (common.Address, error) {
	return _AsyncCall.Contract.Owner(&_AsyncCall.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AsyncCall *AsyncCallCallerSession) Owner() (common.Address, error) {
	return _AsyncCall.Contract.Owner(&_AsyncCall.CallOpts)
}

// PendingCalls is a free data retrieval call binding the contract method 0x32fdd45c.
//
// Solidity: function pendingCalls( bytes32) constant returns(callData bytes, sender address, approved bool, success bool)
func (_AsyncCall *AsyncCallCaller) PendingCalls(opts *bind.CallOpts, arg0 [32]byte) (struct {
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
	err := _AsyncCall.contract.Call(opts, out, "pendingCalls", arg0)
	return *ret, err
}

// PendingCalls is a free data retrieval call binding the contract method 0x32fdd45c.
//
// Solidity: function pendingCalls( bytes32) constant returns(callData bytes, sender address, approved bool, success bool)
func (_AsyncCall *AsyncCallSession) PendingCalls(arg0 [32]byte) (struct {
	CallData []byte
	Sender   common.Address
	Approved bool
	Success  bool
}, error) {
	return _AsyncCall.Contract.PendingCalls(&_AsyncCall.CallOpts, arg0)
}

// PendingCalls is a free data retrieval call binding the contract method 0x32fdd45c.
//
// Solidity: function pendingCalls( bytes32) constant returns(callData bytes, sender address, approved bool, success bool)
func (_AsyncCall *AsyncCallCallerSession) PendingCalls(arg0 [32]byte) (struct {
	CallData []byte
	Sender   common.Address
	Approved bool
	Success  bool
}, error) {
	return _AsyncCall.Contract.PendingCalls(&_AsyncCall.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_AsyncCall *AsyncCallCaller) Version(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _AsyncCall.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_AsyncCall *AsyncCallSession) Version() ([32]byte, error) {
	return _AsyncCall.Contract.Version(&_AsyncCall.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_AsyncCall *AsyncCallCallerSession) Version() ([32]byte, error) {
	return _AsyncCall.Contract.Version(&_AsyncCall.CallOpts)
}

// WaitConfirmation is a free data retrieval call binding the contract method 0xac72bb98.
//
// Solidity: function waitConfirmation() constant returns(bool)
func (_AsyncCall *AsyncCallCaller) WaitConfirmation(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AsyncCall.contract.Call(opts, out, "waitConfirmation")
	return *ret0, err
}

// WaitConfirmation is a free data retrieval call binding the contract method 0xac72bb98.
//
// Solidity: function waitConfirmation() constant returns(bool)
func (_AsyncCall *AsyncCallSession) WaitConfirmation() (bool, error) {
	return _AsyncCall.Contract.WaitConfirmation(&_AsyncCall.CallOpts)
}

// WaitConfirmation is a free data retrieval call binding the contract method 0xac72bb98.
//
// Solidity: function waitConfirmation() constant returns(bool)
func (_AsyncCall *AsyncCallCallerSession) WaitConfirmation() (bool, error) {
	return _AsyncCall.Contract.WaitConfirmation(&_AsyncCall.CallOpts)
}

// BeginCall is a paid mutator transaction binding the contract method 0x203eaf27.
//
// Solidity: function beginCall(publicCallData bytes, privateData bytes) returns()
func (_AsyncCall *AsyncCallTransactor) BeginCall(opts *bind.TransactOpts, publicCallData []byte, privateData []byte) (*types.Transaction, error) {
	return _AsyncCall.contract.Transact(opts, "beginCall", publicCallData, privateData)
}

// BeginCall is a paid mutator transaction binding the contract method 0x203eaf27.
//
// Solidity: function beginCall(publicCallData bytes, privateData bytes) returns()
func (_AsyncCall *AsyncCallSession) BeginCall(publicCallData []byte, privateData []byte) (*types.Transaction, error) {
	return _AsyncCall.Contract.BeginCall(&_AsyncCall.TransactOpts, publicCallData, privateData)
}

// BeginCall is a paid mutator transaction binding the contract method 0x203eaf27.
//
// Solidity: function beginCall(publicCallData bytes, privateData bytes) returns()
func (_AsyncCall *AsyncCallTransactorSession) BeginCall(publicCallData []byte, privateData []byte) (*types.Transaction, error) {
	return _AsyncCall.Contract.BeginCall(&_AsyncCall.TransactOpts, publicCallData, privateData)
}

// ChangeConfirmation is a paid mutator transaction binding the contract method 0x44f5484a.
//
// Solidity: function changeConfirmation(_waitConfirmation bool) returns()
func (_AsyncCall *AsyncCallTransactor) ChangeConfirmation(opts *bind.TransactOpts, _waitConfirmation bool) (*types.Transaction, error) {
	return _AsyncCall.contract.Transact(opts, "changeConfirmation", _waitConfirmation)
}

// ChangeConfirmation is a paid mutator transaction binding the contract method 0x44f5484a.
//
// Solidity: function changeConfirmation(_waitConfirmation bool) returns()
func (_AsyncCall *AsyncCallSession) ChangeConfirmation(_waitConfirmation bool) (*types.Transaction, error) {
	return _AsyncCall.Contract.ChangeConfirmation(&_AsyncCall.TransactOpts, _waitConfirmation)
}

// ChangeConfirmation is a paid mutator transaction binding the contract method 0x44f5484a.
//
// Solidity: function changeConfirmation(_waitConfirmation bool) returns()
func (_AsyncCall *AsyncCallTransactorSession) ChangeConfirmation(_waitConfirmation bool) (*types.Transaction, error) {
	return _AsyncCall.Contract.ChangeConfirmation(&_AsyncCall.TransactOpts, _waitConfirmation)
}

// ContinueCall is a paid mutator transaction binding the contract method 0x411f2351.
//
// Solidity: function continueCall(msgDataHash bytes32) returns()
func (_AsyncCall *AsyncCallTransactor) ContinueCall(opts *bind.TransactOpts, msgDataHash [32]byte) (*types.Transaction, error) {
	return _AsyncCall.contract.Transact(opts, "continueCall", msgDataHash)
}

// ContinueCall is a paid mutator transaction binding the contract method 0x411f2351.
//
// Solidity: function continueCall(msgDataHash bytes32) returns()
func (_AsyncCall *AsyncCallSession) ContinueCall(msgDataHash [32]byte) (*types.Transaction, error) {
	return _AsyncCall.Contract.ContinueCall(&_AsyncCall.TransactOpts, msgDataHash)
}

// ContinueCall is a paid mutator transaction binding the contract method 0x411f2351.
//
// Solidity: function continueCall(msgDataHash bytes32) returns()
func (_AsyncCall *AsyncCallTransactorSession) ContinueCall(msgDataHash [32]byte) (*types.Transaction, error) {
	return _AsyncCall.Contract.ContinueCall(&_AsyncCall.TransactOpts, msgDataHash)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_AsyncCall *AsyncCallTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AsyncCall.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_AsyncCall *AsyncCallSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AsyncCall.Contract.TransferOwnership(&_AsyncCall.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_AsyncCall *AsyncCallTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AsyncCall.Contract.TransferOwnership(&_AsyncCall.TransactOpts, newOwner)
}
