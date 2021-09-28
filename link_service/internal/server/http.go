package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	"github.com/gorilla/mux"
	pb "github.com/raylin666/go-micro-protoc/services/link/v1"
	"link_service/internal/conf"
	"link_service/internal/service"
	nethttp "net/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.ShortLinkService, logger log.Logger) *http.Server {
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

		r := mux.NewRouter()
		// 短链接重定向跳转长链接地址
		r.HandleFunc("/g/{name}", func(writer nethttp.ResponseWriter, request *nethttp.Request) {
			var (
				ident string
			)

			raws := mux.Vars(request)
			for k, v := range raws {
				if k == "name" {
					ident = v
				}
			}

			if len(ident) <= 0 {
				writer.WriteHeader(404)
				return
			}

			reply, err :=  greeter.TransformLongUrl(context.Background(), &pb.TransformLongUrlRequest{Url: ident})
			if err != nil {
				writer.WriteHeader(500)
				return
			}

			nethttp.Redirect(writer, request, reply.GetUrl(), 302)

		}).Methods("GET")
		srv.HandlePrefix("/g/", r)

		return srv
	}

	srv := srvHandler()
	pb.RegisterShortLinkHTTPServer(srv, greeter)

	log.NewHelper(logger).Info("HTTP service started successfully")

	return srv
}
