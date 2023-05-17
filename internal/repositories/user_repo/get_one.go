package user_repo

import (
	"AuthService/configs"
	"AuthService/internal/models"
	"AuthService/internal/repositories/base_repo"
	"AuthService/internal/utils"
	"fmt"
)

// Checks if filters are correct for models.User
func validateUserFilters(filters *map[string]interface{}) error {
	model_fields := utils.GetModelFields(models.User{})
	for key := range *filters {
		field_found := false
		for _, fieldname := range model_fields {
			if key == fieldname {
				field_found = true
				break
			}
		}
		if !field_found {
			return fmt.Errorf("field %s was not found in User model", key)
		}
	}
	return nil
}

func GetUser(filters *map[string]interface{}) (*models.User, error) {
	err := validateUserFilters(filters)
	if err != nil {
		return nil, fmt.Errorf("failed to validate filters: %v", err)
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
