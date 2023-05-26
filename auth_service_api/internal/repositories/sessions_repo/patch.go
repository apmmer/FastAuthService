package sessions_repo

import (
	"auth_service_api/configs"
	"auth_service_api/internal/exceptions"
	"auth_service_api/internal/general_utils"
	"auth_service_api/internal/models"
	"auth_service_api/internal/repositories/base_repo"
	"auth_service_api/internal/repositories/repositories_utils"
	"fmt"
	"log"
	"time"
)

// Updates sessions records in db according provided data and filters.
// Returns:
//
//	updated items (*[]models.UserSession) - a pointer to a list of updated (and parsed) results
//	error
func UpdateSessions(filters *map[string]interface{}, updateData *map[string]interface{}) (*[]models.UserSession, error) {
	// validate filters
	err := repositories_utils.ValidateMapFields(filters, models.UserSession{})
	if err != nil {
		return nil, general_utils.UpdateException("failed to validate filters", err)
	}
	// validate updateData
	err = repositories_utils.ValidateMapFields(updateData, models.UserSession{})
	if err != nil {
		return nil, general_utils.UpdateException("failed to validate filters", err)
	}
	// call update returning updated items (type *[]map[string]interface{})
	results, err := base_repo.Update(
		configs.MainSettings.SessionsTableName,
		filters,
		updateData,
	)
	if err != nil {
		fmt.Println("failed to update sessions")
		return nil, general_utils.UpdateException("failed to update sessions", err)
	}
	if len(*results) == 0 {
		fmt.Println("sessions not found")
		return nil, exceptions.MakeNotFoundError("valid sessions not found.")
	}

	updatedItems := parseListToListOfSessions(results)
	return &updatedItems, nil
}

// Performs single request to update user_sessions and check related user by the way
func OptimizedUpdateWithUserChecking(expires_at *time.Time, token string) (*[]models.UserSession, error) {
	sqlQuery := `
		UPDATE user_sessions
			SET expires_at = $1
			WHERE
				token = $2
				AND deleted_at IS NULL
				AND expires_at > NOW()
				AND user_id IN (
					SELECT id FROM users WHERE deleted_at IS NULL
				)
		RETURNING *;
	`
	args := []interface{}{
		*expires_at,
		token,
	}
	results, err := base_repo.ExecuteRowParseList(sqlQuery, args)
	if err != nil {
		err = general_utils.UpdateException("repo: failed to update sessions", err)
		return nil, err
	}
	if len(*results) == 0 {
		log.Println("sessions not found")
		return nil, exceptions.MakeNotFoundError("sessions not found. ")
	}
	updatedItems := parseListToListOfSessions(results)
	return &updatedItems, nil
}

// converts a list of base_repo results to a list of UserSession.
func parseListToListOfSessions(results *[]map[string]interface{}) []models.UserSession {
	var updatedItems []models.UserSession
	for _, result := range *results {
		updatedItem := models.UserSession{
			ID:        uint(result["id"].(int32)),
			UserID:    uint(result["user_id"].(int32)),
			Token:     result["token"].(string),
			CreatedAt: result["created_at"].(time.Time),
			ExpiresAt: result["expires_at"].(time.Time),
		}
		// check if deleted_at is not null
		if result["deleted_at"] != nil {
			deletedAt := result["deleted_at"].(time.Time)
			updatedItem.DeletedAt = &deletedAt
		}
		updatedItems = append(updatedItems, updatedItem)
	}
	return updatedItems
}
