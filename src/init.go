package src

import (
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB
var Cache *redis.Client

func InitDB() {
	dbConnect, err := sqlx.Open("mysql", DB_SRC)
	if err != nil {
		panic(err)
	}

	DB = dbConnect
}

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     REDIS_HOST,
		Password: "",
		DB:       REDIS_DB,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	Cache = client
}
