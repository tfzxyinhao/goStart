package main

import (
	"flag"
	"github.com/golang/glog"
	"log"
	"net/http"
	"proj/context"
	"proj/user"
	"strings"
)

const (
	MIMEJSON              = "application/json"
	MIMEHTML              = "text/html"
	MIMEXML               = "application/xml"
	MIMEXML2              = "text/xml"
	MIMEPlain             = "text/plain"
	MIMEPOSTForm          = "application/x-www-form-urlencoded"
	MIMEMultipartPOSTForm = "multipart/form-data"
)

type Router struct {
	Handler   map[string]func(c *context.Context)
	SubRouter map[string]*Router
}

func (r *Router) Match(path string) func(c *context.Context) {
	paths := strings.Split(strings.Trim(path, "/"), "/")
	count := len(paths)

	if count == 1 {
		return r.Handler[paths[0]]
	} else {
		tmp := r
		for i := 0; i < count-1; i++ {
			dir := paths[i]
			sub := tmp.SubRouter[dir]
			if sub == nil {
				return nil
			}
			tmp = sub
		}

		last := paths[count-1]
		return tmp.Handler[last]
	}
}

func (r *Router) Register(path string, cb func(c *context.Context)) {
	paths := strings.Split(strings.Trim(path, "/"), "/")
	count := len(paths)

	if count == 1 {
		r.Handler[paths[0]] = cb
	} else {
		tmp := r
		for i := 0; i < count-1; i++ {
			dir := paths[i]
			sub := tmp.SubRouter[dir]
			if sub == nil {
				newSub := &Router{SubRouter: make(map[string]*Router), Handler: make(map[string]func(c *context.Context))}
				tmp.SubRouter[dir] = newSub
				tmp = newSub
				log.Println(i, dir, newSub, &newSub, &tmp)
			} else {
				tmp = sub
			}
		}

		last := paths[count-1]
		tmp.Handler[last] = cb
	}
}

type MyHandler struct {
	Db          *Database
	RouterMatch *Router
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	handler := h.RouterMatch.Match(r.URL.Path)
	c := &context.Context{Req: r, Res: w}

	if handler != nil {
		handler(c)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("Not Found"))
	}
}

func NewMyHanlder() *MyHandler {
	handler := &MyHandler{RouterMatch: &Router{Handler: make(map[string]func(c *context.Context)), SubRouter: make(map[string]*Router)}, Db: &Database{}}
	handler.Db.Init()
	return handler
}

func (h *MyHandler) AddRouter(path string, cb func(c *context.Context)) {
	h.RouterMatch.Register(path, cb)
}

func init() {
	flag.Set("log_dir", "./")
	flag.Set("stderrthreshold", "INFO")
}

func main() {
	handler := NewMyHanlder()
	handler.AddRouter("/user", user.GetProfile)
	handler.AddRouter("/user/edit", user.SetProfile)
	handler.AddRouter("/user/register", user.Register)
	handler.AddRouter("/user/login", user.Login)

	glog.Info("start")
	http.ListenAndServe(":80", handler)
}
