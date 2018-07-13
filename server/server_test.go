package server

import (
	"os"
	"testing"

	"google.golang.org/grpc"

	"github.com/facundomedica/go_grpc_flutter"
	"github.com/lileio/lile"
)

var s = GoGrpcFlutterServer{}
var cli go_grpc_flutter.GoGrpcFlutterClient

func TestMain(m *testing.M) {
	impl := func(g *grpc.Server) {
		go_grpc_flutter.RegisterGoGrpcFlutterServer(g, s)
	}

	gs := grpc.NewServer()
	impl(gs)

	addr, serve := lile.NewTestServer(gs)
	go serve()

	cli = go_grpc_flutter.NewGoGrpcFlutterClient(lile.TestConn(addr))

	os.Exit(m.Run())
}
