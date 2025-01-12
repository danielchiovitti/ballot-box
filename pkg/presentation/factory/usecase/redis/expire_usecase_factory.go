package redis

import (
	redis2 "github.com/danielchiovitti/ballot-box/pkg/domain/usecase/redis"
	"github.com/redis/go-redis/v9"
)

func NewExpireUseCaseFactory(redisClient *redis.Client) ExpireUseCaseFactoryInterface {
	return &ExpireUseCaseFactory{
		redisClient: redisClient,
	}
}

type ExpireUseCaseFactory struct {
	redisClient *redis.Client
}

func (e *ExpireUseCaseFactory) Build() redis2.ExpireUseCaseInterface {
	return &redis2.ExpireUseCase{
		RedisClient: e.redisClient,
	}
}
