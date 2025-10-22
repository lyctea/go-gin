package handlers

import (
	"go-gin/config"
	"go-gin/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetArticles 获取文章列表
func GetArticles(c *gin.Context) {
	var articles []models.Article

	// 从数据库查询所有文章，并预加载关联的用户信息
	if err := config.GetDB().Preload("User").Find(&articles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code:    500,
			Message: "查询文章列表失败",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Code:    200,
		Message: "获取文章列表成功",
		Data:    articles,
	})
}

// GetArticleByID 根据ID获取文章
func GetArticleByID(c *gin.Context) {
	id := c.Param("id")

	// 将字符串ID转换为uint
	articleID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    400,
			Message: "无效的文章ID",
			Data:    nil,
		})
		return
	}

	var article models.Article
	// 从数据库查询指定文章，并预加载关联的用户信息
	if err := config.GetDB().Preload("User").First(&article, articleID).Error; err != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Code:    404,
			Message: "文章不存在",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Code:    200,
		Message: "获取文章信息成功",
		Data:    article,
	})
}

// CreateArticle 创建文章
func CreateArticle(c *gin.Context) {
	var article models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    400,
			Message: "请求参数错误: " + err.Error(),
			Data:    nil,
		})
		return
	}

	// 检查用户是否存在
	var user models.User
	if err := config.GetDB().First(&user, article.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    400,
			Message: "指定的用户不存在",
			Data:    nil,
		})
		return
	}

	// 保存文章到数据库
	if err := config.GetDB().Create(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code:    500,
			Message: "创建文章失败: " + err.Error(),
			Data:    nil,
		})
		return
	}

	// 重新加载文章以包含用户信息
	config.GetDB().Preload("User").First(&article, article.ID)

	c.JSON(http.StatusCreated, models.Response{
		Code:    201,
		Message: "创建文章成功",
		Data:    article,
	})
}

// UpdateArticle 更新文章
func UpdateArticle(c *gin.Context) {
	id := c.Param("id")

	// 将字符串ID转换为uint
	articleID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    400,
			Message: "无效的文章ID",
			Data:    nil,
		})
		return
	}

	var article models.Article
	// 检查文章是否存在
	if err := config.GetDB().First(&article, articleID).Error; err != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Code:    404,
			Message: "文章不存在",
			Data:    nil,
		})
		return
	}

	// 绑定更新的数据
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    400,
			Message: "请求参数错误: " + err.Error(),
			Data:    nil,
		})
		return
	}

	// 保持原有ID
	article.ID = uint(articleID)

	// 更新文章信息
	if err := config.GetDB().Save(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code:    500,
			Message: "更新文章失败: " + err.Error(),
			Data:    nil,
		})
		return
	}

	// 重新加载文章以包含用户信息
	config.GetDB().Preload("User").First(&article, article.ID)

	c.JSON(http.StatusOK, models.Response{
		Code:    200,
		Message: "更新文章成功",
		Data:    article,
	})
}

// DeleteArticle 删除文章（软删除）
func DeleteArticle(c *gin.Context) {
	id := c.Param("id")

	// 将字符串ID转换为uint
	articleID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    400,
			Message: "无效的文章ID",
			Data:    nil,
		})
		return
	}

	// 软删除文章
	if err := config.GetDB().Delete(&models.Article{}, articleID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code:    500,
			Message: "删除文章失败: " + err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Code:    200,
		Message: "删除文章成功",
		Data:    nil,
	})
}
