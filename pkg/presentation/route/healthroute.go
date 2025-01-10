package route

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"sync"
)

var healthLock sync.Mutex
var healthRouteInstance *HealthRoute

type HealthRoute struct{}

func NewHealthRoute() *HealthRoute {
	if healthRouteInstance == nil {
		healthLock.Lock()
		defer healthLock.Unlock()
		if healthRouteInstance == nil {
			healthRouteInstance = &HealthRoute{}
		}
	}

	return healthRouteInstance
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
