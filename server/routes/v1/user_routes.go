package v1

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/rakibultron/talon/db"
	"github.com/rakibultron/talon/db/sqlcgen"
)

func UserRoutes(router *gin.RouterGroup) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/", func(c *gin.Context) {

			dbcon := db.GetDbCon()

			// Use the returned DBTX for queries
			queries := sqlcgen.New(dbcon)
			users, _ := queries.GetAllUsers(context.Background())

			c.JSON(200, gin.H{
				"message": users,
			})
		})

	}
}
