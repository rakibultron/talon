package handlers

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rakibultron/talon/services"
)

// CreateUser handles creating a new user
func RegisterUser(c *gin.Context) {
	var user services.CreateUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser, err := services.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}
