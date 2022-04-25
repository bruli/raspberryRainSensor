package main

import (
	"context"
	"fmt"
	"log"
	http2 "net/http"
	"os"
	"time"

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
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}
	ctx := context.Background()
	logger := log.New(os.Stdout, config.ProjectPrefix, int(time.Now().Unix()))
	definitions, err := handlersDefinition(logger, conf.Environent)
	if err != nil {
		log.Fatalln(err)
	}
	httpHandlers := httpx.NewHandler(definitions)
	if err := httpx.RunServer(ctx, conf.ServerURL, httpHandlers, &httpx.CORSOpt{}); err != nil {
		log.Fatalln(fmt.Errorf("system error: %w", err))
	}
}

func handlersDefinition(log *log.Logger, env env.Environment) (httpx.HandlersDefinition, error) {
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
