package ollama

import (
	"bytes"
	"io"
	"llm-manager/internal/config"
	"net/http"
)

const MimeTypeJson = "application/json"

type Ollama struct {
	client http.Client
}

func (o *Ollama) Query(body []byte, params map[string]interface{}) ([]byte, error) {
	resp, err := o.client.Post(config.AppConfig.Config.Ollama.Server, MimeTypeJson, bytes.NewBuffer(body))

	if err != nil {
		return body, err
	}

	body, err = io.ReadAll(resp.Body)
	return body, err
}

func New() *Ollama {
	return &Ollama{client: http.Client{}}
}
