package app_test

import (
	"context"
	"errors"
	"testing"

	"github.com/bruli/raspberryRainSensor/fixtures"

	"github.com/bruli/raspberryRainSensor/internal/domain/rain"

	"github.com/bruli/raspberryRainSensor/pkg/common/test"
	"github.com/stretchr/testify/require"

	"github.com/bruli/raspberryRainSensor/internal/app"
	"github.com/bruli/raspberryRainSensor/pkg/common/cqs"
)

func TestReadRainHandle(t *testing.T) {
	err := errors.New("")
	ra := fixtures.BuildRaining()
	tests := []struct {
		name                   string
		query                  cqs.Query
		expectedErr, sensorErr error
		result                 rain.Rain
	}{
		{
			name:        "with an invalid query, then returns an invalid query error",
			query:       &invalidQuery{},
			expectedErr: cqs.InvalidQueryError{},
		},
		{
			name:        "with a valid query and sensor returns an error, then returns same error",
			query:       app.ReadRainQuery{},
			expectedErr: err,
			sensorErr:   err,
		},
		{
			name:   "with a valid query and sensor returns a rain, then returns a valid result",
			query:  app.ReadRainQuery{},
			result: ra,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(`Given a ReadRain query handler,
		when Handle method is called `+tt.name, func(t *testing.T) {
			t.Parallel()
			rs := &RainSensorMock{
				ReadFunc: func(ctx context.Context) (rain.Rain, error) {
					return tt.result, tt.sensorErr
				},
			}
			handler := app.NewReadRain(rs)
			result, err := handler.Handle(context.Background(), tt.query)
			if err != nil {
				test.CheckErrorsType(t, tt.expectedErr, err)
				return
			}
			require.Equal(t, tt.expectedErr, err)
			ra, ok := result.(rain.Rain)
			require.True(t, ok)
			require.Equal(t, tt.result, ra)
		})
	}
}

type invalidQuery struct{}

func (i invalidQuery) Name() string {
	return "invalid"
}
