package utils

type VocabList struct {
	Id      int    `json:"id"`
	User_id int    `json:"userid"`
	Name    string `json:"name"`
}

type VocabDetail struct {
	Word_id      int    `json:"word_id"`
	VocabList_id int    `json:"vocablist_id"`
	Vocab        string `json:"vocab"`
	Meaning      string `json:"meaning"`
}
