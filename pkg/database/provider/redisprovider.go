package provider

import (
	"fmt"
	"github.com/danielchiovitti/ballot-box/pkg/domain/model"
	"github.com/danielchiovitti/ballot-box/pkg/shared"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"sync"
)

var lockRedisProvider sync.Mutex
var lockRedisClient sync.Mutex
var redisProviderInstance *RedisProvider
var redisClientInstance *redis.Client

type RedisProvider struct {
	*model.RedisOptions
}

func NewRedisProvider(v *viper.Viper) *RedisProvider {
	if redisProviderInstance == nil {
		lockRedisProvider.Lock()
		defer lockRedisProvider.Unlock()
		if redisProviderInstance == nil {
			opts := []model.RedisOptionsFunc{
				WithRedisAddress(v.GetString(string(shared.REDIS_ADDRESS))),
				WithRedisPort(v.GetInt(string(shared.REDIS_PORT))),
				WithRedisPassword(v.GetString(string(shared.REDIS_PASSWORD))),
				WithRedisDb(v.GetInt(string(shared.REDIS_DATABASE))),
				WithRedisPoolSize(v.GetInt(string(shared.REDIS_POOL_SIZE))),
				WithRedisProtocol(v.GetInt(string(shared.REDIS_PROTOCOL))),
			}
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

func WithRedisPort(port int) model.RedisOptionsFunc {
	return func(opt *model.RedisOptions) {
		opt.Port = port
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
				Addr:     fmt.Sprintf("%s:%d", p.Address, p.Port),
				Password: p.Password,
				DB:       p.Db,
				Protocol: p.Protocol,
				PoolSize: p.PoolSize,
			})
		}
	}

	return redisClientInstance, nil
}
