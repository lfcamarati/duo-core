package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var (
	Db *sql.DB
)

// https://go.dev/play/p/FAiGbqeJG0H
// https://go.dev/doc/database/open-handle
func Init() {
	var host = viper.GetString("database.duo.host")
	var user = viper.GetString("database.duo.user")
	var pass = viper.GetString("database.duo.pass")

	var err error
	Db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/duo", user, pass, host))

	if err != nil {
		log.Fatal("Invalid DB config:", err)
	}

	if err = Db.Ping(); err != nil {
		log.Fatal("DB unreachable:", err)
	}
}
