package server

import (
	"context"

	"github.com/facundomedica/go_grpc_flutter"
	"github.com/facundomedica/go_grpc_flutter/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s TasksServer) UpdateTask(ctx context.Context, r *go_grpc_flutter.Task) (*go_grpc_flutter.Task, error) {
	// UPDATE tasks SET title = ?, completed = ? WHERE owner = ? AND id = ?

	username, ok := ctx.Value("username").(string)

	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Wow, you reached this without logging in? Impressive.")
	}
	stmt, err := database.DB.Prepare("UPDATE tasks SET title = ?, completed = ? WHERE owner = ? AND id = ?")

	if err != nil {
		return nil, status.Error(codes.Internal, "Something is very wrong!")
	}

	var completed int64
	if r.Completed {
		completed = 1
	}

	res, err := stmt.Exec(r.Title, completed, username, r.Id)

	if err != nil {
		return nil, status.Error(codes.Internal, "Some error occurred while trying to update the task! "+err.Error())
	}

	rowsAffected, _ := res.RowsAffected()

	if rowsAffected != 1 {
		return nil, status.Error(codes.Internal, "Couldn't update the task!")
	}

	return r, nil
}
