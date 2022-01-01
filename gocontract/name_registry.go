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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keyRegistryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"NameRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161099f38038061099f83398101604081905261002f91610054565b600080546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b61090c806100936000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806350db6b401461004657806392a296c91461005b578063d37aec9214610084575b600080fd5b610059610054366004610550565b610097565b005b61006e6100693660046105d2565b610316565b60405161007b9190610683565b60405180910390f35b61006e6100923660046106d8565b6103bb565b602181146100a457600080fd5b60038310156100b257600080fd5b60408311156100c057600080fd5b600054604051632d654a9760e11b81526001600160a01b0390911690635aca952e906100f29085908590600401610743565b602060405180830381865afa15801561010f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610133919061075f565b61013c57600080fd5b600060018585604051610150929190610788565b9081526020016040518091039020805461016990610798565b9050111561021e5760005460405133916001600160a01b03169063bd27b241906001906101999089908990610788565b9081526040519081900360200181206001600160e01b031960e084901b1682526101c5916004016107d3565b602060405180830381865afa1580156101e2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610206919061087b565b6001600160a01b03161461021957600080fd5b6102a4565b60005460405163bd27b24160e01b815233916001600160a01b03169063bd27b241906102509086908690600401610743565b602060405180830381865afa15801561026d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610291919061087b565b6001600160a01b0316146102a457600080fd5b8181600186866040516102b8929190610788565b9081526040519081900360200190206102d292909161046e565b507f708bce0ff0fcffae4a3aaa47b6b85d7b433417e5acc7ff1ee5d4937ac585af718484848460405161030894939291906108a4565b60405180910390a150505050565b80516020818301810180516001825292820191909301209152805461033a90610798565b80601f016020809104026020016040519081016040528092919081815260200182805461036690610798565b80156103b35780601f10610388576101008083540402835291602001916103b3565b820191906000526020600020905b81548152906001019060200180831161039657829003601f168201915b505050505081565b6060600183836040516103cf929190610788565b908152602001604051809103902080546103e890610798565b80601f016020809104026020016040519081016040528092919081815260200182805461041490610798565b80156104615780601f1061043657610100808354040283529160200191610461565b820191906000526020600020905b81548152906001019060200180831161044457829003601f168201915b5050505050905092915050565b82805461047a90610798565b90600052602060002090601f01602090048101928261049c57600085556104e2565b82601f106104b55782800160ff198235161785556104e2565b828001600101855582156104e2579182015b828111156104e25782358255916020019190600101906104c7565b506104ee9291506104f2565b5090565b5b808211156104ee57600081556001016104f3565b60008083601f84011261051957600080fd5b50813567ffffffffffffffff81111561053157600080fd5b60208301915083602082850101111561054957600080fd5b9250929050565b6000806000806040858703121561056657600080fd5b843567ffffffffffffffff8082111561057e57600080fd5b61058a88838901610507565b909650945060208701359150808211156105a357600080fd5b506105b087828801610507565b95989497509550505050565b634e487b7160e01b600052604160045260246000fd5b6000602082840312156105e457600080fd5b813567ffffffffffffffff808211156105fc57600080fd5b818401915084601f83011261061057600080fd5b813581811115610622576106226105bc565b604051601f8201601f19908116603f0116810190838211818310171561064a5761064a6105bc565b8160405282815287602084870101111561066357600080fd5b826020860160208301376000928101602001929092525095945050505050565b600060208083528351808285015260005b818110156106b057858101830151858201604001528201610694565b818111156106c2576000604083870101525b50601f01601f1916929092016040019392505050565b600080602083850312156106eb57600080fd5b823567ffffffffffffffff81111561070257600080fd5b61070e85828601610507565b90969095509350505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60208152600061075760208301848661071a565b949350505050565b60006020828403121561077157600080fd5b8151801515811461078157600080fd5b9392505050565b8183823760009101908152919050565b600181811c908216806107ac57607f821691505b602082108114156107cd57634e487b7160e01b600052602260045260246000fd5b50919050565b600060208083526000845481600182811c9150808316806107f557607f831692505b85831081141561081357634e487b7160e01b85526022600452602485fd5b87860183815260200181801561083057600181146108415761086c565b60ff1986168252878201965061086c565b60008b81526020902060005b868110156108665781548482015290850190890161084d565b83019750505b50949998505050505050505050565b60006020828403121561088d57600080fd5b81516001600160a01b038116811461078157600080fd5b6040815260006108b860408301868861071a565b82810360208401526108cb81858761071a565b97965050505050505056fea264697066735822122012a6b53224243678bc8c9d847a99c27713ef1aa3c7ba4dbeae8d54fbe5e11cfb64736f6c634300080b0033",
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
	Arg0 string
	Arg1 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNameRegistered is a free log retrieval operation binding the contract event 0x708bce0ff0fcffae4a3aaa47b6b85d7b433417e5acc7ff1ee5d4937ac585af71.
//
// Solidity: event NameRegistered(string arg0, bytes arg1)
func (_NameRegistry *NameRegistryFilterer) FilterNameRegistered(opts *bind.FilterOpts) (*NameRegistryNameRegisteredIterator, error) {

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "NameRegistered")
	if err != nil {
		return nil, err
	}
	return &NameRegistryNameRegisteredIterator{contract: _NameRegistry.contract, event: "NameRegistered", logs: logs, sub: sub}, nil
}

// WatchNameRegistered is a free log subscription operation binding the contract event 0x708bce0ff0fcffae4a3aaa47b6b85d7b433417e5acc7ff1ee5d4937ac585af71.
//
// Solidity: event NameRegistered(string arg0, bytes arg1)
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

// ParseNameRegistered is a log parse operation binding the contract event 0x708bce0ff0fcffae4a3aaa47b6b85d7b433417e5acc7ff1ee5d4937ac585af71.
//
// Solidity: event NameRegistered(string arg0, bytes arg1)
func (_NameRegistry *NameRegistryFilterer) ParseNameRegistered(log types.Log) (*NameRegistryNameRegistered, error) {
	event := new(NameRegistryNameRegistered)
	if err := _NameRegistry.contract.UnpackLog(event, "NameRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
