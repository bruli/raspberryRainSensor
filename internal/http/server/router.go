package server

import (
	"net/http"

	"github.com/bruli/raspberryRainSensor/internal/spi/humiditySensor"
	log2 "github.com/bruli/raspberryRainSensor/pkg/log"

	"github.com/gorilla/mux"
)

type router struct {
	homepage http.Handler
	rain     http.Handler
}

func newRouter() *router {
	return &router{
		homepage: newHomepage(),
		rain:     newRainHandler(humiditySensor.NewSensor(), log2.NewLogError()),
	}
}

func (r *router) build() *mux.Router {
	rout := mux.NewRouter()
	rout.HandleFunc("/", r.homepage.ServeHTTP).Methods(http.MethodGet)
	rout.HandleFunc("/rain", r.rain.ServeHTTP).Methods(http.MethodGet)

	return rout
}
