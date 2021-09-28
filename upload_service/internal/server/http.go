package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	pb "github.com/raylin666/go-micro-protoc/services/upload/v1"
	"upload_service/internal/conf"
	"upload_service/internal/service"
	"upload_service/repositorie/middleware/validate"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.UploadService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			// Recovery 中间件用于异常恢复，服务出现异常的情况下，防止程序直接退出
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
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	srvHandler := func() *http.Server {
		srv := http.NewServer(opts...)

		// swagger api router	   ---     /q/swagger-ui/
		srv.HandlePrefix("/q/", openapiv2.NewHandler())

		return srv
	}

	srv := srvHandler()

	pb.RegisterUploadHTTPServer(srv, greeter)

	log.NewHelper(logger).Info("HTTP service started successfully")

	return srv
}
