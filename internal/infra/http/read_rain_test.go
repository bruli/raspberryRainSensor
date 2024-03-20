package http_test

import (
	"context"
	"errors"
	http2 "net/http"
	"net/http/httptest"
	"testing"

	"github.com/bruli/raspberryRainSensor/fixtures"

	"github.com/bruli/raspberryRainSensor/pkg/common/cqs"

	"github.com/bruli/raspberryRainSensor/internal/infra/http"
	"github.com/stretchr/testify/require"
)

func TestReadRain(t *testing.T) {
	tests := []struct {
		name         string
		expectedCode int
		result       any
		qhErr        error
	}{
		{
			name:         "and query handler returns an error, then it returns an internal server error",
			qhErr:        errors.New(""),
			expectedCode: http2.StatusInternalServerError,
		},
		{
			name:         "and query handler returns a rain object, then it returns a valid response",
			result:       fixtures.BuildNotRaining(),
			expectedCode: http2.StatusOK,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(`Given a ReadRain http handler,
		when is called `+tt.name, func(t *testing.T) {
			t.Parallel()
			qh := &QueryHandlerMock{
				HandleFunc: func(ctx context.Context, query cqs.Query) (any, error) {
					return tt.result, tt.qhErr
				},
			}
			handler := http.ReadRain(qh)

			req := httptest.NewRequest(http2.MethodGet, "/", nil)
			writer := httptest.NewRecorder()
			handler.ServeHTTP(writer, req)

			resp := writer.Result()
			require.Equal(t, tt.expectedCode, resp.StatusCode)
		})
	}
}
