package main

import (
	"github.com/bruli/raspberryRainSensor/internal/http/server"
	"os"
)

func main() {
	serverAddr := os.Getenv("SERVER_ADDR")
	conf := server.NewConfig(serverAddr)

	s := server.NewServer(conf)
	s.Run()
}
