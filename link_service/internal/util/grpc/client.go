package grpc

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

	instanceGrpcClientConn *grpcClientConn
)

type grpcClientConn struct {
	clientConn map[string]*grpc.ClientConn

	uuidClient uuid.UuidClient
}

func NewGRPCClientConn(registry *registry.Registry) *grpcClientConn {
	grpc_client_conn := &grpcClientConn{
		clientConn: make(map[string]*grpc.ClientConn),
	}

	if endpoint != nil {
		for serverId, serverName := range endpoint {
			conn, err := kratos_grpc.DialInsecure(context.Background(), kratos_grpc.WithEndpoint(serverName), kratos_grpc.WithDiscovery(registry))
			if err != nil {
				panic(err)
			}

			grpc_client_conn.clientConn[serverId] = conn

			switch serverId {
			case SERVER_ID_UUID:
				grpc_client_conn.uuidClient = uuid.NewUuidClient(conn)
				break
			}
		}
	}

	instanceGrpcClientConn = grpc_client_conn
	return grpc_client_conn
}

func (g *grpcClientConn) getClient(serverId string) *grpc.ClientConn {
	return g.clientConn[serverId]
}

/**
	获取生成唯一标识服务
 */
func (g *grpcClientConn) GetUuidClient() uuid.UuidClient {
	return g.uuidClient
}

/**
	获取 GRPC 客户端连接实例容器
 */
func GRPCClientConn() *grpcClientConn {
	return instanceGrpcClientConn
}

/**
	GRPC 客户端关闭连接
 */
func (g *grpcClientConn) GRPCClientConnClose() {
	for serverId, _ := range g.clientConn {
		g.getClient(serverId).Close()
	}
}

