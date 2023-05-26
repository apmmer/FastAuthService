package users_repo

import (
	"auth_service_api/configs"
	"auth_service_api/internal/general_utils"
	"auth_service_api/internal/models"
	"auth_service_api/internal/repositories/base_repo"
	"auth_service_api/internal/repositories/repositories_utils"
	"auth_service_api/internal/repositories/users_repo/users_repo_utils"
	"fmt"
)

// Retrieves a data from DB according filters and returns created User object.
func GetUser(filters *map[string]interface{}) (*models.User, error) {
	err := repositories_utils.ValidateMapFields(filters, models.User{})
	if err != nil {
		return nil, general_utils.UpdateException(
			fmt.Sprintf("failed to validate filters: %v", err),
			err,
		)
	}
	result, err := base_repo.GetOne(
		configs.MainSettings.UsersTableName,
		filters,
	)
	if err != nil {
		return nil, err
	}
	user := users_repo_utils.UserFromResult(&result)
	return &user, nil
}

// Retrieves not deleted user by ID using GetUser function.
// This variation of GetUser() can help with drying and simplifying a code.
func GetActiveUserById(userId int) (*models.User, error) {
	filters := make(map[string]interface{})
	filters["id"] = userId
	filters["deleted_at"] = nil
	user, err := GetUser(&filters)
	return user, err
}
