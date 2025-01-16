package main

import (
	"gateway_service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Register routes
	routes.ProductRoutes(router)
	routes.EmployeeRoutes(router)
	routes.AuthenticationRoutes(router)

	// Start the gateway
	if err := router.Run(":8083"); err != nil {
		panic(err)
	}
}
