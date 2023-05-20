package sessions_repo

import (
	"AuthService/configs"
	"AuthService/internal/exceptions"
	"AuthService/internal/models"
	"AuthService/internal/repositories/base_repo"
	"AuthService/internal/utils"
	"fmt"
	"time"
)

func UpdateSessions(filters *map[string]interface{}, updateData *map[string]interface{}) (*[]models.UserSession, error) {
	// validate filters
	err := utils.ValidateMapFields(filters, models.UserSession{})
	if err != nil {
		return nil, utils.UpdateExceptionMsg("failed to validate filters", err)
	}
	// validate updateData
	err = utils.ValidateMapFields(updateData, models.UserSession{})
	if err != nil {
		return nil, utils.UpdateExceptionMsg("failed to validate filters", err)
	}
	// call update returning updated items (type *[]map[string]interface{})
	results, err := base_repo.Update(
		configs.MainSettings.SessionsTableName,
		filters,
		updateData,
	)
	if err != nil {
		fmt.Println("failed to update sessions")
		return nil, utils.UpdateExceptionMsg("failed to update sessions", err)
	}
	if len(*results) == 0 {
		fmt.Println("sessions not found")
		return nil, &exceptions.ErrNotFound{Message: "sessions not found"}
	}

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

	return &updatedItems, nil
}
