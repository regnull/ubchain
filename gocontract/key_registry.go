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

// KeyRegistryMetaData contains all meta data concerning the KeyRegistry contract.
var KeyRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"KeyDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"KeyOwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"KeyRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"disable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"disabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"registered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610701806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638fc925aa1161005b5780638fc925aa146100d2578063a15d581c146100e5578063bd27b24114610142578063e44ccc651461016d57600080fd5b80635aca952e1461008257806382fbdc9c146100aa5780638a6ec070146100bf575b600080fd5b6100956100903660046104de565b610180565b60405190151581526020015b60405180910390f35b6100bd6100b83660046104de565b6101c7565b005b6100bd6100cd366004610520565b61029f565b6100bd6100e03660046104de565b610377565b6101236100f3366004610599565b805160208183018101805160008252928201919093012091525460ff81169061010090046001600160a01b031682565b6040805192151583526001600160a01b039091166020830152016100a1565b6101556101503660046104de565b61042d565b6040516001600160a01b0390911681526020016100a1565b61009561017b3660046104de565b610467565b6000806001600160a01b03166000848460405161019e92919061064a565b908152604051908190036020019020546001600160a01b03610100909104161415905092915050565b602181146101d457600080fd5b60006001600160a01b0316600083836040516101f192919061064a565b908152604051908190036020019020546001600160a01b03610100909104161461021a57600080fd5b336000838360405161022d92919061064a565b90815260405190819003602001812080546001600160a01b039390931661010002610100600160a81b0319909316929092179091557ff77c9b4a9cc46843db8f22d9e84de14b5900316e5706ca63f7fb447ba07c3dc69061029390849084903390610683565b60405180910390a15050565b602182146102ac57600080fd5b336001600160a01b0316600084846040516102c892919061064a565b908152604051908190036020019020546001600160a01b0361010090910416146102f157600080fd5b806000848460405161030492919061064a565b90815260405190819003602001812080546001600160a01b039390931661010002610100600160a81b0319909316929092179091557f4676a911d74639b4a1cbae32a65b0c9ee0ed453681b4246f164bb4381079b03c9061036a90859085908590610683565b60405180910390a1505050565b6021811461038457600080fd5b336001600160a01b0316600083836040516103a092919061064a565b908152604051908190036020019020546001600160a01b0361010090910416146103c957600080fd5b6001600083836040516103dd92919061064a565b908152604051908190036020018120805492151560ff19909316929092179091557f1eefa765080cc5f57ade4997fa66a727ed30bdef80fe40a6dd5fe0dce8f8abbd9061029390849084906106af565b600080838360405161044092919061064a565b908152604051908190036020019020546001600160a01b0361010090910416905092915050565b600080838360405161047a92919061064a565b9081526040519081900360200190205460ff16905092915050565b60008083601f8401126104a757600080fd5b50813567ffffffffffffffff8111156104bf57600080fd5b6020830191508360208285010111156104d757600080fd5b9250929050565b600080602083850312156104f157600080fd5b823567ffffffffffffffff81111561050857600080fd5b61051485828601610495565b90969095509350505050565b60008060006040848603121561053557600080fd5b833567ffffffffffffffff81111561054c57600080fd5b61055886828701610495565b90945092505060208401356001600160a01b038116811461057857600080fd5b809150509250925092565b634e487b7160e01b600052604160045260246000fd5b6000602082840312156105ab57600080fd5b813567ffffffffffffffff808211156105c357600080fd5b818401915084601f8301126105d757600080fd5b8135818111156105e9576105e9610583565b604051601f8201601f19908116603f0116810190838211818310171561061157610611610583565b8160405282815287602084870101111561062a57600080fd5b826020860160208301376000928101602001929092525095945050505050565b8183823760009101908152919050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60408152600061069760408301858761065a565b905060018060a01b0383166020830152949350505050565b6020815260006106c360208301848661065a565b94935050505056fea26469706673582212206e42d958efbfa898bdaf421fa41dcae34f0818d19bd662566035729828e6979c64736f6c634300080b0033",
}

// KeyRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use KeyRegistryMetaData.ABI instead.
var KeyRegistryABI = KeyRegistryMetaData.ABI

// KeyRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KeyRegistryMetaData.Bin instead.
var KeyRegistryBin = KeyRegistryMetaData.Bin

// DeployKeyRegistry deploys a new Ethereum contract, binding an instance of KeyRegistry to it.
func DeployKeyRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KeyRegistry, error) {
	parsed, err := KeyRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KeyRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KeyRegistry{KeyRegistryCaller: KeyRegistryCaller{contract: contract}, KeyRegistryTransactor: KeyRegistryTransactor{contract: contract}, KeyRegistryFilterer: KeyRegistryFilterer{contract: contract}}, nil
}

