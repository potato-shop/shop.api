package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"shop.go/config"
	"shop.go/middlewares"
	"shop.go/routes"
)

func main() {
	// 初始化設定
	config.LoadEnvFile()
	config.ConnectDB()
	config.ConnectStorage()

	// 創建 Gin 路由器
	router := gin.Default()

	// CORS 設定
	router.Use(middlewares.CORS())

	// 路由設定
	routes.Setup(router)

	// 啟動服務
	router.Run(":" + os.Getenv("APP_PORT"))
}
