package types

var (
	// Default environment variables, values will be replaced with .env file
	KEYCLOAK_BACKEND_URL          = "http://localhost:8080/"
	ADMIN_USER                    = "admin"
	ADMIN_PASSWORD                = "admin"
	GRANT_TYPE                    = "password"
	CLIENT_ID                     = "admin-cli"
	KEYCLOAK_REALM                = "master"
	KEYCLOAK_MASTER_REALM_TOKEN   = ""
	KEYCLOAK_TOKEN_PATH           = "realms/master/protocol/openid-connect/token"
	KEYCLOAK_GET_BY_USERNAME_PATH = "admin/realms/master/users?username="
	KEYCLOAK_GET_BY_USERS         = "admin/realms/master/users"
	USER_SERVICE_PORT             = ":8000"

	USERNAME_PARAM   = "username"
	ORG_ID_PARAM     = "org_id"
	USER_NAMES_PARAM = "usernames"
	EMAILS_PARAM     = "emails"
	USER_IDS_PARAM   = "user_ids"

	ORDER_BY_EMAIL    = "email"
	ORDER_BY_USERNAME = "username"
	ORDER_BY_MODIFIED = "modified"
	ORDER_BY_CREATED  = "created"

	ORDER_BY_DIR_ASC  = "asc"
	ORDER_BY_DIR_DESC = "desc"

	// http codes
	HTTP_CODE_BAD_REQUEST = "400"

	// Error messages
	ERR_NIL_HTTP_CLIENT_OR_REQUEST = "nil http request or http client object"

	RUN_ON_LOCAL           = "local"
	RUN_ON_OPENSHIFT_LOCAL = "openshift.local"
	RUN_ON_DOCKER          = "docker"
)
