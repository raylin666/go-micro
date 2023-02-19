package data

import (
	"fmt"
	"mt/config"
	"mt/pkg/logger"
	"mt/pkg/repositories"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	repositories.NewDataRepo,
	NewHeartbeatRepo,
	NewArticleRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	DbRepo    repositories.DbRepo
	RedisRepo repositories.RedisRepo
}

// NewData .
func NewData(c *config.Data, logger *logger.Logger, repo repositories.DataRepo) (*Data, func(), error) {
	cleanup := func() {
		// 资源关闭
		repo.DB(repositories.DB_CONNECTION_DEFAULT_NAME).Close()
		logger.UseApp().Info(fmt.Sprintf("closing the data resource: %s db.repo.", repositories.DB_CONNECTION_DEFAULT_NAME))
		repo.Redis(repositories.REDIS_CONNECTION_DEFAULT_NAME).Close()
		logger.UseApp().Info(fmt.Sprintf("closing the data resource: %s redis.repo.", repositories.REDIS_CONNECTION_DEFAULT_NAME))
	}

	return &Data{
		DbRepo:    repo.DbRepo(),
		RedisRepo: repo.RedisRepo(),
	}, cleanup, nil
}
