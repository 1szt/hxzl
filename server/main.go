package main

import (
	"hxzl/config"
	"hxzl/gin"
	"hxzl/gininit"
	"hxzl/home"
	"hxzl/motd"
)

func main() {

	motd.Run()

	config.Run()

	gininit.Run()

	home.Run()

	gin.Run()
}
