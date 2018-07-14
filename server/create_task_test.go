package server

import (
	"testing"

	"context"

	"github.com/facundomedica/go_grpc_flutter"
	"github.com/facundomedica/go_grpc_flutter/utils"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

func TestCreateTask(t *testing.T) {
	req := &go_grpc_flutter.Task{Title: "Testing"}
	token, _ := utils.MakeToken("user_that_exists")
	md := metadata.New(map[string]string{"authorization": "Bearer " + token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	res, err := tasksCli.CreateTask(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
