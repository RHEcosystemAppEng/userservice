package token_routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	token_handlers "userservice-go/handlers/token-handlers"
	"userservice-go/types"
)

func GetTokenWithPasswordGrant(c *gin.Context) {
	var tokenRequest types.TokenRequestFormBody
	err := c.Bind(&tokenRequest)
	if err != nil {
		log.Fatal(err)
		error := &types.Error{Error: "Invalid token request", ErrorDescription: "Invalid token request data"}
		c.JSON(http.StatusBadRequest, error)
	}

	err, token := token_handlers.GetTokenWithPasswordGrantHandler(tokenRequest)

	if err != nil {
		log.Fatal(err)
		error := &types.Error{Error: "Error getting token", ErrorDescription: "Error getting token"}
		c.JSON(http.StatusBadRequest, error)
	}
	c.JSON(http.StatusOK, token)
}
