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

func Test_RegisterNameAgain(t *testing.T) {
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

	publicKey1 := bc.PrivateKey1().PublicKey().SerializeCompressed()
	_, err = keyRegistry.Register(bc.Auth1(), publicKey1)
	assert.NoError(err)
	bc.Backend().Commit()

	// First regisration.
	_, err = nameRegistry.Register(bc.Auth(), "spongebob", publicKey)
	assert.NoError(err)
	bc.Backend().Commit()

	// Second registration.
	_, err = nameRegistry.Register(bc.Auth(), "spongebob", publicKey1)
	assert.NoError(err)
	bc.Backend().Commit()

	key, err := nameRegistry.GetKey(nil, "spongebob")
	assert.NoError(err)

	assert.Equal(key, publicKey1)

	// Now if we try to register the name again, we should get an error
	// (since we don't own the key anymore.)
	publicKey2 := bc.PrivateKey2().PublicKey().SerializeCompressed()
	_, err = keyRegistry.Register(bc.Auth2(), publicKey2)
	assert.NoError(err)
	bc.Backend().Commit()

	_, err = nameRegistry.Register(bc.Auth(), "spongebob", publicKey2)
	assert.Error(err)

	// But the new owner should be able to change the registration.
	_, err = nameRegistry.Register(bc.Auth1(), "spongebob", publicKey2)
	assert.NoError(err)
}
