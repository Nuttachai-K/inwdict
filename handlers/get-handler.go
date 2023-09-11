package handlers

import (
	"fmt"
	"inwdic/queries"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetVocab(ctx *gin.Context) {

	wordLists, err := queries.SelectWord(ctx.Query("vocab"))
	if err != nil {
		fmt.Println(err)
	}
	for i := range wordLists {
		wordLists[i].Forms.Dict = wordLists[i].ToDictForm()
		wordLists[i].Forms.Ta = wordLists[i].ToTaForm()
	}
	ctx.JSON(http.StatusOK, wordLists)
}
