package handlers

import (
	"fmt"
	"inwdic/queries"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetVocab is a handler function of GET request that recieve vocab parameter then use it to query.
//Add all the forms into wordList and then encode it to JSON

func GetWord(ctx *gin.Context) {

	words, err := queries.SelectWord(ctx.Query("vocab"))
	if err != nil {
		fmt.Println(err)
	}
	for i, word := range words {
		if word.Type == "คำกริยา (Verb 1)" || word.Type == "คำกริยา (Verb 2)" || word.Type == "คำกริยา (Verb 3)" {
			words[i].Forms.Dict = words[i].ToDictForm()
			words[i].Forms.Ta = words[i].ToTaForm()
			words[i].Forms.Te = words[i].ToTeForm()
			words[i].Forms.Nai = words[i].ToNaiForm()
		}

	}
	ctx.JSON(http.StatusOK, words)
}

func GetWordList(ctx *gin.Context) {
	wordLists, err := queries.SelectWordList(ctx.Query("jlpt"), ctx.Query("row"))
	if err != nil {
		fmt.Println(err)
	}
	for i, word := range wordLists {
		if word.Type == "คำกริยา (Verb 1)" || word.Type == "คำกริยา (Verb 2)" || word.Type == "คำกริยา (Verb 3)" {
			wordLists[i].Forms.Dict = wordLists[i].ToDictForm()
			wordLists[i].Forms.Ta = wordLists[i].ToTaForm()
			wordLists[i].Forms.Te = wordLists[i].ToTeForm()
			wordLists[i].Forms.Nai = wordLists[i].ToNaiForm()
		}

	}
	ctx.JSON(http.StatusOK, wordLists)
}
