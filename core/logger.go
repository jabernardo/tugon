package core

import (
	"log"
	"log/slog"
	"os"
	"sync"
)

var (
	logger     *slog.Logger
	loggerOnce sync.Once
)

func GetLoggerInstance() *slog.Logger {
	loggerOnce.Do(func() {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	})

	if logger == nil {
		log.Fatalln("Pasok")
	}

	return logger
}
