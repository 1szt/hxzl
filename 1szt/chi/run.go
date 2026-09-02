package chi

// Web 引擎启动
// 使用 chi 路由库

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Port 服务端口
const Port = "9081"

// Mux 全局路由底座
var Mux = chi.NewRouter()

// Run 启动服务（同步阻塞）
// 必须在所有模块挂载（Mount）完成之后调用，
// 否则挂载与请求处理会产生数据竞争（DATA RACE）
func Run() {

	// 打印访问地址
	log.Printf("🌐 [Chi] 访问 http://localhost:%s", Port)

	// 同步启动，阻塞当前协程直到出错
	if err := http.ListenAndServe(":"+Port, Mux); err != nil {
		log.Fatalf("❌ [Chi] 致命错误：端口可能被占用或权限不足 | %v", err)
	}
}
