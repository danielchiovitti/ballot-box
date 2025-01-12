package redis

import "github.com/danielchiovitti/ballot-box/pkg/domain/usecase/redis"

type ExpireUseCaseFactoryInterface interface {
	Build() redis.ExpireUseCaseInterface
}
