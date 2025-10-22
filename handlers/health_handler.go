package handlers

import (
	"go-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck 健康检查接口
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, models.Response{
		Code:    200,
		Message: "服务运行正常",
		Data:    gin.H{"status": "ok"},
	})
}
