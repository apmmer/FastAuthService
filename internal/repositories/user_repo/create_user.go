package user_repo

import (
	"AuthService/configs"
	"AuthService/internal/general_utils"
	"AuthService/internal/models"
	"AuthService/internal/repositories/base_repo"
	"AuthService/internal/repositories/repositories_utils"
	"AuthService/internal/schemas"
	"fmt"
	"log"
	"time"
)

// CreateUser creates a new user in the database
func CreateUser(request schemas.CreateUserRequest) (*models.User, error) {

	hashedPassword, err := general_utils.GetHash(request.Password)
	if err != nil {
		return nil, fmt.Errorf("could not encrypt password: %v", err)
	}
	user := models.User{
		ScreenName: request.ScreenName,
		Email:      request.Email,
		Password:   hashedPassword,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// prepare data for db insertion
	ignore_field := "id"
	fields, values := repositories_utils.GetFieldsAndValues(user, ignore_field)

	log.Println("Calling base_repo.CreateOne")
	// Calling the CreateOne function from base_repo
	id, err := base_repo.CreateOne(
		configs.MainSettings.UsersTableName, fields, values)
	if err != nil {
		err = general_utils.UpdateExceptionMsg("failed to create user", err)
		return nil, err
	}
	// Setting the generated ID to the user model
	user.ID = uint(id)
	return &user, nil
}
