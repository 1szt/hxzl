package gin

// Web 引擎启动（gin 版）
// 使用 gin 路由框架

import (
	"log"

	g "github.com/gin-gonic/gin"
)

// Port 服务端口
const Port = "9081"

// Engine 全局路由引擎
var Router = g.New()

// Run 启动服务（同步阻塞）
// 必须在所有模块挂载完成之后调用，
// 否则挂载与请求处理会产生数据竞争（DATA RACE）
func Run() {

	// 打印访问地址
	log.Printf("🌐 [Gin] 访问 http://localhost:%s", Port)

	// 同步启动，阻塞当前协程直到出错
	if err := Router.Run(":" + Port); err != nil {
		log.Fatalf("❌ [Gin] 致命错误：端口可能被占用或权限不足 | %v", err)
	}
}
