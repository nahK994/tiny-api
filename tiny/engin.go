package tiny

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
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
	pathPattern, pathParamKeys, handler, ok := e.router.MatchRoute(r.Method, r.URL.Path)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 not found"))
		return
	}
	urlParams := make(map[string]any)

	re := regexp.MustCompile("^" + string(pathPattern) + "$")
	re.FindStringSubmatch(r.URL.Path)

	matches := re.FindStringSubmatch(r.URL.Path)

	for i := range pathParamKeys {
		val := matches[i+1] // i+1 because matches[0] is the full match
		if intval, err := strconv.Atoi(val); err == nil {
			urlParams[pathParamKeys[i]] = intval
		} else {
			urlParams[pathParamKeys[i]] = val
		}
	}

	ctx := NewContext(w, r, urlParams)
	// fmt.Println("Path -->", r.URL.Path, "| Method -->", r.Method)
	handler(ctx)
}

func (e *Engine) Run(addr string) error {
	fmt.Println("Server is starting on: ", addr)
	return http.ListenAndServe(addr, e)
}
