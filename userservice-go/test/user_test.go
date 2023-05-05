package test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	userroutes "userservice-go/routes/user-routes"
)

func TestGetUserByUserNameBlank(t *testing.T) {
	r := SetUpRouter()
	r.GET("/user/:username", userroutes.GetUserByUserName)
	username := " "
	req, _ := http.NewRequest("GET", "/user/"+username, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetUserByUserName(t *testing.T) {
	r := SetUpRouter()
	r.GET("/user/:username", userroutes.GetUserByUserName)
	username := "admin"
	req, _ := http.NewRequest("GET", "/user/"+username, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
