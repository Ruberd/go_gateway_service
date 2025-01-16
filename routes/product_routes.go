package routes

import (
	"gateway_service/handlers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	productServiceURL := "http://localhost:8082"

	productGroup := router.Group("/products")
	{
		productGroup.GET("/", handlers.ProxyHandler(productServiceURL))
		productGroup.GET("/:id", handlers.ProxyHandler(productServiceURL))
		productGroup.POST("/", handlers.ProxyHandler(productServiceURL))
		productGroup.PUT("/:id", handlers.ProxyHandler(productServiceURL))
		productGroup.DELETE("/:id", handlers.ProxyHandler(productServiceURL))
	}
}
