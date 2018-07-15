package server

import (
	"time"

	"context"

	"github.com/facundomedica/go_grpc_flutter"
	"github.com/facundomedica/go_grpc_flutter/database"
	"github.com/facundomedica/go_grpc_flutter/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s TasksServer) CreateTask(ctx context.Context, r *go_grpc_flutter.Task) (*go_grpc_flutter.Task, error) {
	// Here we ignore id, timestamp, completed and owner. We set them on the server side.
	if r.Title == "" {
		return nil, status.Error(codes.InvalidArgument, "The title can't be empty!")
	}

	stmt, err := database.DB.Prepare("INSERT INTO tasks (id, timestamp, title, completed, owner) VALUES (?,?,?,?,?)")

	if err != nil {
		return nil, status.Error(codes.Internal, "Something is very wrong!")
	}

	timestamp := time.Now().Unix()
	id := utils.RandSeq(12) // Our homemade "UUID" (Don't use this!!!)
	username, ok := ctx.Value("username").(string)

	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Wow, you reached this without logging in? Impressive.")
	}

	_, err = stmt.Exec(id, timestamp, r.Title, 0, username)

	if err != nil {
		return nil, status.Error(codes.Internal, "Some error occurred while trying to insert the task! "+err.Error())
	}

	r.Id = id
	r.Timestamp = timestamp
	r.Completed = false
	r.Owner = username

	return r, nil
}
