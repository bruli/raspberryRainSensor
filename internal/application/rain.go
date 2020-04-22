package application

import (
	"github.com/bruli/raspberryRainSensor/internal/domain"
)

type RainHandler struct {
	repository domain.RainRepository
	Logger     domain.Logger
}

const rainRef = 500

func NewRainHandler(reader domain.RainRepository, logger domain.Logger) *RainHandler {
	return &RainHandler{repository: reader, Logger: logger}
}
func (m *RainHandler) IsRaining() (bool, error) {
	v, err := m.readValues()
	if err != nil {
		return false, err
	}
	isRain := v < rainRef
	return isRain, err
}

func (m *RainHandler) RainValue() (uint16, error) {
	v, err := m.readValues()
	if err != nil {
		return 0, err
	}

	return v, nil
}

func (m *RainHandler) readValues() (uint16, error) {
	v, err := m.repository.Read()
	if err != nil {
		m.Logger.Fatalf("Fatal error reading humidity: %s", err)
		return 0, err
	}
	return v, nil
}
