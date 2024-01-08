package router

import "net/http"

type HandlerFunc func(w http.ResponseWriter, r *http.Request)
type Router interface {
	Routes() []Route
}

type Route interface {
	Handler() HandlerFunc
	Method() string
	Path() string
}
