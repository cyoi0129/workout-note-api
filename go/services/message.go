package services

import (
	"fmt"
	"workout-note-api/models"
)

// チャットのメッセージを取得
func FetchMessages(chat_id int) ([]models.Message, error) {
	var messages []models.Message
	rows, err := models.DB.Query("SELECT id, chatID, sender, receiver, content, date FROM \"workout_messages\" WHERE chatID = $1", chat_id)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var message models.Message
		rows.Scan(&message.Id, &message.ChatID, &message.Sender, &message.Receiver, &message.Content, &message.Date)
		messages = append(messages, message)
	}
	return messages, nil
}

// メッセージ作成
func CreateMessage(input models.Message) (models.Message, error) {
	message := models.Message{
		ChatID:   input.ChatID,
		Sender:   input.Sender,
		Receiver: input.Receiver,
		Content:  input.Content,
		Date:     input.Date,
	}
	err := models.DB.QueryRow("INSERT INTO workout_messages(chatID, sender, receiver, content, date) VALUES($1, $2, $3, $4, $5) RETURNING id", message.ChatID, message.Sender, message.Receiver, message.Content, message.Date).Scan(&message.Id)
	if err != nil {
		fmt.Println(err)
		return message, err
	}
	return message, nil
}
