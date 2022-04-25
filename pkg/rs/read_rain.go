package rs

import (
	"context"
	"encoding/json"
	"io"
	http2 "net/http"

	"github.com/davecgh/go-spew/spew"

	"github.com/bruli/raspberryRainSensor/internal/infra/http"
)

func ReadRain(cl client) ReadRainFunc {
	return func(ctx context.Context) (Rain, error) {
		url := spew.Sprintf("%s/rain", cl.serverURL)
		req, _ := http2.NewRequest(http2.MethodGet, url, nil)
		resp, err := cl.cl.Do(req)
		if err != nil {
			return Rain{}, err
		}
		data, _ := io.ReadAll(resp.Body)
		rainResp := http.RainResponseJson{}
		if err = json.Unmarshal(data, &rainResp); err != nil {
			return Rain{}, err
		}
		return Rain{
			IsRaining: rainResp.IsRaining,
			Value:     rainResp.Value,
		}, nil
	}
}
