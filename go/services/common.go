package services

import (
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

// 文字列の配列から数字配列へ変換
func convert2Int(strings []string) (ints []int) {
	var results []int
	for _, i := range strings {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		results = append(results, j)
	}
	return results
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
