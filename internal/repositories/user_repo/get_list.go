package user_repo

import (
	"AuthService/configs"
	"AuthService/internal/models"
	"AuthService/internal/repositories/base_repo"
	"AuthService/internal/repositories/repositories_utils"
	"AuthService/internal/schemas"
	"fmt"
)

func GetList(params schemas.ListParams) (*[]models.User, error) {
	fmt.Println("Called user_repo.GetManyUsers")

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

	// Get records from db
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
	users := ParseListToListOfUsers(&results)
	return users, nil
}
