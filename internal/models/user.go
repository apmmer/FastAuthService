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
	ID         uint       `db:"id" gorm:"primaryKey"`
	ScreenName string     `db:"screen_name"`
	Email      string     `db:"email" gorm:"unique;not null"`
	Password   string     `db:"password"`
	CompanyId  *int       `db:"company_id"`
	Rank       *int       `db:"rank"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at"`
	DeletedAt  *time.Time `db:"deleted_at" sql:"index"`
}
