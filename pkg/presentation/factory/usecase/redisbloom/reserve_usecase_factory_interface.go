package redisbloom

import "github.com/danielchiovitti/ballot-box/pkg/domain/usecase/redisbloom"

type ReserveUseCaseFactoryInterface interface {
	Build() redisbloom.ReserveUseCaseInterface
}
