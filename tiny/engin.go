package tiny

import (
	"fmt"
	"net/http"
)

type Engine struct {
	router *Router
}

func New() *Engine {
	return &Engine{router: NewRouter()}
}

func (e *Engine) POST(path string, handler HandlerFunc) {
	e.router.AddRoute("POST", path, handler)
}

func (e *Engine) GET(path string, handler HandlerFunc) {
	e.router.AddRoute("GET", path, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urlParams, handler, ok := e.router.ResolveHandler(r.Method, r.URL.Path)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 not found"))
		return
	}

	ctx := NewContext(w, r, urlParams)
	// fmt.Println("Path -->", r.URL.Path, "| Method -->", r.Method)
	handler(ctx)
}

func (e *Engine) Run(addr string) error {
	fmt.Println("Server is starting on: ", addr)
	return http.ListenAndServe(addr, e)
}
