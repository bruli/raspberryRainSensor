package main

import (
	"os"

	"github.com/bruli/raspberryRainSensor/internal/infrastructure/http/server"
)

func main() {
	serverAddr := os.Getenv("SERVER_ADDR")
	conf := server.NewConfig(serverAddr)

	s := server.NewServer(conf)
	s.Run()
}
