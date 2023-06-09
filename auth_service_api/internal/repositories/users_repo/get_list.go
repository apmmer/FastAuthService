package users_repo

import (
	"auth_service_api/configs"
	"auth_service_api/internal/models"
	"auth_service_api/internal/repositories/base_repo"
	"auth_service_api/internal/repositories/repositories_utils"
	"auth_service_api/internal/repositories/users_repo/users_repo_utils"
	"auth_service_api/internal/schemas"
	"fmt"
)

// Retrieves a list of records of models.User and serializes a result
// Returns:
//
//	users (*[]models.User) - a list of retrieved users (models.User) according filters.
func GetList(params schemas.ListParams) (*[]models.User, error) {
	fmt.Println("Called users_repo.GetList")

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
	users := users_repo_utils.ParseListToListOfUsers(&results)
	return users, nil
}
