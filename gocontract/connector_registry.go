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

// ConnectorRegistryMetaData contains all meta data concerning the ConnectorRegistry contract.
var ConnectorRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keyRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nameRegistryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"ConnectorRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"protocol\",\"type\":\"string\"}],\"name\":\"getLocation\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"protocol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161087138038061087183398101604081905261002f9161007c565b600080546001600160a01b039384166001600160a01b031991821617909155600180549290931691161790556100af565b80516001600160a01b038116811461007757600080fd5b919050565b6000806040838503121561008f57600080fd5b61009883610060565b91506100a660208401610060565b90509250929050565b6107b3806100be6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80634cd08d031461003b578063905808b114610050575b600080fd5b61004e610049366004610339565b610079565b005b61006361005e3660046103d3565b61021d565b604051610070919061048f565b60405180910390f35b6001546040516369bd764960e11b81526000916001600160a01b03169063d37aec92906100ac908a908a906004016104d2565b600060405180830381865afa1580156100c9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526100f19190810190610504565b9050805160211461010157600080fd5b60005460405163bd27b24160e01b815233916001600160a01b03169063bd27b2419061013190859060040161048f565b602060405180830381865afa15801561014e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061017291906105b1565b6001600160a01b03161461018557600080fd5b8282600289896040516101999291906105da565b908152602001604051809103902087876040516101b79291906105da565b908152602001604051809103902091826101d2929190610673565b507f6bcbdc9c6b89c8a6b289cc4308c6f6841a0e25d81c5fe34b1dedb2848c5f5ca387878787878760405161020c96959493929190610734565b60405180910390a150505050505050565b6060600285856040516102319291906105da565b9081526020016040518091039020838360405161024f9291906105da565b90815260200160405180910390208054610268906105ea565b80601f0160208091040260200160405190810160405280929190818152602001828054610294906105ea565b80156102e15780601f106102b6576101008083540402835291602001916102e1565b820191906000526020600020905b8154815290600101906020018083116102c457829003601f168201915b50505050509050949350505050565b60008083601f84011261030257600080fd5b50813567ffffffffffffffff81111561031a57600080fd5b60208301915083602082850101111561033257600080fd5b9250929050565b6000806000806000806060878903121561035257600080fd5b863567ffffffffffffffff8082111561036a57600080fd5b6103768a838b016102f0565b9098509650602089013591508082111561038f57600080fd5b61039b8a838b016102f0565b909650945060408901359150808211156103b457600080fd5b506103c189828a016102f0565b979a9699509497509295939492505050565b600080600080604085870312156103e957600080fd5b843567ffffffffffffffff8082111561040157600080fd5b61040d888389016102f0565b9096509450602087013591508082111561042657600080fd5b50610433878288016102f0565b95989497509550505050565b60005b8381101561045a578181015183820152602001610442565b50506000910152565b6000815180845261047b81602086016020860161043f565b601f01601f19169290920160200192915050565b6020815260006104a26020830184610463565b9392505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6020815260006104e66020830184866104a9565b949350505050565b634e487b7160e01b600052604160045260246000fd5b60006020828403121561051657600080fd5b815167ffffffffffffffff8082111561052e57600080fd5b818401915084601f83011261054257600080fd5b815181811115610554576105546104ee565b604051601f8201601f19908116603f0116810190838211818310171561057c5761057c6104ee565b8160405282815287602084870101111561059557600080fd5b6105a683602083016020880161043f565b979650505050505050565b6000602082840312156105c357600080fd5b81516001600160a01b03811681146104a257600080fd5b8183823760009101908152919050565b600181811c908216806105fe57607f821691505b60208210810361061e57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561066e57600081815260208120601f850160051c8101602086101561064b5750805b601f850160051c820191505b8181101561066a57828155600101610657565b5050505b505050565b67ffffffffffffffff83111561068b5761068b6104ee565b61069f8361069983546105ea565b83610624565b6000601f8411600181146106d357600085156106bb5750838201355b600019600387901b1c1916600186901b17835561072d565b600083815260209020601f19861690835b8281101561070457868501358255602094850194600190920191016106e4565b50868210156107215760001960f88860031b161c19848701351681555b505060018560011b0183555b5050505050565b60608152600061074860608301888a6104a9565b828103602084015261075b8187896104a9565b905082810360408401526107708185876104a9565b999850505050505050505056fea2646970667358221220318fbac10b9e099b30259297636e9aed91cc32ee6ab9ef183a15c783425226af64736f6c63430008110033",
}

// ConnectorRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ConnectorRegistryMetaData.ABI instead.
var ConnectorRegistryABI = ConnectorRegistryMetaData.ABI

// ConnectorRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ConnectorRegistryMetaData.Bin instead.
var ConnectorRegistryBin = ConnectorRegistryMetaData.Bin

// DeployConnectorRegistry deploys a new Ethereum contract, binding an instance of ConnectorRegistry to it.
func DeployConnectorRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, keyRegistryAddress common.Address, nameRegistryAddress common.Address) (common.Address, *types.Transaction, *ConnectorRegistry, error) {
	parsed, err := ConnectorRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ConnectorRegistryBin), backend, keyRegistryAddress, nameRegistryAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ConnectorRegistry{ConnectorRegistryCaller: ConnectorRegistryCaller{contract: contract}, ConnectorRegistryTransactor: ConnectorRegistryTransactor{contract: contract}, ConnectorRegistryFilterer: ConnectorRegistryFilterer{contract: contract}}, nil
}

// ConnectorRegistry is an auto generated Go binding around an Ethereum contract.
type ConnectorRegistry struct {
	ConnectorRegistryCaller     // Read-only binding to the contract
	ConnectorRegistryTransactor // Write-only binding to the contract
	ConnectorRegistryFilterer   // Log filterer for contract events
}

// ConnectorRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConnectorRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConnectorRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConnectorRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConnectorRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConnectorRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConnectorRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConnectorRegistrySession struct {
	Contract     *ConnectorRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ConnectorRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConnectorRegistryCallerSession struct {
	Contract *ConnectorRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ConnectorRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConnectorRegistryTransactorSession struct {
	Contract     *ConnectorRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ConnectorRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConnectorRegistryRaw struct {
	Contract *ConnectorRegistry // Generic contract binding to access the raw methods on
}

// ConnectorRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConnectorRegistryCallerRaw struct {
	Contract *ConnectorRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ConnectorRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConnectorRegistryTransactorRaw struct {
	Contract *ConnectorRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConnectorRegistry creates a new instance of ConnectorRegistry, bound to a specific deployed contract.
func NewConnectorRegistry(address common.Address, backend bind.ContractBackend) (*ConnectorRegistry, error) {
	contract, err := bindConnectorRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConnectorRegistry{ConnectorRegistryCaller: ConnectorRegistryCaller{contract: contract}, ConnectorRegistryTransactor: ConnectorRegistryTransactor{contract: contract}, ConnectorRegistryFilterer: ConnectorRegistryFilterer{contract: contract}}, nil
}

// NewConnectorRegistryCaller creates a new read-only instance of ConnectorRegistry, bound to a specific deployed contract.
func NewConnectorRegistryCaller(address common.Address, caller bind.ContractCaller) (*ConnectorRegistryCaller, error) {
	contract, err := bindConnectorRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConnectorRegistryCaller{contract: contract}, nil
}

// NewConnectorRegistryTransactor creates a new write-only instance of ConnectorRegistry, bound to a specific deployed contract.
func NewConnectorRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ConnectorRegistryTransactor, error) {
	contract, err := bindConnectorRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConnectorRegistryTransactor{contract: contract}, nil
}

// NewConnectorRegistryFilterer creates a new log filterer instance of ConnectorRegistry, bound to a specific deployed contract.
func NewConnectorRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ConnectorRegistryFilterer, error) {
	contract, err := bindConnectorRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConnectorRegistryFilterer{contract: contract}, nil
}

// bindConnectorRegistry binds a generic wrapper to an already deployed contract.
func bindConnectorRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConnectorRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConnectorRegistry *ConnectorRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConnectorRegistry.Contract.ConnectorRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConnectorRegistry *ConnectorRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConnectorRegistry.Contract.ConnectorRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConnectorRegistry *ConnectorRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConnectorRegistry.Contract.ConnectorRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConnectorRegistry *ConnectorRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConnectorRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConnectorRegistry *ConnectorRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConnectorRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConnectorRegistry *ConnectorRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConnectorRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetLocation is a free data retrieval call binding the contract method 0x905808b1.
//
// Solidity: function getLocation(string name, string protocol) view returns(string)
func (_ConnectorRegistry *ConnectorRegistryCaller) GetLocation(opts *bind.CallOpts, name string, protocol string) (string, error) {
	var out []interface{}
	err := _ConnectorRegistry.contract.Call(opts, &out, "getLocation", name, protocol)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetLocation is a free data retrieval call binding the contract method 0x905808b1.
//
// Solidity: function getLocation(string name, string protocol) view returns(string)
func (_ConnectorRegistry *ConnectorRegistrySession) GetLocation(name string, protocol string) (string, error) {
	return _ConnectorRegistry.Contract.GetLocation(&_ConnectorRegistry.CallOpts, name, protocol)
}

// GetLocation is a free data retrieval call binding the contract method 0x905808b1.
//
// Solidity: function getLocation(string name, string protocol) view returns(string)
func (_ConnectorRegistry *ConnectorRegistryCallerSession) GetLocation(name string, protocol string) (string, error) {
	return _ConnectorRegistry.Contract.GetLocation(&_ConnectorRegistry.CallOpts, name, protocol)
}

// Register is a paid mutator transaction binding the contract method 0x4cd08d03.
//
// Solidity: function register(string name, string protocol, string location) returns()
func (_ConnectorRegistry *ConnectorRegistryTransactor) Register(opts *bind.TransactOpts, name string, protocol string, location string) (*types.Transaction, error) {
	return _ConnectorRegistry.contract.Transact(opts, "register", name, protocol, location)
}

// Register is a paid mutator transaction binding the contract method 0x4cd08d03.
//
// Solidity: function register(string name, string protocol, string location) returns()
func (_ConnectorRegistry *ConnectorRegistrySession) Register(name string, protocol string, location string) (*types.Transaction, error) {
	return _ConnectorRegistry.Contract.Register(&_ConnectorRegistry.TransactOpts, name, protocol, location)
}

// Register is a paid mutator transaction binding the contract method 0x4cd08d03.
//
// Solidity: function register(string name, string protocol, string location) returns()
func (_ConnectorRegistry *ConnectorRegistryTransactorSession) Register(name string, protocol string, location string) (*types.Transaction, error) {
	return _ConnectorRegistry.Contract.Register(&_ConnectorRegistry.TransactOpts, name, protocol, location)
}

// ConnectorRegistryConnectorRegisteredIterator is returned from FilterConnectorRegistered and is used to iterate over the raw logs and unpacked data for ConnectorRegistered events raised by the ConnectorRegistry contract.
type ConnectorRegistryConnectorRegisteredIterator struct {
	Event *ConnectorRegistryConnectorRegistered // Event containing the contract specifics and raw log

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
func (it *ConnectorRegistryConnectorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectorRegistryConnectorRegistered)
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
		it.Event = new(ConnectorRegistryConnectorRegistered)
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
func (it *ConnectorRegistryConnectorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConnectorRegistryConnectorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConnectorRegistryConnectorRegistered represents a ConnectorRegistered event raised by the ConnectorRegistry contract.
type ConnectorRegistryConnectorRegistered struct {
	Arg0 string
	Arg1 string
	Arg2 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterConnectorRegistered is a free log retrieval operation binding the contract event 0x6bcbdc9c6b89c8a6b289cc4308c6f6841a0e25d81c5fe34b1dedb2848c5f5ca3.
//
// Solidity: event ConnectorRegistered(string arg0, string arg1, string arg2)
func (_ConnectorRegistry *ConnectorRegistryFilterer) FilterConnectorRegistered(opts *bind.FilterOpts) (*ConnectorRegistryConnectorRegisteredIterator, error) {

	logs, sub, err := _ConnectorRegistry.contract.FilterLogs(opts, "ConnectorRegistered")
	if err != nil {
		return nil, err
	}
	return &ConnectorRegistryConnectorRegisteredIterator{contract: _ConnectorRegistry.contract, event: "ConnectorRegistered", logs: logs, sub: sub}, nil
}

// WatchConnectorRegistered is a free log subscription operation binding the contract event 0x6bcbdc9c6b89c8a6b289cc4308c6f6841a0e25d81c5fe34b1dedb2848c5f5ca3.
//
// Solidity: event ConnectorRegistered(string arg0, string arg1, string arg2)
func (_ConnectorRegistry *ConnectorRegistryFilterer) WatchConnectorRegistered(opts *bind.WatchOpts, sink chan<- *ConnectorRegistryConnectorRegistered) (event.Subscription, error) {

	logs, sub, err := _ConnectorRegistry.contract.WatchLogs(opts, "ConnectorRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConnectorRegistryConnectorRegistered)
				if err := _ConnectorRegistry.contract.UnpackLog(event, "ConnectorRegistered", log); err != nil {
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
// Solidity: event ConnectorRegistered(string arg0, string arg1, string arg2)
func (_ConnectorRegistry *ConnectorRegistryFilterer) ParseConnectorRegistered(log types.Log) (*ConnectorRegistryConnectorRegistered, error) {
	event := new(ConnectorRegistryConnectorRegistered)
	if err := _ConnectorRegistry.contract.UnpackLog(event, "ConnectorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
