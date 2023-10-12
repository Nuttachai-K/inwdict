package main

import (
	"inwdic/database"
	"inwdic/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	/*utils.CreateJson("excel/JPN101_Wordlist.xlsx", "L1-L10 Vocabulary", 2, 4, 3, 5, "json/jpn101.json")
	utils.CreateJson("excel/JPN102_Wordlist.xlsx", "L11-L18 Vocabulary", 2, 3, 4, 5, "json/jpn102.json")
	utils.CreateJson("excel/JPN201_Wordlist.xlsx", "L19-L26 Vocabulary", 2, 3, 4, 5, "json/jpn201.json")
	utils.CreateJson("excel/JPN202_Wordlist.xlsx", "L27-L32 Vocabulary", 2, 3, 4, 5, "json/jpn202.json")
	utils.CreateJson("excel/JPN301_Wordlist.xlsx", "L33-L38 Vocabulary", 2, 3, 4, 5, "json/jpn301.json")*/
	database.CreateWordTable()
	database.CreateUserTable()
	database.CreateVocabListTable()
	database.CreateVocabDetailTable()
	// err := queries.InsertAllJsonWordList()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	router := gin.Default()

	router.GET("/", handlers.GetWord)
	router.GET("/wordlist", handlers.GetWordList)
	router.GET("/user", handlers.GetUser)
	router.GET("/vocablist", handlers.GetVocabList)
	router.GET("/vocabdetails", handlers.GetVocabDetails)
	router.POST("/user", handlers.PostUser)
	router.POST("/vocablist", handlers.PostVocabList)
	router.POST("/vocabdetail", handlers.PostVocabDetail)
	router.Run(":8080")

}
