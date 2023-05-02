package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"github.com/gin-gonic/gin"
)

var (
	BACKEND_KEYCLOAK_URL = "http://localhost:8080"
	USER_SERVICE_PORT = ":8000"
	allowedURLs = map[string]bool{
		"/realms/master/protocol/openid-connect/token":  true,
	}
)

func isURLAllowedMiddleware(whitelist map[string]bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		fmt.Println("Path is ", path)
		if !whitelist[path] {
			c.IndentedJSON(http.StatusForbidden, gin.H{
				"message": "This endpoint cannot be used",
			})
			return
		}
	}
}

func proxy(c *gin.Context) {
	remote, err := url.Parse(BACKEND_KEYCLOAK_URL) // Keycloak backend service

	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)

	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("proxyPath")
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}

func main() {

	r := gin.Default()

	r.Use(isURLAllowedMiddleware(allowedURLs))
	//Create a catchall route

	r.Any("/*proxyPath", proxy)

	r.Run(USER_SERVICE_PORT)
}