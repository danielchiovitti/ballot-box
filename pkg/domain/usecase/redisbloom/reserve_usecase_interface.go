package redisbloom

type ReserveUseCaseInterface interface {
	Execute(name string, precision float64, initialCapacity uint64) error
}
