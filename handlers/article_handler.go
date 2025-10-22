package handlers

import (
	"go-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetArticles 获取文章列表
func GetArticles(c *gin.Context) {
	articles := []models.Article{
		{ID: 1, Title: "标题1", Content: "内容1"},
		{ID: 2, Title: "标题2", Content: "内容2"},
		{ID: 3, Title: "标题3", Content: "内容3"},
	}

	c.JSON(http.StatusOK, models.Response{
		Code:    200,
		Message: "获取文章列表成功",
		Data:    articles,
	})
}
