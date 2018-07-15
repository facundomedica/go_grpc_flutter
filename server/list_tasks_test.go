package server

import (
	"io"
	"testing"

	"context"

	"github.com/facundomedica/go_grpc_flutter"
	"github.com/facundomedica/go_grpc_flutter/utils"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

func TestListTasks(t *testing.T) {
	token, _ := utils.MakeToken("user_that_exists")
	md := metadata.New(map[string]string{"authorization": "Bearer " + token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	req := &go_grpc_flutter.Empty{}

	stream, err := tasksCli.ListTasks(ctx, req)
	assert.Nil(t, err)

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			assert.Fail(t, err.Error())
			break
		}

		assert.Nil(t, err)
		assert.NotNil(t, res)
	}
}
