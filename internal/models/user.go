package models

import "time"

// User godoc
// User represents user data
// @Schema
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
	ID         uint       `json:"id" db:"id" gorm:"primaryKey"`
	ScreenName string     `json:"screen_name" db:"screen_name"`
	Email      string     `json:"email" db:"email" gorm:"unique;not null"`
	Password   string     `json:"password" db:"password"`
	CompanyId  *int       `json:"company_id" db:"company_id"`
	Rank       *int       `json:"rank" db:"rank"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at" db:"deleted_at" sql:"index"`
}
