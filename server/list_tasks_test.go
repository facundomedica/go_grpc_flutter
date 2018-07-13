package server

import (
	"io"
	"testing"

	"context"

	"github.com/facundomedica/go_grpc_flutter"
	"github.com/stretchr/testify/assert"
)

func TestListTasks(t *testing.T) {
	ctx := context.Background()
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
