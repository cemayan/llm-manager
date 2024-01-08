package backend

import (
	"git-observer/internal/backend/langchaingo"
	"git-observer/internal/backend/lingoose"
	"git-observer/internal/backend/ollama"
	"git-observer/internal/config"
)

var BackendInstance Backend

type Backend interface {
	Query(body []byte, params map[string]interface{}) ([]byte, error)
}

func Init() {
	if config.AppConfig.Config.Api.Backend == "ollama" {
		BackendInstance = ollama.New()
	} else if config.AppConfig.Config.Api.Backend == "lingoose" {
		BackendInstance = lingoose.New()
	} else if config.AppConfig.Config.Api.Backend == "langchaingo" {
		BackendInstance = langchaingo.New()
	}
}
