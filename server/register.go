package server

import (
	"context"

	"github.com/facundomedica/go_grpc_flutter"
	"github.com/facundomedica/go_grpc_flutter/database"
	"github.com/facundomedica/go_grpc_flutter/utils"
	sqlite3 "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s AuthServer) Register(ctx context.Context, r *go_grpc_flutter.AuthRequest) (*go_grpc_flutter.AuthResponse, error) {
	stmt, err := database.DB.Prepare("INSERT INTO users (username, password) VALUES (?,?)")

	if err != nil {
		return nil, status.Error(codes.Internal, "Something is very wrong!")
	}

	passwordHashBytes, err := bcrypt.GenerateFromPassword([]byte(r.Password), 10)

	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to hash the password.")
	}

	_, err = stmt.Exec(r.Username, string(passwordHashBytes))

	if err != nil {
		// I check if the user already exists or not. Not elegant but works.
		if (err.(sqlite3.Error)).Code == sqlite3.ErrConstraint {
			return nil, status.Error(codes.InvalidArgument, "This user already exists")
		}
		return nil, status.Error(codes.Internal, "Some error occurred while trying to insert the user! "+err.Error())
	}

	tokenString, err := utils.MakeToken(r.Username)

	if err != nil {
		return nil, status.Error(codes.Internal, "There was an error making the token")
	}

	return &go_grpc_flutter.AuthResponse{Token: tokenString}, nil
}
