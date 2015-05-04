package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os/user"
)

var connection *sql.DB = nil

func GetConnection(option string) *sql.DB {
	var err error

	if connection == nil {
		if option == "" {
			u, _ := user.Current()
			option = "user=" + u.Username + " password=" + u.Username + "1234 dbname=" + u.Username + ".test sslmode=disable"
		}
		connection, err = sql.Open("postgres", option)
		if err != nil {
			log.Fatal(err)
		}
	}
	return connection
}
