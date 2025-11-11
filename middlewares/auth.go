package middlewares

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CustomerRequired(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userID := session.Get("user_id")
	userRole := session.Get("user_role")

	if userID == nil {
		ctx.JSON(http.StatusOK, "未登入")
		ctx.Abort()
		return
	}

	if userRole != "customer" {
		ctx.JSON(http.StatusUnauthorized, "非顧客身份")
		ctx.Abort()
		return
	}

	ctx.Next()
}

func AdminRequired(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userID := session.Get("user_id")
	userRole := session.Get("user_role")

	if userID == nil {
		ctx.JSON(http.StatusUnauthorized, "未登入")
		ctx.Abort()
		return
	}

	if userRole != "admin" {
		ctx.JSON(http.StatusUnauthorized, "權限不足")
		ctx.Abort()
		return
	}

	ctx.Next()
}
