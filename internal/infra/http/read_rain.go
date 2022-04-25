package http

import (
	"encoding/json"
	"net/http"

	"github.com/bruli/raspberryRainSensor/internal/app"
	"github.com/bruli/raspberryRainSensor/internal/domain/rain"
	"github.com/bruli/raspberryRainSensor/pkg/common/cqs"
)

func ReadRain(qh cqs.QueryHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := qh.Handle(r.Context(), app.ReadRainQuery{})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ra, _ := result.(rain.Rain)
		resp := RainResponseJson{
			IsRaining: ra.Raining(),
			Value:     ra.Value(),
		}
		data, _ := json.Marshal(resp)
		writeResponse(w, http.StatusOK, data)
	}
}

func writeResponse(w http.ResponseWriter, statusCode int, body []byte) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write(body)
}
