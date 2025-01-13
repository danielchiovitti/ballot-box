package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type AddToStreamUseCase struct {
	RedisClient *redis.Client
}

func (a *AddToStreamUseCase) Execute(ctx context.Context, streamName string, value interface{}) error {
	_, err := a.RedisClient.XAdd(ctx, &redis.XAddArgs{
		Stream: streamName,
		Values: map[string]interface{}{
			"vote": value,
		},
	}).Result()
	if err != nil {
		return err
	}
	return nil
}
