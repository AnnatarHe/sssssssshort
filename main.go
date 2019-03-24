package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/AnnatarHe/sssssssshort/src"

	_ "github.com/go-sql-driver/mysql"
)

type shortModel struct {
	ID  int64
	Src string
}

type payload struct {
	URL string `json:"url"`
}

func getRealLink(id int64) string {
	return src.HOST_URL + src.Encode(id)
}

func setResponse(writer http.ResponseWriter, statusCode int, url string) {
	resp := payload{URL: url}

	buf, err := json.Marshal(resp)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("error"))
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	writer.Write(buf)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	pass, _ := src.IPFilter(r)

	if !pass {
		setResponse(w, http.StatusTooManyRequests, "")
		return
	}

	// TODO: 可以加域名白名单限制呢

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	var req payload
	if err := json.Unmarshal(bodyBytes, &req); err != nil {
		setResponse(w, http.StatusBadRequest, "")
		return
	}

	if _, err := url.ParseRequestURI(req.URL); err != nil {
		setResponse(w, http.StatusBadRequest, "")
		return
	}

	var model shortModel

	expection := src.DB.Get(&model, "SELECT * FROM `short_url` WHERE `src`=? LIMIT 1", req.URL)

	if expection == nil {
		setResponse(w, http.StatusOK, getRealLink(model.ID))
		return
	}

	if expection != nil && expection != sql.ErrNoRows {
		log.Fatal(expection)
		setResponse(w, http.StatusInternalServerError, "")
		return
	}

	result, err := src.DB.Exec("INSERT INTO short_url(src) VALUES(?)", req.URL)
	if err != nil {
		log.Fatal(err)
		setResponse(w, http.StatusInternalServerError, "")
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		setResponse(w, http.StatusInternalServerError, "")
		return
	}

	setResponse(w, http.StatusOK, getRealLink(id))
	return
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	key := strings.Split(r.URL.Path, "/")[1]
	if key == "" {
		setResponse(w, http.StatusBadRequest, "")
		return
	}

	id, err := src.Decode(key)
	if err != nil {
		setResponse(w, http.StatusBadRequest, "")
		return
	}

	log.Println(id)

	var model shortModel

	expection := src.DB.Get(&model, "SELECT * FROM `short_url` WHERE `id`=? LIMIT 1", id)
	if expection != nil {
		log.Println(expection)
		setResponse(w, http.StatusNotFound, "")
		return
	}

	http.Redirect(w, r, model.Src, 302)
}

func init() {
	src.InitDB()
}

func dispatcher(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		createHandler(w, r)
		return
	}

	if r.URL.Path == "/" {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<h1>文档请访问： <a href='https://github.com/AnnatarHe/sssssssshort/tree/master/docs/api.md'>docs</a></h1>"))
		return
	}

	// 读操作
	readHandler(w, r)
	return
}

func main() {
	http.HandleFunc("/", dispatcher)

	if err := http.ListenAndServe(src.API_LISTEN, nil); err != nil {
		panic(err)
	}
}
