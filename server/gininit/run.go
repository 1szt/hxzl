package gininit

import (
	web "server/gin"

	g "github.com/gin-gonic/gin"
)

func Run() {

	// 捕获所有内部 panic
	web.Router.Use(g.Recovery())
	// 记录日志
	web.Router.Use(g.Logger())
}
