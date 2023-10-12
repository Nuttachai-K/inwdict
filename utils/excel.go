package utils

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

// CreateExcel recieve fileName, sheet, cellNum1, cellNum2, cellNum3, cellNum4, jsonFile as parameters
// then use these parameter to get the specific information from excel file and add it to json file
func CreateJson(fileName string, sheet string, cellNum1, cellNum2, cellNum3, cellNum4 int, jsonFile string) {
	//Open spreadsheet
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		//Close spreadsheet
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	//Get all rows from the sheet
	rows, err := f.GetRows(sheet)
	if err != nil {
		fmt.Println(err)
		return
	}
	var sliceWordList []Word
	//Loop through all rows
	for ix, row := range rows {
		wordList := Word{}

		for iy, colCell := range row {
			if ix >= 2 {
				switch iy {
				case cellNum1:
					wordList.Vocab = colCell
				case cellNum2:
					wordList.Hiragana = colCell

				case cellNum3:
					wordList.Type = colCell

				case cellNum4:
					wordList.Meaning = colCell
				}

			}
		}
		wordList.Jlpt = "N5"
		if ix >= 2 {
			sliceWordList = append(sliceWordList, wordList)
		}
	}
	wordListToJson(&sliceWordList, jsonFile)
}
