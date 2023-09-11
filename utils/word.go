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

func wordListToJson(sliceWordList *[]WordList, jsonFile string) {
	file, _ := json.Marshal(sliceWordList)
	_ = os.WriteFile(jsonFile, file, 0644)
}

func (w WordList) ToFutsukei() string {
	var futsuke = ""
	fmt.Printf("hiragana : %s len : %v\n", w.Vocab, len(w.Vocab))
	last3char := w.Vocab[len(w.Vocab)-9:]
	fmt.Printf("Last 3 Character : %s\n", last3char)
	if w.Type == "คำกริยา (Verb 1)" {

		switch last3char {
		case "います":
			futsuke = fmt.Sprintf("%sう", w.Vocab[:len(w.Vocab)-9])

		case "します":
			futsuke = fmt.Sprintf("%sす", w.Vocab[:len(w.Vocab)-9])

		case "きます":
			futsuke = fmt.Sprintf("%sく", w.Vocab[:len(w.Vocab)-9])

		case "ぎます":
			futsuke = fmt.Sprintf("%sぐ", w.Vocab[:len(w.Vocab)-9])

		case "みます":
			futsuke = fmt.Sprintf("%sむ", w.Vocab[:len(w.Vocab)-9])

		case "びます":
			futsuke = fmt.Sprintf("%sぶ", w.Vocab[:len(w.Vocab)-9])

		case "ちます":
			futsuke = fmt.Sprintf("%sつ", w.Vocab[:len(w.Vocab)-9])

		case "ります":
			futsuke = fmt.Sprintf("%sる", w.Vocab[:len(w.Vocab)-9])

		default:
			futsuke = ""
		}

	} else if w.Type == "คำกริยา (Verb 2)" {

		futsuke = fmt.Sprintf("%sる", w.Vocab[:len(w.Vocab)-6])

	} else if w.Type == "คำกริยา (Verb 3)" {

		if last3char == "します" {
			futsuke = fmt.Sprintf("%sする", w.Vocab[:len(w.Vocab)-9])
		} else if last3char == "きます" {
			futsuke = fmt.Sprintf("%sこない", w.Vocab[:len(w.Vocab)-9])
		}
	}
	fmt.Printf("futsukei : %s\n\n", futsuke)
	return futsuke
}
