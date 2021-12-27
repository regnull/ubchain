.PHONY: gencode

ROOT_DIR = $(dir $(realpath $(firstword $(MAKEFILE_LIST))))

gencode:
	mkdir -p $(ROOT_DIR)/build
	solc --optimize --bin contracts/key_registry.sol -o $(ROOT_DIR)/build --overwrite
	solc --abi contracts/key_registry.sol -o $(ROOT_DIR)/build --overwrite
	abigen --bin=build/KeyRegistry.bin --abi=build/KeyRegistry.abi --pkg=gocontract --out=gocontract/KeyRegistry.go