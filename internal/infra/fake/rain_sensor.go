package fake

import (
	"context"

	"github.com/bruli/raspberryRainSensor/fixtures"
	"github.com/bruli/raspberryRainSensor/internal/domain/rain"
)

type RainSensor struct{}

func (r RainSensor) Read(ctx context.Context) (rain.Rain, error) {
	return fixtures.BuildRaining(), nil
}
