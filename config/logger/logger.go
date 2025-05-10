package logger

import (
	"log"

	commonLogger "github.com/LukmanulHakim18/core/logger"
	"github.com/LukmanulHakim18/time2go/config"
)

var logger *commonLogger.Logger

func LoadLogger() {
	appLogger, err := commonLogger.NewLogger(commonLogger.LoggerConfig{
		Level:           config.GetConfig("log_level").GetString(),
		LogDirectory:    config.GetConfig("log_directory").GetString(),
		AppName:         config.GetConfig("app_name").GetString(),
		SamplingEnabled: false,
	})
	if err != nil {
		log.Fatalf("cannot initiate logger, with error: %v", err)
	}
	logger = appLogger
}

func GetLogger() *commonLogger.Logger {
	return logger
}
