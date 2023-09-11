package utils

import "testing"

type testForm struct {
	arg1   WordList
	result string
}

var testDictForms = []testForm{
	{WordList{
		Vocab: "写します",
		Type:  "คำกริยา (Verb 1)",
	}, "写す"},
	{WordList{
		Vocab: "止まります",
		Type:  "คำกริยา (Verb 1)",
	}, "止まる"},
	{WordList{
		Vocab: "みます",
		Type:  "คำกริยา (Verb 2)",
	}, "みる"},
	{WordList{
		Vocab: "きます",
		Type:  "คำกริยา (Verb 2)",
	}, "きる"},
	{WordList{
		Vocab: "します",
		Type:  "คำกริยา (Verb 3)",
	}, "する"},
	{WordList{
		Vocab: "きます",
		Type:  "คำกริยา (Verb 3)",
	}, "くる"},
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

var testNaiForms = []testForm{
	{WordList{
		Vocab: "写します",
		Type:  "คำกริยา (Verb 1)",
	}, "写さない"},
	{WordList{
		Vocab: "止まります",
		Type:  "คำกริยา (Verb 1)",
	}, "止まらない"},
	{WordList{
		Vocab: "みます",
		Type:  "คำกริยา (Verb 2)",
	}, "みない"},
	{WordList{
		Vocab: "きます",
		Type:  "คำกริยา (Verb 2)",
	}, "きない"},
	{WordList{
		Vocab: "します",
		Type:  "คำกริยา (Verb 3)",
	}, "しない"},
	{WordList{
		Vocab: "きます",
		Type:  "คำกริยา (Verb 3)",
	}, "こない"},
}

func TestToNaiForm(t *testing.T) {

	for _, testNaiList := range testNaiForms {
		// Call the function being tested
		result := testNaiList.arg1.ToNaiForm()

		// Define the expected result
		expected := testNaiList.result

		// Compare the actual result with the expected result
		if result != expected {
			t.Errorf("Expected: %s, Got: %s", expected, result)
		}
	}

}
