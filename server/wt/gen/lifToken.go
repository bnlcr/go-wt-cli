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

// LifTokenABI is the input ABI used to generate the binding from.
const LifTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DECIMALS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"approveData\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMinting\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transferData\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transferDataFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SYMBOL\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// LifTokenBin is the compiled bytecode used for deploying new contracts.
const LifTokenBin = `0x60606040526000600360146101000a81548160ff0219169083151502179055506000600360156101000a81548160ff02191690831515021790555033600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611a398061008a6000396000f30060606040526004361061011d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806305d2035b14610122578063095ea7b31461014f57806318160ddd146101a957806323b872dd146101d25780632e0f26251461024b5780633f4ba83a1461027457806340c10f19146102a157806342966c68146102fb5780635c975abb1461031e5780636ef3ef7e1461034b57806370a08231146103e85780637d64bcb4146104355780638456cb59146104625780638da5cb5b1461048f578063a3f4df7e146104e4578063a9059cbb14610572578063c0e37b15146105cc578063dd62ed3e14610669578063efef445b146106d5578063f2fde38b14610791578063f76f8d78146107ca575b600080fd5b341561012d57600080fd5b610135610858565b604051808215151515815260200191505060405180910390f35b341561015a57600080fd5b61018f600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190505061086b565b604051808215151515815260200191505060405180910390f35b34156101b457600080fd5b6101bc61089b565b6040518082815260200191505060405180910390f35b34156101dd57600080fd5b610231600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff169060200190919080359060200190919050506108a1565b604051808215151515815260200191505060405180910390f35b341561025657600080fd5b61025e6108d3565b6040518082815260200191505060405180910390f35b341561027f57600080fd5b6102876108d8565b604051808215151515815260200191505060405180910390f35b34156102ac57600080fd5b6102e1600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190505061099f565b604051808215151515815260200191505060405180910390f35b341561030657600080fd5b61031c6004808035906020019091905050610b21565b005b341561032957600080fd5b610331610cb9565b604051808215151515815260200191505060405180910390f35b341561035657600080fd5b6103ce600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610ccc565b604051808215151515815260200191505060405180910390f35b34156103f357600080fd5b61041f600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610cfe565b6040518082815260200191505060405180910390f35b341561044057600080fd5b610448610d47565b604051808215151515815260200191505060405180910390f35b341561046d57600080fd5b610475610df3565b604051808215151515815260200191505060405180910390f35b341561049a57600080fd5b6104a2610ebb565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156104ef57600080fd5b6104f7610ee1565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561053757808201518184015260208101905061051c565b50505050905090810190601f1680156105645780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561057d57600080fd5b6105b2600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091908035906020019091905050610f1a565b604051808215151515815260200191505060405180910390f35b34156105d757600080fd5b61064f600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610f4a565b604051808215151515815260200191505060405180910390f35b341561067457600080fd5b6106bf600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610f7c565b6040518082815260200191505060405180910390f35b34156106e057600080fd5b610777600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050611003565b604051808215151515815260200191505060405180910390f35b341561079c57600080fd5b6107c8600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611037565b005b34156107d557600080fd5b6107dd61110e565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561081d578082015181840152602081019050610802565b50505050905090810190601f16801561084a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b600360149054906101000a900460ff1681565b6000600360159054906101000a900460ff1615151561088957600080fd5b6108938383611147565b905092915050565b60005481565b6000600360159054906101000a900460ff161515156108bf57600080fd5b6108ca8484846112ce565b90509392505050565b601281565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561093657600080fd5b600360159054906101000a900460ff16151561095157600080fd5b6000600360156101000a81548160ff0219169083151502179055507f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a16001905090565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156109fd57600080fd5b600360149054906101000a900460ff16151515610a1957600080fd5b610a2e8260005461157e90919063ffffffff16565b600081905550610a8682600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461157e90919063ffffffff16565b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885836040518082815260200191505060405180910390a26001905092915050565b6000600360159054906101000a900460ff16151515610b3f57600080fd5b600082111515610b4e57600080fd5b339050610ba382600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461159c90919063ffffffff16565b600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610bfb8260005461159c90919063ffffffff16565b6000819055508073ffffffffffffffffffffffffffffffffffffffff167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5836040518082815260200191505060405180910390a2600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a35050565b600360159054906101000a900460ff1681565b6000600360159054906101000a900460ff16151515610cea57600080fd5b610cf58484846115b5565b90509392505050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610da557600080fd5b6001600360146101000a81548160ff0219169083151502179055507fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0860405160405180910390a16001905090565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610e5157600080fd5b600360159054906101000a900460ff16151515610e6d57600080fd5b6001600360156101000a81548160ff0219169083151502179055507f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a16001905090565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6040805190810160405280600481526020017f4cc3ad660000000000000000000000000000000000000000000000000000000081525081565b6000600360159054906101000a900460ff16151515610f3857600080fd5b610f42838361169e565b905092915050565b6000600360159054906101000a900460ff16151515610f6857600080fd5b610f73848484611839565b90509392505050565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6000600360159054906101000a900460ff1615151561102157600080fd5b61102d85858585611922565b9050949350505050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561109357600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151561110b5780600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b6040805190810160405280600381526020017f4c4946000000000000000000000000000000000000000000000000000000000081525081565b6000808214806111d357506000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054145b15156111de57600080fd5b81600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a36001905092915050565b600080600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506113a283600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461157e90919063ffffffff16565b600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061143783600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461159c90919063ffffffff16565b600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061148d838261159c90919063ffffffff16565b600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef856040518082815260200191505060405180910390a360019150509392505050565b600080828401905083811015151561159257fe5b8091505092915050565b60008282111515156115aa57fe5b818303905092915050565b60003073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16141515156115f257600080fd5b6115fc8484611147565b508373ffffffffffffffffffffffffffffffffffffffff168260405180828051906020019080838360005b83811015611642578082015181840152602081019050611627565b50505050905090810190601f16801561166f5780820380516001836020036101000a031916815260200191505b5091505060006040518083038160008661646e5a03f1915050151561169357600080fd5b600190509392505050565b60006116f282600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461159c90919063ffffffff16565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061178782600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461157e90919063ffffffff16565b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a36001905092915050565b60003073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415151561187657600080fd5b8373ffffffffffffffffffffffffffffffffffffffff168260405180828051906020019080838360005b838110156118bb5780820151818401526020810190506118a0565b50505050905090810190601f1680156118e85780820380516001836020036101000a031916815260200191505b5091505060006040518083038160008661646e5a03f1915050151561190c57600080fd5b611916848461169e565b50600190509392505050565b60003073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415151561195f57600080fd5b8373ffffffffffffffffffffffffffffffffffffffff168260405180828051906020019080838360005b838110156119a4578082015181840152602081019050611989565b50505050905090810190601f1680156119d15780820380516001836020036101000a031916815260200191505b5091505060006040518083038160008661646e5a03f191505015156119f557600080fd5b611a008585856112ce565b50600190509493505050505600a165627a7a72305820d1abfa68227be17b2043290567e53c4c8baf12ec06e72088b85343fa3ea834300029`

