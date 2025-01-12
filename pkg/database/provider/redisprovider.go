package provider

import (
	"github.com/danielchiovitti/ballot-box/pkg/domain/model"
	"github.com/redis/go-redis/v9"
	"sync"
)

var lockRedisProvider sync.Mutex
var lockRedisClient sync.Mutex
var redisProviderInstance *RedisProvider
var redisClientInstance *redis.Client

type RedisProvider struct {
	*model.RedisOptions
}

func NewRedisProvider(opts ...model.RedisOptionsFunc) *RedisProvider {
	if redisProviderInstance == nil {
		lockRedisProvider.Lock()
		defer lockRedisProvider.Unlock()
		if redisProviderInstance == nil {
			o := RedisDefaultOpts()
			for _, fn := range opts {
				fn(o)
			}

			redisProviderInstance = &RedisProvider{
				o,
			}
		}
	}
	return redisProviderInstance
}

func RedisDefaultOpts() *model.RedisOptions {
	return &model.RedisOptions{
		Address:  "localhost:6379",
		Password: "admin",
		Db:       0,
		Protocol: 2,
		PoolSize: 150,
	}
}

func WithRedisAddress(address string) model.RedisOptionsFunc {
	return func(opt *model.RedisOptions) {
		opt.Address = address
	}
}

func WithRedisPassword(password string) model.RedisOptionsFunc {
	return func(opt *model.RedisOptions) {
		opt.Password = password
	}
}

func WithRedisDb(db int) model.RedisOptionsFunc {
	return func(opt *model.RedisOptions) {
		opt.Db = db
	}
}

func WithRedisProtocol(protocol int) model.RedisOptionsFunc {
	return func(opt *model.RedisOptions) {
		opt.Protocol = protocol
	}
}

func WithRedisPoolSize(poolSize int) model.RedisOptionsFunc {
	return func(opt *model.RedisOptions) {
		opt.PoolSize = poolSize
	}
}

func (p *RedisProvider) GetRedisClient() (*redis.Client, error) {
	if redisClientInstance != nil {
		return redisClientInstance, nil
	}

	if redisClientInstance == nil {
		lockRedisClient.Lock()
		defer lockRedisClient.Unlock()
		if redisClientInstance == nil {
			redisClientInstance = redis.NewClient(&redis.Options{
				Addr:     p.Address,
				Password: p.Password,
				DB:       p.Db,
				Protocol: p.Protocol,
				PoolSize: p.PoolSize,
			})
		}
	}

	return redisClientInstance, nil
}
