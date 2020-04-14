package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

type router struct {
	homepage http.Handler
	rain http.Handler
}

func newRouter() *router {
	return &router{
		homepage: newHomepage(),
		rain:     newRain(),
	}
}

func (r *router) build() *mux.Router  {
	rout := mux.NewRouter()
	rout.HandleFunc("/", r.homepage.ServeHTTP).Methods(http.MethodGet)
	rout.HandleFunc("/rain", r.rain.ServeHTTP).Methods(http.MethodGet)

	return rout
}