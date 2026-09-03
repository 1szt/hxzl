package main

import (
	"mcup-server/config"
	"mcup-server/gin"
	"mcup-server/gininit"
	"mcup-server/home"
	"mcup-server/motd"
)

func main() {
	motd.Run()
	config.Run()

	gininit.Run()

	home.Run()

	gin.Run()
}
