package models

import "time"

// Company godoc
// Company represents company data
// @Schema
// @ID Company
// @Required id name
// @Optional created_at updated_at deleted_at
// @Property id integer "ID" example(1)
// @Property name string "Name" example("ProductXCompany")
// @Property created_at string "CreatedAt" format(date-time) example("2023-05-17T18:57:39Z")
// @Property updated_at string "UpdatedAt" format(date-time) example("2023-05-18T18:57:39Z")
// @Property deleted_at string "DeletedAt" format(date-time)
type Company struct {
	ID        uint       `json:"id" db:"id" gorm:"primaryKey"`
	Name      string     `json:"name" db:"name"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at" sql:"index"`
}
