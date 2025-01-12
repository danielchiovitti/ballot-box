package presentation

import (
	"context"
	"github.com/danielchiovitti/ballot-box/pkg/presentation/factory/usecase/redis"
	"github.com/danielchiovitti/ballot-box/pkg/presentation/route"
	"github.com/danielchiovitti/ballot-box/pkg/shared"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	"log"
	"sync"
	"time"
)

var handlerLock sync.Mutex
var handlerInstance *Handler

type Handler struct {
	healthRoute       *route.HealthRoute
	votingRoute       *route.VotingRoute
	viper             *viper.Viper
	setUseCaseFactory redis.SetUseCaseFactoryInterface
	config            shared.ConfigInterface
}

func NewHandler(
	healthRoute *route.HealthRoute,
	votingRoute *route.VotingRoute,
	viper *viper.Viper,
	setUseCaseFactory redis.SetUseCaseFactoryInterface,
	config shared.ConfigInterface,
) *Handler {
	if handlerInstance == nil {
		handlerLock.Lock()
		defer handlerLock.Unlock()
		if handlerInstance == nil {
			handlerInstance = &Handler{
				healthRoute:       healthRoute,
				votingRoute:       votingRoute,
				viper:             viper,
				setUseCaseFactory: setUseCaseFactory,
				config:            config,
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

func (h *Handler) SetCache() {
	useCase := h.setUseCaseFactory.Build()
	timeOut := h.config.GetTimeOut()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Millisecond)
	defer cancel()

	err := useCase.Execute(ctx, string(shared.TIMEOUT), h.config.GetTimeOut(), h.config.GetTimeOut())
	if err != nil {
		log.Fatalf(err.Error())
	}

}
