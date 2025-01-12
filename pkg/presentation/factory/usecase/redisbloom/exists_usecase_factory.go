package redisbloom

import (
	redis_bloom_go "github.com/RedisBloom/redisbloom-go"
	"github.com/danielchiovitti/ballot-box/pkg/domain/usecase/redisbloom"
)

func NewExistsUseCaseFactory(redisBloomClient *redis_bloom_go.Client) ExistsUseCaseFactoryInterface {
	return &ExistsUseCaseFactory{
		redisBloomClient,
	}
}

type ExistsUseCaseFactory struct {
	RedisBloomClient *redis_bloom_go.Client
}

func (e *ExistsUseCaseFactory) Build() redisbloom.ExistsUseCaseInterface {
	return &redisbloom.ExistsUseCase{
		RedisBloomClient: e.RedisBloomClient,
	}
}
