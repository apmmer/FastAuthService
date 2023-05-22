package user_repo

import (
	"AuthService/configs"
	"AuthService/internal/models"
	"AuthService/internal/repositories/base_repo"
	"AuthService/internal/repositories/repositories_utils"
	"AuthService/internal/schemas"
	"fmt"
)

func GetManyUsers(params schemas.ListParams) ([]models.User, error) {
	fmt.Println("Called GetManyUsers")
	var users []models.User

	// ensure SortingField in models field
	if params.SortingField != nil {
		sortingField := *params.SortingField
		err := repositories_utils.FieldInModelFields(
			sortingField,
			repositories_utils.GetModelFields(models.User{}),
		)
		if err != nil {
			return nil, err
		}
	}

	// Call the base_repo.GetMany function
	results, err := base_repo.GetMany(
		configs.MainSettings.UsersTableName,
		params.Limit,
		params.Offset,
		params.SortingField,
		params.SortingDirection,
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
