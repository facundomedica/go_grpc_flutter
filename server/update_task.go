package server

import (
	"errors"

	"context"
	"github.com/facundomedica/go_grpc_flutter"
)

func (s TasksServer) UpdateTask(ctx context.Context, r *go_grpc_flutter.Task) (*go_grpc_flutter.Task, error) {
	return nil, errors.New("not yet implemented")
}
