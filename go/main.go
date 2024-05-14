package main

import (
	"os"
	"time"

	controllers "workout-note/controls"
	"workout-note/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	os.Setenv("TZ", "UTC") // タイムゾーンの設定
	r := gin.Default()

	r.Use(cors.New(cors.Config{

		AllowOrigins: []string{ // アクセスを許可したいアクセス元
			"http://localhost:5173",
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
	api.Use(controllers.AuthMiddleware)
	{
		api.GET("/users", controllers.FetchUsers)
		api.POST("/user", controllers.CreateUser)
		api.PUT("/user/:id", controllers.UpdateUser)
		api.DELETE("/user/:id", controllers.DeleteUserById)
		api.GET("/masters/:id", controllers.FetchMasterList)
		api.POST("/master", controllers.CreateMaster)
		api.PUT("/master/:id", controllers.UpdateMaster)
		api.DELETE("/master/:id", controllers.DeleteMasterById)
		api.GET("/types", controllers.FetchTypeList)
		api.GET("/muscles", controllers.FetchMuscleList)
	}
	r.Run() // listen and serve on 0.0.0.0:8080
}
