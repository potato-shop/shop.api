package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"shop.go/config"
	"shop.go/models"
)

type AddCartItemRequest struct {
	ProductID uint
	Quantity  uint
	UnitPrice float64
}

func AddCartItem(ctx *gin.Context) {
	// 找 user
	session := sessions.Default(ctx)
	userID := session.Get("user_id")
	user := models.User{}
	err := config.DB.First(&user, userID).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, userID)
}
