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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"protocol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"}],\"name\":\"ConnectorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NameOwnershipChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NameRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"PriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"PublicKeyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"Sale\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"changePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"protocol\",\"type\":\"string\"}],\"name\":\"lookupConnector\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"lookupName\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"protocol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"}],\"name\":\"registerConnector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"registerName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"updatePublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061144a806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063639cdb951161005b578063639cdb95146100db578063c0e793c214610106578063d346e3bd14610119578063f11a1a9a1461013957600080fd5b80631fa195961461008d578063208b5f38146100a257806356c792db146100b55780635d2097d2146100c8575b600080fd5b6100a061009b366004610d05565b61014c565b005b6100a06100b0366004610d51565b610234565b6100a06100c3366004610d51565b610330565b6100a06100d6366004610dd3565b6105fb565b6100ee6100e9366004610eae565b6107b1565b6040516100fd93929190610f40565b60405180910390f35b6100a0610114366004610f74565b6108d1565b61012c610127366004610d51565b6109c7565b6040516100fd9190610fd7565b6100a0610147366004610ff1565b610adc565b600061018d84848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610bf392505050565b9050336001600160a01b03166000826040516101a9919061108b565b908152604051908190036020019020600101546001600160a01b0316146101cf57600080fd5b816000826040516101e0919061108b565b9081526020016040518091039020600201819055507f29ecd1e0988af1492e43256007437db411881757b3e6e808c9f04847b264178c81836040516102269291906110a7565b60405180910390a150505050565b600061027583838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610bf392505050565b90506021841461028457600080fd5b336001600160a01b031660008260405161029e919061108b565b908152604051908190036020019020600101546001600160a01b0316146102c457600080fd5b84846000836040516102d6919061108b565b908152604051908190036020019020916102f1919083611152565b507f0ba4a7233a2cfb937aa9644f11f49345b443fe8d5ab3494a1879faa500728d3e816040516103219190610fd7565b60405180910390a15050505050565b6021831461033d57600080fd5b60006001600160a01b03166000838360405161035a929190611213565b908152604051908190036020019020600101546001600160a01b03161461038057600080fd5b600181101561038e57600080fd5b604081111561039c57600080fd5b600082828080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509394505050505b81518110156105405760008282815181106103f4576103f4611223565b01602001516001600160f81b0319169050600360fc1b811080156104265750602d60f81b6001600160f81b0319821614155b8061043e5750603d60f91b6001600160f81b03198216115b806104705750601d60f91b6001600160f81b03198216108015906104705750600160fe1b6001600160f81b0319821611155b806104a25750605b60f81b6001600160f81b03198216108015906104a25750602f60f91b6001600160f81b0319821611155b806104ba5750600360fd1b6001600160f81b03198216145b156104fa5760405162461bcd60e51b815260206004820152600c60248201526b696e76616c6964206e616d6560a01b604482015260640160405180910390fd5b61050381610c6d565b83838151811061051557610515611223565b60200101906001600160f81b031916908160001a9053505080806105389061124f565b9150506103d7565b50600081905033600082604051610557919061108b565b908152602001604051809103902060010160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550858560008360405161059e919061108b565b908152604051908190036020019020916105b9919083611152565b507f1c6eac0e720ec22bb0653aec9c19985633a4fb07971cf973096c2f8e3c37c17f81336040516105eb929190611268565b60405180910390a1505050505050565b600061063c84848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610bf392505050565b90506000808260405161064f919061108b565b9081526020016040518091039020600201541161066b57600080fd5b6000808260405161067c919061108b565b90815260200160405180910390206002015490506000826040516106a0919061108b565b908152604051908190036020018120600101546001600160a01b03169082156108fc029083906000818181858888f193505050501580156106e5573d6000803e3d6000fd5b50336000836040516106f7919061108b565b908152602001604051809103902060010160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508260008360405161073d919061108b565b908152604051908190036020019020906107579082611292565b5060008083604051610769919061108b565b9081526020016040518091039020600201819055507f74d699b6db148dc632a2dbe2c680980f546b0382af2e71a7d4bf51c3193adb9a82823360405161032193929190611352565b600060606000806107f786868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610bf392505050565b90506000808260405161080a919061108b565b9081526040519081900360200190206001810154600282015482549293506001600160a01b03909116918391908290610842906110c9565b80601f016020809104026020016040519081016040528092919081815260200182805461086e906110c9565b80156108bb5780601f10610890576101008083540402835291602001916108bb565b820191906000526020600020905b81548152906001019060200180831161089e57829003601f168201915b5050505050915094509450945050509250925092565b600061091284848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610bf392505050565b9050336001600160a01b031660008260405161092e919061108b565b908152604051908190036020019020600101546001600160a01b03161461095457600080fd5b81600082604051610965919061108b565b90815260405190819003602001812060010180546001600160a01b03939093166001600160a01b0319909316929092179091557f8148cdc9d89df5e18d29f32adde657a148d30c0da1546f561cd58739a2057d02906102269083908590611268565b60606000610a0a86868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610bf392505050565b9050600181604051610a1c919061108b565b90815260200160405180910390208484604051610a3a929190611213565b90815260200160405180910390208054610a53906110c9565b80601f0160208091040260200160405190810160405280929190818152602001828054610a7f906110c9565b8015610acc5780601f10610aa157610100808354040283529160200191610acc565b820191906000526020600020905b815481529060010190602001808311610aaf57829003601f168201915b5050505050915050949350505050565b6000610b1d87878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610bf392505050565b9050336001600160a01b0316600082604051610b39919061108b565b908152604051908190036020019020600101546001600160a01b031614610b5f57600080fd5b8282600183604051610b71919061108b565b90815260200160405180910390208787604051610b8f929190611213565b90815260200160405180910390209182610baa929190611152565b507f6bcbdc9c6b89c8a6b289cc4308c6f6841a0e25d81c5fe34b1dedb2848c5f5ca38186868686604051610be29594939291906113ae565b60405180910390a150505050505050565b60608160005b8151811015610c6657610c2b828281518110610c1757610c17611223565b01602001516001600160f81b031916610c6d565b828281518110610c3d57610c3d611223565b60200101906001600160f81b031916908160001a90535080610c5e8161124f565b915050610bf9565b5092915050565b6000604160f81b6001600160f81b0319831610801590610c9b5750602d60f91b6001600160f81b0319831611155b15610cb857610caf60f883901c60206113f5565b60f81b92915050565b5090565b60008083601f840112610cce57600080fd5b50813567ffffffffffffffff811115610ce657600080fd5b602083019150836020828501011115610cfe57600080fd5b9250929050565b600080600060408486031215610d1a57600080fd5b833567ffffffffffffffff811115610d3157600080fd5b610d3d86828701610cbc565b909790965060209590950135949350505050565b60008060008060408587031215610d6757600080fd5b843567ffffffffffffffff80821115610d7f57600080fd5b610d8b88838901610cbc565b90965094506020870135915080821115610da457600080fd5b50610db187828801610cbc565b95989497509550505050565b634e487b7160e01b600052604160045260246000fd5b600080600060408486031215610de857600080fd5b833567ffffffffffffffff80821115610e0057600080fd5b610e0c87838801610cbc565b90955093506020860135915080821115610e2557600080fd5b818601915086601f830112610e3957600080fd5b813581811115610e4b57610e4b610dbd565b604051601f8201601f19908116603f01168101908382118183101715610e7357610e73610dbd565b81604052828152896020848701011115610e8c57600080fd5b8260208601602083013760006020848301015280955050505050509250925092565b60008060208385031215610ec157600080fd5b823567ffffffffffffffff811115610ed857600080fd5b610ee485828601610cbc565b90969095509350505050565b60005b83811015610f0b578181015183820152602001610ef3565b50506000910152565b60008151808452610f2c816020860160208601610ef0565b601f01601f19169290920160200192915050565b6001600160a01b0384168152606060208201819052600090610f6490830185610f14565b9050826040830152949350505050565b600080600060408486031215610f8957600080fd5b833567ffffffffffffffff811115610fa057600080fd5b610fac86828701610cbc565b90945092505060208401356001600160a01b0381168114610fcc57600080fd5b809150509250925092565b602081526000610fea6020830184610f14565b9392505050565b6000806000806000806060878903121561100a57600080fd5b863567ffffffffffffffff8082111561102257600080fd5b61102e8a838b01610cbc565b9098509650602089013591508082111561104757600080fd5b6110538a838b01610cbc565b9096509450604089013591508082111561106c57600080fd5b5061107989828a01610cbc565b979a9699509497509295939492505050565b6000825161109d818460208701610ef0565b9190910192915050565b6040815260006110ba6040830185610f14565b90508260208301529392505050565b600181811c908216806110dd57607f821691505b6020821081036110fd57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561114d57600081815260208120601f850160051c8101602086101561112a5750805b601f850160051c820191505b8181101561114957828155600101611136565b5050505b505050565b67ffffffffffffffff83111561116a5761116a610dbd565b61117e8361117883546110c9565b83611103565b6000601f8411600181146111b2576000851561119a5750838201355b600019600387901b1c1916600186901b17835561120c565b600083815260209020601f19861690835b828110156111e357868501358255602094850194600190920191016111c3565b50868210156112005760001960f88860031b161c19848701351681555b505060018560011b0183555b5050505050565b8183823760009101908152919050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006001820161126157611261611239565b5060010190565b60408152600061127b6040830185610f14565b905060018060a01b03831660208301529392505050565b815167ffffffffffffffff8111156112ac576112ac610dbd565b6112c0816112ba84546110c9565b84611103565b602080601f8311600181146112f557600084156112dd5750858301515b600019600386901b1c1916600185901b178555611149565b600085815260208120601f198616915b8281101561132457888601518255948401946001909101908401611305565b50858210156113425787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6060815260006113656060830186610f14565b6020830194909452506001600160a01b0391909116604090910152919050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6060815260006113c16060830188610f14565b82810360208401526113d4818789611385565b905082810360408401526113e9818587611385565b98975050505050505050565b60ff818116838216019081111561140e5761140e611239565b9291505056fea26469706673582212200278c0b919e48c61a64f25640c37ff04956267f7806e0d8a5de0cc6702d9850f64736f6c63430008110033",
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

