package boot

// 底座初始化：武装全局中间件
// 必须在任何模块挂载（Mount）之前调用，
// 以确保中间件在请求到达前就绪

import (
"time"

"1szt/chi"

"github.com/go-chi/chi/v5/middleware"
"github.com/go-chi/httprate"
)

// Run 武装底座：设置全局中间件
func Run() {

// 捕获所有内部 panic
chi.Mux.Use(middleware.Recoverer)
// 记录日志
chi.Mux.Use(middleware.Logger)
// 限制每个IP每分钟最多100个请求
chi.Mux.Use(httprate.LimitByIP(100, 1*time.Minute))

}