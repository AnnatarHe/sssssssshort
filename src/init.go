package src

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

func InitDB() {
	dbConnect, err := sqlx.Open("sqlite3", DB_SRC)
	if err != nil {
		panic(err)
	}

	DB = dbConnect
}
