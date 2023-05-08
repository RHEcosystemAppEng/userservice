package main

import (
	tokenroutes "find-users-service-go/routes/token-routes"
	userroutes "find-users-service-go/routes/user-routes"
	"github.com/gin-gonic/gin"
	"userservice-go/types"
)

// InitializeAndStartServer Initializes and starts the server
func InitializeAndStartServer() {
	server := gin.Default()
	initializeRoutes(*server)
	startServer(*server)
}

func initializeRoutes(server gin.Engine) {
	server.GET("/user/:username", userroutes.GetUserByUserName)
	server.POST("/token", tokenroutes.GetTokenWithPasswordGrant)
}

func startServer(server gin.Engine) {
	server.Run(types.USER_SERVICE_PORT)
}
