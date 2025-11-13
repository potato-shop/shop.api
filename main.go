package main

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"shop.go/config"
	"shop.go/handlers"
	"shop.go/middlewares"
)

func main() {
	// 初始化設定
	config.LoadEnvFile()
	config.ConnectDB()

	// 創建 Gin 路由器
	router := gin.Default()

	// 設置 Session
	router.Use(sessions.Sessions("user_session", config.CreateSessionStore()))

	// 路由設定
	setUpPublicRoutes(router)
	setUpWebRoutes(router)
	setUpAdminRoutes(router)

	// 啟動服務
	router.Run(":" + os.Getenv("APP_PORT"))
}

func setUpPublicRoutes(router *gin.Engine) {
	router.Static("/api/uploads", "./uploads")
}

func setUpWebRoutes(router *gin.Engine) {
	customerGroup := router.Group("/api")
	{
		// ======================== 不需 Customer 登入 ========================
		customerGroup.POST("/login", handlers.Login)
		customerGroup.GET("/categories", handlers.ListCategories)
		customerGroup.GET("/products", handlers.ListProducts)
		customerGroup.GET("/products/:productId", handlers.GetProduct)

		// ======================== 需要 Customer 登入 ========================
		// Auth
		customerGroup.GET("/me", middlewares.CustomerRequired, handlers.GetUser)
		customerGroup.POST("/logout", middlewares.CustomerRequired, handlers.Logout)

		// 購物車
		customerGroup.POST("/cart/items", middlewares.CustomerRequired, handlers.AddCartItem)
		customerGroup.PUT("/cart/items/:cartItemId", middlewares.CustomerRequired, handlers.UpdateCartItemQuantity)
		customerGroup.DELETE("/cart/items/:cartItemId", middlewares.CustomerRequired, handlers.DeleteCartItem)
		customerGroup.DELETE("/cart/items/all", middlewares.CustomerRequired, handlers.DeleteAllCartItem)

		// 訂單
		customerGroup.POST("/order", middlewares.CustomerRequired, handlers.CreateOrder)
		customerGroup.GET("/orders", handlers.ListOrdersByCustomer)
		customerGroup.GET("/orders/:orderId", handlers.GetOrder)
	}

}

func setUpAdminRoutes(router *gin.Engine) {
	adminGroup := router.Group("/api/admin")
	{
		// ======================== 不需 Admin 登入 ========================
		adminGroup.POST("/login", handlers.AdminLogin)

		// ======================== 需要 Admin 登入 ========================
		// Auth
		adminGroup.POST("/logout", middlewares.AdminRequired, handlers.Logout)
		adminGroup.POST("/signup", middlewares.AdminRequired, handlers.Signup)

		// 種類
		adminGroup.POST("/categories", middlewares.AdminRequired, handlers.AddCategory)
		adminGroup.GET("/categories", middlewares.AdminRequired, handlers.ListCategories)
		adminGroup.PUT("/categories/:categoryId", middlewares.AdminRequired, handlers.UpdateCategory)
		adminGroup.DELETE("/categories/:categoryId", handlers.DeleteCategory)

		// 商品
		adminGroup.POST("/products", middlewares.AdminRequired, handlers.AddProduct)
		adminGroup.GET("/products", middlewares.AdminRequired, handlers.ListProducts)
		adminGroup.PUT("/products/:productId", middlewares.AdminRequired, handlers.UpdateProduct)
		adminGroup.PUT("/products/:productId/image", middlewares.AdminRequired, handlers.UpdateProductImage)
		adminGroup.DELETE("/products/:productId", middlewares.AdminRequired, handlers.DeleteProduct)
	}
}
