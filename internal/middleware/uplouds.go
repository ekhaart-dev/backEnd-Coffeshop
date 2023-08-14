package middleware

import (
	"backEnd_Coffeshop/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("banner_product")
	if err != nil {
		if err.Error() == "http: no such file" {
			ctx.Set("banner_product", "")
			ctx.Next()
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing file"})
		return
	}

	// Open the file
	src, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer src.Close()

	result, err := pkg.CloudInary(src)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to uploud file"})
		return
	}

	ctx.Set("banner_product", result)
	ctx.Next()
}
