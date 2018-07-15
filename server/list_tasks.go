package server

import (
	"github.com/facundomedica/go_grpc_flutter"
	"github.com/facundomedica/go_grpc_flutter/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s TasksServer) ListTasks(r *go_grpc_flutter.Empty, stream go_grpc_flutter.Tasks_ListTasksServer) error {
	username, ok := stream.Context().Value("username").(string)

	if !ok {
		return status.Error(codes.Unauthenticated, "Wow, you reached this without logging in? Impressive.")
	}

	// id text not null primary key, title text, timestamp integer, completed integer, owner string

	sqlStmt := `SELECT id, title, timestamp, completed, owner FROM tasks WHERE owner = ?`
	rows, err := database.DB.Query(sqlStmt, username)

	if err != nil {
		return status.Error(codes.Internal, "An error occurred while querying the database. "+err.Error())
	}

	for rows.Next() {
		task := &go_grpc_flutter.Task{}
		var completed int64
		err := rows.Scan(&task.Id, &task.Title, &task.Timestamp, &completed, &task.Owner)
		if err != nil {
			return status.Error(codes.Internal, "An error occurred while scanning the data")
		}

		err = stream.Send(task)
		if err != nil {
			return status.Error(codes.Internal, "An error occurred while returning the data")
		}
	}

	return nil
}
