package configs

import (
	"os"
)

func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

type MainSettingsScheme struct {
	UsersDBURL     string
	ProjectPath    string
	UsersTableName string
	Debug          string
}

func GetSettings() *MainSettingsScheme {
	return &MainSettingsScheme{
		UsersDBURL: GetEnv(
			"USERS_DB_URL",
			"postgres://admin:admin@auth_service_postgres:5432/users_db?sslmode=disable",
		),
		ProjectPath:    "/app",
		UsersTableName: "users",
		Debug: GetEnv(
			"DEBUG",
			"true",
		),
	}
}

var MainSettings = GetSettings()
