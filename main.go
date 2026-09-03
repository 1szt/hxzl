package main

import (
	"mcup-server/gin"
	"mcup-server/ginuse"
	"mcup-server/home"
	"mcup-server/motd"
)

func main() {
	motd.Run()
	// 加载中间件
	ginuse.Run()
	// 加载模块
	home.Run()
	// 启动服务
	gin.Run()
}
