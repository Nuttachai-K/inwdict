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

func CreateWordTable() {
	db := ConnectDatabase()
	defer db.Close()
	createTable := `CREATE TABLE IF NOT EXISTS words (
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

func CreateVocabListTable() {
	db := ConnectDatabase()
	defer db.Close()
	createTable := `CREATE TABLE IF NOT EXISTS vocablists (
		id SERIAL PRIMARY KEY,
		user_id INTEGER,
		name VARCHAR(255),
		FOREIGN KEY (user_id) REFERENCES users(id)
		);
		`
	_, err := db.Exec(createTable)
	if err != nil {
		fmt.Println(err)
	}
}

func CreateVocabDetailTable() {
	db := ConnectDatabase()
	defer db.Close()
	createTable := `CREATE TABLE IF NOT EXISTS vocabdetails (
		vocablist_id INTEGER,
		word_id INTEGER,
		PRIMARY KEY (vocablist_id, word_id),
		FOREIGN KEY (vocablist_id) REFERENCES vocablists(id),
		FOREIGN KEY (word_id) REFERENCES words(id)
		);
		`
	_, err := db.Exec(createTable)
	if err != nil {
		fmt.Println(err)
	}
}
