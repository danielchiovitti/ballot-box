package redisbloom

import (
	redis_bloom_go "github.com/RedisBloom/redisbloom-go"
	"github.com/danielchiovitti/ballot-box/pkg/domain/usecase/redisbloom"
)

func NewReserveUseCaseFactory(redisBloomClient *redis_bloom_go.Client) ReserveUseCaseFactoryInterface {
	return &ReserveUseCaseFactory{
		redisBloomClient,
	}
}

type ReserveUseCaseFactory struct {
	RedisBloomClient *redis_bloom_go.Client
}

func (r *ReserveUseCaseFactory) Build() redisbloom.ReserveUseCaseInterface {
	return &redisbloom.ReserveUseCase{
		RedisBloomClient: r.RedisBloomClient,
	}
}
