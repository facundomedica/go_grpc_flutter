package server

import (
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/facundomedica/go_grpc_flutter"
	"github.com/facundomedica/go_grpc_flutter/database"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/lileio/lile"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

var authServer = AuthServer{}
var tasksServer = TasksServer{}
var authCli go_grpc_flutter.AuthClient
var tasksCli go_grpc_flutter.TasksClient

func TestMain(m *testing.M) {
	database.InitDB()
	rand.Seed(time.Now().UnixNano())

	impl := func(g *grpc.Server) {
		go_grpc_flutter.RegisterAuthServer(g, authServer)
		go_grpc_flutter.RegisterTasksServer(g, tasksServer)

	}

	gs := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(AuthFunc)),
		grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(AuthFunc)),
	)
	impl(gs)

	addr, serve := lile.NewTestServer(gs)

	go serve()

	authCli = go_grpc_flutter.NewAuthClient(lile.TestConn(addr))
	tasksCli = go_grpc_flutter.NewTasksClient(lile.TestConn(addr))

	os.Exit(m.Run())
}
