package tiny

type RouteGroup struct {
	path  string
	engin *Engine
}

func (e *Engine) Group(path string) *RouteGroup {
	return &RouteGroup{
		path:  path,
		engin: e,
	}
}

func (rg *RouteGroup) Group(path string) *RouteGroup {
	return &RouteGroup{
		path:  rg.path + path,
		engin: rg.engin,
	}
}

func (rg *RouteGroup) POST(path string, handler HandlerFunc) {
	rg.engin.router.AddRoute("POST", path, handler)
}

func (rg *RouteGroup) GET(path string, handler HandlerFunc) {
	rg.engin.router.AddRoute("GET", rg.path+path, handler)
}
