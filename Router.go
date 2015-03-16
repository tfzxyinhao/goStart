package main

import (
	"proj/context"
	"strings"
)

type Router struct {
	Param     string
	Handler   map[string]func(c *context.Context)
	SubRouter map[string]*Router
}

func (r *Router) Match(path string) (func(c *context.Context), map[string]string) {
	paths := strings.Split(strings.Trim(path, "/"), "/")
	count := len(paths)

	if count == 1 {
		param := make(map[string]string)
		count := len(r.Param)
		if count > 0 {
			param[string(r.Param[1:count])] = paths[0]
		}
		return r.Handler[paths[0]], param
	} else {
		tmp := r
		param := make(map[string]string)
		for i := 0; i < count-1; i++ {
			count := len(r.Param)
			if count > 0 {
				param[string(tmp.Param[1:count])] = paths[i]
				sub := tmp.SubRouter[tmp.Param]
				if sub == nil {
					return nil, nil
				}
				tmp = sub
			} else {
				dir := paths[i]
				sub := tmp.SubRouter[dir]
				if sub == nil {
					return nil, nil
				}
				tmp = sub
			}
		}

		last := paths[count-1]
		count := len(r.Param)
		if count > 0 {
			param[string(tmp.Param[1:count])] = last
			return tmp.Handler[tmp.Param], param
		} else {
			return tmp.Handler[last], param
		}
	}
}

func (r *Router) Register(path string, cb func(c *context.Context)) {
	paths := strings.Split(strings.Trim(path, "/"), "/")
	count := len(paths)

	if count == 1 {
		dir := paths[0]
		if dir[0] == 58 { // check token it's colon
			r.Param = dir
		}
		r.Handler[dir] = cb

	} else {
		tmp := r
		for i := 0; i < count-1; i++ {
			dir := paths[i]
			sub := tmp.SubRouter[dir]
			if sub == nil {
				newSub := &Router{SubRouter: make(map[string]*Router), Handler: nil, Param: ""}
				if dir[0] == 58 { // check token it's colon
					tmp.Param = dir
				}
				tmp.SubRouter[dir] = newSub
				tmp = newSub
			} else {
				tmp = sub
			}
		}

		last := paths[count-1]
		if tmp.Handler == nil {
			tmp.Handler = make(map[string]func(c *context.Context))
		}

		if last[0] == 58 {
			tmp.Param = last
		}
		tmp.Handler[last] = cb
	}
}
