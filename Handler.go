package main

import (
	"github.com/tfzxyinhao/proj/context"
	"log"
	"net/http"
)

// http server request handler
type MyHandler struct {
	Db          *Database
	RouterMatch *Router
	CrossDomain bool
}

// default http request handler
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	handler, param := h.RouterMatch.Match(r.URL.Path)
	c := &context.Context{Params: param, Req: r, Res: w}

	if handler != nil {
		if h.CrossDomain {
			w.Header().Add("Access-Control-Allow-Origin", "*")
		}
		handler(c)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("Not Found"))
	}
}

// register handler for path
func (h *MyHandler) AddRouter(path string, cb func(c *context.Context)) {
	h.RouterMatch.Register(path, cb)
}

// enable or disable crossdomain
func (h *MyHandler) EnableCrossDomain(enable bool) {
	h.CrossDomain = enable
}

// start to handle http request
func (h *MyHandler) Start(addr string) {
	http.ListenAndServe(addr, h)
}

// create and init http hander
func NewHanlder() *MyHandler {
	handler := &MyHandler{RouterMatch: &Router{Handler: make(map[string]func(c *context.Context)), SubRouter: make(map[string]*Router)}, Db: &Database{}, CrossDomain: false}
	//handler.Db.Init()
	return handler
}
