package routes

import (
	"github.com/gin-gonic/gin"
	"shop.go/handlers"
	"shop.go/middlewares"
)

func Setup(router *gin.Engine) {
	api := router.Group("/api")

	// Auth
	api.POST("/user/signup", handlers.Signup("user"))
	api.POST("/admin/signup", handlers.Signup("admin"))
	api.POST("/user/login", handlers.Login([]string{"user"}))
	api.POST("/admin/login", handlers.Login([]string{"admin", "guest"}))

	// 用戶
	api.GET("/me", middlewares.Auth([]string{"admin", "user"}), handlers.GetUser)
	api.GET("/users", middlewares.Auth([]string{"admin"}), handlers.ListUsers)
	api.PUT("/user/avatar", middlewares.Auth([]string{"admin", "user"}), handlers.UpdateUserImage)
	api.PUT("/user/:userId/password", middlewares.Auth([]string{"admin"}), handlers.ResetUserPassword)

	// 種類
	api.GET("/categories", middlewares.Auth([]string{"admin", "user"}), handlers.ListCategories)
	api.POST("/category", middlewares.Auth([]string{"admin"}), handlers.AddCategory)
	api.PUT("/category/:categoryId", middlewares.Auth([]string{"admin"}), handlers.UpdateCategory)
	api.DELETE("/category/:categoryId", middlewares.Auth([]string{"admin"}), handlers.DeleteCategory)

	// 商品
	api.GET("/products", middlewares.Auth([]string{"admin", "user"}), handlers.ListProducts)
	api.GET("/product/:productId", middlewares.Auth([]string{"admin", "user"}), handlers.GetProduct)
	api.POST("/product", middlewares.Auth([]string{"admin"}), handlers.AddProduct)
	api.PUT("/product/:productId", middlewares.Auth([]string{"admin"}), handlers.UpdateProduct)
	api.PUT("/product/:productId/image", middlewares.Auth([]string{"admin"}), handlers.UpdateProductImage)
	api.DELETE("/product/:productId", middlewares.Auth([]string{"admin"}), handlers.DeleteProduct)

	// 訂單
	api.GET("/order/:orderId", handlers.GetOrder)
	api.GET("/user/me/orders", middlewares.Auth([]string{"user"}), handlers.ListOrdersByCustomer)
	api.GET("/orders", middlewares.Auth([]string{"admin"}), handlers.ListOrdersByAdmin)
	api.POST("/order", middlewares.Auth([]string{"user"}), handlers.CreateOrder)
	api.PUT("/order/:orderId", middlewares.Auth([]string{"admin"}), handlers.UpdateOrder)

	// 購物車
	api.POST("/cart/item", middlewares.Auth([]string{"user"}), handlers.AddCartItem)
	api.PUT("/cart/item/:cartItemId", middlewares.Auth([]string{"user"}), handlers.UpdateCartItemQuantity)
	api.DELETE("/cart/item/:cartItemId", middlewares.Auth([]string{"user"}), handlers.DeleteCartItem)
	api.DELETE("/cart/item/all", middlewares.Auth([]string{"user"}), handlers.DeleteAllCartItem)
}
