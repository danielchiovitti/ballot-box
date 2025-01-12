package main

import (
	"fmt"
	ballotbox "github.com/danielchiovitti/ballot-box"
	"github.com/danielchiovitti/ballot-box/pkg/shared"
	"net/http"
)

func main() {
	h := ballotbox.InitializeHandler()
	r := h.GetRoutes()
	v := h.GetViper()
	h.SetBloomFilter()

	port := v.GetInt(string(shared.PORT))
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		panic(err)
	}
}
