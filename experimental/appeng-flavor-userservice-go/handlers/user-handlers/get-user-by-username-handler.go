package user_handles

import (
	"encoding/json"
	"errors"
	tokenhandlers "find-users-service-go/handlers/token-handlers"
	log "github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"userservice-go/types"
)

func GetUserByUserNameHandler(userName string) (error, types.User) {
	var user types.User

	url := types.KEYCLOAK_BACKEND_URL + types.KEYCLOAK_GET_BY_USERNAME_PATH + userName

	err, req, client := tokenhandlers.GetHttpClientAndRequestWithToken(http.MethodGet, url, nil)
	if err != nil {
		log.Error().Msg(err.Error())
		return err, user
	}

	if client != nil && req != nil {
		response, err := client.Do(req)
		if err != nil {
			log.Error().Msg(err.Error())
			return err, user
		}

		if response.StatusCode == http.StatusOK {
			responseData, err := ioutil.ReadAll(response.Body)

			if err != nil {
				log.Error().Msg(err.Error())
				return err, user
			}
			var users []types.User
			err = json.Unmarshal(responseData, &users)
			if err != nil {
				log.Error().Msg(err.Error())
				return err, user
			}
			user = users[0]
		}
		return nil, user
	} else {
		return errors.New(types.ERR_NIL_HTTP_CLIENT_OR_REQUEST), user
	}
}
