package server

import (
	"errors"

	"context"
	"github.com/facundomedica/go_grpc_flutter"
)

func (s AuthServer) Register(ctx context.Context, r *go_grpc_flutter.AuthRequest) (*go_grpc_flutter.AuthResponse, error) {
	return nil, errors.New("not yet implemented")
}
