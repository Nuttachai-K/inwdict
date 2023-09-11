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

func GetFutsukei(ctx *gin.Context) {
	var futsukeis []string
	wordLists, err := queries.SelectWord(ctx.Query("vocab"))
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range wordLists {
		futsukeis = append(futsukeis, v.ToFutsukei())
	}
	ctx.JSON(http.StatusOK, futsukeis)
}
