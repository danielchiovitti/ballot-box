package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type SetUseCase struct {
	RedisClient *redis.Client
}

func (s *SetUseCase) Execute(ctx context.Context, key string, value interface{}, exp int) error {
	err := s.RedisClient.Set(ctx, key, value, time.Duration(exp)*time.Millisecond).Err()
	if err != nil {
		return err
	}
	return nil
}