// KeyRegistry is an auto generated Go binding around an Ethereum contract.
type KeyRegistry struct {
	KeyRegistryCaller     // Read-only binding to the contract
	KeyRegistryTransactor // Write-only binding to the contract
	KeyRegistryFilterer   // Log filterer for contract events
}

// KeyRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type KeyRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KeyRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KeyRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KeyRegistrySession struct {
	Contract     *KeyRegistry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KeyRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KeyRegistryCallerSession struct {
	Contract *KeyRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// KeyRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KeyRegistryTransactorSession struct {
	Contract     *KeyRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// KeyRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type KeyRegistryRaw struct {
	Contract *KeyRegistry // Generic contract binding to access the raw methods on
}

// KeyRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KeyRegistryCallerRaw struct {
	Contract *KeyRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// KeyRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KeyRegistryTransactorRaw struct {
	Contract *KeyRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKeyRegistry creates a new instance of KeyRegistry, bound to a specific deployed contract.
func NewKeyRegistry(address common.Address, backend bind.ContractBackend) (*KeyRegistry, error) {
	contract, err := bindKeyRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KeyRegistry{KeyRegistryCaller: KeyRegistryCaller{contract: contract}, KeyRegistryTransactor: KeyRegistryTransactor{contract: contract}, KeyRegistryFilterer: KeyRegistryFilterer{contract: contract}}, nil
}

// NewKeyRegistryCaller creates a new read-only instance of KeyRegistry, bound to a specific deployed contract.
func NewKeyRegistryCaller(address common.Address, caller bind.ContractCaller) (*KeyRegistryCaller, error) {
	contract, err := bindKeyRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KeyRegistryCaller{contract: contract}, nil
}

// NewKeyRegistryTransactor creates a new write-only instance of KeyRegistry, bound to a specific deployed contract.
func NewKeyRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*KeyRegistryTransactor, error) {
	contract, err := bindKeyRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KeyRegistryTransactor{contract: contract}, nil
}

// NewKeyRegistryFilterer creates a new log filterer instance of KeyRegistry, bound to a specific deployed contract.
func NewKeyRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*KeyRegistryFilterer, error) {
	contract, err := bindKeyRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KeyRegistryFilterer{contract: contract}, nil
}

// bindKeyRegistry binds a generic wrapper to an already deployed contract.
func bindKeyRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KeyRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyRegistry *KeyRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyRegistry.Contract.KeyRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyRegistry *KeyRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyRegistry.Contract.KeyRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyRegistry *KeyRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyRegistry.Contract.KeyRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyRegistry *KeyRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyRegistry *KeyRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyRegistry *KeyRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyRegistry.Contract.contract.Transact(opts, method, params...)
}

