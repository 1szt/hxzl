package ginboot

// 底座初始化（gin 版）：武装全局中间件
// 必须在任何模块挂载之前调用，
// 以确保中间件在请求到达前就绪

import (
"sync"
"time"

web "1szt/gin"

g "github.com/gin-gonic/gin"
)

// Run 武装底座：设置全局中间件
func Run() {

// 捕获所有内部 panic
web.Engine.Use(g.Recovery())
// 记录日志
web.Engine.Use(g.Logger())
// 限制每个IP每分钟最多100个请求
web.Engine.Use(rateLimitByIP(100, time.Minute))

}

// ---------- 简单固定窗口 IP 限流中间件 ----------

// ipBucket 记录某个 IP 在当前窗口内的请求计数
type ipBucket struct {
windowStart time.Time
count       int
}

var (
rateMu   sync.Mutex
buckets  = make(map[string]*ipBucket)
)

// rateLimitByIP 限制每个 IP 在 window 时间内最多发 limit 个请求
func rateLimitByIP(limit int, window time.Duration) g.HandlerFunc {
return func(c *g.Context) {
ip := c.ClientIP()
now := time.Now()

rateMu.Lock()
b, ok := buckets[ip]
if !ok || now.Sub(b.windowStart) >= window {
// 新窗口开始
b = &ipBucket{windowStart: now}
buckets[ip] = b
}
if b.count >= limit {
rateMu.Unlock()
c.AbortWithStatusJSON(429, g.H{"error": "请求过于频繁，请稍后再试"})
return
}
b.count++
rateMu.Unlock()

c.Next()
}
}