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
	SessionsTableName       string
	Debug                   string
	ApiKey                  string
	JwtSecret               string
	JwtRefreshSecret        string
	SessionSecret           string
	TokenLifeMinutes        int
	RefreshTokenLifeMinutes int
	SecureCookies           bool
	HttpOnlyCookies         bool
	CertKeyLocation         string
	СertFileLocation        string
	SwaggerUrl              string
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
		ApiKey: GetEnv(
			"API_KEY",
			"secret",
		),
		JwtSecret: GetEnv(
			"JWT_SECRET",
			"secret",
		),
		JwtRefreshSecret: GetEnv(
			"JWT_REFRESH_SECRET",
			"refresh",
		),
		// SessionSecret MUST CONTAIN 16, 24 or 32 bytes!
		SessionSecret: GetEnv(
			"SESSION_SECRET",
			"sesssesssesssess",
		),
		CertKeyLocation: GetEnv(
			"CERTIFICATE_KEY_LOC",
			"/app/certificates/dev_private_key.pem",
		),
		СertFileLocation: GetEnv(
			"CERTIFICATE_FILE_LOC",
			"/app/certificates/dev_certificate.pem",
		),
		ProjectPath:             "/app",
		UsersTableName:          "users",
		SessionsTableName:       "user_sessions",
		TokenLifeMinutes:        10,
		RefreshTokenLifeMinutes: 1000,
		ServiceName:             "AuthService",
		HttpOnlyCookies:         true,
		SecureCookies:           false,
		SwaggerUrl:              "/swagger",
	}
}

var MainSettings = GetSettings()
