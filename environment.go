package main

import (
	"github.com/joho/godotenv"
	log "github.com/rs/zerolog/log"
	"os"
	"strings"
	"userservice-go/types"
)

// LoadEnvVars Loads environment variables
func LoadEnvVars() {
	runOn := strings.ToLower(os.Getenv("RUN_USER_SERVICE_ON"))
	var err error
	var envFileName = ".env" // also servers purpose for local environment

	switch runOn {
	case types.RUN_ON_DOCKER:
		fallthrough
	case types.RUN_ON_OPENSHIFT_LOCAL:
		envFileName = envFileName + "." + runOn
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
		types.KEYCLOAK_GET_BY_USERS = os.Getenv("KEYCLOAK_GET_BY_USERS")
		types.KEYCLOAK_TOKEN_PATH = os.ExpandEnv(types.KEYCLOAK_TOKEN_PATH)
		types.KEYCLOAK_GET_BY_USERNAME_PATH = os.ExpandEnv(types.KEYCLOAK_GET_BY_USERNAME_PATH)
		types.KEYCLOAK_GET_BY_USERS = os.ExpandEnv(types.KEYCLOAK_GET_BY_USERS)
		log.Debug().Msg("Loaded environment variables from: " + envFileName)
	}
}
