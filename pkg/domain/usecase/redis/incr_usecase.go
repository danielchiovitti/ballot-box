package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type IncrUseCase struct {
	RedisClient *redis.Client
}

func (i *IncrUseCase) Execute(ctx context.Context, key string) (int, error) {
	count, err := i.RedisClient.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
