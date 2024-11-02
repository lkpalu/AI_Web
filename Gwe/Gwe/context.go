package Gwe

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	W          http.ResponseWriter
	R          *http.Request
	Method     string
	Path       string
	StatusCode int
	Params     map[string]string
	Headers    []HandlerFunc
	index      int
}

func (c *Context) Param(key string) string {
	val, _ := c.Params[key]
	return val
}
func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		W:      w,
		R:      r,
		Method: r.Method,
		Path:   r.URL.Path,
		index:  -1,
	}
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.W.WriteHeader(code)
}

func (c *Context) PostForm(key string) string {
	return c.R.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.R.URL.Query().Get(key)
}

func (c *Context) SetHeader(key, val string) {
	c.W.Header().Set(key, val)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.W.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.W)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.W, err.Error(), 500)
	}
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.W.Write([]byte(html))
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.W.Write(data)
}
func (c *Context) Next() {
	c.index++
	s := len(c.Headers)
	for ; c.index < s; c.index++ {
		c.Headers[c.index](c)
	}
}

func (c *Context) ShouldBindJSON(Stu any) error {
	body := c.R.Body
	defer body.Close()
	jsBody := json.NewDecoder(body)
	err := jsBody.Decode(Stu)
	return err
}
