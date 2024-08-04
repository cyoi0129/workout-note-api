package controllers

import (
	"net/http"

	"strconv"
	"workout-note-api/services"

	"github.com/gin-gonic/gin"
)

// 未読の通知を取得
func FetchNotices(c *gin.Context) {
	user_param := c.Param("id")
	id, err := strconv.Atoi(user_param)
	if err != nil {
		id = 0
	}
	notices, err := services.FetchNotices(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": 1, "data": ""})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": 0, "data": notices})
	}
}

// マッチ通知を削除
func DeleteMatchNotice(c *gin.Context) {
	notice_param := c.Param("id")
	target_user, err := strconv.Atoi(notice_param)
	if err != nil {
		target_user = 0
	}
	id, db_error := services.DeleteMatchNotice(target_user)
	if db_error != nil {
		c.JSON(http.StatusOK, gin.H{"status": 1, "data": ""})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": 0, "data": id})
	}
}

// メッセージ通知を削除
func DeleteMessageNotice(c *gin.Context) {
	notice_param := c.Param("id")
	target_id, err := strconv.Atoi(notice_param)
	if err != nil {
		target_id = 0
	}
	id, db_error := services.DeleteMessageNotice(target_id)
	if db_error != nil {
		c.JSON(http.StatusOK, gin.H{"status": 1, "data": ""})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": 0, "data": id})
	}
}
