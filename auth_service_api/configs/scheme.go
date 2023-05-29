package configs

type MainSettingsScheme struct {
	ServiceName             string // Name of this service
	UsersDBURL              string // URL of the users database
	ProjectPath             string // Path to the project root within the container
	UsersTableName          string // Name of the users table in the database
	SessionsTableName       string // Name of the sessions table in the database
	Debug                   string // Flag for enabling debug mode
	ApiKey                  string // API key for service access
	JwtSecret               string // Secret for signing JWT tokens
	JwtRefreshSecret        string // Secret for signing JWT refresh tokens
	SessionSecret           string // Secret for encrypting session data, must be 16, 24 or 32 bytes string
	TokenLifeMinutes        int    // Lifetime of JWT token in minutes
	RefreshTokenLifeMinutes int    // Lifetime of JWT refresh token in minutes
	SecureCookies           bool   // Flag for enabling secure flag in cookies
	HttpOnlyCookies         bool   // Flag for enabling HttpOnly flag in cookies
	CertKeyLocation         string // Path to SSL certificate private key
	СertFileLocation        string // Path to SSL certificate file
	SwaggerUrl              string // URL of Swagger documentation
}
