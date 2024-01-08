package cli

import (
	"encoding/json"
	"fmt"
	"git-observer/api/types"
	"git-observer/api/util"
	"git-observer/internal/backend"
	"git-observer/internal/config"
)

func Exec(prompt string) {
	if prompt != "" {

		switch config.AppConfig.Config.Api.Backend {
		case "ollama":
			query, err := backend.BackendInstance.
				Query(util.PrepareOllamaRequest(config.AppConfig.Config.Ollama.Model, prompt).Marshall(), nil)
			if err != nil {
				fmt.Println(string(util.PrepareErrorResponse(err, 1003).Marshall()))
				return
			}

			var ollamaResp types.OllamaResponse
			err = json.Unmarshal(query, &ollamaResp)

			if err != nil {
				fmt.Println(string(util.PrepareErrorResponse(err, 1001).Marshall()))
				return
			}

			fmt.Println(string(util.PrepareResponse(ollamaResp.Response).Marshall()))
		case "lingoose":
			query, err := backend.BackendInstance.
				Query([]byte(prompt), nil)
			if err != nil {
				fmt.Println(string(util.PrepareErrorResponse(err, 1003).Marshall()))
				return
			}

			fmt.Println(string(util.PrepareResponse(query).Marshall()))
		case "langchaingo":
			query, err := backend.BackendInstance.
				Query([]byte(prompt), nil)
			if err != nil {
				fmt.Println(string(util.PrepareErrorResponse(err, 1003).Marshall()))
				return
			}

			fmt.Println(string(util.PrepareResponse(query).Marshall()))
		}

		return
	}
}
