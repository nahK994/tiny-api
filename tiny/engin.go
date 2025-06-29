package tiny

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	urlPrefix string
	rootPath  string
}

type Engine struct {
	router *Router
	media  *Config
}

func New() *Engine {
	return &Engine{
		router: NewRouter(),
		media:  &Config{},
	}
}

func (e *Engine) SetMediaConfig(mediaUrl, mediaRoot string) {
	e.media.urlPrefix = mediaUrl
	e.media.rootPath = mediaRoot

	e.registerMediaRoute()
}

func (e *Engine) POST(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) {
	e.router.AddRoute("POST", path, handler, middlewares...)
}

func (e *Engine) GET(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) {
	e.router.AddRoute("GET", path, handler, middlewares...)
}

func (e *Engine) registerMediaRoute() {
	e.POST("/upload", func(c *Context) {
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid file upload"})
			return
		}
		defer file.Close()

		err = os.MkdirAll(e.media.rootPath, 0755)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create upload dir"})
			return
		}

		filePath := filepath.Join(e.media.rootPath, header.Filename)
		out, err := os.Create(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create file"})
			return
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not save file"})
			return
		}

		os.Chmod(filePath, 0644)

		c.JSON(http.StatusOK, map[string]string{
			"message": "Uploaded successfully",
			"url":     e.media.urlPrefix + "/" + header.Filename,
		})
	})
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, e.media.urlPrefix) {
		filePath := strings.TrimPrefix(r.URL.Path, e.media.urlPrefix)
		fullPath := filepath.Join(e.media.rootPath, filePath)

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
