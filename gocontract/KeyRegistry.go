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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"KeyDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"KeyParentRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"KeyRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"disable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"disabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"keyOwner\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"parent\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"parentKey\",\"type\":\"bytes\"}],\"name\":\"registerParent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"parent\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610c36806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80638fc925aa1161005b5780638fc925aa14610101578063a15d581c14610114578063b10d937b14610136578063e44ccc651461014957600080fd5b80630f04a82e1461008d578063467d64c9146100b657806379fc09a2146100c957806382fbdc9c146100ec575b600080fd5b6100a061009b36600461089a565b61015c565b6040516100ad9190610929565b60405180910390f35b6100a06100c436600461089a565b610213565b6100dc6100d736600461089a565b61037b565b60405190151581526020016100ad565b6100ff6100fa36600461089a565b6103a9565b005b6100ff61010f36600461089a565b610455565b610127610122366004610959565b6105ae565b6040516100ad93929190610a0a565b6100ff610144366004610a36565b61066b565b6100dc61015736600461089a565b610784565b606060008383604051610170929190610aa2565b9081526020016040518091039020600101805461018c90610ab2565b80601f01602080910402602001604051908101604052809291908181526020018280546101b890610ab2565b80156102055780601f106101da57610100808354040283529160200191610205565b820191906000526020600020905b8154815290600101906020018083116101e857829003601f168201915b505050505090505b92915050565b606060008383604051610227929190610aa2565b9081526040519081900360200190205460ff16610253575060408051602081019091526000815261020d565b6000808484604051610266929190610aa2565b9081526020016040518091039020600101805461028290610ab2565b9050111561033d576000838360405161029c929190610aa2565b908152602001604051809103902060010180546102b890610ab2565b80601f01602080910402602001604051908101604052809291908181526020018280546102e490610ab2565b80156103315780601f1061030657610100808354040283529160200191610331565b820191906000526020600020905b81548152906001019060200180831161031457829003601f168201915b5050505050905061020d565b82828080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929695505050505050565b600080838360405161038e929190610aa2565b9081526040519081900360200190205460ff16905092915050565b602181146103b657600080fd5b600082826040516103c8929190610aa2565b9081526040519081900360200190205460ff16156103e557600080fd5b6001600083836040516103f9929190610aa2565b908152604051908190036020018120805492151560ff19909316929092179091557f2123184afbd7510b176b666deb82298c81b9a8b363e3a88f8cf4b726e4c81ae4906104499084908490610b16565b60405180910390a15050565b6021811461046257600080fd5b60008282604051610474929190610aa2565b9081526040519081900360200190205460ff1661049057600080fd5b600082826040516104a2929190610aa2565b908152602001604051809103902060010180546104be90610ab2565b151590506104f857600082826040516104d8929190610aa2565b60405190819003902060601c90503381146104f257600080fd5b50610545565b600080838360405161050b929190610aa2565b90815260200160405180910390206001016040516105299190610b32565b60405190819003902060601c905033811461054357600080fd5b505b600160008383604051610559929190610aa2565b90815260405190819003602001812080549215156101000261ff0019909316929092179091557f1eefa765080cc5f57ade4997fa66a727ed30bdef80fe40a6dd5fe0dce8f8abbd906104499084908490610b16565b80516020818301810180516000825292820191909301209152805460018201805460ff80841694610100909404169291906105e890610ab2565b80601f016020809104026020016040519081016040528092919081815260200182805461061490610ab2565b80156106615780601f1061063657610100808354040283529160200191610661565b820191906000526020600020905b81548152906001019060200180831161064457829003601f168201915b5050505050905083565b6021831461067857600080fd5b6021811461068557600080fd5b60008484604051610697929190610aa2565b9081526040519081900360200190205460ff166106b357600080fd5b600082826040516106c5929190610aa2565b9081526040519081900360200190205460ff166106e157600080fd5b600084846040516106f3929190610aa2565b60405190819003902060601c905033811461070d57600080fd5b828260008787604051610721929190610aa2565b9081526020016040518091039020600101919061073f9291906107b8565b507f3f629b715f22cf1453603aec7468ddc75788aab619ed052ccba9606638b8bb6d858585856040516107759493929190610bce565b60405180910390a15050505050565b6000808383604051610797929190610aa2565b9081526040519081900360200190205460ff61010090910416905092915050565b8280546107c490610ab2565b90600052602060002090601f0160209004810192826107e6576000855561082c565b82601f106107ff5782800160ff1982351617855561082c565b8280016001018555821561082c579182015b8281111561082c578235825591602001919060010190610811565b5061083892915061083c565b5090565b5b80821115610838576000815560010161083d565b60008083601f84011261086357600080fd5b50813567ffffffffffffffff81111561087b57600080fd5b60208301915083602082850101111561089357600080fd5b9250929050565b600080602083850312156108ad57600080fd5b823567ffffffffffffffff8111156108c457600080fd5b6108d085828601610851565b90969095509350505050565b6000815180845260005b81811015610902576020818501810151868301820152016108e6565b81811115610914576000602083870101525b50601f01601f19169290920160200192915050565b60208152600061093c60208301846108dc565b9392505050565b634e487b7160e01b600052604160045260246000fd5b60006020828403121561096b57600080fd5b813567ffffffffffffffff8082111561098357600080fd5b818401915084601f83011261099757600080fd5b8135818111156109a9576109a9610943565b604051601f8201601f19908116603f011681019083821181831017156109d1576109d1610943565b816040528281528760208487010111156109ea57600080fd5b826020860160208301376000928101602001929092525095945050505050565b83151581528215156020820152606060408201526000610a2d60608301846108dc565b95945050505050565b60008060008060408587031215610a4c57600080fd5b843567ffffffffffffffff80821115610a6457600080fd5b610a7088838901610851565b90965094506020870135915080821115610a8957600080fd5b50610a9687828801610851565b95989497509550505050565b8183823760009101908152919050565b600181811c90821680610ac657607f821691505b60208210811415610ae757634e487b7160e01b600052602260045260246000fd5b50919050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b602081526000610b2a602083018486610aed565b949350505050565b600080835481600182811c915080831680610b4e57607f831692505b6020808410821415610b6e57634e487b7160e01b86526022600452602486fd5b818015610b825760018114610b9357610bc0565b60ff19861689528489019650610bc0565b60008a81526020902060005b86811015610bb85781548b820152908501908301610b9f565b505084890196505b509498975050505050505050565b604081526000610be2604083018688610aed565b8281036020840152610bf5818587610aed565b97965050505050505056fea264697066735822122072f50e38f82eb89926f329c2d045ed8e89532e1f4e9217e34cf40ab227ccc7ef64736f6c634300080b0033",
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

