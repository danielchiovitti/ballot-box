package redisbloom

import redis_bloom_go "github.com/RedisBloom/redisbloom-go"

type AddUseCase struct {
	RedisBloomClient *redis_bloom_go.Client
}

func (a *AddUseCase) Execute(filterName, value string) error {
	_, err := a.RedisBloomClient.Add(filterName, value)
	if err != nil {
		return err
	}
	return nil
}
