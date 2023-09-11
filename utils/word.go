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
	Forms    Form   `json:"forms"`
	Meaning  string `json:"meaning"`
	Jlpt     string `json:"jlpt"`
}

type Form struct {
	Dict     string `json:"dict"`
	Ta       string `json:"ta"`
	Te       string `json:"te"`
	Negative string `json:"negative"`
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

		case "にます":
			dict = fmt.Sprintf("%sぬ", w.Vocab[:len(w.Vocab)-9])

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

func (w WordList) ToTaForm() string {
	var ta = ""
	fmt.Printf("hiragana : %s len : %v\n", w.Vocab, len(w.Vocab))
	last3char := w.Vocab[len(w.Vocab)-9:]
	fmt.Printf("Last 3 Character : %s\n", last3char)
	if w.Type == "คำกริยา (Verb 1)" {

		switch last3char {
		case "います":
			ta = fmt.Sprintf("%sった", w.Vocab[:len(w.Vocab)-9])

		case "します":
			ta = fmt.Sprintf("%sした", w.Vocab[:len(w.Vocab)-9])

		case "きます":
			ta = fmt.Sprintf("%sいた", w.Vocab[:len(w.Vocab)-9])

		case "ぎます":
			ta = fmt.Sprintf("%sいだ", w.Vocab[:len(w.Vocab)-9])

		case "みます":
			ta = fmt.Sprintf("%sんだ", w.Vocab[:len(w.Vocab)-9])

		case "びます":
			ta = fmt.Sprintf("%sんだ", w.Vocab[:len(w.Vocab)-9])

		case "ちます":
			ta = fmt.Sprintf("%sった", w.Vocab[:len(w.Vocab)-9])

		case "ります":
			ta = fmt.Sprintf("%sった", w.Vocab[:len(w.Vocab)-9])

		case "にます":
			ta = fmt.Sprintf("%sんだ", w.Vocab[:len(w.Vocab)-9])

		default:
			ta = ""
		}

	} else if w.Type == "คำกริยา (Verb 2)" {

		ta = fmt.Sprintf("%sた", w.Vocab[:len(w.Vocab)-6])

	} else if w.Type == "คำกริยา (Verb 3)" {

		if last3char == "します" {
			ta = fmt.Sprintf("%sした", w.Vocab[:len(w.Vocab)-9])
		} else if last3char == "きます" {
			ta = fmt.Sprintf("%sきた", w.Vocab[:len(w.Vocab)-9])
		}
	}
	fmt.Printf("ta form : %s\n\n", ta)
	return ta
}

func (w WordList) ToTeForm() string {
	var te = ""
	fmt.Printf("hiragana : %s len : %v word type : %s\n ", w.Vocab, len(w.Vocab), w.Type)
	last3char := w.Vocab[len(w.Vocab)-9:]
	fmt.Printf("Last 3 Character : %s\n", last3char)
	if w.Type == "คำกริยา (Verb 1)" {

		switch last3char {
		case "います":
			te = fmt.Sprintf("%sって", w.Vocab[:len(w.Vocab)-9])

		case "します":
			te = fmt.Sprintf("%sして", w.Vocab[:len(w.Vocab)-9])

		case "きます":
			te = fmt.Sprintf("%sいで", w.Vocab[:len(w.Vocab)-9])

		case "ぎます":
			te = fmt.Sprintf("%sいで", w.Vocab[:len(w.Vocab)-9])

		case "みます":
			te = fmt.Sprintf("%sんで", w.Vocab[:len(w.Vocab)-9])

		case "びます":
			te = fmt.Sprintf("%sんで", w.Vocab[:len(w.Vocab)-9])

		case "ちます":
			te = fmt.Sprintf("%sって", w.Vocab[:len(w.Vocab)-9])

		case "ります":
			te = fmt.Sprintf("%sって", w.Vocab[:len(w.Vocab)-9])

		case "にます":
			te = fmt.Sprintf("%sんで", w.Vocab[:len(w.Vocab)-9])

		default:
			te = ""
		}

	} else if w.Type == "คำกริยา (Verb 2)" {

		te = fmt.Sprintf("%sて", w.Vocab[:len(w.Vocab)-6])

	} else if w.Type == "คำกริยา (Verb 3)" {

		if last3char == "します" {
			te = fmt.Sprintf("%sして", w.Vocab[:len(w.Vocab)-9])
		} else if last3char == "きます" {
			te = fmt.Sprintf("%sきて", w.Vocab[:len(w.Vocab)-9])
		}
	}
	fmt.Printf("te form : %s\n\n", te)
	return te
}
