package ginhome

// 依赖模块：web（gin）
// 主页模块负责处理用户访问根路径的请求
// 挂载路径："/"

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"

	web "1szt/gin"

	g "github.com/gin-gonic/gin"
)

//go:embed templates static
var content embed.FS

func Run() {

	// 静态文件服务
	staticFS, _ := fs.Sub(content, "static")
	web.Engine.StaticFS("/static", http.FS(staticFS))

	// 主页
	web.Engine.GET("/", home)

	// 最后一炮打到 web 底座，搞定！
	log.Print("✅ [GinHome] 主页模块 加载完成！")
}

func home(c *g.Context) {
	tmpl := template.Must(template.ParseFS(content, "templates/index.html"))
	c.Header("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(c.Writer, nil); err != nil {
		http.Error(c.Writer, "模板错误: "+err.Error(), http.StatusInternalServerError)
	}
}
