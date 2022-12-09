.PHONY: gencode

ROOT_DIR = $(dir $(realpath $(firstword $(MAKEFILE_LIST))))

gencode:
	mkdir -p $(ROOT_DIR)/build
	mkdir -p $(ROOT_DIR)/build/v2
	solc --optimize --bin contracts/*.sol -o $(ROOT_DIR)/build --overwrite
	solc --optimize --bin contracts/v2/*.sol -o $(ROOT_DIR)/build/v2 --overwrite
	solc --abi contracts/*.sol -o $(ROOT_DIR)/build --overwrite
	solc --abi contracts/v2/*.sol -o $(ROOT_DIR)/build/v2 --overwrite
	abigen --bin=build/KeyRegistry.bin --abi=build/KeyRegistry.abi --type=KeyRegistry --pkg=gocontract --out=gocontract/key_registry.go
	abigen --bin=build/NameRegistry.bin --abi=build/NameRegistry.abi --type=NameRegistry --pkg=gocontract --out=gocontract/name_registry.go
	abigen --bin=build/ConnectorRegistry.bin --abi=build/ConnectorRegistry.abi --type=ConnectorRegistry --pkg=gocontract --out=gocontract/connector_registry.go	
	abigen --bin=build/v2/NameRegistry.bin --abi=build/v2/NameRegistry.abi --type=NameRegistry --pkg=v2 --out=gocontract/v2/name_registry.go
