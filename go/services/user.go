package services

import (
	"fmt"
	"workout-note-api/models"

	"github.com/lib/pq"
)

// ユーザー一覧の取得
func FetchUsers() ([]models.User, error) {
	var users []models.User
	rows, err := models.DB.Query("SELECT id, email FROM \"users\"")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var user models.User
		rows.Scan(&user.Id, &user.Email)
		users = append(users, user)
	}
	return users, nil
}

// ユーザー追加
func CreateUser(input models.User) (models.User, error) {
	var person models.Person
	var err error
	user := models.User{
		Email:    input.Email,
		Password: input.Password,
	}
	err = models.DB.QueryRow("INSERT INTO users(email, password) VALUES($1,$2) RETURNING id", user.Email, user.Password).Scan(&user.Id)
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	err = models.DB.QueryRow("INSERT INTO persons(userID, name, gender, brith, bp, sq, dl) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id", user.Id, "", "", 1980, 0, 0, 0).Scan(&person.Id)

	if err != nil {
		fmt.Println(err)
		return user, err
	}

	return user, nil
}

// 既ユーザーの更新
func UpdateUser(user_id int, input models.User) (models.User, error) {
	user := models.User{
		Id:       input.Id,
		Email:    input.Email,
		Password: input.Password,
	}
	_, err := models.DB.Query("UPDATE \"users\" SET email = $1, password = $2 WHERE id = $3", user.Email, user.Password, user_id)
	if err != nil {
		return user, err
	}
	return user, nil
}

// 既ユーザーの削除
func DeleteUserById(user_id int) (int, error) {
	var err error
	_, err = models.DB.Query("DELETE FROM \"users\" WHERE id = $1", user_id)
	if err != nil {
		return user_id, err
	}
	_, err = models.DB.Query("DELETE FROM \"persons\" WHERE userID = $1", user_id)
	if err != nil {
		return user_id, err
	}
	return user_id, nil
}

// ユーザーログイン検証
func CheckUserVaildation(email string, password string) (bool, models.Person) {
	var user models.User
	var person models.Person
	row := models.DB.QueryRow("SELECT id, email, password FROM \"users\" WHERE email = $1", email)
	err := row.Scan(&user.Id, &user.Email, &user.Password)

	if err != nil || user.Password != password {
		fmt.Println(err)
		return false, person
	}

	person_row := models.DB.QueryRow("SELECT id, userID, name, gender, brith, stations, areas, gyms, times, bp, sq, dl FROM \"persons\" WHERE userID = $1", user.Id)
	var stationStr, areaStr, gymStr []string
	err = person_row.Scan(&person.Id, &person.UserID, &person.Name, &person.Gender, &person.Brith, pq.Array(&stationStr), pq.Array(&areaStr), pq.Array(&gymStr), pq.Array(&person.Times), &person.Bp, &person.Sq, &person.Dl)
	person.Stations = convert2Int(stationStr)
	person.Areas = convert2Int(areaStr)
	person.Gyms = convert2Int(gymStr)

	if err != nil {
		fmt.Println(err)
		return false, person
	}

	return true, person
}
