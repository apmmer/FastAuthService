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

	// CertKeyLocation: Path to the SSL certificate's private key within the container.
	//
	// CAUTION: Take special care when transferring the certificate file and key into the container.
	// - The certificate and key files should always be kept secure and confidential. They should never be exposed or checked into version control.
	// - Ideally, these files should be provided to the container at runtime using secure orchestration tools, such as Kubernetes Secrets or equivalent.
	// - If using Docker, do NOT use `COPY` in the Dockerfile to transfer these files into the image. This could lead to exposure if the image is pushed to a public repository.
	// - Regularly rotate and change your keys and certificates to reduce the impact if they are compromised.
	// - Always ensure that the certificate is valid and hasn't expired to avoid any service interruption.
	CertKeyLocation  string
	СertFileLocation string // Path to SSL certificate file
	SwaggerUrl       string // URL of Swagger documentation
}
