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

// LifMarketValidationMechanismABI is the input ABI used to generate the binding from.
const LifMarketValidationMechanismABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getBuyPrice\",\"outputs\":[{\"name\":\"price\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalPausedSeconds\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lifToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMaxClaimableWeiAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialBuyPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"secondsPerPeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialWei\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalBurnedTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAccumulatedDistributionPercentage\",\"outputs\":[{\"name\":\"percentage\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinished\",\"outputs\":[{\"name\":\"finished\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pausedTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalReimbursedWei\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalWeiClaimed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentPeriodIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"fund\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"calculateDistributionPeriods\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"weiAmount\",\"type\":\"uint256\"}],\"name\":\"claimWei\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"originalTotalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"periods\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"funded\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"sendTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalPeriods\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"foundationAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"lifAddr\",\"type\":\"address\"},{\"name\":\"_startTimestamp\",\"type\":\"uint256\"},{\"name\":\"_secondsPerPeriod\",\"type\":\"uint256\"},{\"name\":\"_totalPeriods\",\"type\":\"uint8\"},{\"name\":\"_foundationAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"pausedSeconds\",\"type\":\"uint256\"}],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"claimedWei\",\"type\":\"uint256\"}],\"name\":\"ClaimedWei\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"returnedWei\",\"type\":\"uint256\"}],\"name\":\"SentTokens\",\"type\":\"event\"}]"

