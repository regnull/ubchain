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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"protocol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"}],\"name\":\"ConnectorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NameRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"PriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"Sale\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"changePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"protocol\",\"type\":\"string\"}],\"name\":\"lookupConnector\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"lookupName\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"protocol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"}],\"name\":\"registerConnector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"registerName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610d84806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80631fa195961461006757806356c792db1461007c5780635d2097d21461008f578063639cdb95146100a2578063d346e3bd146100cd578063f11a1a9a146100ed575b600080fd5b61007a610075366004610724565b610100565b005b61007a61008a366004610770565b6101aa565b61007a61009d3660046107f2565b6102cf565b6100b56100b03660046108cd565b610450565b6040516100c493929190610955565b60405180910390f35b6100e06100db366004610770565b61052f565b6040516100c49190610989565b61007a6100fb3660046109a3565b610602565b336001600160a01b03166000848460405161011c929190610a3d565b908152604051908190036020019020600101546001600160a01b03161461014257600080fd5b8060008484604051610155929190610a3d565b9081526020016040518091039020600201819055507f29ecd1e0988af1492e43256007437db411881757b3e6e808c9f04847b264178c83838360405161019d93929190610a76565b60405180910390a1505050565b602183146101b757600080fd5b60006001600160a01b0316600083836040516101d4929190610a3d565b908152604051908190036020019020600101546001600160a01b0316146101fa57600080fd5b600181101561020857600080fd5b604081111561021657600080fd5b3360008383604051610229929190610a3d565b908152602001604051809103902060010160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550838360008484604051610272929190610a3d565b9081526040519081900360200190209161028d919083610b23565b507f1c6eac0e720ec22bb0653aec9c19985633a4fb07971cf973096c2f8e3c37c17f8282336040516102c193929190610be4565b60405180910390a150505050565b60008084846040516102e2929190610a3d565b908152602001604051809103902060020154116102fe57600080fd5b6000808484604051610311929190610a3d565b908152602001604051809103902060020154905060008484604051610337929190610a3d565b908152604051908190036020018120600101546001600160a01b03169082156108fc029083906000818181858888f1935050505015801561037c573d6000803e3d6000fd5b503360008585604051610390929190610a3d565b908152602001604051809103902060010160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555081600085856040516103d8929190610a3d565b908152604051908190036020019020906103f29082610c10565b506000808585604051610406929190610a3d565b9081526020016040518091039020600201819055507f74d699b6db148dc632a2dbe2c680980f546b0382af2e71a7d4bf51c3193adb9a848483336040516102c19493929190610cd0565b6000606060008060008686604051610469929190610a3d565b9081526040519081900360200190206001810154600282015482549293506001600160a01b039091169183919082906104a190610a9a565b80601f01602080910402602001604051908101604052809291908181526020018280546104cd90610a9a565b801561051a5780601f106104ef5761010080835404028352916020019161051a565b820191906000526020600020905b8154815290600101906020018083116104fd57829003601f168201915b50505050509150935093509350509250925092565b606060018585604051610543929190610a3d565b90815260200160405180910390208383604051610561929190610a3d565b9081526020016040518091039020805461057a90610a9a565b80601f01602080910402602001604051908101604052809291908181526020018280546105a690610a9a565b80156105f35780601f106105c8576101008083540402835291602001916105f3565b820191906000526020600020905b8154815290600101906020018083116105d657829003601f168201915b50505050509050949350505050565b336001600160a01b03166000878760405161061e929190610a3d565b908152604051908190036020019020600101546001600160a01b03161461064457600080fd5b818160018888604051610658929190610a3d565b90815260200160405180910390208686604051610676929190610a3d565b90815260200160405180910390209182610691929190610b23565b507f6bcbdc9c6b89c8a6b289cc4308c6f6841a0e25d81c5fe34b1dedb2848c5f5ca38686868686866040516106cb96959493929190610d05565b60405180910390a1505050505050565b60008083601f8401126106ed57600080fd5b50813567ffffffffffffffff81111561070557600080fd5b60208301915083602082850101111561071d57600080fd5b9250929050565b60008060006040848603121561073957600080fd5b833567ffffffffffffffff81111561075057600080fd5b61075c868287016106db565b909790965060209590950135949350505050565b6000806000806040858703121561078657600080fd5b843567ffffffffffffffff8082111561079e57600080fd5b6107aa888389016106db565b909650945060208701359150808211156107c357600080fd5b506107d0878288016106db565b95989497509550505050565b634e487b7160e01b600052604160045260246000fd5b60008060006040848603121561080757600080fd5b833567ffffffffffffffff8082111561081f57600080fd5b61082b878388016106db565b9095509350602086013591508082111561084457600080fd5b818601915086601f83011261085857600080fd5b81358181111561086a5761086a6107dc565b604051601f8201601f19908116603f01168101908382118183101715610892576108926107dc565b816040528281528960208487010111156108ab57600080fd5b8260208601602083013760006020848301015280955050505050509250925092565b600080602083850312156108e057600080fd5b823567ffffffffffffffff8111156108f757600080fd5b610903858286016106db565b90969095509350505050565b6000815180845260005b8181101561093557602081850181015186830182015201610919565b506000602082860101526020601f19601f83011685010191505092915050565b6001600160a01b03841681526060602082018190526000906109799083018561090f565b9050826040830152949350505050565b60208152600061099c602083018461090f565b9392505050565b600080600080600080606087890312156109bc57600080fd5b863567ffffffffffffffff808211156109d457600080fd5b6109e08a838b016106db565b909850965060208901359150808211156109f957600080fd5b610a058a838b016106db565b90965094506040890135915080821115610a1e57600080fd5b50610a2b89828a016106db565b979a9699509497509295939492505050565b8183823760009101908152919050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b604081526000610a8a604083018587610a4d565b9050826020830152949350505050565b600181811c90821680610aae57607f821691505b602082108103610ace57634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115610b1e57600081815260208120601f850160051c81016020861015610afb5750805b601f850160051c820191505b81811015610b1a57828155600101610b07565b5050505b505050565b67ffffffffffffffff831115610b3b57610b3b6107dc565b610b4f83610b498354610a9a565b83610ad4565b6000601f841160018114610b835760008515610b6b5750838201355b600019600387901b1c1916600186901b178355610bdd565b600083815260209020601f19861690835b82811015610bb45786850135825560209485019460019092019101610b94565b5086821015610bd15760001960f88860031b161c19848701351681555b505060018560011b0183555b5050505050565b604081526000610bf8604083018587610a4d565b905060018060a01b0383166020830152949350505050565b815167ffffffffffffffff811115610c2a57610c2a6107dc565b610c3e81610c388454610a9a565b84610ad4565b602080601f831160018114610c735760008415610c5b5750858301515b600019600386901b1c1916600185901b178555610b1a565b600085815260208120601f198616915b82811015610ca257888601518255948401946001909101908401610c83565b5085821015610cc05787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b606081526000610ce4606083018688610a4d565b6020830194909452506001600160a01b039190911660409091015292915050565b606081526000610d1960608301888a610a4d565b8281036020840152610d2c818789610a4d565b90508281036040840152610d41818587610a4d565b999850505050505050505056fea26469706673582212203439cf435881b7314a9effb684b966587eb8753d32acffa561770b92538385ce64736f6c63430008110033",
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
// Solidity: function buy(string name, bytes publicKey) returns()
func (_NameRegistry *NameRegistryTransactor) Buy(opts *bind.TransactOpts, name string, publicKey []byte) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "buy", name, publicKey)
}

