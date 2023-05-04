package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"userservice-go/types"
)

func main() {
	loadEnvVars()
	InitializeAndStartServer()
}

func loadEnvVars() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, will fallback to default environment variables")
	}

	types.USER_SERVICE_PORT = os.Getenv("USER_SERVICE_PORT")
	types.KEYCLOAK_BACKEND_URL = os.Getenv("KEYCLOAK_BACKEND_URL")
	types.ADMIN_USER = os.Getenv("ADMIN_USER")
	types.ADMIN_PASSWORD = os.Getenv("ADMIN_PASSWORD")
	types.GRANT_TYPE = os.Getenv("GRANT_TYPE")
	types.CLIENT_ID = os.Getenv("CLIENT_ID")
	types.KEYCLOAK_MASTER_REALM_TOKEN = os.Getenv("KEYCLOAK_MASTER_REALM_TOKEN")
	types.KEYCLOAK_MASTER_TOKEN_PATH = os.Getenv("KEYCLOAK_MASTER_TOKEN_PATH")
	types.KEYCLOAK_GET_BY_USERNAME_PATH = os.Getenv("KEYCLOAK_GET_BY_USERNAME_PATH")
	log.Println("Loaded environment variables")
}
