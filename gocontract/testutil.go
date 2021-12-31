package gocontract

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/regnull/ubchain/testutil"
)

func deployKeyRegistry(ctx context.Context, bc *testutil.SimulatedBlockchain) (*KeyRegistry, common.Address, error) {
	gasPrice, err := bc.Backend().SuggestGasPrice(ctx)
	if err != nil {
		return nil, common.Address{}, err
	}
	bc.Auth().GasPrice = gasPrice
	addr, _, instance, err := DeployKeyRegistry(bc.Auth(), bc.Backend())
	if err != nil {
		return nil, common.Address{}, err
	}
	bc.Backend().Commit()
	return instance, addr, nil
}

func deployNameRegistry(ctx context.Context, bc *testutil.SimulatedBlockchain,
	keyRegistryAddr common.Address) (*NameRegistry, common.Address, error) {

	gasPrice, err := bc.Backend().SuggestGasPrice(ctx)
	if err != nil {
		return nil, common.Address{}, err
	}
	bc.Auth().GasPrice = gasPrice
	addr, _, instance, err := DeployNameRegistry(bc.Auth(), bc.Backend(), keyRegistryAddr)
	if err != nil {
		return nil, common.Address{}, err
	}
	bc.Backend().Commit()
	return instance, addr, nil
}
