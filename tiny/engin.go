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

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := &Context{
		Writer:  w,
		Request: r,
	}
	if handler, ok := e.router.Handle(r.Method, r.URL.Path); ok {
		fmt.Printf("Got %s request on %s\n", r.Method, r.URL.Path)
		handler(ctx)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 not found"))
	}
}

func (e *Engine) Run(addr string) error {
	fmt.Println("Server is starting on: ", addr)
	return http.ListenAndServe(addr, e)
}
