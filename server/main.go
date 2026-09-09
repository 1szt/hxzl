package main

import (
	"server/config"
	"server/gin"
	"server/gininit"
	"server/home"
	"server/motd"
)

func main() {

	motd.Run()

	config.Run()

	gininit.Run()

	home.Run()

	gin.Run()
}
