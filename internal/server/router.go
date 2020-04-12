package server

import (
	"net/http"
	"github.com/gorilla/mux"
)

type router struct {
	homepage http.Handler
}

func newRouter() *router {
	return &router{homepage: newHomepage()}
}

func (r *router) build() *mux.Router  {
	rout := mux.NewRouter()
	rout.HandleFunc("/", r.homepage.ServeHTTP).Methods(http.MethodGet)

	return rout
}