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
	cm     = &sync.Mutex{}
	Client GoGrpcFlutterClient
)

func GetGoGrpcFlutterClient() GoGrpcFlutterClient {
	cm.Lock()
	defer cm.Unlock()

	if Client != nil {
		return Client
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

	cli := NewGoGrpcFlutterClient(conn)
	Client = cli
	return cli
}
