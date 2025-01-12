package model

type RedisBloomOptions struct {
	Address  string
	Password string
	Db       int
	Protocol int
	PoolSize int
}

type RedisBloomOptionsFunc func(option *RedisBloomOptions)
