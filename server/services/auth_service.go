package services

import (
	"context"

	"github.com/rakibultron/talon/db"
	"github.com/rakibultron/talon/db/sqlcgen"
)

// CreateUserRequest defines the request payload for creating a user
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
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
