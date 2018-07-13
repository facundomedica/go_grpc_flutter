package server

import (
	"errors"
	"github.com/facundomedica/go_grpc_flutter"
)

func (s TasksServer) ListTasks(r *go_grpc_flutter.Empty, stream go_grpc_flutter.Tasks_ListTasksServer) error {
	res := &go_grpc_flutter.Task{}
	err := stream.Send(res)
	if err != nil {
		return err
	}

	return errors.New("not yet implemented")
}
