package utils

import (
	"encoding/json"
	"os"
)

type WordList struct {
	Vocab    string `json:"vocab"`
	Hiragana string `json:"hiragana"`
	Type     string `json:"type"`
	Meaning  string `json:"meaning"`
	Jlpt     string `json:"jlpt"`
}

func wordListToJson(sliceWordList *[]WordList, jsonFile string) {
	file, _ := json.Marshal(sliceWordList)
	_ = os.WriteFile(jsonFile, file, 0644)
}
