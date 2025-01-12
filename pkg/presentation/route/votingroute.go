package route

import (
	"encoding/json"
	"github.com/danielchiovitti/ballot-box/pkg/presentation/middleware"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"sync"
)

var votingLock sync.Mutex
var votingRouteInstance *VotingRoute

type VotingRoute struct {
	ratingMiddleware       *middleware.RatingMiddleware
	backPressureMiddleware *middleware.BackPressureMiddleware
}

func NewVotingRoute(
	ratingMiddleware *middleware.RatingMiddleware,
	backPressureMiddleware *middleware.BackPressureMiddleware,
) *VotingRoute {
	if votingRouteInstance == nil {
		votingLock.Lock()
		defer votingLock.Unlock()
		if votingRouteInstance == nil {
			votingRouteInstance = &VotingRoute{
				ratingMiddleware:       ratingMiddleware,
				backPressureMiddleware: backPressureMiddleware,
			}
		}
	}
	return votingRouteInstance
}

func (v *VotingRoute) VotingRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(v.backPressureMiddleware.ServeBackPressure)
	r.Use(middleware.BasicValidationMiddleware)
	r.Use(v.ratingMiddleware.ServeHTTP)
	r.Post("/", postVoting)
	return r
}

func postVoting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := map[string]string{
		"status": "ok",
	}

	resJson, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(resJson)
	if err != nil {
		log.Fatal(err)
	}
}
