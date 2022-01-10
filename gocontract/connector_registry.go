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
	Bin: "0x608060405234801561001057600080fd5b5060405161080638038061080683398101604081905261002f9161007c565b600080546001600160a01b039384166001600160a01b031991821617909155600180549290931691161790556100af565b80516001600160a01b038116811461007757600080fd5b919050565b6000806040838503121561008f57600080fd5b61009883610060565b91506100a660208401610060565b90509250929050565b610748806100be6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80634cd08d031461003b578063905808b114610050575b600080fd5b61004e6100493660046103d1565b610079565b005b61006361005e36600461046b565b61021c565b6040516100709190610533565b60405180910390f35b6001546040516369bd764960e11b81526000916001600160a01b03169063d37aec92906100ac908a908a90600401610576565b600060405180830381865afa1580156100c9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526100f191908101906105a8565b9050805160211461010157600080fd5b60005460405163bd27b24160e01b815233916001600160a01b03169063bd27b24190610131908590600401610533565b602060405180830381865afa15801561014e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101729190610655565b6001600160a01b03161461018557600080fd5b82826002898960405161019992919061067e565b908152602001604051809103902087876040516101b792919061067e565b9081526040519081900360200190206101d19290916102ef565b507f6bcbdc9c6b89c8a6b289cc4308c6f6841a0e25d81c5fe34b1dedb2848c5f5ca387878787878760405161020b9695949392919061068e565b60405180910390a150505050505050565b60606002858560405161023092919061067e565b9081526020016040518091039020838360405161024e92919061067e565b90815260200160405180910390208054610267906106d7565b80601f0160208091040260200160405190810160405280929190818152602001828054610293906106d7565b80156102e05780601f106102b5576101008083540402835291602001916102e0565b820191906000526020600020905b8154815290600101906020018083116102c357829003601f168201915b50505050509050949350505050565b8280546102fb906106d7565b90600052602060002090601f01602090048101928261031d5760008555610363565b82601f106103365782800160ff19823516178555610363565b82800160010185558215610363579182015b82811115610363578235825591602001919060010190610348565b5061036f929150610373565b5090565b5b8082111561036f5760008155600101610374565b60008083601f84011261039a57600080fd5b50813567ffffffffffffffff8111156103b257600080fd5b6020830191508360208285010111156103ca57600080fd5b9250929050565b600080600080600080606087890312156103ea57600080fd5b863567ffffffffffffffff8082111561040257600080fd5b61040e8a838b01610388565b9098509650602089013591508082111561042757600080fd5b6104338a838b01610388565b9096509450604089013591508082111561044c57600080fd5b5061045989828a01610388565b979a9699509497509295939492505050565b6000806000806040858703121561048157600080fd5b843567ffffffffffffffff8082111561049957600080fd5b6104a588838901610388565b909650945060208701359150808211156104be57600080fd5b506104cb87828801610388565b95989497509550505050565b60005b838110156104f25781810151838201526020016104da565b83811115610501576000848401525b50505050565b6000815180845261051f8160208601602086016104d7565b601f01601f19169290920160200192915050565b6020815260006105466020830184610507565b9392505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60208152600061058a60208301848661054d565b949350505050565b634e487b7160e01b600052604160045260246000fd5b6000602082840312156105ba57600080fd5b815167ffffffffffffffff808211156105d257600080fd5b818401915084601f8301126105e657600080fd5b8151818111156105f8576105f8610592565b604051601f8201601f19908116603f0116810190838211818310171561062057610620610592565b8160405282815287602084870101111561063957600080fd5b61064a8360208301602088016104d7565b979650505050505050565b60006020828403121561066757600080fd5b81516001600160a01b038116811461054657600080fd5b8183823760009101908152919050565b6060815260006106a260608301888a61054d565b82810360208401526106b581878961054d565b905082810360408401526106ca81858761054d565b9998505050505050505050565b600181811c908216806106eb57607f821691505b6020821081141561070c57634e487b7160e01b600052602260045260246000fd5b5091905056fea2646970667358221220356a556c28e2ad938c233f5cd841eb233efabda3268eba0833fde0939eb1b75e64736f6c634300080b0033",
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
