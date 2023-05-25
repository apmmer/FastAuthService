package configs

import (
	"os"
)

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func newMainSettings() *MainSettingsScheme {
	// Loading values from the environment or using default
	return &MainSettingsScheme{
		UsersDBURL:              getEnv("USERS_DB_URL", "postgres://admin:admin@auth_service_postgres:5432/users_db?sslmode=disable"),
		Debug:                   getEnv("DEBUG", "true"),
		ApiKey:                  getEnv("API_KEY", "secret"),
		JwtSecret:               getEnv("JWT_SECRET", "secret"),
		JwtRefreshSecret:        getEnv("JWT_REFRESH_SECRET", "refresh"),
		SessionSecret:           getEnv("SESSION_SECRET", "sesssesssesssess"),
		CertKeyLocation:         getEnv("CERTIFICATE_KEY_LOC", "/app/certificates/dev_private_key.pem"),
		СertFileLocation:        getEnv("CERTIFICATE_FILE_LOC", "/app/certificates/dev_certificate.pem"),
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

// create global main settings object
var MainSettings = newMainSettings()
