package langchaingo

import (
	"context"
	"github.com/tmc/langchaingo/llms"
	openAi2 "github.com/tmc/langchaingo/llms/openai"
)

type Langchaingo struct {
	llm       *openAi2.LLM
	stopWords []string
	temp      float64
}

func (l Langchaingo) Query(body []byte, params map[string]interface{}) ([]byte, error) {
	ctx := context.Background()
	completion, err := l.llm.Call(ctx, string(body),
		llms.WithTemperature(l.temp),
		llms.WithStopWords(l.stopWords),
	)
	return []byte(completion), err
}

func New() *Langchaingo {
	llm, _ := openAi2.New()
	return &Langchaingo{llm: llm, temp: 0.7}
}
