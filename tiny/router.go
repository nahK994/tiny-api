package tiny

import (
	"regexp"
	"strconv"
)

type RouteKey struct {
	PathPattern pathPattern
	Method      string
}

type Router struct {
	handlers      map[RouteKey]*HandlerFunc
	pathParamKeys map[RouteKey][]string
}

func NewRouter() *Router {
	return &Router{
		handlers:      make(map[RouteKey]*HandlerFunc),
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

	r.handlers[routeKeys] = &handler
	r.pathParamKeys[routeKeys] = pathParamKeys
}

func (r *Router) extractPathParam(method string, actualPath string) (pathPattern, map[string]any, bool) {
	for pattern := range r.handlers {
		if pattern.Method != method {
			continue
		}

		re := regexp.MustCompile("^" + string(pattern.PathPattern) + "$")
		if re.MatchString(actualPath) {
			urlParams := make(map[string]any)
			matches := re.FindStringSubmatch(actualPath)
			routeKeys := RouteKey{
				PathPattern: pattern.PathPattern,
				Method:      method,
			}

			paramKeysByRoute := r.pathParamKeys[routeKeys]
			for i := range paramKeysByRoute {
				val := matches[i+1] // i+1 because matches[0] is the full match
				if intval, err := strconv.Atoi(val); err == nil {
					urlParams[paramKeysByRoute[i]] = intval
				} else {
					urlParams[paramKeysByRoute[i]] = val
				}
			}

			return pattern.PathPattern, urlParams, true
		}
	}
	return "", nil, false
}

func (r *Router) ResolveHandler(method string, actualPath string) (map[string]any, HandlerFunc, bool) {
	pathPattern, pathParams, ok := r.extractPathParam(method, actualPath)
	if !ok {
		return nil, nil, false
	}

	routeKeys := RouteKey{
		PathPattern: pathPattern,
		Method:      method,
	}

	handler, existsHandler := r.handlers[routeKeys]
	if !existsHandler {
		return nil, nil, false
	}

	return pathParams, *handler, true
}
