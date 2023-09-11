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

var testTaForms = []testForm{
	{WordList{
		Vocab: "写します",
		Type:  "คำกริยา (Verb 1)",
	}, "写した"},
	{WordList{
		Vocab: "いきます",
		Type:  "คำกริยา (Verb 1)",
	}, "いった"},
	{WordList{
		Vocab: "みます",
		Type:  "คำกริยา (Verb 2)",
	}, "みた"},
	{WordList{
		Vocab: "きます",
		Type:  "คำกริยา (Verb 2)",
	}, "きた"},
	{WordList{
		Vocab: "します",
		Type:  "คำกริยา (Verb 3)",
	}, "した"},
	{WordList{
		Vocab: "きます",
		Type:  "คำกริยา (Verb 3)",
	}, "きた"},
}

func TestToTaForm(t *testing.T) {

	for _, testTaList := range testTaForms {
		// Call the function being tested
		result := testTaList.arg1.ToTaForm()

		// Define the expected result
		expected := testTaList.result

		// Compare the actual result with the expected result
		if result != expected {
			t.Errorf("Expected: %s, Got: %s", expected, result)
		}
	}

}

var testTeForms = []testForm{
	{WordList{
		Vocab: "写します",
		Type:  "คำกริยา (Verb 1)",
	}, "写して"},
	{WordList{
		Vocab: "いきます",
		Type:  "คำกริยา (Verb 1)",
	}, "いって"},
	{WordList{
		Vocab: "みます",
		Type:  "คำกริยา (Verb 2)",
	}, "みて"},
	{WordList{
		Vocab: "きます",
		Type:  "คำกริยา (Verb 2)",
	}, "きて"},
	{WordList{
		Vocab: "します",
		Type:  "คำกริยา (Verb 3)",
	}, "して"},
	{WordList{
		Vocab: "きます",
		Type:  "คำกริยา (Verb 3)",
	}, "きて"},
}

func TestToTeForm(t *testing.T) {

	for _, testTeList := range testTeForms {
		// Call the function being tested
		result := testTeList.arg1.ToTeForm()

		// Define the expected result
		expected := testTeList.result

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
		Vocab: "あります",
		Type:  "คำกริยา (Verb 1)",
	}, "ない"},
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
