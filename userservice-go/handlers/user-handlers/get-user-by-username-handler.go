package user_handles

import (
	"userservice-go/types"
)

func GetUserByUserNameHandler(userName string) types.User {
	url := types.KEYCLOAK_BACKEND_URL
	user := types.User{Id: url}
	return user
}
