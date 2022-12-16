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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"configName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"configValue\",\"type\":\"string\"}],\"name\":\"ConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NameOwnershipUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NameRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"PriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"PublicKeyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"Sale\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"buyName\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"configName\",\"type\":\"string\"}],\"name\":\"lookupConfig\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"lookupName\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"registerName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"configName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"configValue\",\"type\":\"string\"}],\"name\":\"updateConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"updateOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"updatePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"updatePublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611228806100206000396000f3fe60806040526004361061007b5760003560e01c806362eaff7d1161004e57806362eaff7d14610118578063639cdb951461012b578063971c72871461015a578063a3ddd6911461017a57600080fd5b8063208b5f38146100805780633aabd19c146100a25780634a432a46146100d857806356c792db146100f8575b600080fd5b34801561008c57600080fd5b506100a061009b366004610b0c565b61019a565b005b3480156100ae57600080fd5b506100c26100bd366004610b0c565b610258565b6040516100cf9190610bbe565b60405180910390f35b3480156100e457600080fd5b506100a06100f3366004610bd8565b61032b565b34801561010457600080fd5b506100a0610113366004610b0c565b6103d5565b6100a0610126366004610c3a565b6106b2565b34801561013757600080fd5b5061014b610146366004610d15565b610862565b6040516100cf93929190610d57565b34801561016657600080fd5b506100a0610175366004610d8b565b610941565b34801561018657600080fd5b506100a0610195366004610e25565b610a0a565b602183146101a757600080fd5b336001600160a01b0316600083836040516101c3929190610e88565b908152604051908190036020019020600101546001600160a01b0316146101e957600080fd5b8383600084846040516101fd929190610e88565b90815260405190819003602001902091610218919083610f21565b507f0ba4a7233a2cfb937aa9644f11f49345b443fe8d5ab3494a1879faa500728d3e828260405161024a92919061100b565b60405180910390a150505050565b60606001858560405161026c929190610e88565b9081526020016040518091039020838360405161028a929190610e88565b908152602001604051809103902080546102a390610e98565b80601f01602080910402602001604051908101604052809291908181526020018280546102cf90610e98565b801561031c5780601f106102f15761010080835404028352916020019161031c565b820191906000526020600020905b8154815290600101906020018083116102ff57829003601f168201915b50505050509050949350505050565b336001600160a01b031660008484604051610347929190610e88565b908152604051908190036020019020600101546001600160a01b03161461036d57600080fd5b8060008484604051610380929190610e88565b9081526020016040518091039020600201819055507f159e83f4712ba2552e68be9d848e49bf6dd35c24f19564ffd523b6549450a2f48383836040516103c893929190611027565b60405180910390a1505050565b602183146103e257600080fd5b60006001600160a01b0316600083836040516103ff929190610e88565b908152604051908190036020019020600101546001600160a01b03161461042557600080fd5b600181101561043357600080fd5b604081111561044157600080fd5b818160008181106104545761045461104b565b909101356001600160f81b031916602d60f81b03905061047357600080fd5b600082828080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250939450600193925050505b82518110156105eb5760008382815181106104cf576104cf61104b565b01602001516001600160f81b0319169050602d60f81b8114806105195750600360fc1b6001600160f81b03198216108015906105195750603960f81b6001600160f81b0319821611155b806105315750605f60f81b6001600160f81b03198216145b806105635750606160f81b6001600160f81b03198216108015906105635750603d60f91b6001600160f81b0319821611155b6105a25760405162461bcd60e51b815260206004820152600c60248201526b696e76616c6964206e616d6560a01b604482015260640160405180910390fd5b602d60f81b6001600160f81b03198216148015906105ce5750605f60f81b6001600160f81b0319821614155b156105d857600092505b50806105e381611061565b9150506104b2565b5080156105f757600080fd5b336000858560405161060a929190610e88565b908152602001604051809103902060010160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550858560008686604051610653929190610e88565b9081526040519081900360200190209161066e919083610f21565b507f1c6eac0e720ec22bb0653aec9c19985633a4fb07971cf973096c2f8e3c37c17f8484336040516106a293929190611088565b60405180910390a1505050505050565b60008084846040516106c5929190610e88565b908152602001604051809103902060020154116106e157600080fd5b34600084846040516106f4929190610e88565b908152602001604051809103902060020154111561071157600080fd5b6000808484604051610724929190610e88565b90815260200160405180910390206002015490506000848460405161074a929190610e88565b908152604051908190036020018120600101546001600160a01b0316903480156108fc02916000818181858888f1935050505015801561078e573d6000803e3d6000fd5b5033600085856040516107a2929190610e88565b908152602001604051809103902060010160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555081600085856040516107ea929190610e88565b9081526040519081900360200190209061080490826110b4565b506000808585604051610818929190610e88565b9081526020016040518091039020600201819055507f74d699b6db148dc632a2dbe2c680980f546b0382af2e71a7d4bf51c3193adb9a8484833360405161024a9493929190611174565b600060606000806000868660405161087b929190610e88565b9081526040519081900360200190206001810154600282015482549293506001600160a01b039091169183919082906108b390610e98565b80601f01602080910402602001604051908101604052809291908181526020018280546108df90610e98565b801561092c5780601f106109015761010080835404028352916020019161092c565b820191906000526020600020905b81548152906001019060200180831161090f57829003601f168201915b50505050509150935093509350509250925092565b336001600160a01b03166000878760405161095d929190610e88565b908152604051908190036020019020600101546001600160a01b03161461098357600080fd5b818160018888604051610997929190610e88565b908152602001604051809103902086866040516109b5929190610e88565b908152602001604051809103902091826109d0929190610f21565b507fcde50e1bbc8495f4d015791042ac8d9b4e45d1cd60159e1fbad5863ee388d8288686868686866040516106a2969594939291906111a9565b336001600160a01b031660008484604051610a26929190610e88565b908152604051908190036020019020600101546001600160a01b031614610a4c57600080fd5b8060008484604051610a5f929190610e88565b90815260405190819003602001812060010180546001600160a01b03939093166001600160a01b0319909316929092179091557f179b5e07786d37fc59a40f1e1e8a697240c6e378cf1e17cf9438b0110abeb6d7906103c890859085908590611088565b60008083601f840112610ad557600080fd5b50813567ffffffffffffffff811115610aed57600080fd5b602083019150836020828501011115610b0557600080fd5b9250929050565b60008060008060408587031215610b2257600080fd5b843567ffffffffffffffff80821115610b3a57600080fd5b610b4688838901610ac3565b90965094506020870135915080821115610b5f57600080fd5b50610b6c87828801610ac3565b95989497509550505050565b6000815180845260005b81811015610b9e57602081850181015186830182015201610b82565b506000602082860101526020601f19601f83011685010191505092915050565b602081526000610bd16020830184610b78565b9392505050565b600080600060408486031215610bed57600080fd5b833567ffffffffffffffff811115610c0457600080fd5b610c1086828701610ac3565b909790965060209590950135949350505050565b634e487b7160e01b600052604160045260246000fd5b600080600060408486031215610c4f57600080fd5b833567ffffffffffffffff80821115610c6757600080fd5b610c7387838801610ac3565b90955093506020860135915080821115610c8c57600080fd5b818601915086601f830112610ca057600080fd5b813581811115610cb257610cb2610c24565b604051601f8201601f19908116603f01168101908382118183101715610cda57610cda610c24565b81604052828152896020848701011115610cf357600080fd5b8260208601602083013760006020848301015280955050505050509250925092565b60008060208385031215610d2857600080fd5b823567ffffffffffffffff811115610d3f57600080fd5b610d4b85828601610ac3565b90969095509350505050565b6001600160a01b0384168152606060208201819052600090610d7b90830185610b78565b9050826040830152949350505050565b60008060008060008060608789031215610da457600080fd5b863567ffffffffffffffff80821115610dbc57600080fd5b610dc88a838b01610ac3565b90985096506020890135915080821115610de157600080fd5b610ded8a838b01610ac3565b90965094506040890135915080821115610e0657600080fd5b50610e1389828a01610ac3565b979a9699509497509295939492505050565b600080600060408486031215610e3a57600080fd5b833567ffffffffffffffff811115610e5157600080fd5b610e5d86828701610ac3565b90945092505060208401356001600160a01b0381168114610e7d57600080fd5b809150509250925092565b8183823760009101908152919050565b600181811c90821680610eac57607f821691505b602082108103610ecc57634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115610f1c57600081815260208120601f850160051c81016020861015610ef95750805b601f850160051c820191505b81811015610f1857828155600101610f05565b5050505b505050565b67ffffffffffffffff831115610f3957610f39610c24565b610f4d83610f478354610e98565b83610ed2565b6000601f841160018114610f815760008515610f695750838201355b600019600387901b1c1916600186901b178355610fdb565b600083815260209020601f19861690835b82811015610fb25786850135825560209485019460019092019101610f92565b5086821015610fcf5760001960f88860031b161c19848701351681555b505060018560011b0183555b5050505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60208152600061101f602083018486610fe2565b949350505050565b60408152600061103b604083018587610fe2565b9050826020830152949350505050565b634e487b7160e01b600052603260045260246000fd5b60006001820161108157634e487b7160e01b600052601160045260246000fd5b5060010190565b60408152600061109c604083018587610fe2565b905060018060a01b0383166020830152949350505050565b815167ffffffffffffffff8111156110ce576110ce610c24565b6110e2816110dc8454610e98565b84610ed2565b602080601f83116001811461111757600084156110ff5750858301515b600019600386901b1c1916600185901b178555610f18565b600085815260208120601f198616915b8281101561114657888601518255948401946001909101908401611127565b50858210156111645787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b606081526000611188606083018688610fe2565b6020830194909452506001600160a01b039190911660409091015292915050565b6060815260006111bd60608301888a610fe2565b82810360208401526111d0818789610fe2565b905082810360408401526111e5818587610fe2565b999850505050505050505056fea2646970667358221220b485a4e0f17ed1f4eff123699c7a3f4e6ab83da846645c999de5a7b57ea306e764736f6c63430008110033",
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

