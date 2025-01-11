package redis

import "github.com/danielchiovitti/ballot-box/pkg/domain/usecase/redis"

type SetStringUseCaseFactoryInterface interface {
	Build() redis.SetStringUseCaseInterface
}
