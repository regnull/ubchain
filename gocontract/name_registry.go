// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gocontract

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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keyRegistryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161091a38038061091a83398101604081905261002f91610054565b600080546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b610887806100936000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806350db6b401461004657806392a296c91461005b578063d37aec9214610084575b600080fd5b610059610054366004610513565b610097565b005b61006e610069366004610595565b6102d9565b60405161007b9190610646565b60405180910390f35b61006e61009236600461069b565b61037e565b602181146100a457600080fd5b60038310156100b257600080fd5b60408311156100c057600080fd5b600054604051632d654a9760e11b81526001600160a01b0390911690635aca952e906100f290859085906004016106dd565b602060405180830381865afa15801561010f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610133919061070c565b61013c57600080fd5b600060018585604051610150929190610735565b9081526020016040518091039020805461016990610745565b9050111561021e5760005460405133916001600160a01b03169063bd27b241906001906101999089908990610735565b9081526040519081900360200181206001600160e01b031960e084901b1682526101c591600401610780565b602060405180830381865afa1580156101e2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102069190610828565b6001600160a01b03161461021957600080fd5b6102a4565b60005460405163bd27b24160e01b815233916001600160a01b03169063bd27b2419061025090869086906004016106dd565b602060405180830381865afa15801561026d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102919190610828565b6001600160a01b0316146102a457600080fd5b8181600186866040516102b8929190610735565b9081526040519081900360200190206102d2929091610431565b5050505050565b8051602081830181018051600182529282019190930120915280546102fd90610745565b80601f016020809104026020016040519081016040528092919081815260200182805461032990610745565b80156103765780601f1061034b57610100808354040283529160200191610376565b820191906000526020600020905b81548152906001019060200180831161035957829003601f168201915b505050505081565b606060018383604051610392929190610735565b908152602001604051809103902080546103ab90610745565b80601f01602080910402602001604051908101604052809291908181526020018280546103d790610745565b80156104245780601f106103f957610100808354040283529160200191610424565b820191906000526020600020905b81548152906001019060200180831161040757829003601f168201915b5050505050905092915050565b82805461043d90610745565b90600052602060002090601f01602090048101928261045f57600085556104a5565b82601f106104785782800160ff198235161785556104a5565b828001600101855582156104a5579182015b828111156104a557823582559160200191906001019061048a565b506104b19291506104b5565b5090565b5b808211156104b157600081556001016104b6565b60008083601f8401126104dc57600080fd5b50813567ffffffffffffffff8111156104f457600080fd5b60208301915083602082850101111561050c57600080fd5b9250929050565b6000806000806040858703121561052957600080fd5b843567ffffffffffffffff8082111561054157600080fd5b61054d888389016104ca565b9096509450602087013591508082111561056657600080fd5b50610573878288016104ca565b95989497509550505050565b634e487b7160e01b600052604160045260246000fd5b6000602082840312156105a757600080fd5b813567ffffffffffffffff808211156105bf57600080fd5b818401915084601f8301126105d357600080fd5b8135818111156105e5576105e561057f565b604051601f8201601f19908116603f0116810190838211818310171561060d5761060d61057f565b8160405282815287602084870101111561062657600080fd5b826020860160208301376000928101602001929092525095945050505050565b600060208083528351808285015260005b8181101561067357858101830151858201604001528201610657565b81811115610685576000604083870101525b50601f01601f1916929092016040019392505050565b600080602083850312156106ae57600080fd5b823567ffffffffffffffff8111156106c557600080fd5b6106d1858286016104ca565b90969095509350505050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b60006020828403121561071e57600080fd5b8151801515811461072e57600080fd5b9392505050565b8183823760009101908152919050565b600181811c9082168061075957607f821691505b6020821081141561077a57634e487b7160e01b600052602260045260246000fd5b50919050565b600060208083526000845481600182811c9150808316806107a257607f831692505b8583108114156107c057634e487b7160e01b85526022600452602485fd5b8786018381526020018180156107dd57600181146107ee57610819565b60ff19861682528782019650610819565b60008b81526020902060005b86811015610813578154848201529085019089016107fa565b83019750505b50949998505050505050505050565b60006020828403121561083a57600080fd5b81516001600160a01b038116811461072e57600080fdfea26469706673582212207086de6139cf01d0bbbdcca5dc971a7d52dfd053fc46973a6f7086b9082d79da64736f6c634300080b0033",
}

// NameRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use NameRegistryMetaData.ABI instead.
var NameRegistryABI = NameRegistryMetaData.ABI

// NameRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NameRegistryMetaData.Bin instead.
var NameRegistryBin = NameRegistryMetaData.Bin

// DeployNameRegistry deploys a new Ethereum contract, binding an instance of NameRegistry to it.
func DeployNameRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, keyRegistryAddress common.Address) (common.Address, *types.Transaction, *NameRegistry, error) {
	parsed, err := NameRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NameRegistryBin), backend, keyRegistryAddress)
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

// GetKey is a free data retrieval call binding the contract method 0xd37aec92.
//
// Solidity: function getKey(string name) view returns(bytes)
func (_NameRegistry *NameRegistryCaller) GetKey(opts *bind.CallOpts, name string) ([]byte, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "getKey", name)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetKey is a free data retrieval call binding the contract method 0xd37aec92.
//
// Solidity: function getKey(string name) view returns(bytes)
func (_NameRegistry *NameRegistrySession) GetKey(name string) ([]byte, error) {
	return _NameRegistry.Contract.GetKey(&_NameRegistry.CallOpts, name)
}

// GetKey is a free data retrieval call binding the contract method 0xd37aec92.
//
// Solidity: function getKey(string name) view returns(bytes)
func (_NameRegistry *NameRegistryCallerSession) GetKey(name string) ([]byte, error) {
	return _NameRegistry.Contract.GetKey(&_NameRegistry.CallOpts, name)
}

// Registry is a free data retrieval call binding the contract method 0x92a296c9.
//
// Solidity: function registry(string ) view returns(bytes)
func (_NameRegistry *NameRegistryCaller) Registry(opts *bind.CallOpts, arg0 string) ([]byte, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "registry", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x92a296c9.
//
// Solidity: function registry(string ) view returns(bytes)
func (_NameRegistry *NameRegistrySession) Registry(arg0 string) ([]byte, error) {
	return _NameRegistry.Contract.Registry(&_NameRegistry.CallOpts, arg0)
}

// Registry is a free data retrieval call binding the contract method 0x92a296c9.
//
// Solidity: function registry(string ) view returns(bytes)
func (_NameRegistry *NameRegistryCallerSession) Registry(arg0 string) ([]byte, error) {
	return _NameRegistry.Contract.Registry(&_NameRegistry.CallOpts, arg0)
}

// Register is a paid mutator transaction binding the contract method 0x50db6b40.
//
// Solidity: function register(string name, bytes key) returns()
func (_NameRegistry *NameRegistryTransactor) Register(opts *bind.TransactOpts, name string, key []byte) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "register", name, key)
}

// Register is a paid mutator transaction binding the contract method 0x50db6b40.
//
// Solidity: function register(string name, bytes key) returns()
func (_NameRegistry *NameRegistrySession) Register(name string, key []byte) (*types.Transaction, error) {
	return _NameRegistry.Contract.Register(&_NameRegistry.TransactOpts, name, key)
}

// Register is a paid mutator transaction binding the contract method 0x50db6b40.
//
// Solidity: function register(string name, bytes key) returns()
func (_NameRegistry *NameRegistryTransactorSession) Register(name string, key []byte) (*types.Transaction, error) {
	return _NameRegistry.Contract.Register(&_NameRegistry.TransactOpts, name, key)
}