// LookupConfig is a free data retrieval call binding the contract method 0x3aabd19c.
//
// Solidity: function lookupConfig(string name, string configName) view returns(string)
func (_NameRegistry *NameRegistryCaller) LookupConfig(opts *bind.CallOpts, name string, configName string) (string, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "lookupConfig", name, configName)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LookupConfig is a free data retrieval call binding the contract method 0x3aabd19c.
//
// Solidity: function lookupConfig(string name, string configName) view returns(string)
func (_NameRegistry *NameRegistrySession) LookupConfig(name string, configName string) (string, error) {
	return _NameRegistry.Contract.LookupConfig(&_NameRegistry.CallOpts, name, configName)
}

// LookupConfig is a free data retrieval call binding the contract method 0x3aabd19c.
//
// Solidity: function lookupConfig(string name, string configName) view returns(string)
func (_NameRegistry *NameRegistryCallerSession) LookupConfig(name string, configName string) (string, error) {
	return _NameRegistry.Contract.LookupConfig(&_NameRegistry.CallOpts, name, configName)
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

// BuyName is a paid mutator transaction binding the contract method 0x62eaff7d.
//
// Solidity: function buyName(string name, bytes publicKey) payable returns()
func (_NameRegistry *NameRegistryTransactor) BuyName(opts *bind.TransactOpts, name string, publicKey []byte) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "buyName", name, publicKey)
}

