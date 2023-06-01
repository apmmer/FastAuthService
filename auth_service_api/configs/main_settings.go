package configs

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// Reads the value of an environment variable identified by
// the provided key.
// If the variable does not exist, it returns a fallback value.
func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

// getCurrentDirectory returns the full directory path of the current file and its parent directory.
func getCurrentDirectory() (string, string, error) {
	// Get the path to the currently running file
	_, filePath, _, ok := runtime.Caller(1)
	if !ok {
		return "", "", errors.New("unable to get the current file path")
	}

	// Get the full directory path
	dirPath := filepath.Dir(filePath)

	// Get the parent directory path
	parentDirPath := filepath.Dir(dirPath)

	return dirPath, parentDirPath, nil
}

// Creates a new instance of MainSettingsScheme, which contains various
// settings for the main configuration.
func NewMainSettings() *MainSettingsScheme {
	dir, projectPath, err := getCurrentDirectory()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Main config directory: %s\n", dir)
	log.Printf("Project directory: %s\n", projectPath)
	// Loading values from the environment or using default
	return &MainSettingsScheme{
		UsersDBURL:       getEnv("USERS_DB_URL", "postgres://admin:admin@auth_service_postgres:5432/users_db?sslmode=disable"),
		Debug:            getEnv("DEBUG", "true"),
		ApiKey:           getEnv("API_KEY", "secret"),
		JwtSecret:        getEnv("JWT_SECRET", "secret"),
		JwtRefreshSecret: getEnv("JWT_REFRESH_SECRET", "refresh"),
		SessionSecret:    getEnv("SESSION_SECRET", "sesssesssesssess"),
		CertKeyLocation: filepath.Join(
			projectPath, getEnv("CERTIFICATE_KEY_LOC", "/certificates/dev_private_key.pem"),
		),
		СertFileLocation: filepath.Join(
			projectPath, getEnv("CERTIFICATE_FILE_LOC", "/certificates/dev_certificate.pem"),
		),
		ProjectPath:             projectPath,
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
