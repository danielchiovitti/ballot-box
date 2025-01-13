package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/danielchiovitti/ballot-box/pkg/domain/model"
	"github.com/danielchiovitti/ballot-box/pkg/presentation/factory/usecase/redis"
	"github.com/danielchiovitti/ballot-box/pkg/presentation/factory/usecase/redisbloom"
	"github.com/danielchiovitti/ballot-box/pkg/shared"
	"log"
	"net/http"
	"sync"
)

var ratingMiddlewareInstance *RatingMiddleware
var lockRating sync.Mutex

func NewRatingMiddleware(
	incrUseCaseFactory redis.IncrUseCaseFactoryInterface,
	expUseCaseFactory redis.ExpireUseCaseFactoryInterface,
	addUseCaseFactory redisbloom.AddUseCaseFactoryInterface,
	config shared.ConfigInterface,
) *RatingMiddleware {
	if ratingMiddlewareInstance == nil {
		lockRating.Lock()
		defer lockRating.Unlock()
		if ratingMiddlewareInstance == nil {
			ratingMiddlewareInstance = &RatingMiddleware{
				incrUseCaseFactory: incrUseCaseFactory,
				expUseCaseFactory:  expUseCaseFactory,
				addUseCaseFactory:  addUseCaseFactory,
				config:             config,
			}
		}
	}
	return ratingMiddlewareInstance
}

type RatingMiddleware struct {
	incrUseCaseFactory redis.IncrUseCaseFactoryInterface
	expUseCaseFactory  redis.ExpireUseCaseFactoryInterface
	addUseCaseFactory  redisbloom.AddUseCaseFactoryInterface
	config             shared.ConfigInterface
}

func (rm *RatingMiddleware) ServeRating(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := fmt.Sprintf("user:%s", r.Header.Get("user"))
		incrUseCase := rm.incrUseCaseFactory.Build()
		res, err := incrUseCase.Execute(r.Context(), key)
		if err != nil {
			res := &model.JsonErrorMessage{
				Message: shared.INTERNAL_ERROR_MESSAGE,
				Code:    shared.INTERNAL_ERROR_CODE,
			}
			resJson, err := json.Marshal(res)
			if err != nil {
				log.Println(err)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write(resJson)
			if err != nil {
				log.Println(err)
			}

			return
		}

		if res > rm.config.GetRateMaxReq() {
			addUseCase := rm.addUseCaseFactory.Build()
			_ = addUseCase.Execute(rm.config.GetBloomName(), key)

			res := &model.JsonErrorMessage{
				Message: shared.MAX_REQ_LIMIT_EXCEEDED,
				Code:    shared.MAX_REQ_LIMIT_EXCEEDED_CODE,
			}
			resJson, err := json.Marshal(res)
			if err != nil {
				log.Println(err)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write(resJson)
			if err != nil {
				log.Println(err)
			}

			return
		}

		if res == 1 {
			expUseCase := rm.expUseCaseFactory.Build()
			_ = expUseCase.Execute(r.Context(), key, rm.config.GetRateWindow())
		}

		next.ServeHTTP(w, r)
	})
}
