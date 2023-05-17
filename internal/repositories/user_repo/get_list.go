package user_repo

import (
	"AuthService/configs"
	"AuthService/internal/models"
	"AuthService/internal/repositories/base_repo"
	"AuthService/internal/schemas"
	"AuthService/internal/utils"
)

func GetManyUsers(params schemas.ListParams) ([]models.User, error) {
	var users []models.User

	// Use the ParseSorting function to extract the sorting field and direction
	sortingField, sortingDirection := utils.ParseSorting(params.Sorting)

	// Prepare parameters for base_repo.GetMany function
	limit := params.Limit
	if limit != nil {
		*limit = int(*limit)
	}

	offset := params.Offset
	if offset != nil {
		*offset = int(*offset)
	}

	// Call the base_repo.GetMany function
	results, err := base_repo.GetMany(
		configs.MainSettings.UsersTableName,
		limit,
		offset,
		sortingField,
		sortingDirection,
		nil,
	)
	if err != nil {
		return nil, err
	}

	// Convert each map result to a User
	for _, result := range results {
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

		users = append(users, user)
	}

	return users, nil
}
