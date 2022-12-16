// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package v2

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// NameRegistryMetaData contains all meta data concerning the NameRegistry contract.
var NameRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"protocol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"}],\"name\":\"ConnectorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NameOwnershipUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NameRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"PriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"PublicKeyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"Sale\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"protocol\",\"type\":\"string\"}],\"name\":\"lookupConnector\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"lookupName\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"protocol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"}],\"name\":\"registerConnector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"registerName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"updateOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"updatePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"updatePublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611228806100206000396000f3fe60806040526004361061007b5760003560e01c8063639cdb951161004e578063639cdb95146100f5578063a3ddd6911461012d578063d346e3bd1461014d578063f11a1a9a1461017a57600080fd5b8063208b5f38146100805780634a432a46146100a257806356c792db146100c25780635d2097d2146100e2575b600080fd5b34801561008c57600080fd5b506100a061009b366004610b0c565b61019a565b005b3480156100ae57600080fd5b506100a06100bd366004610b78565b610258565b3480156100ce57600080fd5b506100a06100dd366004610b0c565b610302565b6100a06100f0366004610bda565b6105df565b34801561010157600080fd5b50610115610110366004610cb5565b61078f565b60405161012493929190610d3d565b60405180910390f35b34801561013957600080fd5b506100a0610148366004610d71565b61086e565b34801561015957600080fd5b5061016d610168366004610b0c565b610927565b6040516101249190610dd4565b34801561018657600080fd5b506100a0610195366004610dee565b6109fa565b602183146101a757600080fd5b336001600160a01b0316600083836040516101c3929190610e88565b908152604051908190036020019020600101546001600160a01b0316146101e957600080fd5b8383600084846040516101fd929190610e88565b90815260405190819003602001902091610218919083610f21565b507f0ba4a7233a2cfb937aa9644f11f49345b443fe8d5ab3494a1879faa500728d3e828260405161024a92919061100b565b60405180910390a150505050565b336001600160a01b031660008484604051610274929190610e88565b908152604051908190036020019020600101546001600160a01b03161461029a57600080fd5b80600084846040516102ad929190610e88565b9081526020016040518091039020600201819055507f159e83f4712ba2552e68be9d848e49bf6dd35c24f19564ffd523b6549450a2f48383836040516102f593929190611027565b60405180910390a1505050565b6021831461030f57600080fd5b60006001600160a01b03166000838360405161032c929190610e88565b908152604051908190036020019020600101546001600160a01b03161461035257600080fd5b600181101561036057600080fd5b604081111561036e57600080fd5b818160008181106103815761038161104b565b909101356001600160f81b031916602d60f81b0390506103a057600080fd5b600082828080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250939450600193925050505b82518110156105185760008382815181106103fc576103fc61104b565b01602001516001600160f81b0319169050602d60f81b8114806104465750600360fc1b6001600160f81b03198216108015906104465750603960f81b6001600160f81b0319821611155b8061045e5750605f60f81b6001600160f81b03198216145b806104905750606160f81b6001600160f81b03198216108015906104905750603d60f91b6001600160f81b0319821611155b6104cf5760405162461bcd60e51b815260206004820152600c60248201526b696e76616c6964206e616d6560a01b604482015260640160405180910390fd5b602d60f81b6001600160f81b03198216148015906104fb5750605f60f81b6001600160f81b0319821614155b1561050557600092505b508061051081611061565b9150506103df565b50801561052457600080fd5b3360008585604051610537929190610e88565b908152602001604051809103902060010160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550858560008686604051610580929190610e88565b9081526040519081900360200190209161059b919083610f21565b507f1c6eac0e720ec22bb0653aec9c19985633a4fb07971cf973096c2f8e3c37c17f8484336040516105cf93929190611088565b60405180910390a1505050505050565b60008084846040516105f2929190610e88565b9081526020016040518091039020600201541161060e57600080fd5b3460008484604051610621929190610e88565b908152602001604051809103902060020154111561063e57600080fd5b6000808484604051610651929190610e88565b908152602001604051809103902060020154905060008484604051610677929190610e88565b908152604051908190036020018120600101546001600160a01b0316903480156108fc02916000818181858888f193505050501580156106bb573d6000803e3d6000fd5b5033600085856040516106cf929190610e88565b908152602001604051809103902060010160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508160008585604051610717929190610e88565b9081526040519081900360200190209061073190826110b4565b506000808585604051610745929190610e88565b9081526020016040518091039020600201819055507f74d699b6db148dc632a2dbe2c680980f546b0382af2e71a7d4bf51c3193adb9a8484833360405161024a9493929190611174565b60006060600080600086866040516107a8929190610e88565b9081526040519081900360200190206001810154600282015482549293506001600160a01b039091169183919082906107e090610e98565b80601f016020809104026020016040519081016040528092919081815260200182805461080c90610e98565b80156108595780601f1061082e57610100808354040283529160200191610859565b820191906000526020600020905b81548152906001019060200180831161083c57829003601f168201915b50505050509150935093509350509250925092565b336001600160a01b03166000848460405161088a929190610e88565b908152604051908190036020019020600101546001600160a01b0316146108b057600080fd5b80600084846040516108c3929190610e88565b90815260405190819003602001812060010180546001600160a01b03939093166001600160a01b0319909316929092179091557f179b5e07786d37fc59a40f1e1e8a697240c6e378cf1e17cf9438b0110abeb6d7906102f590859085908590611088565b60606001858560405161093b929190610e88565b90815260200160405180910390208383604051610959929190610e88565b9081526020016040518091039020805461097290610e98565b80601f016020809104026020016040519081016040528092919081815260200182805461099e90610e98565b80156109eb5780601f106109c0576101008083540402835291602001916109eb565b820191906000526020600020905b8154815290600101906020018083116109ce57829003601f168201915b50505050509050949350505050565b336001600160a01b031660008787604051610a16929190610e88565b908152604051908190036020019020600101546001600160a01b031614610a3c57600080fd5b818160018888604051610a50929190610e88565b90815260200160405180910390208686604051610a6e929190610e88565b90815260200160405180910390209182610a89929190610f21565b507f6bcbdc9c6b89c8a6b289cc4308c6f6841a0e25d81c5fe34b1dedb2848c5f5ca38686868686866040516105cf969594939291906111a9565b60008083601f840112610ad557600080fd5b50813567ffffffffffffffff811115610aed57600080fd5b602083019150836020828501011115610b0557600080fd5b9250929050565b60008060008060408587031215610b2257600080fd5b843567ffffffffffffffff80821115610b3a57600080fd5b610b4688838901610ac3565b90965094506020870135915080821115610b5f57600080fd5b50610b6c87828801610ac3565b95989497509550505050565b600080600060408486031215610b8d57600080fd5b833567ffffffffffffffff811115610ba457600080fd5b610bb086828701610ac3565b909790965060209590950135949350505050565b634e487b7160e01b600052604160045260246000fd5b600080600060408486031215610bef57600080fd5b833567ffffffffffffffff80821115610c0757600080fd5b610c1387838801610ac3565b90955093506020860135915080821115610c2c57600080fd5b818601915086601f830112610c4057600080fd5b813581811115610c5257610c52610bc4565b604051601f8201601f19908116603f01168101908382118183101715610c7a57610c7a610bc4565b81604052828152896020848701011115610c9357600080fd5b8260208601602083013760006020848301015280955050505050509250925092565b60008060208385031215610cc857600080fd5b823567ffffffffffffffff811115610cdf57600080fd5b610ceb85828601610ac3565b90969095509350505050565b6000815180845260005b81811015610d1d57602081850181015186830182015201610d01565b506000602082860101526020601f19601f83011685010191505092915050565b6001600160a01b0384168152606060208201819052600090610d6190830185610cf7565b9050826040830152949350505050565b600080600060408486031215610d8657600080fd5b833567ffffffffffffffff811115610d9d57600080fd5b610da986828701610ac3565b90945092505060208401356001600160a01b0381168114610dc957600080fd5b809150509250925092565b602081526000610de76020830184610cf7565b9392505050565b60008060008060008060608789031215610e0757600080fd5b863567ffffffffffffffff80821115610e1f57600080fd5b610e2b8a838b01610ac3565b90985096506020890135915080821115610e4457600080fd5b610e508a838b01610ac3565b90965094506040890135915080821115610e6957600080fd5b50610e7689828a01610ac3565b979a9699509497509295939492505050565b8183823760009101908152919050565b600181811c90821680610eac57607f821691505b602082108103610ecc57634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115610f1c57600081815260208120601f850160051c81016020861015610ef95750805b601f850160051c820191505b81811015610f1857828155600101610f05565b5050505b505050565b67ffffffffffffffff831115610f3957610f39610bc4565b610f4d83610f478354610e98565b83610ed2565b6000601f841160018114610f815760008515610f695750838201355b600019600387901b1c1916600186901b178355610fdb565b600083815260209020601f19861690835b82811015610fb25786850135825560209485019460019092019101610f92565b5086821015610fcf5760001960f88860031b161c19848701351681555b505060018560011b0183555b5050505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60208152600061101f602083018486610fe2565b949350505050565b60408152600061103b604083018587610fe2565b9050826020830152949350505050565b634e487b7160e01b600052603260045260246000fd5b60006001820161108157634e487b7160e01b600052601160045260246000fd5b5060010190565b60408152600061109c604083018587610fe2565b905060018060a01b0383166020830152949350505050565b815167ffffffffffffffff8111156110ce576110ce610bc4565b6110e2816110dc8454610e98565b84610ed2565b602080601f83116001811461111757600084156110ff5750858301515b600019600386901b1c1916600185901b178555610f18565b600085815260208120601f198616915b8281101561114657888601518255948401946001909101908401611127565b50858210156111645787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b606081526000611188606083018688610fe2565b6020830194909452506001600160a01b039190911660409091015292915050565b6060815260006111bd60608301888a610fe2565b82810360208401526111d0818789610fe2565b905082810360408401526111e5818587610fe2565b999850505050505050505056fea2646970667358221220643fbd611451d690204bb5df431f45da7d4738c0bbf8262fe17a2e16fb923ba064736f6c63430008110033",
}

// NameRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use NameRegistryMetaData.ABI instead.
var NameRegistryABI = NameRegistryMetaData.ABI

// NameRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NameRegistryMetaData.Bin instead.
var NameRegistryBin = NameRegistryMetaData.Bin

// DeployNameRegistry deploys a new Ethereum contract, binding an instance of NameRegistry to it.
func DeployNameRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NameRegistry, error) {
	parsed, err := NameRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NameRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NameRegistry{NameRegistryCaller: NameRegistryCaller{contract: contract}, NameRegistryTransactor: NameRegistryTransactor{contract: contract}, NameRegistryFilterer: NameRegistryFilterer{contract: contract}}, nil
}

// NameRegistry is an auto generated Go binding around an Ethereum contract.
type NameRegistry struct {
	NameRegistryCaller     // Read-only binding to the contract
	NameRegistryTransactor // Write-only binding to the contract
	NameRegistryFilterer   // Log filterer for contract events
}

// NameRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type NameRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NameRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NameRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NameRegistrySession struct {
	Contract     *NameRegistry     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NameRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NameRegistryCallerSession struct {
	Contract *NameRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// NameRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NameRegistryTransactorSession struct {
	Contract     *NameRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// NameRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type NameRegistryRaw struct {
	Contract *NameRegistry // Generic contract binding to access the raw methods on
}

// NameRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NameRegistryCallerRaw struct {
	Contract *NameRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// NameRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NameRegistryTransactorRaw struct {
	Contract *NameRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNameRegistry creates a new instance of NameRegistry, bound to a specific deployed contract.
func NewNameRegistry(address common.Address, backend bind.ContractBackend) (*NameRegistry, error) {
	contract, err := bindNameRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NameRegistry{NameRegistryCaller: NameRegistryCaller{contract: contract}, NameRegistryTransactor: NameRegistryTransactor{contract: contract}, NameRegistryFilterer: NameRegistryFilterer{contract: contract}}, nil
}

// NewNameRegistryCaller creates a new read-only instance of NameRegistry, bound to a specific deployed contract.
func NewNameRegistryCaller(address common.Address, caller bind.ContractCaller) (*NameRegistryCaller, error) {
	contract, err := bindNameRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NameRegistryCaller{contract: contract}, nil
}

// NewNameRegistryTransactor creates a new write-only instance of NameRegistry, bound to a specific deployed contract.
func NewNameRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*NameRegistryTransactor, error) {
	contract, err := bindNameRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NameRegistryTransactor{contract: contract}, nil
}

// NewNameRegistryFilterer creates a new log filterer instance of NameRegistry, bound to a specific deployed contract.
func NewNameRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*NameRegistryFilterer, error) {
	contract, err := bindNameRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NameRegistryFilterer{contract: contract}, nil
}

// bindNameRegistry binds a generic wrapper to an already deployed contract.
func bindNameRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NameRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NameRegistry *NameRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NameRegistry.Contract.NameRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NameRegistry *NameRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NameRegistry.Contract.NameRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NameRegistry *NameRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NameRegistry.Contract.NameRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NameRegistry *NameRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NameRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NameRegistry *NameRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NameRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NameRegistry *NameRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NameRegistry.Contract.contract.Transact(opts, method, params...)
}

