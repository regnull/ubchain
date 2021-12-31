package testutil

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

type SimulatedBlockchain struct {
	privateKey *ecdsa.PrivateKey
	auth       *bind.TransactOpts
	backend    *backends.SimulatedBackend
}

func NewSimulatedBlockchain() (*SimulatedBlockchain, error) {
	// Generate private key.
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	auth := bind.NewKeyedTransactor(privateKey)

	// Create a simulated blockchain.
	alloc := make(core.GenesisAlloc)
	// Balance should be high enough to cover the transaction costs.
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, 10000000)

	// Deploy contract.
	gasPrice, err := blockchain.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice

	return &SimulatedBlockchain{
		privateKey: privateKey,
		auth:       auth,
		backend:    blockchain,
	}, nil
}

func (sbc *SimulatedBlockchain) PrivateKey() *ecdsa.PrivateKey {
	return sbc.privateKey
}

func (sbc *SimulatedBlockchain) Auth() *bind.TransactOpts {
	return sbc.auth
}

func (sbc *SimulatedBlockchain) Backend() *backends.SimulatedBackend {
	return sbc.backend
}
