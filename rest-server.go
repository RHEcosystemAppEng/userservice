package main

import "C"
import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	tokenroutes "userservice-go/routes/token-routes"
	userroutes "userservice-go/routes/user-routes"
	"userservice-go/types"
)

// InitializeAndStartServer Initializes and starts the server
func InitializeAndStartServer() {
	server := gin.New()
	//setupCORS(*server)
	initializeRoutes(*server)
	startServer(*server)
}

//func setupCORS(server gin.Engine) {
//	// CORS Config
//	config := cors.Config{
//		AllowAllOrigins:  true,                                                                                                    // Todo: only allow appropriate domains in production
//		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS", "PATCH", "HEAD" /*, "TRACE", "CONNECT" */}, // Todo: vet this list for production use
//		AllowHeaders:     []string{"Origin", "Content-type", "Accept", "Access-Control-Allow-Origin", "Authorization"},
//		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
//		AllowCredentials: true,
//		AllowWildcard:    true,
//		MaxAge:           1 * time.Hour,
//	}
//	server.Use(cors.New(config))
//	server.Use(addCORSHeaders)
//}
//

func addCORSHeaders(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		c.Next()
	}
	log.Info().Msg("Adding CORS headers")
	c.Header("Access-Control-Allow-Origin", "https://stage.foo.redhat.com:1337,https://stage.foo.redhat.com")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "GET,PUT,POST,PATCH,DELETE,OPTIONS,HEAD")
	c.Header("Access-Control-Allow-Headers", "Origin,Content-type,Accept,Access-Control-Allow-Origin,Authorization")
	c.Header("Access-Control-Max-Age", "3600")
	c.Next()
}

func initializeRoutes(server gin.Engine) {
	server.GET("/users", addCORSHeaders, userroutes.GetUsersByUsersCriteria)
	server.POST("/token", addCORSHeaders, tokenroutes.GetTokenWithPasswordGrant)
}

func startServer(server gin.Engine) {
	server.Run(types.USER_SERVICE_PORT)
}
