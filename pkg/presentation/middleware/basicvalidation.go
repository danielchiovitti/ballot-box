package middleware

import (
	"encoding/json"
	"github.com/danielchiovitti/ballot-box/pkg/domain/model"
	"github.com/danielchiovitti/ballot-box/pkg/shared"
	"github.com/danielchiovitti/ballot-box/pkg/shared/helper"
	"log"
	"net/http"
)

func BasicValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		headers := []string{
			"user",
			"cookie",
			"referer",
			"user-agent",
			"sec-ch-ua",
			"sec-ch-ua-mobile",
			"sec-ch-ua-platform",
		}

		if !helper.CheckMandatoryHeaders(headers, r.Header) {
			res := &model.JsonErrorMessage{
				Message: shared.MANDATORY_HEADERS_MISSING_MESSAGE,
				Code:    shared.MANDATORY_HEADERS_MISSING_CODE,
			}
			resJson, err := json.Marshal(res)
			if err != nil {
				log.Fatal(err)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write(resJson)
			if err != nil {
				log.Fatal(err)
			}

			return
		}

		next.ServeHTTP(w, r)
	})
}
