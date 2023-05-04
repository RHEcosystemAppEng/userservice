package types

var (
	USERNAME_PARAM                = "username"
	KEYCLOAK_BACKEND_URL          = "http://localhost:8080/"
	ADMIN_USER                    = "admin"
	ADMIN_PASSWORD                = "admin"
	GRANT_TYPE                    = "password"
	CLIENT_ID                     = "admin-cli"
	KEYCLOAK_MASTER_REALM_TOKEN   = ""
	KEYCLOAK_MASTER_TOKEN_PATH    = "realms/master/protocol/openid-connect/token"
	KEYCLOAK_GET_BY_USERNAME_PATH = "admin/realms/master/users?username="

	ERR_NIL_HTTP_CLIENT_OR_REQUEST = "nil http request or http client object"
)
