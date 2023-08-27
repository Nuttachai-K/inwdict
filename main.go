package main

import (
	"inwdic/utils"
)

func main() {
	utils.CreateExcel("excel/JPN101_Wordlist.xlsx", "L1-L10 Vocabulary", 2, 4, 3, 5, "jpn101.json")
	utils.CreateExcel("excel/JPN102_Wordlist.xlsx", "L11-L18 Vocabulary", 2, 3, 4, 5, "jpn102.json")
	utils.CreateExcel("excel/JPN201_Wordlist.xlsx", "L19-L26 Vocabulary", 2, 3, 4, 5, "jpn201.json")
	utils.CreateExcel("excel/JPN202_Wordlist.xlsx", "L27-L32 Vocabulary", 2, 3, 4, 5, "jpn202.json")
	utils.CreateExcel("excel/JPN301_Wordlist.xlsx", "L33-L38 Vocabulary", 2, 3, 4, 5, "jpn301.json")
}
