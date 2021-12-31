package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/regnull/ubchain/gocontract"
)

func main() {
	var (
		nodeURL         string
		contractAddress string
		publicKey       string
	)
	flag.StringVar(&nodeURL, "node-url", "http://127.0.0.1:7545", "URL of the node to connect to")
	flag.StringVar(&contractAddress, "contract-address", "", "contract address")
	flag.StringVar(&publicKey, "public-key", "", "public key")
	flag.Parse()

	// Connect to the node.
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		log.Fatal(err)
	}

	instance, err := gocontract.NewKeyRegistry(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatal(err)
	}

	res, err := instance.Registered(nil, common.Hex2Bytes(publicKey))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", res)
}
