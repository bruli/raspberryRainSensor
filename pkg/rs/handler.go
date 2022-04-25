package rs

import "context"

type ReadRainFunc func(ctx context.Context) (Rain, error)

type Handler struct {
	ReadRain func(ctx context.Context) (Rain, error)
}
