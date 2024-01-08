package backend

import (
	"llm-manager/internal/backend/langchaingo"
	"llm-manager/internal/backend/lingoose"
	"llm-manager/internal/backend/ollama"
	"llm-manager/internal/config"
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
