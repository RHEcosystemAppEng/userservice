package user_routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	userhandlers "userservice-go/handlers/user-handlers"
	"userservice-go/types"
)

func ActivateUser(c *gin.Context) {
	id := c.Param(types.ID_PARAM)
	if strings.TrimSpace(id) == "" {
		error := &types.Error{Detail: "Missing id parameter to activate user", Status: types.HTTP_CODE_BAD_REQUEST}
		c.JSON(http.StatusBadRequest, error)
		return
	}
	if err := userhandlers.ActivateUser(id, true); err != nil {
		error := &types.Error{Detail: "Error activating user with id: " + id, Status: types.HTTP_CODE_BAD_REQUEST}
		c.JSON(http.StatusBadRequest, error)
		return
	} else {
		c.JSON(http.StatusNoContent, nil)
	}
}

func DeactivateUser(c *gin.Context) {
	id := c.Param(types.ID_PARAM)
	if strings.TrimSpace(id) == "" {
		error := &types.Error{Detail: "Missing id parameter to activate user", Status: types.HTTP_CODE_BAD_REQUEST}
		c.JSON(http.StatusBadRequest, error)
		return
	}
	if err := userhandlers.ActivateUser(id, false); err != nil {
		error := &types.Error{Detail: "Error activating user with id: " + id, Status: types.HTTP_CODE_BAD_REQUEST}
		c.JSON(http.StatusBadRequest, error)
		return
	} else {
		c.JSON(http.StatusNoContent, nil)
	}
}
