package redis

import "context"

type GetUseCaseInterface interface {
	Execute(ctx context.Context, key string) (interface{}, error)
}
