package redis

import (
	"context"
)

type IncrUseCaseInterface interface {
	Execute(ctx context.Context, key string) (int, error)
}
