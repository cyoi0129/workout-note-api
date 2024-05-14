package controllers

import (
	"fmt"
	"net/http"

	"strconv"
	"time"
	"workout-note/models"
	"workout-note/services"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const SECRET_KEY = "SECRET"

// ユーザー一覧の取得
func FetchUsers(c *gin.Context) {
	users, err := services.FetchUsers()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": 1, "data": ""})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": 0, "data": users})
	}
}

// 新たなユーザーの追加
func CreateUser(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 1, "data": err.Error()})
		return
	}

	user, db_error := services.CreateUser(input)
	if db_error != nil {
		c.JSON(http.StatusOK, gin.H{"status": 1, "data": ""})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": 0, "data": user})
	}
}

// 既存ユーザーの更新
func UpdateUser(c *gin.Context) {
	user_param := c.Param("id")
	id, err := strconv.Atoi(user_param)
	if err != nil {
		id = 0
	}
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 1, "data": err.Error()})
		return
	}

	user, db_error := services.UpdateUser(id, input)
	if db_error != nil {
		c.JSON(http.StatusOK, gin.H{"status": 1, "data": ""})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": 0, "data": user})
	}
}

// 既存ユーザーの削除
func DeleteUserById(c *gin.Context) {
	user_param := c.Param("id")
	target_user, err := strconv.Atoi(user_param)
	if err != nil {
		target_user = 0
	}
	user, db_error := services.DeleteUserById(target_user)
	if db_error != nil {
		c.JSON(http.StatusOK, gin.H{"status": 1, "data": ""})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": 0, "data": user})
	}
}

// ログインAPIがPOSTされたときのプロセス
func LoginHandler(c *gin.Context) {
	var inputUser models.LoginUser

	// リクエストからユーザー情報を取得
	if err := c.BindJSON(&inputUser); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": 1, "data": "Login Error"})
		return
	}

	loginResult, userName := services.CheckUserVaildation(inputUser.Email, inputUser.Password)
	// ユーザー情報の検証
	if !loginResult {
		c.JSON(http.StatusBadRequest, gin.H{"status": 1, "data": "Login Error"})
		return
	}

	// トークンの発行（ヘッダー・ペイロード）
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": inputUser.Email,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 1, "data": "Login Error"})
		return
	}

	// ヘッダーにトークンをセット
	c.Header("Authorization", tokenString)
	c.JSON(http.StatusOK, gin.H{"status": 0, "token": tokenString, "user": userName})
}

// APIのトークンからの認証チェック
func AuthMiddleware(c *gin.Context) {
	// Authorizationヘッダーからトークンを取得
	tokenString := c.GetHeader("Authorization")

	// トークンの検証
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"status": 1, "data": "Auth Error"})
		c.Abort()
		return
	}

	c.Next()
}
