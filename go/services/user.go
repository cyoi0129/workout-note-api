package services

import (
	"fmt"
	"workout-note/models"

	"golang.org/x/crypto/bcrypt"
)

// ユーザー一覧の取得
func FetchUsers() ([]models.User, error) {
	var users []models.User
	// rows, err := models.DB.Query("SELECT * FROM users")
	rows, err := models.DB.Query("SELECT id, name, email FROM \"users\"")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var user models.User
		rows.Scan(&user.Id, &user.Name, &user.Email)
		users = append(users, user)
	}
	fmt.Printf("%v", users)
	return users, nil
}

// ユーザー追加
func CreateUser(input models.User) (models.User, error) {
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	err := models.DB.QueryRow("INSERT INTO users(name, email, password) VALUES($1,$2,$3) RETURNING id", user.Name, user.Email, user.Password).Scan(&user.Id)

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
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	_, err := models.DB.Query("UPDATE \"users\" SET name = $1, email = $2, password = $3 WHERE id = $4", user.Name, user.Email, user.Password, user_id)
	if err != nil {
		return user, err
	}
	return user, nil
}

// 既ユーザーの削除
func DeleteUserById(user_id int) (int, error) {
	_, err := models.DB.Query("DELETE FROM \"users\" WHERE id = $1", user_id)
	if err != nil {
		return user_id, err
	}
	return user_id, nil
}

// ユーザーログイン検証
func CheckUserVaildation(email string, password string) (bool, string) {
	var user models.User
	row := models.DB.QueryRow("SELECT name, password FROM \"users\" WHERE email = $1", email)
	err := row.Scan(&user.Name, &user.Password)
	if err != nil {
		fmt.Println(err)
		return false, ""
	}
	row.Scan(&user.Name, &user.Password)
	if user.Password != password {
		return false, ""
	}
	return true, user.Name
}

// 暗号(Hash)化
func PasswordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// 暗号(Hash)と入力された平パスワードの比較
func CompareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// INSERT INTO students (id, name, age) VALUES ('0001', '川端奈緒', 11);
// UPDATE students SET name = '多田典子' WHERE id = '0001'
// DELETE FROM students WHERE id = '0001'
