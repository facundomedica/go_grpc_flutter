package server

import (
	"errors"
	"math/rand"
	"testing"

	"context"

	"github.com/facundomedica/go_grpc_flutter"
)

func TestRegister(t *testing.T) {
	ctx := context.Background()

	var reqTests = []struct {
		request *go_grpc_flutter.AuthRequest // input
		err     error                        // expected result
	}{
		{&go_grpc_flutter.AuthRequest{
			Username: "already_exists",
			Password: "passwordpassword"},
			errors.New("rpc error: code = InvalidArgument desc = This user already exists"),
		}, // a username that already exists (you have to run the test twice the first time)
		{&go_grpc_flutter.AuthRequest{
			Username: "notexists" + randSeq(6),
			Password: "passwordpassword"},
			nil}, // a username that does not exist
	}

	for _, tt := range reqTests {
		_, err := authCli.Register(ctx, tt.request)

		if err == nil && tt.err == nil {
			continue // we are not expecting errors and it's okay
		}

		if err == nil && tt.err != nil { // we were expecting an error and we got none
			t.Errorf("Register with username %s was expected to fail and it did not", tt.request.Username)
		}

		if err != nil && tt.err == nil { //we were not expecting an error and we got one
			t.Errorf("Register with username %s was expected to succeed and it did not", tt.request.Username)
		}
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
