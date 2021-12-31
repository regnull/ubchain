package gocontract

import (
	"context"
	"testing"

	"github.com/regnull/easyecc"
	"github.com/regnull/ubchain/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_DeployKeyRegistry(t *testing.T) {
	assert := assert.New(t)

	bc, err := testutil.NewSimulatedBlockchain()
	assert.NoError(err)

	ctx := context.Background()
	keyRegistry, err := deployKeyRegistry(ctx, bc)
	assert.NoError(err)
	assert.NotNil(keyRegistry)
}

func Test_RegisterKey(t *testing.T) {
	assert := assert.New(t)

	bc, err := testutil.NewSimulatedBlockchain()
	assert.NoError(err)

	ctx := context.Background()
	keyRegistry, err := deployKeyRegistry(ctx, bc)
	assert.NoError(err)

	myPrivateKey, err := easyecc.NewRandomPrivateKey()
	assert.NoError(err)
	publicKey := myPrivateKey.PublicKey().SerializeCompressed()
	assert.NoError(err)
	_, err = keyRegistry.Register(bc.Auth(), publicKey)
	assert.NoError(err)
	bc.Backend().Commit()

	e, err := keyRegistry.Exists(nil, publicKey)
	assert.NoError(err)
	assert.True(e)

	someOtherKey, err := easyecc.NewRandomPrivateKey()
	assert.NoError(err)
	e, err = keyRegistry.Exists(nil, someOtherKey.PublicKey().SerializeCompressed())
	assert.NoError(err)
	assert.False(e)
}
