package commit

import (
	"llm-manager/api/router"
	"llm-manager/internal/backend"
	"net/http"
)

type commitRouter struct {
	routes []router.Route
}

func (cr commitRouter) Routes() []router.Route {

	h := handler{backend: backend.BackendInstance}

	cr.routes = []router.Route{
		router.CreateRoute(http.MethodPost, "/commit", h.QueryHandler),
	}
	return cr.routes
}

func NewRouter() router.Router {
	return &commitRouter{}
}