// LookupConnector is a free data retrieval call binding the contract method 0xd346e3bd.
//
// Solidity: function lookupConnector(string name, string protocol) view returns(string)
func (_NameRegistry *NameRegistryCaller) LookupConnector(opts *bind.CallOpts, name string, protocol string) (string, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "lookupConnector", name, protocol)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LookupConnector is a free data retrieval call binding the contract method 0xd346e3bd.
//
// Solidity: function lookupConnector(string name, string protocol) view returns(string)
func (_NameRegistry *NameRegistrySession) LookupConnector(name string, protocol string) (string, error) {
	return _NameRegistry.Contract.LookupConnector(&_NameRegistry.CallOpts, name, protocol)
}

// LookupConnector is a free data retrieval call binding the contract method 0xd346e3bd.
//
// Solidity: function lookupConnector(string name, string protocol) view returns(string)
func (_NameRegistry *NameRegistryCallerSession) LookupConnector(name string, protocol string) (string, error) {
	return _NameRegistry.Contract.LookupConnector(&_NameRegistry.CallOpts, name, protocol)
}

// LookupName is a free data retrieval call binding the contract method 0x639cdb95.
//
// Solidity: function lookupName(string name) view returns(address owner, bytes publicKey, uint256 price)
func (_NameRegistry *NameRegistryCaller) LookupName(opts *bind.CallOpts, name string) (struct {
	Owner     common.Address
	PublicKey []byte
	Price     *big.Int
}, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "lookupName", name)

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

