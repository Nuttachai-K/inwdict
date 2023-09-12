package utils

import "testing"

type testForm struct {
	arg1   Word
	result string
}

var testDictForms = []testForm{
	{Word{
		Vocab: "写します",
		Type:  "คำกริยา (Verb 1)",
	}, "写す"},
	{Word{
		Vocab: "止まります",
		Type:  "คำกริยา (Verb 1)",
	}, "止まる"},
	{Word{
		Vocab: "みます",
		Type:  "คำกริยา (Verb 2)",
	}, "みる"},
	{Word{
		Vocab: "きます",
		Type:  "คำกริยา (Verb 2)",
	}, "きる"},
	{Word{
		Vocab: "します",
		Type:  "คำกริยา (Verb 3)",
	}, "する"},
	{Word{
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
	{Word{
		Vocab: "写します",
		Type:  "คำกริยา (Verb 1)",
	}, "写した"},
	{Word{
		Vocab: "いきます",
		Type:  "คำกริยา (Verb 1)",
	}, "いった"},
	{Word{
		Vocab: "みます",
		Type:  "คำกริยา (Verb 2)",
	}, "みた"},
	{Word{
		Vocab: "きます",
		Type:  "คำกริยา (Verb 2)",
	}, "きた"},
	{Word{
		Vocab: "します",
		Type:  "คำกริยา (Verb 3)",
	}, "した"},
	{Word{
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
	{Word{
		Vocab: "写します",
		Type:  "คำกริยา (Verb 1)",
	}, "写して"},
	{Word{
		Vocab: "いきます",
		Type:  "คำกริยา (Verb 1)",
	}, "いって"},
	{Word{
		Vocab: "みます",
		Type:  "คำกริยา (Verb 2)",
	}, "みて"},
	{Word{
		Vocab: "きます",
		Type:  "คำกริยา (Verb 2)",
	}, "きて"},
	{Word{
		Vocab: "します",
		Type:  "คำกริยา (Verb 3)",
	}, "して"},
	{Word{
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
	{Word{
		Vocab: "写します",
		Type:  "คำกริยา (Verb 1)",
	}, "写さない"},
	{Word{
		Vocab: "あります",
		Type:  "คำกริยา (Verb 1)",
	}, "ない"},
	{Word{
		Vocab: "みます",
		Type:  "คำกริยา (Verb 2)",
	}, "みない"},
	{Word{
		Vocab: "きます",
		Type:  "คำกริยา (Verb 2)",
	}, "きない"},
	{Word{
		Vocab: "します",
		Type:  "คำกริยา (Verb 3)",
	}, "しない"},
	{Word{
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
