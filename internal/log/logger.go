package log

import (
	log "github.com/sirupsen/logrus"
	"llm-manager/internal/config"
)

var LoggerInstance *Log

type Log struct {
	Logger *log.Logger
}

// New sets a new log instance
func New() {
	_logger := log.New()

	if config.AppConfig.Config.Environment == "production" {
		_logger.SetFormatter(&log.JSONFormatter{})
	} else {
		_logger.SetFormatter(&log.TextFormatter{
			FullTimestamp: true,
		})
		_logger.SetFormatter(&log.TextFormatter{})
	}

	_logger.SetLevel(config.AppConfig.Config.Log.Level)

	LoggerInstance = &Log{Logger: _logger}
}
