package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/regnull/easyecc"
	"github.com/regnull/ubchain/gocontract"
)

func main() {
	var (
		nodeURL         string
		keystoreDir     string
		address         string
		contractAddress string
		password        string
	)
	flag.StringVar(&nodeURL, "node-url", "http://127.0.0.1:7545", "URL of the node to connect to")
	flag.StringVar(&keystoreDir, "keystore-dir", "keystore", "keystore directory")
	flag.StringVar(&address, "address", "", "account address of the transactions originator")
	flag.StringVar(&contractAddress, "contract-address", "", "contract address")
	flag.StringVar(&password, "password", "", "account password")
	flag.Parse()

	ctx := context.Background()
	// Connect to the node.
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		log.Fatal(err)
	}

	// Open the keystore, find the account, and unlock it.
	ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.Find(accounts.Account{Address: common.HexToAddress(address)})
	if err != nil {
		log.Fatal(err)
	}
	err = ks.Unlock(account, password)
	if err != nil {
		log.Fatal(err)
	}

	// Get nonce.
	nonce, err := client.PendingNonceAt(ctx, common.HexToAddress(address))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("got nonce: %d\n", nonce)

	// Recommended gas limit.
	gasLimit := uint64(300000)

	// Get gas price.
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("gas price: %d\n", gasPrice)

	chainID, err := client.NetworkID(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("chain ID: %d\n", chainID)

	auth, err := bind.NewKeyStoreTransactorWithChainID(ks, account, chainID)
	//auth, err := bind.NewKeyedTransactorWithChainID(privateKey.ToECDSA(), chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	instance, err := gocontract.NewKeyRegistry(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatal(err)
	}

	keyToRegister, err := easyecc.NewRandomPrivateKey()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("public key: %x\n", keyToRegister.PublicKey().SerializeCompressed())

	tx, err := instance.Register(auth, keyToRegister.PublicKey().SerializeCompressed())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
}
