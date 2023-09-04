package handlers

import (
	"fmt"
	"inwdic/queries"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHandler(ctx *gin.Context) {

	wordLists, err := queries.SelectWord(ctx.Query("vocab"))
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range wordLists {
		fmt.Println(v.Meaning)
	}
	ctx.JSON(http.StatusOK, wordLists)
}
