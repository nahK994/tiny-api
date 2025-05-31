package tiny

import (
	"regexp"
)

type HandlerFunc func(*Context)

type RouteKey struct {
	PathPattern pathPattern
	Method      string
}

type RouteEntry struct {
	PathParamKeys pathParamKeys
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
		PathPattern: pathPattern,
		Method:      method,
	}] = &RouteEntry{
		PathParamKeys: pathParamKeys,
		Handler:       handler,
	}
}

func (r *Router) MatchRoute(method string, actualPath string) (pathPattern, pathParamKeys, HandlerFunc, bool) {
	var pathPattern pathPattern
	for pattern := range r.handlers {
		if pattern.Method != method {
			continue
		}
		re := regexp.MustCompile("^" + string(pattern.PathPattern) + "$")
		if re.MatchString(actualPath) {
			pathPattern = pattern.PathPattern
			break
		}
	}

	handler, exists := r.handlers[RouteKey{
		PathPattern: pathPattern,
		Method:      method,
	}]
	if !exists {
		return "", nil, nil, false
	}

	return pathPattern, handler.PathParamKeys, handler.Handler, exists
}
