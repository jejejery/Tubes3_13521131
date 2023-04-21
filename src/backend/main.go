package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type QnA struct {
	id       int    `json: "id"`
	question string `json:"question"`
	answer   string `json:"answer"`
}

type History struct {
	id_sesi int    `json: "id_sesi"`
	urutan  int    `json: "urutan"`
	teks    string `json:"text"`
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/chatty_loli")
	if err != nil {
		println("Trouble in connecting")
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		println("Can not do query")
	}
	results, err := db.Query(("SELECT question FROM qna"))
	if err != nil {
		println("Query tdk bekerja")
	}
	for results.Next() {
		var hasil QnA
		err = results.Scan(&hasil.question)
		if err != nil {
			println("doesnt work")
		}
		println(hasil.question)
	}
}
