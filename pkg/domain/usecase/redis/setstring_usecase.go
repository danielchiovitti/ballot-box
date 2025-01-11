package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type SetStringUseCase struct {
	RedisClient *redis.Client
}

func (s *SetStringUseCase) Execute(ctx context.Context, key string, value string, exp int) error {
	err := s.RedisClient.Set(ctx, key, value, time.Duration(exp)*time.Millisecond).Err()
	if err != nil {
		return err
	}
	return nil
}
