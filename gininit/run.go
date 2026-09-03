package gininit

import (
	web "mcup-server/gin"

	g "github.com/gin-gonic/gin"
)

func Run() {

	// 捕获所有内部 panic
	web.Router.Use(g.Recovery())
	// 记录日志
	web.Router.Use(g.Logger())

	web.Router.LoadHTMLGlob("template/*.html")

	web.Router.StaticFile("/favicon.ico", "./favicon.ico")
	web.Router.Static("/static", "./static")
}
