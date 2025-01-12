package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type CreateStreamGroupUseCase struct {
	RedisClient *redis.Client
}

func (c *CreateStreamGroupUseCase) Execute(ctx context.Context, streamName, groupName string) error {
	err := c.RedisClient.XGroupCreateMkStream(ctx, streamName, groupName, "$").Err()
	if err != nil {
		return err
	}
	return nil
}
