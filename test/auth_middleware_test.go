package test

import (
	httpmock "github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"userservice-go/middlewares"
	userroutes "userservice-go/routes/user-routes"
)

func TestAuthMiddlewareNoBearerToken(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	setupHttpMockForFindUsersNoParams()

	r := SetUpRouter()
	r.GET("/users", middlewares.AuthMiddleware, userroutes.GetUsersByUsersCriteria)

	req, _ := http.NewRequest(http.MethodGet, "/users", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, "Unauthorized - Authorization header is not available with Bearer token. Example- Authorization: Bearer <TOKEN>", w.Body.String())
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthMiddlewareBlankBearerToken(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	setupHttpMockForFindUsersNoParams()

	r := SetUpRouter()
	r.GET("/users", middlewares.AuthMiddleware, userroutes.GetUsersByUsersCriteria)

	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set("Authorization", "Bearer    ")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, " Unauthorized - Token is not valid.", w.Body.String())
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
