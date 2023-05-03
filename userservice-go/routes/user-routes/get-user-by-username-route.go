package user_routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	user_handles "userservice-go/handlers/user-handlers"
	"userservice-go/types"
)

func GetUserByUserName(c *gin.Context) {
	username := c.Param(types.USERNAME_PARAM)
	if strings.TrimSpace(username) == "" {
		error := &types.Error{Error: "username cannot be empty", ErrorDescription: "username cannot be empty"}
		c.JSON(http.StatusBadRequest, error)
	}

	user_handles.GetUserByUserNameHandler(username)
}
