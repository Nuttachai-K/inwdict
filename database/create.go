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

func CreateDictTable() {
	db := ConnectDatabase()
	defer db.Close()
	createTable := `CREATE TABLE IF NOT EXISTS dict (
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

func CreateUserTable() {
	db := ConnectDatabase()
	defer db.Close()
	createTable := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255),
		password VARCHAR(255),
		image TEXT);
		`
	_, err := db.Exec(createTable)
	if err != nil {
		fmt.Println(err)
	}
}
