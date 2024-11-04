package services

import (
	"fmt"
	"workout-note-api/models"

	"github.com/lib/pq"
)

// チャット一覧を取得
func FetchChats(user_id int) ([]models.ChatData, error) {
	var chats []models.ChatData
	rows, err := models.DB.Query("SELECT id, member FROM \"workout_chats\" WHERE member && ($1)", pq.Array([]int{user_id}))
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var chat models.ChatData
		var memberStr []string
		rows.Scan(&chat.Id, pq.Array(&memberStr))
		memberInt := convert2Int(memberStr)

		for _, v := range []int(memberInt) {
			if v != user_id {
				chat.TargetId = v
			}
		}
		person_row := models.DB.QueryRow("SELECT name FROM \"workout_persons\" WHERE userID = $1", chat.TargetId)
		person_row.Scan(&chat.TargetName)
		message_row := models.DB.QueryRow("SELECT content, date FROM \"workout_messages\"WHERE chatID = $1 ORDER BY id DESC", chat.Id)
		message_row.Scan(&chat.Message, &chat.Date)
		chats = append(chats, chat)
	}
	return chats, nil
}

// チャット作成
func CreateChat(input models.Chat) (models.Chat, error) {
	chat := models.Chat{
		Member: input.Member,
	}
	err := models.DB.QueryRow("INSERT INTO workout_chats(member) VALUES($1) RETURNING id", pq.Array(chat.Member)).Scan(&chat.Id)
	if err != nil {
		fmt.Println(err)
		return chat, err
	}
	return chat, nil
}
