package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"github.com/gin-gonic/gin"
)

var (
	BACKEND_KEYCLOAK_URL = "http://localhost:8080"
	USER_SERVICE_PORT = ":8000"
)

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

	//Create a catchall route
	r.Any("/*proxyPath", proxy)

	r.Run(USER_SERVICE_PORT)
}