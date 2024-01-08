package types

import "encoding/json"

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Raw    bool   `json:"raw"`
	Format string `json:"format"`
	Stream bool   `json:"stream"`
}

type OllamaResponse struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Response  string `json:"response"`
}

func (or *OllamaRequest) Marshall() []byte {
	marshal, err := json.Marshal(or)

	if err != nil {
		return []byte("")
	}

	return marshal
}
