package queries

import (
	"fmt"
	"inwdic/database"
	"inwdic/dict"
	"strconv"
)

func SelectWord(vocab string) ([]dict.Word, error) {
	db := database.ConnectDatabase()
	defer db.Close()
	selectWord := fmt.Sprintf(`SELECT * FROM dict WHERE vocab = '%s' or hiragana = '%s' ;`, vocab, vocab)
	rows, err := db.Query(selectWord)
	if err != nil {
		fmt.Println(err)
	}

	var words []dict.Word

	for rows.Next() {
		var word dict.Word
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

func SelectWordList(jlpt string, limitString string) ([]dict.Word, error) {
	db := database.ConnectDatabase()
	defer db.Close()
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		return nil, err
	}
	fmt.Printf("limit : %v type : %T  limitString : %v type : %T\n", limit, limit, limitString, limitString)
	selectWord := fmt.Sprintf(`SELECT * FROM dict WHERE jlpt = '%s' ORDER BY random() LIMIT %v ;`, jlpt, limit)
	rows, err := db.Query(selectWord)
	if err != nil {
		return nil, err
	}

	var words []dict.Word

	for rows.Next() {
		var word dict.Word
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
