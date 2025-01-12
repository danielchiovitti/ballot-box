package redis

import "context"

type SetUseCaseInterface interface {
	Execute(ctx context.Context, key string, value interface{}, exp int) error
}
