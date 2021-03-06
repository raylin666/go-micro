package main

import (
	"flag"
	"link_service/repositorie/pool"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"link_service/internal/conf"

	consul "github.com/go-kratos/consul/registry"
	"github.com/hashicorp/consul/api"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server) *kratos.App {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	registry := consul.New(client)

	// 初始化依赖 GRPC 服务客户端连接池
	pool.NewGRPCClientPool(registry, logger)

	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
		kratos.Registrar(registry),
	)
}

func main() {
	flag.Parse()

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	// 初始化配置保存
	conf.NewStore(&bc)

	id = conf.GetStore().GetService().GetId()
	Name = conf.GetStore().GetService().GetName()
	Version = conf.GetStore().GetService().GetVersion()

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace_id", log.TraceID(),
		"span_id", log.SpanID(),
	)

	app, cleanup, err := initApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer func() {
		// GRPC 客户端连接关闭
		pool.CloseGRPCClientPool(logger)
		log.NewHelper(logger).Info("closing the grpc client connection resources")

		cleanup()
	}()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
