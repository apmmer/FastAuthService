package users_repo_utils

import (
	"AuthService/internal/models"
	"time"
)

// converts a list of base_repo results to a list of User.
func ParseListToListOfUsers(results *[]map[string]interface{}) *[]models.User {
	var users []models.User
	for _, result := range *results {
		user := UserFromResult(&result)
		users = append(users, user)
	}
	return &users
}

// Parses base repository result (map[string]interface{}) to a new models.User object
func UserFromResult(resultPtr *map[string]interface{}) models.User {
	result := *resultPtr
	user := models.User{
		ID:         uint(result["id"].(int32)),
		ScreenName: result["screen_name"].(string),
		Email:      result["email"].(string),
		Password:   result["password"].(string),
		CreatedAt:  result["created_at"].(time.Time),
		UpdatedAt:  result["updated_at"].(time.Time),
	}
	// check if deleted_at is not null
	if result["deleted_at"] != nil {
		deletedAt := result["deleted_at"].(time.Time)
		user.DeletedAt = &deletedAt
	}
	if result["company_id"] != nil {
		companyId := result["company_id"].(int)
		user.CompanyId = &companyId
	}
	if result["rank"] != nil {
		rank := result["rank"].(int)
		user.Rank = &rank
	}
	return user
}
