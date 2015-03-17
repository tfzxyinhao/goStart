package main

import (
	"log"
	"net/http"
	"proj/context"
)

type MyHandler struct {
	Db          *Database
	RouterMatch *Router
	CrossDomain bool
}

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

func (h *MyHandler) AddRouter(path string, cb func(c *context.Context)) {
	h.RouterMatch.Register(path, cb)
}

func (h *MyHandler) EnableCrossDomain(enable bool) {
	h.CrossDomain = enable
}

func (h *MyHandler) Start(addr string) {
	http.ListenAndServe(addr, h)
}

func NewMyHanlder() *MyHandler {
	handler := &MyHandler{RouterMatch: &Router{Handler: make(map[string]func(c *context.Context)), SubRouter: make(map[string]*Router)}, Db: &Database{}, CrossDomain: false}
	//handler.Db.Init()
	return handler
}
