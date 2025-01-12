package redis

import (
	redis2 "github.com/danielchiovitti/ballot-box/pkg/domain/usecase/redis"
	"github.com/redis/go-redis/v9"
)

func NewGetUseCaseFactory(redisClient *redis.Client) GetUseCaseFactoryInterface {
	return &GetUseCaseFactory{
		redisClient: redisClient,
	}
}

type GetUseCaseFactory struct {
	redisClient *redis.Client
}

func (g *GetUseCaseFactory) Build() redis2.GetUseCaseInterface {
	return &redis2.GetUseCase{
		RedisClient: g.redisClient,
	}
}
