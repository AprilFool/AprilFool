package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os/user"
)

var connection *sql.DB = nil
var name string = ""

func SetConnection(n string) {
	name = n
}

func GetConnection() *sql.DB {
	var err error

	if connection == nil {
		if name == "" {
			u, _ := user.Current()
			name = "user="+u.Username+" password="+u.Username+"1234 dbname="+u.Username+".test"
		}
		connection, err = sql.Open("postgres", name)
		if err != nil {
			log.Fatal(err)
		}
	}
	return connection
}
