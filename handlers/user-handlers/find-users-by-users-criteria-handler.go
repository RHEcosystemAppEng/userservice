package user_handles

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	tokenhandlers "userservice-go/handlers/token-handlers"
	"userservice-go/types"
)

func FindUsers(findUsersCriteria types.FindUsersCriteria) (error, types.UserPagination) {
	var usersList []types.UserOut
	var userPagination types.UserPagination
	var err error
	if len(findUsersCriteria.OrgId) == 0 && len(findUsersCriteria.Emails) == 0 && len(findUsersCriteria.UserIds) == 0 && len(findUsersCriteria.Usernames) == 0 {
		err, usersList = findAllUsers()
	} else if len(findUsersCriteria.OrgId) != 0 && len(findUsersCriteria.Emails) == 0 && len(findUsersCriteria.UserIds) == 0 && len(findUsersCriteria.Usernames) == 0 {
		err, usersList = findUsersByOrgId(findUsersCriteria)
	} else if len(findUsersCriteria.Emails) > 0 {
		err, usersList = findUsersByEmails(findUsersCriteria)
	} else if len(findUsersCriteria.Usernames) > 0 {
		err, usersList = findUsersByUserNames(findUsersCriteria)
	} else if len(findUsersCriteria.UserIds) > 0 {
		err, usersList = findUsersByUserIds(findUsersCriteria)
	}

	if err != nil {
		log.Error().Msg(err.Error())
		return err, userPagination
	}

	paginationMeta := getPaginationObject(findUsersCriteria, usersList)
	userPagination = getPagedResults(findUsersCriteria, usersList, paginationMeta)

	return nil, userPagination
}

func findAllUsers() (error, []types.UserOut) {
	var usersList []types.UserOut

	url := types.KEYCLOAK_BACKEND_URL + types.KEYCLOAK_GET_BY_USERS
	log.Info().Msg(url)

	err, users := executeGetUserHttpRequest(url)
	if err != nil {
		log.Error().Msg(err.Error())
		return err, usersList
	}
	usersList = append(usersList, users...)

	return nil, usersList
}

func findUsersByOrgId(findUsersCriteria types.FindUsersCriteria) (error, []types.UserOut) {
	var usersList []types.UserOut

	qPart := "q=org_id:" + findUsersCriteria.OrgId
	url := types.KEYCLOAK_BACKEND_URL + types.KEYCLOAK_GET_BY_USERS + "?" + qPart

	err, users := executeGetUserHttpRequest(url)
	if err != nil {
		log.Error().Msg(err.Error())
		return err, usersList
	}
	usersList = append(usersList, users...)
	return nil, usersList
}

func findUsersByEmails(findUsersCriteria types.FindUsersCriteria) (error, []types.UserOut) {
	var usersList []types.UserOut
	hostPath := types.KEYCLOAK_BACKEND_URL + types.KEYCLOAK_GET_BY_USERS
	url, _ := url.Parse(hostPath)
	queryParams := url.Query()
	if len(findUsersCriteria.OrgId) > 0 {
		queryParams.Set("q", "org_id:"+findUsersCriteria.OrgId)
	}

	for _, email := range findUsersCriteria.Emails {
		if len(email) > 0 {
			queryParams.Set("email", email)
			url.RawQuery = queryParams.Encode()
			log.Info().Msg(url.String())
			err, users := executeGetUserHttpRequest(url.String())
			if err != nil {
				log.Error().Msg(err.Error())
				return err, usersList
			}
			usersList = append(usersList, users...)
		}
	}
	return nil, usersList
}

func findUsersByUserNames(findUsersCriteria types.FindUsersCriteria) (error, []types.UserOut) {
	var usersList []types.UserOut
	hostPath := types.KEYCLOAK_BACKEND_URL + types.KEYCLOAK_GET_BY_USERS
	url, _ := url.Parse(hostPath)
	queryParams := url.Query()
	if len(findUsersCriteria.OrgId) > 0 {
		queryParams.Set("q", "org_id:"+findUsersCriteria.OrgId)
	}

	for _, userName := range findUsersCriteria.Usernames {
		if len(userName) != 0 {
			queryParams.Set("username", userName)
			url.RawQuery = queryParams.Encode()

			log.Info().Msg(url.String())
			err, users := executeGetUserHttpRequest(url.String())
			if err != nil {
				log.Error().Msg(err.Error())
				return err, usersList
			}
			usersList = append(usersList, users...)
		}
	}
	return nil, usersList
}