// LookupConnector is a free data retrieval call binding the contract method 0xd346e3bd.
//
// Solidity: function lookupConnector(string name, string protocol) view returns(string)
func (_NameRegistry *NameRegistryCaller) LookupConnector(opts *bind.CallOpts, name string, protocol string) (string, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "lookupConnector", name, protocol)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LookupConnector is a free data retrieval call binding the contract method 0xd346e3bd.
//
// Solidity: function lookupConnector(string name, string protocol) view returns(string)
func (_NameRegistry *NameRegistrySession) LookupConnector(name string, protocol string) (string, error) {
	return _NameRegistry.Contract.LookupConnector(&_NameRegistry.CallOpts, name, protocol)
}

// LookupConnector is a free data retrieval call binding the contract method 0xd346e3bd.
//
// Solidity: function lookupConnector(string name, string protocol) view returns(string)
func (_NameRegistry *NameRegistryCallerSession) LookupConnector(name string, protocol string) (string, error) {
	return _NameRegistry.Contract.LookupConnector(&_NameRegistry.CallOpts, name, protocol)
}

// LookupName is a free data retrieval call binding the contract method 0x639cdb95.
//
// Solidity: function lookupName(string name) view returns(address owner, bytes publicKey, uint256 price)
func (_NameRegistry *NameRegistryCaller) LookupName(opts *bind.CallOpts, name string) (struct {
	Owner     common.Address
	PublicKey []byte
	Price     *big.Int
}, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "lookupName", name)

	outstruct := new(struct {
		Owner     common.Address
		PublicKey []byte
		Price     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.PublicKey = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.Price = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LookupName is a free data retrieval call binding the contract method 0x639cdb95.
//
// Solidity: function lookupName(string name) view returns(address owner, bytes publicKey, uint256 price)
func (_NameRegistry *NameRegistrySession) LookupName(name string) (struct {
	Owner     common.Address
	PublicKey []byte
	Price     *big.Int
}, error) {
	return _NameRegistry.Contract.LookupName(&_NameRegistry.CallOpts, name)
}

// LookupName is a free data retrieval call binding the contract method 0x639cdb95.
//
// Solidity: function lookupName(string name) view returns(address owner, bytes publicKey, uint256 price)
func (_NameRegistry *NameRegistryCallerSession) LookupName(name string) (struct {
	Owner     common.Address
	PublicKey []byte
	Price     *big.Int
}, error) {
	return _NameRegistry.Contract.LookupName(&_NameRegistry.CallOpts, name)
}

// Buy is a paid mutator transaction binding the contract method 0x5d2097d2.
//
// Solidity: function buy(string name, bytes publicKey) payable returns()
func (_NameRegistry *NameRegistryTransactor) Buy(opts *bind.TransactOpts, name string, publicKey []byte) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "buy", name, publicKey)
}

// Buy is a paid mutator transaction binding the contract method 0x5d2097d2.
//
// Solidity: function buy(string name, bytes publicKey) payable returns()
func (_NameRegistry *NameRegistrySession) Buy(name string, publicKey []byte) (*types.Transaction, error) {
	return _NameRegistry.Contract.Buy(&_NameRegistry.TransactOpts, name, publicKey)
}

// Buy is a paid mutator transaction binding the contract method 0x5d2097d2.
//
// Solidity: function buy(string name, bytes publicKey) payable returns()
func (_NameRegistry *NameRegistryTransactorSession) Buy(name string, publicKey []byte) (*types.Transaction, error) {
	return _NameRegistry.Contract.Buy(&_NameRegistry.TransactOpts, name, publicKey)
}

// RegisterConnector is a paid mutator transaction binding the contract method 0xf11a1a9a.
//
// Solidity: function registerConnector(string name, string protocol, string location) returns()
func (_NameRegistry *NameRegistryTransactor) RegisterConnector(opts *bind.TransactOpts, name string, protocol string, location string) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "registerConnector", name, protocol, location)
}

// RegisterConnector is a paid mutator transaction binding the contract method 0xf11a1a9a.
//
// Solidity: function registerConnector(string name, string protocol, string location) returns()
func (_NameRegistry *NameRegistrySession) RegisterConnector(name string, protocol string, location string) (*types.Transaction, error) {
	return _NameRegistry.Contract.RegisterConnector(&_NameRegistry.TransactOpts, name, protocol, location)
}

// RegisterConnector is a paid mutator transaction binding the contract method 0xf11a1a9a.
//
// Solidity: function registerConnector(string name, string protocol, string location) returns()
func (_NameRegistry *NameRegistryTransactorSession) RegisterConnector(name string, protocol string, location string) (*types.Transaction, error) {
	return _NameRegistry.Contract.RegisterConnector(&_NameRegistry.TransactOpts, name, protocol, location)
}

