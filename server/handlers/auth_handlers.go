package handlers

import (
	// "fmt"
	"fmt"
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

// LoginUser handles user login
func LoginUser(c *gin.Context) {

	var body services.LoginUserRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, jwt, _ := services.LoginUser(body)

	fmt.Println(jwt)

	c.SetCookie(
		"user_token", // Name of the cookie
		jwt,          // Value of the cookie
		31536000,     // Max age in seconds (1 year)
		"/",          // Path
		"localhost",  // Domain
		false,        // Secure (HTTPS only)
		true,         // HttpOnly
	)
	c.JSON(http.StatusOK, gin.H{
		"user": user,
		"jwt":  jwt,
	})

}
