package router

type baseRoute struct {
	method      string
	path        string
	handlerFunc HandlerFunc
}

func (b baseRoute) Method() string {
	return b.method
}

func (b baseRoute) Path() string {
	return b.path
}

func (b baseRoute) Handler() HandlerFunc {
	return b.handlerFunc
}
func CreateRoute(method string, path string, handlerFunc HandlerFunc) Route {
	return NewRoute(method, path, handlerFunc)
}

func NewRoute(method string, path string, handlerFunc HandlerFunc) Route {
	return &baseRoute{method: method, path: path, handlerFunc: handlerFunc}
}
