package handlers

import (
	"go-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers 获取用户列表
func GetUsers(c *gin.Context) {
	users := []models.User{
		{ID: 1, Name: "张三", Age: 25},
		{ID: 2, Name: "李四", Age: 30},
		{ID: 3, Name: "王五", Age: 28},
	}

	c.JSON(http.StatusOK, models.Response{
		Code:    200,
		Message: "获取用户列表成功",
		Data:    users,
	})
}

// GetUserByID 根据ID获取用户
func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	// 模拟用户数据（这里可以根据id查询真实数据）
	user := models.User{ID: 1, Name: "张三", Age: 25}

	// 使用id参数（避免未使用变量警告）
	_ = id

	c.JSON(http.StatusOK, models.Response{
		Code:    200,
		Message: "获取用户信息成功",
		Data:    user,
	})
}

// CreateUser 创建用户
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    400,
			Message: "请求参数错误",
			Data:    nil,
		})
		return
	}

	// 模拟创建用户
	user.ID = 4

	c.JSON(http.StatusCreated, models.Response{
		Code:    201,
		Message: "创建用户成功",
		Data:    user,
	})
}