// Buy is a paid mutator transaction binding the contract method 0x5d2097d2.
//
// Solidity: function buy(string name, bytes publicKey) returns()
func (_NameRegistry *NameRegistrySession) Buy(name string, publicKey []byte) (*types.Transaction, error) {
	return _NameRegistry.Contract.Buy(&_NameRegistry.TransactOpts, name, publicKey)
}

// Buy is a paid mutator transaction binding the contract method 0x5d2097d2.
//
// Solidity: function buy(string name, bytes publicKey) returns()
func (_NameRegistry *NameRegistryTransactorSession) Buy(name string, publicKey []byte) (*types.Transaction, error) {
	return _NameRegistry.Contract.Buy(&_NameRegistry.TransactOpts, name, publicKey)
}

// ChangePrice is a paid mutator transaction binding the contract method 0x1fa19596.
//
// Solidity: function changePrice(string name, uint256 price) returns()
func (_NameRegistry *NameRegistryTransactor) ChangePrice(opts *bind.TransactOpts, name string, price *big.Int) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "changePrice", name, price)
}

// ChangePrice is a paid mutator transaction binding the contract method 0x1fa19596.
//
// Solidity: function changePrice(string name, uint256 price) returns()
func (_NameRegistry *NameRegistrySession) ChangePrice(name string, price *big.Int) (*types.Transaction, error) {
	return _NameRegistry.Contract.ChangePrice(&_NameRegistry.TransactOpts, name, price)
}

