package core

import (
	"context"
	"fmt"
	"net"
	"net/url"

	"github.com/Becklyn/go-fx-core/env"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var GrpcModule = fx.Provide(
	newGrpc,
)

func newGrpc(lifecycle fx.Lifecycle, logger *logrus.Logger) *grpc.Server {
	addr := env.StringWithDefault("GRPC_ADDR", "tcp://0.0.0.0:9000")
	uri, err := url.Parse(addr)
	if err != nil {
		logger.Fatal(err)
	}

	listener, err := net.Listen(uri.Scheme, uri.Host)
	if err != nil {
		logger.Fatal(err)
	}

	server := grpc.NewServer()

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				reflection.Register(server)

				if err := server.Serve(listener); err != nil {
					logger.Fatal(err)
				}
			}()
			logger.Info(fmt.Sprintf("gRPC server running on %s://%s", uri.Scheme, uri.Host))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			server.GracefulStop()
			return nil
		},
	})

	return server
}
