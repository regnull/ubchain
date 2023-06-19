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

// NameRegistryMetaData contains all meta data concerning the NameRegistry contract.
var NameRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"configName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"configValue\",\"type\":\"string\"}],\"name\":\"ConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NameOwnershipUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NameRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"PriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"PublicKeyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"Sale\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"buyName\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"configName\",\"type\":\"string\"}],\"name\":\"lookupConfig\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"lookupName\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"registerName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"configName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"configValue\",\"type\":\"string\"}],\"name\":\"updateConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"updateOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"updatePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"updatePublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506114558061001d5f395ff3fe608060405260043610610079575f3560e01c806362eaff7d1161004c57806362eaff7d14610111578063639cdb9514610124578063971c728714610152578063a3ddd69114610171575f80fd5b8063208b5f381461007d5780633aabd19c1461009e5780634a432a46146100d357806356c792db146100f2575b5f80fd5b348015610088575f80fd5b5061009c610097366004610d51565b610190565b005b3480156100a9575f80fd5b506100bd6100b8366004610d51565b61029c565b6040516100ca9190610dfb565b60405180910390f35b3480156100de575f80fd5b5061009c6100ed366004610e14565b61036d565b3480156100fd575f80fd5b5061009c61010c366004610d51565b610428565b61009c61011f366004610e70565b61087e565b34801561012f575f80fd5b5061014361013e366004610f43565b610a8d565b6040516100ca93929190610f82565b34801561015d575f80fd5b5061009c61016c366004610fb5565b610b67565b34801561017c575f80fd5b5061009c61018b366004611048565b610c42565b602183146101da5760405162461bcd60e51b8152602060048201526012602482015271696e76616c6964207075626c6963206b657960701b60448201526064015b60405180910390fd5b336001600160a01b03165f83836040516101f59291906110a6565b908152604051908190036020019020600101546001600160a01b03161461022e5760405162461bcd60e51b81526004016101d1906110b5565b83835f84846040516102419291906110a6565b9081526040519081900360200190209161025c919083611162565b507f0ba4a7233a2cfb937aa9644f11f49345b443fe8d5ab3494a1879faa500728d3e828260405161028e929190611246565b60405180910390a150505050565b6060600185856040516102b09291906110a6565b908152602001604051809103902083836040516102ce9291906110a6565b908152602001604051809103902080546102e7906110dc565b80601f0160208091040260200160405190810160405280929190818152602001828054610313906110dc565b801561035e5780601f106103355761010080835404028352916020019161035e565b820191905f5260205f20905b81548152906001019060200180831161034157829003601f168201915b50505050509050949350505050565b336001600160a01b03165f84846040516103889291906110a6565b908152604051908190036020019020600101546001600160a01b0316146103c15760405162461bcd60e51b81526004016101d1906110b5565b805f84846040516103d39291906110a6565b9081526020016040518091039020600201819055507f159e83f4712ba2552e68be9d848e49bf6dd35c24f19564ffd523b6549450a2f483838360405161041b93929190611261565b60405180910390a1505050565b6021831461046d5760405162461bcd60e51b8152602060048201526012602482015271696e76616c6964207075626c6963206b657960701b60448201526064016101d1565b5f6001600160a01b03165f83836040516104889291906110a6565b908152604051908190036020019020600101546001600160a01b0316146104f15760405162461bcd60e51b815260206004820152601a60248201527f6e616d6520697320616c7265616479207265676973746572656400000000000060448201526064016101d1565b60018110156105385760405162461bcd60e51b81526020600482015260136024820152726e616d652063616e277420626520656d70747960681b60448201526064016101d1565b604081111561057c5760405162461bcd60e51b815260206004820152601060248201526f6e616d6520697320746f6f206c6f6e6760801b60448201526064016101d1565b81815f81811061058e5761058e611284565b909101356001600160f81b031916602d60f81b0390506105f05760405162461bcd60e51b815260206004820152601c60248201527f6e616d652063616e27742073746172742077697468206120646173680000000060448201526064016101d1565b5f82828080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920182905250939450600193925050505b8251811015610760575f83828151811061064957610649611284565b01602001516001600160f81b0319169050602d60f81b8114806106935750600360fc1b6001600160f81b03198216108015906106935750603960f81b6001600160f81b0319821611155b806106ab5750605f60f81b6001600160f81b03198216145b806106dd5750606160f81b6001600160f81b03198216108015906106dd5750603d60f91b6001600160f81b0319821611155b6107185760405162461bcd60e51b815260206004820152600c60248201526b696e76616c6964206e616d6560a01b60448201526064016101d1565b602d60f81b6001600160f81b03198216148015906107445750605f60f81b6001600160f81b0319821614155b1561074d575f92505b508061075881611298565b91505061062d565b5080156107c65760405162461bcd60e51b815260206004820152602e60248201527f6e616d65206d75737420636f6e7461696e206174206c65617374206f6e65206c60448201526d195d1d195c881bdc88191a59da5d60921b60648201526084016101d1565b335f85856040516107d89291906110a6565b90815260200160405180910390206001015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555085855f868660405161081f9291906110a6565b9081526040519081900360200190209161083a919083611162565b507f1c6eac0e720ec22bb0653aec9c19985633a4fb07971cf973096c2f8e3c37c17f84843360405161086e939291906112bc565b60405180910390a1505050505050565b5f8084846040516108909291906110a6565b908152602001604051809103902060020154116108de5760405162461bcd60e51b815260206004820152600c60248201526b6e6f7420666f722073616c6560a01b60448201526064016101d1565b345f84846040516108f09291906110a6565b90815260200160405180910390206002015411156109455760405162461bcd60e51b8152602060048201526012602482015271696e73756666696369656e742076616c756560701b60448201526064016101d1565b5f8084846040516109579291906110a6565b90815260200160405180910390206002015490505f848460405161097c9291906110a6565b908152604051908190036020018120600101546001600160a01b0316903480156108fc02915f818181858888f193505050501580156109bd573d5f803e3d5ffd5b50335f85856040516109d09291906110a6565b90815260200160405180910390206001015f6101000a8154816001600160a01b0302191690836001600160a01b03160217905550815f8585604051610a169291906110a6565b90815260405190819003602001902090610a3090826112e7565b505f808585604051610a439291906110a6565b9081526020016040518091039020600201819055507f74d699b6db148dc632a2dbe2c680980f546b0382af2e71a7d4bf51c3193adb9a8484833360405161028e94939291906113a3565b5f60605f805f8686604051610aa39291906110a6565b9081526040519081900360200190206001810154600282015482549293506001600160a01b03909116918391908290610adb906110dc565b80601f0160208091040260200160405190810160405280929190818152602001828054610b07906110dc565b8015610b525780601f10610b2957610100808354040283529160200191610b52565b820191905f5260205f20905b815481529060010190602001808311610b3557829003601f168201915b50505050509150935093509350509250925092565b336001600160a01b03165f8787604051610b829291906110a6565b908152604051908190036020019020600101546001600160a01b031614610bbb5760405162461bcd60e51b81526004016101d1906110b5565b818160018888604051610bcf9291906110a6565b90815260200160405180910390208686604051610bed9291906110a6565b90815260200160405180910390209182610c08929190611162565b507fcde50e1bbc8495f4d015791042ac8d9b4e45d1cd60159e1fbad5863ee388d82886868686868660405161086e969594939291906113d7565b336001600160a01b03165f8484604051610c5d9291906110a6565b908152604051908190036020019020600101546001600160a01b031614610c965760405162461bcd60e51b81526004016101d1906110b5565b805f8484604051610ca89291906110a6565b90815260405190819003602001812060010180546001600160a01b03939093166001600160a01b0319909316929092179091557f179b5e07786d37fc59a40f1e1e8a697240c6e378cf1e17cf9438b0110abeb6d79061041b908590859085906112bc565b5f8083601f840112610d1c575f80fd5b50813567ffffffffffffffff811115610d33575f80fd5b602083019150836020828501011115610d4a575f80fd5b9250929050565b5f805f8060408587031215610d64575f80fd5b843567ffffffffffffffff80821115610d7b575f80fd5b610d8788838901610d0c565b90965094506020870135915080821115610d9f575f80fd5b50610dac87828801610d0c565b95989497509550505050565b5f81518084525f5b81811015610ddc57602081850181015186830182015201610dc0565b505f602082860101526020601f19601f83011685010191505092915050565b602081525f610e0d6020830184610db8565b9392505050565b5f805f60408486031215610e26575f80fd5b833567ffffffffffffffff811115610e3c575f80fd5b610e4886828701610d0c565b909790965060209590950135949350505050565b634e487b7160e01b5f52604160045260245ffd5b5f805f60408486031215610e82575f80fd5b833567ffffffffffffffff80821115610e99575f80fd5b610ea587838801610d0c565b90955093506020860135915080821115610ebd575f80fd5b818601915086601f830112610ed0575f80fd5b813581811115610ee257610ee2610e5c565b604051601f8201601f19908116603f01168101908382118183101715610f0a57610f0a610e5c565b81604052828152896020848701011115610f22575f80fd5b826020860160208301375f6020848301015280955050505050509250925092565b5f8060208385031215610f54575f80fd5b823567ffffffffffffffff811115610f6a575f80fd5b610f7685828601610d0c565b90969095509350505050565b6001600160a01b03841681526060602082018190525f90610fa590830185610db8565b9050826040830152949350505050565b5f805f805f8060608789031215610fca575f80fd5b863567ffffffffffffffff80821115610fe1575f80fd5b610fed8a838b01610d0c565b90985096506020890135915080821115611005575f80fd5b6110118a838b01610d0c565b90965094506040890135915080821115611029575f80fd5b5061103689828a01610d0c565b979a9699509497509295939492505050565b5f805f6040848603121561105a575f80fd5b833567ffffffffffffffff811115611070575f80fd5b61107c86828701610d0c565b90945092505060208401356001600160a01b038116811461109b575f80fd5b809150509250925092565b818382375f9101908152919050565b6020808252600d908201526c3737ba103a34329037bbb732b960991b604082015260600190565b600181811c908216806110f057607f821691505b60208210810361110e57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561115d575f81815260208120601f850160051c8101602086101561113a5750805b601f850160051c820191505b8181101561115957828155600101611146565b5050505b505050565b67ffffffffffffffff83111561117a5761117a610e5c565b61118e8361118883546110dc565b83611114565b5f601f8411600181146111bf575f85156111a85750838201355b5f19600387901b1c1916600186901b178355611217565b5f83815260209020601f19861690835b828110156111ef57868501358255602094850194600190920191016111cf565b508682101561120b575f1960f88860031b161c19848701351681555b505060018560011b0183555b5050505050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b602081525f61125960208301848661121e565b949350505050565b604081525f61127460408301858761121e565b9050826020830152949350505050565b634e487b7160e01b5f52603260045260245ffd5b5f600182016112b557634e487b7160e01b5f52601160045260245ffd5b5060010190565b604081525f6112cf60408301858761121e565b905060018060a01b0383166020830152949350505050565b815167ffffffffffffffff81111561130157611301610e5c565b6113158161130f84546110dc565b84611114565b602080601f831160018114611348575f84156113315750858301515b5f19600386901b1c1916600185901b178555611159565b5f85815260208120601f198616915b8281101561137657888601518255948401946001909101908401611357565b508582101561139357878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b606081525f6113b660608301868861121e565b6020830194909452506001600160a01b039190911660409091015292915050565b606081525f6113ea60608301888a61121e565b82810360208401526113fd81878961121e565b9050828103604084015261141281858761121e565b999850505050505050505056fea2646970667358221220ffed26ce9144f81bde129a7b2c5fd56f8ce000843c7f03a48b4e580943fa68da64736f6c63430008140033",
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

