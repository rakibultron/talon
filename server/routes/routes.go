package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/rakibultron/talon/routes/v1"
)

func Routes(router *gin.Engine) {
	api := router.Group("/v1/api")

	v1.AuthRoutes(api)

}
