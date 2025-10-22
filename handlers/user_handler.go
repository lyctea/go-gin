package handlers

import (
	"go-gin/config"
	"go-gin/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUsers 获取用户列表
func GetUsers(c *gin.Context) {
	var users []models.User

	// 从数据库查询所有用户
	if err := config.GetDB().Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code:    500,
			Message: "查询用户列表失败",
			Data:    nil,
		})
		return
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

	// 将字符串ID转换为uint
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    400,
			Message: "无效的用户ID",
			Data:    nil,
		})
		return
	}

	var user models.User
	// 从数据库查询指定用户
	if err := config.GetDB().First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Code:    404,
			Message: "用户不存在",
			Data:    nil,
		})
		return
	}

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
			Message: "请求参数错误: " + err.Error(),
			Data:    nil,
		})
		return
	}

	// 保存用户到数据库
	if err := config.GetDB().Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code:    500,
			Message: "创建用户失败: " + err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, models.Response{
		Code:    201,
		Message: "创建用户成功",
		Data:    user,
	})
}

// UpdateUser 更新用户
func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	// 将字符串ID转换为uint
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    400,
			Message: "无效的用户ID",
			Data:    nil,
		})
		return
	}

	var user models.User
	// 检查用户是否存在
	if err := config.GetDB().First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Code:    404,
			Message: "用户不存在",
			Data:    nil,
		})
		return
	}

	// 绑定更新的数据
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    400,
			Message: "请求参数错误: " + err.Error(),
			Data:    nil,
		})
		return
	}

	// 保持原有ID
	user.ID = uint(userID)

	// 更新用户信息
	if err := config.GetDB().Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code:    500,
			Message: "更新用户失败: " + err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Code:    200,
		Message: "更新用户成功",
		Data:    user,
	})
}

// DeleteUser 删除用户（软删除）
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	// 将字符串ID转换为uint
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    400,
			Message: "无效的用户ID",
			Data:    nil,
		})
		return
	}

	// 软删除用户
	if err := config.GetDB().Delete(&models.User{}, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code:    500,
			Message: "删除用户失败: " + err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Code:    200,
		Message: "删除用户成功",
		Data:    nil,
	})
}
