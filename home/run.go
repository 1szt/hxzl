package home

import (
	web "mcup-server/gin"
	g "github.com/gin-gonic/gin"
	"log"
)


func Run() {
	web.Router.GET("/", home)

	log.Print("✅ [Home] 主页模块 加载完成！")
}

func home(c *g.Context) {
	c.String(200, "你好，世界！")
}
