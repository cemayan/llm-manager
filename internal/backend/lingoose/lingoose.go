package lingoose

import (
	"context"
	"github.com/henomis/lingoose/chat"
	"github.com/henomis/lingoose/llm/openai"
	"github.com/henomis/lingoose/prompt"
	"llm-manager/internal/backend/utils"
)

type Lingoose struct {
	utils.CommonParams
	openAI *openai.OpenAI
}

// Query returns response according to the given prompt
// StopWords can be set
func (l Lingoose) Query(body []byte, params map[string]interface{}) ([]byte, error) {
	l.ConfigureParams(params)
	l.openAI.SetStop(l.StopWords)

	developerChat := chat.New(
		chat.PromptMessage{
			Type:   chat.MessageTypeSystem,
			Prompt: prompt.New(string(body)),
		},
	)

	response, err := l.openAI.Chat(context.Background(), developerChat)
	return []byte(response), err
}

// New returns Lingoose client
// Default values are like below:
// TODO: model, temprature and maxTokens can be changed by os environment
func New() *Lingoose {
	llmOpenAI := openai.New(
		openai.GPT3Dot5Turbo,
		openai.DefaultOpenAITemperature,
		openai.DefaultOpenAIMaxTokens,
		true,
	)
	return &Lingoose{openAI: llmOpenAI}
}
