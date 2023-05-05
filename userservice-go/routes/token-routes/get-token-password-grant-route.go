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
		error := &types.Error{Error: "Invalid token request", ErrorDescription: err.Error()}
		c.JSON(http.StatusBadRequest, error)
		log.Error().Msg(err.Error())
		return
	}

	err, token := tokenhandlers.GetTokenWithPasswordGrantHandler(tokenRequest)

	if err != nil {
		log.Error().Msg(err.Error())
		error := &types.Error{Error: "Error getting token", ErrorDescription: "Error getting token"}
		c.JSON(http.StatusBadRequest, error)
		return
	}
	c.JSON(http.StatusOK, token)
}
