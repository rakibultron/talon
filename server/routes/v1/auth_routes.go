package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rakibultron/talon/handlers"
)

func AuthRoutes(router *gin.RouterGroup) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", handlers.RegisterUser)

	}
}
