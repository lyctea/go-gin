package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// User 用户结构体
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Response 通用响应结构体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ArticleResponse struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    []Article `json:"data"`
}

func main() {
	// 创建Gin路由引擎
	r := gin.Default()

	// 添加中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 健康检查接口
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, Response{
			Code:    200,
			Message: "服务运行正常",
			Data:    gin.H{"status": "ok"},
		})
	})

	// 获取用户列表接口
	r.GET("/users", func(c *gin.Context) {
		users := []User{
			{ID: 1, Name: "张三", Age: 25},
			{ID: 2, Name: "李四", Age: 30},
			{ID: 3, Name: "王五", Age: 28},
		}

		c.JSON(http.StatusOK, Response{
			Code:    200,
			Message: "获取用户列表成功",
			Data:    users,
		})
	})

	// 根据ID获取用户接口
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")

		// 模拟用户数据（这里可以根据id查询真实数据）
		user := User{ID: 1, Name: "张三", Age: 25}

		// 使用id参数（避免未使用变量警告）
		_ = id

		c.JSON(http.StatusOK, Response{
			Code:    200,
			Message: "获取用户信息成功",
			Data:    user,
		})
	})

	// 创建用户接口
	r.POST("/users", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, Response{
				Code:    400,
				Message: "请求参数错误",
				Data:    nil,
			})
			return
		}

		// 模拟创建用户
		user.ID = 4

		c.JSON(http.StatusCreated, Response{
			Code:    201,
			Message: "创建用户成功",
			Data:    user,
		})
	})

	r.GET("/articles", func(c *gin.Context) {
		articles := []Article{
			{ID: 1, Title: "标题1", Content: "内容1"},
			{ID: 2, Title: "标题2", Content: "内容2"},
			{ID: 3, Title: "标题3", Content: "内容3"},
		}

		c.JSON(http.StatusOK, Response{
			Code:    200,
			Message: "获取文章列表成功",
			Data:    articles,
		})
	})

	// 启动服务器
	r.Run(":8080")
}
