package rain_test

import (
	"testing"

	"github.com/bruli/raspberryRainSensor/internal/domain/rain"
	"github.com/stretchr/testify/require"
)

func TestNewRain(t *testing.T) {
	t.Run(`Given a New function,
	when is called,
	then returns a full rain struct`, func(t *testing.T) {
		value := 1023
		raining := true
		obj := rain.New(value, raining)
		require.Equal(t, value, obj.Value())
		require.Equal(t, raining, obj.Raining())
	})
}
