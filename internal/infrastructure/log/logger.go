package log

import (
	"log"
	"os"

	log2 "github.com/bruli/raspberryRainSensor/internal/log"
)

func NewLogError() log2.Logger {
	return log.New(os.Stdout, "ERROR", 1)
}
