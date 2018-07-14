package server

import (
	"context"

	"github.com/facundomedica/go_grpc_flutter"
	"github.com/facundomedica/go_grpc_flutter/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s TasksServer) DeleteTask(ctx context.Context, r *go_grpc_flutter.Task) (*go_grpc_flutter.Empty, error) {
	username, ok := ctx.Value("username").(string)

	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Wow, you reached this with out loggin in? Impressive.")
	}
	stmt, err := database.DB.Prepare("DELETE FROM tasks WHERE id = ? AND owner = ?")

	if err != nil {
		return nil, status.Error(codes.Internal, "Something is very wrong!")
	}

	res, err := stmt.Exec(r.Id, username)

	if err != nil {
		return nil, status.Error(codes.Internal, "Some error occurred while trying to insert the task! "+err.Error())
	}

	rowsAffected, _ := res.RowsAffected()

	if rowsAffected != 0 {
		return nil, status.Error(codes.Internal, "Couldn't delete the task!")
	}

	return &go_grpc_flutter.Empty{}, nil
}
