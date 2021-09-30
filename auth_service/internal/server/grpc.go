package server

import (
	"auth_service/internal/conf"
	"auth_service/internal/service"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	pb "github.com/raylin666/go-micro-protoc/services/auth/v1"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, greeter *service.AuthService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(
				recovery.WithLogger(log.DefaultLogger),
				recovery.WithHandler(func(ctx context.Context, req, err interface{}) error {
					return nil
				}),
			),
			logging.Server(logger),
			logging.Client(logger),
			validate.Validator(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)

	pb.RegisterAuthServer(srv, greeter)

	log.NewHelper(logger).Info("gRPC service started successfully")

	return srv
}
