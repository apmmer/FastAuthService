package user_repo

import (
	"AuthService/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SaveUserToDB(user *models.User) (*models.User, error) {
	// Base code
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// AutoMigrate
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	// user creation
	result := db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
