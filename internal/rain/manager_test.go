package rain_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/bruli/raspberryRainSensor/internal/rain"
	"github.com/stretchr/testify/assert"

	"github.com/bruli/raspberryRainSensor/pkg/log"
)

func TestManager_IsRaining(t *testing.T) {
	tests := map[string]struct {
		result    bool
		readValue uint16
		err       error
		logMsg    string
	}{
		"it should write log when error": {result: false, readValue: 0, err: errors.New("error"), logMsg: "Fatal error reading humidity: error"},
		"it should false without error":  {result: false, readValue: 1000},
		"it should true without error":   {result: true, readValue: 250},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			read := rain.HumidityReaderMock{}
			logger := log.LoggerMock{}
			read.ReadFunc = func() (uint16, error) {
				return tt.readValue, tt.err
			}
			logger.FatalfFunc = func(format string, v ...interface{}) {
				assert.Equal(t, tt.logMsg, fmt.Sprintf(format, v...))
			}
			manager := rain.NewRainManager(&read, &logger)
			result, err := manager.IsRaining()
			assert.Equal(t, tt.result, result)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestRainManager_RainValue(t *testing.T) {
	tests := map[string]struct {
		result uint16
		err    error
	}{
		"it should return error":  {result: 0, err: errors.New("error")},
		"it should return values": {result: 500},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			read := rain.HumidityReaderMock{}
			logger := log.LoggerMock{}
			read.ReadFunc = func() (uint16, error) {
				return tt.result, tt.err
			}
			logger.FatalfFunc = func(format string, v ...interface{}) {
			}
			manager := rain.NewRainManager(&read, &logger)
			result, err := manager.RainValue()
			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.result, result)
		})
	}
}
