package user_handles

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"strconv"
	tokenhandlers "userservice-go/handlers/token-handlers"
	"userservice-go/types"
)

func ActivateUser(id string, activate bool) error {
	url := types.KEYCLOAK_BACKEND_URL + types.KEYCLOAK_USERS_RESOURCE_URI + "/" + id
	log.Info().Msg("Activating user on: " + url + " with value: " + strconv.FormatBool(activate))

	user := types.ActivateUser{Enabled: strconv.FormatBool(activate)}
	userJson, _ := json.Marshal(user)

	err, req, client := tokenhandlers.GetHttpClientAndRequestWithToken(http.MethodPut, url, bytes.NewReader(userJson))
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	if client != nil && req != nil {
		log.Debug().Msg(fmt.Sprintf("Activating user body %s", userJson))
		req.Header.Set("content-type", "application/json")
		response, err := client.Do(req)
		if err != nil {
			log.Error().Msg(err.Error())
			return err
		}

		if response.StatusCode == http.StatusOK || response.StatusCode == http.StatusNoContent {
			body, err := ioutil.ReadAll(response.Body)
			log.Debug().Msg(fmt.Sprintf("Activating user response %s with code %s", body, response.StatusCode))
			if err != nil {
				log.Error().Msg(err.Error())
				return err
			}
		}
	}

	return nil
}
