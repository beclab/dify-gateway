package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var Database *gorm.DB = nil
var StatusActive = 0
var StatusDelete = 1

type AccountApp struct {
	gorm.Model
	AccountID    string `gorm:"column:account_id"`
	AccountEmail string `gorm:"column:account_email"`
	AppID        string `gorm:"column:app_id"`
	Status       int    `gorm:"column:status;default:0"`
}

func InitDB() {
	username := os.Getenv("PG_USERNAME") // 账号
	password := os.Getenv("PG_PASSWORD") // 密码
	host := os.Getenv("PG_HOST")         // 地址
	port := os.Getenv("PG_PORT")         // 端口
	DBname := os.Getenv("PG_DATABASE")   // 数据库名称

	// 连接字符串
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, DBname, port)

	// 连接到 PostgreSQL 数据库
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// 自动迁移模式（创建表）
	err = db.AutoMigrate(&AccountApp{})
	if err != nil {
		log.Fatal(err)
	}

	Database = db

	fmt.Println("Table created successfully!")
}
