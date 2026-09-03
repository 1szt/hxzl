package main

import (
	"mcup-server/gin"
	"mcup-server/home"
	"mcup-server/middleware"
	"mcup-server/motd"
	"mcup-server/static"
	"mcup-server/template"
)

func main() {
	motd.Run()

	middleware.Run()
	static.Run()

	home.Run()

	template.Run()

	gin.Run()
}
