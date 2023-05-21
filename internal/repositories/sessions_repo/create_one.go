package sessions_repo

import (
	"AuthService/configs"
	"AuthService/internal/models"
	"AuthService/internal/repositories/base_repo"
	"AuthService/internal/repositories/repositories_utils"
	"AuthService/internal/utils"
	"log"
	"time"
)

// CreateSession creates a new session for user in the database
func CreateSession(userID uint, token string, expiresAt *time.Time) (*models.UserSession, error) {
	log.Printf("User session creation")
	session := models.UserSession{
		UserID:    userID,
		Token:     token,
		CreatedAt: time.Now(),
		ExpiresAt: *expiresAt,
	}
	// prepare data for db insertion
	ignore_field := "id"
	fields, values := repositories_utils.GetFieldsAndValues(session, ignore_field)
	// insert the data
	id, err := base_repo.CreateOne(
		configs.MainSettings.SessionsTableName, fields, values)
	if err != nil {
		err = utils.UpdateExceptionMsg("failed to create session", err)
		return nil, err
	}
	// id edition
	session.ID = uint(id)
	log.Printf("Created session with ID = %d", session.ID)
	return &session, nil
}
