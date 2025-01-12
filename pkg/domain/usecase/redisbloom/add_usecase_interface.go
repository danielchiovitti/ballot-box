package redisbloom

type AddUseCaseInterface interface {
	Execute(filterName, value string) error
}
