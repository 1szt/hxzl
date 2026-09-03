package setup

import (
	"errors"
	"fmt"
	"mcup-server/config"

	"gorm.io/gorm"
)

func Run() {
	var cfg config.Config
	err := config.DB.Where("name = ?", "port").First(&cfg).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		cfg = config.Config{
			Name:  "port",
			Value: "9081",
		}
		err = config.DB.Create(&cfg).Error
		if err != nil {
			fmt.Printf("❌ [Setup] 初始化配置数据库失败 | %v\n", err)
		}
	} else if err != nil {
		fmt.Printf("❌ [Setup] 查询配置数据库失败 | %v\n", err)
	}
}
