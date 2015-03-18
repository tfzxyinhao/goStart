package main

import (
	"github.com/tfzxyinhao/proj/user"
)

func main() {
	handler := NewHanlder()
	handler.AddRouter("/user", user.GetProfile)
	handler.Start(":80")
}
