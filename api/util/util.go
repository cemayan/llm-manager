package util

import (
	"llm-manager/api/types"
	"llm-manager/internal/config"
	"llm-manager/internal/structs"
)

var errorMap map[int]string

func init() {

	errorMap = make(map[int]string)

	errorMap[1001] = "JSON_UNMARSHALL_ERROR"
	errorMap[1002] = "IO_READ_ERROR"
	errorMap[1003] = "SERVER_ERROR"
}

func PrepareOllamaRequest(model string, commitMessage string) *types.OllamaRequest {
	return &types.OllamaRequest{
		Model:  model,
		Prompt: commitMessage,
	}
}

func PrepareResponse(data interface{}) []byte {

	switch config.AppConfig.Config.Api.Output {
	case structs.Text:
		if val, ok := data.(types.OllamaResponse); ok {
			return []byte(val.Response)
		}

	case structs.Json:
		resp := types.Response{
			Data: data,
		}
		return resp.Marshall()
	case structs.Yaml:
		return nil
	}

	return nil
}

func PrepareErrorResponse(err error, code int) *types.Response {
	resp := types.Response{
		Error: &types.Error{NativeErr: err.Error(), Code: code, ErrGroup: errorMap[code]},
	}
	return &resp
}
