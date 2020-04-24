package rain_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/bruli/raspberryRainSensor/internal/log"

	"github.com/bruli/raspberryRainSensor/internal/rain"

	"github.com/stretchr/testify/assert"
)

func TestNewRainHandler(t *testing.T) {
	tests := map[string]struct {
		value  uint16
		isRain bool
		err    error
		logMsg string
	}{
		"it should write log when error": {err: errors.New("error"),
			logMsg: "Fatal error reading humidity: error"},
		"it should false without error": {value: 200, isRain: true},
		"it should true without error":  {value: 1000, isRain: false},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			repo := rain.RepositoryMock{}
			logger := log.LoggerMock{}
			repo.ReadFunc = func() (uint16, error) {
				return tt.value, tt.err
			}
			logger.FatalfFunc = func(format string, v ...interface{}) {
				assert.Equal(t, tt.logMsg, fmt.Sprintf(format, v...))
			}
			manager := rain.NewHandler(&repo, &logger)
			rain, err := manager.Handle()
			assert.Equal(t, tt.value, rain.Value)
			assert.Equal(t, tt.isRain, rain.IsRain)
			assert.Equal(t, tt.err, err)
		})
	}
}
