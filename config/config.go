package config

import "github.com/bruli/raspberryRainSensor/pkg/common/env"

const (
	ProjectPrefix = "RS_"
	ServerURL     = ProjectPrefix + "SERVER_URL"
	Environment   = ProjectPrefix + "ENVIRONMENT"
)

type Config struct {
	ServerURL  string
	Environent env.Environment
}

func NewConfig() (Config, error) {
	url, err := env.Value(ServerURL)
	if err != nil {
		return Config{}, err
	}
	envStr, err := env.Value(Environment)
	if err != nil {
		return Config{}, err
	}
	environment, err := env.ParseEnvironment(envStr)
	if err != nil {
		return Config{}, err
	}
	return Config{
		ServerURL:  url,
		Environent: environment,
	}, nil
}
