package main

import (
	"flag"
	"log"

	"github.com/bruli/raspberryRainSensor/internal/infrastructure/http/server"
	"github.com/spf13/viper"
)

func main() {
	configFile := flag.String("config", "", "config file")
	flag.Parse()

	viper.SetConfigFile(*configFile)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("invalid config file: %s", err)
	}

	serverAddr := viper.GetString("server_url")
	conf := server.NewConfig(serverAddr)

	s := server.NewServer(conf)
	s.Run()
}
