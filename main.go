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
	flag.Set("log_dir", "./")
	flag.Set("stderrthreshold", "INFO")
}

func main() {
	j := NewJson()
	j.BeginObject("user")
	j.BeginArray("products")
	j.BeginObject("")
	j.Add("id", "123456")
	j.Add("id", 3.1415926)
	j.EndObject()
	j.EndArray()
	j.EndObject()
	glog.Infoln(j.ToString())
	return
	handler := NewMyHanlder()
	handler.AddRouter("/user", user.GetProfile)
	handler.AddRouter("/user/edit", user.SetProfile)
	handler.AddRouter("/user/register", user.Register)
	handler.AddRouter("/user/login", user.Login)

	glog.Info("start")
	http.ListenAndServe(":80", handler)
}