// BuyName is a paid mutator transaction binding the contract method 0x62eaff7d.
//
// Solidity: function buyName(string name, bytes publicKey) payable returns()
func (_NameRegistry *NameRegistrySession) BuyName(name string, publicKey []byte) (*types.Transaction, error) {
	return _NameRegistry.Contract.BuyName(&_NameRegistry.TransactOpts, name, publicKey)
}

// BuyName is a paid mutator transaction binding the contract method 0x62eaff7d.
//
// Solidity: function buyName(string name, bytes publicKey) payable returns()
func (_NameRegistry *NameRegistryTransactorSession) BuyName(name string, publicKey []byte) (*types.Transaction, error) {
	return _NameRegistry.Contract.BuyName(&_NameRegistry.TransactOpts, name, publicKey)
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

// UpdateConfig is a paid mutator transaction binding the contract method 0x971c7287.
//
// Solidity: function updateConfig(string name, string configName, string configValue) returns()
func (_NameRegistry *NameRegistryTransactor) UpdateConfig(opts *bind.TransactOpts, name string, configName string, configValue string) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "updateConfig", name, configName, configValue)
}

// UpdateConfig is a paid mutator transaction binding the contract method 0x971c7287.
//
// Solidity: function updateConfig(string name, string configName, string configValue) returns()
func (_NameRegistry *NameRegistrySession) UpdateConfig(name string, configName string, configValue string) (*types.Transaction, error) {
	return _NameRegistry.Contract.UpdateConfig(&_NameRegistry.TransactOpts, name, configName, configValue)
}

