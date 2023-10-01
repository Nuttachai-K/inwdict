package queries

import (
	"fmt"
	"inwdic/database"
	"inwdic/utils"
	"strconv"
)

func SelectWord(vocab string) ([]utils.Word, error) {
	db := database.ConnectDatabase()
	defer db.Close()
	selectWord := fmt.Sprintf(`SELECT * FROM dict WHERE vocab = '%s' or hiragana = '%s' ;`, vocab, vocab)
	rows, err := db.Query(selectWord)
	if err != nil {
		fmt.Println(err)
	}

	var words []utils.Word

	for rows.Next() {
		var word utils.Word
		var ID int
		if err := rows.Scan(&ID, &word.Vocab, &word.Hiragana,
			&word.Type, &word.Meaning, &word.Jlpt); err != nil {
			return words, err
		}
		words = append(words, word)
	}
	if err = rows.Err(); err != nil {
		return words, err
	}

	return words, nil
}

func SelectWordList(jlpt string, rowString string) ([]utils.Word, error) {
	db := database.ConnectDatabase()
	defer db.Close()
	row, err := strconv.Atoi(rowString)
	if err != nil {
		return nil, err
	}
	fmt.Printf("limit : %v type : %T  limitString : %v type : %T\n", row, row, rowString, rowString)
	selectWord := fmt.Sprintf(`SELECT * FROM dict WHERE jlpt = '%s' ORDER BY random() LIMIT %v ;`, jlpt, row)
	rows, err := db.Query(selectWord)
	if err != nil {
		return nil, err
	}

	var words []utils.Word

	for rows.Next() {
		var word utils.Word
		var ID int
		if err := rows.Scan(&ID, &word.Vocab, &word.Hiragana,
			&word.Type, &word.Meaning, &word.Jlpt); err != nil {
			return words, err
		}
		words = append(words, word)
	}
	if err = rows.Err(); err != nil {
		return words, err
	}

	return words, nil
}
