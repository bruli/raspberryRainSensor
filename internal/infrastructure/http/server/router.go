package server

import (
	"net/http"

	"github.com/bruli/raspberryRainSensor/internal/infrastructure/log"

	"github.com/bruli/raspberryRainSensor/internal/infrastructure/spi/humiditySensor"

	"github.com/gorilla/mux"
)

type router struct {
	homepage http.Handler
	rain     http.Handler
}

func newRouter() *router {
	sensor := humiditySensor.NewSensor()
	logError := log.NewLogError()
	return &router{
		homepage: newHomepage(),
		rain:     newRainHandler(sensor, logError),
	}
}

func (r *router) build() *mux.Router {
	rout := mux.NewRouter()
	rout.HandleFunc("/", r.homepage.ServeHTTP).Methods(http.MethodGet)
	rout.HandleFunc("/rain", r.rain.ServeHTTP).Methods(http.MethodGet)

	return rout
}
