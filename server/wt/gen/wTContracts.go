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

// WTContractsABI is the input ABI used to generate the binding from.
const WTContractsABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_version\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getByAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"total\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pos\",\"type\":\"uint256\"}],\"name\":\"getContract\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getByName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_version\",\"type\":\"string\"}],\"name\":\"edit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractType\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// WTContractsBin is the compiled bytecode used for deploying new contracts.
const WTContractsBin = `0x60606040527f302e302e312d616c7068610000000000000000000000000000000000000000006001906000191690557f7774636f6e747261637473000000000000000000000000000000000000000000600290600019169055336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611561806100a76000396000f3006060604052600436106100a4576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630578dd1c146100a9578063259a1a341461010a5780632ddbd13a1461025b57806354fd4d50146102845780636ebc8c86146102b55780638da5cb5b146103f0578063b336ad8314610445578063c362e82d146105ba578063cb2ef6f71461061b578063f2fde38b1461064c575b600080fd5b34156100b457600080fd5b6101086004808035906020019082018035906020019190919290803573ffffffffffffffffffffffffffffffffffffffff169060200190919080359060200190820180359060200191909192905050610685565b005b341561011557600080fd5b610141600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610914565b60405180806020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001838103835286818151815260200191508051906020019080838360005b838110156101b757808201518184015260208101905061019c565b50505050905090810190601f1680156101e45780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b8381101561021d578082015181840152602081019050610202565b50505050905090810190601f16801561024a5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b341561026657600080fd5b61026e610c04565b6040518082815260200191505060405180910390f35b341561028f57600080fd5b610297610c0a565b60405180826000191660001916815260200191505060405180910390f35b34156102c057600080fd5b6102d66004808035906020019091905050610c10565b60405180806020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001838103835286818151815260200191508051906020019080838360005b8381101561034c578082015181840152602081019050610331565b50505050905090810190601f1680156103795780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156103b2578082015181840152602081019050610397565b50505050905090810190601f1680156103df5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b34156103fb57600080fd5b610403610e06565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561045057600080fd5b6104a0600480803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610e2b565b60405180806020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001838103835286818151815260200191508051906020019080838360005b838110156105165780820151818401526020810190506104fb565b50505050905090810190601f1680156105435780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b8381101561057c578082015181840152602081019050610561565b50505050905090810190601f1680156105a95780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b34156105c557600080fd5b6106196004808035906020019082018035906020019190919290803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001908201803590602001919091929050506111cb565b005b341561062657600080fd5b61062e611321565b60405180826000191660001916815260200191505060405180910390f35b341561065757600080fd5b610683600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611327565b005b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156106e057600080fd5b6000600586866040518083838082843782019150509250505090815260200160405180910390205414801561075457506000600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054145b1561090d5760036000815480929190600101919050555060606040519081016040528086868080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505081526020018473ffffffffffffffffffffffffffffffffffffffff16815260200183838080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505081525060046000600354815260200190815260200160002060008201518160000190805190602001906108329291906113fc565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020190805190602001906108969291906113fc565b509050506003546005868660405180838380828437820191505092505050908152602001604051809103902081905550600354600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b5050505050565b61091c61147c565b600061092661147c565b6000600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541115610bd15760046000600660008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054815260200190815260200160002060000160046000600660008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660046000600660008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548152602001908152602001600020600201828054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b235780601f10610af857610100808354040283529160200191610b23565b820191906000526020600020905b815481529060010190602001808311610b0657829003601f168201915b50505050509250808054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610bbf5780601f10610b9457610100808354040283529160200191610bbf565b820191906000526020600020905b815481529060010190602001808311610ba257829003601f168201915b50505050509050925092509250610bfd565b600060206040519081016040528060008152509060206040519081016040528060008152509250925092505b9193909250565b60035481565b60015481565b610c1861147c565b6000610c2261147c565b60035484111515610dd357600460008581526020019081526020016000206000016004600086815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660046000878152602001908152602001600020600201828054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d255780601f10610cfa57610100808354040283529160200191610d25565b820191906000526020600020905b815481529060010190602001808311610d0857829003601f168201915b50505050509250808054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610dc15780601f10610d9657610100808354040283529160200191610dc1565b820191906000526020600020905b815481529060010190602001808311610da457829003601f168201915b50505050509050925092509250610dff565b600060206040519081016040528060008152509060206040519081016040528060008152509250925092505b9193909250565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610e3361147c565b6000610e3d61147c565b60006005856040518082805190602001908083835b602083101515610e775780518252602082019150602081019050602083039250610e52565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902054111561119857600460006005866040518082805190602001908083835b602083101515610eed5780518252602082019150602081019050602083039250610ec8565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020548152602001908152602001600020600001600460006005876040518082805190602001908083835b602083101515610f6e5780518252602082019150602081019050602083039250610f49565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902054815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600460006005886040518082805190602001908083835b6020831015156110105780518252602082019150602081019050602083039250610feb565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020548152602001908152602001600020600201828054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110ea5780601f106110bf576101008083540402835291602001916110ea565b820191906000526020600020905b8154815290600101906020018083116110cd57829003601f168201915b50505050509250808054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156111865780601f1061115b57610100808354040283529160200191611186565b820191906000526020600020905b81548152906001019060200180831161116957829003601f168201915b505050505090509250925092506111c4565b600060206040519081016040528060008152509060206040519081016040528060008152509250925092505b9193909250565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561122657600080fd5b60006005868660405180838380828437820191505092505050908152602001604051809103902054111561131a5782600460006005888860405180838380828437820191505092505050908152602001604051809103902054815260200190815260200160002060010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550818160046000600589896040518083838082843782019150509250505090815260200160405180910390205481526020019081526020016000206002019190611318929190611490565b505b5050505050565b60025481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561138257600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415156113f957806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061143d57805160ff191683800117855561146b565b8280016001018555821561146b579182015b8281111561146a57825182559160200191906001019061144f565b5b5090506114789190611510565b5090565b602060405190810160405280600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106114d157803560ff19168380011785556114ff565b828001600101855582156114ff579182015b828111156114fe5782358255916020019190600101906114e3565b5b50905061150c9190611510565b5090565b61153291905b8082111561152e576000816000905550600101611516565b5090565b905600a165627a7a7230582061f3d2ab0ebf23e9f5cd46e2834ef33c5f1d164601d7419024610e0226aadaaf0029`

