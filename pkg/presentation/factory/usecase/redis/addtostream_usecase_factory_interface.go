package redis

import "github.com/danielchiovitti/ballot-box/pkg/domain/usecase/redis"

type AddToStreamUseCaseFactoryInterface interface {
	Build() redis.AddToStreamUseCaseInterface
}
