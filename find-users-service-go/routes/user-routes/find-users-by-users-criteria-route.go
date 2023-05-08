package user_routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"userservice-go/types"
)

func GetUsersByUsersCriteria(c *gin.Context) {
	// org_id parameter is required
	orgId := c.Query(types.ORG_ID_PARAM)
	if strings.TrimSpace(orgId) == "" {
		error := &types.Error{Detail: "org_id parameter is required", Status: types.HTTP_CODE_BAD_REQUEST}
		c.JSON(http.StatusBadRequest, error)
		return
	}

	// One of the filters needs to be provided: usernames, emails or user_ids
	userids := c.Query(types.USER_IDS_PARAM)
	emailids := c.Query(types.EMAILS_PARAM)
	usernames := c.Query(types.USER_NAMES_PARAM)

	if len(userids) == 0 && len(emailids) == 0 && len(usernames) == 0 {
		error := &types.Error{Detail: "One of the filter criteria needs to be specified: user_ids, emails or usernames", Status: types.HTTP_CODE_BAD_REQUEST}
		c.JSON(http.StatusBadRequest, error)
		return
	}

	// Process the parameters and get data

}
