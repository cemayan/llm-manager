package commit

import (
	"encoding/json"
	"git-observer/api/types"
	"git-observer/api/util"
	"git-observer/internal/backend"
	"git-observer/internal/config"
	"net/http"
	"regexp"
)

type handler struct {
	backend backend.Backend
}

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
	resp, err := h.backend.Query(
		util.PrepareOllamaRequest(config.AppConfig.Config.Ollama.Model, payload.Message).Marshall(),
		nil)

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
		w.Write(util.PrepareResponse(ollamaResp).Marshall())
	case "langchaingo":
		panic("not implemented!")
	case "linhoose":
		panic("not implemented!")
	}

}
