package server

import (
	"net/http"
)

type rain struct {
}

func (r rain) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}

func newRain() *rain {
	return &rain{}
}