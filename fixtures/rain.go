package fixtures

import "github.com/bruli/raspberryRainSensor/internal/domain/rain"

func BuildRaining() rain.Rain {
	return rain.New(1023)
}

func BuildNotRaining() rain.Rain {
	return rain.New(0)
}
