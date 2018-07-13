package server

import (
	"os"
	"testing"

	"google.golang.org/grpc"

	"github.com/facundomedica/go_grpc_flutter"
	"github.com/lileio/lile"
)

var authServer = AuthServer{}
var tasksServer = TasksServer{}
var authCli go_grpc_flutter.AuthClient
var tasksCli go_grpc_flutter.TasksClient

func TestMain(m *testing.M) {
	impl := func(g *grpc.Server) {
		go_grpc_flutter.RegisterAuthServer(g, authServer)
		go_grpc_flutter.RegisterTasksServer(g, tasksServer)
	}

	gs := grpc.NewServer()
	impl(gs)

	addr, serve := lile.NewTestServer(gs)
	go serve()

	authCli = go_grpc_flutter.NewAuthClient(lile.TestConn(addr))
	tasksCli = go_grpc_flutter.NewTasksClient(lile.TestConn(addr))

	os.Exit(m.Run())
}
