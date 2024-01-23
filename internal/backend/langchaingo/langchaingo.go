package langchaingo

import (
	"context"
	"github.com/tmc/langchaingo/llms"
	openAi2 "github.com/tmc/langchaingo/llms/openai"
	"llm-manager/internal/backend/utils"
)

type Langchaingo struct {
	utils.CommonParams
	llm *openAi2.LLM
}

func (l Langchaingo) Query(body []byte, params map[string]interface{}) ([]byte, error) {
	ctx := context.Background()
	l.ConfigureParams(params)

	completion, err := l.llm.Call(ctx, string(body),
		llms.WithTemperature(l.Temp),
		llms.WithStopWords(l.StopWords),
	)
	return []byte(completion), err
}

func New() *Langchaingo {
	llm, _ := openAi2.New()
	return &Langchaingo{llm: llm}
}
