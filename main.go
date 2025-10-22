package main

import (
	"go-gin/routes"
)

func main() {
	// 配置路由
	r := routes.SetupRouter()

	// 启动服务器
	r.Run(":8080")
}
