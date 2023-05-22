package user_repo

import (
	"AuthService/configs"
	"AuthService/internal/exceptions"
	"AuthService/internal/models"
	"AuthService/internal/repositories/base_repo"
	"AuthService/internal/utils"
	"fmt"
)

// Retrieves a data from DB according filters and returns created User object.
func GetUser(filters *map[string]interface{}) (*models.User, error) {
	err := utils.ValidateMapFields(filters, models.User{})
	if err != nil {
		return nil, &exceptions.ErrInvalidEntity{Message: fmt.Sprintf("failed to validate filters: %v", err)}
	}
	result, err := base_repo.GetOne(
		configs.MainSettings.UsersTableName,
		filters,
	)
	if err != nil {
		return nil, err
	}
	user := models.User{
		ID:         uint(result["id"].(int32)),
		ScreenName: result["screen_name"].(string),
		Email:      result["email"].(string),
		Password:   result["password"].(string),
	}
	if result["company_id"] != nil {
		companyId := result["company_id"].(int)
		user.CompanyId = &companyId
	}

	if result["rank"] != nil {
		rank := result["rank"].(int)
		user.Rank = &rank
	}
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
