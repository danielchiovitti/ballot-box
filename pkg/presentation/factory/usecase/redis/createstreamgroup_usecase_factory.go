package redis

import (
	redis2 "github.com/danielchiovitti/ballot-box/pkg/domain/usecase/redis"
	"github.com/redis/go-redis/v9"
)

func NewCreateStreamGroupUseCaseFactory(redisClient *redis.Client) CreateStreamGroupUseCaseFactoryInterface {
	return &CreateStreamGroupUseCaseFactory{
		redisClient: redisClient,
	}
}

type CreateStreamGroupUseCaseFactory struct {
	redisClient *redis.Client
}

func (c *CreateStreamGroupUseCaseFactory) Build() redis2.CreateStreamGroupUseCaseInterface {
	return &redis2.CreateStreamGroupUseCase{
		RedisClient: c.redisClient,
	}
}