// LookupConfig is a free data retrieval call binding the contract method 0x3aabd19c.
//
// Solidity: function lookupConfig(string name, string configName) view returns(string)
func (_NameRegistry *NameRegistryCaller) LookupConfig(opts *bind.CallOpts, name string, configName string) (string, error) {
	var out []interface{}
	err := _NameRegistry.contract.Call(opts, &out, "lookupConfig", name, configName)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LookupConfig is a free data retrieval call binding the contract method 0x3aabd19c.
//
// Solidity: function lookupConfig(string name, string configName) view returns(string)
func (_NameRegistry *NameRegistrySession) LookupConfig(name string, configName string) (string, error) {
	return _NameRegistry.Contract.LookupConfig(&_NameRegistry.CallOpts, name, configName)
}

// LookupConfig is a free data retrieval call binding the contract method 0x3aabd19c.
//
// Solidity: function lookupConfig(string name, string configName) view returns(string)
func (_NameRegistry *NameRegistryCallerSession) LookupConfig(name string, configName string) (string, error) {
	return _NameRegistry.Contract.LookupConfig(&_NameRegistry.CallOpts, name, configName)
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

// BuyName is a paid mutator transaction binding the contract method 0x62eaff7d.
//
// Solidity: function buyName(string name, bytes publicKey) payable returns()
func (_NameRegistry *NameRegistryTransactor) BuyName(opts *bind.TransactOpts, name string, publicKey []byte) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "buyName", name, publicKey)
}

