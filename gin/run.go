package gin

import (
	"log"

	g "github.com/gin-gonic/gin"
)

var Router = g.New()

func Run() {

	log.Printf("🌐 [Gin] 访问 http://localhost:%s", "9081")

	if err := Router.Run(":" + "9081"); err != nil {
		log.Fatalf("❌ [Gin] 致命错误：端口可能被占用或权限不足 | %v", err)
	}
}
