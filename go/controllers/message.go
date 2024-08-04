package controllers

import (
	"net/http"

	"strconv"
	"workout-note-api/models"
	"workout-note-api/services"

	"github.com/gin-gonic/gin"
)

// チャットのメッセージを取得
func FetchMessages(c *gin.Context) {
	chat_param := c.Param("id")
	id, err := strconv.Atoi(chat_param)
	if err != nil {
		id = 0
	}
	messages, err := services.FetchMessages(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": 1, "data": ""})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": 0, "data": messages})
	}
}

// メッセージ作成
func CreateMessage(c *gin.Context) {
	var input models.Message

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 1, "data": err.Error()})
		return
	}
	notice := models.Notice{
		UserID: input.Receiver,
		ChatID: input.ChatID,
		Type:   "MESSAGE",
	}
	message, db_error := services.CreateMessage(input)
	if db_error != nil {
		c.JSON(http.StatusOK, gin.H{"status": 1, "data": ""})
	} else {
		services.CreateNotice(notice) // 受信者に通知を作成
		c.JSON(http.StatusOK, gin.H{"status": 0, "data": message})
	}
}
