package server

import (
	"github.com/facundomedica/go_grpc_flutter"
)

type AuthServer struct {
	go_grpc_flutter.AuthServer
}

type TasksServer struct {
	go_grpc_flutter.TasksServer
}
