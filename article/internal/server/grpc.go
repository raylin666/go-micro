package server

import (
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"mt/api/v1"
	"mt/config"
	"mt/internal/middleware/auth"
	"mt/internal/middleware/validate"
	"mt/internal/service"
	"mt/pkg/logger"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	articleV1 "github.com/raylin666/go-micro-protoc/api/article/v1"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *config.Server, heartbeat *service.HeartbeatService, article *service.ArticleService, logger *logger.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			validate.Validator(),
			logging.Server(logger),
			metadata.Server(),
			auth.NewAuthServer(),
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
	v1.RegisterHeartbeatServer(srv, heartbeat)
	articleV1.RegisterArticleServer(srv, article)
	return srv
}
