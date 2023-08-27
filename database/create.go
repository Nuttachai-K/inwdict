package database

import (
	"database/sql"
	"fmt"
)

func CreateDatabase(db *sql.DB) {
	createDb := `CREATE DATABASE inwdic;`
	_, err := db.Exec(createDb)
	if err != nil {
		fmt.Println(err)
	}

}

func CreateTable(db *sql.DB) {
	createTable := `CREATE TABLE wordlists (
		id SERIAL PRIMARY KEY,
		vocab VARCHAR(20),
		hiragana VARCHAR(20),
		type VARCHAR(20),
		meaning VARCHAR(50),
		jlpt VARCHAR(2) );		
	)`
	_, err := db.Exec(createTable)
	if err != nil {
		fmt.Println(err)
	}
}
