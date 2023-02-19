package cache

import (
	"context"
	"errors"
	"fmt"
	utils_redis "github.com/raylin666/go-utils/cache/redis"
	"mt/config"
	"time"
)

var _ Redis = (*redis)(nil)

type Redis interface {
	Get() utils_redis.Client
	Close() error
}

type redis struct {
	client utils_redis.Client
}

func NewRedis(name string, config *config.RedisItem) (Redis, error) {
	var rds = new(redis)
	opts := new(utils_redis.Options)
	opts.Addr = fmt.Sprintf("%s:%d", config.Addr, config.Port)
	opts.Network = config.Network
	opts.Username = config.Username
	opts.Password = config.Password
	opts.DB = int(config.Db)
	opts.DialTimeout = time.Duration(config.DialTimeout)
	opts.IdleTimeout = time.Duration(config.IdleTimeout)
	opts.MaxConnAge = time.Duration(config.MaxConnAge)
	opts.MaxRetries = int(config.MaxRetries)
	opts.IdleCheckFrequency = time.Duration(config.IdleCheckFrequency)
	opts.MaxRetryBackoff = time.Duration(config.MinRetryBackoff)
	opts.MinRetryBackoff = time.Duration(config.MinRetryBackoff)
	opts.MinIdleConns = int(config.MinIdleConns)
	opts.WriteTimeout = time.Duration(config.WriteTimeout)
	opts.ReadTimeout = time.Duration(config.ReadTimeout)
	opts.PoolFIFO = config.PoolFifo
	opts.PoolSize = int(config.PoolSize)
	opts.PoolTimeout = time.Duration(config.PoolTimeout)

	client, err := utils_redis.NewClient(context.TODO(), opts)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("new redis to %s client err", name))
	}

	rds.client = client

	return rds, nil
}

func (rds *redis) Get() utils_redis.Client {
	return rds.client
}

func (rds *redis) Close() error {
	return rds.Get().Close()
}
