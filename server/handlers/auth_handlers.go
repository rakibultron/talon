package handlers

import (
	// "fmt"

	"net/http"
	"time"

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

	// Get the current time
	currentTime := time.Now()

	// Format the time as RFC3339 (ISO 8601 format)
	timestamp := currentTime.Format(time.RFC3339)

	var body services.LoginUserRequest

	// Bind the JSON body to the request struct and handle validation errors
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	// Call the service layer to process the login
	user, jwt, err := services.LoginUser(body)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "User not found",
			"data":    nil,
			"error": gin.H{
				"code":    "USER_NOT_FOUND",
				"details": err.Error(),
			},
			"meta": gin.H{
				"timestamp": timestamp,
			},
		})
		return
	}

	// Set the authentication cookie
	c.SetCookie(
		"user_token", // Name of the cookie
		jwt,          // Value of the cookie
		31536000,     // Max age in seconds (1 year)
		"/",          // Path
		"localhost",  // Domain
		false,        // Secure (HTTPS only)
		true,         // HttpOnly
	)

	// Respond with the user data and JWT
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"user": user,
			"jwt":  jwt,
		},
	})
}
