package main

import (
	"net/http"
	"strings"

	_ "github.com/lib/pq"
)

func shortHandler(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func init() {
	initDB()
	initRedis()
}

func main() {

	http.HandleFunc("/short", shortHandler)

	if err := http.ListenAndServe(":9999"); err != nil {
		panic(err)
	}
}