// BuyName is a paid mutator transaction binding the contract method 0x62eaff7d.
//
// Solidity: function buyName(string name, bytes publicKey) payable returns()
func (_NameRegistry *NameRegistrySession) BuyName(name string, publicKey []byte) (*types.Transaction, error) {
	return _NameRegistry.Contract.BuyName(&_NameRegistry.TransactOpts, name, publicKey)
}

// BuyName is a paid mutator transaction binding the contract method 0x62eaff7d.
//
// Solidity: function buyName(string name, bytes publicKey) payable returns()
func (_NameRegistry *NameRegistryTransactorSession) BuyName(name string, publicKey []byte) (*types.Transaction, error) {
	return _NameRegistry.Contract.BuyName(&_NameRegistry.TransactOpts, name, publicKey)
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

// UpdateConfig is a paid mutator transaction binding the contract method 0x971c7287.
//
// Solidity: function updateConfig(string name, string configName, string configValue) returns()
func (_NameRegistry *NameRegistryTransactor) UpdateConfig(opts *bind.TransactOpts, name string, configName string, configValue string) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "updateConfig", name, configName, configValue)
}

// UpdateConfig is a paid mutator transaction binding the contract method 0x971c7287.
//
// Solidity: function updateConfig(string name, string configName, string configValue) returns()
func (_NameRegistry *NameRegistrySession) UpdateConfig(name string, configName string, configValue string) (*types.Transaction, error) {
	return _NameRegistry.Contract.UpdateConfig(&_NameRegistry.TransactOpts, name, configName, configValue)
}

