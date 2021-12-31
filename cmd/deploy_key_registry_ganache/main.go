package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/regnull/ubchain/gocontract"
)

func main() {
	// Generate private key.
	privateKey, err := crypto.HexToECDSA("394e1417beae2d42f31b92fb784f93bb1e1779904ea823bc33d0d7774e818124")
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)

	// Connect to the node.
	ctx := context.Background()
	client, err := ethclient.Dial("HTTP://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}

	// Deploy contract.
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice
	address, tx, instance, err := gocontract.DeployKeyRegistry(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("contract address: %s\n", address.String())
	fmt.Printf("deploy tx: %s\n", tx.Hash().Hex())

	_ = instance

	/*
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
	*/
}
