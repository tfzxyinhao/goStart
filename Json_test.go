package main

import (
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {
	j := NewJson()
	j.BeginObject("user")
	j.BeginArray("products")
	j.BeginObject("")
	j.Add("id", "123456")
	j.Add("id", 3.1415926)
	j.EndObject()
	j.EndArray()
	j.EndObject()
	fmt.Println(j.ToString())
}
