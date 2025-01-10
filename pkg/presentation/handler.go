package presentation

import (
	"github.com/danielchiovitti/ballot-box/pkg/presentation/route"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	healthRoute *route.HealthRoute
	votingRoute *route.VotingRoute
}

func NewHandler(
	healthRoute *route.HealthRoute,
	votingRoute *route.VotingRoute,
) *Handler {
	return &Handler{
		healthRoute: healthRoute,
		votingRoute: votingRoute,
	}
}

func (h *Handler) GetRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Mount("/health", h.healthRoute.HealthRoutes())
	r.Mount("/voting", h.votingRoute.VotingRoutes())
	return r
}
