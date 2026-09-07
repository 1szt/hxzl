package setup

import (
	"errors"
	"fmt"
	"mcup-server/config"
	"os"
	"os/exec"
	"time"

	g "github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Run() {
	var cfg config.Config
	var err error

	// 查询 setup 配置
	err = config.DB.Where("name = ?", "setup").First(&cfg).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 第一次运行，插入默认配置
		cfg = config.Config{
			Name:  "setup",
			Value: "flash",
		}
		err = config.DB.Create(&cfg).Error
		if err != nil {
			fmt.Printf("❌ [Setup] 初始化配置失败 | %v\n", err)
			return
		}

		cfg = config.Config{
			Name:  "port",
			Value: "9081",
		}
		err = config.DB.Create(&cfg).Error
		if err != nil {
			fmt.Printf("❌ [Setup] 初始化配置失败 | %v\n", err)
			return
		}
	} else if err != nil {
		fmt.Printf("❌ [Setup] 查询配置数据库失败 | %v\n", err)
		return
	}

	// 判断状态
	err = config.DB.Where("name = ?", "setup").First(&cfg).Error
	if err != nil {
		fmt.Printf("❌ [Setup] 查询配置数据库失败 | %v\n", err)
		return
	}
	if cfg.Value == "flash" {
		fmt.Println("⚡ [Setup] 初始化 web 服务")

		r := g.Default()
		r.LoadHTMLGlob("template/**/*.html")
		r.StaticFile("/favicon.ico", "./favicon.ico")
		r.Static("/static", "./static")

		r.GET("/", func(c *g.Context) { c.Redirect(302, "/setup") })
		r.GET("/setup", func(c *g.Context) { c.HTML(200, "setup", nil) })
		r.POST("/setup", func(c *g.Context) {
			// 更新为 done
			if err := config.DB.Model(&config.Config{}).
				Where("name = ?", "setup").
				Update("value", "done").Error; err != nil {
				c.String(500, fmt.Sprintf("❌ 保存配置失败 | %v", err))
				return
			}

			// 自我重启
			go func() {
				time.Sleep(1 * time.Second)
				restartSelf()
			}()
			c.String(200, "配置完成，正在重启...")
		})

		err := r.Run(":9081")
		if err != nil {
			fmt.Printf("❌ [Setup] 初始化服务启动失败 | %v\n", err)
		}
	}
}

// 自我重启函数
func restartSelf() {
	path, err := os.Executable()
	if err != nil {
		fmt.Printf("❌ 获取可执行文件路径失败 | %v\n", err)
		return
	}
	cmd := exec.Command(path, os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Start(); err != nil {
		fmt.Printf("❌ 重启进程失败 | %v\n", err)
		return
	}
	os.Exit(0)
}
