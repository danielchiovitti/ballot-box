package redis

import "github.com/danielchiovitti/ballot-box/pkg/domain/usecase/redis"

type IncrUseCaseFactoryInterface interface {
	Build() redis.IncrUseCaseInterface
}
