//go:build wireinject
// +build wireinject

package ballot_box

import (
	"github.com/danielchiovitti/ballot-box/pkg/presentation"
	"github.com/danielchiovitti/ballot-box/pkg/presentation/route"
	"github.com/google/wire"
)

var superSet = wire.NewSet(
	presentation.NewHandler,
	route.NewHealthRoute,
	route.NewVotingRoute,
)

func InitializeHandler() *presentation.Handler {
	wire.Build(
		superSet,
	)
	return &presentation.Handler{}
}
