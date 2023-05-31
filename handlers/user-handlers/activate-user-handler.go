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

	// Add approved custom attribute to the user if needed while activating
	if activate {
		_, hasGroupApprovedAttribute := userGroupHasApprovedAttribute(id)
		_, hasUserApprovedAttribute := userHasApprovedAttribute(id)

		if !hasGroupApprovedAttribute && !hasUserApprovedAttribute {
			userAttribs := make(map[string][]string)
			approvedAttrib := []string{"true"}
			userAttribs["approved"] = approvedAttrib
			user.Attributes = userAttribs
		}
	}

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

func userGroupHasApprovedAttribute(id string) (error, bool) {
	returnVal := false
	var groups []types.Group

	// Get User's groups and see whether group has approved attribute at group level
	url := types.KEYCLOAK_BACKEND_URL + types.KEYCLOAK_USERS_RESOURCE_URI + "/" + id + "/groups"
	log.Info().Msg("Activate user -> get groups: " + url)

	err, req, client := tokenhandlers.GetHttpClientAndRequestWithToken(http.MethodGet, url, nil)

	if err != nil {
		log.Error().Msg(err.Error())
		return err, returnVal
	}

	if client != nil && req != nil {
		response, err := client.Do(req)
		if err != nil {
			log.Error().Msg(err.Error())
			return err, returnVal
		}

		if response.StatusCode == http.StatusOK {
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Error().Msg(err.Error())
				return err, returnVal
			}
			err = json.Unmarshal(body, &groups)
			if err != nil {
				log.Error().Msg(err.Error())
				return err, returnVal
			}
		}
	}

	// Iterate over groups and see if any groups have approved attribute
	if len(groups) > 0 {
		for _, g := range groups {
			if len(g.Attributes) > 0 && g.Attributes["approved"][0] == "true" {
				returnVal = true
				log.Debug().Msg("User with Id: " + id + " has group level approved attribute")
				break
			}
		}
	}

	return nil, returnVal
}

func userHasApprovedAttribute(id string) (error, bool) {
	returnVal := false
	var users []types.UserOut

	// Get User's custom attributes and see if she has approved attribute defined
	url := types.KEYCLOAK_BACKEND_URL + types.KEYCLOAK_USERS_RESOURCE_URI + "/" + id
	log.Info().Msg("Activate user -> get user: " + url)

	err, req, client := tokenhandlers.GetHttpClientAndRequestWithToken(http.MethodGet, url, nil)

	if err != nil {
		log.Error().Msg(err.Error())
		return err, returnVal
	}

	if client != nil && req != nil {
		response, err := client.Do(req)
		if err != nil {
			log.Error().Msg(err.Error())
			return err, returnVal
		}

		if response.StatusCode == http.StatusOK {
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Error().Msg(err.Error())
				return err, returnVal
			}
			err = json.Unmarshal(body, &users)
			if err != nil {
				log.Error().Msg(err.Error())
				return err, returnVal
			}
		}
	}

	// Iterate over user and see if a custom attribute approved is present
	if len(users) > 0 {
		for _, u := range users {
			if len(u.Attributes) > 0 && u.Attributes["approved"][0] == "true" {
				returnVal = true
				log.Debug().Msg("User with Id: " + id + " has approved attribute")
				break
			}
		}
	}

	return nil, returnVal
}
