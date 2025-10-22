package routes

import (
	"go-gin/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRouter 配置路由
func SetupRouter() *gin.Engine {
	// 创建Gin路由引擎
	r := gin.Default()

	// 添加中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 健康检查接口
	r.GET("/health", handlers.HealthCheck)

	// 用户相关接口
	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetUserByID)
	r.POST("/users", handlers.CreateUser)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	// 文章相关接口
	r.GET("/articles", handlers.GetArticles)
	r.GET("/articles/:id", handlers.GetArticleByID)
	r.POST("/articles", handlers.CreateArticle)
	r.PUT("/articles/:id", handlers.UpdateArticle)
	r.DELETE("/articles/:id", handlers.DeleteArticle)

	return r
}
