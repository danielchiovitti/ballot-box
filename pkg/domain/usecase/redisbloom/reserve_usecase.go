package redisbloom

import redis_bloom_go "github.com/RedisBloom/redisbloom-go"

type ReserveUseCase struct {
	RedisBloomClient *redis_bloom_go.Client
}

func (r *ReserveUseCase) Execute(name string, precision float64, initialCapacity uint64) error {
	err := r.RedisBloomClient.Reserve(name, precision, initialCapacity)
	if err != nil {
		return err
	}
	return nil
}
