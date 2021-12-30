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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"KeyDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"KeyParentRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"KeyRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"disable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"parentKey\",\"type\":\"bytes\"}],\"name\":\"registerParent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"parent\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506108e5806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806382fbdc9c146100515780638fc925aa14610066578063a15d581c14610079578063b10d937b146100a4575b600080fd5b61006461005f366004610574565b6100b7565b005b610064610074366004610574565b610163565b61008c6100873660046105cc565b6102bc565b60405161009b9392919061067d565b60405180910390f35b6100646100b23660046106e5565b610379565b602181146100c457600080fd5b600082826040516100d6929190610751565b9081526040519081900360200190205460ff16156100f357600080fd5b600160008383604051610107929190610751565b908152604051908190036020018120805492151560ff19909316929092179091557f2123184afbd7510b176b666deb82298c81b9a8b363e3a88f8cf4b726e4c81ae490610157908490849061078a565b60405180910390a15050565b6021811461017057600080fd5b60008282604051610182929190610751565b9081526040519081900360200190205460ff1661019e57600080fd5b600082826040516101b0929190610751565b908152602001604051809103902060010180546101cc906107a6565b1515905061020657600082826040516101e6929190610751565b60405190819003902060601c905033811461020057600080fd5b50610253565b6000808383604051610219929190610751565b908152602001604051809103902060010160405161023791906107e1565b60405190819003902060601c905033811461025157600080fd5b505b600160008383604051610267929190610751565b90815260405190819003602001812080549215156101000261ff0019909316929092179091557f1eefa765080cc5f57ade4997fa66a727ed30bdef80fe40a6dd5fe0dce8f8abbd90610157908490849061078a565b80516020818301810180516000825292820191909301209152805460018201805460ff80841694610100909404169291906102f6906107a6565b80601f0160208091040260200160405190810160405280929190818152602001828054610322906107a6565b801561036f5780601f106103445761010080835404028352916020019161036f565b820191906000526020600020905b81548152906001019060200180831161035257829003601f168201915b5050505050905083565b6021831461038657600080fd5b6021811461039357600080fd5b600084846040516103a5929190610751565b9081526040519081900360200190205460ff166103c157600080fd5b600082826040516103d3929190610751565b9081526040519081900360200190205460ff166103ef57600080fd5b60008484604051610401929190610751565b60405190819003902060601c905033811461041b57600080fd5b82826000878760405161042f929190610751565b9081526020016040518091039020600101919061044d929190610492565b507f3f629b715f22cf1453603aec7468ddc75788aab619ed052ccba9606638b8bb6d85858585604051610483949392919061087d565b60405180910390a15050505050565b82805461049e906107a6565b90600052602060002090601f0160209004810192826104c05760008555610506565b82601f106104d95782800160ff19823516178555610506565b82800160010185558215610506579182015b828111156105065782358255916020019190600101906104eb565b50610512929150610516565b5090565b5b808211156105125760008155600101610517565b60008083601f84011261053d57600080fd5b50813567ffffffffffffffff81111561055557600080fd5b60208301915083602082850101111561056d57600080fd5b9250929050565b6000806020838503121561058757600080fd5b823567ffffffffffffffff81111561059e57600080fd5b6105aa8582860161052b565b90969095509350505050565b634e487b7160e01b600052604160045260246000fd5b6000602082840312156105de57600080fd5b813567ffffffffffffffff808211156105f657600080fd5b818401915084601f83011261060a57600080fd5b81358181111561061c5761061c6105b6565b604051601f8201601f19908116603f01168101908382118183101715610644576106446105b6565b8160405282815287602084870101111561065d57600080fd5b826020860160208301376000928101602001929092525095945050505050565b8315158152600060208415158184015260606040840152835180606085015260005b818110156106bb5785810183015185820160800152820161069f565b818111156106cd576000608083870101525b50601f01601f19169290920160800195945050505050565b600080600080604085870312156106fb57600080fd5b843567ffffffffffffffff8082111561071357600080fd5b61071f8883890161052b565b9096509450602087013591508082111561073857600080fd5b506107458782880161052b565b95989497509550505050565b8183823760009101908152919050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60208152600061079e602083018486610761565b949350505050565b600181811c908216806107ba57607f821691505b602082108114156107db57634e487b7160e01b600052602260045260246000fd5b50919050565b600080835481600182811c9150808316806107fd57607f831692505b602080841082141561081d57634e487b7160e01b86526022600452602486fd5b81801561083157600181146108425761086f565b60ff1986168952848901965061086f565b60008a81526020902060005b868110156108675781548b82015290850190830161084e565b505084890196505b509498975050505050505050565b604081526000610891604083018688610761565b82810360208401526108a4818587610761565b97965050505050505056fea2646970667358221220f72b2524f9b80741988b0f06c3b197b56b2e15d3bb8c48940bf9b60f033e00b364736f6c634300080b0033",
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
// Solidity: function registry(bytes ) view returns(bool initialized, bool disabled, bytes parent)
func (_Keyregistry *KeyregistryCaller) Registry(opts *bind.CallOpts, arg0 []byte) (struct {
	Initialized bool
	Disabled    bool
	Parent      []byte
}, error) {
	var out []interface{}
	err := _Keyregistry.contract.Call(opts, &out, "registry", arg0)

	outstruct := new(struct {
		Initialized bool
		Disabled    bool
		Parent      []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Initialized = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Disabled = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Parent = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Registry is a free data retrieval call binding the contract method 0xa15d581c.
//
// Solidity: function registry(bytes ) view returns(bool initialized, bool disabled, bytes parent)
func (_Keyregistry *KeyregistrySession) Registry(arg0 []byte) (struct {
	Initialized bool
	Disabled    bool
	Parent      []byte
}, error) {
	return _Keyregistry.Contract.Registry(&_Keyregistry.CallOpts, arg0)
}

// Registry is a free data retrieval call binding the contract method 0xa15d581c.
//
// Solidity: function registry(bytes ) view returns(bool initialized, bool disabled, bytes parent)
func (_Keyregistry *KeyregistryCallerSession) Registry(arg0 []byte) (struct {
	Initialized bool
	Disabled    bool
	Parent      []byte
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

// RegisterParent is a paid mutator transaction binding the contract method 0xb10d937b.
//
// Solidity: function registerParent(bytes publicKey, bytes parentKey) returns()
func (_Keyregistry *KeyregistryTransactor) RegisterParent(opts *bind.TransactOpts, publicKey []byte, parentKey []byte) (*types.Transaction, error) {
	return _Keyregistry.contract.Transact(opts, "registerParent", publicKey, parentKey)
}

// RegisterParent is a paid mutator transaction binding the contract method 0xb10d937b.
//
// Solidity: function registerParent(bytes publicKey, bytes parentKey) returns()
func (_Keyregistry *KeyregistrySession) RegisterParent(publicKey []byte, parentKey []byte) (*types.Transaction, error) {
	return _Keyregistry.Contract.RegisterParent(&_Keyregistry.TransactOpts, publicKey, parentKey)
}

// RegisterParent is a paid mutator transaction binding the contract method 0xb10d937b.
//
// Solidity: function registerParent(bytes publicKey, bytes parentKey) returns()
func (_Keyregistry *KeyregistryTransactorSession) RegisterParent(publicKey []byte, parentKey []byte) (*types.Transaction, error) {
	return _Keyregistry.Contract.RegisterParent(&_Keyregistry.TransactOpts, publicKey, parentKey)
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

// KeyregistryKeyParentRegisteredIterator is returned from FilterKeyParentRegistered and is used to iterate over the raw logs and unpacked data for KeyParentRegistered events raised by the Keyregistry contract.
type KeyregistryKeyParentRegisteredIterator struct {
	Event *KeyregistryKeyParentRegistered // Event containing the contract specifics and raw log

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
func (it *KeyregistryKeyParentRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyregistryKeyParentRegistered)
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
		it.Event = new(KeyregistryKeyParentRegistered)
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
func (it *KeyregistryKeyParentRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyregistryKeyParentRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyregistryKeyParentRegistered represents a KeyParentRegistered event raised by the Keyregistry contract.
type KeyregistryKeyParentRegistered struct {
	Arg0 []byte
	Arg1 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterKeyParentRegistered is a free log retrieval operation binding the contract event 0x3f629b715f22cf1453603aec7468ddc75788aab619ed052ccba9606638b8bb6d.
//
// Solidity: event KeyParentRegistered(bytes arg0, bytes arg1)
func (_Keyregistry *KeyregistryFilterer) FilterKeyParentRegistered(opts *bind.FilterOpts) (*KeyregistryKeyParentRegisteredIterator, error) {

	logs, sub, err := _Keyregistry.contract.FilterLogs(opts, "KeyParentRegistered")
	if err != nil {
		return nil, err
	}
	return &KeyregistryKeyParentRegisteredIterator{contract: _Keyregistry.contract, event: "KeyParentRegistered", logs: logs, sub: sub}, nil
}

// WatchKeyParentRegistered is a free log subscription operation binding the contract event 0x3f629b715f22cf1453603aec7468ddc75788aab619ed052ccba9606638b8bb6d.
//
// Solidity: event KeyParentRegistered(bytes arg0, bytes arg1)
func (_Keyregistry *KeyregistryFilterer) WatchKeyParentRegistered(opts *bind.WatchOpts, sink chan<- *KeyregistryKeyParentRegistered) (event.Subscription, error) {

	logs, sub, err := _Keyregistry.contract.WatchLogs(opts, "KeyParentRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyregistryKeyParentRegistered)
				if err := _Keyregistry.contract.UnpackLog(event, "KeyParentRegistered", log); err != nil {
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

// ParseKeyParentRegistered is a log parse operation binding the contract event 0x3f629b715f22cf1453603aec7468ddc75788aab619ed052ccba9606638b8bb6d.
//
// Solidity: event KeyParentRegistered(bytes arg0, bytes arg1)
func (_Keyregistry *KeyregistryFilterer) ParseKeyParentRegistered(log types.Log) (*KeyregistryKeyParentRegistered, error) {
	event := new(KeyregistryKeyParentRegistered)
	if err := _Keyregistry.contract.UnpackLog(event, "KeyParentRegistered", log); err != nil {
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
