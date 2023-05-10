package test

import (
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	tokenroutes "userservice-go/routes/token-routes"
)

func TestGetToken(t *testing.T) {
	r := SetUpRouter()
	r.POST("/token", tokenroutes.GetTokenWithPasswordGrant)

	data := url.Values{}
	data.Add("username", "admin")
	data.Add("password", "admin")
	data.Add("client_id", "admin-cli")
	data.Add("grant_type", "password")
	req, _ := http.NewRequest(http.MethodPost, "/token", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	log.Info().Msg(w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetTokenNoUserName(t *testing.T) {
	r := SetUpRouter()
	r.POST("/token", tokenroutes.GetTokenWithPasswordGrant)

	data := url.Values{}
	data.Set("username", "")
	data.Set("password", "admin")
	data.Set("client_id", "admin-cli")
	data.Set("grant_type", "password")

	req, _ := http.NewRequest(http.MethodPost, "/token", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