// LifMarketValidationMechanismBin is the compiled bytecode used for deploying new contracts.
const LifMarketValidationMechanismBin = `0x60606040526000600755600060085560006009556000600a556000600c60006101000a81548160ff0219169083151502179055506000600c60016101000a81548160ff0219169083151502179055506000600d55341561005e57600080fd5b60405160a0806119c583398101604052808051906020019091908051906020019091908051906020019091908051906020019091908051906020019091905050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161415151561011a57600080fd5b428411151561012857600080fd5b60008311151561013757600080fd5b60188260ff16148061014c575060308260ff16145b151561015757600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561019357600080fd5b84600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550836004819055508260058190555081600660006101000a81548160ff021916908360ff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505050611773806102526000396000f300606060405260043610610175576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063018a25e81461017a57806335a6c1e0146101a357806338fab8c5146101cc5780633ad6f8ac146102215780633ca6d5a91461024a5780633f4ba83a14610273578063407f8001146102885780635495794b146102b1578063555f323a146102da5780635c975abb146103035780636790f3fe146103305780637b352962146103595780638456cb59146103865780638da5cb5b1461039b578063911ef508146103f0578063a156ce7b14610419578063b30475b614610442578063b4f5a21a1461046b578063b60d428814610494578063c0670d2c1461049e578063ddd7c879146104b3578063df8f4eb7146104d6578063e6fd48bc146104ff578063ea4a110414610528578063f2fde38b1461055f578063f3a504f214610598578063f5c6ca08146105c5578063fea708f6146105e8578063feafb79b14610617575b600080fd5b341561018557600080fd5b61018d61066c565b6040518082815260200191505060405180910390f35b34156101ae57600080fd5b6101b66106be565b6040518082815260200191505060405180910390f35b34156101d757600080fd5b6101df6106c4565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561022c57600080fd5b6102346106ea565b6040518082815260200191505060405180910390f35b341561025557600080fd5b61025d6108ab565b6040518082815260200191505060405180910390f35b341561027e57600080fd5b6102866108b1565b005b341561029357600080fd5b61029b6109ad565b6040518082815260200191505060405180910390f35b34156102bc57600080fd5b6102c46109b3565b6040518082815260200191505060405180910390f35b34156102e557600080fd5b6102ed6109b9565b6040518082815260200191505060405180910390f35b341561030e57600080fd5b6103166109bf565b604051808215151515815260200191505060405180910390f35b341561033b57600080fd5b6103436109d2565b6040518082815260200191505060405180910390f35b341561036457600080fd5b61036c610a1d565b604051808215151515815260200191505060405180910390f35b341561039157600080fd5b610399610a41565b005b34156103a657600080fd5b6103ae610b05565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156103fb57600080fd5b610403610b2a565b6040518082815260200191505060405180910390f35b341561042457600080fd5b61042c610b30565b6040518082815260200191505060405180910390f35b341561044d57600080fd5b610455610b36565b6040518082815260200191505060405180910390f35b341561047657600080fd5b61047e610b3c565b6040518082815260200191505060405180910390f35b61049c610b8e565b005b34156104a957600080fd5b6104b1610d07565b005b34156104be57600080fd5b6104d46004808035906020019091905050611090565b005b34156104e157600080fd5b6104e96111d5565b6040518082815260200191505060405180910390f35b341561050a57600080fd5b6105126111db565b6040518082815260200191505060405180910390f35b341561053357600080fd5b61054960048080359060200190919050506111e1565b6040518082815260200191505060405180910390f35b341561056a57600080fd5b610596600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611205565b005b34156105a357600080fd5b6105ab6112da565b604051808215151515815260200191505060405180910390f35b34156105d057600080fd5b6105e660048080359060200190919050506112ed565b005b34156105f357600080fd5b6105fb6115e6565b604051808260ff1660ff16815260200191505060405180910390f35b341561062257600080fd5b61062a6115f9565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6000806106776109d2565b90506106b8620186a06106aa61069984620186a061161f90919063ffffffff16565b60085461163890919063ffffffff16565b61166b90919063ffffffff16565b91505090565b600d5481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060008060006106fa610a1d565b1561071e573073ffffffffffffffffffffffffffffffffffffffff163194506108a4565b61075e600a54610750620186a061074260095460085461163890919063ffffffff16565b61166b90919063ffffffff16565b61161f90919063ffffffff16565b9350600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318160ddd6000604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15156107ee57600080fd5b6102c65a03f115156107ff57600080fd5b5050506040518051905092506108136109d2565b915061087784610869600b5461085b8761084d620186a061083f8a60035461163890919063ffffffff16565b61166b90919063ffffffff16565b61163890919063ffffffff16565b61166b90919063ffffffff16565b61168690919063ffffffff16565b905060075481111561089f576108986007548261161f90919063ffffffff16565b94506108a4565b600094505b5050505090565b60085481565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561090e57600080fd5b600c60019054906101000a900460ff16151561092657fe5b61093b600e544261161f90919063ffffffff16565b905061095281600d5461168690919063ffffffff16565b600d819055506000600c60016101000a81548160ff0219169083151502179055507faaa520fdd7d2c83061d632fa017b0432407e798818af63ea908589fceda39ab7816040518082815260200191505060405180910390a150565b60055481565b60035481565b60095481565b600c60019054906101000a900460ff1681565b6000806109dd610b3c565b9050600660009054906101000a900460ff1660ff16811015156109fc57fe5b600f81815481101515610a0b57fe5b90600052602060002090015491505090565b6000600660009054906101000a900460ff1660ff16610a3a610b3c565b1015905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610a9c57600080fd5b600c60019054906101000a900460ff16151515610ab557fe5b6001600c60016101000a81548160ff02191690831515021790555042600e819055507f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600e5481565b600a5481565b60075481565b60006004544210151515610b4c57fe5b610b89600554610b7b600d54610b6d6004544261161f90919063ffffffff16565b61161f90919063ffffffff16565b61166b90919063ffffffff16565b905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610be957600080fd5b600c60009054906101000a900460ff16151515610c0257fe5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318160ddd6000604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1515610c9057600080fd5b6102c65a03f11515610ca157600080fd5b50505060405180519050600b8190555034600381905550610ce4600b54610cd6620186a060035461163890919063ffffffff16565b61166b90919063ffffffff16565b6008819055506001600c60006101000a81548160ff021916908315150217905550565b610d0f6116a4565b610d176116cd565b60006018600660009054906101000a900460ff1660ff161480610d4c57506030600660009054906101000a900460ff1660ff16145b1515610d5457fe5b6000600f80549050141515610d6557fe5b6103006040519081016040528060008152602001601281526020016075815260200161015f81526020016102ff815260200161057f81526020016109058152602001610db781526020016113b78152602001611b2881526020016124298152602001612edb8152602001613b5c81526020016149c98152602001615a408152602001616cde81526020016181bf81526020016198fe815260200161b2b5815260200161cf00815260200161edf9815260200162010fb981526020016201345a815260200162015bf4815250925061060060405190810160405280600081526020016003815260200160128152602001603681526020016075815260200160d6815260200161015f815260200161021681526020016102ff8152602001610420815260200161057e815260200161071e81526020016109048152602001610b358152602001610db6815260200161108a81526020016113b6815260200161173e8152602001611b268152602001611f73815260200161242881526020016129498152602001612eda81526020016134df8152602001613b5b815260200161425281526020016149c881526020016151c18152602001615a4081526020016163488152602001616cde815260200161770481526020016181be8152602001618d1081526020016198fd815260200161a588815260200161b2b5815260200161c086815260200161cf00815260200161de25815260200161edf9815260200161fe7e815260200162010fb88152602001620121ab8152602001620134598152602001620147c5815260200162015bf38152602001620170e68152509150600090505b600660009054906101000a900460ff1660ff168160ff16101561108b576018600660009054906101000a900460ff1660ff16141561103f57600f805480600101828161101091906116f6565b91600052602060002090016000858460ff1660188110151561102e57fe5b60200201519091909150555061107e565b600f805480600101828161105391906116f6565b91600052602060002090016000848460ff1660308110151561107157fe5b6020020151909190915055505b8080600101915050610fc4565b505050565b6000600c60019054906101000a900460ff161515156110ab57fe5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561110757600080fd5b61110f6106ea565b905081811015151561111d57fe5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f19350505050151561117f57600080fd5b6111948260075461168690919063ffffffff16565b6007819055507f7995ed8c8bb70e086ac77eabe37bd8742685022b74d12ac20d7629469b5374e5826040518082815260200191505060405180910390a15050565b600b5481565b60045481565b600f818154811015156111f057fe5b90600052602060002090016000915090505481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561126057600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415156112d757806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b600c60009054906101000a900460ff1681565b600080600c60019054906101000a900460ff1615151561130957fe5b60008311151561131857600080fd5b61132061066c565b915061134a620186a061133c848661163890919063ffffffff16565b61166b90919063ffffffff16565b9050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330866000604051602001526040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b151561144d57600080fd5b6102c65a03f1151561145e57600080fd5b5050506040518051905050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166342966c68846040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050600060405180830381600087803b15156114f957600080fd5b6102c65a03f1151561150a57600080fd5b5050506115228360095461168690919063ffffffff16565b6009819055503373ffffffffffffffffffffffffffffffffffffffff167f1e3ea5698ac6d5bb5cde5c6a3764daa2ef39b16b2062c0ded43333188a5851c083858460405180848152602001838152602001828152602001935050505060405180910390a261159b81600a5461168690919063ffffffff16565b600a819055503373ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f1935050505015156115e157600080fd5b505050565b600660009054906101000a900460ff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600082821115151561162d57fe5b818303905092915050565b60008082840290506000841480611659575082848281151561165657fe5b04145b151561166157fe5b8091505092915050565b600080828481151561167957fe5b0490508091505092915050565b600080828401905083811015151561169a57fe5b8091505092915050565b610300604051908101604052806018905b60008152602001906001900390816116b55790505090565b610600604051908101604052806030905b60008152602001906001900390816116de5790505090565b81548183558181151161171d5781836000526020600020918201910161171c9190611722565b5b505050565b61174491905b80821115611740576000816000905550600101611728565b5090565b905600a165627a7a72305820b651eac14a696269f05266b46274bc6a47b35c07e370135b07a7f4245398b9c40029`

