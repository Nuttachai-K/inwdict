package handlers

import (
	"encoding/base64"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostImage(ctx *gin.Context) {
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

	// Here, you can process the encodedImage data or save it to a database as needed

	// Return a success response
	ctx.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully"})
}
