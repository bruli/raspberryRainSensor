package server

import (
	"net/http"

	"github.com/bruli/raspberryRainSensor/internal/application"

	"github.com/bruli/raspberryRainSensor/internal/domain"
	jsoniter "github.com/json-iterator/go"
)

type rainResponseBody struct {
	IsRaining bool   `json:"is_raining"`
	Value     uint16 `json:"value"`
}

func newRainResponseBody() *rainResponseBody {
	return &rainResponseBody{}
}

type rainHandler struct {
	handler *application.RainHandler
	body    *rainResponseBody
}

func newRainHandler(reader domain.RainRepository, logger domain.Logger) *rainHandler {
	return &rainHandler{handler: application.NewRainHandler(reader, logger), body: newRainResponseBody()}
}

func (r *rainHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	raining, err := r.handler.IsRaining()
	if err != nil {
		writeJsonErrorResponse(writer, http.StatusInternalServerError, "Error reading rain sensor.")
		return
	}
	value, err := r.handler.RainValue()
	if err != nil {
		writeJsonErrorResponse(writer, http.StatusInternalServerError, "Error reading rain sensor.")
		return
	}

	r.body.IsRaining = raining
	r.body.Value = value
	response, _ := jsoniter.Marshal(r.body)

	writeJsonResponse(writer, http.StatusOK, response)
}
