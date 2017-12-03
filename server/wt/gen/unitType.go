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

// UnitTypeABI is the input ABI used to generate the binding from.
const UnitTypeABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"amenityId\",\"type\":\"uint256\"}],\"name\":\"removeAmenity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"url\",\"type\":\"string\"}],\"name\":\"addImage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"increaseUnits\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getImagesLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalUnits\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"decreaseUnits\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"images\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unitType\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeImage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_description\",\"type\":\"string\"},{\"name\":\"_minGuests\",\"type\":\"uint256\"},{\"name\":\"_maxGuests\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"string\"}],\"name\":\"edit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractType\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAmenities\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amenityId\",\"type\":\"uint256\"}],\"name\":\"addAmenity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_unitType\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// UnitTypeBin is the compiled bytecode used for deploying new contracts.
const UnitTypeBin = `0x60606040527f302e302e312d616c7068610000000000000000000000000000000000000000006001906000191690557f696d6167657300000000000000000000000000000000000000000000000000006002906000191690557f302e302e312d616c7068610000000000000000000000000000000000000000006004906000191690557f756e69747479706500000000000000000000000000000000000000000000000060059060001916905534156100b757600080fd5b6040516040806110d883398101604052808051906020019091908051906020019091905050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600681600019169055505050610f60806101786000396000f3006060604052600436106100f1576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630ad23923146100f6578063166f177914610119578063362db2911461017657806354fd4d501461018b5780635a9b0b89146101bc57806360839bd8146102c45780636d86acc4146102ed57806372730c0114610316578063848564821461032b57806385ad4f90146103c75780638b25261f146103f85780638da5cb5b1461041b57806391e3118314610470578063cb2ef6f714610522578063d0ceaae814610553578063f240180c146105bd578063f2fde38b146105e0575b600080fd5b341561010157600080fd5b6101176004808035906020019091905050610619565b005b341561012457600080fd5b610174600480803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919050506106c1565b005b341561018157600080fd5b61018961075b565b005b341561019657600080fd5b61019e6107ca565b60405180826000191660001916815260200191505060405180910390f35b34156101c757600080fd5b6101cf6107d0565b604051808060200185815260200184815260200180602001838103835287818151815260200191508051906020019080838360005b8381101561021f578082015181840152602081019050610204565b50505050905090810190601f16801561024c5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b8381101561028557808201518184015260208101905061026a565b50505050905090810190601f1680156102b25780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b34156102cf57600080fd5b6102d7610933565b6040518082815260200191505060405180910390f35b34156102f857600080fd5b610300610940565b6040518082815260200191505060405180910390f35b341561032157600080fd5b610329610946565b005b341561033657600080fd5b61034c60048080359060200190919050506109b6565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561038c578082015181840152602081019050610371565b50505050905090810190601f1680156103b95780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156103d257600080fd5b6103da610a72565b60405180826000191660001916815260200191505060405180910390f35b341561040357600080fd5b6104196004808035906020019091905050610a78565b005b341561042657600080fd5b61042e610afc565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561047b57600080fd5b610520600480803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803590602001909190803590602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610b21565b005b341561052d57600080fd5b610535610bbe565b60405180826000191660001916815260200191505060405180910390f35b341561055e57600080fd5b610566610bc4565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156105a957808201518184015260208101905061058e565b505050509050019250505060405180910390f35b34156105c857600080fd5b6105de6004808035906020019091905050610c22565b005b34156105eb57600080fd5b610617600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610cc6565b005b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561067457600080fd5b600c600d60008381526020019081526020016000205481548110151561069657fe5b9060005260206000209001600090556000600d60008381526020019081526020016000208190555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561071c57600080fd5b600380548060010182816107309190610d9b565b916000526020600020900160008390919091509080519060200190610756929190610dc7565b505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156107b657600080fd5b600760008154809291906001019190505550565b60045481565b6107d8610e47565b6000806107e3610e47565b6008600954600a54600b838054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108825780601f1061085757610100808354040283529160200191610882565b820191906000526020600020905b81548152906001019060200180831161086557829003601f168201915b50505050509350808054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561091e5780601f106108f35761010080835404028352916020019161091e565b820191906000526020600020905b81548152906001019060200180831161090157829003601f168201915b50505050509050935093509350935090919293565b6000600380549050905090565b60075481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156109a157600080fd5b60076000815480929190600190039190505550565b6003818154811015156109c557fe5b90600052602060002090016000915090508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a6a5780601f10610a3f57610100808354040283529160200191610a6a565b820191906000526020600020905b815481529060010190602001808311610a4d57829003601f168201915b505050505081565b60065481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610ad357600080fd5b600381815481101515610ae257fe5b90600052602060002090016000610af99190610e5b565b50565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610b7c57600080fd5b8360089080519060200190610b92929190610dc7565b508260098190555081600a8190555080600b9080519060200190610bb7929190610dc7565b5050505050565b60055481565b610bcc610ea3565b600c805480602002602001604051908101604052809291908181526020018280548015610c1857602002820191906000526020600020905b815481526020019060010190808311610c04575b5050505050905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610c7d57600080fd5b600c80549050600d600083815260200190815260200160002081905550600c8054806001018281610cae9190610eb7565b91600052602060002090016000839091909150555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610d2157600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515610d9857806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b815481835581811511610dc257818360005260206000209182019101610dc19190610ee3565b5b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610e0857805160ff1916838001178555610e36565b82800160010185558215610e36579182015b82811115610e35578251825591602001919060010190610e1a565b5b509050610e439190610f0f565b5090565b602060405190810160405280600081525090565b50805460018160011615610100020316600290046000825580601f10610e815750610ea0565b601f016020900490600052602060002090810190610e9f9190610f0f565b5b50565b602060405190810160405280600081525090565b815481835581811511610ede57818360005260206000209182019101610edd9190610f0f565b5b505050565b610f0c91905b80821115610f085760008181610eff9190610e5b565b50600101610ee9565b5090565b90565b610f3191905b80821115610f2d576000816000905550600101610f15565b5090565b905600a165627a7a72305820a292a5a24a79ad320e2b35f6a9649f34db0b049e055aac615dbd59d86b8b4c7b0029`

