package user_routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
	userhandlers "userservice-go/handlers/user-handlers"
	"userservice-go/types"
)

func GetUsersByUsersCriteria(c *gin.Context) {
	var findUsersCriteria types.FindUsersCriteria
	err := c.Bind(&findUsersCriteria)
	if err != nil {
		error := &types.Error{Detail: err.Error(), Status: types.HTTP_CODE_BAD_REQUEST}
		c.JSON(http.StatusBadRequest, error)
		return
	}

	if len(findUsersCriteria.UserIdsQueryArray) != 0 {
		findUsersCriteria.UserIds = strings.Split(findUsersCriteria.UserIdsQueryArray, ",")
	}
	if len(findUsersCriteria.EmailsQueryArray) != 0 {
		findUsersCriteria.Emails = strings.Split(findUsersCriteria.EmailsQueryArray, ",")
	}
	if len(findUsersCriteria.UserNamesQueryArray) != 0 {
		findUsersCriteria.Usernames = strings.Split(findUsersCriteria.UserNamesQueryArray, ",")
	}

	err, usersList := userhandlers.FindUsers(findUsersCriteria)

	if err != nil {
		log.Error().Msg(err.Error())
		error := &types.Error{Detail: "Error finding users", Status: string(http.StatusBadRequest)}
		c.JSON(http.StatusBadRequest, error)
		return
	}

	c.JSON(http.StatusOK, usersList)
}
