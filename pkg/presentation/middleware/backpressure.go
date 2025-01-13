package middleware

import (
	"encoding/json"
	"github.com/danielchiovitti/ballot-box/pkg/domain/model"
	"github.com/danielchiovitti/ballot-box/pkg/presentation/factory/usecase/redis"
	"github.com/danielchiovitti/ballot-box/pkg/shared"
	"log"
	"net/http"
	"sync"
)

var backPressureMiddlewareInstance *BackPressureMiddleware
var lockBackPressure sync.Mutex

func NewBackPressureMiddleware(
	incrUseCaseFactory redis.IncrUseCaseFactoryInterface,
	expUseCaseFactory redis.ExpireUseCaseFactoryInterface,
	config shared.ConfigInterface,
) *BackPressureMiddleware {
	if backPressureMiddlewareInstance == nil {
		lockBackPressure.Lock()
		defer lockBackPressure.Unlock()
		if backPressureMiddlewareInstance == nil {
			backPressureMiddlewareInstance = &BackPressureMiddleware{
				incrUseCaseFactory: incrUseCaseFactory,
				config:             config,
				expUseCaseFactory:  expUseCaseFactory,
			}
		}
	}
	return backPressureMiddlewareInstance
}

type BackPressureMiddleware struct {
	incrUseCaseFactory redis.IncrUseCaseFactoryInterface
	expUseCaseFactory  redis.ExpireUseCaseFactoryInterface
	config             shared.ConfigInterface
}

func (rm *BackPressureMiddleware) ServeBackPressure(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := "back-pressure"
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

		if res > rm.config.GetRateGlobalMaxReq() {
			res := &model.JsonErrorMessage{
				Message: shared.MAX_GLOBAL_REQ_LIMIT_EXCEEDED_MESSAGE,
				Code:    shared.MAX_GLOBAL_REQ_LIMIT_EXCEEDED_CODE,
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
			_ = expUseCase.Execute(r.Context(), key, rm.config.GetRateGlobalWindow())
		}

		next.ServeHTTP(w, r)
	})
}
