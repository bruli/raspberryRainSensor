package app

import (
	"context"

	"github.com/bruli/raspberryRainSensor/pkg/common/cqs"
)

const ReadRainQueryName = "readRain"

type ReadRainQuery struct{}

func (r ReadRainQuery) Name() string {
	return ReadRainQueryName
}

type ReadRain struct {
	rs RainSensor
}

func (r ReadRain) Handle(ctx context.Context, query cqs.Query) (cqs.QueryResult, error) {
	_, ok := query.(ReadRainQuery)
	if !ok {
		return nil, cqs.NewInvalidQueryError(ReadRainQueryName, query.Name())
	}
	ra, err := r.rs.Read(ctx)
	if err != nil {
		return nil, err
	}
	return ra, nil
}

func NewReadRain(rs RainSensor) ReadRain {
	return ReadRain{rs: rs}
}