// DeployLifMarketValidationMechanism deploys a new Ethereum contract, binding an instance of LifMarketValidationMechanism to it.
func DeployLifMarketValidationMechanism(auth *bind.TransactOpts, backend bind.ContractBackend, lifAddr common.Address, _startTimestamp *big.Int, _secondsPerPeriod *big.Int, _totalPeriods uint8, _foundationAddr common.Address) (common.Address, *types.Transaction, *LifMarketValidationMechanism, error) {
	parsed, err := abi.JSON(strings.NewReader(LifMarketValidationMechanismABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LifMarketValidationMechanismBin), backend, lifAddr, _startTimestamp, _secondsPerPeriod, _totalPeriods, _foundationAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LifMarketValidationMechanism{LifMarketValidationMechanismCaller: LifMarketValidationMechanismCaller{contract: contract}, LifMarketValidationMechanismTransactor: LifMarketValidationMechanismTransactor{contract: contract}}, nil
}

// LifMarketValidationMechanism is an auto generated Go binding around an Ethereum contract.
type LifMarketValidationMechanism struct {
	LifMarketValidationMechanismCaller     // Read-only binding to the contract
	LifMarketValidationMechanismTransactor // Write-only binding to the contract
}

// LifMarketValidationMechanismCaller is an auto generated read-only Go binding around an Ethereum contract.
type LifMarketValidationMechanismCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LifMarketValidationMechanismTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LifMarketValidationMechanismTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LifMarketValidationMechanismSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LifMarketValidationMechanismSession struct {
	Contract     *LifMarketValidationMechanism // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// LifMarketValidationMechanismCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LifMarketValidationMechanismCallerSession struct {
	Contract *LifMarketValidationMechanismCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// LifMarketValidationMechanismTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LifMarketValidationMechanismTransactorSession struct {
	Contract     *LifMarketValidationMechanismTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// LifMarketValidationMechanismRaw is an auto generated low-level Go binding around an Ethereum contract.
type LifMarketValidationMechanismRaw struct {
	Contract *LifMarketValidationMechanism // Generic contract binding to access the raw methods on
}

// LifMarketValidationMechanismCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LifMarketValidationMechanismCallerRaw struct {
	Contract *LifMarketValidationMechanismCaller // Generic read-only contract binding to access the raw methods on
}

// LifMarketValidationMechanismTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LifMarketValidationMechanismTransactorRaw struct {
	Contract *LifMarketValidationMechanismTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLifMarketValidationMechanism creates a new instance of LifMarketValidationMechanism, bound to a specific deployed contract.
func NewLifMarketValidationMechanism(address common.Address, backend bind.ContractBackend) (*LifMarketValidationMechanism, error) {
	contract, err := bindLifMarketValidationMechanism(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LifMarketValidationMechanism{LifMarketValidationMechanismCaller: LifMarketValidationMechanismCaller{contract: contract}, LifMarketValidationMechanismTransactor: LifMarketValidationMechanismTransactor{contract: contract}}, nil
}

// NewLifMarketValidationMechanismCaller creates a new read-only instance of LifMarketValidationMechanism, bound to a specific deployed contract.
func NewLifMarketValidationMechanismCaller(address common.Address, caller bind.ContractCaller) (*LifMarketValidationMechanismCaller, error) {
	contract, err := bindLifMarketValidationMechanism(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &LifMarketValidationMechanismCaller{contract: contract}, nil
}

// NewLifMarketValidationMechanismTransactor creates a new write-only instance of LifMarketValidationMechanism, bound to a specific deployed contract.
func NewLifMarketValidationMechanismTransactor(address common.Address, transactor bind.ContractTransactor) (*LifMarketValidationMechanismTransactor, error) {
	contract, err := bindLifMarketValidationMechanism(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &LifMarketValidationMechanismTransactor{contract: contract}, nil
}

// bindLifMarketValidationMechanism binds a generic wrapper to an already deployed contract.
func bindLifMarketValidationMechanism(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LifMarketValidationMechanismABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LifMarketValidationMechanism *LifMarketValidationMechanismRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LifMarketValidationMechanism.Contract.LifMarketValidationMechanismCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LifMarketValidationMechanism *LifMarketValidationMechanismRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LifMarketValidationMechanism.Contract.LifMarketValidationMechanismTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LifMarketValidationMechanism *LifMarketValidationMechanismRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LifMarketValidationMechanism.Contract.LifMarketValidationMechanismTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LifMarketValidationMechanism.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LifMarketValidationMechanism *LifMarketValidationMechanismTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LifMarketValidationMechanism.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LifMarketValidationMechanism *LifMarketValidationMechanismTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LifMarketValidationMechanism.Contract.contract.Transact(opts, method, params...)
}

// FoundationAddr is a free data retrieval call binding the contract method 0xfeafb79b.
//
// Solidity: function foundationAddr() constant returns(address)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCaller) FoundationAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LifMarketValidationMechanism.contract.Call(opts, out, "foundationAddr")
	return *ret0, err
}

// FoundationAddr is a free data retrieval call binding the contract method 0xfeafb79b.
//
// Solidity: function foundationAddr() constant returns(address)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismSession) FoundationAddr() (common.Address, error) {
	return _LifMarketValidationMechanism.Contract.FoundationAddr(&_LifMarketValidationMechanism.CallOpts)
}

// FoundationAddr is a free data retrieval call binding the contract method 0xfeafb79b.
//
// Solidity: function foundationAddr() constant returns(address)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCallerSession) FoundationAddr() (common.Address, error) {
	return _LifMarketValidationMechanism.Contract.FoundationAddr(&_LifMarketValidationMechanism.CallOpts)
}

// Funded is a free data retrieval call binding the contract method 0xf3a504f2.
//
// Solidity: function funded() constant returns(bool)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCaller) Funded(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LifMarketValidationMechanism.contract.Call(opts, out, "funded")
	return *ret0, err
}

