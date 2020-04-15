package server

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"

	"github.com/bruli/raspberryRainSensor/internal/rain"
	"github.com/bruli/raspberryRainSensor/pkg/log"
)

type rainValueResponseBody struct {
	Value uint16 `json:"value"`
}

func newRainValueResponseBody() *rainValueResponseBody {
	return &rainValueResponseBody{}
}

type rainValueHandler struct {
	body    *rainValueResponseBody
	manager *rain.RainManager
}

func newRainValueHandler(reader rain.HumidityReader, logger log.Logger) *rainValueHandler {
	return &rainValueHandler{body: newRainValueResponseBody(), manager: rain.NewRainManager(reader, logger)}
}

func (r *rainValueHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	val, err := r.manager.RainValue()
	if err != nil {
		writeJsonErrorResponse(writer, http.StatusInternalServerError, "Error reading rain sensor.")
		return
	}

	r.body.Value = val
	response, _ := jsoniter.Marshal(r.body)
	writeJsonResponse(writer, http.StatusOK, response)
}
