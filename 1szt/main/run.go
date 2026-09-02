package main

import (
	"1szt/boot"
	"1szt/chi"
	"1szt/home"
	"1szt/motd"
)

func main() {
	motd.Run()

	// 加载中间件
	boot.Run()
	// 加载模块
	home.Run()
	// 启动服务
	chi.Run()
}
