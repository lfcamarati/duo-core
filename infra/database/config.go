package database

import (
	"database/sql"
	"log"
)

var (
	Db *sql.DB
)

// https://go.dev/play/p/FAiGbqeJG0H
// https://go.dev/doc/database/open-handle
func Init() {
	var err error
	Db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/duo")

	if err != nil {
		log.Fatal("Invalid DB config:", err)
	}

	if err = Db.Ping(); err != nil {
		log.Fatal("DB unreachable:", err)
	}
}
