package handlers

import (
	"fmt"
	"inwdic/queries"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetVocab is a handler function of GET request that recieve vocab parameter then use it to query.
//Add all the forms into wordList and then encode it to JSON

func GetVocab(ctx *gin.Context) {

	wordLists, err := queries.SelectWord(ctx.Query("vocab"))
	if err != nil {
		fmt.Println(err)
	}
	for i := range wordLists {
		wordLists[i].Forms.Dict = wordLists[i].ToDictForm()
		wordLists[i].Forms.Ta = wordLists[i].ToTaForm()
		wordLists[i].Forms.Te = wordLists[i].ToTeForm()
		wordLists[i].Forms.Nai = wordLists[i].ToNaiForm()
	}
	ctx.JSON(http.StatusOK, wordLists)
}
