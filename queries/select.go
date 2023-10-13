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
	selectWord := fmt.Sprintf(`SELECT * FROM words WHERE vocab = '%s' or hiragana = '%s' ;`, vocab, vocab)
	rows, err := db.Query(selectWord)
	if err != nil {
		return nil, err
	}

	var words []utils.Word

	for rows.Next() {
		var word utils.Word
		if err := rows.Scan(&word.Id, &word.Vocab, &word.Hiragana,
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
	selectWord := fmt.Sprintf(`SELECT * FROM words WHERE jlpt = '%s' ORDER BY random() LIMIT %v ;`, jlpt, row)
	rows, err := db.Query(selectWord)
	if err != nil {
		return nil, err
	}

	var words []utils.Word

	for rows.Next() {
		var word utils.Word
		if err := rows.Scan(&word.Id, &word.Vocab, &word.Hiragana,
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
	selectUser := fmt.Sprintf(`SELECT * FROM users WHERE name = '%s' and password = '%s';`, name, password)
	rows, err := db.Query(selectUser)
	if err != nil {
		return utils.User{}, "", err
	}
	var user utils.User
	var image string
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Name, &user.Password, &image); err != nil {
			return utils.User{}, "", err
		}
	}
	if err = rows.Err(); err != nil {
		return utils.User{}, "", err
	}
	fmt.Printf("%v \n %s", user, image)
	return user, image, nil
}

func SelectVocabLists(user_id string) ([]utils.VocabList, error) {
	db := database.ConnectDatabase()
	defer db.Close()
	selectVocabListId := fmt.Sprintf(`SELECT * FROM vocablists WHERE user_id = %s;`, user_id)
	rows, err := db.Query(selectVocabListId)
	if err != nil {
		return []utils.VocabList{}, err
	}
	var vcls []utils.VocabList
	for rows.Next() {
		var vcl utils.VocabList
		if err := rows.Scan(&vcl.Id, &vcl.User_id, &vcl.Name); err != nil {
			return []utils.VocabList{}, err
		}
		vcls = append(vcls, vcl)
	}
	if err = rows.Err(); err != nil {
		return []utils.VocabList{}, err
	}
	return vcls, nil

}

func SelectVocabDetails(vocablist_id string) ([]utils.VocabDetail, error) {
	db := database.ConnectDatabase()
	defer db.Close()
	selectVocabDetail := fmt.Sprintf(`SELECT w.id,vd.vocablist_id,w.vocab,w.meaning FROM words AS w JOIN vocabdetails AS vd ON w.id = vd.word_id JOIN vocablists AS vl ON vd.vocablist_id = vl.id WHERE vd.vocablist_id = %s ;`, vocablist_id)
	rows, err := db.Query(selectVocabDetail)
	if err != nil {
		return []utils.VocabDetail{}, err
	}
	var vocabDetails []utils.VocabDetail
	for rows.Next() {
		var vocabDetail utils.VocabDetail
		if err := rows.Scan(&vocabDetail.Word_id, &vocabDetail.VocabList_id, &vocabDetail.Vocab, &vocabDetail.Meaning); err != nil {
			return []utils.VocabDetail{}, err
		}
		vocabDetails = append(vocabDetails, vocabDetail)
	}
	if err = rows.Err(); err != nil {
		return []utils.VocabDetail{}, err
	}
	return vocabDetails, nil

}
