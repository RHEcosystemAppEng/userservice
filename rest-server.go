package main

import (
	"github.com/gin-gonic/gin"
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

}

func startServer(server gin.Engine) {
	server.Run(USER_SERVICE_PORT)
}