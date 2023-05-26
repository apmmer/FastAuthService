package models

import "time"

// User godoc
// User represents user data
// @Schema(name="User")
// @ID User
// @Required id screen_name email password
// @Optional company_id rank created_at updated_at deleted_at
// @Property id integer "ID" example(1)
// @Property screen_name string "ScreenName" minLength(4) example("JohnDoe")
// @Property email string "Email" format(email) example("johndoe@example.com")
// @Property password string "Password" minLength(7) example("12345678")
// @Property company_id integer "CompanyID" format(int64) example(1)
// @Property rank integer "Rank" format(int64) example(1)
// @Property created_at string "CreatedAt" format(date-time) example("2023-05-17T18:57:39Z")
// @Property updated_at string "UpdatedAt" format(date-time) example("2023-05-18T18:57:39Z")
// @Property deleted_at string "DeletedAt" format(date-time)
type User struct {
	// The auto-created unique identifier for the user.
	// @Example 123
	ID uint `json:"id" db:"id" gorm:"primaryKey"`

	// The desired screenname for the user.
	// This field is required and must contain at least 4 characters.
	// @Example jonhDoe123
	ScreenName string `json:"screen_name" db:"screen_name"`

	// The email address of a user.
	// This field is unique and not-null.
	// @Format email
	// @Example john@example.com
	Email string `json:"email" db:"email" gorm:"unique;not null"`

	// A secret that only the creator of the record knows
	// @Example 1234567
	Password string `json:"password" db:"password"`

	// A link to the user's company ID, nullable.
	// @Example 1
	CompanyId *int `json:"company_id" db:"company_id"`

	// Status of a user, nullable.
	// @Example 1
	Rank *int `json:"rank" db:"rank"`

	// When record was created.
	// @Format format(date-time)
	// @Example 2023-05-18T18:57:39Z
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// When record was updated last time.
	// @Format format(date-time)
	// @Example 2023-05-18T18:57:39Z
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	// When record was attempted to delete, nullable.
	// @Format format(date-time)
	// @Example 2023-05-18T18:57:39Z
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at" sql:"index"`
}
