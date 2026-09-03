package main

import (
	"mcup-server/config"
	"mcup-server/gin"
	"mcup-server/gininit"
	"mcup-server/home"
	"mcup-server/motd"
	"mcup-server/setup"
)

func main() {
	
	motd.Run()

	config.Run()

	setup.Run()

	gininit.Run()

	home.Run()

	gin.Run()
}
