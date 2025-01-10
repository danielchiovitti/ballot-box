package route

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type VotingRoute struct{}

func NewVotingRoute() *VotingRoute {
	return &VotingRoute{}
}

func (VotingRoute) VotingRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", postVoting)
	return r
}

func postVoting(w http.ResponseWriter, r *http.Request) {
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
