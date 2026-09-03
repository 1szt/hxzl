package ginuse

// 底座初始化（gin 版）：武装全局中间件
// 必须在任何模块挂载之前调用，
// 以确保中间件在请求到达前就绪

import (
	web "mcup-server/gin"

	g "github.com/gin-gonic/gin"
)

// Run 武装底座：设置全局中间件
func Run() {

	// 捕获所有内部 panic
	web.Engine.Use(g.Recovery())
	// 记录日志
	web.Engine.Use(g.Logger())

}
