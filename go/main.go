package main

import (
	"os"
	"time"
	"workout-note-api/controllers"
	"workout-note-api/models"
	"workout-note-api/websocket"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	os.Setenv("TZ", "UTC") // タイムゾーンの設定
	r := gin.Default()

	r.Use(cors.New(cors.Config{

		AllowOrigins: []string{ // アクセスを許可したいアクセス元
			"http://localhost:5173",
			"https://workout-note.netlify.app",
		},

		AllowMethods: []string{ // アクセスを許可したいHTTPメソッド
			"POST",
			"GET",
			"PUT",
			"DELETE",
		},

		AllowHeaders: []string{ // 許可したいHTTPリクエストヘッダ
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},

		AllowCredentials: true, // cookieなどの情報を必要とするかどうか

		MaxAge: 24 * time.Hour, // preflightリクエストの結果をキャッシュする時間
	}))
	models.ConnectDatabase() // DB初期化
	models.CreateCache()     // キャッシュ初期化

	r.StaticFile("/favicon.ico", "./resources/favicon.ico")
	api := r.Group("/api")
	api.POST("/login", controllers.LoginHandler)
	api.GET("/masters", controllers.FetchMasterList)
	api.POST("/user", controllers.CreateUser)
	api.Use(controllers.AuthMiddleware)
	{
		api.GET("/users", controllers.FetchUsers)
		api.GET("/person/:id", controllers.FetchPersonByID)
		api.POST("/person/:id", controllers.UpdatePerson)
		api.POST("/persons", controllers.FetchTargetPersons)
		api.PUT("/user/:id", controllers.UpdateUser)
		api.DELETE("/user/:id", controllers.DeleteUserById)
		api.GET("/exists/:id", controllers.FetchExistMatches)
		api.GET("/match/:id", controllers.FetchRequestingMatches)
		api.POST("/match", controllers.CreateMatch)
		api.PUT("/match/:id", controllers.UpdateMatch)
		api.GET("/notice/:id", controllers.FetchNotices)
		api.DELETE("/match_notice/:id", controllers.DeleteMatchNotice)
		api.DELETE("/message_notice/:id", controllers.DeleteMessageNotice)
		api.GET("/chats/:id", controllers.FetchChats)
		api.GET("/messages/:id", controllers.FetchMessages)
		api.POST("/message", controllers.CreateMessage)
		r.GET("/ws/:topic", websocket.ServeWs)
	}
	r.Run()
}
