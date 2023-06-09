package user_routes

import (
	user_handles "find-users-service-go/handlers/user-handlers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"userservice-go/types"
)

func GetUserByUserName(c *gin.Context) {
	username := c.Param(types.USERNAME_PARAM)
	if strings.TrimSpace(username) == "" {
		error := &types.Error{Error: "username cannot be empty", ErrorDescription: "username cannot be empty"}
		c.JSON(http.StatusBadRequest, error)
		return
	}

	err, user := user_handles.GetUserByUserNameHandler(username)

	if err != nil {
		error := &types.Error{Error: "Get by username error", ErrorDescription: "Get by username error"}
		c.JSON(http.StatusBadRequest, error)
		return
	}

	c.JSON(http.StatusOK, user)
}
