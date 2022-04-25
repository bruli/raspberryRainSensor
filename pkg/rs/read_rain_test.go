package rs_test

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/bruli/raspberryRainSensor/fixtures"
	"github.com/bruli/raspberryRainSensor/internal/domain/rain"
	http2 "github.com/bruli/raspberryRainSensor/internal/infra/http"

	"github.com/bruli/raspberryRainSensor/pkg/common/test"
	"github.com/bruli/raspberryRainSensor/pkg/rs"
	"github.com/stretchr/testify/require"
)

func TestReadRain(t *testing.T) {
	err := errors.New("")
	ra := fixtures.BuildRaining()
	tests := []struct {
		name               string
		expectedErr, clErr error
		expectedRain       rs.Rain
		response           *http.Response
	}{
		{
			name:        "and client returns error, then it returns same error",
			clErr:       err,
			expectedErr: err,
		},
		{
			name:        "and read body returns error, then it returns same error",
			expectedErr: &json.SyntaxError{},
			response:    &http.Response{Body: http.NoBody},
		},
		{
			name:     "then it returns a valid rain object",
			response: &http.Response{Body: buildBody(t, ra)},
			expectedRain: rs.Rain{
				IsRaining: ra.Raining(),
				Value:     ra.Value(),
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(`Given a ReadRain function,
		when is called `+tt.name, func(t *testing.T) {
			t.Parallel()
			cl := &HTTPClientMock{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					return tt.response, tt.clErr
				},
			}
			pkg := rs.New("serverURL", cl)
			ra, err := pkg.ReadRain(context.Background())
			if err != nil {
				test.CheckErrorsType(t, tt.expectedErr, err)
				return
			}
			require.Equal(t, tt.expectedErr, err)
			require.Equal(t, tt.expectedRain, ra)
		})
	}
}

func buildBody(t *testing.T, ra rain.Rain) io.ReadCloser {
	resp := http2.RainResponseJson{
		IsRaining: ra.Raining(),
		Value:     ra.Value(),
	}
	data, err := json.Marshal(resp)
	require.NoError(t, err)
	return io.NopCloser(strings.NewReader(string(data)))
}
