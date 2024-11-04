package services

import (
	"fmt"
	"workout-note-api/models"
)

// 未読の通知を取得
func FetchNotices(user_id int) ([]models.Notice, error) {
	var notices []models.Notice
	rows, err := models.DB.Query("SELECT id, userID, chatID, type FROM \"workout_notices\" WHERE userID = $1", user_id)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var notice models.Notice
		rows.Scan(&notice.Id, &notice.UserID, &notice.ChatID, &notice.Type)
		notices = append(notices, notice)
	}
	return notices, nil
}

// 通知を作成
func CreateNotice(input models.Notice) (models.Notice, error) {
	notice := models.Notice{
		UserID: input.UserID,
		ChatID: input.ChatID,
		Type:   input.Type,
	}

	err := models.DB.QueryRow("INSERT INTO workout_notices(userID, type, chatID) VALUES($1, $2, $3) RETURNING id", notice.UserID, notice.Type, notice.ChatID).Scan(&notice.Id)
	if err != nil {
		fmt.Println(err)
		return notice, err
	}
	return notice, nil
}

// 通知を削除
func DeleteMatchNotice(user_id int) (int, error) {
	var err error
	_, err = models.DB.Query("DELETE FROM \"workout_notices\" WHERE userID = $1", user_id)
	if err != nil {
		return user_id, err
	}
	return user_id, nil
}

func DeleteMessageNotice(target_id int) (int, error) {
	var err error
	_, err = models.DB.Query("DELETE FROM \"workout_notices\" WHERE chatID = $1", target_id)
	if err != nil {
		return target_id, err
	}
	return target_id, nil
}
