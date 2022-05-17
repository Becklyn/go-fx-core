package core

import (
	"context"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var GrpcModule = fx.Provide(
	newGrpc,
)

func newGrpc(lifecycle fx.Lifecycle, logger *logrus.Logger) *grpc.Server {
	addr := ":9000"

	listener, err := net.Listen("tcp", addr)
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
			logger.Info(fmt.Sprintf("gRPC server bound on host 0.0.0.0 and port %s", addr))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			server.GracefulStop()
			return nil
		},
	})

	return server
}
