package tiny

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Writer    http.ResponseWriter
	Request   *http.Request
	PathParam map[string]any
}

func NewContext(w http.ResponseWriter, r *http.Request, pathParam map[string]any) *Context {

	return &Context{
		Writer:    w,
		Request:   r,
		PathParam: pathParam,
	}
}

func (c *Context) JSON(code int, obj any) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(code)
	json.NewEncoder(c.Writer).Encode(obj)
}

func (c *Context) BindJSON(obj any) error {
	return json.NewDecoder(c.Request.Body).Decode(obj)
}
