package v1

import (
	"github.com/gin-gonic/gin"

)

func UserRoutes(router *gin.RouterGroup) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "users",
			})
		})

	}
}
