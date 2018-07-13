package go_grpc_flutter

import (
	"sync"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/lileio/lile"
	opentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

var (
	cm      = &sync.Mutex{}
	aClient AuthClient
	tClient TasksClient
)

func GetAuthClient() AuthClient {
	cm.Lock()
	defer cm.Unlock()

	if aClient != nil {
		return aClient
	}

	serviceURL := lile.URLForService("go-grpc-flutter")

	// We don't need to error here, as this creates a pool and connections
	// will happen later
	conn, _ := grpc.Dial(
		serviceURL,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(
				lile.ContextClientInterceptor(),
				otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer()),
			),
		))

	cli := NewAuthClient(conn)
	aClient = cli
	return cli
}

func GetTasksClient() TasksClient {
	cm.Lock()
	defer cm.Unlock()

	if tClient != nil {
		return tClient
	}

	serviceURL := lile.URLForService("go-grpc-flutter")

	// We don't need to error here, as this creates a pool and connections
	// will happen later
	conn, _ := grpc.Dial(
		serviceURL,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(
				lile.ContextClientInterceptor(),
				otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer()),
			),
		))

	cli := NewTasksClient(conn)
	tClient = cli
	return cli
}
