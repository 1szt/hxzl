package main

import (
	"mcup-server/gin"
	"mcup-server/home"
	"mcup-server/motd"
)

func main() {
	motd.Run()

	gin.Middleware()
	gin.Static()
	gin.Template()

	home.Run()

	gin.Run()
}
