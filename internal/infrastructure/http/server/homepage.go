package server

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

type homepageResponseBody struct {
	Status string `json:"status"`
}

func newHomepageResponseBody() *homepageResponseBody {
	return &homepageResponseBody{Status: "OK"}
}

type homepageHandler struct {
	body *homepageResponseBody
}

func (h *homepageHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	b, _ := jsoniter.Marshal(h.body)
	writeJsonResponse(writer, http.StatusOK, b)
}

func newHomepage() *homepageHandler {
	return &homepageHandler{body: newHomepageResponseBody()}
}
