package lingoose

import (
	"context"
	"github.com/henomis/lingoose/llm/openai"
)

type Lingoose struct {
	openAI *openai.OpenAI
}

func (l Lingoose) Query(body []byte, params map[string]interface{}) ([]byte, error) {
	response, err := l.openAI.Completion(context.Background(), string(body))
	return []byte(response), err
}

func New() *Lingoose {
	openAI := openai.NewCompletion().WithVerbose(true)
	return &Lingoose{openAI: openAI}
}
