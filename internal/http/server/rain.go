package server

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"

	"github.com/bruli/raspberryRainSensor/pkg/log"

	"github.com/bruli/raspberryRainSensor/internal/rain"
)

type rainResponseBody struct {
	IsRaining bool `json:"is_raining"`
}

func newRainResponseBody() *rainResponseBody {
	return &rainResponseBody{}
}

type rainHandler struct {
	manager *rain.RainManager
	body    *rainResponseBody
}

func newRainHandler(reader rain.HumidityReader, logger log.Logger) *rainHandler {
	return &rainHandler{manager: rain.NewRainManager(reader, logger), body: newRainResponseBody()}
}

func (r *rainHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	data, err := r.manager.IsRaining()
	if err != nil {
		writeJsonErrorResponse(writer, http.StatusInternalServerError, "Error reading rain sensor.")
		return
	}
	r.body.IsRaining = data
	response, _ := jsoniter.Marshal(r.body)

	writeJsonResponse(writer, http.StatusOK, response)
}
