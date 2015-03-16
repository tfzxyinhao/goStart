package main

import (
	"proj/user"
	"testing"
)

func BenchmarkRouter(b *testing.B) {
	b.StopTimer()
	handler := NewMyHanlder()
	handler.AddRouter("/user", user.GetProfile)
	handler.AddRouter("/user/edit", user.SetProfile)
	handler.AddRouter("/user/register", user.Register)
	handler.AddRouter("/user/login", user.Login)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		handler.RouterMatch.Match("/user/login")
	}
}
