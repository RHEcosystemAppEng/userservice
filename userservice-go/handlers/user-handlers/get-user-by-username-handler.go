package user_handles

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	token_handlers "userservice-go/handlers/token-handlers"
	"userservice-go/types"
)

func GetUserByUserNameHandler(userName string) (error, types.User) {
	var user types.User
	tokenRequestBody := types.TokenRequestFormBody{
		Grant_type: types.GRANT_TYPE,
		Client_id:  types.CLIENT_ID,
		Username:   types.ADMIN_USER,
		Password:   types.ADMIN_PASSWORD,
	}
	err, token := token_handlers.GetTokenWithPasswordGrantHandler(tokenRequestBody)

	if err != nil {
		log.Fatal(err)
		return err, user
	}

	url := types.KEYCLOAK_BACKEND_URL + types.KEYCLOAK_GET_BY_USERNAME_PATH + userName
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
		return err, user
	}
	var bearer = "Bearer " + token.AccessToken
	req.Header.Set("Authorization", bearer)
	client := &http.Client{}
	response, err := client.Do(req)

	if response.StatusCode == http.StatusOK {
		responseData, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatal(err)
			return err, user
		}
		var users []types.User
		err = json.Unmarshal(responseData, &users)
		if err != nil {
			log.Fatal(err)
			return err, user
		}
		user = users[0]
	}
	return nil, user
}
