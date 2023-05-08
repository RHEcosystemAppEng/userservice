package test

import (
	tokenroutes "find-users-service-go/routes/token-routes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestGetToken(t *testing.T) {
	r := SetUpRouter()
	r.POST("/token", tokenroutes.GetTokenWithPasswordGrant)

	data := url.Values{}
	data.Set("username", "admin")
	data.Set("password", "admin")
	data.Set("client_id", "admin-cli")
	data.Set("grant_type", "password")

	req, _ := http.NewRequest(http.MethodPost, "/token", strings.NewReader(data.Encode()))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetTokenNoUserName(t *testing.T) {
	r := SetUpRouter()
	r.POST("/token", tokenroutes.GetTokenWithPasswordGrant)

	data := url.Values{}
	data.Set("username", "admin")
	data.Set("password", "admin")
	data.Set("client_id", "admin-cli")
	data.Set("grant_type", "password")

	req, _ := http.NewRequest(http.MethodPost, "/token", strings.NewReader(data.Encode()))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
