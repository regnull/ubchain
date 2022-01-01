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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keyRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nameRegistryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"ConnectorRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"protocol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161086638038061086683398101604081905261002f9161007c565b600080546001600160a01b039384166001600160a01b031991821617909155600180549290931691161790556100af565b80516001600160a01b038116811461007757600080fd5b919050565b6000806040838503121561008f57600080fd5b61009883610060565b91506100a660208401610060565b90509250929050565b6107a8806100be6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80634cd08d031461003b578063b13ac26414610050575b600080fd5b61004e6100493660046103c0565b610079565b005b61006361005e36600461051f565b61021c565b60405161007091906105df565b60405180910390f35b6001546040516369bd764960e11b81526000916001600160a01b03169063d37aec92906100ac908a908a90600401610622565b600060405180830381865afa1580156100c9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526100f1919081019061063e565b9050805160211461010157600080fd5b60005460405163bd27b24160e01b815233916001600160a01b03169063bd27b241906101319085906004016105df565b602060405180830381865afa15801561014e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061017291906106b5565b6001600160a01b03161461018557600080fd5b8282600289896040516101999291906106de565b908152602001604051809103902087876040516101b79291906106de565b9081526040519081900360200190206101d19290916102de565b507f6bcbdc9c6b89c8a6b289cc4308c6f6841a0e25d81c5fe34b1dedb2848c5f5ca387878787878760405161020b969594939291906106ee565b60405180910390a150505050505050565b815160208184018101805160028252928201948201949094209190935281518083018401805192815290840192909301919091209152805461025d90610737565b80601f016020809104026020016040519081016040528092919081815260200182805461028990610737565b80156102d65780601f106102ab576101008083540402835291602001916102d6565b820191906000526020600020905b8154815290600101906020018083116102b957829003601f168201915b505050505081565b8280546102ea90610737565b90600052602060002090601f01602090048101928261030c5760008555610352565b82601f106103255782800160ff19823516178555610352565b82800160010185558215610352579182015b82811115610352578235825591602001919060010190610337565b5061035e929150610362565b5090565b5b8082111561035e5760008155600101610363565b60008083601f84011261038957600080fd5b50813567ffffffffffffffff8111156103a157600080fd5b6020830191508360208285010111156103b957600080fd5b9250929050565b600080600080600080606087890312156103d957600080fd5b863567ffffffffffffffff808211156103f157600080fd5b6103fd8a838b01610377565b9098509650602089013591508082111561041657600080fd5b6104228a838b01610377565b9096509450604089013591508082111561043b57600080fd5b5061044889828a01610377565b979a9699509497509295939492505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff811182821017156104995761049961045a565b604052919050565b600067ffffffffffffffff8211156104bb576104bb61045a565b50601f01601f191660200190565b600082601f8301126104da57600080fd5b81356104ed6104e8826104a1565b610470565b81815284602083860101111561050257600080fd5b816020850160208301376000918101602001919091529392505050565b6000806040838503121561053257600080fd5b823567ffffffffffffffff8082111561054a57600080fd5b610556868387016104c9565b9350602085013591508082111561056c57600080fd5b50610579858286016104c9565b9150509250929050565b60005b8381101561059e578181015183820152602001610586565b838111156105ad576000848401525b50505050565b600081518084526105cb816020860160208601610583565b601f01601f19169290920160200192915050565b6020815260006105f260208301846105b3565b9392505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6020815260006106366020830184866105f9565b949350505050565b60006020828403121561065057600080fd5b815167ffffffffffffffff81111561066757600080fd5b8201601f8101841361067857600080fd5b80516106866104e8826104a1565b81815285602083850101111561069b57600080fd5b6106ac826020830160208601610583565b95945050505050565b6000602082840312156106c757600080fd5b81516001600160a01b03811681146105f257600080fd5b8183823760009101908152919050565b60608152600061070260608301888a6105f9565b82810360208401526107158187896105f9565b9050828103604084015261072a8185876105f9565b9998505050505050505050565b600181811c9082168061074b57607f821691505b6020821081141561076c57634e487b7160e01b600052602260045260246000fd5b5091905056fea26469706673582212206ca98d92d2cd62dd8426c251a1e5962f509d2cfdab53ccc6a6e54a4f452c042a64736f6c634300080b0033",
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

// Registry is a free data retrieval call binding the contract method 0xb13ac264.
//
// Solidity: function registry(string , string ) view returns(string)
func (_ConnectorRegistry *ConnectorRegistryCaller) Registry(opts *bind.CallOpts, arg0 string, arg1 string) (string, error) {
	var out []interface{}
	err := _ConnectorRegistry.contract.Call(opts, &out, "registry", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0xb13ac264.
//
// Solidity: function registry(string , string ) view returns(string)
func (_ConnectorRegistry *ConnectorRegistrySession) Registry(arg0 string, arg1 string) (string, error) {
	return _ConnectorRegistry.Contract.Registry(&_ConnectorRegistry.CallOpts, arg0, arg1)
}

// Registry is a free data retrieval call binding the contract method 0xb13ac264.
//
// Solidity: function registry(string , string ) view returns(string)
func (_ConnectorRegistry *ConnectorRegistryCallerSession) Registry(arg0 string, arg1 string) (string, error) {
	return _ConnectorRegistry.Contract.Registry(&_ConnectorRegistry.CallOpts, arg0, arg1)
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
