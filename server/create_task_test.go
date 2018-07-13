package server

import (
	"testing"

	"context"

	"github.com/facundomedica/go_grpc_flutter"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	ctx := context.Background()
	req := &go_grpc_flutter.Task{}

	res, err := tasksCli.CreateTask(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