// ChangePrice is a paid mutator transaction binding the contract method 0x1fa19596.
//
// Solidity: function changePrice(string name, uint256 price) returns()
func (_NameRegistry *NameRegistryTransactorSession) ChangePrice(name string, price *big.Int) (*types.Transaction, error) {
	return _NameRegistry.Contract.ChangePrice(&_NameRegistry.TransactOpts, name, price)
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

// NameRegistryPriceChangedIterator is returned from FilterPriceChanged and is used to iterate over the raw logs and unpacked data for PriceChanged events raised by the NameRegistry contract.
type NameRegistryPriceChangedIterator struct {
	Event *NameRegistryPriceChanged // Event containing the contract specifics and raw log

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
func (it *NameRegistryPriceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryPriceChanged)
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
		it.Event = new(NameRegistryPriceChanged)
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
func (it *NameRegistryPriceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryPriceChanged represents a PriceChanged event raised by the NameRegistry contract.
type NameRegistryPriceChanged struct {
	Name  string
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPriceChanged is a free log retrieval operation binding the contract event 0x29ecd1e0988af1492e43256007437db411881757b3e6e808c9f04847b264178c.
//
// Solidity: event PriceChanged(string name, uint256 price)
func (_NameRegistry *NameRegistryFilterer) FilterPriceChanged(opts *bind.FilterOpts) (*NameRegistryPriceChangedIterator, error) {

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "PriceChanged")
	if err != nil {
		return nil, err
	}
	return &NameRegistryPriceChangedIterator{contract: _NameRegistry.contract, event: "PriceChanged", logs: logs, sub: sub}, nil
}

// WatchPriceChanged is a free log subscription operation binding the contract event 0x29ecd1e0988af1492e43256007437db411881757b3e6e808c9f04847b264178c.
//
// Solidity: event PriceChanged(string name, uint256 price)
func (_NameRegistry *NameRegistryFilterer) WatchPriceChanged(opts *bind.WatchOpts, sink chan<- *NameRegistryPriceChanged) (event.Subscription, error) {

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "PriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryPriceChanged)
				if err := _NameRegistry.contract.UnpackLog(event, "PriceChanged", log); err != nil {
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

// ParsePriceChanged is a log parse operation binding the contract event 0x29ecd1e0988af1492e43256007437db411881757b3e6e808c9f04847b264178c.
//
// Solidity: event PriceChanged(string name, uint256 price)
func (_NameRegistry *NameRegistryFilterer) ParsePriceChanged(log types.Log) (*NameRegistryPriceChanged, error) {
	event := new(NameRegistryPriceChanged)
	if err := _NameRegistry.contract.UnpackLog(event, "PriceChanged", log); err != nil {
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
