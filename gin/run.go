package gin

import (
	"log"
	"strconv"

	g "github.com/gin-gonic/gin"
	"mcup-server/config"

)

var Router = g.New()

func Run() {

	// 读取config数据库配置
	var cfg config.Config
	if err := config.DB.First(&cfg).Error; err != nil {
		log.Fatalf("❌ [Gin] 读取配置失败 | %v", err)
	}

	log.Printf("🌐 [Gin] 访问 http://localhost:%d", cfg.Port)

	if err := Router.Run(":" + strconv.Itoa(cfg.Port)); err != nil {
		log.Fatalf("❌ [Gin] 致命错误：端口可能被占用或权限不足 | %v", err)
	}
}
