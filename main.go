package main

import (
	"mcup-server/gin"
	"mcup-server/gininit"
	"mcup-server/home"
	"mcup-server/motd"
)

func main() {
	motd.Run()

	gininit.Run()

	home.Run()

	gin.Run()
}