// Funded is a free data retrieval call binding the contract method 0xf3a504f2.
//
// Solidity: function funded() constant returns(bool)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismSession) Funded() (bool, error) {
	return _LifMarketValidationMechanism.Contract.Funded(&_LifMarketValidationMechanism.CallOpts)
}

// Funded is a free data retrieval call binding the contract method 0xf3a504f2.
//
// Solidity: function funded() constant returns(bool)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCallerSession) Funded() (bool, error) {
	return _LifMarketValidationMechanism.Contract.Funded(&_LifMarketValidationMechanism.CallOpts)
}

// GetAccumulatedDistributionPercentage is a free data retrieval call binding the contract method 0x6790f3fe.
//
// Solidity: function getAccumulatedDistributionPercentage() constant returns(percentage uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCaller) GetAccumulatedDistributionPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LifMarketValidationMechanism.contract.Call(opts, out, "getAccumulatedDistributionPercentage")
	return *ret0, err
}

// GetAccumulatedDistributionPercentage is a free data retrieval call binding the contract method 0x6790f3fe.
//
// Solidity: function getAccumulatedDistributionPercentage() constant returns(percentage uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismSession) GetAccumulatedDistributionPercentage() (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.GetAccumulatedDistributionPercentage(&_LifMarketValidationMechanism.CallOpts)
}

// GetAccumulatedDistributionPercentage is a free data retrieval call binding the contract method 0x6790f3fe.
//
// Solidity: function getAccumulatedDistributionPercentage() constant returns(percentage uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCallerSession) GetAccumulatedDistributionPercentage() (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.GetAccumulatedDistributionPercentage(&_LifMarketValidationMechanism.CallOpts)
}

// GetBuyPrice is a free data retrieval call binding the contract method 0x018a25e8.
//
// Solidity: function getBuyPrice() constant returns(price uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCaller) GetBuyPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LifMarketValidationMechanism.contract.Call(opts, out, "getBuyPrice")
	return *ret0, err
}

// GetBuyPrice is a free data retrieval call binding the contract method 0x018a25e8.
//
// Solidity: function getBuyPrice() constant returns(price uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismSession) GetBuyPrice() (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.GetBuyPrice(&_LifMarketValidationMechanism.CallOpts)
}

// GetBuyPrice is a free data retrieval call binding the contract method 0x018a25e8.
//
// Solidity: function getBuyPrice() constant returns(price uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCallerSession) GetBuyPrice() (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.GetBuyPrice(&_LifMarketValidationMechanism.CallOpts)
}

