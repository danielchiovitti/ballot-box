package redis

import "github.com/danielchiovitti/ballot-box/pkg/domain/usecase/redis"

type SetUseCaseFactoryInterface interface {
	Build() redis.SetUseCaseInterface
}
