package middleware

import "net/http"

func RatingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("user") == "0" {
			http.Error(w, "not valid", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
