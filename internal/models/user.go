package models

import "time"

// User godoc
// User represents user data
// @Schema
// @ID User
// @Required screen_name email password company_id rank
// @Property id integer "ID"
// @Property screen_name string "ScreenName" minLength(4)
// @Property email string "Email" format(email)
// @Property password string "Password" minLength(7)
// @Property company_id integer "CompanyID" format(int64)
// @Property rank integer "Rank" format(int64)
// @Property created_at string "CreatedAt" format(date-time)
// @Property updated_at string "UpdatedAt" format(date-time)
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
