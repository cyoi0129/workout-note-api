package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// APIサーバーのヘルスチェック
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": 0, "data": "OK"})
}
