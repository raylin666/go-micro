package server

import (
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"mt/api/v1"
	"mt/config"
	"mt/internal/middleware/auth"
	"mt/internal/middleware/cors"
	"mt/internal/middleware/encode"
	"mt/internal/service"
	"mt/pkg/logger"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"

	articleV1 "github.com/raylin666/go-micro-protoc/api/article/v1"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *config.Server, heartbeat *service.HeartbeatService, article *service.ArticleService, logger *logger.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
			logging.Server(logger),
			metadata.Server(),
			auth.NewAuthServer(),
			cors.New(),
		),
		http.ResponseEncoder(encode.ResponseEncoder),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterHeartbeatHTTPServer(srv, heartbeat)
	articleV1.RegisterArticleHTTPServer(srv, article)
	return srv
}
