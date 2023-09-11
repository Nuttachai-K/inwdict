package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type WordList struct {
	Vocab    string `json:"vocab"`
	Hiragana string `json:"hiragana"`
	Type     string `json:"type"`
	Meaning  string `json:"meaning"`
	Jlpt     string `json:"jlpt"`
}

// Recieve pointer of slice of WordList and encoding it to json file
func wordListToJson(sliceWordList *[]WordList, jsonFile string) {
	file, _ := json.Marshal(sliceWordList)
	_ = os.WriteFile(jsonFile, file, 0644)
}

// Get word type from w and change its form to dict form according to its type
func (w WordList) ToDictForm() string {
	var dict = ""
	fmt.Printf("hiragana : %s len : %v\n", w.Vocab, len(w.Vocab))
	last3char := w.Vocab[len(w.Vocab)-9:]
	fmt.Printf("Last 3 Character : %s\n", last3char)
	if w.Type == "คำกริยา (Verb 1)" {

		switch last3char {
		case "います":
			dict = fmt.Sprintf("%sう", w.Vocab[:len(w.Vocab)-9])

		case "します":
			dict = fmt.Sprintf("%sす", w.Vocab[:len(w.Vocab)-9])

		case "きます":
			dict = fmt.Sprintf("%sく", w.Vocab[:len(w.Vocab)-9])

		case "ぎます":
			dict = fmt.Sprintf("%sぐ", w.Vocab[:len(w.Vocab)-9])

		case "みます":
			dict = fmt.Sprintf("%sむ", w.Vocab[:len(w.Vocab)-9])

		case "びます":
			dict = fmt.Sprintf("%sぶ", w.Vocab[:len(w.Vocab)-9])

		case "ちます":
			dict = fmt.Sprintf("%sつ", w.Vocab[:len(w.Vocab)-9])

		case "ります":
			dict = fmt.Sprintf("%sる", w.Vocab[:len(w.Vocab)-9])

		default:
			dict = ""
		}

	} else if w.Type == "คำกริยา (Verb 2)" {

		dict = fmt.Sprintf("%sる", w.Vocab[:len(w.Vocab)-6])

	} else if w.Type == "คำกริยา (Verb 3)" {

		if last3char == "します" {
			dict = fmt.Sprintf("%sする", w.Vocab[:len(w.Vocab)-9])
		} else if last3char == "きます" {
			dict = fmt.Sprintf("%sくる", w.Vocab[:len(w.Vocab)-9])
		}
	}
	fmt.Printf("futsukei : %s\n\n", dict)
	return dict
}
