package redis

import "context"

type CreateStreamGroupUseCaseInterface interface {
	Execute(ctx context.Context, streamName, groupName string) error
}
