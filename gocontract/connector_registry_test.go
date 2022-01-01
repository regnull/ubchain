package gocontract

import (
	"context"
	"testing"

	"github.com/regnull/ubchain/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_RegisterConnector(t *testing.T) {
	assert := assert.New(t)

	bc, err := testutil.NewSimulatedBlockchain()
	assert.NoError(err)

	ctx := context.Background()

	// Deploy all registries.
	keyRegistry, keyRegistryAddr, err := deployKeyRegistry(ctx, bc)
	assert.NoError(err)
	bc.Commit()

	nameRegistry, nameRegistryAddr, err := deployNameRegistry(ctx, bc, keyRegistryAddr)
	assert.NoError(err)
	bc.Commit()

	connRegistry, _, err := deployConnectorRegistry(ctx, bc,
		keyRegistryAddr, nameRegistryAddr)
	assert.NoError(err)
	bc.Commit()

	// Register public key.
	publicKey := bc.PrivateKey().PublicKey().SerializeCompressed()
	_, err = keyRegistry.Register(bc.Auth(), publicKey)
	assert.NoError(err)
	bc.Commit()

	// Register name.
	_, err = nameRegistry.Register(bc.Auth(), "spongebob", publicKey)
	assert.NoError(err)
	bc.Commit()

	// Register connector.
	_, err = connRegistry.Register(bc.Auth(), "spongebob", "PL_DMS", "http://somelocation.blah")
	assert.NoError(err)
	bc.Commit()

	location, err := connRegistry.GetLocation(nil, "spongebob", "PL_DMS")
	assert.NoError(err)
	assert.Equal("http://somelocation.blah", location)

	// Attempt to register with wrong auth.
	_, err = connRegistry.Register(bc.Auth1(), "spongebob", "PL_XYZ", "http://somelocation.blah")
	assert.Error(err)
}
