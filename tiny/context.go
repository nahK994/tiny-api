package tiny

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func (c *Context) JSON(code int, obj interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(code)
	json.NewEncoder(c.Writer).Encode(obj)
}

func (c *Context) BindJSON(obj interface{}) error {
	return json.NewDecoder(c.Request.Body).Decode(obj)
}
