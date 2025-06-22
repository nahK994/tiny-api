package tiny

type RouteGroup struct {
	path        string
	engin       *Engine
	middlewares []MiddlewareFunc
}

func (e *Engine) Group(path string) *RouteGroup {
	return &RouteGroup{
		path:        path,
		engin:       e,
		middlewares: []MiddlewareFunc{},
	}
}

func (rg *RouteGroup) Group(path string) *RouteGroup {
	return &RouteGroup{
		path:        rg.path + path,
		engin:       rg.engin,
		middlewares: rg.middlewares,
	}
}

func (rg *RouteGroup) Use(middlewares ...MiddlewareFunc) {
	rg.middlewares = append(rg.middlewares, middlewares...)
}

func (rg *RouteGroup) POST(path string, handler HandlerFunc) {
	rg.engin.router.AddRoute("POST", path, handler, rg.middlewares...)
}

func (rg *RouteGroup) GET(path string, handler HandlerFunc) {
	rg.engin.router.AddRoute("GET", rg.path+path, handler, rg.middlewares...)
}
