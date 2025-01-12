package redis

import "context"

type ExpireUseCaseInterface interface {
	Execute(ctx context.Context, key string, rateTime int) error
}
