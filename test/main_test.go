package test

import (
	"github.com/gin-gonic/gin"
	"os"
	env "userservice-go/env"
)

func SetUpRouter() *gin.Engine {
	env.LoadEnvVars("../")
	os.Setenv(UNIT_TEST_RUN, "true")
	router := gin.Default()
	return router
}