// UpdateConfig is a paid mutator transaction binding the contract method 0x971c7287.
//
// Solidity: function updateConfig(string name, string configName, string configValue) returns()
func (_NameRegistry *NameRegistryTransactorSession) UpdateConfig(name string, configName string, configValue string) (*types.Transaction, error) {
	return _NameRegistry.Contract.UpdateConfig(&_NameRegistry.TransactOpts, name, configName, configValue)
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

// NameRegistryConfigUpdatedIterator is returned from FilterConfigUpdated and is used to iterate over the raw logs and unpacked data for ConfigUpdated events raised by the NameRegistry contract.
type NameRegistryConfigUpdatedIterator struct {
	Event *NameRegistryConfigUpdated // Event containing the contract specifics and raw log

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
func (it *NameRegistryConfigUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryConfigUpdated)
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
		it.Event = new(NameRegistryConfigUpdated)
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
func (it *NameRegistryConfigUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryConfigUpdated represents a ConfigUpdated event raised by the NameRegistry contract.
type NameRegistryConfigUpdated struct {
	Name        string
	ConfigName  string
	ConfigValue string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConfigUpdated is a free log retrieval operation binding the contract event 0xcde50e1bbc8495f4d015791042ac8d9b4e45d1cd60159e1fbad5863ee388d828.
//
// Solidity: event ConfigUpdated(string name, string configName, string configValue)
func (_NameRegistry *NameRegistryFilterer) FilterConfigUpdated(opts *bind.FilterOpts) (*NameRegistryConfigUpdatedIterator, error) {

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "ConfigUpdated")
	if err != nil {
		return nil, err
	}
	return &NameRegistryConfigUpdatedIterator{contract: _NameRegistry.contract, event: "ConfigUpdated", logs: logs, sub: sub}, nil
}

// WatchConfigUpdated is a free log subscription operation binding the contract event 0xcde50e1bbc8495f4d015791042ac8d9b4e45d1cd60159e1fbad5863ee388d828.
//
// Solidity: event ConfigUpdated(string name, string configName, string configValue)
func (_NameRegistry *NameRegistryFilterer) WatchConfigUpdated(opts *bind.WatchOpts, sink chan<- *NameRegistryConfigUpdated) (event.Subscription, error) {

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "ConfigUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryConfigUpdated)
				if err := _NameRegistry.contract.UnpackLog(event, "ConfigUpdated", log); err != nil {
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

// ParseConfigUpdated is a log parse operation binding the contract event 0xcde50e1bbc8495f4d015791042ac8d9b4e45d1cd60159e1fbad5863ee388d828.
//
// Solidity: event ConfigUpdated(string name, string configName, string configValue)
func (_NameRegistry *NameRegistryFilterer) ParseConfigUpdated(log types.Log) (*NameRegistryConfigUpdated, error) {
	event := new(NameRegistryConfigUpdated)
	if err := _NameRegistry.contract.UnpackLog(event, "ConfigUpdated", log); err != nil {
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
