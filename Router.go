package main

import (
	"log"
	"proj/context"
	"strings"
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
