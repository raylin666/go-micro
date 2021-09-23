package server

import (
	"context"
	nethttp "net/http"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	"github.com/gorilla/mux"
	pb "github.com/raylin666/go-micro-protoc/link/v1"
	"link_service/internal/conf"
	"link_service/internal/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.ShortLinkService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
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

			reply, err :=  greeter.ShortUrlToLongUrl(context.Background(), &pb.ShortUrlToLongUrlRequest{Url: ident})
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
	return srv
}
