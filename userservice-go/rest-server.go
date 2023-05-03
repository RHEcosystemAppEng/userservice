package main

import (
	"github.com/gin-gonic/gin"
	tokenroutes "userservice-go/routes/token-routes"
	userroutes "userservice-go/routes/user-routes"
)

var (
	USER_SERVICE_PORT = ":8000"
)

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
	server.Run(USER_SERVICE_PORT)
}
