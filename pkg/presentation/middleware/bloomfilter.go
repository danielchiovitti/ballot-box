package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/danielchiovitti/ballot-box/pkg/domain/model"
	"github.com/danielchiovitti/ballot-box/pkg/presentation/factory/usecase/redisbloom"
	"github.com/danielchiovitti/ballot-box/pkg/shared"
	"log"
	"net/http"
	"sync"
)

var bloomFilterMiddlewareInstance *BloomFilterMiddleware
var lockBloomFilter sync.Mutex

func NewBloomFilterMiddleware(
	existsUseCaseFactory redisbloom.ExistsUseCaseFactoryInterface,
	config shared.ConfigInterface,
) *BloomFilterMiddleware {
	if bloomFilterMiddlewareInstance == nil {
		lockBloomFilter.Lock()
		defer lockBloomFilter.Unlock()
		if bloomFilterMiddlewareInstance == nil {
			bloomFilterMiddlewareInstance = &BloomFilterMiddleware{
				existsUseCaseFactory: existsUseCaseFactory,
				config:               config,
			}
		}
	}
	return bloomFilterMiddlewareInstance
}

type BloomFilterMiddleware struct {
	existsUseCaseFactory redisbloom.ExistsUseCaseFactoryInterface
	config               shared.ConfigInterface
}

func (rm *BloomFilterMiddleware) ServeBloomFilter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		key := fmt.Sprintf("user:%s", r.Header.Get("user"))
		existsUseCase := rm.existsUseCaseFactory.Build()
		res, err := existsUseCase.Execute(rm.config.GetBloomName(), key)
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
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write(resJson)
			if err != nil {
				log.Println(err)
			}

			return
		}

		if res {
			res := &model.JsonErrorMessage{
				Message: shared.BANNED_USER_MESSAGE,
				Code:    shared.BANNED_USER_CODE,
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