// GetCurrentPeriodIndex is a free data retrieval call binding the contract method 0xb4f5a21a.
//
// Solidity: function getCurrentPeriodIndex() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCaller) GetCurrentPeriodIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LifMarketValidationMechanism.contract.Call(opts, out, "getCurrentPeriodIndex")
	return *ret0, err
}

// GetCurrentPeriodIndex is a free data retrieval call binding the contract method 0xb4f5a21a.
//
// Solidity: function getCurrentPeriodIndex() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismSession) GetCurrentPeriodIndex() (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.GetCurrentPeriodIndex(&_LifMarketValidationMechanism.CallOpts)
}

// GetCurrentPeriodIndex is a free data retrieval call binding the contract method 0xb4f5a21a.
//
// Solidity: function getCurrentPeriodIndex() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCallerSession) GetCurrentPeriodIndex() (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.GetCurrentPeriodIndex(&_LifMarketValidationMechanism.CallOpts)
}

// GetMaxClaimableWeiAmount is a free data retrieval call binding the contract method 0x3ad6f8ac.
//
// Solidity: function getMaxClaimableWeiAmount() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCaller) GetMaxClaimableWeiAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LifMarketValidationMechanism.contract.Call(opts, out, "getMaxClaimableWeiAmount")
	return *ret0, err
}

// GetMaxClaimableWeiAmount is a free data retrieval call binding the contract method 0x3ad6f8ac.
//
// Solidity: function getMaxClaimableWeiAmount() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismSession) GetMaxClaimableWeiAmount() (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.GetMaxClaimableWeiAmount(&_LifMarketValidationMechanism.CallOpts)
}

// GetMaxClaimableWeiAmount is a free data retrieval call binding the contract method 0x3ad6f8ac.
//
// Solidity: function getMaxClaimableWeiAmount() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCallerSession) GetMaxClaimableWeiAmount() (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.GetMaxClaimableWeiAmount(&_LifMarketValidationMechanism.CallOpts)
}

// InitialBuyPrice is a free data retrieval call binding the contract method 0x3ca6d5a9.
//
// Solidity: function initialBuyPrice() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCaller) InitialBuyPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LifMarketValidationMechanism.contract.Call(opts, out, "initialBuyPrice")
	return *ret0, err
}

// InitialBuyPrice is a free data retrieval call binding the contract method 0x3ca6d5a9.
//
// Solidity: function initialBuyPrice() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismSession) InitialBuyPrice() (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.InitialBuyPrice(&_LifMarketValidationMechanism.CallOpts)
}

// InitialBuyPrice is a free data retrieval call binding the contract method 0x3ca6d5a9.
//
// Solidity: function initialBuyPrice() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCallerSession) InitialBuyPrice() (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.InitialBuyPrice(&_LifMarketValidationMechanism.CallOpts)
}

// InitialWei is a free data retrieval call binding the contract method 0x5495794b.
//
// Solidity: function initialWei() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCaller) InitialWei(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LifMarketValidationMechanism.contract.Call(opts, out, "initialWei")
	return *ret0, err
}

// InitialWei is a free data retrieval call binding the contract method 0x5495794b.
//
// Solidity: function initialWei() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismSession) InitialWei() (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.InitialWei(&_LifMarketValidationMechanism.CallOpts)
}

// InitialWei is a free data retrieval call binding the contract method 0x5495794b.
//
// Solidity: function initialWei() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCallerSession) InitialWei() (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.InitialWei(&_LifMarketValidationMechanism.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(finished bool)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCaller) IsFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LifMarketValidationMechanism.contract.Call(opts, out, "isFinished")
	return *ret0, err
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(finished bool)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismSession) IsFinished() (bool, error) {
	return _LifMarketValidationMechanism.Contract.IsFinished(&_LifMarketValidationMechanism.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(finished bool)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCallerSession) IsFinished() (bool, error) {
	return _LifMarketValidationMechanism.Contract.IsFinished(&_LifMarketValidationMechanism.CallOpts)
}

// LifToken is a free data retrieval call binding the contract method 0x38fab8c5.
//
// Solidity: function lifToken() constant returns(address)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCaller) LifToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LifMarketValidationMechanism.contract.Call(opts, out, "lifToken")
	return *ret0, err
}

// LifToken is a free data retrieval call binding the contract method 0x38fab8c5.
//
// Solidity: function lifToken() constant returns(address)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismSession) LifToken() (common.Address, error) {
	return _LifMarketValidationMechanism.Contract.LifToken(&_LifMarketValidationMechanism.CallOpts)
}

// LifToken is a free data retrieval call binding the contract method 0x38fab8c5.
//
// Solidity: function lifToken() constant returns(address)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCallerSession) LifToken() (common.Address, error) {
	return _LifMarketValidationMechanism.Contract.LifToken(&_LifMarketValidationMechanism.CallOpts)
}

// OriginalTotalSupply is a free data retrieval call binding the contract method 0xdf8f4eb7.
//
// Solidity: function originalTotalSupply() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCaller) OriginalTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LifMarketValidationMechanism.contract.Call(opts, out, "originalTotalSupply")
	return *ret0, err
}

