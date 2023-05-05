package token_handlers

import (
	"encoding/json"
	log "github.com/rs/zerolog/log"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"userservice-go/types"
)

func GetTokenWithPasswordGrantHandler(tokenRequestFormBody types.TokenRequestFormBody) (error, types.Token) {
	var token types.Token

	data := url.Values{}
	data.Set("username", tokenRequestFormBody.Username)
	data.Set("password", tokenRequestFormBody.Password)
	data.Set("client_id", tokenRequestFormBody.Client_id)
	data.Set("grant_type", tokenRequestFormBody.Grant_type)

	response, err := http.Post(types.KEYCLOAK_BACKEND_URL+types.KEYCLOAK_TOKEN_PATH, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))

	if err != nil {
		log.Error().Msg(err.Error())
		return err, token
	}

	if response.StatusCode == http.StatusOK {
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Error().Msg(err.Error())
			return err, token
		}
		json.Unmarshal(responseData, &token)
	}
	return nil, token
}

func GetKeycloakToken() (error, types.Token) {
	tokenRequestBody := types.TokenRequestFormBody{
		Grant_type: types.GRANT_TYPE,
		Client_id:  types.CLIENT_ID,
		Username:   types.ADMIN_USER,
		Password:   types.ADMIN_PASSWORD,
	}
	return GetTokenWithPasswordGrantHandler(tokenRequestBody)
}

func GetHttpClientAndRequestWithToken(httpMethod string, url string, body io.Reader) (error, *http.Request, *http.Client) {
	req, err := http.NewRequest(http.MethodGet, url, body)
	if err != nil {
		log.Error().Msg(err.Error())
		return err, nil, nil
	}

	err, token := GetKeycloakToken()
	if err != nil {
		log.Error().Msg(err.Error())
		return err, nil, nil
	}

	var bearer = "Bearer " + token.AccessToken
	req.Header.Set("Authorization", bearer)
	return nil, req, &http.Client{}
}
