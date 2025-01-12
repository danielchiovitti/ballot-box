package provider

import (
	redis_bloom_go "github.com/RedisBloom/redisbloom-go"
	"github.com/danielchiovitti/ballot-box/pkg/domain/model"
	redis2 "github.com/gomodule/redigo/redis"
	"sync"
)

var lockRedisBloomProvider sync.Mutex
var lockRedisBloomClient sync.Mutex
var redisBloomProviderInstance *RedisBloomProvider
var redisBloomClientInstance *redis_bloom_go.Client

type RedisBloomProvider struct {
	*model.RedisBloomOptions
}

func NewRedisBloomProvider(opts ...model.RedisBloomOptionsFunc) *RedisBloomProvider {
	if redisBloomProviderInstance == nil {
		lockRedisBloomProvider.Lock()
		defer lockRedisBloomProvider.Unlock()
		if redisBloomProviderInstance == nil {
			o := RedisBloomDefaultOpts()
			for _, fn := range opts {
				fn(o)
			}
			redisBloomProviderInstance = &RedisBloomProvider{
				o,
			}
		}
	}
	return redisBloomProviderInstance
}

func RedisBloomDefaultOpts() *model.RedisBloomOptions {
	return &model.RedisBloomOptions{
		Address:  "localhost:6379",
		Password: "admin",
		Db:       0,
		Protocol: 2,
		PoolSize: 150,
	}
}

func WithRedisBloomAddress(address string) model.RedisBloomOptionsFunc {
	return func(opt *model.RedisBloomOptions) {
		opt.Address = address
	}
}

func WithRedisBloomPassword(password string) model.RedisBloomOptionsFunc {
	return func(opt *model.RedisBloomOptions) {
		opt.Password = password
	}
}

func WithRedisBloomDb(db int) model.RedisBloomOptionsFunc {
	return func(opt *model.RedisBloomOptions) {
		opt.Db = db
	}
}

func WithRedisBloomProtocol(protocol int) model.RedisBloomOptionsFunc {
	return func(opt *model.RedisBloomOptions) {
		opt.Protocol = protocol
	}
}

func WithRedisBloomPoolSize(poolSize int) model.RedisBloomOptionsFunc {
	return func(opt *model.RedisBloomOptions) {
		opt.PoolSize = poolSize
	}
}

func (p *RedisProvider) GetRedisBloomClient() (*redis_bloom_go.Client, error) {
	if redisBloomClientInstance != nil {
		return redisBloomClientInstance, nil
	}

	if redisBloomClientInstance == nil {
		lockRedisBloomClient.Lock()
		defer lockRedisBloomClient.Unlock()
		if redisBloomClientInstance == nil {
			pool := &redis2.Pool{
				MaxIdle:   p.PoolSize,
				MaxActive: p.PoolSize,
				Dial: func() (redis2.Conn, error) {
					return redis2.Dial(
						"tcp",
						p.Address,
						redis2.DialPassword(p.Password),
						redis2.DialDatabase(p.Db),
					)
				},
			}

			redisBloomClientInstance = redis_bloom_go.NewClientFromPool(pool, "bloom-client")
		}
	}

	return redisBloomClientInstance, nil
}
