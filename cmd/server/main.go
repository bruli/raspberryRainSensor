package main

import (
	"context"
	http2 "net/http"
	"os"

	"github.com/rs/zerolog"

	"github.com/bruli/raspberryRainSensor/internal/infra/fake"

	"github.com/bruli/raspberryRainSensor/config"
	"github.com/bruli/raspberryRainSensor/internal/app"
	"github.com/bruli/raspberryRainSensor/internal/infra/http"
	"github.com/bruli/raspberryRainSensor/internal/infra/hunter"
	"github.com/bruli/raspberryRainSensor/pkg/common/cqs"
	"github.com/bruli/raspberryRainSensor/pkg/common/env"
	"github.com/bruli/raspberryRainSensor/pkg/common/httpx"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to build config")
	}
	ctx := context.Background()
	definitions, err := handlersDefinition(&log, conf.Environment())
	if err != nil {
		log.Fatal().Err(err)
	}
	httpHandlers := httpx.NewHandler(definitions)
	if err = httpx.RunServer(ctx, conf.ServerURL(), httpHandlers, &httpx.CORSOpt{}); err != nil {
		log.Fatal().Err(err).Msg("system error")
	}
}

func handlersDefinition(log *zerolog.Logger, env env.Environment) (httpx.HandlersDefinition, error) {
	var rs app.RainSensor
	rs = hunter.RainSensor{}
	if !env.IsProduction() {
		rs = fake.RainSensor{}
	}
	qhErrMdw := cqs.NewQueryHndErrorMiddleware(log)
	return httpx.HandlersDefinition{
		{
			Endpoint: "/",
			Method:   http2.MethodGet,
			HandlerFunc: func(writer http2.ResponseWriter, request *http2.Request) {
				_, _ = writer.Write([]byte(`{"status": "OK"}`))
			},
		}, {
			Endpoint:    "/rain",
			Method:      http2.MethodGet,
			HandlerFunc: http.ReadRain(qhErrMdw(app.NewReadRain(rs))),
		},
	}, nil
}
