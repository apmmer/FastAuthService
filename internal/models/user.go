package models

import "time"

type User struct {
	ID uint `gorm:"primaryKey"`

	Username string `json:"username"`

	Email string `json:"email" gorm:"unique;not null"`

	Password string `json:"password,omitempty"`

	CreatedAt time.Time

	UpdatedAt time.Time

	DeletedAt *time.Time `sql:"index"`
}
