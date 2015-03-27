package main

import (
	"fmt"
	"github.com/tfzxyinhao/proj/user"
	"testing"
)

//func BenchmarkRouter(b *testing.B) {
//	b.StopTimer()
//	handler := NewHanlder()
//	handler.AddRouter("/user", user.GetProfile)
//	handler.AddRouter("/user/edit", user.SetProfile)
//	handler.AddRouter("/user/register", user.Register)
//	handler.AddRouter("/user/login", user.Login)
//	handler.AddRouter("/user/:userid/profile", user.Login)
//	b.StartTimer()
//	for i := 0; i < b.N; i++ {
//		handler.RouterMatch.Match("/user/123456/profile")
//	}
//}

func TestRouter(t *testing.T) {
	handler := NewHanlder()
	handler.AddRouter("/user", user.GetProfile)
	handler.AddRouter("/user/edit", user.SetProfile)
	handler.AddRouter("/user/register", user.Register)
	handler.AddRouter("/user/login", user.Login)
	handler.AddRouter("/user/:userid/profile", user.Login)
	fmt.Println("-----output-----")
	fmt.Println(handler.RouterMatch.Match("/user/123456/profile"))
	fmt.Println(handler.RouterMatch.Match("/user/123456"))
	fmt.Println(handler.RouterMatch.Match("/user"))
	fmt.Println("------end-------")
	fmt.Println()
}

func TestRootRouter(t *testing.T) {
	handler := NewHanlder()
	handler.AddRouter("/user", user.GetProfile)
	handler.AddRouter("/user/edit", user.SetProfile)
	handler.AddRouter("/user/register", user.Register)
	handler.AddRouter("/user/login", user.Login)
	handler.AddRouter("/:userid", user.Login)
	fmt.Println("-----output-----")
	fmt.Println(handler.RouterMatch.Match("/123456"))
	fmt.Println(handler.RouterMatch.Match("/user/login"))
	fmt.Println("------end-------")
	fmt.Println()
}

func TestNodeRouter(t *testing.T) {
	handler := NewHanlder()
	handler.AddRouter("/user", user.GetProfile)
	handler.AddRouter("/user/edit", user.SetProfile)
	handler.AddRouter("/user/register", user.Register)
	handler.AddRouter("/user/login", user.Login)
	handler.AddRouter("/user/:userid", user.Login)
	fmt.Println("-----output-----")
	fmt.Println(handler.RouterMatch.Match("/user/123456"))
	fmt.Println(handler.RouterMatch.Match("/user/123456/profile"))
	fmt.Println(handler.RouterMatch.Match("/user/login"))
	fmt.Println("------end-------")
	fmt.Println()
}
