package route

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type HealthRoute struct{}

func NewHealthRoute() *HealthRoute {
	return &HealthRoute{}
}

func (HealthRoute) HealthRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", getHealth)
	return r
}

func getHealth(w http.ResponseWriter, r *http.Request) {
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
