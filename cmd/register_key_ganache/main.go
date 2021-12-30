package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/regnull/easyecc"
	"github.com/regnull/ubchain/keyregistry"
)

func main() {
	// Generate private key.
	privateKey, err := crypto.HexToECDSA("394e1417beae2d42f31b92fb784f93bb1e1779904ea823bc33d0d7774e818124")
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)

	// Connect to the node.
	//ctx := context.Background()
	client, err := ethclient.Dial("HTTP://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x9EB5b6f4f2FC106760fb42837068AE000CDeC641")
	instance, err := keyregistry.NewKeyregistry(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	keyToRegister, err := easyecc.NewRandomPrivateKey()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("got key: %x\n", keyToRegister.PublicKey().SerializeCompressed())
	tx, err := instance.Register(auth, keyToRegister.PublicKey().SerializeCompressed())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("got tx: %s\n", tx.Hash().Hex())

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
