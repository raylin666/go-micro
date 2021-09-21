package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"link_service/internal/conf"
	"link_service/internal/util/grpc"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewShortLinkRepo)

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		// GRPC 客户端关闭连接
		grpc.GRPCClientConn().GRPCClientConnClose()
		log.NewHelper(logger).Info("closing the grpc client connection resources")
	}
	return &Data{}, cleanup, nil
}
