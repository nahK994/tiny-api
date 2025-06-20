package tiny

import (
	"regexp"
)

type HandlerFunc func(*Context)

type RouteKey struct {
	PathPattern pathPattern
	Method      string
}

type Router struct {
	handler       map[RouteKey]*HandlerFunc
	pathParamKeys map[RouteKey][]string
}

func NewRouter() *Router {
	return &Router{
		handler:       make(map[RouteKey]*HandlerFunc),
		pathParamKeys: make(map[RouteKey][]string),
	}
}

func (r *Router) AddRoute(method string, path string, handler HandlerFunc) {
	pathPattern, _ := getPathPattern(path)
	pathParamKeys, _ := getPathParamKeys(path)
	routeKeys := RouteKey{
		PathPattern: pathPattern,
		Method:      method,
	}

	r.handler[routeKeys] = &handler
	r.pathParamKeys[routeKeys] = pathParamKeys
}

func (r *Router) MatchRoute(method string, actualPath string) (pathPattern, pathParamKeys, HandlerFunc, bool) {
	var pathPattern pathPattern
	for pattern := range r.handler {
		if pattern.Method != method {
			continue
		}
		re := regexp.MustCompile("^" + string(pattern.PathPattern) + "$")
		if re.MatchString(actualPath) {
			pathPattern = pattern.PathPattern
			break
		}
	}

	routeKeys := RouteKey{
		PathPattern: pathPattern,
		Method:      method,
	}

	handler, existsHandler := r.handler[routeKeys]
	if !existsHandler {
		return "", nil, nil, false
	}

	pathParamKeys, existsPathParamKeys := r.pathParamKeys[routeKeys]
	if !existsPathParamKeys {
		return "", nil, nil, false
	}

	return pathPattern, pathParamKeys, *handler, existsPathParamKeys && existsHandler
}
