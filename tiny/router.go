package tiny

type HandlerFunc func(*Context)

type Router struct {
	handlers map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *Router) AddRoute(method string, path string, handler HandlerFunc) {
	key := method + "-" + path
	r.handlers[key] = handler
}

func (r *Router) Handle(method string, path string) (HandlerFunc, bool) {
	handler, exists := r.handlers[method+"-"+path]
	return handler, exists
}
