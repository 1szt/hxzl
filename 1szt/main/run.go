package main

import (
	"1szt/gin"
	"1szt/ginuse"
	"1szt/home"
	"1szt/motd"
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
