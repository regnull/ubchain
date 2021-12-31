.PHONY: gencode

ROOT_DIR = $(dir $(realpath $(firstword $(MAKEFILE_LIST))))

gencode:
	mkdir -p $(ROOT_DIR)/build
	solc --optimize --bin contracts/*.sol -o $(ROOT_DIR)/build --overwrite
	solc --abi contracts/*.sol -o $(ROOT_DIR)/build --overwrite
	abigen --bin=build/KeyRegistry.bin --abi=build/KeyRegistry.abi --type=KeyRegistry --pkg=gocontract --out=gocontract/KeyRegistry.go
	abigen --bin=build/NameRegistry.bin --abi=build/NameRegistry.abi --type=NameRegistry --pkg=gocontract --out=gocontract/NameRegistry.go
	abigen --bin=build/ConnectorRegistry.bin --abi=build/ConnectorRegistry.abi --type=ConnectorRegistry --pkg=gocontract --out=gocontract/ConnectorRegistry.go	