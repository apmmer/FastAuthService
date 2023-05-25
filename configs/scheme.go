package configs

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
