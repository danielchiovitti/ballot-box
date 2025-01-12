package provider

import (
	"fmt"
	redis_bloom_go "github.com/RedisBloom/redisbloom-go"
	"github.com/danielchiovitti/ballot-box/pkg/domain/model"
	"github.com/danielchiovitti/ballot-box/pkg/shared"
	redis2 "github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"sync"
)

var lockRedisBloomProvider sync.Mutex
var lockRedisBloomClient sync.Mutex
var redisBloomProviderInstance *RedisBloomProvider
var redisBloomClientInstance *redis_bloom_go.Client

type RedisBloomProvider struct {
	*model.RedisBloomOptions
}

func NewRedisBloomProvider(v *viper.Viper) *RedisBloomProvider {
	if redisBloomProviderInstance == nil {
		lockRedisBloomProvider.Lock()
		defer lockRedisBloomProvider.Unlock()
		if redisBloomProviderInstance == nil {
			opts := []model.RedisBloomOptionsFunc{
				WithRedisBloomAddress(v.GetString(string(shared.REDIS_BLOOM_ADDRESS))),
				WithRedisBloomPort(v.GetInt(string(shared.REDIS_BLOOM_PORT))),
				WithRedisBloomPassword(v.GetString(string(shared.REDIS_BLOOM_PASSWORD))),
				WithRedisBloomDb(v.GetInt(string(shared.REDIS_BLOOM_DATABASE))),
				WithRedisBloomPoolSize(v.GetInt(string(shared.REDIS_BLOOM_POOL_SIZE))),
				WithRedisBloomProtocol(v.GetInt(string(shared.REDIS_BLOOM_PROTOCOL))),
			}
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

func WithRedisBloomPort(port int) model.RedisBloomOptionsFunc {
	return func(opt *model.RedisBloomOptions) {
		opt.Port = port
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
						fmt.Sprintf("%s:%d", p.Address, p.Port),
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
