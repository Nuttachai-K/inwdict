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
	for _, v := range wordLists {
		fmt.Println(v.Meaning)
	}
	ctx.JSON(http.StatusOK, wordLists)
}

// Handler of GET request to
func GetDictForm(ctx *gin.Context) {
	var dictForms []string
	wordLists, err := queries.SelectWord(ctx.Query("vocab"))
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range wordLists {
		dictForms = append(dictForms, v.ToDictForm())
	}
	ctx.JSON(http.StatusOK, dictForms)
}
