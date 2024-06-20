package backend

import (
	"llm-manager/internal/backend/langchaingo"
	"llm-manager/internal/backend/lingoose"
	"llm-manager/internal/backend/ollama"
	"llm-manager/internal/config"
)

// BackendInstance is used to reach selected backend whenever what you want
var BackendInstance Backend

type Backend interface {
	Query(body []byte, params map[string]interface{}) ([]byte, error)
}

// Init initializes according to given backend
// All backends should be implemented the Backend interface
func Init() {
	if config.AppConfig.Config.Api.Backend == "ollama" {
		BackendInstance = ollama.New()
	} else if config.AppConfig.Config.Api.Backend == "lingoose" {
		BackendInstance = lingoose.New()
	} else if config.AppConfig.Config.Api.Backend == "langchaingo" {
		BackendInstance = langchaingo.New()
	}
}
