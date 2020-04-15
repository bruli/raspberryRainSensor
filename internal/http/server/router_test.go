package server

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	jsoniter "github.com/json-iterator/go"

	"github.com/bruli/raspberryRainSensor/internal/rain"
	"github.com/bruli/raspberryRainSensor/pkg/log"

	"github.com/stretchr/testify/assert"
)

func TestHomepage(t *testing.T) {
	rout := newRouter().build()

	t.Run("it should return homepageHandler", func(t *testing.T) {
		body := newHomepageResponseBody()
		expectedBody, _ := jsoniter.Marshal(body)

		request, err := http.NewRequest(http.MethodGet, "/", nil)
		assert.NoError(t, err)

		writer := httptest.NewRecorder()
		rout.ServeHTTP(writer, request)
		assert.Equal(t, http.StatusOK, writer.Code)
		assert.Equal(t, bytes.NewBuffer(expectedBody), writer.Body)
	})
}

func TestRainHandler(t *testing.T) {
	errorResponse := errorResponse{Error: "Error reading rain sensor."}
	errorMsg, _ := jsoniter.Marshal(&errorResponse)
	response := rainResponseBody{IsRaining: true}
	responseBody, _ := jsoniter.Marshal(&response)

	tests := map[string]struct {
		body         io.Reader
		httpCodeResp int
		readerError  error
		readerValue  uint16
	}{
		"it should return internal server error when reader returns error": {body: bytes.NewBuffer(errorMsg), httpCodeResp: http.StatusInternalServerError, readerError: errors.New("error"), readerValue: 0},
		"it should return Ok": {body: bytes.NewBuffer(responseBody), httpCodeResp: http.StatusOK, readerValue: 200},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			reader := rain.HumidityReaderMock{}
			logger := log.LoggerMock{}
			rout := newRouter()
			rout.rain = newRainHandler(&reader, &logger)
			server := rout.build()

			request, err := http.NewRequest(http.MethodGet, "/rain", tt.body)
			assert.NoError(t, err)

			writer := httptest.NewRecorder()

			logger.FatalfFunc = func(format string, v ...interface{}) {
			}
			reader.ReadFunc = func() (uint16, error) {
				return tt.readerValue, tt.readerError
			}

			server.ServeHTTP(writer, request)
			assert.Equal(t, tt.httpCodeResp, writer.Code)
			assert.Equal(t, tt.body, writer.Body)
		})
	}
}
