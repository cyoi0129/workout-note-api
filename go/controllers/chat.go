package controllers

import (
	"net/http"

	"strconv"
	"workout-note-api/services"

	"github.com/gin-gonic/gin"
)

// チャット一覧を取得
func FetchChats(c *gin.Context) {
	user_param := c.Param("id")
	id, err := strconv.Atoi(user_param)
	if err != nil {
		id = 0
	}
	chats, err := services.FetchChats(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": 1, "data": ""})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": 0, "data": chats})
	}
}
