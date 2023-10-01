package queries

import (
	"errors"
	"inwdic/utils"
	"reflect"
	"testing"
)

type selectWord struct {
	vocab   string
	result1 []utils.Word
	result2 error
}

var testSelects = []selectWord{
	{"きます", []utils.Word{{
		Vocab:    "きます",
		Hiragana: "きます",
		Type:     "คำกริยา (Verb 3)",
		Meaning:  "มา",
		Jlpt:     "N5",
	}, {
		Vocab:    "きます",
		Hiragana: "きます",
		Type:     "คำกริยา (Verb 2)",
		Meaning:  "ใส่ (เสื้อผ้า)",
		Jlpt:     "N5",
	}}, nil},
	{"します", []utils.Word{{
		Vocab:    "します",
		Hiragana: "します",
		Type:     "คำกริยา (Verb 3)",
		Meaning:  "ทำ (do)",
		Jlpt:     "N5",
	},
	}, nil},
	{"とります", []utils.Word{{
		Vocab:    "撮ります",
		Hiragana: "とります",
		Type:     "คำกริยา (Verb 1)",
		Meaning:  "ถ่าย (รูป, วีดีโอ)",
		Jlpt:     "N5",
	},
	}, nil},
}

func TestSelectWord(t *testing.T) {

	for _, testSelect := range testSelects {
		// Call the function being tested
		result1, result2 := SelectWord(testSelect.vocab)

		// Define the expected result
		expected1, expected2 := testSelect.result1, testSelect.result2

		// Compare the actual result with the expected result
		if !reflect.DeepEqual(result1, expected1) {
			t.Errorf("Expected result1: %v, Got result1: %v", expected1, result1)
		}

		if !errors.Is(result2, expected2) {
			t.Errorf("Expected error: %v, Got error: %v", expected2, result2)
		}
	}

}
