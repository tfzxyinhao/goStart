package context

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Req *http.Request       "request"
	Res http.ResponseWriter "response"
}

func (c *Context) Json(data interface{}) error {
	c.Res.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(data)
	if err == nil {
		c.Res.Write(j)
	}
	return err
}

func (c *Context) Jsonp(data interface{}) error {
	c.Res.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(data)
	if err == nil {
		c.Res.Write(j)
	}
	return err
}
