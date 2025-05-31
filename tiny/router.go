package tiny

type HandlerFunc func(*Context)

type RouteKey struct {
	PathPattern string
	Method      string
}

type RouteEntry struct {
	PathParamKeys []string
	Handler       HandlerFunc
}

type Router struct {
	handlers map[RouteKey]*RouteEntry
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[RouteKey]*RouteEntry),
	}
}

func (r *Router) AddRoute(method string, path string, handler HandlerFunc) {
	pathPattern, _ := getPathPattern(path)
	pathParamKeys, _ := getPathParamKeys(path)

	r.handlers[RouteKey{
		PathPattern: string(pathPattern),
		Method:      method,
	}] = &RouteEntry{
		PathParamKeys: pathParamKeys,
		Handler:       handler,
	}
}

func (r *Router) MatchRoute(method string, actualPath string) (pathPattern, pathParamKeys, HandlerFunc, bool) {
	pathPattern, err := getPathPattern(actualPath)
	if err != nil {
		return "", nil, nil, false
	}

	handler, exists := r.handlers[RouteKey{
		PathPattern: string(pathPattern),
		Method:      method,
	}]
	return pathPattern, handler.PathParamKeys, handler.Handler, exists
}
