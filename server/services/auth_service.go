package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rakibultron/talon/db"
	"github.com/rakibultron/talon/db/sqlcgen"
	"github.com/rakibultron/talon/utils"
)

// Define custom errors
var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

// CreateUserRequest defines the request payload for creating a user
type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" validate:"min=8,max=32,alphanum"`
}
type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" validate:"min=8,max=32,alphanum"`
}

// RegisterUser adds a new user to the database
func RegisterUser(request CreateUserRequest) (sqlcgen.CreateUserRow, error) {
	// Get the database connection (assuming db.GetDbCon() provides the connection)
	dbcon := db.GetDbCon()

	// Use the returned DBTX for queries
	queries := sqlcgen.New(dbcon)

	hash, _ := utils.HashPassword(request.Password)

	// Prepare the CreateUserParams with the data from the request
	createParams := sqlcgen.CreateUserParams{
		Name:     request.Name,
		Email:    request.Email,
		Password: hash,
	}

	// Perform the database operation to create a new user
	user, err := queries.CreateUser(context.Background(), createParams)
	if err != nil {
		return sqlcgen.CreateUserRow{}, err
	}

	// Return the created user and nil error if successful
	return user, nil
}

// LoginUser handles user login
func LoginUser(request LoginUserRequest) (sqlcgen.GetUserByEmailRow, string, error) {
	dbcon := db.GetDbCon()
	email, password := request.Email, request.Password

	// Use the returned DBTX for queries
	queries := sqlcgen.New(dbcon)

	// Fetch user by email
	user, err := queries.GetUserByEmail(context.Background(), email)
	if user.Email == "" {
		fmt.Println("user not found")

		return sqlcgen.GetUserByEmailRow{}, "", ErrUserNotFound
	}

	if err != nil {
		// If there's an error fetching the user, return empty user, empty string, and the error
		return sqlcgen.GetUserByEmailRow{}, "", fmt.Errorf("failed to fetch user: %v", err)
	}

	// Verify password
	match := utils.VerifyPassword(password, user.Password)
	if !match {
		// Return a custom error message for invalid credentials
		return sqlcgen.GetUserByEmailRow{}, "", ErrInvalidCredentials
	}

	// Generate JWT if password matches
	jwtSecrate := os.Getenv("JWT_SECRATE")
	key := []byte(jwtSecrate)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Name,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign the token
	tokenString, err := token.SignedString(key)
	if err != nil {
		// If there's an error signing the token, return the user, an empty token, and the error
		return user, "", fmt.Errorf("failed to sign token: %v", err)
	}

	// Return the user, generated token, and nil error if everything succeeds
	return user, tokenString, nil
}
