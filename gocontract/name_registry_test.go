package gocontract

import (
	"context"
	"testing"

	"github.com/regnull/ubchain/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_RegisterName(t *testing.T) {
	assert := assert.New(t)

	bc, err := testutil.NewSimulatedBlockchain()
	assert.NoError(err)

	ctx := context.Background()
	keyRegistry, keyRegistryAddr, err := deployKeyRegistry(ctx, bc)
	assert.NoError(err)

	nameRegistry, _, err := deployNameRegistry(ctx, bc, keyRegistryAddr)
	assert.NoError(err)
	bc.Backend().Commit()

	publicKey := bc.PrivateKey().PublicKey().SerializeCompressed()
	_, err = keyRegistry.Register(bc.Auth(), publicKey)
	assert.NoError(err)
	bc.Backend().Commit()

	_, err = nameRegistry.Register(bc.Auth(), "spongebob", publicKey)
	assert.NoError(err)
	bc.Backend().Commit()

	key, err := nameRegistry.GetKey(nil, "spongebob")
	assert.NoError(err)

	assert.Equal(key, publicKey)

	key, err = nameRegistry.GetKey(nil, "patrick")
	assert.NoError(err)
	assert.Empty(key)
}
