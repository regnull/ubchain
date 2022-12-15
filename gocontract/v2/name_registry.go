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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"protocol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"}],\"name\":\"ConnectorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NameOwnershipUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NameRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"PriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"PublicKeyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"Sale\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"protocol\",\"type\":\"string\"}],\"name\":\"lookupConnector\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"lookupName\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"protocol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"}],\"name\":\"registerConnector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"registerName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"updateOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"updatePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"updatePublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506111ab806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063639cdb951161005b578063639cdb95146100db578063a3ddd69114610106578063d346e3bd14610119578063f11a1a9a1461013957600080fd5b8063208b5f381461008d5780634a432a46146100a257806356c792db146100b55780635d2097d2146100c8575b600080fd5b6100a061009b366004610a8f565b61014c565b005b6100a06100b0366004610afb565b61020a565b6100a06100c3366004610a8f565b6102b4565b6100a06100d6366004610b5d565b610591565b6100ee6100e9366004610c38565b610712565b6040516100fd93929190610cc0565b60405180910390f35b6100a0610114366004610cf4565b6107f1565b61012c610127366004610a8f565b6108aa565b6040516100fd9190610d57565b6100a0610147366004610d71565b61097d565b6021831461015957600080fd5b336001600160a01b031660008383604051610175929190610e0b565b908152604051908190036020019020600101546001600160a01b03161461019b57600080fd5b8383600084846040516101af929190610e0b565b908152604051908190036020019020916101ca919083610ea4565b507f0ba4a7233a2cfb937aa9644f11f49345b443fe8d5ab3494a1879faa500728d3e82826040516101fc929190610f8e565b60405180910390a150505050565b336001600160a01b031660008484604051610226929190610e0b565b908152604051908190036020019020600101546001600160a01b03161461024c57600080fd5b806000848460405161025f929190610e0b565b9081526020016040518091039020600201819055507f159e83f4712ba2552e68be9d848e49bf6dd35c24f19564ffd523b6549450a2f48383836040516102a793929190610faa565b60405180910390a1505050565b602183146102c157600080fd5b60006001600160a01b0316600083836040516102de929190610e0b565b908152604051908190036020019020600101546001600160a01b03161461030457600080fd5b600181101561031257600080fd5b604081111561032057600080fd5b8181600081811061033357610333610fce565b909101356001600160f81b031916602d60f81b03905061035257600080fd5b600082828080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250939450600193925050505b82518110156104ca5760008382815181106103ae576103ae610fce565b01602001516001600160f81b0319169050602d60f81b8114806103f85750600360fc1b6001600160f81b03198216108015906103f85750603960f81b6001600160f81b0319821611155b806104105750605f60f81b6001600160f81b03198216145b806104425750606160f81b6001600160f81b03198216108015906104425750603d60f91b6001600160f81b0319821611155b6104815760405162461bcd60e51b815260206004820152600c60248201526b696e76616c6964206e616d6560a01b604482015260640160405180910390fd5b602d60f81b6001600160f81b03198216148015906104ad5750605f60f81b6001600160f81b0319821614155b156104b757600092505b50806104c281610fe4565b915050610391565b5080156104d657600080fd5b33600085856040516104e9929190610e0b565b908152602001604051809103902060010160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550858560008686604051610532929190610e0b565b9081526040519081900360200190209161054d919083610ea4565b507f1c6eac0e720ec22bb0653aec9c19985633a4fb07971cf973096c2f8e3c37c17f8484336040516105819392919061100b565b60405180910390a1505050505050565b60008084846040516105a4929190610e0b565b908152602001604051809103902060020154116105c057600080fd5b60008084846040516105d3929190610e0b565b9081526020016040518091039020600201549050600084846040516105f9929190610e0b565b908152604051908190036020018120600101546001600160a01b03169082156108fc029083906000818181858888f1935050505015801561063e573d6000803e3d6000fd5b503360008585604051610652929190610e0b565b908152602001604051809103902060010160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550816000858560405161069a929190610e0b565b908152604051908190036020019020906106b49082611037565b5060008085856040516106c8929190610e0b565b9081526020016040518091039020600201819055507f74d699b6db148dc632a2dbe2c680980f546b0382af2e71a7d4bf51c3193adb9a848483336040516101fc94939291906110f7565b600060606000806000868660405161072b929190610e0b565b9081526040519081900360200190206001810154600282015482549293506001600160a01b0390911691839190829061076390610e1b565b80601f016020809104026020016040519081016040528092919081815260200182805461078f90610e1b565b80156107dc5780601f106107b1576101008083540402835291602001916107dc565b820191906000526020600020905b8154815290600101906020018083116107bf57829003601f168201915b50505050509150935093509350509250925092565b336001600160a01b03166000848460405161080d929190610e0b565b908152604051908190036020019020600101546001600160a01b03161461083357600080fd5b8060008484604051610846929190610e0b565b90815260405190819003602001812060010180546001600160a01b03939093166001600160a01b0319909316929092179091557f179b5e07786d37fc59a40f1e1e8a697240c6e378cf1e17cf9438b0110abeb6d7906102a79085908590859061100b565b6060600185856040516108be929190610e0b565b908152602001604051809103902083836040516108dc929190610e0b565b908152602001604051809103902080546108f590610e1b565b80601f016020809104026020016040519081016040528092919081815260200182805461092190610e1b565b801561096e5780601f106109435761010080835404028352916020019161096e565b820191906000526020600020905b81548152906001019060200180831161095157829003601f168201915b50505050509050949350505050565b336001600160a01b031660008787604051610999929190610e0b565b908152604051908190036020019020600101546001600160a01b0316146109bf57600080fd5b8181600188886040516109d3929190610e0b565b908152602001604051809103902086866040516109f1929190610e0b565b90815260200160405180910390209182610a0c929190610ea4565b507f6bcbdc9c6b89c8a6b289cc4308c6f6841a0e25d81c5fe34b1dedb2848c5f5ca38686868686866040516105819695949392919061112c565b60008083601f840112610a5857600080fd5b50813567ffffffffffffffff811115610a7057600080fd5b602083019150836020828501011115610a8857600080fd5b9250929050565b60008060008060408587031215610aa557600080fd5b843567ffffffffffffffff80821115610abd57600080fd5b610ac988838901610a46565b90965094506020870135915080821115610ae257600080fd5b50610aef87828801610a46565b95989497509550505050565b600080600060408486031215610b1057600080fd5b833567ffffffffffffffff811115610b2757600080fd5b610b3386828701610a46565b909790965060209590950135949350505050565b634e487b7160e01b600052604160045260246000fd5b600080600060408486031215610b7257600080fd5b833567ffffffffffffffff80821115610b8a57600080fd5b610b9687838801610a46565b90955093506020860135915080821115610baf57600080fd5b818601915086601f830112610bc357600080fd5b813581811115610bd557610bd5610b47565b604051601f8201601f19908116603f01168101908382118183101715610bfd57610bfd610b47565b81604052828152896020848701011115610c1657600080fd5b8260208601602083013760006020848301015280955050505050509250925092565b60008060208385031215610c4b57600080fd5b823567ffffffffffffffff811115610c6257600080fd5b610c6e85828601610a46565b90969095509350505050565b6000815180845260005b81811015610ca057602081850181015186830182015201610c84565b506000602082860101526020601f19601f83011685010191505092915050565b6001600160a01b0384168152606060208201819052600090610ce490830185610c7a565b9050826040830152949350505050565b600080600060408486031215610d0957600080fd5b833567ffffffffffffffff811115610d2057600080fd5b610d2c86828701610a46565b90945092505060208401356001600160a01b0381168114610d4c57600080fd5b809150509250925092565b602081526000610d6a6020830184610c7a565b9392505050565b60008060008060008060608789031215610d8a57600080fd5b863567ffffffffffffffff80821115610da257600080fd5b610dae8a838b01610a46565b90985096506020890135915080821115610dc757600080fd5b610dd38a838b01610a46565b90965094506040890135915080821115610dec57600080fd5b50610df989828a01610a46565b979a9699509497509295939492505050565b8183823760009101908152919050565b600181811c90821680610e2f57607f821691505b602082108103610e4f57634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115610e9f57600081815260208120601f850160051c81016020861015610e7c5750805b601f850160051c820191505b81811015610e9b57828155600101610e88565b5050505b505050565b67ffffffffffffffff831115610ebc57610ebc610b47565b610ed083610eca8354610e1b565b83610e55565b6000601f841160018114610f045760008515610eec5750838201355b600019600387901b1c1916600186901b178355610f5e565b600083815260209020601f19861690835b82811015610f355786850135825560209485019460019092019101610f15565b5086821015610f525760001960f88860031b161c19848701351681555b505060018560011b0183555b5050505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b602081526000610fa2602083018486610f65565b949350505050565b604081526000610fbe604083018587610f65565b9050826020830152949350505050565b634e487b7160e01b600052603260045260246000fd5b60006001820161100457634e487b7160e01b600052601160045260246000fd5b5060010190565b60408152600061101f604083018587610f65565b905060018060a01b0383166020830152949350505050565b815167ffffffffffffffff81111561105157611051610b47565b6110658161105f8454610e1b565b84610e55565b602080601f83116001811461109a57600084156110825750858301515b600019600386901b1c1916600185901b178555610e9b565b600085815260208120601f198616915b828110156110c9578886015182559484019460019091019084016110aa565b50858210156110e75787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60608152600061110b606083018688610f65565b6020830194909452506001600160a01b039190911660409091015292915050565b60608152600061114060608301888a610f65565b8281036020840152611153818789610f65565b90508281036040840152611168818587610f65565b999850505050505050505056fea2646970667358221220c4edf483df15d840722f97f9c4de5f187c726648152a1f0b2d65a3788c0c2ee864736f6c63430008110033",
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
