package utils

import "testing"

type testDictForm struct {
	arg1   WordList
	result string
}

var testDictForms = []testDictForm{
	{WordList{
		Vocab: "します",
		Type:  "คำกริยา (Verb 3)",
	}, "する"},

	{WordList{
		Vocab: "写します",
		Type:  "คำกริยา (Verb 1)",
	}, "写す"},
	{WordList{
		Vocab: "渡します",
		Type:  "คำกริยา (Verb 1)",
	}, "渡す"},
	{WordList{
		Vocab: "泳ぎます",
		Type:  "คำกริยา (Verb 1)",
	}, "泳ぐ"},
	{WordList{
		Vocab: "止まります",
		Type:  "คำกริยา (Verb 1)",
	}, "止まる"},
	{WordList{
		Vocab: "選びます",
		Type:  "คำกริยา (Verb 1)",
	}, "選ぶ"},
	{WordList{
		Vocab: "きます",
		Type:  "คำกริยา (Verb 3)",
	}, "くる"},
	{WordList{
		Vocab: "みます",
		Type:  "คำกริยา (Verb 2)",
	}, "みる"},
	{WordList{
		Vocab: "きます",
		Type:  "คำกริยา (Verb 2)",
	}, "きる"},
}

func TestToDictForm(t *testing.T) {

	for _, testWordList := range testDictForms {
		// Call the function being tested
		result := testWordList.arg1.ToDictForm()

		// Define the expected result
		expected := testWordList.result

		// Compare the actual result with the expected result
		if result != expected {
			t.Errorf("Expected: %s, Got: %s", expected, result)
		}
	}

}
