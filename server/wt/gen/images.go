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

// ImagesABI is the input ABI used to generate the binding from.
const ImagesABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"url\",\"type\":\"string\"}],\"name\":\"addImage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getImagesLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"images\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeImage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractType\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ImagesBin is the compiled bytecode used for deploying new contracts.
const ImagesBin = `0x60606040527f302e302e312d616c7068610000000000000000000000000000000000000000006001906000191690557f696d616765730000000000000000000000000000000000000000000000000000600290600019169055336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610726806100a76000396000f30060606040526004361061008e576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063166f17791461009357806354fd4d50146100f057806360839bd814610121578063848564821461014a5780638b25261f146101e65780638da5cb5b14610209578063cb2ef6f71461025e578063f2fde38b1461028f575b600080fd5b341561009e57600080fd5b6100ee600480803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919050506102c8565b005b34156100fb57600080fd5b610103610362565b60405180826000191660001916815260200191505060405180910390f35b341561012c57600080fd5b610134610368565b6040518082815260200191505060405180910390f35b341561015557600080fd5b61016b6004808035906020019091905050610375565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101ab578082015181840152602081019050610190565b50505050905090810190601f1680156101d85780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101f157600080fd5b6102076004808035906020019091905050610431565b005b341561021457600080fd5b61021c6104b5565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561026957600080fd5b6102716104da565b60405180826000191660001916815260200191505060405180910390f35b341561029a57600080fd5b6102c6600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506104e0565b005b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561032357600080fd5b6003805480600101828161033791906105b5565b91600052602060002090016000839091909150908051906020019061035d9291906105e1565b505050565b60015481565b6000600380549050905090565b60038181548110151561038457fe5b90600052602060002090016000915090508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104295780601f106103fe57610100808354040283529160200191610429565b820191906000526020600020905b81548152906001019060200180831161040c57829003601f168201915b505050505081565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561048c57600080fd5b60038181548110151561049b57fe5b906000526020600020900160006104b29190610661565b50565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60025481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561053b57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415156105b257806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b8154818355818115116105dc578183600052602060002091820191016105db91906106a9565b5b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061062257805160ff1916838001178555610650565b82800160010185558215610650579182015b8281111561064f578251825591602001919060010190610634565b5b50905061065d91906106d5565b5090565b50805460018160011615610100020316600290046000825580601f1061068757506106a6565b601f0160209004906000526020600020908101906106a591906106d5565b5b50565b6106d291905b808211156106ce57600081816106c59190610661565b506001016106af565b5090565b90565b6106f791905b808211156106f35760008160009055506001016106db565b5090565b905600a165627a7a72305820b144bdd89a2da06085bddd9875aa2023b89de9454e697bf5b3e11c0b5c43c7530029`

