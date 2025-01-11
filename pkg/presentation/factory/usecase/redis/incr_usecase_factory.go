package redis

import (
	redis2 "github.com/danielchiovitti/ballot-box/pkg/domain/usecase/redis"
	"github.com/redis/go-redis/v9"
)

func NewIncrUseCaseFactory(redisClient *redis.Client) IncrUseCaseFactoryInterface {
	return &IncrUseCaseFactory{
		redisClient: redisClient,
	}
}

type IncrUseCaseFactory struct {
	redisClient *redis.Client
}

func (i *IncrUseCaseFactory) Build() redis2.IncrUseCaseInterface {
	return &redis2.IncrUseCase{
		RedisClient: i.redisClient,
	}
}
