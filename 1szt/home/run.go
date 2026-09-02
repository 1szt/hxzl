package home

// 依赖模块 ：web
// 主页模块负责处理用户访问根路径的请求
// 挂载路径："/"

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"

	web "1szt/chi"

	"github.com/go-chi/chi/v5"
)

//go:embed templates static
var content embed.FS

func Run() {
	r := chi.NewRouter()
	// ... 这里可以无脑复制官方文档的代码 ...

	// 静态文件服务
	staticFS, _ := fs.Sub(content, "static")
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	r.Group(func(r chi.Router) {
		r.Get("/", home)
	})

	// 最后一炮打到web底座，搞定！
	web.Mux.Mount("/", r)
	log.Print("✅ [Home] 主页模块 加载完成！")
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFS(content, "templates/index.html"))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "模板错误: "+err.Error(), http.StatusInternalServerError)
	}
}
