package gin

import (
	"log"

	// "hxzl/config"

	g "github.com/gin-gonic/gin"
)

var Router = g.New()

func Run() {

	// // 读取config数据库配置
	// var cfg config.Config

	// // 查询数据库name为"port"的配置
	// if err := config.DB.Where("name = ?", "port").First(&cfg).Error; err != nil {
	// 	log.Fatalf("❌ [Gin] 查询配置数据库失败 | %v", err)
	// }

	// log.Printf("🌐 [Gin] 访问 http://localhost:%s", cfg.Value)

	if err := Router.Run(":" + "9081"); err != nil {
		log.Fatalf("❌ [Gin] 致命错误：端口可能被占用或权限不足 | %v", err)
	}
}
