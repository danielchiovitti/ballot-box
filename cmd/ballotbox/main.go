package main

import (
	"fmt"
	ballotbox "github.com/danielchiovitti/ballot-box"
	"github.com/spf13/viper"
	"net/http"
)

func main() {
	viper.SetEnvPrefix("bb")
	viper.AutomaticEnv()

	h := ballotbox.InitializeHandler()
	r := h.GetRoutes()

	err := http.ListenAndServe(fmt.Sprintf(":%d", viper.GetInt("PORT")), r)
	if err != nil {
		panic(err)
	}
}
