package server

import (
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func writeJsonResponse(w http.ResponseWriter, code int, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}

func writeJsonErrorResponse(w http.ResponseWriter, code int, error string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	message := errorResponse{Error: error}
	body, _ := jsoniter.Marshal(message)
	_, _ = w.Write(body)
}
