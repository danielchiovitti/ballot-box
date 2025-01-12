//go:build wireinject
// +build wireinject

package ballot_box

import (
	"github.com/danielchiovitti/ballot-box/pkg/database/provider"
	"github.com/danielchiovitti/ballot-box/pkg/presentation"
	"github.com/danielchiovitti/ballot-box/pkg/presentation/factory/usecase/redis"
	"github.com/danielchiovitti/ballot-box/pkg/presentation/factory/usecase/redisbloom"
	"github.com/danielchiovitti/ballot-box/pkg/presentation/route"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func LoadEnvConfig() *viper.Viper {
	v := viper.New()
	v.SetEnvPrefix("bb")
	v.AutomaticEnv()
	return v
}

var superSet = wire.NewSet(
	LoadEnvConfig,
	provider.NewRedisProvider,
	presentation.NewHandler,
	route.NewHealthRoute,
	route.NewVotingRoute,
	//a linha abaixo está comentada para mostrar em casos que você não quer criar uma função de provider,
	//como poderia ser feito um binding entre a interface e a struct
	//wire.Bind(new(provider.MongoDbProviderInterface), new(*provider.MongoDbProvider)),
	provider.NewMongoDbProvider,
	redis.NewSetUseCaseFactory,
	redis.NewIncrUseCaseFactory,
	redis.NewExpireUseCaseFactory,
	redisbloom.NewReserveUseCaseFactory,
	redisbloom.NewAddUseCaseFactory,
	redisbloom.NewExistsUseCaseFactory,
	redis.NewGetUseCaseFactory,
)

func InitializeHandler() *presentation.Handler {
	wire.Build(
		superSet,
	)
	return &presentation.Handler{}
}
