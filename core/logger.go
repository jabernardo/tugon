package core

import (
	"log/slog"
	"os"
	"sync"
)

var (
	logger     *slog.Logger
	loggerOnce sync.Once
)

func Logger() *slog.Logger {
	loggerOnce.Do(func() {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	})

	return logger
}