// Disabled is a free data retrieval call binding the contract method 0xe44ccc65.
//
// Solidity: function disabled(bytes publicKey) view returns(bool)
func (_KeyRegistry *KeyRegistryCaller) Disabled(opts *bind.CallOpts, publicKey []byte) (bool, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "disabled", publicKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Disabled is a free data retrieval call binding the contract method 0xe44ccc65.
//
// Solidity: function disabled(bytes publicKey) view returns(bool)
func (_KeyRegistry *KeyRegistrySession) Disabled(publicKey []byte) (bool, error) {
	return _KeyRegistry.Contract.Disabled(&_KeyRegistry.CallOpts, publicKey)
}

// Disabled is a free data retrieval call binding the contract method 0xe44ccc65.
//
// Solidity: function disabled(bytes publicKey) view returns(bool)
func (_KeyRegistry *KeyRegistryCallerSession) Disabled(publicKey []byte) (bool, error) {
	return _KeyRegistry.Contract.Disabled(&_KeyRegistry.CallOpts, publicKey)
}

// Owner is a free data retrieval call binding the contract method 0xbd27b241.
//
// Solidity: function owner(bytes publicKey) view returns(address)
func (_KeyRegistry *KeyRegistryCaller) Owner(opts *bind.CallOpts, publicKey []byte) (common.Address, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "owner", publicKey)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xbd27b241.
//
// Solidity: function owner(bytes publicKey) view returns(address)
func (_KeyRegistry *KeyRegistrySession) Owner(publicKey []byte) (common.Address, error) {
	return _KeyRegistry.Contract.Owner(&_KeyRegistry.CallOpts, publicKey)
}

// Owner is a free data retrieval call binding the contract method 0xbd27b241.
//
// Solidity: function owner(bytes publicKey) view returns(address)
func (_KeyRegistry *KeyRegistryCallerSession) Owner(publicKey []byte) (common.Address, error) {
	return _KeyRegistry.Contract.Owner(&_KeyRegistry.CallOpts, publicKey)
}

// Registered is a free data retrieval call binding the contract method 0x5aca952e.
//
// Solidity: function registered(bytes publicKey) view returns(bool)
func (_KeyRegistry *KeyRegistryCaller) Registered(opts *bind.CallOpts, publicKey []byte) (bool, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "registered", publicKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Registered is a free data retrieval call binding the contract method 0x5aca952e.
//
// Solidity: function registered(bytes publicKey) view returns(bool)
func (_KeyRegistry *KeyRegistrySession) Registered(publicKey []byte) (bool, error) {
	return _KeyRegistry.Contract.Registered(&_KeyRegistry.CallOpts, publicKey)
}

// Registered is a free data retrieval call binding the contract method 0x5aca952e.
//
// Solidity: function registered(bytes publicKey) view returns(bool)
func (_KeyRegistry *KeyRegistryCallerSession) Registered(publicKey []byte) (bool, error) {
	return _KeyRegistry.Contract.Registered(&_KeyRegistry.CallOpts, publicKey)
}

// Registry is a free data retrieval call binding the contract method 0xa15d581c.
//
// Solidity: function registry(bytes ) view returns(bool disabled, address owner)
func (_KeyRegistry *KeyRegistryCaller) Registry(opts *bind.CallOpts, arg0 []byte) (struct {
	Disabled bool
	Owner    common.Address
}, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "registry", arg0)

	outstruct := new(struct {
		Disabled bool
		Owner    common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Disabled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Registry is a free data retrieval call binding the contract method 0xa15d581c.
//
// Solidity: function registry(bytes ) view returns(bool disabled, address owner)
func (_KeyRegistry *KeyRegistrySession) Registry(arg0 []byte) (struct {
	Disabled bool
	Owner    common.Address
}, error) {
	return _KeyRegistry.Contract.Registry(&_KeyRegistry.CallOpts, arg0)
}

// Registry is a free data retrieval call binding the contract method 0xa15d581c.
//
// Solidity: function registry(bytes ) view returns(bool disabled, address owner)
func (_KeyRegistry *KeyRegistryCallerSession) Registry(arg0 []byte) (struct {
	Disabled bool
	Owner    common.Address
}, error) {
	return _KeyRegistry.Contract.Registry(&_KeyRegistry.CallOpts, arg0)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0x8a6ec070.
//
// Solidity: function changeOwner(bytes publicKey, address newOwner) returns()
func (_KeyRegistry *KeyRegistryTransactor) ChangeOwner(opts *bind.TransactOpts, publicKey []byte, newOwner common.Address) (*types.Transaction, error) {
	return _KeyRegistry.contract.Transact(opts, "changeOwner", publicKey, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0x8a6ec070.
//
// Solidity: function changeOwner(bytes publicKey, address newOwner) returns()
func (_KeyRegistry *KeyRegistrySession) ChangeOwner(publicKey []byte, newOwner common.Address) (*types.Transaction, error) {
	return _KeyRegistry.Contract.ChangeOwner(&_KeyRegistry.TransactOpts, publicKey, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0x8a6ec070.
//
// Solidity: function changeOwner(bytes publicKey, address newOwner) returns()
func (_KeyRegistry *KeyRegistryTransactorSession) ChangeOwner(publicKey []byte, newOwner common.Address) (*types.Transaction, error) {
	return _KeyRegistry.Contract.ChangeOwner(&_KeyRegistry.TransactOpts, publicKey, newOwner)
}

// Disable is a paid mutator transaction binding the contract method 0x8fc925aa.
//
// Solidity: function disable(bytes publicKey) returns()
func (_KeyRegistry *KeyRegistryTransactor) Disable(opts *bind.TransactOpts, publicKey []byte) (*types.Transaction, error) {
	return _KeyRegistry.contract.Transact(opts, "disable", publicKey)
}

// Disable is a paid mutator transaction binding the contract method 0x8fc925aa.
//
// Solidity: function disable(bytes publicKey) returns()
func (_KeyRegistry *KeyRegistrySession) Disable(publicKey []byte) (*types.Transaction, error) {
	return _KeyRegistry.Contract.Disable(&_KeyRegistry.TransactOpts, publicKey)
}

// Disable is a paid mutator transaction binding the contract method 0x8fc925aa.
//
// Solidity: function disable(bytes publicKey) returns()
func (_KeyRegistry *KeyRegistryTransactorSession) Disable(publicKey []byte) (*types.Transaction, error) {
	return _KeyRegistry.Contract.Disable(&_KeyRegistry.TransactOpts, publicKey)
}

// Register is a paid mutator transaction binding the contract method 0x82fbdc9c.
//
// Solidity: function register(bytes publicKey) returns()
func (_KeyRegistry *KeyRegistryTransactor) Register(opts *bind.TransactOpts, publicKey []byte) (*types.Transaction, error) {
	return _KeyRegistry.contract.Transact(opts, "register", publicKey)
}

// Register is a paid mutator transaction binding the contract method 0x82fbdc9c.
//
// Solidity: function register(bytes publicKey) returns()
func (_KeyRegistry *KeyRegistrySession) Register(publicKey []byte) (*types.Transaction, error) {
	return _KeyRegistry.Contract.Register(&_KeyRegistry.TransactOpts, publicKey)
}

// Register is a paid mutator transaction binding the contract method 0x82fbdc9c.
//
// Solidity: function register(bytes publicKey) returns()
func (_KeyRegistry *KeyRegistryTransactorSession) Register(publicKey []byte) (*types.Transaction, error) {
	return _KeyRegistry.Contract.Register(&_KeyRegistry.TransactOpts, publicKey)
}

// KeyRegistryKeyDisabledIterator is returned from FilterKeyDisabled and is used to iterate over the raw logs and unpacked data for KeyDisabled events raised by the KeyRegistry contract.
type KeyRegistryKeyDisabledIterator struct {
	Event *KeyRegistryKeyDisabled // Event containing the contract specifics and raw log

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
func (it *KeyRegistryKeyDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyRegistryKeyDisabled)
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
		it.Event = new(KeyRegistryKeyDisabled)
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
func (it *KeyRegistryKeyDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyRegistryKeyDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyRegistryKeyDisabled represents a KeyDisabled event raised by the KeyRegistry contract.
type KeyRegistryKeyDisabled struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterKeyDisabled is a free log retrieval operation binding the contract event 0x1eefa765080cc5f57ade4997fa66a727ed30bdef80fe40a6dd5fe0dce8f8abbd.
//
// Solidity: event KeyDisabled(bytes arg0)
func (_KeyRegistry *KeyRegistryFilterer) FilterKeyDisabled(opts *bind.FilterOpts) (*KeyRegistryKeyDisabledIterator, error) {

	logs, sub, err := _KeyRegistry.contract.FilterLogs(opts, "KeyDisabled")
	if err != nil {
		return nil, err
	}
	return &KeyRegistryKeyDisabledIterator{contract: _KeyRegistry.contract, event: "KeyDisabled", logs: logs, sub: sub}, nil
}

// WatchKeyDisabled is a free log subscription operation binding the contract event 0x1eefa765080cc5f57ade4997fa66a727ed30bdef80fe40a6dd5fe0dce8f8abbd.
//
// Solidity: event KeyDisabled(bytes arg0)
func (_KeyRegistry *KeyRegistryFilterer) WatchKeyDisabled(opts *bind.WatchOpts, sink chan<- *KeyRegistryKeyDisabled) (event.Subscription, error) {

	logs, sub, err := _KeyRegistry.contract.WatchLogs(opts, "KeyDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyRegistryKeyDisabled)
				if err := _KeyRegistry.contract.UnpackLog(event, "KeyDisabled", log); err != nil {
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
func (_KeyRegistry *KeyRegistryFilterer) ParseKeyDisabled(log types.Log) (*KeyRegistryKeyDisabled, error) {
	event := new(KeyRegistryKeyDisabled)
	if err := _KeyRegistry.contract.UnpackLog(event, "KeyDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyRegistryKeyOwnerChangedIterator is returned from FilterKeyOwnerChanged and is used to iterate over the raw logs and unpacked data for KeyOwnerChanged events raised by the KeyRegistry contract.
type KeyRegistryKeyOwnerChangedIterator struct {
	Event *KeyRegistryKeyOwnerChanged // Event containing the contract specifics and raw log

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
func (it *KeyRegistryKeyOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyRegistryKeyOwnerChanged)
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
		it.Event = new(KeyRegistryKeyOwnerChanged)
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
func (it *KeyRegistryKeyOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyRegistryKeyOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyRegistryKeyOwnerChanged represents a KeyOwnerChanged event raised by the KeyRegistry contract.
type KeyRegistryKeyOwnerChanged struct {
	Arg0 []byte
	Arg1 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterKeyOwnerChanged is a free log retrieval operation binding the contract event 0x4676a911d74639b4a1cbae32a65b0c9ee0ed453681b4246f164bb4381079b03c.
//
// Solidity: event KeyOwnerChanged(bytes arg0, address arg1)
func (_KeyRegistry *KeyRegistryFilterer) FilterKeyOwnerChanged(opts *bind.FilterOpts) (*KeyRegistryKeyOwnerChangedIterator, error) {

	logs, sub, err := _KeyRegistry.contract.FilterLogs(opts, "KeyOwnerChanged")
	if err != nil {
		return nil, err
	}
	return &KeyRegistryKeyOwnerChangedIterator{contract: _KeyRegistry.contract, event: "KeyOwnerChanged", logs: logs, sub: sub}, nil
}

// WatchKeyOwnerChanged is a free log subscription operation binding the contract event 0x4676a911d74639b4a1cbae32a65b0c9ee0ed453681b4246f164bb4381079b03c.
//
// Solidity: event KeyOwnerChanged(bytes arg0, address arg1)
func (_KeyRegistry *KeyRegistryFilterer) WatchKeyOwnerChanged(opts *bind.WatchOpts, sink chan<- *KeyRegistryKeyOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _KeyRegistry.contract.WatchLogs(opts, "KeyOwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyRegistryKeyOwnerChanged)
				if err := _KeyRegistry.contract.UnpackLog(event, "KeyOwnerChanged", log); err != nil {
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

// ParseKeyOwnerChanged is a log parse operation binding the contract event 0x4676a911d74639b4a1cbae32a65b0c9ee0ed453681b4246f164bb4381079b03c.
//
// Solidity: event KeyOwnerChanged(bytes arg0, address arg1)
func (_KeyRegistry *KeyRegistryFilterer) ParseKeyOwnerChanged(log types.Log) (*KeyRegistryKeyOwnerChanged, error) {
	event := new(KeyRegistryKeyOwnerChanged)
	if err := _KeyRegistry.contract.UnpackLog(event, "KeyOwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyRegistryKeyRegisteredIterator is returned from FilterKeyRegistered and is used to iterate over the raw logs and unpacked data for KeyRegistered events raised by the KeyRegistry contract.
type KeyRegistryKeyRegisteredIterator struct {
	Event *KeyRegistryKeyRegistered // Event containing the contract specifics and raw log

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
func (it *KeyRegistryKeyRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyRegistryKeyRegistered)
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
		it.Event = new(KeyRegistryKeyRegistered)
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
func (it *KeyRegistryKeyRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyRegistryKeyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyRegistryKeyRegistered represents a KeyRegistered event raised by the KeyRegistry contract.
type KeyRegistryKeyRegistered struct {
	Arg0 []byte
	Arg1 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterKeyRegistered is a free log retrieval operation binding the contract event 0xf77c9b4a9cc46843db8f22d9e84de14b5900316e5706ca63f7fb447ba07c3dc6.
//
// Solidity: event KeyRegistered(bytes arg0, address arg1)
func (_KeyRegistry *KeyRegistryFilterer) FilterKeyRegistered(opts *bind.FilterOpts) (*KeyRegistryKeyRegisteredIterator, error) {

	logs, sub, err := _KeyRegistry.contract.FilterLogs(opts, "KeyRegistered")
	if err != nil {
		return nil, err
	}
	return &KeyRegistryKeyRegisteredIterator{contract: _KeyRegistry.contract, event: "KeyRegistered", logs: logs, sub: sub}, nil
}

// WatchKeyRegistered is a free log subscription operation binding the contract event 0xf77c9b4a9cc46843db8f22d9e84de14b5900316e5706ca63f7fb447ba07c3dc6.
//
// Solidity: event KeyRegistered(bytes arg0, address arg1)
func (_KeyRegistry *KeyRegistryFilterer) WatchKeyRegistered(opts *bind.WatchOpts, sink chan<- *KeyRegistryKeyRegistered) (event.Subscription, error) {

	logs, sub, err := _KeyRegistry.contract.WatchLogs(opts, "KeyRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyRegistryKeyRegistered)
				if err := _KeyRegistry.contract.UnpackLog(event, "KeyRegistered", log); err != nil {
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

// ParseKeyRegistered is a log parse operation binding the contract event 0xf77c9b4a9cc46843db8f22d9e84de14b5900316e5706ca63f7fb447ba07c3dc6.
//
// Solidity: event KeyRegistered(bytes arg0, address arg1)
func (_KeyRegistry *KeyRegistryFilterer) ParseKeyRegistered(log types.Log) (*KeyRegistryKeyRegistered, error) {
	event := new(KeyRegistryKeyRegistered)
	if err := _KeyRegistry.contract.UnpackLog(event, "KeyRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
