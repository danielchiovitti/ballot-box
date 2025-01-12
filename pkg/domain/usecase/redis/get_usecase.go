package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type GetUseCase struct {
	RedisClient *redis.Client
}

func (g *GetUseCase) Execute(ctx context.Context, key string) (interface{}, error) {
	v, err := g.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return v, nil
}
