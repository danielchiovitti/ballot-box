package redis

import "context"

type AddToStreamUseCaseInterface interface {
	Execute(ctx context.Context, streamName string, value interface{}) error
}
