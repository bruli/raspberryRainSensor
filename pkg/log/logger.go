package log

import (
	"log"
	"os"
)

//go:generate moq -out loggerMock.go . Logger
type Logger interface {
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
}

func NewLogInfo() Logger {
	return log.New(os.Stdout, "INFO", 0)
}

func NewLogError() Logger {
	return log.New(os.Stdout, "ERROR", 1)
}
