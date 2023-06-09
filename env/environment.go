package env

import (
	"github.com/joho/godotenv"
	log "github.com/rs/zerolog/log"
	"os"
	"strings"
	"userservice-go/types"
)

// LoadEnvVars Loads environment variables
func LoadEnvVars(relativePath string) {
	runOn := strings.ToLower(os.Getenv("RUN_USER_SERVICE_ON"))
	var err error
	var envFileName = ".env" // also servers purpose for local environment

	if runOn == "" {
		envFileName = relativePath + envFileName
	} else {
		envFileName = relativePath + envFileName + "." + runOn
	}

	err = godotenv.Load(envFileName)
	if err != nil {
		log.Error().Msg("Error loading .env file:" + envFileName + ", will fallback to default environment variables. " + err.Error())
	} else {
		types.USER_SERVICE_PORT = os.Getenv("USER_SERVICE_PORT")
		types.KEYCLOAK_BACKEND_URL = os.Getenv("KEYCLOAK_BACKEND_URL")
		types.ADMIN_USER = os.Getenv("ADMIN_USER")
		types.ADMIN_PASSWORD = os.Getenv("ADMIN_PASSWORD")
		types.GRANT_TYPE = os.Getenv("GRANT_TYPE")
		types.CLIENT_ID = os.Getenv("CLIENT_ID")
		types.KEYCLOAK_REALM = os.Getenv("KEYCLOAK_REALM")
		types.KEYCLOAK_MASTER_REALM_TOKEN = os.Getenv("KEYCLOAK_MASTER_REALM_TOKEN")
		types.KEYCLOAK_TOKEN_PATH = os.Getenv("KEYCLOAK_TOKEN_PATH")
		types.KEYCLOAK_GET_BY_USERNAME_PATH = os.Getenv("KEYCLOAK_GET_BY_USERNAME_PATH")
		types.KEYCLOAK_TOKEN_PATH = os.ExpandEnv(types.KEYCLOAK_TOKEN_PATH)
		types.KEYCLOAK_GET_BY_USERNAME_PATH = os.ExpandEnv(types.KEYCLOAK_GET_BY_USERNAME_PATH)
		types.KEYCLOAK_USERS_RESOURCE_URI = os.ExpandEnv(types.KEYCLOAK_USERS_RESOURCE_URI)
		types.DISABLE_KEYCLOAK_CERT_VERIFICATION = os.Getenv("DISABLE_KEYCLOAK_CERT_VERIFICATION")
		types.USER_SERVICE_TLS_CRT_PATH = os.Getenv("USER_SERVICE_TLS_CRT_PATH")
		types.USER_SERVICE_TLS_KEY_PATH = os.Getenv("USER_SERVICE_TLS_KEY_PATH")
		log.Debug().Msg("Loaded environment variables from: " + envFileName)
	}
}