// DeployUnitType deploys a new Ethereum contract, binding an instance of UnitType to it.
func DeployUnitType(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _unitType [32]byte) (common.Address, *types.Transaction, *UnitType, error) {
	parsed, err := abi.JSON(strings.NewReader(UnitTypeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UnitTypeBin), backend, _owner, _unitType)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UnitType{UnitTypeCaller: UnitTypeCaller{contract: contract}, UnitTypeTransactor: UnitTypeTransactor{contract: contract}}, nil
}

// UnitType is an auto generated Go binding around an Ethereum contract.
type UnitType struct {
	UnitTypeCaller     // Read-only binding to the contract
	UnitTypeTransactor // Write-only binding to the contract
}

// UnitTypeCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnitTypeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnitTypeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnitTypeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnitTypeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnitTypeSession struct {
	Contract     *UnitType         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UnitTypeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnitTypeCallerSession struct {
	Contract *UnitTypeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// UnitTypeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnitTypeTransactorSession struct {
	Contract     *UnitTypeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// UnitTypeRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnitTypeRaw struct {
	Contract *UnitType // Generic contract binding to access the raw methods on
}

// UnitTypeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnitTypeCallerRaw struct {
	Contract *UnitTypeCaller // Generic read-only contract binding to access the raw methods on
}

// UnitTypeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnitTypeTransactorRaw struct {
	Contract *UnitTypeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnitType creates a new instance of UnitType, bound to a specific deployed contract.
