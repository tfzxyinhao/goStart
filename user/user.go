package user

import (
	"proj/context"
)

func Register(c *context.Context) {
	c.Json("Register")
}

func Login(c *context.Context) {
	c.Json("Login")
}

func GetProfile(c *context.Context) {
	c.Json("GetProfile")
}

func SetProfile(c *context.Context) {
	c.Json("SetProfile")
}
