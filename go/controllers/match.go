package controllers

import (
	"net/http"

	"strconv"
	"workout-note-api/models"
	"workout-note-api/services"

	"github.com/gin-gonic/gin"
)

// すでに存在するマッチングを取得
func FetchExistMatches(c *gin.Context) {
	user_param := c.Param("id")
	id, err := strconv.Atoi(user_param)
	if err != nil {
		id = 0
	}
	matches, err := services.FetchExistMatches(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": 1, "data": ""})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": 0, "data": matches})
	}
}

// 未処理のマッチングリクエストを取得
func FetchRequestingMatches(c *gin.Context) {
	user_param := c.Param("id")
	id, err := strconv.Atoi(user_param)
	if err != nil {
		id = 0
	}
	matches, err := services.FetchRequestingMatches(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": 1, "data": ""})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": 0, "data": matches})
	}
}

// マッチングリクエストを作成
func CreateMatch(c *gin.Context) {
	var input models.Match
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 1, "data": err.Error()})
		return
	}
	match, db_error := services.CreateMatch(input)

	if db_error != nil {
		c.JSON(http.StatusOK, gin.H{"status": 1, "data": ""})
	} else {
		notice := models.Notice{
			UserID: input.Approver,
			Type:   "REQUEST",
		}
		services.CreateNotice(notice) // リクエスト先の人に通知を作成
		c.JSON(http.StatusOK, gin.H{"status": 0, "data": match})
	}
}

// マッチングリクエストの更新
func UpdateMatch(c *gin.Context) {
	var input models.Match
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 1, "data": err.Error()})
		return
	}

	match, match_error := services.UpdateMatch(input)

	if match_error != nil {
		c.JSON(http.StatusOK, gin.H{"status": 1, "data": ""})
	} else {
		if input.Status == "APPROVAL" { // 承認された場合、チャットの作成
			chat_data := models.Chat{
				Member: []int{input.Requester, input.Approver},
			}
			chat, chat_error := services.CreateChat(chat_data)
			if chat_error != nil {
				c.JSON(http.StatusOK, gin.H{"status": 1, "data": ""})
			} else {
				notice := models.Notice{
					UserID: input.Requester,
					ChatID: int(chat.Id),
					Type:   "MESSAGE",
				}
				services.CreateNotice(notice) // リクエストされた人に通知を作成
			}
		}
		c.JSON(http.StatusOK, gin.H{"status": 0, "data": match})
	}
}