// LookupName is a free data retrieval call binding the contract method 0x639cdb95.
//
// Solidity: function lookupName(string name) view returns(address owner, bytes publicKey, uint256 price)
func (_NameRegistry *NameRegistrySession) LookupName(name string) (struct {
	Owner     common.Address
	PublicKey []byte
	Price     *big.Int
}, error) {
	return _NameRegistry.Contract.LookupName(&_NameRegistry.CallOpts, name)
}

// LookupName is a free data retrieval call binding the contract method 0x639cdb95.
//
// Solidity: function lookupName(string name) view returns(address owner, bytes publicKey, uint256 price)
func (_NameRegistry *NameRegistryCallerSession) LookupName(name string) (struct {
	Owner     common.Address
	PublicKey []byte
	Price     *big.Int
}, error) {
	return _NameRegistry.Contract.LookupName(&_NameRegistry.CallOpts, name)
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

// RegisterConnector is a paid mutator transaction binding the contract method 0xf11a1a9a.
//
// Solidity: function registerConnector(string name, string protocol, string location) returns()
func (_NameRegistry *NameRegistryTransactor) RegisterConnector(opts *bind.TransactOpts, name string, protocol string, location string) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "registerConnector", name, protocol, location)
}

// RegisterConnector is a paid mutator transaction binding the contract method 0xf11a1a9a.
//
// Solidity: function registerConnector(string name, string protocol, string location) returns()
func (_NameRegistry *NameRegistrySession) RegisterConnector(name string, protocol string, location string) (*types.Transaction, error) {
	return _NameRegistry.Contract.RegisterConnector(&_NameRegistry.TransactOpts, name, protocol, location)
}

