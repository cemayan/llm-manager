package commit

import (
	"encoding/json"
	"llm-manager/api/types"
	"llm-manager/api/util"
	"llm-manager/internal/backend"
	"llm-manager/internal/config"
	"net/http"
	"regexp"
)

type handler struct {
	backend backend.Backend
}

// QueryHandler handles the incoming request
// According to given backend it will be redirected to selected backend
func (h *handler) QueryHandler(w http.ResponseWriter, r *http.Request) {

	var payload types.CommitPayload

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(util.PrepareErrorResponse(err, 1001).Marshall())
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var body []byte

	switch config.AppConfig.Config.Api.Backend {
	case "ollama":
		body = util.PrepareOllamaRequest(config.AppConfig.Config.Ollama.Model, payload.Prompt).Marshall()
	case "langchaingo":
		body = []byte(payload.Prompt)
	case "lingoose":
		body = []byte(payload.Prompt)
	}

	resp, err := h.backend.Query(body, nil)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(util.PrepareErrorResponse(err, 1003).Marshall())
		return
	}

	switch config.AppConfig.Config.Api.Backend {
	case "ollama":

		var ollamaResp types.OllamaResponse
		err = json.Unmarshal(resp, &ollamaResp)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(util.PrepareErrorResponse(err, 1001).Marshall())
			return
		}

		re := regexp.MustCompile(`\r?\n`)
		ollamaResp.Response = re.ReplaceAllString(ollamaResp.Response, "")

		w.WriteHeader(http.StatusOK)
		w.Write(util.PrepareResponse(ollamaResp))
	case "langchaingo":
		w.WriteHeader(http.StatusOK)
		w.Write(util.PrepareResponse(resp))
	case "lingoose":
		w.WriteHeader(http.StatusOK)
		w.Write(util.PrepareResponse(resp))
	}

}
