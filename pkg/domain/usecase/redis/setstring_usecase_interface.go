package redis

import "context"

type SetStringUseCaseInterface interface {
	Execute(ctx context.Context, key string, value string, exp int) error
}
