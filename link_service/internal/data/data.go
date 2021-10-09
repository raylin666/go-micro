package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/raylin666/go-utils/database/gorm"
	"link_service/internal/conf"
	"link_service/internal/data/model"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewShortLinkRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *gorm.DB

	model struct{
		LinkRelation *model.LinkRelationModel
	}
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	dataInstance := &Data{}

	// 初始化数据库
	db, err := gorm.New(
		gorm.Options{
			Driver: c.Database.GetDriver(),
			Dsn: c.Database.GetDsn(),
		}, gorm.PluginConfig{})
	if err != nil {
		return dataInstance, func() {
			log.NewHelper(logger).Info("Database connection failed")
		}, err
	}

	cleanup := func() {
		err = db.Close()
		if err != nil {
			log.NewHelper(logger).Errorf("closing data resources error: %v", err)
		} else {
			log.NewHelper(logger).Info("closing the data resources")
		}
	}

	dataInstance.db = db

	// Model 实例
	dataInstance.model.LinkRelation = model.NewLinkRelationModel(db)

	return dataInstance, cleanup, nil
}
