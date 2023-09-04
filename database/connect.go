package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Dragon_1234"
	dbname   = "inwdic"
)

func ConnectDatabase() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	Db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Connected!")
	return Db
}