// OriginalTotalSupply is a free data retrieval call binding the contract method 0xdf8f4eb7.
//
// Solidity: function originalTotalSupply() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismSession) OriginalTotalSupply() (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.OriginalTotalSupply(&_LifMarketValidationMechanism.CallOpts)
}

// OriginalTotalSupply is a free data retrieval call binding the contract method 0xdf8f4eb7.
//
// Solidity: function originalTotalSupply() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCallerSession) OriginalTotalSupply() (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.OriginalTotalSupply(&_LifMarketValidationMechanism.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LifMarketValidationMechanism.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismSession) Owner() (common.Address, error) {
	return _LifMarketValidationMechanism.Contract.Owner(&_LifMarketValidationMechanism.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCallerSession) Owner() (common.Address, error) {
	return _LifMarketValidationMechanism.Contract.Owner(&_LifMarketValidationMechanism.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LifMarketValidationMechanism.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismSession) Paused() (bool, error) {
	return _LifMarketValidationMechanism.Contract.Paused(&_LifMarketValidationMechanism.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCallerSession) Paused() (bool, error) {
	return _LifMarketValidationMechanism.Contract.Paused(&_LifMarketValidationMechanism.CallOpts)
}

// PausedTimestamp is a free data retrieval call binding the contract method 0x911ef508.
//
// Solidity: function pausedTimestamp() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCaller) PausedTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LifMarketValidationMechanism.contract.Call(opts, out, "pausedTimestamp")
	return *ret0, err
}

// PausedTimestamp is a free data retrieval call binding the contract method 0x911ef508.
//
// Solidity: function pausedTimestamp() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismSession) PausedTimestamp() (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.PausedTimestamp(&_LifMarketValidationMechanism.CallOpts)
}

// PausedTimestamp is a free data retrieval call binding the contract method 0x911ef508.
//
// Solidity: function pausedTimestamp() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCallerSession) PausedTimestamp() (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.PausedTimestamp(&_LifMarketValidationMechanism.CallOpts)
}

// Periods is a free data retrieval call binding the contract method 0xea4a1104.
//
// Solidity: function periods( uint256) constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCaller) Periods(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LifMarketValidationMechanism.contract.Call(opts, out, "periods", arg0)
	return *ret0, err
}

// Periods is a free data retrieval call binding the contract method 0xea4a1104.
//
// Solidity: function periods( uint256) constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismSession) Periods(arg0 *big.Int) (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.Periods(&_LifMarketValidationMechanism.CallOpts, arg0)
}

// Periods is a free data retrieval call binding the contract method 0xea4a1104.
//
// Solidity: function periods( uint256) constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCallerSession) Periods(arg0 *big.Int) (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.Periods(&_LifMarketValidationMechanism.CallOpts, arg0)
}

// SecondsPerPeriod is a free data retrieval call binding the contract method 0x407f8001.
//
// Solidity: function secondsPerPeriod() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCaller) SecondsPerPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LifMarketValidationMechanism.contract.Call(opts, out, "secondsPerPeriod")
	return *ret0, err
}

// SecondsPerPeriod is a free data retrieval call binding the contract method 0x407f8001.
//
// Solidity: function secondsPerPeriod() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismSession) SecondsPerPeriod() (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.SecondsPerPeriod(&_LifMarketValidationMechanism.CallOpts)
}

// SecondsPerPeriod is a free data retrieval call binding the contract method 0x407f8001.
//
// Solidity: function secondsPerPeriod() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCallerSession) SecondsPerPeriod() (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.SecondsPerPeriod(&_LifMarketValidationMechanism.CallOpts)
}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCaller) StartTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LifMarketValidationMechanism.contract.Call(opts, out, "startTimestamp")
	return *ret0, err
}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismSession) StartTimestamp() (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.StartTimestamp(&_LifMarketValidationMechanism.CallOpts)
}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCallerSession) StartTimestamp() (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.StartTimestamp(&_LifMarketValidationMechanism.CallOpts)
}

// TotalBurnedTokens is a free data retrieval call binding the contract method 0x555f323a.
//
// Solidity: function totalBurnedTokens() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCaller) TotalBurnedTokens(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LifMarketValidationMechanism.contract.Call(opts, out, "totalBurnedTokens")
	return *ret0, err
}

// TotalBurnedTokens is a free data retrieval call binding the contract method 0x555f323a.
//
// Solidity: function totalBurnedTokens() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismSession) TotalBurnedTokens() (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.TotalBurnedTokens(&_LifMarketValidationMechanism.CallOpts)
}

// TotalBurnedTokens is a free data retrieval call binding the contract method 0x555f323a.
//
// Solidity: function totalBurnedTokens() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCallerSession) TotalBurnedTokens() (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.TotalBurnedTokens(&_LifMarketValidationMechanism.CallOpts)
}

