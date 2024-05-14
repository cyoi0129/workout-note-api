package models

import (
	"fmt"
	"time"

	"context"
	"encoding/json"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var go_cache *redis.Client
var ctx = context.Background()

func CreateCache() {
	env := os.Getenv("ENV") // 環境変数を取得
	if env == "production" {
		fmt.Println("本番Cache環境接続")
		godotenv.Load(".env.production")
	} else {
		fmt.Println("開発Cache環境接続")
		godotenv.Load(".env.development")
	}

	go_cache = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("CACHE"), // redisの接続先
		Password: "",                 // パスワードなし
		DB:       0,                  // デフォルトDBを使う
		PoolSize: 1000,               // サイズの上限
	})
	if err := go_cache.Ping(ctx).Err(); err != nil {
		panic(err)
	}
	go_cache.FlushAll(ctx)
}

func SetCache(k string, v interface{}) (interface{}, error) { // キャッシュを保存
	data, err := json.Marshal(v) // オブジェクト式をラッパーしたjson形式へ変換
	if err != nil {
		return nil, err
	}
	result, err := go_cache.Set(ctx, k, data, 0).Result()
	go_cache.Expire(ctx, k, 30*time.Second) // 保存期間を指定
	return result, err
}

func GetCache(k string) (interface{}, error) { // キャッシュを取得
	var data interface{}
	result, err := go_cache.Get(ctx, k).Result()
	json.Unmarshal([]byte(result), &data) // キャッシュに保存されたjson形式へまたオブジェクトへ変換
	return data, err
}

func DeleteCache(k string) (interface{}, error) { // キャッシュを削除
	result, err := go_cache.Del(ctx, k).Result()
	return result, err
}
