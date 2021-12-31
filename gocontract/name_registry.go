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
	Bin: "0x608060405234801561001057600080fd5b5060405161089d38038061089d83398101604081905261002f91610054565b600080546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b61080a806100936000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806350db6b401461004657806392a296c91461005b578063d37aec9214610084575b600080fd5b6100596100543660046104ba565b610097565b005b61006e610069366004610595565b610280565b60405161007b9190610645565b60405180910390f35b61006e610092366004610678565b610325565b602181146100a457600080fd5b60038310156100b257600080fd5b60408311156100c057600080fd5b600054604051633cfe04d160e11b81526001600160a01b03909116906379fc09a2906100f290859085906004016106ba565b602060405180830381865afa15801561010f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061013391906106e9565b61013c57600080fd5b600060018585604051610150929190610712565b9081526020016040518091039020805461016990610722565b9050111561024b57600082828080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052508054604051630782541760e11b815295965090946001600160a01b039091169350630f04a82e92506101de9150879087906004016106ba565b600060405180830381865afa1580156101fb573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610223919081019061075d565b805190915015610231578091505b8151602083012060601c33811461024757600080fd5b5050505b81816001868660405161025f929190610712565b9081526040519081900360200190206102799290916103d8565b5050505050565b8051602081830181018051600182529282019190930120915280546102a490610722565b80601f01602080910402602001604051908101604052809291908181526020018280546102d090610722565b801561031d5780601f106102f25761010080835404028352916020019161031d565b820191906000526020600020905b81548152906001019060200180831161030057829003601f168201915b505050505081565b606060018383604051610339929190610712565b9081526020016040518091039020805461035290610722565b80601f016020809104026020016040519081016040528092919081815260200182805461037e90610722565b80156103cb5780601f106103a0576101008083540402835291602001916103cb565b820191906000526020600020905b8154815290600101906020018083116103ae57829003601f168201915b5050505050905092915050565b8280546103e490610722565b90600052602060002090601f016020900481019282610406576000855561044c565b82601f1061041f5782800160ff1982351617855561044c565b8280016001018555821561044c579182015b8281111561044c578235825591602001919060010190610431565b5061045892915061045c565b5090565b5b80821115610458576000815560010161045d565b60008083601f84011261048357600080fd5b50813567ffffffffffffffff81111561049b57600080fd5b6020830191508360208285010111156104b357600080fd5b9250929050565b600080600080604085870312156104d057600080fd5b843567ffffffffffffffff808211156104e857600080fd5b6104f488838901610471565b9096509450602087013591508082111561050d57600080fd5b5061051a87828801610471565b95989497509550505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561056557610565610526565b604052919050565b600067ffffffffffffffff82111561058757610587610526565b50601f01601f191660200190565b6000602082840312156105a757600080fd5b813567ffffffffffffffff8111156105be57600080fd5b8201601f810184136105cf57600080fd5b80356105e26105dd8261056d565b61053c565b8181528560208385010111156105f757600080fd5b81602084016020830137600091810160200191909152949350505050565b60005b83811015610630578181015183820152602001610618565b8381111561063f576000848401525b50505050565b6020815260008251806020840152610664816040850160208701610615565b601f01601f19169190910160400192915050565b6000806020838503121561068b57600080fd5b823567ffffffffffffffff8111156106a257600080fd5b6106ae85828601610471565b90969095509350505050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b6000602082840312156106fb57600080fd5b8151801515811461070b57600080fd5b9392505050565b8183823760009101908152919050565b600181811c9082168061073657607f821691505b6020821081141561075757634e487b7160e01b600052602260045260246000fd5b50919050565b60006020828403121561076f57600080fd5b815167ffffffffffffffff81111561078657600080fd5b8201601f8101841361079757600080fd5b80516107a56105dd8261056d565b8181528560208385010111156107ba57600080fd5b6107cb826020830160208601610615565b9594505050505056fea2646970667358221220cb5bb28df770d0be8372b72a8914b559615b1bf19962beaee71810c8e5c4038564736f6c634300080b0033",
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
