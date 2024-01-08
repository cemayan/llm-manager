package types

import (
	"encoding/json"
)

type Error struct {
	NativeErr string `json:"native_err"`
	Code      int    `json:"code"`
	ErrGroup  string `json:"err_group"`
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error *Error      `json:"error,omitempty"`
}

func (r *Response) Marshall() []byte {
	marshal, err := json.Marshal(r)

	if err != nil {
		return []byte("")
	}

	return marshal
}
