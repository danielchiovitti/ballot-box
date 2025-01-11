package model

type RedisOptions struct {
	Address  string
	Password string
	Db       int
	Protocol int
	PoolSize int
}

type RedisOptionsFunc func(option *RedisOptions)
