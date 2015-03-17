package main

import (
	"flag"
	"github.com/golang/glog"
	"net/http"
	"proj/user"
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

func init() {
	flag.Set("log_dir", "./logs")
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
