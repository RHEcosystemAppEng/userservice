package main

import (
	"github.com/gin-gonic/gin"
	routes "userservice-go/routes/user-routes"
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
	server.GET("/user/:username", routes.GetUserByUserName)
}

func startServer(server gin.Engine) {
	server.Run(USER_SERVICE_PORT)
}
