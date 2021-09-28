package pool

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	uuid "github.com/raylin666/go-micro-protoc/services/uuid/v1"
	"google.golang.org/grpc/connectivity"

	"github.com/go-kratos/consul/registry"
	kratos_grpc "github.com/go-kratos/kratos/v2/transport/grpc"
	grpcpool "github.com/raylin666/go-utils/pool/grpc"
	"google.golang.org/grpc"
)

const (
	SERVER_ID_UUID = "uuid_1"
)

var (
	endpoint = map[string]string{
		SERVER_ID_UUID: "discovery:///micro.uuid.service",
	}

	pool map[string]*grpcpool.Pool
)

func NewGRPCClientPool(registry *registry.Registry, logger log.Logger) map[string]*grpcpool.Pool {
	l := log.NewHelper(logger)
	if endpoint != nil {
		pool = make(map[string]*grpcpool.Pool)
		for serverId, serverName := range endpoint {
			p, err := grpcpool.New(func() (*grpc.ClientConn, error) {
				return kratos_grpc.DialInsecure(context.Background(), kratos_grpc.WithEndpoint(serverName), kratos_grpc.WithDiscovery(registry))
			}, 10, 100, 1)

			if err != nil {
				l.Errorf("The pool returned an error: %s", err.Error())
			}

			pool[serverId] = p
		}
	}

	return pool
}

func CloseGRPCClientPool(logger log.Logger)  {
	l := log.NewHelper(logger)
	for _, p := range pool {
		c, _ := p.Get(context.Background())
		cc := c.ClientConn
		if err := c.Close(); err != nil {
			l.Errorf("Close returned an error: %s", err.Error())
		}

		// Close pool should close all underlying gRPC client connections
		p.Close()

		if cc.GetState() != connectivity.Shutdown {
			l.Errorf("Returned connection was not closed, underlying connection is not in shutdown state")
		}
	}
}

func getPool(serverId string) *grpcpool.Pool {
	return pool[serverId]
}

// GetUuidGRPCClientPool 获取 唯一标识服务 客户端连接池
func GetUuidGRPCClientPool() (uuid.UuidClient, error) {
	conn, err := getPool(SERVER_ID_UUID).Get(context.Background())
	if err != nil {
		return nil, err
	}

	return uuid.NewUuidClient(conn), nil
}