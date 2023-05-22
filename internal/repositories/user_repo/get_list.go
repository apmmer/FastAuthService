package user_repo

import (
	"AuthService/configs"
	"AuthService/internal/models"
	"AuthService/internal/repositories/base_repo"
	"AuthService/internal/schemas"
	"AuthService/internal/utils"
	"fmt"
)

func GetManyUsers(params schemas.ListParams) ([]models.User, error) {
	fmt.Println("Called GetManyUsers")
	var users []models.User

	// Use the ParseSorting function to extract the sorting field and direction
	sortingField, sortingDirection := utils.ParseSorting(params.Sorting)

	// Call the base_repo.GetMany function
	results, err := base_repo.GetMany(
		configs.MainSettings.UsersTableName,
		params.Limit,
		params.Offset,
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
