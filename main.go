package main

import (
	"inwdic/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	/*utils.CreateExcel("excel/JPN101_Wordlist.xlsx", "L1-L10 Vocabulary", 2, 4, 3, 5, "json/jpn101.json")
	utils.CreateExcel("excel/JPN102_Wordlist.xlsx", "L11-L18 Vocabulary", 2, 3, 4, 5, "json/jpn102.json")
	utils.CreateExcel("excel/JPN201_Wordlist.xlsx", "L19-L26 Vocabulary", 2, 3, 4, 5, "json/jpn201.json")
	utils.CreateExcel("excel/JPN202_Wordlist.xlsx", "L27-L32 Vocabulary", 2, 3, 4, 5, "json/jpn202.json")
	utils.CreateExcel("excel/JPN301_Wordlist.xlsx", "L33-L38 Vocabulary", 2, 3, 4, 5, "json/jpn301.json")*/
	// database.CreateTable()
	// err := queries.InsertJson("jpn301")
	// if err != nil {
	// 	fmt.Print(err)
	// }

	/*word, err := queries.SelectWord("ちかい")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%+v", word)*/

	router := gin.Default()

	router.GET("/dict", handlers.GetHandler)
	router.Run(":8080")

}
