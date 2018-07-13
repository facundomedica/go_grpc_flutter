package server

import (
	"errors"

	"context"

	"github.com/facundomedica/go_grpc_flutter"
)

func (s TasksServer) CreateTask(ctx context.Context, r *go_grpc_flutter.Task) (*go_grpc_flutter.Task, error) {
	// Here we ignore Id, timestamp and owner. We set them on the server side.
	return nil, errors.New("not yet implemented")
}
