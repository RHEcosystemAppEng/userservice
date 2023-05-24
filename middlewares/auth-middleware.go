package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"strings"
)

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	fmt.Println("authHeader: " + authHeader)

	if len(authHeader) == 0 || strings.Index(authHeader, "Bearer") != 0 {
		returnErrorResponse(c, http.StatusUnauthorized, "Unauthorized - Authorization header is not available with Bearer token. Example- Authorization: Bearer <TOKEN>")
		return
	}
	reqToken := strings.Split(authHeader, " ")[1]

	SecretKey := "-----BEGIN CERTIFICATE-----\n" + os.Getenv("KEYCLOAK_ACCESS_TOKEN_PUBLIC_KEY") + "\n-----END CERTIFICATE-----"
	key, er := jwt.ParseRSAPublicKeyFromPEM([]byte(SecretKey))
	if er != nil {
		returnErrorResponse(c, http.StatusUnauthorized, " Unauthorized - Token is not valid.")
		return
	}

	token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		returnErrorResponse(c, http.StatusUnauthorized, "Unauthorized - Token is expired or not valid. ")
		return
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Debug().Msg("token is valid")
	}
}

func returnErrorResponse(c *gin.Context, statusCode int, message string) {
	fmt.Println(message)
	c.Writer.WriteHeader(statusCode)
	c.Writer.Write([]byte(message))
	c.Abort()
}
