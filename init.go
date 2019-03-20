package main

import (
	"git.llsapp.com/geryon/dakuan/config"
	"github.com/douban-girls/backend/cfg"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB
var cache *redis.Client

func initDB() {
	dbConnect, err := sqlx.Open("postgres", cfg.CONFIG.DatabaseResourceStr)
	if err != nil {
		panic(err)
	}

	db = dbConnect
}

func initRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisHost,
		Password: "",
		DB:       config.RedisDB,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	cache = client
}
