package queries

import (
	"encoding/json"
	"fmt"
	"inwdic/database"
	"inwdic/utils"
	"io"
	"os"
)

func InsertJsonWordlist(jpn string) error {
	db := database.ConnectDatabase()
	defer db.Close()
	jsonFile, err := os.Open(fmt.Sprintf("json/%s.json", jpn))

	if err != nil {
		return err
	}
	fmt.Printf("Successfully Opened %s.json\n", jpn)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var data []utils.Word

	if err := json.Unmarshal(byteValue, &data); err != nil {
		return err
	}

	for _, word := range data {
		insertJson := fmt.Sprintf(`INSERT INTO dicts (vocab,hiragana,type,meaning,jlpt) VALUES ('%s','%s','%s','%s','%s')`, word.Vocab, word.Hiragana, word.Type, word.Meaning, word.Jlpt)
		_, err = db.Exec(insertJson)
		if err != nil {
			return err
		}
	}
	return nil

}

func InsertAllJsonWordList() error {
	err := InsertJsonWordlist("jpn101")
	if err != nil {
		return err
	}
	err = InsertJsonWordlist("jpn102")
	if err != nil {
		return err
	}
	err = InsertJsonWordlist("jpn201")
	if err != nil {
		return err
	}
	err = InsertJsonWordlist("jpn202")
	if err != nil {
		return err
	}
	err = InsertJsonWordlist("jpn301")
	if err != nil {
		return err
	}
	return nil
}

func InsertProfile(user utils.User, image string) error {
	db := database.ConnectDatabase()
	defer db.Close()
	insertJson := fmt.Sprintf(`INSERT INTO users (name,password,image) VALUES ('%s','%s','%s')`, user.Name, user.Password, image)
	_, err := db.Exec(insertJson)
	if err != nil {
		return err
	}
	return nil
}
