.PHONY: gencode

ROOT_DIR = $(dir $(realpath $(firstword $(MAKEFILE_LIST))))

gencode:
	mkdir -p $(ROOT_DIR)/build
	solc --optimize --bin contracts/*.sol -o $(ROOT_DIR)/build --overwrite
	solc --abi contracts/*.sol -o $(ROOT_DIR)/build --overwrite
	abigen --bin=build/NameRegistry.bin --abi=build/NameRegistry.abi --type=NameRegistry --pkg=gocontract --out=gocontract/name_registry.go
