package redis

import (
	redis2 "github.com/danielchiovitti/ballot-box/pkg/domain/usecase/redis"
	"github.com/redis/go-redis/v9"
)

func NewSetStringUseCaseFactory(redisClient *redis.Client) SetStringUseCaseFactoryInterface {
	return &SetStringUseCaseFactory{
		redisClient: redisClient,
	}
}

type SetStringUseCaseFactory struct {
	redisClient *redis.Client
}

func (s *SetStringUseCaseFactory) Build() redis2.SetStringUseCaseInterface {
	return &redis2.SetStringUseCase{
		RedisClient: s.redisClient,
	}
}
