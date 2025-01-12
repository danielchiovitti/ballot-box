package redisbloom

type ExistsUseCaseInterface interface {
	Execute(filterName, value string) (bool, error)
}
