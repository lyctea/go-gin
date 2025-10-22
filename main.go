package main

import (
	"go-gin/config"
	"go-gin/routes"
	"log"
)

func main() {
	// 初始化数据库
	if err := config.InitDB(); err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}
	defer config.CloseDB()

	// 自动迁移数据库表
	if err := config.AutoMigrate(); err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	// 配置路由
	r := routes.SetupRouter()

	// 启动服务器
	log.Println("服务器启动在 http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