// DeployLifToken deploys a new Ethereum contract, binding an instance of LifToken to it.
func DeployLifToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LifToken, error) {
	parsed, err := abi.JSON(strings.NewReader(LifTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LifTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LifToken{LifTokenCaller: LifTokenCaller{contract: contract}, LifTokenTransactor: LifTokenTransactor{contract: contract}}, nil
}

// LifToken is an auto generated Go binding around an Ethereum contract.
type LifToken struct {
	LifTokenCaller     // Read-only binding to the contract
	LifTokenTransactor // Write-only binding to the contract
}

// LifTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type LifTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LifTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LifTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LifTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LifTokenSession struct {
	Contract     *LifToken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LifTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LifTokenCallerSession struct {
	Contract *LifTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// LifTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LifTokenTransactorSession struct {
	Contract     *LifTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LifTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type LifTokenRaw struct {
	Contract *LifToken // Generic contract binding to access the raw methods on
}

// LifTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LifTokenCallerRaw struct {
	Contract *LifTokenCaller // Generic read-only contract binding to access the raw methods on
}

// LifTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LifTokenTransactorRaw struct {
	Contract *LifTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLifToken creates a new instance of LifToken, bound to a specific deployed contract.
func NewLifToken(address common.Address, backend bind.ContractBackend) (*LifToken, error) {
	contract, err := bindLifToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LifToken{LifTokenCaller: LifTokenCaller{contract: contract}, LifTokenTransactor: LifTokenTransactor{contract: contract}}, nil
}

// NewLifTokenCaller creates a new read-only instance of LifToken, bound to a specific deployed contract.
func NewLifTokenCaller(address common.Address, caller bind.ContractCaller) (*LifTokenCaller, error) {
	contract, err := bindLifToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &LifTokenCaller{contract: contract}, nil
}

// NewLifTokenTransactor creates a new write-only instance of LifToken, bound to a specific deployed contract.
func NewLifTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*LifTokenTransactor, error) {
	contract, err := bindLifToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &LifTokenTransactor{contract: contract}, nil
}

// bindLifToken binds a generic wrapper to an already deployed contract.
func bindLifToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LifTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LifToken *LifTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LifToken.Contract.LifTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LifToken *LifTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LifToken.Contract.LifTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LifToken *LifTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LifToken.Contract.LifTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LifToken *LifTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LifToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LifToken *LifTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LifToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LifToken *LifTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LifToken.Contract.contract.Transact(opts, method, params...)
}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() constant returns(uint256)
func (_LifToken *LifTokenCaller) DECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LifToken.contract.Call(opts, out, "DECIMALS")
	return *ret0, err
}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() constant returns(uint256)
func (_LifToken *LifTokenSession) DECIMALS() (*big.Int, error) {
	return _LifToken.Contract.DECIMALS(&_LifToken.CallOpts)
}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() constant returns(uint256)
func (_LifToken *LifTokenCallerSession) DECIMALS() (*big.Int, error) {
	return _LifToken.Contract.DECIMALS(&_LifToken.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() constant returns(string)
func (_LifToken *LifTokenCaller) NAME(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _LifToken.contract.Call(opts, out, "NAME")
	return *ret0, err
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() constant returns(string)
func (_LifToken *LifTokenSession) NAME() (string, error) {
	return _LifToken.Contract.NAME(&_LifToken.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() constant returns(string)
func (_LifToken *LifTokenCallerSession) NAME() (string, error) {
	return _LifToken.Contract.NAME(&_LifToken.CallOpts)
}

// SYMBOL is a free data retrieval call binding the contract method 0xf76f8d78.
//
// Solidity: function SYMBOL() constant returns(string)
func (_LifToken *LifTokenCaller) SYMBOL(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _LifToken.contract.Call(opts, out, "SYMBOL")
	return *ret0, err
}

// SYMBOL is a free data retrieval call binding the contract method 0xf76f8d78.
//
// Solidity: function SYMBOL() constant returns(string)
func (_LifToken *LifTokenSession) SYMBOL() (string, error) {
	return _LifToken.Contract.SYMBOL(&_LifToken.CallOpts)
}

// SYMBOL is a free data retrieval call binding the contract method 0xf76f8d78.
//
// Solidity: function SYMBOL() constant returns(string)
func (_LifToken *LifTokenCallerSession) SYMBOL() (string, error) {
	return _LifToken.Contract.SYMBOL(&_LifToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_LifToken *LifTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LifToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_LifToken *LifTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _LifToken.Contract.Allowance(&_LifToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_LifToken *LifTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _LifToken.Contract.Allowance(&_LifToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_LifToken *LifTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LifToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_LifToken *LifTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _LifToken.Contract.BalanceOf(&_LifToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_LifToken *LifTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _LifToken.Contract.BalanceOf(&_LifToken.CallOpts, _owner)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_LifToken *LifTokenCaller) MintingFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LifToken.contract.Call(opts, out, "mintingFinished")
	return *ret0, err
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_LifToken *LifTokenSession) MintingFinished() (bool, error) {
	return _LifToken.Contract.MintingFinished(&_LifToken.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_LifToken *LifTokenCallerSession) MintingFinished() (bool, error) {
	return _LifToken.Contract.MintingFinished(&_LifToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_LifToken *LifTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LifToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_LifToken *LifTokenSession) Owner() (common.Address, error) {
	return _LifToken.Contract.Owner(&_LifToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_LifToken *LifTokenCallerSession) Owner() (common.Address, error) {
	return _LifToken.Contract.Owner(&_LifToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_LifToken *LifTokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LifToken.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_LifToken *LifTokenSession) Paused() (bool, error) {
	return _LifToken.Contract.Paused(&_LifToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_LifToken *LifTokenCallerSession) Paused() (bool, error) {
	return _LifToken.Contract.Paused(&_LifToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_LifToken *LifTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LifToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_LifToken *LifTokenSession) TotalSupply() (*big.Int, error) {
	return _LifToken.Contract.TotalSupply(&_LifToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_LifToken *LifTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _LifToken.Contract.TotalSupply(&_LifToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_LifToken *LifTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LifToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_LifToken *LifTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LifToken.Contract.Approve(&_LifToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_LifToken *LifTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LifToken.Contract.Approve(&_LifToken.TransactOpts, _spender, _value)
}

// ApproveData is a paid mutator transaction binding the contract method 0x6ef3ef7e.
//
// Solidity: function approveData(spender address, value uint256, data bytes) returns(bool)
func (_LifToken *LifTokenTransactor) ApproveData(opts *bind.TransactOpts, spender common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _LifToken.contract.Transact(opts, "approveData", spender, value, data)
}

// ApproveData is a paid mutator transaction binding the contract method 0x6ef3ef7e.
//
// Solidity: function approveData(spender address, value uint256, data bytes) returns(bool)
func (_LifToken *LifTokenSession) ApproveData(spender common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _LifToken.Contract.ApproveData(&_LifToken.TransactOpts, spender, value, data)
}

// ApproveData is a paid mutator transaction binding the contract method 0x6ef3ef7e.
//
// Solidity: function approveData(spender address, value uint256, data bytes) returns(bool)
func (_LifToken *LifTokenTransactorSession) ApproveData(spender common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _LifToken.Contract.ApproveData(&_LifToken.TransactOpts, spender, value, data)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_LifToken *LifTokenTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _LifToken.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_LifToken *LifTokenSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _LifToken.Contract.Burn(&_LifToken.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_LifToken *LifTokenTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _LifToken.Contract.Burn(&_LifToken.TransactOpts, _value)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_LifToken *LifTokenTransactor) FinishMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LifToken.contract.Transact(opts, "finishMinting")
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_LifToken *LifTokenSession) FinishMinting() (*types.Transaction, error) {
	return _LifToken.Contract.FinishMinting(&_LifToken.TransactOpts)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_LifToken *LifTokenTransactorSession) FinishMinting() (*types.Transaction, error) {
	return _LifToken.Contract.FinishMinting(&_LifToken.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_LifToken *LifTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LifToken.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_LifToken *LifTokenSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LifToken.Contract.Mint(&_LifToken.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_LifToken *LifTokenTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LifToken.Contract.Mint(&_LifToken.TransactOpts, _to, _amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_LifToken *LifTokenTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LifToken.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_LifToken *LifTokenSession) Pause() (*types.Transaction, error) {
	return _LifToken.Contract.Pause(&_LifToken.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_LifToken *LifTokenTransactorSession) Pause() (*types.Transaction, error) {
	return _LifToken.Contract.Pause(&_LifToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_LifToken *LifTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LifToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_LifToken *LifTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LifToken.Contract.Transfer(&_LifToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_LifToken *LifTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LifToken.Contract.Transfer(&_LifToken.TransactOpts, _to, _value)
}

// TransferData is a paid mutator transaction binding the contract method 0xc0e37b15.
//
// Solidity: function transferData(to address, value uint256, data bytes) returns(bool)
func (_LifToken *LifTokenTransactor) TransferData(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _LifToken.contract.Transact(opts, "transferData", to, value, data)
}

// TransferData is a paid mutator transaction binding the contract method 0xc0e37b15.
//
// Solidity: function transferData(to address, value uint256, data bytes) returns(bool)
func (_LifToken *LifTokenSession) TransferData(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _LifToken.Contract.TransferData(&_LifToken.TransactOpts, to, value, data)
}

// TransferData is a paid mutator transaction binding the contract method 0xc0e37b15.
//
// Solidity: function transferData(to address, value uint256, data bytes) returns(bool)
func (_LifToken *LifTokenTransactorSession) TransferData(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _LifToken.Contract.TransferData(&_LifToken.TransactOpts, to, value, data)
}

// TransferDataFrom is a paid mutator transaction binding the contract method 0xefef445b.
//
// Solidity: function transferDataFrom(from address, to address, value uint256, data bytes) returns(bool)
func (_LifToken *LifTokenTransactor) TransferDataFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _LifToken.contract.Transact(opts, "transferDataFrom", from, to, value, data)
}

// TransferDataFrom is a paid mutator transaction binding the contract method 0xefef445b.
//
// Solidity: function transferDataFrom(from address, to address, value uint256, data bytes) returns(bool)
func (_LifToken *LifTokenSession) TransferDataFrom(from common.Address, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _LifToken.Contract.TransferDataFrom(&_LifToken.TransactOpts, from, to, value, data)
}

// TransferDataFrom is a paid mutator transaction binding the contract method 0xefef445b.
//
// Solidity: function transferDataFrom(from address, to address, value uint256, data bytes) returns(bool)
func (_LifToken *LifTokenTransactorSession) TransferDataFrom(from common.Address, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _LifToken.Contract.TransferDataFrom(&_LifToken.TransactOpts, from, to, value, data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_LifToken *LifTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LifToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_LifToken *LifTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LifToken.Contract.TransferFrom(&_LifToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_LifToken *LifTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LifToken.Contract.TransferFrom(&_LifToken.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_LifToken *LifTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LifToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_LifToken *LifTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LifToken.Contract.TransferOwnership(&_LifToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_LifToken *LifTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LifToken.Contract.TransferOwnership(&_LifToken.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_LifToken *LifTokenTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LifToken.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_LifToken *LifTokenSession) Unpause() (*types.Transaction, error) {
	return _LifToken.Contract.Unpause(&_LifToken.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_LifToken *LifTokenTransactorSession) Unpause() (*types.Transaction, error) {
	return _LifToken.Contract.Unpause(&_LifToken.TransactOpts)
}
