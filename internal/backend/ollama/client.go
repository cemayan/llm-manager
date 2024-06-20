package ollama

import (
	"bytes"
	"io"
	"llm-manager/internal/config"
	"net/http"
)

const MimeTypeJson = "application/json"

type Ollama struct {
	client http.Client // in order to request ollama server it used to http.Client
}

// Query returns response according to the given prompt
func (o *Ollama) Query(body []byte, params map[string]interface{}) ([]byte, error) {
	resp, err := o.client.Post(config.AppConfig.Config.Ollama.Server, MimeTypeJson, bytes.NewBuffer(body))

	if err != nil {
		return body, err
	}

	body, err = io.ReadAll(resp.Body)
	return body, err
}

// New returns Ollama client
// Local or remote ollama server
func New() *Ollama {
	return &Ollama{client: http.Client{}}
}
