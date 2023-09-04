package queries

import (
	"fmt"
	"inwdic/database"
	"inwdic/utils"
)

func SelectWord(vocab string) ([]utils.WordList, error) {
	db := database.ConnectDatabase()
	defer db.Close()
	selectWord := fmt.Sprintf(`SELECT * FROM wordlists WHERE vocab = '%s' ;`, vocab)
	rows, err := db.Query(selectWord)
	if err != nil {
		fmt.Println(err)
	}

	var wordList []utils.WordList
	for rows.Next() {
		var word utils.WordList
		var ID int
		if err := rows.Scan(&ID, &word.Vocab, &word.Hiragana,
			&word.Type, &word.Meaning, &word.Jlpt); err != nil {
			return wordList, err
		}
		wordList = append(wordList, word)
	}
	if err = rows.Err(); err != nil {
		return wordList, err
	}
	return wordList, nil
}