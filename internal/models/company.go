package models

import "time"

type Company struct {
	ID        uint       `db:"id" gorm:"primaryKey"`
	Name      string     `db:"name"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" sql:"index"`
}
