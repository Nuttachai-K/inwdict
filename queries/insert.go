package queries

import (
	"database/sql"
	"fmt"
)

func InsertJson(db *sql.DB) {
	insertJson := `INSERT INTO wordlists (vocab,hiragana,type,meaning,jlpt) VALUES ('きいてください','kiitekudasai','วลี (Phrase)','กรุณาฟัง','N5')`
	_, err := db.Exec(insertJson)
	if err != nil {
		fmt.Println(err)
	}
}
