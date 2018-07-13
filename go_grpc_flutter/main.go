package main

import (
	_ "net/http/pprof"

	"github.com/facundomedica/go_grpc_flutter"
	"github.com/facundomedica/go_grpc_flutter/database"
	"github.com/facundomedica/go_grpc_flutter/go_grpc_flutter/cmd"
	"github.com/facundomedica/go_grpc_flutter/server"
	"github.com/lileio/fromenv"
	"github.com/lileio/lile"
	"github.com/lileio/logr"
	"github.com/lileio/pubsub"
	"github.com/lileio/pubsub/middleware/defaults"
	"google.golang.org/grpc"
)

func main() {
	logr.SetLevelFromEnv()
	as := &server.AuthServer{}
	ts := &server.TasksServer{}

	database.InitDB()

	lile.Name("go_grpc_flutter")
	lile.Server(func(g *grpc.Server) {
		go_grpc_flutter.RegisterAuthServer(g, as)
		go_grpc_flutter.RegisterTasksServer(g, ts)
	})

	pubsub.SetClient(&pubsub.Client{
		ServiceName: lile.GlobalService().Name,
		Provider:    fromenv.PubSubProvider(),
		Middleware:  defaults.Middleware,
	})

	cmd.Execute()
}