// Exists is a free data retrieval call binding the contract method 0x79fc09a2.
//
// Solidity: function exists(bytes publicKey) view returns(bool)
func (_KeyRegistry *KeyRegistryCaller) Exists(opts *bind.CallOpts, publicKey []byte) (bool, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "exists", publicKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x79fc09a2.
//
// Solidity: function exists(bytes publicKey) view returns(bool)
func (_KeyRegistry *KeyRegistrySession) Exists(publicKey []byte) (bool, error) {
	return _KeyRegistry.Contract.Exists(&_KeyRegistry.CallOpts, publicKey)
}

// Exists is a free data retrieval call binding the contract method 0x79fc09a2.
//
// Solidity: function exists(bytes publicKey) view returns(bool)
func (_KeyRegistry *KeyRegistryCallerSession) Exists(publicKey []byte) (bool, error) {
	return _KeyRegistry.Contract.Exists(&_KeyRegistry.CallOpts, publicKey)
}

// KeyOwner is a free data retrieval call binding the contract method 0x467d64c9.
//
// Solidity: function keyOwner(bytes publicKey) view returns(bytes)
func (_KeyRegistry *KeyRegistryCaller) KeyOwner(opts *bind.CallOpts, publicKey []byte) ([]byte, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "keyOwner", publicKey)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// KeyOwner is a free data retrieval call binding the contract method 0x467d64c9.
//
// Solidity: function keyOwner(bytes publicKey) view returns(bytes)
func (_KeyRegistry *KeyRegistrySession) KeyOwner(publicKey []byte) ([]byte, error) {
	return _KeyRegistry.Contract.KeyOwner(&_KeyRegistry.CallOpts, publicKey)
}

// KeyOwner is a free data retrieval call binding the contract method 0x467d64c9.
//
// Solidity: function keyOwner(bytes publicKey) view returns(bytes)
func (_KeyRegistry *KeyRegistryCallerSession) KeyOwner(publicKey []byte) ([]byte, error) {
	return _KeyRegistry.Contract.KeyOwner(&_KeyRegistry.CallOpts, publicKey)
}

// Parent is a free data retrieval call binding the contract method 0x0f04a82e.
//
// Solidity: function parent(bytes publicKey) view returns(bytes)
func (_KeyRegistry *KeyRegistryCaller) Parent(opts *bind.CallOpts, publicKey []byte) ([]byte, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "parent", publicKey)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Parent is a free data retrieval call binding the contract method 0x0f04a82e.
//
// Solidity: function parent(bytes publicKey) view returns(bytes)
func (_KeyRegistry *KeyRegistrySession) Parent(publicKey []byte) ([]byte, error) {
	return _KeyRegistry.Contract.Parent(&_KeyRegistry.CallOpts, publicKey)
}

// Parent is a free data retrieval call binding the contract method 0x0f04a82e.
//
// Solidity: function parent(bytes publicKey) view returns(bytes)
func (_KeyRegistry *KeyRegistryCallerSession) Parent(publicKey []byte) ([]byte, error) {
	return _KeyRegistry.Contract.Parent(&_KeyRegistry.CallOpts, publicKey)
}

// Registry is a free data retrieval call binding the contract method 0xa15d581c.
//
// Solidity: function registry(bytes ) view returns(bool initialized, bool disabled, bytes parent)
func (_KeyRegistry *KeyRegistryCaller) Registry(opts *bind.CallOpts, arg0 []byte) (struct {
	Initialized bool
	Disabled    bool
	Parent      []byte
}, error) {
	var out []interface{}
	err := _KeyRegistry.contract.Call(opts, &out, "registry", arg0)

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
func (_KeyRegistry *KeyRegistrySession) Registry(arg0 []byte) (struct {
	Initialized bool
	Disabled    bool
	Parent      []byte
}, error) {
	return _KeyRegistry.Contract.Registry(&_KeyRegistry.CallOpts, arg0)
}

// Registry is a free data retrieval call binding the contract method 0xa15d581c.
//
// Solidity: function registry(bytes ) view returns(bool initialized, bool disabled, bytes parent)
func (_KeyRegistry *KeyRegistryCallerSession) Registry(arg0 []byte) (struct {
	Initialized bool
	Disabled    bool
	Parent      []byte
}, error) {
	return _KeyRegistry.Contract.Registry(&_KeyRegistry.CallOpts, arg0)
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

// RegisterParent is a paid mutator transaction binding the contract method 0xb10d937b.
//
// Solidity: function registerParent(bytes publicKey, bytes parentKey) returns()
func (_KeyRegistry *KeyRegistryTransactor) RegisterParent(opts *bind.TransactOpts, publicKey []byte, parentKey []byte) (*types.Transaction, error) {
	return _KeyRegistry.contract.Transact(opts, "registerParent", publicKey, parentKey)
}

// RegisterParent is a paid mutator transaction binding the contract method 0xb10d937b.
//
// Solidity: function registerParent(bytes publicKey, bytes parentKey) returns()
func (_KeyRegistry *KeyRegistrySession) RegisterParent(publicKey []byte, parentKey []byte) (*types.Transaction, error) {
	return _KeyRegistry.Contract.RegisterParent(&_KeyRegistry.TransactOpts, publicKey, parentKey)
}

// RegisterParent is a paid mutator transaction binding the contract method 0xb10d937b.
//
// Solidity: function registerParent(bytes publicKey, bytes parentKey) returns()
func (_KeyRegistry *KeyRegistryTransactorSession) RegisterParent(publicKey []byte, parentKey []byte) (*types.Transaction, error) {
	return _KeyRegistry.Contract.RegisterParent(&_KeyRegistry.TransactOpts, publicKey, parentKey)
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

// KeyRegistryKeyParentRegisteredIterator is returned from FilterKeyParentRegistered and is used to iterate over the raw logs and unpacked data for KeyParentRegistered events raised by the KeyRegistry contract.
type KeyRegistryKeyParentRegisteredIterator struct {
	Event *KeyRegistryKeyParentRegistered // Event containing the contract specifics and raw log

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
func (it *KeyRegistryKeyParentRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyRegistryKeyParentRegistered)
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
		it.Event = new(KeyRegistryKeyParentRegistered)
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
func (it *KeyRegistryKeyParentRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyRegistryKeyParentRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyRegistryKeyParentRegistered represents a KeyParentRegistered event raised by the KeyRegistry contract.
type KeyRegistryKeyParentRegistered struct {
	Arg0 []byte
	Arg1 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterKeyParentRegistered is a free log retrieval operation binding the contract event 0x3f629b715f22cf1453603aec7468ddc75788aab619ed052ccba9606638b8bb6d.
//
// Solidity: event KeyParentRegistered(bytes arg0, bytes arg1)
func (_KeyRegistry *KeyRegistryFilterer) FilterKeyParentRegistered(opts *bind.FilterOpts) (*KeyRegistryKeyParentRegisteredIterator, error) {

	logs, sub, err := _KeyRegistry.contract.FilterLogs(opts, "KeyParentRegistered")
	if err != nil {
		return nil, err
	}
	return &KeyRegistryKeyParentRegisteredIterator{contract: _KeyRegistry.contract, event: "KeyParentRegistered", logs: logs, sub: sub}, nil
}

// WatchKeyParentRegistered is a free log subscription operation binding the contract event 0x3f629b715f22cf1453603aec7468ddc75788aab619ed052ccba9606638b8bb6d.
//
// Solidity: event KeyParentRegistered(bytes arg0, bytes arg1)
func (_KeyRegistry *KeyRegistryFilterer) WatchKeyParentRegistered(opts *bind.WatchOpts, sink chan<- *KeyRegistryKeyParentRegistered) (event.Subscription, error) {

	logs, sub, err := _KeyRegistry.contract.WatchLogs(opts, "KeyParentRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyRegistryKeyParentRegistered)
				if err := _KeyRegistry.contract.UnpackLog(event, "KeyParentRegistered", log); err != nil {
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
func (_KeyRegistry *KeyRegistryFilterer) ParseKeyParentRegistered(log types.Log) (*KeyRegistryKeyParentRegistered, error) {
	event := new(KeyRegistryKeyParentRegistered)
	if err := _KeyRegistry.contract.UnpackLog(event, "KeyParentRegistered", log); err != nil {
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
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterKeyRegistered is a free log retrieval operation binding the contract event 0x2123184afbd7510b176b666deb82298c81b9a8b363e3a88f8cf4b726e4c81ae4.
//
// Solidity: event KeyRegistered(bytes arg0)
func (_KeyRegistry *KeyRegistryFilterer) FilterKeyRegistered(opts *bind.FilterOpts) (*KeyRegistryKeyRegisteredIterator, error) {

	logs, sub, err := _KeyRegistry.contract.FilterLogs(opts, "KeyRegistered")
	if err != nil {
		return nil, err
	}
	return &KeyRegistryKeyRegisteredIterator{contract: _KeyRegistry.contract, event: "KeyRegistered", logs: logs, sub: sub}, nil
}

// WatchKeyRegistered is a free log subscription operation binding the contract event 0x2123184afbd7510b176b666deb82298c81b9a8b363e3a88f8cf4b726e4c81ae4.
//
// Solidity: event KeyRegistered(bytes arg0)
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

// ParseKeyRegistered is a log parse operation binding the contract event 0x2123184afbd7510b176b666deb82298c81b9a8b363e3a88f8cf4b726e4c81ae4.
//
// Solidity: event KeyRegistered(bytes arg0)
func (_KeyRegistry *KeyRegistryFilterer) ParseKeyRegistered(log types.Log) (*KeyRegistryKeyRegistered, error) {
	event := new(KeyRegistryKeyRegistered)
	if err := _KeyRegistry.contract.UnpackLog(event, "KeyRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
