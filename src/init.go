package src

import (
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB() {
	dbConnect, err := sqlx.Open("mysql", DB_SRC)
	if err != nil {
		panic(err)
	}

	DB = dbConnect
}
