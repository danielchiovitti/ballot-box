//go:build wireinject
// +build wireinject

package ballot_box

import (
	redis_bloom_go "github.com/RedisBloom/redisbloom-go"
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
	NewViper,
	shared.NewConfig,
	provider.NewRedisProvider,
	NewRedisClient,
	provider.NewRedisBloomProvider,
	middleware.NewRatingMiddleware,
	middleware.NewBackPressureMiddleware,
	presentation.NewHandler,
	route.NewHealthRoute,
	route.NewVotingRoute,
	//a linha abaixo está comentada para mostrar em casos que você não quer criar uma função de provider,
	//como poderia ser feito um binding entre a interface e a struct
	//wire.Bind(new(provider.MongoDbProviderInterface), new(*provider.MongoDbProvider)),
	provider.NewMongoDbProvider,
	redis.NewIncrUseCaseFactory,
	redis.NewExpireUseCaseFactory,
	redisbloom.NewReserveUseCaseFactory,
	redisbloom.NewAddUseCaseFactory,
	redisbloom.NewExistsUseCaseFactory,
	redis.NewGetUseCaseFactory,
	redis.NewSetUseCaseFactory,
)

func InitializeHandler() *presentation.Handler {
	wire.Build(
		superSet,
	)
	return &presentation.Handler{}
}
