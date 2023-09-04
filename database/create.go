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
		vocab VARCHAR(50),
		hiragana VARCHAR(50),
		type VARCHAR(50),
		meaning VARCHAR(200),
		jlpt VARCHAR(2) );		
	`
	_, err := db.Exec(createTable)
	if err != nil {
		fmt.Println(err)
	}
}
