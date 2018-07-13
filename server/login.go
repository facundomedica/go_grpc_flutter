package server

import (
	"context"
	"database/sql"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/facundomedica/go_grpc_flutter"
	"github.com/facundomedica/go_grpc_flutter/database"
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

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain. In this case we are storing the username
	// and the expiration date.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"expires":  time.Now().Add(72 * time.Hour),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("thisismyveysecretkey!!@@33$$5asqweasdqwe"))

	if err != nil {
		return nil, status.Error(codes.Internal, "There was an error making the token")
	}

	return &go_grpc_flutter.AuthResponse{Token: tokenString}, nil
}
