package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/regnull/ubchain/gocontract"
)

func main() {
	client, err := ethclient.Dial("HTTP://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}

	keyToQuery := common.Hex2Bytes("0340a0133eb6cc6f7966cf032f1941f4cced9ff872930dba796483d3307a154a31")

	contractAddress := common.HexToAddress("0x9EB5b6f4f2FC106760fb42837068AE000CDeC641")
	instance, err := gocontract.NewKeyRegistry(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	res, err := instance.Registry(nil, keyToQuery)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
}
