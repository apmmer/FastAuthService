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
	ServiceName             string
	UsersDBURL              string
	ProjectPath             string
	UsersTableName          string
	Debug                   string
	JwtSecret               string
	JwtRefreshSecret        string
	TokenLifeMinutes        int
	RefreshTokenLifeMinutes int
	SecureCookies           bool
	HttpOnlyCookies         bool
}

func GetSettings() *MainSettingsScheme {
	return &MainSettingsScheme{
		UsersDBURL: GetEnv(
			"USERS_DB_URL",
			"postgres://admin:admin@auth_service_postgres:5432/users_db?sslmode=disable",
		),
		Debug: GetEnv(
			"DEBUG",
			"true",
		),
		JwtSecret: GetEnv(
			"JWT_SECRET",
			"secret",
		),
		JwtRefreshSecret: GetEnv(
			"JWT_REFRESH_SECRET",
			"refresh",
		),
		ProjectPath:             "/app",
		UsersTableName:          "users",
		TokenLifeMinutes:        10,
		RefreshTokenLifeMinutes: 1000,
		ServiceName:             "AuthService",
		HttpOnlyCookies:         true,
		SecureCookies:           false,
	}
}

var MainSettings = GetSettings()
