package main

import (
	"proj/user"
)

func main() {
	handler := NewHanlder()
	handler.AddRouter("/user", user.GetProfile)
	handler.AddRouter("/user/edit", user.SetProfile)
	handler.AddRouter("/user/register", user.Register)
	handler.AddRouter("/user/login", user.Login)
	handler.Start(":80")
}
