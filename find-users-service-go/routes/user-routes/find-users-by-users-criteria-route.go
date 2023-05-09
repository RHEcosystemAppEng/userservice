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
	err := c.ShouldBindQuery(&findUsersCriteria)
	if err != nil {
		error := &types.Error{Detail: err.Error(), Status: types.HTTP_CODE_BAD_REQUEST}
		c.JSON(http.StatusBadRequest, error)
		return
	}

	// One of the filters needs to be provided: usernames, emails or user_ids
	if len(findUsersCriteria.UserIdsQueryArray) == 0 && len(findUsersCriteria.EmailsQueryArray) == 0 && len(findUsersCriteria.UserNamesQueryArray) == 0 {
		error := &types.Error{Detail: "One of the filter criteria needs to be specified: user_ids, emails or usernames", Status: types.HTTP_CODE_BAD_REQUEST}
		c.JSON(http.StatusBadRequest, error)
		return
	}

	findUsersCriteria.Emails = strings.Split(findUsersCriteria.EmailsQueryArray, ",")
	findUsersCriteria.UserIds = strings.Split(findUsersCriteria.UserIdsQueryArray, ",")
	findUsersCriteria.Usernames = strings.Split(findUsersCriteria.UserNamesQueryArray, ",")

	err, usersList := userhandlers.FindUsers(findUsersCriteria)

	if err != nil {
		log.Error().Msg(err.Error())
		error := &types.Error{Detail: "Error finding users", Status: string(http.StatusBadRequest)}
		c.JSON(http.StatusBadRequest, error)
		return
	}

	c.JSON(http.StatusOK, usersList)
}