// RegisterName is a paid mutator transaction binding the contract method 0x56c792db.
//
// Solidity: function registerName(bytes publicKey, string name) returns()
func (_NameRegistry *NameRegistryTransactor) RegisterName(opts *bind.TransactOpts, publicKey []byte, name string) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "registerName", publicKey, name)
}

// RegisterName is a paid mutator transaction binding the contract method 0x56c792db.
//
// Solidity: function registerName(bytes publicKey, string name) returns()
func (_NameRegistry *NameRegistrySession) RegisterName(publicKey []byte, name string) (*types.Transaction, error) {
	return _NameRegistry.Contract.RegisterName(&_NameRegistry.TransactOpts, publicKey, name)
}

// RegisterName is a paid mutator transaction binding the contract method 0x56c792db.
//
// Solidity: function registerName(bytes publicKey, string name) returns()
func (_NameRegistry *NameRegistryTransactorSession) RegisterName(publicKey []byte, name string) (*types.Transaction, error) {
	return _NameRegistry.Contract.RegisterName(&_NameRegistry.TransactOpts, publicKey, name)
}

// UpdateOwnership is a paid mutator transaction binding the contract method 0xa3ddd691.
//
// Solidity: function updateOwnership(string name, address newOwner) returns()
func (_NameRegistry *NameRegistryTransactor) UpdateOwnership(opts *bind.TransactOpts, name string, newOwner common.Address) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "updateOwnership", name, newOwner)
}

// UpdateOwnership is a paid mutator transaction binding the contract method 0xa3ddd691.
//
// Solidity: function updateOwnership(string name, address newOwner) returns()
func (_NameRegistry *NameRegistrySession) UpdateOwnership(name string, newOwner common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.UpdateOwnership(&_NameRegistry.TransactOpts, name, newOwner)
}

// UpdateOwnership is a paid mutator transaction binding the contract method 0xa3ddd691.
//
// Solidity: function updateOwnership(string name, address newOwner) returns()
func (_NameRegistry *NameRegistryTransactorSession) UpdateOwnership(name string, newOwner common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.UpdateOwnership(&_NameRegistry.TransactOpts, name, newOwner)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x4a432a46.
//
// Solidity: function updatePrice(string name, uint256 price) returns()
func (_NameRegistry *NameRegistryTransactor) UpdatePrice(opts *bind.TransactOpts, name string, price *big.Int) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "updatePrice", name, price)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x4a432a46.
//
// Solidity: function updatePrice(string name, uint256 price) returns()
func (_NameRegistry *NameRegistrySession) UpdatePrice(name string, price *big.Int) (*types.Transaction, error) {
	return _NameRegistry.Contract.UpdatePrice(&_NameRegistry.TransactOpts, name, price)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x4a432a46.
//
// Solidity: function updatePrice(string name, uint256 price) returns()
func (_NameRegistry *NameRegistryTransactorSession) UpdatePrice(name string, price *big.Int) (*types.Transaction, error) {
	return _NameRegistry.Contract.UpdatePrice(&_NameRegistry.TransactOpts, name, price)
}

// UpdatePublicKey is a paid mutator transaction binding the contract method 0x208b5f38.
//
// Solidity: function updatePublicKey(bytes publicKey, string name) returns()
func (_NameRegistry *NameRegistryTransactor) UpdatePublicKey(opts *bind.TransactOpts, publicKey []byte, name string) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "updatePublicKey", publicKey, name)
}

// UpdatePublicKey is a paid mutator transaction binding the contract method 0x208b5f38.
//
// Solidity: function updatePublicKey(bytes publicKey, string name) returns()
func (_NameRegistry *NameRegistrySession) UpdatePublicKey(publicKey []byte, name string) (*types.Transaction, error) {
	return _NameRegistry.Contract.UpdatePublicKey(&_NameRegistry.TransactOpts, publicKey, name)
}

// UpdatePublicKey is a paid mutator transaction binding the contract method 0x208b5f38.
//
// Solidity: function updatePublicKey(bytes publicKey, string name) returns()
func (_NameRegistry *NameRegistryTransactorSession) UpdatePublicKey(publicKey []byte, name string) (*types.Transaction, error) {
	return _NameRegistry.Contract.UpdatePublicKey(&_NameRegistry.TransactOpts, publicKey, name)
}

