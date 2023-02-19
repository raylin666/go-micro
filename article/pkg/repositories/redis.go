package repositories

import "mt/pkg/cache"

const (
	REDIS_CONNECTION_DEFAULT_NAME = "default"
)

var _ RedisRepo = (*redisRepo)(nil)

type RedisRepo interface {
	Redis(name string) cache.Redis
}

type redisRepo struct {
	resource map[string]cache.Redis
}

func (repo *redisRepo) Redis(name string) cache.Redis {
	return repo.resource[name]
}
