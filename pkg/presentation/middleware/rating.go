package middleware

import (
	"github.com/danielchiovitti/ballot-box/pkg/presentation/factory/usecase/redis"
	"net/http"
)

type RatingMiddleware struct {
	setStringUseCaseFactory redis.SetStringUseCaseFactoryInterface
}

func (r *RatingMiddleware) ServeHTTP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("user") == "0" {
			http.Error(w, "not valid", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
