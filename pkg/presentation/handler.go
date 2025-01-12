package presentation

import (
	"github.com/danielchiovitti/ballot-box/pkg/presentation/factory/usecase/redisbloom"
	"github.com/danielchiovitti/ballot-box/pkg/presentation/route"
	"github.com/danielchiovitti/ballot-box/pkg/shared"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	"sync"
)

var handlerLock sync.Mutex
var handlerInstance *Handler

type Handler struct {
	healthRoute           *route.HealthRoute
	votingRoute           *route.VotingRoute
	viper                 *viper.Viper
	config                shared.ConfigInterface
	reserveUseCaseFactory redisbloom.ReserveUseCaseFactoryInterface
}

func NewHandler(
	healthRoute *route.HealthRoute,
	votingRoute *route.VotingRoute,
	viper *viper.Viper,
	config shared.ConfigInterface,
	reserveUseCaseFactory redisbloom.ReserveUseCaseFactoryInterface,
) *Handler {
	if handlerInstance == nil {
		handlerLock.Lock()
		defer handlerLock.Unlock()
		if handlerInstance == nil {
			handlerInstance = &Handler{
				healthRoute:           healthRoute,
				votingRoute:           votingRoute,
				viper:                 viper,
				config:                config,
				reserveUseCaseFactory: reserveUseCaseFactory,
			}
		}
	}

	return handlerInstance
}

func (h *Handler) GetRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Mount("/health", h.healthRoute.HealthRoutes())
	r.Mount("/voting", h.votingRoute.VotingRoutes())
	return r
}

func (h *Handler) GetViper() *viper.Viper {
	return h.viper
}

func (h *Handler) SetBloomFilter() {
	reserveUseCase := h.reserveUseCaseFactory.Build()
	err := reserveUseCase.Execute(h.config.GetBloomName(), h.config.GetBloomPrecision(), uint64(h.config.GetBloomInitial()))
	if err != nil && err.Error() != "ERR item exists" {
		panic(err)
	}
}
