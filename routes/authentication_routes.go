package routes

import (
	"gateway_service/handlers"

	"github.com/gin-gonic/gin"
)

func AuthenticationRoutes(router *gin.Engine) {
	identityServiceUrl := "http://localhost:8080"

	router.POST("/login", handlers.ProxyHandler(identityServiceUrl))
	router.POST("/register", handlers.ProxyHandler(identityServiceUrl))
	router.POST("/logout", handlers.ProxyHandler(identityServiceUrl))

}
