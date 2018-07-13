package server

import (
	"testing"

	"context"
	"github.com/facundomedica/go_grpc_flutter"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	ctx := context.Background()
	req := &go_grpc_flutter.AuthRequest{}

	res, err := cli.Login(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