func NewUnitType(address common.Address, backend bind.ContractBackend) (*UnitType, error) {
	contract, err := bindUnitType(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UnitType{UnitTypeCaller: UnitTypeCaller{contract: contract}, UnitTypeTransactor: UnitTypeTransactor{contract: contract}}, nil
}

// NewUnitTypeCaller creates a new read-only instance of UnitType, bound to a specific deployed contract.
func NewUnitTypeCaller(address common.Address, caller bind.ContractCaller) (*UnitTypeCaller, error) {
	contract, err := bindUnitType(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &UnitTypeCaller{contract: contract}, nil
}

// NewUnitTypeTransactor creates a new write-only instance of UnitType, bound to a specific deployed contract.
func NewUnitTypeTransactor(address common.Address, transactor bind.ContractTransactor) (*UnitTypeTransactor, error) {
	contract, err := bindUnitType(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &UnitTypeTransactor{contract: contract}, nil
}

// bindUnitType binds a generic wrapper to an already deployed contract.
func bindUnitType(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UnitTypeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnitType *UnitTypeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UnitType.Contract.UnitTypeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnitType *UnitTypeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnitType.Contract.UnitTypeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnitType *UnitTypeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnitType.Contract.UnitTypeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnitType *UnitTypeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UnitType.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnitType *UnitTypeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnitType.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnitType *UnitTypeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnitType.Contract.contract.Transact(opts, method, params...)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() constant returns(bytes32)
func (_UnitType *UnitTypeCaller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _UnitType.contract.Call(opts, out, "contractType")
	return *ret0, err
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() constant returns(bytes32)
func (_UnitType *UnitTypeSession) ContractType() ([32]byte, error) {
	return _UnitType.Contract.ContractType(&_UnitType.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() constant returns(bytes32)
func (_UnitType *UnitTypeCallerSession) ContractType() ([32]byte, error) {
	return _UnitType.Contract.ContractType(&_UnitType.CallOpts)
}

// GetAmenities is a free data retrieval call binding the contract method 0xd0ceaae8.
//
// Solidity: function getAmenities() constant returns(uint256[])
func (_UnitType *UnitTypeCaller) GetAmenities(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _UnitType.contract.Call(opts, out, "getAmenities")
	return *ret0, err
}

// GetAmenities is a free data retrieval call binding the contract method 0xd0ceaae8.
//
// Solidity: function getAmenities() constant returns(uint256[])
func (_UnitType *UnitTypeSession) GetAmenities() ([]*big.Int, error) {
	return _UnitType.Contract.GetAmenities(&_UnitType.CallOpts)
}

// GetAmenities is a free data retrieval call binding the contract method 0xd0ceaae8.
//
// Solidity: function getAmenities() constant returns(uint256[])
func (_UnitType *UnitTypeCallerSession) GetAmenities() ([]*big.Int, error) {
	return _UnitType.Contract.GetAmenities(&_UnitType.CallOpts)
}

// GetImagesLength is a free data retrieval call binding the contract method 0x60839bd8.
//
// Solidity: function getImagesLength() constant returns(uint256)
func (_UnitType *UnitTypeCaller) GetImagesLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UnitType.contract.Call(opts, out, "getImagesLength")
	return *ret0, err
}

// GetImagesLength is a free data retrieval call binding the contract method 0x60839bd8.
//
// Solidity: function getImagesLength() constant returns(uint256)
func (_UnitType *UnitTypeSession) GetImagesLength() (*big.Int, error) {
	return _UnitType.Contract.GetImagesLength(&_UnitType.CallOpts)
}

// GetImagesLength is a free data retrieval call binding the contract method 0x60839bd8.
//
// Solidity: function getImagesLength() constant returns(uint256)
func (_UnitType *UnitTypeCallerSession) GetImagesLength() (*big.Int, error) {
	return _UnitType.Contract.GetImagesLength(&_UnitType.CallOpts)
}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() constant returns(string, uint256, uint256, string)
func (_UnitType *UnitTypeCaller) GetInfo(opts *bind.CallOpts) (string, *big.Int, *big.Int, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _UnitType.contract.Call(opts, out, "getInfo")
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() constant returns(string, uint256, uint256, string)
func (_UnitType *UnitTypeSession) GetInfo() (string, *big.Int, *big.Int, string, error) {
	return _UnitType.Contract.GetInfo(&_UnitType.CallOpts)
}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() constant returns(string, uint256, uint256, string)
func (_UnitType *UnitTypeCallerSession) GetInfo() (string, *big.Int, *big.Int, string, error) {
	return _UnitType.Contract.GetInfo(&_UnitType.CallOpts)
}

// Images is a free data retrieval call binding the contract method 0x84856482.
//
// Solidity: function images( uint256) constant returns(string)
func (_UnitType *UnitTypeCaller) Images(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _UnitType.contract.Call(opts, out, "images", arg0)
	return *ret0, err
}

// Images is a free data retrieval call binding the contract method 0x84856482.
//
// Solidity: function images( uint256) constant returns(string)
func (_UnitType *UnitTypeSession) Images(arg0 *big.Int) (string, error) {
	return _UnitType.Contract.Images(&_UnitType.CallOpts, arg0)
}

// Images is a free data retrieval call binding the contract method 0x84856482.
//
// Solidity: function images( uint256) constant returns(string)
func (_UnitType *UnitTypeCallerSession) Images(arg0 *big.Int) (string, error) {
	return _UnitType.Contract.Images(&_UnitType.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UnitType *UnitTypeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UnitType.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UnitType *UnitTypeSession) Owner() (common.Address, error) {
	return _UnitType.Contract.Owner(&_UnitType.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UnitType *UnitTypeCallerSession) Owner() (common.Address, error) {
	return _UnitType.Contract.Owner(&_UnitType.CallOpts)
}

// TotalUnits is a free data retrieval call binding the contract method 0x6d86acc4.
//
// Solidity: function totalUnits() constant returns(uint256)
func (_UnitType *UnitTypeCaller) TotalUnits(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UnitType.contract.Call(opts, out, "totalUnits")
	return *ret0, err
}

// TotalUnits is a free data retrieval call binding the contract method 0x6d86acc4.
//
// Solidity: function totalUnits() constant returns(uint256)
func (_UnitType *UnitTypeSession) TotalUnits() (*big.Int, error) {
	return _UnitType.Contract.TotalUnits(&_UnitType.CallOpts)
}

// TotalUnits is a free data retrieval call binding the contract method 0x6d86acc4.
//
// Solidity: function totalUnits() constant returns(uint256)
func (_UnitType *UnitTypeCallerSession) TotalUnits() (*big.Int, error) {
	return _UnitType.Contract.TotalUnits(&_UnitType.CallOpts)
}

// UnitType is a free data retrieval call binding the contract method 0x85ad4f90.
//
// Solidity: function unitType() constant returns(bytes32)
func (_UnitType *UnitTypeCaller) UnitType(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _UnitType.contract.Call(opts, out, "unitType")
	return *ret0, err
}

// UnitType is a free data retrieval call binding the contract method 0x85ad4f90.
//
// Solidity: function unitType() constant returns(bytes32)
func (_UnitType *UnitTypeSession) UnitType() ([32]byte, error) {
	return _UnitType.Contract.UnitType(&_UnitType.CallOpts)
}

// UnitType is a free data retrieval call binding the contract method 0x85ad4f90.
//
// Solidity: function unitType() constant returns(bytes32)
func (_UnitType *UnitTypeCallerSession) UnitType() ([32]byte, error) {
	return _UnitType.Contract.UnitType(&_UnitType.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_UnitType *UnitTypeCaller) Version(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _UnitType.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_UnitType *UnitTypeSession) Version() ([32]byte, error) {
	return _UnitType.Contract.Version(&_UnitType.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_UnitType *UnitTypeCallerSession) Version() ([32]byte, error) {
	return _UnitType.Contract.Version(&_UnitType.CallOpts)
}

// AddAmenity is a paid mutator transaction binding the contract method 0xf240180c.
//
// Solidity: function addAmenity(amenityId uint256) returns()
func (_UnitType *UnitTypeTransactor) AddAmenity(opts *bind.TransactOpts, amenityId *big.Int) (*types.Transaction, error) {
	return _UnitType.contract.Transact(opts, "addAmenity", amenityId)
}

// AddAmenity is a paid mutator transaction binding the contract method 0xf240180c.
//
// Solidity: function addAmenity(amenityId uint256) returns()
func (_UnitType *UnitTypeSession) AddAmenity(amenityId *big.Int) (*types.Transaction, error) {
	return _UnitType.Contract.AddAmenity(&_UnitType.TransactOpts, amenityId)
}

// AddAmenity is a paid mutator transaction binding the contract method 0xf240180c.
//
// Solidity: function addAmenity(amenityId uint256) returns()
func (_UnitType *UnitTypeTransactorSession) AddAmenity(amenityId *big.Int) (*types.Transaction, error) {
	return _UnitType.Contract.AddAmenity(&_UnitType.TransactOpts, amenityId)
}

// AddImage is a paid mutator transaction binding the contract method 0x166f1779.
//
// Solidity: function addImage(url string) returns()
func (_UnitType *UnitTypeTransactor) AddImage(opts *bind.TransactOpts, url string) (*types.Transaction, error) {
	return _UnitType.contract.Transact(opts, "addImage", url)
}

// AddImage is a paid mutator transaction binding the contract method 0x166f1779.
//
// Solidity: function addImage(url string) returns()
func (_UnitType *UnitTypeSession) AddImage(url string) (*types.Transaction, error) {
	return _UnitType.Contract.AddImage(&_UnitType.TransactOpts, url)
}

// AddImage is a paid mutator transaction binding the contract method 0x166f1779.
//
// Solidity: function addImage(url string) returns()
func (_UnitType *UnitTypeTransactorSession) AddImage(url string) (*types.Transaction, error) {
	return _UnitType.Contract.AddImage(&_UnitType.TransactOpts, url)
}

// DecreaseUnits is a paid mutator transaction binding the contract method 0x72730c01.
//
// Solidity: function decreaseUnits() returns()
func (_UnitType *UnitTypeTransactor) DecreaseUnits(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnitType.contract.Transact(opts, "decreaseUnits")
}

// DecreaseUnits is a paid mutator transaction binding the contract method 0x72730c01.
//
// Solidity: function decreaseUnits() returns()
func (_UnitType *UnitTypeSession) DecreaseUnits() (*types.Transaction, error) {
	return _UnitType.Contract.DecreaseUnits(&_UnitType.TransactOpts)
}

// DecreaseUnits is a paid mutator transaction binding the contract method 0x72730c01.
//
// Solidity: function decreaseUnits() returns()
func (_UnitType *UnitTypeTransactorSession) DecreaseUnits() (*types.Transaction, error) {
	return _UnitType.Contract.DecreaseUnits(&_UnitType.TransactOpts)
}

// Edit is a paid mutator transaction binding the contract method 0x91e31183.
//
// Solidity: function edit(_description string, _minGuests uint256, _maxGuests uint256, _price string) returns()
func (_UnitType *UnitTypeTransactor) Edit(opts *bind.TransactOpts, _description string, _minGuests *big.Int, _maxGuests *big.Int, _price string) (*types.Transaction, error) {
	return _UnitType.contract.Transact(opts, "edit", _description, _minGuests, _maxGuests, _price)
}

// Edit is a paid mutator transaction binding the contract method 0x91e31183.
//
// Solidity: function edit(_description string, _minGuests uint256, _maxGuests uint256, _price string) returns()
func (_UnitType *UnitTypeSession) Edit(_description string, _minGuests *big.Int, _maxGuests *big.Int, _price string) (*types.Transaction, error) {
	return _UnitType.Contract.Edit(&_UnitType.TransactOpts, _description, _minGuests, _maxGuests, _price)
}

// Edit is a paid mutator transaction binding the contract method 0x91e31183.
//
// Solidity: function edit(_description string, _minGuests uint256, _maxGuests uint256, _price string) returns()
func (_UnitType *UnitTypeTransactorSession) Edit(_description string, _minGuests *big.Int, _maxGuests *big.Int, _price string) (*types.Transaction, error) {
	return _UnitType.Contract.Edit(&_UnitType.TransactOpts, _description, _minGuests, _maxGuests, _price)
}

// IncreaseUnits is a paid mutator transaction binding the contract method 0x362db291.
//
// Solidity: function increaseUnits() returns()
func (_UnitType *UnitTypeTransactor) IncreaseUnits(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnitType.contract.Transact(opts, "increaseUnits")
}

// IncreaseUnits is a paid mutator transaction binding the contract method 0x362db291.
//
// Solidity: function increaseUnits() returns()
func (_UnitType *UnitTypeSession) IncreaseUnits() (*types.Transaction, error) {
	return _UnitType.Contract.IncreaseUnits(&_UnitType.TransactOpts)
}

// IncreaseUnits is a paid mutator transaction binding the contract method 0x362db291.
//
// Solidity: function increaseUnits() returns()
func (_UnitType *UnitTypeTransactorSession) IncreaseUnits() (*types.Transaction, error) {
	return _UnitType.Contract.IncreaseUnits(&_UnitType.TransactOpts)
}

// RemoveAmenity is a paid mutator transaction binding the contract method 0x0ad23923.
//
// Solidity: function removeAmenity(amenityId uint256) returns()
func (_UnitType *UnitTypeTransactor) RemoveAmenity(opts *bind.TransactOpts, amenityId *big.Int) (*types.Transaction, error) {
	return _UnitType.contract.Transact(opts, "removeAmenity", amenityId)
}

// RemoveAmenity is a paid mutator transaction binding the contract method 0x0ad23923.
//
// Solidity: function removeAmenity(amenityId uint256) returns()
func (_UnitType *UnitTypeSession) RemoveAmenity(amenityId *big.Int) (*types.Transaction, error) {
	return _UnitType.Contract.RemoveAmenity(&_UnitType.TransactOpts, amenityId)
}

// RemoveAmenity is a paid mutator transaction binding the contract method 0x0ad23923.
//
// Solidity: function removeAmenity(amenityId uint256) returns()
func (_UnitType *UnitTypeTransactorSession) RemoveAmenity(amenityId *big.Int) (*types.Transaction, error) {
	return _UnitType.Contract.RemoveAmenity(&_UnitType.TransactOpts, amenityId)
}

// RemoveImage is a paid mutator transaction binding the contract method 0x8b25261f.
//
// Solidity: function removeImage(index uint256) returns()
func (_UnitType *UnitTypeTransactor) RemoveImage(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _UnitType.contract.Transact(opts, "removeImage", index)
}

// RemoveImage is a paid mutator transaction binding the contract method 0x8b25261f.
//
// Solidity: function removeImage(index uint256) returns()
func (_UnitType *UnitTypeSession) RemoveImage(index *big.Int) (*types.Transaction, error) {
	return _UnitType.Contract.RemoveImage(&_UnitType.TransactOpts, index)
}

// RemoveImage is a paid mutator transaction binding the contract method 0x8b25261f.
//
// Solidity: function removeImage(index uint256) returns()
func (_UnitType *UnitTypeTransactorSession) RemoveImage(index *big.Int) (*types.Transaction, error) {
	return _UnitType.Contract.RemoveImage(&_UnitType.TransactOpts, index)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_UnitType *UnitTypeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UnitType.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_UnitType *UnitTypeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UnitType.Contract.TransferOwnership(&_UnitType.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_UnitType *UnitTypeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UnitType.Contract.TransferOwnership(&_UnitType.TransactOpts, newOwner)
}
