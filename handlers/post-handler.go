package handlers

import (
	"encoding/base64"
	"fmt"
	"inwdic/queries"
	"inwdic/utils"
	"io"
	"net/http"

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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Image file not provided"})
		return
	}
	defer file.Close()

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

	// Return a success response
	ctx.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully"})
}
