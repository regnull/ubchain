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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"changePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"lookup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061090b806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80631fa19596146100515780635d2097d21461006657806392255fbf14610079578063f67187ac1461008c575b600080fd5b61006461005f36600461046c565b6100b7565b005b6100646100743660046104ce565b610125565b6100646100873660046105a9565b61025a565b61009f61009a366004610615565b610344565b6040516100ae93929190610657565b60405180910390f35b60006001600160a01b0316600084846040516100d49291906106bc565b908152604051908190036020019020600101546001600160a01b0316146100fa57600080fd5b806000848460405161010d9291906106bc565b90815260405190819003602001902060020155505050565b60008084846040516101389291906106bc565b9081526020016040518091039020600201541161015457600080fd5b600083836040516101669291906106bc565b908152604051908190036020018120600101546001600160a01b0316906108fc9060009061019790879087906106bc565b90815260405190819003602001812060020154801590920291906000818181858888f193505050501580156101d0573d6000803e3d6000fd5b5060008084846040516101e49291906106bc565b908152602001604051809103902060010160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550806000848460405161022c9291906106bc565b908152604051908190036020019020906102469082610755565b50600080848460405161010d9291906106bc565b6021831461026757600080fd5b60006001600160a01b0316600083836040516102849291906106bc565b908152604051908190036020019020600101546001600160a01b0316146102aa57600080fd5b60038110156102b857600080fd5b60408111156102c657600080fd5b60008083836040516102d99291906106bc565b908152602001604051809103902060010160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508383600084846040516103229291906106bc565b9081526040519081900360200190209161033d919083610815565b5050505050565b600060606000806000868660405161035d9291906106bc565b9081526040519081900360200190206001810154600282015482549293506001600160a01b03909116918391908290610395906106cc565b80601f01602080910402602001604051908101604052809291908181526020018280546103c1906106cc565b801561040e5780601f106103e35761010080835404028352916020019161040e565b820191906000526020600020905b8154815290600101906020018083116103f157829003601f168201915b50505050509150935093509350509250925092565b60008083601f84011261043557600080fd5b50813567ffffffffffffffff81111561044d57600080fd5b60208301915083602082850101111561046557600080fd5b9250929050565b60008060006040848603121561048157600080fd5b833567ffffffffffffffff81111561049857600080fd5b6104a486828701610423565b909790965060209590950135949350505050565b634e487b7160e01b600052604160045260246000fd5b6000806000604084860312156104e357600080fd5b833567ffffffffffffffff808211156104fb57600080fd5b61050787838801610423565b9095509350602086013591508082111561052057600080fd5b818601915086601f83011261053457600080fd5b813581811115610546576105466104b8565b604051601f8201601f19908116603f0116810190838211818310171561056e5761056e6104b8565b8160405282815289602084870101111561058757600080fd5b8260208601602083013760006020848301015280955050505050509250925092565b600080600080604085870312156105bf57600080fd5b843567ffffffffffffffff808211156105d757600080fd5b6105e388838901610423565b909650945060208701359150808211156105fc57600080fd5b5061060987828801610423565b95989497509550505050565b6000806020838503121561062857600080fd5b823567ffffffffffffffff81111561063f57600080fd5b61064b85828601610423565b90969095509350505050565b60018060a01b038416815260006020606081840152845180606085015260005b8181101561069357868101830151858201608001528201610677565b506000608082860101526080601f19601f83011685010192505050826040830152949350505050565b8183823760009101908152919050565b600181811c908216806106e057607f821691505b60208210810361070057634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561075057600081815260208120601f850160051c8101602086101561072d5750805b601f850160051c820191505b8181101561074c57828155600101610739565b5050505b505050565b815167ffffffffffffffff81111561076f5761076f6104b8565b6107838161077d84546106cc565b84610706565b602080601f8311600181146107b857600084156107a05750858301515b600019600386901b1c1916600185901b17855561074c565b600085815260208120601f198616915b828110156107e7578886015182559484019460019091019084016107c8565b50858210156108055787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b67ffffffffffffffff83111561082d5761082d6104b8565b6108418361083b83546106cc565b83610706565b6000601f841160018114610875576000851561085d5750838201355b600019600387901b1c1916600186901b17835561033d565b600083815260209020601f19861690835b828110156108a65786850135825560209485019460019092019101610886565b50868210156108c35760001960f88860031b161c19848701351681555b505060018560011b018355505050505056fea264697066735822122093fc353cd1355f39fcd7abf3557198a712883d038896c31a37ef2bc9a59a585b64736f6c63430008110033",
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