// RegisterConnector is a paid mutator transaction binding the contract method 0xf11a1a9a.
//
// Solidity: function registerConnector(string name, string protocol, string location) returns()
func (_NameRegistry *NameRegistryTransactorSession) RegisterConnector(name string, protocol string, location string) (*types.Transaction, error) {
	return _NameRegistry.Contract.RegisterConnector(&_NameRegistry.TransactOpts, name, protocol, location)
}

// RegisterName is a paid mutator transaction binding the contract method 0x56c792db.
//
// Solidity: function registerName(bytes publicKey, string name) returns()
func (_NameRegistry *NameRegistryTransactor) RegisterName(opts *bind.TransactOpts, publicKey []byte, name string) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "registerName", publicKey, name)
}

// RegisterName is a paid mutator transaction binding the contract method 0x56c792db.
//
// Solidity: function registerName(bytes publicKey, string name) returns()
func (_NameRegistry *NameRegistrySession) RegisterName(publicKey []byte, name string) (*types.Transaction, error) {
	return _NameRegistry.Contract.RegisterName(&_NameRegistry.TransactOpts, publicKey, name)
}

// RegisterName is a paid mutator transaction binding the contract method 0x56c792db.
//
// Solidity: function registerName(bytes publicKey, string name) returns()
func (_NameRegistry *NameRegistryTransactorSession) RegisterName(publicKey []byte, name string) (*types.Transaction, error) {
	return _NameRegistry.Contract.RegisterName(&_NameRegistry.TransactOpts, publicKey, name)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xc0e793c2.
//
// Solidity: function transferOwnership(string name, address newOwner) returns()
func (_NameRegistry *NameRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, name string, newOwner common.Address) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "transferOwnership", name, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xc0e793c2.
//
// Solidity: function transferOwnership(string name, address newOwner) returns()
func (_NameRegistry *NameRegistrySession) TransferOwnership(name string, newOwner common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.TransferOwnership(&_NameRegistry.TransactOpts, name, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xc0e793c2.
//
// Solidity: function transferOwnership(string name, address newOwner) returns()
func (_NameRegistry *NameRegistryTransactorSession) TransferOwnership(name string, newOwner common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.TransferOwnership(&_NameRegistry.TransactOpts, name, newOwner)
}

// UpdatePublicKey is a paid mutator transaction binding the contract method 0x208b5f38.
//
// Solidity: function updatePublicKey(bytes publicKey, string name) returns()
func (_NameRegistry *NameRegistryTransactor) UpdatePublicKey(opts *bind.TransactOpts, publicKey []byte, name string) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "updatePublicKey", publicKey, name)
}

// UpdatePublicKey is a paid mutator transaction binding the contract method 0x208b5f38.
//
// Solidity: function updatePublicKey(bytes publicKey, string name) returns()
func (_NameRegistry *NameRegistrySession) UpdatePublicKey(publicKey []byte, name string) (*types.Transaction, error) {
	return _NameRegistry.Contract.UpdatePublicKey(&_NameRegistry.TransactOpts, publicKey, name)
}

// UpdatePublicKey is a paid mutator transaction binding the contract method 0x208b5f38.
//
// Solidity: function updatePublicKey(bytes publicKey, string name) returns()
func (_NameRegistry *NameRegistryTransactorSession) UpdatePublicKey(publicKey []byte, name string) (*types.Transaction, error) {
	return _NameRegistry.Contract.UpdatePublicKey(&_NameRegistry.TransactOpts, publicKey, name)
}

// NameRegistryConnectorRegisteredIterator is returned from FilterConnectorRegistered and is used to iterate over the raw logs and unpacked data for ConnectorRegistered events raised by the NameRegistry contract.
type NameRegistryConnectorRegisteredIterator struct {
	Event *NameRegistryConnectorRegistered // Event containing the contract specifics and raw log

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
func (it *NameRegistryConnectorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryConnectorRegistered)
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
		it.Event = new(NameRegistryConnectorRegistered)
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
func (it *NameRegistryConnectorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryConnectorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryConnectorRegistered represents a ConnectorRegistered event raised by the NameRegistry contract.
type NameRegistryConnectorRegistered struct {
	Name     string
	Protocol string
	Location string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterConnectorRegistered is a free log retrieval operation binding the contract event 0x6bcbdc9c6b89c8a6b289cc4308c6f6841a0e25d81c5fe34b1dedb2848c5f5ca3.
//
// Solidity: event ConnectorRegistered(string name, string protocol, string location)
func (_NameRegistry *NameRegistryFilterer) FilterConnectorRegistered(opts *bind.FilterOpts) (*NameRegistryConnectorRegisteredIterator, error) {

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "ConnectorRegistered")
	if err != nil {
		return nil, err
	}
	return &NameRegistryConnectorRegisteredIterator{contract: _NameRegistry.contract, event: "ConnectorRegistered", logs: logs, sub: sub}, nil
}

// WatchConnectorRegistered is a free log subscription operation binding the contract event 0x6bcbdc9c6b89c8a6b289cc4308c6f6841a0e25d81c5fe34b1dedb2848c5f5ca3.
//
// Solidity: event ConnectorRegistered(string name, string protocol, string location)
func (_NameRegistry *NameRegistryFilterer) WatchConnectorRegistered(opts *bind.WatchOpts, sink chan<- *NameRegistryConnectorRegistered) (event.Subscription, error) {

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "ConnectorRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryConnectorRegistered)
				if err := _NameRegistry.contract.UnpackLog(event, "ConnectorRegistered", log); err != nil {
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
// Solidity: event ConnectorRegistered(string name, string protocol, string location)
func (_NameRegistry *NameRegistryFilterer) ParseConnectorRegistered(log types.Log) (*NameRegistryConnectorRegistered, error) {
	event := new(NameRegistryConnectorRegistered)
	if err := _NameRegistry.contract.UnpackLog(event, "ConnectorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistryNameOwnershipChangedIterator is returned from FilterNameOwnershipChanged and is used to iterate over the raw logs and unpacked data for NameOwnershipChanged events raised by the NameRegistry contract.
type NameRegistryNameOwnershipChangedIterator struct {
	Event *NameRegistryNameOwnershipChanged // Event containing the contract specifics and raw log

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
func (it *NameRegistryNameOwnershipChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryNameOwnershipChanged)
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
		it.Event = new(NameRegistryNameOwnershipChanged)
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
func (it *NameRegistryNameOwnershipChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryNameOwnershipChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryNameOwnershipChanged represents a NameOwnershipChanged event raised by the NameRegistry contract.
type NameRegistryNameOwnershipChanged struct {
	Name     string
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNameOwnershipChanged is a free log retrieval operation binding the contract event 0x8148cdc9d89df5e18d29f32adde657a148d30c0da1546f561cd58739a2057d02.
//
// Solidity: event NameOwnershipChanged(string name, address newOwner)
func (_NameRegistry *NameRegistryFilterer) FilterNameOwnershipChanged(opts *bind.FilterOpts) (*NameRegistryNameOwnershipChangedIterator, error) {

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "NameOwnershipChanged")
	if err != nil {
		return nil, err
	}
	return &NameRegistryNameOwnershipChangedIterator{contract: _NameRegistry.contract, event: "NameOwnershipChanged", logs: logs, sub: sub}, nil
}

// WatchNameOwnershipChanged is a free log subscription operation binding the contract event 0x8148cdc9d89df5e18d29f32adde657a148d30c0da1546f561cd58739a2057d02.
//
// Solidity: event NameOwnershipChanged(string name, address newOwner)
func (_NameRegistry *NameRegistryFilterer) WatchNameOwnershipChanged(opts *bind.WatchOpts, sink chan<- *NameRegistryNameOwnershipChanged) (event.Subscription, error) {

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "NameOwnershipChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryNameOwnershipChanged)
				if err := _NameRegistry.contract.UnpackLog(event, "NameOwnershipChanged", log); err != nil {
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

// ParseNameOwnershipChanged is a log parse operation binding the contract event 0x8148cdc9d89df5e18d29f32adde657a148d30c0da1546f561cd58739a2057d02.
//
// Solidity: event NameOwnershipChanged(string name, address newOwner)
func (_NameRegistry *NameRegistryFilterer) ParseNameOwnershipChanged(log types.Log) (*NameRegistryNameOwnershipChanged, error) {
	event := new(NameRegistryNameOwnershipChanged)
	if err := _NameRegistry.contract.UnpackLog(event, "NameOwnershipChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	Name  string
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNameRegistered is a free log retrieval operation binding the contract event 0x1c6eac0e720ec22bb0653aec9c19985633a4fb07971cf973096c2f8e3c37c17f.
//
// Solidity: event NameRegistered(string name, address owner)
func (_NameRegistry *NameRegistryFilterer) FilterNameRegistered(opts *bind.FilterOpts) (*NameRegistryNameRegisteredIterator, error) {

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "NameRegistered")
	if err != nil {
		return nil, err
	}
	return &NameRegistryNameRegisteredIterator{contract: _NameRegistry.contract, event: "NameRegistered", logs: logs, sub: sub}, nil
}

// WatchNameRegistered is a free log subscription operation binding the contract event 0x1c6eac0e720ec22bb0653aec9c19985633a4fb07971cf973096c2f8e3c37c17f.
//
// Solidity: event NameRegistered(string name, address owner)
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

// ParseNameRegistered is a log parse operation binding the contract event 0x1c6eac0e720ec22bb0653aec9c19985633a4fb07971cf973096c2f8e3c37c17f.
//
// Solidity: event NameRegistered(string name, address owner)
func (_NameRegistry *NameRegistryFilterer) ParseNameRegistered(log types.Log) (*NameRegistryNameRegistered, error) {
	event := new(NameRegistryNameRegistered)
	if err := _NameRegistry.contract.UnpackLog(event, "NameRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistryPriceChangedIterator is returned from FilterPriceChanged and is used to iterate over the raw logs and unpacked data for PriceChanged events raised by the NameRegistry contract.
type NameRegistryPriceChangedIterator struct {
	Event *NameRegistryPriceChanged // Event containing the contract specifics and raw log

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
func (it *NameRegistryPriceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryPriceChanged)
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
		it.Event = new(NameRegistryPriceChanged)
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
func (it *NameRegistryPriceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryPriceChanged represents a PriceChanged event raised by the NameRegistry contract.
type NameRegistryPriceChanged struct {
	Name  string
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPriceChanged is a free log retrieval operation binding the contract event 0x29ecd1e0988af1492e43256007437db411881757b3e6e808c9f04847b264178c.
//
// Solidity: event PriceChanged(string name, uint256 price)
func (_NameRegistry *NameRegistryFilterer) FilterPriceChanged(opts *bind.FilterOpts) (*NameRegistryPriceChangedIterator, error) {

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "PriceChanged")
	if err != nil {
		return nil, err
	}
	return &NameRegistryPriceChangedIterator{contract: _NameRegistry.contract, event: "PriceChanged", logs: logs, sub: sub}, nil
}

// WatchPriceChanged is a free log subscription operation binding the contract event 0x29ecd1e0988af1492e43256007437db411881757b3e6e808c9f04847b264178c.
//
// Solidity: event PriceChanged(string name, uint256 price)
func (_NameRegistry *NameRegistryFilterer) WatchPriceChanged(opts *bind.WatchOpts, sink chan<- *NameRegistryPriceChanged) (event.Subscription, error) {

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "PriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryPriceChanged)
				if err := _NameRegistry.contract.UnpackLog(event, "PriceChanged", log); err != nil {
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

// ParsePriceChanged is a log parse operation binding the contract event 0x29ecd1e0988af1492e43256007437db411881757b3e6e808c9f04847b264178c.
//
// Solidity: event PriceChanged(string name, uint256 price)
func (_NameRegistry *NameRegistryFilterer) ParsePriceChanged(log types.Log) (*NameRegistryPriceChanged, error) {
	event := new(NameRegistryPriceChanged)
	if err := _NameRegistry.contract.UnpackLog(event, "PriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistryPublicKeyUpdatedIterator is returned from FilterPublicKeyUpdated and is used to iterate over the raw logs and unpacked data for PublicKeyUpdated events raised by the NameRegistry contract.
type NameRegistryPublicKeyUpdatedIterator struct {
	Event *NameRegistryPublicKeyUpdated // Event containing the contract specifics and raw log

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
func (it *NameRegistryPublicKeyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryPublicKeyUpdated)
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
		it.Event = new(NameRegistryPublicKeyUpdated)
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
func (it *NameRegistryPublicKeyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryPublicKeyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryPublicKeyUpdated represents a PublicKeyUpdated event raised by the NameRegistry contract.
type NameRegistryPublicKeyUpdated struct {
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPublicKeyUpdated is a free log retrieval operation binding the contract event 0x0ba4a7233a2cfb937aa9644f11f49345b443fe8d5ab3494a1879faa500728d3e.
//
// Solidity: event PublicKeyUpdated(string name)
func (_NameRegistry *NameRegistryFilterer) FilterPublicKeyUpdated(opts *bind.FilterOpts) (*NameRegistryPublicKeyUpdatedIterator, error) {

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "PublicKeyUpdated")
	if err != nil {
		return nil, err
	}
	return &NameRegistryPublicKeyUpdatedIterator{contract: _NameRegistry.contract, event: "PublicKeyUpdated", logs: logs, sub: sub}, nil
}

// WatchPublicKeyUpdated is a free log subscription operation binding the contract event 0x0ba4a7233a2cfb937aa9644f11f49345b443fe8d5ab3494a1879faa500728d3e.
//
// Solidity: event PublicKeyUpdated(string name)
func (_NameRegistry *NameRegistryFilterer) WatchPublicKeyUpdated(opts *bind.WatchOpts, sink chan<- *NameRegistryPublicKeyUpdated) (event.Subscription, error) {

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "PublicKeyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryPublicKeyUpdated)
				if err := _NameRegistry.contract.UnpackLog(event, "PublicKeyUpdated", log); err != nil {
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

// ParsePublicKeyUpdated is a log parse operation binding the contract event 0x0ba4a7233a2cfb937aa9644f11f49345b443fe8d5ab3494a1879faa500728d3e.
//
// Solidity: event PublicKeyUpdated(string name)
func (_NameRegistry *NameRegistryFilterer) ParsePublicKeyUpdated(log types.Log) (*NameRegistryPublicKeyUpdated, error) {
	event := new(NameRegistryPublicKeyUpdated)
	if err := _NameRegistry.contract.UnpackLog(event, "PublicKeyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistrySaleIterator is returned from FilterSale and is used to iterate over the raw logs and unpacked data for Sale events raised by the NameRegistry contract.
type NameRegistrySaleIterator struct {
	Event *NameRegistrySale // Event containing the contract specifics and raw log

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
func (it *NameRegistrySaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistrySale)
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
		it.Event = new(NameRegistrySale)
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
func (it *NameRegistrySaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistrySaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistrySale represents a Sale event raised by the NameRegistry contract.
type NameRegistrySale struct {
	Name     string
	Price    *big.Int
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSale is a free log retrieval operation binding the contract event 0x74d699b6db148dc632a2dbe2c680980f546b0382af2e71a7d4bf51c3193adb9a.
//
// Solidity: event Sale(string name, uint256 price, address newOwner)
func (_NameRegistry *NameRegistryFilterer) FilterSale(opts *bind.FilterOpts) (*NameRegistrySaleIterator, error) {

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "Sale")
	if err != nil {
		return nil, err
	}
	return &NameRegistrySaleIterator{contract: _NameRegistry.contract, event: "Sale", logs: logs, sub: sub}, nil
}

// WatchSale is a free log subscription operation binding the contract event 0x74d699b6db148dc632a2dbe2c680980f546b0382af2e71a7d4bf51c3193adb9a.
//
// Solidity: event Sale(string name, uint256 price, address newOwner)
func (_NameRegistry *NameRegistryFilterer) WatchSale(opts *bind.WatchOpts, sink chan<- *NameRegistrySale) (event.Subscription, error) {

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "Sale")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistrySale)
				if err := _NameRegistry.contract.UnpackLog(event, "Sale", log); err != nil {
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

// ParseSale is a log parse operation binding the contract event 0x74d699b6db148dc632a2dbe2c680980f546b0382af2e71a7d4bf51c3193adb9a.
//
// Solidity: event Sale(string name, uint256 price, address newOwner)
func (_NameRegistry *NameRegistryFilterer) ParseSale(log types.Log) (*NameRegistrySale, error) {
	event := new(NameRegistrySale)
	if err := _NameRegistry.contract.UnpackLog(event, "Sale", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
