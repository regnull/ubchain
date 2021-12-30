package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/regnull/easyecc"
	"github.com/regnull/ubchain/keyregistry"
)

func main() {
	// Generate private key.
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
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
	address, tx, instance, err := keyregistry.DeployKeyregistry(auth, blockchain)
	blockchain.Commit()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("contract address: %s\n", address.String())
	fmt.Printf("deploy tx: %s\n", tx.Hash().Hex())

	myPrivateKey, err := easyecc.NewRandomPrivateKey()
	if err != nil {
		log.Fatal(err)
	}
	tx, err = instance.Register(auth, myPrivateKey.PublicKey().SerializeCompressed())
	if err != nil {
		log.Fatal(err)
	}
	blockchain.Commit()
	fmt.Printf("set tx: %s\n", tx.Hash().Hex())

	// Query an item.
	result, err := instance.Registry(nil, myPrivateKey.PublicKey().SerializeCompressed())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("item: %+v\n", result)
}
