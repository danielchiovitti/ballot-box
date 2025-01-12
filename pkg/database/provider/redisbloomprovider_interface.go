package provider

import "github.com/RedisBloom/redisbloom-go"

type RedisBloomProviderInterface interface {
	GetRedisBloomClient() (*redis_bloom_go.Client, error)
}
