package redis

import (
	redis2 "github.com/danielchiovitti/ballot-box/pkg/domain/usecase/redis"
	"github.com/redis/go-redis/v9"
)

func NewAddToStreamUseCaseFactory(redisClient *redis.Client) AddToStreamUseCaseFactoryInterface {
	return &AddToStreamUseCaseFactory{
		redisClient: redisClient,
	}
}

type AddToStreamUseCaseFactory struct {
	redisClient *redis.Client
}

func (a *AddToStreamUseCaseFactory) Build() redis2.AddToStreamUseCaseInterface {
	return &redis2.AddToStreamUseCase{
		RedisClient: a.redisClient,
	}
}
