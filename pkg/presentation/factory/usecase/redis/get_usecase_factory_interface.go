package redis

import "github.com/danielchiovitti/ballot-box/pkg/domain/usecase/redis"

type GetUseCaseFactoryInterface interface {
	Build() redis.GetUseCaseInterface
}
