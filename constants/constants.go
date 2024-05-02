package constants

const (
	ERROR_ACCESS_CREDENTIALS = "access credentials error"
	ERROR_TOKEN              = "token validation error"
	ERROR_GET_TOKEN          = "error searching for token"
	ERROR_LEN_PASSWORD       = "password length too short"
	ERROR_UNHAUTORIZED       = "you do not have privileges to perform this action"
)

const (
	PRIVILEGES_USER = iota
	PRIVILEGES_MODERATOR
	PRIVILEGES_ADMINISTRATOR
	PRIVILEGES_SUPER_ADMIN
)

const (
	LEN_PASSWORD     = 8
	LEN_ACCOUNT_NAME = 3
)

const (
	EXPIRATION_TOKEN = 24
)

const (
	ACCOUNT_SESSION = "Authorization"
)
