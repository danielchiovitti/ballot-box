package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type ExpireUseCase struct {
	RedisClient *redis.Client
}

func (e *ExpireUseCase) Execute(ctx context.Context, key string, rateLimit int) error {
	err := e.RedisClient.Expire(ctx, key, time.Duration(rateLimit)*time.Millisecond).Err()
	if err != nil {
		return err
	}
	return nil
}
