package models

import "time"

// User godoc
// User represents user data
// @Schema
// @ID User
// @Required username email password
// @Property id integer "ID"
// @Property username string "Username" minLength(1)
// @Property email string "Email" format(email)
// @Property password string "Password" minLength(1)
type User struct {
	ID uint `gorm:"primaryKey" json:"-"`

	Username string `json:"username"`

	Email string `json:"email" gorm:"unique;not null"`

	Password string `json:"password,omitempty"`

	CreatedAt time.Time `json:"-"`

	UpdatedAt time.Time `json:"-"`

	DeletedAt *time.Time `sql:"index" json:"-"`
}