// TotalPausedSeconds is a free data retrieval call binding the contract method 0x35a6c1e0.
//
// Solidity: function totalPausedSeconds() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCaller) TotalPausedSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LifMarketValidationMechanism.contract.Call(opts, out, "totalPausedSeconds")
	return *ret0, err
}

// TotalPausedSeconds is a free data retrieval call binding the contract method 0x35a6c1e0.
//
// Solidity: function totalPausedSeconds() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismSession) TotalPausedSeconds() (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.TotalPausedSeconds(&_LifMarketValidationMechanism.CallOpts)
}

// TotalPausedSeconds is a free data retrieval call binding the contract method 0x35a6c1e0.
//
// Solidity: function totalPausedSeconds() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCallerSession) TotalPausedSeconds() (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.TotalPausedSeconds(&_LifMarketValidationMechanism.CallOpts)
}

// TotalPeriods is a free data retrieval call binding the contract method 0xfea708f6.
//
// Solidity: function totalPeriods() constant returns(uint8)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCaller) TotalPeriods(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _LifMarketValidationMechanism.contract.Call(opts, out, "totalPeriods")
	return *ret0, err
}

// TotalPeriods is a free data retrieval call binding the contract method 0xfea708f6.
//
// Solidity: function totalPeriods() constant returns(uint8)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismSession) TotalPeriods() (uint8, error) {
	return _LifMarketValidationMechanism.Contract.TotalPeriods(&_LifMarketValidationMechanism.CallOpts)
}

// TotalPeriods is a free data retrieval call binding the contract method 0xfea708f6.
//
// Solidity: function totalPeriods() constant returns(uint8)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCallerSession) TotalPeriods() (uint8, error) {
	return _LifMarketValidationMechanism.Contract.TotalPeriods(&_LifMarketValidationMechanism.CallOpts)
}

// TotalReimbursedWei is a free data retrieval call binding the contract method 0xa156ce7b.
//
// Solidity: function totalReimbursedWei() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCaller) TotalReimbursedWei(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LifMarketValidationMechanism.contract.Call(opts, out, "totalReimbursedWei")
	return *ret0, err
}

// TotalReimbursedWei is a free data retrieval call binding the contract method 0xa156ce7b.
//
// Solidity: function totalReimbursedWei() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismSession) TotalReimbursedWei() (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.TotalReimbursedWei(&_LifMarketValidationMechanism.CallOpts)
}

// TotalReimbursedWei is a free data retrieval call binding the contract method 0xa156ce7b.
//
// Solidity: function totalReimbursedWei() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCallerSession) TotalReimbursedWei() (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.TotalReimbursedWei(&_LifMarketValidationMechanism.CallOpts)
}

// TotalWeiClaimed is a free data retrieval call binding the contract method 0xb30475b6.
//
// Solidity: function totalWeiClaimed() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCaller) TotalWeiClaimed(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LifMarketValidationMechanism.contract.Call(opts, out, "totalWeiClaimed")
	return *ret0, err
}

// TotalWeiClaimed is a free data retrieval call binding the contract method 0xb30475b6.
//
// Solidity: function totalWeiClaimed() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismSession) TotalWeiClaimed() (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.TotalWeiClaimed(&_LifMarketValidationMechanism.CallOpts)
}

// TotalWeiClaimed is a free data retrieval call binding the contract method 0xb30475b6.
//
// Solidity: function totalWeiClaimed() constant returns(uint256)
func (_LifMarketValidationMechanism *LifMarketValidationMechanismCallerSession) TotalWeiClaimed() (*big.Int, error) {
	return _LifMarketValidationMechanism.Contract.TotalWeiClaimed(&_LifMarketValidationMechanism.CallOpts)
}

// CalculateDistributionPeriods is a paid mutator transaction binding the contract method 0xc0670d2c.
//
// Solidity: function calculateDistributionPeriods() returns()
func (_LifMarketValidationMechanism *LifMarketValidationMechanismTransactor) CalculateDistributionPeriods(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LifMarketValidationMechanism.contract.Transact(opts, "calculateDistributionPeriods")
}

// CalculateDistributionPeriods is a paid mutator transaction binding the contract method 0xc0670d2c.
//
// Solidity: function calculateDistributionPeriods() returns()
func (_LifMarketValidationMechanism *LifMarketValidationMechanismSession) CalculateDistributionPeriods() (*types.Transaction, error) {
	return _LifMarketValidationMechanism.Contract.CalculateDistributionPeriods(&_LifMarketValidationMechanism.TransactOpts)
}

// CalculateDistributionPeriods is a paid mutator transaction binding the contract method 0xc0670d2c.
//
// Solidity: function calculateDistributionPeriods() returns()
func (_LifMarketValidationMechanism *LifMarketValidationMechanismTransactorSession) CalculateDistributionPeriods() (*types.Transaction, error) {
	return _LifMarketValidationMechanism.Contract.CalculateDistributionPeriods(&_LifMarketValidationMechanism.TransactOpts)
}

