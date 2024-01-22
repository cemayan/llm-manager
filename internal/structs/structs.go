package structs

import log "github.com/sirupsen/logrus"

type Output int64

var (
	OutputMap = map[string]Output{
		"json": Json,
		"text": Text,
		"yaml": Yaml,
	}
)

const (
	Text Output = iota + 1
	Json
	Yaml
)

// Serve represents server config of app
type Serve struct {
	Port        int    `yaml:"port" json:"port"`
	PrivateKey  string `yaml:"privatekey" json:"privatekey"`
	Certificate string `yaml:"certificate" json:"certificate"`
	WebRoot     string `yaml:"webroot" json:"webroot"`
}

// Security represents security config of app
type Security struct {
	AllowedOrigins []string `yaml:"allowed_origins" json:"allowed_origins"`
}

// Log represents level config of app
type Log struct {
	Level log.Level `yaml:"level" json:"level"`
}

type Ollama struct {
	Server string `yaml:"server" json:"server"`
	Model  string `yaml:"model" json:"model"`
}

type Api struct {
	Version string `yaml:"version" json:"version"`
	Backend string `yaml:"backend" json:"backend"`
	Output  Output `yaml:"output" json:"output"`
}

// Config represents root config of app
type Config struct {
	Environment string   `yaml:"environment" json:"environment"`
	Api         Api      `yaml:"api" json:"api"`
	Ollama      Ollama   `yaml:"ollama" json:"ollama"`
	Security    Security `yaml:"security" json:"security"`
	Log         Log      `yaml:"log" json:"log"`
	Serve       Serve    `yaml:"serve" json:"serve"`
}

func (o Output) String() string {
	switch o {
	case Json:
		return "json"
	case Text:
		return "text"
	case Yaml:
		return "yaml"
	}
	return ""
}
