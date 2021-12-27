// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package keyregistry

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

// KeyregistryMetaData contains all meta data concerning the Keyregistry contract.
var KeyregistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"KeyDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"KeyRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"disable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061046e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806382fbdc9c146100465780638fc925aa1461005b578063a15d581c1461006e575b600080fd5b6100596100543660046102c0565b6100e1565b005b6100596100693660046102c0565b6101d6565b6100b761007c366004610348565b805160208183018101805160008252928201919093012091525460ff808216916101008104909116906201000090046001600160a01b031683565b60408051931515845291151560208401526001600160a01b03169082015260600160405180910390f35b602181146100ee57600080fd5b600082826040516101009291906103f9565b9081526040519081900360200190205460ff161561011d57600080fd5b6001600083836040516101319291906103f9565b908152604051908190036020018120805492151560ff1990931692909217909155339060009061016490859085906103f9565b90815260405190819003602001812080546001600160a01b0393909316620100000262010000600160b01b0319909316929092179091557f2123184afbd7510b176b666deb82298c81b9a8b363e3a88f8cf4b726e4c81ae4906101ca9084908490610409565b60405180910390a15050565b602181146101e357600080fd5b600082826040516101f59291906103f9565b9081526040519081900360200190205460ff1661021157600080fd5b336001600160a01b03166000838360405161022d9291906103f9565b908152604051908190036020019020546001600160a01b0362010000909104161461025757600080fd5b60016000838360405161026b9291906103f9565b90815260405190819003602001812080549215156101000261ff0019909316929092179091557f1eefa765080cc5f57ade4997fa66a727ed30bdef80fe40a6dd5fe0dce8f8abbd906101ca9084908490610409565b600080602083850312156102d357600080fd5b823567ffffffffffffffff808211156102eb57600080fd5b818501915085601f8301126102ff57600080fd5b81358181111561030e57600080fd5b86602082850101111561032057600080fd5b60209290920196919550909350505050565b634e487b7160e01b600052604160045260246000fd5b60006020828403121561035a57600080fd5b813567ffffffffffffffff8082111561037257600080fd5b818401915084601f83011261038657600080fd5b81358181111561039857610398610332565b604051601f8201601f19908116603f011681019083821181831017156103c0576103c0610332565b816040528281528760208487010111156103d957600080fd5b826020860160208301376000928101602001929092525095945050505050565b8183823760009101908152919050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f1916010191905056fea264697066735822122031ebb23ad3d58189d16a990dc66bc36fb59f8197363e2543ca92a37e1edfee6e64736f6c634300080b0033",
}

// KeyregistryABI is the input ABI used to generate the binding from.
// Deprecated: Use KeyregistryMetaData.ABI instead.
var KeyregistryABI = KeyregistryMetaData.ABI

// KeyregistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KeyregistryMetaData.Bin instead.
var KeyregistryBin = KeyregistryMetaData.Bin

// DeployKeyregistry deploys a new Ethereum contract, binding an instance of Keyregistry to it.
func DeployKeyregistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Keyregistry, error) {
	parsed, err := KeyregistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KeyregistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Keyregistry{KeyregistryCaller: KeyregistryCaller{contract: contract}, KeyregistryTransactor: KeyregistryTransactor{contract: contract}, KeyregistryFilterer: KeyregistryFilterer{contract: contract}}, nil
}

// Keyregistry is an auto generated Go binding around an Ethereum contract.
type Keyregistry struct {
	KeyregistryCaller     // Read-only binding to the contract
	KeyregistryTransactor // Write-only binding to the contract
	KeyregistryFilterer   // Log filterer for contract events
}

// KeyregistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type KeyregistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyregistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KeyregistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyregistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KeyregistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyregistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KeyregistrySession struct {
	Contract     *Keyregistry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KeyregistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KeyregistryCallerSession struct {
	Contract *KeyregistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// KeyregistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KeyregistryTransactorSession struct {
	Contract     *KeyregistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// KeyregistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type KeyregistryRaw struct {
	Contract *Keyregistry // Generic contract binding to access the raw methods on
}

// KeyregistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KeyregistryCallerRaw struct {
	Contract *KeyregistryCaller // Generic read-only contract binding to access the raw methods on
}

// KeyregistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KeyregistryTransactorRaw struct {
	Contract *KeyregistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKeyregistry creates a new instance of Keyregistry, bound to a specific deployed contract.
func NewKeyregistry(address common.Address, backend bind.ContractBackend) (*Keyregistry, error) {
	contract, err := bindKeyregistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Keyregistry{KeyregistryCaller: KeyregistryCaller{contract: contract}, KeyregistryTransactor: KeyregistryTransactor{contract: contract}, KeyregistryFilterer: KeyregistryFilterer{contract: contract}}, nil
}

// NewKeyregistryCaller creates a new read-only instance of Keyregistry, bound to a specific deployed contract.
func NewKeyregistryCaller(address common.Address, caller bind.ContractCaller) (*KeyregistryCaller, error) {
	contract, err := bindKeyregistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KeyregistryCaller{contract: contract}, nil
}

