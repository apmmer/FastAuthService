package models

import "time"

// UserSession represents a user session.
// @Schema
type UserSession struct {
	// ID represents the session ID.
	// @ID id
	// @Required
	// @Property example: 1
	ID uint `json:"id" db:"id" gorm:"primaryKey"`

	// UserID represents the ID of the associated user.
	// @Required
	// @Property example: 123
	UserID uint `json:"user_id" db:"user_id"`

	// Token represents the session token.
	// @Required
	// @Property example: "abc123"
	Token string `json:"token" db:"token"`

	// CreatedAt represents the creation timestamp of the session.
	// @Property format: date-time
	// @Property example: "2023-05-19T12:34:56Z"
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// ExpiresAt represents the expiration time of the session.
	// @Property format: date-time
	// @Property example: "2023-05-19T12:34:56Z"
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`

	// DeletedAt represents the timestamp when session was closed.
	// @Property deleted_at string "DeletedAt" format(date-time)
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at" sql:"index"`
}
