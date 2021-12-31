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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keyRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nameRegistryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"protocol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516107aa3803806107aa83398101604081905261002f9161007c565b600080546001600160a01b039384166001600160a01b031991821617909155600180549290931691161790556100af565b80516001600160a01b038116811461007757600080fd5b919050565b6000806040838503121561008f57600080fd5b61009883610060565b91506100a660208401610060565b90509250929050565b6106ec806100be6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80634cd08d031461003b578063b13ac26414610050575b600080fd5b61004e61004936600461038c565b610079565b005b61006361005e3660046104eb565b6101e8565b60405161007091906105ab565b60405180910390f35b6001546040516369bd764960e11b81526000916001600160a01b03169063d37aec92906100ac908a908a906004016105c5565b600060405180830381865afa1580156100c9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526100f191908101906105f4565b9050805160211461010157600080fd5b6000805460405163467d64c960e01b81526001600160a01b039091169063467d64c9906101329085906004016105ab565b600060405180830381865afa15801561014f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261017791908101906105f4565b8051602082012090915060601c33811461019057600080fd5b848460028b8b6040516101a492919061066b565b908152602001604051809103902089896040516101c292919061066b565b9081526040519081900360200190206101dc9290916102aa565b50505050505050505050565b81516020818401810180516002825292820194820194909420919093528151808301840180519281529084019290930191909120915280546102299061067b565b80601f01602080910402602001604051908101604052809291908181526020018280546102559061067b565b80156102a25780601f10610277576101008083540402835291602001916102a2565b820191906000526020600020905b81548152906001019060200180831161028557829003601f168201915b505050505081565b8280546102b69061067b565b90600052602060002090601f0160209004810192826102d8576000855561031e565b82601f106102f15782800160ff1982351617855561031e565b8280016001018555821561031e579182015b8281111561031e578235825591602001919060010190610303565b5061032a92915061032e565b5090565b5b8082111561032a576000815560010161032f565b60008083601f84011261035557600080fd5b50813567ffffffffffffffff81111561036d57600080fd5b60208301915083602082850101111561038557600080fd5b9250929050565b600080600080600080606087890312156103a557600080fd5b863567ffffffffffffffff808211156103bd57600080fd5b6103c98a838b01610343565b909850965060208901359150808211156103e257600080fd5b6103ee8a838b01610343565b9096509450604089013591508082111561040757600080fd5b5061041489828a01610343565b979a9699509497509295939492505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561046557610465610426565b604052919050565b600067ffffffffffffffff82111561048757610487610426565b50601f01601f191660200190565b600082601f8301126104a657600080fd5b81356104b96104b48261046d565b61043c565b8181528460208386010111156104ce57600080fd5b816020850160208301376000918101602001919091529392505050565b600080604083850312156104fe57600080fd5b823567ffffffffffffffff8082111561051657600080fd5b61052286838701610495565b9350602085013591508082111561053857600080fd5b5061054585828601610495565b9150509250929050565b60005b8381101561056a578181015183820152602001610552565b83811115610579576000848401525b50505050565b6000815180845261059781602086016020860161054f565b601f01601f19169290920160200192915050565b6020815260006105be602083018461057f565b9392505050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b60006020828403121561060657600080fd5b815167ffffffffffffffff81111561061d57600080fd5b8201601f8101841361062e57600080fd5b805161063c6104b48261046d565b81815285602083850101111561065157600080fd5b61066282602083016020860161054f565b95945050505050565b8183823760009101908152919050565b600181811c9082168061068f57607f821691505b602082108114156106b057634e487b7160e01b600052602260045260246000fd5b5091905056fea2646970667358221220d47662d0b437272f0d1ab0b248c892f7c33298071b819b2d428e98dbb3ba65c664736f6c634300080b0033",
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