func findUsersByUserIds(findUsersCriteria types.FindUsersCriteria) (error, []types.UserOut) {
	var usersList []types.UserOut
	hostPath := types.KEYCLOAK_BACKEND_URL + types.KEYCLOAK_GET_BY_USERS
	url, _ := url.Parse(hostPath)
	queryParams := url.Query()
	if len(findUsersCriteria.OrgId) > 0 {
		queryParams.Set("q", "org_id:"+findUsersCriteria.OrgId)
	}

	for _, userId := range findUsersCriteria.UserIds {
		if len(userId) != 0 {
			queryParams.Set("id", userId)
			url.RawQuery = queryParams.Encode()
			err, users := executeGetUserHttpRequest(url.String())
			if err != nil {
				log.Error().Msg(err.Error())
				return err, usersList
			}
			usersList = append(usersList, users...)
		}
	}
	return nil, usersList
}

func executeGetUserHttpRequest(url string) (error, []types.UserOut) {
	var users []types.UserOut

	err, req, client := tokenhandlers.GetHttpClientAndRequestWithToken(http.MethodGet, url, nil)
	if err != nil {
		log.Error().Msg(err.Error())
		return err, users
	}

	if client != nil && req != nil {
		response, err := client.Do(req)
		if err != nil {
			log.Error().Msg(err.Error())
			return err, users
		}

		if response.StatusCode == http.StatusOK {
			responseData, err := ioutil.ReadAll(response.Body)

			if err != nil {
				log.Error().Msg(err.Error())
				return err, users
			}
			err = json.Unmarshal(responseData, &users)
			if err != nil {
				log.Error().Msg(err.Error())
				return err, users
			}
			users = processUsersCustomAttributes(users)
		}
	}
	return nil, users
}

func processUsersCustomAttributes(users []types.UserOut) []types.UserOut {
	for i, user := range users {
		users[i] = processUserCustomAttributes(user)
	}

	return users
}

func processUserCustomAttributes(user types.UserOut) types.UserOut {
	if len(user.Attributes["is_internal"]) > 0 {
		isInternal := user.Attributes["is_internal"]
		user.IsInternal, _ = strconv.ParseBool(isInternal[0])
	}

	if len(user.Attributes["org_admin"]) > 0 {
		orgAdmin := user.Attributes["org_admin"]
		user.OrgAdmin, _ = strconv.ParseBool(orgAdmin[0])
	}

	if len(user.Attributes["type"]) > 0 {
		userType := user.Attributes["type"]
		user.Type_ = userType[0]
	}

	return user
}

func getPaginationObject(findUsersCriteria types.FindUsersCriteria, usersList []types.UserOut) types.PaginationMeta {
	totalUsers := len(usersList)
	pageSize := findUsersCriteria.QueryLimit
	currentIdx := findUsersCriteria.Offset

	first := ""
	previous := ""
	next := ""
	last := ""

	if currentIdx > 0 {
		first = fmt.Sprintf("%s%d", "/users?offset=0&limit=", findUsersCriteria.QueryLimit)
	}

	previousIdx := currentIdx - pageSize
	if previousIdx >= 0 {
		previous = fmt.Sprintf("%s%d%s%d", "/users?offset=", previousIdx, "&limit=", findUsersCriteria.QueryLimit)
	}

	nextIdx := currentIdx + pageSize
	if nextIdx < totalUsers && nextIdx >= pageSize {
		next = fmt.Sprintf("%s%d%s%d", "/users?offset=", nextIdx, "&limit=", findUsersCriteria.QueryLimit)
	}

	lastIdx := totalUsers - (totalUsers % pageSize) - 1
	if lastIdx < totalUsers && currentIdx != lastIdx {
		last = fmt.Sprintf("%s%d%s%d", "/users?offset=", lastIdx, "&limit=", findUsersCriteria.QueryLimit)
	}

	paginationMeta := types.PaginationMeta{
		Total:    int64(len(usersList)),
		First:    first,
		Previous: previous,
		Next:     next,
		Last:     last,
	}

	log.Debug().Msg((fmt.Sprintf("FindUsers Pagination: %+v\n", paginationMeta)))

	return paginationMeta
}

func getPagedResults(findUsersCriteria types.FindUsersCriteria, usersList []types.UserOut, paginationMeta types.PaginationMeta) types.UserPagination {
	totalUsers := len(usersList)
	pageSize := findUsersCriteria.QueryLimit

	beginIdx := findUsersCriteria.Offset
	if beginIdx > totalUsers {
		beginIdx = totalUsers - pageSize
		if beginIdx < 0 {
			beginIdx = 0
		}
	}

	endIdx := beginIdx + pageSize
	if endIdx > totalUsers {
		endIdx = totalUsers
	}

	userPagination := types.UserPagination{
		Meta:  &paginationMeta,
		Users: usersList[beginIdx:endIdx],
	}

	return userPagination
}
