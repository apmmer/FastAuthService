package sessions_repo

import (
	"AuthService/configs"
	"AuthService/internal/models"
	"AuthService/internal/repositories/base_repo"
	"AuthService/internal/utils"
	"time"
)

// CreateSession creates a new session for user in the database
func CreateSession(userID uint, token string, createdAt *time.Time) (*models.UserSession, error) {
	session := models.UserSession{
		UserID:    userID,
		Token:     token,
		CreatedAt: *createdAt,
	}
	// prepare data for db insertion
	ignore_field := "id"
	fields, values := utils.GetFieldsAndValues(session, ignore_field)
	// insert the data
	id, err := base_repo.CreateOne(
		configs.MainSettings.SessionsTableName, fields, values)
	if err != nil {
		err = utils.UpdateExceptionMsg("failed to create user", err)
		return nil, err
	}
	// id edition
	session.ID = uint(id)
	return &session, nil
}