// UpdateConfig is a paid mutator transaction binding the contract method 0x971c7287.
//
// Solidity: function updateConfig(string name, string configName, string configValue) returns()
func (_NameRegistry *NameRegistryTransactorSession) UpdateConfig(name string, configName string, configValue string) (*types.Transaction, error) {
	return _NameRegistry.Contract.UpdateConfig(&_NameRegistry.TransactOpts, name, configName, configValue)
}

// UpdateOwnership is a paid mutator transaction binding the contract method 0xa3ddd691.
//
// Solidity: function updateOwnership(string name, address newOwner) returns()
func (_NameRegistry *NameRegistryTransactor) UpdateOwnership(opts *bind.TransactOpts, name string, newOwner common.Address) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "updateOwnership", name, newOwner)
}

// UpdateOwnership is a paid mutator transaction binding the contract method 0xa3ddd691.
//
// Solidity: function updateOwnership(string name, address newOwner) returns()
func (_NameRegistry *NameRegistrySession) UpdateOwnership(name string, newOwner common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.UpdateOwnership(&_NameRegistry.TransactOpts, name, newOwner)
}

// UpdateOwnership is a paid mutator transaction binding the contract method 0xa3ddd691.
//
// Solidity: function updateOwnership(string name, address newOwner) returns()
func (_NameRegistry *NameRegistryTransactorSession) UpdateOwnership(name string, newOwner common.Address) (*types.Transaction, error) {
	return _NameRegistry.Contract.UpdateOwnership(&_NameRegistry.TransactOpts, name, newOwner)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x4a432a46.
//
// Solidity: function updatePrice(string name, uint256 price) returns()
func (_NameRegistry *NameRegistryTransactor) UpdatePrice(opts *bind.TransactOpts, name string, price *big.Int) (*types.Transaction, error) {
	return _NameRegistry.contract.Transact(opts, "updatePrice", name, price)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x4a432a46.
//
// Solidity: function updatePrice(string name, uint256 price) returns()
func (_NameRegistry *NameRegistrySession) UpdatePrice(name string, price *big.Int) (*types.Transaction, error) {
	return _NameRegistry.Contract.UpdatePrice(&_NameRegistry.TransactOpts, name, price)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x4a432a46.
//
// Solidity: function updatePrice(string name, uint256 price) returns()
func (_NameRegistry *NameRegistryTransactorSession) UpdatePrice(name string, price *big.Int) (*types.Transaction, error) {
	return _NameRegistry.Contract.UpdatePrice(&_NameRegistry.TransactOpts, name, price)
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

// NameRegistryConfigUpdatedIterator is returned from FilterConfigUpdated and is used to iterate over the raw logs and unpacked data for ConfigUpdated events raised by the NameRegistry contract.
type NameRegistryConfigUpdatedIterator struct {
	Event *NameRegistryConfigUpdated // Event containing the contract specifics and raw log

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
func (it *NameRegistryConfigUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryConfigUpdated)
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
		it.Event = new(NameRegistryConfigUpdated)
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
func (it *NameRegistryConfigUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryConfigUpdated represents a ConfigUpdated event raised by the NameRegistry contract.
type NameRegistryConfigUpdated struct {
	Name        string
	ConfigName  string
	ConfigValue string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConfigUpdated is a free log retrieval operation binding the contract event 0xcde50e1bbc8495f4d015791042ac8d9b4e45d1cd60159e1fbad5863ee388d828.
//
// Solidity: event ConfigUpdated(string name, string configName, string configValue)
func (_NameRegistry *NameRegistryFilterer) FilterConfigUpdated(opts *bind.FilterOpts) (*NameRegistryConfigUpdatedIterator, error) {

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "ConfigUpdated")
	if err != nil {
		return nil, err
	}
	return &NameRegistryConfigUpdatedIterator{contract: _NameRegistry.contract, event: "ConfigUpdated", logs: logs, sub: sub}, nil
}

// WatchConfigUpdated is a free log subscription operation binding the contract event 0xcde50e1bbc8495f4d015791042ac8d9b4e45d1cd60159e1fbad5863ee388d828.
//
// Solidity: event ConfigUpdated(string name, string configName, string configValue)
func (_NameRegistry *NameRegistryFilterer) WatchConfigUpdated(opts *bind.WatchOpts, sink chan<- *NameRegistryConfigUpdated) (event.Subscription, error) {

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "ConfigUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryConfigUpdated)
				if err := _NameRegistry.contract.UnpackLog(event, "ConfigUpdated", log); err != nil {
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

// ParseConfigUpdated is a log parse operation binding the contract event 0xcde50e1bbc8495f4d015791042ac8d9b4e45d1cd60159e1fbad5863ee388d828.
//
// Solidity: event ConfigUpdated(string name, string configName, string configValue)
func (_NameRegistry *NameRegistryFilterer) ParseConfigUpdated(log types.Log) (*NameRegistryConfigUpdated, error) {
	event := new(NameRegistryConfigUpdated)
	if err := _NameRegistry.contract.UnpackLog(event, "ConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NameRegistryNameOwnershipUpdatedIterator is returned from FilterNameOwnershipUpdated and is used to iterate over the raw logs and unpacked data for NameOwnershipUpdated events raised by the NameRegistry contract.
type NameRegistryNameOwnershipUpdatedIterator struct {
	Event *NameRegistryNameOwnershipUpdated // Event containing the contract specifics and raw log

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
func (it *NameRegistryNameOwnershipUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryNameOwnershipUpdated)
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
		it.Event = new(NameRegistryNameOwnershipUpdated)
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
func (it *NameRegistryNameOwnershipUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryNameOwnershipUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryNameOwnershipUpdated represents a NameOwnershipUpdated event raised by the NameRegistry contract.
type NameRegistryNameOwnershipUpdated struct {
	Name     string
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNameOwnershipUpdated is a free log retrieval operation binding the contract event 0x179b5e07786d37fc59a40f1e1e8a697240c6e378cf1e17cf9438b0110abeb6d7.
//
// Solidity: event NameOwnershipUpdated(string name, address newOwner)
func (_NameRegistry *NameRegistryFilterer) FilterNameOwnershipUpdated(opts *bind.FilterOpts) (*NameRegistryNameOwnershipUpdatedIterator, error) {

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "NameOwnershipUpdated")
	if err != nil {
		return nil, err
	}
	return &NameRegistryNameOwnershipUpdatedIterator{contract: _NameRegistry.contract, event: "NameOwnershipUpdated", logs: logs, sub: sub}, nil
}

// WatchNameOwnershipUpdated is a free log subscription operation binding the contract event 0x179b5e07786d37fc59a40f1e1e8a697240c6e378cf1e17cf9438b0110abeb6d7.
//
// Solidity: event NameOwnershipUpdated(string name, address newOwner)
func (_NameRegistry *NameRegistryFilterer) WatchNameOwnershipUpdated(opts *bind.WatchOpts, sink chan<- *NameRegistryNameOwnershipUpdated) (event.Subscription, error) {

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "NameOwnershipUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryNameOwnershipUpdated)
				if err := _NameRegistry.contract.UnpackLog(event, "NameOwnershipUpdated", log); err != nil {
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

// ParseNameOwnershipUpdated is a log parse operation binding the contract event 0x179b5e07786d37fc59a40f1e1e8a697240c6e378cf1e17cf9438b0110abeb6d7.
//
// Solidity: event NameOwnershipUpdated(string name, address newOwner)
func (_NameRegistry *NameRegistryFilterer) ParseNameOwnershipUpdated(log types.Log) (*NameRegistryNameOwnershipUpdated, error) {
	event := new(NameRegistryNameOwnershipUpdated)
	if err := _NameRegistry.contract.UnpackLog(event, "NameOwnershipUpdated", log); err != nil {
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

// NameRegistryPriceUpdatedIterator is returned from FilterPriceUpdated and is used to iterate over the raw logs and unpacked data for PriceUpdated events raised by the NameRegistry contract.
type NameRegistryPriceUpdatedIterator struct {
	Event *NameRegistryPriceUpdated // Event containing the contract specifics and raw log

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
func (it *NameRegistryPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NameRegistryPriceUpdated)
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
		it.Event = new(NameRegistryPriceUpdated)
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
func (it *NameRegistryPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NameRegistryPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NameRegistryPriceUpdated represents a PriceUpdated event raised by the NameRegistry contract.
type NameRegistryPriceUpdated struct {
	Name  string
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPriceUpdated is a free log retrieval operation binding the contract event 0x159e83f4712ba2552e68be9d848e49bf6dd35c24f19564ffd523b6549450a2f4.
//
// Solidity: event PriceUpdated(string name, uint256 price)
func (_NameRegistry *NameRegistryFilterer) FilterPriceUpdated(opts *bind.FilterOpts) (*NameRegistryPriceUpdatedIterator, error) {

	logs, sub, err := _NameRegistry.contract.FilterLogs(opts, "PriceUpdated")
	if err != nil {
		return nil, err
	}
	return &NameRegistryPriceUpdatedIterator{contract: _NameRegistry.contract, event: "PriceUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceUpdated is a free log subscription operation binding the contract event 0x159e83f4712ba2552e68be9d848e49bf6dd35c24f19564ffd523b6549450a2f4.
//
// Solidity: event PriceUpdated(string name, uint256 price)
func (_NameRegistry *NameRegistryFilterer) WatchPriceUpdated(opts *bind.WatchOpts, sink chan<- *NameRegistryPriceUpdated) (event.Subscription, error) {

	logs, sub, err := _NameRegistry.contract.WatchLogs(opts, "PriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NameRegistryPriceUpdated)
				if err := _NameRegistry.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
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

// ParsePriceUpdated is a log parse operation binding the contract event 0x159e83f4712ba2552e68be9d848e49bf6dd35c24f19564ffd523b6549450a2f4.
//
// Solidity: event PriceUpdated(string name, uint256 price)
func (_NameRegistry *NameRegistryFilterer) ParsePriceUpdated(log types.Log) (*NameRegistryPriceUpdated, error) {
	event := new(NameRegistryPriceUpdated)
	if err := _NameRegistry.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
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
