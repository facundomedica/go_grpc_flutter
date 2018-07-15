package server

import (
	"context"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/facundomedica/go_grpc_flutter"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type AuthServer struct {
	go_grpc_flutter.AuthServer
}

type TasksServer struct {
	go_grpc_flutter.TasksServer
}

func (AuthServer) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

func AuthFunc(ctx context.Context) (context.Context, error) {
	tokenString, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, grpc.Errorf(codes.Unauthenticated, "Unexpected signing method: %v", token.Header["alg"])
		}

		// Here you have to return the same key that you used to sign the token (utils/token.go)
		return []byte("thisismyveysecretkey!!@@33$$5asqweasdqwe"), nil
	})

	if err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	var username string

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username, ok = claims["username"].(string)
		if !ok {
			return nil, grpc.Errorf(codes.Unauthenticated, "invalid auth token, no username found")
		}
	} else {
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid auth token, invalid claims")
	}

	newCtx := context.WithValue(ctx, "username", username)
	return newCtx, nil
}
