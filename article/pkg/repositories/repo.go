package repositories

import (
	"fmt"
	"go.uber.org/zap"
	"mt/config"
	"mt/pkg/cache"
	"mt/pkg/db"
	"mt/pkg/logger"
)

var _ DataRepo = (*dataRepo)(nil)

type DataRepo interface {
	DB(name string) db.Db
	DbRepo() DbRepo
	Redis(name string) cache.Redis
	RedisRepo() RedisRepo
}

type dataRepo struct {
	db    *dbRepo
	redis *redisRepo
}

func NewDataRepo(logger *logger.Logger, conf *config.Data) DataRepo {
	var dbRepo = new(dbRepo)
	var redisRepo = new(redisRepo)
	var repo = new(dataRepo)

	var (
		dbMap    = make(map[string]*config.DatabaseItem, 1)
		redisMap = make(map[string]*config.RedisItem, 1)
	)

	// 初始化数据库
	var database = conf.GetDatabase()
	if database.GetDefault() != nil {
		dbMap[DB_CONNECTION_DEFAULT_NAME] = database.GetDefault()
	}

	lenDatabase := len(dbMap)
	if lenDatabase > 0 {
		dbRepo.resource = make(map[string]db.Db, lenDatabase)
		for dbName, dbConfig := range dbMap {
			rdb, err := db.NewDb(dbName, dbConfig, logger)
			if err != nil {
				logger.UseApp().Error(fmt.Sprintf("init db.repo %s error", dbName), zap.Error(err))
			} else {
				logger.UseApp().Info(fmt.Sprintf("init db.repo %s success", dbName))
				dbRepo.resource[dbName] = rdb
			}
		}

		repo.db = dbRepo
	} else {
		logger.UseApp().Warn("Currently not db.repo connected.")
	}

	// 初始化 Redis
	var redis = conf.GetRedis()
	if redis.GetDefault() != nil {
		redisMap[REDIS_CONNECTION_DEFAULT_NAME] = redis.GetDefault()
	}

	lenRedis := len(redisMap)
	if lenRedis > 0 {
		redisRepo.resource = make(map[string]cache.Redis, lenRedis)
		for redisName, redisConfig := range redisMap {
			redis, err := cache.NewRedis(redisName, redisConfig)
			if err != nil {
				logger.UseApp().Error(fmt.Sprintf("init redis.repo %s error", redisName), zap.Error(err))
			} else {
				logger.UseApp().Info(fmt.Sprintf("init redis.repo %s success", redisName))
				redisRepo.resource[redisName] = redis
			}
		}

		repo.redis = redisRepo
	} else {
		logger.UseApp().Warn("Currently not redis.repo connected.")
	}

	return repo
}

func (repo *dataRepo) DB(name string) db.Db {
	return repo.db.resource[name]
}

func (repo *dataRepo) DbRepo() DbRepo {
	return repo.db
}

func (repo *dataRepo) Redis(name string) cache.Redis {
	return repo.redis.resource[name]
}

func (repo *dataRepo) RedisRepo() RedisRepo {
	return repo.redis
}