// NameRegistryConnectorRegisteredIterator is returned from FilterConnectorRegistered and is used to iterate over the raw logs and unpacked data for ConnectorRegistered events raised by the NameRegistry contract.
type NameRegistryConnectorRegisteredIterator struct {
	Event *NameRegistryConnectorRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NameRegistryConnectorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryConnectorRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NameRegistryConnectorRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NameRegistryConnectorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryConnectorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryConnectorRegistered represents a ConnectorRegistered event raised by the NameRegistry contract.
type NameRegistryConnectorRegistered struct {
	Name     string
	Protocol string
	Location string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterConnectorRegistered is a free log retrieval operation binding the contract event 0x6bcbdc9c6b89c8a6b289cc4308c6f6841a0e25d81c5fe34b1dedb2848c5f5ca3.
//
// Solidity: event ConnectorRegistered(string name, string protocol, string location)
func (_NameRegistry *NameRegistryFilterer) FilterConnectorRegistered(opts *bind.FilterOpts) (*NameRegistryConnectorRegisteredIterator, error) {

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "ConnectorRegistered")
	if err != nil {
		return nil, err
	}
	return &NameRegistryConnectorRegisteredIterator{contract: _NameRegistry.contract, event: "ConnectorRegistered", logs: logs, sub: sub}, nil
}

// WatchConnectorRegistered is a free log subscription operation binding the contract event 0x6bcbdc9c6b89c8a6b289cc4308c6f6841a0e25d81c5fe34b1dedb2848c5f5ca3.
//
// Solidity: event ConnectorRegistered(string name, string protocol, string location)
func (_NameRegistry *NameRegistryFilterer) WatchConnectorRegistered(opts *bind.WatchOpts, sink chan<- *NameRegistryConnectorRegistered) (event.Subscription, error) {

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "ConnectorRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryConnectorRegistered)
				if err := _NameRegistry.contract.UnpackLog(event, "ConnectorRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConnectorRegistered is a log parse operation binding the contract event 0x6bcbdc9c6b89c8a6b289cc4308c6f6841a0e25d81c5fe34b1dedb2848c5f5ca3.
//
// Solidity: event ConnectorRegistered(string name, string protocol, string location)
func (_NameRegistry *NameRegistryFilterer) ParseConnectorRegistered(log types.Log) (*NameRegistryConnectorRegistered, error) {
	event := new(NameRegistryConnectorRegistered)
	if err := _NameRegistry.contract.UnpackLog(event, "ConnectorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistryNameOwnershipUpdatedIterator is returned from FilterNameOwnershipUpdated and is used to iterate over the raw logs and unpacked data for NameOwnershipUpdated events raised by the NameRegistry contract.
type NameRegistryNameOwnershipUpdatedIterator struct {
	Event *NameRegistryNameOwnershipUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NameRegistryNameOwnershipUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryNameOwnershipUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NameRegistryNameOwnershipUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NameRegistryNameOwnershipUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryNameOwnershipUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryNameOwnershipUpdated represents a NameOwnershipUpdated event raised by the NameRegistry contract.
type NameRegistryNameOwnershipUpdated struct {
	Name     string
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNameOwnershipUpdated is a free log retrieval operation binding the contract event 0x179b5e07786d37fc59a40f1e1e8a697240c6e378cf1e17cf9438b0110abeb6d7.
//
// Solidity: event NameOwnershipUpdated(string name, address newOwner)
func (_NameRegistry *NameRegistryFilterer) FilterNameOwnershipUpdated(opts *bind.FilterOpts) (*NameRegistryNameOwnershipUpdatedIterator, error) {

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "NameOwnershipUpdated")
	if err != nil {
		return nil, err
	}
	return &NameRegistryNameOwnershipUpdatedIterator{contract: _NameRegistry.contract, event: "NameOwnershipUpdated", logs: logs, sub: sub}, nil
}

// WatchNameOwnershipUpdated is a free log subscription operation binding the contract event 0x179b5e07786d37fc59a40f1e1e8a697240c6e378cf1e17cf9438b0110abeb6d7.
//
// Solidity: event NameOwnershipUpdated(string name, address newOwner)
func (_NameRegistry *NameRegistryFilterer) WatchNameOwnershipUpdated(opts *bind.WatchOpts, sink chan<- *NameRegistryNameOwnershipUpdated) (event.Subscription, error) {

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "NameOwnershipUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryNameOwnershipUpdated)
				if err := _NameRegistry.contract.UnpackLog(event, "NameOwnershipUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNameOwnershipUpdated is a log parse operation binding the contract event 0x179b5e07786d37fc59a40f1e1e8a697240c6e378cf1e17cf9438b0110abeb6d7.
//
// Solidity: event NameOwnershipUpdated(string name, address newOwner)
func (_NameRegistry *NameRegistryFilterer) ParseNameOwnershipUpdated(log types.Log) (*NameRegistryNameOwnershipUpdated, error) {
	event := new(NameRegistryNameOwnershipUpdated)
	if err := _NameRegistry.contract.UnpackLog(event, "NameOwnershipUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistryNameRegisteredIterator is returned from FilterNameRegistered and is used to iterate over the raw logs and unpacked data for NameRegistered events raised by the NameRegistry contract.
type NameRegistryNameRegisteredIterator struct {
	Event *NameRegistryNameRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NameRegistryNameRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryNameRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NameRegistryNameRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NameRegistryNameRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryNameRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryNameRegistered represents a NameRegistered event raised by the NameRegistry contract.
type NameRegistryNameRegistered struct {
	Name  string
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNameRegistered is a free log retrieval operation binding the contract event 0x1c6eac0e720ec22bb0653aec9c19985633a4fb07971cf973096c2f8e3c37c17f.
//
// Solidity: event NameRegistered(string name, address owner)
func (_NameRegistry *NameRegistryFilterer) FilterNameRegistered(opts *bind.FilterOpts) (*NameRegistryNameRegisteredIterator, error) {

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "NameRegistered")
	if err != nil {
		return nil, err
	}
	return &NameRegistryNameRegisteredIterator{contract: _NameRegistry.contract, event: "NameRegistered", logs: logs, sub: sub}, nil
}

// WatchNameRegistered is a free log subscription operation binding the contract event 0x1c6eac0e720ec22bb0653aec9c19985633a4fb07971cf973096c2f8e3c37c17f.
//
// Solidity: event NameRegistered(string name, address owner)
func (_NameRegistry *NameRegistryFilterer) WatchNameRegistered(opts *bind.WatchOpts, sink chan<- *NameRegistryNameRegistered) (event.Subscription, error) {

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "NameRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryNameRegistered)
				if err := _NameRegistry.contract.UnpackLog(event, "NameRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNameRegistered is a log parse operation binding the contract event 0x1c6eac0e720ec22bb0653aec9c19985633a4fb07971cf973096c2f8e3c37c17f.
//
// Solidity: event NameRegistered(string name, address owner)
func (_NameRegistry *NameRegistryFilterer) ParseNameRegistered(log types.Log) (*NameRegistryNameRegistered, error) {
	event := new(NameRegistryNameRegistered)
	if err := _NameRegistry.contract.UnpackLog(event, "NameRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistryPriceUpdatedIterator is returned from FilterPriceUpdated and is used to iterate over the raw logs and unpacked data for PriceUpdated events raised by the NameRegistry contract.
type NameRegistryPriceUpdatedIterator struct {
	Event *NameRegistryPriceUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NameRegistryPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryPriceUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NameRegistryPriceUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NameRegistryPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryPriceUpdated represents a PriceUpdated event raised by the NameRegistry contract.
type NameRegistryPriceUpdated struct {
	Name  string
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPriceUpdated is a free log retrieval operation binding the contract event 0x159e83f4712ba2552e68be9d848e49bf6dd35c24f19564ffd523b6549450a2f4.
//
// Solidity: event PriceUpdated(string name, uint256 price)
func (_NameRegistry *NameRegistryFilterer) FilterPriceUpdated(opts *bind.FilterOpts) (*NameRegistryPriceUpdatedIterator, error) {

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "PriceUpdated")
	if err != nil {
		return nil, err
	}
	return &NameRegistryPriceUpdatedIterator{contract: _NameRegistry.contract, event: "PriceUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceUpdated is a free log subscription operation binding the contract event 0x159e83f4712ba2552e68be9d848e49bf6dd35c24f19564ffd523b6549450a2f4.
//
// Solidity: event PriceUpdated(string name, uint256 price)
func (_NameRegistry *NameRegistryFilterer) WatchPriceUpdated(opts *bind.WatchOpts, sink chan<- *NameRegistryPriceUpdated) (event.Subscription, error) {

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "PriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryPriceUpdated)
				if err := _NameRegistry.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePriceUpdated is a log parse operation binding the contract event 0x159e83f4712ba2552e68be9d848e49bf6dd35c24f19564ffd523b6549450a2f4.
//
// Solidity: event PriceUpdated(string name, uint256 price)
func (_NameRegistry *NameRegistryFilterer) ParsePriceUpdated(log types.Log) (*NameRegistryPriceUpdated, error) {
	event := new(NameRegistryPriceUpdated)
	if err := _NameRegistry.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistryPublicKeyUpdatedIterator is returned from FilterPublicKeyUpdated and is used to iterate over the raw logs and unpacked data for PublicKeyUpdated events raised by the NameRegistry contract.
type NameRegistryPublicKeyUpdatedIterator struct {
	Event *NameRegistryPublicKeyUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NameRegistryPublicKeyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryPublicKeyUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NameRegistryPublicKeyUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NameRegistryPublicKeyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryPublicKeyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryPublicKeyUpdated represents a PublicKeyUpdated event raised by the NameRegistry contract.
type NameRegistryPublicKeyUpdated struct {
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPublicKeyUpdated is a free log retrieval operation binding the contract event 0x0ba4a7233a2cfb937aa9644f11f49345b443fe8d5ab3494a1879faa500728d3e.
//
// Solidity: event PublicKeyUpdated(string name)
func (_NameRegistry *NameRegistryFilterer) FilterPublicKeyUpdated(opts *bind.FilterOpts) (*NameRegistryPublicKeyUpdatedIterator, error) {

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "PublicKeyUpdated")
	if err != nil {
		return nil, err
	}
	return &NameRegistryPublicKeyUpdatedIterator{contract: _NameRegistry.contract, event: "PublicKeyUpdated", logs: logs, sub: sub}, nil
}

// WatchPublicKeyUpdated is a free log subscription operation binding the contract event 0x0ba4a7233a2cfb937aa9644f11f49345b443fe8d5ab3494a1879faa500728d3e.
//
// Solidity: event PublicKeyUpdated(string name)
func (_NameRegistry *NameRegistryFilterer) WatchPublicKeyUpdated(opts *bind.WatchOpts, sink chan<- *NameRegistryPublicKeyUpdated) (event.Subscription, error) {

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "PublicKeyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryPublicKeyUpdated)
				if err := _NameRegistry.contract.UnpackLog(event, "PublicKeyUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePublicKeyUpdated is a log parse operation binding the contract event 0x0ba4a7233a2cfb937aa9644f11f49345b443fe8d5ab3494a1879faa500728d3e.
//
// Solidity: event PublicKeyUpdated(string name)
func (_NameRegistry *NameRegistryFilterer) ParsePublicKeyUpdated(log types.Log) (*NameRegistryPublicKeyUpdated, error) {
	event := new(NameRegistryPublicKeyUpdated)
	if err := _NameRegistry.contract.UnpackLog(event, "PublicKeyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistrySaleIterator is returned from FilterSale and is used to iterate over the raw logs and unpacked data for Sale events raised by the NameRegistry contract.
type NameRegistrySaleIterator struct {
	Event *NameRegistrySale // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NameRegistrySaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistrySale)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NameRegistrySale)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NameRegistrySaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistrySaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistrySale represents a Sale event raised by the NameRegistry contract.
type NameRegistrySale struct {
	Name     string
	Price    *big.Int
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSale is a free log retrieval operation binding the contract event 0x74d699b6db148dc632a2dbe2c680980f546b0382af2e71a7d4bf51c3193adb9a.
//
// Solidity: event Sale(string name, uint256 price, address newOwner)
func (_NameRegistry *NameRegistryFilterer) FilterSale(opts *bind.FilterOpts) (*NameRegistrySaleIterator, error) {

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "Sale")
	if err != nil {
		return nil, err
	}
	return &NameRegistrySaleIterator{contract: _NameRegistry.contract, event: "Sale", logs: logs, sub: sub}, nil
}

// WatchSale is a free log subscription operation binding the contract event 0x74d699b6db148dc632a2dbe2c680980f546b0382af2e71a7d4bf51c3193adb9a.
//
// Solidity: event Sale(string name, uint256 price, address newOwner)
func (_NameRegistry *NameRegistryFilterer) WatchSale(opts *bind.WatchOpts, sink chan<- *NameRegistrySale) (event.Subscription, error) {

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "Sale")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistrySale)
				if err := _NameRegistry.contract.UnpackLog(event, "Sale", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSale is a log parse operation binding the contract event 0x74d699b6db148dc632a2dbe2c680980f546b0382af2e71a7d4bf51c3193adb9a.
//
// Solidity: event Sale(string name, uint256 price, address newOwner)
func (_NameRegistry *NameRegistryFilterer) ParseSale(log types.Log) (*NameRegistrySale, error) {
	event := new(NameRegistrySale)
	if err := _NameRegistry.contract.UnpackLog(event, "Sale", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
