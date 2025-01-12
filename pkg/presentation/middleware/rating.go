package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/danielchiovitti/ballot-box/pkg/domain/model"
	"github.com/danielchiovitti/ballot-box/pkg/presentation/factory/usecase/redis"
	"github.com/danielchiovitti/ballot-box/pkg/shared"
	"log"
	"net/http"
)

func NewRatingMiddleware(
	incrUseCaseFactory redis.IncrUseCaseFactoryInterface,
	expUseCaseFactory redis.ExpireUseCaseFactoryInterface,
	config shared.ConfigInterface,
) *RatingMiddleware {
	return &RatingMiddleware{
		incrUseCaseFactory: incrUseCaseFactory,
		config:             config,
		expUseCaseFactory:  expUseCaseFactory,
	}
}

type RatingMiddleware struct {
	incrUseCaseFactory redis.IncrUseCaseFactoryInterface
	expUseCaseFactory  redis.ExpireUseCaseFactoryInterface
	config             shared.ConfigInterface
}

func (rm *RatingMiddleware) ServeHTTP(next http.Handler) http.Handler {
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

		expUseCase := rm.expUseCaseFactory.Build()
		_ = expUseCase.Execute(r.Context(), key, rm.config.GetRateWindow())

		if res > rm.config.GetRateMaxReq() {
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

		next.ServeHTTP(w, r)
	})
}
