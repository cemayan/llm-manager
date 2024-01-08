package log

import (
	"git-observer/internal/config"
	log "github.com/sirupsen/logrus"
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
