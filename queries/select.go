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
	selectWord := fmt.Sprintf(`SELECT * FROM dicts WHERE vocab = '%s' or hiragana = '%s' ;`, vocab, vocab)
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
	selectWord := fmt.Sprintf(`SELECT * FROM dicts WHERE jlpt = '%s' ORDER BY random() LIMIT %v ;`, jlpt, row)
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

func SelectUser(name, password string) (utils.User, string, error) {
	db := database.ConnectDatabase()
	defer db.Close()
	selectUser := fmt.Sprintf(`SELECT name,password,image FROM users WHERE name = '%s' and password = '%s';`, name, password)
	rows, err := db.Query(selectUser)
	if err != nil {
		return utils.User{}, "", err
	}
	var user utils.User
	var image string
	for rows.Next() {
		if err := rows.Scan(&user.Name, &user.Password, &image); err != nil {
			return utils.User{}, "", err
		}
	}
	if err = rows.Err(); err != nil {
		return utils.User{}, "", err
	}
	fmt.Printf("%v \n %s", user, image)
	return user, image, nil
}
