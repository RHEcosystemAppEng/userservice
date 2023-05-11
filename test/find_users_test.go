package test

import (
	httpmock "github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	userroutes "userservice-go/routes/user-routes"
)

func TestFindUsersNoParams(t *testing.T) {
	r := SetUpRouter()
	r.GET("/users", userroutes.GetUsersByUsersCriteria)

	req, _ := http.NewRequest(http.MethodGet, "/users", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestFindUsersNoFilterCriteria(t *testing.T) {
	r := SetUpRouter()
	r.GET("/users", userroutes.GetUsersByUsersCriteria)

	req, _ := http.NewRequest(http.MethodGet, "/users?org_id=rh", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestFindUsersByEmails(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	setupHttpMockForFindUsersByEmail()

	r := SetUpRouter()
	r.GET("/users", userroutes.GetUsersByUsersCriteria)

	req, _ := http.NewRequest(http.MethodGet, "/users?org_id=rh&emails=1@1.com,2@2.com", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestFindUsersByUserNames(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	setupHttpMockForFindUsersByUserName()

	r := SetUpRouter()
	r.GET("/users", userroutes.GetUsersByUsersCriteria)

	req, _ := http.NewRequest(http.MethodGet, "/users?org_id=rh&usernames=mgr1-test,eng1-test", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestFindUsersByUserIds(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	setupHttpMockForFindUsersByUserId()

	r := SetUpRouter()
	r.GET("/users", userroutes.GetUsersByUsersCriteria)

	req, _ := http.NewRequest(http.MethodGet, "/users?org_id=rh&user_ids=c2979a54-b50e-473a-8ff8-0710f701e64f,3c577a73-d15a-4130-968b-1fdab10e0ee0", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func setupHttpMockForGetToken() {
	httpmock.RegisterResponder("POST", KEYCLOAK_GET_TOKEN_URL,
		httpmock.NewStringResponder(200, KEYCLOAK_GET_TOKEN_RESPONSE))
}

func setupHttpMockForFindUsersByEmail() {
	setupHttpMockForGetToken()
	httpmock.RegisterResponder("GET", KEYCLOAK_FIND_USERS_BY_EMAIL1,
		httpmock.NewStringResponder(200, KEYCLOAK_USER_DATA1))

	httpmock.RegisterResponder("GET", KEYCLOAK_FIND_USERS_BY_EMAIL2,
		httpmock.NewStringResponder(200, KEYCLOAK_USER_DATA2))
}

func setupHttpMockForFindUsersByUserName() {
	setupHttpMockForGetToken()
	httpmock.RegisterResponder("GET", KEYCLOAK_FIND_USERS_BY_USERNAME1,
		httpmock.NewStringResponder(200, KEYCLOAK_USER_DATA1))

	httpmock.RegisterResponder("GET", KEYCLOAK_FIND_USERS_BY_USERNAME2,
		httpmock.NewStringResponder(200, KEYCLOAK_USER_DATA2))
}

func setupHttpMockForFindUsersByUserId() {
	setupHttpMockForGetToken()
	httpmock.RegisterResponder("GET", KEYCLOAK_FIND_USERS_BY_USERID1,
		httpmock.NewStringResponder(200, KEYCLOAK_USER_DATA1))

	httpmock.RegisterResponder("GET", KEYCLOAK_FIND_USERS_BY_USERID2,
		httpmock.NewStringResponder(200, KEYCLOAK_USER_DATA2))
}
