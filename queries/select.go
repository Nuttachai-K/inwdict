package queries

import (
	"database/sql"
	"fmt"
	"inwdic/utils"
)

func SelectWord(db *sql.DB, vocab string) ([]utils.WordList, error) {
	selectWord := fmt.Sprintf(`SELECT * FROM wordlists WHERE vocab = '%s' and id = 1;`, vocab)
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
