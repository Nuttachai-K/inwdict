package queries

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"inwdic/utils"
	"io"
	"os"
)

func InsertJson(db *sql.DB) {

	jsonFile, err := os.Open("json/jpn101.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened jpn101.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var data []utils.WordList

	if err := json.Unmarshal(byteValue, &data); err != nil {
		panic(err)
	}

	for _, word := range data {
		insertJson := fmt.Sprintf(`INSERT INTO wordlists (vocab,hiragana,type,meaning,jlpt) VALUES ('%s','%s','%s','%s','%s')`, word.Vocab, word.Hiragana, word.Type, word.Meaning, word.Jlpt)
		_, err = db.Exec(insertJson)
		if err != nil {
			fmt.Println(err)
		}
	}

}
