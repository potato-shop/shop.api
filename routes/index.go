package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"shop.go/config"
)

func SetUpRoutes(router *gin.Engine) {
	router.POST("/upload", func(ctx *gin.Context) {
		file, err := ctx.FormFile("UploadedFile")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		// 打開上傳的文件
		fileContent, err := file.Open()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "無法打開文件"})
			return
		}
		defer fileContent.Close()

		// 上傳到 GCS
		err = config.UploadFile(ctx.Request.Context(), file.Filename, fileContent)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "上傳失敗: " + err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "上傳成功", "filename": file.Filename})
	})
}
