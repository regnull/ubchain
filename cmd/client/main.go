package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const firstAccount = "0x8831b890c788db9B60B3e8a6199B245716A28fca"
const secondAccount = "0x07eC1DebD1A58f4eDbe522B86B5069cDdb8FA61b"

var ether = big.NewInt(1000000000000000000)

func main() {
	//transferFunds()
	getBalances()
	return

	/*
		ks := keystore.NewKeyStore("keystore", keystore.StandardScryptN, keystore.StandardScryptP)
		acc := loadAccount(ks, "0x8831b890c788db9B60B3e8a6199B245716A28fca")
		fmt.Printf("got account: %s\n", acc.Address.Hex())
		err := ks.Unlock(*acc, "123456")
		if err != nil {
			log.Fatal(err)
		}
		return
	*/

	ctx := context.Background()

	ec, _ := ethclient.Dial("http://192.168.86.35:8545")

	getGasPrice(ec)
	return

	head, _ := ec.BlockByNumber(ctx, nil)
	fmt.Println("Chain height:", head.Number())

	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	fmt.Printf("%s\n", account.Hex())
	balance, err := ec.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)
}

func loadAccount(ks *keystore.KeyStore, address string) *accounts.Account {
	for _, acc := range ks.Accounts() {
		if acc.Address.Hex() == address {
			return &acc
		}
	}
	return nil
}

func createNewAccount(ks *keystore.KeyStore) {
	// Create new account.
	newAcc, err := ks.NewAccount("123456")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("New account address: ", newAcc.Address.Hex())
}

func getGasPrice(client *ethclient.Client) {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("gas price: %d\n", gasPrice)
}

func transferFunds() {
	ctx := context.Background()
	client, err := ethclient.Dial("http://192.168.86.35:8545")
	if err != nil {
		log.Fatal(err)
	}

	ks := keystore.NewKeyStore("keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	acc := loadAccount(ks, firstAccount)
	if acc == nil {
		log.Fatal("failed to load account")
	}
	err = ks.Unlock(*acc, "123456")
	if err != nil {
		log.Fatal(err)
	}
	nonce, err := client.PendingNonceAt(ctx, acc.Address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("got nonce: %d\n", nonce)

	gasLimit := uint64(21000)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("gas price: %d\n", gasPrice)

	toAddress := common.HexToAddress(secondAccount)

	value := ether

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	chainID, err := client.NetworkID(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("chain ID: %d\n", chainID)
	// TODO: Chain ID is incorrect. It returns 1, and our chain ID is 101.

	signedTx, err := ks.SignTx(*acc, tx, big.NewInt(101))
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

func getBalances() {
	ctx := context.Background()

	client, err := ethclient.Dial("http://192.168.86.35:8545")
	if err != nil {
		log.Fatal(err)
	}

	for _, h := range []string{firstAccount, secondAccount} {
		account := common.HexToAddress(h)
		balance, err := client.BalanceAt(ctx, account, nil)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s: %d\n", account.Hex(), balance)
	}
}
