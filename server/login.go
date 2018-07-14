package server

import (
	"context"
	"database/sql"

	"github.com/facundomedica/go_grpc_flutter"
	"github.com/facundomedica/go_grpc_flutter/database"
	"github.com/facundomedica/go_grpc_flutter/utils"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s AuthServer) Login(ctx context.Context, r *go_grpc_flutter.AuthRequest) (*go_grpc_flutter.AuthResponse, error) {
	sqlStmt := `SELECT * FROM users WHERE username = ?`
	var username string
	var passwordHash string
	err := database.DB.QueryRow(sqlStmt, r.Username).Scan(&username, &passwordHash)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, status.Error(codes.Internal, "An error occurred when getting the user from the db")
		}

		return nil, status.Error(codes.NotFound, "This user does not exists")
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(r.Password))

	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Wrong password!")
	}

	tokenString, err := utils.MakeToken(username)

	if err != nil {
		return nil, status.Error(codes.Internal, "There was an error making the token")
	}

	return &go_grpc_flutter.AuthResponse{Token: tokenString}, nil
}