// DeployWTContracts deploys a new Ethereum contract, binding an instance of WTContracts to it.
func DeployWTContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WTContracts, error) {
	parsed, err := abi.JSON(strings.NewReader(WTContractsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WTContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WTContracts{WTContractsCaller: WTContractsCaller{contract: contract}, WTContractsTransactor: WTContractsTransactor{contract: contract}}, nil
}

// WTContracts is an auto generated Go binding around an Ethereum contract.
type WTContracts struct {
	WTContractsCaller     // Read-only binding to the contract
	WTContractsTransactor // Write-only binding to the contract
}

// WTContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type WTContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WTContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WTContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WTContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WTContractsSession struct {
	Contract     *WTContracts      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WTContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WTContractsCallerSession struct {
	Contract *WTContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// WTContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WTContractsTransactorSession struct {
	Contract     *WTContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// WTContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type WTContractsRaw struct {
	Contract *WTContracts // Generic contract binding to access the raw methods on
}

// WTContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WTContractsCallerRaw struct {
	Contract *WTContractsCaller // Generic read-only contract binding to access the raw methods on
}

// WTContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WTContractsTransactorRaw struct {
	Contract *WTContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWTContracts creates a new instance of WTContracts, bound to a specific deployed contract.
func NewWTContracts(address common.Address, backend bind.ContractBackend) (*WTContracts, error) {
	contract, err := bindWTContracts(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WTContracts{WTContractsCaller: WTContractsCaller{contract: contract}, WTContractsTransactor: WTContractsTransactor{contract: contract}}, nil
}

// NewWTContractsCaller creates a new read-only instance of WTContracts, bound to a specific deployed contract.
func NewWTContractsCaller(address common.Address, caller bind.ContractCaller) (*WTContractsCaller, error) {
	contract, err := bindWTContracts(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &WTContractsCaller{contract: contract}, nil
}

// NewWTContractsTransactor creates a new write-only instance of WTContracts, bound to a specific deployed contract.
func NewWTContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*WTContractsTransactor, error) {
	contract, err := bindWTContracts(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &WTContractsTransactor{contract: contract}, nil
}

// bindWTContracts binds a generic wrapper to an already deployed contract.
func bindWTContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WTContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WTContracts *WTContractsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WTContracts.Contract.WTContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WTContracts *WTContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WTContracts.Contract.WTContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WTContracts *WTContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WTContracts.Contract.WTContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WTContracts *WTContractsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WTContracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WTContracts *WTContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WTContracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WTContracts *WTContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WTContracts.Contract.contract.Transact(opts, method, params...)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() constant returns(bytes32)
func (_WTContracts *WTContractsCaller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _WTContracts.contract.Call(opts, out, "contractType")
	return *ret0, err
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() constant returns(bytes32)
func (_WTContracts *WTContractsSession) ContractType() ([32]byte, error) {
	return _WTContracts.Contract.ContractType(&_WTContracts.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() constant returns(bytes32)
func (_WTContracts *WTContractsCallerSession) ContractType() ([32]byte, error) {
	return _WTContracts.Contract.ContractType(&_WTContracts.CallOpts)
}

// GetByAddr is a free data retrieval call binding the contract method 0x259a1a34.
//
// Solidity: function getByAddr(_addr address) constant returns(string, address, string)
func (_WTContracts *WTContractsCaller) GetByAddr(opts *bind.CallOpts, _addr common.Address) (string, common.Address, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(common.Address)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _WTContracts.contract.Call(opts, out, "getByAddr", _addr)
	return *ret0, *ret1, *ret2, err
}

// GetByAddr is a free data retrieval call binding the contract method 0x259a1a34.
//
// Solidity: function getByAddr(_addr address) constant returns(string, address, string)
func (_WTContracts *WTContractsSession) GetByAddr(_addr common.Address) (string, common.Address, string, error) {
	return _WTContracts.Contract.GetByAddr(&_WTContracts.CallOpts, _addr)
}

// GetByAddr is a free data retrieval call binding the contract method 0x259a1a34.
//
// Solidity: function getByAddr(_addr address) constant returns(string, address, string)
func (_WTContracts *WTContractsCallerSession) GetByAddr(_addr common.Address) (string, common.Address, string, error) {
	return _WTContracts.Contract.GetByAddr(&_WTContracts.CallOpts, _addr)
}

// GetByName is a free data retrieval call binding the contract method 0xb336ad83.
//
// Solidity: function getByName(_name string) constant returns(string, address, string)
func (_WTContracts *WTContractsCaller) GetByName(opts *bind.CallOpts, _name string) (string, common.Address, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(common.Address)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _WTContracts.contract.Call(opts, out, "getByName", _name)
	return *ret0, *ret1, *ret2, err
}

// GetByName is a free data retrieval call binding the contract method 0xb336ad83.
//
// Solidity: function getByName(_name string) constant returns(string, address, string)
func (_WTContracts *WTContractsSession) GetByName(_name string) (string, common.Address, string, error) {
	return _WTContracts.Contract.GetByName(&_WTContracts.CallOpts, _name)
}

// GetByName is a free data retrieval call binding the contract method 0xb336ad83.
//
// Solidity: function getByName(_name string) constant returns(string, address, string)
func (_WTContracts *WTContractsCallerSession) GetByName(_name string) (string, common.Address, string, error) {
	return _WTContracts.Contract.GetByName(&_WTContracts.CallOpts, _name)
}

// GetContract is a free data retrieval call binding the contract method 0x6ebc8c86.
//
// Solidity: function getContract(_pos uint256) constant returns(string, address, string)
func (_WTContracts *WTContractsCaller) GetContract(opts *bind.CallOpts, _pos *big.Int) (string, common.Address, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(common.Address)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _WTContracts.contract.Call(opts, out, "getContract", _pos)
	return *ret0, *ret1, *ret2, err
}

// GetContract is a free data retrieval call binding the contract method 0x6ebc8c86.
//
// Solidity: function getContract(_pos uint256) constant returns(string, address, string)
func (_WTContracts *WTContractsSession) GetContract(_pos *big.Int) (string, common.Address, string, error) {
	return _WTContracts.Contract.GetContract(&_WTContracts.CallOpts, _pos)
}

// GetContract is a free data retrieval call binding the contract method 0x6ebc8c86.
//
// Solidity: function getContract(_pos uint256) constant returns(string, address, string)
func (_WTContracts *WTContractsCallerSession) GetContract(_pos *big.Int) (string, common.Address, string, error) {
	return _WTContracts.Contract.GetContract(&_WTContracts.CallOpts, _pos)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_WTContracts *WTContractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WTContracts.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_WTContracts *WTContractsSession) Owner() (common.Address, error) {
	return _WTContracts.Contract.Owner(&_WTContracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_WTContracts *WTContractsCallerSession) Owner() (common.Address, error) {
	return _WTContracts.Contract.Owner(&_WTContracts.CallOpts)
}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() constant returns(uint256)
func (_WTContracts *WTContractsCaller) Total(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WTContracts.contract.Call(opts, out, "total")
	return *ret0, err
}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() constant returns(uint256)
func (_WTContracts *WTContractsSession) Total() (*big.Int, error) {
	return _WTContracts.Contract.Total(&_WTContracts.CallOpts)
}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() constant returns(uint256)
func (_WTContracts *WTContractsCallerSession) Total() (*big.Int, error) {
	return _WTContracts.Contract.Total(&_WTContracts.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_WTContracts *WTContractsCaller) Version(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _WTContracts.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_WTContracts *WTContractsSession) Version() ([32]byte, error) {
	return _WTContracts.Contract.Version(&_WTContracts.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_WTContracts *WTContractsCallerSession) Version() ([32]byte, error) {
	return _WTContracts.Contract.Version(&_WTContracts.CallOpts)
}

// Edit is a paid mutator transaction binding the contract method 0xc362e82d.
//
// Solidity: function edit(_name string, _addr address, _version string) returns()
func (_WTContracts *WTContractsTransactor) Edit(opts *bind.TransactOpts, _name string, _addr common.Address, _version string) (*types.Transaction, error) {
	return _WTContracts.contract.Transact(opts, "edit", _name, _addr, _version)
}

// Edit is a paid mutator transaction binding the contract method 0xc362e82d.
//
// Solidity: function edit(_name string, _addr address, _version string) returns()
func (_WTContracts *WTContractsSession) Edit(_name string, _addr common.Address, _version string) (*types.Transaction, error) {
	return _WTContracts.Contract.Edit(&_WTContracts.TransactOpts, _name, _addr, _version)
}

// Edit is a paid mutator transaction binding the contract method 0xc362e82d.
//
// Solidity: function edit(_name string, _addr address, _version string) returns()
func (_WTContracts *WTContractsTransactorSession) Edit(_name string, _addr common.Address, _version string) (*types.Transaction, error) {
	return _WTContracts.Contract.Edit(&_WTContracts.TransactOpts, _name, _addr, _version)
}

// Register is a paid mutator transaction binding the contract method 0x0578dd1c.
//
// Solidity: function register(_name string, _addr address, _version string) returns()
func (_WTContracts *WTContractsTransactor) Register(opts *bind.TransactOpts, _name string, _addr common.Address, _version string) (*types.Transaction, error) {
	return _WTContracts.contract.Transact(opts, "register", _name, _addr, _version)
}

// Register is a paid mutator transaction binding the contract method 0x0578dd1c.
//
// Solidity: function register(_name string, _addr address, _version string) returns()
func (_WTContracts *WTContractsSession) Register(_name string, _addr common.Address, _version string) (*types.Transaction, error) {
	return _WTContracts.Contract.Register(&_WTContracts.TransactOpts, _name, _addr, _version)
}

// Register is a paid mutator transaction binding the contract method 0x0578dd1c.
//
// Solidity: function register(_name string, _addr address, _version string) returns()
func (_WTContracts *WTContractsTransactorSession) Register(_name string, _addr common.Address, _version string) (*types.Transaction, error) {
	return _WTContracts.Contract.Register(&_WTContracts.TransactOpts, _name, _addr, _version)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_WTContracts *WTContractsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WTContracts.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_WTContracts *WTContractsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WTContracts.Contract.TransferOwnership(&_WTContracts.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_WTContracts *WTContractsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WTContracts.Contract.TransferOwnership(&_WTContracts.TransactOpts, newOwner)
}
