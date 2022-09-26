package logger

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

func InitLogger(serviceName string, isPretty bool, logLevel string) *zerolog.Logger {
	logger := getLog(serviceName, isPretty, logLevel)
	return &logger
}

func getLog(serviceName string, isPretty bool, logLevel string) zerolog.Logger {
	logger := zerolog.New(os.Stdout).With().Str("module", serviceName).Timestamp()

	if isPretty {
		logger = zerolog.New(os.Stdout).Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).With().Str("module", serviceName).Timestamp()
	}

	if logLevel != "" {
		level, err := zerolog.ParseLevel(logLevel)
		if err != nil {
			// default log level is debug
			return logger.Logger()
		}
		zerolog.SetGlobalLevel(level)
	}

	return logger.Logger()
}
