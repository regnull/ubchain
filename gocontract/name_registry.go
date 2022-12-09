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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keyRegistryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"NameRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161087e38038061087e83398101604081905261002f91610054565b600080546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b6107eb806100936000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806350db6b401461003b578063d37aec9214610050575b600080fd5b61004e6100493660046103f5565b610079565b005b61006361005e366004610461565b6102f9565b60405161007091906104a3565b60405180910390f35b6021811461008657600080fd5b600383101561009457600080fd5b60408311156100a257600080fd5b600054604051632d654a9760e11b81526001600160a01b0390911690635aca952e906100d4908590859060040161051a565b602060405180830381865afa1580156100f1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101159190610536565b61011e57600080fd5b60006001858560405161013292919061055f565b9081526020016040518091039020805461014b9061056f565b905011156102005760005460405133916001600160a01b03169063bd27b2419060019061017b908990899061055f565b9081526040519081900360200181206001600160e01b031960e084901b1682526101a7916004016105a9565b602060405180830381865afa1580156101c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101e89190610634565b6001600160a01b0316146101fb57600080fd5b610286565b60005460405163bd27b24160e01b815233916001600160a01b03169063bd27b24190610232908690869060040161051a565b602060405180830381865afa15801561024f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102739190610634565b6001600160a01b03161461028657600080fd5b81816001868660405161029a92919061055f565b908152602001604051809103902091826102b59291906106c2565b507f708bce0ff0fcffae4a3aaa47b6b85d7b433417e5acc7ff1ee5d4937ac585af71848484846040516102eb9493929190610783565b60405180910390a150505050565b60606001838360405161030d92919061055f565b908152602001604051809103902080546103269061056f565b80601f01602080910402602001604051908101604052809291908181526020018280546103529061056f565b801561039f5780601f106103745761010080835404028352916020019161039f565b820191906000526020600020905b81548152906001019060200180831161038257829003601f168201915b5050505050905092915050565b60008083601f8401126103be57600080fd5b50813567ffffffffffffffff8111156103d657600080fd5b6020830191508360208285010111156103ee57600080fd5b9250929050565b6000806000806040858703121561040b57600080fd5b843567ffffffffffffffff8082111561042357600080fd5b61042f888389016103ac565b9096509450602087013591508082111561044857600080fd5b50610455878288016103ac565b95989497509550505050565b6000806020838503121561047457600080fd5b823567ffffffffffffffff81111561048b57600080fd5b610497858286016103ac565b90969095509350505050565b600060208083528351808285015260005b818110156104d0578581018301518582016040015282016104b4565b506000604082860101526040601f19601f8301168501019250505092915050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60208152600061052e6020830184866104f1565b949350505050565b60006020828403121561054857600080fd5b8151801515811461055857600080fd5b9392505050565b8183823760009101908152919050565b600181811c9082168061058357607f821691505b6020821081036105a357634e487b7160e01b600052602260045260246000fd5b50919050565b60006020808352600084546105bd8161056f565b808487015260406001808416600081146105de57600181146105f857610626565b60ff1985168984015283151560051b890183019550610626565b896000528660002060005b8581101561061e5781548b8201860152908301908801610603565b8a0184019650505b509398975050505050505050565b60006020828403121561064657600080fd5b81516001600160a01b038116811461055857600080fd5b634e487b7160e01b600052604160045260246000fd5b601f8211156106bd57600081815260208120601f850160051c8101602086101561069a5750805b601f850160051c820191505b818110156106b9578281556001016106a6565b5050505b505050565b67ffffffffffffffff8311156106da576106da61065d565b6106ee836106e8835461056f565b83610673565b6000601f841160018114610722576000851561070a5750838201355b600019600387901b1c1916600186901b17835561077c565b600083815260209020601f19861690835b828110156107535786850135825560209485019460019092019101610733565b50868210156107705760001960f88860031b161c19848701351681555b505060018560011b0183555b5050505050565b6040815260006107976040830186886104f1565b82810360208401526107aa8185876104f1565b97965050505050505056fea264697066735822122081402dd2224887685cb9d71117db14c6146305bca5e6a0bad2df164d1724f22664736f6c63430008110033",
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
