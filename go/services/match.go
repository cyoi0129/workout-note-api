package services

import (
	"errors"
	"fmt"
	"workout-note-api/models"
)

// すでに存在するマッチングを取得
func FetchExistMatches(user_id int) ([]int, error) {
	var matches []int
	rows, err := models.DB.Query("SELECT requester, approver FROM \"workout_matches\" WHERE (requester = $1) OR approver = $1", user_id)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var requester, approver int
		rows.Scan(&requester, &approver)
		if requester != user_id {
			matches = append(matches, requester)
		}
		if approver != user_id {
			matches = append(matches, approver)
		}
	}
	return matches, nil
}

// 未処理のマッチングリクエストを取得
func FetchRequestingMatches(user_id int) ([]models.MatchData, error) {
	var matches []models.MatchData
	rows, err := models.DB.Query("SELECT id, requester, approver, status FROM \"workout_matches\" WHERE (approver = $1) AND status = $2", user_id, "REQUEST")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var match models.MatchData
		rows.Scan(&match.Id, &match.Requester, &match.Approver, &match.Status)
		person_row := models.DB.QueryRow("SELECT name, gender, brith, bp, sq, dl FROM \"workout_persons\" WHERE userID = $1", match.Requester)
		person_row.Scan(&match.Name, &match.Gender, &match.Brith, &match.Bp, &match.Sq, &match.Dl)
		matches = append(matches, match)
	}
	return matches, nil
}

// マッチングリクエストを作成
func CreateMatch(input models.Match) (models.Match, error) {
	match := models.Match{
		Requester: input.Requester,
		Approver:  input.Approver,
		Status:    "REQUEST",
	}
	// 交互のマッチングリクエストの存在有無をまずチェック
	var counter int
	models.DB.QueryRow("SELECT id FROM \"workout_matches\" WHERE ((requester = $1) AND (approver = $2)) OR ((requester = $2) OR (approver = $1))", input.Requester, input.Approver).Scan(&counter)
	if counter == 0 {
		err := models.DB.QueryRow("INSERT INTO workout_matches(requester, approver, status) VALUES($1,$2,$3) RETURNING id", match.Requester, match.Approver, match.Status).Scan(&match.Id)
		if err != nil {
			fmt.Println(err)
			return match, err
		}
	} else {
		err := errors.New("match exist")
		return match, err
	}
	return match, nil
}

// マッチングリクエストの更新
func UpdateMatch(input models.Match) (models.Match, error) {
	match := models.Match{
		Id:        input.Id,
		Requester: input.Requester,
		Approver:  input.Approver,
		Status:    input.Status,
	}
	_, err := models.DB.Query("UPDATE \"workout_matches\" SET status = $1 WHERE Id = $2", match.Status, match.Id)
	if err != nil {
		fmt.Println(err)
		return match, err
	}
	return match, nil
}
