package server

import (
	"net/http"

	"github.com/bruli/raspberryRainSensor/internal/log"

	"github.com/bruli/raspberryRainSensor/internal/rain"

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
	handler *rain.Handler
	body    *rainResponseBody
}

func newRainHandler(reader rain.Repository, logger log.Logger) *rainHandler {
	return &rainHandler{handler: rain.NewHandler(reader, logger), body: newRainResponseBody()}
}

func (r *rainHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	rain, err := r.handler.Handle()
	if err != nil {
		writeJsonErrorResponse(writer, http.StatusInternalServerError, "Error reading rain sensor.")
		return
	}

	r.body.IsRaining = rain.IsRain
	r.body.Value = rain.Value
	response, _ := jsoniter.Marshal(r.body)

	writeJsonResponse(writer, http.StatusOK, response)
}
