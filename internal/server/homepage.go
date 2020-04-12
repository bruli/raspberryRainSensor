package server

import "net/http"

type homepage struct {
}

func (h homepage) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writeJsonResponse(writer, http.StatusOK, nil)
}

func newHomepage() *homepage {
	return &homepage{}
}
