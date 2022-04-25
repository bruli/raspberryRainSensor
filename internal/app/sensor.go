package app

import (
	"context"

	"github.com/bruli/raspberryRainSensor/internal/domain/rain"
)

//go:generate moq -out zmock_sensor_test.go -pkg app_test . RainSensor

type RainSensor interface {
	Read(ctx context.Context) (rain.Rain, error)
}