// DeployImages deploys a new Ethereum contract, binding an instance of Images to it.
func DeployImages(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Images, error) {
	parsed, err := abi.JSON(strings.NewReader(ImagesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ImagesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Images{ImagesCaller: ImagesCaller{contract: contract}, ImagesTransactor: ImagesTransactor{contract: contract}}, nil
}

// Images is an auto generated Go binding around an Ethereum contract.
type Images struct {
	ImagesCaller     // Read-only binding to the contract
	ImagesTransactor // Write-only binding to the contract
}

// ImagesCaller is an auto generated read-only Go binding around an Ethereum contract.
type ImagesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImagesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ImagesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImagesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ImagesSession struct {
	Contract     *Images           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ImagesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ImagesCallerSession struct {
	Contract *ImagesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ImagesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ImagesTransactorSession struct {
	Contract     *ImagesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ImagesRaw is an auto generated low-level Go binding around an Ethereum contract.
type ImagesRaw struct {
	Contract *Images // Generic contract binding to access the raw methods on
}

// ImagesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ImagesCallerRaw struct {
	Contract *ImagesCaller // Generic read-only contract binding to access the raw methods on
}

// ImagesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ImagesTransactorRaw struct {
	Contract *ImagesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewImages creates a new instance of Images, bound to a specific deployed contract.
func NewImages(address common.Address, backend bind.ContractBackend) (*Images, error) {
	contract, err := bindImages(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Images{ImagesCaller: ImagesCaller{contract: contract}, ImagesTransactor: ImagesTransactor{contract: contract}}, nil
}

// NewImagesCaller creates a new read-only instance of Images, bound to a specific deployed contract.
func NewImagesCaller(address common.Address, caller bind.ContractCaller) (*ImagesCaller, error) {
	contract, err := bindImages(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ImagesCaller{contract: contract}, nil
}

// NewImagesTransactor creates a new write-only instance of Images, bound to a specific deployed contract.
func NewImagesTransactor(address common.Address, transactor bind.ContractTransactor) (*ImagesTransactor, error) {
	contract, err := bindImages(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ImagesTransactor{contract: contract}, nil
}

// bindImages binds a generic wrapper to an already deployed contract.
func bindImages(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ImagesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Images *ImagesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Images.Contract.ImagesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Images *ImagesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Images.Contract.ImagesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Images *ImagesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Images.Contract.ImagesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Images *ImagesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Images.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Images *ImagesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Images.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Images *ImagesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Images.Contract.contract.Transact(opts, method, params...)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() constant returns(bytes32)
func (_Images *ImagesCaller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Images.contract.Call(opts, out, "contractType")
	return *ret0, err
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() constant returns(bytes32)
func (_Images *ImagesSession) ContractType() ([32]byte, error) {
	return _Images.Contract.ContractType(&_Images.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() constant returns(bytes32)
func (_Images *ImagesCallerSession) ContractType() ([32]byte, error) {
	return _Images.Contract.ContractType(&_Images.CallOpts)
}

// GetImagesLength is a free data retrieval call binding the contract method 0x60839bd8.
//
// Solidity: function getImagesLength() constant returns(uint256)
func (_Images *ImagesCaller) GetImagesLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Images.contract.Call(opts, out, "getImagesLength")
	return *ret0, err
}

// GetImagesLength is a free data retrieval call binding the contract method 0x60839bd8.
//
// Solidity: function getImagesLength() constant returns(uint256)
func (_Images *ImagesSession) GetImagesLength() (*big.Int, error) {
	return _Images.Contract.GetImagesLength(&_Images.CallOpts)
}

// GetImagesLength is a free data retrieval call binding the contract method 0x60839bd8.
//
// Solidity: function getImagesLength() constant returns(uint256)
func (_Images *ImagesCallerSession) GetImagesLength() (*big.Int, error) {
	return _Images.Contract.GetImagesLength(&_Images.CallOpts)
}

// Images is a free data retrieval call binding the contract method 0x84856482.
//
// Solidity: function images( uint256) constant returns(string)
func (_Images *ImagesCaller) Images(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Images.contract.Call(opts, out, "images", arg0)
	return *ret0, err
}

// Images is a free data retrieval call binding the contract method 0x84856482.
//
// Solidity: function images( uint256) constant returns(string)
func (_Images *ImagesSession) Images(arg0 *big.Int) (string, error) {
	return _Images.Contract.Images(&_Images.CallOpts, arg0)
}

// Images is a free data retrieval call binding the contract method 0x84856482.
//
// Solidity: function images( uint256) constant returns(string)
func (_Images *ImagesCallerSession) Images(arg0 *big.Int) (string, error) {
	return _Images.Contract.Images(&_Images.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Images *ImagesCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Images.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Images *ImagesSession) Owner() (common.Address, error) {
	return _Images.Contract.Owner(&_Images.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Images *ImagesCallerSession) Owner() (common.Address, error) {
	return _Images.Contract.Owner(&_Images.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Images *ImagesCaller) Version(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Images.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Images *ImagesSession) Version() ([32]byte, error) {
	return _Images.Contract.Version(&_Images.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Images *ImagesCallerSession) Version() ([32]byte, error) {
	return _Images.Contract.Version(&_Images.CallOpts)
}

// AddImage is a paid mutator transaction binding the contract method 0x166f1779.
//
// Solidity: function addImage(url string) returns()
func (_Images *ImagesTransactor) AddImage(opts *bind.TransactOpts, url string) (*types.Transaction, error) {
	return _Images.contract.Transact(opts, "addImage", url)
}

// AddImage is a paid mutator transaction binding the contract method 0x166f1779.
//
// Solidity: function addImage(url string) returns()
func (_Images *ImagesSession) AddImage(url string) (*types.Transaction, error) {
	return _Images.Contract.AddImage(&_Images.TransactOpts, url)
}

// AddImage is a paid mutator transaction binding the contract method 0x166f1779.
//
// Solidity: function addImage(url string) returns()
func (_Images *ImagesTransactorSession) AddImage(url string) (*types.Transaction, error) {
	return _Images.Contract.AddImage(&_Images.TransactOpts, url)
}

// RemoveImage is a paid mutator transaction binding the contract method 0x8b25261f.
//
// Solidity: function removeImage(index uint256) returns()
func (_Images *ImagesTransactor) RemoveImage(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _Images.contract.Transact(opts, "removeImage", index)
}

// RemoveImage is a paid mutator transaction binding the contract method 0x8b25261f.
//
// Solidity: function removeImage(index uint256) returns()
func (_Images *ImagesSession) RemoveImage(index *big.Int) (*types.Transaction, error) {
	return _Images.Contract.RemoveImage(&_Images.TransactOpts, index)
}

// RemoveImage is a paid mutator transaction binding the contract method 0x8b25261f.
//
// Solidity: function removeImage(index uint256) returns()
func (_Images *ImagesTransactorSession) RemoveImage(index *big.Int) (*types.Transaction, error) {
	return _Images.Contract.RemoveImage(&_Images.TransactOpts, index)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Images *ImagesTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Images.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Images *ImagesSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Images.Contract.TransferOwnership(&_Images.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Images *ImagesTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Images.Contract.TransferOwnership(&_Images.TransactOpts, newOwner)
}
