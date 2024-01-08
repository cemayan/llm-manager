package cli

import (
	"encoding/json"
	"fmt"
	"llm-manager/api/types"
	"llm-manager/api/util"
	"llm-manager/internal/backend"
	"llm-manager/internal/config"
)

func Exec(prompt string) error {

	switch config.AppConfig.Config.Api.Backend {
	case "ollama":
		query, err := backend.BackendInstance.
			Query(util.PrepareOllamaRequest(config.AppConfig.Config.Ollama.Model, prompt).Marshall(), nil)
		if err != nil {
			fmt.Println(string(util.PrepareErrorResponse(err, 1003).Marshall()))
			return err
		}

		var ollamaResp types.OllamaResponse
		err = json.Unmarshal(query, &ollamaResp)

		if err != nil {
			fmt.Println(string(util.PrepareErrorResponse(err, 1001).Marshall()))
			return err
		}

		fmt.Println(string(util.PrepareResponse(ollamaResp.Response).Marshall()))
	case "lingoose":
		query, err := backend.BackendInstance.
			Query([]byte(prompt), nil)
		if err != nil {
			fmt.Println(string(util.PrepareErrorResponse(err, 1003).Marshall()))
			return err
		}

		fmt.Println(string(util.PrepareResponse(query).Marshall()))
	case "langchaingo":
		query, err := backend.BackendInstance.
			Query([]byte(prompt), nil)
		if err != nil {
			fmt.Println(string(util.PrepareErrorResponse(err, 1003).Marshall()))
			return err
		}

		fmt.Println(string(util.PrepareResponse(query).Marshall()))
	}

	return nil

}
