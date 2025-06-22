package tiny

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Engine struct {
	router       *Router
	staticRoot   string
	staticPrefix string
}

func New() *Engine {
	return &Engine{
		router: NewRouter(),
	}
}

func (e *Engine) SetStaticConfig(staticUrl, staticRoot string) {
	e.staticPrefix = staticUrl
	e.staticRoot = staticRoot
}

func (e *Engine) POST(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) {
	e.router.AddRoute("POST", path, handler, middlewares...)
}

func (e *Engine) GET(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) {
	e.router.AddRoute("GET", path, handler, middlewares...)
}

func (e *Engine) RegisterStaticRoutes() {
	e.POST("/upload", func(c *Context) {
		file, header, _ := c.Request.FormFile("file")
		defer file.Close()

		// Ensure upload folder exists
		os.MkdirAll(e.staticRoot, os.ModePerm)

		// Save file
		filePath := filepath.Join(e.staticRoot, header.Filename)
		os.Create(filePath)

		c.JSON(http.StatusOK, map[string]string{
			"message":  "Uploaded successfully",
			"filename": header.Filename,
			"url":      e.staticPrefix + header.Filename,
		})
	})
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, e.staticPrefix) {
		filePath := strings.TrimPrefix(r.URL.Path, e.staticPrefix)
		fullPath := filepath.Join(e.staticRoot, filePath)

		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, fullPath)
		return
	}

	urlParams, handler, ok := e.router.ResolveHandler(r.Method, r.URL.Path)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 not found"))
		return
	}

	ctx := NewContext(w, r, urlParams)
	handler(ctx)
}

func (e *Engine) Run(addr string) error {
	fmt.Println("Server is starting on: ", addr)
	return http.ListenAndServe(addr, e)
}
