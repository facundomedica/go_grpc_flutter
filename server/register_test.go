package server

import (
	"testing"

	"context"

	"github.com/facundomedica/go_grpc_flutter"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	ctx := context.Background()
	req := &go_grpc_flutter.AuthRequest{}

	res, err := authCli.Register(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
