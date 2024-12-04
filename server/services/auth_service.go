package services

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/rakibultron/talon/db"
	"github.com/rakibultron/talon/db/sqlcgen"
	"github.com/rakibultron/talon/utils"
)

// CreateUserRequest defines the request payload for creating a user
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}
type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" validate:"min=8,max=32,alphanum"`
	Id       string `json:"id" binding:"required"`
}

// RegisterUser adds a new user to the database
func RegisterUser(request CreateUserRequest) (sqlcgen.CreateUserRow, error) {
	// Get the database connection (assuming db.GetDbCon() provides the connection)
	dbcon := db.GetDbCon()

	// Use the returned DBTX for queries
	queries := sqlcgen.New(dbcon)

	// Prepare the CreateUserParams with the data from the request
	createParams := sqlcgen.CreateUserParams{
		Name:  request.Name,
		Email: request.Email,
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
func LoginUser(request LoginUserRequest) (sqlcgen.GetUserByIDRow, string, error) {

	dbcon := db.GetDbCon()

	id, _, _ := request.Id, request.Email, request.Password
	// Use the returned DBTX for queries
	queries := sqlcgen.New(dbcon)

	uuidValue, _ := uuid.Parse(id)

	// Perform the database operation to create a new user
	user, _ := queries.GetUserByID(context.Background(), utils.ConvertUUIDToPGXUUID(uuidValue))

	jwtSecrate := os.Getenv("JWT_SECRATE")

	key := []byte(jwtSecrate)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": user.Name,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(key)
	fmt.Println(tokenString)
	if err != nil {
		return user, tokenString, err
	}

	// Return the created user and nil error if successful
	return user, tokenString, nil
}
