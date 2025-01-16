package routes

import (
	"gateway_service/handlers"

	"github.com/gin-gonic/gin"
)

func EmployeeRoutes(router *gin.Engine) {
	employeeServiceURL := "http://localhost:8081"

	employeeGroup := router.Group("/users")
	{
		employeeGroup.GET("/", handlers.ProxyHandler(employeeServiceURL))
		employeeGroup.GET("/:id", handlers.ProxyHandler(employeeServiceURL))
		employeeGroup.POST("/", handlers.ProxyHandler(employeeServiceURL))
		employeeGroup.PUT("/:id", handlers.ProxyHandler(employeeServiceURL))
		employeeGroup.DELETE("/:id", handlers.ProxyHandler(employeeServiceURL))
	}
}
