package rain

import (
	"github.com/bruli/raspberryRainSensor/internal/log"
)

type Rain struct {
	IsRain bool
	Value  uint16
}

type Handler struct {
	repository Repository
	Logger     log.Logger
}

const rainRef = 500

func NewHandler(reader Repository, logger log.Logger) *Handler {
	return &Handler{repository: reader, Logger: logger}
}

func (h *Handler) Handle() (Rain, error) {
	v, err := h.repository.Read()
	if err != nil {
		h.Logger.Fatalf("Fatal error reading humidity: %s", err)
		return Rain{}, err
	}

	return Rain{Value: v, IsRain: v < rainRef}, nil
}
