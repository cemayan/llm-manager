package config

import (
	"fmt"
	"git-observer/internal/structs"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

var AppConfig *Config

type Config struct {
	Config *structs.Config
}

// configureParameters returns Config according to given env
func configureParameters() *structs.Config {

	cfg := structs.Config{}

	if val, ok := os.LookupEnv("LLM_BACKEND"); ok {
		cfg.Api.Backend = val
	} else {
		cfg.Api.Backend = "ollama"
	}

	if val, ok := os.LookupEnv("BACKEND_ENV"); ok {
		cfg.Environment = val
	} else {
		cfg.Environment = "development"
	}

	if val, ok := os.LookupEnv("BACKEND_API_VERSION"); ok {
		cfg.Api.Version = val
	} else {
		cfg.Api.Version = "v1"
	}

	if val, ok := os.LookupEnv("OLLAMA_MODEL"); ok {
		cfg.Ollama.Model = val
	} else {
		cfg.Ollama.Model = "llama2"
	}

	if val, ok := os.LookupEnv("OLLAMA_SERVER"); ok {
		cfg.Ollama.Server = fmt.Sprintf("%v/api/generate", val)
	} else {
		cfg.Ollama.Server = fmt.Sprintf("%v/api/generate", "http://localhost:11434")
	}

	if val, ok := os.LookupEnv("BACKEND_CERT"); ok {
		cfg.Serve.Certificate = val
	}

	if val, ok := os.LookupEnv("BACKEND_ALLOWED_ORIGINS"); ok {
		cfg.Security.AllowedOrigins = strings.Split(val, ",")
	}

	if val, ok := os.LookupEnv("BACKEND_LOG_LEVEL"); ok {
		if lvl, err := strconv.Atoi(val); err == nil {
			cfg.Log.Level = log.Level(lvl)
		} else {
			cfg.Log.Level = log.InfoLevel
		}
	} else {
		cfg.Log.Level = log.InfoLevel
	}

	if val, ok := os.LookupEnv("BACKEND_PRIVATE_KEY"); ok {
		cfg.Serve.PrivateKey = val
	}

	if val, ok := os.LookupEnv("BACKEND_PORT"); ok {
		if p, err := strconv.Atoi(val); err == nil {
			cfg.Serve.Port = p
		}
	} else {
		cfg.Serve.Port = 8996 // Default port
	}

	return &cfg
}

func New() {
	AppConfig = &Config{Config: configureParameters()}
}