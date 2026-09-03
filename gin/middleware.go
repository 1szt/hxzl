package gin

import (
	g "github.com/gin-gonic/gin"
)

// Middleware 武装全局中间件
// 必须在任何路由挂载、服务启动之前调用，
// 以确保中间件在请求到达前就绪。
func Middleware() {
	// 捕获所有内部 panic
	Router.Use(g.Recovery())
	// 记录日志
	Router.Use(g.Logger())
}
