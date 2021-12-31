package testutil

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/regnull/easyecc"
)

type SimulatedBlockchain struct {
	privateKey  *easyecc.PrivateKey
	privateKey1 *easyecc.PrivateKey
	privateKey2 *easyecc.PrivateKey
	auth        *bind.TransactOpts
	backend     *backends.SimulatedBackend
}

func NewSimulatedBlockchain() (*SimulatedBlockchain, error) {
	// Generate private key.
	privateKey, err := easyecc.NewRandomPrivateKey()
	if err != nil {
		return nil, err
	}

	privateKey1, err := easyecc.NewRandomPrivateKey()
	if err != nil {
		return nil, err
	}

	privateKey2, err := easyecc.NewRandomPrivateKey()
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(privateKey.ToECDSA())

	// Create a simulated blockchain.
	alloc := make(core.GenesisAlloc)
	// Balance should be high enough to cover the transaction costs.
	initialBalance := big.NewInt(1000000000000000000)                     // 1 Ether.
	initialBalance = initialBalance.Mul(initialBalance, big.NewInt(1000)) // 1000 Ether.
	alloc[auth.From] = core.GenesisAccount{Balance: initialBalance}
	alloc[common.HexToAddress(privateKey1.PublicKey().EthereumAddress())] = core.GenesisAccount{Balance: initialBalance}
	alloc[common.HexToAddress(privateKey2.PublicKey().EthereumAddress())] = core.GenesisAccount{Balance: initialBalance}
	blockchain := backends.NewSimulatedBackend(alloc, 10000000)

	// Deploy contract.
	gasPrice, err := blockchain.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice

	return &SimulatedBlockchain{
		privateKey:  privateKey,
		privateKey1: privateKey1,
		privateKey2: privateKey2,
		auth:        auth,
		backend:     blockchain,
	}, nil
}

func (sbc *SimulatedBlockchain) PrivateKey() *easyecc.PrivateKey {
	return sbc.privateKey
}

func (sbc *SimulatedBlockchain) PrivateKey1() *easyecc.PrivateKey {
	return sbc.privateKey1
}

func (sbc *SimulatedBlockchain) PrivateKey2() *easyecc.PrivateKey {
	return sbc.privateKey2
}

func (sbc *SimulatedBlockchain) Auth() *bind.TransactOpts {
	return sbc.auth
}

func (sbc *SimulatedBlockchain) Backend() *backends.SimulatedBackend {
	return sbc.backend
}
