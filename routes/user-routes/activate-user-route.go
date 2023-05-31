package user_routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	userhandlers "userservice-go/handlers/user-handlers"
	"userservice-go/types"
)

func ActivateUser(c *gin.Context) {
	id := strings.TrimSpace(c.Param(types.ID_PARAM))
	activate := strings.TrimSpace(c.Param(types.ACTIVATE_PARAM))

	if id == "" {
		error := &types.Error{Detail: "Missing id parameter to activate user", Status: types.HTTP_CODE_BAD_REQUEST}
		c.JSON(http.StatusBadRequest, error)
		return
	}
	if activate == "" {
		error := &types.Error{Detail: "Missing activate parameter to activate user", Status: types.HTTP_CODE_BAD_REQUEST}
		c.JSON(http.StatusBadRequest, error)
		return
	}

	activateBool := false
	if activate == "true" || activate == "1" {
		activateBool = true
	}

	if err := userhandlers.ActivateUser(id, activateBool); err != nil {
		error := &types.Error{Detail: "Error activating user with id: " + id, Status: types.HTTP_CODE_BAD_REQUEST}
		c.JSON(http.StatusBadRequest, error)
		return
	} else {
		c.JSON(http.StatusNoContent, nil)
	}
}
