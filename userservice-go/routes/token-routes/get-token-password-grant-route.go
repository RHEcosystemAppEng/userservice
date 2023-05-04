package token_routes

import (
	"github.com/gin-gonic/gin"
	log "github.com/rs/zerolog/log"
	"net/http"
	tokenhandlers "userservice-go/handlers/token-handlers"
	"userservice-go/types"
)

func GetTokenWithPasswordGrant(c *gin.Context) {
	var tokenRequest types.TokenRequestFormBody
	err := c.Bind(&tokenRequest)
	if err != nil {
		log.Fatal().Msg(err.Error())
		error := &types.Error{Error: "Invalid token request", ErrorDescription: "Invalid token request data"}
		c.JSON(http.StatusBadRequest, error)
		return
	}

	err, token := tokenhandlers.GetTokenWithPasswordGrantHandler(tokenRequest)

	if err != nil {
		log.Fatal().Msg(err.Error())
		error := &types.Error{Error: "Error getting token", ErrorDescription: "Error getting token"}
		c.JSON(http.StatusBadRequest, error)
		return
	}
	c.JSON(http.StatusOK, token)
}