// ClaimWei is a paid mutator transaction binding the contract method 0xddd7c879.
//
// Solidity: function claimWei(weiAmount uint256) returns()
func (_LifMarketValidationMechanism *LifMarketValidationMechanismTransactor) ClaimWei(opts *bind.TransactOpts, weiAmount *big.Int) (*types.Transaction, error) {
	return _LifMarketValidationMechanism.contract.Transact(opts, "claimWei", weiAmount)
}

// ClaimWei is a paid mutator transaction binding the contract method 0xddd7c879.
//
// Solidity: function claimWei(weiAmount uint256) returns()
func (_LifMarketValidationMechanism *LifMarketValidationMechanismSession) ClaimWei(weiAmount *big.Int) (*types.Transaction, error) {
	return _LifMarketValidationMechanism.Contract.ClaimWei(&_LifMarketValidationMechanism.TransactOpts, weiAmount)
}

// ClaimWei is a paid mutator transaction binding the contract method 0xddd7c879.
//
// Solidity: function claimWei(weiAmount uint256) returns()
func (_LifMarketValidationMechanism *LifMarketValidationMechanismTransactorSession) ClaimWei(weiAmount *big.Int) (*types.Transaction, error) {
	return _LifMarketValidationMechanism.Contract.ClaimWei(&_LifMarketValidationMechanism.TransactOpts, weiAmount)
}

// Fund is a paid mutator transaction binding the contract method 0xb60d4288.
//
// Solidity: function fund() returns()
func (_LifMarketValidationMechanism *LifMarketValidationMechanismTransactor) Fund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LifMarketValidationMechanism.contract.Transact(opts, "fund")
}

// Fund is a paid mutator transaction binding the contract method 0xb60d4288.
//
// Solidity: function fund() returns()
func (_LifMarketValidationMechanism *LifMarketValidationMechanismSession) Fund() (*types.Transaction, error) {
	return _LifMarketValidationMechanism.Contract.Fund(&_LifMarketValidationMechanism.TransactOpts)
}

// Fund is a paid mutator transaction binding the contract method 0xb60d4288.
//
// Solidity: function fund() returns()
func (_LifMarketValidationMechanism *LifMarketValidationMechanismTransactorSession) Fund() (*types.Transaction, error) {
	return _LifMarketValidationMechanism.Contract.Fund(&_LifMarketValidationMechanism.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LifMarketValidationMechanism *LifMarketValidationMechanismTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LifMarketValidationMechanism.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LifMarketValidationMechanism *LifMarketValidationMechanismSession) Pause() (*types.Transaction, error) {
	return _LifMarketValidationMechanism.Contract.Pause(&_LifMarketValidationMechanism.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LifMarketValidationMechanism *LifMarketValidationMechanismTransactorSession) Pause() (*types.Transaction, error) {
	return _LifMarketValidationMechanism.Contract.Pause(&_LifMarketValidationMechanism.TransactOpts)
}

// SendTokens is a paid mutator transaction binding the contract method 0xf5c6ca08.
//
// Solidity: function sendTokens(tokens uint256) returns()
func (_LifMarketValidationMechanism *LifMarketValidationMechanismTransactor) SendTokens(opts *bind.TransactOpts, tokens *big.Int) (*types.Transaction, error) {
	return _LifMarketValidationMechanism.contract.Transact(opts, "sendTokens", tokens)
}

// SendTokens is a paid mutator transaction binding the contract method 0xf5c6ca08.
//
// Solidity: function sendTokens(tokens uint256) returns()
func (_LifMarketValidationMechanism *LifMarketValidationMechanismSession) SendTokens(tokens *big.Int) (*types.Transaction, error) {
	return _LifMarketValidationMechanism.Contract.SendTokens(&_LifMarketValidationMechanism.TransactOpts, tokens)
}

// SendTokens is a paid mutator transaction binding the contract method 0xf5c6ca08.
//
// Solidity: function sendTokens(tokens uint256) returns()
func (_LifMarketValidationMechanism *LifMarketValidationMechanismTransactorSession) SendTokens(tokens *big.Int) (*types.Transaction, error) {
	return _LifMarketValidationMechanism.Contract.SendTokens(&_LifMarketValidationMechanism.TransactOpts, tokens)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_LifMarketValidationMechanism *LifMarketValidationMechanismTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LifMarketValidationMechanism.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_LifMarketValidationMechanism *LifMarketValidationMechanismSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LifMarketValidationMechanism.Contract.TransferOwnership(&_LifMarketValidationMechanism.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_LifMarketValidationMechanism *LifMarketValidationMechanismTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LifMarketValidationMechanism.Contract.TransferOwnership(&_LifMarketValidationMechanism.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LifMarketValidationMechanism *LifMarketValidationMechanismTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LifMarketValidationMechanism.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LifMarketValidationMechanism *LifMarketValidationMechanismSession) Unpause() (*types.Transaction, error) {
	return _LifMarketValidationMechanism.Contract.Unpause(&_LifMarketValidationMechanism.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LifMarketValidationMechanism *LifMarketValidationMechanismTransactorSession) Unpause() (*types.Transaction, error) {
	return _LifMarketValidationMechanism.Contract.Unpause(&_LifMarketValidationMechanism.TransactOpts)
}
