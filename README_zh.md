[![GoDoc](https://godoc.org/github.com/tfzxyinhao/proj?status.svg)](https://godoc.org/github.com/tfzxyinhao/proj)
[![Build Status](https://travis-ci.org/tfzxyinhao/proj.svg?branch=master)](https://travis-ci.org/tfzxyinhao/proj)

# goStart
本项目是用于开发web api的基础代码,包含Router,基础的json序列化等
你可以直接clone一份本代码,然后在此基础上开发你的web api.

#功能：

> * 类似 express 3.x的路由
> * URL参数
> * 基础的json序列化
> * some convenient function

#为什么:
从NodeJs转到go,习惯了使用express框架的我在go中找不到类似的框架,虽然有martin和gin
但是我任然很难习惯,所以就写了一些代码模仿express的功能,让自己尽快的适应go中web api的
开发工作.

#需要做的:

> * Form表达解析(各种MIME)
> * 更丰富的json序列化支持
> * Cookies操作
> * Session支持
> * 中间件
> * 脚手架

#Usage

```go
	handler := NewHanlder()
	handler.AddRouter("/user/login", func(c *context.Context) {
		j := NewJson()
	    j.BeginObject("user")
		j.BeginArray("products")
		j.BeginObject("")
		j.Add("id", "123456")
		j.Add("id", 3.1415926)
		j.EndObject()
		j.EndArray()
		j.EndObject()
		c.Json(j.ToString())
	})
	
	handler.AddRouter("/user/:userid/profile", func(c *context.Context) {
		j := NewJson()
	    j.BeginObject("user")
		j.Add("id", c.Params["userid"])
		j.Add("username", "yourname")
		j.Add("avatar", "http://git.oschina.net/logo.gif")
		j.Add("balance", 3.1415926)
		j.EndObject()
		c.Json(j.ToString())
	})
	handler.Start(":80")
```

#性能测试

 #路由

    PASS
    BenchmarkRouter	 1000000	      1013 ns/op
    ok  	proj	1.064s

