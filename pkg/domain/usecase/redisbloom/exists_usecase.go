package redisbloom

import redis_bloom_go "github.com/RedisBloom/redisbloom-go"

type ExistsUseCase struct {
	RedisBloomClient *redis_bloom_go.Client
}

func (e *ExistsUseCase) Execute(filterName, value string) (bool, error) {
	exists, err := e.RedisBloomClient.Exists(filterName, value)
	if err != nil {
		return false, err
	}
	return exists, nil
}
