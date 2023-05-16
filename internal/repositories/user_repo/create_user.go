package user_repo

import (
	"AuthService/configs"
	"AuthService/database"
	"AuthService/internal/models"
	"AuthService/internal/repositories/base_repo"
	"AuthService/internal/schemas"
	"fmt"
	"log"
	"time"
)

// CreateUser creates a new user in the database
func CreateUser(request schemas.CreateUserRequest) (*models.User, error) {

	user := models.User{
		ScreenName: request.ScreenName,
		Email:      request.Email,
		Password:   request.Password,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	log.Println("Trying to marshal user data")
	// Converting fields values to string
	fields, values, err := database.MarshalToDBString(user)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user data: %v", err)
	}

	// Calling the CreateOne function from base_repo
	log.Println("base_repo.CreateOne fields = " + fields + "	values = " + values)
	id, err := base_repo.CreateOne(configs.MainSettings.UsersTableName, fields, values)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}
	// Setting the generated ID to the user model
	user.ID = uint(id)

	return &user, nil
}
