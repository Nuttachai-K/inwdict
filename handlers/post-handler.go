package handlers

import (
	"encoding/base64"
	"fmt"
	"inwdic/queries"
	"inwdic/utils"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostUser(ctx *gin.Context) {

	var user utils.User
	// if err := ctx.ShouldBindJSON(&user); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	user.Name = ctx.Request.FormValue("name")
	user.Password = ctx.Request.FormValue("password")

	file, _, err := ctx.Request.FormFile("image")
	if err != nil {
		queries.InsertProfile(user, "")
		// Return a success response without image
		ctx.JSON(http.StatusOK, gin.H{"message": "User uploaded successfully without picture"})
		return
	} else {
		// Read the image data
		imageData, err := io.ReadAll(file)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading image data"})
			return
		}

		// Encode the image as Base64
		encodedImage := base64.StdEncoding.EncodeToString(imageData)

		fmt.Printf("%v \n", user)

		// Process the user data and save it to a database
		queries.InsertProfile(user, encodedImage)
		ctx.JSON(http.StatusOK, gin.H{"message": "User uploaded successfully with picture"})

	}
	defer file.Close()

}

func PostVocabList(ctx *gin.Context) {

	var vcl utils.VocabList
	var err error
	vcl.User_id, err = strconv.Atoi(ctx.Request.FormValue("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user id is not match"})
		return
	}
	vcl.Name = ctx.Request.FormValue("name")

	queries.InsertVocabList(vcl)

	ctx.JSON(http.StatusOK, gin.H{"message": "VocabList uploaded successfully"})
}

func PostVocabDetail(ctx *gin.Context) {
	var word_id, vocabList_id int
	var err error
	word_id, err = strconv.Atoi(ctx.Request.FormValue("word_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "word id is not match"})
		return
	}
	vocabList_id, err = strconv.Atoi(ctx.Request.FormValue("vocablist_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "vocablist id is not match"})
		return
	}

	queries.InsertVocabDetail(word_id, vocabList_id)
	ctx.JSON(http.StatusOK, gin.H{"message": "VocabDetail uploaded successfully"})
}
