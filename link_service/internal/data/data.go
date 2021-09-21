package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"link_service/internal/conf"
	"link_service/internal/util/grpc"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewShortLinkRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	db, err := gorm.Open(mysql.Open(c.Database.GetSource()), &gorm.Config{})
	if err != nil {
		return &Data{}, func() {
			log.NewHelper(logger).Info("Database connection failed")
		}, err
	}

	cleanup := func() {
		// GRPC 客户端关闭连接
		grpc.GRPCClientConn().GRPCClientConnClose()
		log.NewHelper(logger).Info("closing the grpc client connection resources")
	}

	return &Data{
		db: db,
	}, cleanup, nil
}
