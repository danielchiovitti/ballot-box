package redisbloom

import "github.com/danielchiovitti/ballot-box/pkg/domain/usecase/redisbloom"

type ExistsUseCaseFactoryInterface interface {
	Build() redisbloom.ExistsUseCaseInterface
}
