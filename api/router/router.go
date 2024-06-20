package router

import "net/http"

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// Router represents the methods  that need to be implemented
type Router interface {
	Routes() []Route
}

// Route represents the methods that need to be implemented
type Route interface {
	Handler() HandlerFunc
	Method() string
	Path() string
}
