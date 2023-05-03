package token_handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

	log.Println("Token data received: ", data)

	response, err := http.Post(types.KEYCLOAK_BACKEND_URL+types.KEYCLOAK_MASTER_TOKEN_PATH, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))

	if err != nil {
		log.Fatal(err)
		return err, token
	}

	if response.StatusCode == http.StatusOK {
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(responseData, &token)
	}
	return nil, token
}
