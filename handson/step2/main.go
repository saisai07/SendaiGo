package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Conn コネクション
type Conn struct {
	Db *sql.DB
}

// グループIDを取得する
type Request struct {
	Group string
}

// データベースからの戻り
type Responce struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

// コネクション
var Db Conn

// エラー
var err error

func handler(w http.ResponseWriter, r *http.Request) {
	req := Request{
		Group: "1",
	}

	resp, err := db.findByGroup(req.Group)

	if err != nil {
		log.Println(err)
	}
	// JSONの生成
	res, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {
	route := http.NewServeMux()
	route.HandleFunc("/", handler)
	http.ListenAndServe(":8080", route)
}

// SQL実行
func (db Conn) findByGroup(group string) (responce []Responce, err error) {
	mess := Responce{}

	db, err = db.conn()
	defer db.Db.Close()

	rows, err := db.Db.Query("SELECT `name`, `message` FROM message WHERE `group` = ?", group)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		if err = rows.Scan(&mess.Name, &mess.Message); err != nil {
			log.Println(err)
		}
		responce = append(responce, mess)
	}
	return
}

// conn コネクションプールする、レシーバ
func (c Conn) conn() (db Conn, err error) {
	c.Db, err = sql.Open("mysql", "{ID}:{PASSWD}@tcp({HOST}:3306)/handson?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatal("db error.")
	}
	db = c
	return
}