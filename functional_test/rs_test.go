//go:build functional
// +build functional

package functional_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/bruli/raspberryRainSensor/pkg/rs"
)

func TestRainSensor(t *testing.T) {
	rsPkg := rs.New("http://localhost:8082", &http.Client{Timeout: 3 * time.Second})
	ra, err := rsPkg.ReadRain(context.Background())
	require.NoError(t, err)
	require.NotEqual(t, 0, ra.Value)
	require.True(t, ra.IsRaining)
}
