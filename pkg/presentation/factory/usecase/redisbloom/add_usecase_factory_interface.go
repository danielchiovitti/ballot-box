package redisbloom

import "github.com/danielchiovitti/ballot-box/pkg/domain/usecase/redisbloom"

type AddUseCaseFactoryInterface interface {
	Build() redisbloom.AddUseCaseInterface
}
