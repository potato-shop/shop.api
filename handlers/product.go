package handlers

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"shop.go/config"
	"shop.go/models"
)

type AddProductRequest struct {
	Name          string  `form:"Name" binding:"required"`
	CategoryID    uint    `form:"CategoryID" binding:"required"`
	Price         float64 `form:"Price" binding:"required"`
	StockQuantity uint    `form:"StockQuantity" binding:"required"`
	Description   string  `form:"Description" binding:"required"`
}

func AddProduct(ctx *gin.Context) {
	req := AddProductRequest{}

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	file, err := ctx.FormFile("UploadedFile")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// 儲存檔案
	ext := filepath.Ext(file.Filename)
	filename := uuid.New().String() + ext
	dst := filepath.Join("uploads", filename)

	log.Println(dst)

	err = ctx.SaveUploadedFile(file, dst)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, "儲存失敗")
		return
	}

	// DB 存紀錄
	product := models.Product{
		CategoryID:    req.CategoryID,
		Name:          req.Name,
		Description:   req.Description,
		Price:         req.Price,
		StockQuantity: req.StockQuantity,
		ImageURL:      dst,
	}

	err = config.DB.Create(&product).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, "cool")
}
