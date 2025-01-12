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
)

var superSet = wire.NewSet(
	presentation.NewHandler,
	route.NewHealthRoute,
	route.NewVotingRoute,
	wire.Bind(new(provider.MongoDbProviderInterface), new(*provider.MongoDbProvider)),
	provider.NewMongoDbProvider,
	wire.Bind(new(provider.RedisProviderInterface), new(*provider.RedisProvider)),
	provider.NewRedisProvider,
	redis.NewSetStringUseCaseFactory,
	redis.NewIncrUseCaseFactory,
	redis.NewExpireUseCaseFactory,
	provider.NewRedisBloomProvider,
	redisbloom.NewReserveUseCaseFactory,
)

func InitializeHandler() *presentation.Handler {
	wire.Build(
		superSet,
	)
	return &presentation.Handler{}
}
