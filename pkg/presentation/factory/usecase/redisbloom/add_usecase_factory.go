package redisbloom

import (
	redis_bloom_go "github.com/RedisBloom/redisbloom-go"
	"github.com/danielchiovitti/ballot-box/pkg/domain/usecase/redisbloom"
)

func NewAddUseCaseFactory(redisBloomClient *redis_bloom_go.Client) AddUseCaseFactoryInterface {
	return &AddUseCaseFactory{
		redisBloomClient,
	}
}

type AddUseCaseFactory struct {
	RedisBloomClient *redis_bloom_go.Client
}

func (a *AddUseCaseFactory) Build() redisbloom.AddUseCaseInterface {
	return &redisbloom.AddUseCase{
		RedisBloomClient: a.RedisBloomClient,
	}
}
