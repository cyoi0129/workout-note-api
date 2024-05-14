package models

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDatabase() {
	env := os.Getenv("ENV") // 環境変数を取得
	if env == "production" {
		fmt.Println("本番DB環境接続")
		godotenv.Load(".env.production")
	} else {
		fmt.Println("開発DB環境接続")
		godotenv.Load(".env.development")
	}
	dbName := os.Getenv("DBNAME")
	dbUser := os.Getenv("DBUSER")
	dbHost := os.Getenv("DBHOST")
	dbPassword := os.Getenv("DBPASSWORD")
	dbPort := 5432

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	fmt.Println("DB接続")
	DB = db
}

// func dbTestSelect(db *sql.DB) {
// 	row := db.QueryRow("SELECT id, name, muscles FROM \"masters\" WHERE id = $1", 1)
// 	var selectedID int
// 	var selectedName string
// 	var muscles []int32
// 	err := row.Scan(&selectedID, &selectedName, pq.Array(&muscles))
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			fmt.Println("指定されたIDのレコードが見つかりませんでした。")
// 		} else {
// 			log.Fatal(err)
// 		}
// 		return
// 	}

// 	for i, v := range muscles {
// 		fmt.Println(i, v)
// 	}
// 	fmt.Printf("ID: %d, 名前: %s\n", selectedID, selectedName)
// }
