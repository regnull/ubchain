package gocontract

import (
	"context"

	"github.com/regnull/ubchain/testutil"
)

func deployKeyRegistry(ctx context.Context, bc *testutil.SimulatedBlockchain) (*KeyRegistry, error) {
	gasPrice, err := bc.Backend().SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}
	bc.Auth().GasPrice = gasPrice
	_, _, instance, err := DeployKeyRegistry(bc.Auth(), bc.Backend())
	if err != nil {
		return nil, err
	}
	bc.Backend().Commit()
	return instance, nil
}
