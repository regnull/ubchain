package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {
	ks := keystore.NewKeyStore("keystore", keystore.StandardScryptN, keystore.StandardScryptP)

	// Create new account.
	newAcc, _ := ks.NewAccount("123456")
	fmt.Println("New:", newAcc.Address.Hex())
}
