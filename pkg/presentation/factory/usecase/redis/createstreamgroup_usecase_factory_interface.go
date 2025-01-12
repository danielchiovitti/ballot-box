package redis

import "github.com/danielchiovitti/ballot-box/pkg/domain/usecase/redis"

type CreateStreamGroupUseCaseFactoryInterface interface {
	Build() redis.CreateStreamGroupUseCaseInterface
}
