package config

import (
	"os"
	"time"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	// "github.com/glebarez/sqlite" // Pure-Go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	// "github.com/libtnb/sqlite" // Pure-Go SQLite driver, checkout https://github.com/libtnb/sqlite for details
	"gorm.io/gorm"
)

var DB *gorm.DB

// github.com/mattn/go-sqlite3

// config 配置表模型
type Config struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	User  string `json:"modifyUser"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func Run() {
	// 确保 data 目录存在
	if err := os.MkdirAll("data", 0755); err != nil {
		panic("failed to create data directory: " + err.Error())
	}

	var err error
	DB, err = gorm.Open(sqlite.Open("data/config.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = DB.AutoMigrate(&Config{})
	if err != nil {
		panic("failed to migrate database")
	}
}

// TableName 设置表名
func (Config) TableName() string {
	return "config"
}
