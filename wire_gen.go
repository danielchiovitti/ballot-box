// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package ballot_box

import (
	"github.com/RedisBloom/redisbloom-go"
	"github.com/danielchiovitti/ballot-box/pkg/database/provider"
	"github.com/danielchiovitti/ballot-box/pkg/presentation"
	"github.com/danielchiovitti/ballot-box/pkg/presentation/factory/usecase/redis"
	"github.com/danielchiovitti/ballot-box/pkg/presentation/factory/usecase/redisbloom"
	"github.com/danielchiovitti/ballot-box/pkg/presentation/middleware"
	"github.com/danielchiovitti/ballot-box/pkg/presentation/route"
	"github.com/danielchiovitti/ballot-box/pkg/shared"
	"github.com/google/wire"
	redis2 "github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

// Injectors from wire.go:

func InitializeHandler() *presentation.Handler {
	healthRoute := route.NewHealthRoute()
	viper := NewViper()
	redisProvider := provider.NewRedisProvider(viper)
	client := NewRedisClient(redisProvider)
	incrUseCaseFactoryInterface := redis.NewIncrUseCaseFactory(client)
	expireUseCaseFactoryInterface := redis.NewExpireUseCaseFactory(client)
	configInterface := shared.NewConfig(viper)
	ratingMiddleware := middleware.NewRatingMiddleware(incrUseCaseFactoryInterface, expireUseCaseFactoryInterface, configInterface)
	backPressureMiddleware := middleware.NewBackPressureMiddleware(incrUseCaseFactoryInterface, expireUseCaseFactoryInterface, configInterface)
	votingRoute := route.NewVotingRoute(ratingMiddleware, backPressureMiddleware)
	redis_bloom_goClient := NewRedisBloomClient(redisProvider)
	reserveUseCaseFactoryInterface := redisbloom.NewReserveUseCaseFactory(redis_bloom_goClient)
	handler := presentation.NewHandler(healthRoute, votingRoute, viper, configInterface, reserveUseCaseFactoryInterface)
	return handler
}

// wire.go:

func NewViper() *viper.Viper {
	v := viper.New()
	v.SetEnvPrefix("bb")
	v.AutomaticEnv()
	return v
}

func NewRedisClient(r *provider.RedisProvider) *redis2.Client {
	res, _ := r.GetRedisClient()
	return res
}

func NewRedisBloomClient(r *provider.RedisProvider) *redis_bloom_go.Client {
	res, _ := r.GetRedisBloomClient()
	return res
}

var superSet = wire.NewSet(
	NewViper, shared.NewConfig, provider.NewRedisProvider, NewRedisClient,
	NewRedisBloomClient, provider.NewRedisBloomProvider, middleware.NewRatingMiddleware, middleware.NewBackPressureMiddleware, presentation.NewHandler, route.NewHealthRoute, route.NewVotingRoute, provider.NewMongoDbProvider, redis.NewIncrUseCaseFactory, redis.NewExpireUseCaseFactory, redisbloom.NewReserveUseCaseFactory, redisbloom.NewAddUseCaseFactory, redisbloom.NewExistsUseCaseFactory, redis.NewGetUseCaseFactory, redis.NewSetUseCaseFactory,
)
