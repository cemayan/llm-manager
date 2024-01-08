package util

import (
	"git-observer/api/types"
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

func PrepareResponse(data interface{}) *types.Response {
	resp := types.Response{
		Data: data,
	}
	return &resp
}

func PrepareErrorResponse(err error, code int) *types.Response {
	resp := types.Response{
		Error: &types.Error{NativeErr: err.Error(), Code: code, ErrGroup: errorMap[code]},
	}
	return &resp
}
