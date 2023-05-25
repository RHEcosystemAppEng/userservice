package token_handlers

import (
	"crypto/tls"
	"encoding/json"
	log "github.com/rs/zerolog/log"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"userservice-go/types"
)

func getHttpClient() http.Client {
	transport := &http.Transport{}

	if len(types.DISABLE_KEYCLOAK_CERT_VERIFICATION) > 0 {
		disableTlsCertVerification, _ := strconv.ParseBool(types.DISABLE_KEYCLOAK_CERT_VERIFICATION)
		if disableTlsCertVerification {
			transport = &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			log.Debug().Msg("http client skipped certificate verification")
		}
	} else {
		log.Debug().Msg("http client enabled certificate verification")
	}

	return http.Client{Transport: transport}
}

func GetTokenWithPasswordGrantHandler(tokenRequestFormBody types.TokenRequestFormBody) (error, types.Token) {
	var token types.Token

	data := url.Values{}
	data.Set("username", tokenRequestFormBody.Username)
	data.Set("password", tokenRequestFormBody.Password)
	data.Set("client_id", tokenRequestFormBody.Client_id)
	data.Set("grant_type", tokenRequestFormBody.Grant_type)

	url := types.KEYCLOAK_BACKEND_URL + types.KEYCLOAK_TOKEN_PATH

	client := getHttpClient()
	if &client != nil {
		response, err := client.PostForm(url, data)
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
			err = json.Unmarshal(responseData, &token)
			if err != nil {
				log.Error().Msg(err.Error())
				return err, token
			}
		} else {
			responseData, _ := ioutil.ReadAll(response.Body)
			log.Debug().Msg("Unsuccessful getting token response from " + url + " " + string(responseData))
		}
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
	client := getHttpClient()
	return nil, req, &client
}
