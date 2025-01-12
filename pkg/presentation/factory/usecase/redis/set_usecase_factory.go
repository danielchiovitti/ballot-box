package redis

import (
	redis2 "github.com/danielchiovitti/ballot-box/pkg/domain/usecase/redis"
	"github.com/redis/go-redis/v9"
)

func NewSetUseCaseFactory(redisClient *redis.Client) SetUseCaseFactoryInterface {
	return &SetUseCaseFactory{
		redisClient: redisClient,
	}
}

type SetUseCaseFactory struct {
	redisClient *redis.Client
}

func (s *SetUseCaseFactory) Build() redis2.SetUseCaseInterface {
	return &redis2.SetUseCase{
		RedisClient: s.redisClient,
	}
}
