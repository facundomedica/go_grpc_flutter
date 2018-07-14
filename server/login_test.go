package server

import (
	"errors"
	"testing"

	"context"

	"github.com/facundomedica/go_grpc_flutter"
	"github.com/facundomedica/go_grpc_flutter/utils"
)

func TestLogin(t *testing.T) {
	ctx := context.Background()

	var reqTests = []struct {
		request *go_grpc_flutter.AuthRequest // input
		err     error                        // expected result
	}{
		{&go_grpc_flutter.AuthRequest{
			Username: "user_that_exists",
			Password: "passwordpassword"},
			nil,
		}, // a username that already exists (you have to run the test twice the first time)
		{&go_grpc_flutter.AuthRequest{
			Username: "notexists" + utils.RandSeq(6),
			Password: "passwordpassword"},
			errors.New("rpc error: code = NotFound desc = This user does not exists"),
		}, // a username that does not exist
	}

	for _, tt := range reqTests {
		_, err := authCli.Login(ctx, tt.request)

		if err == nil && tt.err == nil {
			continue // we are not expecting errors and it's okay
		}

		if err == nil && tt.err != nil { // we were expecting an error and we got none
			t.Errorf("Login with username %s and password %s was expected to fail and it did not", tt.request.Username, tt.request.Password)
		}

		if err != nil && tt.err == nil { //we were not expecting an error and we got one
			t.Errorf("Login with username %s and password %s was expected to succeed and it did not", tt.request.Username, tt.request.Password)
		}
	}
}
