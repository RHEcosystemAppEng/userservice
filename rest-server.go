package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	tokenroutes "userservice-go/routes/token-routes"
	userroutes "userservice-go/routes/user-routes"
	"userservice-go/types"
)

// InitializeAndStartServer Initializes and starts the server
func InitializeAndStartServer() {
	server := gin.Default()
	setupCORS(*server)
	initializeRoutes(*server)
	startServer(*server)
}

func setupCORS(server gin.Engine) {
	// CORS Config
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true                                                                                                 // Todo: only allow appropriate domains in production
	config.AllowMethods = []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS", "PATCH", "HEAD" /*, "TRACE", "CONNECT" */} // Todo: vet this list for production use
	config.AllowHeaders = []string{"Origin", "Content-type", "Accept", "Access-Control-Allow-Origin", "Authorization"}
	config.ExposeHeaders = []string{"Content-Length", "Access-Control-Allow-Origin"}
	config.AllowCredentials = true
	server.Use(cors.New(config))
}

func addCORSHeaders(c *gin.Context) {
	log.Info().Msg("Adding Access-Control-Allow-Origin header")
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Next()
}

func initializeRoutes(server gin.Engine) {
	server.GET("/users", addCORSHeaders, userroutes.GetUsersByUsersCriteria)
	server.POST("/token", addCORSHeaders, tokenroutes.GetTokenWithPasswordGrant)
}

func startServer(server gin.Engine) {
	server.Run(types.USER_SERVICE_PORT)
}
