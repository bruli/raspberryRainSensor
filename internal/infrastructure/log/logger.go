package log

import (
	"log"
	"os"

	"github.com/bruli/raspberryRainSensor/internal/domain"
)

func NewLogError() domain.Logger {
	return log.New(os.Stdout, "ERROR", 1)
}