// NewKeyregistryTransactor creates a new write-only instance of Keyregistry, bound to a specific deployed contract.
func NewKeyregistryTransactor(address common.Address, transactor bind.ContractTransactor) (*KeyregistryTransactor, error) {
	contract, err := bindKeyregistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KeyregistryTransactor{contract: contract}, nil
}

// NewKeyregistryFilterer creates a new log filterer instance of Keyregistry, bound to a specific deployed contract.
func NewKeyregistryFilterer(address common.Address, filterer bind.ContractFilterer) (*KeyregistryFilterer, error) {
	contract, err := bindKeyregistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KeyregistryFilterer{contract: contract}, nil
}

// bindKeyregistry binds a generic wrapper to an already deployed contract.
func bindKeyregistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KeyregistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Keyregistry *KeyregistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Keyregistry.Contract.KeyregistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Keyregistry *KeyregistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Keyregistry.Contract.KeyregistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Keyregistry *KeyregistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Keyregistry.Contract.KeyregistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Keyregistry *KeyregistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Keyregistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Keyregistry *KeyregistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Keyregistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Keyregistry *KeyregistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Keyregistry.Contract.contract.Transact(opts, method, params...)
}

// Registry is a free data retrieval call binding the contract method 0xa15d581c.
//
// Solidity: function registry(bytes ) view returns(bool initialized, bool disabled, address owner)
func (_Keyregistry *KeyregistryCaller) Registry(opts *bind.CallOpts, arg0 []byte) (struct {
	Initialized bool
	Disabled    bool
	Owner       common.Address
}, error) {
	var out []interface{}
	err := _Keyregistry.contract.Call(opts, &out, "registry", arg0)

	outstruct := new(struct {
		Initialized bool
		Disabled    bool
		Owner       common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Initialized = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Disabled = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Registry is a free data retrieval call binding the contract method 0xa15d581c.
//
// Solidity: function registry(bytes ) view returns(bool initialized, bool disabled, address owner)
func (_Keyregistry *KeyregistrySession) Registry(arg0 []byte) (struct {
	Initialized bool
	Disabled    bool
	Owner       common.Address
}, error) {
	return _Keyregistry.Contract.Registry(&_Keyregistry.CallOpts, arg0)
}

// Registry is a free data retrieval call binding the contract method 0xa15d581c.
//
// Solidity: function registry(bytes ) view returns(bool initialized, bool disabled, address owner)
func (_Keyregistry *KeyregistryCallerSession) Registry(arg0 []byte) (struct {
	Initialized bool
	Disabled    bool
	Owner       common.Address
}, error) {
	return _Keyregistry.Contract.Registry(&_Keyregistry.CallOpts, arg0)
}

// Disable is a paid mutator transaction binding the contract method 0x8fc925aa.
//
// Solidity: function disable(bytes publicKey) returns()
func (_Keyregistry *KeyregistryTransactor) Disable(opts *bind.TransactOpts, publicKey []byte) (*types.Transaction, error) {
	return _Keyregistry.contract.Transact(opts, "disable", publicKey)
}

// Disable is a paid mutator transaction binding the contract method 0x8fc925aa.
//
// Solidity: function disable(bytes publicKey) returns()
func (_Keyregistry *KeyregistrySession) Disable(publicKey []byte) (*types.Transaction, error) {
	return _Keyregistry.Contract.Disable(&_Keyregistry.TransactOpts, publicKey)
}

// Disable is a paid mutator transaction binding the contract method 0x8fc925aa.
//
// Solidity: function disable(bytes publicKey) returns()
func (_Keyregistry *KeyregistryTransactorSession) Disable(publicKey []byte) (*types.Transaction, error) {
	return _Keyregistry.Contract.Disable(&_Keyregistry.TransactOpts, publicKey)
}

// Register is a paid mutator transaction binding the contract method 0x82fbdc9c.
//
// Solidity: function register(bytes publicKey) returns()
func (_Keyregistry *KeyregistryTransactor) Register(opts *bind.TransactOpts, publicKey []byte) (*types.Transaction, error) {
	return _Keyregistry.contract.Transact(opts, "register", publicKey)
}

// Register is a paid mutator transaction binding the contract method 0x82fbdc9c.
//
// Solidity: function register(bytes publicKey) returns()
func (_Keyregistry *KeyregistrySession) Register(publicKey []byte) (*types.Transaction, error) {
	return _Keyregistry.Contract.Register(&_Keyregistry.TransactOpts, publicKey)
}

// Register is a paid mutator transaction binding the contract method 0x82fbdc9c.
//
// Solidity: function register(bytes publicKey) returns()
func (_Keyregistry *KeyregistryTransactorSession) Register(publicKey []byte) (*types.Transaction, error) {
	return _Keyregistry.Contract.Register(&_Keyregistry.TransactOpts, publicKey)
}

// KeyregistryKeyDisabledIterator is returned from FilterKeyDisabled and is used to iterate over the raw logs and unpacked data for KeyDisabled events raised by the Keyregistry contract.
type KeyregistryKeyDisabledIterator struct {
	Event *KeyregistryKeyDisabled // Event containing the contract specifics and raw log

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
func (it *KeyregistryKeyDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyregistryKeyDisabled)
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
		it.Event = new(KeyregistryKeyDisabled)
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
func (it *KeyregistryKeyDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyregistryKeyDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyregistryKeyDisabled represents a KeyDisabled event raised by the Keyregistry contract.
type KeyregistryKeyDisabled struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterKeyDisabled is a free log retrieval operation binding the contract event 0x1eefa765080cc5f57ade4997fa66a727ed30bdef80fe40a6dd5fe0dce8f8abbd.
//
// Solidity: event KeyDisabled(bytes arg0)
func (_Keyregistry *KeyregistryFilterer) FilterKeyDisabled(opts *bind.FilterOpts) (*KeyregistryKeyDisabledIterator, error) {

	logs, sub, err := _Keyregistry.contract.FilterLogs(opts, "KeyDisabled")
	if err != nil {
		return nil, err
	}
	return &KeyregistryKeyDisabledIterator{contract: _Keyregistry.contract, event: "KeyDisabled", logs: logs, sub: sub}, nil
}

// WatchKeyDisabled is a free log subscription operation binding the contract event 0x1eefa765080cc5f57ade4997fa66a727ed30bdef80fe40a6dd5fe0dce8f8abbd.
//
// Solidity: event KeyDisabled(bytes arg0)
func (_Keyregistry *KeyregistryFilterer) WatchKeyDisabled(opts *bind.WatchOpts, sink chan<- *KeyregistryKeyDisabled) (event.Subscription, error) {

	logs, sub, err := _Keyregistry.contract.WatchLogs(opts, "KeyDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyregistryKeyDisabled)
				if err := _Keyregistry.contract.UnpackLog(event, "KeyDisabled", log); err != nil {
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

// ParseKeyDisabled is a log parse operation binding the contract event 0x1eefa765080cc5f57ade4997fa66a727ed30bdef80fe40a6dd5fe0dce8f8abbd.
//
// Solidity: event KeyDisabled(bytes arg0)
func (_Keyregistry *KeyregistryFilterer) ParseKeyDisabled(log types.Log) (*KeyregistryKeyDisabled, error) {
	event := new(KeyregistryKeyDisabled)
	if err := _Keyregistry.contract.UnpackLog(event, "KeyDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyregistryKeyRegisteredIterator is returned from FilterKeyRegistered and is used to iterate over the raw logs and unpacked data for KeyRegistered events raised by the Keyregistry contract.
type KeyregistryKeyRegisteredIterator struct {
	Event *KeyregistryKeyRegistered // Event containing the contract specifics and raw log

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
func (it *KeyregistryKeyRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyregistryKeyRegistered)
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
		it.Event = new(KeyregistryKeyRegistered)
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
func (it *KeyregistryKeyRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyregistryKeyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyregistryKeyRegistered represents a KeyRegistered event raised by the Keyregistry contract.
type KeyregistryKeyRegistered struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterKeyRegistered is a free log retrieval operation binding the contract event 0x2123184afbd7510b176b666deb82298c81b9a8b363e3a88f8cf4b726e4c81ae4.
//
// Solidity: event KeyRegistered(bytes arg0)
func (_Keyregistry *KeyregistryFilterer) FilterKeyRegistered(opts *bind.FilterOpts) (*KeyregistryKeyRegisteredIterator, error) {

	logs, sub, err := _Keyregistry.contract.FilterLogs(opts, "KeyRegistered")
	if err != nil {
		return nil, err
	}
	return &KeyregistryKeyRegisteredIterator{contract: _Keyregistry.contract, event: "KeyRegistered", logs: logs, sub: sub}, nil
}

// WatchKeyRegistered is a free log subscription operation binding the contract event 0x2123184afbd7510b176b666deb82298c81b9a8b363e3a88f8cf4b726e4c81ae4.
//
// Solidity: event KeyRegistered(bytes arg0)
func (_Keyregistry *KeyregistryFilterer) WatchKeyRegistered(opts *bind.WatchOpts, sink chan<- *KeyregistryKeyRegistered) (event.Subscription, error) {

	logs, sub, err := _Keyregistry.contract.WatchLogs(opts, "KeyRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyregistryKeyRegistered)
				if err := _Keyregistry.contract.UnpackLog(event, "KeyRegistered", log); err != nil {
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

// ParseKeyRegistered is a log parse operation binding the contract event 0x2123184afbd7510b176b666deb82298c81b9a8b363e3a88f8cf4b726e4c81ae4.
//
// Solidity: event KeyRegistered(bytes arg0)
func (_Keyregistry *KeyregistryFilterer) ParseKeyRegistered(log types.Log) (*KeyregistryKeyRegistered, error) {
	event := new(KeyregistryKeyRegistered)
	if err := _Keyregistry.contract.UnpackLog(event, "KeyRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