// Lookup is a free data retrieval call binding the contract method 0xf67187ac.
//
// Solidity: function lookup(string name) view returns(address owner, bytes publicKey, uint256 price)
func (_NameRegistry *NameRegistryCaller) Lookup(opts *bind.CallOpts, name string) (struct {
	Owner     common.Address
	PublicKey []byte
	Price     *big.Int
}, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "lookup", name)

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

// Lookup is a free data retrieval call binding the contract method 0xf67187ac.
//
// Solidity: function lookup(string name) view returns(address owner, bytes publicKey, uint256 price)
func (_NameRegistry *NameRegistrySession) Lookup(name string) (struct {
	Owner     common.Address
	PublicKey []byte
	Price     *big.Int
}, error) {
	return _NameRegistry.Contract.Lookup(&_NameRegistry.CallOpts, name)
}

// Lookup is a free data retrieval call binding the contract method 0xf67187ac.
//
// Solidity: function lookup(string name) view returns(address owner, bytes publicKey, uint256 price)
func (_NameRegistry *NameRegistryCallerSession) Lookup(name string) (struct {
	Owner     common.Address
	PublicKey []byte
	Price     *big.Int
}, error) {
	return _NameRegistry.Contract.Lookup(&_NameRegistry.CallOpts, name)
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

// ChangePrice is a paid mutator transaction binding the contract method 0x1fa19596.
//
// Solidity: function changePrice(string name, uint256 price) returns()
func (_NameRegistry *NameRegistryTransactor) ChangePrice(opts *bind.TransactOpts, name string, price *big.Int) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "changePrice", name, price)
}

// ChangePrice is a paid mutator transaction binding the contract method 0x1fa19596.
//
// Solidity: function changePrice(string name, uint256 price) returns()
func (_NameRegistry *NameRegistrySession) ChangePrice(name string, price *big.Int) (*types.Transaction, error) {
	return _NameRegistry.Contract.ChangePrice(&_NameRegistry.TransactOpts, name, price)
}

// ChangePrice is a paid mutator transaction binding the contract method 0x1fa19596.
//
// Solidity: function changePrice(string name, uint256 price) returns()
func (_NameRegistry *NameRegistryTransactorSession) ChangePrice(name string, price *big.Int) (*types.Transaction, error) {
	return _NameRegistry.Contract.ChangePrice(&_NameRegistry.TransactOpts, name, price)
}

// Register is a paid mutator transaction binding the contract method 0x92255fbf.
//
// Solidity: function register(bytes publicKey, string name) returns()
func (_NameRegistry *NameRegistryTransactor) Register(opts *bind.TransactOpts, publicKey []byte, name string) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "register", publicKey, name)
}

// Register is a paid mutator transaction binding the contract method 0x92255fbf.
//
// Solidity: function register(bytes publicKey, string name) returns()
func (_NameRegistry *NameRegistrySession) Register(publicKey []byte, name string) (*types.Transaction, error) {
	return _NameRegistry.Contract.Register(&_NameRegistry.TransactOpts, publicKey, name)
}

// Register is a paid mutator transaction binding the contract method 0x92255fbf.
//
// Solidity: function register(bytes publicKey, string name) returns()
func (_NameRegistry *NameRegistryTransactorSession) Register(publicKey []byte, name string) (*types.Transaction, error) {
	return _NameRegistry.Contract.Register(&_NameRegistry.TransactOpts, publicKey, name)
}
