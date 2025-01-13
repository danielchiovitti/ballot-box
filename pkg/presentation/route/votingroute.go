package route

import (
	"encoding/json"
	"github.com/danielchiovitti/ballot-box/pkg/domain/model"
	"github.com/danielchiovitti/ballot-box/pkg/presentation/factory/usecase/redis"
	"github.com/danielchiovitti/ballot-box/pkg/presentation/middleware"
	"github.com/danielchiovitti/ballot-box/pkg/shared"
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

var votingLock sync.Mutex
var votingRouteInstance *VotingRoute

type VotingRoute struct {
	ratingMiddleware          *middleware.RatingMiddleware
	backPressureMiddleware    *middleware.BackPressureMiddleware
	bloomFilterMiddleware     *middleware.BloomFilterMiddleware
	addToStreamUseCaseFactory redis.AddToStreamUseCaseFactoryInterface
	config                    shared.ConfigInterface
}

func NewVotingRoute(
	ratingMiddleware *middleware.RatingMiddleware,
	backPressureMiddleware *middleware.BackPressureMiddleware,
	bloomFilterMiddleware *middleware.BloomFilterMiddleware,
	addToStreamUseCaseFactory redis.AddToStreamUseCaseFactoryInterface,
	config shared.ConfigInterface,
) *VotingRoute {
	if votingRouteInstance == nil {
		votingLock.Lock()
		defer votingLock.Unlock()
		if votingRouteInstance == nil {
			votingRouteInstance = &VotingRoute{
				ratingMiddleware:          ratingMiddleware,
				backPressureMiddleware:    backPressureMiddleware,
				bloomFilterMiddleware:     bloomFilterMiddleware,
				addToStreamUseCaseFactory: addToStreamUseCaseFactory,
				config:                    config,
			}
		}
	}
	return votingRouteInstance
}

func (v *VotingRoute) VotingRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(v.backPressureMiddleware.ServeBackPressure)
	r.Use(v.bloomFilterMiddleware.ServeBloomFilter)
	r.Use(middleware.BasicValidationMiddleware)
	r.Use(v.ratingMiddleware.ServeRating)
	r.Post("/", v.postVoting)
	return r
}

func (v *VotingRoute) postVoting(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var vote model.Vote
	err = json.Unmarshal(body, &vote)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vote.CreatedAt = time.Now()
	addToStreamUseCase := v.addToStreamUseCaseFactory.Build()

	jsonData, err := json.Marshal(vote)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = addToStreamUseCase.Execute(r.Context(), v.config.GetOltpStreamName(), string(jsonData))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = addToStreamUseCase.Execute(r.Context(), v.config.GetOlapStreamName(), string(jsonData))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := map[string]string{
		"status": "ok",
	}

	resJson, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}

	_, err = w.Write(resJson)
	if err != nil {
		log.Println(err)
	}
}
