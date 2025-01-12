package presentation

import (
	"github.com/danielchiovitti/ballot-box/pkg/presentation/route"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	"sync"
)

var handlerLock sync.Mutex
var handlerInstance *Handler

type Handler struct {
	healthRoute *route.HealthRoute
	votingRoute *route.VotingRoute
	viper       *viper.Viper
}

func NewHandler(
	healthRoute *route.HealthRoute,
	votingRoute *route.VotingRoute,
	viper *viper.Viper,
) *Handler {
	if handlerInstance == nil {
		handlerLock.Lock()
		defer handlerLock.Unlock()
		if handlerInstance == nil {
			handlerInstance = &Handler{
				healthRoute: healthRoute,
				votingRoute: votingRoute,
				viper:       viper,
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
