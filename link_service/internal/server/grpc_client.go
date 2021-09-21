package server

import (
	"context"
	"github.com/go-kratos/consul/registry"
	kratos_grpc "github.com/go-kratos/kratos/v2/transport/grpc"
	uuid "github.com/raylin666/go-micro-protoc/uuid/v1"
	"google.golang.org/grpc"
)

const (
	SERVER_ID_UUID = "uuid_1"
)

var (
	endpoint = map[string]string{
		SERVER_ID_UUID: "discovery:///micro.uuid.service",
	}
)

type grpcClientConn struct {
	connPool map[string]*grpc.ClientConn

	uuidClient uuid.UuidClient
}

func NewGRPCClientConn(registry *registry.Registry) *grpcClientConn {
	grpc_client_conn := &grpcClientConn{
		connPool: make(map[string]*grpc.ClientConn),
	}

	if endpoint != nil {
		for serverId, serverName := range endpoint {
			conn, err := kratos_grpc.DialInsecure(context.Background(), kratos_grpc.WithEndpoint(serverName), kratos_grpc.WithDiscovery(registry))
			if err != nil {
				panic(err)
			}

			grpc_client_conn.connPool[serverId] = conn

			switch serverId {
			case SERVER_ID_UUID:
				grpc_client_conn.uuidClient = uuid.NewUuidClient(conn)
				break
			}
		}
	}

	return grpc_client_conn
}

func (g *grpcClientConn) getClient(serverId string) *grpc.ClientConn {
	return g.connPool[serverId]
}

/**
	获取生成唯一标识服务
 */
func (g *grpcClientConn) GetUuidClient() uuid.UuidClient {
	return g.uuidClient
}

/**
	GRPC 客户端关闭连接
 */
func (g *grpcClientConn) GRPCClientConnClose() {
	for serverId, _ := range g.connPool {
		g.getClient(serverId).Close()
	}
}

